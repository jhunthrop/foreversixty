# Sim-accuracy round 24: integration (2026-11-06)

Round 24 started from the landed round 23 (site f05c281fe, fork 3b4d6b054) in build mode, at most two lanes at once, with rate-limit pauses. Main moved under it (bis-heal2, redesign-1 and -2, guild migrations, bis-source-kinds, sim-source-blizzard, spelltips-fix) and is merged in.

**Not a zero round: 70 confirmed bugs, all fixed** (round 23: 19, round 22: 33). The count is high because the hunter lane found the measuring tool judged rank rows on the band floor's talents; the platform lane fixed the tool and every lane re-judged its rows (193 suspect rows listed in `platform-suspect-rank-rows.tsv`).

| Lane | Confirmed | What |
|---|---|---|
| priest | 13 | Discipline's Alliance race is Human; Renew, Heal and Greater Heal rank rows (Greater Heal rank 1 holds to 59); Holy Renew race rows 28-35; Shadow Mind Blast and Devouring Plague rows |
| druid | 12 | Balance, Feral and Restoration rank rows below 20 and after new ranks; the second pass fixed 6 of 40 re-judged rows |
| paladin | 11 | Protection Consecration on a mana projection (manaIncomePerSecond minus manaSpentPerSecond); Retribution Consecration rank 1 at 30-52 and rank 3 at 53-59; Holy Strike, Hammer of Wrath, Holy Light and Flash of Light rows; two Consecration rows that repeated the top rank deleted |
| shaman | 11 | Troll Berserking was six cooldowns on one timer (core racials) and fired with Lightning Mastery; Lightning Bolt rank 3 at 20-33; Flame Shock, Earth Shock and tank Healing Wave rows |
| casters | 8 | Fire Blast, Pyroblast, Scorch and Fireball rows; Arcane and Frost Frostbolt rows; Destruction and Demonology Shadow Bolt rows |
| platform | 8 | the leveling order follows the guide's sequence across trees; talent-owned auras start card rows; a duration constant in number times duration; the web build reader and an empty tree; spec-breakdown -level uses the player's talents at the level; a forced -rank wins at bare 60; rank-tagged card gates; manaSpentPerSecond built |
| hunter | 5 | every Mongoose Bite and Lacerate rank registers; Beast Mastery, Marksmanship and Survival rank rows on the player's talents |
| items | 1 | the BiS ranker's wand pick was decided by noise (Torch of Light in the Alliance Destruction raid set); near-ties are re-measured |
| data | 1 | Defiance 7/13/20% threat (the 70790 push and Blizzard's 3 Nov note) |
| melee, healbisect | 0 | rogue and warrior rank sweep aligned (297 runs, newest rank always); the healer bare-band drops are correct (below) |

## The 70790 hotfix push

The 1.60.1.70790 push layered nothing at activation (wago's search paging was broken). Main's spelltips-fix (80f4a3569) re-layered it; the data lane regenerated the client-derived data through the pipeline from that layer (offline, byte-identical to the committed hotfix-push.json). The full table diff is `hotfix-diff-70790.txt` in this folder. What it moved for the sim: Defensive Stance 7376 at 80% and Bear Form 21178 at 115% now come straight from the hotfixed rows (the hand constants are gone); Defiance 7/13/20%; the Flasks of Natural Accuracy, Aggression and Precision (level 55, 30 Stamina plus 2%), added to the raid preset (melee and hunters Aggression, tanks Precision; casters and healers keep their flasks, which beat every Natural flask). The Elixir and Scroll of Endless spells are level 1, cost 4 silver and have no known source, so they are not in the catalogue. Righteous Fire and the other rows are world-event or NPC spells.

## Healer bare-band drops (BiS lane's question)

Correct, not a regression: round 15's level-gap spell power penalty on low-rank heals (fork 8d9ca6aae; Blizzard's 1 Oct note and two independent tooltip measurements). The "r16 picks" were ranked at engine d1d9f2775 (round 14), before the penalty. On the published gear: Restoration bare 60 627.8 to 415.3, Discipline bare 60 348 to 204; later rank rows, gates and re-ranked gear win part of it back. The ranker and spec-breakdown use the same heal profile. Full bisect: `healbisect.md`.

## Integration

- **Fork:** int/r24 = bf727699f (pushed), with goldens for Defiance, the flasks and the new hunter ranks; conformance keeps amd64 values on rounding-only rows.
- **Site pin:** 67b01a11f, with simproto re-vendoring apl.proto and common.proto for manaSpentPerSecond.
- **Regen at the pin (b57a02515):** simdb, raid-consumes and published presets, addon-builds, dps-per-point, ladder goldens, addon-data and Data.lua, coach starters via persona_starter.py --all.
- **Finish lane (28fb9c1f6):** `loot-merge --buffs-only` (new, with tests) names the four Natural flasks in simbuffs.json without touching loot.json; the manifest refresh re-hashed stale entries; the Survival trap note scoped to the rank trained at 46 and 20 stale card-note allowlist lines removed (main's new card-note-level guard); the guides' consumable "what it gives" lines for Natural Aggression and Precision; the web card matching skips only steps that end at a talent rank (Arcane Heating Up). **Known red:** `web/src/lib/bis/rotation-icons.test.ts` for shaman-elemental: Berserking 20554 (cast by round 24's Elemental line) has no entry in spellicons.json. Fixing it needs the `icons` step on fresh raw tables; a regen from cached raw added 20554 but churned 191 other entries, so it was reverted. The next data run fixes it.
- **Integrator test updates:** healer, Shadow Bolt, Mind Blast race and Holy paladin tests follow the new rows; the Affliction mana potion test asks for at least one drink, since Affliction on the re-ranked BiS barely Life Taps.
- **Not regenerated here:** `simbuffs.json` names for the four Natural flask ids (written by the `loot` step, which also rewrites the nightly-owned loot.json); the BiS (its own run, on this pin: melee, hunter and tank raid sets now rank under the Natural flasks; rogues gained one-hand axes in round 22; healers on the level-gap penalty).
- **Checks run (build mode):** sim request, rankchoices, leveling, addon-builds, spec-breakdown; data test_apl, test_addonrotation, test_addondata, test_genproto; addon-data --check; luac; apl-check.
- **Measuring note:** the hunter lane's finding (spec-breakdown -level used the band floor's talents; a forced -rank lost to the bare-60 row) affected every rank row judged above a band floor in rounds 15 to 23. The fixed tool is on this pin; the rows each lane re-judged are listed in its report.

## Before/after, raid 60

Before: round 23's landed engine and site. After: int/r24 at the pin. Same frozen BiS, guide raid builds, 10,000 iterations at seeds 1 and 50001, both factions. Only rows that moved. Most of the round's fixes act below 60; at raid 60 the moves are the Natural flasks and Defiance.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| warrior-protection (threat) | horde | 1453.2 to 1584.5 | +9.04 / +8.98 |
| warrior-protection (threat) | alliance | 1465.5 to 1595.7 | +8.88 / +8.90 |
| druid-feral-bear (threat) | alliance | 1715.1 to 1795.2 | +4.67 / +4.45 |
| druid-feral-bear (threat) | horde | 1710.5 to 1785.1 | +4.36 / +4.57 |
| paladin-protection (threat) | alliance | 1125.6 to 1166.9 | +3.67 / +3.58 |
| paladin-protection (threat) | horde | 1123.8 to 1161.9 | +3.39 / +3.40 |
| warrior-fury | horde | 878.2 to 899.6 | +2.44 / +2.45 |
| warrior-fury | alliance | 886.9 to 907.7 | +2.35 / +2.38 |
| hunter-marksmanship | horde | 960.0 to 979.7 | +2.05 / +2.06 |
| hunter-marksmanship | alliance | 959.3 to 978.7 | +2.02 / +2.01 |
| hunter-survival | horde | 850.7 to 867.9 | +2.02 / +2.02 |
| hunter-survival | alliance | 851.1 to 868.2 | +2.01 / +2.02 |
| warrior-arms | horde | 830.4 to 846.8 | +1.97 / +2.01 |
| shaman-enhancement | horde | 839.5 to 856.0 | +1.97 / +1.95 |
| hunter-beast-mastery | horde | 865.5 to 882.4 | +1.95 / +2.01 |
| shaman-enhancement | alliance | 833.7 to 849.8 | +1.93 / +1.89 |
| hunter-beast-mastery | alliance | 868.3 to 884.9 | +1.91 / +1.96 |
| warrior-arms | alliance | 828.5 to 844.3 | +1.91 / +1.94 |
| rogue-assassination | horde | 790.2 to 804.6 | +1.82 / +1.92 |
| rogue-assassination | alliance | 792.8 to 807.0 | +1.79 / +1.77 |
| paladin-retribution | horde | 779.8 to 793.6 | +1.77 / +1.76 |
| druid-feral | horde | 871.2 to 886.6 | +1.77 / +1.78 |
| rogue-subtlety | horde | 789.0 to 802.9 | +1.76 / +1.76 |
| druid-feral | alliance | 888.0 to 903.1 | +1.70 / +1.67 |
| paladin-retribution | alliance | 772.7 to 785.7 | +1.68 / +1.68 |
| rogue-combat | horde | 781.6 to 794.5 | +1.65 / +1.66 |
| rogue-subtlety | alliance | 795.8 to 808.7 | +1.62 / +1.63 |
| rogue-combat | alliance | 780.2 to 792.1 | +1.53 / +1.51 |
| paladin-holy (HPS) | horde | 684.6 to 684.6 | +0.01 / -0.00 |
| priest-discipline (HPS) | alliance | 530.4 to 530.4 | +0.00 / +0.01 |
| paladin-holy (HPS) | alliance | 701.6 to 701.6 | +0.00 / +0.01 |
| priest-holy (HPS) | alliance | 771.2 to 771.2 | -0.00 / +0.01 |
| shaman-restoration (HPS) | horde | 740.6 to 740.6 | -0.01 / -0.00 |
| priest-discipline (HPS) | horde | 526.0 to 526.0 | -0.01 / -0.00 |

## Open, for round 25

- **Coach code (guide lane stopped):** talent-shortened cooldowns in F1, Evasion at spell deaths in F7, the Shadowburn shard floor; the coach's "above half mana" Protection override (`tools/coach/starter_overrides.py:260`) and the Discipline/Holy Greater Heal "until 58" lines.
- **Failing test:** the `internal/rating` presence evidence-rule test.
- **Platform:** bare band 20 runs rogues unarmed below 19 and hunters with no ranged weapon below 20, so melee and hunter rows at 10 to 19 cannot be judged; `hotfix-push.json`'s staleness checks; the Natural flask names in simbuffs at the next data run.
- **Race-keyed rows judged on a subset of races** (hunter on Troll and Dwarf): re-check the rest at refresh.
- **Live logs:** none; Warcraft Logs key 401 (owner).
- **Guide lane:** the Fury Berserker Rage note names Magmadar and Nefarian (Forever's first raids are Barrow Deeps, Hyjal Summit and Onyxia); protection.md:69; feral-bear.md:83.
