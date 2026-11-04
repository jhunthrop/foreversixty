# Sim audit round 14: core

Lane `core`. Fork `lane/r14-core` (from int/r14 5195a8889), site `wt/r14-core` (from int/r14 72471fcb3). Sources fetched 2026-11-03: cmangos-classic master dfb37aa32 (`src/game/Maps/Map.cpp`, `Entities/Object.h`, `Entities/Object.cpp`, `Entities/Unit.cpp`, `Spells/Spell.cpp`, `Spells/SpellAuras.cpp`, `Spells/SpellMgr.cpp`, `Spells/UnitAuraProcHandler.cpp`); wago.tools 1.60.1.70753 SpellMisc, SpellDuration, SpellCastTimes; Blizzard's 1.13.7 batch change (warcrafttavern.com/wow-classic/news/patch-1-13-7-now-available-ptr/, massivelyop.com/?p=319302, us.forums.blizzard.com/en/wow/t/wow-classic-version-1137-ptr-is-now-available/838945).

Haste probe: a throwaway copy of each engine (never committed) that multiplies the player's cast speed by `FS_PROBE_HASTE` (0.9999 = −0.01%, 1.0001 = +0.01%, 1.005 = +0.5%) and can print a millisecond log.

## 1. Owner ruling 2: same-instant ties (Discipline and Elemental at 0% haste)

