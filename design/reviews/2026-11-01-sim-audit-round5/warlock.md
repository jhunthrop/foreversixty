# Sim-accuracy round 5: warlock (2026-11-01)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70707. Pin f8a5f94ce. Lane engine `lane/r5-warlock` (fork), site branch `wt/r5-warlock`.

Method: spec-breakdown, raid preset, band 60, horde troll and alliance gnome, 10000 iterations at seeds 1 and 50001. "Before" is the pin; "after" is the lane engine, both on origin/main's data and the committed simdb.bin. Client rows are the 70707 raw tables (ItemSparse, ItemEffect, ItemXItemEffect, SpellEffect, SpellName, SkillLineAbility) from the r4 scratch copy. Server text is the Wowhead Forever tooltip (`nether.wowhead.com/forever/tooltip/...`, read 2026-11-01). Talent searches use `sim/cmd/talent-search -preset raid -faction horde -keep <raid-keeps.json>` with GOWORK on the lane engine (the reports in `warlock-ts/` print the pin sha in their header).

Logs: still no level-55+ Forever log. Our `/v1/reports/recent` has nothing newer than 2026-10-21, foreverlogs.gg's API answers 403, and a search found no warcraftlogs Forever reports (launch is 2026-11-04).

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Healing gave a hidden third as spell damage on top of the damage each source already states.** `addUniversalStatDependencies` added HealingPower/3 to SpellDamage for every character. Forever's sources state that third as their own damage stat, so it counted twice. Torch of Light (every published warlock ranged slot) gave 6 + 17/3 = 11.7 damage instead of 6. Uther's Strength gave 12 instead of 6. A Devout piece gave twice its damage. A "damage and healing 62" staff (Spire of Hakkar, Kindling Stave) gave 82.7 instead of 62. Healing is unchanged, because the engine's healing reads only the healing stat (e6e054930). The talents that leaned on the rule now state both halves. Spiritual Guidance takes its client damage ranks (1/3/5/6/8% of Spirit, not 8.33%). Champion of the Light gives healers the whole share as spell power (same totals as before). Mental Quickness drops to its stated 30% (was 40%). Atiesh's healing aura states its 19 damage. | (1) Client 70707 ItemSparse: 320 rows carry stat 41 (healing) beside stat 42 (spell damage) at a third of it, e.g. Hide of the Wild 42/14. Only 74 rows carry stat 41 alone, and those are AQ, ZG and Naxx items. (2) Server tooltips: Torch of Light 279246 "healing done by up to 17 and damage done by up to 6"; Devout Bracers "11 and 4"; Hide of the Wild "42 and 14"; Benediction "195 and 65". Rejuvenating Gem, a healing-only row, shows no damage. (3) Talent rows state both halves: 14901 aura 174 = 8 and aura 175 = 25; 30812 aura 174 = 30 and aura 175 = 30. Test `TestHealingPowerGrantsNoHiddenSpellDamage` (failed with 150 against 50, then passed). | CONFIRMED-BUG | E fce7dd53d |
| C2 | Assigned item 3, warlock side: do the round-4 school spell damage stats reach the right spells? Yes. On the round-4 regenerated simdb, +40 Shadow (an elixir, or Anathema's 69) moves only Corruption, Bane of Doom, Bane of Agony and Shadowburn. +40 Fire moves only Incinerate, Immolate and Conflagrate. Per point, each spell moves by its coefficient times its own multipliers. `GetSchoolDamage` (`core/spell_result.go:138`) picks the spell's school, and the demon takes the warlock's larger of Fire or Shadow (`pet.go:281`). The Destruction mismatch came from C1. The ranker's weights priced Spire of Hakkar at 62, while its verify sim gave the Spire 82.7 through the hidden third. After C1, Elemental Mage Staff (+100 Fire) beats the Spire by +1.1% at both seeds (672.5 against 665.3), which is the ranker's own pick. Kindling Stave (60 + 60 healing) is the same case, and it is the staff the ledger names for Arcane. | spec-breakdown `-gear main_hand:<id>` and `-consumes` probes on `r4/casters/newdata/simdb.bin`; scratch final-stat dumps (Spire: SpellDamage 41 to 61 from its 62 healing) | Aligned (school routing); the verify loss is C1 | E fce7dd53d |
| C3 | Forever DoTs read spell power at each tick. The engine snapshots it when the DoT is applied. A bug report says Bane of Agony snapshot "unlike other dot spells". Blizzard replied that Agony gained "scripting to adjust every tick to the spell power of the caster", still snapshotting Amplify Curse. The published warlock sets carry one temporary spell-power effect, Chthonic Power (+29, 15 s, once a fight), so the difference is small. | forever-bugs #398 (2026-10-19 report with combat log; AidanZMoon 2026-10-26) | OPEN: two statements in one thread, and nobody says whether damage multipliers update too. A log with an on-use spell-power trinket popped mid-DoT settles it. | - |
| C4 | Benediction (18608) carries no healing in simdb. Its 70707 ItemSparse row has no stat 41 or 42, yet the server tooltip prints 195 healing and 65 damage. This affects priests only. | ItemSparse 18608; Wowhead Forever tooltip | OPEN, for the hybrids and casters lanes (a hotfix row or the Wowhead supplement could carry it) | - |
| C5 | Champion of the Light's 70707 row (1311084) carries only the damage aura (174), and its tooltip says "spell damage". The engine still gives Holy paladins the healing half from an older tooltip. C1 keeps today's totals. | client SpellEffect 1341129 | OPEN, for the hybrids lane | - |
| C6 | Racials: Troll Berserking and Gnome Eureka! are unchanged since round 4. Set bonuses: the horde Affliction and Demonology sets in round 4 carried Vestments of the Devout. Today's published raid sets carry no set bonus. The only on-use is Frozen Heart of the Mountain: Chthonic Power, +29 Frost and Shadow damage for 15 s on a 5 min cooldown (client 1249110 aura 13, misc 48), modeled as written (`common/item_effects.go:3707`). | bis json; client rows | Aligned | - |

## Pets (assigned item 1)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Demon crit. Issue #331 is still open; its last update was 2026-10-14. #371 (racial crit missing from pet spells) was closed on 2026-10-16 as its duplicate. Nothing new. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/331, /371 | OPEN: a launch log settles it. Demon melee critting near 0.7% means today's behaviour is live; near the warlock's spell crit means Blizzard's change shipped (+1.7 to +2.6%). | - |
| P2 | Demon haste. No source. #437 (channels ignore haste) was closed as not a bug and says nothing about pets. | forever-bugs #437 | OPEN: a demon's swing timer under Berserking in a launch log | - |
| P3 | Demon hit. #394 is still the only source, and it is closed. | forever-bugs #394 | OPEN (single source): a launch log with the demon's melee miss rate against a boss | - |
| P4 | Level-60 demon sheets (base attack power and weapon). No sheet has been published. The BlizzCon demo guides give no demon numbers. | zockify.com/forever/warlock (2026-11-01) | OPEN: a level-60 pet sheet or a launch log | - |
| P5 | Imp Firebolt gaps of 400 ms or more between casts. This is pet spell batching, which Fresh shows too; the reply calls it not a Forever bug. | forever-bugs #435 | Aligned (no Forever change) | - |

## Affliction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Raid search on the lane engine. The best variant that keeps every unmodeled guide talent swaps 1 Improved Drains for Improved Bane of Agony: +4.0 ± 1.3 (+0.7%), below the 1% bar. The winners above it drop Grim Reach, an unmodeled guide talent. The build holds, as in round 4. | `warlock-ts/warlock-affliction-raid.md` | Aligned (held by the adoption rule) | - |
| A2 | Ranks: Corruption 7 (25311), Shadow Bolt 10 (25307) and Immolate 8 (25309) are in the trainer table at level 60, but they are Ahn'Qiraj books (`data/curated/book-ranks.json`), so the APLs rightly use 6, 9 and 7. Corruption 1223963 (137 a tick) has no SkillLineAbility row, so it cannot be learned. | trainables/warlock.json; SkillLineAbility; design/reviews/2026-10-07-aq-ranks-off.md | Aligned | - |
| A3 | Life Tap is 420 + 1 a level + Spirit: 424 at 60. The tooltip's 430 is drawn at level 66. Corruption, Bane of Agony, Siphon Life and Shadow Bolt match the tooltips (342 over 18 s, 552 over 24 s, 41 a tick, 246 to 274 at the drawn level). | Wowhead Forever 11689, 11672, 11713, 18881, 11661; `lifetap.go:28` | Aligned | - |
| A4 | Siphon Life tick crit stays held for the owner, unchanged. | owner ruling | held | - |

## Demonology

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | Raid search: the guide build is the best (636.9 ± 1.4). | `warlock-ts/warlock-demonology-raid.md` | Aligned | - |
| D2 | Demonic Brand (assigned item 2) cannot be modeled from the client. Its damage spells 1293697 and 1293698 carry base points 0 and coefficient 0, and the brand proc 1293696 (aura 42) names no trigger spell. Only the talent text gives the amount: (60 - 26) × 1.5 + 14 to 17 + 0.078 × Shadow spell power for each of the pet's next 2/4/6 attacks, plus less threat on Searing Pain. No raid build takes it, and no rotation casts Searing Pain. A rough count puts a Brand build a few percent up, so it is worth a model follow-up built from the text, with a Searing Pain line. It is a single source until a log shows the proc. | client SpellEffect 1293695 to 1293699; talents/warlock.json 105896 | OPEN (server-scripted; not modeled) | - |
| D3 | Fel Vitality (+15% warlock and pet mana) and Demonic Knowledge match their text. | `talents.go:1125`, round 4 D4 | Aligned | - |

## Destruction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Raid search: the guide build is the best (663.6 ± 1.1). | `warlock-ts/warlock-destruction-raid.md` | Aligned | - |
| S2 | Conflagrate (259 to 321 drawn at 66: 282 + 1.3 a level), Incinerate, Shadowburn, Immolate and Bane of Doom match the tooltips. Doom's "1 min cooldown" equals its duration, so the 60 s gate covers it. Shadow and Flame refunds Shadowburn's shard at 100%, so ignoring shards is right. | Wowhead Forever 18932, 1293813, 18871, 11668, 603, 426316 | Aligned | - |
| S3 | Death Coil (17926: 454 + 3 a level, coefficient 0.214, 600 mana, 2 min) is in no APL. Added before the filler, it loses at both seeds: Affliction -0.5%, Demonology -0.5%, Destruction -0.9%. | spec-breakdown `-rotation` with a Death Coil line; client SpellEffect 17926 | Aligned (correctly left out) | - |

## Before/after (relative, raid preset, band 60, 10000 iterations; pin against E fce7dd53d)

| Spec | Alliance seed 1 | Alliance seed 50001 | Horde seed 1 | Horde seed 50001 |
|---|---|---|---|---|
| Affliction | 577.0 to 559.7 (-3.00%) | -3.00% | 572.4 to 558.3 (-2.46%) | -2.48% |
| Demonology | 664.9 to 637.3 (-4.15%) | -4.16% | 655.7 to 636.9 (-2.87%) | -2.86% |
| Destruction | 688.8 to 668.0 (-3.02%) | -3.02% | 682.2 to 665.3 (-2.48%) | -2.49% |

These are measured on today's published sets, which the ranker chose with healing-hybrid pieces in them: Torch of Light, Uther's Strength, Spire of Hakkar, Kindling Stave and the alliance Demonology Devout pieces. The next bis.yml run re-ranks without the phantom third and will win some of this back.

C1 moves every spec that wears healing gear. Same runs, specs outside this group (seed 1, alliance / horde; seed 50001 within 0.02):
- Arcane -4.64 / -4.69%.
- Balance -4.17 / -3.53%.
- Elemental -3.36 / -3.39%.
- Frost -2.83 / -2.32%.
- Fire -2.75 / -2.26%.
- Shadow -2.16 / -2.15%.
- Enhancement -0.55 / -0.55%.
- Protection paladin threat -0.44 / -0.42%.

Every healer's effective HPS and every other spec are unchanged to the decimal. Fork goldens regenerated in the same commit: TestFeralTank (stats only), TestProtection, TestElemental, TestEnhancement. The rest of the fork suite passes with `-tags=with_db` (`sim/web` fails setup on the pin as well).

## Still open

- Demon crit, haste, hit and the level-60 sheets (P1 to P4): a launch-build log settles each.
- DoT spell power read per tick (C3), Benediction's missing healing (C4) and Champion of the Light's healing half (C5).
- Demonic Brand needs a text-built model and a Searing Pain line before a search can price it (D2).
- Affliction's Improved Bane of Agony swap is +0.7%, below the bar (A1).
- Siphon Life tick crit is held for the owner.
- sim/request TestRotationLadder goldens are not regenerated, as the brief requires. C1 moves them for every class with healing gear.

Sources (read 2026-11-01): https://github.com/ClassicWoWCommunity/forever-bugs/issues/331, /371, /394, /398, /435, /437; https://nether.wowhead.com/forever/tooltip/item/279246 (and 16697, 18510, 18608, 19395, 2816, 13031); https://nether.wowhead.com/forever/tooltip/spell/11689 (and 11672, 11713, 18881, 11661, 18932, 1293813, 18871, 11668, 603, 426316, 17926, 18288, 1293696); https://wago.tools/db2/ItemSparse/csv?build=1.60.1.70707 (and SpellEffect, ItemEffect, ItemXItemEffect, SkillLineAbility); https://www.zockify.com/forever/warlock/.
