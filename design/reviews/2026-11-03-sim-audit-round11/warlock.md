# Sim-accuracy round 11: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons, bands 20 to 60, bare and raid presets. Client 1.60.1.70707 with the hotfix push layered in (the layered tables in scratch `r9/casters/raw`, fetched 2026-11-01 21:42 UTC; data/builds/1.60.1.70707/hotfix-push.json changes only Spell text for warlock rows). Pin cfcb76e80, site origin/main 8f047a62f. Lane engine `lane/r11-warlock`, site branch `wt/r11-warlock`.

Method: every spell the three APLs cast and every talent the published builds take (data/builds/1.60.1.70707/bis/warlock-*.json, decoded against talents/warlock.json), read against the client rows (SpellEffect, SpellMisc, SpellAuraOptions, SpellCooldowns, class masks through SpellClassOptions set 5) and the engine code; spec-breakdown at every band and preset; the raid talent search re-run at the pin. Round 1 to 10 reports were grepped for each spell first; the rows below say where a spell was already settled.

Totals: 3 CONFIRMED-BUGs, all fixed. Origin: 3 old, 0 R9, 0 R10, 0 new-data. The bane and curse split moves every bare band (the raid preset already carries the raid's Curse of the Elements and does not move).

## Cross-cutting: banes and curses

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **A bane removed Curse of the Elements; Forever gives banes and curses separate slots.** Forever renamed Agony and Doom to banes and split them from the curses. Every bane's client text reads "Only one Bane per Warlock can be active on any one target" (Bane of Agony 980, 1014, 6217, 11711 to 11713; Bane of Doom 603; Bane of Havoc 1225228). Every curse's reads "Only one Curse per Warlock can be active on any one target" (Curse of the Elements 440892, 1311676, 1311677, 1311680; Recklessness 704 to 11717; Weakness; Tongues; Exhaustion 18223). Guides say the same: a warlock keeps one bane and one curse on a target. The engine kept Era's single curse slot (`ActiveCurseAura`), so Bane of Agony or Doom removed Curse of the Elements and no rotation could keep both. A 2026-10-07 measurement (design/reviews/2026-10-07-warlock-ua-coe.md) compared the Elements *instead of* Agony and found it lower. Now Agony, Doom and Havoc take a bane slot and the curses keep the curse slot. Test `TestABaneAndACurseHoldSeparateSlots` failed first ("Bane of Agony removed Curse of the Elements") and passes; the old test that encoded the single slot (`TestCurseOfTheElementsGivesWayToAnotherCurse`) is replaced by it. | Spell.csv Description_lang for every rank above; mobalytics.gg/wow-forever "warlock class overview" and conquestcapped.com "WoW Forever warlock overview" (read 2026-11-02): one Bane and one Curse on a target at once; `sim/warlock/curses.go` | CONFIRMED-BUG (client text on every rank, two guides) | old: Era's single slot (wowsims 10cd22ddf, 2024), kept when the Forever curse was registered (7c4274241, 2026-10-07) | E 360dcb037 |
| B2 | **The warlock's own Curse of the Elements stacked with the raid's.** The raid debuff (`core.CurseOfElementsAura`) uses exclusive effects; the warlock's own rank multiplied the target's damage taken and removed resistance on top of it, so a target with both took 1.10 x 1.10. No published rotation cast the curse, so nothing published moved; a /sim rotation that cast it did. Both now use `core.CurseOfElementsEffects` (the same exclusive categories, strongest wins), and the warlock does not cast it over a curse already in force. Test `TestOwnCurseOfTheElementsDoesNotStackWithTheRaidCurse` failed first (x1.10 on every school, 75 more resistance) and passes. | the curse is one debuff (same spell family, same "Only one Curse" rule); `core/debuffs.go`, `sim/warlock/curses.go` | CONFIRMED-BUG (latent) | old (7c4274241, 2026-10-07) | E 360dcb037 |
| B3 | With the slots split, each rotation casts Curse of the Elements once, before its bane (five minutes, so once a fight). It raises the magic damage the target takes from the warlock and the Succubus' Lash of Pain. In the raid preset the raid's curse is already up, the engine does not cast over it, and the raid numbers are byte-identical. Gains are in "Before/after"; every bare band gains beyond error and above 1%. The guides' rotation and "Brings" lines say so. | spec-breakdown; `data/curated/apl/warlock-*.json`; fork copy synced (E ccdd59f2f) | follow-on of B1 | - | S ac4937071 |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Cast spells against the client: Shadow Bolt r9 (251 + 1.8 a level, variance 0.108, 0.857, 3 s, 370), Corruption r6 (57 every 3 s for 6 ticks, 0.20 a tick, instant with Improved Corruption 5), Bane of Agony r6 (46 every 2 s for 12 ticks at 0.133, ramping 0.5/1/1.5 by thirds, the server text "dealt slowly at first"), Siphon Life r4 (41 every 3 s for 30 s, 0.05), Immolate r7 (146 + 2.3 a level at 0.20, then 52 every 3 s for 5 at 0.13), Life Tap r6 (420 + 1 a level + Spirit, x1.2 with Improved Life Tap 2). Engine tables are generated from the same rows. | SpellEffect rows; `spellconst_damage_test.go` | Aligned | - | - |
| A2 | Talents of the raid build: Suppression 5 (aura 54 and 55, +5% hit, mask 0), Improved Corruption 5 (-2 s, +10% periodic), Pandemic 3 (crit damage bonus +100% on Corruption, Agony, Doom, Drains, Siphon Life, Wrack: 2.0x), Malevolence 5 (aura 107 on the Shadow list and Lash of Pain), Shadow Mastery 5 (direct half on Shadow Bolt, Shadowburn, Corruption, Death Coil; periodic half on the banes, Corruption, Drains, Siphon Life, Wrack), Improved Shadow Bolt 5 (17794 aura 270, 20%, 12 s, no SpellAuraOptions row so no charges; the warlock's own Shadow damage only), Bane 5, Ruin 4 (Ruin's mask covers Shadow Bolt), Unholy Power 5, Amplify Curse (3 min; text and replies reach Agony, base only). Improved Drains 1 does nothing for this rotation (no drain is cast); see H1. | mask resolution (`r10/warlock/py/mask.py`); talents.go | Aligned | - | - |
| A3 | Nightfall (bare 40 to 60 builds): 2% a rank on Corruption, Drain Life, Drain Soul and Wrack ticks; Shadow Trance 17941 is 10 s (DurationIndex 1), -100% Shadow Bolt cast time, consumed by the instant bolt, a bolt already casting keeps its cast time. | 18094, 17941 rows; `talents.go:311` | Aligned | - | - |
| A4 | Bands 20 to 50 (bare): Shadow Bolt, Corruption, Agony, Immolate, Siphon Life from 40, Amplify Curse from 30, the Succubus at every band; ranks are the band's learned ranks. Out of mana at most 0.7 s. | spec-breakdown per band | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Prepull: Summon Imp, Demonic Sacrifice (Forever: the Imp's sacrifice is +15% Shadow, "the opposing aspect"), Summon Succubus (Demonic Pact keeps the sacrifice: "no longer cancelled by summoning a different Demon pet"), Soul Link. Master Demonologist with the Succubus is +10% Shadow for both; Soul Link +3% damage for both; Demonic Knowledge +100% of level (60) spell power for both while a demon is out; Unholy Power +10% demon damage; Improved Sayaad +10% a rank on Lash of Pain's base (misc 8, all effects, so base points only); Fel Vitality +15% mana. | 18788, 23785, 19028, 412732, 425464, 18754, 18769 rows and text; `talents.go` | Aligned | - | - |
| D2 | Decimation 2: Soul Fire cooldown -90% (60 s to 6 s), below 35% a Shadow Bolt grants 440873 (10 s, -40% Soul Fire cast, no shard) and the bolt deals 6% more. Soul Fire r2 (431 + 1.9 a level, coefficient 1.0, 6 s less Bane's 2 s, so 2.4 s under Decimation). Bane of Doom (1742 at coefficient 4.0 after 60 s, a non-channelled DoT, so it crits by the ruling). | 440870, 17924, 603 rows | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **A missed Conflagrate consumed Immolate.** Conflagrate's only client effect is its Fire damage (18932 effect 0); consuming Immolate is part of the hit, so a miss deals nothing and takes nothing. The engine consumed Immolate, and rolled Shadow and Flame's keep chance, on a miss too. With Shadow and Flame 5 (bands 50 and 60) Immolate is always kept, so only band 40 (Shadow and Flame 1, a 20% keep) loses Immolates, at about one Conflagrate miss in twenty. Test `TestAMissedConflagrateLeavesImmolate` failed first ("a missed Conflagrate consumed Immolate") and passes. | 18932 SpellEffect (one effect); `sim/warlock/conflagrate.go` | CONFIRMED-BUG (client row; the server consumes on the hit) | old (wowsims 843a513ae and 8aa55f434, 2024) | E 18cfcb100 |
| X2 | Incinerate r3 (217 + 1.4 a level, 0.714, 2.5 s less Bane's 0.5, +25% while Immolate is up: effect 1 bp 25), Conflagrate r6 (282 + 1.3 a level, 0.429, 10 s; Forever ranks 1293817 at 25 and 1293818 at 32), Shadowburn r6 (266 + 1.8 a level, 0.429, 15 s), Fire and Brimstone 3 (+25% Conflagrate crit), Agonizing Flames 3 (+10% Searing Pain crit; +10% damage on Shadow Bolt, Searing Pain, Conflagrate, Immolate, Incinerate, Rain of Fire, Shadowburn, Soul Fire, the engine's Destruction flag), Shadow and Flame 5 (+10% Shadow for 20 s on a Conflagrate hit, +10% Fire on a Shadowburn hit, Immolate kept, the Shadowburn shard refunded), Aftermath 4 (+40% Immolate initial damage), Cataclysm 3 (10%), Ruin 5. Bane of Havoc is taken and never cast single target, rightly. | rows and text; incinerate.go, conflagrate.go, shadowburn.go, talents.go | Aligned | - | - |
| X3 | Bands 30 to 50: Shadowburn from 30 (no Shadow and Flame at 30, so each costs a shard; the sim carries no reagents), Incinerate and Conflagrate from 40, Shadow Bolt rank 4 at 30 to 39 by the curated rank choice. | spec-breakdown; data/curated/rank-choices.json | Aligned | - | - |

## Demons

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Succubus melee against a +3 boss: 2.0 s swing, 9% base miss less 8/17 of the warlock's hit (#394), 6.5% dodge, 40% glancing, no crit after suppression (round 10's 0.7%). Lash of Pain r6: 50 flat, 0.429, 12 s (SpellCooldowns CategoryRecoveryTime 12000), 1.5 s GCD, 160 mana; 6.4% crit plus Malevolence. Demons refuse player buffs (round 3), so the raid's Windfury, Battle Shout and the rest do not reach them. | 11780 rows; `pet.go`, `succubus.go` | Aligned | - | - |
| P2 | Round 10's open rows re-checked for new evidence. forever-bugs #331 has had no comment since 2026-10-14; no warlock issue on the tracker was updated after 2026-10-31; api.foreversixty.gg/v1/reports/recent holds three October reports, none level 60. The demon sheets (Era's figures), the warlock's plain crit reaching demons, and the demon crit values at 60 stay open. | gh issue view 331; gh search issues; the API (read 2026-11-02) | OPEN: a level-60 or launch log with a Succubus | - | - |
| P3 | No Felguard: Summon Felguard is not a Forever trainable (trainables/warlock.json has Imp, Voidwalker, Succubus, Incubus, Felhunter). | trainables/warlock.json | Aligned | - | - |

