# Mobile Wave 3 F9: Guides, tier lists, what changed (2026-11-01)

One lane, both platforms and the static data. Spec: `design/specs/2026-10-31-mobile-wave3.md` §5.1 and
§10 (W3-C1). Boards: `design/mocks/2026-10-31-mobile-wave3/renders/f9-guides-{ios,and}`.

## Static files (W3-C1)

Written at site build time by `web/src/lib/app-data` through Astro endpoints under `web/src/pages/data`. Every block
is composed by the same `lib/guides` and `lib/tiers` functions the web's guide page and `/tiers` call, so the apps
and the site never disagree.

| File | What it carries |
|---|---|
| `/data/{build}/guides/index.json` | Classes in `classes.json` order, specs in tree order, each with role, icon, updated day and its file |
| `/data/{build}/guides/{class}-{spec}.json` | Summary, role line, races, the web's stat pill order, hit cap, overview lists, key talents with the game's top-rank text, leveling and raid build totals, opener/priority/cooldowns with ranks and icons, race table, the BiS list the Gear row opens, sources, and the sections the guide has |
| `/data/{build}/tiers.json` | Both factions, every role: the web's rank, gap, bar, tie, letter and the callout sentences, plus the notes |
| `/data/changelog.json` | The changelog newest first, bodies as plain `p`/`h`/`li` blocks |

- Budgets: the site build fails a guide over 60 KB gzipped, tiers over 40 KB, the index over 10 KB (largest real
  files: a guide 3.3 KB, tiers 5 KB, the index 1 KB).
- OpenAPI schemas and `app-static` operations, fixtures `cdn.guides.index`, `cdn.guides.{shaman-elemental,mage-fire,warrior-protection}`,
  `cdn.tiers`, `cdn.changelog` (written by the generator, held to the schema by the Go contract test and to the
  generator by `app-files.test.ts`), the contract doc section, generated Swift and Kotlin models.
- Fixes to shared web code found on the way: a rotation clause no longer ends inside an unclosed parenthesis; the
  guide cooldown list never names a long-cast utility spell (Astral Recall was a "15 min cooldown" on the web too).
- The key talents add the raid-build talents the rotation casts (Elemental Mastery), per spec F9.2. A spell the
  client gives no icon reaches the apps as `null` (neutral tile), not the client's question-mark art.

## Apps

| | iOS (`apps/ios/App/Guides`) | Android (`app/.../guides`) |
|---|---|---|
| Logic | `GuidesLogic`, `ChangesLogic`, tier models: grid, saved and footer, sections, rotation rows, tier rows and labels, callout by character, key dates, changelog tags, new-data decision | `logic/GuidesLogic.kt`, `TiersLogic.kt`, `ChangesLogic.kt`, the same |
| Data | `GuidesStore` over the existing `CachedStaticDataClient` (one offline cache with BiS and Talents); prefetch of the player's classes | the same |
| Screens | Guides tab (fifth tab), class page, guide reader with sticky section strip, tier lists, What changed and the changelog reader | the same |
| Shared pieces | `NewDataStrip` (for F15 to extend), `SegmentBar` | `NewDataStrip` in the design system |

- Switches: `guides` (the tab and the links), `guide_tiers` (the Tier lists segment). Both default false.
- Links: `/app/guides`, `/app/guides/{class}/{spec}`. The links generator renames reserved-word params (`className`)
  and Swift now matches routes from a list rather than one `??` chain the type checker could no longer finish.

## Reviews

Captures (default and largest text, both platforms; the apps are dark-only) in
`design/reviews/2026-11-01-mobile-guides/build/{ios,and}/`.

- Round 1: ux-designer HOLD, wow-player HOLD. Fixed: the missing icons (Fire Nova, the no-spell row), Astral Recall
  as a cooldown, Elemental Mastery missing from the talents, the doubled Name reservation line, iOS largest-text
  chips, tier chip and segments, the tier chip shape. Ruled: race rows keep the faction mark (spec F9.2); Mana
  Spring Totem's five-minute life is printed because the guide and the client data agree (spec F9.2's own rule).
- Round 2: ux-designer SHIP (`ux-verdict-build.md`), wow-player SHIP (`player-verdict-build.md`). Both made
  Mana Spring Totem's duration conditional on the accuracy reviewer confirming Forever ships 300 s, not Era's 60 s;
  if not, the line becomes "Down first. It returns mana for the rest of the fight."
- Open SHOULDs: Fire Nova's neutral tile (wow-player would draw the Fire Nova Totem art; not done, tenet: never
  re-badge an icon), mixed date formats, the 10% hit cap, the "Nature power" and "Spell haste" stat names, the race
  ability text and Fire Nova's condition without its Fire totem: data and accuracy items for the guide owners.

## Gates

On the tree merged with origin/main (F16 guild extras), each run on its own with its exit code checked:

- `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts src/lib/app-data src/lib/guides src/content`: 711 passed, exit 0 (codegen drift, the W3-C1 generator against its fixtures, the guide libs).
- `go test ./internal/contract/`: exit 0 (every fixture against `api/openapi.yaml`). No API handler or migration in this lane.
- `make -C apps/ios test`: TEST SUCCEEDED, exit 0 (574 app tests plus the package suites).
- `./gradlew test :app:verifyRoborazziDebug lint assembleDebug`: exit 0.
- `astro build` emits every W3-C1 file under its budget (checked before the merge).

## Turning it on

Add `guides` to `APP_FEATURES` on the API deploy, after the site deploy that publishes the W3-C1 files. Leave
`guide_tiers` off until the accuracy reviewer verifies every DPS band (spec F9.4: Elemental 20th, Frost Mage and
Balance Druid far behind on build 70658).

## Gaps and follow-ups

1. **Search button on the Guides root bar** waits for F12 search on main (F15 or F12's merge wires it).
2. **Raid build Compare and Open this build in the planner** rows: F6 compare and F13 planner.
3. **Fire Nova has no icon in the client data** (`IconFileDataID` 0 for the Forever ranks); the pipeline should
   resolve one, then the neutral tile goes away by itself.
4. **`tiers.json` is published** with the site (the web's `/tiers` already shows the same ranking); only the app
   segment is gated, which differs from the spec's "lane F9 does not ship tiers.json".
5. **The rotation footer's extra guide note** ("Lava Burst is new in Forever.") needs a guide field.
6. **Not recorded here:** a VoiceOver and TalkBack pass, the haptics unit test, the on-device open-from-cache budget,
   and old builds' files deleted 7 days after a swap (the cache keys by build path; nothing prunes yet).
7. **iOS snapshots do not draw the iOS 26 glass toolbar items** (the guide's share button, the root selector's
   trailing items); they exist in code.
8. **F12's `StaticDataStore`** duplicates the `CachedStaticDataClient` this lane uses; merge them into one when F12 lands.
