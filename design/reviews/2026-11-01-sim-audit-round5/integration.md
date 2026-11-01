# Sim-accuracy round 5: integration (2026-11-01)

The owner's round-5 rulings, the data workflow, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: warlock.md, physical.md, hybrids.md, casters.md, ranker.md and page.md.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| R1 | /tiers/healer never orders healers by a margin within the measurement's error, or by a tuned profile | The ranker publishes each healer band's HPS error (`hps_error`) and its HPS at the heal profile's tuning bounds (`hps_at_bounds`, ×1.5 and ×1.65, now configuration in heal-profile.json). Two neighbours are ordered only when the one above leads by more than two standard errors of the difference at the tuned value and at both bounds. Otherwise they tie and share a rank. A file with no error orders nobody. | page.md P4 |
| R2 | Power Infusion stays out of the raid preset and every default | No default carried it. Tests now hold it out of the preset, every /sim default, an empty API request and the coach. /sim already lists it as a player buff, off by default. No priest list casts it. | page.md P3; casters.md |
| R3 | /sim consumables come from each spec's published set | The page sends the curated raid group plus the class kit, resolved by level the way the published numbers resolve them (`sim/request/raid-consumes.json`, generated, with a staleness test). Consumables above the character's level are cut (RequiredLevel). Combat, Survival, Fury, Fire, Destruction and Marksmanship now read the same on the page as published. Before, Combat was -7.7% and Survival +3.8%. | page.md P1, P1b, P1c |
| R4 | The BiS ranker never publishes a set its own check scores lower than an alternative it considered | The Arcane and Destruction raid losses were the engine's hidden healing-to-damage third (warlock.md, E fce7dd53d), not the ranker. Three ranker causes stand on their own. First, `weight_stats` left out stats Forever's talents turn into damage: Intellect for Enhancement, the hunters and Retribution, and spirit and mp5 for mana casters (S bca6074c7). Second, the stat runner-up never reached verify (S 04e3b11a2). Third, only healers had the published-set guard (S ad919b369). Every role is now guarded. Where linear weights cannot see the set (the hit cap, ties within the weights' error, tank score), the guard is the fix. | ranker.md |
| R4b | Rhok'delar for Marksmanship if obtainable in Phase 1 | Not obtainable, so it is not added. Its chain needs Ancient Petrified Leaf, which drops only in Molten Core (not in Forever's announced content), and a sinew from Onyxia (raids-1, 9 December). Its sibling reward from the same leaf is already gated `opens: later`. It sims correctly against its client row (physical.md). | ranker.md RK1 |
| R5 | Ice Barrier: a second source for +15% / +5%, else the client's table | There are three sources. The raw 31 October hotfix push for 70707 on wago.tools sets Frigid Power 1326891's effects 3 to 15, and Ice Barrier's own effect pairs 3 to 5. The hotfix cache wowforevertalents.com reads agrees, and so does Wowhead Forever. The engine's 15% (Frigid Power) and 5% (while shielded) stand. The shipped tables' 3 / 3 was the 16 October hotfix. | casters.md R1, R1b |

## Integrator findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **data.yml failed on every push from 2026-11-01 00:18.** The committed addon-data.json and Data.lua derive from the curated rotations. Rounds 3 and 4 edited rotations (Fire's Evocation, Retribution's seal twist, the warlock curses) without a regeneration, and only the nightly bis.yml regenerates them, so main stayed red until the next night. The owner rule is that nobody regenerates them by hand. A new `addon-data` job on a push to main now regenerates and commits them as forever-data-bot. The strict `--check` stays on pull requests. Its first run committed 13c8d42e9, and data.yml has been green since. | runs 41332429890 (red) and 41335181189, 41335615696 (green) | CONFIRMED (CI) | S 23ff379a1 |
| I2 | **Two data tests pinned generated state.** test_addondata and test_audit_addon asserted the committed copies were in step. The dps-per-point test pinned a nightly measurement, and the 2026-11-01 re-measure moved Arms band 20's weapon term 2.1%, past the 2% pin. The drift tests now read a session-scoped regeneration in a scratch copy (tests/conftest.py; AuditContext takes the Lua path). The dps-per-point tests read the committed figures the generator carries through, plus shape rules. | memory: data contracts are floors | CONFIRMED (tests) | S 23ff379a1 |
| I3 | **The addon showed abilities before the player can train them.** An unranked spell family took every id spellranks.json lists, including client copies with no SkillLineAbility row at lower levels. 20 rotation lines were affected: Recklessness at bands 30 to 40 for Fury (1719 is trained at 50; the copy 13847 sits at 22), Evocation at 10 for Fire (456397 at 1), Barkskin and Frenzied Regeneration from 10 for the bear, Innervate from 10 for Restoration, and Hammer of the Righteous from 10 for Protection. An unranked family now keeps only the class's trainable ids, when it has any. Placeholder icons fall from 28 to 24. The sim already gated these through its engine overrides, so no number moves. | wago SkillLineAbility 70707 (no row for 456397, 13847, 428713); test_addonrotation (failed first) | CONFIRMED-BUG (addon) | S 576be1e9d |
| I5 | **My I1 change cost the bis.yml run its picks.** Run 41332442067 ranked every spec, then its rebase conflicted on addon-data.json and Data.lua, which the new push job had committed meanwhile. The conflict handler only took tonight's side for data/builds, so the rebase aborted and all five pushes failed. Nothing was published from that run. The handler now renders both files again from the rebased tree, and a failed render aborts rather than commit conflict markers. Pushed alone, so the 06:23 nightly does not repeat it. | run 41332442067's log (CONFLICT in addon/ForeverSixty/Data.lua) | CONFIRMED (CI, self-inflicted) | S 441b96a42, on main as fd6adb354 |
| I4 | **Horde race picks after the racial fix** (E e4d9ff60c). These are the guides' `recommendedRaces`, curated. Horde, band 60 raid, seed 1, 10,000 iterations (error about 0.4):<br>- Orc beats the published Troll by about 1.0% on all three warlocks: Affliction 565.3 against 559.5, Demonology 645.5 against 638.9, Destruction 673.4 against 666.7. Command's +5% pet damage drives this.<br>- Undead beats the published Orc by 0.5 to 0.7% on Arcane (589.2 against 585.1) and Fire (654.2 against 650.8). Frost's Troll is within 0.3% of Undead.<br>The gaps sit at or under the adoption rule's 1% bar, so the picks are left for the owner. | spec-breakdown with `-bis-dir` copies whose horde race is swapped | OPEN (owner) | - |

