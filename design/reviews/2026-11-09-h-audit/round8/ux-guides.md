# Round 8 UX review: guides (1440 and 390)

Handed-off items (handoffs.md) are not counted. Blank spell-icon squares are fixture artefacts.

## guides: SHIP
- 1440: compact title band, 3x3 class plates, class-coloured names on plates, crest + spec icons, 44 px rows. Answer (pick a spec) is above the fold. 390: single column, 9 plates, no wrap.
- Polish: about 60 px dead band between the last plate and the footer at 1440 (Druid has 4 rows). No action needed.

## guides-you: SHIP
- "Your spec" row highlight (accent bar + tint) is clear at 1440. The character block (Felsnap, 34 Dwarf Beast Mastery) sits in the band with the next step as a link. Pill wrap at 390 and the "Your guide:" prefix are handed off.
- Polish: at 390 the character block pushes the first plate to y=250; acceptable.

## guides-class: SHIP WITH FIXES
- Must fix: 1440, the "What changed from 1.12" plate ends at y=690 and leaves about 420 px of empty tan beside the Sources column. Run Races beside the plate and Sources full width below, so no column ends 400 px short.
- Must fix: the three spec cards carry only name, talent split and "Read the guide". No role, no top stats, no BiS DPS target; the player picks blind. Add role and the top two stats from the spec payload. Make the whole card the link and drop the three repeated "Read the guide" lines.
- Polish: last bullet "Mail from 40." is a fragment beside full sentences.
- Polish: 390 Races grid wraps "High Order Skyborne" and throws "new" to the far right; set "new" inline after the name.

## guides-spec: SHIP WITH FIXES
- Hero band at 1440 is right: build tree, rotation and weights side by side with Load/Sim buttons, answer above the fold. 390 stacks the same three panels in order, buttons full width.
- Must fix: 390 page is about 17,800 px. Leveling, Races, Professions and Enchants are all unfolded. Fold Leveling 20-60, Races and Professions by default, with a count in the label.
- Must fix: 1440 Overview columns end at very different heights (Strengths 330 px, Weaknesses 190, In a raid 230), leaving empty plate under two of them. Balance or go to two columns.
- Polish: 390 Overview opens with "In a raid" before Strengths; 1440 opens with Strengths. Keep one order.
- Polish: the hit-cap note mixes bold and grey in one paragraph; one level of emphasis.
- Handed off and still present: broken tree glyphs, "On this page / Change" bar, "1 point of agility..." footer, "DPS · Updated", dead plates beside prose.

## Checklist
- [x] Answer above the fold at 1440 and 390 (all four)
- [x] No em dashes visible
- [x] Compact title band, no hero
- [ ] guides-class dead column fixed
- [ ] guides-class spec cards carry data
- [ ] guides-spec 390 length and folds
