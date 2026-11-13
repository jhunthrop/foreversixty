# Tier share card, round 1 fixes

Applied from the ux-designer, wow-player and ai-slop-reviewer verdicts:

- Damage card: three specs a line, names 22 px, crests 38 px, meta line 22 px.
- Every card names its faction beside the role ("Damage · Alliance"): the numbers are Alliance's, the page default.
- Tank and healer rows carry the figure they are sorted on, with the page's column unit (HPS, Taken / s).

Kept as the page has them (the card repeats the page, it never restates it differently):

- The meta line per role ("3 min", "110 s", the tank line without a length) is tiersCopy.metaLine.
- Tier C is the uncommon quality green, as the page's letter is (TIER_LETTER_TONES).
- Order inside a tier is by DPS, as on the page.
