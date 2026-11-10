# Seal v4: ai-slop-reviewer, 2026-11-09

**Verdict: PASS.** No blocking findings.

First impression: a hand-painted coin by an artist who plays WoW, not a generated badge. v3's gloss, specular streaks and
the drop-shadow wedge under the coin are gone; the finish is matte, with painted wear on the rim.

The full-res cut at 3x on footer brown (#170e09) and cream (#f3eee2): clean on both, no halo, light fringe, mushy feather
or leftover ground colour. About 8.8k of roughly 819k opaque pixels are semi-transparent, consistent with the 1.5 px
feather. The rim band is smooth without looking plastic, so the rim-only denoise is invisible; flecks and rune wear near
the rim are kept and read as intentional.

## Tells considered

- 15 (generated illustration): the art is the owner's choice; the cut adds no gloss or airbrushed sheen.
- 12 and 1 (halo, glow): none. Without v3's shadow wedge, v4 sits flat on nav, footer and light grounds.
- 11 (centred composition): the Discord banner centres seal over wordmark, the same layout as the banner already
  uploaded, so nothing new.

## Notes (non-blocking)

1. The near-black outer rim (about #2a1a0e) sits close to the footer's #170e09, so the coin's outermost edge nearly
   disappears there; the gold ring still defines the circle (footer-after@2x.png). A 1 px lighter keyline if the footer
   ever goes darker.
2. Minimap 32 and list icon 64 stay legible (LX, teal field, gold ring; the rune ring turns to texture, as expected). A
   hand-sharpened 32 px pass would improve the minimap.
3. 16 px favicon is a teal and gold dot with no readable LX, the same as v3; the tab mock holds on dark and light.
4. Discord banner: the seal is small (about 230 px of 1920) and the wordmark carries it; the edge is clean over the
   tavern art with no halo against the lantern tones. A larger seal or an offset layout would be less templated, optional.
5. Header at 1440: the 40 px seal keeps ring and LX, flush beside the serif wordmark, no glow or added ring.
