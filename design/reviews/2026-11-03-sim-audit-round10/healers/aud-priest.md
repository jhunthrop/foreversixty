# Round 10 healer sweep: PRIEST (priest-holy, priest-discipline), level 60

Engine: `scratchpad/r10/healers/engine` at 4d3129cae. Client: `scratchpad/r9/casters/raw` (1.60.1.70707 with the push layered in). Site: `.claude/worktrees/r10-healers`. Sim outputs read: `r10/healers/a0-priest-{holy,discipline}.txt`, `b-priest-holy.txt`. My scripts are in `r10/healers/aud-priest-bin/`: `ranks.py` compares every `heal_ranks.go` row with SpellLevels, SpellPower, SpellMisc cast/duration, SpellCooldowns and the heal or aura effect. `mask.py <spell>:<effect>` decodes an effect's class mask against every SpellClassSet 6 row. `curves.py` dumps the trait curves. `items.py` lists ItemXItemEffect spells.

Published 60 builds (decoder `tools/talents.py`; raid and bare are the same):
- Holy `025003031330000000-03505003030121431`, troll (horde) / gnome (alliance). Wand Spec 2, Twin Disciplines 5, Mental Strength 3, Mental Agility 3, Inner Focus, Improved PW:S 3, Meditation 3 / Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3, Holy Binding, Litany of Light 2, Spirit of Redemption, Spiritual Guidance 4, Spiritual Healing 3, Prayer of Mending.
- Discipline `025005031330011530-03505003030000000`, undead (horde) / human (alliance). The same Discipline points with Mental Strength 5, plus Soul Warding, Penance, Renewed Hope 5 and Divine Aegis 3 / Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3. No Power Infusion.

Casts at 60 (APLs): Holy casts Lightwell r3 (prepull), Prayer of Mending r3, Renew r3, Flash Heal r2 (tank) and r7, Prayer of Healing r4, Heal r4, Inner Focus, Inner Fire, Dark Sacrifice (undead only, so never for Holy), Berserking or the gnome racials, and the Ankh of Life (troll raid). Discipline casts Lightwell r3, Penance r4, PW:S r4, Renew r3, Binding Heal r6, Flash Heal r7, Prayer of Healing r4, Greater Heal r4, Heal r1, Inner Focus, Inner Fire, and Dark Sacrifice or Divine Grace.

## Findings

