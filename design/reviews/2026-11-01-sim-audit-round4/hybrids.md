# Sim audit round 4: hybrids (2026-11-01)

Druid (Balance, Feral cat, Feral bear, Restoration), Shaman (Elemental, Enhancement, Restoration) and Paladin (Retribution, Protection, Holy).

**Setup:**
- Client 1.60.1.70707.
- Engine: pin 640da215f, plus fork branch `lane/r4-hybrids`.
- Site: branch `wt/r4-hybrids`.
- Round 3 is `design/reviews/2026-11-01-sim-audit-round3/hybrids.md`.

**How claims were checked:**
- Client tables: the raw 70707 wago CSVs, fetched 2026-11-01 (`SpellEffect`, `SpellMisc`, `SpellAuraOptions`, `SpellClassOptions`, `SpellCooldowns`, `SkillLineAbility`, `SpellEquippedItems`, `SpellLabel`, `Spell`, `SpellName`).
- Wowhead Forever tooltips (`nether.wowhead.com/forever/tooltip/spell/<id>`), read 2026-11-01.
- Blizzard's "WoW Forever Beta Development Notes" thread, sections of 24 September, 8 October and 16 October 2026 (https://us.forums.blizzard.com/en/wow/t/2360696), read 2026-11-01. The third-party summary of the 15 October notes (https://wowclassicforever.info/patch-notes/) was used as a second, weaker source.
- Sim runs: `sim/cmd/spec-breakdown`, raid preset, band 60, 10000 iterations, seeds 1 and 50001.
- The conformance goldens (`sim/core/testdata/conformance/*.golden.md`) already hold every druid, shaman and paladin rank's amount, cost, cast time and coefficient against the 70707 client file; this round looked at what those tables cannot see: procs, talents, outcomes, racials and rotation.

