# Round 13, lane hotfix: the 1.60.1.70753 server hotfix push (2026-11-03)

A new server hotfix push for 1.60.1.70753 went live on wago.tools on 2026-11-03 at 14:14 UTC. The committed `data/builds/1.60.1.70753/hotfix-push.json` (fetched 11-02 22:28) has zero entries. This lane layered the push with `pipeline.hotfix_push.layer_hotfix_push`, copying the trainables lane's layered raw/ into `$R/hotfix/root/1.60.1.70753`. `check_push` against wago at 09:3x UTC returned `moved=false`, so the copy matches what wago lists now. The lane then diffed every pushed table against `raw/shipped/`, ran `simconst` and `spellranks` over the layered raw/ into scratch, and regenerated the fork's hunter constants from that output.

Wowhead's Forever tooltips (nether.wowhead.com, read 2026-11-03) still show the shipped text: Aimed Shot 20904 "+166", Arcane Shot 14287 "217", Trueshot Aura 20906 "50". The push rows are the only source for these numbers. They are client hotfix rows, so they count as one source (flagged).

Scratch evidence (absolute paths):
- `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r13/hotfix/pushdiff.txt`: every changed, added and removed row, with spell names.
- `.../r13/hotfix/root/1.60.1.70753/spellconst/hunter.json`: simconst over the layered raw/. This is the only class file that differs from the committed spellconst. `spellranks.json` is byte-identical.

