# Sim audit round 9: healers

Lane: wt/r9-healers (site, from origin/main 55a1212f7) and lane/r9-healers (fork, from 40021374a). Specs: Restoration druid, Holy paladin, Discipline priest, Holy priest, Restoration shaman, the heal profile and the effective-HPS metric.

Client truth: 1.60.1.70707 layered with the hotfix push (raw CSVs fetched 2026-11-01 21:42 UTC, data/builds/1.60.1.70707/spellconst). Four delegated engine-vs-client sweeps (one a class) re-read every heal rank by script and every taken talent's class mask; I re-read each reported bug against the client row and the code before confirming it. Their full findings are in [healers/](healers/) (aud-druid.md, aud-paladin.md, aud-priest.md, aud-shaman.md), and the downrank table is [healers/downrank-candidates.md](healers/downrank-candidates.md).

The short version:
- **Twin Disciplines reaches Prayer of Mending's heal after all.** Round 8 removed it after reading only the cast and the jump. The heal is three other spells whose one class-mask bit Twin Disciplines names. Holy +0.7%.
- **Restoration druid spent 4 points on Improved Wrath, which heals for nothing.** They go to Heart of the Wild. +1.4% on both factions, legal by heal-builds' tree check, clears the adoption rule.
- **The addon told players to cast ranks they cannot learn**: Ahn'Qiraj book ranks at 60 (Healing Wave 10, Holy Light 9, Heroic Strike 9), a Burning Crusade Lesser Healing Wave, Water Shield's proc id, and the top rank where the band-60 list deliberately downranks.
- **Small engine fixes:** the priest racial heals had no class mask, so no talent reached them; Nature's Swiftness covered Wild Growth.
- **Downranking (ruling 3):** measured at every band. Below 60 a bare healer runs dry 30 to 50 s into the level-60 profile, so lower ranks win big (Holy paladin up to +59%, shaman +20%, Discipline +15%). The deliberate Renew rank 3 also wins at bands 30 and 40. Handed to the ranks lane; no rank edited here.
- **Settled from the client:** every Ahn'Qiraj book rank (each has a client item that teaches it), Mana Tide's 290 a tick, Chain Heal's jumps.

