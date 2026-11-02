# Round 9 healer sweep: Holy paladin (paladin-holy)

Engine: the fork at scratchpad/r9/healers/engine (HEAD 40021374a). Client: scratchpad/r8/healers/raw (1.60.1.70707 with the hotfix push, fetched 2026-11-01 21:42 UTC). Published build: data/builds/1.60.1.70707/bis/paladin-holy.json. Human (alliance) and undead (horde). Talents 50320215012101251-50322-0503 at 60. Band 20 takes Spiritual Focus 5, Healing Light 3 and Reverence 1. Band 30 adds Reverence 2, Purifying Power 2, Divine Favor and Illumination 5. Band 40 adds Divine Precision 1, Infusion 2, Oath, Holy Shock, Boundless 2, Holy Power 5 and Light's Vigil. Bands 50 and 60 add Divine Intellect 5, Precision 3, Guardian's Favor 2 and Sacred Duty 2 (Protection), then Benediction 5 and Conviction 3 (Retribution).

Check run: `sb -spec paladin-holy -heal -iterations 1000 -seed 1` (raid preset, 60). Effective HPS 662.9, crit 30.07%, 60.9 Illumination procs, 15.3 Holy Shocks. Mana lasts 267.7 s.

## Paladin: Holy

| # | finding | evidence | class | size estimate |
|---|---|---|---|---|
| P1 | **Holy Power's crit-heal reading.** The engine adds 0.10 a rank to the crit bonus, so a crit heal is 1 + 0.5 x 1.5 = 1.75x at rank 5. The client row is aura 50 with MiscValue 2 (Holy school) and no class mask. Points are 10/20/30/40/50 a rank. Aura 50 is not the vanilla aura: mangos-classic `SpellAuraDefines.h` has 50 = SPELL_AURA_MOD_BLOCK_SKILL. So Forever's server uses the modern meaning, MOD_CRITICAL_HEALING_AMOUNT. In TrinityCore that aura multiplies the whole crit heal, which gives 1.5 x 1.5 = 2.25x. The text "critical strike healing bonus by 50%" also fits a flat +50 points, which gives 2.0x. All three readings are open. Round 7 listed this as blocked; the aura-number evidence is new. A side note: being school-wide, it would also reach the Holy heal of Ankh of Life (trivial). | `sim/paladin/healing_talents.go:38-44, 186-196`; `core/spell_outcome.go:1083` (`1 + (1.5*CritMultiplier-1)*CritDamageBonus`). SpellEffect 683158 (spell 5923): EffectAura 50, MiscValue_0 2. TraitDefinitionEffectPoints def 134816, curve 82962. mangos-classic master `SpellAuraDefines.h:185` (fetched 2026-11-02). TrinityCore's aura 50 meaning is from memory, a single source. | OPEN. Settle with one logged Flash of Light crit beside a non-crit, both with known +healing: the ratio is 1.75, 2.0 or 2.25. | At about 30% crit, the average heal multiplier is 1.225 now, against 1.30 or 1.375. That is about +6% or +12% raw healing per mana, less overheal. The largest open item for this spec. |
| P2 | **Concentration Aura is trainable (level 22) and not modelled.** The bare kit runs no aura for a Holy paladin, and the raid preset has no paladin aura, so nothing in the sim competes for the slot. The client row is aura 149 (RESIST_PUSHBACK), 35, all schools ("a 35% chance of ignoring spell interruption when damaged"). mangos-classic sums the caster's SPELLMOD_NOT_LOSE_CASTING_TIME and aura 149 into one chance and rolls it once. Spiritual Focus 5 is 70% (aura 108, misc 9, mask 0xC0000000 covers HL and FoL), so 70 + 35 = 105%: Holy Light and Flash of Light are never pushed back. This answers round 7's "add or roll apart": they add. Pushback is 1 s, then 0.2 s less each hit, as the engine has it. | Client: SpellEffect 697291 (spell 19746): aura 149, base 35, misc 127. SkillLineAbility 11336, AcquireMethod 0. trainables/paladin.json, level 22. Server core: mangos-classic master `src/game/Spells/Spell.cpp:6815-6820`, `resistChance = 100; ApplySpellMod(..., SPELLMOD_NOT_LOSE_CASTING_TIME, resistChance); resistChance += GetTotalAuraModifier(SPELL_AURA_RESIST_PUSHBACK) - 100; if (roll_chance_i(resistChance)) return;` (fetched 2026-11-02). Engine: `core/cast.go:171` rolls one `PushbackReduction`, and nothing adds an aura term. `grep -i concentration` finds nothing in the engine. | OPEN (modelling gap; kit policy for the owner). Settle by adding it to the paladin-holy kit from 22 (a +0.35 aura term on the caster's pushback chance) and measuring. | Round 7 measured pushback, before the focus talents, at -1.3% for the paladin. Spiritual Focus leaves about 30% of it, so the aura would be worth about +0.4% at bands 30-60 (at most +1.3%). |
| P3 | **Oath of Light: scaled by the coefficient (as modelled) or flat.** Not reopened beyond this. Effect 0 is aura 226 (PERIODIC_DUMMY) with a 1500 ms period, and its job is still unexplained. The amounts (6.6667 / 1.9167 / 3.8333 a level, no SpellLevels row) and the cost (PowerCostPct 8) match. Duration index 5 is 300 s. | `oath_of_light.go:19-30, 78-92`; `healing.go:92`. SpellEffect 1367618-1367620. Server cores cannot settle it: the spell is new to Forever and has no core script. | OPEN (round 7 item, unchanged). | Round 7's estimate stands: flat would be larger. |
| P4 | **Holy Light rank 9 (25292) and Blessing of Wisdom rank 6 (25290) as Ahn'Qiraj book ranks.** Their SkillLineAbility rows (13229, 13230) match the trainer ranks' rows (6508 Holy Light 8, 11249 BoW 5) on every column: AcquireMethod 0, ClassMask 2, no RaceMask, no Flags. The known AQ book Greater Heal 5 (25314, row 13214) reads the same. trainables/paladin.json lists both at level 60. The engine leaves both out at 60 (`core/config.go:12 IncludeAQ = false`; `buff_ranks.go:120 AhnQiraj: true`), which is a vanilla rule, not client data. | SkillLineAbility.csv rows above; `data/curated/book-ranks.json` (basis "vanilla" for 25292, "engine" for 25290). | OPEN (unchanged). The client cannot settle it: check a level-60 paladin trainer's window in game. | If both are trainer ranks: Holy Light 1580 against 1217 base (rank 9 costs 660 against 580) on the tank line, and Blessing of Wisdom 40 against 36 mp5 on the bare kit. |
| P5 | **Ruthless and The Human Spirit: the spell rows and the racial table's text disagree.** The Spell rows give 1% crit (20597, base 1) and 10% Spirit (20598, base 10). The engine uses those. ChrRaceRacialAbility 664 and 665 read "Increases critical chance by 2%" and "Increases Spirit by 5%". | `core/racials.go:216-229`; SpellEffect 697219 and 698126; ChrRaceRacialAbility.csv 664/665. | FOREVER-CHANGE as the engine reads it (the spell rows are the newer data, cut on 26 October per the engine's note). Low priority: an in-game character sheet settles it. | About ±1% crit and ±5% Spirit on the human. |
| P6 | **A stale comment, not a sim bug.** `healing_talents.go:46-50` says "Oath of Light is unmodeled ... so ... every Holy Shock heal gets the bonus". The code removes Boundless Light's crit on the sworn ally (`boundlessLightCritOn`, lines 201-206; `holy_shock.go:158-162`). | As quoted. | CONFIRMED (comment only) | 0 |
| P7 | **Sourcing note, outside the engine.** Band 50 human wears Knight's Imbued Armor and Leggings (220813, 220809) from Captain Dirgehammer, a PvP rank vendor (set 1620, Knight-Lieutenant's Imbued Plate). Two pieces, so no set bonus. Whether a bare level-50 character can buy them is a pipeline question. | bis/paladin-holy.json band 50 human; sets.json 1620. | OPEN (for the ranker/pipeline lane) | - |

