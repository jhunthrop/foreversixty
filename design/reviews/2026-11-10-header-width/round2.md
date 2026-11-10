# Page frame, review round 2 (2026-11-10)

On `after/` captures of the second build: BiS side to the frame edge, /b-unavailable filling the
frame with Pick your class, character Upgrades beside the doll.

## ux-designer: SHIP

1. Fixed (was must): `bis-1920`, `bis-2000`: the doll's left edge and the side column's right edge
   sit on the frame (144 and 1776 at 1920, 168 and 1832 at 2000).
2. Fixed (was must): `b-unavailable-1920`: the notice and Pick your class span the frame.
3. Fixed (was should): `character-1920`, `character-2000`: the full Upgrades list beside the doll.
4. should: with no slot chosen the doll's centre is a blank field; consider defaulting to the top
   upgrade's slot as BiS does with the helm. Follow-up.
5. should: at 1440 the BiS control bar wraps "All sources" to a second row; unchanged from before.
   Follow-up.
6. Accepted: the round 1 notes on legal, tiers, the 1440 header and the control row's end.

## ai-slop-reviewer: PASS (with shoulds)

Nothing the frame change introduced reads as templated or AI-made.

1. should: `b-unavailable-2000`: the notice panel is wide for one sentence and two links, and the
   gap above Pick your class is larger than the one above the notice. Follow-up.
2. should: `character-2000`: the doll's blank centre; about 120 px under the list. The list's link
   now sits at the panel's foot.
3. should: `bis-1920`: the right column ends above the doll's foot; a consequence of the doll's
   height, leave it.

## info-density-reviewer: DENSE (with additions)

The widened pages use the frame well enough for this lane; round 1's must items 1 and 4 are done.

1. must: `character-1920-full`: about 170 px of panel under the Upgrades list. Addressed: the
   column runs the doll's height and "Open in best in slot" sits at its foot.
2. must: `b-unavailable-1920-full`: about 220 px of ground between Pick your class and the footer.
   Not changed: that is the site-wide footer on a short page (Base.astro's body is at least one
   screen tall, so the footer sits at the bottom on every short page, the 404 and /login too).
   Changing it is a site-wide footer decision, logged as a follow-up.
3. to 6. Follow-ups: Best per encounter columns and the rail on character (needs guild medians and
   per-encounter percentiles), /b point order and rail, guild roster, loot and overview columns.
   They are the same at 1440 as before.
