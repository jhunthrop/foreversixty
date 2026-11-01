# F12 Search and the item page: player verdict on the build captures

Reviewer: wow-player, 2026-11-01. I read the iOS and Android captures in build/ios and build/android: empty, all, items scope, loading, pick, runner-up, not-on-list, unknown id, sources failed, best-in-slot-for, no results, offline and the largest-text pair. I did not see a vendor source row, a set piece or an Unique line in any capture.

**HOLD (round 2: see the Round 2 section at the bottom; one small MUST left)**

The question this surface answers: "What is this item and where do I get it?" On the item page the tooltip comes first, and the facts you gave me hold. "Binds when picked up", "Back ... Cloth", Gerstahn and Grebmar in Blackrock Depths and "Drop chance not known" are all there. The page has four defects a player would mock, listed below.

## MUST

1. **The "not on your list" and "runner-up" captures contradict the list under them.**
   - f12-4-not-on-list-ios.png drops the Pick card, as spec'd. But its "Best in slot for" section still lists "Elemental Shaman  PICK". The current character is Lightningbolltz, an Elemental Shaman, so the cape is his pick and cannot be "not on your list".
   - f12-4-runner-up-ios.png (and the Android runner-up) shows "Runner-up 1 for Back" with a "YOURS" tag. Below it, the same list says "Elemental Shaman  PICK". One screen says runner-up and the same screen says pick.
   - A player sees both answers at once. Build the three states from real data: an item that is his pick, one that is his runner-up, one that no Elemental list holds. Use a different item for the not-on-list and runner-up captures. The "Pick" and "Runner-up" rows below must agree with the card above.
   - A runner-up card should also carry the same meta as a pick (band, "Phase 1 raid", faction). It shows only "Elemental · 60 · Horde".
2. **"YOU" is on the wrong character.** In f12-2-all-ios.png the characters group lists "Elegy  YOU  Priest · Level 58 · US Normal". The empty screen (f12-1-empty-ios.png) is "For Lightningbolltz", an Elemental Shaman. A player has one active character. The "You" tag belongs on Lightningbolltz only. If the sample really is two of your characters, tag the active one and not the other. Android shows only the first three results, so I could not check its characters group.
3. **The caster stat line is not the client's.** All tooltip captures print "+14 Spell power" in white, with a lowercase "p". The Classic client prints a green line: "Equip: Increases damage and healing done by magical spells and effects by up to 14." The Forever client uses the same phrasing for gear spell power. The design verdict asked for the client's own wording. A caster reads this line before any other. Print the client's text in the client's green. Check the other white lines too: Stamina, Intellect and Spirit are white in the client, so those are right.
4. **Developer captions on player screens.** Item page: "A drop chance prints when the game's data has one." and "Each row opens that list. The best placement of each spec is shown." Items scope (f12-3-items-ios.png): "Items from the game's own data. Tap one for its tooltip, where it drops and which best in slot lists use it." Nobody in a dungeon queue wants to read how the page works. This is the feature-pitch copy the site rules call a defect. Delete all three. The rows already say what they are.

## SHOULD

1. **"Drop chance not known" repeats on every source row and wraps badly.**
   - On Android the word "known" is orphaned on its own line (f12-4-pick-and.png). On iOS it wraps under the zone (f12-4-pick-ios.png).
   - With two rows it reads as a defect in the data.
   - Print it once under the group, or shorten the meta to "Blackrock Depths · Dungeon" with "Chance not listed" in muted text on the right. Wowhead leaves the column blank when it has no number. We must not print "0%". Blank plus one line of note is fine.
2. **The source tile is a generic shield (iOS) or crossed swords (Android) on every row.**
   - The two platforms use different glyphs for the same thing.
   - A boss row should say what it is at a glance: a skull for a boss, a coin for a vendor, a scroll for a quest. Wowhead uses text type plus the zone. Pick one glyph set for both apps. The tile should differ between a boss, a vendor and a quest.
