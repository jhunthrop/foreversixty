# Healer mana audit (2026-10-08)

Why the two priest specs ran out of mana at 152 to 216 seconds of a 300 second fight, the paladin at about 180 and
the druid at 274, with a Phase 1 raid kit that should not be dry at the halfway mark. The rule the whole healer lane
keeps still holds: **a healing sim ranks gear under a stated incoming-damage profile. It never says one healer class
beats another, and every published healer number carries the profile's name.** Nothing below compares the five
specs; each runs its own curated rotation, talents and gear.

## 1. Verdict

The cause was the engine and the rotation lists, not the mana formulas and not the profile.

1. **A healer never drank its mana potion or its Demonic Rune, and never used Inner Focus.** Fork commit
   `f67083d9c` (a harmful major cooldown needs an enemy to land on) refuses every cooldown that is not flagged helpful
   when the current target is a friend, and a healer's current target is the tank. The potion and the rune are
   self-cast spells with no helpful flag, so the major-cooldown pass (the `autocastOtherCooldowns` line the priests',
   the paladin's and the druid's lists open with) skipped them all fight long. They are worth about 7,200 mana over
   300 seconds, nearly half of the 16,100 the Holy priest took in by regeneration and Litany of Light. The preset's own
   reason said "used on cooldown by the engine"; it never was.
2. **The Restoration Shaman's list had no autocast line at all**, so its potion and rune were never reachable either
   (its notes said so: "mana consumables are the preset's").
3. **Three healer inputs contradicted the client** (section 4): Blessing of Wisdom (a flat 30 at every level, where the
   client states 12 to 36 by rank), the three mana oils (4/8/12 mana, where the client states 5/10/15, with healing
   10/20/30), and Nightfin Soup (8 mana per five seconds, where the client turned it into 22 spell damage and no mana).
4. **Top ranks bought mana the profile did not need.** With the engine fixed the priests and the paladin still ran dry
   at 252 to 273 seconds, because a top rank costs two to three times the mana of a rank that heals nearly as much once
   spell power is added (the client states the same spell power coefficient on every rank of these spells). Section 6
   states the rule that picked ranks.
5. **Not causes:** the spirit regeneration formulas, the five-second rule, mp5, Meditation and its siblings, the
   per-spell costs and the talents that touch mana. All were measured against the client or the stated formula and
   hold (sections 2, 3 and 5).
6. **The profile is not above what a Phase 1 healer sustains.** After the fixes every raid-ready healer's mana lasts
   the fight (the horde Holy paladin by the narrowest margin, 299.7 s of 300). No profile change is proposed
   (section 8).

## 2. Where the mana goes: income and spend, five healers

`go run ./cmd/spec-breakdown -repo-root .. -heal -spec <spec> -bis-dir <ranker out>` (new, in the site worktree)
runs the ranker's published set for a spec under the curated profile and prints the final stats, the mana flowing in
by source and out by spell, and the healing each spell landed. Every table is band 60, the raid preset, alliance,
2,000 iterations, 300 seconds, the profile `onyxia-sized`. **Before** is the site and the fork at `main` / `forever`
(commit `72242529` and `583668528`) with their own ranker sets; **after** is this branch with its own ranker sets.
Gear differs between the two columns because the ranker re-chose it: the weights moved with the mana.

Spend is what the bar paid, so a healer that runs dry stops spending: the "before" columns show what a bar that is
empty by the halfway mark could afford, not what the rotation wanted. For the Holy priest over the first 120 seconds,
before the bar empties, the unfixed list spent 115 mana a second against 55 coming in (plus an 8,300 pool): the bar
was empty by 139 seconds on arithmetic alone, and the mana pacing in the list stretched that to 161.

#### Holy priest (band 60, raid preset, alliance)

| | Before | After |
|---|---|---|
| Mana pool at the pull | 8,332 | 8,418 |
| Intellect | 323 | 328 |
| Spirit | 365 | 278 |
| mp5 from gear and buffs | 83 | 112 |
| + Regeneration inside the five-second rule | 10,543 | 12,273 |
| + Regeneration outside the rule | 3,570 | 856 |
| + Major Mana Potion | - | 3,641 |
| + Demonic Rune | - | 3,593 |
| + Litany of Light | 2,327 | 2,743 |
| **Supply over 300 s (pool plus income)** | **24,772** | **31,524** |
| - Heal r4 | 10,707 (41 casts) | 16,667 (65 casts) |
| - Prayer of Mending r3 | 7,130 (24 casts) | 7,802 (28 casts) |
| - Renew r10 | 4,599 (12 casts) | - |
| - Prayer of Healing r5 | 1,035 (1 casts) | 4,044 (4 casts) |
| - Renew r3 | - | 1,788 (19 casts) |
| - Flash Heal r7 | 791 (2 casts) | 136 (0 casts) |
| - Inner Fire r6 | 315 (1 casts) | 315 (1 casts) |
| **Spend over 300 s** | **24,577** | **30,752** |
| Mana lasts | 161 s | 313 s |
| Effective HPS | 337 | 498 |

