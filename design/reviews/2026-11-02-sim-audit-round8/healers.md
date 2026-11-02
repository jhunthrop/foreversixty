# Sim audit round 8: healers

Lane: wt/r8-healers (site, from origin/main 170b7b84c) and lane/r8-healers (fork, from cd4dfdcc0). Specs: Restoration druid, Holy paladin, Discipline priest, Holy priest, Restoration shaman, the heal profile, and the effective-HPS metric.

Client truth: the 1.60.1.70707 tables layered with the hotfix push (fetched 2026-11-01 21:42 UTC). Server rules come from cmangos-classic and vmangos (fetched 2026-11-02).

The short version:
- **Every set bonus that says "damage and healing" gave no healing.** That includes Vestments of the Devout 3P, which all five published healers wear. A heal reads healing power only, and those bonuses were decoded as spell power alone. Eight hand-written sets had the same gap.
- **A Lightwell click healed for all three ranks at once.** The three ranks' renews shared one aura label, so the one aura on a member carried every rank's ticks.
- **No spell learned below level 20 took Classic's coefficient penalty.** Every caster and healer at band 20 was too strong.
- **Innervate was castable before level 40.** This inflated band-20 and band-30 druids: Restoration, and Balance by about 30%.
- **The shaman's spirit regen used the generic 15 + Spirit/5.** The shaman's own base is 17.
- **The priest racial heals were missing.** Divine Grace (human) and Contingency Plan (gnome) were not modelled; the published Discipline priest is human and the published Holy priest a gnome.
- **Smaller fixes.** Twin Disciplines no longer raises Prayer of Mending, and Desperate Prayer is now dwarf-only. A Holy paladin's bare kit now blesses itself with Wisdom, not Might.

I ran four delegated engine-vs-client sweeps (druid, paladin, priest, shaman). Before calling any reported bug confirmed, I re-read it against the client row and the code myself.

