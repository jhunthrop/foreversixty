package seats

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// seatCopy is premium.seat's words: what the guild did, never an upsell.
var seatCopy = struct{ Title, Held, Offered string }{
	Title:   "A Premium seat from your guild",
	Held:    "Your guild gave you a Premium seat.",
	Offered: "Your guild offered you a Premium seat. Accept it on your account page.",
}

// seatAccountPath is where a member sees and accepts a seat.
const seatAccountPath = "/account"

// publishSeat tells toUserID the seat the guild just gave or offered them: offered when they
// pay for their own Premium (the seat waits on their accept), given otherwise. The hand-over
// time holds it to once per hand-over.
func (s *Service) publishSeat(ctx context.Context, guildID int64, toUserID int64, seat Seat, offered bool) {
	if s.Events == nil {
		return
	}
	body, state := seatCopy.Held, "given"
	if offered {
		body, state = seatCopy.Offered, "offered"
	}
	stamp := strconv.FormatInt(s.now().Unix(), 10)
	events.Quietly(context.WithoutCancel(ctx), s.Events, s.logger(), events.Event{
		Kind: events.PremiumSeat, Audience: events.Users(toUserID),
		DedupeKey: fmt.Sprintf("premium.seat:%d:%d:%d:%s:%s", guildID, seat.Slot, toUserID, state, stamp),
		Title:     seatCopy.Title, Body: body, Link: strings.TrimRight(s.PublicBaseURL, "/") + seatAccountPath,
		Data: map[string]string{"guild_id": strconv.FormatInt(guildID, 10)},
	})
}
