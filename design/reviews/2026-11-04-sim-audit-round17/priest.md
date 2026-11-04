# Sim-accuracy round 17: priest

This is lane priest. It works on site wt/r17-priest and fork lane/r17-priest, both cut from int/r17 (site d7fc51fe7, fork f637ec32c). The client is 1.60.1.70755.

Specs: Discipline, Holy and Shadow, plus the shared low-rank penalty (core.LevelPenalty, MaxLevel + 12).

Every sim ran through the lock wrapper on frozen BiS (`$R/bis-frozen`), with the guide raid build at raid 60. Exploration runs used 3,000 iterations. Reported numbers use 10,000 iterations at seeds 1 and 50001.

## 1. The low-rank penalty: no new source settles anything

Forever has been live for a day. I looked again everywhere the brief names.

| Source | What it holds now (checked 2026-11-04) |
|---|---|
| foreverlogs.gg API (`/api/reports/<id>`) | One report since beta: **2814**, "Test", 2026-11-04 12:24 UTC, with **0 encounters**. Ids 2815 to 2849, and spot checks at 3000, 4000, 6000 and 10000, return "Report not found". The list endpoint needs a token. |
| api.foreversixty.gg/v1/reports/recent | Still one report, from 2026-10-20 (42 fights, beta). |
| HealRankAdvisor (github.com/mendeleevforever/HealRankAdvisor) | The last commit is 2026-10-24 ("Create release.yml"). It has no issues. It holds the same level-30 measurement round 16 used. |
| ForeverChanges /downranking | It is unchanged (still the beta table). It says plainly: "If that rule holds above level 30 ...". It holds no damage measurements. |
| foreverupgrades.com (11 Oct), foreverwisp, Ozgar's tool | These are claims only. Damage spells are "unverified", and nothing is measured above 30. Ozgar implies a top-rank exemption but cites no measurement. |
| r/wowforever | **Blocked from here**: reddit.com and old.reddit.com redirect or refuse. A web search found no thread with numbers. |
| Warcraft Logs | Still blocked (expired login). Recorded, not worked around. |

**Verdict:**
- **OPEN, unchanged:** the top-rank exemption, damage spells, levels above 30, and the Healing Wave rank 1 curve.
- **No code change.** The owner's in-game test in round 16 (`2026-11-04-sim-audit-round16/priest.md`, section 1) still settles all of them at once. Lesser Heal rank 3 against Greater Heal rank 4 at 60 is the cleanest single read for the exemption. Mind Flay rank 3 against Mind Blast rank 9 is the one for damage.
- **No other class's rank rows move:** the penalty did not change.

## 2. Round 16 regression targets

| # | target | check | result |
|---|---|---|---|
| R1 | Per-level Renew rows (Holy 14-32 rank 1, 33-39 rank 3; Discipline 14-29 / 30-33 / 34-39 ranks 1 / 2 / 3) | `spec-breakdown -level` at the switch points, horde, seed 1, through each rank's own id | Each number reproduces round 16 exactly:<br>- **Holy:** at 31, rank 1 68.8 against rank 3 67.2. At 33, rank 3 69.5 against rank 1 67.7.<br>- **Discipline:** at 29, rank 1 44.6 against rank 2 44.2. At 30, rank 2 90.0 against rank 3 87.5. At 34, ranks 2 and 3 tie at 89.6.<br>`TestPriestRenewRowsSwitchInsideBandThirty` passes. **Aligned.** |
| R2 | Devouring Plague rows (rank 1 to 51, rank 2 from 52) | the same method, at 51 and 52 | At 51, rank 1 243.3 against rank 2 243.6 (a tie). At 52, rank 2 251.5 against rank 1 247.5. `TestShadowDevouringPlagueMovesToRankTwoAtFiftyTwo` passes. **Aligned.** |
| R3 | Wand Specialization first | `data/curated/ladder-first.json` priest-shadow | Wand Specialization 2/2 is first, then Spirit Tap. **Aligned.** |
| R4 | Power Word: Shield rank 10 | `priest-discipline.json` line 7 is 10901 rank 10, in the site copy and the fork copy | At raid 60 it is cast 5.6 times a fight, all at rank 10. **Aligned.** |

