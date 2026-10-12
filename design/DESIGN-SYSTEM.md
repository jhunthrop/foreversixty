# Forever Sixty design system

Derived from the approved "Cinematic" direction and the live-game homepage mockup in this folder. Source of truth for tokens once the site exists is `web/src/styles/tokens.css`; keep this document and that file in sync.

## Principles

1. **Helper, not pitch.** The site exists to make one character play better. Every page leads with the player's class or character and the next thing that improves it; no marketing hero, no feature-list pitch. The front door is "pick your class" (the nine crests) or "your character", under the one line the site is allowed to say about itself: "Play your class better." Reference facts (dates, sources, the changelog) stay, one level down.
2. **The game's own icons, high definition.** Class crests, spec icons and item icons are the client's own art, upscaled 4x, bezel cropped, framed by us (rounded square with a 1px line, or the circle crest: the whole icon at an 80% fit over a disc made from the icon itself, blurred and darkened, with a class ring). Never redrawn, never restyled; the Forever Sixty design system artifact carries the sets and the rules. **One crest language:** wherever a character or class is pictured (headers, character cards, the current-character bar, account menu and rows, switch lists, report rosters, rankings), it is the circular ringed crest (`ClassCrest.astro` / `ClassCrestRing.svelte`), never a square class icon, a letter square or a Blizzard avatar thumbnail.
3. **Second screen first.** Dark by default, high contrast, 44px minimum hit targets, pages that paint before the player alt-tabs back.
4. **Every fact is dated and sourced, quietly.** A page header carries one "Updated" stamp and nothing else about provenance. Source pills (Blizzard, Datamined, Community, This site) live where the source changes what a player does: on a feed row (a datamined change is not a confirmed one), in the Sources footer of a content page, and beside a single-source claim in prose. Never in a page header, hero or title line; the player came for the answer, and the sites players actually use (Wowhead, Icy Veins) put provenance under the content, not over it.
5. **The game's own colors do the wayfinding.** WoW class colors and item-rarity colors are used consistently and never repurposed.
6. **Atmosphere in the header only.** The night-sky band lives at the top of a page. Content areas are flat, calm, and dense enough to be useful.
7. **Ornament stays out.** Warmth comes from Cinzel at small sizes, gold accents, and the sky. No stone frames, no parchment textures, no beveled buttons.

## Color

### Surfaces

| Token           | Value                 | Use                                                   |
| --------------- | --------------------- | ----------------------------------------------------- |
| `--bg`          | `#07090d`             | Page background                                       |
| `--bg-raised`   | `#0d111a`             | Panels, table bodies                                  |
| `--bg-card-top` | `#131824`             | Card gradient start (cards go `#131824` to `#0d111a`) |
| `--border`      | `#262e40`             | Panel and card borders                                |
| `--border-soft` | `#1c2230`             | Row dividers                                          |
| `--border-warm` | `#3a3326` / `#4a4030` | Inputs and buttons on the sky band                    |

### Text

| Token           | Value     | Use                              |
| --------------- | --------- | -------------------------------- |
| `--text`        | `#e9e4d8` | Body                             |
| `--text-strong` | `#f2eee4` | Headings, primary labels         |
| `--text-muted`  | `#9a9484` | Secondary text (6.3:1 on raised) |
| `--text-nav`    | `#b9b3a4` | Nav links, footer links          |

### Accent

| Token          | Value     | Use                                          |
| -------------- | --------- | -------------------------------------------- |
| `--gold`       | `#e5b955` | Links, focus, live indicator, primary accent |
| `--gold-hover` | `#f5d27a` | Link hover and active (pressed)              |
| `--gold-deep`  | `#a8762a` | Gradient end for gold text and bars          |
| `--gold-light` | `#fbe7a1` | Gradient start for gold text                 |
| `--ember`      | `#d66e28` | Sky horizon glow only                        |
| `--night`      | `#26405c` | Sky zenith only                              |

Gold text is a gradient (`#fbe7a1` → `#e5b955` → `#a8762a`, top to bottom) clipped to text. Use it for section titles only; the wordmark is flat `--text-strong`.

