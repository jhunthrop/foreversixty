# Guide accuracy round 3: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree land/guides-r2 7a2165cb0, client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md` and the card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`.

## Builds and shipped leveling order

Build strings unchanged since round 1 (all five legal on 70755; re-decoded against `talents/shaman.json`, tier gates and prerequisites re-checked). Shipped orders re-decoded from `data/builds/1.60.1.70755/addon-builds.json` (index 0 = level 10); identical to round 2's list:

- Elemental leveling: Convection 10-14, Concussion 15-19, Elemental Devastation 20, Elemental Alacrity 21-23, Call of Flame 24-26, Elemental Focus 27, Elemental Warding 28-30, Call of Thunder 31, Lightning Overload 32-34, Elemental Fury 35-39, Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60. elemental.md:52 matches.
- Enhancement leveling: Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-25, Shamanistic Focus 26, Flurry 27-31, Stormstrike 32, Improved Ghost Wolf 33-34, Mental Quickness 35-36, Maelstrom Weapon 37-39 and 41-42, Rage of the Farseer 40, Convection 43-45, Concussion 46-47, Elemental Alacrity 48-50, Call of Flame 51-53, Tidal Focus 54-58, Mindfulness 59-60. enhancement.md:113 (round 2 fix) matches.
- Restoration leveling: Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide Totem 25, Ancestral Healing 26-27, Healing Focus 28-30, Tidal Mastery 31-35, Restorative Totems 36-40, Healing Way 41-43, Nature's Swiftness 44, Purification 45-49, Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60. restoration.md:117 (round 2 fix) matches.

## What checked out (no finding)