| # | finding | engine | client | class | size |
|---|---|---|---|---|---|
| P1 | **Mental Agility does not reduce Inner Fire's cost.** | `sim/priest/priest.go:114-117`: PriestSpellMaskMentalAgility has no Inner Fire bit. `sim/priest/inner_fire.go:35-53`: Inner Fire has no ClassSpellMask, so no masked mod can reach it. | SpellEffect 687728 (14520 e0: aura 108, op 14, -10 at rank 3 on curve 135372), mask word0 0x8251c1eb. Inner Fire 588/7128/602/1006/10951/10952 have SpellClassOptions mask_0 0x2, and 0x8251c1eb & 0x2 = 0x2. All six ranks decode inside the mask. | CONFIRMED-BUG (negligible; single source: the client masks) | One cast a fight: 315 → 283.5 mana, 31.5 mana a fight. Under 0.01%, both specs. |
| P2 | **Inner Focus's +25% crit reaches heal-over-time ticks and Prayer of Mending charges that land while it is up.** The aura adds the crit to every priest spell with a cost that is not PureDot or Channeled. Renew has neither flag. HoT ticks read BonusCritRating when they tick. So a Renew already ticking, and any PoM charge spent (PoM heals through its cast spell), gets +25% until the next priest cast completes. | `talents.go:415-417` (innerFocusAddsCrit), `431-437`; `heals.go:9` (priestHealFlags, no PureDot); `core/spell_outcome.go:78-86` (the tick crit reads HealingCritCheck live); `prayer_of_mending.go:89`. | SpellEffect 692357 (14751 e1: aura 107, op 7, +25) decoded: no Renew rank in the mask. The mask does name the PoM heals (word3 0x200000), but the aura has 1 charge (SpellAuraOptions ProcCharges 1), and the PoM cast spends it at cast time. Text: "if it is a non-periodic spell". | CONFIRMED-BUG (negligible) | 2 Inner Focus a fight, each window up to 2.5 s (a Heal cast). About 1.6 Renew ticks and 1.3 PoM charges a fight get the bonus: about 60 healing a fight, under 0.05%. Fix: also exclude spells with a Hot config, and do not let a PoM charge read the cast's IF crit. |
| P3 | **Lightwell Renew's crit: the client DefenseType is NONE, and the engine overrides it to Magic.** The engine registers the renew on the priest with DefenseTypeMagic. It rolls the priest's spell-table crit without Holy Specialization, which matches the renew's empty class mask. Holy measures 15.4% tick crit, Discipline 16.3%. | `sim/priest/lightwell.go:96-133` (DefenseType line 101). `core/spell_outcome.go:1072` panics on DefenseTypeNone in CritMultiplier, so the override is what makes the crit path run. | SpellCategories 27874: DefenseType 0. SpellMisc 27874: Attributes_8 0x200 (PERIODIC_CAN_CRIT). SpellClassOptions: none. Lightwell 27871 e0 is Effect 50 (summon gameobject 181106), so the renew is cast by the well, or by its trigger with the priest as original caster, not by a priest spell. In TrinityCore-lineage servers, `SpellCritChanceDone` returns 0 for DmgClass NONE, and a non-player-controlled caster cannot crit. That is an inference, not a Forever source. | OPEN. The rule in force (HoTs follow PERIODIC_CAN_CRIT) is kept. DefenseType 0 may be a second gate. Settled by a Lightwell Renew crit tick in a log. If ticks never crit, drop the crit. If they do, they crit off the priest's spell crit with no Holy Specialization, as now. | At most about 400 raw healing a fight each: Holy 16.1 ticks × 15.4% × 0.5 × 320; Discipline 15.8 × 16.3%. About 0.2% raw, less effective. |
| P4 | **Renewed Hope's Weakened Soul cut is applied once per Penance pulse.** healTarget calls shortenWeakenedSoul for every pulse. Three pulses × 5 s = 15 s, so one Penance always clears Weakened Soul. That reopens PW:S on the tank and ends Renewed Hope's own +10% crit window for the later pulses and heals. | `heals.go:94-101`; `penance.go:39-45`; `healing_talents.go:123-137`. | 425280 e1: aura 4 (dummy, server script), curve 135365 -1000 to -5000. Text: "Your heals from ... Penance ... reduce the remaining duration of Weakened Soul on their target by 5 sec". The pulses are separate heal spells: 1316991, triggered by 1316992 e1 (aura 23, period 1000, DurationIndex 39 = 2 s). Per heal that is 3 cuts; per cast it is 1. The client does not say which. | OPEN. Settled by a log: Weakened Soul time left after a Penance on a shielded target. | Discipline only. PW:S is 8 casts and 2,754 absorbed a fight, and the Renewed Hope crit window is involved. Within ±0.5% of Discipline effective HPS. |
| P5 | **Sudden Insight (Devout 5-piece) rolls once per completed cast. The client proc is magic-class and may roll per target hit.** | `common/item_sets/client_rows.go:130-134` (onSpellcast: OnCastComplete, ProcMaskSpellDamage or ProcMaskSpellHealing); `client_procs.go:28-45`. | SpellAuraOptions 450576: ProcChance 5, ProcTypeMask 0x14000 (DONE_SPELL_MAGIC_DMG_CLASS_POS and _NEG); trigger 450554, 200 mana. The PoM cast (1240827) and the Penance cast (1316995) have DefenseType 0, so they cannot roll on a magic-class mask. Penance's pulses (1316991), the PoM heals (401863/1313379/1313381), Prayer of Healing and Binding Heal have DefenseType 1. A modern server procs DONE flags per target hit: Prayer of Healing ×5, Binding Heal ×2, Penance ×3. | OPEN (cross-class: the S03 mana bonus is shared by other dungeon sets). Settled by a log of Sudden Insight procs against casts. | Holy: −27 rolls (the PoM casts) up to +62 (the PoM heals and Prayer of Healing targets). Discipline: −18 (the Penance casts) up to +68 (pulses, Prayer of Healing targets, the Binding Heal self half). About ±600 mana a fight, about 2% of the mana budget. Discipline is mana-bound (spend 36.6k against 27.9k income plus an 9.0k pool), so it is worth up to about 1% of Discipline effective HPS. |
| P6 | **Litany of Light compares spell codes, so two ranks of one spell count as the same spell.** | `healing_talents.go:229-238` (`previous != spell.SpellCode`). | 1317006: aura 42 with no trigger spell (server script). Text: "if your previous heal was a different spell". Holy casts Flash Heal r2 (9472) and r7 (10917), which are different spell ids. | OPEN. Settled by a log of a Litany gain between two ranks of one spell. | 0.3 r7 casts a fight, under 40 mana. Negligible. |
| P7 | **Trait curves with no client effect.** Holy Specialization's definition 135395 has eff0 and eff1 curves, and a second definition 142601 points at the same spell. The client spell 14889 has one effect (692220). Twin Disciplines 135379 has an eff2 curve, but 1225132 has two effects (1214136, 1239165). A server-only effect would not be in the client. | - | TraitDefinitionEffectPoints / CurvePoint, as above. The Wowhead text has one sentence each. Single source. | OPEN (probably leftovers; nothing to model without a server line) | - |

