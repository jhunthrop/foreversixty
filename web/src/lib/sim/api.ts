// web/src/lib/sim/api.ts
// Every call this lane makes to our API. It composes requestEnvelope from the account module
// rather than re-implementing the envelope, the session cookie and the CSRF header: that
// function is the one place a browser request to our API is built, and a second copy here
// would be a second place for the CSRF header to go wrong.
//
// Two rules the contract sets and this file enforces:
//   * a saved result carries its whole request, character included: it is plain JSON and a
//     few hundred bytes, and a saved sim that could not say what it simmed would be useless;
//   * POST /v1/sims/run answers 402 when the account is not premium. The envelope's error
//     has no machine-readable code in the phase-3 TypeScript mirror, so premium is keyed off
//     the status, and the copy is ours rather than the API's. Whether to offer the control at
//     all is decided from `user.premium` on GET /v1/me, not from a failed call.
import { AccountError, requestEnvelope } from '../account/api';
import type { CharacterPath } from '../characters';
import { API_BASE_URL } from '../planner/config';
import { bulkCopy, simCopy } from './copy';
import type { KindFilter } from './history';
import type { PhaseRow } from './phase';
import type { SimInput, SimListPage, SimProgress, SimRequest, SimResult, SpecFidelity } from './types';

export const PREMIUM_REQUIRED_STATUS = 402;

export class SimApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'SimApiError';
  }
}

/** Every failure this module raises, with the copy the page shows, keyed off the status. */
function asSimError(error: unknown, fallback: string): SimApiError {
  const status = error instanceof AccountError ? error.status : 0;
  if (status === PREMIUM_REQUIRED_STATUS) return new SimApiError(simCopy.premiumRequired, status);
  if (status === 404) return new SimApiError(simCopy.notFound, status);
  return new SimApiError(fallback, status);
}

async function call<T>(
  path: string,
  apiBase: string,
  fallback: string,
  init: { method?: string; body?: unknown; credentials?: RequestCredentials } = {},
): Promise<T> {
  let data: T | null;
  try {
    ({ data } = await requestEnvelope<T>(path, apiBase, { ...init, failureMessage: fallback }));
  } catch (error) {
    throw asSimError(error, fallback);
  }
  if (data === null) throw new SimApiError(fallback, 0);
  return data;
}

/**
 * Saves a browser-run result. The whole request travels; it is plain JSON throughout.
 *
 * `title` is appended after `apiBase` rather than between it and `result`, even though it
 * is the more commonly-passed of the two (Task 17's save form always sends one): every
 * existing caller -- `SharePanel.svelte`, `store.svelte.ts`, `api.test.ts` -- already calls
 * this with `apiBase` as the second positional argument, and inserting a parameter ahead of
 * it would silently turn every one of those into a call that sends a build's own API origin
 * as the sim's title. The contract's `title` column has no field on `SimResult` itself
 * (only the `sims` table does), so it travels as a sibling key on the JSON body rather than
 * an addition to the `SimResult` shape every reader of that type would then have to ignore.
 */
export async function saveSim(
  result: SimResult,
  apiBase: string = API_BASE_URL,
  title: string = '',
): Promise<string> {
  const body: SimResult & { title?: string } = { ...result, lane: 'browser' };
  delete body.sim_id;
  if (title !== '') body.title = title;
  const data = await call<{ sim_id: string }>('/v1/sims', apiBase, simCopy.saveFailed, {
    method: 'POST',
    body,
  });
  return data.sim_id;
}

export function fetchSim(simId: string, apiBase: string = API_BASE_URL): Promise<SimResult> {
  return call<SimResult>(`/v1/sims/${simId}`, apiBase, simCopy.loadFailed, { credentials: 'omit' });
}

export function listMySims(
  page: number = 1,
  apiBase: string = API_BASE_URL,
  kind: KindFilter = 'all',
): Promise<SimListPage> {
  // "all" is the absence of the parameter, not a value: contract 8 gives `kind=` a closed
  // vocabulary of five and adding a sixth for "no filter" would be a word the API has to
  // know about for no reason.
  const filter = kind === 'all' ? '' : `&kind=${kind}`;
  return call<SimListPage>(`/v1/sims?mine=1&page=${page}${filter}`, apiBase, simCopy.loadFailed);
}

/** The premium lane. Throws SimApiError with status 402 when the account is not premium. */
export async function dispatchServerSim(
  request: SimRequest,
  apiBase: string = API_BASE_URL,
): Promise<string> {
  const data = await call<{ sim_id: string }>('/v1/sims/run', apiBase, simCopy.saveFailed, {
    method: 'POST',
    body: request,
  });
  return data.sim_id;
}

export function fetchSimProgress(simId: string, apiBase: string = API_BASE_URL): Promise<SimProgress> {
  return call<SimProgress>(`/v1/sims/${simId}/progress`, apiBase, simCopy.loadFailed, {
    credentials: 'omit',
  });
}

export async function fetchSpecs(apiBase: string = API_BASE_URL): Promise<SpecFidelity[]> {
  const data = await call<{ specs: SpecFidelity[] }>('/v1/specs', apiBase, simCopy.specsFailed, {
    credentials: 'omit',
  });
  return data.specs;
}

/**
 * The character model the API holds, from the newest source it has. The route is three path
 * segments -- `/v1/characters/{region}/{ruleset}/{name}/sim-input` -- exactly as the existing
 * character route is spelled, never one percent-encoded `{character_key}` segment. Each
 * segment is encoded on its own so a name with a character needing escaping still resolves.
 */
export function fetchSimInput(path: CharacterPath, apiBase: string = API_BASE_URL): Promise<SimInput> {
  const segments = [path.region, path.ruleset, path.slug].map(encodeURIComponent).join('/');
  return call<SimInput>(`/v1/characters/${segments}/sim-input`, apiBase, simCopy.characterFailed);
}

/**
 * The content phase table (contract 10.6), straight off the wire. `phase.ts`'s own
 * `fetchPhases` is the single call site: it wraps this in `BUILT_IN_PHASES`, the build-time
 * fallback, so the gate always has an answer even when this call fails. No other module
 * should call this directly -- read the phase table through `phase.ts` instead.
 */
export async function fetchPhases(apiBase: string = API_BASE_URL): Promise<PhaseRow[]> {
  const data = await call<{ phases: PhaseRow[] }>('/v1/phases', apiBase, bulkCopy.phasesFailed, {
    credentials: 'omit',
  });
  return data.phases;
}