## Assigned open rows (round 8)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Mana Tide's per-tick mana.** Rank 4's totem (1325256) summons creature 277993 (EffectMiscValue_0; it is a creature id, not a spell). The restore spell is 1325257: aura 24 (periodic energize), period 3000, base 290, level 58, same name and icon as the totem. Ranks 1-3: 16191/17355/17360 at 88/133/197. The Forever tooltip of 1325256 says "restores 290 mana every 3 seconds ... for 12 sec". The engine (`sim/shaman/mana_tide_totem.go`) reads ManaTideBaseDamage and ticks 4 times every 3 s. | SpellEffect 1362474, 1362475; SpellLevels 154055/154056; https://nether.wowhead.com/forever/tooltip/spell/1325256 (read 2026-11-02). | Aligned (client row + server tooltip) | - |
| A2 | **Healing Wave rank 10 (25357) and every other Ahn'Qiraj book rank.** The client carries the books and links each to its rank with a learn-on-use effect (ItemEffect TriggerType 6): Tablet of Healing Wave X 21291 → 25357, Libram: Holy Light IX 21290 → 25292, Libram: Blessing of Wisdom VI 21288 → 25290, Codex of Greater Heal V 21284 → 25314, Codex of Renew X 21285 → 25315, Codex of Prayer of Healing V 21287 → 25316, Book of Healing Touch XI 21294 → 25297, Book of Rejuvenation XI 21296 → 25299, Manual of Heroic Strike IX 21297 → 25286. None is on Phase 1 loot. This settles the paladin sweep's P4 and round 7's open row the same way. | ItemSparse, ItemXItemEffect (e.g. 154158 → effect 223194), ItemEffect 223193/223194. | Aligned (client items + vanilla) | - |
| A3 | **Chain Heal's party-only jumps.** The client effect is ImplicitTarget 45 (chain-heal ally), 3 targets, 0.5 falloff. cmangos-classic `Spell.cpp` (TARGET_UNIT_FRIEND_CHAIN_HEAL) limits jumps to the caster's GROUP when the target is in it, which in a raid is the raid, not the party. The fake raid is one group, so the engine's "any raid member" is right. One small difference: cmangos orders jumps by missing health, the engine by health share; the members share one health total, so only the healer's own row can sort differently. | `sim/shaman/chain_heal.go:65-91`; https://github.com/cmangos/mangos-classic/blob/master/src/game/Spells/Spell.cpp (read 2026-11-02). | Aligned (one server source) | - |
| A4 | **Power in Light** (1309969) is a priest Discipline talent: "Smite and Penance deal 10% increased damage", one aura-4 dummy effect. Damage only; no healer list casts Smite or Penance's damage. | spellconst/priest.json 1309969. | Aligned (no healing effect) | - |
| A5 | **Concentration Aura** is trainable at 22, aura 149 (resist pushback) 35%. mangos-classic adds it to the caster's "not lose casting time" modifier and rolls once, so Spiritual Focus 70% + 35% means Holy Light and Flash of Light never lose time. Nothing competes for the slot in the kit or the raid preset. About +0.4% at bands 30 to 60. Whether the paladin's kit runs it is a kit choice for the owner. | aud-paladin.md P2; mangos-classic Spell.cpp:6815-6820. | OPEN (owner: kit choice) | - |
| A6 | **forever-bugs #419** (shadow wands ignore Shadow multipliers) is still open with no reply (updated 2026-10-21). No healer list shoots a wand. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/419 (read 2026-11-02). | OPEN (not a healer number) | - |
| A7 | **Swiftmend's amount.** No log at level 40+ exists: foreverlogs beta logs stop near 30, `/v1/reports/recent` holds three October reports with no Swiftmend. Gift of Nature eff0 and Genesis 1223081 eff1 both name Swiftmend; the engine heals the consumed effect's snapshot only. At most about +0.4% for the spec. | aud-druid.md D3. | OPEN (one logged Swiftmend beside the eaten tick) | - |
| A8 | **Discipline's Inspiration hold.** Inspiration's effect is armor on the healed target; the heal profile's damage is after mitigation, so no armor model exists to value it. The adoption rule keeps unmodelled talents, so the hold stays. | `sim/priest` has no Inspiration; heal-profile.json notes. | OPEN (held by rule) | - |

## Downranking (owner ruling 3)

The full table, every single change and the method: [healers/downrank-candidates.md](healers/downrank-candidates.md) (also in the lane scratch at r9/healers/downrank-candidates.md). Confirmed at 10,000 iterations on seeds 1 and 50001, published entry per band, alliance:

| spec | 20 | 30 | 40 | 50 | 60 (raid) |
|---|---|---|---|---|---|
| Holy paladin | none | Flash of Light r1: +22% | Holy Light r2: +57% | Holy Light r3 + Flash of Light r1: +59% | tank Flash of Light r1: +1.1% |
| Restoration shaman | Healing Wave r3: +9% | Healing Wave r3: +14% | Healing Wave r3: +20% | Healing Wave r3: +10% | none |
| Discipline priest | none | Renew r1: +4% | Renew r1 + Binding Heal r1: +15% | Binding Heal r3: +6% | Heal r1: +1.1% |
| Holy priest | none | Renew r3: +1.4% | Renew r3 + party Heal r2: +3% | none | tank Flash Heal r2: +1.4% |
| Restoration druid | Rejuvenation r2: +2% | Rejuvenation r2: +10% | none | none | party Rejuvenation r4: +3% |

- The deliberate band-60 ranks the rule must keep are listed in `sim/request/healer_ranks_test.go` (`chosen`): Holy Renew r3 (6075); Discipline Renew r3 and Power Word: Shield r4 (3747); paladin party Flash of Light r4 (19941) and Holy Shock r1 (1311606). Below 60 the rewrite lifts each to the band's top rank today. For Holy, keeping Renew r3 is itself the best change at bands 30 and 40.
- The profile is the level-60 fight at every band (see X3), so below 60 the answer is the most healing per mana. Ranks learned below 20 here carry the sub-20 penalty and still win.
- Several 60 rows sit at the rule's 1% bar. Holy paladin's and the shaman's band 40-50 rows are not near it.

