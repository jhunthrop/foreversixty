// web/src/lib/character-search-api.ts
// The typeahead's one read: GET /v1/characters/search, through rankings/api.ts's shared
// envelope transport. Nothing is cached -- a suggestion list is as fresh as the last keystroke.
import { API_BASE_URL } from './planner/config';
import { get, search } from './rankings/api';

/** One suggestion row. `class`, `level` and `spec` are null for a character not yet scanned. */
export interface CharacterSuggestion {
  key: string;
  region: string;
  ruleset: string;
  name: string;
  class: string | null;
  level: number | null;
  spec: string | null;
}

export interface CharacterSearchOptions {
  region?: string;
  ruleset?: string;
  limit?: number;
  signal: AbortSignal;
}

export type CharacterSearch = (
  text: string,
  options: CharacterSearchOptions,
) => Promise<CharacterSuggestion[]>;

export async function searchCharacters(
  text: string,
  { region, ruleset, limit, signal }: CharacterSearchOptions,
  apiBase: string = API_BASE_URL,
): Promise<CharacterSuggestion[]> {
  const path = `/v1/characters/search${search({ q: text, region, ruleset, limit })}`;
  const { rows } = await get<{ rows: CharacterSuggestion[] }>(path, apiBase, signal);
  return rows;
}
