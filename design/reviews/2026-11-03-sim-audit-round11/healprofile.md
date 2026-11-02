# Sim audit round 11: heal profile pulls

Lane: wt/r11-healprofile (site, from origin/main 8f047a62f) and lane/r11-healprofile (fork). The fork branch is lane/r11-healers fast-forwarded (89da891cc) and has no commits of its own. At the healers lane's request, the site branch merges wt/r11-healers (dd5840a6b). Every final rank row, band build and coverage figure below was measured on that merged state. A row that moved because heals learned below 20 now keep their whole coefficient is tagged "R11 healers 074edfadf".

Scope:
- Owner ruling 1 (round 11): a pull-size multiplier on the sub-60 heal profile.
- The sub-60 healer rows of data/curated/rank-choices.json.
- The sub-60 healer band builds (data/curated/ladder-first.json).
- Round 10's late Restoration druid band-20 find.
- A note for the owner on coverage against HPS.

Owner CPU and BiS rulings (2026-11-02, relayed by the healers lane): every run went one at a time, under `nice -n 15` with GOMAXPROCS=9. No BiS ranking ran, local or CI. The pick check is "measure at refresh". In its place there is a targeted gear-sensitivity check, below.

## The rule (bands.rule in data/curated/heal-profile.json; applied only by sim/request/healprofile_bands.go)

