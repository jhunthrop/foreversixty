# Sim audit round 7: healers

Lane: wt/r7-healers (site, from origin/main 9cfb91c47) and lane/r7-healers (fork, from b57765fb9). Specs: Restoration druid, Holy paladin, Discipline priest, Holy priest, Restoration shaman, and the heal profile.

The short version:
- **Round 6 let the raid's damage reach the healer, but the healer never lost casting time to it.** Hits now push back casts and cut channels short, as any damage does. The four healer "Focus" talents now resist that, as their text says.
- **Mana Spring Totem was dropped whenever Blessing of Wisdom was up.** Since 86d33a100 that covered every raid-preset request on both factions. This moves every mana user, not only healers.
- **Restoration shaman's Healing Stream healing was never counted in its HPS.** About an eighth of its healing.
- **Holy paladin's Oath of Light did nothing.** The engine's note said the spell was not in the client tables. The 70707 push adds it: Blessing of Light's 400/115 on one sworn ally, with Holy Shock +230.
- With these fixes the strongest raid set covered 97% of the heal profile. The profile's own rule retunes it from 1.65 to 1.85.

Engine files touched:
- `sim/core/{cast,raid_damage,dot,unit,buffs}.go`
- `sim/{druid,priest,shaman}/talents.go`
- `sim/druid/bloom.go`
- `sim/priest/priest.go`
- `sim/paladin/{oath_of_light,healing,holy_shock,healing_talents,paladin,talents}.go`
- `sim/shaman/{water_shield,water_totems}.go`
- `sim/common/item_effects.go`
- two APL copies
- seven fixture goldens (Mana Spring only)

