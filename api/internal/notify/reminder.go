package notify

import "fmt"

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

// maxReminderTitleRunes bounds the escaped title so the line stays far under Discord's limit.
const maxReminderTitleRunes = 400

// ReminderLine is the day-before reminder's text: the bold lead, the start, and the count.
func ReminderLine(in RaidPosted) string {
	unix := in.StartsAt.Unix()
	title := escapedWithin(collapse(in.Title), maxRaidTitleRunes, maxReminderTitleRunes)
	return fmt.Sprintf(raidCopy.ReminderLead, title, unix, unix) + " " + reminderCount(in)
}

func reminderCount(in RaidPosted) string {
	signed := in.Counts.In + in.Counts.Late
	switch {
	case signed == 0:
		return raidCopy.ReminderNobody
	case in.SizeKnown:
		return fmt.Sprintf(raidCopy.ReminderCount, signed, in.Size)
	}
	return fmt.Sprintf(raidCopy.ReminderCountOpen, signed)
}

// WebhookReminderContent is the reminder through a webhook: no buttons, so a link to the raid.
func WebhookReminderContent(in RaidPosted) string {
	return ReminderLine(in) + " " + fmt.Sprintf(raidCopy.ReminderWebhookLink, in.RaidURL)
}
