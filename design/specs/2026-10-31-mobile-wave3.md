# Forever Sixty for iPhone and Android, Wave 3: experience spec

Author: ux-designer, 2026-10-31. It extends `design/specs/2026-10-30-mobile-apps.md` (the v1 spec). Everything in v1 holds unless a section here changes it: tokens (§4), shared controls (§4.5), the screen rules (§5.0: loading, offline strip, error, signed out), the copy rules, the state table (§8), the hierarchy rules (§9), the accessibility rules (§7.1), the budgets (§11) and the capture rules (§15). Tenets: `docs/tenets.md`, with 2, 7, 9 to 11 and 14 applied screen by screen.

Owner direction (2026-10-31): "make the mobile app as feature rich as possible", then "build the server features now, stop deferring things". So:
- **Nothing here is split into "ships now" and "waits".** Each new endpoint is a contract change that its feature lane implements in `api/` together with both apps (§10).
- The held lanes A1 to A3 are merging, so `/v1/app/config`, app sign-in and push exist for these lanes to use.
- A lane still ships dark behind a feature switch in `GET /v1/app/config` (plan lane P1) and turns on in both apps together.

A wow-player review and a ux-designer review follow this one: `design/reviews/2026-10-31-mobile-wave3/`.

## Mock

`design/mocks/2026-10-31-mobile-wave3/`: `gen_wave3.py` (with `w3kit.py`, `guide_data.py`, `w3_tree.py`, `screens_w3a.py`, `screens_w3b.py`, `screens_w3c.py`) writes `renders/*.html`; `render.mjs` screenshots them at 1.5x. `extract_w3.py` rewrites `w3.json` and `assets/` from a checkout holding `data/builds/1.60.1.70658/`. It reuses the v1 kit (`../2026-10-30-mobile-apps/kit.py`), so chrome, tokens, crests and item icons are identical to the v1 boards.

```
python3 design/mocks/2026-10-31-mobile-wave3/gen_wave3.py && node design/mocks/2026-10-31-mobile-wave3/render.mjs
```

Data rules for the boards:
- **Real, from build 1.60.1.70658:** the tier list order, gaps and tier letters (the tier rules of `web/src/lib/tiers/tier-rules.ts` over the BiS band files, Horde, 60, Phase 1 raid); spell ranks (`spellranks.json`); talents and their rank text; the Spritecaster Cape tooltip, drops (`loot.json`) and the specs that list it; key dates (`dates.json`); changelog titles.
- **Real, from the v1 data (build 70532):** the upgrade rows on the widgets and Siri answer, and the Paperdoll items.
- **Invented, and marked so in the captions:** people, guilds, parses, DPS on the meters, rankings, the raid night, sizes of downloads.
- **Guide content is rendered from the guide's own files** (`guide_data.py`): stat priority, overview, strengths and weaknesses from `web/src/content/guides/shaman/elemental.md`; the rotation's spells and order from `data/curated/apl/shaman-elemental.json`; talent names, ranks and text from the build's `talents/shaman.json`. The generator fails if the loop names a spell the board omits.
- **The tier board does not use Elemental Shaman.** The band file for Elemental (`bis/shaman-elemental.json`, 60, raid, Horde, build 70658) puts it 20th of 20 DPS specs, 51.0% behind Fury Warrior. That number is the engine's, and it is not believed (see the F9 accuracy gate). The tier board's sample player is a Fire Mage (6th of 20, 9.1% behind). Every other board keeps Lightningbolltz.
- **Prices are the client's.** Vendor lines print `BuyPrice` from the build's `raw/hotfixes/ItemSparse.csv` (`Vendor: Meilosh · 8g 63s`). The standing a vendor asks for is not in our data (`required_standing_raw` is null for Meilosh and Pix Xizzix items, `MinReputation` is 0), so the boards print none. The board never guesses a standing.
- **Spritecaster Cape's Binds word is the mock's.** The item is in no `ItemSparse` row of our build, base or hotfix, so `Binds when equipped` is drawn from spec C6, not from data. `Cloth` is the item's `Item.csv` subclass.
- **Two timelines.** F10, F11 and F16.3 are Thursday 17 December. F14, F15 and F16.4 to F16.7 are the evening of Thursday 10 December, the first raid night. First raids open Wednesday 9 December (`dates.json`); no Onyxia date is earlier. Dungeon runs (Blackrock Depths, Lower Blackrock Spire) are five players. F16.3 is drawn on Thursday 17 December and F16.4 to F16.7 on Tuesday 8 December.
- **F12 and the Home upgrades come from two builds.** F12.4 and F12.5 read build 70658, where Spritecaster Cape is Elemental's Back pick. The Home, widget and Siri upgrade rows read the v1 data (build 70532), where Arcanoweave Cloak is the Back pick. The lanes read one build; the captures must not show both answers for one slot.
- **One character, one number.** Lightningbolltz kills Onyxia at 398 DPS, 14th of 212 Elemental Shamans, the 93rd percentile, on F10.2, F10.3, F11.1 and the meter. The sim result on F10.8 and F14.4 (412 DPS, Top Gear) is a different quantity and stays above the log. F11.3 is Sparkwick, an alt of the same account, 204th at 210 DPS; the rank card names the character in both states.

Every board is drawn on iPhone 17 (402 x 874 pt) and Pixel 9 (412 x 915 dp). `-ios` and `-and` name the platform.

| Board | What it shows |
|---|---|
| `f9-guides-{ios,and}` | F9.1 Guides tab, F9.2 a guide, F9.3 talents, F9.3b rotation, F9.4 tier lists (sample player: a Fire Mage), F9.5 what changed, F9.6 offline class page, F9.7 loading |
| `f10-reports-{ios,and}` | F10.1 your reports, F10.2 a report, F10.3 a fight's meters, F10.4 coach (Premium), F10.5 report states, F10.6 empty, F10.7 loading, F10.8 Home with the Logs and Rankings rows and the search button |
| `f11-rankings-{ios,and}` | F11.1 rankings, F11.2 filters, F11.3 off the page, F11.4 a character, F11.5 gear, F11.6 no logs, F11.7 not found |
| `f12-search-{ios,and}` | F12.1 empty search, F12.2 all results, F12.3 items, F12.4 an item, F12.5 best in slot for, F12.6 no results |
| `f13-planner-{ios,and}` | F13.1 plan a build, F13.2 a refusal, F13.3 share confirm, F13.4 system share, F13.5 a shared build, F13.6 your builds |
| `f14-surfaces-ios` | F14.1 home screen, F14.2 widget states, F14.3 lock screen, F14.4 Dynamic Island and the sim |
| `f14-surfaces-and` | F14.1 home screen, F14.2 widget states, F14.3 shade and quick-settings tile |
| `f15-system-ios` | F15.1 and F15.2 Siri, F15.3 Spotlight, F15.4 offline data, F15.5 stale and new-data states |
| `f15-system-and` | F15.1 launcher shortcuts, F15.2 Assistant, F15.3 launcher search, F15.4 offline data, F15.5 stale and new-data states |
| `f16-guild-{ios,and}` | F16.1 roster, F16.2 no match, F16.3 raid history, F16.4 calendar, F16.5 access off, F16.6 the event, F16.7 the Raids tab with the Roster and Raid history rows |

# 1. Scope: what is in, what changed, what is dropped

| # | Feature | Player question | Status |
|---|---|---|---|
| F9 | Guides, tier lists, what changed | "How do I play my spec, and where does it stand?" | In. A fifth tab |
| F10 | Your reports | "How did I do last raid?" | In. Pushed from Home |
| F11 | Rankings and character pages | "Where do I stand, and what does that player wear?" | In. Pushed from Home and Search |
| F12 | Search and the item page | "Where does this drop, and does my spec want it?" | In. A button on every root bar |
| F13 | Talent planner | "Let me build it, save it, send it." | In. Opens from Talents and from a guide |
| F14 | Widgets and live surfaces | "Am I in tonight, and what is my next upgrade, without opening the app?" | In |
| F15 | System integration | "Do it by voice, find it from the home screen, work with no signal." | In |
| F16 | Guild extras | "Who is in my guild, how did we do, and put it in my calendar." | In |

**Changes to v1 screens (the lanes make them with the feature):**
- **Tab bar.** Five tabs: Home, Raids, BiS, Guides, You. §3.
- **F3's search button** (filtered the 17 rows) becomes the universal search button (F12). F3 keeps no second search.
- **F2.7 Your gear** adopts the Paperdoll component (`design/specs/2026-10-31-paperdoll.md`), built natively once in lane F11 (tenet 10: one component, two call sites).
- **Home** gains two rows under `Next for {name}`: `Logs` and `Rankings`. §5.2, §5.3.
- **Raids tab** gains two rows under the guild head: `Roster` and `Raid history`. §5.8.
- **You** gains `Calendar` and `Offline data`.

**Dropped, with the reason:**

| Dropped | Why |
|---|---|
| Sim composing in the app | v1 ruling: the app reruns, it does not compose |
| Guild leaderboards (`GET /v1/rankings/guilds`) and public guild pages | A reference leaderboard, not a helper (tenet 14). No guild search endpoint exists. The player's own guild is F4 and F16 |
| Loot picker and loot awards | Officer planning stays on the web (v1 ruling) |
| Approving characters | v1.1, on the officer tools (v1 §17 Q1) |
| Soft reserves | Needs a raid field and a site surface first |
| Parse history charts | Premium and the web's chart component. The app shows one row, `Parse history`, that opens the web (F11.5) |
| Per-fight cast, aura and death tables | Details!-depth tables belong on a desktop. F10 shows meters and opens the web for the rest |
| "What you can win tonight" on the raid page | Checked against build 70658: no Phase 1 spec lists an Onyxia drop in its 60 raid list, so the card would be empty for every player. Revisit when Barrow Deeps and Hyjal drops are in `loot.json` |
| A second changelog reader | One list, opening the same reader as guides (F9) |
| Item share links | There is no web item page, so a link would be a dead end for non-users. Items share nothing |
| Ratings explainer page | `How is this calculated?` opens foreversixty.gg/ratings in the browser |

## 1.1 Free or Premium (decision G, applied)

Decision G: the apps are free and carry the web's exact entitlement rules. The owner asked for maximum features. This is the ruling for every Wave 3 feature. The flag names are the API's `entitlements` (`GET /v1/me`).

| Feature | Free | Premium (flag) | Why |
|---|---|---|---|
| F9 Guides, tiers, what changed | All | none | Free on the web |
| F10 Reports list, report, fights, your parses, meters | All | none | Free on the web |
| F10 Coach card | Never drawn, no prompt | `log_coach` | The web gates it; v1 F4.23 already hides it for free players |
| F11 Rankings, any character page, watching | All | none | Free on the web |
| F11 Parse history | Not shown | `history`: one row opening the web | The web gates the owner's history |
| F12 Search, item page | All | none | |
| F13 Planner, sharing, your builds | All | none | The web planner is free |
| F14 Next-raid widgets, lock-screen widgets, controls, raid Live Activity | All | none | "Raid things free": guild members never pay (decision G) |
| F14 Upgrades widget | All | none | Its data is the free Home answer |
| F14 Sim Live Activity and ongoing notification | Does not exist | `server_sims` | A server run is the thing that is Premium. The surface follows the run. A free player has none |
| F15 Siri, shortcuts, Spotlight, share, offline | All | none | |
| F16 Roster, raid history, calendar, calendar link | All | none | Raid things are free |

**There is no mobile-only Premium extra.** A paid widget or a paid Live Activity would break decision G, and a locked widget in a free player's gallery is a nag (premium-entry spec §2). If the owner later wants one, it needs an `entitlements` flag first. Today every Premium surface in the apps already exists on the web.

# 2. The one question per surface, and what is visible without scrolling

First screen means no scroll on iPhone 17 and Pixel 9 at the default text size with the tab bar drawn (v1 §1). On iPhone SE 3 (375 x 667) and 360 x 640 Android the same answer shows with only lower rows cut.

| Surface | The player's question | Visible without scrolling |
|---|---|---|
| Guides tab | "How do I play my spec?" | `Your guide` card (spec icon, `Elemental Shaman`, updated date, `Saved on this phone`), then the class grid (F9.1) |
| A guide | "What do I do?" | Role, the seven-stat priority as numbered chips, then Overview (F9.2). Talents and rotation are one chip tap away |
| Tier lists | "Where does my spec stand?" | `{Spec} {Class} is {ordinal} of {n} DPS specs.` with its gap, then the top rows (F9.4; the board shows `Fire Mage is 6th of 20 DPS specs.`) |
| Your reports | "Which raid do I look at?" | The newest report, its state and counts (F10.1) |
| A report | "How did I do?" | `Your parses`: boss, DPS, percentile, kill time (F10.2) |
| Rankings | "Where do I stand?" | `Your rank`, `14th` and `of 212 Elemental Shaman characters`, the filter line, the first rows (F11.1) |
| A character | "How good is this player, what do they wear?" | Name, class, level, guild, Forever Score with tier, the five components (F11.4) |
| Search | "Find it." | The field with the keyboard up; empty: `For {name}`, `Watching`, `Recent` (F12.1) |
| An item | "Where does it drop, does my spec want it?" | Name in quality colour, `Your Back pick` when it is, the tooltip, `Where it comes from` (F12.4) |
| Planner | "What does my build look like?" | The tree, points left, the inspector line, Share (F13.1) |
| Home screen | "Am I in tonight? What is my next upgrade?" | The widgets themselves (F14.1) |
| Roster | "Who is in my guild and who shows up?" | Counts, role chips, the first rows with attendance (F16.1) |
| Raid history | "How did the last nights go?" | Zone, date, kills and wipes, raiders, `You were there` (F16.3) |

# 3. Information architecture

## 3.1 Five tabs

