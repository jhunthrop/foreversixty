# Round 10 healer sweep: Holy paladin (paladin-holy), level 60

Engine: scratchpad/r10/healers/engine (HEAD 4d3129cae). Client: scratchpad/r9/casters/raw (1.60.1.70707 plus the hotfix push, fetched 2026-11-01 21:42 UTC). Site: worktree r10-healers.

- **Published build at 60 (raid and bare, both factions):** 50320215012101251-50322-0503, decoded with tools/talents.py.
- **Races:** human (alliance) and undead (horde).
- **Rotation casts:** Oath of Light 1328548 on the tank (before the pull, then on expiry); Divine Favor 20216; Holy Light rank 8 (10329); Holy Shock rank 1 (1311606, heal 1311605); Flash of Light rank 1 (19750) on the tank and rank 4 (19941) on the party; autocast cooldowns (Major Mana Potion, Demonic Rune, Lay on Hands 10310 below 10% of own health).
- **Raid 60 gear:** Living Crown 252561, Glowing Eye of Mordresh 10769, Devout Mantle/Robe/Bracers/Gloves/Belt (5 of set 182), Imperial Red Cloak 8248, Martyr's Legplates 250600, Mooncloth Boots 15802, Band of the Hierophant 13096, Glowing Crystal Ring 18402, Shard of the Splithooves 10659, Uther's Strength 11302, Hand of Edward the Odd 2243, Evergreen Shield 279262.
- **Bare 60 gear:** 6 Devout pieces (with the Sandals), Horizon Choker 13085, Eye of Adaegus 5266, Elunarian Sphere 15968, Death Speaker Scepter 2816 (undead).

Method: scripts in scratchpad/r10/healers/paltools/.
- `mods.py` lists every SpellClassSet-10 effect whose class mask overlaps a cast spell. I then traced each hit's source (talent trait, ItemEffect, ItemSetSpell, SkillLineAbility, trigger).
- `curves.py` prints the trait rank curves.
- `meta.py` prints SpellPower, cooldown, cast, duration, level and category.
- `items.py` and classic-db item_template list the item effects.
- Wowhead Forever tooltips were read for 10310, 20216, 1328548, 1311606 and 20210.

## Findings

