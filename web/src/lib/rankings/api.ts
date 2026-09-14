// web/src/lib/rankings/api.ts
// Everything the rankings, character and guild views read. Shapes are the interface
// contract's, field for field -- including `ruleset` everywhere a realm used to be, since
// Forever has four rulesets per region and no realms. Tasks 19 (/rankings/<encounter-slug>)
// and 20 (character and guild pages) import this module as-is: it is a shared surface, not
// a private helper of the report island's Rankings mode.
//
// These reads are public -- a ranking is not tied to a signed-in visitor -- so unlike
// web/src/lib/account/api.ts's requestEnvelope this sends no credentials and no CSRF
// header. Going through account/api.ts's primitive would mean sending a session cookie an
// anonymous ranking read does not need, so this module owns its own thin GET, parsing the
// same `{ ok, data, error, request_id }` envelope every API response uses.
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
  /** ok | at_risk | removed. */
  state: string;
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
  metric: string;
  value: number;
  percentile?: number;
  spec?: string;
  fought_at: string;
  report_id: string;
  fight_index: number;
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
  metric: string;
  value: number;
  fought_at: string;
}

export interface GuildPage {
  guild: { name: string; region: string; ruleset: string };
  progression: GuildProgressionRow[];
  roster_best: GuildRosterBest[];
  reports: { id: string; title: string; zone: string; created_at: string }[];
}

export interface RankingsQuery {
  /** The numeric encounter id or the `/rankings/<encounter-slug>` slug -- the API accepts either. */
  encounter: string;
  difficulty?: number;
  metric?: string;
  spec?: string;
  class?: string;
  phase?: string;
  region?: string;
  ruleset?: string;
  faction?: string;
  since?: string;
  page?: number;
}

interface Envelope<T> {
  ok: boolean;
  data: T | null;
  error: { message?: string } | null;
}

async function get<T>(path: string, apiBase: string): Promise<T> {
  let response: Response;
  try {
    response = await fetch(new Request(`${apiBase}${path}`));
  } catch {
    throw new RankingsError(RANKINGS_FAILED, 0);
  }
  if (!response.ok) throw new RankingsError(RANKINGS_FAILED, response.status);
  let envelope: Envelope<T>;
  try {
    envelope = (await response.json()) as Envelope<T>;
  } catch {
    throw new RankingsError(RANKINGS_FAILED, response.status);
  }
  if (!envelope.ok || envelope.data === null) {
    throw new RankingsError(envelope.error?.message ?? RANKINGS_FAILED, response.status);
  }
  return envelope.data;
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

export function fetchRankings(query: RankingsQuery, apiBase: string = API_BASE_URL): Promise<RankingsPage> {
  return get<RankingsPage>(`/v1/rankings${search({ ...query })}`, apiBase);
}

export function fetchGuildRankings(
  query: { encounter: string; kind: string; phase?: string },
  apiBase: string = API_BASE_URL,
): Promise<{ rows: GuildRankingRow[] }> {
  return get<{ rows: GuildRankingRow[] }>(`/v1/rankings/guilds${search({ ...query })}`, apiBase);
}

export function fetchCharacter(path: CharacterPath, apiBase: string = API_BASE_URL): Promise<CharacterPage> {
  return get<CharacterPage>(`/v1/characters/${path.region}/${path.ruleset}/${path.slug}`, apiBase);
}

export function fetchGuild(path: CharacterPath, apiBase: string = API_BASE_URL): Promise<GuildPage> {
  return get<GuildPage>(`/v1/guilds/${path.region}/${path.ruleset}/${path.slug}`, apiBase);
}
