// web/src/lib/character/history-api.ts
// The owner-only character history (experience spec 2026-10-24-character-history section 10):
// the score snapshots, the kills per boss, the gaps the rating module names and the fight rows
// behind every point. The types mirror that JSON field for field. The read is private and never
// cached: the route answers `Cache-Control: private, no-store`, and a stale copy would show a
// chart that disagrees with the fights just uploaded.
import { AccountError, requestEnvelope } from '../account/api';
import type { CharacterPath } from '../characters';
import { API_BASE_URL } from '../planner/config';
import type { CharacterFight } from '../rankings/api';
import { historyCopy } from './history-copy';

/** The owner is on the free plan: the API answers 402 (the web never calls without the flag). */
export const HISTORY_PREMIUM_REQUIRED_STATUS = 402;

export type HistoryVisibility = 'public' | 'guild' | 'private';

/** One rated fight whose score snapshot was kept, oldest first. `score` is on the 0 to 2100 scale. */
export interface HistoryScorePoint {
  /** RFC 3339, UTC. */
  at: string;
  score: number;
  tier: string;
  report_id: string;
  fight_index: number;
  encounter: string;
  /** The log's difficulty id, named by `difficultyLabel`. */
  difficulty: number;
  parse_percentile?: number;
  kill: boolean;
  visibility: HistoryVisibility;
}

export interface HistoryKill {
  at: string;
  parse_percentile: number;
  report_id: string;
  fight_index: number;
}

/** One boss and difficulty, kills only, oldest first. */
export interface HistoryEncounter {
  encounter: string;
  encounter_id: number;
  /** The log's difficulty id, named by `difficultyLabel`. */
  difficulty: number;
  points: HistoryKill[];
}

/** A stretch over 14 days between two UTC dates that have a score point. The web draws these
 *  and derives none; `reason` is printed as given. */
export interface HistoryGap {
  /** YYYY-MM-DD. */
  from: string;
  /** YYYY-MM-DD. */
  to: string;
  reason: string;
}

/** A row of the fights lists: the public page's `CharacterFight` plus where its report is visible. */
export interface HistoryFight extends CharacterFight {
  visibility: HistoryVisibility;
}

export interface CharacterHistory {
  /** YYYY-MM-DD. */
  as_of: string;
  score: HistoryScorePoint[];
  encounters: HistoryEncounter[];
  gaps: HistoryGap[];
  /** Newest first, one per (report_id, fight_index). */
  fights: HistoryFight[];
}

export class HistoryApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'HistoryApiError';
  }
}

/** How a fight is keyed everywhere: a score point, a kill, a list row. */
export function fightKey(reportId: string, fightIndex: number): string {
  return `${reportId}:${fightIndex}`;
}

function historyPath(path: CharacterPath): string {
  return `/v1/characters/${path.region}/${path.ruleset}/${path.slug}/history`;
}

/** The signed-in owner's history of one character. Throws HistoryApiError (402 when the plan lapsed). */
export async function fetchCharacterHistory(
  path: CharacterPath,
  apiBase: string = API_BASE_URL,
): Promise<CharacterHistory> {
  try {
    const { data, status } = await requestEnvelope<CharacterHistory>(historyPath(path), apiBase, {
      failureMessage: historyCopy.loadFailed,
    });
    if (data === null) throw new HistoryApiError(historyCopy.loadFailed, status);
    return data;
  } catch (error) {
    if (error instanceof HistoryApiError) throw error;
    if (error instanceof AccountError) throw new HistoryApiError(error.message, error.status);
    throw new HistoryApiError(historyCopy.loadFailed, 0);
  }
}