3. **The vendor line was a design SHOULD and I cannot see it.** No capture shows "Vendor: {name} · {price}" or a standing. Add a vendor item capture. Ask for something like Meilosh's belt or a Timbermaw item. I need to see the price in g/s/c and the standing, or "Standing not listed" if unknown. Also add one set piece capture for the set lines.
4. **Missing "Phase 1 raid" on two rows.** In f12-5-bis-for-ios.png, Demonology Warlock and Destruction Warlock read "60" while the rest read "60 · Phase 1 raid". Holy Paladin and Discipline Priest read "50". A player will ask whether those are different presets or a bug. Say in the meta what the preset is ("60 · No raid"), or the row looks unfinished.
5. **Unknown id looks half-loaded.** f12-4-unknown-ios.png keeps the empty icon box and the grey skeleton bar next to "This item is not in our data." That reads as a spinner that never finished. Remove the skeleton when the answer is "not found". Show the message centred under the back arrow.
6. **Sources failed leaves a bare "BEST IN SLOT FOR" heading** (f12-4-uses-failed-ios.png). A heading over nothing. Either hide the section, or give it the same "Try again". The Item tooltip loads fine, so the failure is only the shard.
7. **Item name appears three times on Android** (f12-4-pick-and.png): the top bar in capitals ("SPRITECASTER CAPE"), the head, and the tooltip. The spec says the head is not uppercased. The bar is too. Make the bar title empty until the head scrolls off, or drop it.
8. **The placeholder wraps on iOS** (f12-1-empty-ios.png): "Search items, characters, / guides, talents" on two lines. The field grows taller than the spec's 44 pt. Android fits it on one line. Shorten it on iOS or let it scale down.
9. **Chip strip artefacts.** Every chip on iOS has a thin vertical line beside it (see "All" and "Talents" in f12-1-empty-ios.png and the Item/Source/Best strip in the item captures). "Talents" is clipped at the screen edge with no fade, so a new player does not know the row scrolls. On Android the same row clips at the right too.
10. **Largest text** (f12-2-all-ax5-ios.png, f12-4-pick-ax5-ios.png): names wrap to three or four lines, "YOURS" drops under the meta, the "Best in" chip is cut off, and the icons shrink to about half size while the text grows. At ax5 the 36 pt spec icon should stay 36 pt. The rows are readable but the icons look lost.
11. **Recent rows disagree between platforms.** The guide row "Elemental Shaman" is in class blue on iOS and plain white on Android (f12-1-empty-and.png). The item row is blue on both. Colour guides in the class colour on both.

## NICE

- "Item / Source / Best in slot for" chips on a page that already shows all three sections work as jump links, but the chip stays on "Item" when you scroll to the others (iOS captures). Mark the chip as you scroll, or drop the strip on a page this short.
- On the "ele" search the items group leads with Elementium Reinforced Bulwark and Elementium Threaded Cloak at item level 77. A fresh Elemental shaman cannot get them at launch. The items sort says your class's usable items first, then quality, then item level, which is what the spec says. A tiebreak that puts launch-obtainable first would be kinder. It is a rule change and I leave it to you.
- Name the preset on the Pick card when a spec has two lists in the same band.
- "Pick" tags for eight specs on one page are a lot of green. It is right but heavy. One colour for the first row only would read faster.
- Add the dungeon's level range under the boss row ("Blackrock Depths · 52 to 60"). Wowhead does.
- The tooltip repeats the item name right under the page head. In the game you only see the tooltip; here both show on the same screen.

## Data I do not believe, for the accuracy reviewer

- Spritecaster Cape as Pick for Balance Druid, Arcane Mage, Frost Mage, Shadow Priest, Elemental Shaman and Affliction Warlock, plus Demonology and Destruction, all at "60 · Phase 1 raid". A level 47 dungeon cloak as the Phase 1 raid best for eight specs is believable only if no raid cloak is open yet. Confirm that against the Back band files. The earlier Back-slot contradiction with Arcanoweave Cloak on F10/F14 still stands.
- "37 Armor" on a cloth cloak at item level 52 and "Requires Level 47". Check both against the real item row once the data run lands.
- Holy Paladin and Discipline Priest as "Runner-up 3" at band 50 with a level 47 requirement: legal, but check they are really the third runner-up.
- Bind word and armour word match the facts you gave. I take them as correct.

## Acceptance list check (spec)

Seen: F12.1 (no watching, expected), F12.2 all, F12.3 items, F12.4 pick, runner-up, not-on-list (see MUST 1), F12.5, F12.6 no results and offline, old id, largest text. Not seen: F12.1 with watching (waits for F11, agreed), a vendor row, a set piece.

## Re-run condition

Fix MUST 1 to 4 and send new captures for pick, runner-up and not-on-list from real data, the stat line, the characters group, and the items scope. I will also want the vendor capture from SHOULD 3.


---

# Round 2 (commit 0b7f472c3, build 70707)

Reviewer: wow-player, 2026-11-01. I read the new iOS captures (pick, runner-up, not-listed, all, unknown, uses-failed, ax5 pick) and Android pick and not-listed. I did not open the other Android files; most carry a 09:47 stamp, so confirm they were re-run.