No CONFIRMED-BUG in the engine.

## Aligned (what I checked)

Each item below was re-read against the client rows; ranks were compared by script, not by eye.

- **Holy Light, all 9 ranks** (635...25292). Base, Variance, per-level growth, SpellLevel/MaxLevel, learn level, ManaCost and coefficient 0.714 all match. Cast index 20 is 2.5 s, GCD 1.5 s.
- **Flash of Light, all 6 ranks.** Same fields, coefficient 0.429, 1.5 s cast.
- **Holy Shock.**
  - Casts 1311606, 20473, 20929, 20930: ManaCost 160/225/275/325, CategoryRecoveryTime 10000, GCD 1500, learn levels 30/40/48/56.
  - Heals 1311605, 25914, 25913, 25903: base 114/156/230/320, Variance 0.0755-0.0789, no per-level growth, MaxLevel 37/47/55/60, coefficient 0.429.
  - The heal ranks 2-4 carry SkillLineAbility AcquireMethod 3. They are child spells of the learned casts, so this does not matter.
  - `clientdamage.Range` (centre ±V/2) reproduces the talent text's 110-118.
- **Class-mask overlaps** (the engine uses its own masks, so I checked each talent's client mask against the spells').
  - Healing Light (0x6000 / mask1 0x10000): HL 0x80004000, FoL 0x40002000 and the HS heal's mask1 0x10000 all overlap.
  - Spiritual Focus (0xC0000000, mask3 0x04000000): HL and FoL overlap; the HS heal does not.
  - Divine Favor (0x80206000): HL, FoL, and the HS cast and heal.
  - Infusion of Light's buff 437063 (0x80000000): HL only.
  - Benediction (0x1BE11FBF...): includes Holy Shock 0x200000. Oath of Light and Divine Favor have no class mask, so Benediction does not apply to them, which matches the engine.
  - Lay on Hands 0x8000 is not in Benediction's mask (and costs nothing anyway).