| # | finding | engine | client | class | size |
|---|---|---|---|---|---|
| P1 | **Holy Power's crit-heal size** (left to the lane lead, as asked). New client note: only two SpellEffect rows in the whole client carry aura 50, Holy Power 5923 (bp 0, misc 2, trait curve 82962 = 10/20/30/40/50) and **21540 "Improved Block Value" (bp 20, misc 0)**. So the client shares aura 50 between a crit-healing talent and a block-named row. That weakens the "Forever uses the modern meaning for aura 50 everywhere" argument a little. It does not decide 1.75, 2.0 or 2.25. | `healing_talents.go:44,185-195` (+0.10 a rank to CritDamageBonus on PaladinSpellMaskIlluminating); `core/spell_outcome.go:1083` | SpellEffect 683158 (5923), and the 21540 row | OPEN (log ratio), unchanged | as round 9: +6% or +12% |
| P2 | **Heals ignore school crit.** `healingCritRating` is spellTableCritRating + BonusCritRating only. It leaves out `PseudoStats.SchoolBonusCritChance`, which `SpellCritChance` (damage) adds. **For the paladin this is zero today:** nothing in sim/paladin, the gear or the preset writes a Holy school crit. Holy Power is no longer a crit-chance talent (aura 50, not 57 with a school). Conviction, Ruthless, Moonkin Aura, Leader of the Pack and Elixir of Sages are all aura 290 and go to stats.Crit. **Cross-lane flag:** `priest/item_sets_pve.go:22` (Vestments of Prophecy 5P, "+2% crit with Holy spells") and `shaman/item_sets_pve.go:53` write SchoolBonusCritChance, so those bonuses never reach a heal. The priest and shaman sweeps should check them. | `core/spell_result.go:277-282` vs `:251-257` | Moonkin Aura 24907 e0 aura 290 bp 3; Leader of the Pack 24932 aura 290 bp 3; Ruthless 20597 aura 290 bp 1; Elixir of Sages 1250941 e1 aura 290 bp 2; Conviction 20117 aura 290 (curve 2/4/6) | Aligned for the paladin; OPEN for the priest and shaman | paladin 0 |
| P3 | **Illumination's refund on Holy Shock rank 1.** The engine refunds 50% of the cast's cost, 80 mana from 1311606's 160, when the heal 1311605 crits. The client heal has no SpellPower row. Vanilla cores make this work with a hard-coded map from heal to cast (mangos-classic's Illumination handler maps 25914→20473, 25913→20929, 25903→20930; **from memory, single source**). 1311605/1311606 are Forever-new ids that no public core maps. If Forever's script reads the heal's own cost, rank 1 refunds 0. | `holy_shock.go:165` (markIlluminating with rank.manaCost); `healing_talents.go:171-180` | SpellPower: 1311606 ManaCost 160, 1311605 none. 20210 e0 aura 42 → 18350 (dummy), ptm 0x4400; refund e2 bp 50 fixed. | OPEN. Settle with a combat log: a rank-1 Holy Shock crit should show a 20272 energize of 80. | Up to about 15 HS a fight × ~60% crit (Boundless 30% on the party) × 80 ≈ 700 mana. About +0.5% HPS at most if the engine is right and the server is not, and only while mana runs out. |
| P4 | **Devout 5P on off-list casts.** The engine procs it on OnCastComplete for spells with SpellDamage or SpellHealing proc masks. Divine Favor (no proc mask) never procs it. The client proc row 450576 is ptm 0x14000 (helpful and harmful magic spells, 5%). DF is DefenseType 1 (magic), so on the server it probably procs. Oath of Light has no SpellCategories row (DefenseType 0) and is correctly excluded. | `common/item_sets/client_rows.go:130-134`; `divine_favor.go:61-76` | SpellAuraOptions 450576 ProcChance 5, ptm 0x14000; SpellCategories 20216 DefenseType 1 | OPEN (trivial) | one cast every 2 min × 5% × 200 mana: under 0.1 mp5, 0.0% |
| P5 | **Raid preset stacks three Forever elixirs and a flask.** The preset carries Greater Mageblood 1250948, Greater Cleric's 1250926, Elixir of Sages 1250941 and Flask of Distilled Wisdom 17627. All three elixirs share SpellLabel 4519, but so do about 60 Forever alchemy spells, potions included, and no effect or restriction row names label 4519. So the label is a family tag, not exclusivity proof. No client row decides whether these stack; vanilla has no elixir categories. This is round 9's auras C5 question, now with this preset's elixirs named. | `core/consumes.go:619-637` | SpellLabel 188152-188162 (label 4519); no SpellAuraRestrictions row | OPEN (beta buff bar with all three) | if only one elixir applies: -40 healing or -20 mp5 or -25 Spirit/-2% crit, about 3-6% |
| P6 | **Concentration Aura** is still not modelled. Round 9 P2/A5, an owner kit choice. Client 19746 e0 aura 149 bp 35, unchanged. | none | SpellEffect 697291 | OPEN (owner) | about +0.4% |
| P7 | **Oath of Light: scaled by the coefficient or flat.** Unchanged since round 7. Client per-level amounts 6.6667/1.9167/3.8333 match `oath_of_light.go:27-29`; no SpellLevels row, so 400/115/230 at 60. Cost is 8% (PowerCostPct 8), duration index → 300000 ms, GCD 1500. | `oath_of_light.go:19-30, 78-92`; `healing.go:92`; `holy_shock.go:157` | SpellEffect 1367618-20, SpellPower 1328548 | OPEN (unchanged) | flat would be larger |

No CONFIRMED-BUG for the Holy paladin.

## Aligned (checked this round, from the client rows)