## Hunter

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The push buffs the hunter's shots, and the engine still had the shipped numbers.** Aimed Shot went 20/34/55/89/125/166 → 40/70/115/185/260/350. Arcane Shot went 20/26/39/65/94/134/170/217 → 25/35/50/85/120/175/220/280. Serpent Sting per tick went 2/6/12/22/34/48/64/83/111 → 3/7/14/26/40/58/75/100/135; at 60 the trainer rank is rank 8 (13555, 100), and rank 9 (25295, 135) is AQ. Sniper Shot went 160/225/295 → 320/450/590. Aspect of the Hawk ranks 1-6 went 15/22/30/38/47/55 → 20/35/50/70/90/110; rank 7 (25296) stays 120 (AQ), so a 60 now runs rank 6 at 110. | SpellEffect EffectBasePointsF rows, e.g. 697764 (20904) 166→350, 691069 (14287) 217→280, 690655 (13555) 83→100, 702403 (25295) 111→135, 692138 (14322) 55→110, 1340619 (1310786) 295→590. Engine: `sim/hunter/client_damage.go` (hand tables), `constants_auto_gen.go`. | CONFIRMED-BUG | new-data | E a3b2af6d4 |
| H2 | **Trueshot Aura doubles at 60, 50 → 100 ranged attack power** (ranks 15/22/30/40/50 → 20/35/50/75/100). Every rank is still one effect: Effect 65 (party area aura), aura 124, ranged attack power. There is no melee row, so the ranged-only aura is right and melee is unaffected. The physical lane asked about this and got that answer. | SpellEffect 1323534, 1323538, 697417, 698474, 697683. `sim/core/buff_ranks.go` TrueshotAuraRanks. | CONFIRMED-BUG | new-data | E a3b2af6d4 |
| H3 | **Aimed Shot now beats Multi-Shot on their shared cooldown when the cast fits.** The guides said to cast Aimed Shot only as the opener. With the push it averages more per hit than Multi-Shot. The new line casts Aimed Shot when the shared cooldown is up and the next auto shot is more than 1.2 s away, and casts Multi-Shot otherwise. Thresholds 0.5/0.8/1.0/1.2/1.5 s were tried; 1.2 is best or tied at both seeds and both factions. "Aimed Shot only, no Multi-Shot" loses 1.9-2.0%. On the shipped values the same line lost 0.8-1.0%, so the push is what flips it. | spec-breakdown at 10000 iterations: BM horde 661.6/661.9 → 663.5/663.4, BM alliance 657.5/658.2 → 658.7/659.3 (seeds 1/50001, error ±0.2). MM gains +0.5% at 3000 iterations on top of H4. | CONFIRMED-BUG | new-data | S 35534af90, E bcb7c2912 |
| H4 | **Marksmanship never cast Serpent Sting.** The line was dropped by the 2026-10-08 rotation search (site fb976bcd2). On this engine, adding the Beast Mastery line back gains +5.5% before the push (622.8 → 657.0, seed 1, 3000 iterations) and +6.2% after it. The guide's "Serpent Sting is not worth a global in this build" was wrong. | `data/curated/apl/hunter-marksmanship.json` history (`git log -S '"spellId": 13555'`); measurements in this lane's scratch. | CONFIRMED-BUG | old (fb976bcd2; no round re-ran Marksmanship with the sting) | S 35534af90, E bcb7c2912 |
| H5 | **Marksmanship's raid build should take Lone Wolf and drop the pet.** The guide said "the cat still does more than the 20%". It does not: Forever's pet copies 10% of the hunter's attack power. The new raidBuild (`5/005355001150305/50005003`) takes Lone Wolf 1, Improved Stings 3, Surefooted 3 (through Survivalist 5) and Deadly Aspects 5, and drops Endurance Training, Focused Fire, Unleashed Fury and Ferocity. talent-search found a close variant (+3.2% seed 1, +2.9% seed 50001; Lone Wolf, Improved Stings 3, Surefooted 1 with Deflection 5, Focused Fire kept). Moving the dead Focused Fire points into Surefooted and taking Survivalist instead of Deflection (same DPS, more health) gives +5.1%: horde 702.2/701.9 → 737.7/738.4, alliance 697.5/697.3 → 732.0/732.4, error ±0.3. It keeps every unmodeled guide talent, holds the Marksmanship tree (33 points) and drops nothing the rotation casts. Before the push the same build gained +2.1%, and talent-search's own winner gained +0.8%, under the 1% bar. The leveling build keeps the pet. | talent-search reports `.../r13/hotfix/ts-mm/`, `ts-mm-50001/`, `ts-mm-before/`; `sim/request/request.go:692` (Lone Wolf → PetNone); `sim/hunter/talents.go` applyLoneWolf. | CONFIRMED-BUG | old (the build predates round 9; the "cat does more" claim is landing 1feda0686) | S daaf8314d |
| H6 | **The guide prose stated the shipped play.** It said "do not cast Aimed Shot after the opener", that Trueshot Aura is 50 at 60, that Marksmanship plays with a pet in raids, and that Serpent Sting is not worth a global for Marksmanship. Updated in the hunter class page, Beast Mastery and Marksmanship. | `web/src/content/guides/hunter/{index,beast-mastery,marksmanship}.md`; `_prose.test.ts` and `builds.test.ts` pass (399 tests). | CONFIRMED-BUG | new-data | S daaf8314d |
| H7 | Lacerating Strikes (1310536): SpellMisc Attributes_4 goes 256 → 1048832, which adds bit 20 (0x00100000). In TrinityCore master that bit is SPELL_ATTR4_AURA_NEVER_BOUNCES (aura reflection and bounce); in 3.3.5 it was NOT_CHECK_SELFCAST_POWER. Neither moves a combat number, and nothing in the engine or the pipeline reads Attributes_4. The bleed already had IGNORE_DAMAGE_TAKEN_MODIFIERS (0x100, the engine's SpellFlagIgnoreTargetModifiers) and Attributes_8 0x200 PERIODIC_CAN_CRIT. periodiccrit and eureka regenerate byte-identical. | SpellMisc 857128; TrinityCore `src/server/game/Miscellaneous/SharedDefines.h` (master, read 2026-11-03). | Aligned | - | - |
| H8 | The rank-choice rows (Aimed Shot rank 1 at bare 60; Arcane Shot rank 3 at 40-49, rank 1 at 50-59 and bare 60) were re-measured on the pushed values and the new rotation. They still win: BM bare 40 168.0 vs 163.0 and bare 60 328.2 vs 315.2 without the rows; MM bare 60 330.3 vs 316.0 (seeds 1 and 50001 agree). | an spec-breakdown built with the hunter rows removed from the embedded rank-choices | Aligned (rows hold) | - | - |
| H9 | Beast Mastery's raid build holds after the push. The best talent-search candidate gains +0.4% at seed 1, and the guide build wins at seed 50001. | `.../r13/hotfix/ts-bm-1/`, `ts-bm-50001/` | Aligned | - | - |
| H10 | Survival is melee. It casts none of the pushed spells, and its DPS is unchanged at every band. | spec-breakdown before = after | Aligned | - | - |

## Cross-cutting: the rest of the push

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Only hunter player spells change. Of the 91 spell ids the push touches, the only ones with a player class family are hunter (family 9). Fireball 21162 (family 3) and Benediction 23236 (family 6) appear only as new ItemEffect rows. | SpellClassOptions over the pushed ids | Aligned | - | - |
| X2 | NPC rows: Super Hogger's event (Meathook, Ground Slam, Call Gnolls, Bloodlust 1325968, Ultra/Giga/Cosmic Hogger), Hateful Strike 1325876, Void Bolt, Warstomp, Cleave and Shadow Bolt Volley 1327241-1327251, and the new Manastorm, Chill, Summon Infernal, Hurricane and Chain Lightning ids. None is a sim target, a pet ability or a referenced id (grep of the fork's sim/ and the site's sim/ and curated). | pushdiff.txt | Aligned (no sim effect) | - | - |
| X3 | Consumables: a new "Endless" set (Intellect +2% x5, Power +3% damage x5, Healing, Rapidity 4% haste x5, Opportunity 4% crit x5, Mana regen 30%, Fortitude scroll, Speed, Growth, Experience) and a "Ding!" aura. Their ItemEffect rows (239511-239689) are linked to no item: ItemXItemEffect is unchanged, and no ItemSparse row names them. They look like event-only buffs, so nothing for the sim to use yet. Noggenfogger gets an ItemEffect row (no combat effect). | ItemEffect, SpellEffect, SpellAuraOptions rows in pushdiff.txt | OPEN (settles when an item carrying them ships) | - | - |
| X4 | ItemEffect adds Thunderfury 21992, Fireball 21162, Immolation 21142, Benediction 23236, Sanctuary 22850 and Call Anathema/Benediction rows, none linked to an item. The existing items keep their shipped effects. | ItemXItemEffect unchanged | Aligned (no sim effect yet) | - | - |
| X5 | 21 CurvePoint rows (curves 82803, 94906, 102986, 102987, 109548) change. No generator, engine file or pipeline reads CurvePoint. | grep of data/pipeline and the fork tools | Aligned | - | - |

## Before/after (raid 60, 10000 iterations, horde and alliance)

Before is $R/before.tsv (int/r12). "Engine" is the push alone (E a3b2af6d4, old rotation). "Final" is the push plus the new rotations (published talents). "MM raid build" is the final engine and rotation plus the H5 raidBuild through `-talents`; the published entry moves at the ranker refresh.

| spec | faction | seed | before | engine | final | final vs before | MM raid build |
|---|---|---|---|---|---|---|---|
| hunter-beast-mastery | horde | 1 | 623.5 | 661.6 | 663.5 | +6.42% | |
| hunter-beast-mastery | horde | 50001 | 623.7 | 661.9 | 663.4 | +6.37% | |
| hunter-beast-mastery | alliance | 1 | 619.4 | 657.5 | 658.7 | +6.34% | |
| hunter-beast-mastery | alliance | 50001 | 620.1 | 658.2 | 659.3 | +6.32% | |
| hunter-marksmanship | horde | 1 | 622.5 | 656.6 | 701.6 | +12.71% | 737.7 (+18.5%) |
| hunter-marksmanship | horde | 50001 | 622.3 | 656.4 | 701.7 | +12.76% | 738.4 (+18.7%) |
| hunter-marksmanship | alliance | 1 | 617.3 | 651.3 | 697.0 | +12.91% | 732.0 (+18.6%) |
| hunter-marksmanship | alliance | 50001 | 617.3 | 651.4 | 697.4 | +12.98% | 732.4 (+18.6%) |
| hunter-survival | both | both | 741.1 / 744.7 | same | same | 0.00% | |

Bare bands (horde, 10000 iterations; before → after at seed 1, and seed 50001 agrees within 0.1%):
- BM: 20: +2.4%, 30: +2.7%, 40: +6.2%, 50: +2.8%, 60: +8.6%.
- MM: 20: +4.2%, 30: +8.1%, 40: +6.8%, 50: +7.8%, 60: +12.4%.
- Survival: 0.0% everywhere.
- On the new engine the new rotation never loses to the old one at any band: BM 20 and 50 tie, BM 40 is +3.6%, BM 60 +3.8%, MM 20-60 +2.0% to +7.3%.

Absolute at seed 1, horde: BM 623.5 → 663.5, MM 622.5 → 701.6 (737.7 on the raid build), Survival 744.7 unchanged.

Captures for the wow-player review (in `design/reviews/2026-11-03-sim-audit-round13/hotfix/`):
- `hunter-*.before.txt` / `.after.txt`: spec-breakdown at seed 1, horde.
- `hunter-marksmanship.after-raidbuild.txt`: the H5 build, with Aimed Shot 1506 a hit against Multi-Shot 967, and Serpent Sting at 8% of damage.
- `marksmanship-talents.png`, `marksmanship-rotation.png`, `beast-mastery-rotation.png`, `hunter-class-page.png`: dev build at 1440. The talent strips had not hydrated in the capture. The page's priority card reads `addon-data.json`, so it shows the old list until that file is regenerated; the prose under it is the new text.

### What changed, in player terms
- Aimed Shot, Arcane Shot, Serpent Sting, Sniper Shot, Aspect of the Hawk (ranks 1-6) and Trueshot Aura now hit at the hotfix numbers. The top Aimed Shot adds 350 instead of 166, and Trueshot Aura gives 100 ranged attack power instead of 50.
- Beast Mastery and Marksmanship: when the shared Aimed Shot/Multi-Shot cooldown comes up right after an auto shot, cast Aimed Shot. If the next auto is about to fire, cast Multi-Shot.
- Marksmanship keeps Serpent Sting up.
- Marksmanship raids with no pet: Lone Wolf, Improved Stings 3, Surefooted 3. Level with the pet.

## Files to regenerate (integrator, in this order)

1. **layer-hotfix-push + spell generators (FIRST).** From `<site>/data`:
   1. `.venv/bin/python -m pipeline layer-hotfix-push --build 1.60.1.70753`. raw/ must hold the fetched CSVs; on a fresh checkout run `.venv/bin/python -m pipeline fetch --product wowf --build 1.60.1.70753`, which layers the push itself. This rewrites `data/builds/1.60.1.70753/hotfix-push.json` (the record with the 11-03 search totals).
   2. `.venv/bin/python -m pipeline simconst --build 1.60.1.70753`: only `spellconst/hunter.json` moves.
   3. `.venv/bin/python -m pipeline levels --build 1.60.1.70753`: writes spellranks.json; this lane saw it byte-identical.
   4. `.venv/bin/python -m pipeline trainables --build 1.60.1.70753` (trainables lane's call).
   5. In the fork, `make spellconst SITE_DIR=<site>`: hunter constants only; already committed as E a3b2af6d4. Also re-vendor `sim/core/testdata/conformance/client/hunter.json` (done in a3b2af6d4), then run `make periodiccrit eureka relicmods` (no change expected).
   6. `make simdb` (no item change expected; the hotfix-push record is a follow skip-hash input).
2. `make apl-sync`: writes `sim/request/apl/hunter-{beast-mastery,marksmanship}.apl.json` (left uncommitted here). The fork copies are already synced by hand in E bcb7c2912.
3. Ladder goldens (`sim/request` TestRotationLadder): hunter-beast-mastery and hunter-marksmanship move.
4. addon-data.json (the guide priority card), addon-builds (Marksmanship now has a raid build), coach starters, dps-per-point.
5. BiS: the hunter-marksmanship raid pass on the new raidBuild, and every hunter band (shot values moved).

`sim/leveling` TestTrueshotAuraRanksMatchTheClient now reads the 70753 spellconst. It fails until step 1.2 lands. It was verified green against the scratch spellconst.

## Measure at refresh
- hunter-marksmanship raid 60 on the published raidBuild (expected about +5% over the final column).
- Every hunter band's BiS after the re-rank: Aspect of the Hawk and Arcane Shot moved the value of attack power and hit.
- Hunter stat weights: Surefooted's hit in the MM raid build moves the hit cap.

## For other lanes
- **ranks:** the hunter rank-choice rows hold on the pushed values (H8). No change requested.
- **auras:** the "Endless" elixirs and scroll (X3) carry real combat auras (+3% damage x5, 4% haste x5, 4% crit x5). If an item carrying them ships, they belong in the consumables list. For now they are event-only and linked to no item.
- **guides:** the Marksmanship guide now has two builds (leveling with the pet, raid with Lone Wolf). The Survival guide's Lone Wolf wording already agrees.
- **physical:** answered: Trueshot Aura is aura 124 only, before and after the push (H2).

## Still open
- E 05b72a3b5 (Arcane Shot, Serpent Sting and Volley take no spell power) stays, pending the owner's tooltip check (ruling 5). The push changes only the flat amounts.
- Wowhead's Forever tooltips still show the shipped text. Re-check them when Wowhead re-scrapes.
- The `sim/leveling/buff_ranks_client_test.go` comment beside `auraModRangedAttackPower` says the engine grants melee too. It does not (buffs.go). The comment is stale and was left for the core/auras owner.
- `web/src/lib/sim/rotations.test.ts` fails 2 mage-frost cases in this worktree with main's generated data; not hunter-related and not investigated.
