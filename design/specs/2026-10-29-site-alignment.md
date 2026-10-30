# Site alignment: every page in the account page's look (experience spec)

Owner, 2026-10-29, on a production capture of /account: "I like the design of this page, however we didn't align the rest of the site to have the same look and feel."

This spec turns /account into exact rules (§3), names the shared changes that carry most pages at once (§4), lists every route's sections with what changes (§5), and fences what must not change (§6). It builds on `design/specs/2026-10-28-polish.md` §4 and §4.5 and `design/DESIGN-SYSTEM.md`. Rows that the polish spec already owns are cited, not restated.

Evidence: the fixture site built from this worktree at 19:30 on 29 Oct (`FOREVER_DATA=fixture`), signed in through stubbed `/v1/*` as a Premium founder with two characters and Discord connected. The captures are `design/reviews/2026-10-29-site-alignment/before-*.png` (1440 for every route, 390 for the main ones). The after-mocks are `design/mocks/2026-10-29-site-alignment/after-*.png`, with before and after side by side in `compare-*.png`. Generators: `capture.spec.ts`, `capture-after.spec.ts`, `after-kit.js`, `compose.py` in the same folder.

Capture notes. Item, spell and spec icons draw as broken images because the fixture ships no icon files (52 icon 404s per page, none on real data). The guide's build tree shows skeleton bars because it is a `client:visible` island and a full-page shot never scrolls it into view. Neither one is a finding. Home was being rebuilt by another lane during this audit, so its rows (§5.1) go to that lane.

## 0. Decisions for the owner

| # | Decision | Default in this spec and the mocks |
|---|---|---|
| D1 | The page-head band. Today seven pages paint their own navy band (`#26405c` gradient: /bis, /logs, /rankings, /tiers, the planner, the guild page, the class pages). The account page has no band. | **Ruled 2026-10-29: no band anywhere**, the account page as it ships. Every `PageHead` sits on `--bg`. Every navy band goes, and no ember band replaces it. /account is untouched. (The mocks still show the ember band; where they differ, this ruling wins.) |
| D2 | IconTile size in a card head. The owner described "44 px"; the account page renders 36 px (measured), inside a 68 px head. | **Ruled:** 36 px, as shipped and as seen. A compact head variant (§3.3.1, about 48 px tall) serves dense data cards. |
| D3 | Tabs. Four tab treatments ship: outline boxes (guild), gold text (sim), underline (report), and pills (BiS band). No shared tab strip exists in `components/ui`. | **Ruled:** one shared `TabStrip`, styled like the report's underline tabs. It lands in stage 1 with the other shared pieces, before the per-route work. |
| D4 | About page copy (**adopted**). "Forever Sixty is a player's reference for World of Warcraft: Forever" contradicts tenet 14. | `Forever Sixty helps you play your class better in World of Warcraft: Forever. It is built and run by Commish LLC, an Arkansas company founded by players, and written and checked by players.` |

## 1. The question and the answer in view

The visitor's question on any page is the page's own question (see each route in §5). This spec does not move any answer. It adds one requirement: the answer's container reads as the same product as /account. That means the same page head, the same card with the same head, and the same row. The test is that a player who moves from /account to /bis, a guide or the planner sees no change of material, only of content.

Fold rule, unchanged: each route's answer stays above the fold at 360 and 1280. A card head adds 68 px. Where that would push the answer below the fold on a phone (§5.3 BiS spec, §5.5 planner), the compact head (§3.3.1) applies.

## 2. Reference

| Reference | What we borrow |
|---|---|
| /account (ours, owner-approved) | Everything in §3: the page head with eyebrow rule, the warm card surface with its gold top edge, the head (tile, title, right meta), the divider, full-bleed rows, tags, button recipes and ActionLinks. |
| The in-game character panel and the Battle.net launcher's game tiles | A framed panel with a title bar and an emblem at the left. The CardHead is our title bar, and the IconTile is the emblem. |
| Details! / ElvUI | Data stays in dense rows inside the frame. The frame never reformats the data (§6). |

## 3. The pattern, as exact rules

All values were measured on the built /account at 1440 (`before-account-1440.png`), or come from the token named.

### 3.1 Page ground and page head (`PageHead`, §4.2)

1. The page ground is `--bg` `#07090d`.
2. The head sits in the site column (`page-column`, 1248 px wide at 1440, gutters 18 px below 768), on the page ground. There is no band behind it (D1).
3. Main padding-top is 28 px from 768 and 16 px below.
4. The eyebrow row is a 24 × 1 px rule in `--gold` at 60%, a 12 px gap, then the eyebrow in Barlow 700, 11 px, `.18em` tracking, uppercase, `--gold`. It is the page's section name (ACCOUNT, BEST IN SLOT · LEVELING, GUIDES, BUILD PLANNER, GUILD).
5. The h1 sits 8 px under the eyebrow in `heading-page`: Barlow Condensed 800, 44/44, `.01em`, uppercase (34 px below 768). A class page's h1 keeps its class colour and its 64 px crest at the left (`lead` slot).
6. One optional job line sits 8 px under the h1: 15/22 `--text`, at most 640 px wide. One optional meta line follows (13 px `--muted`): the Updated stamp, or a descriptor such as "DPS · Updated 29 Oct". Nothing else about provenance goes here (DESIGN-SYSTEM principle 4).
7. The right slot (`aside`) is bottom-aligned to the h1's baseline box from 1024. It holds one of two things. The first is the identity block: a 36 px IconTile or 40 px crest, the name in 15 px 600 `--strong`, a 13 px muted sub, and one `button('secondary','sm')`. The second is the page's one control, for example the /premium Yearly | Monthly SegmentedControl. Below 1024 it stacks under the head at a 16 px gap.
8. A `below` slot holds tabs or a controls row (spec tabs, band and faction pills), 16 px under the head.
9. The head is never a card. It never has a band, navy or ember (D1). It never carries a source pill.

