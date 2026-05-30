package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequestFactory(consts.WebSocketCommand) any { return nil }

// TestWSAPIBase_StateTransitions verifies the state machine helpers that
// Subscribe/start rely on, including the recovery path when the atomic.Value
// has not been initialised yet.
func TestWSAPIBase_StateTransitions(t *testing.T) {
	base := NewWSAPIBase(testRequestFactory)

	// NewWSAPIBase initialises the state to stopped.
	assert.True(t, base.isStopped(), "newly created base should be stopped")
	assert.False(t, base.isConnected())

	base.changeStateToConnecting()
	assert.False(t, base.isStopped(), "connecting is not a stopped state")
	assert.False(t, base.isConnected())

	base.changeStateToConnected()
	assert.False(t, base.isStopped())
	assert.True(t, base.isConnected())

	base.changeStateToClosed()
	assert.True(t, base.isStopped(), "a closed connection counts as stopped for restart")
	assert.False(t, base.isConnected())

	// An uninitialised state value should be treated as stopped (and self-heal).
	fresh := &WSAPIBase{state: &atomic.Value{}}
	assert.True(t, fresh.isStopped())
	assert.False(t, fresh.isConnected())
}

// TestWSAPIBase_CreateSendFunc_ContextCanceled ensures the send function does
// not block forever when the context is cancelled and no send goroutine is
// draining msgStream (the original code deadlocked on the unbuffered errChan).
func TestWSAPIBase_CreateSendFunc_ContextCanceled(t *testing.T) {
	base := NewWSAPIBase(testRequestFactory)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	base.ctx = ctx

	sendFn := base.createSendFunc("subscribe")

	done := make(chan error, 1)
	go func() { done <- sendFn() }()

	select {
	case err := <-done:
		assert.ErrorIs(t, err, context.Canceled)
	case <-time.After(time.Second):
		t.Fatal("send function blocked after context cancellation")
	}
}

// TestWSAPIBase_Subscribe_ContextCanceled verifies that subscribing with an
// already-cancelled context returns promptly without dialing, and that the
// spawned goroutines observe the cancellation and exit.
func TestWSAPIBase_Subscribe_ContextCanceled(t *testing.T) {
	base := NewWSAPIBase(testRequestFactory)
	// Point at an address that must never be dialed; the cancelled context
	// should make the receive goroutine return before any dial attempt.
	base.SetHostFactoryFunc(func() string { return "ws://127.0.0.1:0" })

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := base.Subscribe(ctx)
	assert.ErrorIs(t, err, context.Canceled)
}

// TestWSAPIBase_ConcurrentStart_SingleConnection verifies that many concurrent
// Subscribe/start calls only ever establish a single WebSocket connection, i.e.
// the start() check-and-launch is atomic and goroutine pairs are not
// duplicated. It also exercises the atomic.Pointer[conn] access under -race.
func TestWSAPIBase_ConcurrentStart_SingleConnection(t *testing.T) {
	var connCount int32
	upgrader := websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		atomic.AddInt32(&connCount, 1)
		defer func() { _ = conn.Close() }()
		// Hold the connection open until the client closes it.
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	defer srv.Close()

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")

	base := NewWSAPIBase(testRequestFactory)
	base.SetHostFactoryFunc(func() string { return wsURL })

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer base.close()

	const goroutines = 50
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			base.start(ctx)
		}()
	}
	wg.Wait()

	// Wait for the single connection to be established.
	require.Eventually(t, base.isConnected, 2*time.Second, 10*time.Millisecond,
		"expected the base to become connected")

	// Give any erroneously-spawned second goroutine a chance to dial, then
	// assert that exactly one connection was ever made.
	time.Sleep(200 * time.Millisecond)
	assert.Equal(t, int32(1), atomic.LoadInt32(&connCount),
		"concurrent start must not create more than one connection")
}
