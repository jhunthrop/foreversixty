# Sim audit round 13: core (regen, ticks, timing)

Lane core. Fork branch `lane/r13-core` (from int/r13 @ 076a5a517), site branch `wt/r13-core` (from int/r13 @ 93c5d2510). Client 1.60.1.70753; its hotfix push touches only Item and ItemSparse rows, so none of the spells below.

Method: engine reading (sim/core mana.go, energy.go, rage.go, cast.go, gcd.go, apl.go, sim.go; sim/mage/evocation.go; sim/rogue/talents.go), the 70753 client tables (SpellEffect, SpellMisc + SpellDuration, SpellCooldowns, SpellPower, SkillLineAbility, talents/rogue.json), and the cmangos-classic server source (`src/game/Entities/Player.cpp` Player::Update / RegenerateAll / Regenerate, `StatSystem.cpp` Player::UpdateManaRegen, `Unit.cpp` Unit::OCTRegenMPPerSpirit; fetched 2026-11-03 from github.com/cmangos/mangos-classic master). spec-breakdown at 10,000 iterations, seeds 1 and 50001.

## 1. Owner ruling 4: the regen tick's phase

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **The mana regen tick sat on a fixed grid from the pull.** `initManaTickAction` put the first tick exactly 2 s after the prepull start, so every tick fell on the global's 0.5 s grid and on the idle APL's 50 ms poll grid. On a real server the tick is the player's own timer (cmangos `m_regenTimer`, which runs from login and fires `RegenerateAll` at `REGEN_TIME_FULL`), so its phase against the pull is arbitrary. Each iteration now draws the phase from [0, 2 s) with the sim's seeded RNG (label "Mana Tick Phase"), so a seed still reproduces. The draw happens only when some unit has a mana bar, so rogue and warrior random streams are untouched (their numbers are identical to the decimal). Player and pets share one phase per fight. | `TestManaTickPhaseIsRandomPerIteration` (first tick within [0, 2 s), nearly all distinct over 200 iterations, same seed same phase) and `TestManaTicksStayTwoSecondsApart` failed first (first tick at 2 s), then passed. `sim/core/mana.go` initManaTickAction. | Owner ruling (implemented) | old (upstream anchor) | E e88a33c97 |
| C2 | **Energy has no tick to randomise.** Forever regenerates Energy continuously (10 a second, accrued on demand to the nanosecond; `sim/core/energy.go` header, from Blizzard's 30 Sep 2026 class-highlights video). Classic's 20-per-2 s tick (cmangos `Regenerate(POWER_ENERGY)`: 1 per 100 ms, paid on the 2 s timer) is gone, so there is no phase. The APL's "time to next energy tick" is dropped with a warning (`apl_values_resources.go:249`). | engine reading | Aligned (single source for Forever's continuous energy: the video) | - | - |
| C3 | **What else runs on the old grid, and why it does not follow.** (a) Mana Spring (aura 24 every 2 s, 4/6/8/10 a pulse, SpellEffect 5677/10491/10493/10494) and Blessing of Wisdom (aura 24 every 5 s, 12 to 40, 19742 to 25290; Greater 25894/25918 36/40) are folded into MP5 and paid on the regen tick, so they follow the drawn phase; the real pulses have their own timers, but the mean is the same. (b) The idle APL's 50 ms poll, the 1.5 s global, prepull times, encounter movement windows and heal-profile pulses are not regen ticks and are not ruled; they stay. (c) Side effect: a caster waiting for mana now resumes at the next 50 ms poll after the tick (about 25 ms later on average); before, ticks landed on the poll grid. That is inside a player's reaction time and not modelled as a bug. | engine reading | Aligned / noted | - | - |
| C4 | **Round 12's Evocation tick fix (E 7ccb794fd) and Innervate (round 12 B2).** With a drawn phase, an 8 s Evocation spans exactly four ticks except on a measure-zero tie, and the fix still covers that tie. Innervate's end-tick ordering (round 12 B2, left to auras) can now only bite on the same measure-zero tie, so it needs no fix. Evocation 12051 in 70753: aura 110 +1500% regen, aura 134 100% while casting, aura 226 1 s periodic dummy, 8 s, 8 min cooldown; SkillLineAbility 6564 trains 12051. Innervate 29166: aura 110 +400%, aura 134 100%, 20 s, 6 min. The engine's +15 and +4 multipliers on spirit regen match (aura 110 multiplies spirit regen only, cmangos UpdateManaRegen). | client rows; `TestEvocationThatStartsOnARegenTickGetsFourTicks` still passes (it now resets the global before forcing its on-tick cast) | Aligned | - | - |
| C5 | **Round 12's small ties at 0% haste do not go away; they are not regen ties.** Bare 60, horde, seeds 1 / 50001, with a throwaway probe binary (cast speed x0.9999 / x1 / x1.0001, regen phase fixed or drawn): Elemental 191.9 / 191.9 / 191.3 fixed and 192.0 / 192.1 / 191.4 drawn at seed 1 (191.9 / 191.8 / 191.4 drawn at 50001); Discipline 317.5 / 320.7 / 319.1 HPS fixed and 317.3 / 320.5 / 319.0 drawn (317.3 / 320.4 / 318.9 at 50001). Discipline's exact-0% peak (about +1% over both neighbours) is there with either phase. A 3,000-iteration breakdown shows Renew 18.9 to 18.3 casts at +0.01%: a Renew-refresh condition meeting the heal profile's pulses, not regen. spec-breakdown uses the shared (unlabelled) RNG, so the Elemental step (about −0.3% at +0.01%, both seeds) is near noise. | probe runs, `scratchpad/r13/core/haste-ties.txt` | OPEN, routed (healers: Discipline; casters: Elemental) | old | - |

