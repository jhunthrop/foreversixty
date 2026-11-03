# Sim audit round 13: regression lane, melee and physical side (regmelee)

Scope: everything round 12 changed for physical and hybrid specs, re-checked against client 1.60.1.70753 and in the sim, then a fresh look at the warrior and feral openers and pre-pull lines.

- **Branches:** fork `lane/r13-regmelee` from int/r13 (076a5a517, no fork commit), site `wt/r13-regmelee` from int/r13 (93c5d2510).
- **Client rows:** the layered 70753 raw tables in `scratchpad/r11/c70753/builds/1.60.1.70753/raw/` (SpellMisc, SpellEffect, SpellDuration, SpellCastTimes). A tooltip and its client row count as one source.
- **Measurements:** spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, unless a row says otherwise. Exploratory runs at 300 to 3,000 iterations are marked.
- **Hotfix push (2026-11-03 14:14 UTC):** the new `hotfix` lane owns the hunter amounts it changed (Aimed Shot, Arcane Shot, Serpent Sting, Trueshot Aura, Aspect of the Hawk, Sniper Shot, Lacerating Strikes). This lane did not touch hunter amounts.

**Result: 1 confirmed bug, fixed (old).** Every other round-12 change re-checks as aligned.

## Round-12 changes: re-check

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Poised, rogues (18) and cat (32).** One mechanism, `core.RegisterPoised` (`sim/core/poised.go`), up at the start of every fight and spent by the first ability in the client mask. In the sim: rogue Garrote costs 40 at bare 20 (50 less 10) and 20 at bare 40 (less Dirty Deeds 20, less Poised 10); Subtlety's Premeditation (no Energy, not in the mask) leaves Poised for the Garrote; Stealth at -3 s costs nothing. Cat: Ravage costs 50 at bare 40 (60 less 10), Prowl at -2 s costs nothing, and no cat pre-pull line spends Energy (no Tiger's Fury before the pull), so the buff's "at maximum Energy" holds when Ravage is cast. Bands below the learn level have no Poised (bare 30 cat, round-12 tests). The 110 Energy of Vigor is not a factor: `energyBar.reset` starts at the maximum, whatever it is. | breakdowns at cat bare 30/40/raid 60, Subtlety bare 20/40, Combat bare 60 (500 iterations); passive and buff rows as in round 12 (SkillLineAbility 59888/59889, SpellEffect 1368039/1368040) | Aligned | - | - |
| G2 | **The Arms opener: Charge at -2 s, Battle Shout at -1.5 s, Bloodrage on the pull.** The client attributes settle round 12's open point: Charge (100, 6178, 11578) carries SpellMisc Attributes_0 0x10000000, the "not in combat" attribute, and Bloodrage 2687's effect 1 triggers 29131, whose effect 1 is aura 94 (SpellEffect 704871), the aura cmangos turns into the in-combat flag. Bloodrage itself has no combat attribute (Attributes_0 0x50010). So Bloodrage before Charge is refused by the client's own rows; the round-12 order is right. Each band casts its learned Charge rank (bare 20: 100 for 9 rage; 30 and 40: 6178 for 12; 50 and 60: 11578 for 15), plus Improved Charge's 6. | SpellMisc 100/6178/11578/2687/29131; SpellEffect 679778 (2687 → 29131), 704871; breakdowns at bare 20 to 60 | Aligned (client rows; a launch log would only confirm the server honours them) | - | - |
| G3 | **Pre-pull Bloodrage and the raid-shout skip, all three warriors.** Raid 60: Battle Shout shows only the raid's aura (1.00 proc, no cast) for Arms, Fury and Protection. Bare bands shout. The gate is rank-rewritten with the cast: at band 40 with a raid shout added (`-buffs battle_shout`) the raid aura is 11549 and the shout is skipped; without it Fury casts 11549 1.54 times (Battle Shout lasts 3 minutes, SpellDuration 180000 for every rank). Fury and Protection pull with Bloodrage at -2.5 s; neither casts anything that needs to be out of combat afterwards. | breakdowns, Fury bare 40/50 with and without the raid shout (300 iterations) | Aligned | - | - |
| G4 | **Arms' execute swap before Recklessness.** Berserker Stance and Recklessness: 0 casts at bare 20, 30 and 40, 1.00 each at bare 50 and raid 60. | breakdowns, bare 20 to 50 | Aligned | - | - |
| G5 | **Improved Charge rage only with a Charge.** Arms shows Improved Charge's 6 rage beside each Charge at every band; Fury's and Protection's builds take no Improved Charge, and neither shows the rage. | breakdowns | Aligned | - | - |
| G6 | **The new Fury build (round 12 W6) under the rule.** Nothing the rotation casts was dropped (Improved Charge and Deflection touch no cast; Improved Tactical Mastery 5 is the way into Anger Management, which pays 1 rage every 3 s). Fury still holds the most points (32), every tier gate holds. But the same build kept one more dead point: see W1. | talents/warrior.json; warrior-fury APL | Aligned (W1 found beside it) | - | - |
| G7 | **Thunder Clap's raid debuff at 20% (11581).** `ThunderClapRaidSpellID` 11581, 20%, 30% with Conqueror's 5-piece, read by the warrior's own spell too. No raid preset carries `thunder_clap` (presets.json debuffs: Curse of the Elements, Sunder, Faerie Fire, the three Judgements, Hunter's Mark, Curse of Recklessness), and tank-encounter.json has none, so no published number reads it; the bear and the paladin tank fight without it. Protection's own clap: rank 6 at 60, 7.39 casts a fight (30 s), 11 rage a cast (20 less Improved Thunder Clap 6 less Focused Rage 3), 16% missed against the level-63 spell table. | `sim/core/debuffs.go`; `sim/warrior/thunder_clap.go`; Protection raid breakdown | Aligned | - | - |
| G8 | **Windfury Totem's bonus ends with the totem.** The change touched only the shaman's own tracker aura (8515/10609/10612, no stats; the APL-dropped totem never gives stats, "TODO" in air_totems.go). The raid buff (`core.ApplyWindfury`) is unchanged: a permanent aura at the band's rank (10610 at 52+). Enhancement drops no Windfury Totem (Grace of Air) and its Windfury Weapon keeps the totem off it. Every other melee and the cat and bear in form receive it (`windfuryTotemReaches`); Arms' Hamstring and Survival's Wing Clip gates on 10610 still find the aura (Arms raid: Windfury Totem 10610 28.2 procs, Hamstring 35.6 casts). The old ids survive only in Enhancement's historical notes and `engine-spell-names.json` (platform, round 12 already listed it). | `sim/core/buffs.go:1667-1745`; E 651660849 diff; Arms raid breakdown | Aligned | - | - |
| G9 | **Missile speeds.** Every hunter shot that the engine fires carries the client's Speed (`sim/hunter/missile_speeds.go`: Aimed, Arcane and Serpent Sting 40, Multi-Shot 30, Sniper Shot 60, Summon Hawk 35, traps 0), Auto Shot 40 (`core/attack.go:207`), wand Shoot 20 (`core/wand.go:62`). Volley is a ground channel with no missile. Hammer of Wrath flies at 35 and deals its rolled damage on arrival (`paladin/hammer_of_wrath.go:56,86`). | SpellMisc rows as in round 12 | Aligned | - | - |
| G10 | **Arcane Shot, Serpent Sting and Volley take no spell power (E 05b72a3b5), pending the owner's in-game tooltip check (ruling 5).** The change is complete: those three are 0; the weapon strikes (Aimed, Multi, Sniper, Raptor, Wing Clip, Auto Shot) keep the weapon's coefficient 1, which is bonus weapon damage, not spell power; the traps and Hunter's Mark set none (0); the pet's abilities scale with the pet, which has no spell power. Not reverted. | `grep BonusCoefficient sim/hunter/*.go` | Aligned, **pending owner confirmation** | - | - |
| G11 | **Juju Might for bow hunters.** BM and MM raid 60 carry `juju_might` (40 ranged attack power) and no Juju Power; Survival carries both. | spec-breakdown consume lists | Aligned | - | - |
| G12 | **Dire Bear Enrage 16%.** `bearFormTiers` gives 27 in Bear Form and 16 in Dire Bear Form (from 40), applied to the equipment armor pools for Enrage's 10 s. Enrage's rows: 10 rage at once (effect 1, 100 tenths), 2 a tick (aura 24, 20 tenths), aura 94 (in combat), so the -1 s cast holds its rage. | `sim/druid/bear_form.go:44-64`, `enrage.go`; SpellEffect 683740-683742 | Aligned | - | - |
| G13 | **Primal Bite rank 3 (1238070) at 60.** The bear casts 1238070 29.1 times a fight at raid 60; the curated and synced rotation name only 1238070. | bear raid breakdown; `grep 12380` | Aligned | - | - |
| G14 | **Protection paladin band 30 Holy Shield.** `TestProtectionTakesHolyShieldAtThirty` and `TestCuratedLadderFirstIsSound` pass. | sim/leveling | Aligned | - | - |
| G15 | **Elixir of Holy Power in the paladin raid presets.** Retribution and Protection raid 60 carry `elixir_of_holy_power`; Holy does not. | consume lists | Aligned | - | - |
| G16 | **The Shred/Claw note.** The cat's APL note now says rank 1 Shred hits harder than Claw; unchanged rotation. | data/curated/apl/druid-feral.json | Aligned | - | - |

