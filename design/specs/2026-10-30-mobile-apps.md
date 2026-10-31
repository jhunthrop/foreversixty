# Forever Sixty for iPhone and Android: experience spec

Author: ux-designer, 2026-10-30 (lane D of `docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md`). One spec for two native apps built at the same time: iOS in SwiftUI (iOS 26+) and Android in Kotlin and Compose with Material 3 (API 28+). Tenets: `docs/tenets.md` 1 to 14, with 2 and 9 to 14 applied screen by screen. Owner direction (2026-10-30): native apps, built together; raid sign-up in the app instead of Discord; push for the lifecycle catalogue; free with the web's Premium rules; no in-app purchase, and "See plans" links to web checkout only where the store allows it.

A wow-player review follows this one. Wave 1 builds from this spec and lane C's fixtures; nothing here asks either app to compute BiS, gains, scores or entitlements (plan guard rail "one behaviour, one place").

## Revision 2 (2026-10-30), the fix round on the wow-player HOLD

The verdict is `design/reviews/2026-10-30-mobile-apps/player-verdict.md`. The coordinator's rulings are binding and are applied here.

Blockers:
1. **Talents (F6).** The tab totals are summed from the drawn build (30, 3, 18 = 51), never typed. The count is the planner's `diffCount` (`web/src/lib/planner/band-compare.ts`): your points spent where the guide spends fewer. `Matches` shows only when no cell differs. Cells and prerequisite lines follow the planner's `cellState` and `connectorPath`. The differing example is now a legal build: Lightning Overload's prerequisite (Elemental Alacrity 3) still holds.
2. **Loot.** The loot example is Sapphiron Drape, an Onyxia drop in `loot.json` (`raid:onyxias-lair`). `loot.awarded` is off by default as a standalone push and folds into the after-raid summary.
3. **New event `raid.roster_moved`** (§6.2, boards `f7-moved-*`): time sensitive, the high channel, and it comes through quiet hours.
4. **F4.11** now shows In checked after a refused Late.

Rulings on the shoulds:
5. **Officer tools in v1.** Officers lock and unlock sign-ups and move raiders between In and Bench, or off the raid, from the raid page (§5.4 F4.18 to F4.21). Planning stays on the web.
6. **One after-raid push.** `raid.summary` replaces `report.ready`, `coach.ready`, `ranking.best` and `loot.awarded` as pushes. It opens a new screen, Your raid night (F4.22).
7. **Invite time.** It is a raid field, `invite_at`, shown beside the pull. `raid.starting` fires at the invite time minus the lead.
8. **Tooltip parity.** The tooltip gains the Binds line, Unique, and the armor type, all from the client's ItemSparse and Item rows. Source lines carry reputation standings (`Warsong Outriders (Exalted)`).

Other shoulds folded in:
- the iOS guild subtitle;
- `raid.starting` reads `in 1 hour`;
- the unknown-spec push uses Vorgash, an alt with no saved spec, and draws `Out`;
- F4.12 draws both character blocks with a level-60 alt;
- `Also planned 2` counts only the listed rows;
- the notes say phase 2;
- `band.entered` names the first upgrade;
- `game.milestone` skips the launch push for accounts with a synced character;
- F2.7 names the pick;
- the F3.1 chips wrap;
- F3.2's back line only offers the way back when the picks differ.

Shoulds skipped, with the reason:
- **Client stat wording** (`Equip: Improves your chance to hit…`). The site prints `+10 Hit` by decision, because the repo has no verified per-point-to-percent table (`lib/items/tooltip.ts`, tenet 5). The app matches the site.
- **Soft reserves.** These need a raid field and a site surface first. Planned for v1.x.
- **Item lookup beyond the 17 rows.** It needs a search endpoint over the build's items. Planned for v1.x.
- **Quest giver and zone.** The site does not show them yet. The app follows the site when it does.
- **Two-column gear grid on F5.3.** The grid already runs in character-frame order. A long press shows the slot label (§5.5).
- **The four Phase-1 BiS source questions.** These concern the site's BiS data and are routed separately by the coordinator.

## Mock

`design/mocks/2026-10-30-mobile-apps/`: `gen_mobile_apps.py` (with `kit.py`, `model.py`, `screens_home.py`, `screens_raids.py`, `screens_more.py`) writes `renders/*.html`; `render.mjs` screenshots them with Playwright from `web/node_modules` at 1.5x. `extract_data.py` rewrites `data.json` and `assets/` from a checkout holding the synced static data and the raw client tables. Run from the repo root:

```
python3 design/mocks/2026-10-30-mobile-apps/gen_mobile_apps.py && node design/mocks/2026-10-30-mobile-apps/render.mjs
```

The item data is real: build `1.60.1.70532`, Elemental Shaman, level 60, Phase 1 raid preset, Horde, copied into `data.json` with its icons in `assets/` (134 files). Lightningbolltz wears listed alternatives in eight slots, so each gain on the boards is the band's own measured `dps_delta`. The guild roster, raid times and sim figures are invented for the mock. The talent tree maps the band's talent string to the tree file's list order; the build uses the planner's own decoding.

Every board is drawn on iPhone 17 (402 x 874 pt) and Pixel 9 (412 x 915 dp). `-ios` and `-and` suffixes name the platform.

| Board | What it shows |
|---|---|
| `f1-signin-{ios,and}` | F1.1 signed out, F1.2 system sign-in, F1.3 import, F1.4 sign-in failed, F1.5 the switcher sheet |
| `f2-home-{ios,and}` | F2.1 the home answer, F2.2 loading, F2.3 offline, F2.4 no gear export, F2.5 error |
| `f2-items-{ios,and}` | F2.6 the item tooltip sheet, F2.7 your gear, all slots |
| `f3-bis-{ios,and}` | F3.1 the list, F3.2 the picker sheet, F3.3 offline |
| `f4-raids-{ios,and}` | F4.1 Raids tab, F4.2 raid page, F4.3 spec ask, F4.4 the reminder ask, F4.5 the More menu |
| `f4-raid-full-{ios,and}` | F4.6 the raid page scrolled top to bottom, roster included |
| `f4-states-{ios,and}` | F4.7 unsigned, F4.8 locked, F4.9 cancelled, F4.10 offline, F4.11 save failed, F4.12 moved to Bench |
| `f4-tab-states-{ios,and}` | F4.13 no guild, F4.14 waiting for approval, F4.15 no raid planned, F4.16 error, F4.17 offline |
| `f4-officer-{ios,and}` | F4.18 officer view of an open raid, F4.19 the officer roster, F4.20 the move menu, F4.21 locked with the toast |
| `f4-night-{ios,and}` | F4.22 Your raid night (Premium), F4.23 the same for a free player |
| `f7-moved-{ios,and}` | F7.7 `raid.roster_moved`: moved In, moved to Bench, removed, added |
| `f4-pair` | Raids tab and raid page, iOS and Android side by side |
| `f5-sims-{ios,and}` | F5.1 sims (Premium), F5.2 sims (free), F5.3 to F5.5 the result for Premium, free in a link region, free elsewhere |
| `f6-talents-{ios,and}` | F6.1 matches the guide (totals 30, 3, 18), F6.2 differs by 2 points |
| `f7-push-ios` | F7.1 lock screen, F7.2 long-press with In, Late, Out, F7.3 the replaced answer and loot |
| `f7-push-and` | F7.1 the shade with actions, F7.3 the answer and the unknown-spec case |
| `f7-settings-{ios,and}` | F7.4 notification settings, F7.5 the channel picker, F7.6 permission denied |
| `f8-you-{ios,and}` | F8.1 free in a link region, F8.2 free elsewhere, F8.3 Premium, F8.4 Plan, F1.6 Devices with Revoke |
| `xl-ios`, `xl-and` | iOS AX5 and Android font scale 2.0: home, raid page, switcher, notification settings |

# 1. The one question, and the answer on the first screen

"First screen" means no scroll on an iPhone 17 (402 x 874) and a Pixel 9 (412 x 915) at the default text size, with the tab bar drawn. The smallest supported phones (iPhone SE 3, 375 x 667; Android 360 x 640) must show the same answer, with only the lower rows cut.

| Surface | The player's question | Visible without scrolling |
|---|---|---|
| Home (signed in) | "What should this character do next?" | The character (selector), the answer "Upgrades for your level range" with its two figures, and the top five upgrades, each with its icon, gain and what it replaces (F2.1) |
| Home (signed out) | "What is this for me?" | "Play your class better.", `Sign in with Battle.net` and `Pick your class` (F1.1) |
| Raids tab | "When is the next raid, and am I in?" | The next raid's date, time, countdown, count, and my In / Late / Bench / Out control with my state checked (F4.1) |
| Raid page | "Am I in, as what, and who else is?" | Title, when, countdown, count and role breakdown, then my sign-up card (F4.2). The roster is the first scroll. |
| Raid notification | "Can I answer without opening the app?" | In, Late, Out on the notification (F7.2) |
| Your raid night | "What did I get tonight, and how did I do?" | Kills and wipes, loot awarded to me with its icon, my personal bests (F4.22) |
| Raid page, officer | "Who is coming, and can I set the roster from here?" | The officer row (Lock sign-ups) under the count; the roster's place buttons on the first scroll (F4.18, F4.19) |
| BiS | "What is best in each slot for my spec and range?" | The picker chips and the first seven slots, each with source and the next best (F3.1) |
| Sims | "Did my run finish, and what did it do?" | Running runs with progress, then finished runs with DPS (F5.1) |
| Talents | "Is my build the guide build?" | The Matches or Differs tag with its sentence, and the tree (F6.1) |
| You | "What plan am I on, and what will the app tell me?" | Account, Plan, Notifications (F8.1) |

Tenet 14: no screen opens on a slogan, a feature list or an onboarding carousel. A signed-out first launch is the Home tab in its signed-out state; there is no welcome flow.

# 2. References, and what we borrow

| Reference | Borrowed | Not borrowed |
|---|---|---|
| The in-game character frame | Slot names and order for "Your gear" (F2.7), quality-ringed item icons | The paper doll; it does not fit 402 pt |
| The game's item tooltip | The tooltip sheet's body (F2.6): name in quality colour, slot, armor, stat lines, Requires Level, Item Level in the client's gold `#e5cc80`, on a near-black panel | Hover; a tap opens it |
| The in-game Calendar invite | One tap per state on the event; the list grouped by answer; realm time beside local time | The month grid |
| Raid-Helper's Discord post | Roster grouped by role with each group's count, spec icon then class-coloured name, Late, Bench and Out as their own groups, the per-player note | Emoji walls, numbered order, pings |
| Our own raid page (`design/specs/2026-10-26-guild-raids.md`) and spec ask (`design/specs/2026-10-28-discord-guided-cards.md` §P.10) | Every rule, chip, copy line and the R3 one-character-In rule, verbatim | The officer form; officers plan on the web in v1 (§17 Q1) |
| Our Discord `/char` panel (`design/specs/2026-10-29-discord-gear-panel.md`) | The upgrade order (measured gains high to low, then `Gain not modelled yet`), "Replaces {worn}, ilvl {n}" | Its Discord component limits |
| Our home answer (`web/src/lib/home-panel-copy.ts`, `home/header-answer.ts`) | The label, the two figures at one size, the note, every empty and error sentence | The site's grid |
| iOS Settings and Android system settings | Notification settings as grouped lists with the value trailing each row | Toggle-per-event; a channel is four values, so a toggle is wrong |
| Apple Calendar and Google Calendar notifications | Actionable notifications that answer in the background and replace themselves with the result | Opening the app for a one-word answer |
| Details! | One number per row, right-aligned, tabular | Bars per row; a list of six gains needs no bars |

# 3. Information architecture

## 3.1 Four tabs: Home, Raids, BiS, You

I keep the suggested tabs. The defence:

1. **Home is the character** (tenet 14). The tab's screen title is the character selector itself, not the word "Home". Sims and Talents are about this character, so they push from Home's "Next for {name}" rows instead of taking tabs. One job per page holds: Home links to them, never restates them.
2. **Raids is the headline and the notification target.** Every raid notification lands here. A tab named "Guild" was considered and rejected: the job is the raid, and guild facts (name, faction, roster count) are the header of this tab. The tab says what the player does in it.
3. **BiS is the one surface that works signed out and offline.** It is also the most visited site page. It earns a tab.
4. **You** holds the account, plan, notifications, characters, devices and Discord link. Settings live here and nowhere else.
5. **Five tabs (adding Sims) rejected.** For a free player a Sims tab would be a page about Premium, which breaks the no-nag rule (premium-entry spec §2). Server runs are reachable from Home and from the `sim.finished` notification.

## 3.2 The idiom per platform

| | iOS (SwiftUI) | Android (Compose, Material 3) |
|---|---|---|
| Tabs | `TabView` with four `Tab`s; the iOS 26 floating glass tab bar; minimises on scroll down (`.tabBarMinimizeBehavior(.onScrollDown)`) | `NavigationBar` with four `NavigationBarItem`s, labels always shown, 80 dp tall above the gesture inset |
| Each tab | Its own `NavigationStack`; re-tapping the active tab pops to root, then scrolls to top | Its own back stack (Navigation Compose, one nested graph per tab); re-selecting pops to root, then scrolls to top; switching tabs keeps each stack |
| Title | Root screens: the selector as the leading toolbar item (Home, Raids, BiS) or a large title (`You`). Pushed screens: large title collapsing to inline on scroll | Root screens: the selector in `TopAppBar`'s title slot. Pushed screens: `TopAppBar` with back arrow and the title, `enterAlwaysScrollBehavior` |
| Sheets | `.sheet` with detents `.medium` and `.large`, grabber visible | `ModalBottomSheet`, drag handle, skip-partially-expanded off |
| Menus | `Menu` on a toolbar button | `DropdownMenu` anchored to the overflow icon |
| Confirmations | `.alert` (two buttons) for destructive acts; `.confirmationDialog` for Sign out | `AlertDialog` |
| Back | Edge swipe and the back button; sheets swipe down | System back gesture and predictive back (`android:enableOnBackInvokedCallback="true"`); a sheet closes first |

## 3.3 Screen map

| Tab | Root | Pushes |
|---|---|---|
| Home | F2.1 the answer | F2.7 Your gear (all 17 slots); F6 Talents; F5.1 Sims, then F5.3 a result |
| Raids | F4.1 Raids tab for the current character's guild | F4.2 raid page; F4.22 Your raid night (from the raid page of a done raid, or the push) |
| BiS | F3.1 the list for the current character's spec, range and faction | none; the picker is a sheet |
| You | F8.1 | F8.4 Plan; F7.4 Notifications; Characters; F1.6 Devices; Discord |

Sheets from anywhere: F1.5 the switcher (from the selector), F2.6 the item tooltip (from any item icon or name), F4.3 the spec ask.

## 3.4 One way to switch