## Integration

- Fork `int/r5` merges `lane/r5-warlock`, `lane/r5-casters`, `lane/r5-hybrids` and `lane/r5-physical`. The hybrids and physical lanes carried cherry-picks of the racial, rune and healing commits. Two kinds of conflict needed resolving:
  - Champion of the Light: kept the hybrids lane's spell-damage-only version.
  - Goldens: TestProtection, TestElemental and TestEnhancement, regenerated at the merge.
- The full fork suite passes with `-tags=with_db`, except `sim/web`, which already fails setup without the `binary_dist` package.
- The site merges `wt/r5-warlock`, `wt/r5-casters`, `wt/r5-hybrids`, `wt/r5-physical`, `wt/r5-page` and `wt/r5-ranker` with no conflicts.
- The pin is 8baffe5fc. The rotation ladder goldens are regenerated at the pin with `make simdb` and a scratch go.work.
- Gates, all exit 0:
  - `go vet` and `go test ./...` in sim
  - the coach and sims packages in api
  - `make apl-check` against the pin
  - `make styles-check`
  - the full data pytest
  - the full web vitest
  - astro check
  - `npm run lint`
  - the sim, tiers, planner-share-sim and header-finished-sim e2e specs

## Before/after

Pin f8a5f94ce with origin/main, against pin 8baffe5fc with this branch. Both run on today's published data: the BiS sets, builds and the committed simdb.bin. Raid preset, band 60, 10,000 iterations. Healers are measured in effective HPS on the heal profile (×1.65), tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 524.9 to 501.0 | -4.6 / -4.5 | 509.4 to 489.5 | -3.9 / -3.9 |
| druid-feral | 611.5 to 611.5 | +0.0 / +0.0 | 613.1 to 613.1 | +0.0 / +0.0 |
| druid-feral-bear | 679.7 to 679.7 | +0.0 / +0.0 | 672.5 to 672.5 | +0.0 / +0.0 |
| druid-restoration | 734.4 to 741.5 | +1.0 / +1.0 | 732.0 to 739.2 | +1.0 / +1.0 |
| hunter-beast-mastery | 644.8 to 649.8 | +0.8 / +0.8 | 661.1 to 667.4 | +1.0 / +1.0 |
| hunter-marksmanship | 649.2 to 655.2 | +0.9 / +0.9 | 665.7 to 673.2 | +1.1 / +1.1 |
| hunter-survival | 642.0 to 642.0 | +0.0 / +0.0 | 650.2 to 651.8 | +0.2 / +0.2 |
| mage-arcane | 601.8 to 585.3 | -2.7 / -2.7 | 605.5 to 585.1 | -3.4 / -3.4 |
| mage-fire | 625.8 to 659.3 | +5.4 / +5.1 | 618.9 to 650.8 | +5.2 / +5.0 |
| mage-frost | 547.2 to 532.9 | -2.6 / -2.6 | 543.3 to 534.7 | -1.6 / -1.5 |
| paladin-holy | 549.5 to 551.6 | +0.4 / +0.3 | 541.9 to 541.9 | +0.0 / +0.0 |
| paladin-protection | 685.4 to 1028.1 | +50.0 / +50.1 | 690.2 to 1043.2 | +51.1 / +51.0 |
| paladin-retribution | 635.7 to 635.6 | -0.0 / +0.0 | 578.6 to 583.1 | +0.8 / +0.6 |
| priest-discipline | 526.5 to 536.0 | +1.8 / +1.8 | 543.1 to 545.7 | +0.5 / +0.5 |
| priest-holy | 639.7 to 641.3 | +0.3 / +0.3 | 618.4 to 620.2 | +0.3 / +0.3 |
| priest-shadow | 665.4 to 651.0 | -2.2 / -2.2 | 668.4 to 656.1 | -1.8 / -1.8 |
| rogue-assassination | 601.0 to 601.0 | +0.0 / +0.0 | 614.5 to 616.1 | +0.3 / +0.1 |
| rogue-combat | 608.1 to 608.1 | +0.0 / +0.0 | 622.0 to 621.8 | -0.0 / +0.0 |
| rogue-subtlety | 545.3 to 545.3 | +0.0 / +0.0 | 559.7 to 560.6 | +0.2 / +0.1 |
| shaman-elemental | 464.7 to 449.1 | -3.4 / -3.4 | 466.3 to 447.0 | -4.1 / -4.1 |
| shaman-enhancement | 618.4 to 615.0 | -0.5 / -0.5 | 639.8 to 630.6 | -1.4 / -1.4 |
| shaman-restoration | 640.9 to 640.0 | -0.1 / -0.1 | 639.6 to 638.3 | -0.2 / -0.1 |
| warlock-affliction | 577.0 to 559.7 | -3.0 / -3.0 | 572.4 to 559.5 | -2.3 / -2.3 |
| warlock-demonology | 664.9 to 637.3 | -4.2 / -4.2 | 655.7 to 638.9 | -2.6 / -2.7 |
| warlock-destruction | 688.8 to 668.0 | -3.0 / -3.0 | 682.2 to 666.7 | -2.3 / -2.3 |
| warrior-arms | 726.5 to 726.5 | +0.0 / +0.0 | 751.2 to 745.2 | -0.8 / -0.7 |
| warrior-fury | 763.6 to 763.6 | +0.0 / +0.0 | 778.8 to 780.4 | +0.2 / +0.3 |
| warrior-protection | 1038.1 to 1038.1 | +0.0 / +0.0 | 1026.4 to 1026.4 | +0.0 / +0.0 |