Two independent reads were delegated (druid, paladin, shaman engine-vs-client sweeps). Every bug they reported was re-checked against the client row and the code before it was fixed.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Raid pulses that hit the healer never cost casting time.** Round 6 made pulses land on the healer, as plain health loss with no spell hit. The engine's pushback listened only to spell hits, so a priest hit mid-Heal finished on time. In the game any damage while casting delays the cast. | `sim/core/raid_damage.go` `dealRaidDamage`; `cast.go` `applySpellPushback`. The amounts are the engine's existing rule: 1 s, then 0.2 s less a hit, floor 0.2 s. cmangos-classic's `Spell::GetNextDelayAtDamageMsTime` (`1000 - count*200`, floor 200) is the same rule; that is the second source. The engine's own priest notes already called this a "core gap" (`sim/priest/talents.go`, Twilight Focus). Test first: `TestRaidDamagePulsePushesBackTheHealersCast`. | CONFIRMED-BUG (side effect of round 6 X1) | E 60c8e97d8 |
| X2 | **The chance to keep casting time was read off the spell that hit you, not the spell you were casting.** So Spiritual Focus, Eye of the Storm, Fel Concentration and Intensity never applied. Each of them sets it on the caster's own spells. | `cast.go:101` at b57765fb9: `sim.Roll(0,1) > spell.PushbackReduction`, where `spell` is the attacker's. cmangos rolls the caster's `SPELLMOD_NOT_LOSE_CASTING_TIME`. | CONFIRMED-BUG | E 60c8e97d8 |
| X3 | **Channels never lost duration.** Penance, Tranquility and Mind Flay are an instant cast plus a channelled dot and never set Hardcast. A hit now shortens the channel by the same amounts. Ticks past the new end are lost, and the rotation resumes. | cmangos `Spell::DelayedChannel` (same delay table). Test first: `TestRaidDamagePulseShortensTheHealersChannel`. | CONFIRMED-BUG | E 2ef72fb7a |
| X4 | **A pushed cast could end up longer than its full cast time.** cmangos clamps the remaining time to the cast time, so a hit 0.3 s into a 2.5 s Heal costs 0.3 s, not 1 s. | cmangos `Spell::Delayed`. Test first: `TestPushbackNeverRestartsMoreThanTheFullCast`. | CONFIRMED-BUG | E 2a8ac5ee1 |
| X5 | **Nature's Focus, Twilight Focus and Healing Focus were not modelled.** Each was named as "the sim cannot see it". Each sets PushbackReduction on the spells its client mask names, not on the spells its text names. | Client rows, aura 108 misc 9:<br>- Nature's Focus 17063, 14% a rank: Healing Touch, Regrowth, Tranquility, Wrath, Starfire. Not Wild Growth.<br>- Twilight Focus 14913, 23/47/70%: the priest's cast and channelled spells. Not Lightwell.<br>- Healing Focus 16181, 23/47/70%: Healing Wave, Lesser Healing Wave, Chain Heal.<br>Tests first, three `...KeepsCastTimeUnderDamage` tests. | CONFIRMED-BUG | E e74591d21 |
| X6 | **Mana Spring Totem was cancelled by Blessing of Wisdom.** `applyBuffEffects` used `if BlessingOfWisdom {...} else if ManaSpringTotem {...}`. Upstream that was harmless, because one was Alliance and the other Horde. 86d33a100 (2026-10-07) removed the faction gates and kept the `else if`. Since then every raid-preset request lost Mana Spring's 25 mp5 (31.25 improved). | Two spells from two casters: Blessing of Wisdom 19854 (36 mana every 5 s) and Mana Spring 10494 (10 every 2 s, totem 10497). The preset lists both. Test first: `TestBlessingOfWisdomAndManaSpringBothApply`. Seven fixture goldens move by their MP5. **Moves specs outside this group.** | CONFIRMED-BUG (fresh) | E 68e7ca25b |
| X7 | **Ankh of Life's use was not modelled.** The trinket is in the druid, Discipline, Holy priest and shaman sets. Use: ItemEffect 97284, spell 1317424, heals for 720 with variance 0.2, on a 5 min cooldown. Now a cooldown used on the healer's target under 60% health. | `raw/ItemEffect.csv`; spell 1317424. Test first: `TestAnkhOfLifeHealsOnUse`. Effect about +0.3%. | CONFIRMED-BUG (fresh) | E dbfd3e2ef |
| X8 | **Ledger item: the heal profile's coverage was outside its 60 to 85% rule.**<br>- At round 6's integration it was druid 85.2%, paladin 66.3%. The ledger's 86/57 were the healers lane's own pre-integration numbers, so at that point it was not a profile bug.<br>- After X1 to X7 and the fixes below, the strongest raid set covered 97% (shaman) and the druid 89%. That is a profile bug by its own rule.<br>- Sweep on the published gear (3,000 iterations, alliance; raid coverage), in order shaman / druid / Holy priest / paladin / Discipline:<br>&nbsp;&nbsp;- 1.75: 92 / 84 / 76 / 70 / 67<br>&nbsp;&nbsp;- 1.80: 88 / 81 / 73 / 69 / 65<br>&nbsp;&nbsp;- 1.85: 84 / 79 / 71 / 67 / 63<br>&nbsp;&nbsp;- 1.90: 81 / 77 / 69 / 66 / 62<br>- At 1.85, bare sets cover 20 to 40%, and every bare and raid set runs dry before 300 s. 1.85 is the smallest multiplier that meets the rule. | `data/curated/heal-profile.json`: tank hit 1,898 to 2,128, pulse 495 to 555, bounds [1.65, 1.85]. Embedded copy synced. Tests: `TestCuratedHealProfileStatesItsTuning` and the two that pin the figures. Caveat: the ranker re-picks gear at the new scale, so the nightly's sets must repeat the check. | CONFIRMED-BUG (profile, by its own rule) | S 74905fe4c |
| X9 | Ledger item: /tiers/healer order. Bis run 41363895819 was cancelled (owner ruling). The local order on the published gear, lane engine, new priest builds, both bounds, both factions:<br>- Restoration shaman, then Restoration druid, then Holy priest, then Holy paladin, then Discipline.<br>- Holy paladin vs Discipline: paladin ahead at 1.85 on both factions. At 1.65 horde they tie (624.1 vs 623.9).<br>- **The shaman now passes the druid** (823 vs 773 at 1.85, alliance). This is not a reopening of "the druid's lead stays". It is the result of X6 and S1/S2: the shaman was missing an eighth of its healing and a mana source. | Local runs only, never committed. | Reported | - |
| X10 | **The raid's Mana Spring Totem gave rank 4's 25 mp5 at every level.** It was routed by the integrator. The client teaches four ranks: 5675, 10495, 10496 and 10497 at levels 26, 36, 46 and 56 (SkillLineAbility, skill line 374). Their auras (5677, 10491, 10493, 10494) restore 4, 6, 8 and 10 mana every 2 s: 10, 15, 20 and 25 mp5, and nothing below 26. There is no Ahn'Qiraj rank, so level 60 keeps 25. It now reads `core.BuffRanks`, like Strength of Earth (d6c407a19). | Tests first: `TestManaSpringFollowsTheClientRanks`, `TestManaSpringGivesTheRankOfTheLevel`. No golden moves. Published numbers do not move either: bare carries no Mana Spring, and the shaman's own totem was already ranked. Bands 20/30/40/50 bare, both seeds, for Restoration, Elemental and Enhancement: identical. Only a user's Raid-buffed run below 60 moves: Elemental with Mana Spring at 30/40/50 goes 53.3 to 48.9, 92.3 to 88.4, 115.6 to 113.9 DPS. | CONFIRMED-BUG | E cd4dfdcc0 |

