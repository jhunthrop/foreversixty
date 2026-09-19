// web/src/lib/sim/url.ts
// /sim's URL state. The page is static and the island reads the query, exactly as the
// planner reads ?class= and the report reads ?fight=, so a link to "sim this fight" or
// "sim this build" is a plain URL a player can paste.
//
// Everything here is validated against a fixed vocabulary: the query is attacker-controlled
// and the values reach fetch() paths.
import type { SourceKind } from './types';

const SOURCES: readonly SourceKind[] = ['armory', 'addon', 'build', 'fight', 'manual'];
const MODES = ['sim', 'compare'] as const;
export type SimMode = (typeof MODES)[number];

/** A build id is 12, a character key about 40, a fight ref 14. 128 is generous and finite. */
const MAX_REF = 128;

export interface SimState {
  source: SourceKind | '';
  ref: string;
  mode: SimMode;
  /** The fight compare mode is comparing against: "<report_id>:<fight_index>". */
  fight: string;
}

export function defaultSimState(): SimState {
  return { source: '', ref: '', mode: 'sim', fight: '' };
}

function bounded(value: string | null): string {
  if (value === null || value.length > MAX_REF) return '';
  return value;
}

export function parseSimState(search: string): SimState {
  const params = new URLSearchParams(search);
  const source = params.get('source');
  const mode = params.get('mode');
  return {
    source: SOURCES.includes(source as SourceKind) ? (source as SourceKind) : '',
    ref: bounded(params.get('ref')),
    mode: MODES.includes(mode as SimMode) ? (mode as SimMode) : 'sim',
    fight: bounded(params.get('fight')),
  };
}

export function simSearch(state: SimState): string {
  const params = new URLSearchParams();
  if (state.source !== '') params.set('source', state.source);
  if (state.ref !== '') params.set('ref', state.ref);
  if (state.mode !== 'sim') params.set('mode', state.mode);
  if (state.fight !== '') params.set('fight', state.fight);
  const query = params.toString();
  return query === '' ? '' : `?${query}`;
}

export function withSimState(state: SimState, patch: Partial<SimState>): SimState {
  return { ...state, ...patch };
}

/** Matches the contract's sim_id: 12 lowercase base32 characters. */
export const SIM_ID_PATTERN = /^\/sim\/([a-z2-7]{12})\/?$/;

export function simIdFrom(pathname: string): string {
  return SIM_ID_PATTERN.exec(pathname)?.[1] ?? '';
}
