# Lane P15, owner follow-ups (2026-11-09): the marble centre, the struck guide ledes, talent trees in the lane

Captures: paperdoll-marble/ (before and after, hovered and idle, 390 to 2000), trees/ (guide tops, Leveling and
Talents sections, before and after at 390 to 2560, overflow JSON), guides/ (every section, after-gaps.json).

## ux-designer
VERDICT: SHIP
- MUST: none.
- SHOULD: a darker shadow or vignette under "Open in the sim ›" for a lighter marble crop (it already carries
  `text-shadow: 0 1px 2px black`); the phone list form's "Open in the sim ›" sits flush left (the BiS page's own
  list form, P2's); rank labels crowd 18 px cells at 390; a long-prose Leveling lane leaves the right column empty
  below it at 2560 (it is pinned while it fits the viewport).

## ai-slop-reviewer
PASS
- MUST: none. No pseudo-script, map rings or text-like marks left in the marble.
- SHOULD: a faint vertical seam above the tooltip and a thin curved light stroke below it (read as marble at 1x);
  "51 points spent, all placed" and "Talents by level" could be shorter (GuideLevelingTree copy, P8's); the hint is
  weakest on the brightest cyan patch.
