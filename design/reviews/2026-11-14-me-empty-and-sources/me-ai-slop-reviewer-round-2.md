# ai-slop-reviewer, round 2 (Me with no site activity)

Verdict: FAIL

Round 2 is much closer. The changes that worked: the l60 sim now reads as a real state, the duplicated "on plan" is gone, and the gear rows now name a source. I have no visual tells. The remaining fails are small but concrete.

**Must-fix**
1. **Tell 13: l5 right column (me-empty-l5).**
   - "Right now" has one yellow line, "Pick your spec", and roughly 250 px of empty panel under it.
   - Fix: drop the column and let the left panel span the width. Or fill it with real level 5 data such as the next trainable spells, Hunter starter zone ranges, or the first talent point at level 10.
2. **Tell 22: the Next row repeats the "Next point" row (l25, l26).**
   - "Level 30: your next gear list" has "your next" where "Level 30: gear list" would do. "Level 10: your first talent point" has the same extra words.
   - The reviewer asked to keep this wording, so only the redundant "your" and "next" are the owner's call. I would cut them.
   - Also, "Next point ... Level 26" and "Next ... Level 30" with the same label root read as a labelling clash. Rename the row to "Gear list" with a value of "Level 30".
3. **Tell 22: "No enchant: Enchant Chest - Minor Stats".**
   - "No enchant:" followed by "Enchant" says the same thing twice. Use the red slot name and the enchant: "Chest: Minor Stats".
   - The hyphen in the item name ("Enchant Chest - Minor Stats") is also a dash substitute. Strip the "Enchant Chest -" prefix.
4. **Truncation is visible in three places.**
   - "Captain Green...", "· Che...", "· ..." and "Enchant Cloak - Greater Defense ..." are cut mid-word.
   - The second enchant row at l25 ends with a bare "· ...". This reads as unfinished.
   - Fix: wrap to a second line or shorten, and drop the trailing slot name after the enchant since the enchant already names it.
5. **Tell 13: l25 and l26 left panel dead space.**
   - There is roughly 300 px of empty panel under "Next" at l26 (the site line is gone there because the companion is synced), and the left panel is shorter than the right.
   - Fix: move Professions or the gear-source lines to the left, or collapse the panel height.
6. **Tell 19: the l60 left panel is nearly empty.**
   - It holds Build and Professions plus one grey sentence ("Log a raid and your score and logs come in here.") over about 450 px of blank.
   - The sentence is also the pitch-copy pattern. It is a promise, not information.
   - Fix: delete it, as you did below 60 with the companion. Fill the left side with real raid-prep data (consumables, resistances, attunements) or move the Gear block left.

**Shoulds**
7. "Your sim" over "Not simmed yet" and "Run a sim" says the same thing three times. Keep one line: "Sim: not run" plus a button.
8. "8 of 14 best in slot": the brand reviewer's decision is accepted. Shorten the number to "8/14".
9. The footer right is "No companion" at l25 and l5 but a sync age at l26 and l60. That is state-driven and fine, but "No companion" repeats the site line above it. Leave it in the footer only.
10. "Spend 1 point: Improved Aspect ..." at l26 is truncated and duplicates the Next point row. Remove it.

**Authentic, keep**
- Real item names, quality colours and DPS deltas.
- Source lines such as "Ranged · World drop" and "Feet · Scholomance · Kirtonos the Herald".
- Real talent icon and "Skinning 112, Leatherworking 105".
- The Onyxia's Lair raid line.