## Warriors: fresh re-audit

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Fury's build still spent a point on Improved Intercept and left Raging Blows empty.** Improved Intercept (Fury tier 4) "Reduces the cooldown of your Intercept ability by 5 sec"; the Fury rotation never Intercepts a raid boss, and the guide says you walk in. Raging Blows (Fury tier 3, 1 point) "Reduces the Rage cost of your Cleave and Whirlwind abilities by 3"; the rotation presses Whirlwind on cooldown (13.4 a fight). The engine models it (`sim/warrior/talents.go:267`): Whirlwind's rage per fight falls 336.9 to 297.0. The point moves from one to the other and every tier gate holds (25 Fury points below tier 4 instead of 24; 26 below Flurry, 31 below Bloodthirst). Same Shadow-build precedent as round 12 W6, which moved Fury's Arms points and did not look at this one. Raid 60 +0.46 to +0.60%, bare 60 +0.70 to +0.75%, beyond error at both seeds and factions. The guide's build line and its talent list now say so. | talents/warrior.json (nodes 105977, 105926); `-talents` runs; `TestFuryGuideBuildTakesRagingBlowsNotImprovedIntercept` (sim/leveling) failed first (Raging Blows 0, Improved Intercept 1), then passed; `_prose.test.ts` 15/15 | CONFIRMED-BUG (build defect) | old S fc0d455b9 (2026-10-07 warrior talent-search adoption; round 12 W6 looked at the build and missed it) | S cb1325223 |
| W2 | Arms raid build: Deflection 2 and Blood Craze 2 do nothing for a damage dealer behind the boss. Moving them into Improved Heroic Strike 3 (Arms casts Heroic Strike at 90+ rage) measured +0.2% at 3,000 iterations (731.8 → 733.2, ±0.8): not beyond error. | `-talents` runs, horde seed 1 | Held (within error) | - | - |
| W3 | Charge rage by rank, Bloodrage (10 plus 10 over 10 s; Improved Bloodrage +50% for Protection), Battle Shout 10 rage and 3 minutes, Charge off the global cooldown inside the shout's 1.5 s global, Arms in Battle Stance at the pull, Fury in Berserker Stance, Protection's Sunder 9 rage (15 less Improved Sunder 3 less Focused Rage 3) and Revenge 2. All match. | SpellEffect, SpellDuration, SpellCooldowns; breakdowns | Aligned | - | - |