**Inline links** (a link inside a sentence) are underlined, since colour alone does not mark a link: `--gold` at rest (`--text-muted` for a secondary one such as the profile's guild), `--gold-hover` on hover (`--text-strong` for a muted one), `--gold-hover` while pressed (`:active`, every link), and the 2px gold focus ring.

### Source pills

Placement: feed rows, the Sources footer, and single-source claims in prose. Never in a page header (principle 3). The band header of a class page shows "Updated <date>" only.

| Pill              | Text      | Background              | Border                  |
| ----------------- | --------- | ----------------------- | ----------------------- |
| Blizzard / Hotfix | `#6fb1ff` | `rgba(0,112,221,.18)`   | `rgba(0,112,221,.35)`   |
| Datamined         | `#c98bff` | `rgba(163,53,238,.16)`  | `rgba(163,53,238,.35)`  |
| Community         | `#7bff5c` | `rgba(30,255,0,.10)`    | `rgba(30,255,0,.25)`    |
| This site         | `#e5b955` | `rgba(229,185,85,.14)`  | `rgba(229,185,85,.35)`  |
| Sample            | `#9a9484` | `rgba(154,148,132,.12)` | `rgba(154,148,132,.30)` |

### Faction

Alliance `#6fb1ff` (text) / `#2f6fd6` (bars). Horde `#ff6b5c` (text) / `#c0392b` (bars). Wherever a faction is named it carries the game's own emblem: the client's UI-PVP-Alliance lion shield and UI-PVP-Horde red disc from the unit frame, upscaled 4x with their transparency kept, never boxed or framed. 20px before the word, 36px in a header. In a row descriptor (the home hero's descriptor and the header selector's rows) it is the flat faction logo, the Alliance lion and the Horde crest (`factionLogoSrc`, `{faction}-logo-512.webp`) at 16px, with the faction name as alt text; the 72px unit-frame shield and disc read as a blob and a generic shield at that size. Never a coloured dot alone.

### Class colors (WoW standard)

Warrior `#c69b6d` · Paladin `#f48cba` · Hunter `#aad372` · Rogue `#fff468` · Priest `#ffffff` · Shaman `#0070dd` (use `#3f8fe0` for text on dark) · Mage `#3fc7eb` · Warlock `#8788ee` · Druid `#ff7c0a` · Monk `#00ff98`.

### Item rarity (WoW standard)

Poor `#9d9d9d` · Common `#ffffff` · Uncommon `#1eff00` · Rare `#0070dd` · Epic `#a335ee` · Legendary `#ff8000`. Level-range tags reuse these: green for low brackets, blue for mid, purple for max-level.

Rare and epic are too dark to read as small text on `--color-raised` `#0d111a`: 3.92:1 and 3.87:1, under WCAG AA's 4.5:1 for 13–14px semibold. Text uses two lightened variants instead — rare `#3d94f0` (6.02:1) and epic `#b866f5` (5.63:1) — the same move the Shaman class colour makes. The base colours above are unchanged and stay in use for bars, borders and icons; poor (6.96:1), common (18.88:1), uncommon (13.81:1) and legendary (7.50:1) already clear AA and are used as-is for text.

## Typography

| Role    | Face               | Fallback                          | Notes                                                                                                                                                                                                                                                                                        |
| ------- | ------------------ | --------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Display | Cinzel 600–800     | Trajan Pro, Georgia, serif        | Wordmark 20px, section titles 18px uppercase with 0.10em tracking, card titles 15–17px, stat figures 22px. Never above 22px on content pages. The home hero is the one exception: its h1 ("Play your class better." signed out, the character's name in class colour signed in) is 34px 700. |
| Body    | Barlow 400–700     | Helvetica Neue, Arial, sans-serif | 14–15px body, 13px secondary, 17px search placeholder                                                                                                                                                                                                                                        |
| Numbers | JetBrains Mono 500 | SF Mono, Menlo, monospace         | Dates, timers, counts, keyboard hints; `font-variant-numeric: tabular-nums`                                                                                                                                                                                                                  |
| Labels  | Barlow 700         |                                   | 11px, uppercase, 0.14em tracking, muted or gold                                                                                                                                                                                                                                              |

Self-host all three faces (Google Fonts license permits it) so no third-party request happens on page load.

## Spacing and shape

- Page gutter 48px desktop, 18px phone. Content max width 1344px.
- Section gap 32px desktop, 22px phone. Inside panels 16–20px.
- Grid gaps 12–16px. Row dividers 1px `--border-soft`.
- Radii: panels and cards 6px, buttons and inputs 4px, pills 3px, avatars 999px.
- Cards carry a 1px top inner highlight `rgba(229,185,85,.35)` and a 1px `--border`.
- Hit targets: 44px minimum on phone; buttons 36–48px tall.

## Components in the mockup

- **Sky band**: layered radial gradients (stars, ember horizon, night zenith) over a `#070b12` → `#1b1410` vertical gradient, a two-layer mountain silhouette SVG at the bottom, and a fade to `--bg`. Header only.
- **Search**: 56px tall, warm border, gold icon, `/` keyboard hint. First interactive element on every page.
- **State panel** ("This week"): label with a glowing gold dot, two-column key/value rows, a Sample or Updated stamp on the right.
- **Tool card**: icon, title, one-line description; the whole card is the link.
- **Feed row**: date (mono, muted), source pill, title link, optional one-line note.
- **Class tile**: 84px tall, 2px inset class-color top stroke, class name in class color.
- **Progress bar**: 6px, `--border-soft` track, filled with the relevant accent.
- **Secondary button**: 36px, warm border, uppercase 12px 700 with 0.06em tracking. There is no primary marketing button.
- **Account button**: the one filled gold button (44px, `#f0cc6c` → `#c99a3a`, dark text, 1px `--gold-deep` inset ring, soft gold glow). Used only for "Sign in with Battle.net" on the signed-out home; never for a tool or a link.
- **Character row**: class crest (36px circle with class ring), name in class colour (Cinzel 14px), muted descriptor, and on the right either a mono stat ("2 upgrades") or a progress bar. No page has a character switcher of its own: the header character selector is the one way to change character, and every character surface follows it in place.
- **Stat weights rail**: the convention players know from SimulationCraft's scale factors, in our own words. Rows sorted by scale factor; columns Stat, bar, Scale (per point, normalized so the top per-point stat = 1.00), Per point (absolute DPS per stat point), Error (sim error in the same units). The table holds per-point stats only, sorted, normalized to the top one (Agility = 1.00 for a hunter). Haste has no rating in this client, so it is never a table row: the caption says "Haste: 1.58 per 1%" and, when no item in the band carries haste, adds "not in the table because no item at this band has it". An insignificant stat keeps its row greyed with "not significant". Under the rows: one secondary button "Use these weights in the addon" and the line that the addon rates bag and tooltip items with them. Never a Pawn string or any other addon's name.
- **Class page header** (`ArtPanel` + `ClassHeader`): the tree art behind the night-to-bg gradient, the crest at 64px, the h1 at 22px in the class colour with the band and faction as a muted suffix, then a `SpecTabs` row (three 44px tabs, spec icon 28px plus name, the active one ringed in the class colour) and a controls row of `BandTabs` (five 44px mono pills labelled 20 to 29, 30 to 39, 40 to 49, 50 to 59 and plain 60, never "60 to 60"), `FactionToggle` (two 44px pills with the emblem at 20px) and the race select (44px).
- **Character header** (`CharacterHeader.svelte`, spec `design/specs/2026-10-11-character-header.md`): the one panel at the top of a page about the signed-in player's character, the same on every page that mounts it (home, BiS landing, BiS spec page, planner; Rankings and Logs next). Three fixed regions: **who** (crest 64px, 56 stacked, label always "Your character", name in class colour in Cinzel 700 24px, 22 stacked, the descriptor "Gnome Fury Warrior · Level 60 · [flat faction logo 16px] Alliance" with the race only when known and split into two fixed lines below 1280px and in the card layout, the sync line "Addon · 19 days ago" with only the stale words "Log out in the game to sync" in ember), **the page's answer** (an 11px label, then one figure or two at the same 26px mono size with a 15px muted unit, an optional 6px bar, an optional note; or one 14px sentence) and **the page's one action** (the filled 48px gold button, the only filled gold element on the page). Raised ground, 1px gold border, 6px radius, min height 136px from 1024px; below that it stacks. The states are fixed by the component, never by the page: loading (the ready box hidden under skeletons, so exactly its height), answer loading, stale, no gear on file, under the page's level floor, empty, failed ("Try again"), and one shared signed-out band (a single muted 13px line, "Sign in or paste an export to see your character here", 44px between hairlines, no button: the nav selector is the only sign-in control). A page supplies only its answer's words and its action; the API is in the component's header comment.
- **Profile header** (`ProfileDescriptor.svelte` on the public character profile, `/character/<region>/<ruleset>/<name>`; Task L8.4): the lines under the name, on the name's left edge, with the crest topped against the name block so the owner's extra line never moves it. Line 1, 13px/22px muted with the spec and class in `--text`: "Holy Priest · 0/31/20 · <Olympus XXVII> · Classic Beta PvP", each piece only when known, never a placeholder; the spec word and the talent split come from the same record, the newest build seen at pull. **Trailing separators on wrap:** each "·" attaches to the end of the segment before it (a no-break space, then the dot) and every segment is one `white-space: nowrap` unit, so the only breaks are the plain spaces between units: a wrap never starts a line with a dot and never leaves a dot on its own ("Holy Priest · 0/31/20 · <Olympus XXVII> ·" then "Classic Beta PvP"). **The realm is unbroken** (one nowrap unit that moves whole). The guild is the client's own "<Guild>", an underlined muted inline link, 44px tall on a phone from vertical padding, its focus ring at no offset; its unit is an inline-block at most the line's width, so it moves whole and wraps inside itself only when it is longer than the line. Line 2: "Level 60 Night Elf". **The chip on the level line:** the ruleset-and-region chip ("Hardcore US", muted pill, `--border`) ends the last text line, the level line whenever there is one, 8px after the text, top-aligned to the 22px line box so it is centred on the text, and glued to the last word by a word joiner so it never wraps alone. For the owner a third line, the Character header's sync line. **Faction from race for everyone:** the In game tooltip mock's faction line is derived from the race (the Classic race table in the addon's Rules.lua), so a visitor and the owner see the same mock; absent when the race is unknown. **The 12px sparkline caption** on the Forever Score card: "Up 330 since 13 Sep" at 12px in `--text` leads the date row and the newest date ends it in mono muted; with one point there is no line, only "First score recorded 12 Oct."; with fewer than five points the caption gains ", 2 readings" so a sparse straight line does not read as a steady climb.
- **Character card** (the planner and the BiS spec page): the Character header's stacked variant, `layout="card"` with `emphasis="secondary"`, at 440px on a tool page: identical identity and states, the action is the 44px outline gold button ("Send this build to the addon" on the planner, "Send this list to the addon" on the spec page), and the planner's has no answer slot. Replaces the older 40px-crest card with three stat tiles.
- **Pick row** (the best-in-slot list): slot label 90px; the pick as icon 40px with quality border, name in quality colour, source line and item level; runners-up as compact rows (icon 24px, name, source, mono delta "+0.7 DPS" or "same DPS"); a "You" column with the gain over what you wear or an "equipped" tick. Empty rows keep the slot label, a dashed 40px box and one plain sentence; never a coverage counter.
- **Upgrade row**: slot label, the worn item, an arrow, the best-in-slot item, and a mono gain in green; every item is icon (36px, 1px quality border), name in quality colour, and one source line under it. The table carries a label header row (Slot / You wear / Best in slot / Gain).

## Logo

The mark is the LX seal (`design/logo`, built by `build.py` from the real Cinzel outlines): the unit-frame level badge with sixty as the Roman numeral, a dark disc, a gold-gradient ring with a dotted inner track, and LX in Cinzel 800 gold. Header: seal at 34px beside the flat wordmark (Cinzel 700, 0.10em tracking, `--text-strong`); footer 22px; the favicon and the addon icon are the seal alone. Mono versions ink the whole seal in one colour for light or dark grounds. Never recolour, outline, stretch or add the gradient back to the wordmark.

## Icons

Inline SVG, 24px grid, 2px stroke, round caps and joins, `currentColor` or gold. No emoji anywhere in the UI.

## Motion

One reveal on page load (sky band fades in over 400ms), gold focus rings, 120ms hover transitions on cards and links, and one 160ms fade (`.reveal`) when an island's data lands in place of its skeleton. Nothing else animates. Every island reserves its ready height while loading (a `Skeleton` sized to the content, or a fixed `min-h`), so a fade is the only thing that changes; nothing moves. Reduced-motion turns the shimmer and the fade off.

## Light mode

Not planned. The audience plays at night beside a dark game window. If demand appears later, only the surfaces and text tokens change; accents, class, rarity, and pills stay as they are.
