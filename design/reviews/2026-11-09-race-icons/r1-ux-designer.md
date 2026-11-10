Verdict: SHIP

Scope: race-icon change only. Reviewed after-* captures for hunter, mage (class landing) and hunter-bm, warrior-fury (spec page) at 1440 and 390.

Class landing, Races rail panel (reference: in-game character-creation race list, faction-grouped)
- Alliance then Horde, small muted group labels, 24px portrait with the warm 1px frame, name in faction colour. Matches the spec given. Icons sit on the same left edge as the names, rows are even, no layout shift, no empty sockets.
- 1440: single column in the 340px rail, 8 rows for hunter, answer visible above the prose. Fine.
- 390: two columns, wrapped names ("High Order Skyborne", "Windshaper Skyborne") wrap to two lines with the icon vertically centred and no row collision. Fine.
- Faction colours on the dark panel read clearly.

Spec Races lane (reference: Wowhead racial list)
- 28px portrait leads each race row, the Recommended tag stays in the trail (1440 right-aligned, 390 wraps under the name). The racial icon rows (20px) below are visibly secondary to the 28px race portrait. One level of emphasis kept.
- Portraits do not shift the name baseline on any race.

Findings
- Polish: High Order Skyborne and Windshaper Skyborne both show the Blood Elf portrait, so two different rows in the same list look identical. Accepted given the client ships no Skyborne portrait; the faction colour of the name tells them apart. Revisit if a Skyborne portrait ships.
- Polish: on the hunter landing the prose under the list repeats the same races in a sentence ("Alliance hunters are Dwarf, Night Elf and..."). Pre-existing copy; trim in a content pass, not this change.
- Polish: the race name is the icon's title only, which does nothing on touch; harmless because the name is always printed beside it.

Checklist
- [x] Icon from client portrait, 1px warm frame, same radius language as SpellIcon
- [x] Alliance before Horde, faction colours
- [x] No layout shift, width and height set, lazy loaded
- [x] Missing portrait draws nothing, name alone remains
- [x] 360-390 and 1440 both hold, no overflow
- [x] Hunter, mage, warrior-fury, hunter-bm all render icons

No must-fix items.
