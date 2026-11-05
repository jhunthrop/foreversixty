# Sim-accuracy round 21: warlock (2026-11-05)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70790 with its live hotfix push layered (region 91). Site `wt/r21-warlock` from int/r21 (93c66ebc8); fork `lane/r21-warlock` from int/r21 (d5c708834). Sub-60 runs use the bare band gear from `$R/bis-frozen` with the talents the shipped leveling (or raid) order has at that exact level (`$R/warlock/tools/lvl.py`, read from `addon-builds.json`); raid-60 runs use the guide raid builds (`$R/rt`). The round-19 method rules were applied.

**Totals: 3 CONFIRMED-BUGs, all `old` (two also `guides-handoff`); the third is Demonic Brand, modeled in the engine (E 0994b6f8c, no published number moves). The first two, both fixed in the curated APLs (S 8c345ae68, wording S e61642ae0; fork copies E 60e02285b, E d468e18f6). Raid 60 is unchanged to the decimal on all 12 runs.** The Shadowburn shard seed is decided as aligned with a card-note fix. No live log exists yet (see "Logs").

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Destruction held Conflagrate on cooldown before Shadow and Flame.** <br>**What was wrong:** the card cast Conflagrate on cooldown whenever Immolate was up, while destruction.md:128 said to wait until the burn had done most of its work. Before Shadow and Flame, Conflagrate always consumes Immolate (`conflagrate.go:87`; client 1293817 / 1293818 / 17962 "dealing N Fire damage and consuming your Immolate effect"), so on cooldown it throws away the rest of the burn and forces a recast. <br>**Evidence** (bare, leveling order at each level, 10,000 iterations, seeds 1 / 50001, both factions): holding until Immolate has 6 s or less left against on cooldown: 34 +1.7 / +1.7% (orc and human), 35 +1.5 to +1.7%, 36 +0.9 to +1.0%, 37 +0.4%, 38 -0.1 to +0.1%; from 39 (first Shadow and Flame point) on cooldown wins (hold -1.0% at 39, -0.8% at 42, -1.8% at 43), so the gate switches on the talent. Raid order (Conflagrate 32, Shadow and Flame 37): 32 +1.3 to +1.4%, 33 +1.2 to +1.5%, 34 +0.7 to +0.8%, 35 +0.2 to +0.4%, 36 -0.3%. A 4 s threshold was tried too and is no better (`out/conflag-10k.tsv`). <br>**Fix:** Conflagrate splits into two rows: on cooldown while Immolate is up once `auraIsKnown(426316 tag 1)` (Shadow and Flame), otherwise only with Immolate at 6 s or less; identical from the talent on, 60 included. destruction.md:128 and both notes say the same thing; the Immolate note takes guide round 19's wording (WARLOCK-R19-1). `TestWarlockDestructionHoldsConflagrateUntilShadowAndFlame` was red (14.8 casts at 36), then green. | `out/conflag-10k.tsv`, `out/gated-10k.tsv`; guide r19 WARLOCK-R19-3 | CONFIRMED-BUG | old, guides-handoff (round 20 X4 could not judge Conflagrate below 40 on the band-30 talents) | S 8c345ae68, S e61642ae0 |
| S2 | **Affliction kept Amplify Curse for a last minute below 60.** <br>**What was wrong:** the card pressed Amplify Curse only before the Bane of Agony of a fight's last minute. That is the 60 rule (the last minute is Agony's once Doom has fired); below 60 there is no Doom, so a 3-minute cooldown waited for a minute that a leveling mob never has. <br>**Evidence** (bare, 10,000 iterations, seeds 1 / 50001, both factions, leveling order 30/40/50/58 and raid order 26/30/40/50/58): before each Agony it is ready for gains +0.03 to +0.09% at every level on both seeds (never a loss); a 300 s fight gets 2 casts instead of 1. Tie-level on a boss, so the web rule decides: Classic and Forever warlocks press it with Agony whenever it is up while leveling. <br>**Fix:** the Amplify line fires when Doom is not known, or in the last minute when it is (`spellIsKnown(603)`, registered only at 60); 60 is unchanged. Note and affliction.md's leveling paragraph follow. `TestWarlockAfflictionAmplifyCurseBeforeEachAgonyBelow60` was red (1 cast at 50, 300 s), then green; 60 stays at 1. | `out/amp-screen.txt`, `out/amp-10k.txt` | CONFIRMED-BUG | old, guides-handoff (guide r19 "Sims requested") | S 8c345ae68, S e61642ae0 |
| S3 | **Soul Shards and Shadowburn at 30-43.** The 70790 client confirms the reagent: SpellReagents 20433, 20420, 20421, 20353, 20355, 20511 give every Shadowburn rank (17877 to 18871) one Soul Shard (6265); the live hotfix push changes no warlock row (layered SpellReagents, SpellEffect, SpellMisc, SpellCooldowns, SpellCategories, SpellClassOptions all match). The engine keeps no shard. Decision: no shard model and no APL gate. On a boss the card casts Shadowburn about 8 to 12 times in 180 s, a stock a warlock carries (round 9 F4; in Forever Drain Soul can give a shard on any target it damages), and shards never run out inside one fight, so a model changes no number. A finisher-only gate costs 1.3 to 3% at 34-44 (`out/conflag-screen.txt`, `sbfinish`). What was wrong was the card's silence: its note now says each cast costs a shard, with a 20% a point refund from Shadow and Flame, and destruction.md:128 says to cast it on cooldown on a dungeon boss while you have shards to spend. | `wago/SpellReagents.layered.csv`; spelltips 17877/18871 | Aligned (note fixed) | - | S 8c345ae68, S e61642ae0 |
| S4 | Round 20's warlock log items (Eureka! on Incinerate, demon crit, Soul Siphon's count, Wrack and Doom at 60, the level-50 Succubus weapon): no live log exists (see "Logs"). | `$R/platform/logs.md` | OPEN | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Raid 60 reproduces `before.tsv` to the decimal (753.1 / 753.6 alliance, 750.6 / 750.8 horde), after the S2 change too. | `out/after60.tsv` | Aligned | - | - |
| A2 | Amplify Curse's client mask (SpellEffect 695823: aura 108, misc 8, word 0 0x8000) names only Curse of Weakness; Era 1.15.9.70576 has 0x8400 (Agony too). Forever narrowed the mask, but the tooltip says Bane of Agony and Blizzard's replies on forever-bugs #340 and #398 treat Amplify as amplifying Agony (round 9 F2). Misc 8 multiplies effect base points only, which is what `curses.go:98` does. Rechecked, no change. | `wago/SpellEffect.csv`, `era/SpellEffect.csv`; round 9 F2 | Aligned | - | - |
| A3 | Spell constants (cost, cast, cooldown, GCD, level, duration, roll and coefficient at every rank and preset level) are all `match` in the fork's 70790 conformance golden apart from the documented control durations (Death Coil, Shadowburn's shard marker, Curse of Recklessness). Life Tap's amount (rank flat + 1 a level + Spirit, x Improved Life Tap) matches the 70790 tooltips (1454 to 11689). | `sim/core/testdata/conformance/warlock.golden.md`; spelltips | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D2 | **Demonology filler, Shadow Bolt against Searing Pain** (guide-accuracy round 20 WARLOCK-R20-4). Leveling order at each level, bare band gear, 10,000 iterations, seeds 1 / 50001, both factions; the card as shipped against Shadow Bolt replaced by Searing Pain at its learned rank (cast under its own id: 5676 at 18, 17921 at 42). Searing Pain loses at every level asked and every rank-up level beside it: 18 -9.6 / -8.7% (horde / alliance), 19 -9.1 / -8.6%, 20 -8.1 / -7.7%, 21 -7.6 / -7.3%, 26 -7.6 / -7.0%, 27 -7.4 / -6.8%, 34 -1.8 / -1.3%, 35 -2.2 / -1.8%, 42 -4.8 / -4.4%, 43 -4.8 / -4.2%, 50 -3.8 / -2.8%, 51 -3.5 / -2.7%, 58 -5.1 / -3.9%, 59 -6.0 / -4.7% (seeds agree within 0.2). Searing Pain wins more per second on base damage at the rank-up levels, but it costs more mana per point of damage, and the crit-free leveling build has nothing that favours it; neither filler runs out of mana before 50 (1 to 2 s out of mana at 50 to 59 for both). The card's "It beats Searing Pain" holds at every level, so the card stays. Levels 18 and 19 have no published bare band-10 entry, so they ran on the band-20 gear (a scratch binary that only widens `-level`'s band check; not committed): both fillers on the same gear. The note takes guide round 20's R20-1 wording (Master Demonologist with the Succubus out; the sacrifice's 15% only with Pact in the raid build). | `out/demo-sp-10k.tsv`, `out/demo-sp-screen.txt`; guide r20 WARLOCK-R20-4 | Aligned | - | S (note wording only, see heads) |
| D3 | **Demonic Brand was not modeled** (MythicSim comparison lead). <br>**What was wrong:** the engine skipped the talent (`talents.go`, "no resolvable number"), so a build that takes it and casts Searing Pain lost the demon's branded hits and kept Searing Pain's full threat. <br>**Client 70790 (live hotfix push layered; no row touched):** TraitDefinition 135427 curves give the threat reduction (curve 83195: -17/-33/-50%, effect 0, aura 108 misc 2 on Searing Pain's family bit 0x100) and the charges (curve 103069: 2/4/6, effect 1, aura 107 misc 4 on the brand). The brand is 1293696 (10 s; ProcTypeMask 139944: the target's taken melee, ranged and harmful spells). The consumed hits are 1293697 (school Shadow) and 1293698 (Fire); both carry 0 in SpellEffect (a server script), and their Wowhead Forever tooltips state the amount the talent text gives: ((60 - 26) x 1.5 + 14) to (... + 17) + 0.078 x Shadow (or Fire) spell power, 65 to 68 plus spell power at 60. 1293699, the row the lead cited, is school Physical with 1 damage, a stub no pet uses; it does not disagree with the text. <br>**Fix:** Searing Pain's threat x0.83/0.67/0.5; a landed Searing Pain brands the target with 2/4/6 charges for 10 s; each landed attack of the demon on the branded target spends a charge on that hit (Shadow, or Fire for the Imp; the demon's spell hit and crit, the demon's own multipliers). `TestDemonicBrandDamageFollowsTheTalentText` and `TestDemonicBrandSearingPainBrandsAndThePetSpendsIt` failed to build, then pass. The amount is one source (the client text and its tooltip); no Forever log has a branded demon yet. <br>**Effect:** no guide build and no published band entry takes Demonic Brand (Demonology field 13 is 0 in all six guide strings and every BiS entry), and no card casts Searing Pain, so every published number is unchanged. Illustration (raid 60 horde, 3,000 iterations): the raid build with Brand 3 in place of Improved Health Funnel 2 and Demonic Aegis 1 and Searing Pain as the filler: the brand hits are 67 DPS (10.5%), but the whole is 640.2 against 724.7 for the guide (Shadow Bolt still wins the filler, D2), so no build or card changes. | `wago/TraitDefinitionEffectPoints.csv`, `CurvePoint.csv`, `SpellEffect.layered.csv`, `SpellMisc.layered.csv`, `SpellAuraOptions.csv`; nether.wowhead.com/forever/tooltip/spell/1293695, 1293697, 1293698, 1293699 (2026-11-05) | CONFIRMED-BUG | old (every round listed Demonic Brand as unmodeled; round 21's integrator lead) | E 0994b6f8c |
| D1 | Raid 60 reproduces `before.tsv` (734.1 / 735.1 alliance, 724.3 / 724.0 horde). No hotfix row touches a Demonology spell or talent; the card (prepull sacrifice gated on Demonic Pact, Decimation Soul Fire only in the leveling build) is unchanged from round 20's audit (D1 to D4). | `out/after60.tsv`; layered tables | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | S1 and S3 above. | | | | |
| X2 | Raid 60 reproduces `before.tsv` (795.3 / 794.1 alliance, 792.2 / 792.6 horde): at 60 the raid build has Shadow and Flame, so the new Conflagrate row is the old one. | `out/after60.tsv` | Aligned | - | - |
| X3 | Soul Fire for Destruction at 60 (never cast by the card): a pre-pull Soul Fire is +0.2% at 3,000 iterations and would land on the boss before the tank, and Soul Fire on cooldown in the rotation is -1.2%. The card stays without it. | rotation variants `rot/destro-sfpre.json`, `destro-sfcd.json` | Aligned | - | - |
| X4 | Bane of Havoc in the raid build does nothing on one target but is Incinerate's prerequisite (round 14 X3); it stays. | round 14 X3 | Aligned | - | - |

## Follow-up after the first merge (integrator, guide-accuracy round 20)

The lane merged int/r21 (spellIsKnown now honours RequiredLevel) and took guide round 20's warlock wording where it touches the same lines, so the two landings do not fight: the Demonology Shadow Bolt note (R20-1), the Destruction Immolate note (R20-3: "never runs out ... until Shadow and Flame is full, recast it whenever Conflagrate eats it", which agrees with the held Conflagrate), affliction.md's Wrack sentence (R20-2) and the 70790 source links. D2 above answers R20-4: no card change.

## Logs

Last looked 2026-11-05 03:15 UTC (4 h 15 min after launch): `api.foreversixty.gg/v1/reports/recent` lists only z2opt2x2gs5u (2026-10-20); foreverlogs.gg `reports/public` is empty; Warcraft Logs v1 `/v1/zones` lists no Forever zone on www (66), classic (57), fresh (21) or vanilla (29). The platform lane's harvest (`$R/platform/logs.md`, 02:55 UTC) agrees. The logs still needed: a 40+ Gnome Destruction (Eureka! on Incinerate), 40+ Affliction (Soul Siphon, Wrack), level 60 (Doom, demon crit) and a level-50 Succubus.

## Before/after

Raid 60, 10,000 iterations, seeds 1 / 50001: unchanged to the decimal on all 12 runs (absolute at seed 1, horde: Affliction 750.6, Demonology 724.3, Destruction 792.2).

Bare, leveling order, horde (alliance within 0.2 points of the same percentage), seeds 1 / 50001, Destruction (S1): 34 107.4 to 109.2 (+1.7%), 35 108.5 to 110.2 (+1.6%), 36 115.7 to 116.7 (+0.9%), 37 117.8 to 118.3 (+0.4%), 38 -0.1 / +0.1%, 39 to 60 unchanged. Raid order: 32 +1.3%, 33 +1.3%, 34 +0.7%, 35 +0.3 / +0.4%, 36 -0.3%, 37 on unchanged.

Affliction (S2): +0.03 to +0.09% at 26 to 59 on both orders and seeds; 60 unchanged.

## Files to regenerate

- addon-data.json and Data.lua: the Destruction card gets a second Conflagrate row and the new Conflagrate, Shadowburn and Immolate notes; the Affliction Amplify Curse line and note change.
- sim/request ladder goldens (TestRotationLadder): warlock-destruction at 38 (Conflagrate cast counts) and warlock-affliction at 30, 38, 40 and 50 (Amplify Curse casts).
- Fork: the conformance goldens `warlock.golden.md` (and mage, paladin, priest, shaman) read stale on this arm64 host before and after this lane's change (amd64 rounding plus the platform's summon-level fix); not regenerated here. `TALENTS.md` (hand-written) marks Demonic Brand modeled.
- Any coach starter or apl-sync output that copies the warlock APLs (the `sim/request/apl` and fork `ui/warlock/apls/forever_*` copies were synced by hand).

