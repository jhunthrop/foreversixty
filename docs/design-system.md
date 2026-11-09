# The design system (round 5)

The web's shared layer: the tokens, the faces, the art and the components every page composes from. The look is the owner's approved redesign (`design/specs/2026-11-09-directions.md`, round 5 on round 4); the build plan is `docs/superpowers/plans/2026-11-09-redesign.md`. This file is the reference page lanes build from.

**The rule (owner, 2026-11-09: "we should use shared components and resources as a first class citizen where we can").** A page composes from the pieces below and adds no row, panel, title band, chart frame or colour of its own. When a page needs something new, it goes into the shared library first (a primitive in `web/src/components/ui/`, its props here, a mock primitive in `design/mocks/`), then the page uses it.

## Tokens

One source: `design/tokens/tokens.json`. `node design/tokens/build.mjs` writes `web/src/styles/tokens.css` (the Tailwind `@theme`), `apps/ios/Shared/DesignTokens.swift` and the Android `Tokens.kt`; `design/tokens/tokens.test.mjs` fails CI on a stale output.

- **The web leads.** A token whose web value differs carries it in `$extensions["gg.foreversixty"].platformValue.web`; the apps keep the shared `$value` until the mobile lane adopts the round-5 look and deletes the override. The round-5 groups the apps do not read yet (`color.forever`, `font.nav`, `radius.tooltip`) are `native: false`.
- **Raw colours are a test failure** outside `styles/` and `components/ui/` (`web/src/styles/raw-colours.test.ts`). Pages read a token: a Tailwind utility (`text-muted`, `bg-teal`, `border-white/20`), `var(--color-*)` or a primitive.
- **Surfaces** are drawn only by the primitives (`web/src/components/surface-guard.test.ts`).

### Colour

Blizzard's Forever page, exactly (c computed, s sampled; the spec's section 2 has every source).

| Token (CSS) | Value | Use |
|---|---|---|
| `--color-ground` | `#67563a` | The page ground, under the painted earth tile |
| `--color-raised` | `#170e09` | Every panel, the card brown-black |
| `--color-foot` | `#170e09` | The footer |
| `--color-bg` | `#100a06` | Wells and inputs inside a panel |
| `--color-hover`, `--color-pressed`, `--color-selected` | white lifts on the card | Row states |
| `--color-band` | `#054562` | The section band, under the painted band |
| `--color-nav-bar` | `#0b2f3b` | The floating nav bar |
| `--color-night` | `#071d2e` | Under home's hero art |
| `--color-cream` | `#f3eee2` | Page and section headings |
| `--color-cream-2` | `#fcedcf` | The footer name |
| `--color-text` | `#e3dfd9` | Body ink |
| `--color-strong` | `#ffffff` | Names, figures |
| `--color-muted` | card `#8b8784`, ground `#e2dcd3` | Secondary ink; the body sets the ground's, every surface restores the card's (`ui/panel.css`) |
| `--color-gold`, `--color-gold-hover` | `#ffec71`, `#f7b90a` | Links |
| `--color-teal`, `--color-teal-edge` | `#0f6064`, `#408076` | The primary button; Premium is the edge colour flat |
| `--color-blue`, `--color-blue-edge` | `#114f72`, `#1d6a96` | The secondary button |
| `--color-accent` | `#0093b2` | The current tab's underline, selection, chart accent, bars |
| `--color-gain` | `#32ae88` | A DPS gain |
| `--color-panel-edge`, `--color-frame-line` | warm ring, cream at 42 % | The panel's outer ring, the framed panel's inset line |
| `--color-veil-band`, `--color-veil-night`, `--color-veil-cave` | flat veils | Over the band, the hero and the raid cave |
| `--color-tooltip`, `-tooltip-edge`, `-tooltip-inner`, `-tooltip-yellow` | the client's | The client tooltip only |

Class, faction, rarity, parse, school and outcome colours are unchanged. `tokens.test.ts` holds text, muted, gold and cream to AA on the card and on the ground's darkest and lightest 200 px windows, and cream on the band.

### Type

| Role | Face | Token |
|---|---|---|
| Body, labels, names, figures | Open Sans 400/600/700 | `--font-body` |
| Page and section headings (h1, h2, every `Section` title), the bar's name | Marcellus 400, cream, tracking -0.015em | `--font-title` (`--font-heading` aliases it) |
| The navigation bar's doors | Archivo 600, 17 px | `--font-nav` |
| Literal code only | JetBrains Mono 500 | `--font-mono` |

