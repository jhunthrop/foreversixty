# Sim audit round 11: casters, 2026-11-02

Specs: mage-arcane, mage-fire, mage-frost, druid-balance, priest-shadow, shaman-elemental.

## Method

- **Engine.** `lane/r11-casters` from the pin cfcb76e80. **Site.** `wt/r11-casters` from origin/main 8f047a62f, with `make simdb`. "Before" is the shared baseline (`r11/sb-before` run from `r11/site-before/sim`).
- **Client.** 1.60.1.70707 with the hotfix push layered in (the round-10 layered copy, `r10/crit/raw`, fetched 2026-11-01 21:42 UTC, checked against its `shipped/` copy): SpellEffect, SpellMisc (Attributes_8), SpellCategories, SpellCooldowns, SpellAuraOptions, SpellPower, SpellCastTimes, SpellDuration, SkillLineAbility, ItemEffect; talent texts from `data/builds/1.60.1.70707/talents/`. Every talent mask was decoded against every learned spell of the class.
- **Fresh re-audit.** Spell by spell for every line each APL casts and every talent each band build takes, at raid 60 and bare 20 to 60, both factions: amounts and per-level growth per rank, coefficients, casts, costs, cooldowns, durations, masks, proc rates, crit tables, ranks against trainables; published BiS procs, uses and sets; racials of the published races. The six specs were split three ways (mage; Balance and Shadow; Elemental and the cross-cutting caster kit); every confirmed row was re-read against the client by the lane before fixing.
- **Tests.** Each confirmed bug got a failing test first, seen to fail, then pass. Fork: `go test -tags=with_db ./sim/core/... ./sim/mage/... ./sim/shaman/... ./sim/priest/... ./sim/druid/...` passes; no fork golden moved (TestP1Mage and TestElemental unchanged). Site: `go test ./request -run 'TestShadow|TestMage|Smoke'`, data `tests/test_apl.py test_addonrotation.py test_specs.py` (89 passed).
- **Runs.** spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, raid 60 and bare 20 to 60. Grids: `r11/casters/before.tsv`, `after.tsv`.

