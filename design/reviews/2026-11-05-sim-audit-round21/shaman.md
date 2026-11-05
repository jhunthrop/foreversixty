# Sim-accuracy round 21: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70790. Site branch wt/r21-shaman (off int/r21, int/r21 da6152025 merged in, so guides rounds 19 and 20 are under it), fork branch lane/r21-shaman (off int/r21 d5c708834 + lanes, int/r21 merged in). Gear is the frozen BiS ($R/bis-frozen); raid 60 runs the guide raid build ($R/rt); bare bands run the published band talents unless a row says otherwise. Sources read: the round 20 shaman report and integration "Open, for round 21", guide-accuracy rounds 19 and 20 (shaman), Blizzard's Forever notes thread (https://us.forums.blizzard.com/en/wow/t/2360696, every shaman line of the 24 Sep, 1 Oct, 8 Oct, 16 Oct and 3 Nov posts, read 5 Nov ~04:00 UTC), the 70790 client rows (data/builds/1.60.1.70790/spellconst, talents, hotfix-push.json; 70009 to 70790 compared for Maelstrom Weapon), Wowhead's Forever tooltips (nether.wowhead.com/forever/tooltip/spell/<id>, 5 Nov).

**3 confirmed bugs, all fixed** (Elemental 1, Enhancement 2). Restoration 0.
- Elemental: on two or more targets below 26 the card dropped no Fire totem, so Fire Nova never fired. Origin `old` + `guides-handoff` (guide r19 sim finding).
- Enhancement: Maelstrom Weapon's chance was half the live one (Blizzard, 3 Nov: "Chance to trigger doubled"; the client row never changed). Origin `new-data`.
- Enhancement: with the doubled chance, the five-stack Lightning Bolt below 5/5 and the half-mana Earth Shock reserve that paid for it now lose 1.9% to 3.7% at 38 to 46. Origin `new-data` (follows N1).