- Every talent value in the prose against the 70755 max-rank text: Shamanistic Focus 45% (shocks and Lightning Shield), Stormstrike 150% and +20% on the next Lightning Bolt, Chain Lightning or Earth Shock for 12 s, Improved Ghost Wolf 3.0 s and usable indoors, Elemental Weapons "your Windfury Weapon effect by 40%", Mental Quickness 30% of Intellect, Maelstrom Weapon 20% a stack, five stacks, Elemental Devastation 9% melee crit for 10 s after a direct-damage spell crit, Mana Tide 12 s (13000 ms duration, `${13 sec-1}`), Healing Way 25% flat, Riptide 445 over 15 s plus 25% Chain Heal on its target, Call of Flame +9 s and 15%, Elemental Fury +100% on Searing, Magma and Fire/Frost/Nature spells, Lightning Mastery 0.5 s and no Chain Lightning cooldown for 15 s (`spellconst` 16166, 180 s).
- Trainer levels (`trainables/shaman.json`): Call of the Elements 20, Totemic Call 22 (30 s, 25% base mana), Ghost Wolf 20 (3 s, aura 31 at 40%), Chain Lightning 32, Chain Heal 40, Lesser Healing Wave 20, Healing Stream 20, Strength of Earth 10, Searing 10, Flametongue 10, Rockbiter 1, Windfury Weapon 30, Lightning Shield 8, Cure Poison 16, Cure Disease 22.
- Windfury Totem "nothing to twist" (index :27, enhancement :79): the totem's pulse auras 8516/10608/10610 last 1000 ms, so the bonus ends with the totem. Right.
- Riptide "cheapest heal per mana" (restoration :74, card 11): rank 3 805 over 15 s for 230 mana (3.5 base a mana) against Healing Wave rank 10 1615 x 1.25 for 620 (3.3). Right.
- Chain Heal rank 2 vs 3 crossover (restoration :83): same 0.714 coefficient; rank 2 is better per mana above about 290 healing. Right.
- Stat priorities against the band 60 raid scale factors in `bis/shaman-*.json`: Elemental spell power 2.00 > hit 1.37 > Nature power 1.29 > Intellect 1.00 > crit 0.74; Enhancement hit 2.51 > Intellect 1.54 > crit 1.34 > Agility 1.06 > Strength 1.00 > AP 0.45 > spell power 0.40; Restoration mp5 1.69 > Intellect 1.00 > healing 0.81 > Spirit 0.60 > crit 0.40. All three frontmatter lists match.
- BiS and profession prose: Elemental raid set wears Leatherworking (Blue Dragonscale Helm, Gloves of the Greatfather, Ironfeather Leggings, equip; Skycaller's Leather Waistguard, pickup) and Tailoring (Mantle of the Timbermaw, Runecloth Cuffs, equip) crafts; Enhancement set as round 2 said; Restoration Devout pieces from Stratholme and Blackrock Spire, Living Crown and Mooncloth Boots bind on equip (`items/shaman.json`).
- Consumables: Elemental flask plus Greater Arcane Elixir, Brilliant Wizard Oil; Enhancement Mongoose, Juju Power, Juju Might, R.O.I.D.S., Smoked Desert Dumpling; Restoration Brilliant Mana Oil (`sim/request/raid-consumes.json`).
- Racial text against `races.json`: Blood Fury 10% AP/RAP/SP 15 s / 2 min, Berserking 10% casting and attack speed 15 s / 3 min, Stoneform 10% physical 8 s / 3 min, Artisan's Precision 1% hit, Endurance 5% health, Wind Blessed 2% haste; Windshaper Skyborne is Horde-only for shamans.
- Round 2 fixes re-checked: R2-1 to R2-8 and D1 all still hold against the data. The Chain Lightning card note ("the jumps add most of another bolt") is right on two targets (second hit 0.7 of the first, about 84% of a bolt at 600 Nature power).
- No em or en dashes, no site/sim talk, no quoted nightly numbers in prose.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R3-1 | CONFIRMED | consumable / rotation | elemental.md:74; card `data/curated/apl/shaman-elemental.json` priorityList[1] note | "Press it on the pull with your trinkets, Blood Fury or Berserking and your potion"; card: "Lightning Mastery, trinkets, racials and potions on cooldown" | The only potion the page names is Major Mana Potion, and elemental.md:96 says to drink it "as soon as each fits under your mana cap". At full mana on the pull it does not fit. The preset (`raid-consumes.json` shaman-elemental from 49) carries `major_mana_potion` and no damage potion. The page contradicts itself. | elemental.md:74: "Press it on the pull with your trinkets and Blood Fury or Berserking, and again every 3 minutes." Card note: "Lightning Mastery, trinkets and racials together: for its 15 seconds Lightning Bolt is half a second faster and Chain Lightning has no cooldown, so stack everything inside it." (Major Spellblasting Potion on the pull is a separate question, see Sims requested 3.) |
| SHAMAN-R3-2 | CONFIRMED | race (owner rule) | elemental.md:102, enhancement.md:105, restoration.md:109, index.md:36 | "Windshaper Skyborne's Wind Blessed is 2% haste that is always on" (and the same line on the other two spec pages); index: "Orc, Tauren and Troll shamans are joined by the Windshaper Skyborne" | Owner rule (round 3 brief): paid races are never the recommended pick and are named as the paid option only. Skyborne needs the Skyborne Heroic Pack (`races.json` 96, Blizzard pre-purchase page). None of the four shaman lines says so; warrior/index.md:39 and warrior/fury.md:115 do. | Spec pages: "Windshaper Skyborne, the paid Heroic Pack race, brings Wind Blessed's 2% haste, always on." Index: "On the Horde, Orc, Tauren and Troll shamans are joined by the Windshaper Skyborne, a paid unlock." |
| SHAMAN-R3-3 | LIKELY | leveling | restoration.md:60, :117 | Mana Tide Totem is described only as the talent ("comes at 25"); the page never says the talent teaches rank 1. | `trainables/shaman.json` Mana Tide Totem ranks 16190 (talent), 17354 at 38, 17359 at 48, 1325256 at 58. `spellconst`: rank 1 restores 88 mana every 3 s, rank 4 290. Round 2 added the same line for Riptide (R2-8); the card note already says "Rank 4 at 58", the prose does not. A player who never trains it gets under a third of the mana. | restoration.md:117, after "Mana Tide Totem comes at 25, the first level the fourth row opens.": "The talent teaches its first rank; train the others at 38, 48 and 58, each one a big step up." |

## Decisions

- **D1. Horde race for Elemental and Enhancement: held for the sims below.** Racials are reworked in Forever, so no Classic source settles it (the only web answers are Classic/TBC era: Orc for Enhancement, Troll for Elemental, both on pre-rework racials). Alliance has only the Dwarf, so the Alliance line stands. Apply when the numbers land:
  - If Orc leads the free Horde races beyond error (expected: two Blood Furies in a three minute fight against one Berserking), keep `recommendedRaces: [dwarf, orc]` and the Orc lines as written.
  - If Troll leads, Elemental: `recommendedRaces: [dwarf, troll]` and elemental.md:102 "On the Horde, Troll is the damage pick: Berserking is 10% casting speed for 15 seconds every 3 minutes, pressed with Lightning Mastery. Orc's Blood Fury, 10% spell power for 15 seconds every 2 minutes, is close behind." Enhancement the same shape with attack speed and Rage of the Farseer.
  - If Windshaper Skyborne leads, it is still not the pick (owner rule); name it as the paid option with SHAMAN-R3-2's wording and recommend the best free race.
  - Tauren is not a damage pick in any outcome (War Stomp, Cultivation, Plainsrunning, Endurance).
- **D2. Restoration race: no sim.** The heal sim's shaman does not run out of mana under the site's heal profile (APL notes), so healing is bound by health thresholds and Berserking's casting speed barely moves it; Tauren's 5% health is the survival pick a healer wants. `[dwarf, tauren]` stands; Skyborne line per SHAMAN-R3-2.

## Sims requested

All on band 60, raid preset, sim/cmd/spec-breakdown, 10,000 iterations, seeds 1 and 50001, as round 2's race-sims.md. Pass the guide raid builds with `-talents`, because `bis/shaman-enhancement.json`'s raid talent string is still in the 70753 tree order (round 1 S2).

1. **shaman-elemental race**, `-talents` from raidBuild `FS1:1.60.1.70755:shaman:dwarf:5510331300001351/055/053002:` (engine form `5510331300001351-055000000000000000-0530020000000000`): Horde `orc`, `troll`, `tauren`, `windshaper-skyborne`; Alliance `dwarf` for reference.
2. **shaman-enhancement race**, `-talents` from raidBuild `FS1:1.60.1.70755:shaman:dwarf:0533001/053030310051202051/053:` (engine form `0533001000000000-053030310051202051-0530000000000000`): Horde `orc`, `troll`, `tauren`, `windshaper-skyborne`; Alliance `dwarf` for reference.
3. **shaman-elemental potion** (optional, decides whether SHAMAN-R3-1's pull line gets a potion back): same build and race (dwarf), preset as published against the preset with `major_spellblasting_potion` used on the pull in place of the first Major Mana Potion. If the Spellblasting opener wins beyond error, the fix becomes "Press it on the pull with your trinkets, Blood Fury or Berserking and a Major Spellblasting Potion; a Major Mana Potion and a Demonic Rune come out later in a long fight", and elemental.md:96 names it.

## Sim findings

- None new. Carried: Stormstrike `RequiredLevel: 40` versus talent-taught at 25 (round 1 Q1), Restoration card ranks, Grace of Air and Strength of Earth book ranks (R2 SIM-R2-2), Lightning Overload ratio check on the first level-60 log (R2 D1), Totemic Call positioning (R2 SIM-R2-1).
- Data lane note (not a guide finding): `bis/shaman-enhancement.json` raid band has no ranged (totem) slot filled, and `bis/shaman-elemental.json` raid band leaves the off hand empty beside a one-hand mace. BiS refresh list.

## Sources

- Blizzard, "Pre-Purchase World of Warcraft: Forever Upgrades" (Skyborne Heroic Pack) https://news.blizzard.com/en-us/article/24301508/pre-purchase-world-of-warcraft-forever-upgrades-and-begin-your-next-journey-in-azeroth
- Wowhead, "Skyborne First Look" (shaman only on the Horde for Skyborne) https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Icy Veins, "The racial rework gives every race a reason to matter in Warcraft Forever" https://www.icy-veins.com/wow-forever/news/the-racial-rework-gives-every-race-a-reason-to-matter-in-warcraft-forever/
- Icy Veins, "Classic Enhancement Shaman DPS Guide" (Classic-era Orc recommendation, pre-rework racials, so not used to decide) https://www.icy-veins.com/wow-classic/enhancement-shaman-dps-pve-guide

## Fix log (round 3)

- SHAMAN-R3-1 fixed: elemental.md:74 drops the potion ("Press it on the pull with your trinkets and Blood Fury or Berserking, and again every 3 minutes."); card note `data/curated/apl/shaman-elemental.json`:96 now "Lightning Mastery, trinkets and racials together: ...". Action unchanged. The Spellblasting-on-pull sim (Sims requested 3) moves to the sim round.
- SHAMAN-R3-2 fixed: elemental.md:102, enhancement.md:105, restoration.md:109 "Windshaper Skyborne, the paid Heroic Pack race, brings Wind Blessed's 2% haste, always on."; index.md:36 "... joined by the Windshaper Skyborne, a paid unlock."
- SHAMAN-R3-3 fixed (verified: `trainables/shaman.json` 16190 talent, 17354 at 38, 17359 at 48, 1325256 at 58): restoration.md:117 adds "The talent teaches its first rank; train the others at 38, 48 and 58, each one a big step up."
- D1 applied per integrator race decision (round3/race-sims.md): Elemental Skyborne 633.4 (paid) > Orc 630.5 > Troll 624.8 ~ Tauren 624.6; Enhancement Skyborne 796.0 (paid) ~ Orc 795.6 > Troll 792.8 > Tauren 787.4. Orc kept as the Horde pick on both; `recommendedRaces: [dwarf, orc]` and the FS1 race fields (dwarf) unchanged; Orc lines stand as written with Skyborne named only as the paid option. Build strings unchanged.
- D2 applied: Restoration keeps `[dwarf, tauren]`; Skyborne line per R3-2.
- New SIM findings: none.
