# ux-designer review 1, premium-trial mocks (2026-11-08): SHIP WITH FIXES

The coordinator's copy rulings (spec 0B) come on top of these. All MUSTs are applied in the build round, and the mocks are regenerated before the parity gate.

## MUST
1. The confirmed, account and confirm pages and the emails quote the interval the buyer chose. Regenerate them from one yearly fixture so the $40 /year band and "then $5 a month" never sit together.
2. Used trial: the title becomes `Subscribe to Premium`, the note becomes `Stripe asks for a card now. You are charged $40 today.`, the action stays `Subscribe · $40 a year`, and the status is shown in neutral text, not red.
3. `Due today: $0` is the first line of the chooser's and the confirmed page's Plate.
4. The cutoff is said everywhere. Account: `Trial ends 15 Nov 2026, then $5 a month. Cancel before 15 Nov 2026 starts (UTC).` Reminder: `Cancel before 15 Nov 2026 (UTC) and you pay nothing.` Both match trial-copy.ts and the Go test.
5. Creator code field:
   - `Apply` is aligned to the input.
   - A code that applied shows full-contrast text and swaps `Apply` for `Remove`.
   - The status line replaces the hint.
   - An unknown code gets a red border, an icon and aria-describedby, and keeps focus.
   - The spec gains a state table: hover, focus-visible, active, disabled, `Checking`, `Cancelling`, and where `cancelFailed` goes.
6. The receipt names the next amount (`The next charge is $5 on 15 Dec 2026.`). Whether to add `(plus any tax)` under Managed Payments is for the owner.

## SHOULD
7. Confirmed page:
   - Title `Free Premium is on`.
   - Drop the repeated trial line.
   - 32 px above the footer.
   - `Cancel trial` everywhere.
8. No founder sentence once a trial has started (the confirmed page and guild settings).
9. At 390 the primary actions are full width and 48 px tall.
10. The plan card drops the standalone `Or $40 a year.`.
11. Cancelled: the tag reads `Ends 15 Nov`.
12. Dialog:
   - Default focus is on `Keep my trial`.
   - Esc closes it.
   - The focus ring is visible.
   - At 390 it is a bottom sheet.
13. Emails:
   - Straight apostrophes in plain text, held by the contract test.
   - The receipt names the interval (`Premium, $5 a month`).
14. 1440 chooser: the price summary moves inside the Plate, and the Plate is centred.