## Cross-cutting (metric, profile, core)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **"Damage and healing" set bonuses never reached a heal.** `Spell.HealingPower` reads the HealingPower stat alone. `flatDecoder.finish` turned a bonus with equal damage (aura 13) and healing (aura 135) into SpellPower only. Example: Vestments of the Devout 3P, 9346 "Increases damage and healing done by magical spells and effects by up to 18". Every published healer wears 3 or more Devout pieces. Measured on the shaman: dropping the third piece cost only the skirt's own 24 healing, not 42. | `sim/core/client_sets.go` `finish`; SpellEffect 9346 eff0 aura 13 = 18, eff1 aura 135 = 18; ItemSetSpell set 182 threshold 3. Test first: `TestDecodeFlatBonus` ("equal damage and healing is spell power and healing"). `TestSetBonusConformance` then caught the hand-written ones: Ironfeather Armor 2P, Blue Dragon Mail 3P, Illusionist's Attire 2P, Arcanist Regalia 3P, Nemesis Raiment 3P, Demoniac's Threads 2P, Judgement Armor 5P, Avenger's Battlegear 5P. Also Talisman of Ephemeral Power (23271) and the Restrained Essence of Sapphiron (28779), which both carry aura 135. | CONFIRMED-BUG (fresh) | E 7783a83ad |
| X2 | **Classic's sub-level-20 coefficient penalty was never applied.** A spell learned below 20 keeps 1 - (20 - level) x 0.0375 of its spell power coefficient. The client's coefficients carry no penalty: Healing Touch rank 1 (5185, level 1) is 0.429 = 1.5/3.5, and Lesser Heal 2050 is 0.429. So band-20 Heal rank 1 (level 16) is 0.857 x 0.85, Holy Light rank 3 (level 14) 0.714 x 0.775, Rejuvenation rank 3 0.2 x 0.85, and Healing Wave rank 4 0.857 x 0.925.<br>Applied only to ranks the player casts (a cost or a cast). A proc keeps its coefficient: every Seal of Righteousness and Seal of Fury proc carries client spell level 1, and how the penalty treats a triggered spell is unsettled. Physical abilities are untouched. | cmangos-classic `Unit::CalculateLevelPenalty` and vmangos `SpellCaster::CalculateLevelPenalty`: `1 - (20 - spellLevel) * 0.0375`, applied in both direct and periodic bonus code. This is a server rule, not a client table, from two emulators. Three sweeps reported it independently. Tests first: `TestSpellsLearnedBelowTwentyGainLessFromSpellPower`, `TestAProcKeepsItsCoefficientAtAnyLevel`. The conformance report now expects the penalised value, and the rank tests apply `core.LevelPenalty`. No fixture golden moves. | CONFIRMED-BUG (fresh; server rule) | E 87743552b |
| X3 | **Two spells could share one dot aura on a target.** A dot appends its gain and expire handlers to the aura it wraps, and `GetOrRegisterAura` hands a second spell with the same label the first one's aura. Both dots then tick whenever either is applied. This is the mechanism behind P1. Registration now refuses it and names the aura. Nothing else in the suite trips it. | `sim/core/dot.go` `claimDotAura`. Test first: `TestTwoSpellsCannotShareADotAura`. | CONFIRMED-BUG (engine guard) | E 7b668da02 |
| X4 | **Effective HPS.** Overheal is excluded: `GainHealth` returns what landed. An absorb counts only what it soaks (`AbsorbDamage`), and raw HPS adds shielding once. Heals over time count per tick. Self-heals on the healer count, because pulses really hurt it (the healer has a health bar, `raid.go:292`). The Demonic Rune's 800 health cost is real damage the healer can heal back. | `core/spell_result.go dealHealingInternal`, `core/shield.go`, `core/metrics_aggregator.go:455`. | Aligned | - |
| X5 | **Raid model and round-7 side effects.** Checked:<br>- Pulse picks: 3 of 5, the healer included.<br>- Pushback: 1 s, then 0.2 s less a hit, floor 0.2 s. The resist chance is read off the caster's spell. A cast never extends past its full cast time. Channels are cut.<br>- Mana is spent at cast end, so the five-second rule starts there (`cast.go:326`, `mana.go:363`).<br>- Mana Spring by level stacks with Wisdom.<br>- Water Shield procs on pulses; Healing Stream is counted.<br>- Clearcasting from HoTs, Bloom crit, Ankh, Oath of Light.<br>Uther's Light (8397, struck in melee, proc type 0x28) and Divine Protection (27778) correctly ignore the pulses or are left out. | Code read; shaman and druid sweeps. | Aligned | - |
| X6 | **Spirit regen per class** against cmangos and vmangos `OCTRegenMPPerSpirit`: druid and paladin 15 + Spi/5, priest 12.5 + Spi/4, shaman 17 + Spi/5 (fixed as S1). Mana potions (17531: 1800, variance 0.5) and the Demonic Rune (16666: 1200, variance 0.5) match. So do the Forever healer elixirs: Greater Mageblood 20 mp5, Greater Cleric's 40, Elixir of Sages 25 Spirit and 2% crit, Sage's Tea 44, Distilled Wisdom 2000, Brilliant Mana Oil 15 mp5 and 30 healing. | Client rows; `core/consumes.go`. | Aligned | - |
| X7 | **Heal profile coverage at 1.85 after this lane's fixes** (published raid gear, alliance, seed 1, 10,000 iterations; incoming 977.8 a second). Shaman / druid / Holy priest / paladin / Discipline: 86.0 / 79.7 / 70.2 / 67.8 / 63.3%. Bare sets cover under 40%. The Holy priest's raid set still runs dry. The strongest is at the rule's "about 85", so this is not a retune. | `data/curated/heal-profile.json` rule. | Aligned; repeat after the re-rank | - |
| X8 | **Published bis entries are stale against the engine.** Two cases:<br>- Holy paladin bare bands: held-in-off-hand items (Orb of Lorica 11262, Elunarian Sphere 15968) still carry no hand type in the committed `simdb.bin` (23:47 UTC 1 Nov), so they overwrite the main hand. The pipeline fix 1bd3920b5 landed after it. That is why the band-40 undead paladin publishes 12% over the human.<br>- The priests' published talents predate round 7's Twilight Focus builds.<br>Both need data.yml and bis.yml, which are held. | `aud-paladin` measurements: the human at band 40 with a buckler sims 107.2 against the undead's 105.5. | Known; held by the owner | - |

