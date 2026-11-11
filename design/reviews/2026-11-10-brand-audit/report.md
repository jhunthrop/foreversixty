# Cross-surface brand audit: 2026-11-10

**Goal (owner, 2026-11-10):** reusable components and one seamless brand on every surface. The website is the reference.

## Base

`origin/main` at `fe8796eac`.
- The web unification lanes are all merged into main: A (icons, `ui-icons`), B (formats, `lane/unify-format`) and C (server and companion, `lane/unify-server`). Each is 0 commits ahead of main, so this audit runs after J3.
- `lane/app-brand-polish` is not merged; it is 5 commits ahead.

## Method

- Seven `brand-consistency-reviewer` runs (`.claude/agents/brand-consistency-reviewer.md`), in two waves of at most four.
- Each run compared one surface, side by side, with the web renders in `design/reviews/2026-11-10-unify-format/after/*-{390,1440}.jpg`, and read the code and token sources.
- The verdicts are in this folder as `*-brand-consistency-reviewer.md`. Every finding below cites a capture or a file:line from them.
- New captures are in `captures/`:
  - `identity-sheet.png`: every seal, icon, tray, OG and Discord asset on one sheet.
  - `email-inbox.png`: every transactional mail as a mail client shows it. The bodies were dumped from `billingmail` and from `auth/handler.go:268`.
- Everything else uses the 11-09 and 11-10 review captures, cited by path.
- No fresh simulator or emulator runs, because the laptop is shared. The stale sets are listed at the end.

## Result

Every surface returned `VERDICT: DRIFT`. None reads as a different product end to end. The drift sits in three places:
1. The addon's palette.
2. The identity assets.
3. Everything that hand-writes copy, formats or colours outside the generators.

## Summary

| Surface | BLOCKER | SHOULD | NICE | Top 3 issues |
|---|---|---|---|---|
| Addon (in game) | 2 | 6 | 5 | `Theme.HEX` is navy where the web is umber; raw client rarity colours on the dark panel; Friz and Morpheus fonts with caps eyebrows, where the web is Marcellus and Open Sans in sentence case |
| Store, listing, identity | 2 | 4 | 1 | A third seal (black disc) on the app icons on main and on the Discord server icon; no App Store or Play listing at all; a stale CurseForge and Wago gallery showing the navy addon |
| Android app | 3 | 7 | 0 | Retired styles live on about half the screens (65 `quietCard`, 57 `.uppercase(`); the Home capture is broken or stale; the Raid screen lacks the guild-links block |
| Email | 2 | 3 | 2 | Text only, with no seal, wordmark or footer; sends from `hello@` while the site promises `help@`; copy and dates hand-built in Go |
| iOS app | 0 | 6 | 3 | 39 `uppercased()` titles; quality and crest glows the web doesn't have; 148 SF Symbols while the generated icon set goes unused |
| Discord bot | 0 | 5 | 4 | Server icon, bot avatar and banner disagree; the raid title reads "Thursday 10 December" where the web reads "Wed 13 Jan"; unringed crest and item thumbnails |
| Companion | 0 | 4 | 5 | Tray strings hardcoded, and one job has three names; formats.json reimplemented twice and never read; copy not keyed to apps/shared/copy |
| Installers | 0 | 1 | 0 | A light, baked DMG background under a dark window; Windows setup is on brand |
| Server pages (/b, /app/auth) | 0 | 2 | 6 | Stale app-confirm captures with no seal visible; a local copy struct; hardcoded `#000`, a shadow and the 560 px column |
| Push copy (API) | 0 | 2 | 1 | `pushCopy` is a Go struct that duplicates `home.next.simFinished`, and the same words are written again in `notify/embed.go` |

## What a player notices first

1. **The addon looks like an older product.** It is the surface players see every session.
   - The panel is navy and dull gold; the site is umber and Blizzard gold (`addon/ForeverSixty/Theme.lua:15-98` against `web/src/styles/tokens.css:11-27`).
   - Rare item names are dark 0070dd blue on a dark ground (`Theme.lua:1259`).
   - Headings are set in Morpheus.
2. **The phone home-screen icon is a mark the site never uses.**
   - It is the black disc with a dotted ring: `captures/identity-sheet.png`, and `apps/shared/brand/ios/Brand.xcassets/AppIcon.appiconset/AppIcon-1024.png`, last changed in `44349ebfb` on 10-30.
   - The Discord server icon is the same disc.
   - The fix is already built on `lane/app-brand-polish` (`c669fb8a4`: the painted seal on the painted ground). It only needs to land.