#### Discipline priest (band 60, raid preset, alliance)

| | Before | After |
|---|---|---|
| Mana pool at the pull | 7,374 | 7,684 |
| Intellect | 285 | 306 |
| Spirit | 303 | 297 |
| mp5 from gear and buffs | 108 | 113 |
| + Regeneration inside the five-second rule | 9,994 | 12,473 |
| + Regeneration outside the rule | 4,687 | 1,002 |
| + Major Mana Potion | - | 3,598 |
| + Demonic Rune | - | 3,550 |
| **Supply over 300 s (pool plus income)** | **22,055** | **28,308** |
| - Heal r4 | 9,153 (35 casts) | 12,104 (47 casts) |
| - Penance r4 | 5,959 (22 casts) | 5,844 (23 casts) |
| - Prayer of Healing r1 | - | 5,930 (15 casts) |
| - Power Word: Shield r4 | - | 3,150 (24 casts) |
| - Power Word: Shield r10 | 2,964 (8 casts) | - |
| - Prayer of Healing r5 | 2,313 (2 casts) | - |
| - Flash Heal r7 | 1,155 (3 casts) | 69 (0 casts) |
| - Inner Fire r6 | 315 (1 casts) | 315 (1 casts) |
| **Spend over 300 s** | **21,859** | **27,412** |
| Mana lasts | 216 s | 302 s |
| Effective HPS | 426 | 493 |

#### Holy paladin (band 60, raid preset, alliance)

| | Before | After |
|---|---|---|
| Mana pool at the pull | 8,198 | 8,017 |
| Intellect | 331 | 319 |
| Spirit | 176 | 212 |
| mp5 from gear and buffs | 137 | 95 |
| + Regeneration inside the five-second rule | 10,458 | 8,264 |
| + Regeneration outside the rule | 29 | 30 |
| + Major Mana Potion | - | 3,596 |
| + Demonic Rune | - | 3,594 |
| + Illumination | 2,995 | 3,981 |
| **Supply over 300 s (pool plus income)** | **21,680** | **27,482** |
| - Flash of Light r6 | 21,019 (150 casts) | 19,504 (139 casts) |
| - Holy Shock r1 | - | 3,832 (24 casts) |
| - Flash of Light r4 | - | 2,609 (29 casts) |
| - Holy Light r9 | 13 (0 casts) | 654 (1 casts) |
| - Greater Blessing of Light | 260 (1 casts) | 260 (1 casts) |
| - Holy Shock r4 | 252 (1 casts) | - |
| - Divine Favor | 61 (1 casts) | 8 (0 casts) |
| **Spend over 300 s** | **21,605** | **26,867** |
| Mana lasts | 180 s | 315 s |
| Effective HPS | 310 | 424 |

#### Restoration shaman (band 60, raid preset, alliance)

| | Before | After |
|---|---|---|
| Mana pool at the pull | 7,680 | 7,890 |
| Intellect | 296 | 310 |
| Spirit | 247 | 235 |
| mp5 from gear and buffs | 107 | 108 |
| + Regeneration inside the five-second rule | 6,845 | 8,514 |
| + Regeneration outside the rule | 5,868 | 3,276 |
| + Major Mana Potion | - | 3,600 |
| + Demonic Rune | - | 3,086 |
| + Mana Tide Totem | 1,044 | 1 |
| + Water Shield | 1,577 | 2,140 |
| **Supply over 300 s (pool plus income)** | **23,014** | **28,507** |
| - Chain Heal r3 | 5,072 (13 casts) | 14,082 (37 casts) |
| - Healing Wave r10 | 8,648 (15 casts) | 3,409 (6 casts) |
| - Riptide r3 | 7,224 (20 casts) | 7,184 (20 casts) |
| - Healing Stream Totem r5 | 158 (2 casts) | 159 (2 casts) |
| - Lesser Healing Wave r6 | 0 (0 casts) | 2 (0 casts) |
| **Spend over 300 s** | **21,102** | **24,836** |
| Mana lasts | 404 s | 573 s |
| Effective HPS | 373 | 432 |

#### Restoration druid (band 60, raid preset, alliance)

