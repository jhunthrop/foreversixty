# Sim-accuracy round 15: platform (2026-11-03)

Lane `platform`: web components and web/src/lib. It had two seed items from guide-accuracy round 1 (ledger "Out of scope this round"; warlock.md fix pass 2 item 2; warrior.md WARRIOR-2 and items 106/111; druid.md DRUID-13/S4; paladin.md PALADIN-6). Site `wt/r15-platform` from int/r15 fd320b60e. No engine worktree; the fork is untouched.

Totals: **2 CONFIRMED-BUGs, both origin old.** P1 (guides-handoff) is the raid-build line. P5 is the published hit and expertise figures, which were read on the wrong character. **P3, my first expertise fix, was wrong and is reverted (0e44706db): Forever items do carry expertise.** Nothing changes any sim DPS.

## Cross-cutting (web)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **The "raid build moves N points" line hid talents.** The point count was right, but `listOf` in `web/src/lib/guides/build-diff.ts` cut each list at four talents and ended it with "and N more". So the named talents' bracketed points did not add up to N. Affliction read "moves 15 points from Soul Harvest (2), Improved Drains (3), Nightfall (2), Soul Siphon (3) and 3 more": 10 of 15 named. Demonology named 11 of 12. The same cut hit three more pages: druid Feral (from-list 11 of 12, to-list 10 of 12), rogue Combat (from 8 of 11, to 6 of 11) and shaman Enhancement (to 10 of 11). The fix names every talent. The test runs on every real guide's `build`/`raidBuild` pair and checks two things: every changed talent appears with its points, and both lists add up to the stated count. It also pins the full Affliction sentence. Before the fix, 7 tests failed (the 5 pages, the Affliction sentence, the list unit test). After, 25 pass. | `build-diff.ts:56-62` (old); hand count: an independent Python decode of each FS1 string against `data/builds/1.60.1.70755/talents/<class>.json` agrees with the decoder on every page (counts below) | CONFIRMED-BUG | old (86479f7c4; no sim round reads generated guide lines) + guides-handoff | S e762aaa8f |
| P2 | Hand count of every guide with a raid build (from / to, points and talents): druid Balance 8 (2 into 3), Feral 12 (5 into 5); hunter BM 2 (2 into 1), MM 11 (4 into 4); mage Arcane 2, Fire 3, Frost identical builds (no line); paladin Ret 2; priest Shadow 4; rogue Assassination 7, Combat 11 (6 into 5); shaman Elemental 5, Enhancement 11 (4 into 5); warlock Affliction 15 (7 into 4), Demonology 12 (5 into 4), Destruction 4; warrior Arms 5. Each total is 51 on both builds. The decoder's counts and talent names match the hand decode on every page. The only fault was the list cut in P1. | probe over `web/src/content/guides/*/*.md` frontmatter | Aligned | - | - |
| P3 | **Retracted.** The guide-accuracy lead said no Forever item carries Expertise, and I first agreed: `data/builds/<build>/items/*.json` has no expertise stat key and no "dodged or parried" text on equip lines. That file drops the stat. The engine's item database (`sim/internal/simdb/simdb.bin`, built from the client's ItemSparse) gives expertise to **71 catalogue items**. They are mostly the ilvl-98 level-60 sets: Lightbreaker (warrior), Inquisition (paladin), Waywatcher (druid; the Gauntlets 240081 give 2%), Duskwraith (rogue), Dawnstalker (hunter), Soulcrusher (shaman) and Soulforge pieces, plus Dwarven Tree Chopper 2907 (0.6%) and Servomechanic Sledgehammer 4548. Wowhead's Forever tooltips agree (nether.wowhead.com/forever/tooltip/item/240081: "Equip: Reduces chance to be Dodged or Parried by 2.0%"; 239550 1.0%; 2907 0.6%). With two sources, expertise is a real gear stat. The tank weights keep it, leveling-bis keeps `expertise_to_cap`, and the web keeps main's expertise-cap and talent-hit design. My four commits are reverted. | simdb probe (throwaway test, deleted); Wowhead Forever tooltips | Not a bug (retracted) | - | S 0e44706db (revert of c66c7f176, 1f8ac7821, 25045536b, 2fe9433c6) |
| P4 | **Feral bear 8% expertise vs 0% (coordinator lead).** In the 2026-11-04 nightly (de044125c, engine d1d9f2775), feral bear's level-60 `expertise_to_cap.baseline` is 8 on bare Alliance, bare Horde and raid Horde, and 0 on raid Alliance. Warrior protection reads 8 on raid Horde only. Every one of those published sets wears **0** expertise (summed over the simdb expertise of each band's picks). The 8 is the weights sweep's own character: `gearedBandWeights` re-measures on first-pass gear, which can carry expertise items (Waywatcher pieces) that the final, verified picks drop. Where that sweep is not adopted, the character is the bare ladder (0). See P5 for the cause. | `sim/cmd/leveling-bis/geared_weights.go:92`; `main.go` hit-profile call; per-band sums of the de044125c bis files | CONFIRMED-BUG (see P5) | old | S eb20817ce |
| P5 | **The published hit and expertise figures described the wrong character.** `hit_to_cap` and `expertise_to_cap` came from `runner.HitProfileFor(wreq)`, the weights request. That request wears the sweep's first-pass or ladder gear, and its race is one faction's race for both factions. In the failing test, the Horde band was profiled as a dwarf, not the troll it publishes. So the guide line "8% expertise from BiS gear" was false for the set the page shows. Main's web design is right; its input was wrong. The profile is now read on `publishedSetRequest(wreq, f.race, picks)`: the published picks on the faction's own race. Test first: `TestHitProfileIsReadOnThePublishedSet` failed (gear `ranged:1002` against the published set; race `dwarf` against `troll`) and now passes. The leveling-bis package passes as a whole (fake engine). | `sim/cmd/leveling-bis/main.go`, `verify.go` | CONFIRMED-BUG | old (hit_to_cap since the field was added; no round compared it to the published set) | S eb20817ce |

