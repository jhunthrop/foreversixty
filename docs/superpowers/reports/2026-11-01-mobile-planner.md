# Mobile Wave 3 F13: The talent planner (2026-11-01)

One lane, both platforms, the API change and the shared rules. Spec: `design/specs/2026-10-31-mobile-wave3.md`
§5.5 and §10 (W3-C9 to W3-C11). Boards: `design/mocks/2026-10-31-mobile-wave3/renders/f13-planner-{ios,and}`.

## Shared pieces

| Piece | Where |
|---|---|
| Feature switch `planner` (default false) | `apps/shared/features.json`, generated into both apps |
| Copy `plan.*` (editor, refusals, share confirm, viewer, Your builds) | `apps/shared/copy/en.json`; the refusals are the site planner's own `messages` (`web/src/lib/planner/rules.ts`), plus the app's level-budget line `No points left. Take one off first.` |
| Golden vectors | `apps/shared/fixtures/planner-rules.json`, written by `web/scripts/gen-planner-rules.mjs` (bundles `web/scripts/mobile-codegen/planner-golden.ts` with esbuild) from the site planner's `canAddPoint`, `canRemovePoint`, `validateOrder`, `pointsPerTree`, `pointsAvailableAt`, `decodeFS1` and `orderFromRanks` over the active build's shaman, warrior and druid trees. Seeded walks plus a targeted prerequisite case per talent; the generator fails if a refusal kind goes missing. `src/lib/mobile/planner-rules-codegen.test.ts` fails CI on drift |
| Deep links (W3-C11) | `apps/shared/links.json`: `/app/planner/{class}?code=&race=` (`planner`), `/app/builds` (`builds`), and `/b/{id}` (`sharedBuild`), the first route outside `/app/`. A route marked `site: true` is a real site page; the generator emits it as `site_pages` in `web/app-link-paths.json`, and `web/scripts/app-links.mjs` admits exactly those besides `/app/*` (`web/app-links.json` paths `["/app/*", "/b/*"]`, the AASA regenerated) |

**F11's planner link.** Rankings' `Open in planner` opens `DeepLink.planner(className: <class slug>, code: nil, race: nil)` on iOS
and `DeepLink.Planner(className, null, null)` on Android: the editor on that class (its saved draft, else empty).
`builds_seen` carries only a tree split, which is not a build, so there is no code to pass. A row that ever carries an
FS1 code passes it as `code` and the editor loads it as one legal order.

## API

- **W3-C9 `DELETE /v1/builds/{id}`** (`api/internal/builds`): session or device token; `204`; `404 not_found` when the
  build is not on the caller's list (another player's, nobody's, unknown), never `403`. It clears `user_id` on the
  caller's own row only: the record and its `/b/{id}` link stay, and saving the same build again puts it back on the
  list (the apps' Undo). `builds.user_id` is already nullable (migration 0015), so **no migration**. Store, handler
  and contract tests; OpenAPI `removeBuild`; fixture `builds.delete.error-not-found` (a 204 has no body, so there is
  no `builds.delete` fixture, the convention every other 204 follows).
- **W3-C10** `/data/{build}/races.json` and `/data/{build}/combos.json` as `app-static` operations (`RaceInfo`,
  `RaceClassCombo`; `ForeverChange` hoisted to a shared schema for classes and races), fixtures `cdn.races`,
  `cdn.combos` held to the generator's files by `cdn_test.go`.
- Contract doc: `docs/contracts/app-v1.md` (DELETE, the planner's rules, races and combos).
