// web/src/lib/sim/url.ts
// /sim's URL state. The page is static and the island reads the query, exactly as the
// planner reads ?class= and the report reads ?fight=, so a link to "sim this fight" or
// "sim this build" is a plain URL a player can paste.
//
// Everything here is validated against a fixed vocabulary: the query is attacker-controlled
// and the values reach fetch() paths.
import type { SimRequest, SourceKind } from './types';

const SOURCES: readonly SourceKind[] = ['armory', 'addon', 'build', 'fight', 'manual'];
const MODES = ['sim', 'compare'] as const;
export type SimMode = (typeof MODES)[number];

/** A build id is 12, a character key about 40, a fight ref 14. 128 is generous and finite. */
const MAX_REF = 128;

/**
 * An FS1 code carries three tree strings and up to seventeen gear entries, easily several
 * hundred characters -- far past MAX_REF. `fs1.ts`'s own decoder refuses anything over 2048
 * characters for the same reason (a generous multiple of a real code, high enough to never
 * clip one); this mirrors that bound rather than importing it, since url.ts only ever needs
 * to cap an attacker-controlled query string before the code reaches a decoder at all.
 */
const MAX_CODE = 2048;

/**
 * A base64url request is about a third larger than its JSON, and a real single-run request
 * is two to three kilobytes, so this leaves a comfortable margin under the ~8 KB every
 * browser and proxy handles. A request past it is shared by its saved id instead
 * (contract 9), which the drawer says in so many words.
 */
export const MAX_REQUEST_PARAM = 8192;

export interface SimState {
  source: SourceKind | '';
  ref: string;
  /** An FS1 build code, for a "Sim this build" link to a build that has not been saved. */
  code: string;
  /** A whole request, base64url-encoded: the drawer's share link (design 8). */
  req: string;
  mode: SimMode;
  /** The fight compare mode is comparing against: "<report_id>:<fight_index>". */
  fight: string;
}

export function defaultSimState(): SimState {
  return { source: '', ref: '', code: '', req: '', mode: 'sim', fight: '' };
}

function bounded(value: string | null, max: number): string {
  if (value === null || value.length > max) return '';
  return value;
}

export function parseSimState(search: string): SimState {
  const params = new URLSearchParams(search);
  const source = params.get('source');
  const mode = params.get('mode');
  return {
    source: SOURCES.includes(source as SourceKind) ? (source as SourceKind) : '',
    ref: bounded(params.get('ref'), MAX_REF),
    code: bounded(params.get('code'), MAX_CODE),
    req: bounded(params.get('req'), MAX_REQUEST_PARAM),
    mode: MODES.includes(mode as SimMode) ? (mode as SimMode) : 'sim',
    fight: bounded(params.get('fight'), MAX_REF),
  };
}

export function simSearch(state: SimState): string {
  const params = new URLSearchParams();
  if (state.source !== '') params.set('source', state.source);
  if (state.ref !== '') params.set('ref', state.ref);
  if (state.code !== '') params.set('code', state.code);
  if (state.req !== '') params.set('req', state.req);
  if (state.mode !== 'sim') params.set('mode', state.mode);
  if (state.fight !== '') params.set('fight', state.fight);
  const query = params.toString();
  return query === '' ? '' : `?${query}`;
}

export function withSimState(state: SimState, patch: Partial<SimState>): SimState {
  return { ...state, ...patch };
}

/** UTF-8 bytes to base64url, no padding: a chat client must not mangle a share link. */
function toBase64Url(text: string): string {
  const bytes = new TextEncoder().encode(text);
  let binary = '';
  for (const byte of bytes) binary += String.fromCharCode(byte);
  return btoa(binary).replaceAll('+', '-').replaceAll('/', '_').replaceAll('=', '');
}

function fromBase64Url(value: string): string {
  const padded = value.replaceAll('-', '+').replaceAll('_', '/');
  const binary = atob(padded.padEnd(Math.ceil(padded.length / 4) * 4, '='));
  const bytes = Uint8Array.from(binary, (character) => character.charCodeAt(0));
  return new TextDecoder().decode(bytes);
}

/** The request as a query value, or null when it is past the budget. */
export function encodeRequestParam(request: SimRequest): string | null {
  const encoded = toBase64Url(JSON.stringify(request));
  return encoded.length > MAX_REQUEST_PARAM ? null : encoded;
}

/**
 * A query value back into a request, or null. The query is attacker-controlled and this
 * value reaches the engine, so anything that is not a JSON object is refused here and the
 * engine's own Validate refuses the rest.
 */
export function decodeRequestParam(value: string): SimRequest | null {
  if (value === '' || value.length > MAX_REQUEST_PARAM) return null;
  try {
    const parsed: unknown = JSON.parse(fromBase64Url(value));
    if (typeof parsed !== 'object' || parsed === null || Array.isArray(parsed)) return null;
    return parsed as SimRequest;
  } catch {
    return null;
  }
}

/** Matches the contract's sim_id: 12 lowercase base32 characters. */
export const SIM_ID_PATTERN = /^\/sim\/([a-z2-7]{12})\/?$/;

export function simIdFrom(pathname: string): string {
  return SIM_ID_PATTERN.exec(pathname)?.[1] ?? '';
}
