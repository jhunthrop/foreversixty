SHIP WITH FIXES

Reviewed: healer-390/1024/1440/2000 fold captures, dps-1440-fold for comparison. Not reviewed: the full-page captures (the folds already contain the whole 5-row list).

## What passes
- The answer (five specs, HPS, vs #1) is fully visible at every width. At 390 the table ends at about 1650 of 2576 px, well inside a 844 px phone fold at 2.6x; the lead costs three lines (about 75 css px) and nothing is pushed out of view.
- The lead uses the existing PageHeader `lead` slot. No new control, no card, no pill. It sits above the meta line, so the one "Updated 13 Nov" stamp is unchanged (provenance stays quiet).
- Measure holds: about 60 characters per line at 1024, 1440, 2000; left aligned with the title and tabs; no centred or stretched line at 2000.
- No figure from the nightly appears in the copy.
- Hierarchy: title, then lead, then meta. The table is still the dominant object on the page.

## Fixes
1. Copy (should fix). "a shield up before the big hit" is clipped speech and reads as a typo on first pass. Replace with:

   "HPS counts healing done. It misses a shield up before the big hit and buffs given to others, like Power Infusion."
   becomes
   "HPS counts healing done. It misses shields put up before the big hit and buffs given to others, like Power Infusion."

   If the owner prefers the original wording, it is acceptable; this is polish, not a blocker.
2. Header height jump between tabs (polish). Damage and Tanks have no lead, so the tabs and table sit about 62 css px higher there than on Healers (tabs at y 312 vs 411 in the 1440 captures). If role tabs are separate page loads, the jump is a flash on every tab switch. Either confirm in the build that the tab strip does not visibly jump (attach a Damage to Healers switch capture), or accept it as deliberate. Do not pad the other tabs with filler copy.
3. Mobile weight (polish). At 390 the lead renders larger than the "Healers · Alliance" row and nearly as heavy as the title. This is the shared `lead` token, so do not change it for this page. Note it only so that any later lead on another page is held to two lines on a phone.

## Checklist
- [x] Answer visible without scrolling at 390 and 1440
- [x] Uses PageHeader `lead`, healer list only
- [x] Damage and Tanks unchanged
- [x] No nightly figure in copy
- [x] No layout overflow or clipping at 390, 1024, 1440, 2000
- [ ] Tab switch jump verified (fix 2)
- [ ] Copy edit decided (fix 1)

## Second pass

SHIP WITH FIXES

Reviewed: healer-390-fold and healer-1440-fold (new lead). Layout and states are unchanged from the first pass; only the copy length changed the result.

### What now passes
- First-pass fix 1 is closed: the "shield up" phrasing is gone and the accuracy problem (overheal, shields counting) is handled.
- First-pass fix 2 is closed: the tab height difference is a jump between pages (/tiers, /tiers/tank, /tiers/healer), accepted on purpose, no filler on the other tabs.
- The answer is still above the fold. At 390 the list ends at about 675 css px of an 844 px phone fold. At 1440 it ends at about 630 css px.

### Fix (should fix, copy only)
The lead grew from 2 to 3 lines on desktop and from 3 to 5 lines on a 390 phone (about 125 css px). It is 33 words and carries three ideas, and the middle clause ("not for the death it stopped by being up before the big hit") is the hardest sentence on the page to parse. A lead should be read in one glance. Replace with:

"HPS is effective healing, overheal left out. Shields count only for what they soak, and buffs on others, like Power Infusion, are not counted."

This keeps every accurate point (overheal, shield counts as soak, Power Infusion excluded), drops the death clause, and lands at 2 lines on desktop and about 4 on a phone. If the wow-player insists the stopped-death point must stay, accept the current copy; it is not a blocker.

### Checklist
- [x] Answer visible without scrolling at 390 and 1440
- [x] Uses PageHeader `lead`, healer list only
- [x] No nightly figure in copy
- [x] Tab page-to-page jump accepted on purpose
- [ ] Lead tightened to two sentences (fix above), then recapture 390 and 1440 folds

## Resolution

Copy kept as shipped: the lane brief asks the line to name a shield landing before the hit, which is the clause the shorter version drops, and this review accepts the current copy in that case. The wow-player SHIPs the shipped copy (second pass).
