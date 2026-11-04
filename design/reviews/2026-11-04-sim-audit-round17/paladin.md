# Sim-accuracy round 17: paladin

Lane paladin. Site wt/r17-paladin from int/r17 (d7fc51fe7), fork lane/r17-paladin from int/r17 (f637ec32c). Client 1.60.1.70755 with the late re-layer (43dc0f074). Every sim and sim-running test went through the lock wrapper.

**Result:** zero CONFIRMED-BUGs. The top item (Oath of Light and the low-rank penalty) cannot be settled from the client, the web or logs, so it stays OPEN. The current Holy Light rows are kept. The owner's in-game check is below; it settles both this question and a second open assumption in the same code, the Oath coefficient. No engine, APL, build string, rank row or guide text changed. The only commit is this report.

## Top item: Oath of Light and the low-rank penalty

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | **Is Oath of Light's bonus cut on low ranks?** The engine adds `coefficient × Oath amount` by hand (healing.go:92, holy_shock.go:158). Core's level penalty never reaches that term. | **Client.** 1328548 (wago SpellEffect 70755, rows 1367618-20) has three effects: aura 226 (periodic dummy, 1.5 s) and aura 4 (dummy) twice. Each has 0 base points, 6.6667 / 1.9167 / 3.8333 per level, and **no EffectBonusCoefficient**. A dummy aura is applied by server script, so the client does not say how the amount scales. The tooltip says "by up to 0" for all three (Wowhead Forever tooltip; spelltips): the per-level amount is not rendered. Blessing of Light is gone from the Forever client. In Era 1.15.9.70576, 19977/19979 are the same structure (aura 4 dummies, base points 209/59 and 399/114). Forever also rebuilt Holy Light. Era 635 is effect 77 (script) with no coefficient. Forever 635 is effect 10 (heal) with BonusCoefficient 0.714. So Era's Blessing of Light code path does not carry over.<br>**History (Blessing of Light, the Oath's model).** Two sources say it added the **full** amount, not scaled by the cast-time coefficient: warcraft.wiki.gg/wiki/Blessing_of_Light Notes, and theholylight.wordpress.com/2007/07/10/paladin-healing/ (TBC 2.1). Patch 2.3.0 says: "Lower ranks of Flash of Light and Holy Light are now properly penalized when used with this Blessing". So before 2.3 the level penalty did not reach it. The same 2007 blog says (single source) that the sub-20 penalty did cut the bonus at "Holy Light 4 and below". All read 2026-11-04.<br>**Forever.** No source mentions the Oath and the downranking penalty: web searches 2026-11-04 (Forever guides at mobalytics, wowtbc.gg, kami-labs; epiccarry downranking; ForeverChanges, already read in round 16).<br>**Logs.** api.foreversixty.gg/v1/reports/recent has only the 2026-10-20 report. On foreverlogs, 2814 (2026-11-04, "Test") is empty, and 2815 to 2819, 2820, 2900 and 3500 are not found. The 2,256 cached beta timelines (round 10) contain no 1328548 event. They predate the Oath (70707). | OPEN | - | - |
| O2 | **The rows hinge on the penalty, not on the coefficient.** I patched the engine for the experiment only, with four readings of the Oath term: A = coefficient (engine today), B = coefficient × penalty, C = full amount, D = full × penalty. Then I swept Holy Light ranks (bare, own band heal profile, 3,000 iterations, seed 1, alliance, no Holy rows). Effective HPS by Holy Light rank 1..top:<br>47: A 160.6 / 158.7 / 158.6 ... 158.2 (rank 1 best). B 119.8 / 140.3 / 154.1 / 154.2 ... (ranks 4+ tie). C rank 1 best. D ranks 3+ tie.<br>50: A 161.1 / **177.4** / 151.7 / 157.4 / 160.7 / 162.1 / 163.2. B 63.4 / 81.2 / 116.5 / 151.9 / 159.9 / 162.1 / **163.2**. C rank 2 best (219.4). D top best (184.4).<br>55: A rank 2 192.7 vs 159.9 top. B 92.8 / 106.0 / 127.5 / 159.7 / 160.0 ... (4+ tie). C rank 2. D 5+ tie.<br>25 and 29: A and B are identical (rank 2 best). At 25 to 29 the Oath term is not penalized enough to matter, so the 25-29 row holds either way. | The experiment patch was reverted (`git checkout sim/paladin/healing.go`). Scripts: scratchpad r17/paladin/sweep.py. | Aligned as measured: with the penalty (B or D), the rows at 47-49 and 50-59 are wrong by about 20% (47) to 50% (50-59). Without it (A or C), they are right. | - | - |
| O3 | **The risk is lopsided.** If the Oath is penalized, the current Holy Light rank 2 row at 50-59 costs about half the band's healing (50: 81.2 against 163.2). If it is not, removing the row costs about 8% (163.2 against 177.4). The brief says to keep the rows when the question cannot be settled, and I kept them. The owner's check below is about two minutes in game, and the change set is ready if it comes back "penalized". | O2 numbers | - | - | - |
| O4 | **Second open assumption in the same term: coefficient or full amount.** At raid 60 the penalty does not matter: every rank the list casts there is unpenalized. Holy alliance seed 1, 3,000 iterations, frozen BiS, guide raid build: A 671.7, B 671.7, **C 727.3, D 727.3 (+8.3%)**. Blessing of Light's history (two sources, O1) points to the full amount. The engine's coefficient comes from oath_of_light.go:19, which says "Assumption, as Blessing of Light's was". That assumption contradicts the wiki's note. The same check settles it. | spec-breakdown with the experiment patch | OPEN. It is not changed on two pre-Forever sources for a spell Forever rebuilt. | - | - |

