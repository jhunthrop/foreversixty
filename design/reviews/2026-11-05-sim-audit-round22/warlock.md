# Sim-accuracy round 22: warlock (2026-11-05)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70790 with its live hotfix push. Site `wt/r22-warlock` from int/r22 (ebeec2e0b, which carries the 3% Leader of the Pack / Moonkin crit change); fork `lane/r22-warlock` from int/r22 (4baa020f2). Baseline: spec-breakdown built from this branch before any change (`$R/warlock/sb-base`, `out/base60.tsv`), not `$R/before.tsv` (the crit change already moved every warlock raid-60 number by about -2%). Raid 60 runs use `-bis-dir $R/bis-frozen` and the guide raid build (`$R/rt`). Leveling runs use the round-21 helpers (`tools/lvl.py`, `tools/lrun.sh`: bare band gear, the shipped order's talents at that exact level). The lane took main's round 23/24 warlock wording (both curated APLs' notes and the four guide pages) as the base before editing, so the merge with main is clean.

**Totals: 2 CONFIRMED (both `old`; one `guides-handoff`), 1 lead not adopted, every other seed and the Blizzard notes aligned.** Raid 60 moves only for Affliction (+1.2 to +2.1%).

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Affliction raid build: Demonic Sacrifice (Imp), Nightfall and Improved Shadow Bolt beat the Succubus build** (MythicSim lead, +2.7% claimed). <br>**What was wrong:** the raid build kept the Succubus out (Unholy Power 5) and spent its Destruction points on Bane and Aftermath for Immolate. Sacrificing the Imp before the pull gives Burning Shadow (client 18789: aura 79, 15, school mask 32 = Shadow damage +15%), which is worth more to a spec that casts almost only Shadow than the Succubus's melee; Nightfall (18094, 2% a point on Corruption, Drain Soul, Drain Life and Wrack ticks; Shadow Trance 17941: aura 108 -100% cast time on mask 0x1 = Shadow Bolt, 10 s) turns into instant Shadow Bolts whose crits put up Improved Shadow Bolt (17794, 20% Shadow damage taken, 12 s) for Wrack and the DoTs. Without Bane, Immolate is a 2 s Fire cast Burning Shadow does not touch, and dropping it measured better. <br>**Verified:** the competitor's own build (25500020135201351-0050203001-05) drops Grim Reach, an unmodeled talent, so it fails the adoption rule as is. The adopted build keeps Grim Reach and gives up Improved Life Tap instead: **05520020135201351-0050203001-05** (Affliction 35: Suppression 5, Improved Corruption 5, Grim Reach 2, Improved Bane of Agony 2, Amplify Curse, Pandemic 3, Malevolence 5, Nightfall 2, Siphon Life, Soul Siphon 3, Shadow Mastery 5, Wrack; Demonology 11: Demonic Embrace 5, Demonic Aegis 2, Fel Vitality 3, Demonic Sacrifice; Destruction 5: Improved Shadow Bolt 5). Every raid keep (Amplify Curse, Siphon Life, Suppression, Demonic Aegis) and every unmodeled talent stays; nothing the rotation casts is dropped. Screens at 3,000 iterations: competitor build 745.4, adopted 745.3, Grim Reach kept with Improved Shadow Bolt 3 740.8, Suppression 3 737.7, no Nightfall 723.5, old build without Immolate 718.7 (Immolate stays right for the Succubus build). <br>**10,000 iterations, seeds 1 / 50001:** horde 735.5 / 735.6 to 745.0 / 744.7 (+1.29 / +1.24%), alliance 738.0 / 738.4 to 753.7 / 754.0 (+2.13 / +2.11%). Horde gains less because Orc's Command (5% pet damage) has no pet left. <br>**Fix:** guide raidBuild; card: pre-pull Summon Imp and Demonic Sacrifice gated on knowing the talent, Immolate gated on not knowing it, and a Shadow Bolt row on Shadow Trance gated on knowing Nightfall, ahead of Wrack. The talent gates needed the engine to expose both talents as known auras (E 169fe9493 Demonic Sacrifice 18788, E 6ef87d217 Nightfall 18094; the guide card settles a gate only on a talent name). The leveling build and the old raid build cast exactly as before (bare leveling 30/50/59 unchanged to the decimal; raid 60 old build 735.5). Tests: `TestWarlockAfflictionRaidSacrificesTheImpAndSpendsShadowTrance` (red: 0 sacrifices, 11.8 Immolates, 0 Shadow Bolts; then green), `TestDemonicSacrificeTalentIsAKnownAura`, `TestNightfallTalentIsAKnownAura` (both red, then green). | `out/a3_*`, `out/a10_*`, `out/chk3_*`; wago SpellEffect/SpellMisc/SpellName 70790 (layered), SpellDuration index 1 = 10 s | CONFIRMED (talent build, adoption rule) | old (every raid-build search ran with the Succubus card, which never sacrificed or spent Shadow Trance, so it could not find this) | S 171da0642, S 2b87bd8ed, S 90b39b26c, S e34b46508; E 169fe9493, E 6ef87d217, copies E 5e5216837, E 185f9631c, E aea07f64b |
| S2 | **Demonology lead (MythicSim +3.5%): measured, not adopted.** Their 011-0025003231020301351-0550005003 on our card: 728.8 against 706.2 (3k). The gain is Agonizing Flames 3 (17927 effect 1: +10% damage on mask 0x1a5 / 0x8000c0, which includes Shadow Bolt, bit 0x1); Demonic Brand does nothing on a card with no Searing Pain. A version that keeps Improved Health Funnel and Demonic Aegis (unmodeled): 011-2035103021220001351-0550005003, +3.7 / +3.6% horde, +4.0 / +3.9% alliance (10k, both seeds). It must cut Suppression from 5 to 1 (Pact needs 31 Demonology points, Agonizing Flames 18 Destruction), and `design/reviews/talent-search/raid-keeps.json` protects Suppression for Demonology as threat reduction (20% less threat on the highest-threat warlock spec, Shadow Bolt crits through Ruin). Decision: not adopted; Suppression stays. Every legal build with Suppression 5 leaves no room for Agonizing Flames; Master Summoner 2 into Improved Sayaad and Demonic Embrace loses (701.8). | `out/d3_*`, `out/d10_*` | Lead, not adopted (keep) | - | - |
| S3 | **No warlock rank row relied on a skip-below-20 rule** (platform P1, priest P1). The engine's level penalty has had no sub-20 cut since round 15, and round 15's rank sweep put Shadow Bolt rank 3 (learned at 12) in the race directly (W2, W3: the current 20-21 and 20-23 rows). That sweep tried only Shadow Bolt, so this lane screened the sub-20 DoT ranks every spec still holds through its 20s, leveling order, 3,000 iterations: Corruption rank 2 (14) at 24 to 30, Immolate rank 2 (10) at 20 to 28, Bane of Agony rank 2 (18) at 28 and 30. Every one loses, by 0.9 to 4.9% (`out/rankscreen.tsv`). No row. | `out/rankscreen.tsv` | Aligned | - | - |
| S4 | **R23-4: Demonology raid build, Soul Fire inside Decimation.** <br>**What was wrong:** the card skipped Soul Fire at every level of the raid build. Raid order, band gear, 10,000 iterations: rank 1 loses at 48 (-0.7%), 52 (-1.1%) and 55 (-0.3%); rank 2 (56) wins at 56 (+0.17 / +0.23% horde, +0.5% alliance), 57 (+0.16%), 58 (+0.13%) and 59 (+0.19 / +0.25%, alliance +0.44%); 60 bare ties (404.8 / 404.6 both ways); raid 60 loses 1.1% (705.6 to 697.6). So the note's reason holds at raid power and fails on band gear from 56. <br>**Fix:** the row fires with Pact known only at levels 56 to 59 (`currentLevel`, as the druid cards use); 60 and the leveling build are unchanged. The note and demonology.md :84 and :98 scope the skip to 60 and say to cast them from 56 if you play the raid build before 60. `TestWarlockDemonologyRaidSoulFireInsideDecimationAt56To59` was red at 57, then green. `spellIsKnown(17924)` could not gate this: it reads true from rank 1 at 48 (the site rewrites ids to the learned rank). | `out/r60_*`; lrun raid order 48-60 | CONFIRMED (small) | old, guides-handoff (guide r23 WARLOCK-R23-4) | S 4a1b2e2d8, E 7775a05ef |
| S5 | R23-3 (Gnome Demonology Eureka! on cooldown): the guide's arithmetic decides it and the card already presses it on cooldown (`autocastOtherCooldowns`); no sim needed. | guide r23 | Aligned | - | - |
| S6 | R24 optional sims: the Destruction raid order at 36 (Fire and Brimstone 3, no Shadow and Flame) is the one level where the held Conflagrate loses (round 21 S1: -0.3%). The card cannot read a talent rank, and on the leveling order the same talents (38) tie, so it stays held. The (a)/(b) Shadow and Flame decomposition was not run (optional; R24-2 is decided on arithmetic). | round 21 `out/gated-10k.tsv` | Aligned (one level, -0.3%, not gateable) | - | - |
| S7 | The Demonology Horde race re-run: pending owner (not reopened). | brief | Pending owner | - | - |
| S8 | Round 21's open: Demonic Brand's damage is still one source (talent text and its Wowhead tooltip); no build or card takes it (S2 confirms it does nothing on the card). No log yet (see Logs). | | OPEN | - | - |

