# F11 build review, round 1: player verdict

Reviewer: veteran Classic player. Captures in `build/`, iOS = iPhone 17, Android = Pixel 9. Board comparison and the fixtures were read; the apps were not run.

## Verdict: HOLD

The rankings page answers "where do I stand" first, and the parse colours are right (93 purple, 99 pink, 204th grey). The character page and the Paperdoll are the right idea. Three things would be screenshotted and mocked in guild chat (below), so this does not ship yet.

## MUST

1. **Rank 1 and rank 14 wear the same 17 items.** `character-other-ios.png`, `character-gear-ios.png`, `character-gear-and.png` (Voltaris, 462 DPS, 99th percentile) against `character-yours-ios.png` (Lightningbolltz, 398 DPS). Both show Blue Dragonscale Helm, Archlight Talisman, Ironfeather Shoulders, Hand of Edward the Odd, and so on, slot for slot, with the same gold dots. Yet Voltaris has Gear 412 and Lightningbolltz 371. A raider who inspects two players to see what the better one wears sees nothing. Cause: `character.card.ranked.json` and `character.card.yours.json` have byte-identical `gear` arrays. Give Voltaris his own, better set (the 99th percentile should wear the picks, with few gold dots) and keep the 14th-place set for Lightningbolltz.
2. **The Home Rankings row names the wrong class.** `home-rankings-ios.png` (see also `home-rankings-and.png`): the character is "Orc Beast Mastery Hunter" Velaren, but the row reads "14th of 212 Elemental Shaman characters on Onyxia". A hunter never sees a shaman ladder. The fixture is mixed (Lightningbolltz's rank on Velaren's Home). Fix the data. Velaren's row should read a hunter ladder, or Home should be shown for the shaman.
3. **"Today only" switch has no switch.** `rankings-filters-ios.png`, `filters-and.png`: the control draws as a plain grey blob with no knob and no on/off. A player cannot tell whether it is on. The filter chips also show stray vertical tick marks at their left and right edges (Measure, Ruleset, Region, Phase, Faction rows, iOS). It looks broken. Draw the real toggle and clean the chip borders.
4. **Raw debug text in the off-hand slot.** `home-gear-full-ios.png` (and Android Home): the off-hand reads "none: Legionite Glaive is two-handed." Lowercase, with a colon, in the item-name position. Say what the client says: an empty slot labelled "Off hand", with the line "Two-hand in main hand" (the Paperdoll spec's own wording).
5. **The on-page rank card does not name the character.** `rankings-ios.png`, `rankings-onpage-and.png`: "Your best kill on Onyxia, Normal, Phase 1." The spec (One character, one number) says the card names the character in both states. The off-page card does (`rankings-offpage-ios.png`: "Sparkwick · best kill on Onyxia, Normal, Phase 1."). A player with an alt has to know which toon this is. Add "Lightningbolltz" in the same gold.
6. **The big ordinal is purple on iOS and white on Android.** `rankings-ios.png` "14th" is parse purple; `rankings-onpage-and.png` "14th" is white. The off-page "204th" is grey on iOS. One platform is wrong. The parse colour on the big ordinal is the better reading (the player sees their tier). Colour both and check `rankings-offpage-and.png` too.

## SHOULD

1. **"Gear gap 22.7 points" has no unit a player understands** (`character-gear-ios.png`, `character-gear-and.png`). Home says "+21.2 DPS" per upgrade; here the same idea is "points". Of what? Say "Gear gap 22.7 score points" or show it as DPS if the model supports it. The gold-dot slots also show no per-slot gain, so the 6.2 / 9.4 / 4.1 / 3.0 behind the sum are invisible until a tap. Say what a tap does.
2. **Dates disagree between list and page.** Voltaris's row on `rankings-ios.png` is "2 d ago", but his page (`character-other-ios.png`) says the same kill was "1 d ago". The fixture has the list at 15 Dec and the card at 16 Dec. The kill time is 3:13 in the list and 192 s (3:12) on the card. Same fight, one number.
3. **"2 d ago" is spaced and wraps badly.** `rankings-onpage-and.png`: "2 d / ago" breaks across lines on rows 1, 3 and 4. Write "2d ago" and keep it in one piece. In `rankings-large-ios.png` the sub-line also breaks mid-phrase ("· / 6 d ago"). The kill time "3:13" has no unit; the client and WCL show it as "3:13" under a "Time" label, so say "kill 3:13".
4. **Unknown items do not follow the owner's ruling.** `home-gear-full-ios.png`: Wrist reads "Not in our item data yet" with "Wrist" underneath. The ruling is "<Slot> · Item <id>". Match it, in the same line style as the other rows.
5. **Home fixture is on an old build.** `character.home.json` says data_build 1.60.1.70658; the active build is 1.60.1.70707. Regenerate so the Home captures match the other fixtures.
6. **Home "Your gear" intro** (`home-gear-full-ios.png`): "What you wear, and the pick for 60 · raid. Tap an item for its tooltip." The middle dot reads like a typo. Use "...the pick for level 60 raid gear."
7. **Paperdoll layout.** In `character-gear-ios.png` the right-hand trinkets and the left-hand empty column leave a hole under Wrist; trinkets sit in the right column only. That is the web Paperdoll's own layout, so fine, but at the largest text (`character-gear-large-ios.png`) it falls to one long list and the "Gold dot" legend and Gear gap sit far from the dots. Keep the legend under the grid at every size (it does).
8. **Android has no star or share on the Rankings page** while the character page has both. Fine. But iOS rankings and character pages have no tab bar while Android keeps one (`rankings-ios.png` against `rankings-onpage-and.png`). Pick one behaviour for pushed pages.

## Missing against Wowhead / WCL

- The rankings row has no spec icon, no item level and no talent split. The fixture carries `ilvl` (62, 61, 60) and `talent_split` (30/0/21) for each row. A player wanting "what does that player wear" sees only a name and guild; show the ilvl per row, which is exactly the number we have.
- No rank-1 comparison on the Your Rank card ("48 DPS behind 1st") for a player trying to improve; this is the gap that moves the player to a better character.
- No icon on the encounter or the class on the filter line; the word "Damage" is fine.

## Data I do not believe (for the accuracy reviewer)

- `character.card.*.json` Elemental Shaman picks: **Ironfeather Breastplate** (Leatherworking, replacing Vest of Elements ilvl 63), **Skycaller's Leather Waistguard** (replacing Cord of Elements) and **Uther's Strength** (World drop) for a caster. Uther's Strength is a healing-leaning trinket; check it against Briarwood Reed under the Elemental weights. A leather pick for a shaman is legal at level 60 (mail from 40), but check it beats the mail it replaces.
- The worn set is mostly 51 to 62 ilvl at the 99th percentile on Onyxia in week two, with a ranged slot Totem of Thunder at ilvl 65, higher than every other item. Check that the totem is launch-legal and that 65 is the client's number.
- Emperor Dagran Thaurissan is a five-man boss ranked at "Normal" with a percentile, next to a 40-man raid boss. Confirm the encounter is meant to rank.
- The Home gear for a level-60 hunter: Shadowcraft Spaulders and Windrunner Legguards (leather) are legal as worn, but Legionite Glaive as a hunter weapon pick should be checked against the BiS list.
- Filter regions include KR, TW and CN; confirm Forever has those.

## What is right

- `rankings-offpage-ios.png`: the pinned "YOU" row with the gold edge and "204" in grey, above the list, answers "where am I" without scrolling. Keep it.
- Parse colours follow WCL (99 pink, 93 purple, 81 purple, 204th grey).
- Quality borders and name colours on the Paperdoll read like the character frame; units (DPS, ilvl, percentile) are present on every figure that has one.

---

# Round 2

## Verdict: HOLD (one MUST left, small)

Five of the six round-1 MUSTs are fixed. One new contradiction on Home would still be screenshotted.

Checked and fixed:
- **Gear sets.** `character-gear-ios.png`: Voltaris now wears his own set (Living Crown, Ironfeather Breastplate, Skycaller's Leather Waistguard, one gold dot on Trinket 2, gap 3.0). `character-yours-ios.png`: Lightningbolltz is in Coif of Elements with gold dots. Two different players, two different sets.
- **Filters.** `rankings-filters-ios.png`: a real switch with a knob, clean chip borders.
- **Off hand.** `home-gear-full-ios.png`: an empty "Off hand" socket reading "Two-hand in main hand". The unknown wrist reads "Wrist · Item 999999".
- **Rank card.** `rankings-onpage-and.png`: "Lightningbolltz · best kill on Onyxia, Normal, Phase 1." in gold, and "14th" in parse purple, matching iOS.
- **Ages and times.** Ages no longer break in the middle ("2 d ago" stays together). Voltaris reads 3:13 and "2 d ago" in the list and "Normal · 2 d ago" on his page.

### MUST

1. **Home iOS still shows another character's upgrades.** `home-rankings-ios.png`: the header says "Troll Elemental Shaman · Level 60 · Hordebreakers" and the next section is "NEXT FOR LIGHTNINGBOLLTZ" with "14th of 212 Elemental Shaman characters on Onyxia". But the upgrades card above it lists Legionite Glaive, "Hyper Deluxe Sniper Rifle Mk XVII", "Beaststalker's Pants" and "Beastmaster Boots", and the screen title is still "Velaren". A shaman with a sniper-rifle upgrade is exactly what a player would mock. Your note says this capture has no upgrades. This one does, and it contradicts `home-rankings-and.png` ("Every slot matches the 60 list."). Recapture iOS with the same data as Android, with the title "Lightningbolltz".
2. **Android Home contradicts the character page.** `home-rankings-and.png` says "Every slot matches the 60 list." and the Talents row says "DIFFERS", while `character-yours-ios.png` shows Lightningbolltz with a gold dot on Head (Coif of Elements) and Waist, and a gear gap of 36.0 (your note). Either Home shows his upgrades or the character page loses the dots. This file is also older (11:33) than the iOS captures (12:32); regenerate it with the fresh build.

### SHOULD (carried over, not changed)

- "Gear gap 3.0 points": still no unit. A player does not know what a point is.
- Rows still show no ilvl or talent split, though the data carries both.
- The "pick for 60 · raid" intro line is still on Home.
- Your note calls Voltaris's set a "nearly complete raid set". It reads as dungeon and crafted gear: Band of the Unicorn (ilvl 48), Rune of the Guard Captain (green), Totem of Thunder (ilvl 65). That is fine for a launch-week Onyxia kill, but don't promise a raid set in the copy.
- The data items for the accuracy reviewer (Uther's Strength, leather picks, ilvl-65 totem, Emperor Dagran Thaurissan, KR/TW/CN) still stand.
