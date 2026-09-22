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

/**
 * A character's own guild line (spec 2026-09-22 §6): distinct from `MeGuild`, the
 * account-level "My guilds" entry below -- this is what `GET /v1/me`'s `characters[].guild`
 * carries, with no `region`/`ruleset`/`consent`/`plan` of its own.
 */
export interface MeCharacterGuild {
  id: number;
  name: string;
  rank?: string;
  rank_index?: number;
  verified: boolean;
}

export interface MeCharacter {
  key: string;
  region: string;
  ruleset: string;
  name: string;
  class?: string;
  /** Spec 2026-09-22 §6: omitted when the Battle.net import/refresh has not learned it yet. */
  realm?: string;
  level?: number;
  faction?: 'alliance' | 'horde';
  /** "Night Elf"; from the Battle.net import, omitted until it has run. */
  race?: string;
  gender?: 'male' | 'female';
  /** Equipped item level, from the Battle.net character profile when it answered. */
  item_level?: number;
  /** `'bnet'` or `'export'` -- which path last wrote this row. Optional: older/stubbed
   *  fixtures written before the Battle.net import landed carry no such field. */
  source?: string;
  /** Omitted when the character has no `guild_characters` row at all. */
  guild?: MeCharacterGuild;
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
  /**
   * Non-null only for a verified officer/leader of a guild with an active guild-plan
   * entitlement (spec 1.4) -- billing detail, gated tighter than membership alone. Optional
   * because the live API does not send it yet (the API lane has not landed); every reader
   * treats a missing key the same as null.
   */
  plan?: GuildBillingView | null;
}

/** Spec 1.4: every Can() feature pre-resolved for the caller, plus their own billing state. */
export interface EntitlementsView {
  server_sims: boolean;
  retention: boolean;
  multi_compare: boolean;
  history: boolean;
  notifications: boolean;
  officer_views: boolean;
  roster_check: boolean;
  supporter_mark: boolean;
  billing: BillingView | null;
}

export interface BillingView {
  plan: string;
  status: string;
  current_period_end: string | null;
  cancel_at_period_end: boolean;
}

export interface GuildBillingView {
  status: string;
  current_period_end: string | null;
  cancel_at_period_end: boolean;
  /** Empty for a plan granted through the CLI, never through Stripe -- there is no billing
   *  contact to name (api/internal/auth/handler.go's attachGuildPlans). A reader must not
   *  show "Billed by" with nothing after it. */
  billed_by: string;
  you_are_billing_contact: boolean;
}

export interface Me {
  user: {
    id: number;
    battletag: string | null;
    email: string | null;
    role: string;
    anonymize: boolean;
    /**
     * The premium flag, set by hand until payments exist. Optional and kept only for
     * backward compatibility with e2e fixtures written before the entitlements API landed
     * (web lane ruling A, docs/superpowers/plans/2026-09-21-pay-web.md) -- new code reads
     * `entitlements.server_sims` via `effectiveServerSims`, never this field directly.
     */
    premium?: boolean;
  };
  characters: MeCharacter[];
  guilds: MeGuild[];
  /** Optional: absent until the API lane ships spec 1.4's block. See effectiveServerSims. */
  entitlements?: EntitlementsView;
  /** RFC3339; spec 2026-09-22 §6, omitted when this account has never imported from
   *  Battle.net. */
  bnet_imported_at?: string;
}

/**
 * Whether server-side sims should be offered, reading the new entitlements shape when the
 * API sends it and falling back to the legacy `user.premium` boolean otherwise -- one place
 * for the fallback (web lane ruling A) so SimView.svelte, ToolsView.svelte and RunControl.svelte
 * never each re-derive it differently.
 */
export function effectiveServerSims(me: Me | null): boolean {
  if (me === null) return false;
  if (me.entitlements !== undefined) return me.entitlements.server_sims === true;
  return me.user.premium === true;
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

/** One character's export, matching `POST /v1/me/exports`'s body (spec 2026-09-22 §4.5) --
 *  the same `{name, region, ruleset, export}` shape the companion's `PutExports` reads,
 *  because the export string itself carries no character identity (see the plan's Ruling 1). */
export interface MyExportInput {
  name: string;
  region: string;
  ruleset: string;
  export: string;
}

/**
 * The signed-in paste's write path: the addon-less way to reach the same `characters` rows
 * the companion writes (spec 2026-09-22 §4.5). Returns the `/v1/me` character objects for the
 * keys written, empty when the API answered with none.
 */
export async function postMyExports(
  exports: MyExportInput[],
  apiBase: string = API_BASE_URL,
): Promise<MeCharacter[]> {
  const result = await call<{ characters: MeCharacter[] }>('/v1/me/exports', apiBase, {
    method: 'POST',
    body: { exports },
  });
  return result?.characters ?? [];
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
