// web/src/lib/ui/cells.ts
// The one vocabulary of a coverage cell (web layout spec 2026-11-05 §7.3), shared by ui/Grid,
// ui/Tally and ui/Checklist: met, close, missing, or nothing to count. The primitives turn a tone
// into a colour; pages pass the tone only.

export type CellTone = 'ok' | 'warn' | 'miss' | 'none';

/** A tone with something to count (a Tally row, a Checklist mark). */
export type MarkTone = Exclude<CellTone, 'none'>;

/** The text colour of a tone (outcome tokens). */
export const CELL_TEXT: Readonly<Record<CellTone, string>> = {
  ok: 'text-kill',
  warn: 'text-gold',
  miss: 'text-wipe',
  none: 'text-muted font-normal',
};

/** A Checklist mark's ground and its glyph. */
export const MARK: Readonly<Record<MarkTone, { ground: string; glyph: string; label: string }>> = {
  ok: { ground: 'bg-kill', glyph: '✓', label: 'Done' },
  warn: { ground: 'bg-gold', glyph: '!', label: 'To do' },
  miss: { ground: 'bg-wipe', glyph: '×', label: 'Missing' },
};

/** At the target, within a quarter of it, or short. */
export function tallyTone(have: number, want: number): MarkTone {
  if (have >= want) return 'ok';
  return have >= want * 0.75 ? 'warn' : 'miss';
}

/** One Grid cell: its tone, the text it shows, and the words a screen reader hears for it. */
export interface GridCell {
  tone: CellTone;
  text: string;
  label?: string;
}

/** One Grid row: a label (and a muted sub line) and one cell per column. */
export interface GridRowModel {
  key: string;
  label: string;
  sub?: string;
  cells: readonly GridCell[];
}
