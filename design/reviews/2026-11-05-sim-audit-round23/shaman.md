# Sim-accuracy round 23: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70790. Site branch wt/r23-shaman (off int/r23 80100bc1d), fork branch lane/r23-shaman (off int/r23, no commit). Gear is the frozen BiS ($R/bis-frozen); raid 60 runs the guide raid build ($R/rt). Build mode: touched tests only, no wow-player gate.

Sources read: round 22 shaman report and integration "Open, for round 23"; guide-accuracy round 25 shaman (origin/main ca06416cd: "Sims requested: none new") and round 24's carried list; Blizzard's Forever notes thread https://us.forums.blizzard.com/en/wow/t/2360696 (JSON read 13:20 UTC 5 Nov: still 10 posts, the last 4 Nov 19:35 UTC, UI only; every shaman line of posts 1, 4, 5, 6 and 9 re-read against the engine).

**One confirmed bug, a card bug in shared pipeline code (S 20d45f26b). No engine or rotation bug.** Raid 60 moves for no spec.

## Seed items (round 22 "Still open")

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Maelstrom Weapon flat 100% against 20 a minute: settles nothing published.** The two models differ only on a weapon faster than 3.0 s. Every band that has the talent wears a 3.6 s or 3.3 s two-hander: band 40 Fiery War Axe 3.60, band 50 Thorium Greatmace 3.60, raid 60 Forest Defender's Axe 3.30 (alliance) and The Unstoppable Force 3.60 (horde). Band 30's Manual Crowd Pummeler (2.00) comes before the talent (37). No log; a web search found no post-3 Nov source on the rate. | Wowhead Forever tooltips nether.wowhead.com/forever/tooltip/item/{279260,19323,250613,870,9449}; $R/bis-frozen/shaman-enhancement.md | OPEN in principle, Aligned for every published band | - | - |
| S2 | **Sub-20 damage penalty: nothing new.** Platform's round 23 S1 searched forever-bugs and the web; a fresh search returned only the SpellCoeffsForever add-on page (Classic Era model) and an epiccarry.com blog (403). Elemental's Lightning Bolt 3 rows at 26-27 and 30-33 still stand only while damage has no sub-20 cut (round 22 E1). | https://www.curseforge.com/wow/addons/spellcoeffsforever ; https://epiccarry.com/blogs/?p=41909 (403) | OPEN-low (a level 25 to 30 shaman's Lightning Bolt 3 tooltip at two spell damage values) | - | - |
| S3 | Logs, 13:21 UTC 5 Nov: api.foreversixty.gg/v1/reports/recent lists only the 20 October report; foreverlogs.gg is a single-page app whose /api/* answers 403 "API access not allowed"; the Warcraft Logs v1 key returns 401 "Invalid key specified" on www, classic, fresh, sod and vanilla. | curl | no logs | - | - |

## Round 22's shared changes as they touch shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | Stormstrike row from 25 on the Enhancement card (both orders take it at 25); trainables and spellranks at 25. Mana Tide Totem from 25 and Nature's Swiftness from 44 on the Restoration card, where the order takes them; Water Shield 20, Riptide 50. | addon-data.json rotations; addon-builds.json orders decoded (`$R/shaman/decode.py`) | Aligned | - | - |
| T2 | **The row that USES Nature's Swiftness printed from 10.** Restoration's "Your highest rank of Healing Wave right after Nature's Swiftness, on the tank under 25%" is gated on `auraIsKnown(16188)`, so it can never fire before 44, but the card showed it from 10 (it was the first row a level 10 to 43 shaman read). Same for Enhancement's "Lightning Bolt at five Maelstrom Weapon stacks" (`auraIsKnown(408498)`), on the card from 10, the talent's first point at 37. Round 22's fix floored a talent's own spell at the build's point, not a line whose condition needs the talent. Fix: `build_rotation` floors a line at the build's point level for every `auraIsKnown`/`spellIsKnown` of a talent spell that is the whole condition or a term of its top-level `and` (a negated or or-ed one gates nothing). Test `test_a_row_gated_on_a_talent_starts_where_the_published_build_takes_it` failed (10, 10), then passed (44, 37). **Moves other classes' cards** (all rows that cannot fire before their talent): Balance Wrath 10 -> 35, Restoration druid Healing Touch after Nature's Swiftness 10 -> 30, Arcane Frostbolt with Presence of Mind 10 -> 31, Retribution's four Unsealed Wrath seal rows to 40, Affliction's Nightfall Shadow Bolt 10 -> 34 (`$R/shaman/cardiff.py`). | `data/pipeline/addonrotation.py` `_talents_allow`; `data/tests/test_addonrotation.py` | CONFIRMED (card) | old (round 22 platform's talent-row fix b4a82d5ad / 6c639ea7f covered cast talent spells only) | S 20d45f26b |
| T3 | 3% raid crit: Elemental and Enhancement raid fights list Leader of the Pack and Moonkin Aura with one 3% (druid lane's sourcing); Healing Stream ticks crit on the every-table crit only (11.1% at raid 60 = Thundering Strikes 5 + raid 3 + item crit). | spec-breakdown raid 60 | Aligned | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | Raid 60 fight re-read: Lightning Bolt 84 casts and 23.6 overload hits (30% of landed bolts, Lightning Overload 3/3), crit 25.4% against Flame Shock's 20.9% (Call of Thunder +5% on the bolt only), Flame Shock 7 ticks a cast (12 + 9 s Call of Flame), Clearcasting 9.4 procs over about 93 damage casts (10% on cast, Forever's text, not on crit), Lightning Mastery 1.48 casts (3 min), Searing Totem's attack every 2.55 s and ProcMaskEmpty (it procs nothing of the shaman's). | `$R/shaman/out/ele60.txt`; spellconst 70790 rows 16246, 16166, 29228, 15208 | Aligned | - | - |
| E2 | Bare 60: rank 8 filler (rank row), 73.8 s out of mana of 180, own Mana Spring dropped once. | `$R/shaman/out/elemental-bare60.txt` | Aligned | - | - |
| E3 | Talent texts against the engine: Elemental Focus 10% on any Fire/Frost/Nature damage cast, 15 s, the procing cast keeps it; Elemental Fury 20% a rank; Call of Thunder 5% Lightning Bolt and Chain Lightning; Lightning Mastery -0.5 s bolt, Chain Lightning off cooldown, 15 s. Elemental Warding 2 and Trial by Fire 2 remain the round 20 tolls. | talents/shaman.json 70790; `sim/shaman/talents.go` | Aligned | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | Raid 60 re-read: Stormstrike 29.9 casts (6 s), Earth Shock 29.5 (6 s less Reverberation 3/5's 0.6 s, gated at a quarter mana, 0 s out of mana), Maelstrom 97 stacks, bolt 17.7 casts, Windfury Totem procs on 20% of landed melee hits (16.3), Elemental Devastation 9% for 10 s at 50% uptime, Flurry 78.8%. Shamanistic Focus is Forever's flat 45% shock cost cut; Mental Dexterity 150% Intellect to attack power; Improved Ghost Wolf unmodeled (kept). | `$R/shaman/out/enh60.txt`; talents 70790 | Aligned | - | - |
| N2 | Guide "Stormstrike is your hardest hit" at 60: Stormstrike 1552 average, white 1007, Earth Shock 821. The near tie at 40 stays the guide lane's wording item. | breakdown | Aligned at 60 | - | - |
| N3 | The Fire Nova row (on Clearcasting, `auraIsKnown(16246)`) still prints from 12: Clearcasting is Elemental Focus's aura, not a talent spell id, so T2's rule cannot see it. The leveling order never takes Elemental Focus; the raid order takes it at 52. The note says "Clearcasting comes from Elemental Focus in the raid build", so the card is honest about it. | addon-data.json; addon-builds.json | wording SHOULD (guide lane / platform) | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Mana Tide Totem casts by rank at 40, 44, 48 and 58 (17354, 17359, 1325256; 4 ticks every 3 s). At bare 25 and 29 no cast: the bare band 20 talents (5320000010000000, taken at the band floor) have no Mana Tide; that is the band design, not a bug (card follows the order). Nature's Swiftness casts at 58, not at bare 44/48 (band 40 talents end before it). | `$R/shaman/out/resto*.txt`; bis-frozen band talents | Aligned | - | - |
| R2 | Water Shield: 3 globes, 3.5 s between globes, any spell crit or hit taken (16 Oct note), cast free and re-cast when empty; Restorative Totems on the whole Healing Stream (16 Oct); Riptide no direct heal, 12 s (3 Nov); Improved Healing Wave 0.1 s a rank; Healing Focus 3 and Ancestral Healing 3 are pushback and armour talents the engine does not model (kept by the adoption rule). | `water_shield.go`, `talents_restoration.go`, spellconst 408521/1239243 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Blizzard notes, every shaman line (posts 1, 4, 5, 6, 9) against the engine: Lightning Bolt ranks, Flametongue/Windfury exclusivity, Grace of Air not exclusive and raid-wide totems, Chain Lightning 2.5 s and cost, shocks at Era values, Lightning Overload 10/20/30, Call of Thunder, Lightning Mastery, Maelstrom doubled, Stormstrike 150% / 6 s / 8% base mana, Mental Dexterity 150%, Riptide, Mana Tide ranks 25/38/48/58, Water Shield on any spell crit, Elemental Focus not from or spent on its own Fire Nova, Clearcasting not spent by overloads. No new post since round 22. | thread 2360696 | Aligned | - | - |
| X2 | Grace of Air 3, Strength of Earth 5 and Healing Wave 10 at 60 stay out (Ahn'Qiraj books, round 11 B2; pending owner "level-60 trainer book ranks"). | trainables/shaman.json | pending owner | - | - |

## Before/after (10,000 iterations, frozen BiS, raid 60, guide raid builds)

No commit moves any number: the one fix is a card row floor. Horde seed 1 on this branch: Elemental 650.5, Enhancement 839.5, Restoration 740.6 HPS, identical to $R/before.tsv.

## Files to regenerate

- data/builds/1.60.1.70790/addon-data.json and addon/ForeverSixty/Data.lua (rotations): the rows in T2 move (shaman, druid Balance and Restoration, mage Arcane, paladin Retribution, warlock Affliction). Nothing else.

## Measure at refresh

- Carried from round 22: Enhancement band 30 with Stormstrike (band talents lack it); Restoration bands 40, 50 and bare 60 Chain Heal rows; Elemental band 20/30 talent strings; Thunderbrew's Boot Flask in Elemental and Enhancement raid 60 BiS.

## For other lanes

- **Platform / integrator (FLAG: touches other classes' cards):** S 20d45f26b changes `data/pipeline/addonrotation.py` for every spec; only card start levels move, no DPS. The six moves are listed in T2.
- Platform: a row gated on a talent's *aura* that is not the talent's spell (Enhancement Clearcasting 16246 from Elemental Focus) still prints early (N3); a small aura-to-talent table would cover it.

## For the guide lane

- N3 (Fire Nova on the leveling card), and round 24/25's carried SHOULDs; nothing new in the guide prose. No guide text changed this round (no rotation, build or rank change).

## Still open

- Maelstrom Weapon flat against 20 a minute (S1): needs a log of a 37+ Enhancement shaman on a weapon faster than 3.0 s; no published band wears one.
- Sub-20 damage penalty (S2), Lightning Overload rank spells, level-60 book ranks: as in the brief.
- Logs last looked 13:21 UTC 5 Nov (S3).