- **Talent rank curves** (TraitDefinitionEffectPoints) and modifier ops:
  - Healing Light 4/8/12, aura 108 misc 0 (percent amount).
  - Spiritual Focus 14-70, aura 108 misc 9.
  - Reverence 15/30, aura 134.
  - Illumination: chance 20-100; mana is 50% of base cost (fixed effect index 2); the casting-regen buff 20272 is 50% for 15 s (DurationIndex 8). Reverence and Illumination add: aura 134 is summed.
  - Infusion of Light: -500/-1000 on the talent, the buff 437063 a fixed -1000 (misc 10, flat, 1 charge, 15 s). Rank 2 is taken, so both read 1.0 s.
  - Boundless Light 15/30, dummy, kept off the sworn ally.
  - Benediction -2 to -10, misc 14 (cost).
  - Conviction 2/4/6, aura 290 (all crit).
  - Divine Intellect 2-10, aura 137 with stat mask MiscValue_1 = 8 (Intellect).
  - Divine Favor: +100 (aura 107 misc 7), 1 charge, 120 s, 4% of base mana.
- **Untaken and not cast.** Precision, Guardian's Favor, Sacred Duty, Purifying Power and Divine Precision change no heal. Light's Vigil is registered but not cast. **Power in Light** is a priest talent (1309969, SpellClassSet 6), not a paladin one. **Light's Grace** (428909/428912/429243) is a Season of Discovery rune with no SkillLineAbility or trait row, so it is not in Forever.
- **Lay on Hands.** 633/2800/10310, learned at 10/30/50, cooldown 20 min. Heals the paladin's max health and returns 0/250/550 mana. The drain does not restart the five-second rule (`SpendMana` does not touch `FiveSecondRuleRefreshTime`), which matches "does not interrupt Mana regeneration". The autocast fires only below 10% of the paladin's own health.
- **Blessing of Wisdom.** Ranks 19742-25290 give 12/18/24/30/36/40 every 5 s, learned at 14/24/34/44/54/60 (`core/buff_ranks.go:114-121`). Greater Blessing of Wisdom 25894/25918 gives 36/40 (230/250 mana). The kit gives plain Wisdom (not the improved form), and Forever's tree has no Improved Blessing of Wisdom.
- **Judgement of Wisdom and Judgement of Light** (raid debuffs): 33/46/59 mana and 25/34/49/61 health match 20268... and 20267.... The healer neither attacks nor is attacked in melee, so neither affects the heal sim.
- **Racials.** Human: Ruthless and The Human Spirit (see P5); Will to Survive and Perception do nothing in combat. Undead: Touch of the Grave (1260189, attacks only; the healer deals no damage), Will of the Forsaken, Cannibalize. No Forever healer racial exists for either race on a paladin (SkillLineAbility RaceMasks 1 and 16 scanned).
- **Items on the published gear.**
  - Ankh of Life (1713): use 1317424, 720 with Variance 0.2, giving 648-792; 5 min cooldown; used on the tank below 60%. `common/item_effects.go:3482-3521`.
  - Uther's Strength (11302): its equip 8397 procs when the wearer is struck in melee, and the heal profile's pulses are plain health loss (`core/raid_damage.go:13`), so it is correctly idle. Its +12 healing is item stats.
  - Hand of Edward the Odd and Verigan's Fist: melee procs, idle.
  - Shard of the Splithooves, Evergreen Shield, Living Crown and Martyr's Legplates: stats only.
- **Set bonuses.**
  - Vestments of the Devout (182): 2P resistances; 3P 9346 healing 18 (both auras 13 and 135); 4P struck shield (no-sim); 5P 450576, 5% on cast for 450554's 200 mana (modelled); 6P 18378, 8 mp5. Bare 60 wears 6 pieces, raid 60 wears 5.
  - Stormcloth Regalia (1972), Ironfeather Armor (144) and Knight-Lieutenant's Imbued Plate: one or two pieces worn, below every threshold.
- **Learn levels per band.** Flash of Light from 20. Holy Shock from 30, taken at 40. Concentration Aura from 22. Oath of Light is a talent with no SpellLevels row; its amounts scale with the caster's level.