3. **Android outside Home, Raid and Gear is still in the pre-redesign style.** Guild, Rankings, Reports, Planner and Guides use the retired cards, micro-labels and forced caps (`android-brand-consistency-reviewer.md` BLOCKER 1; spec `design/specs/2026-11-07-mobile-redesign.md:165`).
4. **Billing mail is a bare text notice from an address the site doesn't list** (`captures/email-inbox.png`, `api/internal/mail/resend.go:33`, `api/internal/config/config.go:16`).
5. **iOS has forced caps and glows, and the two apps use different sign-up controls.**
   - Caps and glows: `apps/ios/App/.../RaidPageView.swift:228`, `ItemIcon.swift:113`.
   - Sign-up: four buttons on iOS, a segmented pill on Android (`m2i/side-by-side-raid.png`, `m2a/side-by-side-raid.png`).

## Per surface

### Addon (in game)

Verdict file: `addon-brand-consistency-reviewer.md`.

Captures:
- `design/listing/gallery/01..12-*.png`
- `design/reviews/2026-11-08-addon-ux/panel-raider-60.png`
- `design/reviews/2026-11-10-addon-onboarding/onboarding-*.png`
- `design/reviews/2026-11-10-seal-flat/addon-minimap-button.png`

**BLOCKER: palette.**
- `Theme.HEX` says it is "the site's palette, web/src/styles/tokens.css" (`Theme.lua:17`).
- About 11 keys are wrong and about 14 have no token:

  | Key | Lua | Web |
  |---|---|---|
  | background | 0d111a | 100a06 |
  | gold | e5b955 | ffec71 |
  | border | 262e40 | 3d342d |
  | teal | 2f7a74 | 0f6064 |

- The umber, brass and bone frame keys exist nowhere in `design/tokens/tokens.json`.
- The full table is in the reviewer file.

**BLOCKER: rarity text.**
- `Theme.qualityColor` returns the client's `ITEM_QUALITY_COLORS` (`Theme.lua:1259`).
- The web uses `rarity.rare-text` 3d94f0 and `epic-text` b866f5.
- The addon's own "bright" lift (`Theme.lua:62-68`) matches no token.

**SHOULD:**
- Fonts are FRIZQT and MORPHEUS (`Theme.lua:411-414`), and no Marcellus or Open Sans ships in `media/`.
- At least 13 caps eyebrows in `Locale.lua` (lines 120, 196, 245-251 and 381-422).
- Buttons are client red in `onboarding-*.png`, where the web primary is teal.
- Bevelled frames where the web uses hairline plates.
- "1840" is shown with no thousands separator.
- The listing plates are stale.

**What is right:**
- Hardcodes sit almost entirely in `Theme.lua`. There are 0 numeric `SetTextColor` or `SetVertexColor` literals; outside `Theme.lua` there are only `surfaces/Panel.lua:822-823` and `surfaces/TourExample.lua:28,35`.
- The minimap uses the flat seal at 20 and 32 px, which is correct.
- Ringed crests ship for nine classes (`media/crests`).

### Store, listing and identity

Verdict file: `store-identity-brand-consistency-reviewer.md`. Capture: `captures/identity-sheet.png`.

**BLOCKER: three seals coexist.**
- Painted v4: `web/public/apple-touch-icon.png`, `seal-400.png`, `addon/ForeverSixty/icon.png`.
- Flat turquoise: `web/public/favicon.svg`, the tray icons, `companion/ui/seal.svg`.
- Black disc with a dotted ring: the iOS AppIcon (light), Android `ic_launcher_foreground.png`, `design/community/discord/assets/icon.png`.
- No ruling covers the disc. `design/specs/2026-11-09-directions.md:69` says the app icon is the flat seal in the hammered-gold frame.
- The three iOS variants (light, dark, tinted) are three different marks.

**BLOCKER: no App Store or Play listing.** There is no fastlane, no metadata and no store screenshots anywhere in `apps/`.

**SHOULD:**
- `design/community/discord/assets/banner.png` uses forced, tracked caps in a Cinzel-like face. `design/assets/art/discord-banner-2026-11-09.jpg` (painted seal, Marcellus, sentence case) matches the web.
- `design/listing/gallery/*` was last committed on 2026-11-03 (`6b4e85898`). It shows the navy addon and "(needs bullets)". `design/listing/RECAPTURE.md` itself asks for lossless in-game captures.