## Priest: Holy and Discipline

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **A Lightwell click healed with all three ranks' renews.** Renews 7001, 27873 and 27874 shared the label "Lightwell Renew", so a rank-3 well's click ticked 160 + 233 + 320 a tick instead of 320. The breakdown showed healing under all three ids. | `sim/priest/lightwell.go`; renews 7001/27873/27874 heal 160/233/320 every 2 s. Test first: `TestALightwellClickTicksOnlyTheCastRanksRenew` (4,320 and 6,058 healed through the lower ranks). Small, because the renew is mostly overheal. | CONFIRMED-BUG (fresh) | E 420a8dea7 |
| P2 | **Twin Disciplines raised Prayer of Mending.** The engine applied it to every instant heal. Its client modifiers name spells by mask (1225132 eff0 op 0 `0x19480020/0x18012/0/0x200000`, eff1 op 22 `0xa08060/0/0/0x10`). Neither shares a bit with Prayer of Mending's cast (1240827 `0/0x20/0x400/0`) or its jump (401880 `0/0x20/0/0`). | Test first: `TestTwinDisciplinesLeavesPrayerOfMendingOut`. Holy about -0.6%. | CONFIRMED-BUG (fresh) | E 0d7440978 |
| P3 | **Divine Grace (human) and Contingency Plan (gnome) were not modelled.** Both are trainable Forever racials (SkillLineAbility AcquireMethod 0).<br>- Divine Grace: ranks 1277370-1277378, rank 7 1383 + 8.2 a level at 0.429. It heals a friendly target below 50% health, is not castable on self, is free, and has a 10 minute cooldown.<br>- Contingency Plan: wards 1277462-1277640, 30 s. The next damage below 35% health gives a 926 absorb (1277466) and 134 every 3 s for 15 s at 0.2 a tick (1277460). 10 minute cooldown.<br>Each is now an autocast cooldown on the healer's target. The conformance report lists Divine Grace and finds it matches the client roll. | Client rows; RaceMasks 1 and 64. Tests first: four tests in `sim/priest/healing/racials_test.go`. About +1% each. | CONFIRMED-BUG (fresh) | E 119b49487 |
| P4 | **Desperate Prayer was registered for every race.** Its learn rows (13908, 19243) carry RaceMasks 4, dwarf only. No rotation casts it, so no number moves. | Test first: `TestDesperatePrayerIsTheDwarfPriests`. | CONFIRMED-BUG (no impact) | E 419d53a37 |
| P5 | **Every priest heal rank matches the client.** I re-compared all 76 rows of `heal_ranks.go` plus Binding Heal's 6 by script: learn level, mana, cast, coefficient, base, variance, growth and cap. Also aligned:<br>- Mechanics: Weakened Soul 15 s, Power Word: Shield (4 s cooldown, 0.1), Penance (3 pulses, 0.19 each, 12 s), Prayer of Mending (5 charges, 0.429, 10 s), Prayer of Healing (party, 0.286), Lightwell's costs and cooldown, Dark Sacrifice, Inner Focus, Divine Aegis (non-periodic crits, 20 x level cap).<br>- Talent curves and masks: Mental Agility, Renewed Hope, Twilight Focus, Improved Renew, Holy Specialization, Divine Fury, Improved Healing, Spiritual Healing, Inspiration, Litany of Light, Spiritual Guidance, Meditation.<br>- The gnome, human and troll racials. | `aud-priest/bin/cmp_ranks.py`, `trait.py`, `masks.py`. | Aligned | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **The bare self-buff kit gave a Holy paladin Blessing of Might.** That does nothing for a healer. A Holy paladin blesses itself with Wisdom: rank 1 is 19742, learned at 14, and ranks give 12 to 40 mana every 5 s. One paladin's blessings on a target are one at a time, so Wisdom replaces Might from 14. The kit feeds the ranker, the ladder, /sim and every tool. | `sim/leveling/kit.go`; raid-consumes.json and published-presets.json regenerated from it. Test first: `TestKitBuffsHolyPaladinCarriesBlessingOfWisdomFromFourteen`. Bare band 20/30/40/50: +10.6 / +20.0 / +17.8 / +16.1%. | CONFIRMED-BUG (fresh) | S 58aaee100 |
| H2 | **The list's note said "Holy Light (rank 9)"**, but it casts 10329, rank 8. Rank 9 (25292) is an Ahn'Qiraj book rank in `book-ranks.json`. | Curated list, embedded copy, fork copy. | CONFIRMED (prose) | S 317c35fe7, E 19924faab |
| H3 | **Aligned:**<br>- Every rank of Holy Light, Flash of Light and Holy Shock (cast and heal). Divine Favor. Oath of Light's per-level amounts. Light's Vigil (not cast).<br>- Talents: Spiritual Focus, Healing Light, Reverence, Illumination (50% of base cost), Infusion of Light, Boundless Light, Divine Intellect, Conviction, Benediction.<br>- The Devout 5P and 6P. Brilliant Mana Oil. Hand of Edward the Odd (a melee proc, correctly idle). Lay on Hands as a self-cast (heal and mana both to self, which is right). | `aud-paladin`; client rows. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Spirit regen was 15 + Spirit/5 a tick, not the shaman's 17 + Spirit/5.** That is 1 mana a second out of the five-second rule, and the Mindfulness share while casting. | cmangos-classic and vmangos `Unit::OCTRegenMPPerSpirit`: `case CLASS_SHAMAN: addvalue = (Spirit / 5 + 17)`. Test first: `TestShamanSpiritRegenIsSeventeenPlusAFifthOfSpiritATick`. TestElemental and TestEnhancement goldens move by that mana only. **Moves Elemental and Enhancement.** | CONFIRMED-BUG (fresh; server rule) | E 2fc17d1ae |
| S2 | **Aligned:**<br>- Every rank of Healing Wave, Lesser Healing Wave (rank 6 is 10468, not the Burning Crusade 27624), Chain Heal (3 targets, 0.5), Riptide and Healing Stream. Mana Tide by band.<br>- Talents: Water Shield, Tidal Focus, Tidal Mastery, Healing Focus, Healing Way (flat 8/17/25%), Purification, Restorative Totems, Mindfulness, Ancestral Knowledge, Thundering Strikes.<br>- Totems: the shaman's own Healing Stream versus the preset's Mana Spring, with no double count.<br>- Earth Shield and Healing Rain are rune rows only, so they are rightly not learnable. | `aud-shaman`. | Aligned | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Innervate was cast before level 40.** It is learned at 40 (SpellLevels 29166, skill line 573 trainer row), but nothing gated it. At band 20 and 30 a druid gained about 1,200 to 1,700 mana a fight. The site's rank rewrite took its rank-0 row as always known; it now joins the single-tier level overrides and the ladder mirror. | Tests first: `TestInnervateIsLearnedAtForty` (fork) and `TestSingleTierOverridesGateTheirIDs` (site). Conformance drops the Balance rows below 40. **Moves Balance and Feral at bands 20 and 30.** | CONFIRMED-BUG (fresh) | E 60ef5b013, E 7a2832d8d, S 4e7461350 |
| R2 | **I checked every cast in all five healer lists at bands 20 to 50 against SpellLevels:** nothing else unlearned is cast. Nature's Swiftness, Divine Favor, Inner Focus, Mana Tide, Holy Shock and Swiftmend are talent spells the band's build takes. | Breakdown casts against `SpellLevels.csv`. | Aligned | - |
| R3 | **Aligned:**<br>- Every rank of Healing Touch, Regrowth (direct and 7 ticks), Rejuvenation, Tranquility and Wild Growth. Swiftmend's cost and cooldown. Bloom.<br>- Innervate's amount. Nature's Swiftness. Omen.<br>- Talents: Genesis, Nature's Majesty, Nature's Splendor, Gift of Nature, Tranquil Spirit, Improved Rejuvenation, Living Spirit, Improved Tranquility, Reflection (17 against 16.67, trivial).<br>- Elune's Light. Lifebloom, Nourish and Tree of Life are rune rows, not learnable. | `aud-druid`. | Aligned | - |

