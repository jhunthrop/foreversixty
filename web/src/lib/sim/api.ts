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
import { invalidate, query } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import type { BuildRecord } from '../planner/types';
import type { BulkServerProgress } from './bulk-types';
import { bulkCopy, simCopy } from './copy';
import { landingCopy } from './landing-copy';
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

/**
 * Every failure this module raises, with the copy the page shows, keyed off the status.
 * `notFoundMessage` lets one caller's 404 read differently from every other route's shared
 * "No sim with that id." (Finding 2, 2026-09-24 landing pass: `fetchSimInput`'s 404 -- the
 * character has no build recorded at all -- is not a missing sim, and saying so the same
 * way misleads). Every other call site passes none, so its own 404 message is unchanged.
 */
function asSimError(error: unknown, fallback: string, notFoundMessage?: string): SimApiError {
  const status = error instanceof AccountError ? error.status : 0;
  if (status === PREMIUM_REQUIRED_STATUS) return new SimApiError(simCopy.premiumRequired, status);
  if (status === 404) return new SimApiError(notFoundMessage ?? simCopy.notFound, status);
  return new SimApiError(fallback, status);
}

async function call<T>(
  path: string,
  apiBase: string,
  fallback: string,
  init: { method?: string; body?: unknown; credentials?: RequestCredentials } = {},
  notFoundMessage?: string,
): Promise<T> {
  let data: T | null;
  try {
    ({ data } = await requestEnvelope<T>(path, apiBase, { ...init, failureMessage: fallback }));
  } catch (error) {
    throw asSimError(error, fallback, notFoundMessage);
  }
  if (data === null) throw new SimApiError(fallback, 0);
  return data;
}

// Cache classes (spec 3.2). fetchSim/fetchSimInput are a public report's meta/fights and a
// public sim's own result -- 5 minutes. fetchMyBuilds/listMySims are "my sims", private --
// 10 minutes. fetchSpecs/fetchPhases are reference data, public -- 1 hour.
const SIM_RESULT_TTL_MS = 5 * 60 * 1000;
const MY_SIMS_TTL_MS = 10 * 60 * 1000;
const REFERENCE_TTL_MS = 60 * 60 * 1000;

function myBuildsKey(apiBase: string, page: number): string {
  return `${apiBase}/v1/builds?mine=1&page=${page}`;
}

function mySimsKey(apiBase: string, page: number, kind: KindFilter): string {
  return `${apiBase}/v1/sims?mine=1&page=${page}&kind=${kind}`;
}

/**
 * Saves a browser-run result. The whole request travels; it is plain JSON throughout.
 *
 * `title` is appended after `apiBase` rather than between it and `result`, even though it
 * is the more commonly-passed of the two (Task 17's save form always sends one): every
 * existing caller -- `SharePanel.svelte`, `store.svelte.ts`, `api.test.ts` -- already calls
 * this with `apiBase` as the second positional argument, and inserting a parameter ahead of
 * it would silently turn every one of those into a call that sends a build's own API origin
 * as the sim's title.
 *
 * `body.title` is always decided by this function's own `title` parameter, never by
 * whatever `result.title` already carries: `result` is ordinarily a fresh run's own
 * `SimResult`, with no title of its own to have an opinion about, but `SimResult.title`
 * (defect fix, `types.ts`) is exactly the field `fetchSim` fills in from a saved sim's own
 * `GET` -- and a re-run of one of those, saved again with no new name typed, must not
 * silently carry the old sim's title onto a result the player never named. `delete
 * body.title` clears whatever `{ ...result }` copied over before the `title !== ''` check
 * decides whether to put one back.
 */
export async function saveSim(
  result: SimResult,
  apiBase: string = API_BASE_URL,
  title: string = '',
): Promise<string> {
  const body: SimResult & { title?: string } = { ...result, lane: 'browser' };
  delete body.sim_id;
  delete body.title;
  if (title !== '') body.title = title;
  const data = await call<{ sim_id: string }>('/v1/sims', apiBase, simCopy.saveFailed, {
    method: 'POST',
    body,
  });
  // Every mine=1 sims list, whatever page or kind, shares this prefix -- a bare
  // invalidate() call clears them all at once rather than guessing which page/kind the
  // caller was last looking at.
  invalidate(`${apiBase}/v1/sims?mine=1`);
  return data.sim_id;
}