Fix-sha column: "E" is a fork commit on `lane/r4-hybrids`, "S" is a site commit on `wt/r4-hybrids`.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Heals over time never crit. The client lets them.** `core.Dot` skipped the crit rule for every heal over time, and Renew, Riptide and Lightwell's renew ticked through `spell.OutcomeHealing`, which never crits a tick. | **Client:** every trained rank of Rejuvenation, Regrowth, Tranquility, Wild Growth, Renew, Riptide and Healing Stream carries `SPELL_ATTR8_PERIODIC_CAN_CRIT` (SpellMisc Attributes_8 0x200); the NPC copies (8070, 12160, 21791 and others) do not.<br>**Blizzard, 24 September:** "Corrected an issue and now Rejuvenation can land critical hits."<br>**Fix:** a heal over time follows the client flag (`core.periodicCanCrit`), on the healing crit chance and the spell's own multiplier. `tools/periodiccrit/gen.py` also reads periodic heals (aura 8), and only heal ids join `periodic_crit_auto_gen.go`.<br>**Tests first:** `TestHotsFollowTheClientFlag`, `TestRiptideTicksCanCrit` and `TestRenewTicksCanCrit` failed, then passed. | CONFIRMED-BUG | E 9b51cbd88, E 1dd5e6111 |
| X2 | **Touch of the Grave read build 70009's passive** (10%, 1 s internal cooldown): about 24 drains a fight for an Undead paladin. | **Client 70707:** two passives by class. SkillLineAbility 55923 gives 1260189 (ClassMask 11: warrior, paladin, rogue) at ProcChance 20; 55924 gives 1260201 (ClassMask 400: priest, mage, warlock) at 30. Both have a 25000 ms ProcCategoryRecovery.<br>**Beta notes summary, 15 October:** "a 30% chance for casters and 20% for Paladins, Rogues and Warriors, a 25-second internal cooldown".<br>**Test first:** `TestTouchOfTheGraveMatchesTheClientRows` failed, then passed. | CONFIRMED-BUG | E 537622d34 |
| X3 | **Orc Savage was still Axe Specialization** (1% crit only with an axe). | **Client:** 20574 "Savage: Increases your critical strike chance with all spells and abilities by 1%", aura 290, no SpellEquippedItems row. Wowhead Forever's live text is the same. The 16 October notes summary: "Axe Specialization becomes Savage, 1% critical strike chance".<br>**Effect:** orc casters and non-axe melee gain 1% crit (Elemental horde +0.8%).<br>**Test first:** `TestSavageIsCritWithAnyWeapon`. | CONFIRMED-BUG | E 46b72f0f6 |
| X4 | **Tauren Endurance gave 1% Hit** from a demo transcription. | **Client 20550:** 5% Health and 2% critical strike bonus (aura 163). **Wowhead Forever (live):** "Total Health increased by 5%". Neither gives Hit.<br>**Test first:** `TestEnduranceIsHealthWithoutHit`.<br>The table's 2% critical bonus is not added: the live text no longer shows it (Still open 4). | CONFIRMED-BUG | E 46b72f0f6 |
| X5 | **Blessing of Kings with Heart of the Lion (round 3 C2): still no Forever evidence.** Searched Exa and web search on 2026-11-01 for Reddit, Wowhead comments, Blizzard notes and beta reports. Every hit is Season of Discovery 2023-2024 ("hunter aura and kings dont stack"). | **Who it affects:** the curated raid preset has no Kings, so no published number. The /sim page's "raid-buffed" list (`web/src/lib/sim/settings.ts:160`) has Kings and no Heart of the Lion, so a default /sim run gets one 10%, whichever reading is true. Only a request that ticks both is at stake: the engine multiplies them (1.10 x 1.10, `sim/core/buffs.go:264`, `:302`).<br>**Stakes (Kings on top of the raid preset, horde, seed 1):** Feral +5.0%, Enhancement +4.5%, Retribution +4.4%, Elemental +2.1%, Balance +0.7%.<br>**For the integrator:** if the /sim list carries Heart of the Lion in place of Kings, it matches the curated preset, and the answer stops mattering for default runs. | OPEN: a Forever character sheet with both buffs | - |
| X6 | Blood Fury's cooldown: Blizzard's notes of 16 October say "Cooldown increased to 3 minutes (was 2 minutes)". The 70707 table (SpellCooldowns 55612) and Wowhead Forever's live tooltip say 2 minutes. Wowhead's live data does carry the same day's War Stomp change (1.5 min), which argues it is current. The engine keeps 2 minutes. | Round 3 C4; Wowhead 20572 and 20549, read 2026-11-01. | OPEN (one Forever source each way) | - |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | Tauren lost Endurance's false 1% Hit (X4): horde -1.0%. | X4. | CONFIRMED-BUG (X4) | E 46b72f0f6 |
| B2 | Genesis and Nature's Majesty at 2% a rank. The client row is 1, and Wowhead's spell tooltip shows rank 1 at 1%. Wowhead's live talent payload, which the site lays over the trees (`data/pipeline/normalize/wowhead_overlay.py`), reads 2/4/6/8/10% and 2/4%. The engine follows the live payload, as round 3 did. | `talents/druid.json`; Wowhead 1223081, 1223082. | Aligned (one hotfix source) | - |
| B3 | Moonfire, Insect Swarm and Starfire rows, and Omen of Clarity at 4% a spell, doubled in Moonkin Form, with a 10 s cooldown halved in Moonkin. Clearcasting is not spent by Wrath (16864 text). | Conformance golden; `sim/druid/talents.go`. | Aligned | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Horde -0.6% from X4. Shred 155%, Claw 110% and Savage Fury's mask are unchanged from round 3. | X4; `weapon_percent_test.go`. | CONFIRMED-BUG (X4) | E 46b72f0f6 |
| F2 | Set bonuses on the published band: Shadowcraft Armor 5 pieces. The 5-piece energize (27787) procs at the unstated 1 a minute; the client gives no chance for it. | `client_dungeon_sets.go`, `client_procs.go`. | Aligned (assumption already listed) | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | **Maul never spent Clearcasting.** Clearcasting (16870, ProcCharges 1) is spent in OnCastComplete, but Maul is cast off the swing with `SpellFlagNoOnCastComplete`. Once a Clearcasting landed, every Maul was free until it expired (15 s) or a Swipe or Lacerate spent it. Found on the integrator's request to check every Clearcasting after the mage bug. | Maul is in 16870's mask (0x800). **Test first:** `sim/druid/tank/omen_consume_test.go` failed, then passed. The `TestFeralTank` golden moves. | CONFIRMED-BUG | E 27240ef0d |
| G2 | Tauren Endurance (X4): no change in bear threat at band 60 (694.0 both seeds). | X4. | (measured) | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Bloom's triggered heal rolled cast procs.** Bloom's heal 1328622 fired OnCastComplete with the healing proc mask, so every bloom (about 55 a fight) rolled Vestments of the Devout's "your spellcasts" 5% (Sudden Insight: 6.0 procs a fight against about 4 from casts) and Omen of Clarity, whose 4% round 3 fitted per logged cast. | **Client:** 450576's text reads "spellcasts"; 1328622 is the hidden spell the aura triggers. The engine's other triggered spells (42 of 66 passive spells) carry `SpellFlagNoOnCastComplete`. **Test first:** `TestBloomTriggersNoCastProcs`. **Effect:** -2.3% effective HPS on the old profile. | CONFIRMED-BUG (client text plus the round 3 per-cast fit) | E 3b0ee304f |
| R2 | Heals over time crit (X1): Rejuvenation, Regrowth's heal over time, Tranquility and Wild Growth. Tests in `runs_test.go` now count critical ticks. | X1. | CONFIRMED-BUG (X1) | E 9b51cbd88 |
| R3 | **Assigned item 3, 91% coverage.** Every druid heal row matches the client: Healing Touch 9889 1920 + 6.6 a level, coefficient 1.0; Rejuvenation 9841 161 a tick, 0.2; Regrowth 9858; Tranquility 9863; Swiftmend 20% of base mana. The conformance file is byte-identical to 70707's spellconst. The druid went from 76% (2026-10-08) to 90% mainly through client 70707: Bloom is about 14% of its raw healing, and Genesis doubled. R1 takes 2.3% off and X1 adds about 4%, leaving 91.3% at 1.5. That is not a sim bug, so the profile is retuned by its own rule (S 65199446b): 1.65 is the smallest multiplier at which every raid set is between about 60 and 85% (Before/after). | `healsweep2.tsv` (lane scratch); `data/curated/heal-profile.json`. | Profile retune | S 65199446b |
| R4 | **Assigned item 4, Wild Growth.** The profile already has party damage: three of four members take a pulse every 5 s. With crits and a Wild Growth line (player 1 and 2 under 80% or 90%, Improved Tranquility's point moved to it), the druid measured 659.7 / 659.4 and 646.7 / 646.4 against the guide's 732.2 / 731.8: -10%. Wild Growth heals 4.0 a mana against Rejuvenation's 5.2. Improved Regrowth (2 points from Improved Tranquility, or 5 from Tranquil Spirit) also measures below the guide build (-0.5%, -0.2%). | Rotations `rot/druid-wg80.json` and `-wg90.json` (lane scratch). | Aligned (settled: the guide build holds) | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | The orc gets Savage's 1% crit (X3): horde +0.8%. | X3. | CONFIRMED-BUG (X3) | E 46b72f0f6 |
| E2 | **Clearcasting (integrator's check).** 16246 has one charge, and the first shaman damage spell spends it; a heal does not. Elemental Focus (16164) has no ProcCategoryRecovery, and the engine has no cooldown either. New test `sim/shaman/elemental/clearcasting_test.go`. | Client rows; `sim/shaman/talents.go:113`. | Aligned (test added) | E cff89dba6 |
| E3 | Lightning Bolt ranks 3 to 10 at 2.5 s and 0.714, and Chain Lightning at 2.0 s, 0.571 (0.517 on rank 3) with its cut costs, are the client's. Lava Burst 1326047 is learnable; Lava Lash 408507 is a rune (AcquireMethod 3), and no rotation casts it. | Conformance golden; SkillLineAbility. | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | Savage (X3): horde +0.9%. Stormstrike 8% of base mana and the 8 s cooldown, Windfury Weapon's 20%, and the totem's 1 extra attack are the client's. | X3; SpellEffect 16362, 10610; 17364. | CONFIRMED-BUG (X3) | E 46b72f0f6 |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Riptide's heal over time crits (X1). Healing Stream's ticks (10461, flagged) still cast a separate heal with `OutcomeHealing`; at 11 a tick plus 0.022 the crit is worth under 0.1%, and whether a totem's tick takes the shaman's crit is unread. | X1; `sim/shaman/water_totems.go:81`. | CONFIRMED-BUG (Riptide); Healing Stream OPEN | E 1dd5e6111 |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Seal of Righteousness took spell damage twice.** The swing proc added its 0.1 coefficient to its base, then the spell added it again: 0.2 in all. | **Blizzard, 16 October:** "Fixed an issue causing Seal of Righteousness to gain almost twice as much benefit from Spell Damage than intended." The engine's own comment named the double ("effectively scales with coeff x 2", `sim/paladin/sor.go`). Judgement of Righteousness's 0.429 from the same notes was already in. **Test first:** `sor_coefficient_test.go` measured 0.200, then 0.100. | CONFIRMED-BUG | E bb9912dc0 |
| P2 | **With P1, the Unsealed Wrath twist pays the other way.** Seal of Command kept up, three stored copies released by one swing under Seal of Righteousness: alliance 568.0 / 568.8 to 626.6 / 626.2 (+10.2%), horde 581.1 / 580.8 to 576.7 / 577.3 (-0.7%, before bis.yml re-picks the weapon). Seal of Command alone: 609.6 / 573.6. The old twist is still +3.1% / +1.9% over no twist. A build without Seal of Command or Unsealed Wrath keeps Seal of Righteousness, because Seal of Command alone measured 6% lower at band 40. Bands 20 to 50 are unchanged. The fork's copy is synced. | `-rotation` runs (`rot/ret-soc*.json`, lane scratch). | CONFIRMED-BUG (rotation, follows P1) | S 6494c1a30, E dc3d0a769 |
| P3 | **Touch of the Grave (X2)** on the Undead horde set: 24 drains a fight (27 DPS, 4.2%) fall to the client's rate. | X2. | CONFIRMED-BUG (X2) | E 537622d34 |
| P4 | **Assigned item 2, Holy Strike's flat bonus under the weapon percent:** no new log. Forever logs are still levels 13 to 30, and no training-dummy log exists. One structural point: 10333's flat bonus is effect 121 (normalized weapon damage) beside effect 31 (weapon percent), the same pairing whose flat part round 3's Shred logs showed is multiplied (F1 there). That fits the engine's reading, but it is an analogy, not a measurement. | SpellEffect 10333, 1866. | OPEN | - |
| P5 | **Unsealed Wrath's reading:** Wowhead Forever 1310735 says "Your melee attacks store a copy of your active Seal for 15 sec, stacking up to 3 times. Landing an attack with a different Seal consumes all stacks". Whether a special attack stores or consumes is still unstated. | Wowhead 1310735. | OPEN: a Retribution log | - |
| P6 | Reckoning (16 October: up to 4 charges, all spent by the next swing). The engine grants each extra attack at once, which changes only the timing, never more than four. | `talents.go` notes. | Aligned (timing only) | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Touch of the Grave (X2) on the Undead horde set: threat -4.4%. Seal of Righteousness (P1) does not reach the tank rotation, which keeps Seal of Fury. | X2; `-tank` runs. | CONFIRMED-BUG (X2) | E 537622d34 |
| T2 | Consecration's 0.042 a tick (8 October notes) and Champion of the Light's 10/20/30% (16 October) are in (171b9ff7f). | Blizzard notes; `consecration.go:58`. | Aligned | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | No heal over time; Flash of Light, Holy Light and Holy Shock as round 3 left them. The Human Spirit is 5% in the 70707 table (20598) and 10% in Wowhead's live text; the engine has 5% (Still open 5). | Conformance golden; Wowhead 20598. | Aligned (Spirit OPEN) | - |

## Before/after (relative %)

**Method:**
- Raid preset, band 60, 10000 iterations, mean change at seeds 1 / 50001.
- **Before:** pin 640da215f with origin/main (f2c4446ea).
- **After:** lane engine (head dc3d0a769) with `wt/r4-hybrids`.
- Healers are scored in effective HPS, tanks in threat a second. The healer "after" uses the retuned 1.65 profile. The engine change alone, on the old 1.5 profile, is in the last column.

| Spec | Changes | Alliance | Horde | Absolute, seed 1 (alliance; horde) | Engine only, old profile |
|---|---|---|---|---|---|
| Balance | X4 | 0.0 / 0.0 | -1.0 / -1.0 | 519.2; 518.0 → 512.8 | |
| Feral | X4 | 0.0 / 0.0 | -0.6 / -0.6 | 582.8; 585.8 → 582.3 | |
| Feral bear (threat) | G1, X4 | -0.9 / -1.0 | -1.1 / -1.0 | 674.2 → 667.8; 694.0 → 686.4 | |
| Restoration druid | R1, X1, profile | +2.8 / +2.7 | +2.4 / +2.3 | 712.0 → 732.2; 711.0 → 728.1 | +1.7 (712.0 → 723.9) |
| Elemental | X3 | 0.0 / 0.0 | +0.8 / +0.8 | 464.7; 462.6 → 466.3 | |
| Enhancement | X3 | 0.0 / 0.0 | +0.9 / +0.9 | 618.4; 634.1 → 639.8 | |
| Restoration shaman | X1, profile | +4.7 / +4.7 | +4.7 / +4.6 | 611.9 → 640.9; 610.7 → 639.6 | +1.1 (611.9 → 618.9) |
| Retribution | P1, P2, X2 | +6.3 / +6.1 | -11.2 / -11.1 | 589.5 → 626.6; 649.6 → 576.7 | |
| Protection paladin (threat) | X2 | 0.0 / 0.0 | -4.5 / -4.4 | 693.5; 721.2 → 689.1 | |
| Holy paladin | profile | +2.0 / +2.1 | +2.0 / +1.9 | 538.5 → 549.5; 531.3 → 541.9 | 0.0 |

**Outside the group** (all 18 other specs run; those not listed move 0.0):

| Spec | Cause | Horde | Alliance |
|---|---|---|---|
| Holy priest | X1 (Renew), profile | -1.4 / -1.4 | -2.4 / -2.4 (old profile: +1.7, 651.1 → 662.4) |
| Discipline | X1, profile | +0.5 / +0.5 | +0.7 / +0.8 (old profile: +1.2) |
| Shadow | X2 (Undead) | -0.1 / -0.2 | 0.0 |
| Fire | X2 (Undead) | +1.4 / +1.4 | 0.0 |
| Arcane | X2 (Undead) | +0.8 / +0.8 | 0.0 |
| Arms | X3 (Orc, non-axe) | +0.8 / +0.8 | 0.0 |
| Protection warrior (threat) | X4 (Tauren) | -1.2 / -1.3 | 0.0 |

Heal profile at 1.65 (alliance, raid / bare, seed 1): druid 84.0% / 44.0%, shaman 73.5% / 25.3%, Holy paladin 63.0% / 18.9%, Holy priest 72.9% / 35.3%, Discipline 60.4% / 31.9%. At 1.6 the druid stays at 86.5%, so 1.65 is the smallest multiplier that fits.

Nothing nightly-owned was committed. `TestRotationLadder` fails on this branch as expected; the integrator regenerates it at the pin. Fork goldens regenerated: TestElemental, TestEnhancement, TestFeralTank (twice), TestP1DPSWarrior and TestP1TankWarrior.

## Still open

1. **X5, Kings with Heart of the Lion.** A Forever character sheet with both buffs settles it. The /sim page carries Kings and no Heart of the Lion; listing Heart of the Lion instead makes the question moot for default runs.
2. **X2, Touch of the Grave's cooldown.** The table and the 15 October summary say 25 s; Wowhead Forever's live tooltip says 15 s. A log of an Undead's drains settles it.
3. **X6, Blood Fury 2 or 3 minutes.** Blizzard's notes against the table and Wowhead's live text.
4. **X4, Endurance's 2% critical strike bonus** (in the table, not in the live text).
5. **The Human Spirit 5% or 10%** (table against live text).
6. **P4, Holy Strike's flat bonus** and **P5, Unsealed Wrath's special attacks**: a level-60 Retribution log after launch.
7. **H1, Healing Stream's tick crit.**
8. **Omen's 4%** still rests on levels 20 to 30 (round 3).

## Sources

- Client 1.60.1.70707 wago tables (fetched 2026-11-01): SpellMisc (Attributes_8 of every heal over time), SpellAuraOptions 243729 and 243733, SkillLineAbility 55923 and 55924, SpellEffect 20574, 20550, 1259719, 20597, 20598, 10333, 1866, 20286, 20293, 25713, 450576, 16246, 16164, 16870, 16864.
- Blizzard, WoW Forever Beta Development Notes (24 September, 8 and 16 October 2026): https://us.forums.blizzard.com/en/wow/t/2360696
- Third-party summary of the 15 October notes: https://wowclassicforever.info/patch-notes/
- Wowhead Forever tooltips (read 2026-11-01): 1328621, 1328622, 1223081, 1223082, 17116, 18562, 20572, 20549, 20550, 20574, 20597, 20598, 1259719, 1260189, 1260201, 1310735.
- Season of Discovery, Kings and Heart of the Lion (2023-12-11): https://us.forums.blizzard.com/en/wow/t/blessing-of-kings-doesnt-stack-with-aspect-of-lion-but/1738967