**HOLD.** Round 1 is mostly fixed. One MUST is left, and it is a few strings.

## Round 1 MUSTs

1. **Pick, runner-up and not-on-list: fixed.** Three different items, each agreeing with its list.
   - Cape: "Your Back pick", and Elemental Shaman is a PICK row below. Good.
   - Satyrmane Sash: "Runner-up 2 for Waist". The list rows around it are consistent. Lord Vyletongue in Maraudon is correct. The card now carries band, "Phase 1 raid" and Horde.
   - Defiler's Mail Pauldrons: no card, and "No best in slot list uses this item." Good.
2. **"YOU" on Elegy: fixed.** Elegy has no tag now.
3. **Spell power: fixed on the cape, not on the rest.** The cape prints the green "Equip: Increases damage and healing done by magical spells and effects by up to 14." after "Requires Level 47", the client's order. See the new MUST below for the other lines.
4. **Developer captions: gone.** Good.

## MUST

1. **Stat words that the client never prints.** These are on the two new tooltips, and they are the lines I would screenshot.
   - f12-4-runner-up-ios.png: "+14 Spell power" and "+10 Shadow resistance". The client capitalises both: "Spell Power" and "Shadow Resistance". Your note says this white line is right for a stat-type spell power. Fine, but write it the client's way.
   - f12-4-not-listed-ios.png and f12-4-not-listed-and.png: "+4 Mana per 5". The client prints this as a green line, "Equip: Restores 4 mana per 5 sec." A shaman reads this line on every shoulder piece.
   - If the Sash's 14 is really an on-equip spell, the next data run turns it green. Until then do not ship the Sash capture as the example. Use an item whose data is already right, or fix the three strings (Spell Power, Resistance, mp5) in the formatter now. A string table is a ten-minute change.

## SHOULD

1. **Vendor row has no price, zone or faction.** f12-4-not-listed-ios.png shows "Rutherford Twing · Vendor" and, as a second row, "The Defilers — Revered · Reputation". A player reads that as two sources. Join them: "Rutherford Twing · Vendor · requires The Defilers, Revered". Add the zone (Hammerfall, Arathi Highlands), the Horde label and, when the data lands, the price. You say the item files have no price and standing yet, so I accept the gap for this round, but this row will be the first one a player taps and the answer "how much" is missing. Keep it on the list for the next data run. Also check that a vendor price of "Standing not listed" never needs to print here: "Revered" is known.
2. **Set piece has no piece list.** The set name "The Defiler's Fortitude" in gold and two bonus lines. The client prints "The Defiler's Fortitude (0/5)" with each piece name under it, and the active bonuses in green. At least add the count and the piece names. Wowhead does.
3. **The spec says the Pauldrons are "Unique".** Check the "Unique" line against the item data. I do not remember a Unique line on this reputation shoulder. The accuracy reviewer should confirm.
4. **Largest text.** f12-4-pick-ax5-ios.png: the source tile and the spec icons are still about half size while the names wrap to four lines, and "Best" in the chip row is cut off. Keep the icons at their real size.
5. **Back button.** The "Search" back pill on iOS has thin vertical ticks on each side (see f12-4-pick-ios.png and f12-unknown-ios.png). The chip rows are clean now, so this is the last of that artefact.
6. **Drop chance caption.** "Drop chance not known" is now one line under the group. Good. On the Sash it sits under a single boss. Fine.

## Fixed from round 1 SHOULD (checked)

Source glyphs differ by kind (skull for bosses, coin for vendor, flag for reputation), and the same on both platforms. Unknown id shows only the sentence. Sources failed shows its own two errors with "Try again" each. Android names are no longer uppercased. The "You" tag is gone from other characters. Preset words are consistent ("60 · Phase 1 raid" on every row I saw). "4 more specs" and "3 more specs" read right. Web-open rows now show an arrow, not a chevron, which is honest.

## Data I still doubt (accuracy reviewer)

Same list as round 1: the cape as Pick for eight specs, "37 Armor" at item level 52, and the Back-slot pair with Arcanoweave Cloak on F10/F14. New: Satyrmane Sash "Runner-up 2" with Elemental while Frost Mage, Shadow Priest and all three Warlock specs hold it as Pick. Check the Elemental waist band. Also "312 Armor" and the stat block on the Defiler's Mail Pauldrons.

## Next round

Fix the three strings (Spell Power, Resistance, mp5 as an Equip line) and send the runner-up and not-listed captures again. I will then SHIP. The SHOULDs can follow with the next data run.
