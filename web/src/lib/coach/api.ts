// web/src/lib/coach/api.ts
// Every call the browser makes to the coach routes (Tasks 30 and 31). Calls go through the
// account module's `requestEnvelope` (session cookie, JSON envelope); a failed answer becomes a
// `CoachApiError` carrying the HTTP status and the API's own `error.code`. Answers are
// normalised, so an older API shape never throws here.
import { AccountError, requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import { normalizeFightCoach, normalizeHistory, normalizeNightLine, normalizeRecent } from './normalize';
import type { CoachFull, CoachHistoryPage, CoachRecent, CoachTeaser, NightLine } from './types';

export const COACH_FAILED = 'The coach could not be reached';

export class CoachApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
    /** The API's own `error.code` ("no_plan", "flagged"), when the envelope carried one. */
    readonly code: string | null = null,
  ) {
    super(message);
    this.name = 'CoachApiError';
  }
}

/** GET through the shared envelope; 404 is `null`, any other failure is a `CoachApiError`. */
async function get(path: string, apiBase: string): Promise<unknown> {
  try {
    const { data } = await requestEnvelope<unknown>(path, apiBase, { failureMessage: COACH_FAILED });
    return data;
  } catch (error) {
    if (!(error instanceof AccountError)) throw error;
    if (error.status === 404) return null;
    throw new CoachApiError(error.message, error.status, error.code);
  }
}

/** The path and query of one route; every dynamic part is encoded here. */
function coachPath(path: string, params: Readonly<Record<string, string | undefined>> = {}): string {
  const query = new URLSearchParams();
  for (const [key, value] of Object.entries(params)) if (value !== undefined) query.set(key, value);
  const text = query.toString();
  return text === '' ? path : `${path}?${text}`;
}

const reportPath = (reportId: string): string => `/v1/reports/${encodeURIComponent(reportId)}`;

/** One player's review of one fight: a review, a teaser, or `null` when the API answers 404. */
export async function fetchFightCoach(
  reportId: string,
  fightIndex: number,
  playerKey: string,
  apiBase: string = API_BASE_URL,
  /** `line` is the row strip's read-only form (`line=1`): it never queues a generation. */
  options: { line?: boolean } = {},
): Promise<CoachFull | CoachTeaser | null> {
  const path = coachPath(`${reportPath(reportId)}/fights/${fightIndex}/coach`, {
    player: playerKey,
    line: options.line === true ? '1' : undefined,
  });
  return normalizeFightCoach(await get(path, apiBase));
}

/** The account's last raid line and, for Premium, the newest stored reviews. */
export async function fetchCoachRecent(apiBase: string = API_BASE_URL): Promise<CoachRecent> {
  return normalizeRecent(await get('/v1/me/coach/recent', apiBase));
}

/** One page of a character's stored reviews, newest first. */
export async function fetchCoachHistory(
  region: string,
  ruleset: string,
  name: string,
  cursor?: string,
  apiBase: string = API_BASE_URL,
): Promise<CoachHistoryPage> {
  const base = `/v1/characters/${[region, ruleset, name].map(encodeURIComponent).join('/')}/coach/history`;
  return normalizeHistory(await get(coachPath(base, { cursor }), apiBase));
}

/** The night line for one character across a report; `null` when there is none. */
export async function fetchNightLine(
  reportId: string,
  playerKey: string,
  apiBase: string = API_BASE_URL,
): Promise<NightLine | null> {
  return normalizeNightLine(
    await get(coachPath(`${reportPath(reportId)}/coach/night`, { player: playerKey }), apiBase),
  );
}
