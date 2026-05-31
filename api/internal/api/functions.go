package api

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/common/configuration"
)

// RetrieveStream unmarshals raw WebSocket payloads from rawStream into typed *T
// values delivered on the returned channel. The backing goroutine terminates
// (and closes the returned channel) when ctx is cancelled or rawStream is
// closed, so it never leaks across a Subscribe/Unsubscribe cycle.
func RetrieveStream[T any](ctx context.Context, name string, rawStream <-chan []byte) <-chan *T {
	if ctx == nil {
		ctx = context.Background()
	}
	stream := make(chan *T, 10)
	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-rawStream:
				if !ok {
					return
				}
				if configuration.IsDebug() {
					log.Printf("[%v] received:%v", name, string(v))
				}
				res := new(T)
				if err := json.Unmarshal(v, res); err != nil {
					log.Printf("[%v] unmarshal error:%v", name, err)
					continue
				}
				select {
				case stream <- res:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return stream
}

// RetrieveStreamOnce returns the typed stream for the current Subscribe session,
// creating it on first use and memoizing it so that repeated calls share a
// single backing goroutine instead of spawning competing readers of the raw
// stream. The memo is cleared on every Subscribe (see WSAPIBase.start), so a
// Subscribe -> Unsubscribe -> Subscribe cycle yields a fresh, live stream rather
// than the previous (already closed) one.
func RetrieveStreamOnce[T any](c *WSAPIBase, name string) <-chan *T {
	c.streamMu.Lock()
	defer c.streamMu.Unlock()
	if s, ok := c.typedStream.(<-chan *T); ok {
		return s
	}
	s := RetrieveStream[T](c.ctx, name, c.stream)
	c.typedStream = s
	return s
}