Carried over from round 9, re-read against the client masks, no new evidence: Prayer of Mending's heals (401863...) carry only word3 0x200000, so Spiritual Healing and Holy Specialization name the cast and jump but not the heal (H2). The jump target, "most hurt" or random, is open (H4). Same-op percent stacking is open (H5). Devout 4-piece Divine Protection 27779 is still not simulated (D2); Soul Warding and Improved PW:S both name it, word0 0x1. Also still open: Inner Focus on Penance's pulses (the 692357 mask names 402289/1240723/1316991; the engine excludes channels), the scope of Improved PW:S's op 3, whether Divine Aegis and Holy Binding read the heal before overheal (the engine uses raw result.Damage), and whether Binding Heal rolls one crit or two.

Trivia, no sim effect at 60:
- `lightwell.go:149` says "duration index 26". 27871 is DurationIndex 25 (180,000 ms); index 26 is 240,000. The value is right and the comment is wrong.
- Binding Heal rank 2 (1240770) e1 grows 2.2 a level against e0's 2.5. The engine uses 2.5 for both halves. That is a 1.5 point gap on the self half at level 60, and rank 2 is not cast at 60.

## Crit split (round 9) check

- Heals use `healingCritRating` = spellTableCritRating + BonusCritRating (`core/spell_result.go:277`). HoT ticks use the same through `outcomeHealingCritTick` (`spell_outcome.go:78`).
- Holy Specialization, Renewed Hope and Inner Focus all add to BonusCritRating, so all three reach heals.
- Moonkin Aura 24907, Leader of the Pack 24932, Ruthless 20597 and the Elixir of Sages crit (1250941) are all aura 290, so they give every table, as the rule says.
- Measured crit rates fit:
  - Holy is 7.03 + 8 + 5 = 20.03, against Flash Heal 20.0, Heal 20.4, PoM 20.2 and Renew ticks 20.2.
  - The Ankh of Life is 14.7, with no Holy Specialization, which is right because it has no class mask.
  - Lightwell Renew is 15.4.
  - Discipline is 7.39 + 8 + 5 = 20.39, against Heal 20.6. The rest are higher because of Renewed Hope and Inner Focus.
- The crit-reading effects all read result.DidCrit or a tick outcome, so the split cannot break them: Divine Aegis and Inspiration (OnHealDealt, non-periodic only), and Holy Binding (no crit roll; dummy 1328569).
- The published gear has no school- or ability-limited crit talent and no on-crit item proc.

## Aligned (checked against the client)

