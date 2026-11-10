# ai-slop review, sim pages, round 9 (1440 and 390)

Excluded: handed-off items (per-row "Item level N", timeline toggle label, gear stat strip), fixture artefacts (blank icons), haste-per-1% caption.
Round 8 fixes confirmed: talents diff list no longer overlaps the trees; "0 valid combinations" and "20 talent points" gone; Crafted release line gone; weights error labelled at 390; specs has a DPS head and is sorted by DPS; landing-out inner frame gone.

## sim-landing-out: PASS WITH FIXES
- 390: spec links still wrap one orphan per row (Hunter "Survival >", Warlock "Destruction >"); class list about 900 px. Two columns of class rows at 390.
- Authentic: crests, real spec names, "/fs copy".

## sim-landing-in: PASS
- Real gear, talents, sims, rotation prose. Nit: 390 dead band between the character name and the tabs.

## sim-result: PASS
- Dense and data-led. "Level 60 raid BiS 1,004.0 · 9.9% under" is the point. Nit: "Crit on the boss", "Miss 0% specials · 13.9% white" are qualifier-heavy.

## sim-saved: PASS
- Same body as result; "Open in the sim / Copy link", "Saved 4 Nov".

## sim-gear (Top Gear): PASS WITH FIXES
- 390: "Change source" is a full blue button beside "Open in planner" on its own row. Make it a text link like Open in planner.
- 390: Whole sets "Name" input is half width above a full-width paste field; Candidate names wrap beside ilvl and "Copy and modify" (Lionheart / Helm). Stack the row.
- Nit (tell 8, mild): Candidates carries a second inner frame that Try each / Find an item do not. Same double rule on Talent builds, Sources, Published weights and Damage specs. Border by role or drop the inner rule everywhere.

## sim-talents: PASS WITH FIXES
- Fixed: diff list now sits beside the guide trees at 1440 and under them at 390.
- Trees are still about 130 px wide in a 1,240 px plate: dead field right of the list and under the "Your current build" tree. Scale the trees up.
- 390: the +/- list is 19 plain rows under the tree; put it in two columns.

## sim-drops: PASS WITH FIXES
- Raids is the only group with no count and nothing under it but the grey "First raids ... No date yet ..." line, so the open group reads as an empty expander (tell 13, mild). Give it a count like Dungeons "21 sources", or list Onyxia's Lair with its item count.
- Real professions, item counts, ilvl: authentic.

## sim-weights: PASS
- Normalized 1.00 top stat, bars, "10 Hit rating = 1% Hit", per-point DPS, error labelled at 390. Nothing left that is a tell.

## sim-specs: PASS WITH FIXES
- Healers and tanks is still a bare list of names with no numbers (thin plate under a data-heavy one; 390 wraps ragged). Give it a data column or a stacked list.
- Rank "#n" duplicates the sort order beside the bar; nit.
- 390: the "Raid BiS at 60" column label and bars are gone, DPS head remains; acceptable.

## Summary
PASS: landing-in, result, saved, weights. PASS WITH FIXES: landing-out, gear, talents, drops, specs. FAIL: none.
