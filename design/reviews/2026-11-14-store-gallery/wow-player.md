VERDICT: SHIP (round 2; round 1 was FIX)

Seen (captions re-read after the description file changed: 1, 8, 9 and 11 now read differently and all four still match their plates): all 11 plates at full size, read against the captions in curseforge-description.md and spec 2026-11-14-addon-brand.md (sections 7 and 17). I checked every number that appears on more than one plate.

## What holds up (checked, no change needed)

- Captions match the plates on all 11. Plate 3 counts out: 23 points spent, 1 unspent, 22 on plan of 24 earned, 1 off plan, BM badges 23, 24, 26-30, 31 add to 31, then MM 32-36, 37-41, 42-46, 47 and Survival 48-51 close at 51.
- Plate 4 builds all sum to 51 (31/16/4, 31/20, 31/15/5, 33/13/5, 33/13/5).
- Plate 5 boss percentages average 85 (681/8), downtime 31+24+15+18+4 = 92 s, dead early 42 s, all repeated correctly on plates 10.
- Plate 10 parts add to 1184 (410+380+228+120+46), 1500-1184 = 316, median 74, best Golemagg 92, worst Garr 58. Plate 9 parts add to 1240.
- Plate 6 header "8 of 16 best in slot" matches the rows (6 "Best in slot" plus 2 "Better than the list"), "+21.2 DPS in all" matches the sum of the slot gains and the 21.2 gap on plate 8.
- Plates 7 and 8: 11 In + 1 Late + 2 Bench + 1 Out + 16 not signed = 31, Onyxia role targets 4+10+12+14 = 40, Velisandre is 9th by score in both, 5 missing enchants = Kethra 1 + Velisandre 2 + Saffa 2.
- Class rules: a hunter in leather (Dawn Armor), mail at 60, "Not for you" on Legplates of Might (plate), no raid drops in the lists ("Raid 0"). Dates and weekdays are right for 2026 (15 Dec Tue, 17 Dec Thu).
- Tooltip order is the client's (name, Binds, slot and type on one line, armor, durability, Requires Level, Equip, then our block after a gap). The header line with the seal follows spec 7.

## Findings

1. BLOCKER. Plates 1, 2, 9, 11, every tooltip. The tooltip face has old-style figures, so a zero is drawn as a letter O and a one as a capital I. On screen: "Best for 2O to 29" and "Upgrade +2.O DPS" (plate 1), "Durability 4O / 4O" and "Requires Level 54" (plate 2), "Level 6O Human Priest" and "Forever Score 124O" (plate 9), "Onyxia's Lair 2O:OO" (plate 11). Players will read these as typos, and the client's own tooltip face has plain lining numbers. Fix: draw tooltips in a face with lining, tabular figures (Friz Quadrata look-alike, or the current face with lining-nums switched on) so 0 is a zero everywhere; check "Sell Price: 7 63" and "10 75" gold lines too.

2. SHOULD. Plate 7, "Thornwake Company" panel: "3rd of all guilds · 41 raid nights in 90 days". Forever launched 4 Nov and the first raid phase opens 9 Dec, so on 17 Dec a guild cannot have 41 raid nights, let alone in a 90-day window. Same problem on plate 9: Orlaith's "Forever Score 1563 (6 Dec)" is an Epic-tier score dated three days before any raid exists, and plate 10 has 22 fights and an 8-boss Molten Core by 15 Dec. Fix: change to something true for a week-old raid phase ("7 raid nights since Phase 1"), date Orlaith's record 13 Dec or later, or move the whole fixture clock to a date where the numbers are possible.

3. SHOULD. Plate 4, "To the site" box: the code reads "FS1:70938:hunter:60:Thornmaw:Thornwake Co / mpany:Forever:..." but the character is Velisandre. The name in the code must be Velisandre. The line also breaks in the middle of "Company". Fix: put the real name in the code and cut the string with an ellipsis before it wraps, or break after a colon.

4. SHOULD. Plate 3 versus plates 4 to 11. The talent window shows "Unspent Talents 1" and "22 of 24", which is a level 33 hunter, while the same hunter crest appears on plates where Velisandre is level 60 with "51 of 51 on plan" and "Every point is placed". Nothing on the plate or in the caption says it is a different character. Fix: add to the caption "Shown for a level 33 hunter", or show a level-60 state.

5. SHOULD. Plate 1, right panel, roll frames. The three buttons are a square-in-circle, a ring and a cross. Need is a die and Greed is a coin in the client, Pass is the grey cross; the frame also shows "Bind on Pickup" under the item name. A player cannot tell which button is Need. Fix: use the client's own Need dice, Greed coin and Pass art in the same positions, and add the Bind line. The right half below the two frames is empty, so a third roll row (a green-arrow one for Velisandre with a plain "Sidegrade") would fill it and show the caption's other states.

