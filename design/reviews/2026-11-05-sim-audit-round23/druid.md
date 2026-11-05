# Sim-accuracy round 23: Druid

## Raid crit sourcing (coordinator seed)

**Exclusivity is STATED by Blizzard, word for word, for both auras. Only one small part is inference:** that a raid with both gets exactly 3%. Both auras are 3% (client and tooltips), so whichever one the server keeps, the raid gets 3%. No live log shows the two auras together yet.

1. **Blizzard's Forever notes.**
   - Thread: "WoW Forever Beta Development Notes – Updated November 2", https://us.forums.blizzard.com/en/wow/t/2360696, post #9 by Kaivax.
   - Created 2026-11-03 03:20:33 UTC (2 Nov, 19:20 Pacific, hence "Updated November 2"). Last edited 2026-11-04 19:35 UTC.
   - The lines, verbatim:
     - Balance section: "Moonkin Form: The critical strike bonus from this talent is again an increase to all types of critical strikes and is again exclusive with Leader of the Pack."
     - Feral section: "Leader of the Pack: The critical strike bonus from this talent is again an increase to all types of critical strikes and is again exclusive with Moonkin Aura."
   - These lines reverse the 16 October post (#6, 2026-10-16 00:15 UTC):
     - "Moonkin Aura now only increases Spell Critical Strike chance, and is Raid wide."
     - "Leader of the Pack now only increases Melee and Ranged Critical Strike chance, and is Raid wide."
   - "Again" refers to the beta builds before 15 October. Community guides from that time describe the two auras as exclusive: see 4.
   - Re-read 5 Nov ~10:55 UTC: 10 posts, the last at 2026-11-04 19:35 UTC (post #10, UI only). Nothing new for druids. Blizzard's "Info & Known Issues" thread (2396179, 4 Nov) has no class lines.
2. **Client 1.60.1.70790 (wago, re-fetched 5 Nov ~10:56 UTC).**
   - Moonkin Aura 24907: SpellEffect 701335. Leader of the Pack 24932: SpellEffect 702727.
   - The two rows are identical: Effect 65 (area aura, raid), EffectAura 290 (crit chance of every kind), EffectBasePointsF 3, radius index 11/11, ImplicitTarget 1.
   - Moonkin Form 24858 effect 2 (700928) triggers 24907. The Leader of the Pack talent 17007 is the passive (695211, aura 4).
   - **Hotfix rows:** wago lists push 113283 for both records, in every region including region 91 (Forever), on builds 70707, 70753, 70755 and 70790. The pushed values equal the client's (290, 65, 3). This is the server confirming all-crit at 3%. It does not encode exclusivity.
   - No client table encodes exclusivity. That is a server spell-group rule, invisible to DB2. Both spells carry SpellLabel 21, but so do 361 other druid spells: it is the druid class label, not a stacking group.
   - SpellAuraOptions and SpellCategories have no row for either spell. SpellMisc 332787 and 333864 have nothing stack-related.
   - `data/builds/1.60.1.70790/hotfix-push.json` has no row for either spell.
3. **Wowhead Forever tooltips (5 Nov).**
   - https://nether.wowhead.com/forever/tooltip/spell/24907 (Moonkin Aura): "Increases Critical Strike chance by 3%."
   - /24932 (Leader of the Pack): buff "Increases Critical Strike chance by 3%."
   - /24858 (Moonkin Form): "all party or raid members within 45 yards have their Critical Strike chance increased by 3%."
   - /17007 (Leader of the Pack talent): "increases the Critical Strike chance of all party or raid members within 45 yards by 3%."
   - None of the four mentions stacking or exclusivity. They confirm all-crit at 3% each. A tooltip and its client row are one source.
4. **Community sources.**
   - **Warcraft Tavern**, Forever druid guide (pre-15 Oct beta state, undated page): "Both Feral and Balance bring exclusive 3% Crit buffs via Leader of the Pack or Moonkin Aura." https://www.warcrafttavern.com/forever/guides/druid/
   - **Conquest Capped**, raid buffs guide ("Updated 02 Nov 2026", the day before the notes): "The game files no longer mark them as exclusive, so a raid with both Druids gets both." https://conquestcapped.com/guides/wow-forever/wow-forever-raid-buffs/
     - This describes the 16 October split. Blizzard's later 3 Nov post supersedes it.
     - Its "game files" claim cannot be checked: no DB2 table carries exclusivity either way (2 above).
   - **Not found:** a Forever player, forever-bugs issue or log showing the two auras together.
     - ClassicWoWCommunity/forever-bugs: one Leader of the Pack issue (#275, form-change flicker, closed), none on stacking.
     - Searched: r/wowforever and the Blizzard druid threads ("Druid is a hot mess" 2397126, "Resto Druid Completely Gutted" 2396851).
5. **Live logs (5 Nov ~10:58 UTC).** None anywhere:
   - api.foreversixty.gg/v1/reports/recent lists only the 20 Oct report.
   - foreverlogs.gg answers 403 "API access not allowed".
   - The Warcraft Logs v1 key returns 401 "Invalid key specified" on classic, fresh, vanilla and www.
   - **To settle it:** a raid log with a feral and a moonkin in range of one player, showing that player's buffs (one of 24907/24932, not both), or a character sheet's crit with and without the second druid.
6. **Other all-crit (aura 290) buffs in the raid preset: none.**
   - Every aura 290 area aura in the 70790 SpellEffect table: Moonkin Aura and Leader of the Pack (both effect 65), plus Demonic Tactics 1243030 (effect 143, the warlock's own pet only).
   - The preset's other buffs give no crit: Arcane Intellect, Mark of the Wild, Fortitude, Divine Spirit, Might, Wisdom, Battle Shout, Trueshot Aura (ranged AP), Heart of the Lion (stats and AP), Strength of Earth, Windfury and Mana Spring. Its debuffs are school or armor debuffs.
   - The other aura-290 buffs in the client reach no raid preset player. They are personal, proc or set effects: Ancestor's Vengeance 16939 (a party proc on 1314143, which nothing trains), Mind's Eye 1219232 (Naxxramas priest set), Critical Aim 1218574, and Fervor of the Temple Explorer 446695 (a SoD world buff).
   - So the exclusivity question is only between the two druid auras, which is what `core.RaidHasDruidCritAura` models.
   - A balance druid's own Moonkin Form beside a raid Leader of the Pack adds nothing (`druid.raidHasMoonkinAura`). In a bare fight its own aura still gives the druid 3%.
7. **Engine check (this round).**
   - `sim/core/buff_ranks.go` `RaidHasDruidCritAura` and `sim/core/buffs.go:247`: 3% `stats.Crit` (every table) once.
   - Level gates: Leader of the Pack from 30 (talent tier 4: 10 + 5 x 4) and Moonkin Aura from 40 (tier 6). These match `talents/druid.json` and round 22's learn-level rule.
   - **Fixed:** the curated preset's `leader_of_the_pack` reason still said "Melee and ranged crit", the 16 October state. Text only, no number moves. S 7a3d68662.

**Verdict for the integration doc:** the 3% raid crit stands, and its source is Blizzard's own words. Exclusivity is stated in two lines of post #9 (2026-11-03 03:20 UTC). Both auras are 3% all-crit in the client and in the region-91 hotfix push 113283. The only inference is mechanical: the "one 3% wins" result follows from both being 3%. Nothing in the client can confirm or deny exclusivity, and no log exists yet.

## Lane

- Site: wt/r23-druid from int/r23 25432c0f1. Fork: lane/r23-druid from int/r23 bff826207.
- Client 1.60.1.70790. Wago SpellEffect, SpellLabel, SpellClassOptions, ItemEffect and ItemXItemEffect re-fetched 5 Nov ~10:56 UTC. Blizzard thread 2360696 re-read ~10:55 UTC. Wowhead Forever tooltips read 5 Nov.
- Measured with spec-breakdown on frozen BiS (`$R/bis-frozen`) and guide builds (`$R/rt`). Explorations ran 3,000 iterations; before/after ran 10,000 at seeds 1 and 50001.

**Summary: 1 confirmed sim bug, fixed (origin old). No published number moves.**
- Wolfshead Helm still gave its 5 energy on a shift into Cat Form. Blizzard's 8 October notes and the client moved it to Shifting Power, and added 5 rage to Enrage. No published set wears the helm.
- Two text fixes, not sim bugs: the raid preset's Leader of the Pack reason, and the feral Bite idle reason.
- Re-checked and holding: round 22's four druid fixes, the crit exclusivity code, and the feral finisher thresholds after the stronger Rip and Bite.

## Round 22's druid changes, re-checked first

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Crit exclusivity** (E 5d23a8747): `RaidHasDruidCritAura` adds 3% to every table once. The moonkin's own form adds nothing when the raid carries either aura. Gate levels are 30 and 40, the talent tiers. | `sim/core/buff_ranks.go`, `sim/core/buffs.go:247`, `sim/druid/druid.go:249`, `sim/druid/forms.go:322`; Raid crit sourcing above | Aligned | - | - |
| P2 | **Rip and Bite attack power** (E c7a5fa846): 1.25% a point a tick on Rip (four points at most) and 3.75% a point on Bite. Rip stays 12 s at every combo point count: client duration 12000 ms, and the live tooltip for 9896 reads "5 points: 1071 damage over 12 sec" = (19 + 5 x 31.9) x 6. | `sim/druid/rip.go`, `ferocious_bite.go`; nether.wowhead.com/forever/tooltip/spell/9896 | Aligned | - | - |
| P3 | **No Cat Form below 20** (E f749164f6): `tryPowershift`, `maxShifts`, the Metamorphosis Rune and Reset all guard a nil `CatForm`. | `sim/druid/feral/rotation.go`, `items.go`, `feral.go` | Aligned | - | - |
| P4 | **Caster Faerie Fire swing reset** (E 655bee16c): the flag is gone, and no card casts the caster spell. | `sim/druid/faerie_fire.go` | Aligned | - | - |
| P5 | **Round 22 "Measure at refresh": feral finishers after the stronger Rip and Bite.**<br>Raid 60, horde, seed 1, 3,000 iterations, frozen BiS, guide raid build:<br>- Shipped (Bite at 5 points with Rip at 5 s or more; Rip under 2 s): 871.7.<br>- Bite gate at 2 s: 871.9. At 3 s: 873.5 (+0.2%). At 4 s: 872.4. At 7 s: 866.3. With no Bite: 864.5 (-0.8%).<br>- Rip refresh at 0 s: 870.5. At 1 s: 868.1. At 3 s: 863.8.<br>The shipped rows are within 0.2% of the best variant. A 0.2% gate tweak is not worth a card change that would also need checking level by level. | `$R/druid/rot/feral-*.json` | Aligned (rows stay) | - | - |
| P6 | **Feral Bite `expected_idle` reason** described a 3+combo-point Rip (8 to 16 s) and a 6 s gate. Engineering text, not player-facing: rewritten to state the 12 s Rip, the 5 s gate and P5. | `data/curated/apl/druid-feral.json` | Text fix (not a sim bug) | - | S 57023a545 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Raid preset reason text.** `leader_of_the_pack` said "Melee and ranged crit", the 16 October state. Both reasons now say 3% crit of every kind, exclusive since 3 Nov. | `data/curated/presets.json:68,73` | Text fix (not a sim bug) | - | S 7a3d68662 |
| X2 | **Other aura-290 raid buffs:** none in the raid preset (Raid crit sourcing, point 6). | wago SpellEffect 70790 | Aligned | - | - |

## Feral cat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Wolfshead Helm's energy was on the wrong trigger.**<br>**What was wrong:** the engine gave Wolfshead Helm's 5 energy on a shift into Cat Form, which a feral using Shifting Power never does mid-fight. It gave Enrage nothing.<br>**What the sources say:**<br>- Blizzard, 8 Oct (post #5): "Wolfshead Helm's effect has changed to granting 5 additional Energy when shifting power is used. Rage portion of the helm's effect remains the same."<br>- Item 8345's equip spell 17768 (ItemEffect 98456): effect 0 is aura 108 on Enrage's class mask 0x80000 (SpellClassOptions 43824), 50 tenths of a point. Effect 1 is a dummy for 5, answered by the server's energize 1310990 (5 energy).<br>- Wowhead Forever tooltip 17768: "You gain an additional 5 Rage from activating Enrage and an additional 5 Energy from activating Shifting Power."<br>**What changed:**<br>- Shifting Power with the helm now gives 45 energy.<br>- A shift into Cat Form gives only Furor's energy.<br>- Enrage gives 15 instant rage with the helm.<br>Tests `TestWolfsheadHelmAddsFiveEnergyToShiftingPower` (40, want 45), `TestWolfsheadHelmGivesNoEnergyOnAShiftIntoCatForm` (5, want 0) and `TestWolfsheadHelmAddsFiveRageToEnrage` (10, want 15): red, then green. TestP1Feral and TestFeralTank are unchanged, because their gear sets do not wear the helm.<br>**Effect:** no published set wears the helm, so no published number moves. With the helm forced on (`-gear head:8345`, horde, seed 1, 3,000 iterations), it now loses less to the published head at each band:<br>- Raid 60: 845.3 to 861.9, against 871.7 with the published head.<br>- Bare 60: 337.5 to 342.1, against 348.0.<br>- Bare 50: 199.9 to 202.0, against 204.3.<br>- Bare 40: 140.0 to 141.7, against 143.9.<br>It still loses at every band, so no BiS head should change. | Blizzard thread 2360696 post #5 (8 Oct); wago ItemXItemEffect 93383, ItemEffect 98456, SpellEffect 694584/694585/1340972, SpellClassOptions 43824 (70790); nether.wowhead.com/forever/tooltip/spell/17768, /item/8345 | CONFIRMED-BUG | old: 2cdf50a70 (the build-70291 follow) set 5 energy on the Cat Form shift. The 8 Oct line was missed by every round, including round 22's line-by-line read. | E f52b25882 |
| C2 | **Re-audit, raid build point by point** (Genesis 5, Nature's Majesty 2, Nature's Reach 2, Ferocity 5, Feral Swiftness 2, Feral Instinct 3, Shredding Attacks 3, Savage Fury 2, Sharpened Claws 2, Predatory Strikes 3, Predatory Discipline 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 5, Shifting Power 1, Heart of the Wild 5, Naturalist 5):<br>- Each matches its 70790 text.<br>- Natural Instinct's mask {0x839800, 0x540} holds every cat and bear special and Rip; the engine's `DruidSpellMaskMeleeAbilities` and Rip match.<br>- Feral Swiftness and Feral Instinct do nothing for a single-target cat. They are prerequisite fill, and the keep rule leaves them as in rounds 20 to 22.<br>- Spells against their Forever tooltips: Shred 200% plus 80, Claw 140% plus 115, Ravage 435% plus 343, Bite rank 5 3.1 per energy, Shifting Power 40 energy and 8 s.<br>- Bite rank 6 (31018, at 60) is a level-60 trainer book rank. That stays pending with the owner. | talents/druid.json; wago SpellEffect / SpellClassOptions; tooltips 9830, 9850, 9867, 22829, 1322605 | Aligned | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Wolfshead Helm's Enrage rage: C1. The bear BiS does not wear the helm. | as C1 | (C1) | - | E f52b25882 |
| B2 | **Re-audit:**<br>- Primal Bite casts rank 3 (1238070) at 60, matching trainables (40/50/60).<br>- Lacerate is 25/34/42/50/58. Maul is 7 ranks and Swipe 5 (trainables).<br>- The prepull Enrage pays 10 + 20 rage.<br>- Natural Instinct reaches Maul, Swipe, Lacerate and Primal Bite (C2).<br>- The threat constants stay OPEN (round 21 B4, round 22 B3: they need a live tank log). | trainables/druid.json; raid 60 breakdown | Aligned / OPEN (threat) | - | - |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Re-audit, raid 60 breakdown:**<br>- Starfire and Wrath miss 11.3/11.5%: 17% base against a +3 boss, less Nature's Reach's 4% and about 1.6% gear hit.<br>- Moonfire's 2.3% is tick-weighted, and Insect Swarm's 0% is the no-hit-counter column (round 19 L5).<br>- Eclipse is up 77.6% of the fight, and Nature's Grace 20.3% (13.3 procs).<br>- Poised shows 99.9% uptime on a moonkin. It is a level-32 passive for every druid that changes only energy costs, so it does nothing here.<br>- Moonkin Form's 3% is not added again beside the raid aura (P1). | raid 60 breakdown | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Re-audit, raid 60 heal fight:**<br>- Crit: the every-table crit is 9.0% (raid aura 3 + Nature's Majesty 4 + base), plus 7.44% spell crit from Intellect. That is 16.5% on Rejuvenation, Healing Touch, Bloom, Swiftmend and Tranquility ticks, and Regrowth's 25.9% with Improved Regrowth.<br>- Bloom: 77 a fight, against 80.6 Rejuvenation and Regrowth casts (refreshes do not bloom).<br>- Tranquility: 3 casts in 300 s (Improved Tranquility, 2 min).<br>- Read Leyline does nothing for horde, which is correct.<br>As round 22 R1 for the rest. | raid 60 heal breakdown | Aligned | - | - |

## Blizzard's druid lines

- **Since round 22:** none. The thread's last post is #10 (4 Nov 19:35 UTC, UI only), which round 22 had already read at 07:30 UTC 5 Nov.
- **Re-read in full (24 Sep to 3 Nov) against the engine.** One line was missed by every earlier round: Wolfshead Helm (8 Oct), C1.
- **8 Oct "Rage portion ... remains the same":** the 70790 client and its tooltip put that rage on Enrage. Client data wins.
- Everything else is as round 22's list.
- **Endurance (Tauren):** 5% health only since 26 Oct. Already as the engine has it (`core/racials.go`).

## Before/after

Raid 60, 10,000 iterations, frozen BiS, guide builds (`ALLRT=1`). "Before" is `$R/before.tsv`; "after" is this lane's engine and site (`$R/druid/after.tsv`).

| Spec | alliance s1 | horde s1 | relative, both seeds and factions |
|---|---|---|---|
| Balance | 635.7 to 635.7 | 628.7 to 628.7 | 0.00% |
| Feral cat | 888.0 to 888.0 | 871.2 to 871.2 | 0.00% |
| Feral bear (threat/s) | 1715.1 to 1715.1 | 1710.5 to 1710.5 | 0.00% |
| Restoration (effective HPS) | 927.3 to 927.3 | 923.9 to 923.9 | 0.00% |

No other class moves: C1 touches only druid code.

## Files to regenerate

- None from this lane. No fork golden moved (TestP1Feral and TestFeralTank pass unchanged). No ladder, card, APL action, rank row or guide build changed.

## Measure at refresh

- Wolfshead Helm for feral at 40 to 59 in the BiS re-rank. It now gets its Shifting Power energy, but still measures about 1% under the published head at bands 40, 50 and 60 (C1), so no change is expected.

## For other lanes

- **Integrator:** the "Raid crit sourcing" section at the top is the citation for integration.md. The exclusivity is Blizzard's own words. The rest is client rows and tooltips. No log exists yet.
- **Hunter / warlock lanes (check only):** pets inherit the owner's crit in Forever (16 Oct note). If a pet also received the raid's 3% aura directly, the crit would count twice. This lane did not trace pet buffs.

## For the guide lane

Carried from round 22, still open on main:
- **feral-bear.md:83** "Swipe moves to the top on a pack": true only in a raid at 60 on three or more targets (round 22 G2).
- **balance.md:63, feral.md:51, feral-bear.md:48:** could add that Leader of the Pack and Moonkin Form's aura do not stack (a Should; the cited source is above).

No guide build, APL action, card note or rank row changed in this lane, so no prose was due and no wow-player pass was needed.

## Still open

- **Exclusivity in practice:** a log or character sheet with both druids in range (Raid crit sourcing, point 5).
- **Bear threat constants** (Maul, Swipe, Lacerate, Primal Bite): a live tank threat capture.
- **Eclipse on a missed Wrath; the unarmed paw:** a live Balance log at 40+, and a bare-handed cat.
- **The exact Rip and Bite attack power shares** ("about 25%"): a 60 cat's Rip ticks and Bites at a known attack power.
- **Logs** last tried 5 Nov ~10:58 UTC: api.foreversixty.gg lists only the 20 Oct report; foreverlogs.gg 403; Warcraft Logs v1 401.
- **Pending owner, not reopened:** as in the brief, including the level-60 book ranks (Ferocious Bite 6, Healing Touch 11, Rejuvenation 11, Starfire 7).

## Sources

- Blizzard, "WoW Forever Beta Development Notes – Updated November 2", https://us.forums.blizzard.com/en/wow/t/2360696 (posts #4 1 Oct, #5 8 Oct, #6 16 Oct, #9 3 Nov 03:20 UTC; read 5 Nov ~10:55 UTC). "World of Warcraft: Forever - Info & Known Issues", https://us.forums.blizzard.com/en/wow/t/2396179 (4 Nov).
- wago.tools, build 1.60.1.70790 (5 Nov ~10:56 UTC):
  - Tables: SpellEffect, SpellLabel, SpellClassOptions, SpellMisc, SpellAuraOptions, SpellCategories, ItemEffect, ItemXItemEffect.
  - Hotfixes: https://wago.tools/hotfixes?search=701335 and ?search=702727 (push 113283, region 91).
- Wowhead Forever tooltips: spells 24907, 24932, 24858, 17007, 9896, 9830, 9850, 9867, 22829, 1322605, 1328827, 17768, 1310990; item 8345.
- Warcraft Tavern, Forever druid guide, https://www.warcrafttavern.com/forever/guides/druid/ (pre-15 Oct state).
- Conquest Capped, WoW Forever raid buffs, https://conquestcapped.com/guides/wow-forever/wow-forever-raid-buffs/ (updated 2 Nov 2026).
- ClassicWoWCommunity/forever-bugs issue #275, https://github.com/ClassicWoWCommunity/forever-bugs/issues/275.
