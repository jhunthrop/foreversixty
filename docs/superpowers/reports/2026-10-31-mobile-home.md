# Mobile lane F2: Home, your character's answer (2026-10-31)

Branch `wt/mobile-f2-home`. Not pushed. Spec: `design/specs/2026-10-30-mobile-apps.md` §5.2 (F2.1 to F2.7).

## Commits

| Commit | What |
|---|---|
| 7bb25a4d5 | Contract and copy: hunter item, BiS and set fixtures; the tooltip's C6 facts; spec drift fixed; `items.*` and `home.*` copy |
| 2c2ef67d8 | iOS: Home, its states, Your gear, the tooltip sheet |
| 51e479a17 | Merge origin/main (the shared-device rule, 5d86a8dc7) |
| 8b7e7ce42 | Android: the same, at parity |
| (last) | iOS gear-list fixes, the two extra parity tests, the captures and this report |

## What was built (both platforms)

- **Home (F2.1).** Renders `GET /v1/characters/{region}/{ruleset}/{name}/home` and nothing else: the hero meta, the two figures (`11 upgrades`, `+76.2 DPS`), the note with the unmodelled count, the sync age, the Your upgrades card with the first five rows **in the server's order**, the footer (`6 more upgrades · all 17 slots`), the Already line, and Next for {name} (Talents with its Matches or Differs tag, Sims). Gains are the server's: a modelled row shows its figure, any other row shows the server's `gain_words`. A worn item our data does not know shows the server's `label` (`Not in our item data yet`).
- **States.** F2.2 loading skeletons at real heights with `Loading your upgrades…`; F2.3 offline (the cached answer under the strip, `Offline. Showing what you had at {time}.`); F2.4 no gear export (Get the addon, the list button, the note); under 20; no spec; no list; every slot matches; F2.5 error with Try again, the Next rows still drawn from the cache.
- **Caching.** The answer per character on disk (iOS Application Support, excluded from backup; Android `noBackupFilesDir`), drawn at once, refreshed on pull and on foreground after five minutes. A transport failure with a cache is offline; a server error is the error state.
- **Your gear (F2.7).** 17 rows in character-frame order: an upgrade row is worn → pick, an already-best row carries the BiS tag, any other slot shows the export's worn item looked up in the items file, or Empty.
- **Tooltip sheet (F2.6).** iOS sheet at medium and large detents; Android `ModalBottomSheet`. Reads `items/{class}.json`, `sets.json` and `bis/{spec}.json` from the CDN, cached per build (memory plus disk). Lines in the game's order: name, Binds, Unique, slot with armor type or weapon handedness, armor, damage and speed, DPS, stats, the effect in green, the set and its bonuses, Requires Level, Item Level in gold. Then Where it comes from (with `{Faction} ({Standing})` for a rep reward and ` · requires {Standing}` for a vendor), The next best (two runners-up, `You wear this · ` on the worn one) and the pick footer.
- **Icons.** `/data/{build}/icons/{icon}.webp`, quality border, black inner line, glow; the tile draws at once. iOS: a URLCache with a 200 MB disk cap plus a decoded-image cache. Android: Coil 3.0.4 with a 200 MB disk cache.
- **Gate.** `character_home`, on in fixture builds and wherever the API turns it on.

## Contract changes (commit 7bb25a4d5)

- New fixtures cut from the active build's generator output: `cdn.items.hunter` (86 items), `cdn.bis.hunter-beast-mastery` (the four level-60 bands), `cdn.sets`. The home fixture is a hunter, and fixture mode had no hunter files.
- `Item` gains optional `bonding`, `armor_type`, `min_standing` (spec C6). The hunter fixture carries them, read from the client's own Item and ItemSparse rows (1.60.1.70532 full tables, overlaid with the active build's hotfix rows). Items absent from those tables carry null.
- `GET /data/{build}/sets.json` joins the spec.
- Drift the new fixtures exposed, fixed to the real generator shapes: `Item.stats_source` is nullable (null on most rows), BiS `set_bonus` is an object, `ties` are objects, `new_at_band` is strings. **Without these fixes neither app could decode the real hunter items or BiS files.** Lane F3 reads the same BiS model.
- Go contract tests hold the three fixtures to the generator.

## Tests

One list, implemented on both: 23 logic tests (iOS `HomePresenterTests.swift`, Android `HomePresenterTest.kt`): server order and top five, hero figures and note, gain words, true minus, replaces lines, all match, no gear, the three empty states, talents, sims, stamps, the store's four outcomes, the cache round trip, fixture mode end to end, the tooltip (pick, reputation, You wear this, weapon, unknown item) and the gear rows. Screenshot tests on both at the default size and the largest (iOS AX5, Android 2.0).

| Check | Result |
|---|---|
| `make -C apps/ios test` (shared simulator, behind the lock) | TEST SUCCEEDED: 128 tests, app-hosted 40 |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` | BUILD SUCCESSFUL, 130 JVM tests, 0 failures, lint clean |
| `go test ./internal/contract/` and the server OpenAPI test | ok |
| `vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | 28 passed |

## Captures

`design/reviews/2026-10-31-mobile-home/`, `<screen>-<state>-ios.png` and `-and.png`: home default, default-xl, loading, offline, nogear, under20, nospec, nolist, allmatch, error; tooltip medium, large, large-xl, rep, weapon; gear default (plus gear full on iOS).

They are the screenshot tests' own renders: iOS on the shared iPhone 17 simulator at 402 x 874 pt, Android in Robolectric at Pixel 9 (412 x 915 dp). They are not app screenshots: the fixture app's `/v1/me` main character is a warrior while the home fixture is a hunter, so a live fixture-mode screen would mix the two. The iOS renders leave out the tab bar (the iOS 26 glass bar does not render off-screen); the tooltip captures draw the sheet at its detent height over the dimmed Home.

## Concerns

1. **The selector is lane F1's.** Home picks `/v1/me`'s main character and titles the screen with the name; F1's selector replaces both. `HomeCharacter` is the seam.
2. **Next raid is not drawn.** It needs C2's `next_raid`, not in the contract yet. Talents and Sims rows do not navigate: F6 and F5 own those screens.
3. **C6 is pending in the data lane.** Until `items/{class}.json` publishes `bonding`, `armor_type` and `min_standing`, the live tooltip shows no Binds line, no armor type and no standing. The site's tooltip should gain them in the same change.
4. **Effect lines have no `Equip:` / `Use:` prefix.** `effect_text` does not say which trigger it is; the apps print it as the site does. The data lane would need to publish the trigger.
5. **Where it comes from** is the BiS row's words, else the home row's `source_line`. A worn item that is in neither shows no source: the apps do not port the site's loot.json source ranking (`lib/items/tooltip.ts`). Drop chances (`{n}% from {boss}`) come with that port.
6. **JetBrains Mono is not bundled** (apps/shared/fonts); mono figures use the system mono on both.
7. **Shared controls live in the Home folder** (primary and secondary buttons, tag, offline strip, cards). Other lanes need the same; they belong in DesignSystem once the lanes merge (tenet 10).
8. **Copy keys added beyond §13:** `home.answer.under20Button`, `home.answer.openPlanner`, `home.next.talentsNoBuild`, `home.hero.level`, `home.upgrades.bandRaid` (`60 · raid`; the boards read `60 · Phase 1 raid`, which the home answer cannot name), `home.age.*`, and the `items.*` family.
9. **Coil** joins the Android version catalog (3.0.4, already in the Gradle cache). Lane F3 needs it too; take one entry at merge.
