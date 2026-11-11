# Page frame, review round 1 (2026-11-10)

On `after/` captures of the first build (frame unified, no page changes).

## ux-designer: NO SHIP

The two complaints are fixed at the owner's width: every bar is the same width, door icons, "/" hint
and spec line show on every page at 1920 and 2000, and the edges match on all 34 pages.

1. must: `bis-1920`: the BiS body (doll and side) sat inset about 126 px from each frame edge (a
   760 + 560 px pair centred). Fixed: the side takes the rest of the frame from 1760
   (BisSpecPage.svelte).
2. must: `b-unavailable-1920`: a 560 px card pinned top left with empty ground. Fixed: the panel
   fills the frame, then Pick your class (NotFoundClassPicker), with the mock generated to match
   (ws_premium.py).
3. should: `character-2000`: about 280 px of empty panel each side of the paperdoll. Fixed: from a
   980 px panel the full Upgrades list stands beside the doll (CharacterGear.svelte).
4. should: `legal-1920`: prose measure left-aligned with empty space to its right. Not changed: the
   owner's rule keeps prose at the reading measure inside the frame.
5. should: `tiers-1920`: the right column holds only "Sign in to mark your spec". This predates the
   change. Follow-up.
6. should: at 1440 the bar drops the icons, hint and spec on every page. This is consistent across
   pages and by design (nav.css, the bar has no room at 1248 px).
7. should: `bis-1920`: "All sources" ends about 38 px short of the frame's right edge. That is the
   end of the control row's content, not padding.

## ai-slop-reviewer: FAIL

1. must: `login-2000`, `b-unavailable-2000`: a 560 px card at the far left of the frame.
   b-unavailable is fixed (see above). Login's form measure was the same before the change.
2. must: `legal-2000`: prose and contents rail fill 60 percent of the frame. Prose measure, by the
   owner's rule.
3. must: `b-2000`: the "Read only. Edit a copy" row floats under the head band, and the rail ends
   early. This predates the change (the planner's shared-build layout).
4. should: `character-2000`: dead space beside the doll. Fixed.
5. should: `home-2000`: "Latest on Forever" stretches to the rail's height. Predates the change.
6. should: `guild-2000`: the notice row is indented about 11 px. Predates the change.
7. should: `b-1440`: same as 3.

## info-density-reviewer: THIN

1. must: `character-1920-full`: doll centred in a 1200 px panel. Fixed: Upgrades beside it.
2. must: `character-1920-full`: empty ground under the rail. Needs new data (guild medians,
   percentiles). Follow-up.
3. must: `character-1920-full`: Best per encounter rows cross 900 px of nothing. Follow-up.
4. must: `b-unavailable-1920-full`: card top left, page blank. Fixed.
5. should: `b-1920-full`: rail and Point order leave space. Follow-up.
6. to 8. should: guild overview, roster and loot table columns. Follow-up.
