# Task 9 report

Status: DONE.

- New auth/me_billing.go: BillingView (+founder, interval, amount_cents via PriceFacts), GuildBillingView (+tier, you_hold_seat), seat/seat_offer/ended_seat views (omitted when absent), SeatFactsSource, attachGuildPlans (moved), personal billing helpers. handler.go 795 -> 710 lines.
- Service.Seats nil leaves seat keys out. Seat facts are read once per /v1/me; guild plans take HeldGuildIDs from them.
- SeatFacts gains HeldGuildIDs (seats/me.go fills it for every held seat, Seat still the first).
- Store.Guilds reads verified_count (distinct verified accounts, leader included).
- Wiring: billingWiring gains accounts; billing_wire.go sets accounts.Seats = seatStore; main.go passes accounts.
- Ended seat price comes from the catalog by the comeback's lookup key; an unknown key gives null price and a link without interval.
- Concern: priceFactsOf duplicates billing/guild_view.go's (auth cannot import billing); could move into catalog later.

Tests: me_billing_test.go (founder, list price, owner grant, omitted keys, seat, offer own facts, ended seat, tier/seat, verified count); auth, seats, cmd/api, server, notify pass with -p 1.