- **Band 10 (levels 14-19):** the Renew rows there cannot be simmed. `spec-breakdown` has no band-10 entry ("no priest-holy band 10 bare horde entry"). They are left as round 16 set them.
- **Fork test:** `TestHealerFiresASpellPowerTrinket` (`sim/priest/healing/cooldowns_test.go`) panics on "No item with id: 18820" on the round-start engine too. This is the same missing-test-item family as round 16's open item 8. Every other priest fork test passes.

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **The tank gets no Renew before the pull.** The guide said "Renew the tank as he engages". The list's only pre-pull cast was Lightwell, so the tank's first Renew went out on the first global after the pull.<br>Round 15 measured a pre-pull Renew (its D2). It held the line only because band 20 lost (-0.3% at 3,000 iterations). On the round-17 engine:<br>- **Band 20** is unchanged to 0.1 HPS.<br>- **Every other band gains,** at both seeds and both factions: raid 60 +0.24 to +0.36%, bare 60 +0.85 to +1.01%, bare 30 to 50 +0.14 to +0.67% (table below).<br>A pre-HoT on the tank as he pulls is ordinary Classic healing practice.<br>**The fix:** Renew rank 7 (10927, the rank the tank Renew line casts) on the tank (player 5) at -1 s, after Lightwell. The guide's "Before the pull" now says "a second before he engages". | sims below; `TestDisciplineRenewsTheTankBeforeThePull` (sim/request) failed, then passed; `make apl-check` passes | CONFIRMED (rotation) | old (round 15 D2 held it; the band-20 loss is gone) | S 47e6058a6, E 5dd716a48 |
| D2 | **Fresh re-audit, raid 60** (horde, guide build `00500503133001143-3350500303-0`):<br>- Penance 16.9 casts a fight, the best direct healing per mana (5.98).<br>- Renew 6.58 healed per mana. Power Word: Shield is cast at rank 10.<br>- Inner Focus 2.0 casts in 300 s (180 s cooldown).<br>- Major Mana Potion and Demonic Rune 3.0 each.<br>- Lightwell Renew ticks crit through the client PERIODIC_CAN_CRIT flag (7001, 27873, 27874 all carry it).<br>- The build is legal and has no dead points. Renewed Hope is 4/5 (8%, client rank 4). | spec-breakdown; `talents/priest.json` rank tables; `periodic_crit_auto_gen.go` | Aligned | - | - |
| D3 | **Lightwell Renew ticks crit at 14.6%,** against Renew's 19.5%. Holy Specialization (5%) does not reach the well's renew, because it has no class mask in the engine. In the game the well's renew is applied for the clicker, so whose crit it uses is not knowable from the client rows. The effect is 16 ticks of 320 a fight. | `sim/priest/lightwell.go` | OPEN (negligible) | - | - |
| D4 | **The list's top-level `notes` field is stale (internal, not rendered).** It still says rank-choices names "Power Word: Shield at rank 4", and that the tank's Renew is the top trainer rank. No Power Word: Shield row exists, and the list casts Renew rank 7. Nothing reads this field: `tools/apl_paths.py` DISPLAY_KEYS, and no page renders it. | `data/curated/apl/priest-discipline.json` notes | Note | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The same pre-pull Renew does nothing for Holy:** raid 60 793.0 to 792.2 (-0.1%), bare 60 403.1 to 403.7 (+0.15%), 3,000 iterations. Round 15 found the same. Holy keeps its list. | sims | Aligned (not adopted) | - | - |
| H2 | **Fresh re-audit, raid 60** (guide build `00500503103-33505003030121431-0`):<br>- Prayer of Mending 24.4 casts and 79.3 heals a fight, the best healing per mana (10.56).<br>- Renew ticks crit at 19.3%.<br>- Litany of Light gives 11.4 mana a second.<br>- Holy Binding fires off every Flash Heal (26.6 against 26.6).<br>- Inner Focus 2.0 casts.<br>- Spiritual Guidance 4/5 is 20% of Spirit as healing and 6% as damage, matching the client's non-linear rank table (`healing_talents.go:19`, `talents.go:78`). | spec-breakdown; `talents/priest.json` | Aligned | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **The Shadow Word: Pain cutoff** (APL `remainingTime >= 10`, against the page's "6 seconds"; open since guide-accuracy round 9). 10,000 iterations, horde:<br>- 10 s: 735.5 / 735.0.<br>- 6 s: 735.6 / 735.0.<br>- 7.5 s: 735.7 / 735.1.<br>It is a tie inside error at both seeds. The list stays as it is. The page's 6 seconds is as right as the list's 10 (a refresh with two ticks left still beats a Mind Flay global), so the prose stays too. | sims | Aligned (tie) | - | - |
| S2 | **Eureka! masks** (guide-accuracy S-R9-1). The engine follows the client row 1259823's three class masks, `sim/core/eureka_auto_gen.go`: damage/healing, periodic damage (it names Mind Flay and Penance) and cost. The tooltip says "non-periodic". The client table outranks the tooltip text, and only a Gnome Shadow priest's log can show whether Mind Flay ticks gain. The published races are Undead and Human, so no published number moves. | client SpellEffect rows; Wowhead 1259823 | OPEN (log) | - | - |
| S3 | **Fresh re-audit, raid 60** (guide build `00510503103-0-50052201211312251`):<br>- Mind Flay 39%, Shadow Word: Pain 20%, Mind Blast 19%, Devouring Plague 12%, Shadow Word: Death 9%. No time out of mana.<br>- Mind Flay shows 3.59 "hits" a cast: the landing outcome plus three ticks, less misses. That is a display count, not a fourth tick (`MindFlayTicks = 3`, 1 s ticks).<br>- The Improved Mind Flay ramp compounds 1.2 a tick at 2/2, as the client text says ("20% more than its previous instance").<br>- Shadowform is +10% damage, -50% cost and +100% crit bonus, matching the client.<br>- Inner Focus is consumed by the sequenced Mind Blast (1.35 casts, a 1.5 s aura).<br>- The build has no dead points: Silent Resolve 1/3 cuts threat, and Silence and Shadow Reach are utility a raid uses. | spec-breakdown; `talents/priest.json`; `sim/priest/{mind_flay,talents}.go` | Aligned | - | - |
| S4 | **Bands 20 to 50** (bare, horde): the cast ranks match the rows and the trainer levels.<br>- Shadow Word: Pain rank 3 at 30 and rank 5 at 40.<br>- Mind Blast rank 2, 4, 4 (the Undead row) and 7.<br>- Mind Flay ranks 2 to 4.<br>- Devouring Plague rank 1 at 50 (the Undead row).<br>- Shadow Word: Death 1309633, which is rank 2 (learned at 40), at 40 and 50.<br>**But** the frozen BiS bands carry the pre-round-15 ladder talents (`542...`, with Shadowform at band 40). Band numbers simmed from the published entries therefore describe the old ladder until the BiS refresh rewrites them. | `bis-frozen/priest-shadow.json` (engine d1d9f2775) | Note (refresh) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The low-rank penalty** (section 1): there is no new evidence and no change. | see section 1 | OPEN | - | - |
| X2 | **70755 late re-layer (43dc0f074):** for priest it changes only item and set text (whitespace, a resolved `$a1`). No priest spellconst, talent or trainable row moved. Priest conformance has no "differs" row. | `git show 43dc0f074`; `sim/core/testdata/conformance/priest.golden.md` | Aligned | - | - |
| X3 | **Talent rank values** for every point the three builds take, checked against the client's per-rank text (non-linear ones: Mental Agility 3/7/10, Spiritual Guidance, Improved Power Word: Shield 7/14/20, Wand Specialization 13/25). Each is pinned by an existing fork test (`talent_rates_test.go`, `healing/*_test.go`). | `talents/priest.json` | Aligned | - | - |

