// web/src/lib/rankings/api.ts
// Everything the rankings, character and guild views read. Shapes are the interface
// contract's, field for field -- including `ruleset` everywhere a realm used to be, since
// Forever has four rulesets per region and no realms. Tasks 19 (/rankings/<encounter-slug>)
// and 20 (character and guild pages) import this module as-is: it is a shared surface, not
// a private helper of the report island's Rankings mode.
//
// These reads are public -- a ranking is not tied to a signed-in visitor -- so this module
// goes through account/api.ts's requestEnvelope the same way every other API-reading
// module in this codebase does (account/api.ts itself, and web/src/lib/upload/
// multipart.ts), but passes `credentials: 'omit'`: a ranking read does not need the
// session cookie, and there is no reason to send it. Fix round 1 (Task 17 review) replaced
// this module's own duplicate fetch/envelope-parse implementation with that shared one --
// see `get()` below for how the two modules' slightly different failure readings are
// reconciled.
import { AccountError, requestEnvelope, type EnvelopeResult } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import type { CharacterPath } from '../characters';

export const RANKINGS_FAILED = 'Rankings did not load';

export class RankingsError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'RankingsError';
  }
}

/** The contract's Amendments metric enum, plus `execution` from the simulator contract. */
export type RankingMetric = 'dps' | 'hps' | 'damage_taken' | 'execution';

/** A ranking row's moderation state. */
export type RankingState = 'ok' | 'at_risk' | 'removed';

export interface RankingRow {
  rank: number;
  player: { key: string; name: string; class: string; spec: string };
  guild?: { name: string; ruleset: string; region: string };
  value: number;
  size: number;
  fought_at: string;
  duration_ms: number;
  /** "31/20/0". */
  talent_split: string;
  build_id?: string;
  trinkets: number[];
  buff_count: number;
  report_id: string;
  fight_index: number;
  state: RankingState;
  /** actual / simulated, clamped to [0, 2]. Null when the spec is not validated or the
   *  fight predates scoring. */
  execution_score: number | null;
}

export interface RankingsPage {
  rows: RankingRow[];
  total: number;
  page: number;
  per_page: number;
  updated_at: string;
}

/** No moderation state here: guild rows are aggregates, not a single ranked kill. */
export interface GuildRankingRow {
  rank: number;
  guild: { name: string; ruleset: string; region: string };
  value: number;
  fought_at: string;
  report_id: string;
}

/** One ranked fight, in `best[]` and in `history[]` alike: the API returns the same row shape for both. */
export interface CharacterFight {
  /** The encounter's display name. */
  encounter: string;
  encounter_id: number;
  difficulty: number;
  metric: RankingMetric;
  value: number;
  percentile?: number;
  spec?: string;
  fought_at: string;
  report_id: string;
  fight_index: number;
  /** actual / simulated, clamped to [0, 2]. Null when the spec is not validated or the
   *  fight predates scoring. */
  execution_score: number | null;
}

/**
 * A build seen at pull. There is no build_id: a planner build hashes the order points were
 * spent in, and the combat log does not record that. The split is text, not a link.
 */
export interface BuildSeen {
  /** "31/20/0". */
  talent_split: string;
  spec?: string;
  first_seen: string;
}

export interface CharacterPage {
  character: { name: string; region: string; ruleset: string; class?: string };
  best: CharacterFight[];
  history: CharacterFight[];
  builds_seen: BuildSeen[];
}

export interface GuildProgressionRow {
  encounter: string;
  encounter_id: number;
  difficulty: number;
  kills: number;
  pull_count: number;
  /** Absent until the boss dies. */
  first_kill_at?: string;
}

export interface GuildRosterBest {
  player: { key: string; name: string; class: string; spec: string };
  encounter: string;
  encounter_id: number;
  metric: RankingMetric;
  value: number;
  fought_at: string;
  /** actual / simulated, clamped to [0, 2]. Null when the spec is not validated or the
   *  fight predates scoring. There is no `report_id`/`fight_index` here -- `roster_best`
   *  rows are per-encounter aggregates, not a single fight, so this score is never a link
   *  into compare mode the way the rankings and character rows' scores are. */
  execution_score: number | null;
}

export interface GuildPage {
  // `id` (plan ruling 2): every session-gated guild endpoint (spec section 2.6) is
  // addressed by numeric id, and this public read is the only way a browser learns which
  // guild a region/ruleset/name resolves to. Not in the spec's own GuildPage row — flagged
  // for the coordinator to route to api/internal/rankings/guilds.go.
  guild: { id: number; name: string; region: string; ruleset: string };
  progression: GuildProgressionRow[];
  roster_best: GuildRosterBest[];
  reports: { id: string; title: string; zone: string; created_at: string }[];
}

