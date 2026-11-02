# Sim audit round 10: healers

Lane: wt/r10-healers (site, from origin/main bb619a188) and lane/r10-healers (fork, from 4d3129cae). Specs: Restoration druid, Restoration shaman, Holy priest, Discipline priest, Holy paladin, at the level-60 heal profile (raid and bare). The sub-60 heal profile and the sub-60 healer rows of rank-choices.json belong to the healprofile lane and are untouched here.

Client truth: 1.60.1.70707 with the hotfix push layered in (raw CSVs fetched 2026-11-01 21:42 UTC). Four delegated engine-vs-client sweeps, one per class, re-derived every cast spell, every taken talent's curve and mask, the published gear and the racials. I re-read each reported bug against the client row and the code before confirming it. Their notes are in [healers/](healers/) (aud-druid.md, aud-paladin.md, aud-priest.md, aud-shaman.md).

**New this round: event-level beta logs.** foreverlogs.gg's report timeline (`/api/reports/<id>/encounters/<enc>/timeline`) gives every heal, overheal, cast and buff event, with an is_critical flag. I fetched it for 64 healers at levels 20 to 35 (2,100 encounter files, 2026-11-02). The amounts are effective healing, so a raw heal is the heal event plus the overheal event at the same timestamp, target and spell. The scripts are in [healers/logtools/](healers/logtools/) (fetch.py, critsize.py, rates.py, swift2.py). The site rate-limits the endpoint, so the fetch is partial.

