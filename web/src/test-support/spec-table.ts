// web/src/test-support/spec-table.ts
// Reads a copy table out of a design spec so a string test pins the code to the spec itself:
// a drift in either file fails. The table is the markdown table under a heading; a row is found
// by its first-column label and its strings are the backticked segments of the second column.
import { readFileSync } from 'node:fs';

/** The backticked segments of a table row's string cell, by the row's first-column label. */
export function specCell(specPath: URL, heading: string, labelStart: string): string[] {
  const spec = readFileSync(specPath, 'utf8');
  const start = spec.indexOf(heading);
  if (start === -1) throw new Error(`spec has no heading "${heading}"`);
  const rows = spec.slice(start).split('\n');
  // An exact label wins, so `Remove` is not read as `Remove question, connected`.
  const row =
    rows.find((line) => line.startsWith(`| ${labelStart} |`)) ??
    rows.find((line) => line.startsWith(`| ${labelStart}`));
  if (row === undefined) throw new Error(`spec "${heading}" has no row starting "${labelStart}"`);
  const cell = row.split(/\|/)[2] ?? '';
  return [...cell.matchAll(/`([^`]+)`/g)].map((match) => match[1] ?? '');
}

/** One backticked segment of a row, with `{placeholder}` markers filled in. */
export function specString(
  specPath: URL,
  heading: string,
  label: string,
  index = 0,
  fill: Readonly<Record<string, string>> = {},
): string {
  const segment = specCell(specPath, heading, label)[index];
  if (segment === undefined) throw new Error(`spec "${heading}" row "${label}" has no segment ${index}`);
  return Object.entries(fill).reduce((text, [key, value]) => text.replace(`{${key}}`, value), segment);
}
