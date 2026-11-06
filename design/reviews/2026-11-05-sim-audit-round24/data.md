# Sim-accuracy round 24: data lane (the 70790 server push)

Site branch wt/r24-data, fork branch lane/r24-data, both from int/r24 (site 52bf8f23c, fork 215970648).

The 1.60.1.70790 hotfix push layered nothing at activation (the wago paging bug fixed on main in 80f4a3569, which re-layered hotfix-push.json and regenerated spelltips only). This lane ran the rest of the client-derived data through the pipeline from the layered push and followed it into the engine.

## How it was regenerated

1. `pipeline fetch --product wow_classic_era --build 1.60.1.70790` downloaded the raw tables into data/builds/1.60.1.70790/raw (gitignored). The live push search was still paging Spell after 18,000 requests, so I stopped it after the downloads and ran `pipeline layer-hotfix-push --build 1.60.1.70790 --offline` over the push cache the spelltips fix fetched (scratchpad/gr26/tipgen/.../raw/wago-push). The layer reproduces the committed hotfix-push.json byte for byte (same sha256, file unchanged), so it is the same push.
2. Site: `normalize`, `simconst`, `trainables`, `levels`, `spelltips`, `gametables`, `simdb`. Raw fill-in caches (classicdb, wowhead items) were restored to their committed versions first, so every output diff traces to the push. Not run: loot, BiS, weights, addon-data (nightly-owned or not push-fed).
3. Fork: `tools/follow.sh <site> 1.60.1.70790` with GOWORK=off. A go.work in the lane's parent directory broke the follow's build step on the first run.
4. Every changed file was diffed against the push. Only the files the push moves are committed (below). The rest is listed under "Files to regenerate".

Committed, each traced to a push row:

| file | change | push row |
|---|---|---|
| spellconst/warrior.json, fork client/warrior.json | Defensive Stance Passive 7376 aura 10: 30 -> 80 | SpellEffect 686475, push 113326 |
| spellconst/druid.json, fork client/druid.json, fork druid constants_auto_gen.go | Bear Form (Passive2) 21178 aura 10: 50 -> 115 (BearFormPassive2BaseDamage) | SpellEffect 700189, push 113326 |
| talents/warrior.json, fork talentgen/testdata/warrior.json | Defiance 12792 texts 5/10/15 -> 7/13/20 | CurvePoint 328755-328757 (curve 109358, TraitDefinitionEffectPoints 24567), push 113326 |
| spellconst/mage.json, fork client/mage.json | DNT 1321337 gains an effect 77 row (script, unused) | SpellEffect, added |
| spells.json | the 60 new spell names (NPC and world spells, the Endless line, mage portals) | Spell/SpellName, added |
| zones.json | 12 AreaTable rows (The Drowned City, Blackmaw Hold, ...) | AreaTable, added |

Unchanged by the push: trainables, spelltips (main already had them), levels (timestamp only), simconsumes.json. The Endless items' ItemEffect rows (239511-239689) and the new rows for Thunderfury, Benediction, Sulfuras/Immolation and others (2396xx) are not linked by any ItemXItemEffect row, so `build_consumables` and the item-proc paths do not read them.

