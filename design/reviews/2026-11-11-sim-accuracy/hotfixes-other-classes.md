# Sim accuracy 2026-11-11: the 6, 9 and 10 Nov hotfixes for the other eight classes

Lane brief: `design/reviews/sim-accuracy/brief.md`. Method: the warlock lane's (`warlock.md`), for warrior, paladin, hunter, rogue, priest, shaman, mage and druid. Site from origin/main 5a97ebd12, fork from master 278876977 (the pin). Client on the site: 1.60.1.70790. Live client: 1.60.1.70888 (wago `wowf`, 10 Nov; still the newest wago lists).

**Totals: 11 confirmed engine bugs, all new data from Blizzard's 9 and 10 Nov hotfixes, all fixed (shaman 6, hunter 3, mage 1, druid 1). Two site rules follow from them: Enhancement keeps Windfury Weapon under the raid's Windfury Totem, and every Arcane Shot rank row is gone. One card change was adopted (Enhancement's five-stack Lightning Bolt at any rank). Warrior, paladin, priest and rogue: nothing the engine models changed. Engine pin 278876977 to 7d0e129ad.**

## Method

1. **Blizzard since 4 Nov.**
   - Read "WoW: Forever Hotfixes" (t/2403040) line by line: posts 4 (6 Nov), 5 (9 Nov) and 6 (10 Nov).
   - Also "Class Tuning Incoming" (t/2404092) and the development notes thread (t/2360696, posts 9 and 10; nothing new for these classes since round 21).
2. **Client and push.**
   - From wago.tools, for 70790 and 70888, both hotfixed (`useHotfixes=1`, the push layered) and shipped: SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellCategories, SpellClassOptions. Also SpellName, SpellLevels, Spell (text), SpellInterrupts, SpellAuraRestrictions, SpellTargetRestrictions, SpellCastingRequirements, SpellEquippedItems, SpellDuration, SpellCastTimes, SpellProcsPerMinute and Talent.
   - Every changed row was diffed, keyed to its spell and grouped by SpellClassSet. The 1,571 rows with no class set were read by spell name.
   - The method reproduces the warlock lane's six rows exactly (Wrack 1349556/1349557, Firebolt 3110 to 7801, Bane of Havoc), which checks it.
   - wago's push log (`/hotfixes`) was read record by record for the rows the notes name but the tables do not show.
     - Stormstrike 17364 effect 2 is still 150 in push 113380.
     - Pushes up to 1139xx carry no row for the Enhancement changes.
   - Wowhead's live Forever talent payload and tooltips also still read Stormstrike 150% on 6 s, Mental Dexterity 50/100/150 and Maelstrom Weapon 4% a rank.
   - So the Enhancement, Arcane Concentration and hawk changes are server side. They are in the engine as hand values from Blizzard's notes, each comment naming the post.
3. **Engine read against both.** Each finding has a failing test first (fork d848dc8f4, red on the box), then the fix.
4. **Sims** on the Windows box (`tools/remote-test.sh`, one `spec-breakdown` at a time, `nice -n 15`, GOMAXPROCS=6):
   - raid 60 on both factions at seeds 1 and 50001, 10,000 iterations;
   - bare 20 to 60 for Horde at seed 1;
   - the old pin on origin/main against the new engine with this landing's site changes.
5. **Logs.**
   - foreverlogs.gg's public reports since the deploy: 2845, 2847 and 2849, plus 2848 (recorded 9 Nov) as a control. All are 15 to 45 dungeons.
   - Our API has no new report.

## Confirmed and fixed

| # | class | finding | source | fix (fork) |
|---|---|---|---|---|
| S1 | shaman | Stormstrike hit for 150% weapon damage on a 6 s cooldown. Now 100% on 8 s. | t/2403040 post 6; server side (push 113380 still 150, RecoveryTime 6000) | `stormstrike.go` |
| S2 | shaman | Mental Dexterity gave 50/100/150% of Intellect as attack power. Now 33/66/100%. | post 6; server side (trait curve 85272 unchanged) | `talents.go` |
| S3 | shaman | Maelstrom Weapon took 4% a rank per stack. Now 20% a stack at every rank: five stacks are an instant, free bolt from the first point. | post 6; server side (408498 still 4/8/12/16/20) | `talents.go` |
| S4 | shaman | Maelstrom Weapon procced on every melee hit (100%, round 21's reading of "chance doubled"). Now 6/7/8/9/10 procs a minute by rank, on the weapon's speed. The tooltip example matches: 5 points on a 3.6 weapon is 60%. Totem of the Storm's bolt rolls half the main hand's chance. | post 6 | `talents.go`, `items.go` |
| S5 | shaman | Windfury and Flametongue Weapon on the main hand shut out the matching totem. Now both work together. | post 6 ("Weapon imbues and Weapon totems can now be used at the same time") | `core/buffs.go`, `core/flametongue_totem.go` (the reach rule deleted) |
| S6 | shaman | Lightning Overload's rows were the 70755 ones (rank 10: 98 + 0.6 a level at 0.357; Chain Lightning 62 + 1.2 at 0.2855). Client 70888 makes every overload row half its bolt: 176.5 + 1.05 a level, Chain Lightning 255 + 1.65 at 0.357. | post 5 ("Lightning Overload now deals the correct amount of damage for Lightning Bolt"); client 70888 SpellEffect 1072563-1072638 | `lightning_bolt.go`, `chain_lightning.go` (hand rows until the client move) |
| H1 | hunter | Arcane Shot cost a flat 25 to 190 mana by rank. Now 11% of base mana at every rank. | post 6; client 70888 SpellPower (PowerCostPct 11, ManaCost 0) | `arcane_shot.go` |
| H2 | hunter | Summon Hawk cost a flat 80 to 190. Now 11% of base mana. | post 6; client 70888 SpellPower | `summon_hawk.go` |
| H3 | hunter | The hawk's swing was the beta logs' 0.32 a level. Now 1.8 times that: "Melee damage dealt by the Hawk pet increased by about 80%". | post 6; server side | `summon_hawk_assault.go` |
| M1 | mage | Arcane Concentration rolled 2% a rank on every Arcane Missiles missile (a 5/5 channel left Clearcasting 38% of the time on the box). Now one roll of the rank's chance is shared by the cast and the missiles, about 9.6% a channel. | post 5 ("splits across the cast and every bolt"); server side (11213 unchanged) | `talents.go`, `arcane_missiles.go` |
| D1 | druid | Nature's Swiftness did not reach Wild Growth. Client 70888 widens 17116's mask_1 from 0x2000000 to 0x6000000. | post 6; client 70888 SpellEffect 695239 | `druid.go` |

**Follow-on site changes, measured under the adoption rule:**

- **Windfury Weapon stays under the raid's Windfury Totem.**
  - `request/preset.go` loses round 15's swap to Rockbiter, and `raid-consumes.json` is regenerated.
  - Windfury Weapon beats Rockbiter by +10.9, +10.7, +10.8 and +10.7% at raid 60 (Horde and Alliance, seeds 1 and 50001; errors 0.4 to 0.6).
  - Raid presets exist only at 60.
- **Every Arcane Shot rank row is removed** (BM and MM, `data/curated/rank-choices.json`).
  - Each rank now costs the same, so a lower rank is strictly worse. The rows were all mana-driven down-ranks.
  - The newest rank wins at every level. A new test pins it (`hotfix_arcane_shot_rank_test.go`).
- **Enhancement card.** The five-stack Lightning Bolt line no longer waits for 5/5.
  - The gate keeps the talent, untagged, so the addon starts the row at the first point (35).
  - Gains +4.8, +5.8, +6.5 and +9.1% at levels 35 to 38 (seeds 1 and 50001, errors 0.07 to 0.08). Identical from 39, at bare 40 and 60, and at raid 60.

**Fork also carries** a conformance "server hotfix" verdict for Stormstrike's 8 s, which no client row states yet, and the conformance goldens (hunter costs; the warlock lane's stale Wrack golden is regenerated too).

## Checked, nothing to change

| class | item | why no change |
|---|---|---|
| druid | Natural Instinct now reaches Lacerate and the Pounce bleed (9 Nov, 70888 mask_3 +0x88) | the engine already applied it to both |
| druid | Idol of Ferocity no longer discounts Mangle (27851 mask_1 0x400 removed) | the engine discounts Claw and Rake only |
| druid | Bloom heals the target (6 Nov) | `bloom.go` already heals the HoT's target for the Spirit amount |
| druid | Poised rows rebuilt (1328825/1328829, new 1330741-1330765) | trigger plumbing only; the buff's -10 Energy and its text are unchanged |
| druid | Ravage, Bear Form and Swiftmend tooltips; Feral Charge combat; Omen of Clarity after a spec swap; Cooldown Manager | no sim effect |
| mage | Arcane Blast's debuff is 18 s (6 Nov) | already 18 s |
| mage | Flamestrike ticks no longer spend Clearcasting (9 Nov) | ticks never spent it; only casts do |
| paladin | Consecration's bonus could stop after a despawn (1280353) | multi-target only |
| paladin | Unsealed Wrath needs a weapon (SpellEquippedItems added) | a sim character always has one |
| priest | Revelation enchant procced from Vampiric Embrace; Power Infusion on allies; Shadowform and Fear Ward; Holy Nova unlearn | none of these is modelled or on a card |
| rogue | Quick and the Dead legacy talent; Shank tooltips | the legacy talent is not modelled; tooltips only |
| warrior | Sweeping Strikes keeps its charges without a second target (10 Nov) | single target: no second target either way |
| warrior | Reckless Swings tooltip; Thunderclap in the Cooldown Manager | no sim effect |
| hunter | pet health 115%, not 150%; Great Stamina rows (4187-5042); pet families; Frost Trap trigger | the engine models no pet health |
| all | new 70888 rows Shoot 1283679, Shadow Shot 1283365, Charge 1283073, Shield Charge 1283407, Frost Armor 1283419, Blast Wave 1291176 | new ids no card casts and the engine registers for no class; the client move's trainables check (MR46) says whether a player learns any |

Items, for the item owner and the client move, not class rows:

- Flamebreak Talisman (1315878): 50 Agility to 33, cooldown 3 to 5 min.
- Poison Shock (1312870): 750 to 75.
- Searing Dagger: damaging spells only.
- Violet Sorcerer's Vestments 5-piece (1324323): a 15 s proc cooldown. Set 2134 is still unregistered.
- None has hand engine code; simdb reads them at the client move.

## Before and after

DPS, 10,000 iterations, published BiS. The old pin 278876977 on origin/main against the new engine with this landing:

| spec | raid 60 Horde s1 / s50001 | raid 60 Alliance s1 / s50001 | bare 20 / 30 / 40 / 50 / 60 (Horde) |
|---|---|---|---|
| Enhancement | -9.0 / -9.0% | -9.0 / -9.2% | -5.2 / -19.6 / -19.2 / -19.5 / -19.9% |
| Elemental | +2.35 / +2.35% | +2.33 / +2.34% | 0 / 0 / +1.4 / +1.1 / +1.8% |
| Arcane | -1.44 / -1.31% | -1.26 / -1.33% | -13.1 / -3.3 / -2.5 / -3.1 / -3.0% |
| Fire | -0.40 / -0.48% | -0.27 / -0.26% | 0 / 0 / 0 / 0 / -1.6% |
| Frost | -0.07 / -0.04% | -0.06 / -0.01% | 0 / 0 / 0 / 0 / -1.0% |
| Beast Mastery | 0.0 / 0.0% | 0.0 / 0.0% | -2.1 / -5.9 / -7.3 / -6.5 / -5.4% |
| Marksmanship | 0.0 / 0.0% | 0.0 / 0.0% | -2.6 / -2.3 / -2.4 / -4.2 / -4.5% |

- **Enhancement at raid 60** is net of Windfury Weapon under the totem. On Rockbiter, the hotfixes alone were -18%.
- **The hunter's leveling losses** are Arcane Shot's new price: a leveling hunter no longer has a cheap low rank.
- **Arcane at bare 20** spends 52 s of the 180 s fight out of mana (21 s before). The leveling card leans on Missiles, which now rarely leaves Clearcasting.
- **Unchanged:** Restoration druid's ladder (its card casts Healing Touch right after Nature's Swiftness and never Wild Growth). No warrior, paladin, priest or rogue number moves.

