// web/src/lib/levels.ts
// Rarity colour bracketing for level ranges, kept here rather than inline so it stays in
// sync with the thresholds below.
// design/DESIGN-SYSTEM.md: "Level-range tags reuse these: green for low brackets, blue for
// mid, purple for max-level."

/** Lowest `levelMin` that counts as a mid bracket (rare blue). */
export const MID_BRACKET_MIN_LEVEL = 30;
/** Lowest `levelMin` that counts as an end-game bracket (epic purple). */
export const HIGH_BRACKET_MIN_LEVEL = 55;

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