## Blizzard's Forever notes (24 Sep to 4 Nov), every warlock line

Read from the cached notes thread (`us.forums.blizzard.com/en/wow/t/2360696`, as harvested in `$R/hunter/web/notes-all.txt`):
- Life Tap tooltip shows life converted, scales with Spirit: engine amount matches the tooltips (round 21 A3). Aligned.
- Voidwalker Sacrifice scales with 10% spell healing; Hellfire can crit; Soul Harvest renamed, 50/100% regen; Incubus model; pet stunned while summoning; Pets 2 Health per Stamina; Summon Incubus shard fix; Curse of Recklessness stacking fix; Soul Link castable: no card casts these or they do not touch damage. Aligned.
- Shadow Bolt ranks 1 to 4 rebalanced: spellconst 70790 matches (conformance golden, round 21 A3). Aligned.
- Demonic Brand: high threat only on Shadow damage: threat only; no card uses it.
- Grim Reach replaces Malediction; Shadow Mastery 2% a point: engine matches (TALENTS.md).
- Spellstones only Shadow damage: `TestSpellstoneIsShadowDamageAndCastSpeed`. Aligned.
- Bane of Agony updates dynamically (`dynamic_dots_test.go`); Amplify Curse affects Exhaustion and Agony (round 21 A2); Improved Drains no longer reaches Wrack (engine excludes it); Wrack coefficient 0.286 and 2.5 a level, 86 at 60 (`wrack.go`); Drain Soul shard chance cut (no shard model, round 21 S3). Aligned.
- Shadow and Flame 4% a point (was 2%): `shadowAndFlameDamagePerRank = 0.04`. Aligned.
- Command re-added (5% pet damage): `racials.go` Command. Aligned (and why horde gains less in S1).

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | S1 above. | | | | |
| A2 | Re-audit of what S1 now casts: Nightfall procs on every Corruption and Wrack tick (`OnPeriodicDamageDealt`, 4%), Shadow Trance 10 s and ends on the instant cast (17941 duration index 1); Burning Shadow lasts the fight and any summon cancels it (no Pact); Improved Shadow Bolt 20% for 12 s with no charges (17794). About 6.4 instant bolts a 180 s fight, against about 7.4 procs from roughly 185 ticks. DoT applications roll spell hit (`OutcomeMagicHitNoHitCounter`), so the 0% miss in a breakdown is the ticks only. | engine read; `out/chk3_d.json` | Aligned | - | - |
| A3 | Races on the new build (10k, seeds 1 / 50001): Human 756.7 / 757.0, Gnome 753.7 / 754.0, Orc 745.0 / 744.7, Undead 744.5 / 745.2, Troll 737.3 / 736.8. Human and Orc stay the picks (within 1% of the next race in each faction); Undead is now level with Orc, and Command does nothing once the Imp is gone. affliction.md Races says so. | `out/rc_*` | Aligned (prose updated) | - | S 171da0642 |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | S2, S4 above. Raid 60 unchanged (705.6 / 705.2 horde, 715.6 / 716.6 alliance). | `out/after60.tsv` | | | |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Round 21's Conflagrate hold re-read against guide round 24 (R24-2's arithmetic) and S6; the card's two rows behave as measured. Raid 60 unchanged (774.7 / 775.1 horde, 778.3 / 776.8 alliance). | `out/after60.tsv` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **An APL row on an aura the unit never registered runs unconditionally.** `newValueAuraIsActive` returns nil for an unknown aura and `and` drops nil values, so `auraIsActive(17941)` on a build without Nightfall made Shadow Bolt the filler (Affliction 735.8 to 604.6 in a probe). No shipped warlock row is exposed (Decimation is known wherever Soul Fire is; the new Shadow Trance row is also gated on knowing Nightfall). Shared code: listed for platform. | `out/a3_o_sbt.json`; `sim/core/apl_values_aura.go:36`, `apl_values_operators.go:581` | OPEN (platform) | old | - |