**What is right:** OG passes. `web/src/lib/og.ts` puts the painted seal, the Marcellus wordmark and token colours on every per-page card.

### Android app

Verdict file: `android-brand-consistency-reviewer.md`.

Captures:
- `design/reviews/2026-11-10-m0/*`, `m1a`, `m2a`, `m3a`, `m4a-m5a`.
- `design/reviews/2026-10-31-mobile-wave3-widgets/build/and-*.png` (stale).

**BLOCKER: retired styles.**
- `quietCard` 65, `primaryCard` 22, `surfaceBand` 12, `MicroLabel` 62, `.uppercase(` 57.
- Worst: `raids/ui/RaidPage.kt:167,223,300,350` and `core/designsystem/.../Surfaces.kt:49,66`.
- `home/GearScreen.kt:34` is dead code.

**BLOCKER:** `m0/android-home-34.png` shows empty icon squares and a bare person glyph.

**BLOCKER:** the Raid screen lacks the guild-links block and search (`m2a/side-by-side-raid.png`).

**SHOULD:**
- 30 hardcoded colours.
- `res/values/notification.xml:5,8` and `values/themes.xml:10,12` are not generated. For example, `#FF2C8A9E` is used for the accent.
- 228 Material icon references and about 40 hand-written `ic_*.xml`. `gen-ui-icons.mjs` emits only `ic_tab_*` and `ic_frame_corner`.
- The live-raid lock-screen card is light grey with a black timer; iOS is dark with a gold timer (`m4a-m5a/side-by-side-lock.png`).
- Widgets default to `FontFamily.SansSerif` (`widgets/glance/WidgetDesign.kt:56`).

### Email

Verdict file: `email-server-brand-consistency-reviewer.md`. Capture: `captures/email-inbox.png`.

**BLOCKER:** every mail is `text/plain` only (`api/internal/mail/resend.go:33`). Money mail has no seal, wordmark, button or legal footer.

**BLOCKER:** the sender default is `hello@foreversixty.gg` (`api/internal/config/config.go:16`). The site and the shared copy use `help@foreversixty.gg` (20 uses, against 5 for `hello@`).

**SHOULD:**
- The bodies are hand-built Go sentences (`api/internal/billingmail/billingmail.go:89`), not the account page's "Premium · renews {date}" (`apps/shared/copy/en.json:783`).
- The long date from `EmailDate` (`billingmail/price.go:60`, Pacific time) has no `formats.json` case.

The reviewer file proposes a minimal branded shell: one `render.go` reading `api/internal/brand/tokens_gen.go`, with a hosted flat seal, a teal button and the web footer line.

### iOS app

Verdict file: `ios-brand-consistency-reviewer.md`.

Captures:
- `design/reviews/2026-11-10-m1i..m5i/*`
- `restyle-app/*-ios.png`
- `m5i/ios-live-raid-night.png` (Live Activity)

**SHOULD:**
- 39 `uppercased()` calls. The worst is `RaidPageView.swift:228`; `RaidClock.swift:42` gives "WED" and "JAN".
- Glows at `ItemIcon.swift:113`, `RaidNightView.swift:222` and `BrandImages.swift:43`.
- 148 `systemName` icons, while the generated `Brand.xcassets/Vectors` go unused.
- 0 references to `apps/shared/formats.json`, and 47 local `DateFormatter` or `String(format:)` calls.
- 100 font-size literals.
- `gear.worn` reads "Equipped" while a related string says "over your worn" (`en.json:703`).

**What is right:** copy is fully generated, with 1835 `Copy.` uses and no literal sentences.

### Discord bot

Verdict file: `discord-brand-consistency-reviewer.md`.

Captures:
- `design/reviews/2026-11-10-unify-server/after/discord-*.png` (current).
- `design/reviews/2026-11-09-f-discord/discord/*` (stale for formats).

**What is right:** colours are fully tokenised. There are 0 colour ints, and every colour goes through `brand.RGB(brand.Color*)`.

