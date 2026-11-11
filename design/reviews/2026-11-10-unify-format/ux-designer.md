# ux-designer: lane B, house formats (2026-11-10)

Verdict: SHIP (round 2, on the captures in `after/`, commit 37fa6da59 and later)

## Round 1: HOLD
1. **guild-390, the week card:** the three-unit countdown squeezed the title to about 50 px, and the date and time line broke one word per line.
2. **rankings:** "Updated 13 Dec" is in the future.
3. **report:** "Started 23:49:25" is a 24-hour clock, which breaks the locale clock rule.

## Round 2: SHIP
1. **guild-390:** on phones the countdown now has its own line under the time (`NextRaidRow.svelte`) and reads "in 4 days, 4 hours" (zero units dropped). "Onyxia's Lair" and "Wed 13 Jan · 8 pm CST" each sit on one line.
2. **rankings:** accepted as fixture data. The parity fixture's `updated_at` is 2026-12-13, the mock world clock, so this is not a formatter bug.
3. **report:** the stamp reads "11:49:25 pm". It is the house clock with seconds in the viewer's hour cycle, read in UTC like the log, and is a named case in `apps/shared/formats.json`.

## Shoulds, not blocking
- Capture a signed-in reports list so a "0:04.2" row is on screen.
- Dates either side of New Year show the year only outside the current year ("Wed 13 Jan" beside "30 Dec 2026"). This is the one rule in `dates.ts`.