## Seed

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Elemental card, two or more targets below 26: no Fire totem, so no Fire Nova.**<br>What was wrong: Searing Totem was gated on `numberTargets == 1` and Magma Totem is trained at 26. In Forever Fire Nova (408341, trained at 12) is a spell that "inflicts fire damage to enemies within 10 yd of your active Fire totem" (`fire_nova.go` ExtraCastCondition), so a leveling shaman's pack rotation cast no totem and no Fire Nova from 12 to 25. The Fire Nova note (guide r19) and elemental.md's Levels 1 to 19 ("Fire Nova at 12 bursts around that Searing Totem when two mobs reach you") already said what a shaman does: Searing Totem goes down in the pack and Fire Nova bursts from it.<br>Decision: Searing Totem on a single target OR while Magma Totem is not known (`not spellIsKnown(10587)`, which the rank rewrite resolves to rank 1 at 26 and leaves unknown below). From 26 nothing changes.<br>Fix: APL line 3 condition and note; Fire Nova note takes guide r20's wording; elemental.md Levels 20 to 60 gets "On packs keep dropping Searing Totem for Fire Nova until Magma Totem takes its place at 26" (the reviewer's SHOULD).<br>Test `TestElementalDropsSearingTotemForFireNovaOnPacksBelow26` (3 targets, 14 / 20 / 25, first 30 s): 0 Searing and 0 Fire Nova before; 1.0 to 1.4 Searing and 1.0 to 2.4 Fire Nova after. Single-target ladder and raid 60 unchanged (the new branch needs 2+ targets and level < 26). | sim/request/shaman_pack_totem_test.go; trainables (Searing 10, Fire Nova 12, Magma 26) | CONFIRMED | old (the gate predates rounds 1 to 20); guides-handoff (guide r19 "Sim findings") | S 5d59b485a, S 0a192ab69 (prose), S (report commit), E 25b658fe0 (fork copy) |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | S1 above. | | CONFIRMED | old, guides-handoff | S 5d59b485a |
| E2 | Blizzard's 3 Nov shaman lines against the 70790 client and the engine: Lightning Bolt 10 base 353 (was 196 in 70753), Chain Lightning 4 2.5 s cast, 495 mana (225% of the bolt's 220), 510 base (145% of 353), shocks back to Era (Earth Shock 7 531, Flame Shock 6 292 + 80 a tick), Lightning Overload 10/20/30%, Call of Thunder 5%, Lightning Mastery. All in the 70755/70790 rows and the engine (round 20 E3). Wowhead's Forever tooltip for 10605 still prints 2 s / 315 mana (the 70753 values): the client rows and the note agree, so Wowhead is stale there, not the engine. | spellconst 15208, 10605, 10414, 29228 for 70753 / 70755 / 70790; wh10605 tooltip | Aligned | - | - |
| E3 | Older notes lines: Fire Nova coefficient 0.143 and lower base (16 Oct; `fire_nova.go`), Elemental Focus not spent by the Fire Nova that procs it (8 and 16 Oct; `clearcastingProc.IsFrom`), Clearcasting not spent by Lightning Overload (16 Oct; overloads land inside ApplyEffects, no cast event), Trial by Fire 10% / 2 s / 10 s a point (`fire_nova.go:44-48`). | talents.go:123-170, fire_nova.go | Aligned | - | - |
| E4 | Round 20's change (Trial by Fire 2 and Elemental Warding 2 as the tolls) re-read: the engine applies Trial by Fire to Fire Nova damage and cooldown and to Magma Totem's duration; the APL casts both on packs. No regression. | fire_nova.go, fire_totems.go | Aligned | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Maelstrom Weapon rolled 50% a melee hit; Forever doubled it on 3 November.**<br>What was wrong: Blizzard's 3 Nov notes: "Maelstrom Weapon: Chance to trigger doubled." The client never shows the chance: 408498's dummy reads 50 in every build from 70009 to 70790, so the doubling is server side, and the engine kept 0.5. Before the push a two-hander gained a stack on about half its hits: beta report 2748 (before 23 October; reports 2811 to 2813 are 23 to 29 October) has 51 stacks over 94 hits on a 3.2 s weapon (round 14 N1), and a 15 October forum post reads five stacks "every 9 swings or so" (https://us.forums.blizzard.com/en/wow/t/maelstrom-weapon-useless/2382747). Doubled, that is every hit on a shaman's two-hander.<br>Fix: `maelstromWeaponProcChance` 1.0; Totem of the Storm's Lightning Bolt chance (half) 0.5. Test `TestMaelstromWeaponChanceIsDoubledByTheThirdNovemberPush`: failed (0.5 / 0.25), then passed.<br>Model note: a flat 100% and 20 a minute only differ on a swing faster than 3.0 s (20 a minute on Manual Crowd Pummeler's 2.0 s is 67%); the published BiS uses that mace at bands 30 to 39. OPEN for a fast-weapon log. | relics_test.go; notes thread post 9; spellconst 408498 across builds | CONFIRMED | new-data (3 Nov server change, no client column; rounds 14 and 20 kept 50% open on the old rate) | E b7481b365 |
| N2 | **Below 5/5 the five-stack bolt now costs more than it earns, and the half reserve with it.**<br>What was wrong: the card cast Lightning Bolt at five stacks with any rank of the talent and held Earth Shock at half mana from the first point to the fourth to pay for those bolts. At 2/5 or 3/5 the bolt still has a 1.0 to 1.5 s cast that holds up the next swing; with stacks every few seconds the card spent its time and mana on them. Without the partial-rank bolt and with a quarter reserve everywhere (10,000 iterations, horde / alliance, seeds 1 / 50001): level 38 (2/5, ladder talents) +3.3 / +3.3 / +1.9 / +1.9%, 40 +3.3 / +3.4 / +2.9 / +3.2%, 41 +3.2 / +3.5 / +3.4 / +3.7%, 46 (band 40 talents, 3/5) +3.2 / +3.2 / +2.7 / +2.9%; bands 50 and 60 and raid 60 identical (5/5). The bolt-only change without the reserve change is about half of it (40: 155.2 against 152.6 and 157.4).<br>Fix: Lightning Bolt asks `auraIsKnown(408498 tag 5)`; Earth Shock `mana > 25%` only; notes, enhancement.md loop steps 3 and 6, the Mana paragraph (the half-reserve sentence removed) and Levels 20 to 60 ("Maelstrom Weapon starts at 37 and is full at 42; spend its stacks on Lightning Bolt from 42 ..."). Test `TestEnhancementShocksDownToAQuarterOfItsMana` + `TestEnhancementBoltsMaelstromOnlyAtFiveOfFive` (replace the round 19 half-reserve test): failed, then passed. | `$R/shaman/enh-lb55.tsv`, rotation variants in `$R/shaman/rot` | CONFIRMED | new-data (follows N1; on the old engine the partial bolt was a tie at 40 and -0.6% at 46) | S 0a192ab69, E e7eb0dfac |
| N3 | Notes lines: Stormstrike 150% normalized, 6 s, 8% of base mana (`stormstrike.go`, client 17364), its debuff not consumed by a spell that does not land (`result.Landed()`), Mental Dexterity 150%, Flurry 30%, Rage of the Farseer melee haste only, Improved Ghost Wolf moved to row 4 (carried by name). | stormstrike.go:40-125; round 20 N6 | Aligned | - | - |
| N4 | "Shamans never dual wield": every Enhancement BiS band has an empty off hand weapon slot (shield or nothing); Dual Wield Specialization (408496) is the Season of Discovery rune and has no Forever learn row. | bis/shaman-enhancement.json | Aligned | - | - |
| N5 | Rotation rechecked on the new engine at raid 60: Lightning Bolt at 4 stacks loses (-0.5%, 860.0 against 864.6, 3,000 iterations). | | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Notes lines: Riptide crits (24 Sep; 408521, 1239242, 1239243 are in `periodic_crit_auto_gen.go`, so the heal over time rolls crit), 12 s cooldown, about 40% cheaper, no direct heal (3 Nov; `riptide.go`), Mana Tide rank at 38 (16 Oct), Restorative Totems on the whole Healing Stream (16 Oct), Water Shield from any spell crit (16 Oct; `water_shield.go:84-96`), Tidal Focus / Totemic Focus rows swapped (talent carry by name). | engine files named; notes thread | Aligned | - | - |
| R2 | Round 20's rows (Chain Heal rank 1 and 2, tank Healing Wave rank 7) were judged level by level in round 20; nothing in 70790 or this round touches healing. | | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Fork conformance golden (reported stale at d5c708834).** `go test ./sim/conformance -run 'TestConformance$/shaman'` fails at d5c708834 and on this branch; regenerating it changes one cell: Healing Wave rank 5 at level 50, the low-level penalty coefficient 0.4714 to 0.4713 (0.857 x 0.55 = 0.47135, a rounding tie). It hides no shaman change. Not regenerated (not this lane's change). | sim/core/testdata/conformance/shaman.golden.md | note | - | - |
| X2 | `data/tests/test_apl.py::test_every_spell_the_rotations_name_exists_with_that_rank` fails on int/r21 on mage-arcane (18809 rank 0); not shaman. | | For the mage lane | - | - |