## Held under the adoption rule

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Affliction raid search at the pin (cfcb76e80, 3,000 final iterations, seed 7, keeps from raid-keeps.json): the variant that keeps every unmodeled guide talent is 1 Improved Drains to Improved Bane of Agony, +4.1 DPS (+0.72%, ±0.8). Under 1%. The overall winner (+1.9%) drops Grim Reach (range, unmodeled). Held. | `talent-search -spec warlock-affliction -preset raid` | Aligned (held) | - | - |
| H2 | Destruction's Immolate line below Conflagrate, 10,000 iterations, horde: 636.9 to 637.6 (seed 1), 637.6 to 638.5 (seed 50001), +0.11% / +0.14%. Held. The raid search finds no build beyond error. Demonology's raid search: the guide is best. | spec-breakdown `-rotation` (lines 5 and 6 swapped); talent-search | Aligned (held) | - | - |

## Gear and races

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | Items with an effect beyond plain stats in any published warlock band: Frozen Heart of the Mountain (Chthonic Power), Kindling Stave (Ignition), Rune of the Guard Captain (Authority: attack power, nothing for a warlock or its demon), Robe of the Void (Heal Pet), Mark of the Chosen and Uther's Light (both "when struck"), Cookie's Stirring Rod (cooking speed), the band-30 Archmagister's Wisdom (experience). None but the first two touches DPS, and those were settled in rounds 8 and 9. No set bonus is in a published set. | ItemXItemEffect, ItemEffect, SpellEffect | Aligned | - | - |
| G2 | Troll Berserking (20554: +10% cast and attack speed flat in Forever, 15 s, 3 min) and Gnome Eureka! (round 10 F5). | 20554 rows; `core/racials.go` | Aligned | - | - |