| | Before | After |
|---|---|---|
| Mana pool at the pull | 6,699 | 6,729 |
| Intellect | 249 | 251 |
| Spirit | 368 | 351 |
| mp5 from gear and buffs | 91 | 88 |
| + Regeneration inside the five-second rule | 10,654 | 10,813 |
| + Regeneration outside the rule | 7,815 | 6,588 |
| + Major Mana Potion | - | 3,606 |
| + Demonic Rune | - | 3,595 |
| **Supply over 300 s (pool plus income)** | **25,168** | **31,331** |
| - Rejuvenation r10 | 10,227 (30 casts) | 12,978 (39 casts) |
| - Healing Touch r10 | 7,672 (11 casts) | 8,179 (12 casts) |
| - Regrowth r9 | 5,199 (10 casts) | 5,515 (10 casts) |
| - Tranquility r4 | 1,666 (2 casts) | 2,299 (3 casts) |
| - Swiftmend | 186 (1 casts) | 0 (0 casts) |
| - Innervate | 62 (1 casts) | 62 (1 casts) |
| **Spend over 300 s** | **25,012** | **29,033** |
| Mana lasts | 274 s | 472 s |
| Effective HPS | 435 | 509 |

How to read the Holy priest rows: the pool is the class's base mana at 60 (1,376 in `gametables/basemp.txt`) plus 15
a point of intellect, less the engine's 280 offset, plus the 2,000 of Flask of Distilled Wisdom. Regeneration inside the five-second rule is mp5 plus 50 percent of the spirit regeneration
(Meditation rank 3 in the guide build: measured 0.500 of it), outside the rule it is all of the spirit regeneration
plus mp5. The spirit regeneration is the engine's priest formula, `6.25 + spirit / 8` a second (Classic's
`12.5 + spirit / 4` every two seconds), which reproduces the measured rate to under one percent (section 3). The
fixed Holy priest is inside the rule for 95 percent of the fight, which is why its outside-the-rule row is small; the
unfixed one spent half the fight with an empty bar, which casts nothing and so lets the rule lapse.

## 3. What was checked and held

| Check | Result |
| --- | --- |
| Priest spirit regeneration, `6.25 + spirit/8` a second | Matches the measured outside-the-rule rate to under 0.1 per second in the unfixed Holy and Discipline runs |
| Other healers' spirit regeneration, `7.5 + spirit/10` a second | Matches the measured outside-the-rule rate to within 1 to 6 mana a second of 46 to 64 (the stat table is read before the fight's auras, and the gap has not been traced further) |
| Five-second rule | The engine spends mana when a hardcast completes and starts the rule there (`cast.go`, `mana.go`), which is the rule's definition |
| Share of regeneration while casting | Priest 50 percent (Meditation 3), paladin 30 percent (Reverence 3), shaman 49 percent (Mindfulness 3), druid 89 to 95 percent (Reflection 3 plus Innervate), each as the talent text states |
| Innervate | Regeneration times five (effect 110: +400 percent) and 100 percent continuing while casting (effect 134), as spell 29166 states |
| Per-spell costs | Every spell the five lists cast pays its client cost less its talents (Improved Healing, Mental Agility, Soul Warding, Tidal Focus, Tranquil Spirit): see the cost audit below |
| Mana Spring Totem | Rank 4 restores 10 mana every two seconds (spells 10494, 10497): the engine's 25 mp5 |
| Divine Spirit, Arcane Intellect, Mark of the Wild | 40 spirit (spells 27681, 27841), 31 intellect, 16 to all attributes: the client's amounts (the last two are pinned in `sim/leveling/buff_ranks_client_test.go`) |
| Major Mana Potion, Demonic Rune, Flask of Distilled Wisdom, Mageblood Elixir | 1,350 to 2,250 mana, 900 to 1,500 mana, 2,000 maximum mana, 12 mp5: the client's, unchanged |

Cost audit, the spells each list casts after the fixes (`mana spent / casts` against the client's list cost, in
`spellconst/<class>.json`; the gap is the talent reductions and the free cast Inner Focus gives):