The short version:
- **Swiftmend does not use up the heal over time in Forever.** In the logs, 55 of the 57 effects up at a Swiftmend ticked on to their natural end. With both effects up, all 16 Swiftmends healed for the Regrowth. The engine cancelled the effect and read the one with the least time left. With the effect kept, the rotation's "last-ditch under 30%" Swiftmend became a heal of its own: under 75% is +4.7% more on top. Together they are the largest healer move of the round.
- **Gift of Nature does not reach Tranquility's ticks.** Two level-35 druids' Tranquility ticks fit their Healing Touch's healing power only without it.
- **Healing Stream's ticks crit only on every-table crit**, not the shaman's Intellect crit (flagged by the crit lane's log tally).
- **Two crit relics were dead weight in every Restoration and Holy set:**
  - Totem of Urgency: +4% Lesser Healing Wave crit, never registered.
  - Libram of Infusion: +6% Holy Shock crit, which never reached the heal.
- **Nature's Grace procs on heal crits too** (latent: no published healer takes it).
- **Two small priest fixes:** Mental Agility on Inner Fire; Inner Focus's crit off Renew's ticks.
- **Holy Power's crit-heal size is still open.** No logged paladin can have it: it needs 25 Holy points, and the logs stop at level 35. The 26-point Holy paladins crit at exactly 1.50x.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Heal crit reads spell crit, per the client masks.** Measured heal crit rates at raid 60 (horde, seed 1; spec-breakdown's new crit table):<br>- Holy priest: 7.03 spell + 8 every-table + 5 Holy Specialization = 20.03. Flash Heal 20.1, Renew ticks 20.0, Prayer of Mending 20.1.<br>- Discipline: 20.39 expected, Heal 20.6 and Renew ticks 20.3. Renewed Hope and Inner Focus lift the others.<br>- Shaman: 9.48 + 13 + 5 Tidal Mastery = 27.48. Healing Wave 27.9, Chain Heal 27.6, Riptide 27.3. Healing Stream is 22.5 (no Tidal Mastery; it falls to 13 after A7).<br>- Druid: 7.48 + 12 = 19.48, every heal 19.2 to 19.9.<br>- Paladin: 10.03 + 15 = 25.03. Flash of Light 25.2. Holy Light 27.3 and Holy Shock 51.2 carry Divine Favor and Boundless Light. | `sim/core/spell_result.go:277-282` (Crit + SpellCrit + the spell's bonus). Every crit buff in these presets is aura 290: Leader of the Pack 24932, Moonkin Aura 24907, Elixir of Sages 1250941, Conviction 20117, Thundering Strikes 16255, Nature's Majesty 1223082. | Aligned | S c589a2429 (tooling) |
| A2 | **Crit talents land in the right spells.** Each one's client effect mask matches the engine's spell set:<br>- Holy Specialization (14889, aura 107 op 7, 0x9541ec0/0x18024) and Tidal Mastery (16194, 0x5c0/0/0x10) both match.<br>- Lightwell Renew has no class mask, so it gets no Holy Specialization (15.3%), which is right.<br>- Healing Stream's 0x2000 is outside Tidal Mastery's mask, which is right.<br>- Divine Favor (20216, +100 on 0x80206000) covers Holy Light, Flash of Light and both Holy Shock spells, and is spent by the first of them.<br>- Boundless Light, Renewed Hope and Inner Focus are checked in aud-paladin.md and aud-priest.md.<br>- Improved Regrowth (17074, 0x40) and Holy Power (5923) are not crit-chance talents in this client. | SpellEffect 692220, 694886, 696789, 693957, 683158; `sim/priest/talents.go:262`, `sim/shaman/talents_restoration.go:69`. | Aligned | - |
| A3 | **On-crit procs read the right crits.** Illumination, Infusion of Light, Divine Aegis and Inspiration fire on a non-periodic heal crit only. Water Shield (ProcTypeMask 0x362A8, no periodic bit) ignores tick crits. Divine Aegis and Inspiration say "non-periodic" and use OnHealDealt, which ticks never reach. Ancestral Healing and Inspiration are armor, which is unmodelled. | `core/aura.go:101-103` (OnHealDealt vs OnPeriodicHealDealt); SpellAuraOptions 223904. | Aligned | - |
| A4 | **Heals over time crit, at 1.5x.** Beta logs: Renew ticks crit 7.1% (156 ticks) at 1.518x, Rejuvenation ticks 3 to 10% at 1.50x, and Riptide ticks crit in the engine by the client flag. Every heal-over-time rank the specs cast carries attr8 0x200. | critsize.py on reports 2747, 2752, 2760. | Aligned (client flag and logs) | - |
| A5 | **Heal crit size is 1.5x everywhere it can be measured.** Holy Light 1.496 to 1.518 (47 and 55 heals), Flash of Light 1.46 to 1.55 (four paladins), Healing Wave 1.48 to 1.52, Prayer of Healing 1.49, Swiftmend 1.50. | critsize.py. | Aligned | - |
| A6 | **Holy Power's crit-heal size** (5923: aura 50, misc 2, curve 10 to 50). The paladin sweep found that the only other client row with aura 50 is 21540 "Improved Block Value", which weakens the "modern aura 50" reading. Holy Power sits in tier 5 and needs 25 Holy points first. The logs' best Holy paladins (171783, 132337, 34190: 26 points at level 35) can hold at most one rank, and their crits are 1.496 ± 0.011 and 1.494 ± 0.010. Rank 1 under any model would be at least 1.55, so they took none. | critsize.py; aud-paladin.md P1. | OPEN (a level-40+ Holy paladin's crit beside a non-crit) | - |
| A7 | **Healing Stream's ticks crit only on the shaman's every-table crit.** In the crit lane's tally, 21,134 ticks from five shamans with no Enhancement points crit 0 times, while their own heals crit 4 to 6%. The Enhancement shamans, who take Thundering Strikes 5 (aura 290), crit 4.4 to 6.4% whatever their Intellect. So the totem's heal rolls item crit and aura-290 crit, and not the Intellect crit or the spell base. The engine rolled the full heal crit. | `r10/crit/py/hst.py` over the r9 crit lane's ability tallies; flag 0x200 on 5672 to 10461. Test `TestHealingStreamIgnoresTheShamansSpellTableCrit` failed first. | CONFIRMED-BUG (logs) | E 30003fb0d |
| A8 | **Swiftmend does not use up the effect it reads.** In 33 logged Swiftmends from three druids (levels 25 and 35):<br>- Every non-critical one is exactly 4 Rejuvenation or 7 Regrowth ticks of the effect on the target (208 = 4 × 52, 280 = 7 × 40, 394 = 4 × 98.5, 406 = 7 × 58). A crit is 1.50x.<br>- 55 of the 57 effects up at the cast ticked on to their natural end.<br>- With both up, all 16 healed for the Regrowth, including 12 where the Rejuvenation had less time left.<br>- The client row 18562 has one effect, the heal, and nothing that removes an aura. Only the tooltip says "consumed".<br>The engine cancelled the effect and read the one with the least time left. It now leaves the effect up, so it still ticks and Blooms, and reads Regrowth first. | swift2.py on reports 2747, 2752, 2760; https://nether.wowhead.com/forever/tooltip/spell/18562 (read 2026-11-02). Tests rewritten first: `TestSwiftmendNeedsAnEffectAndLeavesItUp`, `TestSwiftmendReadsRegrowthWhenBothAreUp`, `TestBloomStillFollowsASwiftmendedRejuvenation`. | CONFIRMED-BUG (logs + client row) | E 9badd243f |
| A9 | **Restoration's Swiftmend line under 75%.** The rotation kept Swiftmend as a last-ditch button under 30% because it used to eat the tank's heal over time. The old engine was right to hold it there: at 70% it measures -4%. On the fixed engine, thresholds from 0.3 to 1.0 measure 799 → 836 at 0.7 to 0.8 (plateau). Under 75% is adopted: +4.7% at raid 60 on both factions and both seeds, +1.4% bare 60. At bare 40 and 50 it is within 0.4% (the healprofile lane is rebuilding those fights). | spec-breakdown -rotation variants, 3,000 then 10,000 iterations. | CONFIRMED (rotation follows the fix) | S 03690eeab, E 245898b7d |
| A10 | **Swiftmend's own modifiers.** A druid with about two Gift of Nature ranks still heals exactly 7 × 40, so Swiftmend does not add Gift of Nature a second time. That matches the engine, which takes the snapshot's own multiplier and ignores attacker mods. Genesis cannot be tested: the logged druids have no Balance points. | aud-druid.md D2. | Aligned (logs) | - |
| A11 | **Gift of Nature on Tranquility.** Gift of Nature's DAMAGE-op effect names Tranquility; its DOT-op effect (Forever's 1291839) does not. Logs:<br>- Druid 127978 (report 2748, level 35): Healing Touch rank 5 (5189, 546 + 3.5 a level to 31, coefficient 1.0; 18 casts, median 851) gives healing power about 210 with Gift of Nature 5. Its Tranquility rank 1 ticks (740: 87 + 0.6 a level, 0.067; 28 non-crit ticks, 100 to 104) read 103. That fits 104 without Gift of Nature, against 114 with it. For any rank, the with-it reading is 109 or more.<br>- Druid 152866 (report 2770) agrees: 97 to 98 observed, 100 without, 110 with.<br>So the DAMAGE op reaches direct heals only. | Logs as above; SpellEffect 695243, 1291839, 680112. Test `TestGiftOfNatureRaisesEveryHealingSpellButTranquility` failed first. | CONFIRMED-BUG (logs + client masks; two characters) | E 041bed648 |
| A12 | **Oath of Light's scaling** (coefficient or flat). No logged paladin casts it: 1328548 arrived with the 70707 push, after the beta logs. | aud-paladin.md P7. | OPEN (one sworn Holy Light) | - |
| A13 | **Discipline's Inspiration** stays held. No fix this round values armor on the healed target. | - | OPEN (held by rule) | - |
| A14 | **Nature's Grace procs on heal crits.** 16880's ProcTypeMask 0x14000 takes helpful magic spells, and the text says "All non-periodic spell criticals". The proc's global-cooldown cut (16886 effect 1, 0x7407f7/0x4000000) names Healing Touch, Regrowth, Rejuvenation, Tranquility and Wild Growth. The engine procced it on damage crits only and cut only the Balance spells' cooldown. No published healer takes it: Restoration holds 17 Balance points and Nature's Grace needs 20. Balance casts no heal, so its numbers are unchanged. | `sim/druid/talents.go:119-160`. Test `TestNaturesGraceProcsOnACriticalHealNotATick` failed first. | CONFIRMED-BUG (latent; client proc row) | E d6c9fe1d7 |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | Swiftmend, its line, Gift of Nature on Tranquility, Nature's Grace | A8 to A11, A14 | see above | see above |
| D2 | **The rotation's note said "Top ranks only"** and the guide "The rotation does not downrank". The list casts Rejuvenation rank 4 (2090) on every target on purpose, as its own line says. | `data/curated/apl/druid-restoration.json`; guide. | CONFIRMED (prose) | S 73dbf55fc |
| D3 | **Rejuvenation's three DOT-op bonuses.** Gift of Nature, Improved Rejuvenation and Genesis are all aura 108 misc 22. The engine adds the first two and multiplies by Genesis, so 1.375 matches neither the additive (1.35) nor the multiplicative (1.392) rule. The logs can't separate the two rules. | aud-druid.md D4. | OPEN (same-op stacking, engine-wide) | - |
| D4 | **Sub-20 coefficient penalty on heals over time.** One druid's four heal-over-time ranks fit with no penalty. A low-rank Healing Touch fits only with it. | aud-druid.md D6. | OPEN (for the ranks and healprofile lanes) | - |
| D5 | Tauren Endurance: the Wowhead server text says only "+5% health", which matches the spell row. | aud-druid.md D7. | Aligned (leans) | - |
| D6 | **Aligned:** every rank of Healing Touch, Regrowth, Rejuvenation and Tranquility; every taken talent's value and mask; Innervate; Nature's Swiftness; Omen of Clarity (logs 3.8 to 5% against the engine's 4%); Bloom; the Devout set; Ankh of Life; the raid consumables; the night elf and tauren racials. The published 60 entries still carry Improved Wrath until the nightly runs (round 9's Heart of the Wild change). | aud-druid.md. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Healing Stream's crit table | A7 | CONFIRMED-BUG | E 30003fb0d |
| S2 | **Totem of Urgency (279249) did nothing.** It is an Enchanting relic, crafted, on Phase 1 loot, with equip spell 1306448: aura 107 op 7 +4 on 0x80 (Lesser Healing Wave). The relic table mapped only Lightning Bolt, Chain Lightning and Flame Shock, so no healing totem could register. All four Restoration sets at 60 carried an empty totem slot ("no_sourced_item"). The table now names Healing Wave (0x40) and Lesser Healing Wave (0x80). The totem is registered, and it joins the fork's item database the way the other Forever relics did. Equipped, it measures +0.2% HPS (horde and alliance, raid 60). Tidal Totem (272431, -5% Healing Wave cost) stays out: it is not in the fork's item database and its vendor opens later. | `sim/shaman/items.go:26-35`; SpellEffect 1333901; loot.json source "crafted:enchanting". Test `TestTotemOfUrgencyRaisesLesserHealingWaveCritOnly` failed first (unregistered). | CONFIRMED-BUG (client row) | E 999f1085b |
| S3 | `purificationHealingModifier` was dead code whose comment contradicted the client (Purification's masks don't reach Healing Stream). | aud-shaman.md R2. | CONFIRMED (dead code) | E 30003fb0d |
| S4 | **Earth Shield (974):** no SkillLineAbility row, so not learnable, which matches the rotation. A level-50+ trainer list would confirm. Purification and Healing Way adding or multiplying, and Chain Heal's jump order, are unchanged from round 9. | aud-shaman.md R3 to R5. | OPEN | - |
| S5 | **Aligned:** Healing Wave 9, Lesser Healing Wave 6, Chain Heal 3 and its jumps, Riptide 3 (direct, ticks and the Chain Heal bonus), Healing Stream 5 with Restorative Totems, Mana Tide 4, Mana Spring, Water Shield, Nature's Swiftness; every talent value and mask; the Devout set, trinkets and Ankh of Life; the tauren and dwarf racials; the consumables. | aud-shaman.md. | Aligned | - |

## Priest: Holy and Discipline

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Mental Agility skipped Inner Fire.** Its cost mask (SpellEffect 687728, word 0 0x8251c1eb) names every Inner Fire rank (mask 0x2). Inner Fire had no engine class mask. About 31 mana a fight. | Test `TestMentalAgilityFollowsTheClientMask` extended first. | CONFIRMED-BUG (client mask; under 0.01%) | E 8a241f75a |
| P2 | **Inner Focus's +25% crit reached Renew's ticks** while it was up. Its crit modifier (692357) names no Renew rank, and the text says "non-periodic". | Test `TestInnerFocusLeavesRenewsTicksAlone` failed first. | CONFIRMED-BUG (client mask; under 0.05%) | E 8a241f75a |
| P3 | **Prayer of Mending's heal and its talents.** The heal spells carry only word-3 0x200000, named by Twin Disciplines and Inner Focus. Holy Specialization and Spiritual Healing name the cast and jump (word 1 0x20), and no other priest spell has that bit. So the designers named the cast for a reason, and the engine's "apply both" stands. Prayer of Mending is a tier-6 talent, so no logged priest has it. | aud-priest.md; masks.py. | OPEN (round 9 P4) | - |
| P4 | Lightwell Renew's crit (DefenseType 0, cast by a gameobject); Renewed Hope's Weakened Soul cut per Penance pulse or per cast; Sudden Insight per cast or per target; Litany of Light across ranks. | aud-priest.md P3 to P6. | OPEN (each small; one log line each) | - |
| P5 | **Aligned:** every heal rank by script; every taken talent's curve and mask, both builds; Holy Binding; the gnome, troll, human and undead racials; the Devout set, Ankh of Life, Uther's Strength, Truefaith Vestments; the consumables. | aud-priest.md. | Aligned | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Libram of Infusion (279248) never reached the Holy Shock heal.** Its spell 1306429 is +6 crit on families 0x200000/0x10000. The heal 1311605 carries both bits, the cast only the first, but the relic table mapped bit 21 to the cast alone. Every Holy set scored the libram at zero ("no_dps_value"). Equipped, it measures +0.3% HPS. | `sim/paladin/items.go:30-36`; SpellEffect for 1306429; SpellClassOptions 1311605. Test `TestLibramOfInfusionRaisesTheHolyShockHealsCrit` failed first. | CONFIRMED-BUG (client masks) | E ac11976ed |
| H2 | Holy Power (A6) and Oath of Light (A12) | see above | OPEN | - |
| H3 | **Illumination's refund on Holy Shock rank 1.** The engine refunds 80 mana, half the cast's cost. The heal spell has no cost row of its own. | aud-paladin.md P3. | OPEN (a logged rank-1 crit and its 20272 energize) | - |
| H4 | Divine Favor and Devout 5P (0.0%); Concentration Aura (owner's kit choice). | aud-paladin.md P4, P6. | OPEN | - |
| H5 | **Aligned:** every rank of Holy Light, Flash of Light and Holy Shock (cast and heal); Divine Favor's mask, charge and cooldown; Spiritual Focus, Healing Light, Reverence, Illumination, Infusion of Light (logged buff 437063 after Flash of Light and Holy Shock crits), Boundless Light, Benediction, Divine Intellect, Conviction; the human and undead racials; the Devout set, Uther's Strength, Hand of Edward; the consumables. | aud-paladin.md. | Aligned | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Heals ignore school crit.** The paladin sweep flagged that `healingCritRating` left out `SchoolBonusCritChance`, so Vestments of Prophecy 5P (Holy) and The Ten Storms 5P (Nature) never reached a heal. No published healer wears either set. The crit lane owns the plumbing and has since let "Holy/Nature spells" bonuses reach heals (E b4546cdb6 on lane/r10-crit). | aud-paladin.md P2. | for the crit lane | - |
| X2 | **The raid preset stacks three Forever elixirs and a flask** (Greater Mageblood, Greater Cleric's, Elixir of Sages, Flask of Distilled Wisdom). No client row decides whether they stack. | aud-paladin.md P5. | OPEN (a beta buff bar; round 9 auras C5) | - |

## Before/after

Before: the shared baseline (origin/main bb619a188 on fork 4d3129cae). After: lane/r10-healers HEAD with wt/r10-healers. Published entries, raid 60, heal profile, effective HPS, 10,000 iterations.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| Restoration druid | +6.5% | +6.4% | +6.5% | +6.4% | 779.2 → 830.0 |
| Restoration shaman | -0.4% | -0.5% | -0.5% | -0.5% | 838.0 → 834.0 |
| Holy priest | +0.0% | +0.1% | +0.1% | +0.0% | 693.4 → 693.8 |
| Discipline priest | +0.1% | +0.1% | +0.1% | +0.1% | 624.9 → 625.3 |
| Holy paladin | 0.0% | 0.0% | 0.0% | 0.0% | 641.4 → 641.4 |

Where it comes from:
- **Druid:** Swiftmend kept its effect, +2.5% (778.7 → 798.4 at 3,000 iterations). The Swiftmend line under 75% adds +4.7%. Gift of Nature off Tranquility takes back about 1%.
- **Shaman:** Healing Stream's crit falls from 22.5% to 13%.
- **Priests:** Inner Focus off Renew and Mental Agility on Inner Fire are noise-level. The small gain is Inner Fire's cost.
- **Paladin:** nothing moves until the libram is equipped. With it, +0.3% (horde 641.4 → 643.3, alliance 653.6 → 655.2, seed 1). Totem of Urgency on the shaman likewise adds +0.2% (838.0 → 839.9).

Bare 60 (horde, seed 1, 3,000 iterations): druid 386.1 → 412.3 (+6.8%), shaman 375.0 → 372.5 (-0.7%), Holy priest 312.2 → 312.9, Discipline 294.7 → 295.2, paladin 299.0 → 299.0.

No DPS spec moves:
- Nature's Grace's heal half and Gift of Nature are druid-heal only.
- The Inner Fire and Inner Focus changes left Shadow's goldens unchanged.
- Totem of Urgency and the Holy Shock heal mask reach heals only.

## For other lanes

- **Crit lane:**
  - Healing Stream rolls stats.Crit only (A7, E 30003fb0d). The fix subtracts the shaman's SpellCrit around the tick's roll. If the crit lane adds a per-spell crit-table switch, this should move onto it.
  - Your "Holy/Nature spells reach heals" change (b4546cdb6) and this lane's changes both touch heal crit. Re-measure the healers with both merged.
- **Ranker (held refresh):**
  - Totem of Urgency (S2) and Libram of Infusion (H1) now have value. The Restoration shaman and Holy paladin totem and libram slots should fill at the refresh.
  - Holy paladin raid 60 lists the libram slot as "no_sourced_item" while bare lists the libram as an alternative. Check why the raid pass drops the Enchanting source.
  - Glowing Brightwood Staff's site item row omits its 150 healing and 50 spell power, while the sim's own row is right (aud-shaman.md, display only).
- **Ranks and healprofile lanes:**
  - Druid D4: one druid's heal-over-time ranks learned below 20 show no sub-20 coefficient penalty, which would move round 9's sub-20 downrank candidates.
  - The Swiftmend threshold (A9) is within 0.4% at bare 40 and 50 on today's profile. Re-measure it on the rebuilt sub-60 profile.
- **Integrator:** TestRotationLadder rows for Restoration druid and shaman, Holy paladin and both priests move (A7 to A11, H1 only when equipped). Not regenerated here.

## Still open

- Holy Power's crit-heal size (A6), Oath of Light (A12), Illumination on Holy Shock rank 1 (H3): each needs a level-40+ paladin log.
- Same-op percent stacking (druid D3, shaman S4), Prayer of Mending's talents (P3), Lightwell Renew's crit, Renewed Hope per pulse, Sudden Insight per target (P4).
- The elixir stack (X2); Concentration Aura (owner); Discipline's Inspiration (held).