**SHOULD:**
- The identity assets disagree; see the store section.
- The post title uses `textfmt.LongDay` (`cmd_report.go:137`), while the web raid row uses `RaidDay`.
- "When" and "Guild time" are two fields where the web shows one cell.
- The item thumbnail has no quality ring (`item_embed.go:60`), and the character thumbnail has no ring (`char-card-own-desktop.png`).

**NICE:**
- An ASCII ellipsis `"..."` at `card/card.go:322`.
- Hand-built plurals at `cmd_report.go:206`.
- All copy lives in `copy_*.go`.

### Companion and installers

Verdict file: `companion-brand-consistency-reviewer.md`.

Captures:
- `design/reviews/2026-11-10-unify-server/after/companion-*.png` (current).
- `2026-11-09-companion/round-4/*` (stale).
- `2026-11-10-seal-flat/tray-*.png`.
- `2026-11-10-companion-installers/*.png`.

**What is right:**
- `ui/tokens.css` is generated, and `assets_test.go:95` is a hex test.
- `ui/icons.js` is generated.
- The fonts are the site's.

**SHOULD:**
- The tray strings are hardcoded (`internal/shell/shell_gui.go:94-97`). The reports job has three names: "Open my reports on the site", "Open all my reports on the site", and the web's "Your reports".
- `copy.go:78-88` and `ui/format.js` both reimplement `formats.json`, and neither reads it.
- 0 shared copy keys.
- The DMG background is a light, baked bitmap (`mac-dmg-window-dark.png`, `packaging/macos/dmg-background*.png`).

**NICE:**
- 24 literal font sizes and 4 literal radii in `ui/styles.css`.
- rgba veils and a text shadow with no token (`styles.css:29,97,110,129`).
- `border: 1px solid black` at `styles.css:447`.
- The lead panel's corner filigree is the most decorative element in the product. It is an owner call.

The Windows setup (`win-1-setup.png`) matches the web setup page (`page-setup.png`). The light title bar comes from the OS and is allowed.

### Server pages and push

Verdict file: `email-server-brand-consistency-reviewer.md`.

Captures:
- `design/reviews/2026-11-10-h-api-html/after/*`
- `design/reviews/2026-11-10-unify-server/after/{b-*,app-confirm-*}`

**Server pages:**
- `/b` (`api/internal/site/page.go:130-137`) renders from the web build's variables and is close to the web.
- The app-confirm captures predate the current template (`auth/app_confirm.go:48-51`). They show no seal, although `:45` renders one, so they need a recapture.
- `app_confirm.go` hardcodes `#000` at `:141` and `:147`, and a text shadow at `:145`.
- The 560 px column is hardcoded in both `app_confirm.go` and `page.go`.

**Push:**
- The push copy is a Go struct (`notify/events.go:79-90`) that duplicates `home.next.simFinished`.
- It is written a second time in `notify/embed.go:75-86`.

## Shared design sources

### What each surface reads from a generator, and what it hardcodes

| Surface | Tokens (`design/tokens/build.mjs`) | Icons (`web/scripts/gen-ui-icons.mjs`) | Copy (`apps/shared/copy`) | Formats (`apps/shared/formats.json`) |
|---|---|---|---|---|
| Web | `web/src/styles/tokens.css` (generated) | `src/lib/ui/icons/*` (generated) | web copy modules (the source) | replayed in `web/src/lib/formats.test.ts` |
| iOS | `DesignTokens.swift` (generated); 5 raw colours, 97 raw fonts, 100 size literals | Vectors generated but **unused**; 148 SF Symbols | `App/Generated/Copy.swift` (generated) | **not read**; 47 local formatters |
| Android | `Tokens.kt` (generated); 30 raw colours, 61 `.sp`; `notification.xml` and `themes.xml` **not generated** | 5 generated, about 40 hand-written `ic_*.xml`, 228 Material | `strings.xml` (generated) | partly: `Formats` in 3 files, 89 local formatters |
| Companion | `ui/tokens.css` (generated); 24 font sizes, 4 radii, rgba veils | `ui/icons.js` (generated) | **by hand**: `ui/copy.js` (265 lines), `internal/app/copy.go` (307 lines), tray literals | **reimplemented**: `format.js`, `copy.go:78` |
| Go (Discord, /b, confirm, push, mail) | `api/internal/brand/tokens_gen.go` (generated); card font sizes 26-72 px (`card/card.go:110-116`); `#000` in `app_confirm.go` | emoji synced from app assets; crest PNGs by hand | **by hand**: `discordbot/copy_*.go`, `notify/events.go`, `billingmail.go`, `app_confirm.go:50` | replayed in `api/internal/textfmt/textfmt_test.go`, except mail's `LongDate` |
| Addon (Lua) | **no target**; `Theme.HEX` is hand-copied and stale | **no target**; client textures plus `media/*.tga` | **by hand**: `Locale.lua` | **by hand** |
| Email | **none** (text only) | **none** | by hand (Go) | `textfmt.LongDate` (no fixture) |
| Identity assets | n/a | `design/logo/build.py` for marks and favicon; `og.ts` for OG cards | n/a | n/a |

