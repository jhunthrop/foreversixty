# Sim-accuracy round 15: platform (2026-11-03)

Lane `platform`: web components and web/src/lib. It had two seed items from guide-accuracy round 1 (ledger "Out of scope this round"; warlock.md fix pass 2 item 2; warrior.md WARRIOR-2 and items 106/111; druid.md DRUID-13/S4; paladin.md PALADIN-6). Site `wt/r15-platform` from int/r15 fd320b60e. No engine worktree; the fork is untouched.

Totals: **2 CONFIRMED-BUGs, both origin old, both guides-handoff.** Nothing changes any sim DPS. The expertise fix changes what the BiS weights run measures for the three tanks from the next refresh on (see "Measure at refresh").

## Cross-cutting (web)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **The "raid build moves N points" line hid talents.** The point count was right, but `listOf` in `web/src/lib/guides/build-diff.ts` cut each list at four talents and ended it with "and N more". So the named talents' bracketed points did not add up to N. Affliction read "moves 15 points from Soul Harvest (2), Improved Drains (3), Nightfall (2), Soul Siphon (3) and 3 more": 10 of 15 named. Demonology named 11 of 12. The same cut hit three more pages: druid Feral (from-list 11 of 12, to-list 10 of 12), rogue Combat (from 8 of 11, to 6 of 11) and shaman Enhancement (to 10 of 11). The fix names every talent. The test runs on every real guide's `build`/`raidBuild` pair and checks two things: every changed talent appears with its points, and both lists add up to the stated count. It also pins the full Affliction sentence. Before the fix, 7 tests failed (the 5 pages, the Affliction sentence, the list unit test). After, 25 pass. | `build-diff.ts:56-62` (old); hand count: an independent Python decode of each FS1 string against `data/builds/1.60.1.70755/talents/<class>.json` agrees with the decoder on every page (counts below) | CONFIRMED-BUG | old (86479f7c4; no sim round reads generated guide lines) + guides-handoff | S e762aaa8f |
| P2 | Hand count of every guide with a raid build (from / to, points and talents): druid Balance 8 (2 into 3), Feral 12 (5 into 5); hunter BM 2 (2 into 1), MM 11 (4 into 4); mage Arcane 2, Fire 3, Frost identical builds (no line); paladin Ret 2; priest Shadow 4; rogue Assassination 7, Combat 11 (6 into 5); shaman Elemental 5, Enhancement 11 (4 into 5); warlock Affliction 15 (7 into 4), Demonology 12 (5 into 4), Destruction 4; warrior Arms 5. Each total is 51 on both builds. The decoder's counts and talent names match the hand decode on every page. The only fault was the list cut in P1. | probe over `web/src/content/guides/*/*.md` frontmatter | Aligned | - | - |
| P3 | **Expertise was shown as a stat players can gear for.** No Forever item carries an Expertise stat. Across every class list in `data/builds/1.60.1.70755/items/` the stat keys are agility through strength with no expertise key, and the only text hit is the item name "Hammer of Expertise". `simitems.json` has no expertise either. Expertise reaches a character in only three ways. (1) Set bonuses: Deathbone Guardian 5P (1213289 "Increased Expertise 2", aura 593, -2% dodged/parried); the Tier 1 tank 4Ps, Battleplate of Glory, Justice Battleplate and Grovekeeper Rage (1301096/1301083/1301076, aura 189 expertise rating 12 = 1.2%); and the SoD Bronze Dragonflight signet. The ranker scores set bonuses by full sim, never by the weight. (2) The rogue's Weapon Expertise talent, which the talent build decides. Even so, the pages printed an "Expertise cap: 6.5% (dodge)" note under every melee guide stat table and an "Expertise to cap" line on the BiS rail. The tank tables also carried an "Expertise per 1%" weight that priced nothing a player can equip. Decision: Expertise leaves both the cap line and the weights. The fix is made at the source. `data/curated/specs.json` drops expertise from the three tank weight lists, and the generated Go, TS and JSON copies follow (`python -m pipeline specs`, diff is only those lines). leveling-bis stops publishing `expertise_to_cap` (`expertisecap.go` deleted). The web BiS loader keeps only the weights the spec's catalogue entry lists and drops `expertise_to_cap` from older files, so today's published files render correctly before the refresh. The cap line, its copy and its props are deleted. Tests that failed first: `TestNoSpecWeighsExpertise` (sim/specs), `TestBandReportPublishesNoExpertiseFigure` (leveling-bis), and four vitest cases (load.test.ts x2, GuideStatTable.test.ts x2). | items/*.json stat counter; wago SpellEffect 1.60.1.70755 (aura 593 and aura 189 misc 8388608 rows); `sets.json` 124/2103/2108/2113; `sim/core/client_sets.go:417-424` | CONFIRMED-BUG | old (expertise cap line added 2026-11-02; weights list predates it) + guides-handoff | S c66c7f176, S 1f8ac7821, S 25045536b, S 2fe9433c6 (e2e expectation) |
| P4 | The published tank BiS files carry `expertise_to_cap.baseline: 8` (warrior, paladin, bear; dodge 0, parry 6). The current engine gives these characters 0 expertise. I read `core.ComputeHitProfile` on the band 60 raid and bare requests at d1d9f2775: expertise 0.00, boss dodge 6.5, parry 14. So the 8 is stale output from engine 8baffe5fc (files generated 2026-11-01). It is moot now that the field is retired, and the refresh replaces those files. | throwaway probe test (deleted) | Aligned (stale data; refresh owns) | - | - |