Sizes: the title band's title is `--text-title` 26/30 (21/25 below 768), a section head `--text-section` 22/28, a reading section head `--text-section-reading` 24/30 (22/28 on a phone). An h3 is a row's or block's title and stays the body face's bold. Labels (`text-label`, 12 px) are sentence case in the body face: no uppercase eyebrow anywhere. A number takes its unit and nothing else (ai-slop tell 22). The subsets are in `web/src/assets/fonts/` (`web/fonts/README.md`).

### Radius

`--radius-plate` 10 (the nav, every panel, overlays), `--radius-field` 6 (buttons, controls, the selected row), `--radius-icon` 3 (item icons), `--radius-tooltip` 4 (the client tooltip only).

## Art and seals

One manifest: `web/src/lib/ui/assets.ts`. The files are in `web/public/art/`, written by `python3 tools/art/build_art.py` from the owner's sources in `design/assets/` (the ground tile comes from the mock pipeline's `design/mocks/directions/art.py`, so mocks and site share it). Generated art is the owner's choice: background only, never captioned as a screenshot.

| Export | What |
|---|---|
| `GROUND_TILE`, `GROUND_TILE_PX` | The painted earth, periodic, tiled at 1:1 (the body's background) |
| `BAND_ART`, `HERO_ART`, `CAVE_ART` | The blue band, home's camp road, the raid cave; each `{ phone, desktop }` |
| `OG_ART` | The link preview's background (`lib/og.ts`) |
| `SEAL`, `sealFor(size)` | The flat seal at 48 px and under, the painted seal above |
| `PAPERDOLL_FRAME` | The paperdoll frame at 1x and 2x, with its size |
| `FAVICON` | The favicon set at the site root |
| `ART_CSS_VARS`, `artStyle()` | The art as `--art-*` custom properties, set once on `<html>` by `Base.astro` so stylesheets never name a file |

## Shell

- **`layouts/Base.astro`**: the earth ground on `<body>`, the band (`.site-top`) the bar floats over, the faces preloaded. Props: `title`, `description`, `path`, `jsonLd?`, `wide?`, `art?: 'band' | 'hero'` (home passes `hero`, with PageHeader's `art="hero"`).
- **`components/Header.astro`** (styles in `styles/nav.css`; doors in `lib/nav.ts`): the floating rounded bar, 64 px, inset to the content edge: the flat seal and "Forever Sixty" in Marcellus, then the mock's flat doors in Archivo: Best in slot, Guides, Tier list, Sim, Planner, Logs, Guild. A door is a plain link (`door(label, href)`) unless its group has sub-pages (`panel: true`: Guides by class, Logs, Guild), which open a chevron panel. The current door carries a 3 px cyan underline. Then search, the one-line account control (`nav/CharacterSelector`: medallion, name in class colour, short line, chevron) and Premium last, as their teal Subscribe (`.site-door-premium`, `data-testid="nav-premium"`). "Get set up" is not a door: it is a row of the account menu (`nav/SelectorPanel`, signed in or out) and a footer link. Fit: from 1280 every piece shows; 1024 to 1279 the bar drops the name beside the seal (selector 184 px, name 15 px); the account control is a fixed slot, its label and chevron together at the left, the level beside the name; the selector's short line and the "/" search hint show only on the wide column's bar (from 1760, selector 264 px), where they fit whole. Below 1024 the bar is 56 px with the seal, name, account medallion and Menu; the menu sheet lists the same doors as 44 px rows and Premium as a full-width row.
- **`components/Footer.astro`** (copy in `lib/footer-copy.ts`): the compact footer on `--color-foot`. Row one: the painted seal (88 px from `lg`, 64 below) beside the name in Marcellus, and the site links (`nav[aria-label=Footer]`: Get set up, Premium, Join the Discord, Report a bug). Row two: the legal links (Terms, Privacy, Refunds), the not-affiliated line and "© Commish LLC". No sitemap columns; `FOOTER_COLUMNS` survives only as the 404 page's suggestions.

## Components

All in `web/src/components/ui/` unless named otherwise. Props marked `?` are optional.

### Bands and heads

- **`PageHeader`** (the compact title band): `title`, `tone?` (`'strong'` paints cream; `{ classSlug }`, `{ faction }`), `crest?` snippet (the medallion), `meta?` snippet (one fact line), `lead?`, `answer?` snippet (the figure and its unit at the right), `actions?`, `tabs?` (under a rule on the band), `art?: 'band' | 'hero'`, `minRowHeight?`, `titleTestid?`, `testid?`. It continues the bar's band (`ui/band.css`); its black rule hands over to a ControlBar under it.
- **`ControlBar`**: the page's controls as nav-tab groups on the band (unchanged props). Groups are `Segmented` look: labels white at .72, the current one white over a 3 px cyan underline.
- **`Band`**: a band of a page's own: `art?`, `y?` (its top in the page, so the image continues), `testid?`, children.
- **`Page`** (the layouts): `split` puts a 340 px rail beside main from 1024 (`--rail-w` from 1760); `splitFrom?: 1024 | 1280` (split only, 1024 by default) is the width the two sit side by side from, and `railFirst?` (split only) puts the rail above main where they stack (the tier list: side by side from 1280, its spec panel first below).

### Panels

- **`Plate`** (the page's primary object on a panel): `frame?: 'plain' | 'framed' | 'lead'`, `flush?` (default true), `label?`, `testid?`. `framed` adds the cream inset line, `lead` the line and corner flourishes.
- **`Frame`** (a featured panel with its own padding): `ornate?` (the page's one lead panel), `art?: 'cave'` (the raid plaque's cave under its veil), `label?`, `testid?`, `region?` (the parity region it carries: the tier list's spec panel is `REGION.headAnswer`, the page's answer, in the rail).
- **The rule:** two or three framed panels a page, flourishes on the lead only. Home: the raid plaque (lead), Your upgrades (framed). BiS: the paperdoll (lead). The tier list: its spec panel, framed (the mock draws it plain, no flourishes).
- **`Section`**: a Marcellus heading on the ground (`title`, `level?`, `meta?`, `action?`, `head?`, `reading?`, `id?`, `anchor?`, `testid?`, `focusId?`, `titleTestid?`) and its body. **`boxed`** puts the body on the panel 12 px under the heading, as every list in the mocks sits: class and quality colours only clear contrast on the panel, so page lanes box every list, table and row set (and pass `boxed=True` to the mock's `Section`). A boxed body that already holds a panel (a Plate, a grid of panels) is not boxed twice. Open (the default until each page moves its sections), the body sits on the ground under the heading's hairline: prose only.
- **`Section panel`** (the logs, report and rankings mocks' `prims.panel`): the heading row inside the panel. `panel` wraps the head and the body in one `Plate`, so the surface, the frame line and the corners are Plate's; `frame?: 'plain' | 'framed' | 'lead'` is the Plate's frame (the page's first panel `framed`, its one lead panel `lead`); `flush` runs the body edge to edge for a table that pads its own cells, else the body is padded 20 px (14 on a phone). The region stays `section-<slug>`, with a `plate` region inside it. The mock twin is `ws_prims.Section(..., panel=True, frame=..., flush=...)`.
- **`Page panels`**: main and rail hold Section panels: `--section-gap` is 22 px inside the Page, the columns 28 px apart (mock: `ws_prims.Page(..., panels=True)`).
- The surface classes live in **`ui/panel.css`**: `.ui-surface` (the panel), `.ui-framed`, `.ui-lead`; **`FrameCorners`** draws the flourishes.

### Controls

- **Buttons** (`lib/ui/button.ts`, `Button`): `tone: 'primary' | 'secondary' | 'ghost' | 'danger'`, `size: 'md' | 'sm'`. Primary is their flat teal with the teal edge, secondary their blue, ghost a gold text link, danger the wipe outline; 15 px 600, sentence case, the 6 px radius. One primary per view. `pressed` (aria-pressed) marks the current choice of a set with a white edge (home's I'm in / Late / Can't come).
- **`Segmented`**, **`TabStrip`**, **`Select`**, **`Switch`** (on is the teal), **`Field`**, **`Tag`** (`tone: 'gold' | 'success' | 'muted' | 'danger' | 'warn'`, `dot?`).
- **One tab look.** `TabStrip` (in-page and link tabs) and `Segmented` (ControlBar groups, in-row choices) draw the same tab: label-wide items (12 to 14 px each side, 4 px apart; phones share a Segmented row evenly), the current one white over the cyan `--color-accent` underline, the rest white at .72. A filter that is on shows a count, never a different fill. Selects are the only boxed controls on a band.
- **`Meter`**: `tone?: 'accent' | 'gold' | 'kill' | 'muted' | 'class'` (accent by default), `color?`, `segments?`, `label?`.

### Game surfaces

- **The client tooltip**: `ItemTooltip`, `GameTooltip`, on `Overlay` `kind="tooltip"`: the client's ground, edge, inner line and 4 px corner; descriptions and the gain line in the client yellow.
- **`ItemIcon`**, **`ItemRow`**, **`Row`**: real icons with the quality border on the 3 px icon radius; with no icon, a dark socket (never a dashed box); a selected row is the white lift and a cyan bar.
- **`PaperdollFrame`** (on the owner's frame, `lib/ui/paperdoll-frame.ts` reads `paperdoll-slots.json`): `cells` (the `DollCell[]` the grid paperdoll takes), `build`, `title`, `meta?`, `figure?: { value, unit }`, `avatar?` (the Battle.net `-avatar.jpg`, else the painted seal, never the class crest), `heading?: 'h1' | 'h2'`, `selected` (bindable), `onselect?(slot)`, `centre?` snippet (the selected slot's tooltip), `label?`, `testid`. Under 423 px it is the one-line strip and a two-column list, and the frame image is never fetched. Regions `paperdoll-medallion` and `paperdoll-centre`.
- **`Countdown`**: the countdown panel (the guild's first-raids card). Home's raid plaque is an ornate `Frame art="cave"` holding **`BigFigures`** (`groups` of `{ value, label }`, a hairline between groups, `timer?` naming the first group as one live clock): days and hours at 44 px, then signed up and late.
- **`FeedRow`**: a dated feed line (`lead`, the sentence, `when`, `kind?`: the tag under the time, left out on home).
- **The ranked list** (`RankList`, `RankGroup`, `RankRow`, `ColumnHeads`; `lib/ui/rank-list.ts`): one grid for the heads and every row. `RankList` takes `heads?`, `figures` (1 to 4), `figureWidth?`, `links` (0 to 4), `leads?`, `grouped?`, `label?`, and `compact?`: round 5's compact list (the tier list beside its rail): elastic name and bar tracks that fit a split page's main column, 48 px rows, the letter column at every width (44 px, 64 px from 1280) with the letter alone in the title face, the accent tint and a cyan rail slice on the selected row, the one-line row and its heads from 1024, and below 768 rank, crest, name and figure only. `RankGroup` takes `letter`, `label` and, outside a compact list, `tone`, `band` and `count` (each optional: left out, the column carries the letter alone). `RankList` `phoneFigures?: 1 | 2` (compact only; the log rankings) keeps that many extra figures beside the sub line below 768, so a phone row stays two lines. `RankRow` takes `barTone?: 'class' | 'accent'` (class by default; the tier list's relative bar is the accent), and a figure's `unit` is optional where the column head names it. Without `compact` and `barTone` every list draws as before.

### Charts

`web/src/lib/chart/theme.ts` is the one chart theme: `CHART_INK` roles (`grid`, `frame`, `label`, `accent`, `selection`, `highlight`, `surface`, `edge`, ...) as `{ token, alpha }`, `inkCss(ink)` for CSS and SVG, `inkCanvas(el, ink)` for canvas (it resolves the token through `getComputedStyle`), `chartFont(el, px, weight)` from `--font-body`, and `CHART_ACCENT`. Every report, sim and character chart draws through it; no chart names a colour or a font.

## Link preview

`lib/og.ts`: the owner's road art under the night veil, the painted seal and "Forever Sixty", the page's kicker and title in Marcellus, and "Fan-made player tool for WoW: Forever" (`OG_FAN_LINE`), so a shared link never reads as an official page.

## Parity

The parity references (`design/mocks/ws_*.py`) draw the same system: `design/mocks/ws_lib.py`'s `ROUND5_CSS` comes last in every mock and gives each shared mock primitive the round-5 type, sizes and look; `web/tests/e2e/support/parity.ts` serves the mocks the site's own faces. The header and footer in the mocks are captured from the build (`CHROME_CAPTURE=1 npx playwright test tests/e2e/chrome-capture.spec.ts --project=desktop`). The visual reference for each page is its directions mock in `design/mocks/directions/`.