**Heals the rotation casts (values at 60):**
- **Holy Light r8, 10329.** bp 1217, Variance 0.10782, 5.2 a level, SpellLevel 54 to MaxLevel 59, coefficient 0.714, ManaCost 580, cast 2500 ms, GCD 1500. Matches `holy_light.go:27`.
- **Holy Light r9 (25292).** 1580 / 0.10714 / 5.8 / 660. Matches the table, though the rotation does not cast it.
- **Flash of Light r1, 19750.** 46 / 0.14925 / 1.0 a level, 20-25, cost 35. **r4, 19941:** 151 / 0.11483 / 1.9, 42-47, cost 90. Both coefficient 0.429, cast 1500. Matches `flash_of_light.go:19,22`. Level 20 is not below 20, so no sub-20 penalty.
- **Holy Shock r1 cast, 1311606.** ManaCost 160; CategoryRecoveryTime 10000, category 892 shared by every rank; GCD 1500; instant; friendly range 40 yd (RangeIndex 578: 20 hostile / 40 friendly).
- **Holy Shock r1 heal, 1311605.** bp 114, Variance 0.07547, no growth, 30-37, coefficient 0.429.
- Matches `holy_shock.go:34-35, 59-61, 109`.

**Divine Favor, 20216:**
- aura 107 misc 7 (crit chance) +100, mask 0x80206000. That covers Holy Light (0x80000000 and 0x4000), Flash of Light (0x2000) and Holy Shock (0x200000), which is on both the cast and heal 1311605. The engine's PaladinSpellMaskDivineFavor equals that set.
- ProcCharges 1, so the first matching spell spends it (`divine_favor.go:49-58`).
- Cost: PowerCostPct 4 (base mana).
- RecoveryTime 120000 with StartRecovery 0, so it is off the GCD; the engine has no GCD on it.
- Attributes_0 carries 0x2000000, as Nature's Swiftness, Inner Focus and Presence of Mind do. Those are the "cooldown starts when spent" spells, so the engine starting the cooldown on spend matches (`:27-31`).
- Duration index 21 (-1, until used) matches NeverExpires.

**Talents (trait curves at the taken rank):**
- **Spiritual Focus 5:** curve 82976, 70; aura 108 misc 9; mask 0xC0000000 / mask3 0x04000000 covers Holy Light, Flash of Light and Light's Vigil. Engine 0.14 × 5 on that mask.
- **Healing Light 3:** curve 82974, 12; aura 108 misc 0; mask 0x6000 / mask1 0x10000 covers Holy Light, Flash of Light and the Holy Shock heal. The engine multiplies by 1.12 on the same three.
- **Reverence 2:** curve 110321, 30; aura 134. The engine adds 0.30 to SpiritRegenRateCasting.
- **Illumination 5:**
  - Proc chance 100 (e0 curve 82971).
  - Refund 50% of base cost (e2, fixed; the text reads e2, because e1's 10-50 curve would read 10% at rank 1, against the text's 50%).
  - 20272 e1 aura 134 at 50 for 15 s (DurationIndex → 15000).
  - The engine has 1.0 / 0.5 / +0.5 casting regen for 15 s, and the refund is from flat base cost, before Benediction.
- **Infusion of Light 2:** curve 110323, -1000. Buff 437063: aura 107 misc 10, flat -1000 ms, mask 0x80000000 (Holy Light), 1 charge, 15 s. The engine has -500 × 2 on Holy Light for 15 s and spends it on the Holy Light cast. Triggers on Flash of Light and Holy Shock heal crits (text).
- **Boundless Light 2:** curve 127943, 30. The engine gives +30 on the Holy Shock heal and removes it on the sworn ally.
- **Divine Intellect 5:** curve 82973, 10; aura 137 with stat mask 8. The engine multiplies Intellect by 1.10; aura-137 percents multiply.
- **Benediction 5:** curve 82990, -10; misc 14 (cost). Mask 0x1BE11FBF... includes Holy Shock's 0x200000 but not Holy Light (0x80004000 / mask2 0x400) or Flash of Light (0x40002000). Divine Favor and Oath have no mask. The engine applies 90% to Holy Shock only.
- **Conviction 3:** curve 82985, 6; aura 290. The engine adds 6% to stats.Crit, every table.
- **Divine Precision 1, Precision 3, Purifying Power 2, Guardian's Favor 2, Sacred Duty 2:** no heal or cost effect on any cast spell. Divine Precision is aura 107 misc 16, hit on Holy Shock; a heal does not miss.
- **Light's Vigil:** registered, not in the list.