6. SHOULD. Plate 2, tooltip line "Equip: Improves your chance to hit by 0.3%." Trueaim Gauntlets carry 3% hit in Classic (30 rating at 10 per 1%), and the round 3 review noted the data holds 3%. 0.3% reads as a tenfold error to any hunter and nothing else on the tooltip (no primary stats on a level 54 rare) tells them otherwise. Accuracy reviewer to confirm the client string; if the client truly prints 0.3%, the plate is right but a note in the caption is worth it.

7. NICE. Plate 11. The caption was re-read after it changed to "The Forever Sixty line" and now matches the plate (the chip with its tooltip). The old "By the minimap" problem is gone; a minimap circle above the chip would still show where it sits.

8. SHOULD. Plate 8, roster table. Column "Ench." with "0" for most raiders reads as zero enchants when it means zero missing; "Gap" has no unit in the header (the unit is only in the footer, "Gap in DPS"); "Not signed" in the Sign-up column is cut from "Not signed up" used on plate 7. Fix: headers "No ench." and "Gap DPS", widen the Sign-up column.

9. SHOULD. Plate 6, Upgrades panel. The caption says "The top upgrades say where they drop". Marksman's Girdle does (Blackrock Spire · Urok Doomhowl), but "Cutthroat's Signet, Finger · Pix Xizzix" gives a name with no zone and no word for drop, quest or vendor, "Legionite Glaive, Main Hand · Crafted · Blacksmithing" has no skill level, and Gloves of Holy Might only says "In your bags". Fix: add the zone or the kind after each name, and the required skill on the crafted row.

10. NICE. Plate 6, filter chip "Bare" beside "Raid-ready". A veteran does not know what Bare means (no consumables? no enchants?). Use the word the condition is.

11. NICE. Plate 4, "Up next" panel: "Beast Mastery raid build" over "Every point is placed." The heading promises a next step and the body says there is none. Retitle the empty state ("Nothing to take") or show the next build.

12. NICE. Plate 8, readiness line "Not covered: Faerie Fire" while Saffa, an orange (druid) name, is signed In. Either Saffa is not a druid or the line is wrong; make it Curse of Recklessness or another real gap.

13. NICE. Plate 2, "Character" panel shows three loose slots and the bag has no "Backpack" title, no money and no close button. Fine as a crop, but a veteran will see it is not a real frame. A cropped corner of the real paperdoll column would read truer.

14. NICE. Plate 5 sorts bosses by followed percent (74 to 95) while plate 10 lists the same raid in kill order. Both are fine; a short "worst first" tag on plate 5 would stop a double-take.

## Data I do not believe (for the accuracy reviewer)

- Plate 5 and 6 weights: hit rating 0.42 DPS per rating and crit rating 0.18 per rating on a 941 DPS sim. At 14 rating per 1%, 1% crit is +2.5 DPS, about 0.27% of 941; ranged crit adds close to 1% of damage per 1% crit, so the crit number looks 3 to 4 times low, and hit looks about half of what a ranged hunter under the cap loses. Check the sim's DPS base (pet and Bestial Wrath inside it?) against the per-point figures.
- Plate 10, "941.2 DPS" for a 61 item level hunter in Molten Core gear is high for Classic; fine if the sim assumes full raid buffs, but the plate does not say so.
- Plate 5, "Aimed Shot: On cooldown, shared with Multi-Shot": Classic's Aimed Shot and Multi-Shot do not share a cooldown, check the Forever spell category before this ships.
- Plate 2 and 6, Trueaim hit value (finding 6).

## Verdict

FIX. The plates are consistent with each other to a rare degree and the class and gear rules hold, but the tooltip numerals (finding 1) will be mocked on sight, and findings 2 to 4 are visible contradictions a careful player will catch. Fix 1 to 9 and re-render; 10 to 14 can follow.

## Round 2

VERDICT: SHIP