## Logs

- **Before vs after the deploy.** foreverlogs.gg reports 2845, 2847 and 2849 start after the 10 Nov deploy window. 2848 is from 9 Nov.
- **Lightning Bolt.** One Lightning Bolt in all four (an Enhancement shaman in 2847). Too thin for overload rates.
- **The hawk (H3).**
  - Report 2847 has one Beast Mastery hunter, between level 20 and 25 by the ranks he casts (Serpent Sting rank 3, Aimed Shot rank 1, no Serpent Sting rank 4). His hawk lands 220 swings on Wetlands crocolisks, median 9 non-crit, range 6 to 12.
  - That sits between the beta logs' reading at that level (about 7 to 8) and the hotfix's 1.8 times (about 13 to 14).
  - With one hunter, an uncertain level and no deploy time to the hour, it neither confirms nor refutes the 80%. **Open:** a level-known hawk after 11 Nov.
- **Warcraft Logs** has no Forever zone.

## Gates

- **Fork, on the box, in a writable copy:** `go test -tags=with_db` over every `sim/...` package but `sim/web` (41 ok).
  - `sim/web` fails setup on the old pin too: `binary_dist` is not in the module.
  - Mage, Elemental and Enhancement suite results regenerated.
- **Site, on the box:**
  - `go test` for `sim/request`, `sim/leveling`, `sim/rankchoices` and `sim/specs`.
  - `npm test` for `src/content/guides`, `src/lib/guides`, `src/lib/sim` and `nav-guides` (132 files).
  - `astro check` (0 errors) and `make apl-check`.
  - `data/tests/test_addonrotation.py` and `test_apl.py`.