**Every client modifier naming a cast spell** (SpellClassSet 10) is either applied above or comes from a source the build does not have:
- **Items not worn:** 1220645 (item 236271), 1306429 (279248), 459602 (227152/231643), 436284 (231328, set 1592), 1291089 (272436), 1293544 (273329...), 1214058 (234000), 468913 (231811).
- **Sets not worn:** T1 5P 1301694 (set 2107), TAQ 1213349 (1843), T2 467506 (1809), Improved Holy Light 24457 (set 475).
- **Not in Forever's tree:** Sanctified Light 1224553 (trait 134817, not in the Forever tree).
- **Old rows, no class mask:** Improved Flash of Light 20249-20251 (SkillLineAbility ClassMask 0, like Divine Favor's own pre-tree row; not in trainables).
- **No source at all:** Revelation, Emergency, The Quick and the Dead, Shielded in Righteousness, Sacred Shield, Infusion of Light 53672, The Art of War. Noxious Breath and Corrupted Mind are creature spells.
- No SpellLabel modifier (aura 507/508) names label 26, the label these spells share.

**Other casts:**
- **Lay on Hands 10310:** CategoryRecoveryTime 1200000 (20 min), GCD 1500, effect 67 (max health) and effect 30 at 550 mana. The Wowhead Forever tooltip reads "Drains all of the Paladin's remaining Mana ... does not interrupt Mana regeneration". The engine drains, heals to max, then returns 550 (`lay_on_hands.go:54-60`).
- **Oath of Light:** see P7 for amounts and cost.

**Crit rules after the round-9 split:**
- Every paladin heal is DefenseTypeMagic and rolls spell-table crit (`healing.go:67`, `holy_shock.go:142`), with a 1.5x crit (`spell_outcome.go:1083`).
- Crit procs (Illumination, Infusion) read `result.DidCrit()` on the heal.
- Intellect → SpellCrit at 0.0167% a point at 60 (`base_stats_levels_auto_gen.go:621`, from the client's playerexpectedstat SpellCritPerIntellect).
- Base spell crit 3.5% is Era's value; the engine already says no client table exists.
- Evergreen Shield's item crit (stat 32) reaches every table.

**Gear:**
- **Devout set 182:**
  - 2P 18679, resistances.
  - 3P 9346, aura 13 and 135 at 18.
  - 4P 27778, struck; no sim effect.
  - 5P 450576: 5% on cast, 450554 energize 200. The Holy Shock heal has NoOnCastComplete, so a Holy Shock procs once.
  - 6P 18378, aura 85 at 8.
  - All in `core/client_sets_gen.go:24` and `common/item_sets/client_dungeon_sets.go:64-68`.
- **Uther's Strength:** 8397 ptm 0x28 (taken melee and ranged), 4%, gives 10368, a 216 absorb. The healer is never struck in melee, so it stays idle.
- **Hand of Edward the Odd:** 18803 Focus is a melee-hit proc, idle.
- The other worn items carry no ItemEffect rows in the client. Their stats belong to the items lane.

**Race:**
- Human: Ruthless 20597 is aura 290 at 1, on stats.Crit (`racials.go:216-223`). The Human Spirit 20598 is aura 137 at 10 with stat mask 16 (Spirit). Round 9's text mismatch (P5 there) is unchanged.
- Undead: no combat racial for a healer.

**Raid preset:**
- Brilliant Mana Oil 25116: aura 85 at 15 and aura 135 at 30. Matches `consumes.go:136`.
- Greater Cleric's Elixir 1250926: aura 135 at 40.
- Greater Mageblood Elixir 1250948: aura 85 at 20.
- Elixir of Sages 1250941: aura 29 at 25 Spirit and aura 290 at 2.
- Flask of Distilled Wisdom 17627: aura 418 at 2000.
- Major Mana Potion 17531: 1800 with variance 0.5, which is 1350-2250.
- Demonic Rune 16666: 1200 ± 0.5 mana (900-1500) and 800 ± 0.5 health. It sits on the conjured timer, apart from potions.
- Moonkin Aura and Leader of the Pack: aura 290 at 3, so stats.Crit is right.
- Judgement of the Crusader, Judgement of Light and Judgement of Wisdom are enemy debuffs and change no paladin heal.