## Side effects of round 8's fixes, every band

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R8a | **Casts by band** (1,000 iterations, bands 20-50 bare): Innervate first at 40; Lightwell first at 40 (a Holy trainable in Forever, trainables/priest.json 724/27870/27871 at 40/50/60); one Lightwell renew rank a click; Divine Grace rank by band (1277371 at 20, 1277372 at 30, 1277374 at 40, 1277377 at 50) with the sub-20 penalty on rank 2 (level 18); Contingency Plan ward and heal by band; Mana Tide rank 2 at 40, 3 at 50; Water Shield from 30. Holy Shock first at 40 (the band-30 build cannot reach it). | spec-breakdown -heal per band. | Aligned | - |
| R8b | **Shaman regen 17 + Spirit/5, the paladin kit's Wisdom, Devout 3P, stat 45:** bare and raid numbers at every band reproduce round 8's after-figures exactly on the pin (bands table below). | before/after runs. | Aligned | - |
| R8c | **Twin Disciplines and Prayer of Mending** (round 8 P2) was wrong; see P1. | - | see P1 | - |
| R8d | **Desperate Prayer dwarf-only:** no published priest is a dwarf (gnome/troll Holy, human/undead Discipline). | bis files. | Aligned (no impact) | - |

## Crit sources (for the crit lane)

- In this client, Leader of the Pack (24932), Moonkin Aura (24907), Elixir of Sages (1250941), Nature's Majesty, Elune's Light and Thundering Strikes are all aura 290, which is "critical strike chance" of every kind. The Forever tooltips say "Critical Strike chance" with no melee/spell split (24907, 17007, read 2026-11-02). They must stay in heal crit when spell crit is split off. Aura 52 (115 rows) is melee crit and aura 57 (88 rows) spell crit.
- Heal crit today is `Spell.HealingCritChance` = stats.Crit + the spell's BonusCritRating (`sim/core/spell_result.go:253-258`). It leaves out `GetSchoolBonusCritChance`, so school-crit bonuses (aura 71; the hand-written Vestments of Prophecy 5P Holy and The Ten Storms 5P Nature) never reach a heal. No published healer wears either set.
- When the split lands, heals should read spell crit plus aura 290 sources, not aura 52 sources or Agility.