## Before/after

No sim DPS, TPS or HPS moves. P1 and P3 are display fixes. P3's specs-catalogue change takes effect only when the BiS weights run next runs.

## Files to regenerate

- None by hand. The next BiS refresh publishes tank weights without `expertise` and all bands without `expertise_to_cap`. Until then the web loader filters both.
- `python -m pipeline specs` was run, since it is codegen from `data/curated/specs.json` and not nightly data. It changed sim/specs/specs.go, web/src/lib/sim/specs.ts and web/src/data/specs.json, only the three expertise entries (in c66c7f176). `check_specs` should pass.

## Measure at refresh

- Tank BiS weights (warrior-protection, paladin-protection, druid-feral-bear) are measured without the expertise sweep. Picks should not move, because no item carries expertise. Confirm the refresh's tank picks match the previous run apart from unrelated changes, and that the weights run takes less time.

## For other lanes

- Mobile apps (apps/ios, apps/android) bundle `web/src/data/specs.json`, so its tank `weight_stats` lose expertise. I found no app code that reads `expertise` or `expertise_to_cap`. The ios.yml and android.yml path filters will fire on this file.
- Coordinator: P3 touches shared code outside the web (sim/specs, sim/cmd/leveling-bis, data/curated/specs.json). No sim engine path changes.

## For the guide lane

- Until land/guides-r1 merges, int/r15 still lists `Expertise` in statPriority for druid/feral-bear.md:9, paladin/protection.md:22 and warrior/protection.md:22. With no expertise weight, the stat table shows an "Expertise: Not significant" row and a "Expertise shows as Not significant" note. guides-r1 already removes these lines (DRUID-13, PALADIN-6, WARRIOR-2), so the merge settles it. paladin/protection.md:114 "Expertise and Defense come first" on int/r15 is also already rewritten on guides-r1.
- No prose depended on the old "and N more" wording.

## Pages for a later capture

- Raid-build line (P1): /guides/warlock/affliction, /guides/warlock/demonology, /guides/druid/feral, /guides/rogue/combat, /guides/shaman/enhancement.
- No expertise note or weight (P3): the guide stat tables of /guides/warrior/{arms,fury,protection}, /guides/paladin/{retribution,protection}, /guides/druid/{feral,feral-bear}, /guides/rogue/{assassination,combat,subtlety}, /guides/shaman/enhancement and /guides/hunter/survival, plus the weights rail of the same specs' /bis pages.

## Still open

- Not run per the brief: Playwright (`tests/e2e/bis-tank.spec.ts` updated, not run) and the full vitest suite. `astro check` gives 1 error, a missing `src/data/generated/rotations.json` in this worktree (environment, unrelated).