### 3.2 The card (`Card`, §4.1)

| Property | Value |
|---|---|
| Surface | `card-surface`: `--surface-card` (a gold 7.5% to 1.2% gradient over `#0d0f15`) |
| Border | 1 px `--color-line-card` `rgba(229,185,85,.17)` |
| Radius | `--radius-panel` 12 px, `overflow: hidden` |
| Top edge | inset `0 1px 0 --card-highlight` `rgba(245,210,122,.45)` on a primary card. `card-surface-quiet` drops it for secondary cards (prose asides, tables inside prose, tool-canvas frames). |
| Shadow | `--shadow-card` `0 16px 38px rgba(0,0,0,.55)` |
| Padding | 0 on the card itself. The head, body sections, rows and foot carry their own padding, so rows and dividers run edge to edge. |

### 3.3 The card head (`CardHead`, exists)

| Part | Value |
|---|---|
| Box | flex, `align-items: center`, gap 12 px, padding 16 px 20 px, min-height 68 px |
| Lead | `IconTile` 36 px (radius 8, ground `--tile` `rgba(229,185,85,.12)`, glyph 18 px in `--tile-ink` `#f5d27a`). Tones: gold by default, `discord`, `muted`. The lead is the item's own icon instead when the card is about one item (§6.4). |
| Micro-label (optional) | 11 px Barlow 700, `.14em`, uppercase, `--muted`, 4 px above the title. It is used only when the title is a value ("PLAN" over "Premium"). |
| Title | `h2` (`h3` inside a section that has its own `h2`). Alone, it is `heading-plan` uppercase: Barlow Condensed 800, 24/24, `.01em`, `--strong` ("DEVICES"). Under a micro-label, it keeps its copy's own case ("Premium"). |
| Right slot (optional, one of these) | a 13 px `--muted` meta ("Updated 3 hours ago", "20 to 29"); a `Tag` ("CONNECTED"); an `ActionLink` ("See all"); or one `button('secondary' or 'ghost','sm')` ("Reset…", "Approve all", "Plan a raid" may be primary when it is the view's next step). It is right-aligned with `margin-left: auto`. Below 480 px it wraps under the title, indented to the title's left edge (68 px). |
| Divider | the first body section's 1 px `--color-line-soft` `#1c2230` top rule. The head itself draws no border. |

#### 3.3.1 The compact head (dense data cards)

A data card whose body is a dense table or a short list of figures takes the compact head, so the page's main list stays near the fold. Stat weights, This set and Play it on /bis, the planner's spec trees and the planner's right-column cards (one rule for both columns), and similar data cards use it.

| Part | Value |
|---|---|
| Box | padding 10 px 16 px, min-height 48 px, gap 10 px |
| Lead | `IconTile` 28 px (glyph 14 px), or the game's own icon at 28 px |
| Title | `section-title` 18 px, uppercase (no micro-label) |
| Right slot | as §3.3, 13 px |

### 3.4 Body, rows and foot

1. A text body section has padding 16 px 20px 20 px, a 1 px `--color-line-soft` top rule, and 12 px between its children. A card's lone sentence (an empty or not-yet state) is 14/22 `--text`.
2. A row is `SettingsRow`: min-height 56 px, padding 10 px 20 px, gap 14 px, a 1 px `--color-line-soft` top rule. Its slots are an optional 32 px IconTile or 40 px crest, a label (15 px 600 `--strong`), an optional sub (13 px `--muted`), a right value (14 px), and a trailing control or a 16 px chevron when the row navigates. Hover is `--row-hover`. Focus is a 2 px inset gold ring.
3. A table inside a card keeps its own grid. Its header row is `quiet-label` (13 px 600 `--muted`, sentence case) and its rows are divided by the same 1 px `--color-line-soft`. Its body padding is `0 20px 8px`. The table is never rebuilt as SettingsRows (§6.1).
4. The foot has padding 16 px 20 px and a 1 px `--color-line-soft` top rule. Buttons sit left-aligned, 12 px apart (account "Manage billing", "Send test message" + "Unlink Discord").
5. Phone head: superseded by the compact head (§3.3.1), which This set and the planner's cards take at every width.

### 3.5 Tags, buttons, links

1. A state on a thing is a `Tag`: 22 px, 11 px 700 `.08em` uppercase, tones gold, success, muted and danger, with an optional dot. Examples are SUPPORTER, FOUNDER, CURRENT, MAIN, KILL, WAITING, BETA, and a role (DPS muted, TANK gold, HEALER success). Source pills stay `pill` (DESIGN-SYSTEM principle 4).
2. Buttons come from `button(kind, size)` (`lib/ui/button.ts`) only. There is one primary per view. The legacy `PRIMARY_BUTTON(_FIXED)` and `SECONDARY_BUTTON(_FIXED)` strings (`lib/planner/styles.ts`, used in 80 files) move to the recipe route by route in §5.
3. A standalone action that is not a button is an `ActionLink`: 14 px 600 `--gold`, a trailing 16 px chevron, 44 px tall, no underline, sentence case ("See all", "Pair another computer", "Open the Fury guide"). Tracked-caps gold links ("SEE ALL", "OPEN THE FURY GUIDE", "SEE 30 TO 39", "TIER LIST") and arrow-text links ("Get the addon →") become ActionLinks.
4. An inline link inside a sentence stays underlined (`INLINE_LINK`).

### 3.6 Spacing scale

| Step | px | Use |
|---|---|---|
| 4 | 4 | micro-label to title |
| 8 | 8 | eyebrow to h1; h1 to job line |
| 12 | 12 | head gap; children in a body; foot buttons; tag to name |
| 14 | 14 | row gap |
| 16 | 16 | head and body vertical padding; side-by-side cards in a grid of three or more; head to tabs |
| 20 | 20 | card side padding at every width |
| 22 / 32 | 22 below 768, 32 from 1024 (24 between) | between stacked cards, between columns, from the page head to the first card |

Two-column pages use the account's split: a main column and a 440 px rail with a 32 px gap. The rail is the `CharacterHeader` card's width. Below 1024 the rail follows the main column, except the `CharacterHeader`, which always leads. Cards in a row align to the top. A card is never stretched to fill an empty body (/account's columns are independent).

### 3.7 When a section is a card and when it stays bare

1. **Card.** Any bounded block with a title that does one job: a summary, a tool's input, a settings group, a list of rows, a table, a feed, a lookup form, or an empty or error state.
2. **Bare.** These sit on the page ground with `heading-section` (26 px uppercase) titles:
   - the `PageHead`;
   - long-form prose at the 720 px measure (guide bodies, legal, about, changelog);
   - a page's table of contents;
   - the home hero sky.
3. **Framed canvas.** A tool's work area (planner trees, report chart and timelines) sits in a `card-surface-quiet` frame whose head is the game's own icon and name, not an IconTile (§6.2).
4. **Never.** A section title floating above a separate box ("YOUR UPGRADES" over a panel, "THE LIST" over a table, "1. SIGN IN WITH BATTLE.NET" over a box). The title moves into the CardHead. This is the most common deviation on the site: home (5 sections), /bis (2), /sim (3), /setup (3), the guild tabs, and the /premium FAQ.

## 4. Shared changes first (ranked by how many sections each converts)

| # | Change | API | Converts |
|---|---|---|---|
| S1 | **`Card`**: `web/src/components/ui/Card.svelte` plus a `Card.astro` twin for static pages (no client JS). It wraps `CardHead` and the body rules of §3.2 to §3.4. | `glyph: GlyphName`, `tone?: TileTone`, `lead?: Snippet` (replaces the tile: an item icon, a spec icon, a crest), `title: string`, `label?: string`, `meta?: string`, `aside?: Snippet` (a Tag, an ActionLink or a button), `action?: { label: string; href: string }` (renders the ActionLink), `quiet?: boolean`, `body?: 'text' \| 'rows' \| 'table' \| 'none'` (padding per §3.4; `rows` and `none` are full-bleed), `foot?: Snippet`, `headingLevel?: 2 \| 3`, `compact?: boolean` (§3.3.1), `testid?: string`, `id?: string`, `children: Snippet` | 49 of the 95 ranked rows in §5. It replaces `Panel.astro` (3 pages), the account's last `StatePanel` (Your ratings), the four scoped `.panel-box` recipes (`PlayItPanel`, `StatWeightsPanel`, `BandPanelBody`, `GuideActionRail`; one duplicated style block each), and the hand-rolled `bg-raised border … rounded-panel` boxes (132 files) route by route. |
| S1b | **`card.ts`**: `CARD` becomes `'card-surface'` and `CARD_SURFACE` is renamed `CARD_QUIET = 'card-surface-quiet'` (the polish spec's §4.2 rename, not yet done). | none, a value change | 7 components at once: `ForeverScoreCard`, `RatingCard`, `CoachPanel`, `PlanCard`, `GuildStepCard`, `CheckoutGuildStep`, `PremiumOwnedPanel` |
| S2 | **`PageHead`**: `web/src/components/ui/PageHead.astro` plus a `.svelte` twin. It is extracted from `AccountTitle.svelte` and the account head row, which then use it. | `eyebrow: string`, `title: string`, `titleClass?: string` (class colour), `lead?` slot (crest 64), `job?: string`, `meta?` slot, `aside?` slot (identity block or one control), `below?` slot (tabs, controls row). No band prop: no page has a band (D1). | 17 heads: account, /bis, /bis/[class]/[spec] (via `ClassHeader` + `ArtPanel`), guides index, class and spec guides, planner (`PlannerHeader` + `PlannerHeaderBand`), /sim and every sim tool, /rankings (`RankingsHero`), /tiers (`TierPage`), /guild, the guild page (`GuildShell`), report, /logs, /premium, /setup, `Content.astro` (about, legal), 404 |
| S3 | **`CharacterHeader` surface**: `ch` takes `card-surface` (primary). The 1 px gold border becomes `--color-line-card`. Its emphasis comes from the crest, the class-coloured name and its one primary button. Regions, states and API are unchanged. | none | 6 pages: home, /bis, /bis spec, planner, /rankings, /logs |
| S4 | **Glyphs** added to `lib/ui/glyphs.ts` (24 grid, 2 px stroke, round caps). The paths are in `after-kit.js`. | `scale`, `bolt`, `list`, `pin`, `sword`, `tree`, `book`, `upload`, `download`, `share`, `layers`, `target`, `trophy`, `search`, `news` | every new card head below |
| S7 | **`TabStrip`** (D3): `web/src/components/ui/TabStrip.svelte` plus an `.astro` twin for link tabs. Underline tabs; links (`href`) or in-page tabs (`role="tablist"`, roving tabindex, arrows and Home/End); scrolls with a 24 px edge fade and the current tab scrolled into view. | stage 1 report | guild tabs, sim tools, report tables, BiS band (where it is tabs, not a control) |
| S5 | **`SettingsRow` for settings and lists of things**, and `Switch` for every on/off checkbox (18 files still render `type="checkbox"`) | existing | guild settings, sim characters, home Get started, guild overview rows, setup character line, planner "Include a simmed DPS" |
| S6 | **`ActionLink` for standalone actions** (§3.5.3) | existing | every tracked-caps, arrow-text and underlined standalone link named in §5 |

Order of work: S1b (one line, seven cards), then S2 and S3 (every page head and the identity card change in one pass), then S1 with S4 and S7 (stage 1, the shared pieces), then the routes in §5.17's blocker order and §5's order. S5 and S6 land per route.

## 5. Per route

Ranks: **Blocker** breaks a tenet a player sees, or is the owner's named gap. **Should** means the section does not read as the account pattern. **Polish** is a smaller step. "Polish §n" means the row is already owned by `2026-10-28-polish.md` and must be built in this pattern.

### 5.1 Home, signed in (`/`, `before-home-signed-in-*`, mock `after-home-signed-in-*`)

Another lane is editing `index.astro` and the home components now. These rows go to that lane, or follow its merge.

| Section | What changes | Rank | Component |
|---|---|---|---|
| Launch line under the hero cards | At 1440 the line "Launch Nov 4 · 15:00 PST · First raids Dec 9" is clipped under the Talents and Simulator cards, half off the band (this build). | Blocker | home lane's in-flight work; verify after merge |
| Your character | Moves onto the card surface (S3). The rest is unchanged. | Should | `CharacterHeader` |
| Talents, Simulator | Each becomes a Card: `tree` "TALENTS", `bolt` "SIMULATOR". Today's line is the body. A simmed figure, when present, is a `StatTile`. | Should | `Card` |
| Your upgrades | The floating "YOUR UPGRADES" title moves into a Card head: `list` "YOUR UPGRADES", meta "Level 60". `UpgradeRow`s keep their columns (body `table`). | Should | `Card` |
| Reserved voids | 160 px of empty ground under Your upgrades and about 200 px under Recent reports and Top guilds at 1440. The reserve holds the ready height while the shorter empty state is showing. The reserve must follow the state that rendered (DESIGN-SYSTEM Motion: nothing moves, and nothing is empty either). | Should | the islands' reserve |
| Look up a character | A Card: `search` "LOOK UP A CHARACTER", meta "See any character's Forever Score." The fields are unchanged, and `Look up` stays the lookup's primary. | Polish | `Card` |
| Get started | A Card: `star` "GET STARTED". The three arrow rows become `SettingsRow`s with a 32 px IconTile (`tree`, `bolt`, `upload`), label, sub and chevron. | Should | `Card`, `SettingsRow` |
| Another class | A Card: `user` "ANOTHER CLASS". The right slot holds two ActionLinks, "Tier list" and "All 28 specs" (today tracked gold caps). | Should | `Card`, `ActionLink` |
| Recent reports, Top guilds | Cards `chart` "RECENT REPORTS" and `trophy` "TOP GUILDS", each with the ActionLink "See all" in the head | Should | `Card` |
| Latest | A Card: `news` "LATEST" with the ActionLink "Everything we know". `FeedRow`s run full-bleed with row hover. | Should | `Card`, `FeedRow` |

### 5.2 Best in slot landing (`/bis`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `PageHead`. "Which spec is strongest? Tier list →" becomes the line and the ActionLink "Tier list". | Should | `PageHead`, `ActionLink` |
| Your character | S3 | Should | `CharacterHeader` |
| "ANOTHER CLASS OR SPEC" and the band box | One Card: `layers` "LEVEL RANGE", meta "Spec links open at this range". The five band pills are a `SegmentedControl` (Polish §5.4). The floating title goes. | Should | `Card` |
| Nine class cards | These keep the `BisClassCard` recipe (2 px class-colour top stroke, 56 px crest head, name in class colour) on `card-surface-quiet` with the `--color-line-card` border. Spec rows become 48 px rows with a chevron. HEALER and TANK become `Tag` muted and gold. | Should | `BisClassCard` |

### 5.3 Best in slot spec page (`/bis/warrior/fury`, mock `after-bis-warrior-fury-*`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Where to get it | Item names are plain white text with no icon, no quality colour and no tooltip (tenet 2). Each name becomes the shared `ItemPill`. The panel is a Card: `pin` "WHERE TO GET IT". Each source is a row: the instance as the label with its count, and the items as pills under it. | Blocker | `Card`, `ItemPill` |
| Page head | `ClassHeader` inside `PageHead`; `ClassHeader` and `ArtPanel` lose their navy gradient (D1). Crest, h1, spec tabs and the controls row are unchanged. | Should | `PageHead` |
| Your character | S3 | Should | `CharacterHeader` |
| This set | A Card (compact head): `bolt` "THIS SET", meta "Simmed 29 Oct" (replaces the ISO date, Polish §5.4). 37.5 becomes a `StatTile` (Polish §5.4). The two buttons stay secondary. The compact head (§3.3.1) keeps the figure in the first screen. | Should | `Card`, `StatTile` |
| Stat weights | A Card (compact head): `scale` "STAT WEIGHTS", meta "20 to 29". The table keeps Stat, bar, Scale, Per point and Error. The bar is 8 px (Polish §5.4). | Should | `Card` |
| Play it | A Card (compact head): `sword` "PLAY IT", meta "20 to 29", at natural height. "OPEN THE FURY GUIDE" becomes the ActionLink "Open the Fury guide". | Should | `Card`, `ActionLink` |
| The list | "THE LIST" and "Hover or tap an item for its stats" move into one Card head (`list`, meta = that line). Every column, the runners-up and the You column are unchanged (body `table`). | Should | `Card` |
| New at this level range | A Card: `star` "NEW AT THIS RANGE". "SEE 30 TO 39" becomes the ActionLink "See 30 to 39". | Polish | `Card` |

### 5.4 Guide (`/guides/warrior/fury`, mock `after-guide-warrior-fury-*`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `ClassHeader` (guide mode) inside `PageHead`; its navy gradient goes (D1). "DPS · Updated 29 Oct" and the BETA `Tag` stay. | Should | `PageHead` |
| Build, Rotation, Stat priority (the rail) | Three Cards replace the scoped `.panel-box`. Build: `tree` "BUILD", meta "Arms 17 · Fury 32 · Protection 2" (the per-tree line moves to the meta). Rotation: `sword` "ROTATION", meta "Level 60". Stat priority: `target` "STAT PRIORITY", meta "Hit = 1.00". "FULL PRIORITY ↓" becomes the ActionLink "Full priority". | Should | `Card` |
| "Sim this build" in Talents and builds | Renders underlined: the prose link rule leaks into the `a` button. Buttons inside `.prose` carry `text-decoration: none`. | Should | `button()` |
| Beta callout, Strengths, Weaknesses, In a raid | These move to `card-surface-quiet` with the 12 px radius. They keep their 2 px tone top stroke and micro-label, with no IconTile (prose asides). | Polish | `Card quiet` |
| Key talents, stat table, BiS table, race table, leveling strip | These move to `card-surface-quiet` (`GuideTableBox`). Columns are unchanged. | Polish | `GuideTableBox` |
| Prose body and On this page | Bare, unchanged (§3.7) | none | |

### 5.5 Planner (`/planner`, mock `after-planner-*`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Level 60 build card at 390 | It runs edge to edge with no 18 px gutter (`before-planner-390.png`, second screen). It needs the column's gutter. | Should | `BandCompare` |
| Page head | `PlannerHeaderBand` becomes `PageHead`. Class and Race `Select`s and the four `StatTile`s stay in the `below` slot. | Should | `PageHead` |
| Your character | S3 | Should | `CharacterHeader` |
| "Your gear has not been read yet." | A quiet Card: `device` muted tile, the sentence as the head title in its own case, and the two lines as the body. | Should | `Card quiet` |
| Arms and Fury tree panels | A framed canvas (§3.7.3) on `card-surface-quiet`. The head keeps the spec's game icon (28 px), the tree name in 18 px heading face, and points at the right. Grid and art are unchanged. | Polish | `TreeGrid` frame |
| Level 60 build | A Card (compact head): `layers` "LEVEL 60 BUILD". The duplicated "your level's range" line is dropped, from the meta and the body alike. | Should | `Card` |
| Point order | A Card: `list` "POINT ORDER", meta "0 spent" (today's bare "0" moves into the meta). | Should | `Card` |
| Share this build | A Card: `share` "SHARE THIS BUILD", with "Reset…" in the head's right slot as `button('ghost','sm')`. "Include a simmed DPS on the card" becomes a `Switch` row. | Should | `Card`, `Switch` |
| Import from addon | A Card: `download` "IMPORT FROM ADDON". "Don't have an export? Get the addon." becomes the line plus the ActionLink "Get the addon". | Should | `Card`, `ActionLink` |

### 5.6 Simulator (`/sim`) and tool pages (`/sim/weights`, `/sim/gear` and the rest)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Character strip at 390 | "PLAN SIM LOGS RANKINGS FORGET" starts 5 px from the screen edge, past the 18 px gutter, and is cut on the left (tenet 4). | Blocker | `CharacterStrip` |
| Character strip, every width | A row of tracked-caps words. It becomes the first row of the "YOUR CHARACTERS" card: crest 40, name, then ActionLinks "Plan", "Logs", "Rankings". "Forget" becomes `button('ghost','sm')`. | Should | `SettingsRow`, `ActionLink` |
| Page head, /sim | `PageHead` with eyebrow "TOOLS". The tool tabs go in `below` (D3). | Should | `PageHead` |
| Page head, tool pages | The tabs sit above the h1, the reverse of /sim. Use the same `PageHead`, with tabs in `below`. | Should | `PageHead` |
| Repeated notes, tool pages | "The simulator runs damage specs only…" shows twice, plus "models level 60" and the tool line: three 15 px muted paragraphs 53 px apart. Keep one job line in the head and one muted note in the inputs card foot. Delete the rest. | Should | copy |
| Your server runs | A Card: `server` "SERVER RUNS". `LoadError` and `EmptyState` sit in the body. | Should | `Card` |
| Quick run card | The page's primary Card: crest 40 row, name, descriptor, addon `Tag`, and `Run sim` as the view's one primary in the row's right slot. | Should | `Card`, `SettingsRow` |
| Your characters | A Card: `user` "YOUR CHARACTERS". Each character is the account's `CharacterRowLink` row (crest 40, `Current` `Tag`, "Sim" secondary sm, "Not simulated yet" as a muted value). "SIM SOMETHING ELSE" goes to the card foot. | Should | `Card`, `SettingsRow` |
| Your sims | A Card: `chart` "YOUR SIMS". The Show `Select` goes in the head's right slot (44 px filter size). | Should | `Card`, `Select` |
| Input panels (From the addon, From a build, From a logged fight, Your characters) | Four Cards: `download`, `tree`, `chart`, `user`. Fields and Load buttons are unchanged. | Should | `Card` |
| Stat weights caveat | A quiet Card with an `ActionLink` "Open Top Gear" in place of the outline button | Polish | `Card quiet` |
| Stats to weigh | A Card: `scale` "STATS TO WEIGH". The 24 native checkboxes need the site's multi-select recipe, which is a new control and gets its own review (tenet 10). Out of this lane. | Polish | n/a |

### 5.7 Rankings (`/rankings`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `RankingsHero` becomes `PageHead`. "Comparing specs, not characters? Tier list →" becomes the line plus an ActionLink. | Should | `PageHead` |
| Your character | S3 | Should | `CharacterHeader` |
| "FURY WARRIORS" block (title, "Another class or spec" underlined link, legend, Characters and Guilds) | One Card: `trophy` "FURY WARRIORS". The Characters and Guilds `SegmentedControl` goes in the head's right slot. The filter bar is the first body section. The legend's six `Tag`s go in the foot. "Another class or spec" becomes an ActionLink under the title. | Should | `Card` |
| Empty state | `EmptyState` inside that card (it already has the ActionLink) | Polish | `EmptyState` |

### 5.8 Tier list (`/tiers`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `TierPage` band becomes `PageHead`. The role and faction pills go in `below`. | Should | `PageHead` |
| "Fury Warrior is 2nd of 20" | A primary Card row: crest 40, the sentence, and "YOUR BEST IN SLOT →" as `button('secondary','sm')` "Your best in slot" | Should | `Card` |
| Explanations above the table | Polish §5.14 (one line, the rest under "How the list is made") | Should | copy |
| The table | One Card: `list` "DPS · LEVEL 60". The tier column keeps its tone stroke. Rows, bars and figures are unchanged. Row links BIS, GUIDE and PLANNER stay compact text links (a table, §6.1). | Polish | `Card` |

### 5.9 Guilds landing (`/guild`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `PageHead`, eyebrow "GUILDS" | Should | `PageHead` |
| Lookup | A Card: `search` "OPEN A GUILD". The guild-name field gets a glyph and the placeholder "Guild name". Selects become `Select` (Polish §5.12). | Should | `Card`, `Select` |
| Top guilds | A Card: `trophy` "TOP GUILDS" (today a gold 11 px label over a loose error). "Run a guild? Open its page and choose Claim this guild." moves to the foot. | Should | `Card`, `LoadError` |

### 5.10 Guild page (`/guild/us/pvp/olympus-xxvii`): head and the overview, roster, raids, loot and settings tabs

| Section | What changes | Rank | Component |
|---|---|---|---|
| Settings tab, a second h1 | The tab renders the full settings page under the guild band, with its own 44 px h1 "OLYMPUS XXVII · GUILD SETTINGS". That makes two h1s and a repeated name. Inside the guild page the settings render without their h1. The standalone `/guild/…/settings` keeps it. | Blocker | `GuildSettings` |
| Settings, every group | Billing, Who pays, Records, Discord, Default report visibility, Officers rank, Invite link and Guild crest are bare headings over loose text, outline buttons, a native select and an underlined link. This is the content the account page solved, and the owner's named gap. Use four Cards in the account's two-column split. **Plan** (main column) uses the account Plan card recipe (§4.3 of the polish spec): `crown` tile, micro-label "GUILD PLAN" over "Guild", meta "Renews 30 Nov 2026", the rows Price, Billed to and Who pays (Transfer the plan, disabled with today's sub), and a foot with Manage billing secondary and Move to Raid guild ghost. **Discord** (main column) is the account Discord card recipe verbatim: `discord` tone, Connected `Tag`, destination row, "Use a different channel" link row, foot with Send test message and Remove (danger). **Members and access** (rail) holds three rows: Default report visibility (trailing `Select`), Officers are ranks at or above (trailing 72 px field and Save sm), and Invite link (sub "Anyone with this link can join as a member.", trailing Rotate secondary sm). **Records and crest** (rail) holds Download attendance (CSV) and Download loot (CSV) as link rows with the `download` glyph, and the Guild crest row (crest 40, "Default: your faction's logo", Choose an image secondary sm). | Blocker | `Card`, `SettingsRow`, `Select`, `Tag` |
| Head | `GuildShell` becomes `PageHead`. The faction art at the band's right is dropped with the band (D1). The next-raid line and In, Late, Bench and Out stay in `aside`. | Should | `PageHead` |
| "Guild settings" link and the "Claimed by … Invite link · SHOW" bar in the head | These move into the Settings tab (one job per region). The head keeps identity and the next raid. | Should | `GuildOfficerStrip` |
| Tabs at 390 | "PROGRESSIO" is cut at the right edge with no affordance (tenet 4). The strip scrolls with a 24 px fade and the current tab scrolls into view (the Polish §5.3 rule). | Blocker | tabs (D3) |
| Tabs, every width | Seven outline boxes become `TabStrip` (D3) | Should | `TabStrip` |
| Overview cards (Roster, This week's raid nights, Progression, Readiness, Loot) | Already cards. They take the surface and a CardHead with tiles `user`, `chart`, `trophy`, `guild`, `star`. "SEE ALL" becomes the ActionLink "See all". Raid-night and readiness lines become `SettingsRow`s with values at the right. | Should | `Card`, `SettingsRow` |
| Roster: summary, Filter and Sort | A Card: `user` "ROSTER", meta "26 members". Filter and Sort chips form the first body section. | Should | `Card` |
| Roster: waiting for approval | A Card: `bell` "WAITING FOR APPROVAL", meta "3", with "Approve all" as secondary sm in the head. Rows are crest 40, name and sub, with Approve (secondary sm) and Decline (ghost sm). The per-row WAITING FOR APPROVAL `Tag` goes (the card says it). | Should | `Card`, `SettingsRow` |
| Roster: the table | It stays a table inside a Card. Twenty 44 px REMOVE outline buttons, one per row at equal weight (tenet 11), become `button('ghost','sm')`. At 390, "Remove" is an underlined text link and becomes the same ghost button. | Should | `button()` |
| Raids | "RAIDS" and "PLAN A RAID" form a Card head: `chart` "RAIDS", with Plan a raid as the tab's primary. "PLANNED" and "RAID NIGHTS" become two Cards. Rows are unchanged. | Should | `Card` |
| Loot: item cards | Each item's card title ("Scale of Onyxia", "Stormrage Cover") is plain heading text with no icon, no quality colour and no tooltip (tenet 2). The head's lead is the item's icon (36 px, quality border), the name in its quality colour, and the shared tooltip. AWARD becomes secondary sm. | Blocker | `Card` (`lead`), `ItemPill` |
| Loot: tab head | A Card: `star` "LOOT", with the Boss chooser as a `Select` in the head. The caveat line goes in the foot. | Should | `Card` |

### 5.11 Report (`/reports/fixture2abcd?fight=3`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | One 1,060 px meta line runs zone, date, fights, state, the whole crediting sentence, the boss and deaths together. Use `PageHead`: eyebrow "REPORT", the h1, a meta line "Sanguine Depths · 26 Sept 2026 · 4 fights", then a `Tag` muted "Not credited" and the sentence as a 13 px muted line under it. Sim this fight and Copy link go in `aside`. | Should | `PageHead`, `Tag` |
| Fight list | A Card: `list` "FIGHTS", meta "2 pulls". Fights are rows with a `Kill` success Tag. "SHOW 2 TRASH FIGHTS" becomes a row button. | Should | `Card`, `SettingsRow` |
| Player table | Bare on the page ground today. It goes into a Card: `chart` "DAMAGE DONE" (the active table's name), meta "5 players · 1 death · 1:00". Columns are unchanged. | Should | `Card` |
| Damage taken by ability, Deaths, Performance rating | Cards with tiles `chart`, `target`, `star`. "Damage Taken tab" and the other jump links become the ActionLink in the head. | Should | `Card`, `ActionLink` |
| At pull | A Card: `user` "AT PULL". Rows keep "Gear in the planner" and "Sim" as ActionLinks. | Polish | `Card` |
| Chart and timelines | A framed canvas (§3.7.3), unchanged | none | |

### 5.12 Logs (`/logs`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `.band` becomes `PageHead` | Should | `PageHead` |
| Your character | S3 | Should | `CharacterHeader` |
| Your reports, Recent public reports, The companion, Upload a log | Each already has a title inside its box. Each takes `Card` with tiles `chart`, `news`, `device` and `upload`. "Live, while you raid" and "For a night already logged" stay as meta. "Show pairing code" (underlined) becomes an ActionLink. "PAIR ANOTHER DEVICE" becomes a body sub-label in `label` muted. | Should | `Card`, `ActionLink` |

### 5.13 Plans (`/premium`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `PageHead`. Yearly and Monthly are already the head's aside, which is the account pattern. | Polish | `PageHead` |
| Plan cards, coach sample | Surface via S1b. The `PlanCard` recipe is unchanged (polish §3). | Polish | S1b |
| How the guild plans work | A label and value grid (the shape the account Plan card replaced). It becomes a Card: `book` "HOW THE GUILD PLANS WORK", with rows of label (15 px 600) and sub (the answer, 14/22 `--text`), no chevron. | Should | `Card`, `SettingsRow` |
| FAQ | The h2 renders at about 30 px, larger than every other section title on the page. It becomes a Card: `book` "QUESTIONS", with each question as a full-bleed `Disclosure` row. | Should | `Card`, `Disclosure` |

### 5.14 Get set up (`/setup`)

| Section | What changes | Rank | Component |
|---|---|---|---|
| Page head | `PageHead`, eyebrow "ADDON AND COMPANION". The three step chips go in `aside`. | Should | `PageHead` |
| Steps 1 to 3 | "1. SIGN IN WITH BATTLE.NET" floats over a box. Each step becomes a Card whose lead is a 36 px numbered tile (the `IconTile` geometry with the step's digit in `figure-stat` 18 px). Titles are "SIGN IN", "THE ADDON" and "THE COMPANION". The right slot holds a state `Tag`: "Done" success or "To do" muted. | Should | `Card` |
| Character line in step 2 | "Obnoxious Yell · Fury Warrior · Open in planner · Open in simulator · Forget" becomes a `SettingsRow`: crest 40, name, then ActionLinks, then "Forget" ghost sm. | Should | `SettingsRow` |
| Install buttons | Polish §5.7 (one primary) | Should | `button()` |
| Try an export, Manual install, Everything the addon does | Full-bleed `Disclosure` rows inside the step's card | Polish | `Disclosure` |

### 5.15 About, legal, 404, guides index

| Route | Section | What changes | Rank | Component |
|---|---|---|---|---|
| /about | Page head meta | "Confirmed by Blizzard or by multiple sources" on our own about page is a false provenance claim (tenets 7 and 8; the Polish §5.11 rule extends to every `Content.astro` page with no external source). Show "Updated 24 Oct" only. | Blocker | `Content.astro` |
| /about | First sentence | D4 | Should | copy |
| /about, /terms, /privacy, /refunds | Head | `PageHead` with eyebrow "ABOUT" or "LEGAL",. Prose stays bare at 720. | Polish | `PageHead` |
| 404 | Body | Polish §5.13 in this pattern: `PageHead` (eyebrow "NOT FOUND", "NO PAGE HERE"), then a Card `user` "PICK YOUR CLASS" with the nine crests, then the lookup Card (§5.9). | Should | `Card` |
| /guides | Class cards | A second class-card design for the same job as /bis. Use the `BisClassCard` recipe (§5.2) with role `Tag`s by tone (Polish §5.14). | Should | `BisClassCard` |
| /guides | "Your guide: Fury Warrior →" | A primary Card row: crest 40, micro-label "YOUR GUIDE" over "Fury Warrior", and the ActionLink "Open" | Should | `Card` |

### 5.16 Counts

| Route | Blocker | Should | Polish |
|---|---|---|---|
| Home (signed in) | 1 | 8 | 1 |
| /bis | 0 | 4 | 0 |
| /bis spec | 1 | 6 | 1 |
| Guide | 0 | 3 | 2 |
| Planner | 0 | 8 | 1 |
| /sim and tools | 1 | 9 | 2 |
| /rankings | 0 | 3 | 1 |
| /tiers | 0 | 3 | 1 |
| /guild | 0 | 3 | 0 |
| Guild page and tabs | 4 | 9 | 0 |
| Report | 0 | 4 | 1 |
| /logs | 0 | 3 | 0 |
| /premium | 0 | 2 | 2 |
| /setup | 0 | 4 | 1 |
| About, legal, 404, /guides | 1 | 4 | 1 |
| **Total** | **8** | **73** | **14** |

/account (the reference) changes only through S2 (it renders `PageHead`, pixel for pixel). D1 leaves it untouched.

### 5.17 Blocker order (ruled 2026-10-29)

1. Guild settings groups (§5.10, the owner's named gap), with the second h1.
2. Tenet-2 item icons: guild loot item names (§5.10) and /bis "Where to get it" names (§5.3), each with icon and quality colour.
3. The rest: the /sim strip at 390, the guild tabs at 390, the /about provenance meta.
4. The home launch line (§5.1) is tagged **verify after merge** (the home lane owns it).

## 6. What must not change

1. **Data-dense tables stay tables.** The BiS list, tier list, roster, report tables, stat weights, guide tables and rankings rows keep their columns, order, alignment and `numeric` figures. They gain a Card around them and nothing inside them changes. A table is never rebuilt as SettingsRows or cards per row.
2. **Tool canvases stay full width and keep the game's art.** Planner trees, report charts and timelines, and sim result charts sit in a quiet frame. Their head is the game's spec or spell icon and name, never an IconTile.
3. **The home hero sky stays the home's alone.** No legacy navy band survives on any page, and no other page has a band (D1).
4. **The game's icons are never replaced by glyphs.** An IconTile names a card's job (a scale for weights). A card about one item, spec, class or character leads with that thing's own icon or crest (DESIGN-SYSTEM principle 2, one crest language).
5. **One job per page.** No card restates another page's job. Moving the guild officer strip into Settings follows this rule.
6. **One primary per view.** A Card never adds a filled gold button. The `CharacterHeader`'s button stays the page's one primary where it is mounted.
7. **`CharacterHeader` regions, states and copy** are unchanged. Only its surface changes (S3).
8. **Long-form prose stays prose** at the 720 px measure with `heading-section` titles: guide bodies, legal, about, changelog.
9. **Source pills stay pills.** Provenance stays quiet: no pill in a page head or a card head.
10. **Copy is unchanged** except the lines this spec gives verbatim (§5.6 repeated notes deleted, §5.11 report meta, D4).

## 7. States (every Card; touch and keyboard parity)

| State | Head | Body |
|---|---|---|
| Default | as §3.3 | as §3.4 |
| Loading | renders at once (static title, tile and meta) | a `Skeleton` sized to the ready body for that data's state, so no reserve outlives an empty state (the §5.1 voids) |
| Empty | unchanged; a count meta reads "0" or is omitted | `EmptyState`: one 14/22 line, then at most one ActionLink or secondary button |
| Error | unchanged | `LoadError`: the short line and "Try again" as `button('secondary','sm')` |
| Hover | none. A card is never a link as a whole. | rows: `--row-hover`; ActionLink: `--gold-hover`, chevron 2 px right |
| Focus-visible | head actions: 2 px `--gold` ring at 2 px offset | rows: 2 px inset gold ring |
| Active | head actions per the button and link recipes | rows: 10% gold ground |
| Disabled | per recipe (45%, not-allowed) | row: label and value muted, no chevron |

Keyboard: a card head's action is the first tab stop after the title. Rows are links or buttons in reading order. Touch: every row and head action is 44 px or has a 44 px hit area below 1024.

## 8. Phone and desktop

1. Card side padding is 20 px at every width (as /account at 390).
2. The head title stays 24 px. Below 480 px the right slot wraps under the title at a 68 px indent.
3. The compact head (§3.3.1) applies only where named.
4. Two columns collapse below 1024: the `CharacterHeader` first, then the main column, then the rail.
5. Grids of three cards go 1-up below 768 and 2-up from 768 to 1023, with 16 px gaps.
6. These never collapse: the CardHead title, a table's columns (tables scroll horizontally inside the card body with a 24 px edge fade, never off the page), and the one primary button.
7. At 2000 the site column stays centred at its max width. Cards never stretch past it (wide-screen rule).

## 9. Performance and polish limits

- `Card.astro` and `PageHead.astro` ship no client JavaScript. Static pages stay static, and /premium still ships no `_astro` script.
- No layout shift. The CardHead has a fixed 68 px min-height (48 px compact). Tiles are inline SVG (no icon request). Measure CLS on /, /bis/warrior/fury, a guide and /planner before and after. The target is 0, at most 0.05 (`web/lighthouserc.json`).
- No missing-icon flash: item, spell and spec icons keep their width and height and hide on error.
- Reduced motion turns off the chevron nudge and the row-hover transition.
- Lighthouse budgets are unchanged and binding.

## 10. Acceptance screenshots

Every capture is of the built site beside the matching board at the same scale.

| # | Surface and state | 360 | 390 | 1024 | 1440 | 2000 |
|---|---|---|---|---|---|---|
| A1 | Home signed in, beside `compare-home-signed-in-*` | x | x | x | x | x |
| A2 | /bis/warrior/fury, beside `compare-bis-warrior-fury-*`; Where to get it with item pills and one tooltip open | x | x | x | x | x |
| A3 | /guides/warrior/fury, beside `compare-guide-warrior-fury-*` | x | x | x | x | x |
| A4 | /planner, beside `compare-planner-*`; Share with the Switch on and off | x | x | x | x | x |
| A5 | Guild settings, standalone and as the tab (one h1), Plan, Discord, Members and access, Records and crest | x | x | | x | |
| A6 | Guild tabs at 360 and 390 scrolled, with the fade; Loot item heads with tooltips | x | x | | x | |
| A7 | /sim at 360 and 390 (no strip past the gutter), /sim/weights | x | x | | x | |
| A8 | Every other route in §5, after | | x | | x | |
| A9 | One Card in each of loading, empty and error on home, /sim and the guild overview; a head action hovered and focused | | x | | x | |
| A10 | /account after S2, beside `before-account-*` (the reference must not regress) | | x | | x | x |

## 11. Files

| Path | What |
|---|---|
| `design/reviews/2026-10-29-site-alignment/before-*.png` | 42 before captures (every route at 1440; the main routes at 390) |
| `design/mocks/2026-10-29-site-alignment/after-*.png` | after-mocks: home signed in, /bis spec, guide, planner at 1440 and 390 |
| `design/mocks/2026-10-29-site-alignment/compare-*.png` | before and after side by side, same scale |
| `design/mocks/2026-10-29-site-alignment/capture.spec.ts`, `capture-after.spec.ts`, `stubs.ts`, `after-kit.js`, `compose.py`, `playwright.capture.config.ts` | the generators. Run from `web/` against a fixture build served on `CAPTURE_PORT` (default 4391), after linking `web/node_modules` into the folder (uncommitted). |