## Mage (Arcane, Fire, Frost)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **Clearcasting was -100% cost; Forever's mage Clearcasting is -1000%.** Percent cost mods add, so under the engine's -100 a Clearcasting Arcane Blast at four stacks (+400%) still cost four times its price, and a cast under Arcane Power still cost 30%. At -1000 the next damage spell is free in every case, which is what "your next damage spell costs no mana" means. The druid's and shaman's Clearcastings really are -100; only the mage's changed. Fire and Frost do not move (nothing raises their cost above the -100). | 12536 effect 0, SpellEffect 689319: aura 108, misc 14 (cost %), base points **-1000**, damage-spell mask; the same in the shipped table and the push. 16870 (695684) and 16246 (692327) are -100. `sim/mage/talents.go:190`. Test `TestClearcastingMakesAStackedArcaneBlastFree` failed (727.8 mana), then passed. Rounds 4 and 10 (A11) read the row as -100. | CONFIRMED-BUG | old (E c6d0c1f46, round 4) | E 6b15a7c56 |
| M2 | **A missed Scorch still added an Improved Scorch stack.** Improved Scorch is a trigger the spell adds to the target it hits; a miss applies nothing. Raid Scorch misses about 8%, so the debuff stacked a little faster than it can. | 11095: aura 109 (add target trigger) of 22959, 100%. cmangos applies target triggers only on a hit (SPELL_MISS_NONE). `sim/mage/scorch.go:70` rolled after every cast. Test `TestImprovedScorchStacksOnlyOnAHit` failed (cast 5 missed and stacked), then passed; the refresh test now casts until a Scorch lands. | CONFIRMED-BUG | old (upstream 029a226e2, 2024) | E 8736735a2 |
| M3 | **Rank-1 Arcane Missiles dealt only spell power.** Its missile is 7268, which the client names "Arcane Missile" (singular), so the generator filed it apart and rank 1 read the channel 5143's 0 damage and 0.2 coefficient. No published band casts rank 1 (levels 8 to 15, the /sim custom levels). | 7268: 24 a missile at level 8, +0.2 a level to 12, coefficient 0.286 (`constants_auto_gen.go` ArcaneMissile*[1]). Test `TestArcaneMissilesRankOneFiresTheClientMissile` failed ([0 0], 0.2), then passed. | CONFIRMED-BUG (no published number) | old (generator) | E c9073bba1 |
| M4 | Ice Lance ranks 2 to 6 share Fire Blast's cooldown in the engine. The push gave them CategoryRecoveryTime 8000 but their SpellCategories Category is 0 (only rank 1, 1312002, is in Fire Blast's category 19). Whether a category cooldown with no category is shared, private or nothing is server behaviour. No build casts both spells. | SpellCooldowns 83266, 94686-94689 (shipped 0, push 8000); SpellCategories 148169, 173563-173566 Category 0 | OPEN (no published number) | new-data | - |
| M11 | **Robe of the Archmage could fire within a minute of a mana gem or Demonic Rune** (routed from the items lane). The robe's use, the gems' and the rune's are one client spell category, each with a 60 s category recovery, so each holds the others. The engine held the gems and rune after the robe but not the robe after a gem. Category 30 is now one shared timer the gems and the rune start and the robe reads; the healthstones sit in category 30 with no recovery row and do not hold it. Only the three mage specs wear the robe in a published set (raid 60 and bare 60). | SpellCategories Category 30 for 18385, 5405, 10052, 10057, 10058, 16666; SpellCooldowns CategoryRecoveryTime 60000 on each (shipped and push); healthstones 5720/5723/11732 category 30, no SpellCooldowns row. `common/item_effects.go` robe, `mage/mana_gem.go`, `core/consumes.go` registerConjuredCD. Test `TestAManaGemHoldsRobeOfTheArchmageForAMinute` failed, then passed; TestP1Mage golden regenerated in the same commit. Robe-only effect (10,000 iterations, seeds 1 / 50001, horde / alliance): Arcane bare 60 +0.2 / +0.4 / +0.3 / +0.1%, Fire raid 60 +0.1 / +0.2 / +0.4 / +0.4%, Fire bare 60 -0.1 / -0.1 / -0.0 / -0.0%, Arcane raid and Frost within 0.25%. | CONFIRMED-BUG | old (E 7ae2a6a9b, round 8) | E eca6175f9 |
| M5 | Spell rows, rank by rank against SpellEffect, SpellPower, SpellCastTimes and SpellCooldowns: Scorch 1 and 7 (178, +1.7, 0.429), Fire Blast 7 (438, +3, 0.429; 8 s less 1 s a rank of Improved Fire Blast), Frostbolt 7/9/10 (386, +2.9, 0.814), Pyroblast 8 (583, +4.6, 1.0; dot 53 a tick, 0.15, 4 x 3 s; 6 s cast), Arcane Blast 5 (236, +1.6, 0.429; 15% of base mana; 400573 +100% cost and +12% Intellect as spell power a stack, 4 stacks, 18 s), missile 10274 (171, +0.9, 0.286), Ice Lance 6 (145, +0.9), Fireball 11. Every rank cast at each band is trainable there; the rank-11/12/8 AQ books stay gated. | client rows; `mage/*.go` | Aligned | - | - |
| M6 | Talents and buffs: Heating Up (400625, -25% a stack, three stacks, 20 s), Combustion (28682 mask, three non-periodic crits, Ignite outside it, cooldown from the buff's end), Presence of Mind and Arcane Power (cooldown on use / on cast; +30% damage and cost), Missile Barrage (400589 -50% duration, -100% cost; 20% a rank from Arcane Blast, 10% from others), Fire Power (1% a rank a Fire effect, four at most), Master of Elements (30% of base cost on direct crits), Ignite (8% a rank, 2 x 2 s rolling, item procs feed it, its ticks crit by the owner's rule, no double dip), Ice Shards and Arcane Mind (+100% of the crit bonus), Shatter (17/33/50%), Fingers of Frost and Ice Lance after round 10's A7, Winter's Chill (12579: +20% and +7% a stack on Cone of Cold only), Ice Barrier rank 4 (13033), Critical Mass, Incineration, Arcane Impact, Frost Channeling, Improved Frostbolt, Accelerant masks, Precision (aura 55, every school, Shoot included), Arcane Concentration (2% a rank, 1 s proc cooldown, missiles can proc it). | client rows, masks decoded | Aligned | - | - |
| M7 | Arcane Missiles crit: the missiles are direct hits (7268...25346) that roll their own crit; the channel's PERIODIC_CAN_CRIT does not apply. | `arcane_missiles.go`; Attributes_8 | Aligned | - | - |
| M8 | Mana: Evocation (12051, round 10 R1), Mage Armor 50% and Arcane Meditation 17/33/50% while casting, mana gems 400/600/850/1100 sharing item category 1153 (120 s) with Demonic Rune, Robe of the Archmage (500 ± 50%, 5 min), Major Mana Potion, Major Spellblasting Potion (47, 30 s). | client rows | Aligned | - | - |
| M9 | Gear and racials: Frozen Heart of the Mountain (+29 Frost and Shadow, 15 s, 5 min; idle for Fire, which wears it for hit), Magister's and Devout set procs, Kindling Stave's Ignition (rate still open from round 9), Troll Berserking at the client's 10% copy, Orc Blood Fury and Gnome as round 10. | ItemEffect, set rows | Aligned | - | - |
| M10 | Fire Power on dot ticks: 11124 is a server-side dummy; the engine applies it to direct hits only, so Pyroblast's dot and Ignite take none. | 11124 aura 4 | OPEN (log or tooltip test) | - | - |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Improved Moonfire's +10% is on Moonfire's base damage only (`BaseDamageMultiplierAdditive`), not its spell-power share. The client row is aura 108 misc 0 and misc 22 (damage and periodic %), which cmangos applies after the spell-power bonus (SpellDamageBonusDone). Every other aura-108 damage talent in the engine reaches full damage. But the upstream sim chose base-only on purpose (4faba5634 "Improved Moonfire/Moonfury tweak"), and round 4's caster model left it open. A full-damage fix would be worth about +0.9% raid Balance. | 16821 (mask 2); `druid/talents.go:510`; design/reviews/2026-10-07-caster-model.md:95 | OPEN (one Moonfire hit and tick with known spell power settles it) | old | - |
| B2 | Starfire 1-7 (coefficient 1.0, 3.5 s less 0.1 s a rank of Improved Starfire), Wrath 1-8 (0.429/0.486/0.571; 1.5/1.7/2.0 s; Improved Wrath 0.1 s and 10% cost a rank, additive with Moonglow's 25%), Moonfire 1-10 (direct 0.15, ticks 0.13, Nature's Splendor +1 tick), Insect Swarm (6 x 2 s, 0.158, -2% hit on the target): amounts, growth and costs equal the client. | SpellEffect, SpellPower, SpellCastTimes | Aligned | - | - |
| B3 | Moonfury (aura 79, Arcane and Nature), Vengeance (2.0x at 5/5, mask includes Insect Swarm), Nature's Majesty (aura 290), Nature's Reach, Genesis, Naturalist, Nature's Grace (16886: 10% cast speed and 150 ms off the global, 3 s; the engine does not also haste the global), Eclipse (0.17/0.33/0.5 s, two charges a Wrath, four at most, 15 s), Moonkin Form (35% base mana; Omen chance x2 and its cooldown x0.5; Moonkin Aura not doubled under the raid's), Omen of Clarity (4%, 10 s; Wrath outside 16870's mask), Innervate (29166: +400% spirit regen, full regen while casting, 20 s, 6 min). | client rows | Aligned | - | - |
| B4 | Gear and racials: Chthonic Power (1249110, Frost and Shadow only: Balance gains only its hit), Talons of Wrath (50%, 35 mana), Idol of the Moon (+33 Moonfire), Hand of Edward the Odd (melee-only proc, never fires), Rune of the Guard Captain, Ironfeather 2, Magister's 3, Elune's Light (1259799, aura 290, 10%, 15 s, 3 min), Tauren Endurance. | ItemEffect, set rows | Aligned | - | - |
| B5 | Improved Moonfire and Genesis on Moonfire's ticks multiply (1.10 x 1.04) where same-op percent mods add (1.14): about 0.03%, and only meaningful once B1 is decided. | engine | OPEN (rides on B1) | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Shadow never cast Dark Sacrifice.** The published horde Shadow priest is Undead, whose racial trades health for mana over 15 s plus Spirit (10 min cooldown). Holy and Discipline cast it since 2026-10-08; Shadow's rotation never had the line, while its bare bands run out of mana 40 to 93 s of 180. It is now cast below a fifth of mana when mana is behind the fight's remaining time. A 50% threshold without the pace check gained more at 30 but cost raid 60 0.65%; the paced line leaves raid 60 within error. Gnomes skip it. | SkillLineAbility 57097-57101 (1277324-1277328, levels 20-60, RaceMask 16 Undead); `priest/mana_abilities.go:74`; Holy/Discipline APLs. Test `TestShadowCastsDarkSacrificeWhenManaIsBehind` failed, then passed. Measured below. | CONFIRMED-BUG (rotation) | old | S db8fc067e, E 54caa19c2 (fork copy) |
| S2 | Mind Flay: 3 x 1 s, 0.167 a tick, 205 mana; Attributes_8 0x1200 carries PERIODIC_CAN_CRIT, so its ticks crit at the spell crit (16.8% observed against 17.3% stat x hit). spec-breakdown's 12.6% counts the zero-damage landing hit of every cast as a hit. Improved Mind Flay compounds on the snapshot. | SpellMisc; `priest/mind_flay.go` | Aligned | - | - |
| S3 | Mind Blast 1-9 (0.429, 1.5 s, 8 s; Improved Mind Blast 0.5 s a rank), Shadow Word: Pain 8 (10894: 127 a tick, 0.2), Devouring Plague (106 x 8, 0.1, 24 s, every Forever priest race learns it), Shadow Word: Death (1309636: 448 +2.5 a level, 0.429, 18 s; Early Demise 15/30%), Shadowform (+10% aura 79; half cost as a school multiplier after Mental Agility; +100% crit damage on mask 0x280a010/0x2, every Shadow spell), Twin Disciplines, Mental Agility, Shadow Weaving (15257: 100%, 2% x 5, 15 s), Darkness, Shadow Focus, Meditation, Wand Specialization; Vampiric Embrace is a party heal, so not casting it is right; Spirit Tap never triggers on one target. | client rows | Aligned | - | - |
| S4 | Inner Focus's crit mask (14751 effect 1) names Mind Flay, Mind Blast and Vampiric Embrace, not Shadow Word: Death; the engine gives it to every non-periodic priest spell. The rotation always spends Inner Focus on Mind Blast, so nothing published moves. | 14751 masks 0xd9f43e90/0x98434/0/0x10200000 | Noted (no published number) | old | - |
| S5 | Shadow Weaving's stack is added before the cast's own damage lands (15257 procs on hit); only the first five hits of a fight gain, far under 0.05%. | 15257 ProcTypeMask 0x10000 | Noted (under error) | old | - |
| S6 | Gear and racials: Devout 2-6 thresholds (resist, +18 spell power, the 5%/200-mana proc, 8 mp5), Frozen Heart (+29 Shadow), Touch of the Grave (round 10), Gnome Eureka! (1259823 from its mask, 2 casts a fight). | set rows, ItemEffect | Aligned | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **The own Mana Spring Totem was a dead global under the raid preset.** Two Mana Spring auras do not stack, so since round 8 the shaman's own totem adds only its excess over the raid's Mana Spring, which at the same rank is zero. Elemental's first rotation line still dropped it: the fight's opening global and 100 mana for nothing. The totem is now not castable when it would add nothing; with no raid Mana Spring (every bare band, where it is worth 3.6 to 7.2%) it is cast as before. | `shaman/water_totems.go:181` (`max(own - raidManaSpringMP5, 0)`); raid preset `mana_spring_totem`. Test `TestOwnManaSpringIsNotCastUnderTheRaidsManaSpring` failed, then passed. | CONFIRMED-BUG | old (round 8 E b092aee70 left the rotation line) | E a74c5e368 |
| E2 | Lightning Bolt (915 to 15207: amounts, growth, 0.714, 2.5 s, costs; rank-choice rows), Lava Burst 1-4 (rank 4 220 +1.3, 0.714; 2.5 s; 10 s category 1224; +20% only with Flame Shock, no forced crit), Flame Shock (166 +2.1 direct at 0.214; 44 a tick at 0.1 over 12 s; Call of Flame +9 s; ticks crit), Fire Nova 408341-5 / 408423-8 (0.143, 10 s), Searing Totem Attack 10436 (47, 0.017, 55 s), Mana Spring (10 a 2 s pulse, 300 s). | client rows | Aligned | - | - |
| E3 | Call of Flame (damage effect on Fire Nova, Searing, Flame Shock, Lava Burst; periodic effect on Flame Shock's dot), Concussion (mask 0x100003), Convection (shocks, bolts, Lava Burst; not Fire Nova), Elemental Alacrity (170/330/500 ms), Elemental Fury (no heals), Elemental Focus / Clearcasting (10%, one charge, 15 s; mask 0x98100003/0x1000, no totems), Lightning Overload (3/7/10% at half damage and coefficient; the overload keeps Lightning Bolt's mask, so Totem of Thunder and Concussion reach it), Elemental Devastation (melee only), Tidal Focus (+5 hit, raid build), Mindfulness, Ancestral Knowledge, Thundering Strikes (aura 290), Improved Reincarnation (owner keep). | client rows | Aligned | - | - |
| E4 | Relic and gear: Totem of Thunder 461295 (+1% on mask 0x1), Hand of Edward the Odd (melee proc, never fires), Ironfeather 2. | client rows | Aligned | - | - |
| E5 | Searing Totem's attack interval. The engine attacks every 2.5 s; round 10's cached beta timelines (11 encounters, ranks 1 and 2) show a 2.2 s Attack cast and about 0.23 s before the next, a median cadence of 2.43 s (73 landed-to-landed gaps). It would add about 3% more attacks, about +0.2% raid. One source (logs) plus the 2,200 ms cast row. | `shaman/fire_totems.go:62`; `r11/casters/ele/py/searing2.py` | OPEN (now measurable without level 60; a second source settles it) | old | - |
| E6 | Elemental Mastery (16166): ProcCharges 1 with a proc mask, the shape of Presence of Mind and Inner Focus, against the engine's 15 s buff. One charge would cost about 1.8%. | SpellAuraOptions | OPEN (round 6 E1, unchanged) | old | - |
| E7 | Bare 50 and 60 build the guide's leveling line with 5 points in Improved Healing Wave instead of Tidal Focus (+5% spell hit). Swapping measures about +5.7% at both bands (3,000 iterations, both seeds). Improved Healing Wave is the solo-leveling self-heal by design (2026-10-07-raid-builds.md, 2026-10-08-tidal-focus.md); the raid build already takes Tidal Focus. | band builds in `data/builds/1.60.1.70707/bis/shaman-elemental.json` | Held (owner: leveling build) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The "Thrash (21919)" aura at 99.9% on every caster is the engine's stored-extra-attack bookkeeping the raid's Windfury Totem registers (`core/buffs.go:1614`, `core/attack.go:795`). Casters never swing; it does nothing. | engine | Aligned | - | - |
| X2 | Caster consumables: Flask of Supreme Power 150, Greater Arcane Elixir 35, Elixir of Firepower +10, Cerebral Cortex Compound 25 Intellect, Brilliant Wizard Oil, Major Mana Potion 1350-2250 (Forever's 30 s potions), Demonic Rune 900-1500 sharing category 1153 with mana gems. | client rows | Aligned | - | - |
| X3 | Raid buffs on casters: Arcane Brilliance 31, Divine Spirit 40, Blessing of Wisdom 36, Curse of the Elements by band (-30/-45/-60/-75 resist, 4/6/8/10%), Moonkin Aura 3% every table, Judgement of Wisdom 59 at 50% on direct hits (Fire's procs equal about half its direct hits). | client rows | Aligned | - | - |
| X4 | Regen: shaman 17 + Spirit/5, mage and priest 12.5 + Spirit/4, druid 15 + Spirit/5 a tick (no client table; emulators, round 8). | emulators | Aligned (single source class) | - | - |

## Held under the adoption rule (re-run at the current pin)

| # | item | measurement | verdict |
|---|---|---|---|
| H1 | Frost's raid search winner (`251000203-5002-253510130101301221`: Arcane Impact, Improved Channeling and Arcane Resilience for Arcane Concentration and Accelerant). talent-search at cfcb76e80 with `-keep`, 10,000 final iterations, seed 7: +2.8 ± 0.7 (+0.5%). spec-breakdown 10,000: horde 623.0 / 623.9 to 625.8 / 625.6 (+0.4 / +0.3%), alliance 617.0 / 617.4 to 620.0 / 619.2 (+0.5 / +0.3%), seeds 1 / 50001. | under 1% | Held |
| H2 | Fire's Improved Scorch upkeep line dropped: horde 660.1 / 660.9 to 662.4 / 662.7 (+0.35 / +0.27%), alliance 651.9 / 652.7 to 655.7 / 655.6 (+0.58 / +0.44%). Fire's raid search at the pin: the guide is the best build. Measured before M2, which only touches the opening stack-up. | under 1% | Held |
| H3 | Frost bare 20: Ice Shards 3/5 for Frost Warding 2/2 measures +1.1% (horde 44.4 to 44.9, alliance 44.7 to 45.2, both seeds). Frost Warding is an unmodeled guide talent (armour and Ice Barrier absorb) and on raid-keeps, so the "keeps every unmodeled guide talent" bar fails. | +1.1%, drops an unmodeled keep | Held |

## Before/after

Before: pin cfcb76e80 on origin/main 8f047a62f. After: lane engine c9073bba1 on site db8fc067e (M11, eca6175f9, measured on its own in the M11 row). 10,000 iterations.

| Spec, band | before (horde, seed 1) | after | horde (seed 1 / 50001) | alliance (seed 1 / 50001) |
|---|---|---|---|---|
| mage-arcane raid60 | 600.4 | 610.0 | +1.6 / +1.6% | +1.7 / +1.7% |
| mage-arcane bare30 | 83.1 | 85.7 | +3.1 / +3.1% | +2.8 / +2.8% |
| mage-arcane bare40 | 159.6 | 163.0 | +2.1 / +2.1% | +2.0 / +1.9% |
| mage-arcane bare50 | 206.0 | 211.2 | +2.5 / +2.6% | +2.3 / +2.3% |
| mage-arcane bare60 | 329.9 | 338.1 | +2.5 / +2.4% | +1.9 / +2.0% |
| mage-fire raid60 | 660.1 | 659.6 | -0.1 / -0.2% | -0.1 / -0.2% |
| mage-fire bare30 | 70.4 | 70.0 | -0.6 / -0.7% | -1.1 / -0.8% |
| mage-fire bare40 | 153.4 | 153.1 | -0.2 / -0.4% | -0.4 / -0.5% |
| mage-fire bare50 | 189.9 | 189.2 | -0.4 / -0.4% | -0.4 / -0.7% |
| mage-fire bare60 | 329.8 | 329.2 | -0.2 / -0.2% | -0.4 / -0.2% |
| priest-shadow raid60 | 717.8 | 717.3 | -0.1 / -0.1% | +0.0 / +0.0% |
| priest-shadow bare20 | 34.9 | 36.4 | +4.3 / +4.3% | +0.0 / +0.0% |
| priest-shadow bare30 | 58.3 | 62.1 | +6.5 / +6.5% | +0.0 / +0.0% |
| priest-shadow bare40 | 128.7 | 148.9 | +15.7 / +15.6% | +0.0 / +0.0% |
| priest-shadow bare50 | 198.8 | 229.9 | +15.6 / +15.7% | +0.0 / +0.0% |
| priest-shadow bare60 | 354.1 | 365.1 | +3.1 / +3.1% | +0.0 / +0.0% |
| shaman-elemental raid60 | 466.9 | 470.0 | +0.7 / +0.6% | +0.7 / +0.6% |

Unchanged (under 0.05% at every seed and faction): mage-arcane bare20, mage-fire bare20, mage-frost raid60, mage-frost bare20, mage-frost bare30, mage-frost bare40, mage-frost bare50, mage-frost bare60, druid-balance raid60, druid-balance bare20, druid-balance bare30, druid-balance bare40, druid-balance bare50, druid-balance bare60, shaman-elemental bare20, shaman-elemental bare30, shaman-elemental bare40, shaman-elemental bare50, shaman-elemental bare60.

Causes:
- Arcane, every band from 30 and raid: M1 (Clearcasting frees stacked Arcane Blasts and Arcane Power casts).
- Fire, raid and bare 30 to 60: M2 (a missed Scorch no longer stacks Improved Scorch); bare 30 moves most because bare Scorch misses most.
- Shadow, horde bare 20 to 60: S1 (Undead Dark Sacrifice); alliance is a gnome and does not change. Raid 60 horde -0.1% at both seeds, within error (Dark Sacrifice is rarely cast there).
- Elemental raid 60: E1 (no wasted Mana Spring global). The bare bands keep their own totem and do not move.


## For other lanes

- **Integrator.**
  - No fork golden moved. sim/request TestRotationLadder goldens move for priest-shadow (Dark Sacrifice at the bare rungs), mage-arcane (M1) and mage-fire (M2); I did not regenerate them and deselected nothing.
  - No commit moves a spec outside this group. E1 is Elemental-only (no other published shaman APL drops Mana Spring); M1 to M3 are in sim/mage; M11 touches sim/core and sim/common, but only the mages wear the robe (Demonic Rune users without the robe are unchanged); S1 is the Shadow APL.
  - Arcane's stat weights and BiS at raid 60 and the bare bands should be re-ranked after M1 (mana matters less).
- **Ranks lane.** S1 makes bare Shadow less mana-bound, so Shadow's curated downranks in data/curated/rank-choices.json (Mind Blast 4 at 40, 2 at 50; Devouring Plague 1 at 30 to 60; Shadow Word: Death 2 at 50/60; Shadow Word: Pain 3 at 30) should be re-measured on this branch. M1 does the same for Arcane's rows, if any.
- **Items lane.** Elemental Mage Staff (944, +100 Fire and +100 Frost, the Fire and Frost raid main hand) has no hotfix row to check against; confirm it in 70707 ItemSparse/ItemEffect.
- **Ranker.** Every caster at 60 wears a one-hander (Hand of Edward the Odd) with an empty off hand, and bare Elemental has no relic: round 8's known simdb artifact (round 8 ranker R10); re-check after the refresh.

## Still open

- Improved Moonfire's reach (B1): a Moonfire hit and tick at known spell power with and without the talent.
- Searing Totem's 2.43 s cadence (E5) and Elemental Mastery as one charge (E6).
- Ice Lance's push-only category cooldown (M4); Fire Power on dot ticks (M10).
- Carried: Master of Elements on a Clearcasting crit (now always free for the mage after M1, so the question is only the refund), Lightning Overload's 3/7/10% (talent text and spell row), Kindling Stave's rate, Ignite's 0.73, Ice Lance's coefficient, Missile Barrage on a miss, Shadowfiend (no learn level), Starfire rank 7 at trainers (book-gated), Wrath's cost stacking.

## Sources

- 1.60.1.70707 client tables with the hotfix push (round-10 layered copy, fetched 2026-11-01 21:42 UTC, and its shipped copy). Rows read for the fixes: SpellEffect 689319, 692327, 695684 (12536, 16246, 16870), 687731 (11095); SpellCategories 148169, 173563-173566, 187717; SpellCooldowns 83266, 94686-94689, 101119; SkillLineAbility 57097-57101; spellconst ArcaneMissile 7268.
- cmangos-classic `Unit::SpellDamageBonusDone`, `Spell::DoAllEffectOnTarget` (target triggers on hit).
- foreverlogs.gg beta timelines cached by round 10 (Searing Totem cadence).
