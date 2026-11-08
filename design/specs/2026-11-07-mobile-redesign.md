# Native apps on the flat system: design spec

Author: MOBILE-DESIGN lane, 2026-11-07, ship mode. The owner said "do it": restyle the iOS (SwiftUI) and Android (Compose) apps to the system the website now uses, the way the website was done: mock first, then build to match exactly, under a parity gate. Owner addition the same day: "can you make it where both apps design can be changed from a single place?" (section 2). This spec is the design half; the iOS and Android build lanes run after it, in parallel (section 8).

System it carries: `design/specs/2026-11-05-web-layout.md` (type §3.3, primitives §3A, the shared paperdoll §4), `2026-11-07-home-b.md`, `2026-11-06-character.md` (§6, Activity), the page specs listed in `design/reviews/2026-11-07-redesign-rollout.md`, and `design/parity.md`. Memory rules that bind: one crest language, faithful HD icons, the best-addon standard (every item has its icon, quality colour and a real tooltip), dense and clean, side by side with the mock, build mode ships fast. No em dashes in any copy.

## 0. Artifacts

| What | Path |
|---|---|
| This spec | `design/specs/2026-11-07-mobile-redesign.md` |
| App frame (status bar, navigation bar, tab bar, the native type deltas) | `design/mocks/ws_app_frame.py` |
| Screen catalogue and app-only screens | `design/mocks/ws_app.py` (`python3 design/mocks/ws_app.py [filter]`) |
| Captures | `design/mocks/ws_app_shoot.mjs` (`node design/mocks/ws_app_shoot.mjs [filter]`; `PLAYWRIGHT_FROM=<package.json>` when `web/node_modules` is absent) |
| Mock HTML | `design/mocks/app/<platform>-<screen>.html`: the app-only screens are committed; the web-derived ones are generated (gitignored, 51 MB of inlined icons) and rebuilt by the command above |
| Captures (iOS 393 x 852, Android 412 x 915, the whole scroll at 1x) | `design/reviews/2026-11-07-mobile/<platform>-<screen>.png` |
| Native parity reference | `design/mocks/app_regions.mjs` → `design/mocks/app/regions/<platform>-<screen>.json` (region boxes and inventories); fonts from the repo (`design/mocks/app_fonts.mjs`) |
| Fonts | `apps/shared/fonts/Marcellus-Regular.ttf` (the unmodified upstream file, so the Reserved Font Name clause of the OFL is kept) and `OFL-Marcellus.txt`; Barlow 400/600/700 and `OFL-Barlow.txt` as before |

How a screen that matches a web page is mocked: its web phone mock (`design/mocks/web/<mock>.html`, the 390 px layout the parity gate already holds the site to) inside the app frame, with the site header and footer hidden. The phone deltas the rulings add (home order, the level 34 home, the feed filters) are CSS over the same markup in `ws_app.py`, so the web mock stays the one source of structure and data and an app mock moves when its web mock moves. App-only screens (switcher, devices, deletion, offline data, update required, raid night, widgets, Live Activity, search, You, sign in again) are composed from `ws_prims` primitives alone; `assert_no_surface_styles` holds for them.

Data: the home and guild fixtures (`web/tests/e2e/fixtures/home/*.json`): Thandrel (Fury Warrior 60, Officer of Olympus XXVII), Felsnap (Beast Mastery Hunter 34, no guild); the character, BiS landing, guide, sim, report and guild fixtures their web mocks already read. Clock: Fri 6 Nov 2026, 18:00 UTC (`ws_home.NOW`).

## 0A. Rulings

Persona round: wow-player, 2026-11-07, on the questions in section 0B. Where the lane overruled the persona it says so and why. The owner delegated every call.