### Hardcodes that should come from a source

All of these come from the reviewer files and were checked by grep.

**Colour:**
- Addon: `addon/ForeverSixty/Theme.lua:15-98` (all of `Theme.HEX`).
- Addon parse grey: `surfaces/Panel.lua:822`, 9d9d9d against `parse.greyText` 8c8c8c.
- Android: `apps/android/.../designsystem/GameColors.kt:45`, `paperdoll/Paperdoll.kt:165`, `guild/platform/AndroidCalendar.kt:116`, `res/values/notification.xml:5,8`, `res/values/themes.xml:10,12`.
- iOS: `apps/ios/.../ReportView.swift:219-222`, `EventKitCalendarAdapter.swift:36`.
- Companion: `companion/ui/styles.css:29,97,110,129,447,537,563`.
- Go: `api/internal/auth/app_confirm.go:141,145,147`, `api/internal/sheet/palette.go:23`.

**Font:**
- `addon/ForeverSixty/Theme.lua:411-414` (client fonts).
- `apps/android/.../ScaledType.kt:36-40` (system mono).
- `widgets/glance/WidgetDesign.kt:56`.

**Size and spacing:**
- `companion/ui/styles.css`: font sizes at `:93,126,132,203,213,273,283,349,460`, radii at `:306,309,461,563`.
- `api/internal/card/card.go:110-116`.
- The 560 px column in `api/internal/site/page.go` and in `app_confirm.go`.
- iOS: 100 font-size literals, for example `HomeFlowView.swift:183`, `NextRaidBlock.swift:215`, `ItemPageView.swift:81,321,429`.
- Android: 61 `.sp` literals.

**Icon:**
- iOS: 148 SF Symbols (`chevron.right` ×17, `magnifyingglass`, `person`, `flask`, `lock`).
- Android: 228 Material references (for example `Roster.kt:21-25`, `Surfaces.kt:65`, `HomeScreen.kt:204`) and about 40 `res/drawable/ic_*.xml`.
- Addon: the `UI-Classes-Circles` crest fallback.
- Discord: `crest-*.png` exported by hand, 9 classes against the addon's 10 in `media/crests`.

**String:**
- Companion: `companion/internal/shell/shell_gui.go:94-97`, `companion/ui/copy.js`, `companion/internal/app/copy.go`.
- Go: `api/internal/discordbot/copy_*.go`, `api/internal/notify/events.go:79-90`, `notify/embed.go:75-86`, `api/internal/billingmail/billingmail.go:80-131`, `api/internal/auth/handler.go:268-272`, `auth/app_confirm.go:50`.
- Addon: `addon/ForeverSixty/Locale.lua`.
- The mail sender: `api/internal/config/config.go:16`.
- The em-dash glyph at `apps/shared/copy/en.json:1183`.

**Format:**
- `apps/ios`: 47 local formatters, for example `RaidClock.swift:42`.
- `apps/android`: 89 local formatters.
- `companion/ui/format.js` and `companion/internal/app/copy.go:78-88`.
- `api/internal/billingmail/price.go:60`.

### Missing generator targets, and how to add each so every surface reads one source

**Lua tokens.**
- Add `lua: addon/ForeverSixty/TokensGen.lua` to `OUTPUTS` in `design/tokens/build.mjs`. It returns `{ color = {...}, font = {...}, radius = {...} }`, with colours as six-hex strings.
- Load it in the .toc before `Theme.lua`.
- `Theme.HEX` becomes a thin key map onto it, so no caller changes.
- Promote the frame keys that stay (`umber`, `brass`, `bone`) into `tokens.json` with a `lua` extension, and delete the navy keys.
- `tokens.test.mjs` fails on a stale file, and a busted test fails on any hex literal in `addon/**/*.lua` outside `TokensGen.lua`.
- Files: `design/tokens/*`, `addon/ForeverSixty/{TokensGen,Theme}.lua`, the `.toc`.

