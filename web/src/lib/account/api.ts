// web/src/lib/account/api.ts
// Every call the browser makes to the account half of the API. Sessions are an opaque
// HttpOnly cookie, so nothing here reads or writes one: `credentials: 'include'` is the
// whole mechanism. CSRF is the double-submit pair the contract specifies -- the readable
// fs_csrf cookie echoed in X-CSRF-Token -- on every state-changing request.
//
// DELETE /v1/sessions and PATCH /v1/me { anonymize } come from the contract's Amendments
// section; the first draft required both behaviours on /account without naming a route.
import { API_BASE_URL } from '../planner/config';

export const SIGN_IN_REQUIRED = 'Sign in to continue';
export const ACCOUNT_FAILED = 'That did not work; try again';
export const EMAIL_SENT = 'Check your email for a sign-in link. It works once, for twenty minutes.';

export class AccountError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'AccountError';
  }
}

export interface MeCharacter {
  key: string;
  region: string;
  ruleset: string;
  name: string;
  class?: string;
}

export interface MeGuild {
  id: number;
  region: string;
  ruleset: string;
  name: string;
  rank?: string;
  /** Spec section 2.6: GET /v1/me gains Consent and Verified per Guild entry. */
  consent?: 'roster' | 'gear' | 'gear_bags';
  verified: boolean;
}

export interface Me {
  user: {
    id: number;
    battletag: string | null;
    email: string | null;
    role: string;
    anonymize: boolean;
    /**
     * The premium flag, set by hand until payments exist (simulator contract). The web
     * reads it to decide whether to offer the server lane at all, rather than offering
     * the control and letting POST /v1/sims/run answer 402 -- a button that always fails
     * is worse than no button.
     */
    premium: boolean;
  };
  characters: MeCharacter[];
  guilds: MeGuild[];
}

export interface Device {
  id: string;
  name: string;
  platform: string;
  created_at: string;
  last_seen_at: string | null;
}

export interface PairingCode {
  code: string;
  /** Seconds the code stays valid for, not an instant: the contract's Amendments section. */
  expires_in: number;
}

/** The readable half of the double-submit pair. Empty when there is no session. */
export function csrfToken(): string {
  const match = /(?:^|;\s*)fs_csrf=([^;]*)/.exec(document.cookie);
  return match === null ? '' : decodeURIComponent(match[1]);
}

interface Envelope<T> {
  ok: boolean;
  data: T | null;
  error: { message?: string } | null;
}

export interface EnvelopeResult<T> {
  status: number;
  data: T | null;
  /** The API's own `error.message`, when the envelope carried one -- null otherwise. */
  message: string | null;
}

/**
 * The one place a browser call to our API builds the request: `credentials: 'include'`
 * (by default -- see `init.credentials` below) for the session cookie, `X-CSRF-Token`
 * from the readable `fs_csrf` cookie on every non-GET, and the JSON envelope parsed and
 * turned into an `AccountError` (carrying the API's own `error.message`) on a failed
 * response. Every module that talks to our API -- this one, `web/src/lib/upload/
 * multipart.ts` for the two upload routes, and `web/src/lib/rankings/api.ts` for the
 * public rankings/character/guild reads -- goes through this, so the CSRF header and the
 * credentials mode can only go wrong in one place. It never throws for a *successful*
 * response with no `data`; each caller decides what "no data" means for its own endpoint.
 * Failure is read off the HTTP status alone (`!response.ok`), not the envelope's own `ok`
 * flag: every response this API has ever sent keeps the two in agreement, and the
 * account module's callers (`fetchMe`'s 401-means-signed-out, `listMyReports`'s
 * empty-page-when-signed-out) rely on reading a non-2xx status themselves. A caller that
 * needs `data` to be non-null on success (rankings, the two upload routes) checks that
 * itself, the same way `multipart.ts`'s `post()` already does; the pattern in
 * `rankings/api.ts`'s `get()` matches it.
 *
 * `init.credentials` overrides the session cookie's `'include'` default. A public read
 * that answers the same way for every visitor -- a ranking is not tied to who is asking
 * -- passes `'omit'` so it does not send a cookie or the CSRF header the API does not
 * need for it (CSRF is skipped on every GET regardless, since it only guards
 * state-changing requests, but `'omit'` also keeps the session cookie itself off a
 * request that has no business carrying one).
 *
 * `failureMessage` overrides the fallback shown when a failed response -- including one
 * the browser never reached at all -- carries no `error.message` of its own; it defaults
 * to `ACCOUNT_FAILED`, the account module's own generic copy.
 */
