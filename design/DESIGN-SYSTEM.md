# Forever Sixty design system

Derived from the approved "Cinematic" direction and the live-game homepage mockup in this folder. Source of truth for tokens once the site exists is `web/src/styles/tokens.css`; keep this document and that file in sync.

## Principles

1. **Reference, not pitch.** No hero slogans, no calls to action in marketing voice. The front door is a search box and the current state of the game.
2. **Second screen first.** Dark by default, high contrast, 44px minimum hit targets, pages that paint before the player alt-tabs back.
3. **Every fact is dated and sourced.** Source pills (Blizzard, Datamined, Community, This site) and "updated" stamps are part of the UI, not an afterthought. Single-source claims say so.
4. **The game's own colors do the wayfinding.** WoW class colors and item-rarity colors are used consistently and never repurposed.
5. **Atmosphere in the header only.** The night-sky band lives at the top of a page. Content areas are flat, calm, and dense enough to be useful.
6. **Ornament stays out.** Warmth comes from Cinzel at small sizes, gold accents, and the sky. No stone frames, no parchment textures, no beveled buttons.

## Color

### Surfaces

| Token | Value | Use |
|---|---|---|
| `--bg` | `#07090d` | Page background |
| `--bg-raised` | `#0d111a` | Panels, table bodies |
| `--bg-card-top` | `#131824` | Card gradient start (cards go `#131824` to `#0d111a`) |
| `--border` | `#262e40` | Panel and card borders |
| `--border-soft` | `#1c2230` | Row dividers |
| `--border-warm` | `#3a3326` / `#4a4030` | Inputs and buttons on the sky band |

### Text

| Token | Value | Use |
|---|---|---|
| `--text` | `#e9e4d8` | Body |
| `--text-strong` | `#f2eee4` | Headings, primary labels |
| `--text-muted` | `#9a9484` | Secondary text (6.3:1 on raised) |
| `--text-nav` | `#b9b3a4` | Nav links, footer links |

### Accent

| Token | Value | Use |
|---|---|---|
| `--gold` | `#e5b955` | Links, focus, live indicator, primary accent |
| `--gold-hover` | `#f5d27a` | Link hover |
| `--gold-deep` | `#a8762a` | Gradient end for gold text and bars |
| `--gold-light` | `#fbe7a1` | Gradient start for gold text |
| `--ember` | `#d66e28` | Sky horizon glow only |
| `--night` | `#26405c` | Sky zenith only |

Gold text is a gradient (`#fbe7a1` → `#e5b955` → `#a8762a`, top to bottom) clipped to text. Use it for the wordmark and section titles only.

### Source pills

| Pill | Text | Background | Border |
|---|---|---|---|
| Blizzard / Hotfix | `#6fb1ff` | `rgba(0,112,221,.18)` | `rgba(0,112,221,.35)` |
| Datamined | `#c98bff` | `rgba(163,53,238,.16)` | `rgba(163,53,238,.35)` |
| Community | `#7bff5c` | `rgba(30,255,0,.10)` | `rgba(30,255,0,.25)` |
| This site | `#e5b955` | `rgba(229,185,85,.14)` | `rgba(229,185,85,.35)` |
| Sample | `#9a9484` | `rgba(154,148,132,.12)` | `rgba(154,148,132,.30)` |

### Faction

Alliance `#6fb1ff` (text) / `#2f6fd6` (bars). Horde `#ff6b5c` (text) / `#c0392b` (bars).

### Class colors (WoW standard)

Warrior `#c69b6d` · Paladin `#f48cba` · Hunter `#aad372` · Rogue `#fff468` · Priest `#ffffff` · Shaman `#0070dd` (use `#3f8fe0` for text on dark) · Mage `#3fc7eb` · Warlock `#8788ee` · Druid `#ff7c0a` · Monk `#00ff98`.

### Item rarity (WoW standard)

Poor `#9d9d9d` · Common `#ffffff` · Uncommon `#1eff00` · Rare `#0070dd` · Epic `#a335ee` · Legendary `#ff8000`. Level-range tags reuse these: green for low brackets, blue for mid, purple for max-level.

Rare and epic are too dark to read as small text on `--color-raised` `#0d111a`: 3.92:1 and 3.87:1, under WCAG AA's 4.5:1 for 13–14px semibold. Text uses two lightened variants instead — rare `#3d94f0` (6.02:1) and epic `#b866f5` (5.63:1) — the same move the Shaman class colour makes. The base colours above are unchanged and stay in use for bars, borders and icons; poor (6.96:1), common (18.88:1), uncommon (13.81:1) and legendary (7.50:1) already clear AA and are used as-is for text.

## Typography

| Role | Face | Fallback | Notes |
|---|---|---|---|
| Display | Cinzel 600–800 | Trajan Pro, Georgia, serif | Wordmark 20px, section titles 18px uppercase with 0.10em tracking, card titles 15–17px. Never above 22px on content pages. |
| Body | Barlow 400–700 | Helvetica Neue, Arial, sans-serif | 14–15px body, 13px secondary, 17px search placeholder |
| Numbers | JetBrains Mono 500 | SF Mono, Menlo, monospace | Dates, timers, counts, keyboard hints; `font-variant-numeric: tabular-nums` |
| Labels | Barlow 700 | | 11px, uppercase, 0.14em tracking, muted or gold |

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
- **Character row**: class-colored 36px square, name in class color, muted descriptor, right-aligned mono stat, progress bar.
- **Class tile**: 84px tall, 2px inset class-color top stroke, class name in class color.
- **Progress bar**: 6px, `--border-soft` track, filled with the relevant accent.
- **Secondary button**: 36px, warm border, uppercase 12px 700 with 0.06em tracking. There is no primary marketing button.

## Icons

Inline SVG, 24px grid, 2px stroke, round caps and joins, `currentColor` or gold. No emoji anywhere in the UI.

## Motion

One reveal on page load (sky band fades in over 400ms), gold focus rings, 120ms hover transitions on cards and links, and one 160ms fade (`.reveal`) when an island's data lands in place of its skeleton. Nothing else animates. Every island reserves its ready height while loading (a `Skeleton` sized to the content, or a fixed `min-h`), so a fade is the only thing that changes; nothing moves. Reduced-motion turns the shimmer and the fade off.

## Light mode

Not planned. The audience plays at night beside a dark game window. If demand appears later, only the surfaces and text tokens change; accents, class, rarity, and pills stay as they are.
