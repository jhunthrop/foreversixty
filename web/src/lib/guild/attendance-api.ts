// web/src/lib/guild/attendance-api.ts
// The attendance calls (design/specs/2026-10-27-guild-attendance.md section 9): the match picker's
// read and the officer writes. Same transport and error class as raids-api.ts, so a failure is a
// `GuildApiError` carrying the API's `code` (`report_matched`, `conflict`, `invalid`,
// `already_posted`, `post_failed`, `guild_plan_required`, `forbidden`). Every write answering a
// raid view drops the guild's cached reads through `invalidateGuild`; the post does too, since the
// raid view carries `report.posted_at`.
import { API_BASE_URL } from '../planner/config';
import {
  GuildApiError,
  call,
  invalidateGuild,
  type GuildRaidView,
  type MatchCandidates,
  type PostReportResult,
} from './api';
import { characterKeyPath, raidPath, type RaidId } from './raids-api';

export { fetchGuildReadiness } from './api';

const ALREADY_POSTED = 'already_posted';

/** The picker's read: fetched when it opens, never cached (the list changes as logs upload). */
export function fetchMatchCandidates(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<MatchCandidates> {
  return call<MatchCandidates>(`${raidPath(guildId, raidId)}/match-candidates`, apiBase);
}

async function raidWrite(
  guildId: number,
  path: string,
  apiBase: string,
  init: { method: string; body?: unknown },
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(path, apiBase, init);
  invalidateGuild(guildId, apiBase);
  return raid;
}

/** 409 `report_matched` when another raid holds the log; 409 `conflict` before the raid ends. */
export function setMatch(
  guildId: number,
  raidId: RaidId,
  reportId: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  return raidWrite(guildId, `${raidPath(guildId, raidId)}/match`, apiBase, {
    method: 'PUT',
    body: { report_id: reportId },
  });
}

/** The raid returns to its `pre_done_status`; the auto-matcher never picks that log for it again. */
export function clearMatch(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  return raidWrite(guildId, `${raidPath(guildId, raidId)}/match`, apiBase, { method: 'DELETE' });
}

export function setExcused(
  guildId: number,
  raidId: RaidId,
  characterKey: string,
  excused: boolean,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  return raidWrite(
    guildId,
    `${raidPath(guildId, raidId)}/attendance/${characterKeyPath(characterKey)}`,
    apiBase,
    { method: 'PUT', body: { excused } },
  );
}

/** 409 `conflict` (raid not done), 409 `already_posted`, 502 `post_failed` (retryable). */
export async function postReport(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<PostReportResult> {
  try {
    const result = await call<PostReportResult>(`${raidPath(guildId, raidId)}/report/post`, apiBase, {
      method: 'POST',
    });
    invalidateGuild(guildId, apiBase);
    return result;
  } catch (error) {
    // Another officer posted first: the cached raid view still lacks `report.posted_at`.
    if (error instanceof GuildApiError && error.code === ALREADY_POSTED) invalidateGuild(guildId, apiBase);
    throw error;
  }
}
