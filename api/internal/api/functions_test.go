package api

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sample struct {
	Name string `json:"name"`
}

// TestRetrieveStream_UnmarshalsAndDelivers verifies that raw payloads are
// decoded into typed values and delivered on the returned channel.
func TestRetrieveStream_UnmarshalsAndDelivers(t *testing.T) {
	raw := make(chan []byte, 1)
	out := RetrieveStream[sample](context.Background(), "sample", raw)

	raw <- []byte(`{"name":"btc"}`)

	select {
	case v := <-out:
		require.NotNil(t, v)
		assert.Equal(t, "btc", v.Name)
	case <-time.After(time.Second):
		t.Fatal("expected a value")
	}
}

// TestRetrieveStream_SkipsInvalidPayload ensures a malformed payload is dropped
// without terminating the stream.
func TestRetrieveStream_SkipsInvalidPayload(t *testing.T) {
	raw := make(chan []byte, 2)
	out := RetrieveStream[sample](context.Background(), "sample", raw)

	raw <- []byte(`not-json`)
	raw <- []byte(`{"name":"eth"}`)

	select {
	case v := <-out:
		require.NotNil(t, v)
		assert.Equal(t, "eth", v.Name)
	case <-time.After(time.Second):
		t.Fatal("expected the valid value after skipping the invalid one")
	}
}

// TestRetrieveStream_ContextCancelClosesChannel is the core leak-fix guarantee:
// cancelling the context terminates the goroutine and closes the output, even
// though the raw stream is never closed.
func TestRetrieveStream_ContextCancelClosesChannel(t *testing.T) {
	raw := make(chan []byte) // never closed, mirroring WSAPIBase.stream
	ctx, cancel := context.WithCancel(context.Background())
	out := RetrieveStream[sample](ctx, "sample", raw)

	cancel()

	select {
	case _, ok := <-out:
		assert.False(t, ok, "channel should be closed after context cancellation")
	case <-time.After(time.Second):
		t.Fatal("goroutine did not terminate after context cancellation")
	}
}

// TestRetrieveStreamOnce_MemoizesWithinSession verifies repeated calls return
// the same channel so the raw stream is not split across competing readers.
func TestRetrieveStreamOnce_MemoizesWithinSession(t *testing.T) {
	base := NewWSAPIBase(testRequestFactory)
	base.initContext(context.Background())

	first := RetrieveStreamOnce[sample](base, "sample")
	second := RetrieveStreamOnce[sample](base, "sample")

	assert.Equal(t, first, second, "expected the memoized stream to be reused")
}

// TestRetrieveStreamOnce_FreshAfterResubscribe verifies the Subscribe ->
// Unsubscribe -> Subscribe lifecycle: resetTypedStream (called from start on a
// new session) drops the old, already-closed stream so a live one is built.
func TestRetrieveStreamOnce_FreshAfterResubscribe(t *testing.T) {
	base := NewWSAPIBase(testRequestFactory)

	// Session 1.
	ctx1, cancel1 := context.WithCancel(context.Background())
	base.initContext(ctx1)
	first := RetrieveStreamOnce[sample](base, "sample")

	// Unsubscribe: the session context is cancelled and the stream closes.
	cancel1()
	select {
	case _, ok := <-first:
		assert.False(t, ok, "session 1 stream should close on unsubscribe")
	case <-time.After(time.Second):
		t.Fatal("session 1 stream did not close")
	}

	// Session 2: start() resets the memo before spawning new goroutines.
	base.resetTypedStream()
	ctx2 := context.Background()
	base.initContext(ctx2)
	second := RetrieveStreamOnce[sample](base, "sample")

	assert.NotEqual(t, first, second, "expected a fresh stream for the new session")

	// The new stream must be live: a payload is delivered.
	base.stream <- []byte(`{"name":"xrp"}`)
	select {
	case v := <-second:
		require.NotNil(t, v)
		assert.Equal(t, "xrp", v.Name)
	case <-time.After(time.Second):
		t.Fatal("session 2 stream is not live")
	}
}