## Priest: Holy and Discipline

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Spiritual Healing skipped Binding Heal.** 14898's direct-heal mask carries word 1 bit 4, which is Binding Heal's own class bit (1240774). The engine's mask comment listed it, but the constant left it out when round 6 added the spell. No published build casts both today. | Client SpellEffect and SpellClassOptions. Test first: `TestSpiritualHealingRaisesBindingHeal`. | CONFIRMED-BUG (fresh) | E c148d92a1 |
| P2 | **Holy build: Twilight Focus.** Once X1 and X5 made casts lose time, Twilight Focus 3/3 is worth taking. Its points come from Wand Specialization 2 (no Holy priest wands) and Improved Power Word: Shield 3 (Holy never casts the shield). Mental Strength goes 3 to 5.<br>- At the old profile: +2.7% alliance / +2.4% horde.<br>- At 1.85: +2.0% / +1.8%.<br>- The adoption rule holds: over 1%, Holy tree deepest, no unmodelled talent dropped, nothing the list casts dropped. | heal-builds grid of 33 variants. 10,000 iterations, seeds 1 and 50001 agree within 0.1. | Follows X1 and X5 | S eaaff70f6 |
| P3 | **Discipline build: Twilight Focus.** Twilight Focus 3/3 takes Wand Specialization 2 and Renewed Hope's fifth rank.<br>- +1.8% / +1.9% at the old profile; +1.4% / +1.6% at 1.85.<br>- Taking the points from Inspiration instead scored higher. It was not adopted: the sim cannot see Inspiration's armor, and the rule keeps unmodelled talents. | heal-builds grid; same seeds. | Follows X1 and X5 | S eaaff70f6 |
| P4 | **Stale prose.**<br>- The Holy guide said Holy Binding was unmodelled "(the fight never damages the healer)" and that the client "removed Binding Heal". Both stopped being true in round 6: the talent became Holy Binding, and the spell is trainable.<br>- The Holy list's Prayer of Mending note said it ignores damage; the engine spends it on damage, as the client text says.<br>- The Discipline guide repeated "the client removed Binding Heal". | Guides, the APL note, and the fork copy. | CONFIRMED (prose) | S eaaff70f6, S 062ccc07a (embedded copy), E 719232b80 |
| P5 | **Every priest heal rank matches the 70707 layered tables.** I compared all 69 rows of `sim/priest/heal_ranks.go` (Lesser Heal, Heal, Flash Heal, Greater Heal, Desperate Prayer, Prayer of Healing, Renew, Power Word: Shield, Holy Nova, Penance, Prayer of Mending) by script: mana, cast time, base points, coefficient, variance, growth per level and learn level. Binding Heal's six ranks match too. | `raw/*.csv` from the r7 layered fetch; the script is in the lane scratch. | Aligned | - |
| P6 | **Talent masks and values match the client.** Checked by decoding each client mask: Holy Specialization, Divine Fury, Improved Healing, Twin Disciplines (direct and periodic), Mental Agility, Spiritual Healing (Lightwell excluded), Renewed Hope (2% / 1 s a rank), Divine Aegis (5% a rank, 20 a level cap), Litany of Light (5% a rank), Spiritual Guidance (5% / 1-8%), Meditation (17/33/50%), Inner Focus (25% non-periodic crit), Holy Binding (Flash Heal only). | Client rank descriptions and masks. | Aligned | - |
| P7 | **Holy's Renew at rank 3 is not hit by a downrank penalty.** Classic's only penalty is for spells learned below level 20, and Renew rank 3 is learned at 20. | Client SpellLevels; the list notes. | Aligned | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Oath of Light was taken but did nothing.** `talents.go` said its spell "is not in the client tables yet". The 70707 push adds 1328548: "increasing the healing of your Holy Light on them by up to $s1, your Flash of Light by up to $s2, and your Holy Shock by up to $s3".<br>- Its three effects scale with the caster's level only: 6.6667, 1.9167 and 3.8333 a level, which is 400, 115 and 230 at 60.<br>- 8% of base mana, 5 min, one ally. These are exactly the removed Blessing of Light's amounts, plus Holy Shock.<br>- Now a spell, with Boundless Light's crit kept off the sworn ally, as its text says. The list swears it on the tank before the pull (site and fork copy). | Push-only row; wowforevertalents and foreverarmoury print 400/115/230, but they read the same data, so this is one source. The bonus is scaled by each heal's coefficient: the assumption the removed Blessing of Light made. Tests first: three `TestOathOfLight...` tests. +8.6% / +9.3% alone. | CONFIRMED-BUG (fresh) | E 5ea6a47f3, S 062ccc07a |
| H2 | Guide said "the simulator counts [Oath of Light] as nothing". | `web/src/content/guides/paladin/holy.md`. | CONFIRMED (prose) | S 590585ec7 |
| H3 | All of these match the client, checked against it by a separate read:<br>- Flash of Light 19943 and 19941, Holy Light 10329, Holy Shock 1311606/1311605 (with Benediction's 10%), and Divine Favor.<br>- The Holy talents' rank curves: Healing Light, Spiritual Focus, Reverence, Illumination's mana half, Infusion of Light.<br>- Divine Intellect, Conviction, the Devout set, and the human racials. | Client rows; `flash_of_light.go`, `holy_light.go`, `holy_shock.go`, `healing_talents.go`. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Healing Stream's healing never counted in the shaman's HPS.** Its tick spell (10461) carried `SpellFlagNoMetrics`, which upstream added to tidy the DPS view. Core skips such a spell's metrics, so about 25,000 effective healing a fight was missing from the shaman's figure. The party really was healed. | `sim/shaman/water_totems.go:72`; `core/spell.go` `doneIteration`. Test first: `TestHealingStreamCountsInTheShamansHealing`. About +12% alone. | CONFIRMED-BUG (fresh) | E ac7d5fe99 |
| S2 | **Water Shield never procced on the raid pulses.** Its proc row (SpellAuraOptions 223904 for 408510) has ProcTypeMask 0x362A8, which includes the taken melee, ranged and magic hits. The tooltip says "When a spell, melee, or ranged attack hits the caster". The engine listened only to spell hits, and the pulses are plain health loss. | `water_shield.go`. Test first: `TestWaterShieldReturnsManaWhenARaidPulseHitsTheShaman`. About +7% alone. | CONFIRMED-BUG (side effect of round 6 X1) | E 080eccf1b |
| S3 | Guide said Water Shield returns mana only "on a critical heal". | `guides/shaman/restoration.md`. | CONFIRMED (prose) | S 590585ec7 |
| S4 | All of these match the client: every rank of Healing Wave, Lesser Healing Wave, Chain Heal and Riptide; Chain Heal's 3 targets at 0.5 falloff; Riptide's Chain Heal bonus; Healing Stream with Restorative Totems 50%; Mana Tide rank 4 (290 every 3 s, 4 ticks); Water Shield's globes and 3.5 s gap; every taken talent's mask and value; the Devout set; dwarf and tauren racials. | Client rows; separate read. | Aligned | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Heal-over-time casts could never grant Clearcasting.** Omen rolled only on direct hits and direct heals. Rejuvenation (about half the druid's casts), Tranquility and Wild Growth raise neither. The proc row (16864, ProcTypeMask 0x14014) includes 0x4000, "deal helpful spell".<br>The engine was already inconsistent here: Insect Swarm, a DoT with no hit, procs Omen, and the Devout five-piece (mask 0x14000) procs from Rejuvenation. | `sim/druid/talents.go` Omen aura. Test first: `TestRejuvenationCastsCanGrantClearcasting`. The same 4% a cast and 10 s cooldown. | CONFIRMED-BUG (fresh) | E e4eebd16e |
| R2 | **Bloom could not crit.** It used `OutcomeHealing`, and the code marked that unconfirmed. 1328622's Attributes_2 (0x10004000) lacks SPELL_ATTR2_CANT_CRIT (0x20000000), which the client's triggered heals that cannot crit carry: Judgement of Light 20267, Living Seed 414683. | SpellMisc. One source (the client table). Test first: `TestBloomCanCrit`. About +1%. | CONFIRMED-BUG (fresh) | E 5be4f0fa4 |
| R3 | Guide kept Nature's Focus "because the sim cannot see" it. It now measures level with Heart of the Wild (HPS 772.7 vs 771.0 at 10,000 iterations), and the line says so without numbers. | heal-builds. | CONFIRMED (prose) | S 590585ec7 |
| R4 | All of these match the client: all 38 heal ranks in `healing_tables.go`; tick counts and coefficients; Nature's Splendor; Naturalist; Tranquil Spirit; Improved Tranquility (120 s); Living Spirit; Reflection; Nature's Swiftness; Innervate; dynamic HoTs recomputing at each tick with nothing one-shot lost; the Devout set; Elune's Light. | Separate read; `healing_tables.go`. | Aligned | - |

## Before/after

Before:
- origin/main 9cfb91c47 on the fork at b57765fb9.
- Published talents, heal profile at 1.65.

After:
- Fork lane/r7-healers 2a8ac5ee1 with site wt/r7-healers.
- The new priest builds; the other specs keep their published talents.
- Shown at the old profile (1.65) and at the retuned one (1.85).

All runs: raid preset, band 60, 10,000 iterations, effective HPS, seed 1. The percentages at seed 50001 are within 0.1.

| spec | alliance s1 | after at 1.65 | after at 1.85 | horde s1 | after at 1.65 | after at 1.85 |
|---|---|---|---|---|---|---|
| Restoration druid | 743.3 | 777.0 (+4.5%) | 772.7 (+4.0%) | 741.8 | 775.8 (+4.6%) | 770.3 (+3.8%) |
| Holy priest | 680.7 | 706.1 (+3.7%) | 696.3 (+2.3%) | 662.6 | 692.1 (+4.5%) | 683.3 (+3.1%) |
| Discipline priest | 597.9 | 614.2 (+2.7%) | 618.6 (+3.5%) | 608.1 | 623.9 (+2.6%) | 623.9 (+2.6%) |
| Holy paladin | 578.5 | 634.0 (+9.6%) | 656.9 (+13.6%) | 564.6 | 624.1 (+10.5%) | 642.0 (+13.7%) |
| Restoration shaman | 690.9 | 847.9 (+22.7%) | 823.3 (+19.2%) | 687.7 | 846.3 (+23.1%) | 820.7 (+19.3%) |

The largest shares, alliance seed 1:
- Pushback, alone, before the focus talents and the cap: Holy priest -2.6%, Discipline -2.5% (with channels), paladin -1.3%, druid -0.8%, shaman 0.
- Oath of Light: paladin +8.6%.
- Healing Stream counted: shaman +12%.
- Water Shield on pulses: shaman +7%.
- Mana Spring: every mana healer.

Outside the group, from X6 alone (raid preset, 3,000 iterations, seed 1; the integrator re-measures):
- Fire +1.0, Elemental +1.3, Affliction +0.7, Retribution +0.6, Destruction +0.5, Arcane +0.5.
- Every other spec checked is within 0.3.

## Still open

- **Oath of Light's bonus: scaled by the heal's coefficient (as modelled) or flat?** Flat would be larger. One Flash of Light on a sworn target at 60 settles it. Effect 0's 1.5 s periodic dummy is unexplained.
- **Concentration Aura** is trainable (SkillLineAbility 11336) and not modelled. With Spiritual Focus it would remove most of the paladin's remaining pushback, worth at most about +1.3%. Whether the two add (to 105%) or roll apart is unsettled.
- **Gift of Nature on Tranquility.** Effect 0, the direct modifier, names Tranquility. Whether Tranquility's ticks count as direct heals decides it: about -1% if they do not. One in-game tick settles it.
- **Percent stacking on the same modifier op.** Two cases are unsettled:
  - Rejuvenation: Gift of Nature, Improved Rejuvenation and Genesis are added then multiplied.
  - Healing Wave: Healing Way and Purification are added.
  - Either reading is within about 2%. One measured tick or heal settles each.
- **Healing Wave rank 10 (25357).** Its SkillLineAbility row is indistinguishable from the trainer ranks' rows; it is treated as an Ahn'Qiraj book rank. A trainer check settles it.
- **Chain Heal's party-only jump rule** is not modelled. It does not matter while the list casts Chain Heal only on the tank.
- **Still blocked, no new evidence:** Holy Power's crit heal, Prayer of Mending's bounce mask, Binding Heal's crit rolls, Wild Growth under burst damage, Illumination's casting regen. Illumination has one new client hint: its talent's curve of 10 a rank matches the buff's 50.
- **TestRotationLadder** fails for the healers and for every mana spec (X6). It is not regenerated here: the integrator does that at the pin.
- **The nightly is disabled** (owner). The 1.85 retune and the new priest builds reach published numbers only when bis.yml runs again. The coverage check must then be repeated on the re-picked sets.
