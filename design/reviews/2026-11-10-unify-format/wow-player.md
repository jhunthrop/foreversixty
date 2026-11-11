# wow-player: lane B, house formats (2026-11-10)

Verdict: SHIP (round 2, on the captures in `after/`)

## Round 1: HOLD
1. **Changelog prose:** still read "Sept 17 at 10:00 a.m. PDT" beside "20 Sep" in the date column.
2. **12-hour clock:** not one format. "8:00 pm CST" (Planned rows, planner) sat beside "8 pm CST" (member view), and the changelog had "p.m.".
3. **Character feed:** read "10 bosses in 92 min" where the spoken style is "1 h 32 min".
4. **Should, countdown:** drop the zero units ("in 4 days, 4 hours, 0 min").
5. **Should, report "Started 23:49:25":** say which clock it is and follow the locale.

## Round 2: SHIP
1. **Changelog:** reads "17 Sep at 10 am PDT", "Beta 17 Sep, launch 4 Nov" and "4 Nov 2026 at 3 pm PST", and a content guard fails on "Sept", "a.m./p.m." and month-first dates. The 16 Oct title ("to 21 Oct") was fixed in 7cc359c70.
2. **Clock:** "8 pm" on the hour everywhere: Planned rows, "8 pm to 11 pm CST" in the planner, "Posted … 12 pm". The member's own free-text note "work till 8:30" is left as typed.
3. **Character feed:** the web fixture reads "1 h 32 min". The real sentence is written in Go (`api/internal/activity`), so it is a **blocker for lane C before production**.
4. **Countdown:** "in 4 days, 4 hours" and "in 11 days, 4 hours". The "1 day, 5 min" case is in `apps/shared/formats.json` and replayed by `formats.test.ts`.
5. **Parse colours:** still the Warcraft Logs ladder. Blue and purple use their AA text variants and stay in the same tier.

## Shoulds, not blocking
- Captures run on different fixture clocks ("Updated 13 Dec" on rankings).
- The Onyxia sample's 5:22 kill is a fixture artefact.