## Before/after (10,000 iterations, frozen BiS, seeds 1 / 50001)

Before = $R/before.tsv and $R/sb-before (engine d5c708834, site 93c66ebc8); after = engine e7eb0dfac on site 0a192ab69.

| Spec, level | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Enhancement raid 60 | horde | 859.1 to 864.5 | +0.63 / +0.76 |
| Enhancement raid 60 | alliance | 856.6 to 858.0 | +0.16 / +0.07 |
| Enhancement bare 40 | horde | 154.5 to 157.2 | +1.7 / +1.8 |
| Enhancement bare 50 | horde | 219.0 to 234.1 | +6.9 / +6.8 |
| Enhancement bare 60 | horde | 322.6 to 341.9 | +6.0 / +6.0 |
| Elemental raid 60 | both | 666.1 / 669.4 unchanged | 0.00 |
| Restoration raid 60 | both | 757.1 / 759.4 unchanged | 0.00 |

Elemental single target is unchanged at every level (S1 needs two or more targets below 26).

## Files to regenerate

- data/builds/1.60.1.70790/addon-data.json and addon/ForeverSixty/Data.lua: the Elemental Searing Totem and Fire Nova notes, the Enhancement Lightning Bolt and Earth Shock notes (the web card's Earth Shock row reads addon-data and still shows the half reserve until then).
- sim/request/testdata/ladder/shaman-enhancement.golden.md: DPS at 38, 40, 50, 60 moves (Maelstrom chance; at 38 and 40 the rotation too).
- Coach starters, if they carry the Enhancement rotation.
- Enhancement BiS and stat weights move with the refresh (Maelstrom now stacks on every hit; Lightning Bolt's share grows at 50 and 60).
- Not this lane's: the fork's shaman conformance golden (X1).

## Measure at refresh

- Enhancement bands 40 to 60 and the relic (Totem of the Storm's bolt-triggered stacks are worth more at the doubled chance).
- Carried: Elemental band 20/30 talent strings, Restoration bands 40 / 50 / bare 60 Chain Heal rows, Enhancement band 30 with Stormstrike.

## For other lanes

- Nothing moves a spec outside shaman. Engine commits touch sim/shaman only; site commits touch shaman APLs, guides and sim/request shaman tests.

## For the guide lane

Shipped with this lane, wow-player on the rendered pages ($R/shaman/cap):
- Elemental Searing Totem / Fire Nova notes and the Levels 20 to 60 clause: SHIP (ele-card, ele-rotation, ele-leveling). SHOULDs: the addon row prints the Searing note's "until Magma Totem at 26" at 60 too (harmless); a mana line for Fire Nova at 10 to 25 (the Fire Nova note already says "two or more targets").
- Enhancement loop steps 3 and 6, Packs and Mana paragraphs, Levels 20 to 60, Lightning Bolt and Earth Shock notes: HOLD on the first capture (the Packs paragraph still said to spend five stacks with no five-point condition; two captures cut the new sentences), fixed (Packs: "once you have all five points of Maelstrom Weapon") and re-captured (enh-packs, enh-leveling3, ele-leveling3): SHIP, on condition that the integrator re-captures the Priority table after the addon-data regen. Rows 4 and 7 must then read "At five Maelstrom Weapon stacks once you have all five points" and "On cooldown while your mana is above a quarter". The old half-reserve clause must not survive.

## Still open

- Maelstrom Weapon flat 100% against 20 a minute: settle with a log of a 34+ Enhancement shaman on a two-hander faster than 3.0 s (Manual Crowd Pummeler).
- Pending owner, not reopened: Lightning Overload rank spells; Healing Wave 2 tooltip; level-60 book ranks.
- Logs: last looked 04:06 UTC 5 Nov. api.foreversixty.gg lists only the 20 October report; foreverlogs.gg public list empty and every probed id "Report not found"; Warcraft Logs v1 zones on classic, fresh, vanilla and www list no Forever zone. $R/platform/logs.md says the same.