## Before/after

No sim DPS, TPS or HPS moves. P1 is a display fix. P5 changes only the published `hit_to_cap` and `expertise_to_cap` figures, from the next BiS refresh on.

## Files to regenerate

- None by hand. The next BiS refresh publishes `hit_to_cap` and `expertise_to_cap` read on the published set. Until then, the published bear and warrior-protection Horde-raid expertise lines read 8% where the set has 0%. Hit lines may also be off where the sweep's gear or race differs from the published set.

## Measure at refresh

- Check that every melee band's `expertise_to_cap.baseline` equals the simdb expertise its published picks wear (0 for every current tank and DPS set), plus Weapon Expertise for Combat (2). Check that `hit_to_cap.baseline` matches the published set plus talents and buffs on both factions.

## For other lanes

- **Data lane:** `data/builds/<build>/items/*.json` drops the expertise stat (ItemSparse expertise rating) that `simdb.bin` carries for 71 items. The site's item tooltips and filters therefore never show "Reduces chance to be Dodged or Parried". Guide-accuracy round 1 also read "no expertise" from these files.
- **Coordinator:** eb20817ce touches `sim/cmd/leveling-bis` (shared BiS tool). It changes no sim number and no pick, only the two published cap figures.

## For the guide lane

- Guide-accuracy round 1's WARRIOR-2, DRUID-13 and PALADIN-6 rest on "no Forever item carries expertise", which is wrong (P3). The ilvl-98 level-60 sets carry 1-2% each. Dropping Expertise from the tank statPriority lists should be revisited. The weights at refresh will say where it ranks.
- No prose depended on the old "and N more" wording.

## Pages for a later capture

- Raid-build line (P1): /guides/warlock/affliction, /guides/warlock/demonology, /guides/druid/feral, /guides/rogue/combat, /guides/shaman/enhancement.
- Expertise cap lines (P5, after the refresh): /guides/druid/feral-bear, /guides/warrior/protection and /guides/rogue/combat (the line should name Weapon Expertise 2%).

## Still open

- Not run, per the brief: Playwright and the full vitest suite. Touched vitest files pass (268 tests) and `astro check` reports 0 errors after the int/r15 merge (f20eabe7a).
