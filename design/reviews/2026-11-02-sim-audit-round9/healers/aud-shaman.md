# Round 9 healer sweep: Restoration shaman (engine vs client)

Engine: r9/healers/engine at 40021374a. Site: wt r9-healers. Client rows: r8/healers/raw (1.60.1.70707 + hotfix push, fetched 2026-11-01 21:42 UTC). Wowhead Forever tooltips fetched 2026-11-02. Read-only: no sim runs were needed.

Published build (bis/shaman-restoration.json): dwarf (alliance) / tauren (horde). Bands 20 `5510`, 30 `5530003311`, 40 `5530003315510000`, 50 `5530003315513151`, 60 adds Enhancement `055` (Thundering Strikes 5, Ancestral Knowledge 5). Talents taken: Improved Healing Wave 5, Tidal Focus 5, Mindfulness 1-3, Ancestral Healing 3, Healing Focus 3, Water Shield, Tidal Mastery 5, Restorative Totems 5, Mana Tide Totem, Healing Way 3, Nature's Swiftness, Purification 5, Riptide. Totemic Focus, Natural Grace not taken.

## Shaman: Restoration

| # | finding | evidence | class | size estimate |
|---|---|---|---|---|
| S1 | **The addon's level-60 Restoration rotation names three spells a Forever shaman cannot cast.** `data/pipeline/addonrotation.py` `_resolve` takes `learned[-1]` of the family sorted by (level, id) and never reads spellranks' `book` flag. So the 60 band ships Healing Wave 25357 (AQ book rank), Lesser Healing Wave 27624 (a Burning Crusade id) and Water Shield 408511 (the proc spell, not the buff). The level-50 band is correct (10395, 10466, 10622). The sim is not affected: `sim/internal/spellranks` drops book ranks (`RankAvailable`) and resolves a shared tier to its first id (10468). | `addon/ForeverSixty/Data.lua:2063-2078` (commit c3c9affdb, 2026-11-01).<br>- 25357: ItemSparse 21291 "Tablet of Healing Wave X" ("Teaches Healing Wave (Rank 10)"). ItemXItemEffect 154158 points to ItemEffect 223194, TriggerType 6 (learn), SpellID 25357. `data/curated/book-ranks.json` already marks it `book`, and spellranks.json:17427 carries `"book": true`.<br>- 27624: no SkillLineAbility row; only SpellLevels 105043. The Wowhead Forever tooltip reads "Heals a friendly target for 15743". The engine's own comment says it is not learnable (`lesser_healing_wave.go:8-13`).<br>- 408511: Effect 137 (energize 2%), no SkillLineAbility row; 408510 is the castable buff (SLA 49567).<br>- `trainables/shaman.json` also lists 27624 (line 995) and 25357 (line 916) as level-60 ranks. | CONFIRMED-BUG (site/addon data, not engine) | No HPS change. Player-facing: wrong tooltip and an unlearnable spell on the addon's level-60 card. Same generator rule elsewhere: Holy Light 25292 (book) is on the paladin 60 band at `Data.lua:1220-1221`. |
| S2 | The APL note on the Water Shield line still says "kept up for its mana return on critical heals". The client's procs are hits taken (spell, melee, ranged) plus the caster's critical spells. Round 7 S3 fixed only the guide. The note ships to the addon at `Data.lua:2063`. | `data/curated/apl/shaman-restoration.json` priorityList[2].notes.<br>- Spell 408510 description.<br>- SpellAuraOptions 223904, ProcTypeMask 0x362A8: taken melee/ranged/harmful spell, and dealt helpful/harmful spell. | CONFIRMED (prose) | none |
| S3 | **Percent-modifier stacking on Healing Wave.** The client gives both talents on Healing Wave the same op:<br>- Purification: 16178 effect 0, aura 108, misc 0, 10% (curve 85323).<br>- Healing Way: 29206, aura 108, misc 0, mask 64, 25% (curve 85322).<br>The engine adds them (`talents_restoration.go:77-91`, `SpellMod_DamageDone_Flat` into `DamageMultiplierAdditive`): ×1.35. If Forever's server multiplies percent spell modifiers the way the modern server code does, it is ×1.375. This is an engine-wide convention, not shaman-only. | Client rows as quoted. The stacking rule itself is not in any client table. | OPEN. An in-game Healing Wave rank 9 at 60 with 3/3 Healing Way and 5/5 Purification, against the same heal with one talent dropped, settles it. | +1.85% on Healing Wave if multiplicative, roughly +0.5-0.7% HPS. Other classes with two percent mods on one spell move too. |
| S4 | **Chain Heal jump choice.** The engine jumps to the lowest health percent (`chain_heal.go:78-91`). The client gives only ImplicitTarget 45 (chain-heal ally), EffectChainTargets 3 and amplitude 0.5. Server cores (TrinityCore `SearchChainTargets`) pick the largest absolute missing health. In the profile, the tank is always the primary target. So the rule only decides between the healer and the four 400-pulse members, whose maximum health differs. | `SpellEffect` for 1064/10622/10623. The server-core rule is single source, from memory of TrinityCore. | OPEN. In-game: a jump order with unequal-health party members settles it. | Small (under 0.5%). Mostly an overheal split. |
| S5 | Chain Heal's "If cast on a party member, the heal will only jump to other party members" is not modelled (round 7 noted it). The heal profile is one five-member party, so nothing changes. | Spell 1064 description; `chain_heal.go:82` iterates `Env.Raid.AllPlayerUnits`. | OPEN (no effect in the shipped profile) | 0 today. It matters only if the profile grows past one party. |