The selector (memory rule "one way to switch") is the only character switcher. It sits on Home, Raids and BiS, in the same place and the same look. It shows the class crest (34 pt), the full name in the class colour, the short description (`Elemental · 60`) and a chevron. A tap opens F1.5. Picking a character changes Home, Raids (its guild) and BiS (its spec, range and faction, while BiS shows `Yours`). You has no selector; it is the account, not a character.

The full name shows at every width and every text size. Names are one word, so they cannot wrap at a space. From iOS AX1 and Android 1.3x the selector's name stops growing (iOS: `.dynamicTypeSize(...DynamicTypeSize.accessibility1)` on the selector, 28 pt; Android: the name in `sp` with a 22 sp ceiling) and the description line drops. On iOS a long press shows the Large Content Viewer with the name (`.accessibilityShowsLargeContentViewer()`). Inside the switcher sheet the name keeps growing and breaks inside the word (`overflow-wrap: anywhere` in the mock; `lineBreakMode` character wrapping in SwiftUI, `LineBreak.Simple` with `Hyphens.None` and `softWrap` in Compose) rather than being cut (XL.3).

The current character persists per device: the last one picked, else `main_character_key`, else the most recently synced character. Picking writes nothing to the server.

**Deviation (integration C, 2026-10-31).** On Android the selector is a full-width row at the top of each root screen (Home, Raids, BiS), not the TopAppBar title slot: 16 dp side padding, 8 dp above and below, a 48 dp minimum height at every font scale. While a signed-in session's characters are still arriving the row is reserved at that height, so the screen never shifts when the selector appears. On BiS the search button shares the row. iOS keeps the toolbar's leading slot.

# 4. Brand on mobile

## 4.1 Tokens

Lane T generates `DesignTokens.swift` and `Tokens.kt` from `web/src/styles/tokens.css`. The apps use those names and nothing else. Colours used here:

| Use | Token | Value |
|---|---|---|
| Screen ground | `--color-bg` | `#07090d` |
| Bars, list cards, sheets | `--color-raised` | `#0d111a`; sheets `#10141d` is `--color-raised` mixed 4% white, which lane T adds as `sheet` (one new derived token, flagged for review) |
| Primary card | `--surface-card`, `--color-line-card`, `--card-highlight` | the warm gradient, its 17% gold border and 1 px inset highlight |
| The hero band | `--surface-band` | the ember radial over `#14100b` |
| Dividers | `--color-line-soft` | `#1c2230` |
| Text | `--color-strong`, `--color-text`, `--color-muted`, `--color-nav` | `#f2eee4`, `#e9e4d8`, `#9a9484`, `#b9b3a4` |
| Accent | `--color-gold`, `--color-gold-hover`, `--tile`, `--tile-ink` | `#e5b955`, `#f5d27a`, gold 12%, `#f5d27a` |
| Good and bad | `--color-kill`, `--color-wipe` | `#7fd48a`, `#f0736b` |
| Classes | `--color-class-*` | Shaman text `#3f8fe0` (the site's AA variant) |
| Quality borders | `--color-rarity-*` | rare `#0070dd`, epic `#a335ee` |
| Quality text | `--color-rarity-rare-text`, `--color-rarity-epic-text`, others as their border | `#3d94f0`, `#b866f5` |
| Factions | `--color-horde`, `--color-alliance` | marks only |

Radii: `--radius-panel` 12 for cards and sheet content cards; `--radius-control` 8 for buttons, inputs, the segments and date blocks; `--radius-pill` 6 for tags and item icons. Sheets use the platform's own top radius (iOS system sheet, Material 28 dp).

Spacing: 16 pt/dp screen margin on both platforms; 12 between an icon and its text; 14 to 22 between sections (22 before a section label); list rows min 56, item rows min 72.

## 4.2 Type

Barlow for body, Barlow Condensed 800 uppercase (`.01em`) for headings and figures, JetBrains Mono for every time, count, DPS figure and countdown, Cinzel only in the wordmark (F1.1, F1.3). All four ship in the app bundle (OFL), so there is never a font swap.

Each role scales with the system setting. iOS: `Font.custom(_:size:relativeTo:)`. Android: `sp`.

| Role | Face | Default size | iOS `relativeTo` | iOS AX5 | Android 2.0 (non-linear) |
|---|---|---|---|---|---|
| `title` (large title, pushed screens) | Barlow Condensed 800 | 34 (`--text-h1-phone`) | `.largeTitle` | 60 | 50 |
| `h2` (raid title on Raids, section heads like `Sign-ups`) | Barlow Condensed 800 | 26 (`--text-h2`) | `.title` | 56 | 42 |
| `stat` (the answer's figures) | Barlow Condensed 800 | 28 (`--text-stat`) | `.title` | 58 | 44 |
| `sec` (card titles, sheet titles) | Barlow Condensed 800 | 18 (`--text-card-title`) | `.headline` | 53 | 34 |
| `name` (selector, switcher rows) | Barlow 700 | 17 | `.headline` | 53 (28 cap in the bar) | 32 (22 cap in the bar) |
| `body` | Barlow 400/600 | 15/22 (`--text-body`) | `.subheadline` | 49 | 30 |
| `meta` | Barlow 400 | 13 (`--text-meta`) | `.footnote` | 44 | 26 |
| `micro` (section labels, uppercase .14em) | Barlow 700 | 11 (`--text-micro`) | `.caption2` | 40 | 22 |
| Segments | Barlow 600 | 14 | `.subheadline` | 44 | 28 |
| Tab labels | system | system | not scaled; iOS Large Content Viewer | system | system |

## 4.3 Images

- **Crests.** Every character and class image is the circular ringed HD crest (`web/public/icons/hd/crests/{class}.webp`, bundled), with the 2 pt ring at 60% class colour and the 14 pt glow (`.crest` in the site). Never a square class icon, a letter tile or a Blizzard avatar (memory rule "one crest language"). Sizes: 34 selector, 44 switcher rows and BiS picker, 36 sign-up cards, 30 the classes strip.
- **Spec icons** are the spec's talent-tab icon (`sim/specs` `Icon`), square with `--radius-pill`, 1 pt `#3a4256` border. They mark roster rows (28), the sign-up line (18), the spec ask (40), the BiS chip (28), and sims (36). This is the owner's ruling for roster marks (guided-cards spec §P.9.1).
- **Item icons** come from `/data/{build}/icons/{icon}.webp` on the CDN, square with `--radius-pill`, a 1.5 pt border in the quality colour, a 1 pt inner black line and a 10 pt quality glow at 22%. Sizes: 44 upgrade and BiS rows, 36 gear rows and next-best rows, 56 tooltip header, 40 sim gear grid.
- **Item names** are always in their quality text colour. An item is never a bare name (tenet 2).
- **Guild mark** is the guild crest (`GET /v1/guilds/{id}/crest.webp`), circular 40; with none, the faction mark, as the site's `GuildCrest` falls back.
- **App icon** is `design/logo/export/foreversixty-icon-1024.png`. The Android notification small icon is the seal as a monochrome silhouette (lane T draws `ic_stat_seal`).

## 4.4 Dark only

The apps are dark only, on both platforms. iOS sets `UIUserInterfaceStyle` to `Dark` in Info.plist. Android uses a dark `MaterialTheme` colour scheme built from the tokens and ignores the system setting. Reasons:
1. The site and the in-game UI are dark only. A light theme would be a second design system to verify (tenet 10).
2. The quality and class colours are calibrated against `#0d111a`. Epic `#a335ee` and the game's Shaman blue fail AA on white, and the game's own tooltip is dark.
3. Raid use is at night, in a dark room, beside a game on a dark screen.

Increase Contrast (iOS) and high-contrast text (Android) swap `--color-muted` for `--color-nav` and raise the line tokens one step (`--color-line-soft` to `--color-line`).

## 4.5 Shared controls (tenet 10)

The apps reuse the site's controls under the same names; each platform draws them natively. A lane never invents another.

| Control | Site source | iOS | Android |
|---|---|---|---|
| Primary button | `PRIMARY_BUTTON_FIXED` | `Button` with a gold gradient fill, 48 pt tall, 13 pt 700 uppercase .07em | `Button` with the same fill, 48 dp |
| Secondary button | `SECONDARY_BUTTON_FIXED` | 4.5% white fill, gold 17% border | `OutlinedButton` |
| Text link | `.alink` | `Button` with `.plain` style in gold 15 pt 600, 44 pt hit | `TextButton` |
| Tag | `.pill` / `.tag` | 22 pt capsule, 11 pt 700 uppercase | same |
| Sign-up segments | `guild/SignupSegments.svelte` | a custom `HStack` of four 48 pt buttons (§5.4.3) | a `Row` of four 48 dp `selectable`s in a `selectableGroup` |
| Progress bar | the 8 px bar | 8 pt, `--color-line-soft` track, gold fill | same |
| Date block | Raids tab date block | 52 x 56, `--radius-control`, mono day | same |
| Switch | the site's gold switch | `Toggle` with `.tint(gold)` | `Switch` with gold track colours |
| Item tooltip | `ItemTooltip.svelte` model (`lib/items/tooltip.ts`) | a sheet (F2.6) | a bottom sheet |
| Skeleton | `Skeleton` | `redacted(reason: .placeholder)` on real layout | shimmer boxes at the real sizes |
| Offline strip | new on mobile; see §5.0 | | |

The offline strip is the one new shared component this spec introduces. It needs its own small design review per tenet 10; its spec is §5.0.

# 5. Screens

## 5.0 Rules for every screen

**States.** Each screen has, at least: loading, content, empty, signed out, error and offline.
- **Loading.** On a first load with nothing cached, skeletons at the real row heights. A refresh never replaces content with skeletons. A refresh shows the system pull-to-refresh spinner, or nothing when automatic.
- **Offline.** When the last fetch failed for lack of a connection and a cached answer exists, the cached answer shows under the offline strip. The strip is 36 pt tall, `#11151f` on a `--color-line` bottom border, with a 16 pt wifi-off glyph and the text `Offline. Showing what you had at {time}.` (`{time}` is `9:41 pm` today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`). Writes are disabled with their own sentence. The strip leaves when a fetch succeeds. No strip for a server error: that is the error state.
- **Error.** The region that failed says so in `--color-wipe` with `role` alert, plus `Try again` (secondary button). Other regions keep their content.
- **Signed out.** Home, Raids and You show their signed-out state. BiS works fully.

**Caching.** Each API answer is cached on disk per account (iOS: a file per endpoint in Application Support, excluded from backup; Android: DataStore or Room). The static data (`/data/{build}/…`) is cached per build and kept until a newer build is in use. A cached answer draws at once on launch, then refreshes. Refresh on foreground when the cache is older than 5 minutes, on pull, and on a push for that object.

**Navigation into a screen from a notification** opens the right tab, pushes the target and keeps the tab's root beneath it, so back goes where a player expects.

## 5.1 F1. Sign-in, account and the switcher

### F1.1 Signed out (Home tab)
- Ground: `--surface-band` full bleed.
- From 60 pt below the bar: the seal (64), the wordmark `Forever Sixty` in Cinzel 22, then `Play your class better.` in `title`, then the body line `Sign in and your characters arrive with their gear, talents and guild.` at 17/25.
- `Sign in with Battle.net`: primary, full width, 52 tall. The screen's one primary.
- `Pick your class`: secondary, full width, 52 tall. It switches to the BiS tab with the picker sheet open.
- Meta line: `Pick your class opens best in slot, talents and sims for any spec, no account needed.`
- The tab bar shows. Raids signed out reads the same sign-in panel with the heading `Sign in to sign up for raids` and the line `Your guild's raids show here once you sign in with Battle.net.` You signed out shows the account card replaced by the same two buttons, then Help, Privacy, Terms and About.

### F1.2 System sign-in
- `Sign in with Battle.net` opens `GET /v1/auth/app/start` (lane A1) in the system browser: iOS `ASWebAuthenticationSession` with `callback: .https(host: "foreversixty.gg", path: "/app/auth")` and `prefersEphemeralWebBrowserSession = false`, so Battle.net remembers the player; Android a Custom Tab, returned by the verified App Link `https://foreversixty.gg/app/auth`.
- iOS shows the system consent (`"Forever Sixty" Wants to Use "foreversixty.gg" to Sign In`). The app never draws its own copy of it.
- The button shows its pressed state, then a 20 pt spinner in place of its label while the session is open. It is disabled until the session returns.
- Cancel (the player closes the sheet or tab): back to F1.1 with no message.

### F1.3 Import
After the code exchange (`POST /v1/auth/app/exchange`), the app stores the device token (Keychain; Android encrypted DataStore with Tink) and shows the seal, `Bringing in your characters` (`h2`) and `Reading Battle.net. This takes a few seconds.` with an indeterminate 8 pt bar. It ends when `GET /v1/me` answers. With characters: Home for the current character (§3.4). With none: Home's no-characters state:
- Heading `No characters yet`.
- Body `Battle.net lists a character once it has logged in on Forever. Log in once, then pull to refresh.`
- Secondary button `Get the addon`, which opens foreversixty.gg/setup in the browser.

### F1.4 Sign-in failed
F1.1 with a toast at the top (12 pt below the status bar, `--color-wipe` 45% border, `role` alert), held until the next tap. The text is the API's error message. Known texts:
- `Battle.net could not confirm that sign-in. Try again.`
- `This sign-in has expired. Start again.`
- `No connection. Signing in needs the internet.` (no request made)

The button is enabled again.

### F1.5 The switcher sheet
- Opens from the selector. The selector shows its open state: a full gold border and the gold glow.
- Sheet title `Your characters`. Under it, meta: `{battletag} · Signed in with Battle.net`.
- One 68 pt row per character: crest 44, the name (`name`, class colour), `{Race} {Spec} {Class} · Level {n}` and the guild name (or `No guild`). The current character has the `--row-hover` wash and a gold check.
- Footer meta: `A character shows here after it logs in on Forever, or after the addon syncs it.`
- A tap picks the character, closes the sheet, plays the selection haptic and redraws the tab under it.
- The sheet is never empty: with no characters the selector is not drawn.

### F1.6 Devices (You › Devices)
- One row per device from `GET /v1/devices`: tile, name, meta. Names: `This iPhone` or `This phone` first, then the others (`Pixel 9`, `Companion on Windows`). Meta: `Last used 2 h ago`, or `Signed in 30 Oct` for this device.
- Other devices carry `Revoke` (secondary, `--color-wipe` text and border). Revoke asks:
  - Title `Revoke {device}?`
  - Body `It is signed out of Forever Sixty at once.`
  - Buttons `Cancel` and `Revoke` (destructive role).
- On success the row leaves with the standard list deletion animation. On failure: `Could not revoke that device. Try again.` under the row.
- Footer meta: `Revoke signs a device out at once. To sign this one out, use Sign out on You.`

### Sign out
- You › `Sign out` (secondary, `--color-nav` text) asks `Sign out of Forever Sixty on this {iPhone|phone}?` with `Sign out` (destructive) and `Cancel`.
- Sign out calls `DELETE /v1/devices/self`, unregisters the push token (lane A2) and deletes the token and every cache.
- The app returns to F1.1. A 401 on any request does the same, without asking, and shows the toast `You were signed out. Sign in again.`

## 5.2 F2. Your character (Home)

### F2.1 The answer
Top to bottom on Home, under the selector bar:
1. **Hero band** (`--surface-band`, 14/16/18 padding, `--color-line-soft` bottom border):
   - meta `{Race} {Spec} {Class} · Level {n} · {Guild}`.
   - 14 pt gap, then the `micro` label `Upgrades for your level range`.
   - **The two figures**, the region's one primary element: `{n}` in `stat` with `upgrades` in body muted; then `+{x.x}` in `stat` with `DPS` (or `points` for a tank band) in body muted. Both are the same size; a 22 pt gap between them. With no measured gain, only the count.
   - The note, meta muted: `Best in slot for {band label}, together`, plus ` · {n} gains not modelled yet` when any upgrade has no figure.
   - The sync line, meta `--color-nav` with a 14 pt refresh glyph: `Synced {age} from the addon` (or `from Battle.net`).
2. **Your upgrades** card (`--surface-card`):
   - Head row: `micro` `Your upgrades` left; meta muted band label right (`60 · Phase 1 raid`).
   - Five rows in the bot's order: measured gains high to low, then picks whose gain is not modelled. Each row (min 72, padding 12/16, top-aligned):
     - the pick's icon (44);
     - the pick's name, body, quality text colour;
     - meta muted `{Slot} · {source line}`;
     - meta muted `Replaces {worn}, ilvl {n}`, or `Empty slot` when nothing is worn;
     - right-aligned: the gain in mono body `--color-strong` with `DPS` in `micro` under it; or, for an unmodelled pick, the meta line `Gain not modelled yet` under the others and no right column.
   - Footer row (52, gold body 600, chevron): `{n} more upgrades · all 17 slots`, which pushes F2.7. With five or fewer upgrades it reads `All 17 slots`.
3. **Already best in slot.** Meta muted, under the card: `Already best in slot: {slots}.`
4. **Next for {name}.** `micro` label, then a quiet card of three rows (56, tile or spec icon 36, label body 600, meta line, trailing tag or chevron):
   - `Talents`: tag `Matches` (kill) with `Your build is the {spec} {band} guide build`, or tag `Differs` (gold) with `{n} of your {m} points differ from the {band} build`. The rule is F6's. Pushes F6.
   - `Sims`: `Your {spec} sim finished: {dps} DPS` for the newest unseen finished sim, else the newest sim's headline and DPS, else `No sim yet.`. Pushes F5.1.
   - `Next raid`: tag `{State} · {Role}` (the Raids tab's viewer chip) and `{title} · {when}`. Pushes the raid page in the Raids tab. Shown only for a verified member of a guild with a planned raid.

Data: one new endpoint, §14 C1. The app shows the server's rows in the server's order and never re-ranks them.

States:
- **F2.2 Loading.** No cache: the band and the card as skeletons at real heights, five skeleton rows, and meta `Loading your upgrades…` (also the VoiceOver and TalkBack announcement).
- **F2.3 Offline.** The strip under the bar; the content as cached.
- **F2.4 No gear export.** The answer becomes the sentence `Add the addon to see which slots you can upgrade.` (17 pt, strong). It has two buttons: primary `Get the addon` (opens foreversixty.gg/setup in the browser) and secondary `{Spec} {band} list` (switches to BiS on that list). Under them: meta `The addon sends your gear when you log out. Get the addon opens foreversixty.gg/setup in your browser.` The Next rows read `Not available yet: no talent export for this character. Open the addon once to send it.` and `No sim yet.`.
- **Under level 20.** The answer reads `Best in slot starts at 20.` with secondary `{Spec} gear for 20 to 29 →`.
- **No spec known.** `Set a spec in the planner to see this.` with secondary `Open the planner` (browser).
- **No published list.** `No best in slot list published yet for this spec and level range.`
- **Every slot matches.** `Every slot matches the {band} list.` in place of the figures; the card is replaced by the Already line.
- **F2.5 Error.** `Your upgrades did not load.` in `--color-wipe` plus `Try again`; the Next rows still draw.

### F2.6 The item tooltip sheet
A tap on any item icon or item name anywhere in the app (Home, Your gear, BiS, sim gear, loot) opens this sheet. Detents medium and large; it opens at medium.
- Header: icon 56, the name in `sec` 700 quality text colour, meta `{Quality} · {Slot}`. The item level appears once, in the panel.
- The tooltip panel (`#0a0d14`, 1 pt `#3a4256`, `--radius-control`, padding 12/14), in the game's order and the site's tooltip model (`lib/items/tooltip.ts`):
  - the name in quality colour;
  - the Binds line (`Binds when picked up`, `Binds when equipped`, `Binds when used`) from ItemSparse `Bonding`, left out when 0;
  - `Unique` when ItemSparse `MaxCount` is 1;
  - the slot on the left and, on the same line at the right, the armor type (`Cloth`, `Leather`, `Mail`, `Plate`, `Shield`) from Item `SubclassID` when `ClassID` is 4. No armor type for Back, as the Classic client shows a cloak. A weapon shows its type when exact (`Two-Handed Weapon`);
  - armor;
  - the weapon line (`{min} - {max} Damage`, `Speed {s}`, `({dps} damage per second)`);
  - each stat as `+{n} {Stat}`;
  - the effect text in green `#1eff00` prefixed `Equip:` or `Use:`;
  - set lines;
  - `Requires Level {n}`;
  - `Item Level {n}` in `#e5cc80`.
- `micro` `Where it comes from`, then the source line in body strong. A quest adds ` · requires level {n}`; a drop adds its chance, `{n}% from {boss}`. A reputation reward reads `{Faction} ({Standing})` (`Warsong Outriders (Exalted)`). A vendor item with a reputation requirement (ItemSparse `MinReputation`) adds ` · requires {Standing}`. Knowledge of the Timbermaw has none in this build (its hotfix row reads 0), so it reads `Vendor: Meilosh`.
- `micro` `The next best`, then up to two alternatives (icon 36, quality name, meta `{delta} DPS vs the pick`). The worn one leads its meta with `You wear this · `. With none: `Nothing else ranked`.
- Footer meta `Best in slot for {Slot} · {Spec} {band} · {preset}`, only when the item is a pick.
- Close: swipe down, tap the scrim, or VoiceOver escape (two-finger Z) / TalkBack back.

The sheet reads the build's static files (`items/{class}.json`, `loot.json`, `sets.json`) and the BiS file already cached. It needs no request. The Binds, Unique, armor type and reputation fields need the data lane to publish `bonding`, `unique`, `armor_type` and `min_standing` in `items/{class}.json` (§14 C6). The site's tooltip gains them in the same change, so the two never differ.

### F2.7 Your gear
Pushed screen `Your gear`. Meta `What you wear, and the pick for {band}. Tap an item for its tooltip.`. Then one quiet card of 17 rows in character-frame order. Each row:
- `micro` slot label, 62 wide;
- the worn icon (36);
- the worn name (meta, quality colour) with `ilvl {n}` under it, then `→ {pick}` in the pick's quality colour when the pick differs;
- right: a kill tag `BiS` with a check when it is the pick, else a chevron and the pick's icon (36). Tapping the pick's icon opens its tooltip.

An empty slot shows a 36 dashed `--color-line` box and `Empty`.

## 5.3 F3. Best in slot

### F3.1 The list
- Bar: the selector; a trailing search button (`Search an item`, opens the system search field filtering the 17 rows by item name).
- Title `Best in slot` (`title`) with the gold tag `Yours` when the list is the current character's spec, range and faction.
- Three chips (40 tall, pill, 4.5% fill, `--color-line-card` border, chevron), each opening F3.2. They wrap onto a second line; a chip is never cut:
  - spec icon 28 with `{Spec} {Class}`;
  - the range, e.g. `60 · Phase 1 raid`;
  - the faction mark 24 with `Horde`.
- Meta: `{set_dps} DPS at best in slot for your level range · Updated {d Mon yyyy}`, the figure in mono strong.
- One quiet card, 17 rows in character-frame order. Each row (min 72):
  - icon 44;
  - name in quality colour;
  - meta `{Slot} · {source line}`;
  - meta `Next: {item}, {delta} DPS` (`same DPS` when the delta rounds to 0; `Nothing else ranked` with no alternative);
  - `Gain not modelled yet` for a pick with `effect_unmodelled`;
  - a kill tag `You wear it` when the current character wears the pick.

  An off hand under a two-handed main reads `Off hand · none: {main hand} is two-handed.`; a slot with no pick reads `{Slot} · no pick for this range.`.

  Tapping a row opens F2.6.

### F3.2 The picker sheet
Title `Choose a list`, large detent:
- `Class`: a 3 x 3 grid of crests (44) with class-coloured names; the current class carries the gold ring and 10% wash.
- `Spec`: the class's specs as 48 tall buttons with spec icon 32, name and role.
- `Level range`: pills `20 to 29`, `30 to 39`, `40 to 49`, `50 to 59`, `60`, `60 · Phase 1 raid`.
- `Faction`: two pills with marks.
- Secondary full width `Back to {name}'s list`, shown when the picks differ from the character's.

Every choice applies at once. The list under the sheet redraws live. The sheet stays open until dismissed.

States:
- First load with no cache: 17 skeleton rows.
- **F3.3 Offline:** any list fetched before shows under the strip. A list never fetched shows `This list is not on your phone yet. Connect once to save it.`
- Error: `This list did not load.` with `Try again`.
- Signed out: no selector and no `Yours`; the picker opens on first visit with Warrior preselected.

## 5.4 F4. Guild and raids (the headline)

All rules are the raid page's (guild raids spec 2026-10-26 §3 to §6, with §P.10 for the spec ask). The app adds the notification actions, the classes strip and calendar export. One roster serves the site, the Discord post and the app: a sign-up here is the same `PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}`, and lane A2 refreshes the Discord post on it.

### F4.1 Raids tab
For the current character's guild (`characters[].guild` in `/v1/me`; a character in no guild shows F4.13):
1. **Guild head.** Guild crest 40, the name in Barlow Condensed 22, and meta `{Faction} · {Region} {Ruleset} · {n} on the roster`.
2. **Next raid card** (`--surface-card`), the tab's primary region. It shows the next signable raid (`home.next_raid` rules: planned or locked, not ended).
   - The date block (`THU` / `10` / `DEC`, the viewer's local date).
   - `micro` gold `Next raid`, then the title in Barlow Condensed 22.
   - meta `Invites 7:45 pm · Pull 8:00 pm` (times in mono; `Pull` alone, as `8:00 pm to 11:00 pm`, when the guild sets no invite time), then a second meta line `Invites in 2 days 2 h` with the countdown in mono gold (`Starts in …` with no invite time).
   - The count block: the 120 pt bar, mono body strong `24 of 40 signed up` (`11 signed up` with no size, and no bar), and the meta breakdown `3 tanks · 6 healers · 14 DPS · 1 late`.
   - A 1 pt divider, then `micro` `Your sign-up`.
   - The who line: crest 36; name in class colour; spec icon 18 with `{Spec} · {Role}`.
   - The segments, then the 20 pt status line.
   - Gold link `Who is coming ›`, which pushes F4.2 scrolled to the roster.
   - The whole card, outside the segments and the link, pushes F4.2.
3. **Also planned `{n}`.** `micro`, the count in mono, counting only the rows listed under it (the next raid above is not counted), then a quiet card with the remaining planned raids. Each row (min 76):
   - date block 48 x 52;
   - title body 600;
   - meta `{when}`;
   - count in mono meta strong with a 64 pt bar;
   - one chip right, the first that applies: `Cancelled`, `Locked`, `Ended`, else the viewer chip (`In · DPS`, `Late · DPS`, `Bench`, `Out`, or gold `Not signed up`).

   A row pushes F4.2.
4. **Raid nights.** The guild's past raids (done, with a report), newest first, three rows, each pushing the raid page with its report link (`Open the report` opens the site's report in the browser in v1).

Officers see one more row above Also planned when `Can` includes `approve` and characters are waiting: tile, `{n} characters waiting for approval`, and the meta `Approve them on foreversixty.gg.` It opens the guild roster in the browser. Planning raids and approving characters stay on the web in v1 (coordinator ruling 2026-10-30); locking and moving raiders are in the app (F4.18).

### F4.2 The raid page
Pushed. Bar: back (`Raids`); trailing `Add to calendar` (glyph button) and More (`…` iOS, `⋮` Android). Then:
1. `micro` gold `Raid`, the status chip (`Planned` gold, `Locked` muted, `Cancelled` wipe, `Ended` muted, `Done` kill) and the guild name in meta.
2. The title in `title` (iOS large title; Android shows it in the bar). When the title differs from the zone, ` · {Zone}` follows in body muted.
3. The when lines:
   - body `Thursday 10 December`;
   - body `Invites 7:45 pm · Pull 8:00 pm · Ends 11:00 pm`, with the labels muted and the times in mono. With no invite time it reads `Pull 8:00 pm · Ends 11:00 pm`;
   - body `Invites in 2 days 2 h` (or `Starts in …`), with the countdown in mono gold;
   - `8:00 pm CST guild time` (muted) on its own line when the viewer's offset differs from the guild's.

   The countdown runs to the invite time when there is one, else to the pull. It ticks each minute while visible and is muted once started or cancelled.
4. The count block as F4.1.
5. **Your sign-up** card (`--surface-card`, padding 16), the page's primary region. Title `Your sign-up` (`sec`). Then one block per verified character of the viewer in this guild, the current character first:
   - The who line, with a gold `Change spec` text button right when the character has a saved spec for this guild.
   - **The segments** (§5.4.3).
   - The 20 pt status line.
   - The note field: 48 tall, `Note for the officers (optional)`, 140 characters. A counter `{n} left` appears from 120 characters. It saves on Return or on blur.
   - A waiting character instead shows the muted line `Waiting for approval. You can sign it up once an officer approves it.`, with no controls.
6. **Classes coming** (`micro`). Nine crests (30) in the game's class order, each with its In plus Late count in mono under it; a class at 0 is at 40% opacity. Tapping a crest scrolls to the roster. This answers "do we have a shaman" at a glance. It is a count of the view's own rows, not a rule.
7. **Notes** (`h2`): the API's sanitised `notes_html` as attributed text, 15/22, links in gold underlined (they open in the browser). Empty: `No notes for this raid.`
8. **Sign-ups** (`h2`) with `24/40` mono muted. Groups in order: Tanks, Healers, DPS, Late, Bench, Out, each a `micro` label with its count, then a quiet card of rows.
   - Row (min 48): spec icon 28; the name, body 600, class colour; the spec, meta muted; the note under them, meta `--color-text`, wrapping and never cut.
   - Within a group: class in the game's order, then name.
   - The viewer's own rows carry the `--row-hover` wash.
   - Empty role: `No tanks yet.`, `No healers yet.`, `No DPS yet.`. An empty Late, Bench or Out group is left out.
   - Then `Not signed up {n}`: one wrapped meta line of class-coloured names.

Live: while F4.2 is on screen and the raid starts within 24 h, it refetches every 30 s. Otherwise it refetches on appear, on pull, and on a push for this raid. A refetch never moves the scroll position or a focused control.

### F4.3 The segments and the press
- Four equal segments, `In`, `Late`, `Bench`, `Out`, each at least 48 tall, 14 pt 600, inside a `--radius-control` box with `--color-line` dividers.
- Rest: `--color-muted` text on `--color-raised`.
- Checked In or Late: kill at 15% ground, `--color-strong` text, a 2 pt kill bottom line.
- Checked Bench or Out: `--color-line-warm` ground, strong text, a 2 pt muted bottom line.
- Unsigned: none checked.
- **A press saves at once.** It is optimistic: the segment draws checked at once, the block is busy (VoiceOver value and TalkBack `stateDescription` read `Saving`), and the PUT goes out with `state` and, when known, `spec`.
- **Success:** the success haptic; the status line `Saved.` in kill for 4 s; the view is replaced by the PUT's `raid`. When `moved` is not null, the status line on the moved character's block reads `{name} moved to Bench.` (R3).
- **Failure:** the segment returns to the previous state; the error haptic; the status line in `--color-wipe` with the API's sentence (§13, `raid.*`).
- **Spec ask (F4.3 board).** In, Late or Bench pressed when the character's `raid_spec` is null and its class has two or more specs:
  - The segment draws pending: a 1.5 pt dashed gold outline inset 4, not checked.
  - The spec sheet opens: title `Which spec tonight?`, then the lead `The pick signs **{name}** up as {State} for {title}, {date}, and is kept for the next raid.`
  - Then one 64 pt row per spec (spec icon 40, name, role).
  - A pick closes the sheet and sends the PUT with `spec`. Success adds ` {Spec} is kept for {name}.` to `Saved.`.
  - Closing the sheet drops the press: the pending outline goes and nothing is saved. Meta under the rows: `Close the sheet and nothing is saved.`
  - Out never asks.
- **Change spec** opens the same sheet with `The pick changes the spec **{name}** brings to {title}, {date}, and is kept for the next raid.` and the saved spec checked. A pick sends the PUT with the current state and the new spec.
- **Locked, ended, done, cancelled:** the segments are disabled (50% opacity, not focusable as controls; VoiceOver reads `dimmed`). The status line reads the state's sentence (F4.8).

### F4.4 The reminder ask (notification permission, in context)
After the first successful In or Late on this device, when permission is not determined (iOS `.notDetermined`; Android 13+ without `POST_NOTIFICATIONS`), the sign-up card grows a quiet inset card under the status line. It contains:
- the bell tile;
- body 600 `Get a reminder before raids start?`;
- meta `One an hour before the start, and one when sign-ups open. Change them any time in You, Notifications.`;
- secondary (gold border and text) `Turn on reminders`;
- text button `Not now` (`--color-nav`).

`Turn on reminders` calls the system prompt (`UNUserNotificationCenter.requestAuthorization([.alert, .sound, .badge])`; Android `ActivityResultContracts.RequestPermission`). On grant it registers the token (lane A2). Either answer, or `Not now`, removes the card with a 200 ms collapse. The same ask comes back at most once more, after the next In at least 30 days later. Then it is never shown again; the player can still turn notifications on in F7.4.

The other in-context moment is the first `Run again` on a server sim (F5.3), with the title `Get told when this sim finishes?` and the body `One notification when the run ends. Change it any time in You, Notifications.`

The app never asks at first launch, at sign-in or on a timer. Android 9 to 12 need no permission, so no card shows there; preferences default as the catalogue says.

### F4.5 Add to calendar and the More menu
- **`Add to calendar`.** iOS: `EKEventEditViewController` (write-only access, iOS 17+). Android: `Intent.ACTION_INSERT` on `CalendarContract.Events`. The event has the title, `starts_at`, `duration_minutes`, the location `Olympus XXVII · Onyxia's Lair` and the raid's link in the notes.
- **More menu:** `Add to calendar`, `Share the raid` (the system share sheet with the raid's https link), `Open on foreversixty.gg`.

### Raid states
| State | What shows |
|---|---|
| F4.7 Unsigned | No segment checked; the viewer chip in lists reads `Not signed up` |
| F4.8 Locked | Chip `Locked`; segments disabled; status line `Sign-ups are locked. An officer can still change yours.` |
| F4.9 Cancelled | Chip `Cancelled`; countdown and bar muted; the sign-up card replaced by `This raid was cancelled.` and, when the viewer had signed, `You were {State} as {Role}. Nothing else to do.` |
| Ended | Chip `Ended`; segments disabled; `This raid has ended.` |
| Done | Chip `Done`; segments hidden; the report row `Open the report` |
| F4.10 Offline | Strip; segments disabled; status line `You are offline. Sign-ups need a connection.` No queued writes: a raid can lock while the phone is offline, and a queued answer would land wrong |
| F4.11 Save failed | Previous state restored; `That did not save. Try again.` in wipe |
| F4.12 Moved | `{name} moved to Bench.` on the other character's block |
| Race: locked mid-press | `This raid was locked while you were signing up.` |
| Race: cancelled mid-press | `This raid was cancelled while you were signing up.` |
| Rate limited | `Too many sign-up changes. Wait a minute, then try again.` |
| Plan ended | `Not saved. This guild's officer tools have ended.` |
| F4.13 No guild (tab) | Tile, `h2` `No guild yet`, body `{name} is not in a guild on Forever Sixty. When your guild's officers add it, its raids show here and you sign up in one tap.`, meta `Your guild is on Forever Sixty once one officer claims it at foreversixty.gg/guilds.` |
| F4.14 Waiting (tab) | Guild head; card with the who line, tag `Waiting`, body `You are waiting for approval. Planned raids show once an officer approves you.` |
| F4.15 No raid planned (tab) | `No raid planned yet.` and meta `When an officer plans one, it shows here and we tell you.` |
| F4.16 Error (tab) | `Planned raids did not load.` and `Try again` |
| Raid page error | `This raid did not load.` and `Try again` |
| Raid not found or not a member | `This raid is not available. It may have been removed, or it is for members of another guild.` with `Back to Raids` |
| F4.17 Offline (tab) | Strip; the next raid card with disabled segments and the offline line |

### F4.18 to F4.21 Officer tools (officers only)
Shown only when the raid view says the viewer may act (§14 C2: `viewer.can` holds `lock` and `move`). They are never inferred from a rank. A member sees none of this.
- **Officer row** (F4.18), under the count block. `micro` `Officer`, then two secondary buttons: `Lock sign-ups` (`Unlock sign-ups` once locked) with the lock glyph, and `Copy the list` (the site's `copyList`, onto the clipboard, then `Copied` for 4 s). Under them, meta: `Locking tells everyone signed up whether they are In or on the Bench.`, or when locked `Locked. Everyone signed up was told the roster; moves you make now tell the raider at once.`
- **Lock.** It needs no confirmation, because it is undone with one tap. It calls `POST /v1/guilds/{id}/raids/{raid}/lock` (`unlock`), with the busy and error states of §8. On success the toast `Sign-ups locked. {n} In and {b} on the Bench were told.` shows for 4 s (F4.21), and `raid.roster_set` goes to every signed-up raider. Members then see F4.8; the officer's own segments stay enabled.
- **Place buttons** (F4.19). Every roster row ends in a 44 tall button showing its place (`In`, `Late`, `Bench`, `Out`) with a chevron. Not signed up becomes rows with `Add`.
- **Move menu** (F4.20). iOS `Menu`, Android `DropdownMenu`. The head shows `{name} · {State} as {Role}`. The items depend on the place:

  | Place | Menu items |
  |---|---|
  | In | `Move to Bench`, then `Remove from the raid` (destructive) |
  | Bench | `Move to In`, `Remove from the raid` |
  | Late | `Move to In`, `Move to Bench` |
  | Out, Not signed up | `Add as In`, `Add as Bench` |

  A pick writes the officer `PUT` (state `in`, `bench` or `out`, with the character's spec role) and moves the row to its group with the standard list animation. Failure restores the row with `That did not save. Try again.` under it.
- **After the lock,** every move by an officer sends `raid.roster_moved` to the raider (§6.2).
- **Officer-read-only** (a guild whose plan ended): the buttons stay, disabled, with the site's sentence `Planning and running raids are part of the guild plan.` and no link.

### F4.22 Your raid night
Opened from `raid.summary`, or from a done raid's page (`Your raid night` row, for a viewer in its log). Pushed on the Raids stack, title `Your raid night`.
- Meta: `{Zone} · {Weekday} {d Month} · {Guild}`.
- Primary card: three figures in `stat` (`1` kill, `0` wipes, `1` new best) and, right, the character in class colour with `{Spec} · {Role}`.
- `Loot awarded to you`: item rows (icon 44, quality name, meta `{Slot} · {Boss} · awarded by {officer}`, meta `Replaces {worn}, ilvl {n}`); each opens F2.6. With none, the section is left out.
- `Personal bests`: one row per fight with a new best: boss, meta `Your best on this fight · {kill time} kill`, and right the DPS in mono with its percentile in the parse colour (`--color-parse-*`). With none, left out.
- `Coach notes`: Premium only. Up to three numbered notes from the coach. A free player sees no section and no prompt (F4.23).
- Ghost `Full report on foreversixty.gg`.
- States: loading skeleton; `Your raid night did not load.` with `Try again`; offline strip.

## 5.5 F5. Sims

### F5.1 Sims (pushed from Home)
Two sections:
- **`Server runs`** shows only when `entitlements.server_sims` is true or the account has server rows. Rows: spec icon 36; title (the sim's title, else headline); then a state line, one of:
  - `Starting`;
  - `Running · takes about 4 minutes` (`· taking longer than usual` when overdue), with an 8 pt progress bar (combinations done of total, or iterations);
  - `Finished {age}`, plus the DPS right in mono and an unseen gold dot (`Not opened yet`);
  - `Failed: {reason}` in wipe with `Run again`.

  Progress polls `GET /v1/sims/{id}/progress` every 5 s while the screen is visible.
- **`Your sims`** lists the newest 20 of `GET /v1/sims` (headline, `Saved from the browser · {date}` or `Ran on the server · {date}`, DPS right), with infinite scroll by page.
- Footer meta, Premium: `Start a sim on foreversixty.gg/sim. Server runs keep going after you close the app.` Free: `Start a sim on foreversixty.gg/sim and it shows here.`
- Empty: `No sims yet. Start one on foreversixty.gg/sim and it shows here.`

### F5.3 A result
- Title: the sim title, else `{Spec} sim`.
- Meta: `{Spec} {Class} · level {n} · {preset words} · Finished {age}`.
- The DPS in Barlow Condensed 52, `DPS` muted.
- Mono meta: `± {error} · {iterations} iterations · engine {version}`.
- Meta: `{set_dps} DPS at best in slot for your level range`.
- `micro` `The gear it ran`, then a 6-column grid of 40 icons (each opens F2.6).
- Then the action:
  - Premium: secondary `Run again on the server`. It shows `Starting…`, then `Started` and pushes the new run into Server runs; on failure, `Could not start it again. Try again.`
  - Free, link region (F5.4): meta `Server sims are part of Premium.` and a gold underlined `See plans` that opens the browser. One link, below the answer (premium-entry spec §2.3, §2.4).
  - Free elsewhere (F5.5): nothing.
- Ghost `Full breakdown on foreversixty.gg` with the external glyph.

Opening a result posts `POST /v1/sims/{id}/seen`.

## 5.6 F6. Talents (read-only)
Pushed from Home.
- **Status line:** tag `Matches` (kill) with `Your build is the {spec} {band} guide build.`, shown only when no cell differs. Otherwise tag `Differs` (gold) with `{n} of your {m} points differ from the {band} build.`. `{n}` is the planner's `diffCount` (`band-compare.ts`): the sum over cells of `max(your rank - guide rank, 0)`, your points spent where the guide spends fewer. `{m}` is your points spent. Under it, `{k} points unspent.` when you have spent fewer than your level allows (level minus 9).
- **Segmented:** `Guide build` and `Your build` (iOS `Picker(.segmented)`; Android `SingleChoiceSegmentedButtonRow`). It opens on `Your build` when the character has an export, else `Guide build`.
- **Tree tabs:** `{Tree} {points}` for the three trees; the spec's tree first. Each `{points}` is summed from the build shown, never stored or typed; the three always add to `{m}`.
- **The tree:** a 4-column grid, 22 pt between tiers, each talent its icon at 48, `--radius-control`.
  - The planner's `cellState`. A cell holding points (filled or maxed) has a `--color-kill` border. An empty cell that could take a point (its tier open at 5 points per tier, its prerequisite met) has a `--color-gold` border. A locked cell has a `--color-line` border with its face at 40% opacity.
  - Rank badge bottom right: mono 11 `{rank}/{max}`. Its text is `--color-kill` at max rank, `--color-text` when partly spent, gold when available, muted when locked.
  - Prerequisite lines follow the planner's `connectorPath`: straight down a column, across a row, or down then across. Each runs edge to edge, never over an icon: a 5 pt `--color-bg` halo under a 2 pt line, gold once the prerequisite holds the rank its dependent needs, `--color-line` before.
  - On `Your build`, a talent whose rank differs from the guide gets a 2 pt dashed `--color-wipe` outline, offset 3. The legend under the tree: `Dashed: your rank differs from the guide. Tap a talent for its rank text.`
- A tap opens a sheet with the talent's name, `Rank {n}/{max}`, the rank's description, and the next rank's description under `Next rank`.
- **Footer meta:** `Change talents in the game; the addon sends them on your next logout. Plan a build on foreversixty.gg/talents.`
- **Data:** the guide build string from the BiS band; the character's from `build.talents`. The difference count comes from the API (§14 C1); the app only draws both trees.
- **States:**
  - no export: `Guide build` only, with the line `Not available yet: no talent export for this character. Open the addon once to send it.`;
  - no published build: `No talent build published yet for this spec and level range.`;
  - offline: strip.

## 5.7 F7. Notifications

### F7.4 Settings (You › Notifications)
- A grouped list.
- **Header card**, only when the system permission is denied (F7.6): bell tile; `Notifications are off for Forever Sixty`; `Turn them on in Settings to get raid reminders here.` (Android: `in Android settings`); secondary `Open settings`. iOS opens `UIApplication.openNotificationSettingsURLString`; Android opens `Settings.ACTION_APP_NOTIFICATION_SETTINGS`.
- One section per stage (§6). Each row: the event's label (body strong), an optional meta line, and the channel value trailing in meta muted with a picker glyph.
  - iOS: the row is a `Picker` with `.menu` style, so a tap opens the system menu in place (F7.5).
  - Android: the row opens a `DropdownMenu` anchored to the value.
- **Values:** `Push`, `Discord`, `Push and Discord`, `Off`.
  - `Discord` and `Push and Discord` show only when `discord_link` is set. Otherwise the section footer reads `Link Discord on foreversixty.gg to get these as a Discord message too.`
  - A change saves at once (lane A2's preferences endpoint) with the selection haptic. On failure the value returns and `That did not save. Try again.` shows under the row.
- **Rows.**
  - Raids: `Sign-ups open`, `Not signed up, a day before`, `Time moved or cancelled`, `Roster set`, `Moved on or off the roster` (meta `Comes through quiet hours`), `Raid starting` (meta `1 hour before invites, set by your guild`).
  - After the raid: `Your raid night` (meta `Report, loot, bests and coach notes, in one`), `Loot, as it is awarded` (default `Off`).
  - Your characters: `First sync`, `Approved to a guild roster`, `New level range`, `Best in slot changes` (default `Off`, meta `One digest a week`).
  - Sims: `Server run finished`.
  - Officers: the three officer events.
  - Account and game: `Premium seat from your guild`, `Payment failed`, `Phase and raid dates`.
- **Who sees which sections:** the Officers section only when any guild's `Can` includes `approve`. `Server run finished` only when the entitlement is true. `Payment failed` only for a payer.
- **Quiet hours** section:
  - `Quiet hours` switch (default off);
  - `From` and `To` time rows (default 11:00 pm and 8:00 am; iOS `DatePicker(.hourAndMinute)` in a compact style, Android `TimePicker` dialog);
  - footer meta `In your phone's time zone, {zone name}. Raid starting reminders still come through.` The zone is sent with the preferences so the server applies it.
- **Footer meta:** with Discord linked, `Discord messages come as a DM from the Forever Sixty bot to {display_name}.`
- **States:** loading (skeleton rows); error `Your notification settings did not load.` with `Try again`; offline (strip, pickers disabled).

### F7.1 to F7.3 Delivery
§6 holds the copy, actions, threads and deep links.

## 5.8 F8. Premium

### F8.1 You
- **Account head:** user tile 44; BattleTag in `name` 700; meta `Signed in with Battle.net`; the gold `Premium` tag for any Premium source.
- **Card one rows:**
  - `Plan` (§5.8 below);
  - `Notifications` (`Raids, logs, sims and dates`);
  - `Characters` (`{n} on this account`; pushes a list of crests with `Remove from this account` deferred to the web);
  - `Devices` (`This {device} and {n} more`);
  - `Discord` (`Linked as {name}`, or `Not linked`; pushes a screen that opens foreversixty.gg/account#discord in the browser to link or unlink).
- **Card two:** `Help on the Discord server`, `Privacy`, `Terms` (each opens the browser, with the external glyph) and `About Forever Sixty` (pushes version, build, licences, and `Not affiliated with or endorsed by Blizzard Entertainment.`).
- Then `Sign out`, and the micro version line `Version {v} ({build}) · data {build}`.

### The Plan row by region
"Link region" means the storefront country is in `GET /v1/app/config`'s list for this platform (§14 C3). iOS reads `Storefront.current?.countryCode` (alpha-3); Android reads `BillingClient.getBillingConfigAsync` `countryCode` (alpha-2). If the config, the storefront or the billing config cannot be read, the app treats the region as not a link region (fail closed).

| Account | Link region (F8.1) | Elsewhere (F8.2) |
|---|---|---|
| Free | `Plan` · `Free` · gold `See plans` with the external glyph; the row opens `https://foreversixty.gg/premium` in the default browser (iOS `openURL`, Android `ACTION_VIEW`), not an in-app web view | `Plan` · `Free`. No chevron, not a button, no screen behind it. No other Premium word anywhere in the app |
| Premium (any source) | `Plan` · `Premium · renews {date}` (or the site's other plan lines) · `Founder` tag when founder; pushes F8.4 | Same row and screen; F8.4's billing button becomes the meta line `Billing is managed on the web.` |

### F8.4 Plan
- Card: `Premium` (`h2`) with `Founder`; the plan line (`Renews 27 Oct 2027 · $50.00 a year`, the site's `accountPlanCopy`); the founder line.
- `micro` `What your plan includes`, then kill-checked lines from the entitlements that are true:
  - `Server sims that keep running after you close the app`;
  - `Coach notes on your raid logs`;
  - `Score and parse history charts`;
  - `Sim results and logs kept for a year`.
- Link region only: secondary `Manage billing on foreversixty.gg`, which opens the browser.
- A seat reads `A seat from {guild}` and `Given by {name}. It lasts while the guild's plan does; nothing is billed to you.`.
- Footer meta: `Raid sign-ups, best in slot, talents and your character are free for everyone, in the app and on the web.`

### Anti-nag, on mobile
Everything in premium-entry spec §2 holds, plus:
- No Premium push, ever.
- No Premium mention outside a link region.
- `See plans` appears in at most two places: the Plan row and the free F5.4 result line.

# 6. Notifications: the catalogue as delivered

## 6.1 Mechanics
- **Payload.** Each push carries `event`, `link` (an `https://foreversixty.gg/app/…` universal link, §6.4), `thread` and, for raids, `raid_id`, `guild_id`, `character_key` and `ask_spec`. Text is composed on the server in the player's locale and time zone. The app never builds notification text.
- **iOS.**
  - Categories carry the actions (§6.3).
  - `thread-identifier` = `thread`.
  - `apns-collapse-id` = `{event}:{raid_id}` for raid events, so a later notification of the same kind for the same raid replaces the earlier one (plan rule "one thread per raid").
  - `interruption-level`: `time-sensitive` for `raid.starting`, `raid.roster_moved` and `raid.changed` (cancelled); `passive` for `bis.changed` and `game.milestone`; else `active`.
  - `raid.summary` and `loot.awarded` attach the first loot item's icon through a Notification Service Extension.
- **Android.**
  - Channels: `Raids` (high: `raid.posted`, `raid.signup_nudge`, `raid.changed`, `raid.roster_set`, `loot.awarded`); `Raid starting` (high, own sound: `raid.starting`, `raid.roster_moved`); `After the raid` (default: `raid.summary`); `Characters and gear` (default: `character.linked`, `guild.approved`, `band.entered`, `bis.changed` at low); `Sims` (default); `Officers` (default); `Account` (default); `Game dates` (low). A player can silence a channel in Android settings; the app's F7.4 is the per-event control.
  - `setGroup("raid:{id}")`, a notification tag per `{event}:{raid_id}`, the seal small icon tinted `--color-gold`, and `setLargeIcon` with the item icon for loot.
- **Background actions.**
  - iOS: `UNNotificationAction` with no `.foreground` option. The app wakes in the background, sends the PUT with the stored token, and posts the result as a replacement notification with the same identifier.
  - Android: a `BroadcastReceiver` with `goAsync`, then a WorkManager expedited job for the PUT. The result replaces the notification by tag.
  - Neither requires unlocking (`.authenticationRequired` off): the answer is the player's own, it is reversible, and it is the point of the feature.
- **Which character.** The server picks the character for the actions:
  - the one already signed for this raid;
  - else `main_character_key` when it is a verified member;
  - else the only verified character.

  The body names it (`Sign up Lightningbolltz?`). With two or more verified and none chosen, the notification has no In / Late / Out; its one action is `Choose a character`, which opens the raid page.
- **Unknown spec.** When the chosen character has no `raid_spec` and its class has two or more specs, the server sends `ask_spec: true`. The category then carries one foreground action, `Choose a spec`, which opens the raid page with the spec sheet up and the In press pending. Out stays a background action. On Android this is drawn as `Choose a spec` and `Out`.
- **Quiet hours.** These are server-side. A notification due inside them is held and delivered at the end. `raid.starting` and `raid.roster_moved` are the exceptions (plan rule and coordinator ruling). Held notifications for one raid collapse into the newest.
- **Amendment to the plan's catalogue** (coordinator rulings, 2026-10-30):
  - `raid.summary` replaces `report.ready`, `coach.ready`, `ranking.best` and `loot.awarded` as pushes. It is sent once per raider when the night's log is processed.
  - `loot.awarded` stays as an event, off by default; a player may turn it on as `Loot, as it is awarded`.
  - `raid.roster_set` fires when an officer locks sign-ups.
  - `raid.roster_moved` is new. It fires on every officer change to a raider's place after the lock: In to Bench, Bench to In, added, removed. It goes to that raider only.
  - `raid.starting` fires at `invite_at` minus the guild's lead (`starts_at` minus the lead with no invite time).

## 6.2 Copy, verbatim
`{when}` is the raid copy's `whenPhrase` in the player's zone: `Tonight 8 pm`, `Tomorrow 8 pm`, `Thursday 8 pm`, `Thu 17 Dec, 8 pm`. `{Title}` is the raid title. `{n of size}` is `24 of 40` or `24` when the size is unknown. On iOS the subtitle line carries the guild name for raid and officer events; on Android it goes in `setSubText`.

| Event | Title | Body | iOS actions | Android actions | Tap opens |
|---|---|---|---|---|---|
| `character.linked` | `{Name} is on Forever Sixty` | `Level {n} {Spec} {Class}, synced from the {addon\|Battle.net}. See what to upgrade first.` | none | none | Home on that character |
| `guild.approved` | `Approved in {Guild}` | `{Name} is on the roster. You can sign up for raids now.` | none | none | Raids tab, that guild |
| `band.entered` | `Level {n}: new best in slot` | `{Name} reached level {n}. First upgrade: {item} from {source}.` | `Open the list` | `Open the list` | BiS, that band |
| `bis.changed` (weekly) | `Best in slot changed this week` | `{n} picks changed in your {Spec} {band} list: {slots}.` (`1 pick changed … : Head.`) | none | none | BiS, that band |
| `raid.posted` | `{Title} · {when}` | `Sign-ups are open, {n} signed up so far. Sign up {Name}?` | `In`, `Late`, `Out` | `In`, `Late`, `Out` | The raid page |
| `raid.signup_nudge` | `{Title} · {when}` | `You have not signed up yet. {n of size} are in. Sign up {Name}?` | `In`, `Late`, `Out` | `In`, `Late`, `Out` | The raid page |
| `raid.changed` (moved) | `{Title} moved to {when}` | `You are {State} as {Role}. Change your sign-up if the new time does not work.` | `Out` | `Out` | The raid page |
| `raid.changed` (cancelled) | `{Title}, {when}, is cancelled` | `An officer cancelled it.` | none | none | The raid page |
| `raid.roster_set` (In) | `You are on the roster for {Title}` | `{when}, as {Role} on {Name}.` | none | none | The raid page |
| `raid.roster_set` (Bench) | `You are on the Bench for {Title}` | `{when}. An officer may still call you in.` | none | none | The raid page |
| `raid.roster_moved` (to In) | `You are In for {Title} {tonight\|when}` | `{Officer} moved you from the Bench. Invites {clock}, pull {clock}, as {Role} on {Name}.` | none | none | The raid page |
| `raid.roster_moved` (to Bench) | `You are on the Bench for {Title} {tonight\|when}` | `{Officer} moved you from In. An officer may still call you in before the pull at {clock}.` | none | none | The raid page |
| `raid.roster_moved` (added) | `You are In for {Title}` (or `on the Bench for`) | `{Officer} added you to the roster for {when}, as {Role} on {Name}.` | none | none | The raid page |
| `raid.roster_moved` (removed) | `You are off the roster for {Title}` | `{Officer} removed {Name} from {tonight's raid\|the raid on {when}}. Ask an officer if that is a mistake.` | none | none | The raid page |
| `raid.starting` | `Invites for {Title} in {lead}` (`1 hour`, `30 min`, `2 hours`); `{Title} starts in {lead}` with no invite time | `Invites {clock}, pull {clock}. You are {In\|Late} as {Role} on {Name}. {n of size} signed up.` (`{clock}.` alone with no invite time) | none | none | The raid page |
| `raid.summary` | `Your raid night: {Zone}` | `{k} kills, {w} wipes.` then, each only when true: ` Loot: {item}.` (`Loot: {item} and {n} more.`), ` A new best on {Boss}.` (` {n} new bests.`), ` Coach notes are ready.` (Premium) | none | none | F4.22 Your raid night |
| `loot.awarded` (off by default) | `Loot: {Item}` | `Awarded to {Name} by {Officer}.` (item icon attached) | `Open` | `Open` | F4.22 for that night |
| `sim.finished` | `Your {Spec} sim finished` | `{dps} DPS · {title}. Open the result.` | none | none | F5.3 |
| `sim.finished` (failed) | `Your {Spec} sim stopped` | `{reason}` | none | none | F5.1 |
| `roster.pending` (officers, batched) | `{n} characters waiting for approval` | `{Guild}: {Name}, {Name} and {k} more.` (`{Name} is waiting for approval.` for one) | none | none | The guild roster (web) |
| `raid.short` (officers) | `{Title} is short` | `{when}: {n of size} signed up, {t} tanks, {h} healers. A day to go.` | none | none | The raid page |
| `raid.late_change` (officers) | `{Name} is now Out of {Title}` | `Changed {age} before the start. {n of size} signed up.` | none | none | The raid page |
| `premium.seat` | `{Guild} gave you a Premium seat` | `Server sims, coach notes and history are on for your account.` | none | none | F8.4 |
| `billing.failed`, link region | `Your payment did not go through` | `Update your card on foreversixty.gg to keep Premium.` | `Open billing` | `Open billing` | foreversixty.gg/account#billing in the browser |
| `billing.failed`, elsewhere | `Your payment did not go through` | `We emailed you the details.` | none | none | F8.4 |
| `game.milestone` | `{Milestone} today` (`Phase 2 starts today`, `Onyxia's Lair opens today`; `Forever launches today` only to accounts with no synced character) | the date's own line from `dates.json` | none | none | The date's guide (web) |

Answers posted in place of an actioned notification:

| After | Title | Body |
|---|---|---|
| In | `You are In for {Title}` | `{when}, as {Role} on {Name}. {n of size} signed up.` |
| Late | `You are Late for {Title}` | `{when}, as {Role} on {Name}. Add a note for the officers on the raid page.` |
| Out | `You are Out of {Title}` | `{when}. Change it on the raid page any time before sign-ups lock.` |
| Moved (R3) | as above | the body plus ` {Other} moved to Bench.` |
| Locked | `Sign-ups for {Title} are locked` | `An officer can still change yours.` |
| Cancelled | `{Title} was cancelled` | `Nothing was saved.` |
| Failed or offline | `That did not save` | `Open the raid to sign up. {Title}, {when}.` |
| Spec unknown | `Pick a spec for {Title}` | `{Name} has no spec saved for {Guild}. Open to choose one; nothing is signed up yet.` |

The answer for In and Late carries no actions. It also replaces any earlier `raid.posted` or `raid.signup_nudge` for the raid.

## 6.3 Action identifiers
iOS categories (registered at launch, once):
- `RAID_ANSWER` (`IN`, `LATE`, `OUT`, all background).
- `RAID_ASK_SPEC` (`ASK_SPEC` foreground `Choose a spec`; `OUT` background).
- `RAID_CHOOSE` (`CHOOSE` foreground `Choose a character`).
- `RAID_MOVED` (`OUT`).
- `OPEN_LIST` (`OPEN` foreground `Open the list`).
- `LOOT` (`OPEN` foreground `Open`).
- `BILLING` (`OPEN` foreground `Open billing`).

Android uses the same identifiers as the intent actions.

## 6.4 Deep links
`https://foreversixty.gg/app/…`, served as iOS universal links (`apple-app-site-association`, path `/app/*`) and Android App Links (`assetlinks.json`, `autoVerify`). The site answers the same path with a redirect to its own page, so a link opened without the app still works (§14 C4).

| Path | Opens |
|---|---|
| `/app/character/{region}/{ruleset}/{name}` | Home on that character (it becomes the current character) |
| `/app/raid/{guild_id}/{raid_id}` | Raids tab › raid page; `?spec=ask` opens the spec sheet with In pending; `?state=in` presses In on open (used by `Choose a character` after the pick) |
| `/app/raid/{guild_id}/{raid_id}/night` | Raids tab › raid page › Your raid night |
| `/app/guild/{guild_id}` | Raids tab for that guild |
| `/app/bis/{spec}/{band}` | BiS on that list (`band` is `20` to `60` or `p1`) |
| `/app/sim/{id}` | Home › Sims › the result |
| `/app/you/notifications`, `/app/you/plan` | You › that screen |
| `/app/auth?code=…` | The sign-in return (F1.2), never opened by a notification |

A link to a character, guild or raid the account cannot see opens the matching not-available state, never a blank screen.

# 7. Platform behaviour

| | iOS | Android |
|---|---|---|
| Back | The bar's back button and the edge swipe; a sheet swipes down | The system back gesture with predictive back; a sheet, a menu, then the stack |
| Sheets vs dialogs | Sheets for choosing and reading (switcher, tooltip, spec ask, BiS picker, talent rank). Alerts only for destructive confirmation (Revoke, Sign out) | Bottom sheets for the same; `AlertDialog` for the same confirmations |
| Menus | `Menu` (More, channel picker) | `DropdownMenu` |
| Haptics | `.sensoryFeedback(.selection)` on a segment press, picker change and character pick; `.success` on a saved sign-up; `.error` on a refused one | `HapticFeedbackConstants.SEGMENT_TICK` (API 34+, else `CLOCK_TICK`) on a press; `CONFIRM` (API 30+) on save; `REJECT` (API 30+) on refusal; none below API 30 |
| Pull to refresh | `.refreshable` on Home, Raids, raid page, BiS, Sims | `PullToRefreshBox` on the same |
| Text size | Dynamic Type to AX5; the rules of §4.2 and §9 | Font scale to 2.0, non-linear (Android 14) |
| Reduce Motion | Countdown ticks without animation; sheets cross-fade | `ANIMATOR_DURATION_SCALE` 0 respected |
| Orientation | Portrait on iPhone. iPad runs the iPhone app in v1 | Portrait on phones. Large screens follow `android:resizeableActivity` with the phone layout at max 600 dp centred |
| Haptic and sound for pushes | System | Channel settings |

## 7.1 VoiceOver and TalkBack

Every icon that sits beside its own name is hidden from VoiceOver and TalkBack (`accessibilityHidden(true)`; `contentDescription = null`). Composite rows are one element with a combined label, so a swipe moves row by row.

| Element | Label | Value / state | Hint / action |
|---|---|---|---|
| Selector | `Character, Lightningbolltz, Elemental Shaman, level 60` | | `Opens your characters` / custom action `Switch character` |
| Answer figures | `Upgrades for your level range: 8 upgrades, plus 18.3 DPS. Best in slot for 60, together. 2 gains not modelled yet.` | | |
| Upgrade row | `Waist. Knowledge of the Timbermaw, Epic, item level 65. Vendor: Meilosh. Replaces Belt of the Archmage, item level 62. Plus 5.3 DPS.` | | `Shows the item` |
| Unmodelled row | the same, ending `Gain not modelled yet.` | | |
| BiS row | `Head. Blue Dragonscale Helm, Epic, item level 61. Crafted: Leatherworking. Next: Warlord's Mail Helm, minus 1.3 DPS. You wear it.` | | `Shows the item` |
| Item icon alone (sim grid, gear) | `{Item}, {Quality} {slot} item` (the bot's `TileAlt`) | | `Shows the item` |
| Segments | group `Lightningbolltz's sign-up`; each `In`, `Late`, `Bench`, `Out` | iOS `.isSelected` trait on the checked one, `.notEnabled` when disabled; Android `Role.RadioButton`, `selected`, `stateDescription` `Saving` while busy and `Pending, choose a spec` while pending | |
| Status line | announced politely when it changes (`AccessibilityNotification.Announcement`; `LiveRegionMode.Polite`) | | |
| Countdown | `Starts in 2 days 3 hours` (whole words, never `h`) | not announced on tick | |
| Count | `24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late` | | |
| Classes strip | one element: `Classes coming: Warrior 4, Hunter 2, Rogue 2, Priest 4, Shaman 4, Mage 2, Warlock 3, Druid 3, no Paladins` | | |
| Roster row | `Kraggor, Protection Warrior` and the note | | |
| Date block | hidden; the title row reads `Thursday 10 December` | | |
| Offline strip | `Offline. Showing what you had at 9:41 pm.` announced once when it appears | | |
| Channel row | `Sign-ups open` | `Push` | `Changes how you hear about this` |
| Talent | `Convection, rank 5 of 5` and, when different, `guide rank 3` | | `Shows the talent` |
| Tab bar | the system's labels | badge: Raids shows a dot when a planned raid is unanswered; read `Not signed up for 2 raids` | |

Touch targets are at least 44 x 44 pt (iOS) and 48 x 48 dp (Android). Rows meet this by height. Text buttons meet it by padding. The note field and segments are 48.

Focus order follows the visual order. A sheet traps focus until closed and returns focus to the control that opened it. Hardware keyboards: Tab moves between controls, Space or Return presses, arrow keys move within the segments (iOS through the button group; Android `selectableGroup`). Escape closes a sheet.

# 8. Every interactive element, every state

| Element | Default | Pressed (touch down) | Focused (keyboard, VoiceOver, TalkBack) | Disabled | Loading | Error |
|---|---|---|---|---|---|---|
| Primary button | gold gradient, `#0b0e14` label | 8% darker fill, 0.98 scale over 120 ms (iOS); Material ripple in `#0b0e14` 12% (Android) | 2 pt gold ring offset 2 (iOS hardware focus; Android `focusable` indication) | 50% opacity | label replaced by a 20 pt spinner, same width | the region's error line; button stays enabled |
| Secondary button | 4.5% white, gold 17% border | `--row-hover` x2 (10% gold) | gold ring | 50% | spinner | as above |
| List row | `--color-raised` | `rgba(229,185,85,.10)` (the `pressed` class in `kit.py`) | gold ring inset 2 | n/a | skeleton | n/a |
| Segment | rest look | the segment's checked ground at 60% | gold 2 pt ring inset 3 | 50%, no hit | busy: checked look, `Saving` state | restored, wipe line |
| Selector | glass capsule (iOS) / 4% fill with gold border (Android) | 10% gold wash | gold ring | never | n/a | n/a |
| Selector open | full gold border, gold glow | | | | | |
| Chip (BiS) | pill, card border | 10% gold | ring | n/a | n/a | n/a |
| Switch | system with gold tint | system | system | system 50% | n/a | value returns, line under |
| Channel value | muted with glyph | menu opens | ring | 50% offline | n/a | value returns, line under |
| Note field | 35% black, card border, placeholder muted | caret, gold border | gold border | 50% | `Saving…` in the counter slot | wipe border, `That note did not save. Try again.` |
| Item icon | quality border | 0.96 scale | ring | n/a | tile drawn at once with quality border and `#131824`; image fades in over 120 ms | the tile stays, no broken-image glyph; the name carries the item |
| Tab | system | system | system | n/a | n/a | n/a |
| Notification action | system | system | system | n/a | replaced on result | replaced with the failure answer |

# 9. Hierarchy

| Region | The one primary | Secondary, visibly |
|---|---|---|
| Home hero | The two figures (`stat`, strong) | Label (micro muted), note and sync (meta muted) |
| Home upgrades card | Each row's item name in quality colour | Slot, source and replaces (meta muted); the gain (mono strong, smaller than the name's weight) |
| Raids tab | The next raid card (the only `--surface-card` on the tab) | Planned list (quiet cards), raid nights |
| Raid page | Your sign-up card (the only `--surface-card`) | The head (text on ground), classes, notes, roster (quiet cards) |
| Sign-up card | The segments | Who line, status, note, Change spec (gold text, small) |
| BiS | The rows' names | Chips, set DPS line |
| Sim result | The DPS figure at 52 | Error line, gear grid, actions |
| You | The account head | Rows; Sign out muted |

There is one gold-filled button per screen at most. Gold text marks links and the countdown only. Kill green marks a good state: checked In or Late, `Matches`, `You wear it`, `Saved.`. Whitespace: 22 before a section label, 14 inside cards, 16 margins.

# 10. Phone behaviour and what collapses

| Change | iPhone SE 3 (375 x 667) | iPhone 17, Pixel 9 | Largest text (XL boards) |
|---|---|---|---|
| Home rows above the fold | 3 | 5 | 1 |
| Upgrade gain | right column | right column | moves under the name as a mono meta line `+5.3 DPS` (XL.1) |
| Segments | 4 across | 4 across | 2 x 2 grid from iOS AX1 and Android 1.5x (XL.2) |
| Selector | full name, description | full name, description | full name capped at 28 pt (iOS) / 22 sp (Android), no description |
| Classes strip | 9 across at 26 | 9 across at 30 | 5 then 4, crests at 36 |
| Date block | 48 x 52 | 52 x 56 | grows with its text; the title wraps under it |
| Settings value | trailing | trailing | trailing; the label wraps (XL.4) |
| Tab labels | system | system | system (iOS Large Content Viewer on long press) |

These never collapse:
- the full character name;
- every item's icon and quality colour;
- the In / Late / Bench / Out choice;
- the count and the countdown;
- `Gain not modelled yet` where it applies;
- a runner-up line on BiS (tenet 4).

Nothing is ever cut with an ellipsis except a raid title in a planned row, which clamps at two lines and shows in full on the raid page.

# 11. Performance and polish limits

These are measured on device in Wave 2 and are part of done (tenet 13).

| Budget | iOS (iPhone 12) | Android (Pixel 6a) |
|---|---|---|
| Cold start to first frame | ≤ 400 ms | ≤ 500 ms |
| Cold start to the cached answer drawn | ≤ 1.0 s | ≤ 1.2 s |
| Answer refreshed over 4G | ≤ 2.5 s | ≤ 2.5 s |
| Sign-up press to checked look | ≤ 1 frame | ≤ 1 frame |
| Sign-up round trip to `Saved.` (p75) | ≤ 800 ms | ≤ 800 ms |
| Notification action to answer posted (p75) | ≤ 4 s | ≤ 4 s |
| Scroll, raid page with 40 rows | no hitches over 16 ms at p95 (Instruments Hitches) | janky frames < 1% (Macrobenchmark `FrameTimingMetric`) |
| Download size | ≤ 30 MB | AAB download ≤ 15 MB |
| Memory, raid page | ≤ 120 MB | ≤ 150 MB |

- **No layout shift.** Every region reserves its height on first draw: skeletons at real sizes, item tiles at real sizes, the 20 pt status line always present.
- **No missing-icon flash.** The quality-bordered tile draws at once and the icon fades in. Icons are cached on disk per build (`URLCache` with a 200 MB disk cap; Coil with a 200 MB disk cache). The upgrades and BiS icons for the current character are prefetched after launch.
- **No font swap.** The fonts are bundled.
- **No spinner over cached content.**

# 12. Interaction details that make sign-up the best in WoW

1. **One tap from the lock screen** (F7.2), with the answer posted back where the question was.
2. **One tap on the Raids tab**, without opening the raid.
3. **The spec is asked once per guild membership,** then never again (§P.10).
4. **R3 is shown, not hidden:** a second character going In says which one moved to Bench.
5. **The roster answers "who is coming" in two ways:** by role, as Raid-Helper does, and by class at a glance.
6. **Add to calendar** puts the raid where a player's evening is planned.
7. **Live updates in the last 24 h,** so the count is true on raid night.
8. **The guild's Discord post updates on an app sign-up,** so nobody has to choose a door.

# 13. Copy keys

Keys live in one string catalogue per app: `Localizable.xcstrings` on iOS and `strings.xml` on Android. Each key below has the same name on both platforms, with dots becoming underscores on Android. `{x}` are format arguments, named in the catalogue comment. Strings marked "site" are byte-equal to the named web module; a test in each app reads this table (as `raid-copy.test.ts` reads the raid spec).

| Key | Text | Source |
|---|---|---|
| `tab.home` | `Home` | |
| `tab.raids` | `Raids` | |
| `tab.bis` | `BiS` | |
| `tab.you` | `You` | |
| `offline.strip` | `Offline. Showing what you had at {time}.` | |
| `common.tryAgain` | `Try again` | |
| `common.saved` | `Saved.` | site `raidCopy.stripStatus.saved` |
| `common.failed` | `That did not save. Try again.` | site |
| `signin.line` | `Play your class better.` | site |
| `signin.body` | `Sign in and your characters arrive with their gear, talents and guild.` | site |
| `signin.button` | `Sign in with Battle.net` | site `homePanelCopy.signInButton` |
| `signin.pickClass` | `Pick your class` | |
| `signin.pickClassNote` | `Pick your class opens best in slot, talents and sims for any spec, no account needed.` | |
| `signin.importTitle` | `Bringing in your characters` | |
| `signin.importBody` | `Reading Battle.net. This takes a few seconds.` | |
| `signin.error.bnet` | `Battle.net could not confirm that sign-in. Try again.` | API |
| `signin.error.expired` | `This sign-in has expired. Start again.` | API |
| `signin.error.offline` | `No connection. Signing in needs the internet.` | |
| `signin.noChars.title` | `No characters yet` | |
| `signin.noChars.body` | `Battle.net lists a character once it has logged in on Forever. Log in once, then pull to refresh.` | |
| `signin.getAddon` | `Get the addon` | site `landingCopy.noSyncLink` |
| `signin.signedOut` | `You were signed out. Sign in again.` | |
| `raids.signedOut.title` | `Sign in to sign up for raids` | |
| `raids.signedOut.body` | `Your guild's raids show here once you sign in with Battle.net.` | |
| `switcher.title` | `Your characters` | |
| `switcher.who` | `{battletag} · Signed in with Battle.net` | |
| `switcher.noGuild` | `No guild` | |
| `switcher.footer` | `A character shows here after it logs in on Forever, or after the addon syncs it.` | |
| `selector.a11y` | `Character, {name}, {spec} {class}, level {n}` | |
| `home.answer.label` | `Upgrades for your level range` | site `headerAnswerLabel` |
| `home.answer.upgrades` | `upgrades` / `upgrade` | site |
| `home.answer.note` | `Best in slot for {band}, together` | site `headerNote` |
| `home.answer.notModelled` | ` · {n} gains not modelled yet` / ` · 1 gain not modelled yet` | |
| `home.answer.synced` | `Synced {age} from the addon` / `Synced {age} from Battle.net` | |
| `home.answer.noGear` | `Add the addon to see which slots you can upgrade.` | site `headerNoGear` |
| `home.answer.noGearNote` | `The addon sends your gear when you log out. Get the addon opens foreversixty.gg/setup in your browser.` | |
| `home.answer.listButton` | `{Spec} {band} list` | |
| `home.answer.under20` | `Best in slot starts at 20.` | site `landingCopy.under20` |
| `home.answer.pickSpec` | `Set a spec in the planner to see this.` | site |
| `home.answer.noList` | `No best in slot list published yet for this spec and level range.` | site |
| `home.answer.allMatch` | `Every slot matches the {band} list.` | site |
| `home.answer.failed` | `Your upgrades did not load.` | site `headerFailed` |
| `home.answer.loading` | `Loading your upgrades…` | |
| `home.upgrades.title` | `Your upgrades` | site |
| `home.upgrades.replaces` | `Replaces {item}, ilvl {n}` | bot `Replaces` |
| `home.upgrades.empty` | `Empty slot` | |
| `home.upgrades.notModelled` | `Gain not modelled yet` | bot `NotModelled` |
| `home.upgrades.more` | `{n} more upgrades · all 17 slots` / `1 more upgrade · all 17 slots` / `All 17 slots` | |
| `home.upgrades.already` | `Already best in slot: {slots}.` | site `alreadyBestInSlotPrefix` |
| `home.next.title` | `Next for {name}` | |
| `home.next.talents` | `Talents` | site |
| `home.next.matches` | `Matches` | site planner `matches` |
| `home.next.differs` | `Differs` | |
| `home.next.talentsMatch` | `Your build is the {spec} {band} guide build` | |
| `home.next.talentsDiffer` | `{n} of your {m} points differ from the {band} build` | site planner `differs` with its band |
| `home.next.talentsNone` | `Not available yet: no talent export for this character. Open the addon once to send it.` | site |
| `home.next.sims` | `Sims` | |
| `home.next.simFinished` | `Your {spec} sim finished: {dps} DPS` | site `bandLineParts` |
| `home.next.noSim` | `No sim yet.` | site |
| `home.next.raid` | `Next raid` | |
| `gear.title` | `Your gear` | |
| `gear.intro` | `What you wear, and the pick for {band}. Tap an item for its tooltip.` | |
| `gear.bis` | `BiS` | |
| `gear.emptySlot` | `Empty` | |
| `tooltip.bind.pickup` / `.equip` / `.use` | `Binds when picked up` / `Binds when equipped` / `Binds when used` | client |
| `tooltip.unique` | `Unique` | client |
| `tooltip.armorType.*` | `Cloth`, `Leather`, `Mail`, `Plate`, `Shield` | client |
| `tooltip.repSource` | `{faction} ({standing})` | site `repSourceLabel` |
| `tooltip.requiresStanding` | ` · requires {standing}` | |
| `tooltip.requires` | `Requires Level {n}` | client |
| `tooltip.itemLevel` | `Item Level {n}` | client |
| `tooltip.armor` | `{n} Armor` | client |
| `tooltip.equip` | `Equip: {text}` | client |
| `tooltip.use` | `Use: {text}` | client |
| `tooltip.source` | `Where it comes from` | |
| `tooltip.nextBest` | `The next best` | |
| `tooltip.youWear` | `You wear this · ` | |
| `tooltip.delta` | `{delta} DPS vs the pick` | bot `BisAltDelta` |
| `tooltip.nothingElse` | `Nothing else ranked` | bot `BisOnly` |
| `tooltip.pickOf` | `Best in slot for {slot} · {spec} {band} · {preset}` | |
| `tooltip.a11yTile` | `{item}, {quality} {slot} item` | bot `TileAlt` |
| `bis.title` | `Best in slot` | site |
| `bis.yours` | `Yours` | site `landingCopy.yours` |
| `bis.setDps` | `{dps} DPS at best in slot for your level range · Updated {date}` | site `simulatorBandLine` |
| `bis.next` | `Next: {item}, {delta} DPS` | |
| `bis.nextSame` | `Next: {item}, same DPS` | |
| `bis.youWear` | `You wear it` | |
| `bis.offHandTwoHand` | `Off hand · none: {item} is two-handed.` | bot |
| `bis.noPick` | `{slot} · no pick for this range.` | bot |
| `bis.search` | `Search an item` | |
| `bis.picker.title` | `Choose a list` | |
| `bis.picker.class` / `.spec` / `.range` / `.faction` | `Class` / `Spec` / `Level range` / `Faction` | |
| `bis.picker.back` | `Back to {name}'s list` | |
| `bis.range.p1` | `60 · Phase 1 raid` | |
| `bis.notCached` | `This list is not on your phone yet. Connect once to save it.` | |
| `bis.failed` | `This list did not load.` | |
| `raids.guildLine` | `{faction} · {region} {ruleset} · {n} on the roster` | |
| `raids.next` | `Next raid` | |
| `raids.whoIsComing` | `Who is coming` | |
| `raids.planned` | `Also planned` | |
| `raids.nights` | `Raid nights` | site |
| `raids.pending` | `{n} characters waiting for approval` | |
| `raids.pendingLine` | `Approve them on foreversixty.gg.` | |
| `raids.noGuild.title` | `No guild yet` | |
| `raids.noGuild.body` | `{name} is not in a guild on Forever Sixty. When your guild's officers add it, its raids show here and you sign up in one tap.` | |
| `raids.noGuild.note` | `Your guild is on Forever Sixty once one officer claims it at foreversixty.gg/guilds.` | |
| `raids.waiting` | `You are waiting for approval. Planned raids show once an officer approves you.` | site `notVisible.unverified` |
| `raids.empty` | `No raid planned yet.` | site `bandNone.member` |
| `raids.emptyNote` | `When an officer plans one, it shows here and we tell you.` | |
| `raids.failed` | `Planned raids did not load.` | site |
| `raid.eyebrow` | `Raid` | site |
| `raid.status.planned` / `.locked` / `.done` / `.cancelled` / `.ended` | `Planned` / `Locked` / `Done` / `Cancelled` / `Ended` | site |
| `raid.chip.notSigned` | `Not signed up` | site |
| `raid.timeTo` | `{start} to {end}` | site |
| `raid.invites` / `.pull` / `.ends` | `Invites {clock}` / `Pull {clock}` / `Ends {clock}` | |
| `raid.countdown.invites` / `.starts` | `Invites {countdown}` / `Starts {countdown}` | |
| `officer.label` | `Officer` | |
| `officer.lock` / `.unlock` | `Lock sign-ups` / `Unlock sign-ups` | site |
| `officer.copy` / `.copied` | `Copy the list` / `Copied` | |
| `officer.lockNote` | `Locking tells everyone signed up whether they are In or on the Bench.` | |
| `officer.lockedNote` | `Locked. Everyone signed up was told the roster; moves you make now tell the raider at once.` | |
| `officer.lockedToast` | `Sign-ups locked. {n} In and {b} on the Bench were told.` | |
| `officer.memberLocked` | `Sign-ups are locked. Members cannot change theirs; officers can.` | |
| `officer.menu.head` | `{name} · {state} as {role}` | |
| `officer.move.in` / `.bench` / `.remove` / `.addIn` / `.addBench` | `Move to In` / `Move to Bench` / `Remove from the raid` / `Add as In` / `Add as Bench` | |
| `officer.add` | `Add` | |
| `officer.gate` | `Planning and running raids are part of the guild plan.` | API `refusal.RaidPlanRequired` |
| `night.title` | `Your raid night` | |
| `night.meta` | `{zone} · {weekday} {date} · {guild}` | |
| `night.kill` / `.wipes` / `.best` | `kill` / `kills`, `wipe` / `wipes`, `new best` / `new bests` | |
| `night.loot` | `Loot awarded to you` | |
| `night.lootLine` | `{slot} · {boss} · awarded by {officer}` | |
| `night.bests` | `Personal bests` | |
| `night.bestLine` | `Your best on this fight · {time} kill` | |
| `night.percentile` | `{n}th percentile` (with the English ordinal) | |
| `night.coach` | `Coach notes` | |
| `night.report` | `Full report on foreversixty.gg` | |
| `night.failed` | `Your raid night did not load.` | |
| `raid.guildTime` | `{clock} guild time` | site |
| `raid.count` | `{n} of {size} signed up` / `{n} signed up` | site `countLine` |
| `raid.breakdown` | `{t} tanks · {h} healers · {d} DPS · {l} late` | site `roleBreakdown` |
| `raid.yourSignup` | `Your sign-up` | site |
| `raid.state.in` / `.late` / `.bench` / `.out` | `In` / `Late` / `Bench` / `Out` | site |
| `raid.role.tank` / `.healer` / `.dps` | `Tank` / `Healer` / `DPS` | site |
| `raid.note.placeholder` | `Note for the officers (optional)` | site |
| `raid.note.left` | `{n} left` | site |
| `raid.note.failed` | `That note did not save. Try again.` | |
| `raid.moved` | `{name} moved to Bench.` | site |
| `raid.locked` | `Sign-ups are locked. An officer can still change yours.` | site |
| `raid.cancelled` | `This raid was cancelled.` | site |
| `raid.cancelledYou` | `You were {state} as {role}. Nothing else to do.` | |
| `raid.ended` | `This raid has ended.` | bot `Ended` |
| `raid.offline` | `You are offline. Sign-ups need a connection.` | |
| `raid.race.locked` | `This raid was locked while you were signing up.` | site |
| `raid.race.cancelled` | `This raid was cancelled while you were signing up.` | site |
| `raid.race.toolsEnded` | `Not saved. This guild's officer tools have ended.` | site |
| `raid.rateLimited` | `Too many sign-up changes. Wait a minute, then try again.` | bot `RateLimited` |
| `raid.waitingCharacter` | `Waiting for approval. You can sign it up once an officer approves it.` | site |
| `raid.classes` | `Classes coming` | |
| `raid.notes` | `Notes` | site |
| `raid.notesEmpty` | `No notes for this raid.` | site |
| `raid.signups` | `Sign-ups` | site |
| `raid.group.tanks` / `.healers` / `.dps` / `.late` / `.bench` / `.out` / `.notSigned` | `Tanks` / `Healers` / `DPS` / `Late` / `Bench` / `Out` / `Not signed up` | site |
| `raid.emptyRole.tank` / `.healer` / `.dps` | `No tanks yet.` / `No healers yet.` / `No DPS yet.` | site |
| `raid.failed` | `This raid did not load.` | site |
| `raid.notAvailable` | `This raid is not available. It may have been removed, or it is for members of another guild.` | |
| `raid.back` | `Back to Raids` | |
| `raid.report` | `Open the report` | |
| `raid.menu.calendar` | `Add to calendar` | |
| `raid.menu.share` | `Share the raid` | |
| `raid.menu.web` | `Open on foreversixty.gg` | |
| `raid.countdown.*` | `in {span}`, `starting now`, `started {span} ago`, `ended {span} ago` | site `countdown` |
| `raid.a11y.countdown` | `Starts in {span words}` | |
| `spec.title` | `Which spec tonight?` | site |
| `spec.lead` | `The pick signs {name} up as {state} for {title}, {date}, and is kept for the next raid.` | site |
| `spec.lead.change` | `The pick changes the spec {name} brings to {title}, {date}, and is kept for the next raid.` | site |
| `spec.change` | `Change spec` | site |
| `spec.kept` | `{spec} is kept for {name}.` | site |
| `spec.closeNote` | `Close the sheet and nothing is saved.` | |
| `ask.raid.title` | `Get a reminder before raids start?` | |
| `ask.raid.body` | `One an hour before the start, and one when sign-ups open. Change them any time in You, Notifications.` | |
| `ask.sim.title` | `Get told when this sim finishes?` | |
| `ask.sim.body` | `One notification when the run ends. Change it any time in You, Notifications.` | |
| `ask.turnOn` | `Turn on reminders` (raid) / `Turn on` (sim) | |
| `ask.notNow` | `Not now` | |
| `sims.title` | `Sims` | |
| `sims.server` | `Server runs` | site |
| `sims.mine` | `Your sims` | site |
| `sims.queued` / `.running` | `Starting` / `Running` | site |
| `sims.takes` | ` · takes {estimate}` / ` · taking longer than usual` | site |
| `sims.failed` | `Failed: {reason}` / `Failed` | site |
| `sims.finished` | `Finished {age}` | site |
| `sims.unread` | `Not opened yet` | site |
| `sims.savedBrowser` | `Saved from the browser · {date}` | |
| `sims.ranServer` | `Ran on the server · {date}` | |
| `sims.footer.premium` | `Start a sim on foreversixty.gg/sim. Server runs keep going after you close the app.` | |
| `sims.footer.free` | `Start a sim on foreversixty.gg/sim and it shows here.` | |
| `sims.empty` | `No sims yet. Start one on foreversixty.gg/sim and it shows here.` | |
| `sims.runAgain` | `Run again on the server` (result) / `Run again` (row) | site |
| `sims.runAgainPending` / `Done` / `Failed` | `Starting…` / `Started` / `Could not start it again. Try again.` | site |
| `sims.premiumLine` | `Server sims are part of Premium.` | site `planLapsed` |
| `sims.gear` | `The gear it ran` | |
| `sims.breakdown` | `Full breakdown on foreversixty.gg` | |
| `sims.error` | `± {error} · {n} iterations · engine {version}` | |
| `talents.title` | `Talents` | |
| `talents.matches` / `.differs` | `Matches` / `Differs` | site planner |
| `talents.match` | `Your build is the {spec} {band} guide build.` | |
| `talents.differ` | `{n} of your {m} points differ from the {band} build.` | site planner `differs` |
| `talents.unspent` | `{n} points unspent.` / `1 point unspent.` | |
| `talents.guide` / `.yours` | `Guide build` / `Your build` | |
| `talents.legend` | `Dashed: your rank differs from the guide. Tap a talent for both ranks and its text.` | |
| `talents.rank` | `Rank {n}/{max}` | |
| `talents.nextRank` | `Next rank` | |
| `talents.footer` | `Change talents in the game; the addon sends them on your next logout.` | |
| `talents.noList` | `No talent build published yet for this spec and level range.` | site |
| `notif.title` | `Notifications` | |
| `notif.section.raids` / `.after` / `.characters` / `.sims` / `.officers` / `.account` | `Raids` / `After the raid` / `Your characters` / `Sims` / `Officers` / `Account and game` | |
| `notif.event.raidPosted` | `Sign-ups open` | |
| `notif.event.signupNudge` | `Not signed up, a day before` | |
| `notif.event.raidChanged` | `Time moved or cancelled` | |
| `notif.event.rosterSet` | `Roster set` | |
| `notif.event.rosterMoved` | `Moved on or off the roster` | |
| `notif.event.rosterMovedNote` | `Comes through quiet hours` | |
| `notif.event.raidStarting` | `Raid starting` | |
| `notif.event.raidStartingNote` | `{lead} before invites, set by your guild` | |
| `notif.event.raidSummary` | `Your raid night` | |
| `notif.event.raidSummaryNote` | `Report, loot, bests and coach notes, in one` | |
| `notif.event.lootAwarded` | `Loot, as it is awarded` | |
| `notif.event.characterLinked` | `First sync` | |
| `notif.event.guildApproved` | `Approved to a guild roster` | |
| `notif.event.bandEntered` | `New level range` | |
| `notif.event.bisChanged` | `Best in slot changes` | |
| `notif.event.bisChangedNote` | `One digest a week` | |
| `notif.event.simFinished` | `Server run finished` | |
| `notif.event.rosterPending` | `Characters waiting for approval` | |
| `notif.event.raidShort` | `A raid is short a day before` | |
| `notif.event.lateChange` | `Someone drops to Out inside a day` | |
| `notif.event.premiumSeat` | `Premium seat from your guild` | |
| `notif.event.billingFailed` | `Payment failed` | |
| `notif.event.milestone` | `Phase and raid dates` | |
| `notif.channel.push` / `.discord` / `.both` / `.off` | `Push` / `Discord` / `Push and Discord` / `Off` | |
| `notif.linkDiscord` | `Link Discord on foreversixty.gg to get these as a Discord message too.` | |
| `notif.discordFooter` | `Discord messages come as a DM from the Forever Sixty bot to {name}.` | |
| `notif.quiet` / `.from` / `.to` | `Quiet hours` / `From` / `To` | |
| `notif.quietNote` | `In your phone's time zone, {zone}. Raid starting reminders still come through.` | |
| `notif.denied.title` | `Notifications are off for Forever Sixty` | |
| `notif.denied.body` | `Turn them on in Settings to get raid reminders here.` / `Turn them on in Android settings to get raid reminders here.` | |
| `notif.denied.button` | `Open settings` | |
| `notif.failed` | `Your notification settings did not load.` | |
| `push.*` | every title and body of §6.2, keyed `push.{event}.title` and `push.{event}.body`, plus `push.answer.{in,late,out,locked,cancelled,failed,askSpec}.{title,body}` | the server's catalogue; the apps hold only the action labels below |
| `push.action.in` / `.late` / `.out` | `In` / `Late` / `Out` | |
| `push.action.askSpec` | `Choose a spec` | |
| `push.action.choose` | `Choose a character` | |
| `push.action.openList` | `Open the list` | |
| `push.action.openLoot` | `Open` | |
| `push.action.openBilling` | `Open billing` | |
| `you.title` | `You` | |
| `you.signedInWith` | `Signed in with Battle.net` | site |
| `you.plan` | `Plan` | site |
| `you.plan.free` | `Free` | site |
| `you.seePlans` | `See plans` | site |
| `you.notifications.sub` | `Raids, logs, sims and dates` | |
| `you.characters` / `.sub` | `Characters` / `{n} on this account` | |
| `you.devices` / `.sub` | `Devices` / `This {device} and {n} more` | |
| `you.discord` / `.linked` / `.unlinked` | `Discord` / `Linked as {name}` / `Not linked` | |
| `you.help` / `.privacy` / `.terms` / `.about` | `Help on the Discord server` / `Privacy` / `Terms` / `About Forever Sixty` | |
| `you.signOut` | `Sign out` | site |
| `you.signOut.confirm` | `Sign out of Forever Sixty on this {device}?` | |
| `you.version` | `Version {v} ({build}) · data {data}` | |
| `about.notAffiliated` | `Not affiliated with or endorsed by Blizzard Entertainment.` | site footer |
| `plan.includes` | `What your plan includes` | |
| `plan.feature.serverSims` / `.coach` / `.history` / `.retention` | `Server sims that keep running after you close the app` / `Coach notes on your raid logs` / `Score and parse history charts` / `Sim results and logs kept for a year` | |
| `plan.manage` | `Manage billing on foreversixty.gg` | |
| `plan.managedOnWeb` | `Billing is managed on the web.` | |
| `plan.freeForAll` | `Raid sign-ups, best in slot, talents and your character are free for everyone, in the app and on the web.` | |
| `plan.*` lines | `Premium · renews {date}`, `Premium · ends {date}`, `A seat from {guild}`, `Given by {name}. It lasts while the guild's plan does; nothing is billed to you.`, `Founder`, `Founder price, kept while your plan stays active.` | site `accountPlanCopy` |
| `devices.title` | `Devices` | |
| `devices.this` | `This iPhone` / `This phone` | |
| `devices.lastUsed` | `Last used {age}` | |
| `devices.signedIn` | `Signed in {date}` | |
| `devices.revoke` | `Revoke` | |
| `devices.revoke.title` | `Revoke {device}?` | |
| `devices.revoke.body` | `It is signed out of Forever Sixty at once.` | |
| `devices.revoke.failed` | `Could not revoke that device. Try again.` | |
| `devices.footer` | `Revoke signs a device out at once. To sign this one out, use Sign out on You.` | |
| `common.cancel` | `Cancel` | |

Banned words, in every string and every push: `Go Premium`, `Unlock`, `Pro`, `Limited time`, `!`. A test in each app greps the catalogue for them, as `premium-entry-copy.test.ts` does.

# 14. What the apps need from the API (for lanes C, A1 and A2)

The apps re-derive nothing. Five contract needs are not met today:

- **C1. `GET /v1/me/characters/{region}/{ruleset}/{name}/answer`.** The home answer, computed in Go from the code the Discord `/char` panel already uses (`discordbot/char_upgrades.go`, `rating.ScoreAndGear`). Fields:
  - `band`, `band_label`, `preset`, `spec`, `unit` (`dps` or `points`), `total_gain`, `not_modelled_count`, `synced_at`, `synced_from`;
  - `upgrades[]`: `slot`, `pick {id, name, quality, icon, item_level, source_line}`, `worn {id, name, quality, icon, item_level} | null`, `gain | null`, `modelled`;
  - `already_best[]` (slot labels);
  - `talents {state: optimized|differs|no_export|no_build, points_differ, band_points}`;
  - `state` (`ok`, `no_gear`, `under_20`, `no_spec`, `no_list`).

  The site's home could move to it later (one behaviour, one place). Note the site computes the same answer in TypeScript today (`lib/home/upgrades.ts`); C1 must match it on lane C's fixtures.
- **C2. The raid view** (`RaidView`, `signupAnswerView`), plus `next_raid` per guild on `/v1/me` or `GET /v1/guilds/{id}/home`. Two additions:
  - `invite_at` (nullable instant; a new raid field the web form sets and the scheduler reads);
  - `viewer.can` (`lock`, `move`) so the apps draw officer tools from words, never from a rank.

  Officer moves use the existing officer `PUT`; lock and unlock use the existing status actions. `raid.roster_moved` needs the acting officer's display name (their main character in the guild).
- **C5. `GET /v1/me/raid-nights/{raid_id}`.** The summary behind F4.22: kills, wipes, the viewer's loot (item id, boss, awarding officer), personal bests (boss, DPS, percentile, kill time), and coach notes when entitled.
- **C6. Item facts in the static data:** `bonding`, `unique`, `armor_type` and `min_standing` per item in `items/{class}.json`, from ItemSparse and Item, for the app and the site tooltip alike.
- **C3. `GET /v1/app/config`.** `{ "plans_link": { "ios": ["USA"], "android": ["US"] }, "plans_url": "https://foreversixty.gg/premium", "min_version": { "ios": "1.0", "android": "1.0" } }`. The `min_version` answer drives a blocking `Update Forever Sixty to keep signing up for raids.` screen when the API breaks old builds.
- **C4. The `/app/*` paths on the site.** `apple-app-site-association` and `assetlinks.json` at `foreversixty.gg/.well-known/`, and a redirect from each `/app/…` path to its page.
- **A2. Preferences and pushes.** Preferences read and write per event and channel, plus quiet hours and the zone. The push payload fields of §6.1, including `ask_spec`, the chosen `character_key`, the collapse ids and the per-locale text.

# 15. Acceptance captures

Each Wave 1 lane attaches captures from the iPhone 17 simulator (iOS 26) and a Pixel 9 emulator (API 35), at the default text size and the largest (iOS AX5 via `UIContentSizeCategory.accessibilityExtraExtraExtraLarge`, Android font scale 2.0). Reviewers compare them side by side with the boards and with each other (plan: "each pair's captures are reviewed side by side"). Fixture mode drives every state.

| Lane | Default size | Largest size |
|---|---|---|
| F1 | F1.1, F1.2 (consent or Custom Tab), F1.3, F1.4, F1.5 with three characters, F1.6 with the Revoke confirmation, Sign out confirmation, the 401 toast | F1.5 |
| F2 | F2.1, F2.2, F2.3, F2.4, under 20, no spec, no list, all match, F2.5, F2.6 at medium and large detents, F2.7 | F2.1, F2.6 |
| F3 | F3.1 signed in and signed out, F3.2, F3.3, not cached, error, the search field with a query | F3.1 |
| F4 | F4.1, F4.2, F4.3 pending and picked, F4.4, F4.5 menu, calendar sheet, F4.6 full scroll, F4.7 to F4.17, the rate-limited and race lines, keyboard focus on a segment, F4.18 to F4.21 as an officer and the same raid as a member (no officer controls), F4.22 and F4.23, a raid with no invite time | F4.1, F4.2 (2 x 2 segments), F4.3, F4.20 |
| F5 | F5.1 with a running, finished-unseen and failed run, F5.2, F5.3, F5.4, F5.5, empty, Run again pending and failed | F5.3 |
| F6 | F6.1, F6.2, a talent rank sheet, no export, no build | F6.2 |
| F7 | every §6.2 notification as delivered (lock screen or shade), each action's posted answer including askSpec and failure, `raid.roster_moved` in all four forms delivered inside quiet hours, F7.4 for a member and for an officer with Premium, F7.5, F7.6, quiet hours on with the time pickers | F7.4 |
| F8 | F8.1 (storefront US), F8.2 (storefront GB), F8.3, F8.4 in both regions, a seat, About | F8.1 |

Each lane also records:
- a VoiceOver and a TalkBack pass over its screens (a screen recording with captions on);
- the haptic calls in a unit test;
- the §11 budget it owns.

# 16. Checklist

- [ ] Four tabs in each platform's own control; each tab keeps its stack; re-tap pops to root.
- [ ] The selector on Home, Raids and BiS only; full name at every size; the cap and the Large Content Viewer at large text.
- [ ] Dark only; Increase Contrast swaps muted for nav.
- [ ] Barlow, Barlow Condensed and JetBrains Mono bundled; Cinzel only in the wordmark; every role scales.
- [ ] Every character or class image is the ringed crest; roster rows lead with the spec icon.
- [ ] Every item shows icon, quality border and quality-coloured name; every item opens the tooltip sheet.
- [ ] Home answer visible without scrolling on iPhone 17 and Pixel 9; rows in the server's order.
- [ ] BiS shows a runner-up line on every row.
- [ ] Sign-up saves on press, with haptics, `Saved.`, R3 moved line, spec ask once, and every refusal sentence.
- [ ] No queued writes offline; the offline strip on every cached screen.
- [ ] The reminder ask only in context, at most twice, never at first launch.
- [ ] Notification text from the server; In / Late / Out in the background with the answer posted in place; one thread per raid.
- [ ] Quiet hours apply except `raid.starting` and `raid.roster_moved`.
- [ ] One `raid.summary` per raider per night; no standalone report, coach, best or loot pushes by default.
- [ ] Officer tools drawn from `viewer.can` only; lock, unlock, and moves between In, Bench and off the raid.
- [ ] Invite and pull times both on the raid page; `raid.starting` keyed to invites.
- [ ] Talent tab totals summed from the drawn build; the differ count is the planner's `diffCount`; `Matches` only with no differing cell; prerequisite lines drawn.
- [ ] Tooltip shows Binds, Unique and armor type from the client tables, and reputation standings on sources.
- [ ] `See plans` only in link regions, fail closed; no Premium word elsewhere; no Premium push.
- [ ] Every element has the §8 states; 44 pt and 48 dp targets; VoiceOver and TalkBack labels per §7.1.
- [ ] §11 budgets measured on device.
- [ ] The copy catalogue matches §13; the banned-word test passes.
- [ ] Captures per §15 attached; ux-designer and wow-player SHIP on them.

# 17. Open questions only the owner can answer

1. **Approving characters in the app.** Locking and moving are in v1 (ruled). Approving waiting characters stays on the web. Do you want it in v1.1?
2. **Google Play in the US.** The plan reads the US storefront as a link region on both stores. Play's US external-link rules have changed during 2025. Confirm that the Play Console account is enrolled in whatever programme Play now requires for a link to web checkout, and that any fee is accepted. Until then, should `plans_link.android` ship empty?
3. **Premium features outside link regions.** Apple guideline 3.1.3(b) lets an app use a subscription bought on the web only when it is also offered as an in-app purchase, unless an exception applies. The spec shows Premium status and Premium features (Run again, coach pushes) to paying players everywhere and says nothing about buying. Do you accept the review risk the plan names, or should Premium-only actions hide outside link regions until App Review rules?
4. **Signed-in handoff to checkout.** `See plans` opens the browser, where the player may not be signed in. Should the API mint a one-time web session from the app token (`/v1/auth/app/handoff`), so checkout opens signed in? It is a security-sensitive endpoint, and its own review.
5. **Discord DMs as a channel.** The catalogue offers `Discord` per event. Today the bot posts to channels and the account's webhook, not DMs. Do you want the bot to send DMs (it needs the member's DM permission and a user-install or shared server)? Otherwise the Discord option is the existing channel events only.
6. **iPad and Android tablets.** v1 runs the phone layout. Is that acceptable for launch, or is a two-column raid page wanted?
