# Home body, round 3 (2026-11-09)

Captures: `r3-1440.png`, `r3-2000.png`, `r3-2576.png`, `r3-390.png`.

## ai-slop-reviewer: VERDICT: FAIL

Blocking:
1. Phone Guild sims: the gap line wraps one word per line, and the bar sits over the own row's name.

Non-blocking:
- "ahead of you" repeats on every row.
- Two identical "climbed 6 places to #3" fixture rows.
- The plaque's middle.
- Two "Loot not out yet" lines.
- The phone section order.

## ux-designer: VERDICT: NO SHIP

Blocking (one root cause):
1. Phone Guild sims gap line wraps one word per line.
2. The bar overlaps the "Felsnap" name.
3. The bar clips the gap text.

Non-blocking:
- The own row has no second line.
- The highlight reaches past the dividers.
- The column bottoms differ.
- On phone, upgrades before the feed is right.

## wow-player: VERDICT: SHIP WITH FIXES

Blocking:
1. Phone Guild sims (same as above).
2. "Highlander's Chain Girdle" still reads as mail on a level 34 hunter.

Non-blocking:
- Two rankings (by item level, by DPS) sit close together.
- A "Level 60" tag on the raids.
- No weapon type on upgrades.
- The bare tier tile.

## info-density-reviewer: VERDICT: THIN

Blocking:
1. Phone Guild sims (same as above).

Non-blocking:
- Drill-downs.
- A feed guild marker.
- Sim age.
- Item tooltips.
- Whose 14 raiders.

## Changes for round 4

- **Guild sims on a phone panel:** below a 520 px panel the bar gives way, and the gap line ("12.8 DPS ahead of you") stays on one line (`white-space: nowrap`). The name, its gap and the figure keep the row.
- **Highlander's Chain Girdle:** it is leather in this client build, but a 34 hunter reads the name as mail, so the leveler fixture now wears its band's own waist pick.
  - The generator does this with `WORN_PICK` in `design/mocks/home_fixtures.py`.
  - It leaves 13 upgrades and 1 slot already best in slot.
  - The 34 spec's expectations follow.
