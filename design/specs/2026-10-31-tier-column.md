# Tier list: the tier column (owner fix, 2026-10-31)

Amends design/specs/2026-10-09-tier-list.md, "tier chip". Surfaces: `web/src/components/tiers/TierPage.astro` (`.tier-group`, `.tier-chip`) and `TierRow.astro` (`.tier-row`). Nothing else on the page changes.

Mocks: `design/mocks/2026-10-31-tier-column/` (`gen.py`; before-/after- at 1440, 2000, 360, now S/B/C/D with the 360 you-row; `after-horde-1440.png` shows S, A and B; `after-1440-hover.png` shows hover). v1 abutted a 3px edge at x=132 and was a NO-SHIP from wow-player; v2 is below.

## 1. The defect

The chip is `#0a0d15`, darker than the panel (`--color-raised` #0d111a) and bluer than any token. With a 3px tier rail on its left and a 1px right border it reads as a foreign dark block. The "Your spec" gold edge (2px) starts at x=132, where that block ends, so two edges sit in two places with two widths.

Reference: the in-game raid frame group header and the Wowhead tier-list label cell: the label is the same surface as its rows, with one coloured edge.

## 2. Decision (v2, after the wow-player NO-SHIP)

One surface, one rail column, one edge column. The chip is the panel with the tier colour at 3%. The rail lives at x=0 and is the only left edge on the list. A row's state (hover, "Your spec") does not draw its own edge: it recolours that row's slice of the rail.

| Piece | Owner | x | Width | Colour |
|---|---|---|---|---|
| Tier rail | `.tier-chip` `border-left`, full group height | 0 | 3px | `var(--tier)` |
| Row state slice | `.tier-row::before`, the row's full height, painted over the rail | 0 on desktop (`left: calc(-1 * var(--tier-chip-width))`), 0 of the row on phone | 3px | transparent at rest; `var(--c)` on hover; gold on "Your spec" |

- The rows have no `border-left` at all. Their content x is the shipped 16px (desktop) and 14px (phone).
- The 6% gold tint stays on the row only (from x=148). It does not extend across the chip: the chip is shared by the whole group.
- On the lists with no chip (tank, healer) the same `::before` sits at the row's own x=0.
- On phone the chip is a header strip above the rows, so the slice is at the rows' x=0, the same column as the strip's rail.
- `--tier-chip-width` goes 132px to 148px so "Within 5% of the top" and "30% or more behind" stay on one line (11px letter-count line unchanged). `.tier-heads` padding derives from it.

## 3. CSS delta from v1 (the built state)

`TierPage.astro`

```css
.tier-list { --tier-chip-width: 148px; }                                  /* was 132px */
.tier-chip { background: color-mix(in srgb, var(--tier) 3%, var(--color-raised)); }   /* was 5% */
```

`TierRow.astro`

```css
.tier-row {
  border-left: 0;                          /* v1: 3px solid transparent; original: 2px */
  padding: 14px 16px 14px 16px;            /* v1: 13px left */
}
.tier-row::before {                        /* new */
  content: '';
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  width: 3px;
  background: transparent;
  transition: background-color 120ms ease-out;
  pointer-events: none;
}
:global(.tier-list[data-role='dps']) .tier-row::before {
  left: calc(-1 * var(--tier-chip-width));
}
.tier-row:hover { background: var(--color-card-top); }           /* border-left-color line deleted */
.tier-row:hover::before { background: var(--c); }
.tier-row.is-you,
.tier-row.is-you:hover { background: color-mix(in srgb, var(--color-gold) 6%, transparent); }
.tier-row.is-you::before,
.tier-row.is-you:hover::before { background: var(--color-gold); }
.tier-row.is-you:hover { background: color-mix(in srgb, var(--color-gold) 10%, transparent); }  /* as v1 */
```

Phone media block: revert v1's `padding: 12px 14px 0 11px` to `12px 14px 0 14px` and the link strip margin `0 -14px 0 -11px` to `0 -14px 0 -14px`. `.tier-row` is already `position: relative`. The slice is inside `.tier-panel` (`overflow: hidden`), so nothing clips; the chip is not positioned, so the slice paints above it.

## 4. Legibility (letter, count, band text on the 3% tint, WCAG)

Contrast only improves over the 5% figures below (a lighter tint is closer to the panel). Computed at 5%: letters 5.4:1 (A, the worst) to 12.7:1 (C), count 5.8:1 or better, band text 9.8:1 or better. All pass AA. At 8% the S chip turned brown and read as a block; at 3% it is within a shade of the rows.

## 5. States

- Rest: as above. Chip is not interactive (no hover, focus or active state; nothing to add).
- Row hover: card-top bg, class-colour 3px edge, bar `brightness(1.18)` (unchanged).
- Row focus-visible: links keep the 2px gold outline, 2px offset (unchanged).
- You-row: gold 3px edge plus 6% gold tint plus the "Your spec" pill. Never colour alone.
- Loading, empty, error: unchanged (skeleton rows and the panel-held error have no chip).
- Touch: no hover on touch; the gold edge and tint carry "you" alone.

## 6. Acceptance screenshots

Built page, signed in as a Fury Warrior and signed out:
1. 2000 and 1440: S, B, C, D (Alliance) and S, A, B (Horde), with the you-row inside S. The chip reads as the same surface as the rows; the gold slice is on the rail, level with the you-row.
2. 1099 (chip is a header strip) and 1100 (chip is a column): both show the same tint and 3px edges.
3. 360: S header strip, you-row; rail and gold edge on one x column.
4. Hover on a non-you row at 1440 (card-top bg, class-colour slice on the rail) and on the you-row (gold stays).
5. 360 with the you-row visible: gold slice at x=0.
6. Tank and healer tabs at 1440: slice at the row's own x=0, no layout shift.
7. Grep: no `#0a0d15` or `#101624` left under `web/src/components/tiers`.

Verdict rule: any dark strip next to the rows, or any row edge not at the rail's x, any width other than 3px, is NOT SHIPPABLE.