## Before/after

10,000 iterations, seeds 1 and 50001. Before: the shared baseline (`r11/sb-before` from `r11/site-before/sim`, origin/main 8f047a62f on cfcb76e80). After: wt/r11-warlock on lane/r11-warlock 18cfcb100. Seed 1 absolute, then % at seed 1 / 50001.

| Spec, band, preset | Horde s1 | Horde % (1 / 50001) | Alliance s1 | Alliance % (1 / 50001) |
|---|---|---|---|---|
| Affliction 20 bare | 53.4 to 54.7 | +2.43 / +2.43 | 53.6 to 55.0 | +2.61 / +2.61 |
| Affliction 30 bare | 86.6 to 90.1 | +4.04 / +4.16 | 86.7 to 90.5 | +4.38 / +4.26 |
| Affliction 40 bare | 132.2 to 140.1 | +5.98 / +5.98 | 132.5 to 140.1 | +5.74 / +5.81 |
| Affliction 50 bare | 192.6 to 207.5 | +7.74 / +7.84 | 194.1 to 209.3 | +7.83 / +7.88 |
| Affliction 60 bare | 285.6 to 307.9 | +7.81 / +7.84 | 285.0 to 306.6 | +7.58 / +7.61 |
| Affliction 60 raid | 571.0 to 571.0 | +0.00 / +0.00 | 569.3 to 569.3 | +0.00 / +0.00 |
| Demonology 20 bare | 47.4 to 48.4 | +2.11 / +2.11 | 48.6 to 49.8 | +2.47 / +2.47 |
| Demonology 30 bare | 79.0 to 82.2 | +4.05 / +4.05 | 78.9 to 82.1 | +4.06 / +4.18 |
| Demonology 40 bare | 140.1 to 148.1 | +5.71 / +5.79 | 140.0 to 148.1 | +5.79 / +5.86 |
| Demonology 50 bare | 206.0 to 222.3 | +7.91 / +7.96 | 205.4 to 222.0 | +8.08 / +8.03 |
| Demonology 60 bare | 322.2 to 347.1 | +7.73 / +7.76 | 321.5 to 347.1 | +7.96 / +7.99 |
| Demonology 60 raid | 640.4 to 640.4 | +0.00 / +0.00 | 635.6 to 635.6 | +0.00 / +0.00 |
| Destruction 20 bare | 52.0 to 53.2 | +2.31 / +2.31 | 52.2 to 53.4 | +2.30 / +2.30 |
| Destruction 30 bare | 87.7 to 91.1 | +3.88 / +3.88 | 88.1 to 91.7 | +4.09 / +4.20 |
| Destruction 40 bare | 138.0 to 147.2 | +6.67 / +6.44 | 138.8 to 148.0 | +6.63 / +6.70 |
| Destruction 50 bare | 233.9 to 251.5 | +7.52 / +7.62 | 234.8 to 253.4 | +7.92 / +7.83 |
| Destruction 60 bare | 350.7 to 378.3 | +7.87 / +7.81 | 353.9 to 382.0 | +7.94 / +7.91 |
| Destruction 60 raid | 636.9 to 636.9 | +0.00 / +0.00 | 634.6 to 634.6 | +0.00 / +0.00 |

