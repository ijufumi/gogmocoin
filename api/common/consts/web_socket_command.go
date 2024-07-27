package consts

// WebSocketCommand ...
type WebSocketCommand string

const (
	// WebSocketCommandSubscribe ...
	WebSocketCommandSubscribe = WebSocketCommand("subscribe")
	// WebSocketCommandUnsubscribe ...
	WebSocketCommandUnsubscribe = WebSocketCommand("unsubscribe")
)