**Android XML colours.**
- `build.mjs` also emits `core/designsystem/src/main/res/values/colors_gen.xml`.
- `notification.xml` and `themes.xml` then reference `@color/fs_*`.
- Files: `design/tokens/build.mjs`, `apps/android/.../res/values`.

**Email.**
- Add an `HTML` field to `mail.Message`.
- One `api/internal/mail/render.go` shell reads `brand.Color*` and `brand.Font*` from the existing Go target, so no new generator is needed.
- It uses a hosted flat seal PNG and the web footer line.
- The bodies come from shared copy keys (next item).
- Files: `api/internal/mail`, `billingmail`, `auth/handler.go`.

**Copy for Go, the companion and Lua.**
- Generalise `web/scripts/gen-mobile-copy.mjs` into a shared-copy generator with more targets.
- Add `mail.*`, `push.*`, `discord.*`, `companion.*` and `confirm.*` namespaces to `apps/shared/copy/en.json` and `shared.json`.
- Emit `api/internal/copygen/copy_gen.go`, `companion/internal/app/copy_gen.go`, `companion/ui/copy_gen.js` and `addon/ForeverSixty/LocaleGen.lua`.
- Move the strings namespace by namespace, so `copy_*.go`, `pushCopy`, `copy.js`, the tray literals and `Locale.lua` shrink to lookups.
- Files: `web/scripts/gen-mobile-copy.mjs`, `apps/shared/copy/*`.

**Formats for the apps, the companion and the addon.**
- No generator is needed: `formats.json` is a fixture file. Each surface replays every case in a test, as the web and Go already do.
- Add `FormatsReplayTests.swift` and `FormatsReplayTest.kt`.
- `companion/ui/format.test.js` and `copy_test.go` assert against the file itself, not copies of it.
- Add a `long_date_year` case for mail.
- The addon gets a busted replay of the `numbers` and `dates` cases it prints.

**Icons for the addon and Discord.**
- Add an addon target to `gen-ui-icons.mjs` that rasterises the needed glyphs to `addon/ForeverSixty/media/icons/*.tga`: 64 px, white, tinted in Lua.
- Discord already syncs app emoji; feed `emoji_sync.go` from the same rasterised set.
- On iOS, wire up the existing `Brand.xcassets/Vectors` through an `FSIcon` view.
- On Android, extend the target beyond the tab icons.
- File: `web/scripts/gen-ui-icons.mjs`.

**Crests and seals.**
- One export script writes every surface's crest and seal sizes from the HD sources:
  - Discord `assets/crest-*.png`, adding the tenth crest;
  - addon `media/crests/*.tga`;
  - the app assets;
  - the seal variant by size: painted at 64 px and up, flat at 44 px and under.
- `design/logo/build.py` is the natural home.
- The Discord `icon.png` and `banner.png` and the app icons become outputs, not hand exports.
- Files: `design/logo/build.py` and `web/scripts/build-app-icon.mjs` (on `lane/app-brand-polish`).

**Tray icons.**
- `companion/tools/icons/render.py` reads `--tone-*` from `companion/ui/tokens.css` (generated), so the dot colours are bound to tokens.
- Add a test.
- File: `companion/tools/icons`.

### Components duplicated across surfaces

Code can't be shared across Lua, Swift, Kotlin, Go and Svelte. Each component instead gets:
- one visual spec in `design/specs/components/`;
- one fixture in `apps/shared/fixtures/`, mapping inputs to the exact strings and states to render.

Each surface replays the fixture in its own tests, the way `planner-rules.json` already works.

**Item line and tooltip.**
- Where it lives:
  - web: `web/src/components/ItemTooltip.svelte`
  - iOS: `apps/ios/App/Items/ItemTooltipModel.swift`, `ItemTooltipSheet.swift`
  - Android: `items/ItemTooltipModel.kt`, `ItemTooltipSheet.kt`
  - addon: `surfaces/ItemTooltip.lua`
  - Discord: `discordbot/item_embed.go`
- Drift today:
  - Android drops "drop chance" and the raid day, and uses a straight apostrophe (`m3a/side-by-side-item.png`).
  - iOS adds "Equipped" and a glow.
  - The addon colours names by raw client rarity.
  - Discord has no icon ring.