## 2. Fresh re-audit: core regen and timing

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Spirit regen per class at 60, per 2 s tick: priest and mage 12.5 + Spirit/4 (`priest.go:251`, `mage.go:261`), shaman 17 + Spirit/5 (`shaman.go:34`), druid, paladin, hunter and warlock 15 + Spirit/5 (`mana.go` SpiritManaRegenPerSecondDefault). | cmangos-classic `Unit::OCTRegenMPPerSpirit` lists exactly these; the shaman comment cites vmangos too. The Classic client carries no regen table. | Aligned (server source; Forever states no change) | - | - |
| A2 | Five-second rule: spending mana sets `FiveSecondRuleRefreshTime = max(now + 5 s, cast end)`; a tick before that pays MP5 plus the casting share of spirit regen, a tick at or after pays full. Hardcasts spend at completion, channels at the start. | cmangos `IsUnderLastManaUseEffect` (strictly under 5 s), `m_modManaRegenInterrupt = mp5 + spirit x interrupt% (capped 100)`; `mana.go:209`, `mana.go:363`. | Aligned | - | - |
| A3 | MP5 pays mp5/5 a second on the 2 s tick, during the five-second rule too. | cmangos `power_regen_mp5 = MOD_POWER_REGEN / 5` in both rates | Aligned | - | - |
| A4 | Mana Spring and Blessing of Wisdom amounts (see C3). | SpellEffect rows above; `buff_ranks.go:111`, `:203` | Aligned | - | - |
| A5 | Adrenaline Rush 13750: +100% Energy regen (aura 110, misc 3), 15 s, 5 min, 1 s global, no cost. Blade Flurry 13877: 20% attack speed (aura 319), 15 s, 2 min, 25 Energy, 1 s global. Vigor 14983: 2 ranks, +5 / +10 maximum Energy. | SpellEffect, SpellMisc + SpellDuration, SpellCooldowns, SpellPower; talents/rogue.json; `rogue/talents.go:395-505`, `rogue/client_values.go:26` | Aligned | - | - |
| A6 | Rage from damage taken uses the receiver's level conversion (230.6 at 60) at 2.5 x damage / c; rage from swings is Forever's normalized rage. The off-hand factor (0.5) and the crit multiplier (2.0) still rest on assumption and one note; nothing new to settle them. | `rage.go`; Kalgan formula; cmangos RewardRage; memory forever-rage-normalization-gap | Aligned / OPEN as before (needs a log) | - | - |
| A7 | Round 12's cast-timing fix (E 1e7e01e4b, SetGCDTimer gives a cast that ends before a new global its own completion). Checked: the new hardcast action and the global's action both test `Hardcast.Expires != startingCDTime` and clear it, so a cast cannot complete twice; at cast start the condition sees only the previous, already-cleared cast; pushback reschedules through the same action. No side effect found. | `gcd.go:43-70`, `cast.go:280-345`, `character.go:436` | Aligned | - | - |
| A8 | The global: 1.5 s, floor 1 s, not shortened by spell haste (`cast.go:243`, "Vanilla has no natural GCD reduction"). `Spell.EffectiveCastTime` (hasted global) and `Unit.TimeUntilManaRegen` have no callers, so they cannot skew a result. Spell batching: `SpellBatchWindow` 10 ms, used only for swing timers after a stop/restart and extra attacks. | engine reading; round 12 casters H5 | Aligned (dead code noted, not removed: outside a fix) | - | - |

