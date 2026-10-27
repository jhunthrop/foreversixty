// web/src/lib/guild/figures.ts
// Figures in mono, words in the body face (attendance spec 3.1, 3.3): a copy string split into its
// numerals and the words between them, so a line like `2 h 41 min` sets only `2` and `41` in mono.

export interface FigurePart {
  text: string;
  figure: boolean;
}

/** A numeral with its sign and separators: `4:12`, `+53`, `−12`, `1,234`. */
const FIGURE = /([+−]?\d[\d:.,]*)/;
const STARTS_WITH_FIGURE = /^[+−]?\d/;

export function figureParts(text: string): FigurePart[] {
  return text
    .split(FIGURE)
    .filter((part) => part !== '')
    .map((part) => ({ text: part, figure: STARTS_WITH_FIGURE.test(part) }));
}