| # | Ruling |
|---|---|
| R0 | (Owner, 2026-11-08) **The app is the away-from-PC raid and notifications companion.** The primary flow, in order: (1) raid sign-up and raid night, one tap from a push or the Live Activity; (2) "what changed for me", the feed and push (a loot upgrade, a tier move, a sim done, a guildmate passing you); (3) the next upgrade and where it drops; (4) reference opened from links (BiS, tiers, shared builds). Every existing screen stays and is restyled (planner, sims, reports, rankings); none is deleted ("we shouldn't throw the work away"). They are secondary: reached from Next rows, You and deep links, never from a tab or the top of Home. R1's five tabs stand. Home, signed in, in a guild: the next raid and your sign-up, then the next upgrade, then the feed; without a guild: the next upgrade, then the feed (R2 as amended). Push is first-class and not live yet (no `aps-environment` entitlement): the push lane is the first item of section 8 on both platforms. |
| R1 | **Five tabs: Home, BiS, Guides, Guild, You.** Lane over persona (the persona wanted four, with Guild as a row on You): the lane plan names guild and raids "likely the flagship mobile use", sign-up from the app replaces Discord, and the raid push notifications land there; a tab one tap from the lock-screen deep link beats a row two taps in. With no guild the tab shows the guild landing (find or claim a guild). Guides stays behind its feature switch (four tabs while it is off). Sims, Talents and the Planner, Rankings and Tiers are reached from Home's Next rows and from the screens that name them; Logs and Reports from You (a Logs row) and from Home's Next rows; Tiers and What changed are segments inside Guides. Search stays a toolbar button on Home, BiS and Guides. |
| R2 | **Home, signed in, default order (server-driven, section 2.3), amended by R0 on 2026-11-08:** in a guild, head (three figures), countdowns (the next raid and your sign-up), Your upgrades, Forever right now, Guild sims, Your spec this week, Top movers, Next rows; without a guild, head, Your upgrades, Forever right now, countdowns (first raids), the Join a guild row, Your spec this week, Top movers, Next rows. The phone has no rail, so the player's own next move comes first. (Before R0 the countdowns sat after the upgrades and moved up only on raid day.) |
| R3 | **Home, signed out:** title `Play your class better.`, `Sign in with Battle.net` (the one primary), Pick your class, the raid countdown, the feed (All, no filters). How to start becomes one footer link. A crest goes to that class's BiS at 60 raid, the guide one tap away. |
| R4 | **Level 34:** the head's third figure is levels to the next band; Guild sims is one line, `Join a guild to see how you sim against it. Guilds ›`; the first-raids countdown is one line under the feed. |
| R5 | **Feed filters:** three segments full width under the feed title (All, Your spec, Your guild), Logs behind a filter button at the end (four segments do not fit 360 dp). Your guild hides with no guild. A row opens its object (the sim, the report, the item, the character). Filters do not swipe. |
| R6 | **BiS:** the paperdoll is the default view of a spec's list, one column of 17 cells, a tap expands the inspector in place (pick, two runners-up, source, drop chance); Where to go is the second view (a segmented toggle, swipeable). Signed out the BiS tab opens on 60 raid (as the app does today; app users know their level, the web's 20 to 29 serves search visitors); signed in, the character's band. The band is always visible in the control bar summary. |
| R7 | **Character:** head answer (+DPS available), the standing line, Gear (the paperdoll, worn mode), Best per encounter, Forever Score, Builds seen, Activity, the owner's History. Persona over the web order: Forever Score drops below the fights on a phone. |
| R8 | **Guides:** the tab's top row is a segmented `Guide`, `Tiers`, `What changed` (swipeable pages). A spec guide opens on At a glance under the head, then the build, rotation and gear. |
| R9 | **Sim result:** the DPS answer first with its error and, when a previous sim of the same character exists, `vs last sim`; then what to do about it (the upgrade picks), the ability table, the stat weights in the SimC convention (top stat 1.00). Share is a toolbar button that sends a link. |
| R10 | **Guild tab:** a member sees the next raid first with its countdown and their own In, Late, Bench, Out buttons; the guild rank is a text label on each roster row; the guild sims ladder lives here in full (home shows the top 5 and links here). |
| R11 | **No swipe between tabs** (it fights horizontal scrolls and the iOS back gesture). Swipe between segments only where the segments are pages: Guides (Guide, Tiers, What changed), BiS (By slot, Where to go), report tabs. |
| R12 | **Pull to refresh:** Home, BiS, Character, Guild and raid pages, Reports and Logs, Rankings, Sims list, You. Not on a sim in progress, not on Guides (static data refreshes on launch). **Share:** builds, sims, guides, reports, raids, a character; the text is one line plus the link: `Survival Hunter build, level 60` / `Sim: 842 DPS ±6, Survival Hunter` / `Onyxia's Lair, Wed 9 Dec, 8:00 pm. Sign up:`. |
| R13 | **Widgets:** small: the next guild raid (name, time, your sign-up) or, with no guild or before raids open, `First raids` with the days and `Wed 9 Dec`; second small: `+48.1 DPS`, `from 3 upgrades`. Medium: the countdown left, the upgrades right. Lock screen: the countdown; the upgrades. Never empty: with no character the upgrade side reads `Pick a character`. **Live Activity** on raid night: raid icon and name, `You are In`, the clock, `18 in, 3 late`; Dynamic Island compact `Ony 42:00`, expanded adds the four sign-up buttons; a sim shows its progress and its DPS when done. Android: the same as an ongoing notification. |
| R14 | **Haptics (light):** a sign-up change, a segmented change, sign-in done, a sim finished (success), the pull-to-refresh trigger point. Never on scroll, tab taps, feed arrival or a repeating error. |
| R15 | **Low-level filter and the stamp:** the switcher sheet lists level 10 and up, `Show low-level characters (N)` at the bottom (the current character and the main always show); `Updated 2 h ago from Battle.net · Refresh` sits at the top of the switcher and under the character head on You. |
| R16 | **System screens:** deletion confirms with `This removes your characters, sims and guild sign-ups. It cannot be undone.`, names what goes, and for Premium says to cancel the store subscription first; Devices adds `Sign out everywhere else` (needs an API endpoint, section 6); offline shows `Offline. Showing what you had at 4:12 pm.` and disables sign-up with its reason; Update required is never a dead screen (store button, `Your characters and sign-ups are safe.`); an expired token gets `Your Battle.net link expired. Sign in again to bring your characters back.` |
| O1 | (Owner) One token source for web, iOS and Android (section 2.1). |
| O2 | (Owner) All UI copy in `apps/shared`; where the web already has the copy, one module both read (section 2.2). |
| O3 | (Owner) The home sections and their order come from `/v1/home` (section 2.3). |
| N1 | (Lane, from the brief) **Labels are sentence case on native**: 12/16 Barlow 700 `muted`, tracking .01em. The web still sets its 12 px label in caps (`Figure`, `Inspector`, `ColumnHeads`, `ControlBar`); the mocks apply the native case through `ws_app_frame.APP_CSS`. The parity gate compares structure, not case, so the two do not conflict; moving the web to sentence case is a one-token change after section 2.1 lands, proposed, not ruled here. |
| N2 | (Lane) The gutter is 16 pt / dp; the section gap is 32 (the web's 48 is for a page with a rail). |
| N3 | (Lane) The page title is Marcellus once per screen in the PageHeader, never in the navigation bar. The navigation bar has no title at rest; it fades in the title in Barlow 600 when the PageHeader scrolls under it (iOS inline title, Android `TopAppBar` on scroll). |
| N4 | (Lane) Nothing under 12, tab labels included (12/14 600). |

### 0B. Questions put to the persona
Tabs and where Sims, Logs, Rankings, Tiers and the Planner live; the home order on a phone; signed-out home; 34 against 60; feed filters; BiS default view and band; character order; guides segments; sim result order and share; the guild tab; swipe; pull to refresh and share text; widgets and the Live Activity; haptics; the low-level filter and the Battle.net stamp; deletion, devices, offline and update required.

## 1. What is old today (audit 2026-11-07)

Both apps were built 30 to 31 October on the previous look. Measured on origin/main:
- **Type:** titles, section heads, stat figures and card heads in Barlow Condensed ExtraBold caps (`ScaledType.title/h2/stat/sec/cardHead`, Android `FsType`); `micro` 11 pt caps labels (`MicroLabel`); Cinzel for the wordmark. Marcellus is not bundled. JetBrains Mono is referenced but not bundled (falls back to the system mono).
- **Surfaces:** `primaryCard()` (gold-edged gradient card, 12 radius), `surfaceBand()` (ember radial glow), `quietCard()`; the primary button is a gold gradient with caps text; `TagView` caps; raid message titles `uppercased()`. Every list sits in a card; the flat system has one Plate per screen.
- **Tokens:** generated from `web/src/styles/tokens.css` by `web/scripts/gen-mobile-tokens.mjs` (colours, a type scale, radii, a 4 pt spacing scale), so the colours are right and the roles are the old ones.
- **Features missing against the redesigned web** (both platforms unless noted): the home B feed (`/v1/feed` is generated, no screen reads it), `/v1/home` (import sim, guild sims, spec week, movers), the first-raids countdown (the app counts down only to a guild raid), the guild sims ladder, the guild rank as a standing line, the character Activity section and standing line, the BiS paperdoll (BiS is a list; the doll lives on Gear and the character page), Where to go, the BiS landing answer, the guide's At a glance as a panel in the new system (both have it as a card), the sim result redesign (answer, ability table, weights, Your sims), the report restyle (summary, player, coach), the level-10 character filter, `Updated 2 h ago from Battle.net` (both say `Synced ... from Battle.net` in the home hero), the addon wording in two places (`Get the addon` on Home signed out; must be `Get the addon and companion`, `cta-copy.ts`), pull to refresh on Android Raids and Guides, light haptics on iOS beyond sign-up and the planner.
- **Already right and kept:** account deletion (`DeleteAccountView` / `DeleteAccountScreen`, the account-delete lane; restyled only), Live Activities and widgets (restyled), the update gate, offline data, devices, the switcher sheet, share on the main objects, Dynamic Type through `relativeTo:` on iOS and sp on Android, 44 / 48 targets in most views, snapshot suites on both (394 iOS goldens, 378 Android).

## 2. One source for the design, the copy and the home

### 2.1 Tokens (O1)
`design/tokens/tokens.json` becomes the only hand-edited token file, in the W3C design-tokens format (`$value`, `$type`, `$description`; the format Style Dictionary 4 reads). A small in-repo generator, `design/tokens/build.mjs` (no dependency; Style Dictionary is not needed for three outputs), replaces `web/scripts/gen-mobile-tokens.mjs` and writes:

| Output | Holds |
|---|---|
| `web/src/styles/tokens.css` | the `@theme static` block and the 1760 px block exactly as today (each token's `$description` becomes its comment). First run: the generated file defines every custom property today's file defines with the same value; the parity specs prove no visual change (no mock moves) |
| `apps/ios/Shared/DesignTokens.swift` | `FSColor` (`Color` statics), `FSFont` (PostScript names), `FSType` (each role: face, size, line height, tracking, the Dynamic Type style it scales with), `FSSpace`, `FSRadius`, `FSStroke`, `FSElevation`, `FSMotion`, `FSTarget` |
| `apps/android/core/designsystem/.../Tokens.kt` | `FsColor`, `FsFont` (`FontFamily` from `res/font`), `FsType` (`TextStyle` per role, sp), `FsSpace`, `FsShape` (`Shapes`), `FsStroke`, `FsElevation`, `FsMotion`, `FsTarget` |

Groups in `tokens.json` (every value the three surfaces use; a platform reads only its keys):
- `color`: base (`bg #07090d`, `raised #0d111a`, `hover #131824`, `pressed #1a2030`, `selected` gold 10%, `line #262e40`, `lineSoft #1c2230`, `text #e9e4d8`, `strong #f2eee4`, `muted #9a9484`, `nav #b9b3a4`, `gold #e5b955`, `goldHover #f5d27a`), outcome (`kill`, `wipe`, `death`), parse (exactly Warcraft Logs', with `parseGreyText`), school, faction (`alliance`, `horde`, bars), class (nine, plus monk kept for data), rarity (game colours for borders) and rarity text (`rareText #3d94f0`, `epicText #b866f5`, the AA-safe text colours), tier (`S #ff8000`, `A #b866f5`, `B #3d94f0`, `C #1eff00`, `D #9d9d9d`).
- `font`: `title` Marcellus 400; `body` Barlow 400, 600, 700; `code` JetBrains Mono 500 (literal codes only: export strings, build strings, slash commands). Barlow Condensed and Cinzel leave the apps (the wordmark is the SVG seal and word image, as on the site).
- `type` (role: size / line, weight, tracking; `phone` variants where the web has them; `scalesWith` for iOS): `title` 40/46 (phone 32/38) Marcellus, `.title`; `sectionReading` 22/28 (phone 20/26) 700; `section` 18/24 700 `.headline`; `itemLg` 18/22 700; `item` 15/20 700 `.body`; `lead` 17/26 400 `.body`; `prose` 16/26 400 `.body`; `body` 15/22 400 and 600 `.body`; `small` 13/20 400 `.footnote`; `label` 12/16 700 tracking .01em native (.12em caps web) `.caption`; `button` 13/16 700 `.subheadline`; `figureSm` 16/20, `figure` 28/32, `figureBig` 36/40, `figureXl` 44/48 (56/60 at 768 up), all 700 tabular, `.title2`; `tab` 12/14 600 (native).
- `space`: the 4 pt steps (`1` to `16`), `gutterPhone 16`, `gutter 48`, `gutterWide 64`, `sectionGap 48`, `sectionGapCompact 32` (native), `gap 24`, `rowPadX 16`, `rowPadY 10`.
- `size`: `rowH 58` (paperdoll cell), `row 56`, `rowSlim 44`, `barH 56`, `barPhone 48`, `segment 40` (44 phone), `icon` 28, 32, 36, 40, 56, `crest` 36, 40, 44, 56, 64, 84, `targetIos 44`, `targetAndroid 48`.
- `radius`: `plate 6`, `field 4`, `pill 3`, `control 4`, `widget 22`, `sheet 12`.
- `stroke`: `hairline 1`, `focus 2`, `selection 3`, `crestRing 2`.
- `elevation`: `overlay` (0 12 32 black 60%); native sheets use the platform's own scrim and no extra shadow.
- `motion`: `fast 120ms`, `base 200ms`, `slow 320ms`, `tooltipDelay 150ms`, `livePulse 1600ms`, `ease standard`; every animation is skipped under Reduce Motion / Remove animations.

Check: `design/tokens/tokens.test.mjs` (run by the web unit job and `mobile-codegen.yml`) regenerates the three outputs in memory and fails on any difference, naming the file and the command; each generated file opens with `GENERATED by design/tokens/build.mjs from design/tokens/tokens.json. Do not edit.`, and the test also fails if that line is missing (a hand edit that removed it). `mobile-tokens.test.ts` and `gen-mobile-tokens.mjs` are deleted.

### 2.2 Copy (O2)
- `apps/shared/copy/en.json` stays the app catalog (about 1,766 keys, generated into `Localizable.xcstrings` + `Copy.swift` and `strings.xml` by `web/scripts/gen-mobile-copy.mjs`, drift-checked by `mobile-codegen.yml`).
- Copy the web already owns and the apps show moves into one JSON both read: `apps/shared/copy/shared.json` (same value grammar as `en.json`, keys `shared.*`). The web modules keep their names and exports and read the JSON (`import copy from '../../../../apps/shared/copy/shared.json'`, typed by a generated `web/src/lib/copy/shared.d.ts`), so no web call site changes. First moves (the ones this redesign shows on both): `lib/addon/cta-copy.ts` (`getAddonAndCompanion`, `installAddonAndCompanion`, the inline form), `lib/account/level-filter.ts` (`lowLevelCopy`), `lib/account/character-list-copy.ts` (`updatedFrom`), the home B words (`lib/home/*` head labels, countdown eyebrows, feed tags and filter names, the floor and empty lines, Guild sims, spec week, movers), the character standing line and Activity title, `lib/character-selector/copy.ts` (`refreshing`), the BiS landing answers (`landingCopy`), the paperdoll cell words (`over next`, `Worn`, `Best in slot`, `Not scored`, `Nothing better yet`, `Two-hand in main hand`), and `premiumBlockCopy.supporterChip`. Account deletion is unified here too (account-delete lane, 2026-11-07): the app said `That doesn't match. Type your BattleTag or delete.` and the web `Type {BattleTag} to confirm.`; both now read the one key `shared.account.delete.mismatch` = `Type {battletag} to confirm.` (the web's words, the BattleTag filled in), and the delete screens' other words (title, what goes, the Premium store warning, failed, done) move with it.
- The feed sentences are not copy: the API writes them as segments (home B §2), so both apps render the segments and own no feed words.
- Rule for the lanes: a string a native screen shows is a key in `en.json` or `shared.json`, never a literal in Swift or Kotlin (the existing lint holds); a word that also appears on the site is in `shared.json`.

### 2.3 Server-driven home (O3)
`GET /v1/home` gains `sections`, the home's blocks in order, computed for the viewer:

```json
{ "sections": [
    { "type": "head" },
    { "type": "upgrades", "limit": 3 },
    { "type": "countdowns", "style": "full" },
    { "type": "feed", "filters": ["all", "spec", "guild"], "more_filters": ["logs"], "limit": 7 },
    { "type": "guild_sims", "limit": 5 },
    { "type": "spec_week" },
    { "type": "movers" },
    { "type": "next" } ],
  "import_sim": {}, "spec_week": {}, "movers": {}, "guild_sims": {} }
```
- Types: `head`, `upgrades`, `countdowns` (`style`: `full` or `line`), `feed`, `guild_sims`, `guild_sims_cta` (the one-line `Join a guild` row), `spec_week`, `movers`, `next`, `notice` (`text`, `href`: a standing message the owner can switch on). An app renders the types it knows in the given order and skips any other (logged once per type per session through the telemetry it already sends). An unknown `style` falls back to the type's default.
- The API composes the default from R0, R2 and R4: in a guild, `countdowns` is second (the next raid and your sign-up), before `upgrades` and `feed`; with no guild, `countdowns` follows the feed and `guild_sims_cta` replaces `guild_sims`; under 60 the countdown is `line`. (Lane T shipped the pre-R0 order, countdowns second only on the guild's raid day; the R0 order is a follow-up of lane T.) Each rule is a server switch (`HOME_SECTIONS` env, a JSON override, like `APP_FEATURES`), so the owner reorders or switches off a block without a release.
- Signed out (no `/v1/home`): the apps use the bundled default `apps/shared/home-sections.json` (`head_out`, `class_picker`, `countdowns`, `feed`, `how_to_start_link`), which the API also serves at `GET /v1/app/config` `home_sections_out` so it is switchable too.
- The web keeps its own composition (its rail layout); reading `sections` there is a later choice, not part of this work.

## 3. The native system

### 3.1 Colour and surfaces
Dark only, as the site. Three surfaces: S0 `bg` (the screen), S1 `raised` (the one Plate, sheets, widgets), S2 `hover` (pressed rows on touch: `pressed`). Hierarchy comes from type, spacing and hairlines (`lineSoft` between rows, `line` around the Plate and controls). No gradients, no glows, no gold-edged cards. Status bar light content; Android edge to edge with a transparent status bar and the navigation bar drawn by the tab bar.

### 3.2 Type and Dynamic Type
Roles from `tokens.json` (`FSType` / `FsType`). iOS: `Font.custom(postScript, size:, relativeTo:)` with each role's `scalesWith`, line height through `lineSpacing` from the role's line minus the face's natural line; tabular figures with `.monospacedDigit()`. Android: `TextStyle` in sp with `lineHeight`, `LineHeightStyle(Center, Trim.None)`, `fontFeatureSettings = "tnum"` on figures. Both scale with the system size, with these rules:
- Up to the iOS `xxxLarge` / Android 1.3 the layouts are the mocks'.
- From `accessibility1` / 1.5: a Row's trail moves under its body (the web's below-600 rule), Figure rows wrap to two per line, the paperdoll cell's gain moves under the name, the segmented control becomes a vertical list, the tab bar keeps icons and drops labels to the large-content viewer (iOS) / the tooltip (Android).
- Figures cap at 1.5x (`.dynamicTypeSize(...DynamicTypeSize.accessibility1)` on the figure, `fontScale.coerceAtMost(1.5f)` on Android) so a number never truncates; text never caps.
- Nothing truncates that carries meaning: an item name wraps to two lines and only a boss name in a source line ellipsizes (web §4.3).

### 3.3 Targets, safe areas, gestures
44 pt (iOS) / 48 dp (Android) for every control and row; a 13 px link keeps its look with a padded hit area (web `Section` action rule). Content respects the safe areas; the tab bar sits on the home indicator inset; nothing scrolls under the Dynamic Island. Back swipe on iOS everywhere a screen is pushed; Android predictive back. Pull to refresh per R12 (the system control, tinted gold). Swipe per R11 (`TabView(.page)` / `HorizontalPager` only for segment pages). Haptics per R14 (`UIImpactFeedbackGenerator(.light)`, `.success` notification for a finished sim; Android `HapticFeedbackConstants.CONFIRM` / `CLOCK_TICK`).

### 3.4 Primitives: web to SwiftUI and Compose
Names are the web's, so a grep finds all three. iOS in `apps/ios/Sources/DesignSystem/`, Android in `apps/android/core/designsystem/`. Every primitive reads tokens only, carries its parity region (section 5) and its accessibility role.

| Primitive | Web | SwiftUI | Compose | Native notes |
|---|---|---|---|---|
| Label | `text-label` (caps 12) | `Label12(_:)` (`FSType.label`, `muted`) | `Label12(text)` | Sentence case (N1). Never a button. |
| Tag | `ui/Tag` | `TagView(_:tone:)` restyled: 20 high, 3 radius, 1 pt border in tone, 12/16 600, sentence case | `Tag(text, tone)` | Tones gold, kill, wipe, muted. Not tappable. |
| Button | `ui/Button` | `FSButton(_:kind:)` kinds `primary` (gold fill, `#17110a` text), `secondary` (1 pt `line`), `danger` (wipe text and hairline), `ghost` (gold text) | `FsButton(text, kind)` | 44 / 48 high, 13/16 700 sentence case, 4 radius; one primary per screen; pressed `pressed` ground; disabled 50% text, no hover; loading keeps width with a spinner. |
| ActionLink | `ui/ActionLink` | `ActionLink(_:)` gold 600 with `›` | `ActionLink` | Padded to the target size. |
| PageHeader | `ui/PageHeader` | `PageHeader(title:titleColor:crest:meta:lead:answer:actions:)` | same | Marcellus title 32/38 in the class, spec or faction colour (`strong` without one), once per screen; meta one wrapping line of facts and Tags; answer below on phones (a Figure and its links); actions below the answer. |
| ControlBar | `ui/ControlBar` | `ControlBar(summary:groups:)` | same | Phone form only: one 48 row with the selection in words and `Change ▾`; tapping expands the groups in place (labelled Segmented and Selects, full width, 44 / 48), collapses on a choice; never a sheet (parity with the web phone bar). Sticky under the navigation bar. |
| Segmented | `ui/Segmented` | `Segmented(items:selection:)` | same | 1 pt `line` outline, 4 radius, hairline dividers, 14/16 600; selected gold text, `selected` ground, 2 pt gold underline; light haptic on change; a radio group to VoiceOver / TalkBack. |
| Select | `ui/Select` | `FSSelect` | `FsSelect` | Opens the platform menu (`Menu` / `DropdownMenu`) styled as the overlay. |
| Section | `ui/Section` | `Section(title:meta:action:)` | same | Title 18/24 700 `strong` over a `lineSoft` hairline, meta `small` muted, action an ActionLink at the right; 32 above (N2). Never inside a Plate. |
| Plate | `ui/Plate` | `Plate { }` | `Plate { }` | `raised` ground, 1 pt `line`, 6 radius, no shadow. One per screen (the paperdoll, the tier list, the raid nights, the result). |
| Row | `ui/Row` | `FSRow(title:sub:lead:trail:titleColor:href or action:selected:density:)` | `FsRow` | Lead, body, trail; 56 min (44 slim); hairline under; a link row or an action row, never both; selected: `selected` ground and a 3 pt gold bar; pressed `pressed`; chevron only when it navigates on iOS (system convention), none on Android. |
| ItemRow | `ui/ItemRow` | `ItemRow(item:source:chance:gain:worn:badge:slot:)` | same | Icon in a 1 pt quality border, 5 radius; name 15/20 700 in the AA-safe quality colour; source one line with its kind glyph; long-press (and a tap on the icon) opens the item tooltip sheet (best-addon tenet). |
| Figure | `ui/Figure` | `FigureView(value:unit:label:caption:tone:size:)` | `Figure` | Barlow 700 tabular; unit 13 600 muted; label above (N1), caption below; tones strong, gold, kill, wipe. Caps at 1.5x. |
| Meter | `ui/Meter` | `MeterBar(value:tone:label:)` | `Meter` | 8 high, 3 radius, `lineSoft` track; accessibility value in percent. |
| Countdown | `ui/Countdown` | `CountdownPanel(eyebrow:target:aside:tone:live:)` with `TimelineView(.periodic(by: 60))` | `Countdown` with a minute ticker | Days, hours, minutes; seconds only in the final hour; one timer to VoiceOver ("32 days, 6 hours"); the live dot pulses `livePulse`, still under Reduce Motion; `style: line` is one row (`First raids · Wed 9 Dec`). The raid tone keeps its quiet purple wash, the guild tone a gold wash: the two allowed grounds, as on the web. |
| FeedRow | `ui/FeedRow` | `FeedRow(lead:segments:at:tag:)` | same | 40 crest or icon, the API's segments (name in class colour, item in quality colour with its tooltip, strong), the age (`4 min ago`, refreshed each minute) over the tag (12, sentence case); the whole row opens the object. `FeedSkeleton` 5 rows at final height. |
| GradeTile | `ui/GradeTile` | `GradeTile(letter:)` | same | 56 square, 8 radius, tier colour ground, 32 700 `#17110a`. |
| Tally | `ui/Tally` | `TallyList(items:)` | same | Counts against a target with a Meter each (roles on a raid, readiness). |
| Overlay / sheet | `ui/Overlay`, `UiSheet` | `.sheet` with `.presentationDetents([.medium, .large])`, `raised` ground, the system grabber | `ModalBottomSheet` with `raised` container, drag handle | Item tooltip, switcher, filters, meters, sign-up note. Tooltips on a phone are sheets (long-press or tap the icon). |
| Paperdoll | `ui/Paperdoll` | `PaperdollView(mode:slots:selected:onSelect:view:)` | `Paperdoll` | One column of 17 cells on phones (head to feet, then weapons, as web §4.3), cell 58 min: 40 icon, name, source, gain (Figure sm) at the right; modes pick, compare, worn, report; tap selects and expands the Inspector in place under the cell (3 pt gold bar joins them), a second tap collapses; By slot / Where to go is a Segmented in the summary row; loading 17 skeleton cells at final height; error one LoadError row. |
| Inspector | `ui/Inspector` | `InspectorInline` | same | In place only on phones: stats, You wear, Where to get it, Runners-up (pick pinned, deltas in Figure), one link. |
| CharacterAnswer | `ui/CharacterAnswer` | `CharacterAnswer(label:state:...)` | same | The head's "your character" block: crest, name in class colour, line, links; loading, message and error states. |
| Crest | `ClassCrestRing` | `ClassCrest(class:size:)` | same | Circle clip, 2 pt class ring at 55% (100% pressed), `raised` backing; the only picture of a character or class (one crest language). |
| ItemIcon / SpellIcon | `ui/ItemIcon` | `ItemIcon(icon:quality:size:)` | same | The HD game icon from the site origin, quality border, cached by URL. |
| Skeleton, EmptyState, LoadError, Notice | `ui/*` | same names | same names | Skeletons at final size (no shift); EmptyState one sentence and at most one action; LoadError one row `We could not load ... ` and `Try again`; Notice a standing sentence (the offline strip is a Notice, neutral tone). |
| Switch | `ui/Switch` | system `Toggle` tinted gold | `Switch` with gold track | |
| TabStrip | `ui/TabStrip` | `TabStrip` (scrolling, 2 pt gold underline) | same | Route-level tabs inside a screen (report tabs, guild tabs). |

Retired on both: `primaryCard`, `surfaceBand`, `quietCard`, `MicroLabel`, the gradient primary button, every `uppercased()` on a title, `ScaledType.title/h2/stat/sec/cardHead` in Barlow Condensed, `Strips.kt`'s caps strips. A surface guard like the web's (`surface-guard.test.ts`) lands on both: outside `DesignSystem` / `designsystem`, a `background(` with a token colour, `.overlay(RoundedRectangle...stroke)`, `Modifier.border(` or `Card(` fails a unit test (`SurfaceGuardTests.swift`, `SurfaceGuardTest.kt`), with a shrinking allowlist like the web had.

## 4. The frame and every screen

### 4.1 Devices and frame
iPhone 393 x 852 pt (iPhone 16 / 17 Pro class), Pixel 412 x 915 dp (Pixel 9 class). Status bar 54 / 32, navigation bar 44 / 64, tab bar 49 + 34 home indicator / 80 + 16 gesture area. iOS: a standard `TabView`, gold tint, icons from SF Symbols drawn at the tab glyph's weight (`house`, `shield.checkered`, `book`, `flag`, `person`); Android: Material 3 `NavigationBar` on `raised`, the active indicator `gold` at 18%, icons gold when active. Labels 12/14 600 (N4).

### 4.2 Screen inventory, web counterpart, what changes, data
Each row is one screen (iOS view, Android composable). "Old" lists the old-style pieces it carries today; "Missing" what the redesigned web has that it lacks; "Data" the reads. Mock ids are the capture names in `design/reviews/2026-11-07-mobile/` (`ios-<id>.png`, `android-<id>.png`).

| # | Screen (iOS / Android) | Web counterpart | Old | Missing | Data | Mock ids |
|---|---|---|---|---|---|---|
| 1 | Root, tab bar (`RootView` / `ForeverTabScaffold`, `ForeverTab`) | site header doors | a Raids tab, system tab labels at 10 pt, the Android bar on `surface` | Guild tab (R1), 12 pt labels, the switcher entry on Home's head | `/v1/app/config`, `/v1/me` | every screen |
| 2 | Home (`HomeTab`, `HomeCharacterScreen`, `HomeAnswerView` / `HomeRoute`, `HomeAnswer`) | `/` home B | hero band (`surfaceBand`), Upgrades card, caps heads | sections array (O3), import-sim figure, countdowns (first raids and guild), feed with filters, guild sims, spec week, movers, Next rows restyled | `/v1/home`, `/v1/feed`, `/v1/me`, `/v1/guilds/{id}/home`, `/data/dates.json` | home-60, home-60-refresh, home-34, home-new, home-feed-empty, home-feed-loading, home-loading, home-error |
| 3 | Home signed out (`HomeAccountGate`, `SignedOutView`, `ImportView`, `NoCharactersView`) | `/` signed out | seal hero, `Get the addon` | class picker, raid countdown, feed (All), `Get the addon and companion`, How to start link | `/v1/feed`, dates | home-out |
| 4 | Sign in again (new) | none | | R16 expired token | session | expired |
| 5 | Switcher sheet (`CharacterSwitcher` both) | header selector | card rows | level-10 fold, `Updated 2 h ago from Battle.net · Refresh`, crest rows, Main tag | `/v1/me`, `POST /v1/me/refresh` | switcher, switcher-low |
| 6 | Search (`SearchScreen` / `SearchScreen.kt`) | header search | grouped cards | flat rows, crest and quality colours | `/v1/characters/search`, `/data/{build}/search.json` | search, search-results |
| 7 | Item page and item tooltip sheet (`ItemPageView`, `ItemTooltipSheet` / `ItemPage`, `ItemTooltipSheet`) | the item tooltip | pick card | ItemRow and Inspector sections (Where to get it, Runners-up) | static items, item-uses, loot | (the Inspector in bis-spec-*) |
| 8 | Character (`CharacterView`, `CharacterScreen` / `CharacterPage`) | `/character/...` | Score card, caps heads, paperdoll in a card | head answer +DPS, standing line, Paperdoll worn + Inspector, Best per encounter rows, Forever Score section, Builds seen, Activity, owner History, not-found with Same name elsewhere | `/v1/characters/{...}`, `/rating`, `/history`, `/v1/home`, `/v1/feed?character=` | character-owner-60, character-owner-34, character-visitor-60, character-no-score, character-not-found, character-loading, character-error |
| 9 | Gear (`GearView`, `GearScreen` / `GearList`) | the character Gear section | separate screen | folds into Character (worn paperdoll); the screen is deleted | | character-* |
| 10 | BiS (`BisView`, `BisScreen`, `BisPickerSheet` / `BisContent`, `BisPicker`) | `/bis` landing and `/bis/[class]/[spec]` | picker card, rows card, `You wear` tag | landing answer for your character, ControlBar (spec, band, faction, race), Paperdoll pick/compare + Inspector, Where to go, notes, stat weights section | `/data/{build}/bis/{spec}.json`, loot, items, `/v1/me`, `/v1/home` | bis-out, bis-24, bis-34, bis-60, bis-new, bis-nogear, bis-loading, bis-error, bis-spec-60, bis-spec-healer-60, bis-spec-20 |
| 11 | Guides home and class page (`GuidesHomeView`, `GuidesScreen` / `GuidesHome`, `ClassPage`) | `/guides`, `/guides/[class]` | Your guide card, class cards | Guide / Tiers / What changed segments (R8), crest rows | guides index, dates | guides, guides-you, guides-class |
| 12 | Guide reader (`GuideReaderView`, `GuideBody` / `GuideReader`, `GuideSections`) | `/guides/[class]/[spec]` | At a glance card, caps strip | At a glance panel on the Plate, ReadingSections, CharacterAnswer for your character | `/data/{build}/guides/{spec}.json` | guide, guide-loading, guide-error (guide-you removed 2026-11-08: the reader has no per-character answer) |
| 13 | Tiers (`TiersView` / `TierLists`) | `/tiers` | caps bands | RankList rows with GradeTile groups, ControlBar role and faction, your spec marked | `/data/{build}/tiers.json` | tiers, tiers-you |
| 14 | What changed (`ChangesView` / `WhatChanged`) | `/changelog?show=game` | cards | Dated groups | `/data/changelog.json` | changes |
| 15 | Talents and Planner (`TalentsView`, `TalentTreeView`, `PlannerEditorView`, `YourBuildsView`, `SharedBuildView` / `TalentsPage`, `PlannerEditorPage`, `YourBuildsPage`, `SharedBuildPage`) | `/planner`, `/b/:id` | caps heads, card trees | TalentEditor (one tree behind tree tabs on phone), PointOrder, share line (R12) | `/data/{build}/talents/*`, `/v1/builds` | planner-you, planner-guide, planner-loading, planner-error, build-shared |
| 16 | Sims list (`SimsListView` / `SimsList`) | `/sim` landing | cards | ControlBar, Your sims rows, server runs restyle | `/v1/sims` | sim, sim-you, sim-running (sim-saved removed 2026-11-08: a saved sim opens sim-result) |
| 17 | Sim result (`SimResultView` / `SimResultContent`) | `/sim` result | result card, gear grid | answer Figure xl with error, vs last sim (R9), upgrade picks, ability StatTable, weights, Your sims; share | `/v1/sims/{id}`, `/progress` | sim-result |
| 18 | Reports list (`ReportsListView` / `ReportsList`) | `/logs` (yours) | This week / Earlier cards | Rows, empty and failed states, `Upload with the companion` wording | `/v1/reports` | reports, reports-empty, reports-failed |
| 19 | Report, meters, coach (`ReportView`, `MetersView`, `CoachCardView` / `ReportContent`, `MetersSheet`, `CoachCard`) | `/reports/...` | coach card, parses card, fights card | TabStrip, summary, player view, coach restyle, processing state | `/v1/reports/{id}`, meters, coach | report, report-player, report-coach, report-processing, report-loading, report-error |
| 20 | Rankings and filters (`RankingsView`, `FiltersSheet` / `RankingsBoard`, `FiltersSheet`) | `/rankings` | answer card, caps | ControlBar, RankList, your row pinned | `/v1/rankings`, `/percentile` | rankings, rankings-you, rankings-empty |
| 21 | Guild tab: Raids tab today (`RaidsTabView` / `RaidsTab`) | `/guild`, Overview | guild head band, Next raid card | Guild landing out/none, member and officer Overview: next raid countdown with sign-up first (R10), progression, guild sims ladder, readiness | `/v1/guilds/{id}/home`, `/raids`, `/v1/home guild_sims` | guild-out, guild-none, guild-member, guild-officer, guild-no-raid, guild-low-level |
| 22 | Raids list and raid page (`RaidPageView`, `CompositionView` / `RaidPage`, `Composition`, `SignupViews`) | `/guild/.../raids`, raid | count block card, signup card | Tally roles, SignupRow, In/Late/Bench/Out Segmented with haptic, roster rows with rank labels | `/v1/guilds/{id}/raids/{raid}` | guild-raids, guild-raid |
| 23 | Raid night (`RaidNightView` / `RaidNightScreen`) | none (app only) | night summary card | countdown, your sign-up buttons, who is coming, before you pull | `/v1/me/raid-nights/{raid}`, guild home `needs_before_next_raid` | raid-night-live, raid-night-signup |
| 24 | Roster, History, Calendar (`RosterScreen`, `HistoryScreen`, `CalendarScreen` both) | guild Roster, Loot, calendar link | attendance bars in cards | Rows with crest, rank label, attendance Meter; loot rows | guild home, loot, calendar feed | guild-roster, guild-loot |
| 25 | You (`YouView`, `YouScreen`) | `/account` | seal header, cards | character head with the stamp (R15), Guild and Logs rows, Account rows, Delete account row | `/v1/me`, `/v1/devices` | you, you-out |
| 26 | Account detail and Plan (`PlanView` / `PlanScreen`) | `/account`, `/premium` | caps plan card | PlanHead, Columns, store-rule copy (US link only) | `/v1/me`, `/v1/app/config` | account, account-premium, account-loading, account-error, premium |
| 27 | Devices (`DevicesView` / `DevicesScreen`) | `/account` devices | card | Rows, Revoke danger buttons, Sign out everywhere else | `/v1/devices` | devices |
| 28 | Delete account (`DeleteAccountView` / `DeleteAccountScreen`, the account-delete lane) | `/account` delete | lane's first cut | restyle only: What goes, When, the Premium store warning, failed and done | `DELETE /v1/me` | delete-free, delete-premium, delete-failed, delete-done |
| 29 | Offline data (`OfflineDataView` / `OfflineDataScreen`) | none | cards | Rows, Switch rows, the offline Notice | local | offline-data, offline-data-empty, offline-strip |
| 30 | Live surfaces settings (`LiveSurfacesScreen` / `LiveSurfacesCard`) | none | card | Rows with Switches on You > Notifications | local | (you) |
| 31 | Update required (`UpdateRequiredView` / `UpdateRequiredScreen`) | none | seal and card | the store button, the safe line | `/v1/app/config` | update-required |
| 32 | Widgets (WidgetKit `NextRaidViews`, `UpgradeViews` / Glance `NextRaidWidget`, `NextUpgradeWidget`) | none | caps, gradients | R13 content, flat `raised` widget, Barlow, first-raids fallback | `/v1/me/widget` (adds `first_raids_at`) | widgets-guild, widgets-noguild, widgets-open |
| 33 | Live Activity and ongoing notification (`LiveActivityViews` / `LiveNotifier`) | none | | R13 content, the four sign-up buttons in the expanded island | `/v1/me/live-activities` | live-activity |
| 34 | Control widget, QS tile, shortcuts (iOS `RaidControl`; Android tile, `SignupShortcutActivity`) | none | | restyle of the one line they draw | widget data | (as widgets) |

Counts: **34 screen groups, 31 on the old style** (all but the account-delete screens' copy, which only needs the restyle, and the two pure system surfaces with no styled body, the Control widget and the QS tile). **Missing features: 22** across them: sections array, import sim, first-raids countdown, feed and filters, guild sims ladder (home and guild), spec week, movers, standing line, character Activity, paperdoll on BiS, Where to go, BiS landing answer, ControlBar, At a glance panel, sim result answer and vs last sim, report restyle and processing state, the level-10 filter, the Battle.net stamp, the companion wording, Sign out everywhere else, sign in again, widget first-raids fallback; plus pull to refresh on Android Raids and Guides.

## 5. Mocks
`python3 design/mocks/ws_app.py && node design/mocks/ws_app_shoot.mjs` writes 96 screen states per platform (192 captures). States per area: signed out, loading, error, empty, signed in at 34 and at 60 in a guild, wherever the area has them (section 4.2's last column). Every capture is the whole scroll at 1x, the frame drawn: the build compares the content under the navigation bar (section 6), not the frame.

Captures are packed to a 256-colour palette (`python3 design/mocks/ws_app.py pack`); the gate reads the region files, not pixels.

Known mock notes:
- Two web phone mocks overflow the page at 393: `web-guild-raid` (666 px wide) and `web-guild-raids` (736 px). The Android and iOS lanes repair them in `ws_guild2.py` (stack the raid controls and the role chips below 768) before building those two screens, then regenerate the app mocks and region files. The report tabs, the sim timeline and the BiS cells scroll or ellipsize inside their own regions, which is intended.
- The web home mock keeps its two countdown panels side by side at 390 and overflows by 25 px (the built web stacks them below its breakpoint); the app frame stacks them (`ws_app_frame.APP_CSS`). The web reference mock should stack too (a follow-up for `ws_home.py`; it changes no web parity state at 1440).

## 6. API and data the apps need (named for the build lanes)
- `/v1/home`: `sections` (2.3). `/v1/feed` as built. `/v1/me/widget`: `first_raids_at` (from `dates.json`) for the no-guild widget.
- `POST /v1/devices/revoke-others` (Sign out everywhere else, R16): revokes every device token of the account but the caller's.
- `GET /v1/app/config`: `home_sections_out`.
- `vs last sim` (R9): the sim result needs the character's previous sim figure (`previous_dps` on `/v1/sims/{id}` when the sim has a character and a previous one exists); without it the line is left out.
- Item tooltips on phones load the item from the static items file already cached; no new endpoint.

## 7. Native parity
In `design/parity.md`, section "Native". Summary: each mocked state renders natively in fixture mode with the mock's data and clock on the snapshot tool each app already has (swift-snapshot-testing 1.19.7, Roborazzi 1.77.0); the primitives report the web's `data-region` names; a test compares every region's presence, inventory and box (8 pt / dp, or 5%) with `design/mocks/app/regions/<platform>-<screen>.json`; goldens stay as the self-regression gate; a side-by-side PNG per state is uploaded. Stale region files fail the web `parity` job. Pixels are not compared with the mock (three text renderers), the same choice the web gate made. Text is compared strictly against the build's own fixture render: every state also records a text golden (the accessibility tree's labels and values in reading order), so a copy change fails until it is re-recorded on purpose; the image goldens tighten to `precision: 1, perceptualPrecision: 0.98` on iOS and `changeThreshold = 0` on Android (the account-delete lane saw a signed-out snapshot pass with changed text).

## 8. Build plan
Three lanes. Lane T first (both app lanes depend on it); then iOS and Android in parallel, screen by screen in the order below, each screen landing with its parity states green.

### Lane T: tokens, copy, home sections (first; one lane)
1. `design/tokens/tokens.json` with every group of 2.1; `design/tokens/build.mjs`; the generated `tokens.css` (no visual change: run every `parity-*` spec, all green, no mock moves), `DesignTokens.swift`, `Tokens.kt`; `tokens.test.mjs` in the web unit job and `mobile-codegen.yml`; delete `gen-mobile-tokens.mjs` and `mobile-tokens.test.ts`; update `apps/shared/README.md`.
2. Fonts: Marcellus into `UIAppFonts` (app and widget extension) and the Android shared-fonts plugin; remove Barlow Condensed and Cinzel from both apps (keep the files only if the API card renderer needs them; it has its own copies).
3. `apps/shared/copy/shared.json`, the web modules of 2.2 reading it (no call site changes), the generator emitting it into both apps; the copy fixes (`Get the addon and companion`, `Updated ... from Battle.net`).
4. API: `/v1/home` `sections` with the R2/R4 rules and the `HOME_SECTIONS` switch; `/v1/app/config` `home_sections_out`; `/v1/me/widget` `first_raids_at`; `POST /v1/devices/revoke-others`; `previous_dps` on sims; `api/openapi.yaml`, regenerated clients, contract fixtures.
5. (Follow-up, R0) `/v1/home` `sections` composes R0's order (countdowns second in a guild, after the feed without one) and the Home mocks follow it, with their region files rendered on the CI runner.
6. The region modifier (`parityRegion`) and the region-capture test helpers on both platforms; the text goldens and the tightened image tolerances (parity.md Native), then re-record every golden on CI (`ios.yml` `record`, `:app:recordRoborazziDebug`) so the account-delete screens' signed-out snapshot carries its real text; wire `app_regions.mjs --check` into the web `parity` job.

### iOS lane (after T)
Shared work first: the primitives of 3.4 in `Sources/DesignSystem` (each with a snapshot of its states), the surface guard with its allowlist, delete the retired styles as screens move off them. Then, per screen, restyle to the mock, wire the missing data, record goldens on CI, pass the parity states:
0. **Push (first, R0).** The `aps-environment` entitlement and the Push Notifications capability in `project.yml` (the TODO beside the entitlements), registration and the device token through `POST /v1/me/push-tokens`, the notification categories that open a raid and its sign-up in one tap, and the APNs sender in the API (the GSM secrets `APNS_KEY_P8`, `APNS_KEY_ID` and `APNS_TEAM_ID` exist). The pushes: a raid posted or moved, the sign-up deadline, raid night starting, a sim done, a loot upgrade, a tier move, a guildmate passing you.
1. Tab bar and frame (R1: Guild replaces Raids, five tabs, labels 12), navigation bar title behaviour (N3), pull to refresh (R12), haptics (R14).
2. Home: sections renderer (known types, skip unknown), head figures, upgrades, countdowns, feed with filters and floor, guild sims, spec week, movers, Next; signed out; switcher sheet with the level-10 fold and the stamp; sign in again. States: home-*, switcher*, expired.
3. Character (absorbs Gear): head answer, standing line, Paperdoll worn + Inspector, fights, score, builds, Activity, History; not found. States: character-*.
4. BiS: landing answer, ControlBar, Paperdoll pick/compare + Inspector, Where to go, notes, weights; item tooltip sheet; search. States: bis-*, search*.
5. Guides (segments), guide reader, tiers, what changed. States: guides*, guide*, tiers*, changes.
6. Sims list and result (share), planner and talents, shared build. States: sim*, planner*, build-shared.
7. Reports list, report, player, coach, meters; rankings. States: reports*, report*, rankings*.
8. Guild tab (landing, member, officer, no raid, low level), raids list, raid page, raid night, roster, loot, calendar. States: guild-*, raid-night-*.
9. You, account, plan, devices, delete account (restyle the account-delete lane's view only after it lands), offline data, update required, live surfaces. States: you*, account*, premium, devices, delete-*, offline-*, update-required.
10. Widgets, Live Activity, Control widget. States: widgets-*, live-activity.

### Android lane (after T, same order and states)
0. **Push (first, R0):** FCM (`google-services`, the token through `POST /v1/me/push-tokens`, the notification channels, a raid and its sign-up in one tap) and the FCM sender in the API, the same pushes as iOS. Then the same list against the Compose screens (`ForeverTabScaffold`, `HomeRoute`, `CharacterPage`, `BisContent`, `GuidesHome`/`GuideReader`/`TierLists`/`WhatChanged`, `SimsList`/`SimResultContent`, `PlannerEditorPage`, `ReportsList`/`ReportContent`, `RankingsBoard`, `RaidsTab` → Guild, `RaidPage`, `RaidNightScreen`, `RosterScreen`/`HistoryScreen`/`CalendarScreen`, `YouScreen`, `DevicesScreen`, `DeleteAccountScreen`, `OfflineDataScreen`, `UpdateRequiredScreen`, Glance widgets, `LiveNotifier`, the QS tile), in `core/designsystem` for the primitives. Android adds pull to refresh on Guild and Reports (R12), predictive back, and the Material 3 `NavigationBar` styling of 4.1.

### Rules for both app lanes
Ship mode: locally the touched unit tests, the touched screens' snapshot and parity tests, lint and format; CI is the full check. One land at a time through `tools/land.sh`. A screen lands only with its parity states green; a feature the old screen has and the mock lacks is added to the mock (`ws_app.py`) and kept (parity never costs features). Every character image is the crest; every item has its icon, quality colour and tooltip. No em dashes. Captures at the default and the largest text size are attached to the land for each screen.
