// web/src/lib/sim/history.ts
// The saved-sim list, by kind. The API composes each row's headline (contract 8: run ->
// "1,204 DPS"; gear -> "+41 DPS from Vis'kag"; drops -> "3 upgrades on Ragnaros"), because
// only the API has the result blob in front of it -- the list route does not send one. The
// web renders whatever sentence arrives and falls back to the figure for a row saved
// before the column existed.
import { formatAmount } from '../report/format';
import { SIM_KINDS, type SimKind } from './kind';
import { specDisplayName } from './spec-label';
import type { SimListRow } from './types';

export type KindFilter = SimKind | 'all';

export const KIND_FILTERS: readonly KindFilter[] = ['all', ...SIM_KINDS];

/** A row's kind. Anything outside the vocabulary -- including absent -- reads as a run. */
export function kindOf(row: SimListRow): SimKind {
  return SIM_KINDS.find((kind) => kind === row.kind) ?? 'run';
}

/** The API's own sentence, or the DPS figure for a row saved before it composed one. */
export function headlineOf(row: SimListRow): string {
  const headline = row.headline ?? '';
  return headline === '' ? `${formatAmount(Math.round(row.dps))} DPS` : headline;
}

/** The player's own title, or the bare spec name when they never gave one. */
export function titleOf(row: SimListRow): string {
  return row.title === '' ? specDisplayName(row.spec) : row.title;
}
