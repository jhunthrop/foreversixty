// web/src/lib/chart/palette.ts
// Series colours (controls and charts spec §B5.5). A player is their class token; enemies and
// abilities take this palette in order, and a seventh line repeats it with a 4/3 dash. Every
// colour is read from the tokens at draw time, never a hex in chart code.

export const ENTITY_PALETTE = [
  'var(--color-gold)',
  'var(--color-school-frost)',
  'var(--color-school-nature)',
  'var(--color-school-arcane)',
  'var(--color-school-fire)',
  'var(--color-school-shadow)',
] as const;

/** The n-th entity line's colour and dash. */
export function entityStyle(index: number): { colour: string; dash?: readonly number[] } {
  const colour = ENTITY_PALETTE[index % ENTITY_PALETTE.length];
  return index >= ENTITY_PALETTE.length ? { colour, dash: [4, 3] } : { colour };
}

/** A `var(--token)` (or any CSS colour) resolved against an element's computed style. */
export function resolveColour(element: Element, colour: string): string {
  const token = /^var\((--[\w-]+)\)$/.exec(colour.trim())?.[1];
  if (token === undefined) return colour;
  return getComputedStyle(element).getPropertyValue(token).trim() || colour;
}

/** A resolved colour at an alpha: hex (`#rrggbb`) or `rgb()` in, `rgba()` out. */
export function withAlpha(colour: string, alpha: number): string {
  const hex = /^#([0-9a-f]{6})$/i.exec(colour)?.[1];
  if (hex !== undefined) {
    const [r, g, b] = [0, 2, 4].map((at) => Number.parseInt(hex.slice(at, at + 2), 16));
    return `rgba(${r}, ${g}, ${b}, ${alpha})`;
  }
  const rgb = /^rgba?\(([^)]+)\)$/.exec(colour)?.[1];
  if (rgb !== undefined) {
    const [r, g, b] = rgb.split(/[\s,/]+/).filter(Boolean);
    return `rgba(${r}, ${g}, ${b}, ${alpha})`;
  }
  return colour;
}