## Priest: Holy and Discipline

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Twin Disciplines raised nothing on Prayer of Mending, but its client mask names the heal.** Round 8 checked the cast (1240827, mask 0/0x20/0x400/0) and the jump (401880). The heal is three separate spells, 401863, 1313379 and 1313381 (Effect 10, mask 0/0/0/0x200000). Twin Disciplines eff0 (1225132, aura 108 op 0, +5%, mask 0x19480020/0x18012/0/0x200000) names that word-3 bit, and no other priest spell carries it. Inner Focus's crit mask names it too. The engine heals through the cast, so the cast takes the modifier. | `sim/priest/priest.go` PriestSpellMaskTwinDisciplines; SpellClassOptions 70399, 84501, 84502; SpellEffect 1214136. Test rewritten first: `TestTwinDisciplinesRaisesPrayerOfMendingsHeal` failed, then passed. Single source (client masks); the one inference is that the server applies a caster modifier to the triggered heal. Holy +0.7%. | CONFIRMED-BUG (round 8 regression) | E 181784d8d |
| P2 | **The racial heals carried no class mask, so no talent reached them.** Divine Grace (1277370-1277378) has mask 0x1000000: named by Holy Specialization eff0, Twin Disciplines eff0 and Spiritual Healing eff0. Contingency Plan's heal over time (1277456-1277460) has mask 0x20: named by Spiritual Healing eff0 and both Twin effects. They now carry engine masks those talents include; Holy Specialization's crit also reaches Divine Grace. | `sim/priest/racials.go`, `talents.go` applyHolySpecialization; SpellClassOptions 1277378 mask_0 16777216, 1277460 mask_0 32. Test first: `TestTheTalentsReachThePriestRacialHeals`. Discipline (human) +0.1%; Holy (gnome) under 0.05%. | CONFIRMED-BUG (small) | E 7990f1de6 |
| P3 | **Discipline's Prayer of Healing note said "Rank 1 is the rank this profile pays for"**, but the line has cast rank 4 (10961) since 2026-10-21 (c9bdef048). A test now holds every healer note that names its rank to the rank the line casts. | Test first: `TestHealerListNotesNameTheRankTheyCast`. | CONFIRMED (prose) | S 678e5833c, E 442dbec71 |
| P4 | **Which talents reach Prayer of Mending's heal.** With P1, the client says modifiers read the heal spell's own mask. Spiritual Healing and Holy Specialization name the cast but not the heal; the engine applies both. If the server folds only the heal spell's modifiers, Holy is up to about 1.5% high. | aud-priest.md H2. | OPEN (a Holy log: PoM crit rate against Flash Heal's) | - |
| P5 | Prayer of Mending's jump goes to the most hurt member; the client gives an area-ally target (22/30), which may be random. Percent cost modifiers are added, not multiplied (about 2% of PoM's cost). Devout 4P Divine Protection is not simulated (2% when struck, 60 s, 488 absorb). | aud-priest.md H4, H5, D2. | OPEN (small) | - |
| P6 | **Aligned:** every heal rank re-read by script (Lesser Heal, Heal, Greater Heal, Flash Heal, Renew, Prayer of Healing, Prayer of Mending, Power Word: Shield, Penance, Binding Heal, Lightwell's renew); Weakened Soul, Soul Warding, Improved PW:S, Renewed Hope, Divine Aegis, Inner Focus, Inner Fire, Dark Sacrifice; the talent curves and masks of Inspiration, Spiritual Guidance, Spiritual Healing, Holy Specialization, Divine Fury, Improved Healing, Improved Renew, Litany of Light, Mental Agility, Mental Strength, Meditation, Silent Resolve, Twilight Focus; Eureka!, Sudden Insight; the gnome, troll, human and undead racials; the Devout bonuses and the trinket uses. | aud-priest.md "Aligned". | Aligned | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Holy Power's crit-heal size**, the largest open paladin item. The engine adds 0.10 a rank to the crit bonus (1.75x at rank 5). The row is aura 50, Holy school, no mask. In the vanilla core aura 50 is block skill, so Forever uses the modern meaning (crit healing amount). TrinityCore would multiply the whole crit (2.25x); the text also fits a flat +50 points (2.0x). +6% or +12% healing per mana if either is right. | aud-paladin.md P1. | OPEN (one logged Flash of Light crit beside a non-crit) | - |
| H2 | A stale engine comment said Oath of Light was unmodelled and every Holy Shock got Boundless Light; the code already takes it away on the sworn ally. | `sim/paladin/healing_talents.go:46-50`. | CONFIRMED (comment only) | E 33afa6882 |
| H3 | Ruthless and The Human Spirit: the spell rows give 1% crit and 10% Spirit (engine), the racial table's text says 2% and 5%. Oath of Light's scaling (coefficient or flat) unchanged from round 7. | aud-paladin.md P3, P5. | OPEN (character sheet; one sworn Flash of Light) | - |
| H4 | **Aligned:** every rank of Holy Light, Flash of Light and Holy Shock (cast and heal) by script; Divine Favor, Illumination, Infusion of Light, Boundless Light, Benediction, Conviction, Divine Intellect, Spiritual Focus, Healing Light, Reverence; Lay on Hands; Blessing of Wisdom ranks; Judgement of Wisdom and Light; human and undead racials; Ankh of Life, Uther's Strength, the Devout set. Light's Grace is a Season of Discovery rune, not in Forever. | aud-paladin.md "Aligned". | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **The addon's rotation named ranks the sim does not cast.** `data/pipeline/addonrotation.py` `_resolve` took the last (level, id) of a family: Ahn'Qiraj book ranks at 60 (Healing Wave 25357, Holy Light 25292, Heroic Strike 25286), Lesser Healing Wave's Burning Crusade copy 27624 over 10468, Water Shield's proc 408511 over the buff 408510. It also lifted the band-60 deliberate downranks (Data.lua shows Holy Shock rank 4, 20930, where the list casts rank 1). It now skips book rows, keeps the named id where two share a level, and keeps the list's rank at 60, as the site does. Data.lua is the nightly's to regenerate. | `addon/ForeverSixty/Data.lua:2063-2078, 1222`; spellranks.json `book: true`. Tests first: three in `data/tests/test_addonrotation.py`; the icon pin now reads the active build. | CONFIRMED-BUG (addon data, all classes) | S a65bd8111 |
| S2 | Water Shield's note said its mana comes on critical heals. Its globes (408510, ProcTypeMask 0x362A8, 3 charges, 3.5 s) also go off when the shaman is hit, which the engine models. | APL note. | CONFIRMED (prose) | S 819342bea, E 33afa6882 |
| S3 | Purification and Healing Way are the same modifier op on Healing Wave; the engine adds them (x1.35), a multiplying server gives x1.375 (+0.5 to 0.7% HPS). Same engine-wide question as the druid's D5. | aud-shaman.md S3. | OPEN (one Healing Wave at 60 with and without one talent) | - |
| S4 | **Aligned:** every rank of Healing Wave, Lesser Healing Wave, Chain Heal, Riptide (with its 25% Chain Heal bonus and tick crits), Healing Stream, Mana Tide; Water Shield; Nature's Swiftness; every taken talent's mask and value; Mana Spring; spirit regen; Devout and trinkets; dwarf and tauren racials. | aud-shaman.md "Aligned". | Aligned | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **The build spent 4 points on Improved Wrath, which only names Wrath** (16814: cast-time and cost modifiers, mask 0x1). Moonglow (mask 0x700007) and Improved Moonfire are damage-only too, but they buy the points that open Nature's Splendor. Improved Wrath 4 and Improved Moonfire's second point go to Heart of the Wild 5 (10% Intellect). heal-builds' tree check passes it. 10,000 iterations, published raid gear: alliance 779.6 → 790.6, horde 777.3 → 788.3 (+1.4%; seed 50001 within 0.1). Keeps every unmodelled talent, Restoration stays deepest, nothing the list casts is dropped. The guide's line that Nature's Focus beats Heart of the Wild no longer applies: the build takes both. Band 50's own better split (Genesis 5 + Heart of the Wild 2, +2.0%) follows the nightly's band derivation. | SpellEffect 695185/1338352 (16814); heal-builds output in the lane scratch. | CONFIRMED (build; adoption rule) | S 9fb3887ff |
| D2 | **Nature's Swiftness covered Wild Growth.** The client's mask (17116, SpellEffect 695239: 0x1000261/0x2000000/0x2000000) names Wrath, Healing Touch and Regrowth among this package's spells; Wild Growth is mask_1 0x4000000. No build casts Wild Growth, so nothing moves. | `sim/druid/druid.go` DruidSpellMaskNatureCasts. Test first: `TestNaturesSwiftnessCoversOnlyTheSpellsItsClientMaskNames`. | CONFIRMED-BUG (no impact) | E 7cdc573c6 |
| D3 | Gift of Nature on Tranquility's ticks (about -1% if not), same-op percent stacking on Rejuvenation (±0.5%), Tauren Endurance's "+2% critical healing" text against its 5% health row (+0.3%). | aud-druid.md D4-D6. | OPEN (one tick each; one tauren crit) | - |
| D4 | **Aligned:** every rank of Healing Touch, Regrowth (direct and 7 ticks), Rejuvenation, Tranquility, Wild Growth (not cast, right for the build); Swiftmend's cost and cooldown; Innervate (400%, learned at 40); Nature's Swiftness; Omen; Bloom; every taken talent's per-rank value (Genesis and Nature's Majesty really are 2% a rank); night elf and tauren racials; Devout set; Ankh of Life. | aud-druid.md "Aligned". | Aligned | - |

