# L3 Gear tab: ai-slop-reviewer

## Round 1: FIX

First impression: reads as a WoW addon made by a player. Paperdoll with real quality colours, real item icons, real names (Black Dragonscale Helm, Cape of the Black Baron, Pix Xizzix), real zones, no gradients, glows, cards, eyebrows, icon grids or hero line. The tells that remain are copy and layout, not visual.

Figures checked (all agree unless listed under Must):
- Hunter 60: 16 rows = 16 slots; 9 "Best in slot" cells counted = header 9; 2 "No enchant" (Bloodmail Boots, Dawn Armor) = "2 enchants missing". Doll gains sum 17.1 = "+17.1 DPS in all" (spec: total over every slot, so the four rows, 13.0, rightly differ).
- Warrior 34: 5 "Best in slot" = header 5; 1 "No enchant" = "1 enchant missing"; denominator 15 = 17 rows minus the two target-less trinkets (GearTab.lua `scored`). Doll gains sum 4.3 = "+4.3 DPS in all".
- Weights: every DPS column equals weight x top-stat DPS within rounding, both renders. Source counts match the "Raid 0" rule in spec 4.3.

### Must

1. Truncation reads broken. Warrior Upgrades: "Boulderfist Enforc...", "Scarlet Mon...", "Razorfen Kr...", "Charlga Razorfla..." (3 of 4 rows), plus the doll's "Brawler's Leather Leggua...". Cutting boss and zone names to a stub is the one place a player cannot read what the row says. Drop the "World" word before the drop chance, abbreviate the zone not the NPC, or wrap to a second line; the Upgrades column has dead width at right of the gain only on rows that are short.
2. "Beats the pick" is a third state the header ignores. Hunter Slashclaw Bracers and warrior Tusken Helm / Wolfmaster Cape are worn items better than the list's best, yet "9 of 16 best in slot" and "5 of 15" do not count them. A veteran adds "best in slot" cells, sees 9, and sees a tenth row that beats it. Count them in the header, or say "Beats the list" and count it separately. Also "the pick" is internal jargon; a player says "list".
3. Warrior Sources says the same thing twice: right note "No rep above Friendly" and the footer "Off: Reputation above Friendly". Keep one (the footer carries the Off/Default state; the note is the redundant one when exactly one kind is off).
4. Hunter Sources right note "All sources" under a heading "Sources" with every box ticked. Tell 22: a qualifier the context already gives. Show nothing when nothing is off.
5. `L.panelTabOpens` = "Opens Gear: what you wear, what beats it, where to get it." A triple-clause pitch that explains the tab to a veteran who can read "Gear" (tells 19, 21). Cut to "Gear" or drop the tooltip.

### Should

6. "shipped, per point" (Stat weights note) is a dev word. A player does not know what ships. "default, per point" or "Forever's, per point"; the sim variant "your sim" is fine.
7. Shipped weights drop the error column and, on both renders, the haste line, where spec 4.3 lists error and haste per 1%. If shipped weights carry neither, fine, but then the table has no way to show Haste at all; a hunter or warrior without a Haste row reads as an omission only if the data has it.
8. Hunter 60 render has no Conditions row (`Raid-ready` / `Bare`) although spec 4.3 and the approved board show it at 60. Confirm the fixture has no presets and not a regression.
9. Dead space: the hunter doll ends at y 960 and the left column is empty down to 1290 (about a quarter of the window), while the right column runs full. Not a tell, but the window looks lopsided; anchor the sources or let the doll rows breathe.
10. Hunter Upgrades row 2 "Finger · Pix Xizzix · Vendor": the NPC and the word Vendor say one thing. The board's "Finger · Pix Xizzix" was enough.
11. "Best in slot" repeated nine times is correct but flat. Mute it harder than the gains (it already is), or print nothing for best and let gains and "No enchant" be the only text; the doll then scans.

### What reads authentic (keep)
- Quality-coloured names and real icon edges; item rows two-column mirrored like the character pane.
- "in your bags" in client gold, "No enchant" in the warning colour: state by role.
- Gain figures that take only "+4.9 DPS" and a unit; no "vs your item" qualifiers.
- "Upgrades +17.1 DPS in all", the single total, summed from the same gain function.
- Raid 0 greyed at 60 (true to the phase), the Off: line with Default.
- Warrior caption: the leveling default dropping the rep ring to its quest ring (Band of the Fist) is a detail only this product has.
- Locale block: no em dashes, no marketing verbs.

