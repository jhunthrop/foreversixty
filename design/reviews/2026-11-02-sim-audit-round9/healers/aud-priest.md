# Round 9 healer sweep: PRIEST (priest-holy, priest-discipline)

Engine: `scratchpad/r9/healers/engine` at 40021374a. Client: `scratchpad/r8/healers/raw` (1.60.1.70707 layered with the push, fetched 2026-11-01 21:42 UTC). Published builds: `data/builds/1.60.1.70707/bis/priest-{holy,discipline}.json`. Holy is gnome / troll; Discipline is human / undead. Helper scripts: `scratchpad/r9/healers/aud-priest-bin/` (sp.py dumps a spell's rows and decodes class masks; ranks.py compares heal_ranks.go with the client). Sim runs: 300 iterations, seed 1, raid preset, in the same folder (`*-alliance.txt`, `holy-horde.txt`).

Published talents (decoded against talents/priest.json, sorted by tier and then column):
- Holy 60: Wand Specialization 2, Twin Disciplines 5, Mental Strength 3, Mental Agility 3, Inner Focus, Improved PW:S 3, Meditation 3 / Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3, Holy Binding, Litany of Light 2, Spirit of Redemption, Spiritual Guidance 4, Spiritual Healing 3, Prayer of Mending.
- Discipline 60: the same Discipline points plus Mental Strength 5, Soul Warding, Penance, Renewed Hope 5 and Divine Aegis 3 / Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3.
- Neither build takes Twilight Focus. The published talents predate round 7, as round 8 X8 notes. Neither build takes Silent Resolve.

## priest-holy

| # | finding | evidence | class | size estimate |
|---|---|---|---|---|
| H1 | **Twin Disciplines no longer raises Prayer of Mending's heal, but its client mask names that heal.** Round 8 P2 removed Twin from Prayer of Mending after checking only the cast (1240827, `0/0x20/0x400/0`) and the jump (401880, `0/0x20/0/0`). The heal is a third spell: 401863, 1313379 and 1313381 for ranks 1 to 3 (Effect 10 HEAL, base 1, mask `0/0/0/0x200000`). Twin's eff0 (aura 108, op 0 healing and damage, +5%) has mask `0x19480020/0x18012/0/0x200000`. Its word-3 bit 0x200000 is carried by no priest spell except those three heals, so the bit exists only to name them. Inner Focus's crit mask (word 3 `0x10200000`) names the same heals. Round 6 H2b recorded that the heals carry only that bit but did not test Twin against it. | Engine `sim/priest/priest.go:96-102` (`PriestSpellMaskTwinDisciplines = ... &^ PriestSpellMaskPrayerOfMending`), applied at `sim/priest/talents.go:139-160`. Test `healing_talents_test.go:102-119` pins the exclusion. Client: SpellEffect for 1225132 (eff0 masks 424149024 / 98322 / 0 / 2097152), SpellClassOptions for 401863 / 1313379 / 1313381 (SpellClassMask_3 2097152). Single source (client masks). There is no log. | CONFIRMED-BUG (round 8 regression; the server applying the modifier to the triggered heal is the one inference) | Holy about +0.6% (round 8 measured -0.6% when it was removed). Prayer of Mending heals about 28k raw a fight. Discipline does not take Prayer of Mending. |
| H2 | **Which talents reach Prayer of Mending's heal, reopened by H1.** The engine heals with the cast spell (`prayer_of_mending.go:89`, `pom.cast`), so Spiritual Healing and Holy Specialization apply, because their masks name the cast and jump. H1 shows the server's modifiers do read the heal spell's own mask. Under that reading Spiritual Healing (`0x9041e20/0x10024/0/0`) and Holy Specialization (`0x9541ec0/0x18024/0/0`) reach the heal only if the server folds the cast's modifiers into the amount it passes on. That is how the tooltip `${($m1+($bh*$bc))*$<mult>}` suggests the amount is built, but `$<mult>` is unread. | Client masks above. Round 6 H2b. | OPEN. Settled by a Holy log showing Prayer of Mending heal crit rate against Flash Heal crit rate (Holy Specialization), or the heal size against Heal's at known healing power (Spiritual Healing). The SpellDescriptionVariables row for `$<mult>` would also settle it. | At most -10% healing and -5% crit on Prayer of Mending, about -1.5% Holy |
| H3 | **The racial heals carry no class mask, so no talent reaches them.** Contingency Plan's heal over time (1277456-1277460, mask `0x20/0/0/0`) is named by Spiritual Healing eff0 and Twin eff0/eff1. The gnome Holy build takes both. | `sim/priest/racials.go:87-104` and `169-193`: Flags `SpellFlagHelpful` only, no ClassSpellMask. Client SpellClassOptions 1277460: mask_0 32. | CONFIRMED-BUG (negligible) | Under 0.05%: about 334 healed a fight, 0.2 procs |
| H4 | **The PoM jump goes to the most hurt other raid member.** The client jump (401880 / 1240919) uses targets 22/30 (area ally around the source) in radius index 9. Nothing in the client says the target is the most hurt one; it may be random. | `sim/priest/prayer_of_mending.go:97-109` (`mostHurtOther`). | OPEN. Settled by a log of jump targets against the members' health. | Changes Prayer of Mending's overheal; unmeasured, probably at most 1% |
| H5 | Spell-modifier stacking: the engine adds percent cost modifiers (Improved Healing -15 and Mental Agility -10 make Prayer of Mending 390 x 0.75). A server that multiplies them gives 390 x 0.85 x 0.9. The same question applies to Renew's three healing percentages. | `core` Cost.Multiplier is additive. cmangos-classic adds percent modifiers; a modern server multiplies them. | OPEN (engine-wide). Settled by a combat-log mana cost of Prayer of Mending with both talents. | About 2% of Prayer of Mending's cost; negligible |

## priest-discipline

| # | finding | evidence | class | size estimate |
|---|---|---|---|---|
| D1 | **Divine Grace (the human racial) gets no Holy Specialization, Twin Disciplines or Spiritual Healing.** Divine Grace 1277370-1277378 has mask `0x1000000/0/0/0`. Holy Specialization eff0, Twin eff0 and Spiritual Healing eff0 all carry word-0 bit 0x1000000. The human Discipline build takes Holy Specialization 5 and Twin 5. | `sim/priest/racials.go:87-141` (no SpellFlagPriest, no ClassSpellMask). Client SpellClassOptions 1277378 mask_0 16777216. | CONFIRMED-BUG (negligible) | One cast a fight, about 1,809 healed: +5% healing and +5% crit is about +130 a fight, under 0.1% |
| D2 | **Devout 4-piece (Divine Protection 27778) is not simulated** (`noSimStruckUtility`). It is a 2% chance when struck, with a 60 s cooldown, to absorb 488 on the priest (27779). Round 6 made the raid pulses reach the healer, so it can now proc. Both priests wear six Devout pieces. | `sim/common/item_sets/client_dungeon_sets.go:68`. SpellAuraOptions 27778: ProcChance 2, ProcCategoryRecovery 60000, ProcTypeMask 139944. | OPEN (low). Settled by whether a raid pulse counts as "struck" in the proc mask, which needs a log. | Under 488 absorbed a minute on the priest; negligible |
| D3 | Known open items, unchanged. Inner Focus's crit on Penance: the client crit mask names the heal pulses 1316991 and others, but `talents.go:412-414` excludes channels. Improved PW:S scope: op 3 (EFFECT1) may raise the base only, not the base plus 0.1 x healing power. Holy Binding and Divine Aegis may read the heal before overheal. Binding Heal may roll one crit or two. | Rounds 7 and 8. | OPEN (no new evidence) | - |

## Aligned (what I checked)

Heal ranks, re-read by script (`ranks.py`) against SpellLevels, SpellPower, SpellMisc cast time and the heal effect's base, variance, growth and coefficient. All of these match:
- Lesser Heal 3 ranks, Heal 4, Flash Heal 7, Greater Heal 5, Desperate Prayer 7, Prayer of Healing 5.
- Renew 10 ranks (3 s period, 15 s duration, 0.2 a tick).
- PW:S 10 ranks (0.1 coefficient, growth to MaxLevel, 4 s cooldown, 30 s duration).
- The Holy Nova heal 6 ranks.
- Penance 4 ranks. The heal helpers 402289, 1240723, 1240724 and 1316991 have MaxLevels 39, 49, 59 and 60 and growth 0.87, 0.98, 1.27 and 0. The channel 1316992 lasts 2 s with a 1 s period, giving 3 pulses. The cooldown is 12 s.
- Prayer of Mending 3 ranks: 172, 298 and 413 at 0.429; 10 s cooldown; 5 charges (aura CumulativeAura 5); 30 s after each jump.
- Binding Heal 6 ranks: 247 to 825, 0.429, growth to MaxLevel.
- Learn levels against `trainables/priest.json`: Binding Heal 25-56, Penance 30/40/50/60, Prayer of Mending 40/50/60, Lightwell 40/50/60, Divine Grace 10-58, Contingency Plan 20-60, Dark Sacrifice 20-60.

Mechanics:
- Weakened Soul 6788: 15 s.
- Soul Warding 402000: flat -4000 ms on the cooldown (op 11) and -15% cost (op 14) on mask 0x1, PW:S only. A measured cost of 131 for rank 4 matches 175 x (1 - 0.15 - 0.10).
- Renewed Hope: 2% crit a rank and 1 s of Weakened Soul a rank (trait curves 2..10 and -1000..-5000). It names Flash Heal, Lesser Heal, Heal, Greater Heal and Penance, not Binding Heal.
- Divine Aegis: 5/10/15% (curve 83134), cap 20 x level.
- Inner Focus 14751: -100% cost, +25% crit, 1 charge, 3 min cooldown.
- Inner Fire: armor only (aura 674), 10 min; Improved Inner Fire 45%.
- Lightwell 27871: 365 mana, 1.5 s cast, 10 min cooldown, 180 s or 5 charges. Its renew 27874 is 320 every 2 s for 10 s with no coefficient, carries the PERIODIC_CAN_CRIT flag, and no talent mask names it (the engine registers it without the priest flag, which matches).
- Dark Sacrifice 1277328: 320 health and 320 mana every 3 s for 15 s, plus Spirit; 10 min cooldown; undead only (RaceMasks 16). Holy's list line never fires for gnome or troll, which is correct.
- Spirit of Redemption 20711: a dummy aura with no passive, so not modelled correctly.
- Silent Resolve 14523: threat -30% on Holy and Shadow, plus control durations. Not taken; threat only.

Talent values (TraitDefinitionEffectPoints curves) match the engine tables: Spiritual Healing 3/7/10, Mental Strength 2/rank (stat mask 8, Intellect), Twin 1/rank, Spiritual Guidance 5/rank healing and 1/3/5/6/8 damage, Holy Specialization 1/rank, Litany of Light 5/10, Meditation 17/33/50, Mental Agility 3/7/10, Improved PW:S 7/14/20, Inspiration 8/17/25, Improved Healing 5/rank, Divine Fury -100 ms/rank, Twilight Focus 23/47/70.

Talent masks, decoded against every priest SpellClassOptions row:
- Divine Fury: Heal, Greater Heal, Smite, Holy Fire.
- Improved Healing: Lesser Heal, Heal, Greater Heal, Penance cast, Prayer of Mending cast.
- Mental Agility: includes Renew, PW:S, Holy Nova and Prayer of Mending; excludes Penance.
- Spiritual Healing: direct heals plus Renew via eff1 op 22; excludes Lightwell.
- Holy Specialization: heals including Renew, Binding Heal and Penance.
- Twin Disciplines: Renew via eff1, Penance heals, Holy Nova, Desperate Prayer. Prayer of Mending's heal is the exception, see H1.
- Improved Renew: op 22 on Renew.
- Holy Binding 1328569: dummy, Flash Heal only.

Racials:
- Gnome Eureka! 1259823: 8 s, 2 min, +10% healing and -10% cost. The generated set covers Flash Heal, Heal, Greater Heal, Prayer of Healing, Penance and Binding Heal, but not Renew, Prayer of Mending or PW:S. It is cast 3 times a fight.
- Expansive Mind: covered by round 4.
- Contingency Plan: ward 1277640 lasts 30 s at 35%; absorb 1277466 is 926 for 15 s; heal 1277460 is 134 every 3 s at 0.2; 10 min cooldown; rank levels match.
- Troll Berserking 20554: 10%, 15 s (DurationIndex 8), 3 min, free. The 10% variant fires first, 2 casts a fight. Hex of Weakness and Shadowguard affect enemies or melee attackers, so they have no healing effect.
- Human Divine Grace: matches round 8 apart from D1.
- Undead Dark Sacrifice: listed above.

Sets and items:
- Vestments of the Devout 182: 2-piece +8 resistances, 3-piece +18 healing (9346), 5-piece Sudden Insight (5% on cast for 200 mana, measured 5.9 procs against about 6 expected), 6-piece 8 mp5. The 4-piece is D2.
- Ankh of Life use: covered by round 7.
- Uther's Strength (8397, a melee-struck shield) and Truefaith Vestments (18388, Fade cooldown) do nothing for healing.
- Priest T1 (Vestments of Prophecy 202) is not on any published set.
