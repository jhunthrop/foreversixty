# Web counterparts of the app screens (390 px)

The reference set for the brand-consistency review of the iOS and Android apps. Captured 2026-11-10 from the live site
(https://foreversixty.gg, signed out) and from the static web mocks in `design/mocks/`.

- Viewport 390 x 844, deviceScaleFactor 2 (PNGs are 780 px wide), Chromium with `isMobile` and `hasTouch`.
- Two passes: `colorScheme: dark` and `colorScheme: light`. The light pass renders the same dark site (the site sets
  `color-scheme: dark`), which proves it does not follow the OS theme.
- Per page: `<name>-<scheme>-top.png` is the first viewport, `<name>-<scheme>.png` the full page. A page taller than
  3000 CSS px (6000 px in the PNG) also has `<name>-<scheme>-crop.png`, its top 1200 CSS px (2400 px).
- The script: `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/web-shots/shoot.mjs`
  (scratch; it uses the Playwright install of another worktree because `web/node_modules` in the main checkout was incomplete).

## Pages

Each row has the `-dark`, `-light`, `-top` variants, plus `-crop` where noted.

| Name | Source | State | Counterpart of |
|---|---|---|---|
| `web-home` | `https://foreversixty.gg/` | live, signed out: class picker, Latest on Forever feed, How to start, First raids open countdown | 1 Home, 4 Activity (feed) |
| `web-home-mock` (crop) | `design/mocks/web/web-home-60.html` | signed-in mock, level 60 Fury Warrior: raid card with countdown, sign-up counts, I'm in / Late / Can't come, Before the pull, Your upgrades | 1 Home (raid day), 2 sign-up |
| `web-guild-overview` | `https://foreversixty.gg/guild/us/pvp/caffeine` | live public guild page, signed out (Overview tab; empty: no raid logged) | 2 Raid tab |
| `web-guild-raids` | `.../guild/us/pvp/caffeine#raids` | live, Raids tab, signed out ("No public raid nights yet") | 2 Raid tab |
| `web-raid-mock-member` (crop) | `design/mocks/web/web-guild-dense-raid-member.html` | signed-in member mock of a planned raid: Your sign-up (spec and role, note), Sign-ups list by role | 2 Raid + sign-up |
| `web-raid-mock-officer` (crop) | `design/mocks/web/web-guild-raid.html` | officer mock of the same raid: Edit, Lock sign-ups, Ping on Discord, Swap in, per-row Edit | 2 Raid (officer) |
| `web-raid-direction` (crop) | `design/mocks/directions/html/b3-guild-raid.html` | redesign direction mock for the guild raid page | 2 Raid |
| `web-bis` (crop) | `https://foreversixty.gg/bis/warrior/fury` | live, signed out, 20-29 Alliance: paperdoll card, What to stack, lists | 3 Gear |
| `web-bis-mock` (crop) | `design/mocks/web/web-bis-dps-60.html` | signed-in BiS mock, level 60 DPS | 3 Gear |
| `web-character` | `https://foreversixty.gg/character/us/pvp/pink-shorts` | live public character with imported gear: Gear paperdoll, upgrades, fights | 3 Gear, 4 Activity (character fights) |
| `web-character-mock` (crop) | `design/mocks/web/web-character-owner-60-premium.html` | owner view mock, level 60, Premium | 3 Gear, 4 Activity |
| `web-account-menu-direction` | `design/mocks/directions/html/b4-account-menu.html` | direction mock of the signed-in account sheet: account, Your characters, Paste an export, Character page, Account, Plan, Get set up, Sign out. It still shows an old fixture character name; the current app fixture is Gorvash | 5 Profile drawer |

## Interactions (dark only)

| File | Source | State | Counterpart of |
|---|---|---|---|
| `web-menu-step1-closed.png` | live `/` | header, menu closed | 5 Profile drawer |
| `web-menu-step2-pressed.png` | live `/` | finger down on the menu button | 5 |
| `web-menu-step3-open.png`, `-open-full.png` | live `/` | menu open: Sign in row, character lookup, BiS, Guides class grid, Tier list, Sim, Planner, Logs ... | 5 |
| `web-profile-step1-closed.png` | live guild page | header with the signed-out character button | 5 Profile drawer |
| `web-profile-step2-pressed.png` | live guild page | finger down on the character button | 5 |
| `web-profile-step3-open.png` | live guild page | bottom sheet "Sign in to Forever Sixty" (Battle.net, Paste an export, Get the addon and companion, Get set up) | 5 |
| `web-bis-tooltip-step1-before.png` | live `/bis/warrior/fury` | paperdoll, head selected by default | 3 Gear item sheet |
| `web-bis-tooltip-step2-pressed.png` | live BiS | finger down on the Chest cell | 3 |
| `web-bis-tooltip-step3-after.png` | live BiS | Chest selected (cell ring moves); the tooltip opens inline under the paperdoll, below the fold | 3 |
| `web-bis-tooltip-step4-scrolled-to-tooltip.png` | live BiS | scrolled to the inline tooltip: item, stats, +DPS vs, source, Open in the sim | 3 item sheet |
| `web-bis-band-step1-before.png` | live BiS | control bar "20-29 · Alliance  Change" | segmented controls |
| `web-bis-band-step2-pressed.png` | live BiS | finger down on Change | segmented controls |
| `web-bis-band-step3-open.png`, `-open-full.png` | live BiS | bar open: Spec tabs, Level segments, Faction segments, Race select, Sources toggle chips | segmented controls |
| `web-guild-tab-step1-overview.png` | live guild page | Overview tab selected | 2 Raid tab switch |
| `web-guild-tab-step2-pressed.png` | live guild page | finger down on Raids | 2 |
| `web-guild-tab-step3-loading.png` | live guild page | right after the tap: already Raids; no loading state is visible (the switch is client-side and instant) | 2 |
| `web-guild-tab-step4-raids.png` | live guild page | Raids tab settled | 2 |
| `web-raid-countdown.png` | live `/` | element capture of the First raids open panel (date, days/hours, Barrow Deeps, Hyjal Summit, Onyxia's Lair) | 6 Notifications (nearest web surface) |
| `web-signup-mock-step1-before.png` | `web-home-60.html` | raid card with I'm in / Late / Can't come | 2 sign-up |
| `web-signup-mock-step2-hover.png` | same | I'm in hovered (CSS :hover) | 2 |
| `web-signup-mock-step3-pressed.png` | same | I'm in pressed (CSS :active) | 2 |
| `web-signup-mock-step4-late-hover.png` | same | Late hovered | 2 |

## Not captured

- **Signed-in flows on the live site** (raid sign-up, change sign-up, the signed-in character menu, home raid day,
  activity feed for your own character): they need a session. They exist only in the mocks above.
- **After-state of a sign-up**: the web mocks are static HTML with no script, so the buttons show hover and pressed
  from CSS but cannot change state. No mock in the repo shows a web "after sign-up" or "change sign-up" state.
- **The signed-in header menu in the web mocks**: it is script-driven on the site, so it cannot be opened in a static mock;
  `web-account-menu-direction` is the closest.
- **Live guild raid content**: the only public guild found (Caffeine, US PvP) has no raid nights or progression, so
  the live guild captures show the empty states. No public guild is listed on the rankings guild board yet.
- **Notifications**: no web page; `web-raid-countdown.png` stands in.

## Observations from the 390 px pass (for the reviewers, not verified as bugs)

- `web-guild-dense-raid-member.html` and `web-guild-raid.html` overflow to 681 CSS px: the officer note field and the
  sign-ups table run off the right edge (visible in `web-raid-mock-member-dark-top.png`).
- `web-bis-dps-60.html` has one block (`div.tl`) reaching 773 CSS px; it may sit in a horizontal scroller.
- The light pass is byte-identical to the dark pass in the first viewport of every page checked (home, guild,
  BiS, character, the home, raid-member and BiS mocks, the account-menu direction).
