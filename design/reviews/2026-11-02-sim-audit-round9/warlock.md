# Sim-accuracy round 9: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons. Client 1.60.1.70707 with the hotfix push layered in (layered tables in scratch `r7/casters/raw70707`, fetched 2026-11-01 21:42 UTC; `python -m pipeline push-check --build 1.60.1.70707` re-run 2026-11-02 reports `moved=false`, so no newer push). Pin 40021374a, site origin/main 55a1212f7. Lane engine `lane/r9-warlock`, site branch `wt/r9-warlock`.

Method: spec-breakdown on the published band entries (`data/builds/1.60.1.70707/bis/warlock-*.json`), 10000 iterations, seeds 1 and 50001, horde troll and alliance gnome. "Before" is a detached origin/main on a detached 40021374a, each with its own `make simdb`.

Totals: 1 CONFIRMED-BUG (the Demonology prepull, owner ruling 4), fixed. It moves Demonology bare bands 20, 40, 50 and 60 by +0.1 to +0.4%; the raid build and band 30 do not move. No other warlock finding this round.

## Assigned item 1: Demonology's sacrifice (owner ruling 4)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **Demonology sacrificed the Imp and then summoned the Succubus at every band, and without Demonic Pact the summon cancels the sacrifice.** Demonic Pact's text: "Your Demonic Sacrifice effect is no longer cancelled by summoning a different Demon pet" (talent 105890, spell 425464). Only the raid build takes Pact. The bare builds at 30, 40, 50 and 60 take Demonic Sacrifice (it is Soul Link's prerequisite) but stop short of Pact, so Burning Shadow (18789) was applied once and up 0.0% of the fight at all four. Band 20 has no Demonic Sacrifice and still summoned an Imp a minute before the pull only to replace it. What a player with that build does is keep the Succubus out and never sacrifice: round 8 measured keeping the sacrifice instead at -9 to -24%. The APL could not say this because the engine ignored a prepull action's condition (`sim/core/apl.go`, the prepull `execute`), and Pact registered nothing the APL could read. Now: a prepull action runs only when its condition holds, Demonic Pact registers a permanent aura (425464), and the curated prepull gates the Imp summon and the sacrifice on `auraIsKnown(425464)`. Without Pact the prepull is Summon Succubus, then Soul Link. Tests: fork `TestPrepullSacrificesOnlyUnderDemonicPact` (failed: the sacrifice and the Imp were cast 5 times in 5 iterations without Pact; passes) and site `TestDemonologyNeverSummonsAwayItsSacrifice` in sim/cmd/spec-breakdown (failed at 30, 40, 50 and 60 bare with Burning Shadow applied but up 0.0%; passes at 20, 30, 40, 50, 60 bare and 60 raid). The builds keep their talents: the measured gain of the fix is +0.0 to +0.4%, and dropping Demonic Sacrifice would cost Soul Link, so nothing reaches the adoption rule's 1% bar. | talents/warlock.json (105890, 425464); spec-breakdown aura table, before: `Burning Shadow (18789) 0.0% uptime, 1.00 procs` at band 40 | CONFIRMED-BUG | E 17cc0ec01, E 5b2324d76 (fork APL copy), S ddea9cfcc |

No other spec's prepull carries a condition (checked every curated APL and every fork `ui/**/apls` preset), so the core change moves only Demonology.

## Assigned item 2: held adoption rows, re-measured

Lane engine, horde, raid preset, band 60, 10000 iterations.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Affliction: 1 Improved Drains to Improved Bane of Agony (`25520010135001050-0005200000000000000-0550004000000000`): 585.2 to 589.4 (seed 1), 585.3 to 589.6 (50001), +0.72% / +0.73%. Below the 1% bar. Held. | spec-breakdown `-talents` | Aligned (held) | - |
| A2 | Destruction: Immolate below Conflagrate: 649.7 to 650.8, 649.9 to 650.7, +0.17% / +0.12%. Held. | spec-breakdown `-rotation` (lines 5 and 6 swapped) | Aligned (held) | - |