export function fetchSim(simId: string, apiBase: string = API_BASE_URL): Promise<SimResult> {
  return query<SimResult>(
    `${apiBase}/v1/sims/${simId}`,
    () => call<SimResult>(`/v1/sims/${simId}`, apiBase, simCopy.loadFailed, { credentials: 'omit' }),
    { scope: 'public', ttlMs: SIM_RESULT_TTL_MS },
  );
}

/**
 * The signed-in player's saved planner builds, for Top Gear's talent candidate list
 * (contract 10.6: `builds.user_id` plus this route). Every failure -- a deployment older
 * than the migration answers 404, and an empty list is not an error at all -- is the
 * caller's to treat as "no saved builds" in one line (`bulkCopy.talentsSavedUnavailable`)
 * rather than an error banner on a page whose other numbers are all correct; this function
 * itself only throws the ordinary `SimApiError` every other read here throws.
 */
export function fetchMyBuilds(
  page: number = 1,
  apiBase: string = API_BASE_URL,
): Promise<{ rows: BuildRecord[]; total: number; page: number; per_page: number }> {
  return query(
    myBuildsKey(apiBase, page),
    () => call(`/v1/builds?mine=1&page=${page}`, apiBase, simCopy.loadFailed),
    { scope: 'private', ttlMs: MY_SIMS_TTL_MS },
  );
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
  return query<SimListPage>(
    mySimsKey(apiBase, page, kind),
    () => call<SimListPage>(`/v1/sims?mine=1&page=${page}${filter}`, apiBase, simCopy.loadFailed),
    { scope: 'private', ttlMs: MY_SIMS_TTL_MS },
  );
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

// Not cached: a poll's whole point is a fresh answer every call (see the plan's Task 8 ruling).
export function fetchSimProgress(simId: string, apiBase: string = API_BASE_URL): Promise<SimProgress> {
  return call<SimProgress>(`/v1/sims/${simId}/progress`, apiBase, simCopy.loadFailed, {
    credentials: 'omit',
  });
}

/**
 * The same route `fetchSimProgress` reads, typed for a bulk job's three extra columns. The
 * premium bulk dispatch itself is `dispatchServerSim` unchanged: a `BulkRequest` is a
 * `SimRequest`, the API derives the kind from the body, and a second POST helper would be a
 * second place for the CSRF header to go wrong.
 */
// Not cached: a poll's whole point is a fresh answer every call (see the plan's Task 8 ruling).
export function fetchBulkProgress(
  simId: string,
  apiBase: string = API_BASE_URL,
): Promise<BulkServerProgress> {
  return call<BulkServerProgress>(`/v1/sims/${simId}/progress`, apiBase, simCopy.loadFailed, {
    credentials: 'omit',
  });
}

export async function fetchSpecs(apiBase: string = API_BASE_URL): Promise<SpecFidelity[]> {
  const data = await query<{ specs: SpecFidelity[] }>(
    `${apiBase}/v1/specs`,
    () => call<{ specs: SpecFidelity[] }>('/v1/specs', apiBase, simCopy.specsFailed, { credentials: 'omit' }),
    { scope: 'public', ttlMs: REFERENCE_TTL_MS },
  );
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
  return query<SimInput>(
    `${apiBase}/v1/characters/${segments}/sim-input`,
    () =>
      call<SimInput>(
        `/v1/characters/${segments}/sim-input`,
        apiBase,
        simCopy.characterFailed,
        {},
        landingCopy.buildMissingFallback,
      ),
    { scope: 'public', ttlMs: SIM_RESULT_TTL_MS },
  );
}

/**
 * The content phase table (contract 10.6), straight off the wire. `phase.ts`'s own
 * `fetchPhases` is the single call site: it wraps this in `BUILT_IN_PHASES`, the build-time
 * fallback, so the gate always has an answer even when this call fails. No other module
 * should call this directly -- read the phase table through `phase.ts` instead.
 */
export async function fetchPhases(apiBase: string = API_BASE_URL): Promise<PhaseRow[]> {
  const data = await query<{ phases: PhaseRow[] }>(
    `${apiBase}/v1/phases`,
    () => call<{ phases: PhaseRow[] }>('/v1/phases', apiBase, bulkCopy.phasesFailed, { credentials: 'omit' }),
    { scope: 'public', ttlMs: REFERENCE_TTL_MS },
  );
  return data.phases;
}