### Verdict
FIX. No visual tells 1, 2, 5, 8, 12, 13, 15. Items 1 to 5 must land before PASS.

## Round 2: FIX

Figures re-checked, all agree. Hunter: 9 "Best in slot" + Slashclaw "Better than the list" = 10 of 16; 2 enchants = Chest (Dawn Armor) and Feet (Bloodmail Boots); Upgrades 13.0 + More upgrades 2.3 + 1.8 = 17.1 = "+17.1 DPS in all" = doll gains. Warrior: 5 + 2 "Better than the list" = 7 of 15; 1 enchant = Legs; top four 3.6 + More upgrades 0.4 + 0.1 + 0.1 + 0.1 = 4.3 = header total = doll gains. Each More upgrades row matches a doll gain by slot. Weight DPS columns still consistent. Round 1 must items 2 to 5 are closed (counts, one Sources note, empty note at All sources, "Opens Gear."), and the "per rating point" and Conditions answers are accepted.

### Must

1. Source tails still get cut to unreadable stubs. Warrior More upgrades: "Veteran's Silvered Chain Leggings · Legs · Crafted · Bl..." (Blacksmithing cut to two letters) and "Brawler's Leather Boots · Feet · Crafted · Leatherwor...". Hunter: "Kirtonos t..." (the boss). Upgrades: "Scarlet M..." and "Razorfen K..." are cut to a letter after the 8-letter floor. Cut whole words only: drop the profession or the trailing field entirely rather than leaving "Bl...", and let "Scarlet M..." become "Scarlet" or the full name by dropping the chance first (the line has "90% drop" after it, so the chance should go before the place is cut). Doll "Brawler's Leather Legguar..." is the same defect, shorter to read but still mid-word; the hover is the fallback, not the layout.

### Should

2. "No enchant: Chest, Feet" is a third statement of the same fact: the header says "2 enchants missing" and each affected cell says "No enchant". Keep the per-cell flag and the header count; the new red line adds a list nobody needs and costs a row. If it stays, drop the per-cell flag.
3. Upgrades and More upgrades are two different row styles for one list (icon 28 with two lines versus a small icon, one line, grey tail). Acceptable as rank 1 to 4 versus rank 5+, but the small rows put quality-coloured names at body size next to a grey tail; keep.

### Verdict
FIX, one must (item 1). No visual tells; the dead left column from Round 1 is gone.

## Round 3: FIX

Closed: no mid-word cuts anywhere. Upgrades, More upgrades and the doll all break at whole words; "Scarlet Monastery" and "Razorfen Kraul" now read in full, the chance is dropped before the place, More upgrades shows a source only when it fits. The red "No enchant: Chest, Feet" / "No enchant: Legs" line is accepted on the info-density argument (the header gives a count, the cell flags are scattered). Figures still agree (hunter 10 of 16, 17.1; warrior 7 of 15, 4.3; rows match).

### Must

1. The word cut is now too eager on the doll and made one name worse than Round 2. Hunter right column: "Sentinel's Chain..." (Sentinel's Chain Leggings). In Round 2 the full name fit in that cell (it rendered whole at the same width), so the cutter now trips before the room runs out; the hunter's legs are the identifying word and it is the one that was dropped. Measure against the real cell width, not a conservative budget, and only cut when the full name does not fit. Warrior "Brawler's Leather..." (Brawler's Leather Leggings) is a true overflow, so a word cut is right there, but it now collides with "Brawler's Leather Boots" in More upgrades; if the cell must cut, drop the first word's possessive or the middle word rather than the last, or leave it to the hover and accept the one overflow.

### Verdict
FIX, one must (item 1, the regression on Sentinel's Chain Leggings). Everything else is clean; no visual tells.

## Round 4: PASS

"Sentinel's Chain Leggings" renders whole again; the early cut was the board generator's width estimate, and the in-game fit measures the real 157 px cell. Warrior "Brawler's Leather..." is a true overflow cut at a whole word with the full name on hover; the More upgrades "Brawler's Leather Boots" is a different slot with its own icon, so the collision is accepted. Figures unchanged and agreeing (hunter 10 of 16, +17.1; warrior 7 of 15, +4.3). No visual tells, no copy tells left. Round 1 to 3 items all closed.

Optional, not blocking: the one remaining truncation on the doll (Brawler's Leather...) is the only stub on either render; if the legs name ever matters, a two-line cell would remove it.
