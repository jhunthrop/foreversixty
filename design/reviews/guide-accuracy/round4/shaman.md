# Guide accuracy round 4: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree land/guides-r4 752c02e43, client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, the card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, and `round4/shaman-levels.md`.

## Builds and shipped leveling order

Build strings unchanged since round 1; all five re-checked legal on 70755 (51 points, tier gates, prerequisites). Shipped orders re-decoded from `data/builds/1.60.1.70755/addon-builds.json` (index 0 = level 10), identical to rounds 2 and 3:

- Elemental: Convection 10-14, Concussion 15-19, Elemental Devastation 20, Elemental Alacrity 21-23, Call of Flame 24-26, Elemental Focus 27, Elemental Warding 28-30, Call of Thunder 31, Lightning Overload 32-34, Elemental Fury 35-39, Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60.
- Enhancement: Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-25, Shamanistic Focus 26, Flurry 27-31, Stormstrike 32, Improved Ghost Wolf 33-34, Mental Quickness 35-36, Maelstrom Weapon 37-39 and 41-42, Rage of the Farseer 40, Convection 43-45, Concussion 46-47, Elemental Alacrity 48-50, Call of Flame 51-53, Tidal Focus 54-58, Mindfulness 59-60.
- Restoration: Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide Totem 25, Ancestral Healing 26-27, Healing Focus 28-30, Tidal Mastery 31-35, Restorative Totems 36-40, Healing Way 41-43, Nature's Swiftness 44, Purification 45-49, Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60.

## Levels 1 to 19 (round 4 focus): what checked out

Every level, spell and talent claim in the three new subsections against `trainables/shaman.json`, `spellconst/shaman.json`, `spelltips/shaman.json` and `talents/shaman.json`:

- Elemental (elemental.md:112): Lightning Bolt 1, Earth Shock 4, Flame Shock 10, Searing Totem 10, Fire Nova 408341 at 12 ("within 10 yd of your active Fire totem"); Convection 10-14 at 2% a point on shocks and Lightning Bolt, Concussion 15-19 at 1% a point on Lightning Bolt and Earth Shock (Chain Lightning, also named in the talent, is not learned until 32). No wand: right, shamans have no wand skill. "First quest staff or mace that carries spell damage": `items/shaman.json` has 32 staves and maces under level 20 with spell power (Staff of Orgrimmar 9, Crescent Staff 10, Ceranium Rod 10, Staff of the Purifier 13, Gnarled Necromancer's Staff 15, Emberstone Staff 17, Living Root 18 and others), so the advice is real in Forever.
- Enhancement (enhancement.md:115): Rockbiter 1, Earth Shock 4, Lightning Shield 8, Searing and Strength of Earth 10, Windfury Weapon 30. Rockbiter over Flametongue below 20: Rockbiter rank 2 (8) +58 AP and rank 3 (16) +88 AP, about 4.1 and 6.3 DPS, against Flametongue rank 1 (10) and rank 2 (18) at about 3.3 and 4.8 Fire DPS (tooltip per-hit range scales with speed); keeping Rockbiter to 30 also holds (rank 4 at 24, +129 AP, and Elemental Weapons from 23 adds 20% to it). Thundering Strikes 1% and Ancestral Knowledge 2% a point. Starting weapon skills (staves and one-handed maces; two-handed maces and axes from a weapon master) match Warcraft Tavern's Forever leveling guide. Leather to 40, no dual wield.
- Restoration (restoration.md:119): Healing Wave ranks 2/3/4 at 6/12/18; Improved Healing Wave 0.1 s a point; Mindfulness 15-16; Tidal Focus from 17 at 1% a point; Tremor Totem 8143 at 18 ("removing Fear, Charm and Sleep effects"), right for the Druids of the Fang in Wailing Caverns.
- No contradiction with the "Levels 20 to 60" text below each or with the strips (Water Shield 20, Tidal Focus full 22, Mindfulness 23, Mana Tide 25; Mental Dexterity from 20; Elemental's Devastation at 20). Totems still need their totem items in Forever (Warcraft Tavern Forever leveling guide; items 5175 to 5178 in `items.json`).

## Rest of the pages: what checked out

- Every talent value against the 70755 max-rank text, as in rounds 2 and 3; trainer levels and ranks; Searing Totem rank 6 55 s and 20 yards; Magma Totem 20 s, 8 yards; Fire Nova 10 s; Chain Lightning 6 s category cooldown; Lightning Bolt rank 9/10 190/220 mana (spell levels 50/56, both full bonus at 60); Water Shield 600000 ms; Healing Stream and Mana Spring 300000 ms, "group members"; Strength of Earth, Grace of Air and Windfury Totem "party or raid members within 30 yards"; Rage of the Farseer 25 s / 180 s; Totemic Call 437009 30 s.
- Round 3 fixes hold (R3-1 no potion on the Elemental pull, R3-2 Skyborne named only as the paid race, R3-3 Mana Tide ranks 38/48/58). Round 2 fixes hold. The Enhancement pull line's "your potion" is backed by the preset's Major Frenzy Potion at 55.
- Race lines against `round3/race-sims.md`: Dwarf + Orc (Elemental, Enhancement), Dwarf + Tauren (Restoration); racial texts match `races.json`.
- Professions: Living Crown 252561 is leather (Pattern: Living Crown), bind on equip; Hearts and Totem of Thunder bind on pickup.
- Card notes: all re-read; the Lightning Bolt rank note ("about 20% more base damage", 353 vs 295) agrees with the prose's "about 16% less"; the Healing Wave floor note ("ranks 1 to 4 keep none of it at 60") holds (rank 5, learned 24, keeps 5%).
- No em or en dashes; no site, sim, APL, preset or band talk in prose; no quoted nightly numbers.
- Source links: Blizzard Deep Dive recap 200 (names the Dwarf shaman and the hit/crit change for Enhancement, as cited); talentsforever.com/data.json 200; Wowhead Skyborne First Look 200; foreversixty.gg/sim/weights and /planner 200; Icy Veins talent calculator article 403 to automated fetches (unknown, not dead; the article exists under that title in search results).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R4-1 | CONFIRMED | source link | index.md:14, elemental.md:17, enhancement.md:17, restoration.md:16 | label `'Talents Forever (demo transcription)'` | The linked file's own `_readme` (https://talentsforever.com/data.json, fetched 2026-11-03): "talents, racials, Legacy perks and the class spellbooks come from the WoW Forever beta client's own data files (the site began as frame-by-frame reads of the BlizzCon 2026 demo)". The label names what the source used to be, not what it says it is now. Same label sits on the other classes' pages and in `races.json` sources (integrator: fix everywhere at once). | `label: 'Talents Forever, talent data from the beta client'` |
| SHAMAN-R4-2 | CONFIRMED | gear (item name) | enhancement.md:99 | "there is no room for a sharpening stone or weighted stone on it" | There is no "weighted stone": the items are Rough/Coarse/Heavy/Solid/Dense Weightstone (`items.json`), the blunt-weapon counterpart of a sharpening stone. Introduced by round 1's fix (607ab7de7). | "so there is no room for a sharpening stone or a weightstone on it." |
| SHAMAN-R4-3 | NEEDS-SIM | rotation / race | elemental.md:102 ("Troll's Berserking ... goes in the same window"), elemental.md:74, card `shaman-elemental.json` priorityList[1] ("stack everything inside it") | Berserking pressed inside Lightning Mastery | Inside Lightning Mastery a Lightning Bolt is 2.5 - 0.5 (Elemental Alacrity) - 0.5 = 1.5 s, exactly its 1.5 s GCD (`spellconst` 15208 cast 2500, gcd 1500). If Forever's GCD is not shortened by haste (the Classic 1.12 rule), 10% casting speed in that window buys no single-target bolts. The measured races agree with that: Troll 624.8 vs Tauren 624.6 (±0.30, `round3/race-sims.md`), where 10% casting speed for 15 s every 3 min should be worth about 0.8% (about 5 DPS) if it did anything. Either the window wastes it or the engine is not applying it; the guide line depends on which. | If the sim confirms Berserking gains nothing inside Lightning Mastery but gains outside it: elemental.md:102 "Troll's Berserking, 10% casting speed for 15 seconds every 3 minutes, goes outside Lightning Mastery: inside it your bolts are already at the global cooldown." and :74 "Press it on the pull with your trinkets and Blood Fury, and again every 3 minutes; a Troll saves Berserking for after it ends." Card note: "Lightning Mastery, trinkets and Blood Fury together ... (a Troll presses Berserking after Mastery ends)". If Berserking gains nothing anywhere, it is an engine bug (Sim findings) and the prose stands. |
| SHAMAN-R4-4 | LIKELY | leveling | elemental.md:112, enhancement.md:115, restoration.md:119 | "Do the Call of Earth class quest at 4 for Stoneskin Totem and Call of Fire at 10 for Searing Totem." | Stoneskin rank 1 is a level 4 trainer row (`trainables/shaman.json` 8071, skill_line_ability); what the quest gives is the Earth Totem item (5175), which every Earth totem needs, including the Strength of Earth Totem the Enhancement text drops from 10 and the Tremor Totem the Restoration text drops at 18. Warcraft Tavern's Forever leveling guide: the quests give "physical items that will need to remain in your bags so you can cast your totems", Call of Earth at 3 to 4 by start zone. Not false, but a new player who skips the quest and buys Stoneskin cannot cast it. | "Do the Call of Earth class quest at 4 and Call of Fire at 10: they give the Earth and Fire Totems you carry in your bags, and without them no Earth or Fire totem casts." |

## Decisions

- **D1. Call of Earth level.** Warcraft Tavern's Forever table lists 3 for some starts and 4 for others; Stoneskin, the first Earth totem, trains at 4 (`trainables`). "At 4" stays.
- **D2. Rockbiter to 30.** Kept: Rockbiter leads Flametongue at every trainer step from 8 to 29 on the 70755 tooltips (see above), and Elemental Weapons widens the gap from 23. No change.
- **D3. Enhancement pull potion.** "Your potion" stays: the 55 preset carries Major Frenzy Potion; the consumables table renders it.

## Sims requested

1. **shaman-elemental, Berserking timing** (decides SHAMAN-R4-3): band 60, raid preset, sim/cmd/spec-breakdown, 10,000 iterations, seeds 1 and 50001, `-talents` from raidBuild `FS1:1.60.1.70755:shaman:dwarf:5510331300001351/055/053002:` (engine form `5510331300001351-055000000000000000-0530020000000000`), race `troll`: (a) the published rotation; (b) the same with Berserking held until Lightning Mastery is not active (a Berserking line gated on `not(auraIsActive(Lightning Mastery 16166))` ahead of autocast, Berserking removed from autocast); (c) `tauren` for reference. Also read one combat log of (a): confirm Berserking is cast and that Lightning Bolt's effective cast under both auras is 1.5 s or 1.36 s.
2. If (b) beats (a) beyond error, re-run the Horde race order with (b) for Troll; Orc is expected to stay the pick (630.5).

## Sim findings

- SIM-R4-1 Berserking under Lightning Mastery (SHAMAN-R4-3): the engine reduces the spell GCD with cast speed down to 1 s (`sim/core/unit.go:374` `SpellGCD`, `sim/core/constants.go:25` `GCDMin = 1 s`), the later-expansion rule. If Forever keeps Classic's fixed 1.5 s GCD, the engine over-values every cast-speed effect that pushes a cast below 1.5 s (Berserking or Wind Blessed inside Lightning Mastery, Maelstrom Weapon is instant anyway). Yet Troll measured no gain at all, so check the log first (Sims requested 1).
- Carried, unchanged: Stormstrike `RequiredLevel: 40` (stormstrike.go:41) versus talent-taught at 25; Restoration card ranks; Grace of Air / Strength of Earth book ranks; Lightning Overload ratio on the first level-60 log; Totemic Call positioning; spelltips for Lightning Mastery (16166) and Lightning Bolt (15208) still carry pre-70755 text ("Lava Burst", 196 damage), a data-lane hover-tooltip item, not a guide finding.

## Sources

- Warcraft Tavern, "Shaman Leveling Guide" (Forever; totem quests at 3-4/10/20, totem items stay in your bags, default and trainable weapon skills, Tremor Totem removes fear, charm and sleep) https://www.warcrafttavern.com/forever/guides/shaman-leveling/
- Talents Forever, data export readme (data from the beta client's files; began as BlizzCon demo reads) https://talentsforever.com/data.json
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" (Dwarf Shaman; hit and crit combined helps Enhancement) https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; exists per search) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/
- Warcraft Wiki, "Casting speed" (later-expansion rule that haste lowers the GCD to 1 s; cited only to show the engine follows that rule) https://warcraft.wiki.gg/wiki/Casting_speed

## Fix log (round 4)

- SHAMAN-R4-1 fixed: label "Talents Forever, talent data from the beta client" at index.md:14, elemental.md:17, enhancement.md:17, restoration.md:16.
- SHAMAN-R4-2 fixed: enhancement.md:99 "a sharpening stone or a weightstone".
- SHAMAN-R4-3 fixed per integrator decision (Classic and Forever: spell haste does not shorten the 1.5 s global cooldown, so Berserking inside Lightning Mastery buys nothing): elemental.md:74 drops Berserking from the Mastery stack and adds "A Troll saves Berserking for after it ends."; elemental.md:102 Berserking "goes outside Lightning Mastery"; card note shaman-elemental.json:96 now "trinkets and Blood Fury", a Troll presses Berserking after Mastery ends (notes only, actions unchanged).
- SHAMAN-R4-4 fixed per integrator decision: elemental.md:112, enhancement.md:115, restoration.md:119 now say Call of Earth (4) and Call of Fire (10) give the Earth and Fire Totems every Earth and Fire totem needs, carried in your bags.
- D1, D2, D3: kept as decided, no change.
- Round 4 sources: restoration.md "Forever Sixty, stat weights" entry deleted (healer page); elemental and enhancement weights entries already point at /sim/weights; no /sources, rotation, tank fight or wago.tools links on any shaman page. "### Levels 20 to 60" headings kept.
- SIM finding (integrator): SIM-R4-1 stands as an engine bug. The engine lets cast speed shorten the spell GCD to 1 s (`sim/core/unit.go:374` SpellGCD, `sim/core/constants.go:25` GCDMin); Forever keeps the fixed 1.5 s global, so every haste effect that pushes a cast under 1.5 s (Berserking, Wind Blessed, inside Lightning Mastery) is over-valued. The Berserking-timing sims requested above are moot for the guide and go to the sim round.