## Assigned item 3: downranking (owner ruling 3)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Warlock fights at 60 are not mana-bound: Life Tap turns health into mana, and out-of-mana time is 0.0 to 0.7 s raid and 0.7 to 4.5 s bare. One rank down on the filler loses at both presets: Affliction Shadow Bolt 9 to 8 -1.4% raid / -1.7% bare; Demonology Shadow Bolt 9 to 8 -1.0% / -1.5%; Destruction Incinerate 3 to 2 -1.8% / -2.7% (3000 iterations, seed 1). No warlock candidate at 60. Below 60 a `-rotation` probe cannot test it: `request.RewriteRotationRanks` rewrites every rank to the band's top rank before the run, so the per-band probe belongs to the ranks lane's mechanism. | spec-breakdown `-rotation` | Aligned (no candidate) | - |

## Assigned item 4: crit tables

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Demons keep their own crit (Succubus melee 2.8% at 60 raid, Lash of Pain 6.8% plus Malevolence's 5% for Affliction), as `makeStatInheritance` records from 29 beta logs (#331). Brilliant Wizard Oil's +1% is spell crit on the warlock only. No warlock or demon consumable gives melee crit. For the crit lane: after the split, the warlock's spells should read only spell crit (Intellect, gear "spells"), and the Succubus' melee only melee crit (Agility at the warrior rate, `pet.go` warrior scaling). | spec-breakdown; `sim/warlock/pet.go` | Aligned (note for the crit lane) | - |

## Fresh re-audit

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Cast spells against the client and the Wowhead Forever server text (read 2026-11-02): Shadow Bolt r9 251, 1.8 a level, 0.857, 3 s, 370 mana ("246 to 274"); Incinerate r3 217, 0.714, 2.5 s, 325, +25% on Immolate (effect 1, dummy 25); Conflagrate r6 282, 0.429, 10 s, 255, consumes Immolate; Shadowburn r6 266, 0.429, 15 s, 365, a Soul Shard; Soul Fire r2 431 + 1.9 a level, 1.0, 6 s, 1 min, a Soul Shard; Lash of Pain r6 50, 0.429, 12 s, 160 ("55 / 50" with Unholy Power's 10%). Engine matches (`shadowburn.go`, `conflagrate.go`, conformance golden). | spellconst/warlock.json; nether.wowhead.com/forever/tooltip/spell/11661, /1293813, /18932, /18871, /17924, /11780 | Aligned | - |
| F2 | Amplify Curse (18288): 3 min, 30 s, one charge. Its client masks name Curse of Weakness (+50%) and Curse of Exhaustion (-20 flat), not Bane of Agony (family mask 1024), but the talent text and the server tooltip say "Curse of Weakness or Bane of Agony by 50%", and Blizzard's replies on #340 and #398 treat it as amplifying Agony. The engine raises the first Agony's base by 50%. The masks lag the text, as for Improved Drains in round 5. | SpellEffect 18288 effects 0 and 1; talents/warlock.json 105916; nether.wowhead.com/forever/tooltip/spell/18288 | Aligned (text over mask, single source plus Blizzard replies) | - |
| F3 | Troll Berserking: 10% cast and attack speed, 15 s (Blizzard's 27 October notes over the shipped 10 s), 3 min. 1.49 casts a fight, 11.5% uptime, which is 1.49 x 15 s over the fight. | `sim/core/racials.go:556`; spec-breakdown | Aligned | - |
| F4 | Soul Shards are not counted: Destruction casts Shadowburn about 10 times a fight and Demonology's Soul Fire runs in Decimation's window (no shard there, aura 256). A raid warlock carries a full shard bag into a boss, so ten a fight is a stock a player has. | spec-breakdown; Shadowburn tooltip | Aligned | - |
| F5 | Rotations against the builds: Affliction (Amplify Curse, Corruption, Agony, Siphon Life, Immolate, Shadow Bolt), Demonology (D1's prepull, Soul Link, Doom then Agony in the last minute, Corruption, Immolate, Decimation Soul Fire, Shadow Bolt), Destruction (Immolate at the pull, Corruption, Doom/Agony, Immolate, Conflagrate, Shadowburn, Incinerate). Each casts what its build takes; Life Tap below 10% mana; the potion and the Demonic Rune on each cooldown (round 7). | APLs; spec-breakdown at 60 raid | Aligned | - |
| F6 | Side effects of round 8 on warlocks: Ignition procs about 2.2 times a fight where Kindling Stave is worn; Chthonic Power once a fight, 8.3% uptime; Curse of Recklessness reaches the Succubus' melee as armor only; buffs do not reach the demons (`RefusePlayerBuffs`, Blizzard's hunter deep dive wording, `pet_buffs_test.go`). Uptime rows now read from the pull (round 8 C1): Burning Shadow and Demonic Pact at 99.8% on the raid build. | spec-breakdown | Aligned | - |
| F7 | spec-breakdown divides aura uptime by the nominal 180 s while the encounter varies its length, so a whole-fight aura reads 99.8 to 100.9%. A display quirk of the tool, not of the engine. | `sim/cmd/spec-breakdown/main.go:368` | Aligned (tooling) | - |

## Before/after

10000 iterations. Before: origin/main 55a1212f7 on 40021374a. After: wt/r9-warlock on lane/r9-warlock 5b2324d76. Seed 1 absolute, then % at seed 1 / 50001.

| Spec, band, preset | Horde s1 | Horde % (1 / 50001) | Alliance s1 | Alliance % (1 / 50001) |
|---|---|---|---|---|
| Demonology 20 bare | 48.2 to 48.4 | +0.41 / +0.41 | 49.4 to 49.5 | +0.20 / +0.20 |
| Demonology 30 bare | 80.5 to 80.5 | 0.00 / 0.00 | 80.4 to 80.4 | 0.00 / 0.00 |
| Demonology 40 bare | 142.0 to 142.4 | +0.28 / +0.35 | 141.9 to 142.4 | +0.35 / +0.35 |
| Demonology 50 bare | 208.8 to 209.1 | +0.14 / +0.19 | 208.0 to 208.6 | +0.29 / +0.38 |
| Demonology 60 bare | 326.1 to 327.1 | +0.31 / +0.28 | 325.6 to 326.5 | +0.28 / +0.37 |
| Demonology 60 raid | 654.9 to 654.9 | 0.00 / 0.00 | 649.2 to 649.2 | 0.00 / 0.00 |
| Affliction 40 bare, 60 raid | 135.1 to 135.1, 585.2 to 585.2 | 0.00 / 0.00 | - | - |
| Destruction 40 bare, 60 raid | 141.2 to 141.2, 649.7 to 649.7 | 0.00 / 0.00 | - | - |

Band 30 does not move: the Succubus summon and Soul Link cost the same either way, and the warlock's mana is full again by the pull at that level.

The fork suite `GOWORK=off go test -tags=with_db ./sim/warlock/... ./sim/core/...` passes; no fork golden moved. sim/request TestRotationLadder's warlock-demonology golden moves (the ladder's guide build has no Pact, so 688 and 18788 now show as zero casts from level 10 up and 425464 as unresolved); it is not regenerated here (the integrator's job at the pin). The smoke test's expected warnings gain 425464 for its talentless build.

## For other lanes

- Integrator: E 17cc0ec01 changes `sim/core/apl.go` (prepull actions honour their condition). No curated or fork preset other than Demonology's has a prepull condition, so no other spec moves; re-run the ladder golden for warlock-demonology at the pin.
- Ranks lane: no warlock downrank candidate at 60 (R1). Below 60 the band probe needs your mechanism, since `RewriteRotationRanks` restores the top rank.
- Crit lane: X1.

## Still open

- From earlier rounds, unchanged: Ignition's one-a-minute rate (round 8 C2), Improved Corruption with Shadow Mastery's periodic half, demon crit and the level-60 sheets (#331), shadow wands and Shadow multipliers (#419), Siphon Life tick crit (held for the owner).
