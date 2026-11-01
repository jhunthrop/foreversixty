# Mobile Wave 3 F12: Search and the item page (2026-10-31 to 2026-11-01)

Lane `mobile-search`, branch `wt/mobile-search`, with `wt/mobile-search-ios` and `wt/mobile-search-and` merged in. Spec: `design/specs/2026-10-31-mobile-wave3.md` §5.4, and contract W3-C8 in §10. Boards: `design/mocks/2026-10-31-mobile-wave3/renders/f12-search-{ios,and}.png`. Feature switch: **`search`** (`apps/shared/features.json`), default off.

## Data gate (spec C6)

- **Binds line.** The tooltip prints the client's Binds line from `bonding`:
  - A client row takes it from its own ItemSparse `Bonding`.
  - An item the client's ItemSparse lacks takes it from cmangos classic-db's `item_template.bonding`, which uses the same enum.
  - A wowhead row with no binding takes the same id's classic-db value.
  - `data.yml` re-reads the pinned classic-db dump before `normalize`. The same commit is used, so every other field is unchanged.
- **Armour word.** A cloak's `armor_type` is `cloth`, so it prints `Back` with `Cloth` on the right.
- **Coverage.** The active build 1.60.1.70707 carries both fields: 97% of rows state a binding. The 3% with none are wowhead-only rows with no classic-db twin; they print no line rather than a guess.
- **Floor test.** `TestCDNItemsCarryTheBindsLine` holds that share to a 90% floor, and `bonding`/`armor_type` left the pending list.
- **Spritecaster Cape.** It reads `Binds when picked up` from classic-db. The board's "equipped" was a guess.
- **Clients.** All three treat both fields as optional: old data draws no line.

## Contract W3-C8 (static, emitted by the site build)

- **`/data/{build}/search.json`.**
  - Contents: every item the class files carry, with name, icon, quality, slot, levels and usable classes; every talent; every written guide.
  - Built by `web/src/pages/data/[build]/search.json.ts` over `web/src/lib/search/search-index.ts`.
  - The site build fails over 300 KB gzipped. It is 215 KB on 70658.
- **`/data/{build}/item-uses/{shard}.json`.**
  - Split into 64 shards by item id.
  - Each item lists its sources and one row per spec at its best placement (Pick or Runner-up n).
  - Built by `web/src/lib/search/item-uses.ts`.
  - Sources are `web/src/lib/items/item-sources.ts`' structured rows. The site tooltip now reads the same rows, so a source is worded one way everywhere. A boss drop names every boss that drops it.
- **`/data/{build}/loot.json`.** Documented as an app contract.
- **Contract files and tests.**
  - OpenAPI paths and schemas, plus fixtures `cdn.search`, `cdn.item-uses` and `cdn.loot`.
  - `docs/contracts/app-v1.md` sections.
  - `TestCDNLootFixtureMatchesTheGenerator`.
  - `web/src/lib/search/contract-fixtures.test.ts`, which holds the generator to the fixtures.
  - Generated Swift and Kotlin clients.
- **No API endpoint was added.** Character search uses the existing `GET /v1/characters/search`.

## Tooltip accuracy (both apps and the site)

- **Equip sentences.** The Classic client prints spell power, healing, attack power, ranged attack power, mp5, spell penetration, fire power and block value only as green `Equip:` sentences.
  - One rule is keyed on the stat id in `web/src/lib/items/tooltip.ts` and in both apps' `ItemTooltipModel`.
  - One "Increase Spell Dam" spell, with equal spell power and healing, prints as one sentence.
- **Resistances** use the client's casing: `+10 Shadow Resistance`.

## Apps

| | iOS | Android |
|---|---|---|
| Static data store (download, swap, delete, size) | `Sources/.../StaticDataStore` | `core/network/.../StaticDataStore.kt` |
| Search | `App/Search/` | `app/.../search/` |
| Item page | `ItemPageScreen` | `search/ui/ItemPage.kt` |

- **Store.** `ItemCatalog` reads through the store. Sign-out clears the store and the Recent list.
- **Search button.** It appears on the Home, Raids and BiS root bars. F3's old filter gives way to it when `search` is on.
- **Item page entry.** F2.6's tooltip sheet gains `More about this item`, which opens the item page.

## Reviews

Captures: `design/reviews/2026-10-31-mobile-wave3-search/build/{ios,android}/`. Verdicts: `ux-verdict-build.md`, `player-verdict-build.md`.

## Gaps

- **Watching and the star on character rows** wait for F11's `/v1/me/watching`.
- **Characters, guides, Rankings and the planner** open the website, marked with an external-link glyph, until F9, F11 and F13 land.
- **Vendor price and standing** are not shown: C6 `buy_price`/`min_standing` are not published.
- **Budgets and screen-reader passes.** The §11 search budgets were not measured on device. No VoiceOver or TalkBack recording was made.
- **iOS F12.5 capture.** It is a full-page capture with the chip on, because a scrolled snapshot rendered blank.