## Before/after

E e88a33c97 against `$R/before.tsv` (raid 60) and `$R/sb-before` (bare 60, horde). 10,000 iterations. Every move is within ±0.30%; rogues and warriors are identical.

| spec | raid 60 alliance s1 / s50001 | raid 60 horde s1 / s50001 | raid 60 horde s1 absolute | bare 60 horde s1 / s50001 | bare 60 horde s1 absolute |
|---|---|---|---|---|---|
| druid-balance | +0.02% / +0.02% | -0.05% / -0.03% | 572.2 -> 571.9 | -0.04% / -0.07% | 280.4 -> 280.3 |
| druid-feral-bear | +0.01% / -0.01% | +0.05% / -0.01% | 1010.7 -> 1011.2 | +0.07% / -0.07% | 443.7 -> 444.0 |
| druid-feral | -0.02% / -0.11% | +0.06% / -0.05% | 660.0 -> 660.4 | -0.04% / -0.04% | 266.9 -> 266.8 |
| druid-restoration | -0.07% / +0.11% | -0.04% / +0.05% | 833.7 -> 833.4 | -0.12% / -0.12% | 416.9 -> 416.4 |
| hunter-beast-mastery | +0.05% / -0.13% | +0.03% / +0.02% | 623.5 -> 623.7 | -0.03% / -0.03% | 301.3 -> 301.2 |
| hunter-marksmanship | +0.10% / +0.00% | -0.10% / -0.11% | 622.5 -> 621.9 | +0.24% / +0.27% | 293.3 -> 294.0 |
| hunter-survival | -0.01% / -0.07% | -0.09% / -0.08% | 744.7 -> 744.0 | +0.03% / +0.03% | 295.0 -> 295.1 |
| mage-arcane | +0.08% / -0.08% | -0.10% / -0.03% | 619.8 -> 619.2 | -0.15% / +0.18% | 343.7 -> 343.2 |
| mage-fire | +0.18% / +0.21% | +0.00% / +0.10% | 673.0 -> 673.0 | +0.30% / +0.24% | 334.1 -> 335.1 |
| mage-frost | +0.11% / -0.18% | +0.05% / -0.11% | 632.0 -> 632.3 | +0.14% / -0.09% | 346.7 -> 347.2 |
| paladin-holy | +0.05% / +0.03% | -0.08% / -0.06% | 645.1 -> 644.6 | +0.06% / +0.12% | 323.0 -> 323.2 |
| paladin-protection | +0.06% / -0.01% | -0.02% / +0.07% | 1125.7 -> 1125.5 | -0.24% / -0.17% | 467.9 -> 466.8 |
| paladin-retribution | -0.01% / +0.01% | -0.07% / -0.11% | 701.5 -> 701.0 | +0.00% / -0.04% | 247.3 -> 247.3 |
| priest-discipline | -0.05% / -0.11% | -0.03% / -0.03% | 655.9 -> 655.7 | -0.06% / -0.09% | 320.7 -> 320.5 |
| priest-holy | +0.01% / +0.10% | -0.01% / -0.10% | 698.9 -> 698.8 | +0.03% / +0.09% | 317.8 -> 317.9 |
| priest-shadow | +0.07% / +0.03% | -0.01% / +0.03% | 720.0 -> 719.9 | +0.13% / +0.08% | 389.2 -> 389.7 |
| rogue-assassination | +0.00% / +0.00% | +0.00% / +0.00% | 622.1 -> 622.1 | +0.00% / +0.00% | 217.1 -> 217.1 |
| rogue-combat | +0.00% / +0.00% | +0.00% / +0.00% | 628.2 -> 628.2 | +0.00% / +0.00% | 213.1 -> 213.1 |
| rogue-subtlety | +0.00% / +0.00% | +0.00% / +0.00% | 550.2 -> 550.2 | +0.00% / +0.00% | 218.1 -> 218.1 |
| shaman-elemental | +0.00% / +0.04% | -0.02% / +0.02% | 479.6 -> 479.5 | +0.10% / +0.05% | 191.9 -> 192.1 |
| shaman-enhancement | -0.02% / -0.10% | -0.05% / -0.15% | 586.0 -> 585.7 | +0.09% / -0.09% | 212.0 -> 212.2 |
| shaman-restoration | -0.06% / +0.04% | -0.02% / +0.01% | 840.2 -> 840.0 | +0.00% / -0.02% | 430.0 -> 430.0 |
| warlock-affliction | -0.02% / +0.07% | +0.02% / +0.02% | 571.9 -> 572.0 | +0.00% / -0.07% | 307.3 -> 307.3 |
| warlock-demonology | -0.02% / +0.06% | +0.00% / +0.02% | 639.3 -> 639.3 | +0.00% / +0.00% | 347.0 -> 347.0 |
| warlock-destruction | -0.03% / +0.03% | +0.09% / +0.09% | 636.2 -> 636.8 | +0.00% / +0.13% | 377.6 -> 377.6 |
| warrior-arms | +0.00% / +0.00% | +0.00% / +0.00% | 731.2 -> 731.2 | - (no mana, no draw) | - (no mana, no draw) |
| warrior-fury | +0.00% / +0.00% | +0.00% / +0.00% | 758.2 -> 758.2 | - (no mana, no draw) | - (no mana, no draw) |
| warrior-protection | +0.00% / +0.00% | +0.00% / +0.00% | 1068.8 -> 1068.8 | - (no mana, no draw) | - (no mana, no draw) |

