// web/src/lib/report/load.ts
// The report page's whole data boundary. Two sources, two shapes:
//
//   the API           envelope { ok, data, error, request_id }, session cookie, may 401/403
//   data_base_url     plain JSON on the edge, public, no credentials
//
// report.json and live.json carry `public, max-age=5`, which is right for a shared edge
// cache and wrong for a poller five seconds later in the same browser: `cache: 'no-cache'`
// makes the browser revalidate instead of serving its own copy, while summary.json and
// events.parquet are immutable and take the default.
import { API_BASE_URL } from '../planner/config';
import { asArray, type ReportFile, type ReportMeta, type Summary } from './types';

export const REPORT_LOAD_FAILED = 'Report data did not load';
export const REPORT_NOT_FOUND = 'No report with that id';
export const REPORT_FORBIDDEN = 'That report is not public';

/** Every failure the island renders comes out of here, so components never see a raw fetch error. */
export class ReportLoadError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'ReportLoadError';
  }
}

interface Envelope<T> {
  ok: boolean;
  data: T | null;
  error: { message?: string } | null;
  request_id: string;
}

function apiFailure(status: number): ReportLoadError {
  if (status === 404) return new ReportLoadError(REPORT_NOT_FOUND, status);
  if (status === 401 || status === 403) return new ReportLoadError(REPORT_FORBIDDEN, status);
  return new ReportLoadError(REPORT_LOAD_FAILED, status);
}

async function apiGet<T>(path: string, apiBase: string): Promise<T> {
  let response: Response;
  try {
    response = await fetch(new Request(`${apiBase}${path}`, { credentials: 'include' }));
  } catch {
    throw new ReportLoadError(REPORT_LOAD_FAILED, 0);
  }
  if (!response.ok) throw apiFailure(response.status);
  let envelope: Envelope<T>;
  try {
    envelope = (await response.json()) as Envelope<T>;
  } catch {
    throw new ReportLoadError(REPORT_LOAD_FAILED, response.status);
  }
  if (!envelope.ok || envelope.data === null) throw apiFailure(response.status);
  return envelope.data;
}

export function fetchReportMeta(id: string, apiBase: string = API_BASE_URL): Promise<ReportMeta> {
  return apiGet<ReportMeta>(`/v1/reports/${encodeURIComponent(id)}`, apiBase);
}

/**
 * Private and guild reports are refused at /logs-data/, so the island asks the API for a
 * signed base url instead. It expires in ten minutes; the island re-asks on a 403 rather
 * than holding a timer, because a viewer who leaves the page open overnight should get one
 * failed fetch and a retry, not a background request loop.
 */
export async function fetchAccessUrl(id: string, apiBase: string = API_BASE_URL): Promise<string> {
  const access = await apiGet<{ data_base_url: string }>(
    `/v1/reports/${encodeURIComponent(id)}/access`,
    apiBase,
  );
  return access.data_base_url;
}

/**
 * That re-ask, as the one place it happens.
 *
 * A signed base url lasts ten minutes and a report page is left open for hours, so every
 * read after the first ten minutes is a 403: selecting an uncached fight shows "That
 * report is not public" to someone who can plainly see it, and the live poll swallows the
 * failure and keeps the Live badge on a page that has silently stopped updating.
 *
 * So the refusal is a retry rather than a message: `resign` is asked for a fresh base and
 * the same request is made once more against it. Once -- a second 403 is the answer the
 * visitor should actually be shown, and so is a `resign` that is itself refused, which is
 * what a report that has genuinely been made private looks like. The base that worked
 * comes back with the answer, because the caller has to keep it for the next read.
 */
export async function withFreshBase<T>(
  base: string,
  resign: () => Promise<string>,
  work: (dataBaseUrl: string) => Promise<T>,
): Promise<{ value: T; base: string }> {
  try {
    return { value: await work(base), base };
  } catch (thrown) {
    const refused = thrown instanceof ReportLoadError && (thrown.status === 401 || thrown.status === 403);
    if (!refused) throw thrown;
    const fresh = await resign();
    return { value: await work(fresh), base: fresh };
  }
}

async function dataGet<T>(url: string, cache: RequestCache): Promise<T> {
  let response: Response;
  try {
    response = await fetch(new Request(url, { cache }));
  } catch {
    throw new ReportLoadError(REPORT_LOAD_FAILED, 0);
  }
  if (!response.ok) throw apiFailure(response.status);
  try {
    return (await response.json()) as T;
  } catch {
    throw new ReportLoadError(REPORT_LOAD_FAILED, response.status);
  }
}

export async function fetchReportFile(dataBaseUrl: string): Promise<ReportFile> {
  const file = await dataGet<ReportFile>(`${dataBaseUrl}/report.json`, 'no-cache');
  return { ...file, fights: asArray(file.fights), units: asArray(file.units) };
}

export function fetchSummary(dataBaseUrl: string, fightIndex: number): Promise<Summary> {
  return dataGet<Summary>(`${dataBaseUrl}/fights/${fightIndex}/summary.json`, 'default');
}

/** Null means the fight is not open; anything else is a real failure and throws. */
export async function fetchLive(dataBaseUrl: string, fightIndex: number): Promise<Summary | null> {
  try {
    return await dataGet<Summary>(`${dataBaseUrl}/fights/${fightIndex}/live.json`, 'no-cache');
  } catch (error) {
    if (error instanceof ReportLoadError && error.status === 404) return null;
    throw error;
  }
}

/** The Parquet file is fetched by the query layer, not here; this only names it. */
export function eventsUrl(dataBaseUrl: string, fightIndex: number): string {
  return `${dataBaseUrl}/fights/${fightIndex}/events.parquet`;
}

/** Spec section 1: "five-second page poll". */
export const POLL_INTERVAL_MS = 5000;

export interface Poller {
  start(): void;
  stop(): void;
}

/**
 * A poll that cannot pile up. `setInterval` would keep firing while a slow tick is still in
 * flight, so a laggy network would queue requests behind each other and the newest answer
 * would arrive last. This waits for the tick, then schedules the next one, and a tick that
 * throws is swallowed so a single bad poll does not end live updates for the session.
 */
export function createPoller(tick: () => Promise<void>, intervalMs: number = POLL_INTERVAL_MS): Poller {
  let timer: ReturnType<typeof setTimeout> | null = null;
  let live = false;

  function schedule(): void {
    if (!live) return;
    timer = setTimeout(() => {
      void run();
    }, intervalMs);
  }

  async function run(): Promise<void> {
    try {
      await tick();
    } catch {
      /* One failed poll is not the end of live updates; the next one may succeed. */
    }
    schedule();
  }

  return {
    start(): void {
      if (live) return;
      live = true;
      schedule();
    },
    stop(): void {
      live = false;
      if (timer !== null) clearTimeout(timer);
      timer = null;
    },
  };
}