v1 §3.1 rejected a fifth tab for Sims because it would be a page about Premium. A fifth tab for **Guides** has no such problem: it is free, it is about the player's class (tenet 14: the tab opens on `Your guide`), and the BiS tab has no room for it (BiS owns gear; one job per page).

| Tab | Root | Pushes |
|---|---|---|
| Home | F2.1 | v1 pushes, plus F10 Your reports, F11 Rankings, F13 Plan a build (from Talents) |
| Raids | F4.1 | v1 pushes, plus F16 Roster, F16 Raid history |
| BiS | F3.1 | none; item rows open F2.6, whose `More about this item` row opens F12.4 |
| Guides | F9.1 | a class page, a guide, a tier list, a changelog entry |
| You | F8.1 | v1 pushes, plus F16.4 Calendar, F15.4 Offline data |

Sheets from anywhere, unchanged from v1: F1.5 the switcher, F2.6 the item tooltip, F4.3 the spec ask.

**iOS.** Five `Tab`s in the iOS 26 glass bar. Labels stay 9.5 pt in the mock only because the mock bar is a picture; the real labels use the system size and do not scale (v1 §4.2).
**Android.** Five `NavigationBarItem`s. Material allows three to five.

## 3.2 Search is a button, not a tab

A magnifier sits at the trailing edge of the root bar on Home, Raids, BiS and Guides (not You). It opens F12 full screen over the tab bar. The back or Cancel control returns to where the player was.

- iOS: the glass magnifier button opens a full-screen cover with a `.searchable`-style field at the top (not the iOS 26 detached search tab: it would be a sixth control).
- Android: the 48 dp icon button opens a Material 3 `SearchBar` in full-screen mode.

## 3.3 Where the new screens live

| Screen | Entry points |
|---|---|
| F9 a guide | Guides tab; Search results; a deep link `/guides/{class}/{spec}`; Spotlight; a guide's `Open in planner` row goes the other way |
| F10 Your reports | Home `Logs` row; the push `report.ready`'s Open. A report also opens from a raid history row and a character's recent fight |
| F11 Rankings | Home `Rankings` row; Search `Rankings` quick link |
| F11 a character | A rankings row; a Search result; the roster; a deep link `/character/{region}/{ruleset}/{name}`; Spotlight |
| F12 an item | F2.6 `More about this item`; a Search result; a BiS row's tooltip sheet |
| F13 the editor | The Talents screen (F6) `Plan a build`; a guide's `Open this build in the planner`; a shared build's `Edit a copy`; a deep link `/b/{id}` opens the viewer |
| F16 roster, history | The Raids tab guild head rows |
| F16 Calendar | You › Calendar |

## 3.4 Deep links

Decision B: links are `https://foreversixty.gg/app/…`. Wave 3 adds a few real web paths to the table (`apps/shared/links.json`, which generates the association files and both routers) because a link a player pastes in Discord should open the app when it is installed and the web page when it is not:

| Path | Opens |
|---|---|
| `/b/{id}` | F13.5 the shared build |
| `/character/{region}/{ruleset}/{name}` | F11.4 |
| `/guides/{class}/{spec}` | F9.2 |
| `/tiers`, `/tiers/tank`, `/tiers/healer` | F9.4 on that role |
| `/rankings/{encounter}` | F11.1 on that encounter |
| `/reports/{id}` | F10.2, read only |

An unknown path inside these prefixes opens the system browser. A deep link is never a dead end. The player can still choose `Open in Safari` (iOS long press) or `Open in Chrome` (Android).

# 4. Design system: what is new

No new tokens. These shared pieces are new and each is a new shared component, so each needs its own small design review (tenet 10). The boards are that review's input.

| Piece | Where | Spec |
|---|---|---|
| **Paperdoll** | F11.5, F2.7 | `design/specs/2026-10-31-paperdoll.md`, phone frame: two mirrored columns, 40 pt sockets, names toward the centre in quality colour, `ilvl` micro line, a gold 10 pt dot on a slot with a better BiS item. Weapons in a row under the columns |
| **Meter row** | F10.3 | 44 pt tall, spec icon 28, class-coloured name, value in mono right-aligned. A bar behind the row at 22% of the class colour, length = value over the top value. Your row: `--row` ground and a 3 pt gold left edge. The reference is Details!, one number per row |
| **Tier chip** | F9.4 | 28 pt square, `--radius-pill`, 1.5 pt border and letter in the tier colour: S `--color-rarity-legendary`, A `--color-rarity-epic-text`, B `--color-rarity-rare-text`, C `--color-rarity-uncommon`, D `--color-muted`. The page's own mapping (`TierPage.astro`) |
| **Numbered chip** | F9.2 | 36 pt pill, a 22 pt number disc (`--tile`, `--tile-ink`, mono 12), the stat name. Wraps; never cut |
| **Parse number** | F10.2, F11.4 | Mono 700 coloured by the parse ladder: grey `#8c8c8c` under 25, green `#1eff00` 25, blue `#3d94f0` 50, purple `#b866f5` 75, orange `#ff8000` 95, pink `#f08cc0` 99. The text variants of `--color-parse-*`; the rankings page uses the same floors |
| **Stepper strip** | F13.1 | 96 pt tall bar over the action bar: talent icon 44, name and `{rank}/{max}`, the rank text, a 48 pt minus and plus |
| **Home-screen widget** | F14 | A container: `--surface-card` gradient, 17% gold border, radius 22 (iOS) or 28 (Android). Its internals use only the type roles of v1 §4.2 at widget sizes |
| **Stale strip** | F15.5 | The v1 offline strip, with a second text and a trailing text button. §5.7 |

The tier colours, parse colours and quality colours never mean anything but themselves (tenet 10). Gold marks links, the countdown and the one primary action; kill green marks a good state.

# 5. Features

Rules for every Wave 3 screen are v1 §5.0. In short: a refresh never replaces content with skeletons; offline shows cached content under the strip; an error shows in the failing region only, in `--color-wipe`, with `Try again`; writes are never queued offline. Each feature below lists: the screens and states, the copy (verbatim), the interactions, the data, and the acceptance captures. Contract changes are `W3-C{n}` and are written out in §10.

## 5.1 F9. Guides, tier lists, what changed

**Reference.** The web's guide pages (`web/src/pages/guides/[class]/[spec].astro`) for content and section order; Wowhead's class guides for the sticky section strip; the web's tier list (`TierPage.astro`) for rows, tier letters and the callout.

### F9.1 Guides tab (board F9.1)
- Root bar: the selector and the search button, as Home, Raids and BiS (v1 §3.4). Title `Guides` (`title`).
- Segments (`seg_n`, 3): `Guides`, `Tier lists`, `What changed`. The segment persists per device.
- **`Your guide` card** (the one `--surface-card` on the screen): spec icon 56, micro label `Your guide` (gold), the spec in the class colour (`sec` 24), meta `{Role} · Updated {d Mon yyyy}`, a trailing chevron, and under it the tag `Saved on this phone` (kill) when the guide is cached. The whole card is one 80 pt target. With no character: the card is absent and the class grid leads.
- **`All classes`**: one quiet card, 3 x 3 grid of ringed crests (44) with the class name in its colour (`meta` 600). A tap opens the class page (F9.6). Nine cells, always.
- Footer meta: `Saved on this phone: {classes} · {size}. Guides update by themselves when the game data does.` With nothing saved: `Nothing is saved on this phone yet. Open a guide once and it stays.`

### F9.2 A guide (board F9.2, F9.3)
- Bar: back (`Guides`), share. The share link is `https://foreversixty.gg/guides/{class}/{spec}`. Title is the spec in the class colour.
- Head band (`--surface-band`): spec icon 56, `{Spec} {Class}` (`h2`, class colour), meta `Updated {d Mon yyyy} · Recommended races: {race}, {race}`.
- Section strip: horizontally scrolling chips, `Overview`, `Talents`, `Rotation`, `Stats`, `Gear`, `Races` (only the sections the guide has). 36 pt pills; the current section is gold. A tap scrolls to the section; scrolling updates the chip. The strip sticks under the bar.
- **At a glance card** (the one `--surface-card`): `Role` and its line; `Stat priority` as numbered chips in the guide's order. This is the answer: it is visible first.
- **Overview** (quiet card): the guide's first paragraph, then `Strengths` and `Weaknesses` lists.
- **Talents** (F9.3): the key talents, each a row with the talent icon (40), name, `{n} ranks` (mono, `1 rank` for one), and the game's own text for its top rank (`talents/{class}.json`). The key talents are the guide's named ones plus every talent of the raid build the spec casts through (Elemental adds `Elemental Mastery`). Then a `Raid build` row with the tree totals in mono (`30 / 3 / 18`, summed from the build) and `Level 60`, with `Compare` (gold, opens F6 compare against your build). The last row, `Open this build in the planner`, opens F13 with the build loaded. The guide never restates the planner's tree (one job per page: link, never restate).
- **Rotation** (F9.3): numbered rows. Each has the number (mono, muted), the spell icon (40), the spell name, `Rank {n}` (mono, the top rank available at level 60 from `spellranks.json`), and the condition in plain words (`Two or more targets.`). The rows and their order come from the guide's own rotation data, never typed per screen. A spell the guide says is not in the loop still has a row, with the guide's reason (Elemental: `Flame Shock`, `Not in this site's loop. In our sim it costs more mana than it adds.`). A duration is printed only when the client's spell data and the guide agree; the Elemental sample prints none for Mana Spring Totem (`Down first. It returns mana for the rest of the fight.`) because the client's spell data reads 300 s while Era players know 60 s, and the accuracy reviewer has not settled which Forever ships. Tenet 3 is the rule: an ability is never just a name. Footer meta `Ranks are the ones you have at 60.` followed by a one-line note the guide itself carries (the mock shows `Lava Burst is new in Forever.`).
- **Stats**: the stat priority again with the hit-cap line when the guide has one. No weights (the sims own numbers; one job per page).
- **Gear**: a link row `Best in slot for {spec} · 60 · Phase 1 raid` that switches to the BiS tab with that list. No item list here (BiS owns gear).
- **Races**: rows of race, faction mark (24), and the guide's line.
- **Sources**: one quiet disclosure at the foot, `Sources`, listing the guide's sources. Provenance stays quiet: no source pills anywhere else in a guide. One `Updated` stamp, in the head.
- Prose never quotes nightly numbers (house rule). Numbers on this page are rendered from data at build time: tree totals, spell ranks, ranks of a talent.

