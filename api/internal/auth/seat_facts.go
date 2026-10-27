package auth

import "time"

// SeatFacts are the seat facts /v1/me carries: the seat the account holds,
// the seat offered to it, and a seat that recently ended.
type SeatFacts struct {
	Seat  *HeldSeat
	Offer *SeatOffer
	Ended *EndedSeat
}

// HeldSeat is a Premium seat the account holds.
type HeldSeat struct {
	GuildName string
	GivenBy   string
	Since     time.Time
	StartsAt  *time.Time
}

// SeatOffer is a seat offered to an account that pays for its own Premium.
type SeatOffer struct {
	GuildID   int64
	GuildName string
	FromName  string
	OfferedAt time.Time
}

// EndedSeat is a seat that replaced the account's own plan and has ended;
// it shows for 30 days.
type EndedSeat struct {
	GuildName      string
	EndedAt        time.Time
	ComebackID     int64
	PriceLookupKey string
}
