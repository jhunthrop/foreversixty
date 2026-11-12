# Hunter audit (BM / MM / Survival) - engine 5f1b39450 vs client 1.60.1.70888 (+hotfix push)

Client tables: scratchpad/melee/wago/*_70888h.csv. Helpers: scratchpad/melee/hunt/{sp.py,tab.py,masks.py,dec.py}.
Builds decoded against data/builds/1.60.1.70790/talents/hunter.json (see dec.py output in the session).

## Confirmed bugs

| # | Item | Engine | Client | Verdict |
|---|---|---|---|---|
| B1 | Resourcefulness cost cut on Lacerate | talents.go:691-699 (`applyResourcefulness`): -30%/rank on every spell with ProcMask MeleeMHSpecial/OHSpecial; lacerate.go:36 gives Lacerate ProcMaskMeleeMHSpecial, so Lacerate pays 40% (95 -> 38 at 60) | SpellEffect 440529 effect 0 (aura 108, misc 14 = cost, bp -60): class mask w0 194 (Raptor/Mongoose b1, Wing Clip b6, traps b7), w1 524288 (b19 Counterattack only), w2 2097152 (b21 Strider Kick). Lacerate 24118/24119/24120/1299332 SpellClassOptions mask w1 4096 (b12) is NOT in it. Same rows in 70790. Contrast: Savage Strikes 19159 and Predator's Edge 1310627 use w1 528384 = b19 + b12, i.e. the designers name Lacerate there and leave it out here. | BUG. Survival (build takes Resourcefulness 2/2; leveling order takes it from 31-32; card casts Lacerate from 30). Lacerate true cost 40/65/80/95, engine 16/26/32/38 (rank 1-4). At raid 60: ~9.5 casts (round-13 breakdown, -362 mana) -> about +540 mana a fight; DPS effect only where Survival runs dry (the 26-49 bands are mana-bound per rank-choices.json). Prior rounds 12/14/15 marked it aligned from the tooltip ("melee abilities"), not the mask. |

## Open questions

| # | Item | Engine | Client / source | Note |
|---|---|---|---|---|
| Q1 | Immolation Trap hit | immolation_trap.go:104-108 strips the hunter's whole stats.Hit (upstream vanilla hack "traps gain no benefit from hit") so the trap rolls a bare spell miss (~17% vs 63) | 13797/14301 DefenseType 1 (magic). Forever merges spell/melee/ranged hit into one stat (core/stats/stats.go:34-37) and Surefooted 19290 carries aura 55 (spell hit) +3 as well as aura 54; Survival takes Surefooted 3/3. | Likely BUG under Forever's unified hit; no prior ruling found. Survival, trap ~4-6% of damage, miss 17% -> ~8-11% => ~0.3-0.6% DPS. Needs one owner/log check. |
| Q2 | Mongoose Bite window charge | mongoose_bite.go:64: only Mongoose Bite spends the window | 5302 and 1310726: ProcCharges 1, ProcTypeMask 16 (done melee ability). Read literally, Raptor Strike / Strider Kick / Lacerate / Wing Clip landing inside the window would spend it first. 2026-10-07-mongoose-bite.md recorded the mask but modeled Mongoose-only. | Survival only; if the server spends the charge on any melee special, Mongoose Bite casts drop a lot (~19.5/fight now). Needs a log (5302 removed by a Raptor Strike). |
| Q3 | Pets and Expertise | pet.go:246-256 inheritance has no Expertise channel | Patch notes 2360696 line 1299: "Pets now correctly benefit from Expertise." | BM only, small; unclear whether it means inheritance. |

## Latent (no current build/card affected)

- Efficiency mask (19416) includes Lacerate (w1 b12) and Summon Hawk (w2 b13); talents.go:507 `efficiencyMeleeSpellCodes` omits Lacerate. No build that takes Efficiency casts Lacerate.
- Intimidation (BM takes it) unmodeled: pet's next attack +100% crit (19577 eff1 aura 290 bp 100), 60 s, 8% base mana. Card does not cast it; tiny.
- Frenzy proc aura ActionID 19625 does not exist; client buff is 19615 (aura 319 30, 8 s). Numbers match; id cosmetic.

## Aligned (checked line by line)

Shots / ranged
- Aimed Shot 19434/20900-20904: +40/70/115/185/260/350 (client_damage.go:50-58), 75-310 mana, levels 20-60, 2.0 s cast, category 2 6 s shared with Multi-Shot, GCD 1.5. Generated rank 6 (27632, 600, 3 s) correctly bypassed.
- Multi-Shot 2643: one rank (14288-14290 rows have no spell), 13.9% base mana, instant, cat 2 6 s, normalized weapon, 3 targets, missile 30.
- Arcane Shot 3044/14281-14287: 25/35/50/85/120/175/220/280, 11% base mana all ranks (70888 hotfix, arcane_shot.go:16), cat 1173 6 s, +15% RAP (Blizzard 3 Nov), missile 40.
- Serpent Sting 1978/13549-13555: 3/7/14/26/40/58/75/100 a tick, 15-230 mana, 5 x 3 s, nature, can crit (periodic_crit list), +5% RAP a tick, Mortal Shots in mask (b14). Rank 9 25295 held back by IncludeAQ=false (settled 2026-10-07-aq-ranks-off).
- Sniper Shot 1310687/1310785/1310786: 320/450/590, 365 mana, 4 s cast, 15 s RecoveryTime, normalized; Mortal Shots and Efficiency masks include it (w3 b23), Barrage does not.
- Auto Shot 75: 0.5 s aim, fires during Aimed/Sniper casts (notes 2360696), ammo DPS x weapon speed (upstream ruling, commit 443499c80).
- Hunter's Mark 1130/14323/14324/14325: 26/59/98/143 RAP, 15-60 mana, 2 min, ranged table only.
- Rapid Fire 3045: 40% ranged (aura 140) and melee (aura 319), 15 s, 100 mana, 300 s category, off GCD; Rapid Killing -60 s a rank (415405 bp -120000 at 2/2, mask b5).
- Aspect of the Hawk 13165-14322: 20/35/50/70/90/110 RAP, costs; strongest-rank pick. Quick Shots 6150 30% / 12 s; Auto-Shot-only proc settled by logs (round 16 H3, round 30 L10).
- Trueshot Aura 1299346/1299348/19506/20905/20906: 20/35/50/75/100 RAP (core/buff_ranks.go:128).
Melee (Survival)
- Raptor Strike 2973/14260-14266: +5/11/21/30/35/40/55/70 non-normalized (effect 58), 10-100 mana, cat 40 6 s, no GCD, next swing.
- Mongoose Bite 1495/14269-14271: +15/22/37/57 normalized (121), 30/40/50/65 mana, cat 65 5 s; Expose Prey window 10 s.
- Lacerate 24118/24119/24120/1299332: 17/28/40/58 a tick +0.5/0.6/0.7/0.8 per level to 38/48/58/68, 7 x 3 s, 40/65/80/95 mana, ticks crit (periodic_crit list), Savage Strikes and Predator's Edge masks include it.
- Lacerating Strikes 1310533/1310536: 40% of the bite over 7 x 3 s, crits.
- Strider Kick 1317257: 121 + 31 at 100% (normalized weapon), 8 s, 5.81% base mana, Resourcefulness mask includes it (w2 b21).
- Wing Clip 2974/14267/14268: 5/25/50 physical, 40/60/80 mana.
- Immolation Trap 13795/14302-14305 -> 13797/14298-14301: 21/43/68/102/138 a tick, 5 x 3 s, 50-245 mana, cat 411 30 s, +2% RAP a tick; Clever Traps +30% (19239 eff1 misc 8 mask b2); Survivalist's Discipline -40% CD (1310496, mask b7) -> 18 s.
- Aspect of the Beast 13161/1299445-1299447: 50/70/90/110 melee AP, costs; Quick Strikes 1299448 30%/12 s, melee auto proc (mask 4).
Talents (rank values vs talent text and client masks)
- BM: Deadly Aspects 2%/rank (masks Hawk w0 b20, Beast w2 b4), Focused Fire 1%/rank hunter+pet, Unleashed Fury 3%/rank, Ferocity 2%/rank, Frenzy 20%/rank 30% 8 s, Bestial Discipline 25%/rank regen + 10%/rank focus, Bestial Wrath 50% 18 s 120 s 12% off-GCD, Endurance Training, Bestial Swiftness (unmodeled, no DPS).
- MM: Lethal Attacks 1%/rank, Improved Stings 6/13/20% (SPELLMOD_DOT, mask b14), Efficiency 3%/rank (shots/stings/melee), Careful Aim 20%/rank Int to melee+ranged AP (auras 580/598), Rapid Killing CD half, Lone Wolf 20% no pet (site drops the pet when taken), Mortal Shots 6%/rank crit bonus (mask = Auto, Concussive, Arcane, Multi, Volley, Serpent, Aimed, Scatter, Sniper), Rapid Recuperation 25/50% 15 s on Serpent Sting, Barrage 5/10/15% (mask Multi, Volley, Aimed), Ranged Weapon Specialization 1%/rank, Trueshot Aura, Sniper Shot.
- SV: Improved Tracking 1%/rank, Savage Strikes 2%/rank (mask Raptor, Mongoose, Wing Clip, Lacerate, Counterattack, Strider), Survivalist, Clever Traps, Surefooted 3% hit, Predator's Edge 6%/rank crit bonus + 10%/rank OH, Resourcefulness proc (mask 272, 50%/rank, 1242688 50% regen 30 s), Expose Prey 5/10% 10 s (mask 340, needs Mark), Survivalist's Discipline, Strider Kick, Lightning Reflexes 3%/rank Agility, Lacerating Strikes.
Pet (cat)
- Claw 16827-16832/3010/3009: 4-6 ... 43-59 (bp x (1 +/- var/2)), 25 focus; Bite 17253-17261: 7-9 ... 81-99, 35 focus, cat 19 10 s; Happiness x1.25, cat x1.10, 10% hunter AP, full crit/hit inheritance (notes 16 Oct), Command +5%.

Count: 1 confirmed bug, 3 open questions, 3 latent notes, ~115 items aligned.
