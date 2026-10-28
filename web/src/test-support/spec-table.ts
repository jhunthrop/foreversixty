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

/**
 * The backticked segments of the String column (the third) of a three-column `Where | Key | String`
 * table, for the row whose Key column lists `key` as one of its backticked keys (an exact match,
 * so `trial` never reads the `trialEmail` row). The search stops at the next `### ` heading.
 */
export function specKeyCell(specPath: URL, heading: string, key: string): string[] {
  const spec = readFileSync(specPath, 'utf8');
  const start = spec.indexOf(heading);
  if (start === -1) throw new Error(`spec has no heading "${heading}"`);
  const rest = spec.slice(start + heading.length);
  const end = rest.search(/^### /m);
  const section = (end === -1 ? rest : rest.slice(0, end)).split('\n');
  const row = section.find((line) => backticked(line.split('|')[2] ?? '').includes(key));
  if (row === undefined) throw new Error(`spec "${heading}" has no row keyed "${key}"`);
  return backticked(row.split('|')[3] ?? '');
}

function backticked(cell: string): string[] {
  return [...cell.matchAll(/`([^`]+)`/g)].map((match) => match[1] ?? '');
}

/** Replaces every `{slot}` in a spec string with its value. */
export function fillSlots(text: string, slots: Readonly<Record<string, string>>): string {
  return Object.entries(slots).reduce((t, [slot, value]) => t.replaceAll(`{${slot}}`, value), text);
}

/** The backticked line that follows a bold paragraph label, e.g. `**Held until built`. */
export function specLineAfter(specPath: URL, label: string): string {
  const lines = readFileSync(specPath, 'utf8').split('\n');
  const at = lines.findIndex((line) => line.startsWith(label));
  if (at === -1) throw new Error(`spec has no paragraph "${label}"`);
  const line = lines.slice(at + 1).find((l) => l.startsWith('`'));
  if (line === undefined) throw new Error(`spec paragraph "${label}" has no backticked line`);
  return line.replace(/^`|`$/g, '');
}