## Findings

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| D1 | **Defiance was 5/10/15% threat; Forever's is 7/13/20%.** Round 21 (S2) held 5/10/15% because the client curve still read 5/10/15. The push now carries 7/13/20, so the client and Blizzard's 3 Nov note ("Defiance: Threat bonus increased to 7/13/20% from 5/10/15%") agree. The engine had a hand constant of 0.05 a rank. It is now a 0/7/13/20 table, and a test reads the per-rank values from the client's talent texts. Protection with a shield at 3/3: 1.8 x 1.20 threat (was 1.8 x 1.15). Wowhead's tooltip for 12792 still renders 15 (its live data has the stance's 80 but not this curve). | CurvePoint 328755-7, push 113326 (the same push as the stance's 80 and Bear Form's 115); Blizzard Forever notes 3 Nov | CONFIRMED-BUG | new-data (push layered late; round 21 decided on the shipped curve) | E a703a82d5, S 2091c640f, guide S 225b3c7c0 |
| D2 | **Defensive Stance 80% (round 21) against the hotfixed 7376 row.** The push carries 80 (SpellEffect 686475). The engine's 0.8 matches. The "cached client reads 30" constant and its test are gone: the test now checks the engine's 0.8 against the client row. | push 113326 | Aligned (cleanup) | - | E 74e228f07 |
| D3 | **Bear Form 115% (round 21) against the hotfixed 21178 row.** The push carries 115. The engine's hand 2.15 now reads the generated `BearFormPassive2BaseDamage` (1 + 115/100), with a test against the client row and the note. | push 113326 | Aligned (engine now reads it) | - | E 74e228f07 |
| D4 | **Natural flasks, usable at 60.** Flask of Natural Accuracy, Aggression, Precision and Swiftness (items 274273-274276; Alchemy recipes 1293559-1293562 on SkillLine 171, trivial 320-330; level 55; no faction). Push 113570 (later than 113326) lowers the shipped 60 Stamina / 5% (4% crit) to **30 Stamina and 2%** of hit / crit / dodge-and-parry reduction / haste in Mount Hyjal, Hyjal Summit and the Barrow Deeps. The 2 is the dummy aura 4 row; the stat auras 54/55, 290, 240, 342/65 have 0 base points, which the zone script fills. Wowhead's item tooltips still render 60/5%, the shipped rows. Client data beats web (single source: the push). They are in the engine (Flask enum 5-8, consumes.go `naturalFlaskStats`) and in the raid preset (below). | push 113570; SkillLineAbility 58410-58413; warcraft.wiki.gg Flask_of_Natural_* (beta values) | Added | - | E 2df1ff74c, S 80b28289b |
| D5 | **Elixirs and Scrolls of Endless: not added.** Elixir of Endless Intellect/Power/Healing/Rapidity/Opportunity/Mana/Growth/Experience, Scroll of Endless Fortitude/Speed (items 289082-289334, spells 1327736-1328127). They are required level 1, 4 silver, and stack 5 times (Power +3% damage a stack, Rapidity +4% haste, Opportunity +4% crit), with Experience at +200% XP and Growth an enlarge effect. They share consumable category 2593 with these sets: no recipe, vendor or drop names them in any client table or web source, and they ship next to "Ultra/Giga/Cosmic Hogger" and "Ding!". This looks like an event or test vendor kit. Putting a 5x stacking +15% damage elixir in the raid preset with no source would be wrong. | ItemSparse hotfix rows 289082-289334, ItemEffect 239511-239689 (no ItemXItemEffect), Wowhead spell tooltips; web search (6 Nov) finds nothing | OPEN: settle with an NPC or source for item 289083 (an in-game vendor, or Wowhead's "Sold by" tab) | - | - |
| D6 | **Righteous Fire 48198** ("Strike a zombie with righteous fire", target 25), with Righteous Purge 50146, Infect! 48953 and the "You're a Zombie!" rows: a zombie world event's spells, no class or spec. The engine reads none of them. | push SpellEffect changed/added | Aligned (no spec) | - | - |
| D7 | **The other changed rows:** Meathook, Ground Slam, Warstomp, Hateful Strike (+Primer), Bloodlust 1325968, Void Bolt, Shadow Bolt Volley, Cleave 1327248, Call Gnolls, Mangle! 57596, Makeshift Stickin' Spear, Impaled, Study, Caught Sheep, Retch!. These are all NPC or boss abilities. SpellLevels/SpellCooldowns/AuraOptions/SpellPower rows are on the new NPC and Endless spells. Noggenfogger's new ItemEffect row (cooldown 3 s) is unlinked. No spec, and the engine reads none of them. | push | Aligned (no spec) | - | - |

Total CONFIRMED-BUGs: 1 (new-data, plus guides-handoff in the sense that round 21's S2 asked for it when the client carried it).

## Raid preset flask choice

The raid preset fights Barrow Deeps and Hyjal content, where the 2% applies (at Onyxia only the 30 Stamina does). Raid 60, frozen BiS, guide raid builds, 10,000 iterations, seeds 1 and 50001, both factions, change against the group's current flask (ranges over specs, factions and seeds):

| group | current | Accuracy | Aggression | Precision | Swiftness | pick |
|---|---|---|---|---|---|---|
| melee (no flask before) | none | +0.5 to +2.9 | **+1.5 to +2.5** | +1.2 to +2.4 | +0.7 to +1.4 | Aggression |
| hunters (no flask before) | none | 0.0 to +1.9 | **+1.9 to +2.1** | 0 (BM/MM), +2.5 (SV) | 0 to +1.1 | Aggression |
| tanks (threat) | Titans | +1.0 to +2.6 | +1.3 to +1.8 | **+3.1 to +4.7** | +0.1 to +0.9 | Precision |
| casters | Supreme Power | -8 to -14 | -8 to -14 | -10 to -15 | -10 to -15 | Supreme Power |
| healers | Distilled Wisdom | -4 to -9 | -3 to -8 | -4 to -9 | -4 to -9 | Distilled Wisdom |

The preset groups are by role, so one flask fills each group. Aggression is the best flask for Feral, Combat, Subtlety, Enhancement, Fury, Beast Mastery and Marksmanship. It is within about 1% of the best flask for Retribution and Assassination (Precision, +0.6/+0.4) and Arms (Accuracy, +1.0), and within 0.5% for Survival (Precision). Tanks are judged on threat (owner ruling), so Precision replaces Flask of the Titans for all three.

## Before/after

Before: sb-pre built from int/r24 (52bf8f23c + fork 215970648). After: wt/r24-data + lane/r24-data. Frozen BiS, guide raid builds (`$R/rt`), 10,000 iterations. Columns: Alliance seed 1, Alliance 50001, Horde 1, Horde 50001 (%), then absolute Alliance seed 1.

Raid 60 (DPS; tanks threat/s):

| spec | A1 | A50001 | H1 | H50001 | A1 abs |
|---|---|---|---|---|---|
| warrior-protection | +8.88 | +8.90 | +9.04 | +8.97 | 1465.5 -> 1595.7 |
| druid-feral-bear | +4.67 | +4.45 | +4.36 | +4.57 | 1715.1 -> 1795.2 |
| paladin-protection | +3.34 | +3.40 | +3.09 | +3.07 | 1129.2 -> 1166.9 |
| warrior-fury | +2.35 | +2.38 | +2.44 | +2.45 | 886.9 -> 907.7 |
| hunter-marksmanship | +2.02 | +2.01 | +2.05 | +2.06 | 959.3 -> 978.7 |
| hunter-survival | +2.01 | +2.02 | +2.02 | +2.02 | 851.1 -> 868.2 |
| hunter-beast-mastery | +1.91 | +1.96 | +1.95 | +2.01 | 868.3 -> 884.9 |
| warrior-arms | +1.91 | +1.94 | +1.97 | +2.01 | 828.5 -> 844.3 |
| shaman-enhancement | +1.93 | +1.89 | +1.97 | +1.95 | 833.7 -> 849.8 |
| rogue-assassination | +1.79 | +1.77 | +1.82 | +1.92 | 792.8 -> 807.0 |
| paladin-retribution | +1.68 | +1.68 | +1.77 | +1.76 | 772.7 -> 785.7 |
| druid-feral | +1.70 | +1.67 | +1.77 | +1.78 | 888.0 -> 903.1 |
| rogue-subtlety | +1.62 | +1.63 | +1.76 | +1.76 | 795.8 -> 808.7 |
| rogue-combat | +1.53 | +1.51 | +1.65 | +1.66 | 780.2 -> 792.1 |
| every caster and healer (14 specs) | 0.00 | 0.00 | 0.00 | 0.00 | unchanged |

Warrior Protection's raid gain is Defiance (+4.3%) times the Precision flask (+4.4%). The other tanks' and the melee and hunter gains are the flask alone.

Bare 60: Warrior Protection +4.28% threat at every seed and faction (887.7 -> 925.7, Defiance). The other 27 specs are unchanged (0.00): the bare kit carries no flask, and no DPS or heal spell moved.

TestP1TankWarrior.results (fork) regenerated for Defiance (TPS +4.3%), the only fork golden this change explains.

## Files to regenerate (integrator, at the pin)

- **Fork, full `make follow SITE_DIR=<site> BUILD=1.60.1.70790`** (run with GOWORK=off). Beyond the push rows committed here, it moves the conformance trainables copies (all nine classes: the site's talent-learn levels from round 22 and the rank rows of later lanes), the hunter/mage/paladin/priest/shaman/warlock/warrior goldens and SUMMARY.md (Lacerate and Mongoose Bite rows, Reckless Swings), client_sets_gen.go (Battlegear of Undead Slaying / Undead Slayer's Armor PhaseOnePieces 2 -> 1), and every constants_auto_gen.go Source header. None of that is from the push, so I did not commit it.
- **Site:** manifest.json (70790: rehash spellconst, talents, spells, zones); simbuffs.json (names and icons for the four `flask_of_natural_*` ids now in IDS.md, via loot-merge or the next full loot); then spelltips and icons (the flasks' on-use lines and art); addon-data and addon-builds (the Defiance text in talents/warrior.json); coach rules and starters and the planner rules if they read talent texts; ladder goldens (TestRotationLadder: Protection threat, Defiance) and dps-per-point.
- **BiS re-rank:** the raid preset now carries Aggression (melee and hunters) and Precision (tanks), and Protection warrior's threat moved. Every melee, hunter and tank raid set is ranked under different consumables.
- **Not regenerated, for the data owners:** a fresh `normalize` on today's raw item tables rewrites items/*.json, items.json (it also blanks the loot columns until `loot` runs), itemnames.json (+3 items), simitems.json and simdb.bin (18 item ids change source; 287072 added, 3334, 18610 and 18611 dropped). These come from the item pair and wowhead/classic-db caches, not the push. They are left as committed.

## Measure at refresh

- The per-spec flask: Retribution and Assassination may prefer Precision (+0.6/+0.4 over Aggression), Arms Accuracy (+1.0), Survival Precision (+0.4). A per-spec flask needs a per-spec preset slot, which the preset groups lack (a class kit would also reach the bare preset). Class lanes can decide whether it is worth one.
- Wowhead's flask tooltips still read 60 Stamina / 5%. If a later push or a log shows 60/5%, the engine table and the preset reasons change (one table, `naturalFlaskStats`).

## For other lanes

- Warrior: D1 changes Protection threat everywhere a shield and Defiance are taken (raid and leveling bands). The guide lines are updated (protection.md:52 and :71).
- Every melee, hunter and tank lane: the raid preset now carries a flask (S 80b28289b). Raid-60 baselines move by the table above.

## For the guide lane

- protection.md Defiance lines now say 20% and "about 2.2 times your threat" (S 225b3c7c0). `_prose.test.ts` was not run locally: the worktree has no node_modules, and a symlinked one does not resolve astro. The new wording has no dashes and no meta talk.

## Still open

- D5 Endless elixirs and scrolls: need a source (vendor, drop or event) for items 289082-289334 before any preset uses them.
- Logs: GET https://api.foreversixty.gg/v1/reports/recent at 10:52 UTC 6 Nov lists only the 20 Oct report (no kills). Nothing can check the 2% flask value or Defiance's 20% yet. A Protection warrior's Sunder threat from a Barrow Deeps log with and without the flask would settle both.