### F9.4 Tier lists (board F9.4)
- Segment `Tier lists`. Under it: role chips `DPS`, `Tank`, `Healer` (a chip is a 40 pt pill; the role is one value, so a chip opens a menu on a second tap) and the faction chip with its mark. The role and faction default to the current character's.
- **Callout card** (`--surface-card`): spec icon 40, micro `Your spec` (gold), `{Spec} {Class} is {ordinal} of {n} {Role} specs.` and a detail line `{x}% behind {Spec} {Class}.` or `Level with the top spec.` or, for a tank, `{x}% more damage taken than {Spec} {Class}.` These are the web's strings (`tier-callout.ts`).
- Meta under the callout: `Ranked at best in slot with the Phase 1 raid buffs. Each spec is simmed as its best Horde race, named under it.` (the web's note for the role, from `tier-copy.ts`; the preset label and faction are filled from the file).
- **Rows** (quiet card, 68 min): rank (mono, muted), spec icon 36 (a spec whose icon is missing draws its class crest, never a letter tile), `{Spec} {Class}` in the class colour, the race, an 8 pt bar (fraction of the top, in the tier colour). Right column: the tier chip and the gap (`−2.8%`, `Top` on the first row, mono; the percent sign is its unit). A tie is marked `≈` before the gap. Your spec's row has the `--row` ground, the 3 pt gold edge and the tag `Yours`.
- Tank and healer lists carry no letters (the web's rule); their right column is the metric in mono.
- A tap on a row opens that spec's guide.
- Order, gaps and letters come from `tiers.json` (W3-C1), generated by the web's own `rankRole`. Neither app sorts, gaps or letters anything.
- **Accuracy gate (tenet 8): lane F9 does not ship `tiers.json`, and the app does not show a tier list, until the accuracy reviewer has verified every DPS band in it.** Open items on build 70658: Elemental Shaman 20th, 51.0% behind Fury Warrior (the sim's loop leaves out Flame Shock, casts Lava Burst without its Flame Shock bonus, runs out of mana, and Earth Shock is a mana dump: all modelling choices recorded in `data/curated/apl/shaman-elemental.json`); Frost Mage 44.1% behind; Balance Druid 47.8% behind. Until verified, `Tier lists` is hidden behind a server flag and F9.4 is not part of acceptance.

### F9.5 What changed (board F9.5)
- Segment `What changed`. Section `Key dates`: rows of a date block (`dateblk`, 56 x 48, the date from `key`), the label, the note, and, while a date is in progress, its `running` sentence in gold. Past dates are omitted. Sorted by `iso`.
- Section `What changed`: the changelog newest first. Row: date `10/28` (mono, muted), the title, a `Site` or `Game` tag (muted). A tap opens the entry in the guide reader (blocks, no new component).

### F9.6 A class page, and offline (board F9.6)
- Pushed from a class cell. Title the class in its colour. One quiet card of spec rows: spec icon 40, name, `{Role} · Updated {d Mon}`, trailing `Saved` (kill tag with a check) or `Not saved` (muted, cloud-down glyph 18), chevron.
- Offline with a saved guide: opens it under the offline strip. Offline with an unsaved guide: the screen says `This guide is not on your phone yet. Connect once to save it.`
- Loading (F9.7): skeletons at the real heights; the segments stay live.

### Sync with the nightly data
- On launch, on foreground after an hour, and in a daily background refresh, the app reads `/data/active-build.json`. A new build downloads the changed files for the classes the player has characters in, and for any class they have opened. Other classes download when opened.
- The new files are written beside the old and swapped when complete. Nothing on screen changes under the reader: while a screen is open, a new build shows the **new-data banner** (§5.7) and applies on `Update` or the next open.
- Old builds are deleted 7 days after the swap.
- Budgets: a guide at most 60 KB gzipped, `tiers.json` at most 40 KB, the guide index at most 10 KB. A build step fails over them.

### Interactions
- Segment change: `.selection` haptic (v1 §7). Section chip: selection. Tier row press: row `pressed` state.
- Share: the system share sheet with the guide link. Same for a tier list (`/tiers`).

### Copy (verbatim)
| Key | Text |
|---|---|
| `guides.title` | `Guides` |
| `guides.segments` | `Guides`, `Tier lists`, `What changed` |
| `guides.yours` | `Your guide` |
| `guides.saved` | `Saved on this phone` |
| `guides.notSaved` | `Not saved` |
| `guides.classes` | `All classes` |
| `guides.footer` | `Saved on this phone: {classes} · {size}. Guides update by themselves when the game data does.` |
| `guides.footerNone` | `Nothing is saved on this phone yet. Open a guide once and it stays.` |
| `guides.notCached` | `This guide is not on your phone yet. Connect once to save it.` |
| `guides.error` | `This guide did not load.` |
| `guides.sources` | `Sources` |
| `tiers.yours` | `Your spec` |
| `tiers.top` | `Top` |
| `tiers.loading` | `Loading the tier list…` |
| `tiers.error` | `The tier list did not load.` |
| `tiers.noChar` | the callout is absent |
| `changes.dates` | `Key dates` |
| `changes.list` | `What changed` |
| `changes.error` | `What changed did not load.` |

### Data
- **Existing:** `GET /data/active-build.json`, `GET /data/dates.json`, `GET /data/{build}/classes.json`, `GET /data/{build}/talents/{class}.json`, icons, and the BiS files (for the BiS link and the band label).
- **Contract change W3-C1 (static files, no auth, no migration):** `/data/{build}/guides/index.json`, `/data/{build}/guides/{class}-{spec}.json`, `/data/{build}/tiers.json`, `/data/changelog.json`. The site's `sync-data.mjs` emits them from `web/src/content/guides`, `web/src/content/changelog` and `lib/tiers`.
- **Lane F9 builds W3-C1 and the Guides tab, both apps.**

### Acceptance captures
F9.1 with and without a character; F9.2 at the top and scrolled to Rotation; F9.4 for DPS, Tank and Healer and for a spec that is first; F9.5; F9.6 offline with one unsaved guide; F9.7; the new-data banner over a guide; the largest text size on F9.2 and F9.4.

## 5.2 F10. Your reports

**Reference.** The web's `/reports` page and report page (`lib/reports`, `lib/report/fights.ts`); Warcraft Logs' collapsed trash; Details! for the meters. The v1 contract said the app shows the summary and opens the web for tables. Wave 3 keeps that for casts and deaths and adds meters, which fit a phone.

### F10.1 Your reports (board F10.1)
- Pushed from the Home `Logs` row. Large title `Your reports`; trailing refresh button.
- Sections `This week` and `Earlier` (micro labels). One quiet card per section.
- **Row** (68 min): the logging character's ringed crest 36, the zone (`body` 600), meta `{d Mon} · {n} boss(es) · {n} kill(s)`, a state tag, a visibility glyph 16 (muted), a chevron. `title` is used when the report has one, the zone otherwise.
- State tags: `Live` (kill, with a 7 pt dot) when `status=live` and the heartbeat is recent; meta becomes `Updated {n} s ago`. `Processing` (gold). `Failed` (wipe); meta `{d Mon} · could not be read`.
- Visibility glyphs, with their VoiceOver words: public `globe` (`Public`), unlisted `link` (`Unlisted`), guild `users` (`Guild`), private `lock` (`Private`).
- Footer meta: `Newest first. Logs come from the companion on your PC; the app reads them, it does not upload.`
- Paging: 100 a page (`page`); loading the next page shows a 56 pt spinner row at the bottom.
- **Empty (F10.6):** the web's sentence, `No reports yet. Upload a log, or let the companion send your next raid.`, and one secondary button `Get the companion` (opens foreversixty.gg/setup).
- **Loading (F10.7):** four skeleton rows. **Error:** `Your reports did not load.` + `Try again`. **Signed out:** `Sign in to see your reports.` + `Sign in with Battle.net`.

### F10.2 A report (board F10.2)
- Bar: back (`Reports`), trailing `Open on foreversixty.gg`. Title is the zone. Meta line: `{Weekday} {d Month yyyy} · logged by {character}`, and a tag with the visibility word.
- **`Your parses` card** (the one `--surface-card`): for each boss kill by one of your characters in the report, a 72 pt row: ringed crest 36, boss name (700), meta `Kill {m:ss} · {name}, {Spec}`, and on the right the figure (`{value} DPS`, mono 20) with the parse number and the word `percentile` under it. A new best adds the tag `New best` (gold). Parses count kills only. If none of your characters were in a kill: the card says `No kills of yours in this report.`
- **`Fights` card** (quiet), label with the right-aligned count `{n} boss pulls` (`1 boss pull`, `No boss pulls`; trash is never counted):
  - a boss pull: outcome dot (kill green, wipe red), boss name, meta `Kill · {m:ss} · {n} players · {n} death(s)` or `Wipe · {m:ss} · boss at {n}%`;
  - **trash is one entry per zone**, after that zone's boss pulls, the way the web folds it: dot muted, `Trash · {zone}`, meta `{n} pulls · {m:ss}` (the summed length). A tap opens the combined view. A zone with no zone name keeps its pulls as separate rows;
  - a tap on a pull opens the fight sheet (F10.3).
  - a boss pulled more than once in the report is numbered per boss in fight order, `{Boss} · pull {n} of {m}` (the web's `pullNumbers`), on its row and in the fight sheet's head; a boss pulled once keeps its plain name. (Added at the build review.)
- **Coach card** (Premium only, `log_coach`): see F10.4.
- Footer meta: `Meters, casts and deaths for every fight are on foreversixty.gg. Tap Open on foreversixty.gg.`
- **Live:** a `Live` report polls `GET /v1/reports/{id}` every 20 s while the screen is open. The fight in progress shows as a row `In progress` and joins the list when it ends. No live meters.

### F10.3 A fight (board F10.3)
- A large-detent sheet over the report. Title `Damage done`; head row: outcome dot, boss name, meta `Kill · {m:ss} · {n} players`.
- Segments: `Damage`, `Healing`, `Damage taken`. The title follows (`Damage done`, `Healing done`, `Damage taken`).
- Meter rows as in §4, sorted high to low (damage taken sorts high to low as well). Each row shows one number in mono with its unit beside it in `micro` (`DPS`, `HPS`, `DTPS`): the per-second figure, the way Details! lists it. Your characters' rows are marked. The sheet shows 11 rows, then `Show all {n}`.
- A tap on a name opens that player's character page (F11.4) when the report is public or guild-visible and the player has a character page; otherwise nothing opens.
- Trash combined view: the same sheet with the title `Trash · {zone}` and no parse numbers.
- States: loading is 8 skeleton meter rows; failed is `The meters did not load.` + `Try again`; the flagged report is `This report cannot be shown.`.

### F10.4 Coach (Premium) (board F10.4)
- Shown only when `entitlements.log_coach` and the viewer has a coach note for the night. A free player sees **no card and no prompt** (v1 F4.23, anti-nag).
- Card: title `Your {Spec} coach` (`sec`) and the tag `Premium`; sections `Fix first`, `Practice`, `Then`, `Doing well`, `Couldn't check` (only those present); footer `Reviewed {d Mon} · Based on: {labels}.`
- Busy, preparing and error words are the coach's own (`lib/coach/copy.ts`): `Your coach is writing…`, `Still writing. Check back in a few minutes.`, `The coach couldn't read this pull just now.`
- The mock's coach text is an invented example. Real notes are server-written; the app prints them as given and never edits them.

### F10.5 Report states (board F10.5)
| State | Copy |
|---|---|
| Live | `Live. Updated {n} s ago.` and `The fight in progress shows when it ends.` |
| Processing | `This report is still being read.` and `It opens here when it is ready. You do not need to stay.` |
| Failed | `This report could not be read.` and `Upload the log again from the companion.` |
| Not yours to see (404) | `This report is not available to you.` and `It is private, or it was removed.` |
| Offline | the v1 strip; a report opened before shows from cache |

### Interactions
- Pull to refresh on the list and the report. Row press: `pressed` ground. Segment change: selection haptic. Open on foreversixty.gg: the system browser.
- Home gains the row `Logs`: icon tile, label `Logs`, sub `Your last raid: {zone}, {n} kill(s)` or `No reports yet.`, chevron.
- `raid.summary` (v1) still opens F4.22 Your raid night. F4.22 gets a row `Open the full report` that opens F10.2.

### Data
- **Existing:** `GET /v1/reports?mine=1`, `GET /v1/reports/{id}`, `GET /v1/characters/{region}/{ruleset}/{name}` (the `history[]` rows carry `report_id`, `fight_index`, `percentile`, for the parse figures of the viewer's own characters), `GET /v1/me/raid-nights/{raid}`, `GET /v1/rankings/percentile`.
- **Contract change W3-C2:** `GET /v1/reports/{id}/me`, the viewer's results in one call, so the app never joins history rows to fights itself.
- **Contract change W3-C3:** `GET /v1/reports/{id}/fights/{n}/meters`, the roster rows of one fight (name, class, spec, role, active time, damage, healing, damage taken, per-second figures, deaths).
- **Contract change W3-C4:** the existing coach routes become documented app operations: `GET /v1/reports/{id}/coach/night`, `GET /v1/me/coach/recent`. No behaviour change; fixtures and `426` added.
- **Lane F10 builds W3-C2 to W3-C4 and both apps.**

### Acceptance captures
F10.1 with all four states and a live report; F10.2 with a kill, with only wipes, and with trash in two zones; F10.3 on each segment and the trash view; F10.4 as Premium and the same report as a free player (no card); F10.5; F10.6; F10.7; a report opened from a deep link; largest text on F10.2 and F10.3.

## 5.3 F11. Rankings and character pages

**Reference.** The web's `/rankings` (`design/specs/2026-10-11-rankings.md`) for the answer, the pin, the filters and the colours; the web's character page (`lib/character-profile.ts`) for the score card and sections; the in-game character frame for the gear (Paperdoll).

### F11.1 Rankings (board F11.1)
- Pushed from the Home `Rankings` row (label `Rankings`, sub `{ordinal} of {M} {Spec} {Class} characters on {encounter}`, or `Not ranked yet.`). Title `Rankings`.
- **Answer card** (`--surface-card`): micro `Your rank`, the ordinal at `stat` size, the unit `of {M} {Spec} {Class} characters` (the API dedupes to each character's best kill, so it says characters; if `deduped` is false it says `kills`), meta `Your best kill on {encounter}, {ruleset}, {phase}.`
- **Filter line:** `{Encounter} · {Ruleset} · {Region} · {Phase} · {Measure}`, wrapping, never cut, and a 44 pt `Filters` secondary button. This is the web's phone rule.
- **Board:** a quiet card of rows (68 min): rank numeral in the parse-ladder colour (mono 700), character in the class colour (`You` tag on yours), meta `{guild} · {size} · {m:ss} · {n} d ago`, the value (mono) with `DPS`, `HPS`, `DTPS` or `Executed` under it, a chevron. A single-class board has no crest per row; a mixed board adds a 28 pt ringed crest.
- **Your row:** when it is on a loaded page, the same gold edge and `--row` ground mark it in place, and a text link `Jump to your row, {ordinal}` sits above the board. When it is not (F11.3), it is pinned as the first row under the label `Your best kill`, with its true rank, then a 1 pt rule, then the board from #1.
- Paging: 100 a page, loading on scroll.
- **Empty (F11.1 empty):** the answer says `Not ranked yet on {encounter}.` and meta `A logged kill puts you on the board.`; the board says `No ranked fights yet for this filter. Rankings fill in as reports are uploaded.`
- **Signed out:** no answer card; the filter line starts with `Choose a class and spec`, which opens the F3.2 picker. The board follows.
- A row tap opens the character page; a long press opens a menu `Watch`, `Open the report`, `Share`.

### F11.2 Filters (board F11.2)
A sheet (large detent), live like the BiS picker (every choice applies at once):
- `Encounter`: a row that opens a list of encounters with ranked kills.
- `Measure`: chips `Damage`, `Healing`, `Damage taken`, `Execution`.
- `Ruleset`: `Normal`, `PvP`, `RP`, `Hardcore`, `Every ruleset`.
- `Region`: `All`, `US`, `EU`, `KR`, `TW`, `CN`.
- `Phase`: the phases from `GET /v1/phases`, defaulting to the current phase.
- `Faction`: `Both`, `Alliance`, `Horde`.
- `Today only`: a switch.
- `Reset filters` (secondary).
There is no class or spec filter when signed in: the character sets it (the web's rule). Realm does not exist in the data; the filter is Region.

### F11.4 A character (board F11.4)
- Pushed. Bar: back, a `Watch` star toggle (44 pt), share. Share link `/character/{region}/{ruleset}/{name}`.
- Head: ringed crest 64, name in the class colour (`h2`), `{Race} {Spec} {Class} · Level {n}`, `{Guild} · {Region} {Ruleset}`. A `You` tag on your own characters; the supporter mark exactly as the API resolves it (decision U).
- **Forever Score card** (`--surface-card`): micro `Forever Score`, the score in mono 36 with the tier tag in the quality colour of the tier, `As of {d Mon}`, a sparkline of `score_history[]` (oldest first), the five components as 8 pt bars with their values (`Gear`, `Experience`, `Parses`, `Attendance`, `Discipline`; each bar's length is value over that component's cap from the API), and the text link `How is this calculated?` (opens foreversixty.gg/ratings). **A character with no rated fight shows `No logs yet.`, never a zero** (F11.6).
- **`Bests`** (quiet card): per encounter, the boss, `{ruleset} · {age}`, the value (mono, with unit) and the parse number with `percentile`.
- **`Recent fights`:** kills and wipes with the value; a tap opens the report read-only (F10.2 without `Your parses` unless the character is yours).
- No logs: `No logs yet.` and `{name} has no rated fight in a public report, so there is no score to show.`
- Not found (F11.7): `We have no character called {name}.` and `Check the spelling, the region and the ruleset.` with `Search again`.

### F11.5 Gear (board F11.5)
- `Gear` card: the Paperdoll (§4), with a gold dot where the character's BiS list has a better item, and `Gold dot: the BiS list has a better item for this slot.`
- Under it, one row `Gear gap` with the points in mono. A tap on a socket opens the F2.6 tooltip sheet.
- A character with no gear export: `{name} has not sent gear to Forever Sixty.`
- `Builds seen`: rows with the tree totals in mono, `{Spec} · first seen {d Mon}`, and `Open in planner` (gold) which opens F13 with that build.
- **Parse history:** for your own character with `entitlements.history`, one row `Parse history` that opens the web. Nothing for anyone else.

### Interactions
- **Watching.** The star adds the character to the account's watch list (server-synced, up to 50). Watching shows in Search (F12.1) and in Spotlight (F15). `Watch` and `Watching` are the labels; the star's VoiceOver label is `Watch {name}` and `Stop watching {name}`. Haptic: `.success` on watch.
- Pull to refresh; share; long-press a rankings row.
- The character page of a character whose account is anonymous answers 404 and shows the not-found state.

### Copy (verbatim)
| Key | Text |
|---|---|
| `rank.title` | `Rankings` |
| `rank.label` | `Your rank` |
| `rank.unit` | `of {M} {Spec} {Class} characters` |
| `rank.unitKills` | `of {M} {Spec} {Class} kills` |
| `rank.meta` | `Your best kill on {encounter}, {ruleset}, {phase}.` |
| `rank.none` | `Not ranked yet on {encounter}.` |
| `rank.noneMeta` | `A logged kill puts you on the board.` |
| `rank.empty` | `No ranked fights yet for this filter. Rankings fill in as reports are uploaded.` |
| `rank.pin` | `Your best kill` |
| `rank.jump` | `Jump to your row, {ordinal}` |
| `rank.filters` | `Filters` |
| `rank.reset` | `Reset filters` |
| `rank.error` | `The rankings did not load.` |
| `char.score` | `Forever Score` |
| `char.how` | `How is this calculated?` |
| `char.bests` | `Bests` |
| `char.recent` | `Recent fights` |
| `char.gear` | `Gear` |
| `char.gap` | `Gear gap` |
| `char.builds` | `Builds seen` |
| `char.noLogs` | `No logs yet.` |
| `char.noScore` | `{name} has no rated fight in a public report, so there is no score to show.` |
| `char.noGear` | `{name} has not sent gear to Forever Sixty.` |
| `char.notFound` | `We have no character called {name}.` |
| `char.notFoundHelp` | `Check the spelling, the region and the ruleset.` |
| `char.watch` / `char.watching` | `Watch` / `Watching` |

### Data
- **Existing:** `GET /v1/rankings` (with `distinct=character` and `character=`), `GET /v1/rankings/percentile`, `GET /v1/phases`, `GET /v1/characters/{region}/{ruleset}/{name}`, `GET /v1/characters/search`, `GET /v1/me` (entitlements).
- **Contract change W3-C5:** the existing `/v1/rankings` operations become documented app operations (X-App headers, `426`, fixtures). No behaviour change.
- **Contract change W3-C6:** `GET /v1/rankings/encounters`, the encounters that have ranked kills (id, slug, name, zone, size, kills).
- **Contract change W3-C7:** `GET /v1/me/watching`, `PUT` and `DELETE /v1/me/watching/{region}/{ruleset}/{name}`; a migration adds `watched_characters`.
- **Lane F11 builds W3-C5 to W3-C7, the Paperdoll (and F2.7's adoption), and both apps.**

### Acceptance captures
F11.1 with your row on the page, off the page (F11.3), unranked, signed out, and offline; F11.2; F11.4 for your own character and another; F11.5 with a gear export and without; F11.6; F11.7; the watch toggle on and off; largest text on F11.1 and F11.4.

## 5.4 F12. Search and the item page

**Reference.** Wowhead's item page (tooltip, `Dropped by`, `Used in`), the in-game tooltip, and iOS and Android system search. One field finds items, characters, guides and talents.

### F12.1 The field and the empty state (board F12.1)
- Full screen over the tab bar. iOS: a glass field (44 pt) with a magnifier, placeholder `Search items, characters, guides, talents`, a clear button, and `Cancel` (gold). Android: a 56 dp `SearchBar` pill with back arrow, the same placeholder and a clear button. The keyboard opens at once.
- **Scope chips** under the field: `All`, `Items`, `Characters`, `Guides`, `Talents` (36 pt; one is on). A scope narrows the results and lifts the group's cap.
- **Empty field** (no query): three groups, each a quiet card:
  - `For {name}`: `Best in slot for {name}`, `{Spec} {Class} guide`, `Rankings` (icon tile 32, label, chevron);
  - `Watching`: the account's watched characters (F11), crest 36, name, `{Spec} {Class} · {level} · {guild}`, a filled star;
  - `Recent`: the last 20 things opened on this phone, with the label's trailing `Clear`.
  A group with nothing in it is not drawn.

### F12.2 Results (boards F12.2, F12.3)
- Typing filters at once (no debounce) for items, guides and talents, which are searched on the phone. Characters are asked of the API after 250 ms and at least 2 characters, up to 8, in the current character's region and ruleset first.
- **All** shows up to 5 per group, in this order: `Guides`, `Talents`, `Items`, `Characters`. Each group's last row is `Show all {n}` when there are more. Order inside a group: exact name, then prefix, then word prefix, then contains. Items rank your class's usable items before the rest, then quality, then item level.
- Row shapes:
  - **Item:** icon 40, name in quality colour, meta `{Slot} · item level {n} · requires level {n}`.
  - **Character:** ringed crest 36, name in the class colour (`You` tag on yours), meta `{Class} · Level {n} · {Region} {Ruleset}`, a star (44 pt target) that watches.
  - **Guide:** spec icon 36, `{Spec} {Class}` in the class colour, meta `Guide · {Role}`.
  - **Talent:** talent icon 36, name, meta `{Spec} {Class} · {Tree} · {n} ranks`. A tap opens the talent sheet (v1 F6 rank sheet) with `Open in planner`.
- A tap on an item opens F12.4; on a character F11.4; on a guide F9.2.
- **No results (F12.6):** `Nothing matches “{query}”.` and `Check the spelling. Characters need at least two letters and a connection.` Offline, characters cannot be asked: a strip inside the results says `Offline. Items, guides and talents still search; characters need a connection.`
- Opening a result adds it to Recent. Sign out clears Recent and Watching from the phone.

### F12.4 An item (boards F12.4, F12.5)
- Pushed from Search, from F2.6's `More about this item` row, and from a BiS row's sheet. Bar: back. No share (§1, no web item page).
- Head: item icon 64 with its quality border, the name in quality colour (`h2`, not uppercased), meta `{Quality} · {Slot} · item level {n}`.
- **Your pick card** (`--surface-card`, shown only when this item is on the current character's list): the character's crest, `Your {Slot} pick` or `Runner-up {n} for {Slot}`, meta `{Spec} · {band} · {preset} · {Faction}`, tag `Yours`. An item that is not on the list shows no card and no `not on your list` line.
- **Tooltip panel:** the F2.6 tooltip body exactly (name, `Binds`, `Unique`, slot and armor type, armor, stats, set bonus lines, `Requires Level`, item level in the client's gold). A cloak reads `Back` with `Cloth` on the right, as the Classic client prints it. **The Binds line is a ship blocker for F12 and F2.6:** the capture must show `Binds when picked up`, `Binds when equipped` or `Binds when used` from the client's `ItemSparse.Bonding`. The data dependency is C6: `bonding`, `unique`, `armor_type` and `min_standing` in `items/{class}.json`, plus `buy_price` from `ItemSparse.BuyPrice`. On build 70658 some items (Spritecaster Cape, 11623) are in no ItemSparse row; the data lane must source them (the base table, not the hotfix file) before the lane ships. An item with no `bonding` shows no line and fails the lane's data test.
- **`Where it comes from`:** one row per source from `loot.json`, never a dump of mobs (tenet 7): tile icon, `{boss}` or `{vendor}` or `{quest}`, meta `{zone} · {kind}`. A world drop says `World drop`. A drop chance prints only when the data has one (`{n}% from {boss}`); an unknown chance prints `Drop chance not known`, never `0%` and never a blank. A vendor source prints the price (`Vendor: Meilosh · 8g 63s`) and, when `min_standing` is in the data, the standing (`Timbermaw Hold, Revered`); the same line is used on the Siri answer and the widgets' upgrade source.
- **`Best in slot for`** (the third section; strip chips `Item`, `Source`, `Best in slot for` as on a guide): one row per spec that lists the item, the spec's best placement first: spec icon 36, `{Spec} {Class}` in the class colour, meta `{band} · Phase 1 raid` or `{band}`, a tag `Pick` (kill) or `Runner-up {n}` (muted), a chevron that opens that list in the BiS tab. Rows are ordered Pick before Runner-up, then class, then spec. Six rows show, then `{n} more specs`.
- Set pieces list the set and its bonuses under the tooltip, from `sets.json`.
- Loading: tile at the real size with the quality border, the name from the index, the rest skeleton. Unknown id (an old link): `This item is not in our data.` (the web's `unknown-item-copy`).

### Data
- **Existing:** `GET /v1/characters/search`, `GET /data/{build}/items/{class}.json`, `GET /data/{build}/sets.json`, `GET /data/{build}/bis/{spec}.json`, `GET /data/{build}/talents/{class}.json`, `GET /v1/me/watching` (F11).
- **Contract change W3-C8 (static files):** `/data/{build}/loot.json` (already synced; now documented, with fixtures), `/data/{build}/search.json` (names, quality, slot, item level, usable classes and icon for every item the class files carry, plus talents and guides; at most 300 KB gzipped, a build step fails over it) and `/data/{build}/item-uses/{shard}.json` (for each item: its source lines in the BiS formatter's words and the specs that list it, sharded by id so a phone fetches one shard on demand).
- **Lane F12 builds W3-C8, the static data store both F9 and F12 use (download, swap, delete, size accounting), and both apps.**

### Acceptance captures
F12.1 with and without watching and recents; F12.2 for a query with all four groups; F12.3 items scope; F12.4 for an item that is your pick, one that is a runner-up and one that is not on your list; F12.5; F12.6 online and offline; an old item id; largest text on F12.2 and F12.4.

## 5.5 F13. The talent planner

**Reference.** The web planner (`/planner`, `lib/planner`) and the in-game talent frame. The planner owns talents (one job per page): there is no gear panel, no weights and no band control here, and compare against a guide stays on F6.

### F13.1 Plan a build (board F13.1)
- Full screen, no tab bar (a modal task). Bar: back (`Talents`; `Cancel` on iOS when opened from a guide), an `Undo` button (48 pt), title `Plan a build` on Android.
- **Chips row:** the class crest 30 (the class is fixed by where you opened it), a `Race` chip (`Orc`; opens a sheet of the legal races for the class, from `combos.json`), a `Level` chip (`Level 60`; a sheet with a 10 to 60 stepper), and, right-aligned, the points left (`{n} left`, mono 20 gold; kill green at 0).
- **Tabs:** `{Tree} {points}` for the three trees, the totals summed from the drawn build, never typed.
- **The tree** is v1 F6's drawing: a 4-column grid, icons 48, the planner's `cellState` borders (kill green holding points, gold available, line colour locked at 40%), rank badges in mono, prerequisite lines from `connectorPath`.
- **Stepper strip** pinned over the action bar: the last touched talent's icon, name and `{rank}/{max}`, its rank text (the client's), and `−` and `+` buttons (48 pt). With no talent touched yet it shows the first talent.
- **Action bar:** `Reset` (secondary) and `Share` (primary, with the share glyph), 48 pt, one primary.
- Points: a build has `level − 9` points (51 at 60).

### Interactions
- **Tap a talent: add one point.** Haptic `.selection` (`SEGMENT_TICK`). The strip updates.
- **Long press: remove one point.** Haptic `.impact(.medium)` (`LONG_PRESS`).
- `+` and `−` in the strip do the same for the strip's talent; they are the precise path and the one for VoiceOver and TalkBack (the tree cell's accessibility actions are `Add a point` and `Remove a point`).
- **Undo** steps back through the last 100 changes. Each point is one step.
- **A refusal (F13.2)** changes nothing, plays `.error` (`REJECT`), and the strip's line goes `--color-wipe` and says why. The refusals:
  - tier: `Needs {n} points in {Tree}. You have {m}.`
  - prerequisite: `Needs {k} points in {Talent} first.`
  - no points left: `No points left. Take one off first.`
  - removal that would break a talent below: `{Talent} needs this point. Take its points off first.`
  The planner's own strings win if they differ: the planner owns them, and the app copies them from `lib/planner/copy.ts` into the app catalog.
- **Reset** asks once: `Clear all {n} points?` with `Clear` (destructive) and `Cancel`.
- **Draft.** The editor saves its state per class on the phone as you go, and `Plan a build` reopens it. A shared build is a copy: editing a shared build never changes the shared one (a link cannot be edited).
- **Rules live in one place.** Tier gating, prerequisites, the point cap and `diffCount` are implemented natively from one golden-vector file generated from `web/src/lib/planner` (`apps/shared/fixtures/planner-rules.json`). CI fails if either app disagrees with a vector. The server also validates on save (`400` with `fields` keyed `point_order[n]`).

### F13.3 and F13.4 Share
- `Share` opens the web's confirm step as a sheet, `Share this build`: `Sharing posts the following to a link anyone can open. Once created, that link cannot be edited:` then the list `Class and race`, `Talent order`, `Title: “{title}”`, a `Title (optional)` field (60 characters, with a counter), `Share anyway` (primary) and `Cancel`. Under it: `Signed in, it also appears in Your builds.`
- `Share anyway` calls `POST /v1/builds` with the class, race, `tree_version` (the active build), `point_order`, an empty `gear` and the title. `201` or `200` returns `{ id, url }`; the system share sheet opens with the URL and a preview of `/b/{id}/card.png`. Haptic `.success` (`CONFIRM`).
- Errors: `400` shows the field's sentence under the editor; `429` is `You have shared a lot of builds. Try again in an hour.`; offline is `Sharing needs a connection. Your build stays here.`

### F13.5 A shared build (board F13.5)
- A `/b/{id}` link opens the viewer. Bar: back, share. Title `Shared build`.
- Head: crest 44, `{Race} {Class} {totals}` (mono totals), meta `{title} · Saved {d Mon} · {n} views`.
- Two buttons: `Edit a copy` (primary; opens F13.1 with the points) and `Compare with mine` (secondary; shown only when you have an export for that class; it opens F6's compare with the shared build as the guide side).
- The trees are drawn read-only from `/data/{tree_version}/talents/{class}.json`: **the build's own `tree_version`, not the active build.**
- Unknown id: `This build does not exist.` Offline with no cache: `This build is not on your phone yet. Connect once to open it.`

### F13.6 Your builds (board F13.6)
- Pushed from the Talents screen. Rows: crest 36, the title or `{Race} {Class} {totals}`, meta `{Race} {Class} · {totals} · {d Mon}`, `{n} views` (mono), chevron. Below: `Plan a new build` with `{Class}, level {n}`.
- Swipe (iOS) or swipe-to-dismiss (Android) removes a row from your list. A snackbar says `Removed. Undo` for 5 s. Footer: `Remove only takes a build off your list. The link keeps working for anyone who has it.`
- Empty: `No builds yet. Share one and it shows here.` Error: `Your builds did not load.` + `Try again`.

### Copy keys
`plan.title` `Plan a build`; `plan.left` `{n} left`; `plan.share` `Share`; `plan.reset` `Reset`; `plan.undo` `Undo`; `plan.resetConfirm` `Clear all {n} points?`; `plan.shareTitle` `Share this build`; `plan.shareIntro`, `plan.shareFields`, `plan.shareProceed` `Share anyway`, `plan.rateLimited`, `plan.offline` as above; `plan.edit` `Edit a copy`; `plan.compare` `Compare with mine`; `plan.removed` `Removed. Undo`.

### Data
- **Existing:** `POST /v1/builds`, `GET /v1/builds/{id}`, `GET /v1/builds?mine=1`, `GET /b/{id}/card.png`, `GET /data/{build}/talents/{class}.json`, trees art.
- **Contract change W3-C9:** `DELETE /v1/builds/{id}` removes a build from the caller's list (session or device token; `204`; `404` if it is not theirs; the record and its link stay). If `builds.user_id` is not nullable the lane adds a migration for a nullable owner link.
- **Contract change W3-C10 (static):** `/data/{build}/races.json` and `/data/{build}/combos.json` documented with fixtures (they are already synced).
- **Contract change W3-C11:** `apps/shared/links.json` gains the rows of §3.4, regenerating the association files and the routers.
- **Lane F13 builds W3-C9 to W3-C11, the golden-vector generator, and both apps.**

### Acceptance captures
F13.1 with points left, at zero left and over a legal 51; F13.2 for each refusal; F13.3; F13.4 on each platform; F13.5 with and without an export, an old `tree_version`, and an unknown id; F13.6 with the swipe state and the snackbar; the Reset alert; the draft restored after a force quit; largest text on F13.1; VoiceOver and TalkBack adding and removing a point.

## 5.6 F14. Home-screen and lock-screen surfaces

**Reference.** Apple Calendar's and Weather's widgets for the shape; Raid-Helper's reminder for the content; Details! for restraint. A widget answers one question and opens the app on the right screen.

### Widgets (boards F14.1, F14.2)
| Widget | Sizes | Question |
|---|---|---|
| **Next raid** | iOS small, medium, `accessoryInline`, `accessoryRectangular`, `accessoryCircular`; Android 2 x 2, 4 x 2 | "Am I in tonight?" |
| **Next upgrade** | iOS small, medium, `accessoryRectangular`; Android 2 x 2, 4 x 2 | "What should I get next?" |

**Next raid, small:** the seal 16 and micro `Next raid` (gold); the zone in `sec` 22 (Barlow Condensed, uppercase); `Thu 8:00 pm`; the countdown `in 2 days 2 h` (mono, gold; whole words, never `h` alone for days); at the foot a tag with your answer: `In · DPS` (kill), `Not signed up` (gold), `Late`, `Bench`, `Out`.
**Next raid, medium:** the same on the left; on the right `24 of 40` and `signed up`, then three 44 pt buttons `In`, `Late`, `Out` with your answer checked (kill ground and a 2 pt kill underline). A press answers without opening the app (iOS interactive widget through an App Intent; Android Glance `actionRunCallback`) by calling the existing sign-up `PUT`. The widget redraws with the saved answer at once, then with the server's. A failed save redraws the previous answer and shows `Not saved` in `--color-wipe` on the button for 10 s.
**Lock screen (iOS):** inline `Onyxia · Thu 8 pm · In`; rectangular zone, `Thu 8:00 pm · in 2 d` (`Tonight 8:00 pm` on the day), `You: In · DPS`; circular a ring of signed up over size with the count in the middle. A second circular shows the time to the raid (`2d`, then `53m` on the day). On the day every surface agrees: the board is 7:07 pm, the Live Activity reads `38:00` to invites at 7:45 pm and `38 of 40`, the rings read `38 of 40` and `53 min` (a ring always carries its unit). The upgrade rectangular shows `Next upgrade · {Slot}`, the item name, `+{gain} DPS`. Lock-screen widgets are monochrome: no quality colour, no icon colour (the system tints them). Every surface computes its figures from the one snapshot with one clock through WidgetFigures, held to `apps/shared/widgets/figures.json` on both platforms, so the lock-screen widgets, the home-screen widgets and the Live Activity never disagree (wow-player design SHOULD, fixed in the build). With no invite time the Live Activity counts to the pull: `{zone} starts in {n} minutes`, `until the pull at {time}`.
**Android lock screen:** the widget declares the `keyguard` category and shows where the OS allows lock-screen widgets; elsewhere it is home screen only.
**Next upgrade, small:** crest 16 and micro `Next upgrade`; the item icon 44 with its quality border; the name in quality colour; `+5.3 DPS` (mono). **Medium:** `Upgrades · {name}`, three rows (icon 32, name in quality colour, gain in mono), and a footer `{n} upgrades · +{total} DPS`. A tap opens the item's F2.6 sheet; a tap on the header opens Home. The rows and gains come from the server's order (v1: measured gains high to low, then `Gain not modelled yet`, shown as `—`).

**Time.** Widgets, Live Activities, Siri answers and calendar events use the phone's time zone, as v1's notifications do (§6.1: text is composed in the player's time zone). Where the guild's zone differs from the phone's, the Raids tab already shows the guild's time beside it (v1 §2); a widget cannot fit both, so it shows the phone's time only and the raid page shows both.

**Configuration.** Both widgets take one setting, the character (`Current character` by default). iOS `AppIntentConfiguration` with a `CharacterEntity`; Android a configuration activity listing your characters.

**States (board F14.2), every small state:**
| State | Content |
|---|---|
| Signed out | `Sign in to see your next raid.` and `Open the app` (gold) |
| No guild | `No guild yet` and `Your raids show here once you are on a roster.` |
| No raid planned | `No raid planned` and `Last: {zone}, {d Mon}.` |
| Sign-ups locked | the raid with the tag `Locked`; no buttons |
| Cancelled | the raid with the tag `Cancelled` (wipe); no countdown |
| Older than 12 h | the normal widget with `Updated {n} h ago` in muted at the foot; buttons stay (a press asks the server) |
| Not signed up | the tag `Not signed up` (gold); medium shows no answer checked |
| No spec saved | the tag `Choose a spec`; every button opens the app's spec ask (F4.3) instead of answering |
| Upgrades: no gear export | `Add the addon to see your upgrades.` |
| Upgrades: under level 20 or no list | the Home answer's own sentence, from the server |
A widget never shows a spinner and never shows an empty box: every state has words.

**Refresh.** The app writes one snapshot to the shared container (iOS App Group; Android DataStore read by the Glance widget). The snapshot is `GET /v1/me/widget` (W3-C12). Triggers: the app refresh task about every 30 minutes, the silent `widget.refresh` push on `raid.changed`, `raid.roster_set`, a sign-up answered elsewhere, `band.entered` and a nightly BiS change, and every foreground. Countdown entries are timeline entries at 30-minute steps (hourly past 2 days), so the widget needs no network to tick.

### Live Activity and ongoing notification (boards F14.3, F14.4)
Two kinds. Both start and update from pushes (lane A2's senders; the app never polls in the background).

**Raid starting.** Starts at `invite_at` minus the guild's lead (default 60 min) for players who are In or Late. Content: the zone, `Olympus XXVII`, a countdown `38:00` in mono 36 with `until invites at 7:45 pm`, a 6 pt gold progress bar over the lead, and `You are In · DPS · Lightningbolltz` with `38 of 40`. It ends at the pull time plus 15 minutes, or at once if the raid is cancelled. Cancelled replaces the countdown with `Cancelled` for 5 minutes. Free.
- iOS Dynamic Island: compact leading the seal, trailing a ticking `38:12` timer (amended 2026-10-31, F14 build review: it fits the compact frame and is more useful than `38m`); expanded shows the zone, `Invites 7:45 pm · You are In · DPS`, the countdown, and buttons `Late` and `Out` (App Intents, which answer in place).
- Android: an ongoing notification on the channel `Raid starting` with a chronometer counting down to the invites, actions `Late` and `Out`, the title `Invites for {zone} in {n} minutes`. It uses the v1 notification small icon.

**Server sim.** Starts when a server run starts (from the app: at the press; from the web: by push-to-start). Content: `{Spec} · {kind}` (`Elemental · Top Gear`), `{n} of {m} combinations` (bulk runs; a plain run shows `Running`), a progress bar, `about {n} min left`. Updates at most every 15 s. On finish it becomes `Finished` with `412 DPS · tap for the result` and stays 1 hour. A failed run says `Failed` with `Open the run`. **It exists only for a server run, which is Premium (`server_sims`);** there is nothing to gate.
- Android: an ongoing notification on the channel `Server sims` with a determinate progress bar, replaced by the v1 `sim.finished` notification at the end (one thread per run).

**Quick-settings tile (Android).** `Next raid`, subtitle `Thu 8 pm · In`, active when a raid is planned; a tap opens the raid page. With no raid the tile is inactive and says `No raid planned`.

**Rules.** Quiet hours do not suppress a Live Activity (it is not a sound or a buzz). The player can turn each off in You › Notifications under `Live surfaces`: `Raid starting countdown` and `Server sim progress`, both on by default. A raid in the past never starts one.

### Copy (verbatim)
| Key | Text |
|---|---|
| `widget.nextRaid` | `Next raid` |
| `widget.nextUpgrade` | `Next upgrade` |
| `widget.upgrades` | `Upgrades · {name}` |
| `widget.signedOut` | `Sign in to see your next raid.` |
| `widget.openApp` | `Open the app` |
| `widget.noGuild` | `No guild yet` |
| `widget.noGuildSub` | `Your raids show here once you are on a roster.` |
| `widget.noRaid` | `No raid planned` |
| `widget.lastRaid` | `Last: {zone}, {d Mon}.` |
| `widget.locked` / `widget.cancelled` | `Locked` / `Cancelled` |
| `widget.updated` | `Updated {n} h ago` |
| `widget.chooseSpec` | `Choose a spec` |
| `widget.notSaved` | `Not saved` |
| `widget.signedUp` | `{n} of {size} signed up` |
| `widget.noGear` | `Add the addon to see your upgrades.` |
| `live.invites` | `until invites at {time}` |
| `live.you` | `You are {state} · {role} · {name}` |
| `live.notifTitle` | `Invites for {zone} in {n} minutes` |
| `live.sim` | `{n} of {m} combinations` |
| `live.simLeft` | `about {n} min left` |
| `live.finished` | `Finished` |
| `live.finishedSub` | `{dps} DPS · tap for the result` |
| `live.failed` | `Failed` |
| `live.settings` | `Live surfaces` |
| `live.raidSetting` | `Raid starting countdown` |
| `live.simSetting` | `Server sim progress` |

### Data
- **Existing:** `PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}`, `GET /v1/guilds/{id}/raids`, `GET /v1/sims/{id}/progress`, `POST /v1/me/push-tokens`, `GET /v1/me/notifications`, the lane A2 senders.
- **Contract change W3-C12:** `GET /v1/me/widget`, one small answer for both widgets: the character, the next raid (zone, `starts_at`, `invite_at`, `pull_at`, size, signed up, status, your answer, spec and role, `can_sign_up`), the top three upgrades with the totals, and `generated_at`. Session or device token. ETag, `private, no-cache`. It reuses the Home answer's and the raid view's builders; no new rule.
- **Contract change W3-C13:** `POST /v1/me/live-activities` registers a Live Activity or ongoing-notification subscription (`kind` raid or sim, `platform`, `push_token`, optional `push_to_start_token`, `object_id`) and `DELETE /v1/me/live-activities/{id}` ends it; a migration adds `live_activity_tokens`. A `sim` registration answers `402 premium_required` without `server_sims`.
- **Contract change W3-C14:** the event catalogue gains `widget.refresh` (silent, collapses per user), `live.raid.start`, `live.raid.update`, `live.raid.end`, `live.sim.update`, `live.sim.end`, produced by the scheduler and the sim runner, sent through the A2 senders.
- **Lane F14 builds W3-C12 to W3-C14, the widget and Live Activity extensions on iOS and the Glance widgets and notifications on Android.**

### Acceptance captures
F14.1 on a light and dark wallpaper, medium with each answer checked, the failure line; F14.2 every state; F14.3 each lock-screen widget with the system tint; the Live Activity at 38 minutes, at 5 minutes and cancelled; F14.4 compact and expanded Island, the sim running, finished and failed; Android F14.1, F14.2, F14.3 for the raid and the sim, and the tile; the widget gallery entry on both platforms; largest text on the medium widgets.

## 5.7 F15. System integration

**Reference.** Apple's own apps' Siri answers (Calendar, Weather), Spotlight's app entities, Android's launcher shortcuts and Assistant app actions, the system share sheet.

### Voice and Shortcuts (boards F15.1, F15.2)
**iOS App Intents**, all in the app's `AppShortcutsProvider`, so they work with no setup:

| Intent | Phrases (each ends with `in Forever Sixty`) | Answer |
|---|---|---|
| Sign up for a raid (`In`, `Late`, `Out`) | `Sign me up for tonight's raid`, `Tell Forever Sixty I'm late`, `Mark me out of the next raid` | `You are In for Onyxia's Lair, Thursday at 8 pm, as DPS on Lightningbolltz. 25 of 40 signed up.` |
| Next upgrade | `What's my next upgrade` | a snippet: item icon, name in quality colour, `{Slot} · item level {n} · {source}`, `+5.3 DPS for {name}. Replaces {worn}.` |
| Next raid | `When is the next raid` | `Onyxia's Lair, Thursday at 8 pm. You are In. 25 of 40 signed up.` |
| Who is coming | `Who is coming to raid` | `24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.` |
| Open my best in slot | `Open my best in slot list` | opens the BiS tab |

- **Which raid.** The next raid starting within 24 hours; else the next planned one. If two raids are within 24 hours Siri asks `Which raid?` and lists them. The character is the current character; a `Character` parameter lets a Shortcut choose another.
- **Confirmation by voice only.** Run by voice, a sign-up asks once (F15.1: `Sign Lightningbolltz up as In for Onyxia's Lair, Thursday 8 pm?` with `Sign up` and `Cancel`). Run from a Shortcut, the Action button, a widget or a Control it does not ask. A sign-up is reversible and one tap away, which is why one confirmation is enough.
- **The same rules as the app.** An intent calls the same `PUT`, shows the same refusal sentences (`Sign-ups are locked. An officer can still change yours.`, `That did not save. Try again.`) and asks the spec question when the character has no spec (`Which spec is Lightningbolltz on Olympus XXVII?` with the specs as options).
- **Controls (iOS 18+).** One Control Center control, `In for the next raid`, that answers `In`. It is the same intent.

**Android.**
- **Static shortcuts** (long-press the icon, F15.1): `Next raid`, `Sign up In`, `My upgrades`, `Search`. `Sign up In` runs a transparent activity that does the `PUT` and shows a toast, `You are In for Onyxia's Lair.`; it never opens the app.
- **Dynamic shortcuts** (up to 4): the current character, the first two watched characters, and the next raid. They update with the data and are removed on sign out.
- **Assistant:** the `OPEN_APP_FEATURE` built-in intent for `next raid`, `my upgrades`, `best in slot`, `guides` and `search` (F15.2). Android has no built-in intent for a raid sign-up, so the app does not pretend to offer voice sign-up on Android; `Sign up In` is the shortcut.

### Spotlight and launcher search (board F15.3)
- **iOS:** the app's characters (your own up to 9, and the first 50 watched), the planned raids (next 5, title only: `Onyxia's Lair · Thu 8 pm`), the guides saved on the phone, and the five shortcuts. They are `IndexedEntity` values, so Spotlight, Siri and Shortcuts see the same things. Items are not indexed (thousands of rows nobody asks Spotlight for). The index is on the device and is deleted on sign out.
- **Android:** dynamic shortcuts surface in launcher search on Pixel and other launchers that index them; each carries the crest, the name and `Forever Sixty · character`.
- A tap opens the same screen as in-app.

### Share (v1 §F4.5 extended)
Only things that have a working web page share. Each uses the system share sheet with the page link, whose Open Graph card the site already makes.

| Share | Link | Text |
|---|---|---|
| A raid (v1) | the raid's `/app/…` link | `{zone} · {Weekday} {d Mon} {time} · {n} of {size} signed up` |
| A build | `/b/{id}` | `{Race} {Class} · {totals}` (the web's `shareTextSummary`) |
| A character | `/character/{region}/{ruleset}/{name}` | `{name}, {Spec} {Class}, level {n}` |
| A guide | `/guides/{class}/{spec}` | `{Spec} {Class} guide` |
| A tier list | `/tiers` | `{Role} tier list` |
| A public or unlisted report | `/reports/{id}` | `{zone} · {d Mon}` |
Guild and private reports have no share action. Items share nothing.

### Haptics (v1 §7 extended)
| Moment | iOS | Android |
|---|---|---|
| Segment, chip, picker, tab section, character pick | `.selection` | `SEGMENT_TICK` (API 34+), else `CLOCK_TICK` |
| Planner: add a point | `.selection` | same |
| Planner: remove a point (long press) | `.impact(.medium)` | `LONG_PRESS` |
| Planner: refused | `.error` | `REJECT` (API 30+) |
| Saved: sign-up, share link ready, watch | `.success` | `CONFIRM` (API 30+) |
| Refused sign-up | `.error` | `REJECT` |
| Pull to refresh finished, scrolls, meters, rows | none | none |
| Widget and Control buttons | the system's | the system's |
Haptics are off under Reduce Motion no more than the system says; a unit test records every call.

### Offline mode and stale data (boards F15.4, F15.5)
Three states, one strip component (v1's offline strip, 36 pt):

| State | Strip | When |
|---|---|---|
| **Offline** | `Offline. Showing what you had at {time}.` (v1, unchanged) | no connection and a cached answer exists |
| **Stale** | `Forever Sixty did not answer. Showing what you had at {time}.` with a trailing `Retry` (gold, 44 pt target) | there is a connection, the request failed or timed out, the cache is older than its limit, and the failure is not a `4xx` |
| **New data** | `New game data is ready.` with a trailing `Update` (gold), on a gold 10% ground | a newer build was found while a screen showing build data is open |
Cache limits: the Home answer and the raids 5 minutes (v1); rankings and reports 10 minutes; guides, tiers, BiS, talents and indexes: until the active build changes. A stale strip never covers a `4xx` or `5xx` error; that is the region's own error state.
- **Never saved and offline:** the screen says its `not cached` sentence (`This guide is not on your phone yet. Connect once to save it.`, `This list is not on your phone yet. Connect once to save it.`).
- **Writes** are never queued (v1): sign-ups, shares and removals say `{Thing} needs a connection.` and stay disabled with their reason.
- **New data never moves what you are reading.** It waits for `Update` or the next open.

**You › Offline data (F15.4).** Pushed from You. Title `Offline data`.
- Section label `On this phone, {size}`. A quiet card of rows: tile, name, `{what} · {size} · Updated {when}`:
  - `Guides`: `Shaman, Warrior · 2.4 MB · Updated today`
  - `Best in slot lists`, `Talents`, `Search index`, `Item icons`.
- Section `Updates`: switch `Keep my classes ready` (`Saves the guide, lists and talents for every class you have a character in.`, on) and switch `Download on Wi-Fi only` (`Icons and new guides wait for Wi-Fi.`, on).
- `Update now` (primary, 48 pt) and `Remove downloaded data` (secondary, `--color-wipe` label; asks `Remove downloaded data? It downloads again when you open a guide or list.` with `Remove` and `Cancel`).
- Footer: `Checked today at {time}. Your characters, sign-ups and sims are not stored here; they load when you are online.`
- No build ids, hashes or file counts anywhere (tenet 7).

### Copy keys
`offline.stale` `Forever Sixty did not answer. Showing what you had at {time}.`; `offline.retry` `Retry`; `offline.newData` `New game data is ready.`; `offline.update` `Update`; `offlineData.title` `Offline data`; `offlineData.section` `On this phone, {size}`; `offlineData.keep` `Keep my classes ready`; `offlineData.wifi` `Download on Wi-Fi only`; `offlineData.updateNow` `Update now`; `offlineData.remove` `Remove downloaded data`; `offlineData.removeConfirm` `Remove downloaded data? It downloads again when you open a guide or list.`; `offlineData.footer`; `intent.confirmSignup` `Sign {name} up as {state} for {zone}, {weekday} {time}?`; `intent.which` `Which raid?`; `intent.spec` `Which spec is {name} on {guild}?`; `intent.signedUp` `You are {state} for {zone}, {weekday} at {time}, as {role} on {name}. {n} of {size} signed up.`; `shortcut.signupToast` `You are In for {zone}.`

### Data
- **Existing:** every endpoint of F4, F2, F3, F9 to F12 and F16. Intents add no endpoint.
- **Contract change W3-C15:** none new beyond W3-C11 (links) and the `features` switches of `GET /v1/app/config`: `guides`, `reports`, `rankings`, `search`, `planner`, `widgets`, `live_activities`, `calendar_feed`. Each lane ships dark behind its switch.
- **Lane F15 builds the intents, entities, index, share actions, the stale and new-data strip, and Offline data on both apps.** It reuses the static data store that lane F12 builds.

### Acceptance captures
F15.1 the sign-up by voice with confirmation and the refused answer; F15.2; a Shortcuts-app run without confirmation; F15.3 Spotlight and launcher search; the Android shortcut menu and the toast; F15.4 with data and empty; F15.5 each strip; the share sheet for a character, a guide and a report on both platforms; haptic calls recorded in a unit test.

## 5.8 F16. Guild extras

**Reference.** The web's guild roster and attendance (`design/specs/2026-10-27-guild-attendance.md`), Raid-Helper's roster, Apple and Google Calendar.

### F16.1 Roster (board F16.1, F16.2)
- Pushed from the Raids tab guild head row `Roster`. Large title `Roster`; trailing search button (filters the list by name).
- **Counts card** (quiet): three figures at `stat` size with `meta` captions: `{n}` `on the roster`, `{n}` `in game, no account`, `{n}` `waiting`. They are the API's `counts`; officers see the same.
- Chips: `All`, `Tanks`, `Healers`, `DPS`, then `Sort: {Rank}` (opens a menu: `Rank`, `Class`, `Name`, `Attendance`, `Level`).
- **Row** (64 min): the spec icon 32 (a class crest when the spec is not known, never a letter tile), name in the class colour with the supporter mark when it applies, meta `{Spec} · {level} · {rank}`, and on the right attendance: a bar of 8 segments (7 x 12 pt, kill for a night present, `--color-line` for absent) over the mono caption `{n} of 8 nights` (build review: a bare `8 of 8` had no unit). A character waiting for approval shows the tag `Waiting` instead; a name the addon sees in game that no account has claimed shows the tag `No account` and has no tap target. A chevron. Rows with the same sort key keep name order.
- Attendance is the API's `{present, nights}`: the last 8 raid nights with a log. If `nights` is under 8 the bar has `nights` segments. With no nights: `No logged nights yet.`
- Footer: `Attendance is the last 8 raid nights with a log. Tap a raider for their page.`
- A row tap opens the character page (F11.4).
- **No match (F16.2):** `No raiders match.` and `Change the role or clear the search.`
- Loading: eight skeleton rows. Error: `The roster did not load.` + `Try again`. Not a verified member: the Raids tab's own states (F4.13, F4.14).

### F16.3 Raid history (board F16.3)
- Pushed from `Raid history`. Title `Raid history`. Label `Logged raid nights`. One quiet card.
- **Row** (72 min): a date block (48 x 52, `THU 10 DEC`), the zone, meta `{n} boss(es) · {n} kill(s) · no wipes` or `· {n} wipe(s)`, meta `{n} raiders · {h} h {m} m`, or `{n} players` when the night is a dungeon of 10 or fewer (mono figures; the size is the night's real size, never below the kill's player count; lengths differ night to night), and the tag `You were there` (kill) when one of your characters is in `present[]`.
- Paged by `next_cursor` (20 a page): `Load more` text row at the foot; spinner while loading.
- A tap opens `Your raid night` (F4.22) when you were there and the night has a planned raid, else the report (F10.2, guild-visible).
- Footer: `A night opens your raid night when you were in it, otherwise its report.`
- Empty: `No logged raid nights yet.` and `Nights show here once someone uploads a log.`

### F16.4 Calendar (boards F16.4 to F16.6)
You › Calendar. Two parts:
1. **A calendar the app keeps up to date.**
   - Switch `Raids in my calendar` (`A calendar called Forever Sixty raids. The app keeps it up to date.`). Off by default.
   - Row `Which raids`: `All planned` or `Only ones I am In or Late for`.
   - Row `Alert`: `None`, `15 minutes before`, `1 hour before invites` (default), `1 day before`.
   - Turning it on asks for calendar access in context (iOS full access, `NSCalendarsFullAccessUsageDescription`: `Forever Sixty keeps a calendar of your guild's raids up to date.`; Android `READ_CALENDAR` and `WRITE_CALENDAR`). The system prompt is never shown before the player flips the switch.
   - The app owns one local calendar, `Forever Sixty raids`. Each planned raid is one event: title `{zone} · {In|Late|Bench|Out|Not signed up}`, start the invite time (else the start), end the raid's end, location `{guild} · {zone}`, notes `Pull {time}. {n} of {size} signed up.` and the raid's link. Your sign-up changes the title.
   - It syncs on every refresh, on `raid.posted`, `raid.changed` and your own answer, and in the background refresh. A cancelled raid's event is deleted. Past events stay.
   - **Access off (F16.5):** a card, `Calendar access is off for Forever Sixty.` and `Turn it on in Settings to add raids to your calendar.` with `Open Settings`. The calendar was deleted by the player: the switch shows off with `The calendar was removed. Turn it on again.`
2. **A calendar link for other calendars.** Row `Calendar link`: `For Google Calendar on the web, Outlook and other people. Anyone with the link sees your guild's raid times.` with `Copy link` and `Reset link` (asks `Reset the link? The old link stops working.` with `Reset` and `Cancel`). The link is a revocable feed of the planned raids of your guilds, titled as above. It carries no personal data beyond your sign-up answers. iOS also offers `Add to Calendar` through `webcal://`.
- Footer: `Your sign-up shows in the event title: “Onyxia's Lair · In”. Changing it in the app changes the event.`
- **The event (F16.6)** as Apple Calendar and Google Calendar show it, board F16.6.
- v1's single `Add to calendar` on the raid page stays: it adds one event once, with no permission beyond write-only.

### Copy (verbatim)
| Key | Text |
|---|---|
| `roster.title` | `Roster` |
| `roster.counts` | `on the roster` / `in game, no account` / `waiting` |
| `roster.sort` | `Sort: {by}` |
| `roster.empty` | `No raiders match.` / `Change the role or clear the search.` |
| `roster.footer` | `Attendance is the last 8 raid nights with a log. Tap a raider for their page.` |
| `roster.noNights` | `No logged nights yet.` |
| `roster.attendance` | `{present} of {nights} nights` |
| `roster.waitingTag` / `roster.noAccountTag` | `Waiting` / `No account` |
| `roster.error` | `The roster did not load.` |
| `history.title` | `Raid history` |
| `history.label` | `Logged raid nights` |
| `history.there` | `You were there` |
| `history.more` | `Load more` |
| `history.empty` | `No logged raid nights yet.` / `Nights show here once someone uploads a log.` |
| `calendar.title` | `Calendar` |
| `calendar.switch` | `Raids in my calendar` |
| `calendar.link` | `Calendar link` |
| `calendar.copy` / `calendar.reset` | `Copy link` / `Reset link` |
| `calendar.off` | `Calendar access is off for Forever Sixty.` |

### Data
- **Existing:** `GET /v1/guilds/{id}/home` (`roster[]` with `attendance`, `members[]` with state, level and spec, `counts`), `GET /v1/guilds/{id}/raids` (`rows[]` with `present[]`, cursor), `GET /v1/me/raid-nights/{raid}`, `GET /v1/reports/{id}`, the raid views.
- **Contract change W3-C16:** `members[].role` and `roster[].role` on `GET /v1/guilds/{id}/home` (additive; `tank`, `healer`, `dps` or null), so the role chips never derive a role from a spec name in the app.
- **Contract change W3-C17:** `GET /v1/me/calendar-feed` (session or device token) returns `{ url, created_at }` and creates the feed on first call; `POST /v1/me/calendar-feed/reset` issues a new link and revokes the old; `GET /v1/calendar/{token}.ics` is public by token, `text/calendar`, `Cache-Control: private, max-age=300`, `404` for an unknown or revoked token, 60 reads an hour per token. A migration adds `calendar_feeds` (user, a hash of the token, created and revoked times). The token is shown once and stored hashed.
- **Lane F16 builds W3-C16 and W3-C17, the roster, history, calendar sync and link, on both apps.**

### Acceptance captures
F16.1 for each chip and each sort; F16.2; F16.3 with a first page and a second; F16.4 on, off, and each Which raids value; F16.5; F16.6 on both calendars; the link reset alert; the roster of a guild with 40 members scrolling without a hitch; largest text on F16.1 and F16.3.

# 6. Platform behaviour and accessibility

Everything in v1 §7 and §7.1 holds. Additions:

| Element | VoiceOver / TalkBack label | Action or state |
|---|---|---|
| Search button | `Search` | opens F12 |
| Scope chip | `Items`, selected | `.isSelected` / `selected` |
| Guide section chip | `Talents, section 2 of 6` | scrolls |
| Numbered stat chip | `Priority 1, Intellect` | |
| Tier row | `Fury Warrior, rank 1 of 20, tier S, top spec, Troll` or `Destruction Warlock, rank 2, tier S, 2.8 percent behind the top, Troll` | `Opens the guide` |
| Tier row tie | adds `Within 1 percent of the next spec` | |
| Report row | `Onyxia's Lair, Thursday 10 December, 1 boss, 1 kill, guild visible, live, updated 20 seconds ago` | |
| Parse figure | `Onyxia, kill, 3 minutes 12 seconds, 412 damage per second, 78th percentile` | |
| Fight row | `Onyxia, wipe, 1 minute 12 seconds, boss at 61 percent` | `Shows the meters` |
| Trash row | `Trash in Onyxia's Lair, 9 pulls, 14 minutes 20 seconds` | |
| Meter row | `Rank 6, Lightningbolltz, Elemental Shaman, 412, your character` | |
| Rank row | `Rank 14 of 212, Lightningbolltz, 496 damage per second, Hordebreakers, your character` | |
| Score | `Forever Score 1,830, Epic. Gear 412 of 500, Experience 620 of 700, Parses 540 of 600, Attendance 180 of 200, Discipline 78 of 100.` | |
| Paperdoll socket | `Head, Blue Dragonscale Helm, Epic, item level 61. A better item is on your best in slot list.` | `Shows the item` |
| Watch star | `Watch Voltaris` / `Stop watching Voltaris` | toggle |
| Tree talent | `Convection, rank 5 of 5` | custom actions `Add a point`, `Remove a point` |
| Stepper | `Elemental Fury, rank 3 of 5` with `Add a point` and `Remove a point` buttons | `stateDescription` of the refusal line |
| Planner refusal | announced assertively | `AccessibilityNotification.Announcement` (`LiveRegionMode.Assertive`) |
| Widget | the whole widget is one element: `Next raid, Onyxia's Lair, Thursday 8 pm, in 2 days 2 hours. You are In as DPS. 24 of 40 signed up.`; each button is its own `In`, `Late`, `Out` | |
| Live Activity | `Invites for Onyxia's Lair in 38 minutes. You are In.` | countdown not announced per tick |
| Roster row | `Kraggor, Protection Warrior, level 60, Guild Master, attended 8 of 8 recent raid nights` | |
| Attendance bar | hidden; the row carries the words | |
| Raid history row | `Onyxia's Lair, Thursday 17 December, 1 boss, 1 kill, no wipes, 38 raiders, 1 hour 47 minutes. You were there.` | |
| Strip | `Forever Sixty did not answer. Showing what you had at 9:41 pm.` announced once | `Retry` is a button |

Touch targets: 44 pt (iOS) and 48 dp (Android), met by row height or padding, including every star, chip, segment and the stepper's buttons (48). The 8 pt gap between the paperdoll's columns belongs to no target.
Focus order follows the visual order; a sheet traps focus and returns it. Hardware keyboard: Tab, Space, Return, arrows in segments, Escape closes a sheet. In the planner, arrow keys move between talents and `+` and `-` add and remove a point.
Dynamic Type to AX5 and font scale 2.0: see §9.
Reduce Motion: sheets cross-fade; widgets and Live Activities never animate.
Increase Contrast: `--color-muted` becomes `--color-nav` (v1 §4.4); the bars' tracks go one line step brighter.

# 7. Every interactive element, every state

v1 §8's table holds for buttons, rows, segments, chips, switches, fields and tabs. New elements:

| Element | Default | Pressed | Focused | Disabled | Loading | Error |
|---|---|---|---|---|---|---|
| Search field | glass or pill, placeholder muted | caret, gold border | gold border | n/a | n/a | n/a |
| Scope chip | 36 pt pill, 1 pt line | 10% gold wash | gold ring | n/a | n/a | n/a |
| Section chip (guide) | as scope chip | as above | ring | n/a | n/a | n/a |
| Watch star | outline, muted | 10% gold wash | ring | 50% offline | filled at once, reverts on failure | reverts, line `{Name} was not saved. Try again.` under the row |
| Tier row | `--raised` | `rgba(229,185,85,.10)` | ring inset 2 | n/a | skeleton | n/a |
| Meter row | bar and number | `pressed` | ring | n/a | skeleton at height | n/a |
| Tree cell | v1 F6 borders | 0.96 scale | gold 2 pt ring | n/a | n/a | refusal: wipe line in the strip, `.error` haptic |
| Stepper button | 48 pt circle, 4.5% fill | 10% gold wash | ring | 40% at 0 or at the max | n/a | n/a |
| Share anyway | primary | v1 | ring | 50% offline | spinner in the label | the field's sentence |
| Widget button | tonal, muted label | system | n/a (system) | the whole row dims when locked | n/a | `Not saved` for 10 s |
| Live Activity buttons | system | system | n/a | n/a | n/a | n/a |
| Stale strip `Retry` | gold text | 10% gold wash | ring | n/a | spinner replaces the text | the strip stays |
| New-data `Update` | gold text | wash | ring | n/a | spinner | `The update did not finish. Try again.` |
| Calendar switch | v1 switch | system | system | 50% when access is off | n/a | value returns, line under |
| Roster sort menu | muted value with chevron | menu opens | ring | n/a | n/a | n/a |
| Swipe remove (build) | hidden | red action `Remove` | n/a | n/a | row fades | row returns, `That build was not removed. Try again.` |

# 8. Hierarchy

| Region | The one primary | Visibly secondary |
|---|---|---|
| Guides tab | `Your guide` card | the class grid, the footer |
| A guide | the at-a-glance card | strip, sections, sources |
| Tier lists | the callout | rows (the top spec is not louder than yours) |
| A report | `Your parses` | fights, coach, footer |
| A fight | your row | the meters |
| Rankings | the answer card | filter line, board |
| A character | the Forever Score card | bests, gear, builds |
| An item | the name and your pick card | tooltip, sources, specs |
| Planner | the tree | chips, strip, action bar (`Share` is the one gold button) |
| Widgets | the zone and countdown | the foot tag |
| Roster | the counts card then the rows | chips |

One gold-filled button per screen at most (v1 §9). Gold text marks links and the countdown. Whitespace: 22 before a section label, 14 inside cards, 16 margins.

# 9. Phone behaviour and what collapses

| Change | iPhone SE 3 / 360 dp | iPhone 17, Pixel 9 | Largest text |
|---|---|---|---|
| Tab labels | 5 across, system size | same | system; iOS Large Content Viewer |
| Guide section strip | scrolls | scrolls | scrolls; the current chip is kept in view |
| Class grid | 3 x 3 | 3 x 3 | 2 columns, crests 44 |
| Tier row | gap under the name | gap in the right column | the gap and the chip stack under the name |
| Meter rows | 9 | 11 | 4; value under the name |
| Paperdoll | 2 columns, sockets 36 | sockets 40 | one column of 17 rows |
| Planner tree | 4 columns at 44 icons, strip two lines | 4 columns at 48 | icons stay 44, the tree never scrolls sideways, the strip grows |
| Rankings row | meta wraps to two lines | one line | value under the name |
| Roster attendance | the bar is hidden below 340 and `{n} of 8` stays | bar and number | number only, under the name |
| Widgets | follow the system | | text caps at the largest non-accessibility size on iOS (WidgetKit's own rule) and at 1.3x on Android, while the app scales fully; nothing in a widget truncates at that cap; the medium widget's buttons keep 44 pt (amended 2026-10-31, F14 build review) |

These never collapse: the full character name; every item's icon and quality colour; every spell's and talent's icon and rank; the points left; the parse number's colour; the answer checked on a widget; the unit on every figure; `Gain not modelled yet` where it applies. Nothing is cut with an ellipsis except a widget's item name on one line.

# 10. Contract changes (the C-list, spec §14 pattern)

Each change is implemented by the lane named, in `api/` (and the site for static files), in the same lane as the apps. Each has an entry in `api/openapi.yaml` (`app` tag, the shared `X-App` and `X-App-Install` parameters and the `426` response), fixtures under `docs/contracts/fixtures/app/`, a section in `docs/contracts/app-v1.md`, a Go contract test that pins the handler to its fixture, and a generated Swift and Kotlin model (`web/scripts/gen-mobile-api.mjs`). Nothing merges with a drift. Contracts are floors: no fixture pins a count or an id.

| ID | Lane | Change | Auth | Migration | Fixtures |
|---|---|---|---|---|---|
| W3-C1 | F9 | Static: `/data/{build}/guides/index.json`, `/data/{build}/guides/{class}-{spec}.json`, `/data/{build}/tiers.json`, `/data/changelog.json` | public | no | `cdn.guides.index`, `cdn.guides.shaman-elemental`, `cdn.tiers`, `cdn.changelog` |
| W3-C2 | F10 | `GET /v1/reports/{id}/me` | session or device | no | `reports.me`, `reports.me.empty`, `reports.me.error-not-found` |
| W3-C3 | F10 | `GET /v1/reports/{id}/fights/{n}/meters` | as the report (public, unlisted, guild, owner) | no | `reports.meters`, `reports.meters.error-not-found` |
| W3-C4 | F10 | Document `GET /v1/reports/{id}/coach/night`, `GET /v1/me/coach/recent` as app operations | session, `log_coach` | no | `coach.night`, `coach.night.error-premium-required` |
| W3-C5 | F11 | Document `GET /v1/rankings`, `GET /v1/rankings/percentile`, `GET /v1/phases` as app operations | public | no | `rankings.list`, `rankings.list.empty`, `rankings.percentile` |
| W3-C6 | F11 | `GET /v1/rankings/encounters` | public | no | `rankings.encounters` |
| W3-C7 | F11 | `GET /v1/me/watching`; `PUT`, `DELETE /v1/me/watching/{region}/{ruleset}/{name}` | session or device | yes: `watched_characters` | `watching.list`, `watching.list.empty`, `watching.put`, `watching.error-limit` |
| W3-C8 | F12 | Static: `/data/{build}/loot.json`, `/data/{build}/search.json`, `/data/{build}/item-uses/{shard}.json` | public | no | `cdn.loot`, `cdn.search`, `cdn.item-uses` |
| W3-C9 | F13 | `DELETE /v1/builds/{id}` | session or device | maybe: nullable `builds.user_id` | `builds.delete`, `builds.delete.error-not-found` |
| W3-C10 | F13 | Document `/data/{build}/races.json`, `/data/{build}/combos.json` | public | no | `cdn.races`, `cdn.combos` |
| W3-C11 | F13 | `apps/shared/links.json` rows (§3.4) | n/a | no | the generated association files |
| W3-C12 | F14 | `GET /v1/me/widget` | session or device | no | `widget.get`, `widget.get.no-guild` |
| W3-C13 | F14 | `POST /v1/me/live-activities`; `DELETE /v1/me/live-activities/{id}` | session or device (a sim registration needs `server_sims`) | yes: `live_activity_tokens` | `live.register`, `live.register.error-premium-required` |
| W3-C14 | F14 | Events `widget.refresh`, `live.raid.*`, `live.sim.*` and their push payloads | server | no | `push.widget-refresh`, `push.live-raid-start`, `push.live-sim-update` |
| W3-C15 | F15 | `GET /v1/app/config` gains the `features` switches | public | no | `app.config` updated |
| W3-C16 | F16 | `roster[].role`, `members[].role` on `GET /v1/guilds/{id}/home` | session | no | `guild.home` updated |
| W3-C17 | F16 | `GET /v1/me/calendar-feed`; `POST /v1/me/calendar-feed/reset`; `GET /v1/calendar/{token}.ics` | session or device; the feed by token | yes: `calendar_feeds` | `calendar.feed`, `calendar.ics` |

**Shapes of the new answers** (the OpenAPI schema is written from these):

- **W3-C2** `{ report_id, characters: [ { key, name, class, spec, spec_name, role, fights: [ { index, encounter_id, encounter, kill, duration_ms, value, unit ("dps"|"hps"|"dtps"), percentile (int|null), rank (int|null), ranked (int|null), new_best (bool), deaths } ] } ] }`. A character is one of the caller's characters in the report. `percentile` is null when the bracket has no ranked kills; the app prints no parse number then. `404` when the report is not the caller's to see.
- **W3-C3** `{ fight_index, kill, duration_ms, players: [ { key (nullable), name, class, spec, spec_name, role, active_ms, damage_done, healing_done, damage_taken, dps, hps, dtps, deaths, mine (bool) } ] }`, sorted by the server for the requested `?metric=damage|healing|taken` (default `damage`). A player with no character page has `key` null.
- **W3-C6** `{ encounters: [ { id, slug, name, zone, size, kills } ] }`, most kills first.
- **W3-C7** `GET` returns `{ rows: [ { key, region, ruleset, name, class, level, added_at } ] }`; `PUT` returns `204` (idempotent), `404` for an unknown or anonymised character, `409 limit` with `fields` at 50. A watched character whose account becomes anonymous disappears from the list.
- **W3-C9** `204`; `404`.
- **W3-C12** `{ generated_at, character: { key, name, class, spec_name } | null, next_raid: { raid_id, guild_id, guild_name, zone_name, starts_at, invite_at, pull_at, size, signed_up, status ("open"|"locked"|"cancelled"), my: { state ("in"|"late"|"bench"|"out"|null), spec, role }, can_sign_up } | null, upgrades: { count, total_gain, unit, rows: [ { slot, pick: { id, name, quality, icon, item_level }, gain (number|null) } ] } | null, live: { sim_runs_running } }`. `my.spec` null means the widget shows `Choose a spec`. A signed-out call is the shared `401`; the widget shows its signed-out state.
- **W3-C13** request `{ kind ("raid"|"sim"), platform ("ios"|"android"), push_token, push_to_start_token?, object_id }`, answer `{ id }`. A registration is removed when the object ends or the token fails.
- **W3-C17** `GET /v1/me/calendar-feed` returns `{ url, created_at }` and creates the feed on first call; `POST …/reset` returns the same with a new URL. The `.ics` is RFC 5545 with one `VEVENT` per planned raid, `UID` `raid-{id}@foreversixty.gg`, `SUMMARY` `{zone} · {state}`, `DTSTART` the invite time, `DTEND` the raid's end, `LOCATION`, `DESCRIPTION`, `URL`, and no `VALARM` (the player's calendar app decides).

**Security rules:** the calendar token is 128 bits, shown once, stored hashed, revocable, and never logged. Live-activity and push tokens never appear in logs or metrics. No new endpoint returns a token other than the caller's own. `GET /v1/reports/{id}/me` and `/meters` follow the report's visibility rules exactly and answer `404`, never `403`.

**Launch freeze.** The owner overruled it (2026-10-31): these lanes build the server work now. The rule that only the deployed API migrates (incident 2026-10-04) still holds: a lane's migration ships with its handler in the same deploy, and no seed or tool migrates ahead of it.

# 11. Performance and polish limits

v1 §11 holds, plus:

| Budget | iOS (iPhone 12) | Android (Pixel 6a) |
|---|---|---|
| Search keystroke to results (items, guides, talents) | ≤ 50 ms | ≤ 50 ms |
| Search index load, warm | ≤ 150 ms | ≤ 200 ms |
| Planner tap to the drawn point | ≤ 1 frame | ≤ 1 frame |
| Planner tree scroll and tap with 51 points | no hitches over 16 ms at p95 | janky frames under 1% |
| A guide opens from cache | ≤ 300 ms to the at-a-glance card | ≤ 400 ms |
| Meters, 40 rows | no hitches at p95 | janky frames under 1% |
| Widget timeline reload | within the system's budget | within the system's budget |
| Widget snapshot size | ≤ 8 KB | ≤ 8 KB |
| Offline data, default set (one class) | ≤ 8 MB without icons | same |
| Download size added by Wave 3 | ≤ 6 MB | AAB ≤ 3 MB |

- No layout shift: the tab bar's five slots, the segments, the tier list's rows and the planner's tree reserve their size on first draw.
- No missing-icon flash: spell, talent and item tiles draw the bordered tile at once and fade the image in (v1 §11).
- No spinner over cached content. The stale strip is drawn above the content from the first frame of a failed refresh, so it pushes nothing down after the content has drawn: its 36 pt are reserved on screens that can show it.
- Static packs download in the background, resume, and never block the UI.

# 12. Acceptance captures

Every lane attaches captures from the iPhone 17 simulator (iOS 26) and a Pixel 9 emulator (API 35) at the default text size and the largest (iOS AX5, Android font scale 2.0), driven by fixtures, as v1 §15. They are compared side by side with the boards in `design/mocks/2026-10-31-mobile-wave3/renders/`. Each section 5.x lists its captures. Every lane also records a VoiceOver and a TalkBack pass (screen recording with captions), the haptic calls in a unit test, the budget it owns, and the contract test for each new endpoint.

# 13. Checklist

- [ ] Five tabs; the search button on Home, Raids, BiS and Guides; F3's old filter gone.
- [ ] Guides, tier lists and what changed read from the `guides`, `tiers` and `changelog` static files; the order and letters are never computed in the app.
- [ ] Every spell, talent and item shows its icon; rank on every rotation line; talent totals summed from the drawn build.
- [ ] A new game build never changes a screen under the reader; the new-data strip asks first.
- [ ] Reports: trash is one entry per zone; parses count kills; `Your parses` first; coach drawn only with `log_coach`, never an upsell.
- [ ] Rankings: your row pinned only off the page; a single-class board has no per-row crest; parse colours from the ladder.
- [ ] A character with no rated fight shows `No logs yet.`, never a zero.
- [ ] Search: items, guides and talents work offline; characters need a connection and say so.
- [ ] An item shows only drop chances the data has; no `0%`; a source is one line a player recognises.
- [ ] Planner: tap adds, long press removes, every refusal in words, the shared confirm text verbatim; a link cannot be edited; the editor copies, never edits, a shared build.
- [ ] Planner rules match the golden vectors on both platforms.
- [ ] Widgets: every state has words; In, Late, Out answer without opening the app; stale after 12 h says so; no spinner, no empty box.
- [ ] Live Activity and ongoing notification start from pushes; the sim one exists only for a server run.
- [ ] No mobile-only Premium; no Premium word to a free player; coach and sim surfaces follow the web's flags.
- [ ] Siri confirms once, by voice only; intents give the app's refusal sentences.
- [ ] Spotlight indexes characters, raids, guides; nothing is indexed after sign out.
- [ ] Share links exist only for things with a web page.
- [ ] Offline, stale and new-data strips are distinct; no write is queued offline.
- [ ] The roster's role chips come from the server; attendance is `{present} of {nights}`.
- [ ] The calendar permission is asked only after the switch; a link can be reset; the token is shown once.
- [ ] Every element has the §7 states; 44 pt and 48 dp targets; labels per §6.
- [ ] §11 budgets measured on device.
- [ ] Every new endpoint has OpenAPI, fixtures, a contract test, generated clients and a feature switch.
- [ ] ux-designer and wow-player SHIP on the captures before merge.

# 14. Build lanes

One lane is one feature on **both** platforms (one iOS lane and one Android lane working from this section, the same fixtures and the same boards), plus the server and static-data changes of §10 for that feature. A lane merges by direct push to main behind its switch. Every lane follows the house rules: design before build, tests first for logic, captures at both sizes, two SHIP verdicts on the captures.

| Lane | Feature | Depends on | Builds (contract) | Shared pieces it owns |
|---|---|---|---|---|
| **L-F14** | Widgets and live surfaces | v1 F4 and F2; A2 senders and push tokens (merging now) | W3-C12, W3-C13, W3-C14 | widget snapshot container; App Intent target; `CharacterEntity` |
| **L-F16** | Guild extras | v1 F4; a character row opens the web page until L-F11 lands | W3-C16, W3-C17 | calendar sync service |
| **L-F12** | Search and the item page | v1 F2.6, F3 | W3-C8 | the static data store (download, swap, delete, size accounting); the search index reader |
| **L-F10** | Your reports | v1 F4.22; A2 `report.ready`, `raid.summary` | W3-C2, W3-C3, W3-C4 | the meter row |
| L-F9 | Guides, tier lists, what changed | L-F12's static data store | W3-C1 | the fifth tab; the guide reader; tier chip; numbered chip |
| L-F11 | Rankings and character pages | L-F12 (search characters open this page) | W3-C5, W3-C6, W3-C7 | the Paperdoll (and F2.7); parse number |
| L-F13 | Talent planner | v1 F6; L-F9 (the guide's `Open in the planner` row) | W3-C9, W3-C10, W3-C11 | the planner rules and golden vectors; the stepper strip |
| L-F15 | System integration | L-F12 (entities, index), L-F14 (intents, widgets), L-F16 (calendar), L-F11 (character entity) | W3-C15 | the stale and new-data strip; the Offline data screen |

**Order, at most four features at once, most valuable first.** Value is how often a player uses it times what only a phone does well.
1. **L-F14.** The flagship: it answers "am I in tonight?" from the home screen and lock screen, and the Live Activity and the sign-up intent share one target. Nothing else in the app is used as often or as quickly.
2. **L-F16.** The raid is the headline (v1). Roster, history and the calendar are what a guild member does around a raid, and the calendar link reaches players who never open the app.
3. **L-F12.** Search with the item page is the most obvious missing tool, and it builds the offline data store that F9 and F15 reuse.
4. **L-F10.** The after-raid loop: `raid.summary` already opens F4.22; reports and parses are what a player checks next.

Start these four together. They share no screen: F14 touches widgets and the intent target, F16 the Raids and You screens, F12 the root bars and a new screen, F10 Home's row list and the report screens. The shared files are Home's row list (F10 adds both new rows) and the root bars (F12 adds the button); the lanes rebase on each other.

**Second wave, as first-wave lanes land:** L-F9 and L-F11 may start as soon as L-F12's store merges. Then L-F13 (needs L-F9's guide row). L-F15 last, because it wires everything the others built into Siri, Spotlight and Shortcuts.

**Cross-lane rules:**
- Each lane ships dark behind its `features` switch; a feature turns on in both apps together.
- Static data lanes (F9, F12) deploy with the site; a lane with a migration (F11, F13, F14, F16) ships it with its own handler in one API deploy, and only the deployed API migrates.
- Home's two new rows (`Logs`, `Rankings`) appear only when their switches are on.
- The Paperdoll and the stale strip each get their own small design review from these boards before the first lane that uses them merges.

# 15. Open questions

None blocking. The owner's rulings of 2026-10-31 (maximum features; build server work now; free with the web's entitlement rules) settle scope, Premium and order. Two items are for the data lane, not for design:
- Elemental Shaman is last on the DPS tier list in build 70658. The boards use it unchanged; the data lane should confirm the band.
- Drop chances are null for most items in `loot.json`. The item page prints a chance only when the data has one.