- **Already red on origin/main, not this lane's:**
  - The ladder goldens for paladin-protection, paladin-retribution, priest-shadow, the three rogues and the three warriors: the warlock lane's gear-data drift.
  - `data/tests/test_apl.py::test_every_spell_the_rotations_name_exists_with_that_rank` (rogue-assassination, spell 1329549).
- **Regenerated here:** the ladder goldens for the classes this lane touches, Elemental, Enhancement, the three hunters and the three mages. The hunter and mage ones were among the 15 red on origin/main and now carry the gear drift too.

## Guide and card

wow-player read every touched page and card.

**Shaman: NO-SHIP, then SHIP.** Three MUSTs were taken:
- Stormstrike is "your hardest hit" no more;
- the mana paragraph's "hits harder than the shock you skip";
- Maelstrom Weapon described as a per-hit chance.

**Hunter: SHIP.** "far more for its mana" was trimmed and the rank paragraph compressed.

**Mage and druid: SHIP after fixes.**
- Arcane Concentration's Missiles line and the 15 to 19 leveling line.
- Nature's Swiftness now names Wild Growth.

Pages and cards changed:

- **Enhancement (`enhancement.md`):** Stormstrike, Mental Dexterity (talent and 15 to 17), Maelstrom Weapon (talent, packs and 35 to 39), the imbue (before the pull, Which imbue, Which Air totem, Enchants), the mana paragraph and the stat paragraph. Card: the Grace of Air, Stormstrike and Lightning Bolt notes, and the bolt's condition.
- **Elemental (`elemental.md`):** Lightning Overload at half damage.
- **Hunter (`beast-mastery.md`, `marksmanship.md`):** Arcane Shot costs the same at every rank. The rank lists are now Serpent Sting only. Card: "does more for its mana".
- **Arcane (`arcane.md`):** Arcane Concentration and Missiles; 15 to 19.
- **Restoration druid (`restoration.md`):** Nature's Swiftness and Wild Growth.