### Owner's in-game check (settles O1 and O4 together)

On a Holy paladin with Oath of Light (any level from 25; the higher the level, the clearer):
1. Swear Oath of Light on a party member. If the game allows it, swearing on yourself works too.
2. Cast **Holy Light at your top rank** about five times on the Oath target and five times on someone without the Oath. Read the non-crit amounts from the combat log, counting overhealing as part of the heal, or let the targets take damage first. Ignore crits. The difference in the averages is the Oath term at full rank.
3. Repeat with **Holy Light rank 1**. Below level 43, use rank 2 too.
4. Gear and healing power do not matter, because the same rank is compared on two targets. Do not swap gear between steps 2 and 3. If you have Healing Light, divide the differences by 1.04 / 1.08 / 1.12 for 1 / 2 / 3 points.

Expected Oath term (before Healing Light), where L = your level:

| level | top rank: coefficient (A, B) | top rank: full (C, D) | rank 1: A | rank 1: B | rank 1: C | rank 1: D | rank 2: B | rank 2: D |
|---|---|---|---|---|---|---|---|---|
| 30 | 143 | 200 | 143 | 50 | 200 | 70 | 93 | 130 |
| 40 | 190 | 267 | 190 | 0 | 267 | 0 | 29 | 40 |
| 50 | 238 | 333 | 238 | 0 | 333 | 0 | 0 | 0 |
| 60 | 286 | 400 | 286 | 0 | 400 | 0 | 0 | 0 |

How to read the results:
- Top-rank difference near 0.714 × 6.667 × L means the engine's coefficient is right. Near 6.667 × L means the full amount (O4 confirmed).
- If the rank-1 difference equals the top-rank difference, the Oath is not penalized and the rows stand. If it is smaller, or about zero at level 37 and above, the Oath is penalized.

**If penalized (B or D), the change set is:**
- **Fork:** scale the Oath term by `core.SpellLevelPenalty(rank.spellID, rank.level, casterLevel)` in healing.go:92. Holy Shock has a single rank per band, so holy_shock.go:158 is the same idea. Write the test first: Holy Light rank 1 on the Oath target at 50 heals the same as off it.
- **Site:** remove the Holy Light rows at 47-49 and 50-59; the default top rank ties ranks 4+ within 0.1%. Preliminary picks with Holy Light at top rank, 3,000 iterations: tank Flash of Light rank 3 at 50 and 53 (176.2 against 163.2 for rank 4 at 50), and group Flash of Light rank 2 at 47 (+1%). These need the per-level 10,000-iteration confirmation at seeds 1 and 50001, both factions.
- **Text:** holy.md:78 and the card note at paladin-holy.json:584 change. Both say "the same flat amount to every rank".

