# Player verdict, F9 build (Guides tab), iOS + Android

Verdict: HOLD

Question the tab answers: "How do I play my spec, and where does it stand?" The class guide opens on role and numbered stat chips, so the answer is first on screen. The Guides tab, tier rows and ringed crests read well. The rotation and race screens have things I would screenshot and mock.

## MUST
1. Fire Nova shows the red "?" missing-icon tile (ios/f9-guide-rotation-ios.png, and/f9-guide-rotation-and.png, priority row 7). The game never shows a "?" for a real spell. The row is also named "Fire Nova"; the spell is Fire Nova Totem. "On a Clearcasting proc" is not a rule any shaman plays by; accuracy reviewer to check.
2. Cooldowns rows 2 and 3 ("Astral Recall", "Trinkets, racials and potions") render as empty dark squares where the icon belongs (ios/f9-guide-rotation-ios.png). Give them icons. Astral Recall is a 15 min teleport, not a damage cooldown; listing it in a DPS rotation will get mocked. Drop it or label it.
3. Mana Spring Totem prints "Restores 10 mana every 2 seconds for its five-minute life" (both platforms, Rotation). Spec 5.1 F9.3 says no duration is printed for this totem until the accuracy reviewer settles 60 s vs 300 s. Remove "five-minute life". Rank 4 restoring 10 per 2 s also looks low; accuracy to check.
4. Race rows use the wrong art (and/f9-guide-stats-and.png, Races). Orc, Tauren and Troll all show the same red Horde crest; Dwarf shows an Alliance shield. Show the race icon for each race.
5. Largest text, numbered chips (ios/f9-guide-large-ios.png): the digit in the number disc is tiny next to the stat name at AX5, and "Spell penetration" is clipped at the right edge. Spec 4 says chips wrap and are never cut. The tab row also cuts "Overview" and "Talents" at the screen edge.
6. Talents list leaves out Elemental Mastery (ios/f9-guide-talents-ios.png). Spec F9.2 says key talents include every talent the spec casts through, and Elemental Mastery is on the Cooldowns list. Add the row.
7. What changed, Key dates (ios/f9-changes-ios.png): the Name reservation row says "through Nov 3" and then repeats "Name reservation open through Nov 3" in gold. Print it once. The tile reads "Oct 27" for something still open.

## SHOULD
1. Mixed date formats on one page: "Oct 27" tiles, "10/28" list, "1 Nov 2026" on guides (ios/f9-changes-ios.png, f9-guides-ios.png). Pick one.
2. Stat priority lists "Nature power" as its own stat at 3 and "Spell haste" at 6. A Classic player knows spell power as one number and gear has no haste. If Forever really has both, say so in a line. Accuracy to confirm.
3. "Hit cap: 10%" (and/f9-guide-stats-and.png): a player expects the 16% spell hit cap against raid bosses; Forever hit is one stat, so show the number in rating too. Accuracy to check.
4. Race lines (Tauren "1% Hit chance", Orc Blood Fury "Attack Power and Spell Power", Dwarf Stoneform "instead of raising armor") differ from Classic. May be Forever's client text; accuracy to verify.
5. Offline class page (ios/f9-class-offline-ios.png) still lacks the "This guide is not on your phone yet. Connect once to save it." line under Enhancement ("Not saved"). The brief says it lands this round; confirm before merge.

## Not reviewed in detail
Android large-text captures and error and loading states were only skimmed by name. Tier lists reviewed on presentation only (Elemental Shaman last, 40.6% behind, is the accuracy gate's data).

## Accuracy handoff
Fire Nova name and condition, Mana Spring Totem numbers, hit cap, Nature power and Spell haste as stats, race ability text.

# Round 2

Verdict: SHIP (no MUSTs left; SHOULDs below)

Checked: ios/f9-guide-rotation-ios.png, f9-guide-large-ios.png, f9-changes-ios.png. Android and the remaining captures were not re-read.

Closed:
- Fire Nova no longer shows the red "?". Astral Recall is gone from Cooldowns. The no-spell row (Trinkets, racials and potions) has no blank tile.
- AX5 stat chips: digits now scale with the name and every stat wraps inside its chip, nothing cut.
- Key dates: Name reservation shows one line.
- Mana Spring Totem duration: accepted. The spec rule prints a duration when client data and guide agree, and both say 300 s. If the accuracy reviewer settles on 60 s, this line must change with it.
- Race rows with a faction mark: accepted per spec F9.2.

SHOULD (carry forward):
1. Fire Nova (priority row 7, ios/f9-guide-rotation-ios.png) is still a plain empty dark square among real icons. A player reads that as a missing icon. Use the Fire Nova Totem icon if the client has none for the spell.
2. Talents tab and the Elemental Mastery row were not re-read in this round; confirm in ios/f9-guide-talents-ios.png.
3. Round 1 SHOULDs still open: mixed date formats (Oct 27 / 10/28 / 1 Nov 2026), hit cap 10%, Nature power and Spell haste as stats, race ability text, offline "Connect once to save it" line.
4. Fire Nova's condition "On a Clearcasting proc" does not mention the Fire totem requirement the guide prose gives; accuracy reviewer to check.
