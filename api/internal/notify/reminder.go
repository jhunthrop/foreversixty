package notify

// The destinations a raid reminder can have been sent through.
const (
	ReminderViaBot     = "bot"
	ReminderViaWebhook = "webhook"
)

// ReminderRef points at a sent reminder so it can be withdrawn again: how it went out and
// the Discord message id.
type ReminderRef struct {
	Via       string
	MessageID string
}