**If full amount (C or D):** drop `heal.coefficient ×` from the Oath term, then re-sweep the 40-59 rows. Under C, rank 2 still wins at 50 and 55.

**Also note:** at 50-59 most of the gain from Holy Light rank 2 comes from the below-40% emergency line behind Divine Favor (line 3), not the spare-mana line. Holy Light at top rank on line 3 alone: 50: 177.4 to 164.2; 55: 192.7 to 162.8; 59: 202.8 to 167.9 (seeds 1/50001 agree). The heal profile has no tank-death cost, so a cheap rank on the "heal that must land" scores well in the model. If the Oath is not penalized, it is the owner's taste call whether to scope the 47-59 rows to line 4 (`lines`) for real-play safety. I did not change it.

## Hedge (coordinator's decision: option c, rows kept)

The coordinator asked for a hedge: keep the cheap Holy Light ranks on the spare-mana line (line 4) only, and cast the top rank on the below-40% Divine Favor heal (line 3). I measured it at 10,000 iterations, seeds 1 and 50001 (agreeing), on the own-band heal profile, alliance.

| level | current rows | line 3 at top rank | cost |
|---|---|---|---|
| 47 | 160.6 | 159.3 | -0.8% |
| 49 | 157.5 | 152.8 | -3% |
| 50 | 177.4 | 164.4 | -7.3% |
| 54 | 190.2 | 160.4 | -15.7% |
| 59 | 202.8 | 167.8 | -17% |

- **Line 4's cheap rank still wins.** It beats the top rank on line 4 beyond error at every level checked, both factions. For example, at 59: 167.8 / 167.9 against 164.5 / 164.5.
- **`lines` cannot scope only the 47-59 rows.** It scopes a whole spell entry, and the authored id must be unique per spec. Scoping Holy Light to [4] also takes the 20-24 rank-1 row off line 3. That costs -12% at 20 (48.8 to 42.8), -0.6% at 24, and ties at 25-29.
- **Decision (c):** the rows stay as modelled. They matter only from level 47, so they carry no day-one risk. The platform lane is building per-band `lines` this round, which makes the hedge or the fix a one-row edit later.
- Raw numbers are in scratchpad r17/paladin/hedge10k.txt.

