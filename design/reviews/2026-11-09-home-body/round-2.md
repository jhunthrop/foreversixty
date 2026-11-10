# Home body, round 2 (2026-11-09)

Captures: `r2-1440.png`, `r2-2000.png`, `r2-2576.png`, `r2-390.png`.

## ai-slop-reviewer: VERDICT: PASS

No blockers. Both round-1 blockers are fixed.

Non-blocking:
- Two "climbed 6 places to #3" rows.
- "Climbing this week" repeats its heading.
- "ahead/behind" has no stated reference.
- Four plates in a 2x2.
- The plaque's middle block floats.
- "#2 of 3 guild Beast Mastery Hunters" is long.
- On phone, the gap between the guild facts and the raids is tight.

## ux-designer: VERDICT: SHIP WITH FIXES

Round-1 blockers closed. Blocking:
1. Guild sims: the bar is a fixed 120 px track, so a dead gap between name and bar grows with width (about 650 px at 2000).
2. Phone Guild sims drops the bar and the ahead/behind gap.

Non-blocking:
- The right edges of the upgrade rows and the footer link differ.
- The column bottoms differ.
- The own row's blank gap.
- On phone, the gain sits centred while the text starts at the top.
- The time is centred against a wrapped row.

## wow-player: VERDICT: SHIP WITH FIXES

Round-1 blockers closed: item 20090 is leather and required level 28, and the attendance line is gone. Blocking:
1. "12.8 ahead" has no unit.
2. You can't tell whom each gap is against.
3. "32 days 6 hours" should be 30 days if the clock is live.

Non-blocking:
- Upgrades first on desktop.
- Check the first-raid list.
- Drop chances on upgrades.
- "14 raiders" should say whose.
- Item tooltips.
- On phone, Guild sims loses its gaps.

## info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

No blockers. Additions:
- A guild or "you" marker on feed rows.
- Drill-down links on the ladder and the movers.
- The sim age and spec on Guild sims.
- "0 of 9 bosses down" once pulled.
- Links on the raid rows.
- An icon strip for the raids on phone.

## Changes for round 3

- **Guild sims:**
  - Each gap carries its unit and its reference ("12.8 DPS ahead of you", "3.4 DPS behind you"), as a muted line under the name at every width, so the phone keeps it too.
  - The bar is as wide as the panel allows, `clamp(56px, 36cqi, 360px)`, at every width including the phone.
- **The countdown:** the captures run on the fixtures' frozen clock (`HOME_NOW`, 2026-11-06 18:00 UTC), and that is 32 days 6 hours to 9 December. The page counts live from the browser clock (`minuteClock`).
- **Data that doesn't exist:**
  - The previous-sim baseline is the API's own earlier row for that character.
  - Drop chances on the upgrade rows are not in the BiS payload.