## Before/after

Before: origin/main 170b7b84c on the fork at cd4dfdcc0. After: lane/r8-healers 419d53a37 with wt/r8-healers. Published entries, heal profile at 1.85, effective HPS.

Raid preset, band 60, 10,000 iterations. Seed 50001 is within 0.1%.

| spec | alliance s1 before | after | horde s1 before | after |
|---|---|---|---|---|
| Restoration druid | 772.7 | 779.4 (+0.9%) | 770.3 | 777.1 (+0.9%) |
| Holy paladin | 656.9 | 662.9 (+0.9%) | 642.0 | 648.1 (+1.0%) |
| Discipline priest | 610.0 | 619.1 (+1.5%) | 614.2 | 619.2 (+0.8%) |
| Holy priest | 682.8 | 686.6 (+0.6%) | 670.8 | 673.5 (+0.4%) |
| Restoration shaman | 823.3 | 841.4 (+2.2%) | 820.7 | 838.6 (+2.2%) |

Bare bands, alliance, 3,000 iterations, seed 1:

| spec | 20 | 30 | 40 | 50 |
|---|---|---|---|---|
| Restoration druid | 55.6 → 42.3 (-24%) | 120.1 → 101.0 (-16%) | 177.4 → 177.4 | 274.3 → 274.3 |
| Holy paladin | 28.4 → 31.4 (+11%) | 60.9 → 73.1 (+20%) | 97.3 → 114.6 (+18%) | 124.8 → 144.9 (+16%) |
| Discipline priest | 35.0 → 36.0 | 84.7 → 86.5 | 141.1 → 143.0 | 200.6 → 203.1 |
| Holy priest | 38.0 → 39.2 | 68.7 → 71.0 | 125.4 → 128.5 | 185.4 → 189.3 |
| Restoration shaman | 50.2 → 50.7 | 85.5 → 86.6 | 144.9 → 146.0 | 225.3 → 226.9 |

