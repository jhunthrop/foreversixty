// web/src/lib/account/api.ts
// Every call the browser makes to the account half of the API. Sessions are an opaque
// HttpOnly cookie, so nothing here reads or writes one: `credentials: 'include'` is the
// whole mechanism. CSRF is the double-submit pair the contract specifies -- the readable
// fs_csrf cookie echoed in X-CSRF-Token -- on every state-changing request. DELETE
// /v1/sessions and PATCH /v1/me { anonymize } come from the contract's Amendments section.
import { API_BASE_URL } from '../planner/config';
import { forgetPrivate, invalidate, query, refetch, setQueryData } from '../data/query';
import { forgetGuildReserves } from '../guild/reserve-keys';
import { normalizeMeBilling } from '../billing/normalize';
export type {
  BillingView,
  DiscordDisabledReason,
  DiscordEvents,
  DiscordNotifications,
  DiscordSource,
  EntitlementsView,
  FinishedSim,
  GuildBillingView,
  Me,
  MeCharacter,
  MeCharacterGuild,
  MeGuild,
} from './me-types';
import { clearPointerOnSignOut } from './session-pointer';
import type { Me, MeCharacter } from './me-types';
import { premiumSource } from './plan-line';
import { clearMeGuildCount, clearPremiumHint, rememberMeHints, writePremiumHint } from './premium-hint';

export const SIGN_IN_REQUIRED = 'Sign in to continue';
export const ACCOUNT_FAILED = 'That did not work; try again';
export const EMAIL_SENT = 'Check your email for a sign-in link. It works once, for twenty minutes.';