Seen: all 11 plates at full size in the new order, and the whole curseforge-description.md (table, captions, Description). Checked against api/internal/guilds/readiness_core.go (enchantableSlots is main_hand, chest, back, feet: the description's "chest, back, boots and weapon" is exact), api/internal/coach/f8_adherence.go (F8 runs for DPS and tanks only, healers get nothing: "DPS and tanks" is exact), and api/internal/discordbot (In, Late, Bench, Out buttons exist: the Discord claim holds).

### Round 1 findings, re-checked

- 1 (O for zero): fixed. "Onyxia's Lair 20:00", "Forever Score 1240", "Durability 40 / 40", "Requires Level 54", "Best for 20 to 29" all read as plain digits.
- 2 (dates): fixed. "3 raid nights in 90 days" is possible on 17 Dec, Orlaith is dated 13 Dec and greyed (four days, over the three-day rule), MC on Tue 15 Dec is fine.
- 3 (export code): fixed. "FS1:70938:hunter:60:Velisandre:Thornwake Company:Forever:1,2:..." wraps at the space, not mid-word.
- 4 (level 33 caption), 7 (plate 11 title), 9 (Gear caption wording): fixed and each caption matches its plate. Plate order in the caption list matches the files 1 to 11.
- 5 (roll buttons): still a square-in-circle, a ring and a cross. I hold it as SHOULD, not a blocker: a player can still read the frame as a loot roll. I still cannot tell which is Need.
- 6 (Trueaim 0.3%): not blocking as long as the client prints the string; see finding 2 below.
- 8 (roster headers): accepted as the addon's own words. "Not signed" cut from "Not signed up" stays a NICE.
- Numbers re-checked and still add up: parts 410+380+228+120+46 = 1184 and 410+352+268+150+60 = 1240; sign-ups 11+1+2+1+16 = 31 and role lines 2+4+2+4 = 12.

### Findings

1. SHOULD. Name and plate 5 against plate 4 and the Gear panel. The title says "Sim and BiS" and caption 5 says "Best in slot", but plate 5's Sources shows "Raid 0" ticked, and the same hunter is shown on 17 Dec, three weeks into a raid he is clearing (ilvl 61 on the roster, Molten Core parses on plate 7). Plate 4 then offers Striker's Mark (+10.1 DPS, an epic bow, item level 69-71) while plate 5 says the Eaglehorn Long Bow is "Best in slot". A raider reads "BiS" as the raid list; the Gear tab is a pre-raid list. Fix: say "pre-raid best in slot" in caption 5 and in the Gear section of the description, or drop "BiS" from the store name; or make the fixture honest (raid items counted once raids are open, or the roll item not a raid drop).
2. SHOULD. Plates 3 and 5: Trueaim Gauntlets "Improves your chance to hit by 0.3%." A hunter will see a tenfold error. Still the open data question you are taking to the owner. Until answered, swap the plate's tooltip to an item whose stat nobody will question (the same plate works with Gloves of Holy Might), or keep it and accept the mockery. Not a blocker for a gallery.
3. SHOULD. The pitch is raid-first, but the first bullet is "Upgrades show the DPS you gain" and the heading is "Raid night for Forever guilds." Put the sign-ups bullet first. Also say in the first screen that the Guild and Officer tabs and the sign-up buttons need the companion app (Windows and Mac) and a guild on the site; today that is only in the companion section far down. A player who installs from the top and opens an empty Guild tab will call the store page misleading. "From Discord, the site, the addon or the phone app (in beta)" is true, but the addon route goes through the companion, which the sentence does not say.
4. SHOULD. Plate 4 shows a roll line "+10.1 DPS" while the description quotes "+5.8 DPS · Best at level 60" and "Sells for most". Either show the "Best at level 60" tail on the plate or change the description example to match what the plate shows.
5. SHOULD. Plate 5 Upgrades panel: "Cutthroat's Signet, Finger · Pix Xizzix" still gives a name, not a zone or a verb (drop, quest, vendor). Marksman's Girdle ("Blackrock Spire · Urok Doomhowl") is the model; use it on every row, and give "Crafted · Blacksmithing" its skill level. Caption 5 promises "say where they come from".
6. SHOULD. Plate 2 readiness: "Not covered: Faerie Fire" with Saffa (orange, a druid) signed In. Carried from round 1 (finding 12); pick a real gap such as Curse of Recklessness.
7. SHOULD. Categories: seven are listed (Guild, Boss Encounters, Class, Quests & Leveling, Tooltip, Bags & Inventory, Combat). CurseForge caps how many a project can carry; confirm the dashboard accepts all seven before the upload step, and keep Guild primary.
8. NICE. The enchant check covers four slots. A veteran raider also enchants bracers, gloves, shield and a two-hander. The description says "chest, back, boots and weapon", which is true; do not let a caption say "missing enchants" with no list.
9. NICE. Plate 5 "Bare" is now explained in the description ("no raid buffs or consumables"), but the chip on its own still needs the word "consumables" or a tooltip.
10. NICE. "What's new in 2.0.55" belongs in the store changelog, not the description; it will be stale at the next release. The Settings and Commands tables are reference length for a store page; keep them below the fold as they are.
11. NICE. Plate 5 level chips: no band is shown as selected, so the "60" the list is for is not visible.

### Data I do not believe (for the accuracy reviewer)

- Weights on plates 5 and 6: crit rating 0.18 DPS per rating and hit rating 0.42, on a 941 DPS sim. At 14 rating per 1% crit that is 2.5 DPS per 1%, about 0.27% of damage; the crit figure looks low for a ranged hunter. Still open from round 1.
- "Aimed Shot: On cooldown, shared with Multi-Shot": confirm the Forever spell category; Classic's two shots do not share a cooldown. Still open.
- "941.2 DPS" at item level 61 assumes full raid buffs; the plate does not say so.
- Striker's Mark as a roll from a Molten Core boss: confirm the source and that its level-60 listing is correct (two ids, 228252 and 228519, share the name at item level 71).

### Verdict

SHIP. The tooltip numerals, the calendar and the export code are fixed and every number I can add up adds up. The store text's claims match the code: four enchant slots, DPS and tanks for the cast check, four sign-up answers from Discord, site, addon, app. Fix 1 and 3 before the paste (the BiS wording and the companion note); the rest can follow.