The largest shares:
- At 60: the Devout 18 healing for every healer, the shaman's regen, and Divine Grace (human Discipline) and Contingency Plan (gnome Holy).
- Below 60: Innervate (druid at 20 and 30), the kit's Wisdom (paladin), and the level penalty at 20 (every caster and healer, -2 to -4%).

## For other lanes

- **Moves specs outside this group (the integrator re-measures).** Bare, 3,000 iterations, seed 1, unless marked raid:
  - Shaman spirit regen (E 2fc17d1ae): Elemental raid 60 464.7 → 465.2, Enhancement raid 60 unchanged.
  - Level penalty (E 87743552b), every caster at band 20: Fire 34.7 → 33.4, Shadow 36.3 → 35.4, Affliction 56.6 → 54.8, Elemental 22.6 → 22.9 (with the regen). Retribution unchanged.
  - Innervate at 40 (E 60ef5b013): Balance band 20 33.1 → 22.9, band 30 61.8 → 47.4 (both with the penalty). Feral bands 20 and 30 as far as it cast Innervate.
  - "Damage and healing" (E 7783a83ad) and the dot-aura guard (E 7b668da02) move no DPS.
- **Ranker and pipeline (held workflows):** the re-rank must repeat the X7 coverage check. Devout's healing, the level penalty, the kit's Wisdom and the held-item simdb fix (1bd3920b5) will all re-pick healer gear and race.
- **TestRotationLadder** rows move for every healer, every band-20 caster, and Balance and Feral at 20 and 30. Not regenerated here.

## Still open

- **The sub-20 level penalty (X2) is server-core evidence only.** Its formula comes from cmangos and vmangos; the Forever client carries no coefficient penalty field, and no Forever log below level 20 has been fitted. Treat the -2 to -4% at band 20 as unverified until a level-15-to-19 log settles it (wow-player review, integration).
- **Swiftmend's amount.** The engine heals the consumed effect's snapshot with its periodic modifiers and skips Swiftmend's own. The client names Swiftmend in Gift of Nature's direct effect and in a second Genesis effect (1223081 eff1, mask Swiftmend only). vmangos multiplies the consumed tick (modifiers included) by Swiftmend's own modifiers. Either model is within about ±1%. One logged Swiftmend beside the consumed Rejuvenation's tick settles it.
- **The rank rewrite below 60 lifts the deliberate band-60 downranks.** Renew rank 3 and Power Word: Shield rank 4 become the top learned rank at bands 30 to 50. The downrank rule was only ever measured at 60, so this is a rotation policy question for the owner, not a client mismatch.
- **Improved Power Word: Shield scope** (base only, or base plus healing power), **Holy Binding and Divine Aegis on overheal**, and **Inner Focus's crit on Penance's first pulse.** Each needs an in-game log.
- **Racial descriptions that disagree with their spell rows:** Tauren Endurance and dwarf "Earth Fury" say "+2% critical damage and healing", but the rows give health and hit. A tauren or dwarf crit heal settles it.
- **Holy Light rank 9 (25292)** and **Healing Wave rank 10** as book ranks; the Holy priest's **negative haste weight**; and every round-7 blocked item, unchanged.