## Logs

Last looked 2026-11-05 09:36 UTC: `api.foreversixty.gg/v1/reports/recent` still lists only z2opt2x2gs5u (2026-10-20); foreverlogs.gg's report API answers "API access not allowed"; the Warcraft Logs key could not be read from this shell (gcloud returned nothing), so WCL was not checked. Still needed: 40+ Gnome Destruction (Eureka! on Incinerate), 40+ Affliction (Soul Siphon, Wrack, and now Nightfall from Wrack ticks), level 60 (Doom, demon crit), a level-50 Succubus, any branded demon (Demonic Brand).

## Before/after

Raid 60, guide raid builds, frozen BiS, 10,000 iterations, seeds 1 / 50001, against this branch's own baseline:

| spec | horde | alliance |
|---|---|---|
| Affliction | 735.5 / 735.6 to 745.0 / 744.7 (+1.29 / +1.24%) | 738.0 / 738.4 to 753.7 / 754.0 (+2.13 / +2.11%) |
| Demonology | 705.6 / 705.2 unchanged | 715.6 / 716.6 unchanged |
| Destruction | 774.7 / 775.1 unchanged | 778.3 / 776.8 unchanged |

Demonology raid order, band gear, horde seed 1: 56 302.8 to 303.3, 57 305.6 to 306.1, 58 314.8 to 315.2, 59 316.7 to 317.3; 48 to 55 and 60 unchanged. Leveling orders of all three specs unchanged.