Ranks, by script (all rows match base, variance, growth, max level, coefficient, mana, cast time, learn level):
- Lesser Heal 1-3, Heal 1-4, Flash Heal 1-7, Greater Heal 1-5, Prayer of Healing 1-5, Desperate Prayer 1-7 (dwarf only; no published priest is a dwarf).
- Renew 1-10: aura 8, 3 s period, 15 s, 0.2 a tick.
- PW:S 1-10: aura 69, 0.1, 30 s, 4 s category cooldown.
- Holy Nova damage and heal 1-6.
- Penance 1-4: heal helpers 402289/1240723/1240724/1316991; 0.19 a pulse; 12 s category; the channel 1316992 lasts 2 s with a 1 s period, giving 3 pulses.
- Prayer of Mending 1-3: 172/298/413 at 0.429; 10 s cooldown; 5 charges (401859 e1 bp 5, 401877 CumulativeAura 5); holder aura 30 s (DurationIndex 9).
- Binding Heal 1-6: two heal effects, 0.429, 1.5 s, mana 155 to 380; see trivia for rank 2's self half.
- Inner Fire 1-6: aura 674 armor 315 to 1,395, 10 min.
- Heal r1 (level 16) takes the sub-20 penalty of 0.85 (`core/spell.go:340`, 787-806). Flash Heal r2 (26) and Renew r3 (20) take none.

