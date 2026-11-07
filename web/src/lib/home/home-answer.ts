// web/src/lib/home/home-answer.ts
// GET /v1/home (design/specs/2026-11-07-home-b.md §2): the signed-in home's own blocks for one
// character: the sim on import, the spec's grade this week, the top movers and the guild's sims.
// Types follow api/openapi.yaml; a block the API has nothing true for is null.
import { requestEnvelope } from '../account/api';
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';

export type TierGrade = 'S' | 'A' | 'B' | 'C' | 'D';

export interface ImportSim {
  state: 'pending' | 'done' | 'failed';
  dps: number | null;
  simmed_at: string | null;
  sim_id: string | null;
}

export interface SpecWeek {
  spec: string;
  band: number;
  faction: string;
  role: 'dps' | 'tank' | 'healer';
  /** Null for a tank or healer list, which carries a place but no letter. */
  grade: TierGrade | null;
  /** Null until last week's grade exists: never a guessed delta. */
  last_week_grade: TierGrade | null;
  place: number;
  of: number;
}

export interface MoverRow {
  rank: number;
  name: string;
  class: string;
  key: string;
  /** Places gained since last week's board; null with no baseline. */
  delta: number | null;
}

export interface Movers {
  spec: string;
  band: number;
  /** False before the first weekly board: the rows are the top 3, not movers. */
  baseline: boolean;
  rows: MoverRow[];
}

export interface GuildSimRow {
  rank: number;
  name: string;
  class: string;
  key: string;
  value: number;
  delta: number | null;
  simmed_at: string;
  you: boolean;
}

export interface GuildSims {
  guild_id: number;
  guild_name: string;
  role: 'dps' | 'tank' | 'healer';
  band: number;
  preset: 'bare' | 'raid';
  unit: string;
  of: number;
  rows: GuildSimRow[];
  /** The viewer's own row when it is outside the top 5; null inside it, or with no sim. */
  you: GuildSimRow | null;
}

export interface HomeAnswer {
  import_sim: ImportSim | null;
  spec_week: SpecWeek | null;
  movers: Movers | null;
  guild_sims: GuildSims | null;
}

const HOME_TTL_MS = 60_000;

export const EMPTY_HOME: HomeAnswer = { import_sim: null, spec_week: null, movers: null, guild_sims: null };

export function homePath(characterKey: string): string {
  return `/v1/home?character=${encodeURIComponent(characterKey)}`;
}

export function fetchHomeAnswer(characterKey: string, apiBase: string = API_BASE_URL): Promise<HomeAnswer> {
  const path = homePath(characterKey);
  return query<HomeAnswer>(
    `${apiBase}${path}`,
    async () => (await requestEnvelope<HomeAnswer>(path, apiBase)).data ?? EMPTY_HOME,
    { scope: 'private', ttlMs: HOME_TTL_MS },
  );
}
