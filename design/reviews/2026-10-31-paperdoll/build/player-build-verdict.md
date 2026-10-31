# Paperdoll build: wow-player verdict (2026-10-31)

Seen: side-beta-1280, side-raid-1280, report-1280, tip-bind-1280, tip-right-1280, character-beta-390, discord-sheet. Not looked at: the 360/1024/1440/2000 captures, longest-*, loading/failed, tip-phone*, tip-flip/left.

## Verdict: HOLD

Desktop matches the mock and reads like a character screen. Quality colours, slot order, "Best in slot", the two-hander empty off hand and the beta "Item NNNNNN" slots are all right. Three things a player would mock are left.

## Blockers

1. Item level contradicts itself on one item (tip-bind-1280, report-1280). The slot row says "Lionheart Helm / Item level 60 · Enchanted". The tooltip on the same item says "Item Level 65". Both come from the same record. Pick one source. Also check that the item level is right at all, because Lionheart Helm is a level-60 crafted piece.
2. The report header shows a number that breaks the rule (report-1280). It reads "Baelgrim  Protection · item level 183  8 items" over a panel that says "Item level shows once every item is known". 183 is not an average of anything here. The eight known items are about 60 to 65. It looks like a sum or a score, and it appears while two items are unknown. Drop it, or show the real average (about 63) only when every item is known. The same goes for "Sunwick  Holy · item level 179".
3. Phone character page is hard to read (character-beta-390). Worn items and picks are split in two. A row of two worn items ("Ghostshroud" and "Devilsaur Gauntlets") comes first. Then separate rows carry the slot word, the pick and its gain ("Head · Lionheart Helm +60.0"). The worn item has no slot word and the pick has no worn item beside it. A player cannot tell which pick replaces which item, and left and right worn items are paired as if they were related. On phone, put each slot on one row: slot word, worn item, then pick and gain below it. Keep the desktop layout.

## Shoulds

1. Wrong profession and a bare source line on the tooltip (tip-bind-1280). It reads "Crafted: Tailoring" and then a lone "Quests". Lionheart Helm is Blacksmithing in Classic. "Quests" under a crafted item is meaningless; a source line should name the quest or say nothing. Check this against the data. This is fixture data, but the wrong profession should not exist in a fixture a reviewer sees.
2. Percent stats on the tooltip (tip-bind-1280, tip-right-1280). They print "+2 Crit" and "+2 Hit", "+14 Crit" and "+28 Attack power" in green, as if they were base stats. The client writes "Equip: Improves your chance to get a critical strike by 1%." or the rating form. Per the owner's brief these are ratings, so say "+14 Crit rating" or use the Equip wording. A bare "+2 Hit" has no unit.
3. Set line without progress (tip-right-1280). "Devilsaur Armor" is in bright green with no "(0/2)" and no piece list or bonus. The client shows the set name, the pieces and the bonus text. Either show them or drop the green line.
4. Mock and build differ on the export stamp (side-beta-1280). The mock says "ADDON EXPORT, 12 MIN AGO"; the build says "ADDON EXPORT" with no age. The age matters to a player who is deciding if the gear is current.
5. Discord sheet footer (discord-sheet). "Gold: gain over what is worn" is gold-coloured, and so is "Gain not modelled yet" on several rows, so the key and the exception look the same. Make "not modelled" muted. Also "Trinket 2 · Upgrade available" with an empty socket is fine, but it should say what to buy, or point to the BiS page.
6. The Discord sheet's "Signet Ring of the Bronze Dragonflight" wraps to two lines and shrinks its font against its neighbours. Acceptable, but the font change looks odd. Keep the size the same and wrap.
7. Item level on the tooltip does not match what the player would check on Wowhead for Devilsaur Gauntlets (shows 58, Requires Level 53). Accuracy to verify, not a layout issue.
8. The tooltip still lacks the armour type and "Binds when ..." on real data (tip-right-1280: "Hands" alone). This is accepted until the nightly run, but it must be confirmed in the real build after the run. tip-bind-1280 shows the intended layout and it is right.

## Data not believed (for the accuracy reviewer)
- Lionheart Helm: profession, item level (60 or 65) and the "+2 Crit / +2 Hit" rows.
- Devilsaur Gauntlets item level 58 and Requires Level 53.
- Ghostshroud as the Arms warrior's worn head, if the armour type is legal by the owner's rules.
- Marshal's Chain Legplates and Belt of Preserved Heads (item level 70 and 71) on the Discord hunter sheet: confirm the rank and that the belt is launch-obtainable.

## Fixed since the spec check, confirmed
- Two-hander: empty off hand reads "Two-hand in main hand" on the site and "Off hand · Two-hand in main hand" on Discord. The item level of 63.5 on the sheet counts it as known. Right.
- Item level hidden on beta-1280 with the explanatory note. Right.

## Re-check (fix round 1)

Seen: report-1280, tip-bind-1280, tip-right-1280. Not re-looked at: character-beta-390 (answered from the spec), the other re-shot captures.

### Verdict: SHIP (no blockers left)

1. Item level mismatch: closed. Row and tooltip both read 65 for Lionheart Helm (report-1280, tip-bind-1280). I accept that the real item is 61, required 56. The "Requires Level 60" line on tip-bind-1280 is fixture data and does not match that. Check that the real tooltip says 56.
2. "item level 183": closed as a number. report-1280 now says "item level 66", which matches the injected gear. One thing remains: the header still shows a figure while the same panel says "Item level shows once every item is known" with two unknown items. That reads as a contradiction. It is the log's own figure and outside this lane, so I am not blocking, but do not let it ship unexplained. See should 1.
3. Phone layout: I withdraw it. It is the approved §3.3 design and not a spec breach. I still find the unlabelled worn row hard to scan, but that is a preference.

### Shoulds, none blocking
1. Label the header figure as the log's ("log item level") or hide it while any item is unknown.
2. The tooltip set line has no (0/2) and no bonus text. The agent has noted it as a follow-up. Track it.
3. Re-check real-data tooltips after the nightly run for "Binds when ..." and the armour type; tip-right-1280 does not show them yet.
4. Discord "Gain not modelled yet" is still the same gold as the key. Follow-up.
5. Confirm Lionheart Helm's real stats rows ("+2 Crit / +2 Hit") once the real record shows.