Spell by spell:
- **Flash Heal, Heal, Greater Heal, Lesser Heal.** The heal goes through healTarget. Renewed Hope names exactly Flash Heal, Lesser Heal, Heal, Greater Heal and Penance. Divine Fury's 18530 mask names exactly Heal, Greater Heal, Smite and Holy Fire, at -100 ms a rank. Improved Healing's 688700 mask names exactly Lesser Heal, Heal, Greater Heal, the Penance cast, and the PoM cast and jump. Holy Specialization (692220) and Spiritual Healing (692104) name all four.
- **Prayer of Healing.** Targets 132/31 (the target's party); the engine heals partyOf(target), each member with its own roll and crit. Holy Specialization and Spiritual Healing reach it. Mental Agility, Improved Healing and Twin do not.
- **Renew.** Twin reaches it through e1 (op 22, mask 0xa08060). So do Improved Renew 691708 (op 22, 0x40) and Spiritual Healing e1 1239404 (op 22, 0x40). Mental Agility reaches it. Holy Specialization reaches the ticks; 6075 attr8 0x1000200 carries PERIODIC_CAN_CRIT.
- **Prayer of Mending.** Twin (the round 9 fix is present: PriestSpellMaskInstantHealing includes PoM), Mental Agility and Improved Healing reach it. A charge is spent by any non-periodic heal taken (OnHealTaken, not OnPeriodicHealTaken) or by raid damage; 401877's ProcTypeMask 0x800AAAA8 includes taken-positive and taken-damage.
- **PW:S.** Improved PW:S 692336 (op 3, mask 0x1) at 7/14/20. Soul Warding 1048187/1048188: -4,000 ms cooldown and -15% cost, mask 0x1. Mental Agility reaches it. Weakened Soul 6788 lasts 15 s (DurationIndex 8). Divine Grace clears it (1277378 e1, Effect 77).
- **Penance.** Holy Specialization, Twin e0 and Spiritual Healing e0 name the pulses (word1 0x10000). Improved Healing names the cast (word1 0x800000). Mental Agility excludes it. The channel locks casting (`core/cast.go:276`).
- **Binding Heal.** Holy Specialization and Spiritual Healing reach it (word1 0x4). Twin, Mental Agility and Renewed Hope do not.
- **Lightwell.** 724/27870/27871: 225/295/365 mana; 1.5 s; 10 min category; 180 s; 5 charges. Renew 7001/27873/27874: 160/233/320 every 2 s for 10 s, no coefficient. No talent mask names the renew, and the engine registers it without the priest flag. Crit is P3.
- **Inner Focus 14751.** -100% cost (692356, op 14). Its mask names every cast here, Lightwell and Inner Fire included. 3 min cooldown, 1 charge, consumed on cast completion.
- **Holy Binding 1328569.** Dummy 50%, Flash Heal only, skipped when the target is the priest, no crit and no modifiers.
- **Litany of Light.** Curve 135385 gives 5/10% of the base cost (the flat cost). It fires on healing casts and not on PW:S or Lightwell. See P6.
- **Divine Aegis.** 431622 curve 5/10/15; cap 20 × level; shield 431624 lasts 12 s (DurationIndex 29). It fires on non-periodic crits only (ProcTypeMask 0x4400, no periodic bit). Penance pulses and Prayer of Healing members fire it.
- **Renewed Hope.** 2%/rank crit and 1 s/rank; the per-pulse question is P4.
- **Other talent values (curves).**
  - Twin Disciplines: 1/rank.
  - Mental Strength: 2/rank, aura 137, stat 8 (Intellect).
  - Mental Agility: 3/7/10.
  - Meditation: 17/33/50, aura 134.
  - Holy Specialization: 1/rank, op 7, Holy school priest spells. Every cast the engine reaches is inside 692220's mask. The mask excludes Lightwell, PW:S, Inner Fire, the Contingency Plan heal over time and Divine Aegis, and the engine does not apply it to them.
  - Spiritual Guidance: 5/rank healing (aura 175, Spirit) and 1/3/5/6/8 damage (aura 174). HealingPower is the healing stat alone, so nothing is counted twice.
  - Spiritual Healing: 3/7/10.
  - Inspiration: 8/17/25 armor. Held for Discipline; armor is not modelled.
  - Spirit of Redemption: no passive.
- **Racials.**
  - Troll Berserking 20554: 10% (auras 319, 140, 65), 15 s (DurationIndex 8 with the push), 3 min.
  - Gnome Eureka! 1259823: masks as round 9 recorded; the e2 periodic mask covers no Renew. Gnome Expansive Mind 20591: aura 178, +5% mana.
  - Gnome Contingency Plan: ward 1277640 is 35%, 30 s; shield 1277466 absorbs 926; heal 1277460 is 134 every 3 s at 0.2 with the crit flag. No SpellPower row, so it is free. 10 min category.
  - Human Divine Grace 1277378: 1,383, variance 0.165, 0.429, below 50%, clears Weakened Soul, 10 min, free. Human Spirit 20598 is 10% Spirit with the push.
  - Undead Dark Sacrifice 1277328: 320 health and 320 mana every 3 s for 15 s, plus Spirit; 2,042 mana measured, which matches 1,600 + 442 Spirit.
  - Undead Touch of the Grave: damage only.
- **Gear.**
  - Vestments of the Devout 182 (six pieces for Holy, Discipline bare alliance and Discipline raid alliance; five for Discipline raid undead): 2-piece 18679 (resistances 8), 3-piece 9346 (aura 13 and aura 135 at 18), 5-piece 450576 (P5), 6-piece 18378 (8 mp5). The 4-piece is carried (D2).
  - Ankh of Life 1713 (ItemEffect 97284 → 1317424): 720, variance 0.2, DefenseType 1, 5 min item cooldown, 120 s category 1153.
  - Uther's Strength (8397 → 10368): 4%, melee-struck only (ProcTypeMask 0x28), so no healing effect.
  - Truefaith Vestments 18388: Fade cooldown only.
  - Shard of the Splithooves, Torch of Light, Glowing Brightwood Staff and the other picks have no ItemEffect spells, so they are stats only.
- **Raid consumables.**
  - Elixir of Sages 1250941: Spirit 25 and 2% crit (aura 290).
  - Greater Cleric's Elixir 1250926: 40 healing.
  - Greater Mageblood 1250948: 20 mp5.
  - Gizzard Gum 10693: 25 Spirit.
  - Sage's Tea 1249925 → 1249927: 44 healing.
  - Brilliant Mana Oil, enchant 2629 → 25116: 15 mp5 and 30 healing.
  - Flask of Distilled Wisdom 17627: aura 418, 2,000 mana.
  - Major Mana Potion 17531: 1,800, variance 0.5; measured 1,817 a use.
  - Demonic Rune 16666: 1,200 mana, variance 0.5, and 800 damage; measured 1,192 a use.