How the engine orders one instant: `Simulation.advance` expires auras first (so a channel's last tick, run by the Dot's CleanUp, always comes first), then weapon swings and tasks, then pending actions by priority (DoT ticks, raid pulses and the GCD are all priority 0, first-in-first-out). `Timer.IsReady` is inclusive (`readyAt <= now`); `PushBackCast` does nothing to a cast whose end is now (`IsCasting` is `Expires > now`).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **A hit a hair before a cast ends, or before a channel's last tick, cost the cast a full second or the tick; a hit at the very same instant cost nothing.** Raid pulses push casts back, and the cap (a cast never has more than its full cast left) lines cast ends up on the pulses. At exactly 0% haste a pulse lands on the cast's last instant and does nothing. At −0.01% the cast ends 0.25 ms after the pulse, and the pulse delayed it by the full 1 s. A pulse 1 ns before Penance's last tick cut that tick. That made Discipline dip 0.9% at −0.01% haste (raid 60 650.1 HPS against 655.8 at 0%).<br>The server does not resolve 0.25 ms: it handles events in update batches. Inside a batch it updates players (casts finishing, channel ticks) before creatures (cmangos Map::Update: players' `Update` before the grid objects). A periodic aura ticks when its timer reaches 0 in that update (SpellAuras.cpp Aura::Update: "tick also at m_periodicTimer==0 to prevent lost last tick"). A cut channel keeps whatever remains (Spell::DelayedChannel, Unit::DelaySpellAuraHolder). Blizzard's Classic batch is 10 ms since 1.13.7 ("usually 1 to 10 ms").<br>**Fix:** a hit inside the 10 ms batch before a cast's end does not push it back. A hit inside the batch before a channel's next tick keeps that tick. The chance falls from 1 at no gap to 0 at a whole batch, since the batch edge can fall anywhere. Exact ties draw no random number and keep today's result. | `sim/core/cast.go` PushBackCast and pushBackChannel; `server_batch.go`; tests `server_batch_test.go` (failed first: 5 ticks for 6, a 1 s pushback for 0 s); probe table below | CONFIRMED-BUG | old (round 13 core C5 measured it and left it open; healers X3 put it down to tie order) | E cbdc60fd5 |
| C2 | **Tie (a): Penance's 12 s cooldown ending on the same instant as a cast or global.** The engine casts Penance at the tie, and that is the server's rule. The cooldown is over when `now >= m_expireTime` (cmangos Object.h `CooldownData::IsSpellCDExpired`). The global is over when `m_globalCooldown <= now` (Object.cpp `CooldownContainer::IsGlobalCooldownExpired`). A spell cast at the instant both end is allowed. So 0% is not "high" here. I tested whether the +0.01% dip comes from Penance being 0.25 ms short: a Discipline list with a wait line ahead of the fillers ("wait for Penance when it is ready within 0.1 s", the same shape as Elemental's Lava Burst wait, E 25fca77c0) moved nothing at +0.01% (652.7 / 652.5 against 652.7 / 653.1). One logged fight's Penance gaps at +0.01% are global and priority steps (12.5 s = a 3 s Prayer of Healing then a 1.5 s global), not 0.25 ms misses. | Object.h:154-160, Object.cpp:3135-3138; scratch rotation `scratchpad/r14/core/disc-penance-wait.json`; logs `scratchpad/r14/core/dlog-*.txt` | Aligned (server rule; random order would be wrong here) | - | - |
| C3 | **Tie (b): a raid pulse on the exact instant of a channel tick.** Either order keeps the tick. Tick first: the channel is over before the pulse lands. Pulse first: `newExpires = max(now, ...)` is now, and a tick due at now still counts (floor of 1 period). cmangos keeps it in both orders as well (above). The tick was lost only when the pulse came a hair *before* it, which is C1. | `cast.go` pushBackChannel; SpellAuras.cpp:395-407 | Aligned (server rule) | - | - |
| C4 | **What is left: 0% still sits 0.45% above +0.01% for Discipline raid 60 (0.2% bare).** It is not a same-instant tie: the Penance wait probe and the batch rule leave it, and −0.01% now reads with 0% (654.9 against 655.8). It is the 0.5 s step grid at 0% (unhasted 1.5 s globals with 2.5 s and 3.0 s casts, 1.0 / 0.8 s pushbacks, a 12 s cooldown, 5 s pulses), which any haste breaks up. Elemental shows no step beyond noise (raid 479.7 / 479.6 at 0%, 479.2 / 479.2 at +0.01%). | probe table below | OPEN: a log diff of many paired iterations at 0% against +0.01% (cast counts per spell and Penance gap histogram) would show which priority line takes the slot | old | - |

Haste probe, horde, 10,000 iterations, seeds 1 / 50001 (effective HPS, DPS):

| spec, preset | −0.01% before → after | 0% before → after | +0.01% before → after | +0.5% before → after |
|---|---|---|---|---|
| Discipline raid 60 | 650.2 / 650.0 → 654.9 / 654.9 | 655.7 / 655.9 → 655.7 / 655.9 | 652.7 / 653.0 → 652.7 / 653.1 | 653.4 / 653.7 → 653.5 / 653.7 |
| Discipline bare 60 | 318.5 / 318.4 → 320.2 / 320.1 | 320.3 / 320.2 → 320.4 / 320.2 | 319.8 / 319.5 → 319.8 / 319.5 | 320.7 / 320.2 → 320.7 / 320.3 |
| Elemental raid 60 | 480.2 / 479.9 → 479.7 / 479.5 | 480.3 / 480.0 → 479.7 / 479.6 | 479.6 / 479.6 → 479.2 / 479.2 | 481.0 / 481.1 → 480.4 / 480.9 |
| Elemental bare 60 | 192.4 / 192.0 → 192.1 / 192.1 | 192.4 / 192.0 → 192.1 / 192.1 | 192.3 / 191.9 → 192.0 / 191.9 | 192.7 / 192.4 → 192.5 / 192.4 |

(The Elemental "after" column also carries C5's Searing Totem change. C1 alone moved no Elemental number.)

## 2. Searing Totem's duration under the 2.437 s cadence

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C5 | **The totem's attacks ran out 1.386 s before the totem (0.756 s at rank 1), and the rotations dropped a new one then.** The client's totems last 30 / 35 / 40 / 45 / 50 / 55 s (SpellMisc DurationIndex of 3599, 6363, 6364, 6365, 10437, 10438). The Attack is a 2.2 s cast, missile speed 19 (3606 … 10436). The engine modelled the totem as a Dot of floor(55 / 2.437) = 22 attacks, so the Dot ended at 53.614 s. Elemental and Enhancement redrop Searing on `not dotIsActive`, so they recast it 1.386 s early every time: 2.6% more totem casts, mana and globals than a shaman who drops it when it expires. The number of attacks is right. A full totem lands floor((D − 2.2 − first-cast delay) / 2.437) + 1 attacks in game: 12 / 14 / 16 / 18 / 20 / 22 for the six ranks, whether a bolt still in flight at expiry lands or not. The engine gets the same counts. No attack is lost.<br>**Fix:** the Dot now lasts the totem's full client duration and still lands 22 attacks. Before round 13's cadence fix (2.5 s), 55 / 2.5 = 22 ended exactly at 55 s, so the gap came with that fix. I edited sim/shaman/fire_totems.go for this one item, as the coordinator allowed. | `sim/shaman/fire_totems.go`; test `TestSearingTotemStandsForTheClientDuration` (failed first: "run out at 53.614s, want 55s"); conformance package still passes | CONFIRMED-BUG | R13 5ca93cfa3 | E 2fe671caa |

## 3. Regression over rounds 12, 13 and the landing (fork sim/core, 2606bc05e..5195a8889)

At the lane head, with `--tags=with_db` and `GOWORK=off`: `./core/`, `./paladin/...`, `./priest/...`, `./druid/...`, `./shaman/...`, `./hunter/...`, `./rogue/...`, `./mage/...`, `./warlock/...`, `./common/...`, `./conformance/` and the sim root all pass. Every failure list matches base-engine (5195a8889) line for line. Note: without `--tags=with_db`, relic and golden tests panic on missing items (12640, 15684, 18817, 228176) in base and lane alike. They are not regressions.

| # | commit | what it said | checked now | result |
|---|---|---|---|---|
| R1 | e88a33c97 random regen tick phase | phase drawn per iteration from [0, 2 s), only with a mana bar | `mana.go:263-276`; TestManaTickPhaseIsRandomPerIteration, TestManaTicksStayTwoSecondsApart pass; C1 adds no draw at exact ties and none for units that are not hit | Aligned |
| R2 | 1e7e01e4b a cast ending on the global completes on time | SetGCDTimer gives a pending cast its own completion | `gcd.go:50-56`; PushBackCast calls SetGCDTimer with `hc.Expires` (equal, so the clause stays off); healer prepull Lightwell/Inner Fire unchanged (no healer golden moved) | Aligned |
| R3 | 7ccb794fd Evocation keeps the tick at the channel end | regen ends just after the end tick | mage package passes; with R1's random phase the tie is rarer and still handled | Aligned |
| R4 | 8a54e923e missile speeds | Auto Shot 40, Shoot (5019) 20 | `attack.go:204-208`, `wand.go:60-62`; SpellMisc Speed column agrees | Aligned |
| R5 | 0a879f7f6 item haste counts once | equip haste only in stats | `character.go:252-266`, TestAnItemsHasteCountsOnce passes. One note: items add up, while the server multiplies separate haste auras (Unit::ApplyCastTimePercentMod); the gap is under 0.1% for any two Forever items | Aligned |
| R6 | 2fde66add Read Leyline | 50% regen 15 s, 2 min | racial tests pass | Aligned |
| R7 | 3322675af Poised; d347b5094 consumables; 9291833a7 Thunder Clap; 46aa4f61a / 8e4de235c Command; f085a6c29 Blood Fury; 46be83bcd Holy Power; 1c2190b4d category 1153 | as stated | their tests pass; C1 and C5 touch none of them | Aligned |
| R8 | 5ca93cfa3 Searing cadence + 5195a8889 conformance | 2.437 s cadence; the shaman golden carries no mismatch | cadence kept; the duration gap it opened is C5 | C5 |
| R9 | 25fca77c0 Elemental waits for Lava Burst | wait line within 0.1 s | TestASliverOfHasteDoesNotCostLavaBursts passes | Aligned |
| R10 | landing 54ea3bfb1 (vendored trainables / races) | data only | conformance passes | Aligned |

## 4. Fresh re-audit of the shared mechanics

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Melee table at 300 skill against a level-63 boss: miss 8% plus 1% hit suppression, dodge 6.5%, parry 14% (from the front only), glancing 40% at 55-75% damage, crit suppression 4.8%, block 5% (from the front only). One roll, in the order miss, dodge, parry, glancing, block, crit, hit. | `target.go:451-494`, `spell_outcome.go:345-369`; magey Era attack table (round 11 P3). cmangos puts block ahead of glancing (Unit.cpp:2709-2735, the die enum) and uses skill × 0.2% crit suppression with no 1.8% aura part (Unit.cpp:3733-3770). Order only matters when the table overflows, and it does not. Blizzard-server testing (magey) wins for Forever. | Aligned (cmangos differs on crit suppression; Era testing is the source kept) | - | - |
| A2 | Spell hit: 17% against +3, minimum 1%; no spell crit suppression (the line is commented out); level resistance 2% per level (6% at +3); resist coefficient against level × 5. | `target.go:470,494`, `spell_result.go:252-258`, `spell_resistances.go:125-170`; cmangos CalculateSpellCritChance has no level term (Unit.cpp:3828-3853); level resistance 0.4% per skill point = 6% at +3 (Unit.cpp:3943-3950) | Aligned (two sources) | - | - |
| A3 | Armor: armor / (armor + 400 + 85 × attacker level) | `spell_resistances.go:84-89`; vanilla formula | Aligned | - | - |
| A4 | Crit multipliers: spells 1.5, melee and ranged 2.0, talents scale the bonus (Ice Shards / Ruin 2.0, Impale 2.2) | `spell_outcome.go:1071-1087` | Aligned | - | - |
| A5 | Global: 1.5 s, not hasted (makeCastFunc leaves the GCD alone), 1 s floor. `Spell.EffectiveCastTime` hastes the global but nothing calls it. | `cast.go:56-63,235-239`, `unit.go:373-375`, `spell.go:659-664` | Aligned (dead helper noted) | - | - |
| A6 | Haste: cast speed = 1 / (multiplier × (1 + haste% / 100)); buffs multiply, item stats add | `unit.go:377-379`; cmangos multiplies per aura (Unit.cpp:10501-10507) | Aligned (see R5) | - | - |
| A7 | Procs per minute: chance = PPM × the weapon's base speed / 60, unhasted, per hand | `attack.go:965-1010`; cmangos GetPPMProcChance with GetAttackTime, which divides out the speed mods (Unit.cpp:7731-7737, Unit.h:1353, UnitAuraProcHandler.cpp:596-601) | Aligned (two sources) | - | - |
| A8 | Dual wield: +19% white miss, none on specials; Forever removed the Heroic Strike queue off-hand bonus | `spell_outcome.go:742-746`; cmangos +19 (Unit.cpp CalculateEffectiveMissChance); TestQueuedHeroicStrikeKeepsTheDualWieldPenalty | Aligned | - | - |
| A9 | Cooldowns and the global are over at the exact instant (inclusive) | C2 | Aligned | - | - |
| A10 | Cast pushback: 1.0 s, then 0.2 s less per hit, never under 0.2 s, never more than the full cast; channels lose the same amount | `cast.go:107-215`; cmangos Spell::Delayed / DelayedChannel; plus C1 for the batch | Aligned after C1 | - | E cbdc60fd5 |

## Before/after

Raid 60, 10,000 iterations, seeds 1 and 50001, both factions, against `$R/before.tsv`. Nothing else moved:

| spec | change |
|---|---|
| shaman-elemental | −0.13 / −0.06% alliance, −0.12 / −0.08% horde (480.3 → 479.7 horde seed 1) — C5 |
| shaman-enhancement | −0.15 / −0.03% alliance, −0.19 / −0.07% horde (587.0 → 585.9 horde seed 1) — C5 |
| shaman-restoration | +0.04% alliance 50001, rest same — C1 |
| paladin-holy, priest-holy, priest-discipline, druid-restoration | within ±0.02% — C1 |

Bare 60 horde: Elemental 192.4 / 192.0 → 192.1 / 192.1; Enhancement 212.7 / 212.7 → 212.3 / 212.3. Discipline's published number (0% haste) does not move. Its −0.01% haste point rises 0.7% (raid) and 0.5% (bare).

Fork goldens regenerated (C5 only, in sim/shaman): `sim/shaman/elemental/TestElemental.results`, `sim/shaman/enhancement/TestEnhancement.results`. C1 moves no fork golden (the core, paladin, priest, druid and shaman restoration packages pass unchanged).

## Files to regenerate

- sim/request TestRotationLadder goldens for the shaman and healer rows (elemental, enhancement; tiny healer moves)
- data/builds/1.60.1.70753/dps-per-point.json (Discipline's and the healers' haste points around 0%; Elemental, Enhancement)
- addon-data / coach starters only if they carry the shaman or healer numbers above

## Measure at refresh

- Elemental and Enhancement BiS and talent picks: the totem now recasts every 55 s instead of 53.6 s (−0.1%). No pick is expected to change.
- Discipline haste weight: the −0.01% side no longer dips; C4 is still open.

## For other lanes

- healers: C4 (Discipline 0% sits 0.45% over +0.01% raid 60). It is not Penance's cooldown tie: a 0.1 s Penance wait moved nothing (`scratchpad/r14/core/disc-penance-wait.json`). Look at which line takes the slot at +0.01% before writing any APL change.
- casters: Elemental and Enhancement redrop Searing Totem at 55 s now (C5). The APLs' `dotIsActive(10438)` condition is unchanged and still right.

## Still open

- C4 (above).
- Should cooldown readiness also follow the 10 ms batch? A cooldown that comes up inside the batch the rotation acts in would count as ready. I did not do this: it would move every spec, and the C2 probe found no Penance case that needs it.
- WCL-dependent checks are blocked (gcloud login expired); none were needed here.