| Spec | Spell (rank) | Client cost | Engine pays | Ratio | Reason for the ratio |
| --- | --- | --- | --- | --- | --- |
| Holy priest | Heal (4) | 305 | 257 | 0.84 | Improved Healing 3 (-15 percent), Inner Focus |
| Holy priest | Prayer of Mending (3) | 390 | 279 | 0.71 | Improved Healing -15 and Mental Agility -10 (additive; it is an instant), Inner Focus free casts |
| Holy priest | Renew (3) | 105 | 94 | 0.90 | Mental Agility 3 (-10 percent) |
| Discipline priest | Power Word: Shield (4) | 175 | 131 | 0.75 | Mental Agility -10 and Soul Warding -15 |
| Discipline priest | Penance (4) | 355 | 251 | 0.71 | Improved Healing -15 and Mental Agility -10 (it is an instant), Inner Focus free casts |
| Holy paladin | Flash of Light (6), Holy Shock (1), Holy Light (9) | 140, 160, 660 | 140, 160, 654 | 1.00 | No cost talent applies |
| Restoration shaman | Chain Heal (3), Riptide (3), Healing Wave (10) | 405, 385, 620 | 385, 367, 588 | 0.95 | Tidal Focus 5 (-5 percent) |
| Restoration druid | Healing Touch (10), Tranquility (4) | 755, 925 | 682, 821 | 0.90 | Tranquil Spirit 5 (-10 percent) |
| Restoration druid | Rejuvenation (10), Regrowth (9) | 335, 525 | 335, 525 | 1.00 | No cost talent applies |

## 4. Where the client contradicted the engine, and what was fixed

