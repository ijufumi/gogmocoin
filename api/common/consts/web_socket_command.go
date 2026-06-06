package consts

// WebSocketCommand is the command sent to a WebSocket channel to start or stop a subscription.
type WebSocketCommand string

const (
	// WebSocketCommandSubscribe starts a subscription to a channel.
	WebSocketCommandSubscribe = WebSocketCommand("subscribe")
	// WebSocketCommandUnsubscribe stops a subscription to a channel.
	WebSocketCommandUnsubscribe = WebSocketCommand("unsubscribe")
)