## Files to regenerate

- addon-builds.json (Affliction raid order from the new raidBuild), addon-data.json and Data.lua: the Affliction card (new opener, Shadow Trance row, Immolate gate and note, round 23/24 notes from main) and the Demonology Soul Fire row and note.
- sim/request ladder goldens (TestRotationLadder): warlock-affliction only if a rung uses the raid order (rows are inert without the talents); warlock-demonology raid-order rungs at 56 to 59 if any.
- `sim/request/raid-consumes.json` and `published-presets.json` (`go run ./internal/genraidconsumes` from sim/): Affliction's raid preset no longer carries Elixir of Fire Power (S e34b46508). **`TestRaidConsumesJSONIsCommitted` is red on this branch until that regen**; also `data/curated/simbuffs.json` if it mirrors the preset.
- Coach starters and apl-sync output that copy the warlock APLs (the `sim/request/apl` and fork `ui/warlock/apls` copies were synced by hand).
- Fork: no golden regenerated. `TALENTS.md` (hand-written) still says Nightfall's Drain Soul and Wrack procs are not wired; they are (`nightfallTriggerSpellCodes`), worth a one-line edit at the next refresh.

## Measure at refresh

- Affliction band 60 BiS and its race picks on the new raid build (Human first; Undead and Orc level for Horde); the published BiS talents are still the stale pre-Wrack ones.
- Carried: the Affliction bare 40, 50 and 60 BiS entries should carry the Malevolence string.

## For other lanes

- **Moves DPS outside warlock: nothing.** Engine changes are in `sim/warlock` only (two talent auras get ids; the Imp sacrifice aura's label is now its Forever name, Burning Shadow).
- **Platform:** C1 (a row on an unregistered aura runs unconditionally). A validation warning exists; the card pipeline could reject such rows, or `auraIsActive` could return a false constant instead of nil.
- **Coordinator:** the Demonology Agonizing Flames build (+3.6 to +4.0%) is blocked only by the Suppression keep in raid-keeps.json; if that keep is ever relaxed for Demonology, the build is 011-2035103021220001351-0550005003.

## For the guide lane

- Done here: affliction.md raid build, talents, opener, rotation, consumables (no Fire elixir for the sacrifice build) and races; demonology.md :84 and :98 (Soul Fire skipped at 60, cast from 56 on the raid build). Main's round 23/24 wording was taken first.

## Guide text not done

- None.

## Guide review

- wow-player, on the rendered Affliction (1440, 390) and Demonology (1440) pages of this branch's dev build, with the new card notes as text (the card itself renders after the regen). First pass: Affliction NO-SHIP on two MUSTs (the consumables table listing Elixir of Fire Power; the pre-40 Shadow Bolt filler row showing on the 60 card beside the Shadow Trance bolt), plus wording SHOULDs; Demonology SHIP with one SHOULD (sentence order). All taken (S e34b46508: Fire Power out of the Affliction raid preset, the filler row gated on not knowing Wrack, the wording). Kept with reasons: Burning Shadow's "two hours" (Forever's 18788 text says 2 hrs), Soul Fire's 2.4 s (Bane 5 by 56 on the raid order), no Undead badge (within 1%). **Second pass: Affliction SHIP, conditional on the regenerated consumables table and raid card dropping the Fire Power row and the filler Shadow Bolt row (integrator: check both after the regen); Demonology SHIP.** `_prose.test.ts` and `builds.test.ts` pass.

## Still open

- C1 (platform). S8 and the log list above.
- Pending owner, not reopened: level-60 trainer book ranks (Shadow Bolt 25307, Corruption 25311, Immolate 25309 stay off the cards), Siphon Life tick crit, the Demonology Horde race, Heart of the Lion 280/276, warlock Orc.

## Sources

- Client 1.60.1.70790 (wago.tools SpellEffect, SpellMisc, SpellName layered with the live push via round 21's `layer.py`; SpellDuration fetched 2026-11-05); `data/builds/1.60.1.70790/` talents and trainables.
- Blizzard, WoW Forever development notes thread, https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696 (posts 24 Sep to 4 Nov, cached copy).
- MythicSim comparison: `scratchpad/competitor/lanes/casters.md` (verified here, not taken as given).
- Guide-accuracy rounds 22 to 24 warlock reports (origin/main).