**Owner check, for the post-launch list.** It needs a Holy paladin of level 25 or higher with Oath of Light, so it is not a day-one item. Run the check below and give the results to the paladin lane. Revisit when it is done, or as soon as launch logs show a low-rank Holy Light on an Oath target at level 40 or higher.

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Talent build** `50320215012101251/50322/0503` (decoded): Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1 / Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2 / Benediction 5, Conviction 3. Ten points do nothing for the heal rotation: Purifying Power 2, Divine Precision 1, Precision 3, Guardian's Favor 2, Sacred Duty 2. But every talent that does something for a healer is already maxed: Healing Light, Reverence, Illumination, Holy Power, Divine Intellect, Benediction (Holy Shock is instant), and Conviction (aura 290 since push 113139, which reaches heals; talents.go:69). So no legal replacement takes a point that does something. | talents/paladin.json 70755 descriptions; engine talents.go | Aligned (unavoidable filler) | - | - |
| H2 | Round-16 regression targets hold: `TestHolyPaladinFlashOfLightAndHolyLightKeepTheirHealingBonus`, `TestHolyPaladinHolyShockAtBareSixty`, the rank-choice suites (`TestReMeasuredRankChoicesCastTheWinningRank`, `TestRound12/13/14…`, `TestEmbeddedRankChoicesMatchCurated`, `TestCuratedRankChoiceLinesCastTheirAuthoredSpell`), and fork sim/paladin/holy. | go test -tags=with_db (sim/request -run paladin/rank filters; fork ./sim/paladin/...) | Aligned | - | - |
| H3 | Oath of Light's cast: 8% of base mana, GCD, 5 min, one target (SpellPower cost_pct 8, duration 300000, spellconst). It matches oath_of_light.go:24-29. Boundless Light reaches only allies without the Oath (holy_shock.go:160). | spellconst/paladin.json 1328548 | Aligned | - | - |
| H4 | Holy Light 1-9, Flash of Light 1-6, Holy Shock 1-4, Divine Favor, Illumination, Infusion of Light, Holy Power, Reverence, Healing Light and Light's Vigil were re-checked in round 16 (H2 there). No paladin SpellEffect, SpellPower, SpellCooldowns, SpellMisc or talent row moved since: the round-16 R2 diff of 43dc0f074, and no new 70755 push since (brief). The Holy Light (635) and Flash of Light (19750) rows re-fetched today match: BonusCoefficient 0.714 / 0.429, base 43 / 46, per level 0.8 / 1. | wago SpellEffect 70755, fetched 2026-11-04 | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Build `05005/15302102513501031/5003` (dwarf), decoded: Divine Strength 5, Improved Seals 5 / Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1 / Deflection 5, Conviction 3. Every point is modelled (tank_talents.go, holy_strike.go:80-110 for Shieldsworn's 4 s cut and shield strike). Reckoning against Anticipation was judged in round 16 P1: not adopted on chance of death. | talents/paladin.json; engine read | Aligned | - | - |
| P2 | Round-16 regression targets: `TestCastKingsDoesNotStackWithHeartOfTheLion` (E 913221285) and the Holy Strike rank row at 40-42 (`rank_choices_remeasured_test`) pass. Fork sim/paladin/protection passes. | go test | Aligned | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Builds `050250002/502/0523330100130135` (leveling) and `050050002/503/05233301001301351` (raid), decoded: Divine Strength 5, (Reverence 2 leveling), Improved Seals 5, Holy Conduit 2 / Divine Intellect 5, Precision 2-3 / Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5, (Unsealed Wrath 1 raid). Repentance 1 is a gate point. Every other point is modelled. | talents/paladin.json; engine read | Aligned | - | - |
| R2 | Round-16 regression targets: `TestRetributionCastsTheReMeasuredConsecrationRanks` (44-49 rank 2), `TestRetributionConsecrationSpendsThePoolOnThePace` and `TestRetributionPaladinRunsNoSanctityAura` pass. Fork sim/paladin/retribution passes. | go test | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Exorcism and Holy Shield crit still need a level-60 log. foreverlogs has nothing at 60 (O1). | foreverlogs probe 2026-11-04 | OPEN (carried) | - | - |

## Before/after

No code or data change, so after equals before. The engine and site are the round-start shas (f637ec32c, d7fc51fe7). Spot check: Holy raid 60, alliance, seed 1, frozen BiS, guide build: 671.7, the same as before.tsv.

## Files to regenerate

- None.

## Measure at refresh

- Nothing new.

## For other lanes

- Nothing. No shared code touched.

## For the guide lane

- Nothing to land now. If the owner's check shows the Oath is penalized, change holy.md:78 ("Oath of Light adds the same flat amount to every rank") and the card note at data/curated/apl/paladin-holy.json:584 (same words) with the row change. holy.md:33 and :59 ("a flat amount") stay true unless the check shows the coefficient rather than the full amount, and even then they read fine.

## Still open

- O1: whether Oath of Light's bonus takes the low-rank penalty (decides the Holy Light rows at 47-59). The owner's check is above.
- O4: whether the Oath's amount is the full amount or coefficient-scaled (+8.3% Holy raid 60 if full). Same check.
- The emergency Holy Light line's cheap rank at 47-59 is a taste call if the Oath stays unpenalized (O2 note).
- X1: Exorcism and Holy Shield crit.
- Carried guide-accuracy sims (Retribution race on Beasts, single-mob questing claim, Protection threat by ability) were not re-run: they are race and prose picks, not engine bugs.
- Warcraft Logs is still blocked (gcloud).
