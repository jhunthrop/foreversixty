# Sim-accuracy round 9: platform lane

Site branch `wt/r9-platform` from origin/main 55a1212f7. No engine change. Scope: `.github/`
(never bis.yml, data.yml or engine.yml), the custom /sim panel, the coach catalogue, the Discord
bot. Nothing nightly-owned was regenerated or committed.

**Found: 3 CONFIRMED-BUG rows (P3, P4, P5), all fixed, plus the two owner-ruled CI items (P1, P2).**
No row moves any spec's DPS: nothing here touches the engine, the presets or the sim request.

## CI

| # | finding | evidence and source | confirmed? | fix sha |
| --- | --- | --- | --- | --- |
| P1 | **actionlint rejected the `xcode-27` runner label** (ios.yml:32, apps-release.yml:55). tools.yml also linted only six of the 17 workflows; api, companion and sim carried shellcheck findings. | actionlint 1.7.12 (the version tools.yml pins, darwin build in the lane scratch): before, the CI set exited 1 (`label "xcode-27" is unknown`) and the whole tree exited 1 with 13 findings (api 2, apps-release 1, companion 1, data 7, ios 1, sim 1). After: `actionlint` over every workflow exits 0. New `.github/actionlint.yaml`: `self-hosted-runner.labels: [xcode-27]`, and data.yml's shellcheck findings ignored by path because data.yml is held and nobody edits it. api.yml and companion.yml quote `$IMAGE` and `${ASSET}`; sim.yml uses `sha256sum -- *` (same file names in SHA256SUMS). tools.yml runs `actionlint` with no file list and triggers on the config. | Owner item, done | S f02694e28 |
| P2 | **`real-engine` did not gate deploy.** deploy had `needs: [build]` only. | web.yml before: `deploy: needs: [build]`. Now `needs: [build, real-engine-scope, real-engine]` with `if: always() && !cancelled() && needs.build.result == 'success' && (needs.real-engine.result == 'success' \|\| (needs.real-engine.result == 'skipped' && needs.real-engine-scope.outputs.run == 'false')) && <the old main/push/dispatch clause>`. A failed or cancelled real-engine blocks; a scope skip deploys; a failed scope (which also skips real-engine) blocks. On push the scope always says run, so every main deploy now waits for the real wasm and its four specs. Header diagram and the BUILD MODE note updated. No test asserted the job graph; `web/src/ci-deploy-gate.test.ts` now evaluates deploy's `if:` over every result the jobs can report (red before: 3 of 7 cases failed; green after). actionlint exits 0. | Owner item, done | S 8d9ea7ec5 |

## Discord bot

| # | finding | evidence and source | confirmed? | fix sha |
| --- | --- | --- | --- | --- |
| P3 | **An unranked PvP reward had no source line in the bot.** Round 8 relabelled 467 Forever PvP listings as `pvp` sources with no `rank` ("PvP reward (Alliance/Horde)", `title_rank` for eligibility only). The bot's `pvpLine` needed a rank: a `pvp` pick fell back to the bare badge "PvP", and a quartermaster pick read "Vendor: Sergeant Thunderhorn". The site reads "PvP reward · Horde" (`web/src/lib/items/source-labels.ts` pvpSourceLabel, S 840f51d97). | `data/pipeline/loot/pvp_faction.py:370-380` (rank None, name "PvP reward (<Side>)"); `api/internal/discordbot/item_source.go` resolvedSource/vendorLine. Test first: `TestSourceLinesAreTheBisPagesOwn` gains item 227198 as an unranked Horde reward, read as `pvp` and as `vendor`; it failed with "PvP" and "Vendor: Sergeant Thunderhorn". Now both read "PvP reward · Horde"; ranked lines unchanged. | CONFIRMED-BUG | S 64d471b65 |

## Coach catalogue