export class AccountError extends Error {
  constructor(
    message: string,
    readonly status: number,
    /** The API's own `error.code` ("not_found", "flagged"), when the envelope carried one. */
    readonly code: string | null = null,
    /** A 422's per-field lines (`{starts_at: "Pick a time after now."}`), when the envelope carried them. */
    readonly fields: Readonly<Record<string, string>> | null = null,
  ) {
    super(message);
    this.name = 'AccountError';
  }
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

/** A desktop companion paired through /setup, or a native app signed in with Battle.net. */
export type DeviceKind = 'companion' | 'app';

export interface Device {
  id: string;
  name: string;
  /** The companion's own label (`windows/amd64`), or `ios` / `android` for an app. */
  platform: string;
  /** Absent from an API older than the app sign-in, which only ever held companions. */
  kind?: DeviceKind;
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
  error: { code?: string; message?: string; fields?: Record<string, string> } | null;
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

/**
 * The ETag `requestEnvelope` last saw for a GET, per `${apiBase}${path}`, and the body it
 * came with. Never touched for a non-GET request. `query.ts`'s stale-while-revalidate
 * background reload is what actually exercises this: a 304 makes that reload a cheap,
 * bodiless round trip, resolved transparently to the same data the caller already had.
 */
const remembered = new Map<string, { etag: string; data: unknown }>();

/** Tests only: a fresh module has remembered no ETags. */
export function forgetRemembered(): void {
  remembered.clear();
}

/**
 * Builds the request headers, applying `If-None-Match` for a GET that has a remembered
 * ETag. Returns the `known` entry too (rather than making `requestEnvelope` look it up
 * again) so the 304 short-circuit below can resolve to the same body without a second
 * `remembered.get`.
 */
function buildRequestHeaders(
  method: string,
  body: unknown,
  rememberedKey: string,
): { headers: Headers; known: { etag: string; data: unknown } | undefined } {
  const headers = new Headers({ accept: 'application/json' });
  if (method !== 'GET') {
    headers.set('x-csrf-token', csrfToken());
    // A FormData body (the guild crest upload's multipart field) must keep the browser's
    // own `multipart/form-data; boundary=...` header -- setting `content-type` ourselves
    // would drop the boundary and the server could never split the parts.
    if (body !== undefined && !(body instanceof FormData)) headers.set('content-type', 'application/json');
    return { headers, known: undefined };
  }
  const known = remembered.get(rememberedKey);
  if (known !== undefined) headers.set('if-none-match', known.etag);
  return { headers, known };
}

/** Records (or forgets) the ETag a successful GET came back with. */
function rememberEtag(rememberedKey: string, response: Response, data: unknown): void {
  const etag = response.headers.get('etag');
  if (etag !== null) remembered.set(rememberedKey, { etag, data });
  else remembered.delete(rememberedKey);
}

/** A response with no JSON body (or a malformed one) parses to `null`, never throws. */
async function parseEnvelope<T>(response: Response): Promise<Envelope<T> | null> {
  try {
    return (await response.json()) as Envelope<T>;
  } catch {
    return null;
  }
}

export async function requestEnvelope<T>(
  path: string,
  apiBase: string,
  init: {
    method?: string;
    body?: unknown;
    failureMessage?: string;
    credentials?: RequestCredentials;
    /** Aborts the request: a superseded typeahead lookup. */
    signal?: AbortSignal;
  } = {},
): Promise<EnvelopeResult<T>> {
  const method = init.method ?? 'GET';
  const rememberedKey = `${apiBase}${path}`;
  const { headers, known } = buildRequestHeaders(method, init.body, rememberedKey);

  let response: Response;
  try {
    response = await fetch(
      new Request(`${apiBase}${path}`, {
        method,
        headers,
        credentials: init.credentials ?? 'include',
        signal: init.signal,
        body:
          init.body === undefined
            ? undefined
            : init.body instanceof FormData
              ? init.body
              : JSON.stringify(init.body),
      }),
    );
  } catch {
    throw new AccountError(init.failureMessage ?? ACCOUNT_FAILED, 0);
  }

  if (method === 'GET' && response.status === 304) {
    return { status: 304, data: (known?.data as T | undefined) ?? null, message: null };
  }

  const envelope = await parseEnvelope<T>(response);

  if (!response.ok) {
    // The API's own message is shown verbatim when it has one, since it can say things
    // like "too many sign-in links" that no generic fallback can.
    throw new AccountError(
      envelope?.error?.message ?? init.failureMessage ?? ACCOUNT_FAILED,
      response.status,
      envelope?.error?.code ?? null,
      envelope?.error?.fields ?? null,
    );
  }

  if (method === 'GET') rememberEtag(rememberedKey, response, envelope?.data ?? null);

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
    const me = normalizeMeBilling(await call<Me>('/v1/me', apiBase));
    if (me === null) clearPremiumHint();
    else writePremiumHint(premiumSource(me, new Date()));
    return rememberMeHints(me);
  } catch (error) {
    if (error instanceof AccountError && (error.status === 401 || error.status === 403)) {
      clearPremiumHint();
      return null;
    }
    throw error;
  }
}

/** Exported (with `meKey`/`ME_QUERY_VERSION` below) for Base.astro's pre-paint selector
 *  script: that script reads this exact query's persisted entry out of `query.ts`'s
 *  localStorage cache (for the level) before the character selector hydrates, and
 *  needs this module's own TTL/version rather than a second, drifting copy of either. */
export const ME_TTL_MS = 10 * 60 * 1000;

/** `query.ts`'s `getOrCreateEntry` defaults a caller's omitted `version` to `0`; named here
 *  (and passed explicitly below) so Base.astro's pre-paint script has one real export to
 *  import instead of having to know that implicit default itself. */
export const ME_QUERY_VERSION = 0;

export function meKey(apiBase: string): string {
  return `${apiBase}/v1/me`;
}

/**
 * One `/v1/me` per key, shared by every island that needs to know who is signed in, through
 * the shared client cache (web/src/lib/data/query.ts): a stale answer renders instantly and
 * revalidates in the background; a signed-out answer (`null`) is cached exactly like a
 * signed-in one, since `query.ts` tracks "has an answer" via its own `status`, not `data
 * !== null`.
 */
export function fetchMeOnce(apiBase: string = API_BASE_URL): Promise<Me | null> {
  return query<Me | null>(meKey(apiBase), () => fetchMe(apiBase), {
    scope: 'private',
    ttlMs: ME_TTL_MS,
    version: ME_QUERY_VERSION,
  });
}

/** Drops the cached `/v1/me`: membership, rank, plan and every character card's guild line
 *  come from it, so each guild mutation (lib/guild/api.ts's `invalidateGuild`) calls this. */
export function invalidateMe(apiBase: string = API_BASE_URL): void {
  invalidate(meKey(apiBase));
}

/** `fetchMeOnce` for a guild page: an answer older than a minute is never painted, the
 *  cached one is revalidated on every read, and `onFresh` gets the newer answer when it
 *  differs. */
export function fetchMeFresh(
  onFresh?: (me: Me | null) => void,
  apiBase: string = API_BASE_URL,
): Promise<Me | null> {
  return query<Me | null>(meKey(apiBase), () => fetchMe(apiBase), {
    scope: 'private',
    ttlMs: 60 * 1000,
    version: ME_QUERY_VERSION,
    revalidateOnRead: true,
    onRevalidated: onFresh,
  });
}

/** `/v1/me` asked again right now (window focus on the Settings page). */
export function refetchMe(onData: (me: Me | null) => void, apiBase: string = API_BASE_URL): void {
  refetch(
    meKey(apiBase),
    () => fetchMe(apiBase),
    { scope: 'private', ttlMs: 60 * 1000, version: ME_QUERY_VERSION },
    onData,
  );
}

/** Tests only: a fresh page has no remembered session. */
export function forgetSession(): void {
  forgetPrivate();
  forgetRemembered();
}

export function battlenetStartUrl(next: string, apiBase: string = API_BASE_URL): string {
  return `${apiBase}/v1/auth/battlenet/start?next=${encodeURIComponent(next)}`;
}

/** `next` rides in the body so the emailed link can land back on the page that asked (/connect). */
export async function requestEmailLink(
  email: string,
  apiBase: string = API_BASE_URL,
  next?: string,
): Promise<void> {
  await call('/v1/auth/email', apiBase, {
    method: 'POST',
    body: next === undefined ? { email } : { email, next },
  });
}

const DEVICES_TTL_MS = 10 * 60 * 1000;

function devicesKey(apiBase: string): string {
  return `${apiBase}/v1/devices`;
}

export function listDevices(apiBase: string = API_BASE_URL): Promise<Device[]> {
  return query<Device[]>(
    devicesKey(apiBase),
    async () => (await call<Device[]>('/v1/devices', apiBase)) ?? [],
    {
      scope: 'private',
      ttlMs: DEVICES_TTL_MS,
    },
  );
}

/** The companions among an account's devices: what uploads logs and carries builds into the
 *  game. A signed-in phone app is listed on /account but never counts as a companion. */
export function companionsOf(devices: readonly Device[]): Device[] {
  return devices.filter((device) => device.kind !== 'app');
}

export async function listCompanions(apiBase: string = API_BASE_URL): Promise<Device[]> {
  return companionsOf(await listDevices(apiBase));
}

/**
 * Logs landing spec (2026-10-04) §4.D.2's pairing poll needs a real read every tick, not
 * `listDevices`'s own cached one (`query()`'s "an entry this page has already loaded is
 * served as is" rule -- query.ts's `Entry.validated` -- would otherwise hand every poll
 * tick back the same stale array forever). Invalidating first forces exactly that.
 */
export async function refreshDevices(apiBase: string = API_BASE_URL): Promise<Device[]> {
  invalidate(devicesKey(apiBase));
  return listDevices(apiBase);
}

/**
 * Queues a saved build in the signed-in account's addon inbox (`POST /v1/addon/inbox`). The
 * companion fetches the inbox at start and every ten minutes and writes it for the addon,
 * which shows it on the next login. `buildId` must be a saved build's id (the share flow's
 * `POST /v1/builds`); `characterKey` aims it at one character, omitted for the whole account.
 * Throws `AccountError`: 401 for a visitor who is not signed in, 400 for an id or key the API
 * refuses.
 */
export async function queueBuildToAddon(
  buildId: string,
  characterKey?: string,
  apiBase: string = API_BASE_URL,
): Promise<void> {
  await call('/v1/addon/inbox', apiBase, {
    method: 'POST',
    body:
      characterKey === undefined ? { build_id: buildId } : { build_id: buildId, character_key: characterKey },
  });
}

export async function pairDevice(apiBase: string = API_BASE_URL): Promise<PairingCode> {
  const code = await call<PairingCode>('/v1/devices/pair', apiBase, { method: 'POST' });
  if (code === null) throw new AccountError(ACCOUNT_FAILED, 0);
  invalidate(devicesKey(apiBase));
  return code;
}

/** What the companion's connect asked to be called, as the signed-in page shows it. */
export interface ConnectDevice {
  name: string;
  platform: string;
}

export interface ConnectLookup extends ConnectDevice {
  expires_in: number;
}

/** Companion connect (ruling R-C1). 404 means unknown, expired or already decided. */
export async function connectDeviceLookup(
  code: string,
  apiBase: string = API_BASE_URL,
): Promise<ConnectLookup> {
  const found = await call<ConnectLookup>(
    `/v1/devices/connect/lookup?code=${encodeURIComponent(code)}`,
    apiBase,
  );
  if (found === null) throw new AccountError(ACCOUNT_FAILED, 0);
  return found;
}

/** 404 unknown/expired, 409 already decided. */
export async function connectDeviceApprove(
  code: string,
  apiBase: string = API_BASE_URL,
): Promise<ConnectDevice> {
  const device = await call<ConnectDevice>('/v1/devices/connect/approve', apiBase, {
    method: 'POST',
    body: { code },
  });
  if (device === null) throw new AccountError(ACCOUNT_FAILED, 0);
  invalidate(devicesKey(apiBase));
  return device;
}

export async function connectDeviceDeny(code: string, apiBase: string = API_BASE_URL): Promise<void> {
  await call('/v1/devices/connect/deny', apiBase, { method: 'POST', body: { code } });
}

export async function revokeDevice(id: string, apiBase: string = API_BASE_URL): Promise<void> {
  await call(`/v1/devices/${encodeURIComponent(id)}`, apiBase, { method: 'DELETE' });
  invalidate(devicesKey(apiBase));
}

export async function signOut(apiBase: string = API_BASE_URL): Promise<void> {
  forgetRemembered();
  forgetPrivate();
  forgetGuildReserves();
  clearPointerOnSignOut();
  clearPremiumHint();
  clearMeGuildCount();
  await call('/v1/sessions', apiBase, { method: 'DELETE' });
}

/** PATCH /v1/me; the answer is the new `/v1/me`, normalized and cached. */
async function patchMe(body: Record<string, unknown>, apiBase: string): Promise<void> {
  const me = normalizeMeBilling(await call<Me>('/v1/me', apiBase, { method: 'PATCH', body }));
  if (me !== null) setQueryData(meKey(apiBase), me);
}

export async function setAnonymize(value: boolean, apiBase: string = API_BASE_URL): Promise<void> {
  await patchMe({ anonymize: value }, apiBase);
}

/** Chooses the account's main character (one of its own); the rest become alts. */
export async function setMainCharacter(key: string, apiBase: string = API_BASE_URL): Promise<void> {
  await patchMe({ main_character_key: key }, apiBase);
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
  invalidate(meKey(apiBase));
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
  /** Distinct bosses killed (the API's own count; absent from an API before it). */
  boss_count?: number;
  /** The character the log is credited to (`{ key, name }`); null when none is. */
  logging_character?: string | { key?: string; name?: string } | null;
  /** The newest fight upload, else creation; with `heartbeat_at` it says Quiet or Live as the report meta does. */
  last_upload_at?: string | null;
  /** The companion's last still-tailing ping; null before the first. */
  heartbeat_at?: string | null;
  /** The night's length; the list endpoint does not send it yet, and the row leaves it out until it does. */
  duration_ms?: number;
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

const MY_REPORTS_TTL_MS = 10 * 60 * 1000;

function myReportsKey(apiBase: string, page: number): string {
  return `${apiBase}/v1/reports?mine=1&page=${page}`;
}

/**
 * The "Your reports" list on /logs, per the contract's Amendments section. A signed-out
 * visitor gets an empty page rather than an error, because /logs renders for them too --
 * it just tells them to sign in.
 */
export async function listMyReports(page: number = 1, apiBase: string = API_BASE_URL): Promise<MyReportPage> {
  return query<MyReportPage>(
    myReportsKey(apiBase, page),
    async () => {
      try {
        return (await call<MyReportPage>(`/v1/reports?mine=1&page=${page}`, apiBase)) ?? EMPTY_REPORT_PAGE;
      } catch (error) {
        if (error instanceof AccountError && (error.status === 401 || error.status === 403)) {
          return EMPTY_REPORT_PAGE;
        }
        throw error;
      }
    },
    { scope: 'private', ttlMs: MY_REPORTS_TTL_MS },
  );
}

function guildReportsKey(apiBase: string, guildId: number, page: number): string {
  return `${apiBase}/v1/guilds/${guildId}/reports?page=${page}`;
}

/**
 * `Your reports`' guild tab (logs landing spec 2026-10-04 §4.C.1): a guild's own
 * `guild`-visibility reports, in the same `MyReportPage` shape (and the same page-number
 * pagination) `listMyReports` returns, from the API lane's new `GET /v1/guilds/{id}/reports`
 * endpoint -- authorized the same way a report `patch` already requires guild standing for
 * (`api/internal/reports/handler.go`'s `guildReports`). A signed-out visitor never reaches
 * this call at all (the tab itself only exists for a signed-in member, §4.C.1), so there is
 * no "empty page for no session" case to swallow the way `listMyReports` does.
 */
export async function listGuildReports(
  guildId: number,
  page: number = 1,
  apiBase: string = API_BASE_URL,
): Promise<MyReportPage> {
  return query<MyReportPage>(
    guildReportsKey(apiBase, guildId, page),
    async () =>
      (await call<MyReportPage>(`/v1/guilds/${guildId}/reports?page=${page}`, apiBase)) ?? EMPTY_REPORT_PAGE,
    { scope: 'private', ttlMs: MY_REPORTS_TTL_MS },
  );
}