Left open by the reviewers:

- Enhancement's front-matter stat order still ranks Intellect second. Blizzard means Strength to lead. Re-measured at refresh (MR42).
- The 25% Earth Shock reserve, tuned on the old Stormstrike.
- The leveling imbue line "below 20 ... worth more than Flametongue Weapon", which predates the hotfix.
- The hunter "keep Arcane Shot firing on cooldown even then" line under a quarter of mana.
- Arcane's leveling Missiles line ("more damage for its mana than Arcane Blast").

## Measure at refresh (ledger MR42 to MR46)

- **MR42, Enhancement:** every band and raid 60. BiS, weights, the stat order (Strength against Intellect after Mental Dexterity's cut), races, talents (Maelstrom Weapon's ranks now buy procs, not size), the leveling order and the 25% Earth Shock reserve.
- **MR43, Elemental:** Lightning Bolt's rank rows (the overload is half the bolt again), weights and BiS.
- **MR44, Beast Mastery and Marksmanship:** the Serpent Sting rank rows, judged jointly with the old Arcane Shot rows, and the leveling mana lines. Beast Mastery's builds, with Summon Hawk after +80% (neither build takes it today).
- **MR45, Arcane:** the leveling card at 20 to 29 (Missiles against Frostbolt now that Missiles rarely leaves Clearcasting), weights and BiS. Fire and Frost weights.
- **MR46, the client move to 70888 (issue #21):**
  - Read the overload rows, Nature's Swiftness' mask and the Arcane Shot and Summon Hawk costs from the generated tables again. Their tests stay.
  - The server-side values (Stormstrike, Mental Dexterity, Maelstrom Weapon, the imbue and totem rule, Arcane Concentration's split, the hawk's swing) stay hand values until a client or push row states them. Drop the conformance server-hotfix entry then.

## For the round-30 lane

- **Rank rows on moved levels:** the Arcane Shot rows are gone and nothing replaces them. Elemental's Lightning Bolt rows move with the overload (MR43).
- **Arcane Precision timing:** Arcane's leveling numbers moved (bare 20 -13%), so re-measure the Precision question on the new engine.
- **Untouched by any hotfix:** the sub-60 tank boss, Master Summoner and the bear's Maul gate. Natural Instinct was already right.

## Sim changes in this landing

- **Fork `lane/hotfix-1111`, fast-forward of master 278876977:**
  - d848dc8f4: tests, red on the box.
  - 871c9a0ab to 0882e1a80: the mage test's harness.
  - a8a54522e: the fixes.
  - 1c2c9e3b2 and 856efc5f4: leftovers of the deleted reach rule.
  - 16823e9c7: older tests and suite results.
  - 712d8fcc3 and 7d0e129ad: the Enhancement rotation synced.
  - 589e742fb and 4dcf0e8b7: conformance.
- **Site pin:** 7d0e129ad.
- **Shared code:** `core/buffs.go` no longer checks an imbue before applying a weapon totem. Only shamans wear Windfury or Flametongue Weapon, so no other class moves.

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 4 to 6).
- Blizzard, Class Tuning Incoming, https://us.forums.blizzard.com/en/wow/t/2404092.
- Blizzard, development notes, https://us.forums.blizzard.com/en/wow/t/2360696 (posts 9 and 10).
- wago.tools DB2 CSVs for 1.60.1.70790 and 1.60.1.70888, hotfixed and shipped; wago.tools/hotfixes, pushes 113380, 113702, 113766 and 1137xx to 1139xx.
- Wowhead Forever talent payload and spell tooltips (nether.wowhead.com/forever).
- foreverlogs.gg reports 2845, 2847, 2848 and 2849.