## Feral: fresh re-audit

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Cat: Prowl at -2 s at the band's rank (6783 at 40, 9913 at 60), Ravage from Prowl with Poised, nothing else before the pull. Bear: Enrage at -1 s in Dire Bear Form (16%), held rage, Primal Bite rank 3, Faerie Fire (Feral) skipped under the raid's Faerie Fire. | breakdowns, APL dumps | Aligned | - | - |

## Before/after

Only W1 moves a number, and only through the guide build (the published BiS entries still carry the round-11 Fury talents until the refresh). Before is the round-12 guide build `30305013002000000-05153105022011501-200000000000000000`, after is `30305013002000000-05153105122010501-200000000000000000`, both through `-talents` on this lane. Seed 1 / seed 50001:

| Fury | horde % | alliance % | seed 1 horde / alliance |
|---|---|---|---|
| raid 60 | +0.60 / +0.40 | +0.46 / +0.55 | 764.4 → 769.0 / 761.6 → 765.1 |
| bare 60 | +0.70 / +0.74 | +0.75 / +0.68 | 285.7 → 287.7 / 281.2 → 283.3 |

No engine change: every other spec is identical to `before.tsv`.

## Files to regenerate

- `sim/request` TestRotationLadder golden for warrior-fury (the walk from the new guide build reaches Raging Blows below 60).
- `data/builds/1.60.1.70753/addon-builds.json` (Fury) and the web `leveling-builds.test.ts` expectation that reads it.
- Coach starter `data/curated/coach/warrior-fury.md` (the talent prose).
- Published BiS talents for Fury, bare and raid 60, at the refresh.

## Measure at refresh

- Fury picks and weights on the guide build with Raging Blows (rage-limited, so rage-side stats move slightly).

## For other lanes

- **Guides:** this lane edited `web/src/content/guides/warrior/fury.md` (build line, a Raging Blows bullet, the Improved Intercept sentence). If the guides lane touched the same file, merge around the talent section.
- **Hotfix:** no hunter amount touched here; the Arcane Shot spell-power change (E 05b72a3b5) still stands, so any new Arcane Shot or Serpent Sting amount should be checked against it with no spell power coefficient.

## Still open

- Arcane Shot's spell power (ruling 5): pending the owner's tooltip check with spell power up. If it moves, revert E 05b72a3b5.
- Arms Charges the boss at -2 s while every other melee waits for the tank. The client allows it and the order is right (G2); whether a raid Arms warrior should open with Charge at all is a realism call for the owner (worth 0.1%).
- Carried from round 12: Raptor Strike's off-hand penalty, normalized rage's off-hand factor, Lacerating Strikes' refresh, the one-a-minute proc rate, Holy Shield crit.

## Origin totals

1 confirmed bug: **old 1** (W1, S fc0d455b9). R12 0, landing 0, R9-R11 0, new-data 0.

Sources (read 2026-11-03): the 70753 layered raw tables (SpellMisc, SpellEffect, SpellDuration, SpellCastTimes); data/builds/1.60.1.70753/talents/warrior.json; cmangos-classic `SpellAuras.cpp` (aura 94) as cited in round 12.
