// web/src/lib/levels.ts
// Level ranges and their rarity colours are shown on the dungeon and zone indexes and on
// every dungeon and zone page, so the brackets live here rather than in each route.
// design/DESIGN-SYSTEM.md: "Level-range tags reuse these: green for low brackets, blue for
// mid, purple for max-level."

/** Lowest `levelMin` that counts as a mid bracket (rare blue). */
export const MID_BRACKET_MIN_LEVEL = 30;
/** Lowest `levelMin` that counts as an end-game bracket (epic purple). */
export const HIGH_BRACKET_MIN_LEVEL = 55;

const UNKNOWN_SHORT = '—';
const UNKNOWN_LONG = 'Level range not yet known';

/** Table-cell form: `30–40`, or an em dash when the range is not known. */
export function levelRange(min?: number, max?: number): string {
  return min && max ? `${min}–${max}` : UNKNOWN_SHORT;
}

/** Sentence form for page leads and meta descriptions: `Level 30–40`. */
export function levelRangeLong(min?: number, max?: number): string {
  return min && max ? `Level ${min}–${max}` : UNKNOWN_LONG;
}

/**
 * Rarity colour token for a range, keyed on its lower bound. These land on text, so rare and
 * epic take the `-text` variants that clear WCAG AA on --color-raised; see tokens.css.
 */
export function levelColorClass(min?: number): string {
  if (!min) return 'text-muted';
  if (min >= HIGH_BRACKET_MIN_LEVEL) return 'text-rarity-epic-text';
  if (min >= MID_BRACKET_MIN_LEVEL) return 'text-rarity-rare-text';
  return 'text-rarity-uncommon';
}