**For a Classic player:**
- Most of a dungeon run is trash, and a hard trash pull often hits harder than the last boss. So below 60 each band now fights the harder of two things: its median last boss (round 10's figure) and its hard trash pull.
- The hard trash pull is the pull that a quarter of the run's mobs are killed in or in something harder. That is the 75th percentile, weighted by mobs, because a run spends its time killing mobs.
- **What counts as trash:** every creature the band's dungeons spawn that is hostile to players and can be fought. Bosses, vendors, critters and hidden triggers are left out.
- **How trash is cut into pulls** (the way the server pulls):
  - A pulled creature calls every creature of its own faction within its assistance radius, and each of those calls the next.
  - Spawn groups that aggro together, and aggro-linked creatures, come with them.
  - The radius is the creature's own CallForHelp where its template states one, else 10 yards.
- **A pull's damage** is the sum of its mobs' damage, each counted the way round 10 counts a boss: melee less the time spent casting, plus repeating in-combat spells, once a cast on one target.
- **The multiplier** is that pull's damage over the median last boss's, never below 1. It scales the tank hit and the raid pulse together.
- Level 60 is unchanged (it is a raid; no pulls, multiplier 1).

**Choices and their reasons:**
- **75th percentile, weighted by mobs.** The median pull per pull is a single mob at every band, because lone pats and singles dominate the count. Weighting by mobs gives the pull a typical mob is killed in. A healer gears for the hard pulls, and the 75th percentile is "a quarter of the run's mobs come in pulls at least this hard". The median (weighted by mobs) is stated in every row too.
- **10 yards.** TrinityCore's CreatureFamilyAssistanceRadius default is 10. cmangos ships 5, which pulls nine in ten Deadmines and Razorfen Kraul mobs alone, and no Classic player accepts that. Sensitivity is in the table below.
- **Bosses are left out of the trash.** Any pull holding an instance_encounters credit entry is a boss pull, and round 10's boss figure already covers it.
- **Tank against group:** trash melee and aimed spells are the tank's, because the tank holds the pull. Spells at a random or non-top target, or at an area, are the group's. That puts 96 to 99% of trash damage on the tank. The tank-to-raid ratio is kept as at 60 (see "Still open").

| band | dungeons (trash mobs / pulls, mob-weighted median pull size) | median pull (mobs, DPS) | 75th-percentile pull, the one fought | pull DPS | median last boss DPS | multiplier | tank / pulse at 1.85 (was) |
|---|---|---|---|---|---|---|---|
| 20 | Deadmines 184/138, 1; Wailing Caverns 234/141, 2 | 2, 37.3 | Wailing Caverns guids 4300133, 4300202, 4300203: Druid of the Fang + 2 Deviate Lasher | 82.2 | 51.0 | 1.61 | 143 / 37 (89 / 23) |
| 30 | RFK 179/136, 1; Gnomeregan 380/168, 3; SM Graveyard 84/38, 3 | 2, 92.6 | Gnomeregan 9000201, 9000302 to 9000305: Mechanized Sentry + 4 Leprous Assistant | 151.2 | 73.3 | 2.06 | 263 / 69 (128 / 33) |
| 40 | SM Armory 66/29, 3; RFD 197/52, 4; SM Cathedral 94/56, 2 | 3, 189.7 | RFD 1290160, 1290165 to 1290167, 1290231, 1290244: Splinterbone Skeleton, 3 Splinterbone Warrior, Thorn Eater Ghoul, Frozen Soul | 270.1 | 88.7 | 3.05 | 470 / 123 (154 / 40) |
| 50 | Maraudon 499/201, 4; Sunken Temple 186/80, 3 | 4, 382.3 | Maraudon 3490164, 3490165, 3490247, 3490276 to 3490278: 2 Putridus Trickster, Poison Sprite, 3 Corruptor | 580.6 | 314.3 | 1.85 | 1,010 / 263 (547 / 143) |
| 60 | Onyxia's Lair (raid) | - | - | - | 1,468.1 | 1 | 2,128 / 555 (unchanged) |

Each row states every mob of its pull: creature_template MinMeleeDmg, MaxMeleeDmg, DamageMultiplier, MeleeAttackPower and MeleeBaseAttackTime, plus each spell with its creature_ai_scripts or creature_spell_list row. For example, the Druid of the Fang's Lightning Bolt is creature_ai_scripts 384006, 75 a cast every 3.1 s at the tank, and it spends nearly all its time casting. The Corruptor's Corruption is creature_ai_scripts 1221702, 160 every 17.5 s at a random target, so it hits the group.

Loading refuses three things:
- a pull whose mobs do not sum to its stated damage (to 0.05 a mob)
- a multiplier that is not the rule's
- a stated figure that does not follow from the rule

**Sensitivity (multiplier against the median boss, by assistance radius and percentile, both weighted by mobs):**

| radius | 20: q50 / q75 / q90 | 30 | 40 | 50 |
|---|---|---|---|---|
| 5 yd (cmangos) | 0.56 / 0.63 / 1.07 | 0.66 / 1.12 / 1.87 | 0.86 / 1.50 / 1.91 | 0.52 / 1.01 / 2.01 |
| 10 yd (chosen) | 0.73 / **1.61** / 2.73 | 1.26 / **2.06** / 2.98 | 2.14 / **3.05** / 4.60 | 1.22 / **1.85** / 2.44 |
| 15 yd | 1.82 / 3.29 / 8.26 | 1.77 / 2.45 / 3.57 | 2.37 / 3.77 / 7.44 | 1.35 / 2.08 / 2.60 |

At 10 yards and the 75th percentile, the multipliers land close to what round 10's sweep said a level-60-like stress would need (about x2, x4, x3 and x2). They were derived from spawns, not tuned to that sweep.

**Sources (checked 2026-11-02):**
- cmangos classic-db at ec4f596: `creature`, `creature_spawn_entry`, `spawn_group`, `spawn_group_spawn`, `spawn_group_entry`, `creature_linking`, `creature_linking_template`, `creature_template`, `creature_spell_list`, `creature_ai_scripts`, `instance_encounters`, `spell_template`.
- mangos-classic ff13624:
  - `Creature::MarkCallAssistanceOnPull` and `CanAssistTo` (Creature.cpp 2178-2270): own faction, within CallForHelp, else the config radius.
  - `Unit::TriggerAggroLinkingEvent` (Unit.cpp 559-583): creature links and spawn groups with CREATURE_GROUP_AGGRO_TOGETHER come too.
  - `mangosd.conf.dist.in` 963-970: CreatureFamilyAssistanceRadius = 5.
- TrinityCore 3.3.5 `worldserver.conf.dist` 1678-1684: CreatureFamilyAssistanceRadius = 10.
- Client 1.60.1.70707:
  - FactionTemplate decides which creatures are hostile.
  - Creature and CreatureDifficulty hold 184 creatures and none of these dungeons' trash, so the client gives no trash rows.
  - The x1.2 dungeon modifier (round 10's ContentTuningXExpected) applies to trash and boss alike, so it cancels in the multiplier.
- The Scarlet Monastery's four wings share map 189 and are told apart by position.
- Tools: scratch `r11/healprofile/tools/{extract,pulls,trash,bands,derive,apply}.py`, output `pulls.json`.

## Findings

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Sub-60 bands fought one boss, not a dungeon.** Round 10's band figure is the median last boss, so bare healers covered 72 to 98% below 60 against 29 to 42% at 60, and gear tied. A 75th-percentile trash pull is 3 to 6 mobs dealing 1.6 to 3.1 times the boss. | Table above; bare coverage tables below. Tests first: `TestHealBandPullMultiplier`, `TestHealBandPullDamageIsItsMobs`, `TestCuratedHealBandsStateTheirPulls`, `TestHealProfileRefusesAPullThatIsNotItsMobs`, and new figures in `TestHealProfileScalesEachBandByItsContent` and the Attach and AtScale tests. They failed (no pull_rule), then passed. | CONFIRMED-BUG (owner ruling 1) | R10 S 64a1c9717 (the boss-only rule; round 10 listed pulls as open) | S eaa2b2285 |
| P2 | **The BiS page's profile disclosure would say nothing of the pull.** A sub-60 healer band whose row states a pull now shows "Hard pull: 3 mobs in Wailing Caverns (...), 1.61 times the last boss". Files published before this are unchanged. Vitest written first. | `web/src/lib/bis/heal-view.ts`, `copy.ts`, `types.ts` | CONFIRMED (page follows P1) | R10 S 8ad647829 | S 213beda2a |
| P3 | **Round 10's band mechanism, re-read.** Checked: `bandFor` (below 20 goes to band 20; between bands goes to the band below; 60 and above goes to the reference); `ForLevel`/`AtScale` composition; `validateBands` against the tuned scale; `healerLevel` (first player); `AttachWeights` (player level); heal-view's `bandRowFor`. No bug found. One measuring hazard, not a bug: spec-breakdown takes a healer's band talents from the published BiS entry, which is stale while the bis workflow is held. For druid, Discipline and Holy it is not the current ladder, so every figure here passes the ladder's build with `-talents`. | sim/cmd/spec-breakdown/main.go:238 | Aligned | - | - |

## Sub-60 healer rank rows (rank-choices.json; S 7f4387eeb, S 456f3475f)

**Method:**
- Each band's ladder build.
- Greedy search over every learned rank of the rotation's ranked heals (Lightwell and Dark Sacrifice added for the priests), 3,000 iterations.
- Then default, current, the greedy set, each row alone and the set less each row: 10,000 iterations, seeds 1 and 50001, alliance and horde.
- A row stays if it beats the default rule by more than 0.3% at both seeds, alone and inside the set.
- No row may turn a fight that runs out of mana into one that ends with mana unspent.
- Rows were checked again after the band builds changed (the shaman 20 row fell to that check).

| spec | levels | was | now | gain vs default, seed 1 / 50001 (horde) | origin |
|---|---|---|---|---|---|
| druid-restoration | 50-59 | Rejuvenation 8 (now -5.4%) | Rejuvenation 1 | +10.96% / +11.10% (+10.3% / +10.6%) | R10 S 90d16b4f0; R11 healers 074edfadf |
| druid-restoration | 20-49 | - | none: Rejuvenation 2, 2 and 1 gain 3 to 14% but each ends the fight with mana unspent (443, 442, 879 s) | - | - |
| paladin-holy | 20-29, 30-39, 40-49 | - | Holy Light 1 at each | 20: +15.15%; 40: +5.24% (horde +15%, +0.6%) | R11 healers 074edfadf |
| paladin-holy | 30-39 | - | + Flash of Light 1 | set +8.86% / +8.86% on the final build (horde +9.3%) | R10 S 90d16b4f0 |
| paladin-holy | 50-59 | - | none: Holy Light 2 gains 22% but ends with mana unspent (1,394 s) | - | - |
| priest-discipline | 20-29 | - | Shield 2, Renew 1 | set +10.06% / +10.06% (horde +5.6%) | R11 healers 074edfadf |
| priest-discipline | 30-39 | - | Shield 5, Renew 1 | +4.79% / +4.79% on the final build (horde +5.5%) | R11 healers 074edfadf |
| priest-discipline | 40-49 | - | Heal 4, Renew 1, Prayer of Healing 1, Penance 1 | set +9.48% / +9.55% (horde +9.3%) | R10 S 90d16b4f0; R11 healers 074edfadf |
| priest-discipline | 50-59 | Shield 8, Prayer of Healing 1, Binding Heal 1 (+6.3%) | Heal 4, Renew 1, Prayer of Healing 1 | +12.12% / +12.18% (horde +10.6%) | R10 S 90d16b4f0; R11 healers 074edfadf |
| priest-holy | 20-29, 30-39 | - | Renew 1 | +3.2% / +3.5%; +8.3% / +8.1% (horde +2.7%, +7.8%) | R11 healers 074edfadf |
| priest-holy | 40-49 | Renew 5 (now -1.9%) | Renew 1, Flash Heal 4, Prayer of Healing 1 | set +13.22% / +13.13% on the final build (horde +13.7%) | R10 S 90d16b4f0; R11 healers 074edfadf |
| priest-holy | 50-59 | Flash Heal 6, Prayer of Healing 1 (+9.1%) | + Heal 1, + Renew 1 | set +18.3% (horde +16.3%) | R10 S 90d16b4f0; R11 healers 074edfadf |
| shaman-restoration | 40-49, 50-59 | - | Healing Wave 2 | 40: +33.15% / +33.15% (horde +33.6%); 50: +23.50% / +23.46% (horde +23.7%) | R11 healers 074edfadf |
| shaman-restoration | 20-29, 30-39 | - | none: Healing Wave 1 at 20 turns 204 s into 313 s on the final build; Healing Wave 2 at 30 ends with mana unspent (702 s) | - | - |

Level-60 rows are untouched.

Why the large gains: below-20 ranks now carry their whole coefficient (E 074edfadf), and the pull profile runs every healer dry, so cheap low ranks win wherever mana binds. Rotations may downrank when mana-bound (round 9 law). A wow-player read of "Healing Wave 2 at 40" is worth having at integration.

## Sub-60 healer band builds (ladder-first.json; S d450cf8ec)

**Method:**
- Hill climb (one-point moves inside the guide's ranks, 2,000 iterations, 1% a step) from each band's ladder build, with the new rank rows.
- Each move confirmed at 10,000 iterations, seeds 1 and 50001, both factions.
- A move is adopted only if all of these hold:
  - it can be written as a ladder-first reservation that passes `TestCuratedLadderFirstIsSound` (legal and never taking back a point at any level 10 to 60)
  - no band loses
  - no unmodeled talent is dropped
  - nothing the rotation casts is dropped

| spec | reservation (first) | band build change | gain seed 1 / 50001 (horde) | origin |
|---|---|---|---|---|
| Restoration druid | Heart of the Wild 5, Reflection **3**, Bloom 1, Nature's Swiftness 1 | 20: Nature's Focus 3 to 2, Reflection 2 to 3 | +8.50% / +8.50% (+8.62%) | R10 S a09dd3b99 |
| Holy paladin | Oath of Light 1 | 30: Infusion of Light 2 to 1, Oath of Light 0 to 1 | +10.43% / +10.55% (+10.7%) | old |
| Discipline | Meditation 1, **Penance 1** | 30: Soul Warding 1 to 0, Penance 0 to 1 | +14.29% / +14.29% (+12.4% / +12.5%) | R10 S a09dd3b99 |
| Holy priest | **Spiritual Healing 3**, Prayer of Mending 1 | 40: Spiritual Guidance 4 to 1, Spiritual Healing 0 to 3 | +5.12% / +5.28% (+5.0% / +4.8%) | R10 S a09dd3b99 |
| Restoration shaman | Mindfulness 2 | 20: Tidal Focus 5 to 4, Mindfulness 1 to 2 | +9.8% with the old row; +4.9% / +4.9% (+4.7%) with no row | old |

Not adopted:
- **Druid 30 and 40** (Improved Rejuvenation 2 or Living Spirit 2 over Nature's Focus, +2.3% and +3.2%): no sound reservation, because the build takes back a Nature's Focus point between 29 and 30.
- **Druid 50**, Genesis over Improved Tranquility (+2.5%): it moves a point into another tree.
- **Holy 30** (+13.6%): it drops Inspiration, which a level-60 probe shows the engine does not model.
- **Shaman 30 and 40** (Healing Way, Restorative Totems; +7% and +1 to 2%): every reservation that reaches them drops Mana Tide Totem at 40, and the rotation casts it.
- **No move at all:** paladin 20, 40 and 50; Discipline 20, 40 and 50; Holy 20 and 50; shaman 50.

`addon-builds.json` follows, as its committed-copy test requires.

## Round 10's late find: Restoration druid band 20, "5 Furor points against 5051"

The ranker's "Furor" was the Restoration tree's first digits (Nature's Focus 5, Heart of the Wild 5, Naturalist 1), not the Feral talent Furor. Measured on the merged ladder and the pull profile, 10,000 iterations:

| band-20 build | vs current ladder (seed 1 / 50001, horde) |
|---|---|
| current ladder before this round (NF 3, HotW 5, Reflection 2, Bloom 1) | baseline |
| "5051" (NF 5, Naturalist 5, Reflection 1) | -24.03% / -23.84% (-25.7%) |
| the old plain walk "5510" | -26.62% / -26.62% (-28.4%) |
| adopted (NF 2, HotW 5, Reflection 3, Bloom 1) | +8.50% / +8.50% (+8.6%) |

5051 is not adopted. It loses beyond error and drops Bloom and Heart of the Wild.

## Gear sensitivity (BiS pick check: measure at refresh)

Small spec-breakdown runs, bare, alliance, 3,000 iterations, the band's published set against one swapped item. Before is the round-10 profile; after is this lane's state. The figure is the change of effective HPS from the swap at seeds 1 / 50001.

| spec, band | swap | before | after |
|---|---|---|---|
| Holy priest 30 | head: Holy Shroud to Pristine Circlet | +0.19% / +0.19% | -4.07% / -3.92% |
| Holy priest 30 | feet: Gilded Slippers to Frothing Slippers | +0.19% / 0 | -4.54% / -4.55% |
| Holy priest 30 | neck: to Crystal Starfire Medallion | 0 / 0 | -0.16% / -0.16% |
| Restoration shaman 40 | main hand: Staff of Jordan to Mograine's Might | -0.15% / -0.15% | -14.9% / -14.9% |
| Restoration shaman 40 | head: Papal Fez to Whitemane's Chapeau | 0 / 0 | -0.82% / -0.77% |
| Restoration shaman 40 | back: to Blackforge Cape (a published tie) | 0 / 0 | 0 / +0.05% |
| Restoration druid 20 | wrist: to Drakewing Bands | -0.80% / -0.80% | -1.19% / -1.36% |
| Restoration druid 20 | shoulder: to Serpent's Shoulders (a published tie) | -0.27% / -0.27% | **+0.85% / +0.85%** |
| Restoration druid 20 | back: to Heavy Woolen Cloak (a published tie) | 0 / 0 | 0 / 0 |

- Before, swaps moved effective HPS by 0 to 0.2%: the gear tied.
- After, healing-power swaps separate by 1 to 15%.
- Two published ties stay ties: those items differ in stats that do not heal.
- Restoration druid 20's shoulder alternative now beats the published pick, so that pick will likely change at refresh.
- **The full pick check (every slot, every band, five healers) is "measure at refresh".**

## Bare coverage per band (effective HPS over incoming, seed 1)

**Profile alone** (base engine, published talents; ba1, both seeds agree to 0.4%):

| band | druid | paladin | Discipline | Holy | shaman |
|---|---|---|---|---|---|
| 20 | 79% to 60% | 81% to 50% | 74% to 47% | 85% to 56% | 94% to 77% |
| 30 | 93% to 82% | 92% to 75% | 93% to 53% | 92% to 55% | 93% to 74% |
| 40 | 89% to 92% | 89% to 70% | 93% to 59% | 98% to 57% | 92% to 68% |
| 50 | 95% to 60% | 72% to 42% | 77% to 44% | 87% to 40% | 98% to 54% |
| 60 | 42% | 32% | 29% | 33% | 38% (unchanged) |

**Final state.** Before is origin/main 8f047a62f with round 10's band builds. After is this lane: the pull profile, the healers engine, the new rows and the band builds. Alliance; horde within 8 points:

| band | druid | paladin | Discipline | Holy | shaman |
|---|---|---|---|---|---|
| 20 | 92% to 90% | 81% to 58% | 77% to 55% | 85% to 59% | 94% to 82% |
| 30 | 92% to 81% | 92% to 88% | 93% to 69% | 92% to 53% | 93% to 74% |
| 40 | 90% to 89% | 89% to 73% | 93% to 67% | 98% to 65% | 92% to 90% |
| 50 | 95% to 71% | 72% to 42% | 79% to 47% | 87% to 43% | 98% to 67% |

- Bare healers are now stressed at most bands: 42 to 90%, and mana runs out before 300 s in every cell but druid 40, which lasts 290 s.
- Druid at 20 and 40, and shaman at 40, still cover about 90%. Those specs heal efficiently enough (HoTs and Bloom; Healing Wave 2) that the hard pull does not tax them as much.

## Before/after (effective HPS, bare, 10,000 iterations)

Before: $R/sb-before from site-before (round-10 band builds passed with `-talents`). After: this lane's merged state (ladder-final builds). Relative % at seeds 1 / 50001, absolute at seed 1. Higher numbers here mostly mean a harder fight (more to heal), not a stronger healer. The coverage tables are the comparable figure.

| spec | faction | 20 | 30 | 40 | 50 |
|---|---|---|---|---|---|
| druid-restoration | alliance | 37.5 to 58.7, +56.5% / +56.5% | 54.0 to 97.9, +81.3% / +81.1% | 63.3 to 192.5, +204.1% / +204.3% | 239.1 to 329.0, +37.6% / +37.8% |
| druid-restoration | horde | 37.0 to 56.7, +53.2% / +53.2% | 53.9 to 96.0, +78.1% / +78.1% | 63.2 to 192.7, +204.9% / +205.1% | 239.0 to 323.5, +35.4% / +35.7% |
| paladin-holy | alliance | 33.1 to 38.0, +14.8% / +14.8% | 53.9 to 106.9, +98.3% / +98.3% | 63.2 to 158.6, +151.0% / +151.0% | 180.1 to 196.2, +8.9% / +8.9% |
| paladin-holy | horde | 32.2 to 36.9, +14.6% / +14.6% | 53.9 to 103.5, +92.0% / +91.8% | 63.5 to 168.2, +164.9% / +164.9% | 190.5 to 209.4, +9.9% / +9.8% |
| priest-discipline | alliance | 31.4 to 36.1, +15.0% / +15.0% | 54.3 to 83.2, +53.2% / +53.2% | 66.0 to 144.4, +118.8% / +118.9% | 197.7 to 217.3, +9.9% / +9.9% |
| priest-discipline | horde | 30.2 to 32.3, +7.0% / +7.0% | 54.5 to 87.1, +59.8% / +60.3% | 66.1 to 147.9, +123.8% / +123.8% | 195.6 to 218.2, +11.6% / +11.6% |
| priest-holy | alliance | 34.6 to 38.7, +11.9% / +11.8% | 53.8 to 63.9, +18.8% / +18.6% | 69.1 to 139.6, +102.0% / +102.0% | 218.5 to 198.0, -9.4% / -9.4% |
| priest-holy | horde | 32.7 to 33.9, +3.7% / +3.7% | 53.8 to 60.5, +12.5% / +12.3% | 69.0 to 135.4, +96.2% / +96.2% | 211.1 to 191.8, -9.1% / -9.1% |
| shaman-restoration | alliance | 38.5 to 54.0, +40.3% / +40.3% | 54.2 to 89.7, +65.5% / +65.7% | 65.1 to 194.4, +198.6% / +198.6% | 245.5 to 309.5, +26.1% / +26.0% |
| shaman-restoration | horde | 38.3 to 48.8, +27.4% / +27.4% | 54.2 to 89.0, +64.2% / +64.4% | 65.1 to 193.5, +197.2% / +197.2% | 245.3 to 308.8, +25.9% / +25.9% |

- Holy at 50 heals less on the harder fight: mana runs out at about 100 s, against 266 s before. The rank search's best set is in place, and the new set beats the default rule by 18%.
- Level 60: +0.00% for every healer on the profile alone (ba1). Restoration shaman 60 moves 38% to 42% from the healers engine, not from this lane.
- **DPS specs are byte-identical.** Effective DPS was equal to the last bit before and after for rogue-combat 40, mage-fire 20, paladin-retribution 30, shaman-enhancement 50, priest-shadow 50, druid-balance 40, warrior-fury 20, hunter-marksmanship 30, warlock-affliction 40 and paladin-protection 50, with the site-only build and the merged build. spec-breakdown's `-json` key order is not stable run to run, so DPS was compared, not file hashes.

## Note for the owner (not decided): show sub-60 healers as coverage?

- **What HPS means below 60.** It is the band's fight, so a band-30 figure is "healing a Gnomeregan leper pack". It rises when the fight is harder. HPS against HPS across bands, or against 60, compares fights, not healers.
- **What coverage would show.** Coverage (effective HPS over the band's incoming damage) is a share that reads the same at every band: 42 to 90% now, against 29 to 42% at 60.
- **Coverage's weakness.** It caps at 100%. Near the cap gear stops mattering, which is why round 10's 72 to 98% tied. On the pull profile, three cells are still at 89 to 90% (druid 20 and 40, shaman 40). There, a coverage view would show small gear gains as almost nothing, even though HPS still separates them.
- **Option:** show coverage as the headline below 60 ("covers 58% of a hard Wailing Caverns pull"), with HPS as the line beneath. Keep HPS as the score the ranker sorts by, since it does not saturate.

## For other lanes

- **Integrator:**
  - sim/request TestRotationLadder goldens move for druid-restoration, paladin-holy, priest-discipline, priest-holy and shaman-restoration (pull profile, rank rows, band builds). I did not regenerate them. No other test was deselected.
  - My site branch merges wt/r11-healers. My fork branch equals lane/r11-healers (89da891cc).
  - Nothing here moves a non-healer spec.
- **UX / wow-player:** the BiS page's new "Hard pull" disclosure row (S 213beda2a) is player-facing and wants a screenshot review when BiS next publishes. Each P1, rank row and build row is written for a Classic player's read.
- **Healers lane:** the below-20 coefficient fix makes rank 1 and 2 heals the best choice wherever mana binds (Healing Wave 2 at 40 and 50, Holy Light 1 at 20 to 49, Renew 1 everywhere below 60). If that fix is reverted, those rows (tagged "R11 healers 074edfadf") need re-measuring.
- **Ranker / BiS refresh:**
  - Measure every healer pick at refresh.
  - Restoration druid 20's shoulder alternative now beats its pick by 0.9%.
  - spec-breakdown and the published band talents are stale for druid, Discipline and Holy until refresh. Measure with `-talents` from the ladder.

## Still open

- **Tank against group** (owner): trash puts 96 to 99% of its damage on the tank, but the profile keeps level 60's tank-to-raid ratio, so the pulse is scaled with the pull. A pull-derived split would make sub-60 fights almost tank-only and devalue group heals.
- **Line of sight** is not checked, so a pull can join through a wall. And 10 yards is one emulator's default (see sensitivity).
- **Players' own pulls are not modelled:** chain pulls, pats walking in and crowd control.
- **Unexpressible band moves:** druid 30, 40 and 50, Holy 30, and shaman 30 and 40 (above). They need either a reservation that can drop a top-row point or a guide change.