| # | finding | evidence and source | confirmed? | fix sha |
| --- | --- | --- | --- | --- |
| P4 | **The coach could not see Forever food at the pull.** Each food entry was the item's use spell (Nutritious Food or Drink). In Forever that spell is a 10 second aura whose effect 1 (aura 227) triggers the 15 minute Well Fed, so a fed player has Well Fed up at the pull, not Nutritious Food. A fed raider read as unfed, or "could not check". | wago.tools 1.60.1.70707 (fetched 2026-11-02): SpellMisc DurationIndex 1 (10 s) for 1249522, 1249513, 1248401, 1249925; 347 (900 s) for their Well Fed 1249523, 1249520, 1248422, 1249927. SpellEffect: e.g. 1249522 eff1 aura 227 bp 1 trigger 1249523 (aura 290). Single source (client tables; no beta log carries any food aura, 238 cached foreverlogs reports grepped). Each food now carries both ids: the use spell (still a cast the adherence check excludes) and the Well Fed. Test first: `TestCatalogue_EachFoodCarriesItsUseSpellAndItsWellFed` (failed for all three roles). | CONFIRMED-BUG | S 7fa74a25f |
| P5 | **The catalogue and the personas named the old foods.** dps listed only Grilled Squid (a caster eating Nightfin Soup, or a warrior eating Dumplings, got no food credit), tanks Tender Wolf Steak (15 Agility in Forever), healers Nightfin Soup (the preset's healer food is Sage's Tea). The starter personas skipped food for every caster and named Squid for Strength users and the Steak for tanks. | Round 8 auras C1/C2 and `sim/request/raid-consumes.json` top-band food per spec. Test first: `TestCatalogue_FoodsAreThePublishedPresetsFoods` (4 misses) and pytest `test_food_is_the_published_presets_food`. Now dps: Squid, Dumplings, Nightfin; healer: Sage's Tea and Nightfin; tank: Dumplings. `tools/coach/starter_food.py` reads each spec's preset food; the food lines of 19 starters were spliced from the generator and the three full personas (warrior-fury, priest-holy, warrior-protection) edited by hand. A shared Well Fed id (1249520 is every spell-damage food's) is named for the preset food; a Sagefish eater is credited as fed under that name. | CONFIRMED-BUG | S 7fa74a25f |

## Custom /sim panel

| # | finding | evidence and source | confirmed? | fix sha |
| --- | --- | --- | --- | --- |
| P6 | **The Debuffs section offered five boxes that do nothing.** Shadow Weaving, Improved Scorch, Improved Shadow Bolt and Stormstrike are the caster's own in Forever (round 8 D3, E a56f0d992, an ancestor of pin 40021374a); at the pin no engine code reads those four `Debuffs` fields (`git grep` over sim/, excluding protos and tests). Curse of Shadow is deleted from the client (casters lane, E 35fa13a5a, integrator's request). | Before: 24 boxes, shots `sim-custom-debuffs-1440-before.png`, `-390-before.png`. After: 19, `-after.png` (Playwright on a local fixture build; every server stopped). `INERT_BUFFS` in `web/src/lib/sim/buffs.ts`; `rowsIn` leaves them out. The catalogue and the request vocabulary keep them, so a saved sim that named one still runs. Test first: buffs.test.ts "offers none of the debuffs the engine ignores". Note: Curse of Shadow does nothing only once 35fa13a5a is pinned. | Panel item (round 8 open), fixed | S c55381ed1 |

## Before/after

No DPS change: no engine, preset or request change in this lane.

## For other lanes

- **Coach (whoever owns the starters next):** `tools/coach/test_persona_starter.py` has seven failures on origin/main 55a1212f7 too (guides and APLs moved since the starters were written; e.g. druid-balance's Moonfire, mage-arcane's priority, warrior-arms' Recklessness). No workflow runs it. This lane spliced only the food lines rather than regenerate the rotation text.
- **Casters / auras (engine):** the raid's `winters_chill` flag. Forever's Winter's Chill 12579 is eff0 aura 271 misc 16 bp 5 (Frost damage taken from the caster, per stack; wago 70707), and the engine's `WintersChillAura` (core/debuffs.go:508-535 at 40021374a) is a bare stack counter the mage's own Cone of Cold reads. As a raid flag it changes nothing for any other class and could hand a mage stacks it never applied. If the engine makes the flag inert, add `winters_chill` to `INERT_BUFFS` in web/src/lib/sim/buffs.ts (platform).
- **Page:** the panel's labels come from humanised ids where simbuffs.json has no name ("Hunters mark", "Curse of elements", "Winters chill").

## Still open

- P4 rests on one source (the client's spell durations). A level-60 Forever log after launch will show whether COMBATANT_INFO lists Well Fed (expected) or Nutritious Food at the pull.

## wow-player review (on the screenshots and files, 2026-11-02)

- **Panel (P6): ship.** Removing the four caster-own debuffs and Curse of Shadow is right; Curse of the Elements covers every school. Asked to check Winter's Chill (routed above). Label nits, separate from this change: "Hunters mark", "Curse of elements", "Winters chill" and sentence case where the client uses title case.
- **Discord (P3): ship.** Matches the site; no "PvP rank 0". Nit: append the vendor when loot.json names one. Round 8's ruling is that the client does not say how these are sold, so there is no vendor to name yet.
- **Coach (P4, P5): ship, with two wording notes.** The foods match what a Forever raider eats in Phase 1, and the Well Fed check is right. Notes: the healer prose names Sage's Tea while the catalogue also credits Nightfin Soup; and the dps food credit is role-wide, so Dumplings count as food for a mage. Both are by design of the catalogue (credit for any listed food of the role; it never calls food missing because food is marked incomplete). Left as is; a class-aware food credit is a coach design change, not a sim fix.