| Input | Engine before | Client | Source | Fix |
| --- | --- | --- | --- | --- |
| Blessing of Wisdom | 30 mana per five seconds at every level (33 with the Ahn'Qiraj book) | 12, 18, 24, 30, 36 for the trainer ranks learned at 14, 24, 34, 44 and 54; 40 for the book rank at 60 | Spells 19742, 19850, 19852, 19853, 19854, 25290 (effect 6, aura 24, every five seconds); Greater Blessing 25894, 25918 | `BlessingOfWisdomRanks`, level-aware, pinned to the client by `TestBlessingOfWisdomRanksMatchTheClient` in `sim/leveling`. A level 60 character without the book gets 36 |
| Minor, Lesser, Brilliant Mana Oil | 4/8/12 mana per five seconds; healing 0/0/25 | 5/10/15 mana per five seconds; healing 10/20/30 | Items 20745, 20747, 20748 and their tooltips; enchants 2624, 2625, 2629 apply spells 25114, 25115, 25116 | `manaOilStats` |
| Nightfin Soup | 8 mana per five seconds | 22 spell damage while well fed, no mana | Item 13931, spell 1249513 (aura 227 of 22 into spell 1249520, aura 13) | `nightfinSoupStats`; the healer preset no longer carries the soup (a healer gains nothing from spell damage) |

The client rows for the foods I read are stat meals in the same way (Smoked Sagefish 4 and Sagefish Delight 7 spell
damage, Lobster Stew 28), so no food in the engine's list gives a healer anything. Fixing Nightfin Soup and
Blessing of Wisdom moves other specs, not only healers: every Alliance mana user at level 60 gains 6 mp5 (7.2 with
the improved form the unit tests use), lower bands lose mp5 they never had, and casters that eat Nightfin Soup gain
22 spell damage. Six fork `.results` goldens moved for exactly these reasons and were adopted (section 10).

## 5. The mana-touching talents in the live trees

The Holy and Discipline priest trees, from `data/builds/1.60.1.70009/talents/priest.json`; the other four healers' are
in the second table. "Measured" is the engine against the talent text above.

| Talent | Text (rank 3 or 5 unless noted) | In the engine | Measured / note |
| --- | --- | --- | --- |
| Meditation | 17/33/50 percent of regeneration while casting | Yes, `SpiritRegenRateCasting` | 0.500 of spirit regeneration inside the rule at rank 3 |
| Mental Agility | -3/-7/-10 percent cost of Smite, Holy Fire and instant spells | Yes, on Renew, Prayer of Mending, Power Word: Shield | -10 percent measured on Renew and Power Word: Shield |
| Improved Healing | -5/-10/-15 percent cost of Lesser Heal, Heal, Greater Heal, Penance, Prayer of Mending | Yes | -15 percent measured on Heal |
| Soul Warding | -15 percent Power Word: Shield cost, 4 s shorter cooldown | Yes | Power Word: Shield pays 75 percent with Mental Agility |
| Inner Focus | Next spell free, +25 percent crit | Yes; now used (it was refused, section 1) | Two casts per fight (at the pull and at 180 s) |
| Litany of Light | 5/10 percent of the cost back when a heal follows a different heal | Yes | 2,500 mana per fight in the Holy priest run |
| Mental Strength | +3 to +15 percent intellect | Yes | Intellect in the stat table |
| Spiritual Guidance | Healing from spirit | Yes | Healing, not mana |
| Spirit Tap | Spirit doubles for 15 s after a kill | No: a healing fight has no kill | Shadow tree |
| Searing Light | A free Holy Nova from Holy Fire ticks | No: needs Holy Fire | A healer casts no Holy Fire |
| Blessed Recovery | Heals after a big hit | No: a health effect | Not a mana source |
| Improved Mana Burn, Devouring Contagion, Shadowform | Shadow and Mana Burn costs | Not used by a healer | |
| Shadowfiend (an ability, not a talent) | Summons a fiend for 15 s; the caster receives 5 percent of maximum mana each time it attacks; 5 minute cooldown, learned at level 1 from Shadow Magic (spells 401977 and 401988) | **No**: the generated constants name it, no spell registers it | The largest unmodelled priest mana source; see section 9 |
| Dark Sacrifice (an ability) | Costs 320 health every 3 s for 15 s at rank 5 (learned at 60) and gives that much mana plus the priest's spirit; 10 minute cooldown (spells 1277324 to 1277328) | **No**: constants only | About 1,600 mana plus spirit once a fight, for 1,600 of the priest's own health; see section 9 |

| Healer | Talents that touch mana | In the engine |
| --- | --- | --- |
| Holy paladin | Divine Intellect (+intellect), Reverence (30 percent regeneration while casting at rank 3), Illumination (half a spell's base cost back on a heal crit) | All three; Illumination returned 3,981 mana in the fixed run, the paladin's second source after regeneration |
| Restoration shaman | Tidal Focus (-5 percent heal cost), Mindfulness (50 percent while casting at rank 3), Water Shield, Mana Tide Totem, Restorative Totems (+Mana Spring), Totemic Focus (-totem cost) | All of them; Water Shield returned 2,140 mana in the fixed run and Mana Tide Totem 1,044 in the unfixed one (the fixed shaman does not fall to the 25 percent the list waits for) |
| Restoration druid | Reflection (50 percent while casting at rank 3), Tranquil Spirit (-10 percent Healing Touch and Tranquility), Living Spirit (+spirit), Innervate | All of them; Innervate is cast once, on the list's own line |

Judgement of Wisdom (a raid debuff in the preset) returns mana only to a character that hits the target, which a
healer never does.

## 6. Ranks: the rule and what each rank returned

The client states one spell power coefficient for every rank of Heal (0.857), Greater Heal (0.857), Flash Heal
(0.429), Prayer of Healing (0.286), Flash of Light (0.429), Holy Light (0.714) and Holy Shock (0.429), while the base
heal and the cost both grow with rank. With around 500 healing power the unchanging spell power term is most of a low
rank's heal, so a low rank returns more health per mana than a high one, and a top-rank list spends more than the
profile asks for. The engine applies the client's
numbers to every rank and models no downrank penalty (the client states none), so a lower rank in a list is a
statement about this profile and the client's published ranks, **not a measurement of the game**. To keep from
choosing ranks only because the sim cannot price a penalty, ranks learned below level 20 are never used: Classic cuts
the spell power such a spell receives and the client publishes no table for the cut.

**The rule.** Start from the top ranks. For each ranked spell, with the tank's casts and the party members' casts
counted apart, try every rank learned at level 20 or above in its place, on the ranker's own guarded score (effective
healing per second, scaled by the share of the fight the mana lasts, squared) at band 60 raid gear, 2,000 iterations
or more. Take the move with the largest gain if it raises the score by more than two percent, and repeat. Gear is the
gear the ranker published for the list being tried; when a new list makes the ranker pick different gear the
comparison is repeated on that gear until no spell moves. The site runs a list as written at band 60 and lifts every
ranked spell to the highest learned rank below it (`sim/request/rotation_ranks.go`), so a lowered rank is a band 60
choice only.

First pass, the single moves from the all-top-rank lists on the gear the ranker published for them (rank, score
against the top rank). After the largest move is taken the others are measured again, which is why the Holy priest
ends with one lowered rank though three groups of rows below clear two percent on their own: once Renew is at rank 3
the mana lasts the fight and the rest are worth little. Shaman and druid rows are summarised after the tables.

##### priest-holy: all spells at the top rank: 480 HPS, mana lasts 265 s, score 376

| Spell (casts on) | Rank | Level | HPS | Mana lasts | Score | Against the top rank |
|---|---|---|---|---|---|---|
| Renew (tank) | 9 | 56 | 483 | 270 s | 391 | +4.1% |
| Renew (tank) | 8 | 50 | 488 | 278 s | 421 | +12.0% |
| Renew (tank) | 7 | 44 | 493 | 288 s | 455 | +21.1% |
| Renew (tank) | 6 | 38 | 498 | 300 s | 498 | +32.4% |
| Renew (tank) | 5 | 32 | 501 | 312 s | 501 | +33.4% |
| Renew (tank) | 4 | 26 | 503 | 324 s | 503 | +33.9% |
| Renew (tank) | 3 | 20 | 506 | 345 s | 506 | +34.5% |
| Heal (member) | 4 | 34 | 480 | 265 s | 376 | +0.0% |
| Heal (member) | 3 | 28 | 488 | 274 s | 406 | +8.0% |
| Heal (member) | 2 | 22 | 493 | 282 s | 434 | +15.6% |
| Prayer of Healing (member) | 4 | 60 | 479 | 265 s | 374 | -0.5% |
| Prayer of Healing (member) | 3 | 50 | 486 | 275 s | 410 | +9.1% |
| Prayer of Healing (member) | 2 | 40 | 492 | 285 s | 444 | +18.1% |
| Prayer of Healing (member) | 1 | 30 | 498 | 296 s | 484 | +28.8% |

##### priest-discipline: all spells at the top rank: 487 HPS, mana lasts 273 s, score 404

| Spell (casts on) | Rank | Level | HPS | Mana lasts | Score | Against the top rank |
|---|---|---|---|---|---|---|
| Prayer of Healing (member) | 4 | 60 | 483 | 274 s | 404 | +0.1% |
| Prayer of Healing (member) | 3 | 50 | 484 | 280 s | 420 | +4.2% |
| Prayer of Healing (member) | 2 | 40 | 483 | 282 s | 429 | +6.2% |
| Prayer of Healing (member) | 1 | 30 | 485 | 288 s | 446 | +10.4% |
| Heal (member) | 4 | 34 | 487 | 273 s | 404 | +0.0% |
| Heal (member) | 3 | 28 | 487 | 273 s | 403 | -0.3% |
| Heal (member) | 2 | 22 | 485 | 273 s | 402 | -0.3% |
| Power Word: Shield (tank) | 10 | 60 | 487 | 273 s | 404 | +0.0% |
| Power Word: Shield (tank) | 9 | 54 | 489 | 273 s | 404 | +0.1% |
| Power Word: Shield (tank) | 8 | 48 | 490 | 273 s | 405 | +0.3% |
| Power Word: Shield (tank) | 7 | 42 | 490 | 273 s | 407 | +0.8% |
| Power Word: Shield (tank) | 6 | 36 | 490 | 274 s | 410 | +1.5% |
| Power Word: Shield (tank) | 5 | 30 | 490 | 276 s | 415 | +2.9% |
| Power Word: Shield (tank) | 4 | 24 | 492 | 277 s | 419 | +3.9% |

##### paladin-holy: all spells at the top rank: 420 HPS, mana lasts 252 s, score 297

| Spell (casts on) | Rank | Level | HPS | Mana lasts | Score | Against the top rank |
|---|---|---|---|---|---|---|
| Flash of Light (tank) | 6 | 58 | 420 | 252 s | 297 | +0.0% |
| Flash of Light (tank) | 5 | 50 | 412 | 268 s | 330 | +11.1% |
| Flash of Light (tank) | 4 | 42 | 413 | 276 s | 350 | +17.8% |
| Flash of Light (tank) | 3 | 34 | 416 | 274 s | 348 | +17.2% |
| Flash of Light (tank) | 2 | 26 | 424 | 275 s | 356 | +20.0% |
| Flash of Light (tank) | 1 | 20 | 430 | 276 s | 364 | +22.7% |
| Flash of Light (member) | 6 | 58 | 420 | 252 s | 297 | +0.0% |
| Flash of Light (member) | 5 | 50 | 418 | 257 s | 308 | +3.7% |
| Flash of Light (member) | 4 | 42 | 416 | 262 s | 318 | +7.0% |
| Flash of Light (member) | 3 | 34 | 416 | 266 s | 326 | +10.0% |
| Flash of Light (member) | 2 | 26 | 417 | 270 s | 337 | +13.4% |
| Flash of Light (member) | 1 | 20 | 417 | 273 s | 346 | +16.5% |
| Holy Shock (member) | 4 | 56 | 420 | 252 s | 297 | +0.0% |
| Holy Shock (member) | 3 | 48 | 418 | 257 s | 308 | +3.8% |
| Holy Shock (member) | 2 | 40 | 417 | 267 s | 331 | +11.5% |
| Holy Shock (member) | 1 | 30 | 420 | 305 s | 420 | +41.6% |
| Holy Light (tank) | 8 | 54 | 419 | 253 s | 297 | -0.0% |
| Holy Light (tank) | 7 | 46 | 418 | 253 s | 297 | +0.2% |
| Holy Light (tank) | 6 | 38 | 419 | 253 s | 299 | +0.6% |
| Holy Light (tank) | 5 | 30 | 420 | 254 s | 300 | +1.1% |
| Holy Light (tank) | 4 | 22 | 419 | 254 s | 300 | +1.1% |

Shaman and Druid: no rank clears two percent. The shaman's best is Chain Heal on the tank at rank 1 (+1.0 percent) and
the druid's is within 0.7 percent of the top rank everywhere; the shaman does not run short of mana and the druid
barely, so no rank answers a mana question there.

Result of the rule, with the gear it was repeated on:

| Spec | Lowered ranks | Why the rest keep their top rank |
| --- | --- | --- |
| Holy priest | Renew on the tank 3 | On the second pass Prayer of Healing rank 1 is +1.2 percent, Heal rank 2 +0.5 percent and Heal rank 3 +0.7 percent: under two percent |
| Discipline priest | Prayer of Healing 1, Power Word: Shield 4 | On the second pass Heal rank 2 is +0.5 percent and Heal rank 3 +0.2 percent |
| Holy paladin | Holy Shock 1, Flash of Light on the four party members 4 | Second pass: with Holy Shock at rank 1 the ranker re-chose gear that left the bar dry at 265 s, and the party members' Flash of Light at rank 4 moved the guarded score from 349 to 440 |
| Restoration shaman, Restoration druid | none | |

The 2026-10-07 review said "no downrank penalty is modelled" and left the rank question open; the Holy paladin's
notes had refused downranking for the same honesty reason. The level 20 floor and the two percent bar are this
audit's answer to it, and they are stated in each list's notes.

## 7. Before and after (band 60, raid preset, both sides, 1,000 iterations per set)

"Fixes only" is the engine and preset fixes with the old top-rank lists (the shaman's with its new autocast line):
what the potion, the rune, Inner Focus and the client's consumable values are worth without any rank change.

| Spec | Side and race | Effective HPS before / fixes only / after | Mana lasts (s) before / fixes only / after | Overheal before / after | Healing per mana before / after |
|---|---|---|---|---|---|
| priest-holy | alliance gnome | 337 / 480 / 499 | 161 / 266 / 313 | 1.6% / 5.0% | 4.11 / 4.86 |
| priest-holy | horde troll | 328 / 464 / 501 | 152 / 255 / 325 | 1.6% / 5.2% | 4.08 / 4.94 |
| priest-discipline | alliance human | 426 / 487 / 493 | 216 / 273 / 302 | 4.9% / 8.7% | 5.84 / 5.39 |
| priest-discipline | horde undead | 423 / 484 / 499 | 214 / 271 / 317 | 4.9% / 9.7% | 5.80 / 5.65 |
| paladin-holy | alliance human | 309 / 419 / 424 | 180 / 252 / 316 | 0.0% / 0.1% | 4.30 / 4.73 |
| paladin-holy | horde undead | 318 / 431 / 434 | 179 / 252 / 300 | 0.0% / 0.1% | 4.45 / 4.94 |
| shaman-restoration | alliance dwarf | 373 / 432 / 432 | 405 / 570 / 570 | 1.4% / 1.8% | 5.31 / 5.22 |
| shaman-restoration | horde tauren | 373 / 432 / 432 | 405 / 569 / 569 | 1.4% / 1.8% | 5.31 / 5.22 |
| druid-restoration | alliance night-elf | 435 / 509 / 509 | 274 / 470 / 470 | 4.1% / 5.5% | 5.22 / 5.26 |
| druid-restoration | horde tauren | 438 / 509 / 509 | 274 / 423 / 423 | 4.1% / 5.2% | 5.32 / 5.14 |

The bare preset (the character and its class kit, no consumables) still runs dry: the Holy priest at 103 s (86 s
before) and the Holy paladin at 90 s (89 s before). That is the profile working as designed, one healer's share of a Phase 1
boss fight asked of a character with no flask, elixir, oil, potion or rune; it is not a finding.

Effective healing per second now sits at 80 to 96 percent of the profile's incoming damage (the profile asks for
about 528 a second: a 1,150 tank hit every 3.3 s, 350 a second, and 300 on three members every 5 s, 180 a second), and
the priests' raw healing is at or just above it. Overheal rose to 5 to 10 percent for the priests as a result. The
guard (mana lasts, squared) now binds only the horde Holy paladin, by a fraction of a percent (299.7 s of 300 in the
ranker's run, 302 s in a 300-iteration breakdown of the same set).

## 8. The profile

`data/curated/heal-profile.json` is **not** above what a Phase 1 healer sustains, and this audit proposes no change
to it. After the fixes every raid-ready set's mana lasts the 300 seconds, within noise for two of the five (the
alliance Discipline priest at 302 s, the horde Holy paladin at 299.7 s), and effective healing sits just under the
profile's damage. The halving that produced the present figures was made against an engine that never drank a potion;
with the engine fixed, the halved figure is what a raid-ready healer just covers.

One thing to watch, not to change: once a healer's effective healing reaches the profile's damage, extra healing power
only overheals and the ranking loses its signal. Phase 2 gear will get there. When it does, raise the tank hit or the
pulses together, state the new figures in the file's reasons, and re-run; do not retune the lists to the profile.

## 9. Not changed, and what remains

- **Non-helpful use effects stay refused for a healer.** The friendly-target gate (`shouldActivateHelper`) still
  skips racial buffs (Berserking, Blood Fury) and stat trinkets, which are damage-class cooldowns. No healer set picks
  an on-use trinket the engine models (Serenity Field has no use effect in the engine), so nothing is lost there; a
  Troll Holy priest's Berserking is not cast either (it casts only Inner Focus, the potion and the rune of its
  cooldowns). A fix would flag self-buff cooldowns helpful at their constructors
  (`RegisterTemporaryStatsOnUseCD`, `core/racials.go`) and needs the DPS goldens reviewed.