Where it comes from: Curse of the Elements beside the bane (B1, B3) is all of it except Destruction 40, where the missed-Conflagrate fix alone (X1, the new engine with the old rotation) is 138.0 to 139.2 and 138.3 to 139.4 (+0.87% / +0.80%, horde). The raid preset is byte-identical: the raid already carries the curse and the warlock does not cast over it.

The fork suite `GOWORK=off go test -tags=with_db ./sim/warlock/... ./sim/core/...` passes and no fork golden moved. Site: `go test ./request -skip TestRotationLadder`, `make apl-check`, data/tests/test_apl.py, test_addonrotation.py, test_icons_wowhead.py, tools/coach/test_persona_starter.py and the guides vitest pass (src/pages/guides/_spec.test.ts needs generated data this worktree has not synced; environmental). sim/request TestRotationLadder: the three warlock goldens move at the bare rungs; not regenerated (integrator).

## For other lanes

- Integrator: E 360dcb037 adds `core.CurseOfElementsEffects` and `core.CurseOfElementsOn` in `sim/core/debuffs.go` and moves `core.CurseOfElementsAura` onto them with the same effects; no spec outside the warlock group moves (checked: mage-fire and priest-shadow, raid and bare, identical before and after). The three warlock rotation ladder goldens in sim/request move at the bare rungs (Curse of the Elements); not regenerated here.
- Ranks lane: the warlock bare rank choices (Destruction's Shadow Bolt rank 4 at 30 to 39) were measured without Curse of the Elements; re-check at the pin.
- Ranker: the bare warlock picks were ranked without the curse; the held refresh re-ranks them.

## Still open

- From round 10, unchanged: the warlock's plain crit reaching demons (single source), the demon sheets at 60, demon crit at 60.
- Percent spell modifiers adding or multiplying (round 6 A4): Improved Corruption's +10% and Aftermath's +40% to +50% are separate multipliers in the engine, while Shadow Mastery, Agonizing Flames and Improved Bane of Agony add. If the server adds them, Corruption reads about 0.8% high and Immolate's first hit about 3% high (0.1 to 0.2% of DPS). No new source either way.
- From earlier rounds: Ignition's one-a-minute rate, shadow wands (#419), Siphon Life tick crit (held for the owner), Demonic Brand (no build takes it).
