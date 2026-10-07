# Rogue curation: engine model review and hand-curated rotations (2026-10-07)

Branches: site `rogue-curate` (this repo), fork `rogue-curate` (`/Users/jh/code/wowsims-forever`, from `forever` at 1aed40820). Measurements are `sim/cmd/rotation-search -probe-only` (800 iterations, seed 7, BiS band gear at 60, the guide's FS1 build, Instant Poison on both weapons, no raid buffs), baseline DPS ± standard error. Cast tallies are casts per iteration over the 180 s default fight.

## 1. What the model gets right (checked against the client data)

| Mechanic | Engine | Client | Verdict |
|---|---|---|---|
| Energy | 20.2 per 2.02 s, random first-tick phase (sim/core/energy.go) | 10 per second | same rate |
| Slice and Dice duration | 6 + 3 s per combo point, times 1.15/1.30/1.45 from Improved Slice and Dice (slice_and_dice.go) | 6000 ms base, rank text | match |
| Slice and Dice haste | `MultiplyMeleeSpeed(1.30)`, which calls `AutoAttacks.UpdateSwingTimers` | aura 319, +30% (rank 2) | applies to both hands' auto-attacks (pinned by `TestSliceAndDiceHastesMeleeSwingsForTheClientDuration`) |
| Eviscerate, Rupture per-CP numbers | already corrected (c45705102) | spellconst | match |
| Cold Blood | five spells, 100% crit | rank text | spells match; Mutilate was broken, see below |
| Venom | +30% poison damage, +10 points proc chance, 9..21 s | rank text | match |
| Adrenaline Rush, Blade Flurry | +100% energy tick for 15 s; +20% melee speed 15 s, 25 energy, 2 min | spellconst | match |
| Ruthlessness, Relentless Strikes, Seal Fate | 20% per rank, 20% per CP for 25 energy, 20% per rank | rank text | match |
| Poisons | `KitConsumes` gives every rogue spec Instant Poison on both weapons from level 20 (sim/leveling/kit.go), and rotation-search, the ladder and the BiS ranker all use it | n/a | poisons ARE applied; Deadly Poison is not used by any spec |

**Why Slice and Dice measured badly: not a defect.** Poisons are on and the haste is real. The three finishers simply compete for the same combo points, and per combo point the model values them (finisher only, all points into it, 180 s fight):

- Combat: Sinister Strike only 192.7; Slice and Dice only 223.1; Eviscerate only 229.9.
- Subtlety: no finisher 164.4; Slice and Dice 185.7; Eviscerate 196.6; Rupture 207.0 (all with Hemorrhage, Ghostly Strike).
- Assassination (before the Cold Blood fix, so slightly low): Mutilate only 174.1; Venom 192.9; Slice and Dice 195.2; Eviscerate 207.3.

Slice and Dice's +30% only multiplies the haste-scaled share of damage (auto-attacks and the poison procs they trigger, about 45% of the total for Subtlety), so in Subtlety a fully kept-up Slice and Dice is worth roughly +13%, against about +20% for the same points on Eviscerate and about +26% on Rupture. Mixed rotations land between the pure ones; none beats the best pure one. That is the reason every earlier search "winner" dropped finishers.

## 2. Engine defects found and fixed (fork `rogue-curate`)

Each was pinned by a test written first (RED confirmed), then fixed.

Commit `c043f79cc` "talent and strike numbers follow the Forever client text":

| Defect | Was | Client | Files |
|---|---|---|---|
| Dual Wield Specialization | 10% per rank (+50% off hand at 5) | 5% per rank (+25%) | `sim/rogue/talents.go`, `client_values.go` |
| Improved Eviscerate | 5/10/15% | 7/13/20% | `eviscerate.go`, `client_values.go` |
| Lethality | 6% crit-damage bonus per rank | 4% per rank | `talents.go` |
| Opportunity (Backstab, Ambush, Garrote) | 4/8% (a five-rank table indexed by a two-rank talent; Mutilate already had 5/10) | 5/10% | `backstab.go`, `ambush.go`, `garrote.go`, `mutilate.go` |
| Initiative | 25/50/75% | 33/67/100% | `talents.go` |
| Vigor | +10 energy at either rank | +5 / +10 | `rogue.go` |
| Hemorrhage | vanilla "+7 physical damage on 30 hits" (core.HemorrhageAura), 100% weapon damage always | 15% increased Rupture damage for 15 s (read at each Rupture tick), 145% weapon damage with a dagger | `hemorrhage.go`, `rupture.go`, `rogue.go` |
| Ghostly Strike | 125% weapon damage always | 125%, 180% with a dagger | `ghostly_strike.go` |

Commit `f5c4c62fe` "Cold Blood reaches Mutilate's two hits and is spent by them": the Cold Blood flag sat on the Mutilate parent spell, which rolls no hits, so Cold Blood neither made a Mutilate crit nor was consumed by one. The two hand sub-spells now carry the flag, and the off-hand hit spends it so both hands crit.

New tests: `sim/rogue/client_values_test.go`, `talents_test.go` (Initiative), `sim/rogue/dps_rogue/hemorrhage_test.go`, `cold_blood_test.go`. `go test --tags=with_db ./sim/...` passes (only `sim/web`, which needs a binary_dist module, fails to set up, as before).

Effect of the fixes on the unchanged old rotations: Assassination 193.7 -> 192.9, Combat 236.8 -> 228.9 (off hand now +25% not +50%), Subtlety 183.1 -> 191.4 (dagger Hemorrhage).

Commit `5623963fc` syncs the adopted rotations into the fork's `ui/rogue/apls/forever_*.apl.json` (`make apl-sync ENGINE_DIR=<worktree>`; `make apl-check` passes).

## 3. Candidates measured (site `rogue-curate`)

### Combat (baseline of the committed rotation on the fixed engine: 228.9 ± 0.4)

| Candidate | DPS ± | Notes |
|---|---|---|
| Slice and Dice first, 1/3/4/5 CP floor, refresh <2 s / <4 s, either order (16 variants) | 219.9 - 228.9 | none beats the committed Eviscerate-first order |
| Sinister Strike only, no finishers | 192.7 ± 0.3 | |
| Slice and Dice only, 5 CP / 3 CP | 222.7 / 223.1 | |
| Eviscerate only | 229.9 ± 0.4 | drops Slice and Dice: not adoptable |
| Committed structure + Backstab ahead of Sinister Strike | 230.1 | Backstab rarely reaches its 60 energy before Sinister Strike spends it |
| Committed structure, Backstab only | 251.9 ± 0.4 | the band's weapons are daggers; Backstab 150% weapon damage, Puncturing Wounds |
| Same + Sinister Strike fallback held to 60 energy when Backstab is uncastable | 251.9 | sword gear keeps working |
| Eviscerate at 4+ / Slice and Dice opener at 1 CP + refresh at 5 | 254.2 / 254.8 | best on the band, but 5-8% worse than the old rotation at ladder levels 10-50 (Slice and Dice at 1 CP eats the combo points) and -1.6% at ladder 60 |
| **Adopted: Eviscerate 5 CP, Slice and Dice opener (3 CP, first 15 s) + refresh (5 CP, <3 s), Adrenaline Rush, Backstab, Sinister Strike fallback at 60 energy** | **251.7 ± 0.4** | ladder DPS within 1% of the old rotation at every level |

Final tally (full search run): Eviscerate 3.67, Slice and Dice 5.17, Rupture 0 per iteration. The search found nothing better ("already the best found"). Inserting Rupture costs -11.1.

### Assassination (baseline on the fixed engine: 194.5 ± 0.4; the committed rotation casts Venom 8.7 and never Eviscerate or Slice and Dice)

| Candidate | DPS ± | Eviscerate / Slice and Dice / Venom / Cold Blood casts |
|---|---|---|
| Committed: Cold Blood, Venom 5 CP, Eviscerate 5 CP, Slice and Dice 5 CP | 194.5 ± 0.4 | 0 / 0 / 8.7 / 1.3 |
| Slice and Dice, Venom, Eviscerate at 5 CP (any order) | 194.2 - 200.4 | Eviscerate 0 in most |
| Eviscerate at 4+ CP, Slice and Dice and Venom behind it (not reached: Eviscerate takes every 4 CP) | 207.8 ± 0.4 | 13.0 / 0 / 0 / 1.4 (held: drops both) |
| Mutilate + Rupture instead of Slice and Dice and Venom | 206.0 | 5.7 / 0 / 0 (Rupture 8.0) (held) |
| Backstab builder, no Mutilate | 208.2 | 4.0 / 4.8 / 0.7 (held: drops the spec's talent builder) |
| **Adopted: Cold Blood (4 CP), Eviscerate 4+ CP, Slice and Dice (>=4 CP, down or <3 s), Venom (>=4 CP, down only), Mutilate** | **204.6 ± 0.4** | **5.1 / 4.8 / 2.8 / 1.5** |

The 405-point grid of finisher thresholds and orders is summarised above; the adopted one is the best of those that cast all three finishers. It works because Eviscerate costs 35 energy and Slice and Dice and Venom 25: on the beats with the combo points but under 35 energy the cheaper finishers go, so every finisher fires.

Final full search: "better rotation +3.6" (insert Backstab, which displaces Venom 2.79 -> 0.74): held.

Sensitivity to the poison kit (local rebuild with Deadly Poison main hand, Instant Poison off hand for Assassination; NOT committed): committed rotation 229.1, adopted 236.3, Eviscerate-only 239.9, Mutilate beats Backstab as builder (236.3 vs 232.9). Ranking and the adopted structure are unchanged, and everything is about +17% higher.

### Subtlety (baseline on the fixed engine: 191.4 ± 0.4)

| Candidate | DPS ± | Eviscerate / Slice and Dice / Rupture |
|---|---|---|
| Committed: Ghostly Strike, Eviscerate 5, Slice and Dice 5, Hemorrhage | 191.4 | 4.7 / 4.5 / 0 |
| Rupture at 5 CP only (all three finishers except Rupture dropped) | 207.0 ± 0.3 | 0 / 0 / 10.8 (held) |
| Rupture + Eviscerate at 4+ while Rupture is up + Slice and Dice opener only | 203.8 | 1.0 / 1.0 / 8.8 (best with all three, but token casts) |
| Backstab builder instead of Hemorrhage | 203.8 | 0.5 / 0.8 / 8.7 (no Hemorrhage debuff, Rupture loses its 15%; held) |
| Slice and Dice before Rupture, 3-5 CP floors | 185.5 - 189.2 | Rupture starved |
| **Adopted: Ghostly Strike, Rupture (>=4 CP, not running), Eviscerate (>=3 CP, Rupture running), Slice and Dice (>=3 CP, Rupture running, down or <3 s), Hemorrhage** | **200.8 ± 0.3** | **2.6 / 2.8 / 8.3** |

Final full search: "better rotation +7.8" (the Rupture-only rotation: Eviscerate 0, Slice and Dice 0, Rupture 13.9): held per precedent.

Ladder DPS (sim/request goldens, new vs old rotation on the fixed engine): Assassination 160.5 vs 154.4 at 60, Subtlety 154.9 vs 146.2 at 60, both equal or better at every level from 40; Combat within 1% at every level.

## 4. Guides

`web/src/content/guides/rogue/{assassination,combat,subtlety}.md` "Rotation and priority" rewritten for the adopted rotations (relative figures only: "about 10%" for dagger Backstab over Sinister Strike). The Subtlety Gear paragraph also still described Hemorrhage as a charge-based bleed; corrected. The Combat Gear sentence now names Backstab as well as Sinister Strike.

## 5. Not resolved / for the owner

- **Assassination kit.** `leveling.KitConsumes` gives every rogue Instant Poison on both weapons. Mutilate's "+20% against poisoned targets" only counts Deadly or Wound Poison (Instant leaves no aura), so it never applies in any sim or ranking, and a real Assassination rogue runs Deadly Poison. Deadly main hand adds about 17% DPS (measured above). Changing it moves ladder goldens and the nightly BiS inputs, so I did not.
- **Rotation vs weapon.** The APL has no weapon-type condition. Combat uses Backstab then a Sinister Strike fallback held to 60 energy (`not spellCanCast(Backstab)`), which costs a sword rogue nothing measurable but is a workaround; an `mainHandIsDagger` APL value would be cleaner.
- **Mixed finishers are a price.** The model prefers pure Rupture (Subtlety) and pure Eviscerate (Assassination, Combat). The adopted rotations keep every finisher casting, at -3% (Subtlety), -2% (Assassination) against the single-finisher optimum, as instructed.
- **Murder and Serrated Blades** not checked in depth: client Murder is 2%/4% against Humanoid and Giant only (engine 1%/2%, also Beast and Dragonkin); Serrated Blades' armor-penetration formula uses integer division (`5/3`). Neither moves these measurements (default target type).
- **Tests outside this lane.** In the site worktree `go test ./request/` still fails `TestRotationLadder/hunter-survival` (fork base 1aed40820 predates f39303113, "give the Lacerating Strikes bleed a DefenseType") and the warrior-fury golden; both are independent of rogues and untouched. `TestCombatDaggers`/`TestCombatSinisterStrike` `.results` are skipped ("awaits its Forever talent rewrite"), so they were not refreshed.
- Rotation-search reports under `design/reviews/rotation-search/rogue-*.md` were regenerated by the full-search runs of the adopted rotations.