- **Forever consumables the engine does not list**, each a protobuf addition and an engine pin: Greater Mageblood
  Elixir (20 mana per five seconds, item 250341), Sage's Tea (44 healing power well fed, 249870), Greater Cleric's
  Elixir (40 healing, 250333), Major Mender's Potion (88 healing, 250949). The 2026-10-07 review's assumption 4 ("no
  healing-power elixir exists in the client's item table") is wrong. The healer preset has no food as a result.
- **Wizard oils** give healing as well as spell damage in the client (Brilliant Wizard Oil: 36 of each); the engine
  gives them spell power only, and adding healing would move every caster through the one-third healing rule.
- **Spirit regeneration is unconfirmed.** The client's gametables carry no regeneration table, so the engine's
  Classic formulas stand; the measurements in section 3 show the engine implements them, not that Forever uses them.
- **Mana Spring Totem and Blessing of Wisdom do not stack** in the engine (`else if` in `buffs.go`), and both are
  periodic energize auras in the client, so they may stack in the game. The raid preset gives each side one of the two
  (the faction gates), so nothing here depends on it.
- **Two Forever priest mana abilities are not in the engine: Shadowfiend and Dark Sacrifice** (section 5). Shadowfiend
  gives 5 percent of maximum mana per fiend attack for 15 s every 5 minutes; the fiend's attack speed and hit chance
  are creature data the client tables here do not carry, so a model would have to invent them. At a pool of 7,700, each
  attack that lands is worth 385 mana, and ten of them would be worth about 3,850, a little over half of what the
  potion and the rune bring together. Dark Sacrifice trades 1,600 health for 1,600 mana plus spirit, once per fight. Both would raise the
  priests' supply, not lower it. Neither was added: the fiend's attack rate has no source here, and Dark Sacrifice
  spends the priest's own health, which the profile does not model for the healer. Whether a healer may use them in
  the profile's fight is a ruling for the owner.
- **The shaman's healing is bound by its health thresholds**, not its bar (it ends with 570 s of mana); a retune of the
  list against this profile would raise its effective healing by a few percent and was not done.
- Superseded in the 2026-10-07 review: section 4's mana figures and assumptions 4, 7, 11 and 12.

## 10. Verification

- Fork `healer-mana`: `go test --tags=with_db ./sim/priest/... ./sim/paladin/... ./sim/shaman/... ./sim/druid/...
  ./sim/core/ ./sim/conformance/` passes. New tests: each healer's written list drinks its potion and rune
  (`healsim.UnusedManaConsumables`, red before the fix), the Holy priest uses Inner Focus, the Blessing of Wisdom
  ranks, the mana oils and Nightfin Soup follow the client. Six `.results` goldens moved only by the mp5 and spell
  damage above (final stats index 12 and 39) and were adopted: `paladin/protection`, `shaman/elemental`,
  `shaman/enhancement`, `druid/tank`, `warrior/dps_warrior`, `warrior/tank_warrior`; `mage` fails at the base commit too and was left alone.
- Site `healer-mana`: `go test ./...` under `sim/` passes; `make apl-check ENGINE_DIR=<fork worktree>` passes;
  the ladder goldens of the three changed lists were regenerated and show the lowered ranks lifted to the learned rank
  at lower bands.
- `sim/go.mod`'s replace was pointed at the fork worktree locally and not committed; `data/builds/*/bis/*.json` are
  nightly-owned and were not committed (the before and after sets above were ranked into a scratch directory).
