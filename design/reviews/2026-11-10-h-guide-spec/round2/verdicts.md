# h-guide-spec round 2 verdicts (captures: round2/captures, real-data build at 17f0d13f4)

## ai-slop-reviewer
VERDICT: PASS WITH FIXES

1. Priest Holy priority rows 7-18: Greater Heal rank 4, Greater Heal rank 1 and Heal rank 4 each repeated per
   threshold in one sentence frame (tell 19/21). Fix: one row per spell and rank.
2. Paladin Retribution Enchants and consumables: no prose, so an empty column with the lists beside it.
3. Retribution off-hand row "The Unstoppable Force is a two-hander; the off hand is taken." beside a blank box.
4. Survival 2000: Races and Enchants columns uneven (Horde under the prose).
5. Stat priority: ~200 px of dead prose column beside the table (unchanged from round 1).
6. sim-specs 1440: "Healers and tanks" read as a small grey label.
Out of lane, logged: the in-game panel wording "mana % is over boss health %", the hit-cap detail line.

## ux-designer
VERDICT: SHIP

Round 1 must-fixes resolved (Professions one block at the measure; factions whole; consumables under their
own labels). Shoulds: Priest priority rows 7-18 near-identical (collapse to one row per spell); Horde under
the prose at 2000; Stat priority dead space (allowed for now); Orc's two racials on separate lines.

## wow-player (round 2 content diff and the rotation-collapse change)
VERDICT: SHIP

Restores confirmed; Survival's Mark and Movement cuts lose no fact the page does not still carry. Caution:
rotation-collapse dropped all generated words when only the first clause matched a note; fixed in round 3
(a note that is the target alone keeps the words' later clauses).

## Round 3 changes (f615da540)
One row per run of a spell and rank, rules joined in order; extras with no prose side by side; no empty off
hand row under a two-hander; the collapse keeps unmatched generated clauses. Items 4 and 5 left per the
ux ruling (whole blocks wrap under the prose by the site-wide reading-wrap rule; the stat table is one block).