export async function requestEnvelope<T>(
  path: string,
  apiBase: string,
  init: {
    method?: string;
    body?: unknown;
    failureMessage?: string;
    credentials?: RequestCredentials;
  } = {},
): Promise<EnvelopeResult<T>> {
  const method = init.method ?? 'GET';
  const headers = new Headers({ accept: 'application/json' });
  if (method !== 'GET') {
    headers.set('x-csrf-token', csrfToken());
    if (init.body !== undefined) headers.set('content-type', 'application/json');
  }

  let response: Response;
  try {
    response = await fetch(
      new Request(`${apiBase}${path}`, {
        method,
        headers,
        credentials: init.credentials ?? 'include',
        body: init.body === undefined ? undefined : JSON.stringify(init.body),
      }),
    );
  } catch {
    throw new AccountError(init.failureMessage ?? ACCOUNT_FAILED, 0);
  }

  let envelope: Envelope<T> | null = null;
  try {
    envelope = (await response.json()) as Envelope<T>;
  } catch {
    envelope = null;
  }

  if (!response.ok) {
    // The API's own message is shown verbatim when it has one: it is the only thing that
    // can say "too many sign-in links" or name the field that was wrong.
    throw new AccountError(
      envelope?.error?.message ?? init.failureMessage ?? ACCOUNT_FAILED,
      response.status,
    );
  }
  return { status: response.status, data: envelope?.data ?? null, message: envelope?.error?.message ?? null };
}

async function call<T>(
  path: string,
  apiBase: string,
  init: { method?: string; body?: unknown } = {},
): Promise<T | null> {
  const { data } = await requestEnvelope<T>(path, apiBase, init);
  return data;
}

/** Null means "not signed in", which is a state the header renders, not a failure. */
export async function fetchMe(apiBase: string = API_BASE_URL): Promise<Me | null> {
  try {
    return await call<Me>('/v1/me', apiBase);
  } catch (error) {
    if (error instanceof AccountError && (error.status === 401 || error.status === 403)) return null;
    throw error;
  }
}

/**
 * One `/v1/me` per page, shared by every island that needs to know who is signed in: the
 * header, the pairing block, the upload form and the reports list are separate islands, and
 * each asking on its own was four identical requests. A failure is not remembered, so the
 * next caller asks again rather than inheriting a dead promise.
 */
const sessions = new Map<string, Promise<Me | null>>();

export function fetchMeOnce(apiBase: string = API_BASE_URL): Promise<Me | null> {
  const known = sessions.get(apiBase);
  if (known !== undefined) return known;
  const pending = fetchMe(apiBase).catch((error: unknown) => {
    sessions.delete(apiBase);
    throw error;
  });
  sessions.set(apiBase, pending);
  return pending;
}

/** Tests only: a fresh page has no remembered session. */
export function forgetSession(): void {
  sessions.clear();
}

export function battlenetStartUrl(next: string, apiBase: string = API_BASE_URL): string {
  return `${apiBase}/v1/auth/battlenet/start?next=${encodeURIComponent(next)}`;
}

export async function requestEmailLink(email: string, apiBase: string = API_BASE_URL): Promise<void> {
  await call('/v1/auth/email', apiBase, { method: 'POST', body: { email } });
}

export async function listDevices(apiBase: string = API_BASE_URL): Promise<Device[]> {
  return (await call<Device[]>('/v1/devices', apiBase)) ?? [];
}

export async function pairDevice(apiBase: string = API_BASE_URL): Promise<PairingCode> {
  const code = await call<PairingCode>('/v1/devices/pair', apiBase, { method: 'POST' });
  if (code === null) throw new AccountError(ACCOUNT_FAILED, 0);
  return code;
}

export async function revokeDevice(id: string, apiBase: string = API_BASE_URL): Promise<void> {
  await call(`/v1/devices/${encodeURIComponent(id)}`, apiBase, { method: 'DELETE' });
}

export async function signOut(apiBase: string = API_BASE_URL): Promise<void> {
  await call('/v1/sessions', apiBase, { method: 'DELETE' });
}

export async function setAnonymize(value: boolean, apiBase: string = API_BASE_URL): Promise<void> {
  await call('/v1/me', apiBase, { method: 'PATCH', body: { anonymize: value } });
}

/** One row of the "Your reports" list. */
export interface MyReport {
  id: string;
  title: string;
  zone: string;
  status: string;
  visibility: string;
  created_at: string;
  fight_count: number;
  kill_count: number;
}

export interface MyReportPage {
  rows: MyReport[];
  total: number;
  page: number;
  per_page: number;
}

/** The contract's Amendments section fixes the page size at 100. */
export const REPORTS_PER_PAGE = 100;

const EMPTY_REPORT_PAGE: MyReportPage = { rows: [], total: 0, page: 1, per_page: REPORTS_PER_PAGE };

/**
 * The "Your reports" list on /logs, per the contract's Amendments section. A signed-out
 * visitor gets an empty page rather than an error, because /logs renders for them too --
 * it just tells them to sign in.
 */
export async function listMyReports(page: number = 1, apiBase: string = API_BASE_URL): Promise<MyReportPage> {
  try {
    return (await call<MyReportPage>(`/v1/reports?mine=1&page=${page}`, apiBase)) ?? EMPTY_REPORT_PAGE;
  } catch (error) {
    if (error instanceof AccountError && (error.status === 401 || error.status === 403)) {
      return EMPTY_REPORT_PAGE;
    }
    throw error;
  }
}