What this table does not hold, because each needs a nightly regeneration:
- **The item-data fixes** (physical.md):
  - the wago hotfix pushes (201 rows wrong and 569 items missing at 70707)
  - shield block value
  - Wowhead's on-equip stats where the client lost the link (Benediction's +65 / +195 among 153 items)
  - expertise and haste rating
  They reach published numbers once data.yml regenerates simdb.bin (its fetch job now runs `fetch-wago-hotfixes`) and bis.yml re-ranks. The hotfix rows and the on-equip stats must land together: the hotfix rows alone took Restoration shaman to -8%, and the pair nets +0.3%.
- **The re-ranks:**
  - the warlock sets picked under the hidden healing-to-damage third
  - Fire's new raid build (+3.3%)
  - Survival's off hand
  - Horde Retribution's weapon
  - the ranker fixes

## wow-player review

First pass, on the lane reports, the before/after table and the page screenshots:
- **SHIP:** the warriors, rogues, Beast Mastery and Marksmanship (Hunter's Mark 143 matches SkillLineAbility, SpellEffect and Wowhead), Balance, Elemental, Enhancement, the druids, Holy paladin and Restoration shaman.
- **SHIP WITH FIXES:**
  - Warlocks: the sets picked under the hidden third re-rank at the bis run.
  - Survival: the off hand re-ranks at the bis run.
  - Mages: horde race picks, I4.
  - Priests: Benediction's healing. Physical f52842d53 restores +65 / +195 at the data regeneration.
  - Protection paladin: stacking is unproven. Hybrids T4 now has TrinityCore and cmangos, which both multiply SPELL_AURA_MOD_THREAT effects, and the guide says it awaits a launch threat log (S d2a3e74d7).
- **HOLD** on three surfaces, each settled before the push:
  - **/sim.** The shots showed a leftover warrior result. That was the fake engine's fixture, but the lane then found two real ways one character's result could show under another, and fixed both (S d24be0efc): a run in flight when another character loads, and an unowned Droptimizer upgrade. The headline line now merges ranks of one ability and leaves out always-on, never-cast and queue auras (S a3e032c56, 6a7466fb3, d1d396d43). The reshoots are real wasm runs on 8baffe5fc: Combat 621 against published 620.4, Survival 652 against 653.7. SHIP WITH FIXES: re-shoot Survival after the bis run puts Heartseeker in the off hand.
  - **/tiers/healer.** Until the nightly publishes the error, the page fell back to a 1% margin. Then, with no error, it showed five healers tied over a 25% spread. With no error published, the page now shows no rank column and no tie marks, only the gaps and a note (S 2f1b502a7, 594d62e3a). The tie mark is legible, and one sentence gives the reason. SHIP. After the nightly, check that the published error yields 1, =2, =2, 4, 5, and re-shoot Horde.
  - **Ranker.** This hold was only for the missing report; ranker.md has landed since.
- **Island sizes.** Not a review item, found while re-shooting. The per-level consumables table took sim-island to 93,013 bytes gzipped against a 92,160 budget, and planner-island over its budget too, so deploy's postbuild would have failed. The table now loads lazily, and no budget was raised (S 26dc630da): 91,875 and 83,598.