FOREVER-CHANGE:
- Riptide is not consumed by Chain Heal. Season of Discovery's rune text consumed the heal over time; the Forever client reads "increases the effectiveness of your Chain Heal casts directly on that target by 25%". The aura text is "Increases the effectiveness of the caster's Chain Heal spells cast directly on this target". The engine keeps the heal over time and boosts only the primary target (`chain_heal.go:52-54`, `riptide.go:114-121`). Aligned with the client.
- Racials: ChrRaceRacialAbility 668 ("Dwarf - Passive - Earth Fury") and 685 (Endurance) still say "Critical Damage and Healing increased by 2%". But the spell rows and tooltips say:
  - Endurance 20550: aura 133, 5% health only. Tooltip: "Total Health increased by 5%."
  - The dwarf racial skill line 101 now teaches Artisan's Precision 1259719 (1% hit, aura 54/55), and no Earth Fury spell exists.
  - The engine follows the spell rows (`racials.go:311-318`, `:428-437`). The UI table is stale.
  - Hygiene: the Endurance comment at `racials.go:433-435` mentions an aura-163 crit row that 20550 no longer has.

Hygiene, no effect: `healing_spell.go:83` registers Healing Wave rank 10 (25357) at level 60 whatever `core.IncludeAQ` says. The request rewrite never reaches it, because it aliases the book rank to 10396. A guard on `IncludeAQ` would match `buff_ranks.go`.

## Aligned (what I checked)