- Spec:
  - a quality-bordered icon with no glow;
  - the name in `rarity.*-text`;
  - "Slot · Quality · item level N";
  - the source line "{boss}, {zone}, {pct}% drop chance · {day} with {guild}";
  - the gain as a signed figure with a muted unit.
- Fixture: `item-line.json`.

**Crest.**
- Where it lives:
  - web: `ClassCrest.astro`, `character/ClassCrestRing.svelte`, `lib/class-crest.ts`
  - addon: `media/crests` (10)
  - Discord: `assets/crest-*.png` (9)
  - apps: the crest views (`BrandImages.swift:43` adds a glow)
- Spec:
  - the circular ringed HD crest;
  - size and ring weight from tokens;
  - no glow;
  - an LX crest fallback, never a person glyph. Today iOS shows a person glyph in `restyle-app/gear-ios.png`, and Android does in `m0/android-home-34.png`.
- Assets come from the crest export above.

**Raid row.**
- Where it lives:
  - web: `guild/PlannedRaidRow.svelte`, `guild/overview/NextRaidRow.svelte`
  - apps: the raid cards
  - Discord: the `post-*` embeds
  - push: `notify/events.go`
- Spec:
  - the date as `RaidDay` ("Wed 13 Jan"), in sentence case;
  - time with zone and relative time in one cell;
  - filled against needed per role, as "N of M".
- Fixture: `raid-row.json`, mapping one raid to the exact strings for each surface slot.

**Sign-up control.**
- Where it lives:
  - web: `guild/SignupSegments.svelte`, `SignupRow.svelte`
  - iOS: `Raids/Logic/SignupController.swift` and `SignupViews.swift` (four buttons)
  - Android: `raids/ui/SignupViews.kt` (segmented)
  - Discord: a button row
- Spec:
  - a segmented In / Late / Bench / Out control on both apps, as the mock and the web show;
  - state labels from shared copy keys;
  - Discord's native button row is allowed.

**Figures.**
- Where it lives: web `ui/Figure.svelte`, `ui/BigFigures.svelte`, `guild/Figures.svelte`.
- Spec:
  - a tabular figure with a small muted unit;
  - thousands separators per `formats.json` `numbers`.
- Drift: the addon prints "1840", and "+6.4 DPS" as a single green run. Discord uses `%d` counts.

**Buttons, plates, and empty and error states.**
- The companion matches the web.
- The addon's client-red buttons and bevelled frames do not.
- The server app-confirm page is close.
- This belongs to the queued lane E spec, extended to the addon and the companion.

## Lanes

Ranked by what a player notices first. At most three lanes run at a time, with no file overlap inside a wave.

**Precondition, before wave 1: land `lane/app-brand-polish`.**
- It carries the painted app icon (`c669fb8a4`).
- It touches `design/tokens/{build.mjs,tokens.json,tokens.test.mjs}` and `apps/android/app`, so it must merge before any lane below edits those files.

### Wave 1 (parallel)

**BA1: the addon on the tokens.**
- Scope:
  - the Lua token target, with `Theme.HEX` as a map onto it;
  - rarity text tokens;
  - ship Marcellus and Open Sans in `media/fonts`;
  - sentence-case the `Locale.lua` eyebrows;
  - teal primary buttons and hairline frames;
  - thousands separators in figures;
  - the Android `colors_gen.xml` target (same generator, same lane).
- Files (exclusive): `design/tokens/*`, `addon/ForeverSixty/**`, and the new `apps/android/core/designsystem/src/main/res/values/colors_gen.xml`.
- Fixes: addon BLOCKERs 1 and 2 and the addon SHOULDs; Android SHOULD 4 (the XML colours).

**BA2: identity and listing.**
- Scope:
  - the painted seal as the Discord server icon;
  - adopt the 11-09 tavern banner and delete `banner.png`;
  - a crest and seal export script in `design/logo/build.py`, including the tenth Discord crest;
  - listing copy picks up "Play your class better.";
  - App Store and Play listing metadata (text, subtitle, keywords) and a screenshot plan.
- Files (exclusive): `design/logo/*`, `design/community/discord/assets/*`, `design/listing/*.md`, and a new `apps/store/`.
- Fixes: the store BLOCKERs and Discord SHOULD 1. The gallery recapture waits for BA1 (wave 2).