/** One row of GET /v1/encounters: every encounter any report has ever ranked. */
export interface EncounterOption {
  id: number;
  name: string;
  slug: string;
}

export interface RankingsQuery {
  /** The numeric encounter id or the `/rankings/<encounter-slug>` slug -- the API accepts either. */
  encounter: string;
  difficulty?: number;
  metric?: RankingMetric;
  spec?: string;
  class?: string;
  phase?: string;
  region?: string;
  ruleset?: string;
  faction?: string;
  since?: string;
  page?: number;
}

/**
 * The `/rankings/<encounter-slug>` identifier a fight's own display name maps to -- the
 * same slug `GET /v1/rankings?encounter=` accepts in place of the numeric id. Lowercased,
 * every run of non-alphanumeric characters (spaces, apostrophes, punctuation) collapsed to
 * one hyphen, and any leading or trailing hyphen trimmed: "Warden Kelthas" becomes
 * "warden-kelthas", "Skolex the Insatiable" becomes "skolex-the-insatiable".
 *
 * This is the one place that derivation lives. The report island's own Rankings mode
 * (ReportView.svelte's `encounterSlug` value, computed from the selected fight's name) and
 * Task 19's `/rankings/<encounter-slug>` route both need exactly this identifier for
 * exactly the same encounter name, and two independent derivations of one identifier drift
 * apart at the first name with punctuation in it -- an apostrophe becoming its own hyphen
 * in one copy and not the other, say. Both import this function rather than reimplementing
 * the regex.
 */
export function encounterSlug(name: string): string {
  return name
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '');
}

/** Only the filters that are set reach the URL, so a bare rankings link stays short. */
function search(query: Record<string, string | number | undefined>): string {
  const params = new URLSearchParams();
  for (const [key, value] of Object.entries(query)) {
    if (value === undefined || value === '') continue;
    params.set(key, String(value));
  }
  const text = params.toString();
  return text === '' ? '' : `?${text}`;
}

/**
 * The shared transport (`account/api.ts`'s `requestEnvelope`) reads failure off the HTTP
 * status alone, unchanged from how the account module has always used it -- a caller that
 * needs `data` to be non-null on success checks that itself, the same way
 * `web/src/lib/upload/multipart.ts`'s `post()` already does. Every type this module
 * exports is non-nullable, so that check happens here: a resolved-but-null `data` (an
 * envelope that was `ok` with nothing in it, or a 200 response whose body did not parse)
 * becomes a `RankingsError` using the envelope's own message when the response carried
 * one, the same message fidelity the original standalone implementation had.
 */
async function get<T>(path: string, apiBase: string): Promise<T> {
  let result: EnvelopeResult<T>;
  try {
    result = await requestEnvelope<T>(path, apiBase, {
      credentials: 'omit',
      failureMessage: RANKINGS_FAILED,
    });
  } catch (error) {
    if (error instanceof AccountError) throw new RankingsError(error.message, error.status);
    throw new RankingsError(RANKINGS_FAILED, 0);
  }
  if (result.data === null) throw new RankingsError(result.message ?? RANKINGS_FAILED, result.status);
  return result.data;
}

export function fetchRankings(query: RankingsQuery, apiBase: string = API_BASE_URL): Promise<RankingsPage> {
  return get<RankingsPage>(`/v1/rankings${search({ ...query })}`, apiBase);
}

export function fetchGuildRankings(
  query: { encounter: string; kind: string; phase?: string },
  apiBase: string = API_BASE_URL,
): Promise<{ rows: GuildRankingRow[] }> {
  return get<{ rows: GuildRankingRow[] }>(`/v1/rankings/guilds${search({ ...query })}`, apiBase);
}

/**
 * Every encounter any report has ever ranked, for the /rankings picker: a visitor who
 * has not been handed an `/rankings/<encounter-slug>` link has no way to name one for
 * GET /v1/rankings, so the picker reads this first and offers what exists.
 */
export function fetchEncounters(apiBase: string = API_BASE_URL): Promise<{ rows: EncounterOption[] }> {
  return get<{ rows: EncounterOption[] }>('/v1/encounters', apiBase);
}

export function fetchCharacter(path: CharacterPath, apiBase: string = API_BASE_URL): Promise<CharacterPage> {
  return get<CharacterPage>(`/v1/characters/${path.region}/${path.ruleset}/${path.slug}`, apiBase);
}

export function fetchGuild(path: CharacterPath, apiBase: string = API_BASE_URL): Promise<GuildPage> {
  return get<GuildPage>(`/v1/guilds/${path.region}/${path.ruleset}/${path.slug}`, apiBase);
}