## Before/after

Before = `$R/before.tsv` (raid 60) and `$R/sb-before` from `$R/site-before/sim` (bands). After = this branch. 10,000 iterations, frozen BiS, guide builds at raid 60.

**Raid 60**

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| priest-discipline | 504.6 → 506.2 (+0.32%) | +0.28% | 494.5 → 495.7 (+0.24%) | +0.36% |
| priest-holy | 773.8 (0.00%) | 0.00% | 793.5 (0.00%) | 0.00% |
| priest-shadow | 742.6 (0.00%) | 0.00% | 735.5 (0.00%) | 0.00% |

**Discipline, bare bands**

| band | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| 20 | 43.2 (0.00%) | 0.00% | 41.7 (0.00%) | 0.00% |
| 30 | 89.4 → 89.8 (+0.45%) | +0.45% | 90.0 → 90.6 (+0.67%) | +0.67% |
| 40 | 122.4 → 122.8 (+0.33%) | +0.33% | 128.0 → 128.3 (+0.23%) | +0.31% |
| 50 | 134.7 → 135.3 (+0.45%) | +0.45% | 144.0 → 144.2 (+0.14%) | +0.14% |
| 60 | 200.1 → 201.8 (+0.85%) | +1.00% | 198.8 → 200.8 (+1.01%) | +1.01% |