**BA3: one voice from the server.**
- Scope:
  - Go targets in the shared-copy generator;
  - the `mail.*`, `push.*`, `discord.*` and `confirm.*` keys;
  - an HTML mail shell in `api/internal/mail`;
  - sender or Reply-To set to `help@`;
  - a `long_date_year` case in `formats.json`;
  - the Discord post title as `RaidDay`, with a single time cell;
  - the `card.go` ellipsis;
  - `app_confirm.go`'s `#000` and shadow onto tokens;
  - recapture app-confirm.
- Files (exclusive): `web/scripts/gen-mobile-copy.mjs`, `apps/shared/copy/*`, `apps/shared/formats.json`, and `api/internal/{mail,billingmail,notify,discordbot,auth,config,card,site}`.
- Fixes: the email BLOCKERs, Discord SHOULDs 2 and 3, push P1 and P2, and server S1 to S7.

### Wave 2 (after wave 1)

**BA4: companion copy and formats.**
- Depends on: BA3 (the generator) and BA1 (the `tokens.json` additions).
- Scope:
  - tray strings and copy onto the shared-copy companion targets;
  - one name for the reports job;
  - `format.js` and `copy.go` replay `formats.json`;
  - the DMG ground from tokens;
  - the 24 font sizes and 4 radii onto the scale;
  - veil and heading-shadow tokens.
- Files: `companion/**`, `companion/packaging/macos/*`.

**BA5: listing and capture refresh.**
- Depends on: BA1, BA2 and BA3. It needs the owner's game client and the simulator, so schedule it when the laptop is free.
- Scope:
  - a lossless in-game gallery per `RECAPTURE.md`;
  - re-render `2026-11-09-f-discord`;
  - widgets and the Live Activity on iOS and Android;
  - the Android screens never captured: Guild, Rankings, Reports, Planner and Guides.
- Files: `design/listing/gallery`, `design/reviews/*`.

**Queued F (mobile), extended.**
- Depends on: BA1 (the Android XML colours) and app-brand-polish.
- Scope:
  - remove the retired styles from every Android screen, plus a surface guard;
  - delete `GearScreen.kt`;
  - iOS and Android drop forced caps and glows;
  - generated icons replace SF Symbols and Material (through the BA1 and BA2 icon and crest outputs);
  - both apps replay `formats.json`;
  - the Android lock-screen card goes dark with a gold timer;
  - widget fonts.
- Files: `apps/ios/**`, `apps/android/**`.

### Hand-offs to the queued lanes

**D (crests, names, sign-up)** takes the crest spec above for every surface:
- ringed, with no glow, and the LX fallback in place of the person glyph;
- Discord character and item thumbnails ringed (`item_embed.go:60`, via `sheet/palette.go` `qualityRing`);
- one segmented sign-up control on both apps.

**E (controls and states):** extend the spec to the addon's buttons and frames, and to the companion's lead-panel filigree (an owner call).

**F (mobile):** absorbs the app findings above. It must not start Android work until app-brand-polish has merged.

### Conflicts to watch

- BA1 and app-brand-polish both edit `design/tokens/*`, which is why app-brand-polish lands first.
- BA3 and F both touch `apps/shared/copy`. F only consumes the generated copy, so F rebases after BA3.
- BA4 consumes the shared-copy generator, so it cannot start until BA3 lands.

## Stale or missing captures

- **iOS:** these predate the redesign or the restyle:
  - `design/reviews/2026-10-31-mobile-wave3-widgets/build/ios-*.png`
  - `m5i/ios-live-raid-night.png`
  - the `m1i` side-by-sides
- **Android:**
  - `m0/android-home-34.png` is broken.
  - `2026-10-31-mobile-wave3-widgets/build/and-*.png` predate the redesign.
  - Guild, Rankings, Reports, Planner and Guides have never been captured.
- **Discord:** `2026-11-09-f-discord/discord/*` predates `5b8e8a02c`, the house format.
- **Companion:** `2026-11-09-companion/round-4/*` predates `80e3da9f4`. It shows "Today 18:40", which the code no longer prints.
- **Server:** `2026-11-10-h-api-html/after/app-confirm-*` predates `app_confirm.go:48-51`.
- **Addon:** `design/listing/gallery/*` (11-03) and `2026-11-08-addon-ux/*` predate the umber panel and the tour.
- **Push:** not rendered anywhere.
