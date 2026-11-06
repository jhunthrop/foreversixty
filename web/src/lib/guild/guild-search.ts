// web/src/lib/guild/guild-search.ts
// The guild search a 404 guild page asks for its "Did you mean" (guild loop round 3): public
// `GET /v1/guilds/search?q=&region=`, at most 5 rows `{name, region, ruleset, path}`. A stranger has
// no guilds of their own and may find the guild on no board, so this is the source that reaches
// them. Rows that do not carry the three names are dropped; the link is built from them
// (`guildHref`), never taken from the row's `path`.
import { API_BASE_URL } from '../planner/config';
import { call } from './api';
import type { GuildCandidate } from './guild-suggest';

interface SearchRow {
  name?: unknown;
  region?: unknown;
  ruleset?: unknown;
}

function candidateOf(row: SearchRow): GuildCandidate | null {
  const { name, region, ruleset } = row;
  if (typeof name !== 'string' || typeof region !== 'string' || typeof ruleset !== 'string') return null;
  return { name, region, ruleset };
}

/** The search's rows as candidates, in the API's order; `rows` or a bare array both read. */
export function searchCandidates(data: unknown): GuildCandidate[] {
  const rows = Array.isArray(data) ? data : (data as { rows?: unknown } | null)?.rows;
  if (!Array.isArray(rows)) return [];
  return rows
    .map((row) => candidateOf((row ?? {}) as SearchRow))
    .filter((c): c is GuildCandidate => c !== null);
}

export async function searchGuilds(
  q: string,
  region: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildCandidate[]> {
  const params = new URLSearchParams({ q, region });
  return searchCandidates(await call<unknown>(`/v1/guilds/search?${params.toString()}`, apiBase));
}