- **Healing Wave, ranks 1-10** (331 to 25357). Base, variance, per-level growth and MaxLevel, coefficient (0.429/0.571/0.714/0.857, with the sub-20 penalty on ranks 1-4), cost, 1.5-3.0 s cast, 1.5 s GCD, learn levels. Checked in the conformance golden (`shaman.golden.md:1210-1219`) and by hand for 331, 10396 and 25357. ImplicitTarget 45 with no EffectChainTargets, so a single target.
  - Rank 10 is an AQ book rank (settles round 7's open item): item 21291, its ItemEffect learns 25357, and its SkillLineAbility row 13228 has the same shape as the trainer rows. Only 10396 is cast at 60.
- **Lesser Healing Wave, ranks 1-6.** Rank 6 is 10468 (base 820, 4.2 a level, MaxLevel 65, 380 mana, 1.5 s, 0.429), not 27624. Golden rows 1220-1225, plus 8004 and 10468 by hand.
- **Chain Heal, ranks 1-3** (1064, 10622, 10623). Base 251/335/490, variance, growth 2.5/2.8/3.3, coefficient 0.714, 2.5 s, 260/315/405 mana, learned 40/46/54. EffectChainTargets 3 and amplitude 0.5 match `chainHealJumps=2`, `chainHealJumpFactor=0.5`. Each target rolls its own crit. The falloff scales the spell-power share too.
- **Riptide** (talent; 408521, 1239242, 1239243).
  - Direct heal: 480/612/841 base, growth 3.7/4.2/4.9, coefficient 0.214, variance 0.1, MaxLevel 48/58/68.
  - Heal over time: 89/115/161 a tick, only rank 3 grows (3.7 a level), coefficient 0.1 a tick, 3 s period over 15 s for 5 ticks.
  - Chain Heal bonus: effect 2, aura 4, 25.
  - Cooldown: CategoryRecoveryTime 6000, shared by the ranks. Cost 245/300/385.
  - Every rank carries Attributes_8 0x200 (PERIODIC_CAN_CRIT), so its ticks crit (`periodic_crit_auto_gen.go:324,344,345`). The ticks are dealt as periodic heals, and Water Shield's mask has no periodic flag, so crit ticks rightly do not proc it.
  - Purification effect 1 (misc 22, mask_2 16) reaches the heal over time through `DamageMultiplierAdditive` in `CasterHealingMultiplier`.
- **Healing Stream Totem, ranks 1-5.**
  - Totem rows 5394 to 10463: 40-80 mana, 1 s GCD, 300 s.
  - Heal rows 5672/6371/6372/10460/10461: 5/6/7/9/11 every 2 s, coefficient 0.022, area party aura (effect 35, aura 8), PERIODIC_CAN_CRIT.
  - Restorative Totems effect 1: misc 22, mask 8192, 50% (curve 85319). In the engine it is `DamageMultiplier` 1.5.
  - Purification and Tidal Mastery masks do not include 8192, and the engine applies neither.
- **Mana Tide Totem, ranks 1-4.**
  - 16190/17354/17359/1325256, learned 1/38/48/58, 10/20/30/60 mana, 300 s category cooldown, duration 13000 (the tooltip's "${d-1}" gives 12 s).
  - Restore auras 16191/17355/17360/1325257 (aura 24, period 3000): 88/133/197/290, so 4 ticks. Rank 4's totem summons creature 277993 (EffectMiscValue_0). The engine reads `ManaTideBaseDamage` and gives the party's mana users 4 ticks every 3 s.
  - By band: 40 is rank 2, 50 is rank 3, 60 is rank 4, all through the spellranks chain.
- **Water Shield** (408510).
  - 3 globes (ProcCharges 3), ICD 3500 (ProcCategoryRecovery), 10 min, no cost, GCD 1.5 s.
  - Mana: 408511 effect 137 gives 2% of maximum mana.
  - ProcTypeMask 0x362A8: taken melee swing/ability, ranged, and harmful spell, plus dealt helpful and harmful spells. The engine procs on spell hits taken, raid pulses, crit heals and crit damage spells.
  - Zero stacks deactivates the aura, so the APL recasts it.
- **Nature's Swiftness** (16188). 180 s; COOLDOWN_ON_EVENT (Attributes_0 0x2000000), so the cooldown starts when it is consumed. That matches `talents.go:405-410`. Its mask covers Lightning Bolt, Chain Lightning, Healing Wave, Lesser Healing Wave, Chain Heal and Healing Rain; the engine's rule (Nature school with a cast time) covers the same set for this kit.
- **Talents**, each value read from TraitDefinitionEffectPoints curves:
  - Improved Healing Wave: -100 ms a rank (aura 107, misc 10, mask 64).
  - Tidal Focus: -1% cost a rank (aura 108, misc 14, mask 448 plus Riptide).
  - Mindfulness: 17/33/50 (aura 134).
  - Healing Focus: 23/47/70 pushback (misc 9).
  - Tidal Mastery: 1% crit a rank (misc 7, mask 1472 plus Riptide). The extra bit is Lightning Shield, not Healing Stream.
  - Restorative Totems: Mana Spring 5% a rank (misc 8, mask 16384) and Healing Stream 10% a rank.
  - Healing Way: 8/17/25 (mask 64).
  - Purification: 2% a rank, direct and periodic.
  - Ancestral Knowledge: 2% Intellect a rank (aura 137, misc_1 8).
  - Thundering Strikes: 1% all crit a rank (aura 290). This is the unified Crit stat.
  - Ancestral Healing: armor only (no effect on the sim).
- **Mana Spring.** Ranks 4/6/8/10 every 2 s; Restorative Totems adds 25%. The shaman's own totem adds only its excess over the raid buff. It is not cast by the list, whose water slot is Healing Stream or Mana Tide.
- **Spirit regen:** 17 + Spirit/5 every 2 s (`shaman.go:32-35`, 8.5 + Spirit/10 a second). Mindfulness applies while casting; mp5 always counts.
- **Gear.**
  - Devout set 182 (7 pieces at 60 bare, 6 in the raid set): every bonus 2-6 is active. 9346 is damage and healing 18 (auras 13 and 135). 450576 has a 5% chance, ProcTypeMask 0x14000, for 200 mana. 18378 is 8 mp5.
  - The Devout pieces carry no class restriction in the Wowhead Forever tooltip.
  - Ankh of Life use (1317424, 648-792, 5 min) was modelled in round 7.
  - Uther's Strength: 18 healing; its absorb proc is not a heal.
  - Shard of the Splithooves: 6 mp5, matching the site's item data.
  - Living Crown 73, Glowing Brightwood Staff 150, Mooncloth Boots 24 healing; Staff of Jordan has damage and healing 60.
- **Racials:** dwarf (Stoneform, Find Treasure, Artisan's Precision 1% hit, Big Game Hunter) and tauren (War Stomp, Cultivation, Endurance 5% health, Plainsrunning). None changes healing.
- **Learn levels by band.**
  - 20: Healing Wave 1-4, Lesser Healing Wave 1, Healing Stream 1.
  - 30: Water Shield (talent).
  - 40: Chain Heal 1, Mana Tide 2.
  - 50: Riptide 2, Nature's Swiftness.
  - 60: Healing Wave 9, Lesser Healing Wave 6 (10468), Chain Heal 3, Riptide 3, Mana Tide 4, Healing Stream 5.
  - The request rewrite drops Healing Wave rank 10 (`RankAvailable(book)`).