## Cross-cutting (metric, profile)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Effective HPS** unchanged from round 8 X4: overheal excluded, absorbs count what they soak, HoTs per tick, self-heals count. | `core/spell_result.go`, `core/shield.go`. | Aligned | - |
| X2 | **Coverage at 1.85 after this lane** (raid, alliance, seed 1; incoming 977.8 a second): shaman 86.1%, druid 80.9% with the new build (79.7% on today's bis talents), Holy priest 70.8%, paladin 67.8%, Discipline 63.4%. Inside the rule's 60 to 85% band but for the shaman at 86. | before/after table. | Aligned; repeat after the re-rank | - |
| X3 | **The heal profile does not scale with the band.** Bands 20 to 50 fight the level-60 profile (2,128 tank hits on 9,500 health). Bare healers cover 3 to 25% and run dry 30 to 50 s in (the druid lasts 200+ s). It is a stated assumption, not a client mismatch, but it makes every sub-60 healer number and rank choice a healing-per-mana figure. | `sim/request/healprofile.go` (no band input); spec-breakdown -heal per band. | OPEN (owner: profile design below 60) | - |

## Before/after

Before: origin/main 55a1212f7 on the fork at 40021374a. After: lane/r9-healers 33afa6882 with wt/r9-healers. Published entries, heal profile at 1.85, effective HPS, 10,000 iterations. The druid's build change shows only with its new talents (the published entries keep the old ones until the nightly runs).

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| Restoration druid (published talents) | 779.6 → 779.6 | 779.2 → 779.2 | 777.3 → 777.3 | 777.5 → 777.5 |
| Restoration druid (new build) | 779.6 → 790.6 (+1.4%) | 779.2 → 791.0 (+1.5%) | 777.3 → 788.3 (+1.4%) | 777.5 → 788.4 (+1.4%) |
| Holy paladin | 662.7 → 662.7 | 663.1 → 663.1 | 648.2 → 648.2 | 648.6 → 648.6 |
| Discipline priest | 619.6 → 620.3 (+0.1%) | 619.5 → 620.3 (+0.1%) | 620.0 → 620.0 | 620.5 → 620.5 |
| Holy priest | 687.0 → 691.8 (+0.7%) | 687.5 → 692.1 (+0.7%) | 674.1 → 678.6 (+0.7%) | 674.2 → 678.8 (+0.7%) |
| Restoration shaman | 841.5 → 841.5 | 841.8 → 841.8 | 838.9 → 838.9 | 838.7 → 838.7 |

Bare bands, alliance, 3,000 iterations, seed 1: Holy priest 40 125.5 → 125.8, 50 187.4 → 188.8 (Prayer of Mending from 40); Discipline 20-50 up 0.1 to 0.5 effective HPS (Divine Grace); every other spec and band unchanged.

## For other lanes

- **Ranks lane:** the downrank candidates and the deliberate ranks to keep: [healers/downrank-candidates.md](healers/downrank-candidates.md). If the rank mechanism changes, `data/pipeline/addonrotation.py` `_resolve` must follow it (it now mirrors today's rewrite: top learned below 60, as written at 60).
- **Crit lane:** the crit sources above. Aura 290 is all crit in this client, and heal crit ignores school crit.
- **Ranker / nightly (held):** the druid's new build (S 9fb3887ff) reaches published entries only when bis.yml runs; band 50's best split differs (D1). Data.lua regenerates from the addon fix (S a65bd8111) and changes the 60 rotation for every class that had a book rank there (Heroic Strike for warriors). The paladin sweep notes band-50 human gear from a PvP rank vendor (Captain Dirgehammer, Knight's Imbued pieces) for the ranker lane.
- **No DPS moves outside this group.** Twin Disciplines' and the racials' changes are heals; Nature's Swiftness' mask touches no cast spell.
- **TestRotationLadder:** Holy and Discipline priest rows move (P1, P2). Not regenerated here.

## Still open

- Holy Power's crit-heal size (H1), the largest open healer item; Prayer of Mending's talents (P4); Oath of Light's scaling; Swiftmend's amount; Gift of Nature on Tranquility; same-op percent stacking (druid, shaman, priest costs). Each needs one level-60 log line; launch is 2026-11-04.
- Concentration Aura in the paladin kit (owner), the heal profile below 60 (owner), Discipline's Inspiration (held).
- Racial texts against their spell rows (Tauren Endurance, dwarf, human). The sub-20 level penalty stays server-core evidence only.