The only steady moves at both seeds are bare Marksmanship (+0.24 / +0.27%) and Fire (+0.30 / +0.24%; raid +0.18 / +0.21% alliance), both mana-bound: the old grid's ties came out against them.

## Fork goldens regenerated (E e88a33c97)

`sim/mage/TestP1Mage.results`, `sim/paladin/protection/TestProtection.results`, `sim/shaman/elemental/TestElemental.results`, `sim/shaman/enhancement/TestEnhancement.results`. Each package was run alone, niced. The other goldens (Balance, Feral, Feral Tank, Hunter, Retribution, Shadow, Warden, both warlocks, rogues, warriors) do not move: the fork tests use labelled RNG, so the new draw shifts no other roll, and those fights never wait on mana. Three mage unit tests had stepped onto the old whole-second grid and now own their moment: Heating Up and Evocation reset the global before a forced cast; Ice Barrier steps exactly to its expiry at the fight's 60 s end, because a step aimed past the end does not move the clock and no tick sits at 60 s any more. Also passing: sim/core, the druid, paladin, priest, shaman, warlock, rogue and warrior packages, sim/conformance and sim/common/*.

## Files to regenerate

- `sim/request` TestRotationLadder goldens: every mana spec, every band.
- `sim/adapter` golden_test, if its fixtures sim a mana user.
- Nightly-owned: BiS json, dps-per-point.json, addon-builds and Data.lua numbers (all mana specs, small moves).

## Measure at refresh

- Haste scale factors for Discipline and Elemental: they still carry C5's exact-0% ties.

## For other lanes

- healers: Discipline's +1% peak at exactly 0% haste (C5). Renew is cast 3% less at ±0.01% haste; look for a Renew refresh or heal-profile pulse condition that lines up at exact 1.5 s / 3 s steps.
- casters: bare Elemental's −0.3% at +0.01% haste (C5), near noise on the shared RNG; a labelled-RNG pair would settle it.
- auras: round 12 B2 (Innervate's end tick) can be closed; the drawn phase makes the tie measure-zero.

## Still open

- C5 (routed above).
- A6: normalized rage's off-hand factor and crit multiplier need the owner's combat log.
- Per-unit regen phase: player and pets share one drawn phase. A server gives each its own timer; the effect on a pet's mana is far below error, so it is left.

Housekeeping: the shared disk was full (417 MB free) at the start of this lane; Go build-cache entries untouched for more than two days were deleted (`find ~/Library/Caches/go-build -mtime +2 -delete`), which freed about 21 GB without touching any running build's entries.