No other class moves (site APL and fork APL copy only).

## Files to regenerate

- `data/builds/1.60.1.70755/addon-data.json`, then `addon/ForeverSixty/Data.lua` (the Discipline pre-pull line).
- TestRotationLadder goldens for priest-discipline, bands 30 to 60.
- dps-per-point for priest-discipline.

## Measure at refresh

- priest-discipline, every band: the pre-pull Renew.
- priest-shadow bands: the published band talents are the pre-round-15 ladder until the BiS refresh.

## For other lanes

- None from this lane. Nothing shared changed.
- Fork test data: `TestHealerFiresASpellPowerTrinket` needs item 18820. It belongs with round 16 open item 8 (missing test items).

## For the guide lane

Done in this lane: discipline.md:64 ("Renew the tank a second before he engages, so it is already ticking when the first hits land"). The opener card renders "Lightwell, 2.5 s before the pull" and "Renew, 1 s before the pull". The wow-player verdict is below.

## Still open

- **The low-rank penalty:** the top-rank exemption, damage spells, levels above 30, and Healing Wave rank 1. The owner's in-game test in the round-16 report settles them.
- **Eureka!'s periodic mask** against the tooltip (needs a Gnome Shadow log).
- **Lightwell Renew's crit source.**
- **Logs:** foreverlogs.gg has one empty launch report. r/wowforever and Warcraft Logs are unreachable from here.

## Sources

- foreverlogs.gg API, reports 2813 to 2849, fetched 2026-11-04.
- api.foreversixty.gg/v1/reports/recent, fetched 2026-11-04.
- HealRankAdvisor: https://github.com/mendeleevforever/HealRankAdvisor (commits API, last commit 2026-10-24).
- ForeverChanges: https://foreverchanges.pro/downranking (fetched 2026-11-04, beta table unchanged).
- Forever Upgrades, 2026-10-11: https://foreverupgrades.com/news-article.html?a=downranking-nerfed-less-spell-power-fewer-procs
- Ozgar: https://wowdownrank.forever.ozgar.se/

## wow-player verdict

**SHIP.** The reviewer looked at the rendered Discipline "Rotation and priority" section (dev server, 1440 wide).

- The card and the sentence agree: Lightwell at 2.5 s, then Renew rank 7 at 1 s. They are consistent with the "Keep Renew on the tank all fight" row.
- A pre-pull Renew is normal practice.
- One SHOULD, applied: "so its first tick lands on the first hits he takes" was loose, because the first tick comes 2 s into the fight. It now reads "so it is already ticking when the first hits land".
