SHIP

Pass 2. Reviewed: healer-gear and healer-raid at 390 fold and 1440 fold (the new captures). Pass 1 is in ux-designer-pass1.md.

## Pass 1 fixes
1. Phone lead length: fixed. Gear state is 4 lines at 390 (was 6), raid state 5 lines (was 5, now with the length gone and the meter line in). Table starts at y 392 in both (was 440 and 416) and ends at y 688 in an 844 px fold. At 1440 the lead is 3 lines in both states and the table is fully in the fold.
2. Length stated once: fixed. The lead carries no length. Meta line reads "5 min" (gear) and "110 s" (raid), one formatter, same position.
3. Schedule wording: fixed. "with the next update" is player-facing and names no date.
4. Raid fight named: fixed. "a Molten Core or Onyxia kill".
5. Meta formatter: fixed.

## What passes
- Both states say which fight they rank on in the first sentence. The fallback now says it is a stand-in and that it runs every healer dry, so 1,028.3 HPS is not read as a raid number. The raid state says it reads lower than your meter, which prepares the 131.6 figure. Both are player voice, no em or en dashes, no nightly number in prose.
- Hierarchy: title, lead, meta, tabs, table. No new control, no pill, one "Updated 13 Nov" stamp. Left aligned with the title; about 60 characters per line at 1440.
- No overflow, clipping or shift between states at 390 or 1440.

## Polish (does not block)
1. Raid state at 390 ends on a one-word line ("counted."). Optional: "Ranked on one healer's share of a Molten Core or Onyxia kill, so it reads lower than your meter. HPS is effective healing. Shields count for what they soak; buffs on others do not." Check it lands on four lines; if it does not, leave as is.
2. Carried over from pass 1: the Damage and Tanks tabs have no lead, so their tab strip sits higher than on Healers (about 26 css px at 1440 now, down from 62 in the earlier lead). Accept as deliberate; do not pad them.

## Checklist
- [x] Fight named in both states
- [x] Phone lead short (4 and 5 lines; answer in the fold)
- [x] Length once, in the meta line, one formatter
- [x] No schedule date, no nightly figure, no dashes
- [x] No layout shift or overflow at 390 and 1440
- [ ] Orphan word at 390 in the raid state (polish 1, optional)