## Measure at refresh

- Carried: the Affliction bare 40, 50 and 60 BiS entries should carry the Malevolence string (`25522300132001351`).

## For other lanes

- Nothing in this lane moves DPS outside warlock; no shared code was touched.
- Seen in passing (mage lane or integrator): `web/src/lib/sim/rotations.test.ts` fails 2 mage-frost tests (the Frostbolt filler note) on int/r21 with or without this lane's changes.
- A stray astro dev server from `.claude/worktrees/r18-rogue/web` is still listening on port 4391 (not started by this lane; left running).

## For the guide lane

- Done here: destruction.md:128 (Conflagrate before and after Shadow and Flame; Shadowburn on a dungeon boss), affliction.md "Levels 20 to 60" (Amplify Curse from 29), and the card notes above, which also close WARLOCK-R19-1 with guide round 19's Immolate wording. The wago 70790 source-link fix (WARLOCK-R19-2) is guide round 19's and was not touched.

## Guide text not done

- None.

## Guide review

- wow-player, on the rendered Destruction and Affliction "Levels 20 to 60" paragraphs (dev build of this branch, 1440 and 390) and on the card notes as text (they render after the regen). First pass: CHANGES (Shadow and Flame's 4% a point and the 20% a point shard refund, clearer Amplify wording, "for 20 seconds"). All of it was taken. Second pass: SHIP on every note and the Affliction paragraph. The one item left open was the Destruction paragraph's readability at 390; I checked the 390 capture and it reads cleanly. `_prose.test.ts` passes.

## Rendered-card follow-up (integrator)

The card shows only the Conflagrate row the guide's raid build lights, so its note now carries the hold as well: "once Immolate has 6 seconds or less left until you take Shadow and Flame, then on cooldown while Immolate is up". No semicolon: `clauseOf` (`rotation-lists.ts:81`) cuts a card clause at one. The page's rotation bullet says the same. The two row conditions were re-read and already express this (hold before `auraIsKnown(426316)`, on cooldown after). The engine at 1 to 4 points keeps Immolate 20% a point, so "from the first point, on cooldown" stands (S1's numbers); the Shadow and Flame talent bullet now gives the per-point values. Notes and prose only: no number moves.

## Still open

- S4: the log-gated items above.
- Pending owner, not reopened: level-60 trainer book ranks, Siphon Life tick crit, the Demonology Horde race, Heart of the Lion 280/276, warlock Orc.

## Sources

- Client 1.60.1.70790: wago.tools SpellReagents, SpellEffect, SpellClassOptions, SpellMisc, SpellCooldowns, SpellCategories (`https://wago.tools/db2/<Table>/csv?build=1.60.1.70790`, read 2026-11-05), layered with the live hotfix push through `data/pipeline/hotfix_push` (`$R/warlock/tools/layer.py`); `data/builds/1.60.1.70790/` spelltips, talents and addon-builds (warlock).
- Era 1.15.9.70576 SpellEffect (wago.tools) for Amplify Curse's old mask.
- forever-bugs #340 and #398 (via round 6 and round 9 warlock reports).
- Guide-accuracy round 19 warlock report (origin/land/guides-r19).
- Lane outputs under `$R/warlock/out/` and rotation variants under `$R/warlock/rot/`.
