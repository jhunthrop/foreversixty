// web/src/lib/character/standing.ts
// The owner's standing line and the character's activity (design/specs/2026-11-06-character.md
// section 6, rulings C14 to C22): the character's place in its guild's sim ladder, its spec's
// grade this week as a link to the tier list (never home's tile), the next guild raid as a
// short clock, and its own rows of the site feed (GET /v1/feed?character=).
import { requestEnvelope } from '../account/api';
import { query } from '../data/query';
import type { NextRaidView } from '../guild/api';
import type { FeedPage } from '../home/feed';
import type { GuildSims, SpecWeek } from '../home/home-answer';
import { API_BASE_URL } from '../planner/config';
import { clockLabel, clockParts } from '../ui/countdown';
import { fillShared } from '../copy/shared';
import { sharedCharacterActivity, sharedCharacterActivityFailed } from '../copy/shared/character.generated';
import {
  sharedStandingGraded,
  sharedStandingPlace,
  sharedStandingRaid,
  sharedStandingRank,
  sharedStandingRoleDps,
  sharedStandingRoleHealer,
  sharedStandingRoleTank,
} from '../copy/shared/standing.generated';

/** C18: the activity rows the page shows. */
export const ACTIVITY_ROWS = 5;
const FEED_TTL_MS = 60_000;
const MAX_LEVEL = 60;

const ROLE_WORD = {
  dps: sharedStandingRoleDps,
  healer: sharedStandingRoleHealer,
  tank: sharedStandingRoleTank,
} as const;

export const standingCopy = {
  activity: sharedCharacterActivity,
  failed: sharedCharacterActivityFailed,
} as const;

export function activityPath(characterKey: string, limit = ACTIVITY_ROWS): string {
  const params = new URLSearchParams({ character: characterKey, limit: String(limit) });
  return `/v1/feed?${params.toString()}`;
}

/** The character's own feed rows, cached a minute; a session adds its guild's rows. */
export function fetchActivity(
  characterKey: string,
  signedIn: boolean,
  apiBase: string = API_BASE_URL,
): Promise<FeedPage> {
  const path = activityPath(characterKey);
  return query<FeedPage>(
    `${apiBase}${path}`,
    async () => {
      const { data } = await requestEnvelope<FeedPage>(path, apiBase, {
        credentials: signedIn ? 'include' : 'omit',
      });
      return data ?? { rows: [], next_cursor: null };
    },
    { scope: signedIn ? 'private' : 'public', ttlMs: FEED_TTL_MS },
  );
}

/** C16: "#3 of 14 DPS in Olympus XXVII, 60 raid"; null with no own row on the ladder. */
export function guildRankLine(sims: GuildSims | null): string | null {
  if (sims === null) return null;
  const own = sims.you ?? sims.rows.find((row) => row.you) ?? null;
  if (own === null) return null;
  const band = sims.band >= MAX_LEVEL ? String(MAX_LEVEL) : `${sims.band} to ${sims.band + 9}`;
  const preset = sims.preset === 'raid' ? ' raid' : '';
  return fillShared(sharedStandingRank, {
    rank: own.rank,
    of: sims.of,
    role: ROLE_WORD[sims.role],
    guild: sims.guild_name,
    band: `${band}${preset}`,
  });
}

/** C19: "Holy tier A", linked to the tier list; a tank or healer list has no letter: "Holy #4 of 9 on the tier list". */
export function gradeLine(week: SpecWeek | null, specName: string | null): string | null {
  if (week === null || specName === null) return null;
  return week.grade === null
    ? fillShared(sharedStandingPlace, { spec: specName, place: week.place, of: week.of })
    : fillShared(sharedStandingGraded, { spec: specName, grade: week.grade });
}

/** C21: "Onyxia's Lair in 33 days, 8 hours", the clock to its two largest units; null once it has
 *  started. The sign-up count stays on the raid's own page: after the clock it read as part of it
 *  (review 2026-11-09). */
/** The standing line's raid clock keeps its two largest units. */
const CLOCK_UNITS = 2;

export function raidLine(raid: NextRaidView | null, now: Date): string | null {
  if (raid === null) return null;
  const parts = clockParts(new Date(raid.starts_at), now);
  if (parts === null) return null;
  const values = { raid: raid.title, clock: clockLabel(parts.slice(0, CLOCK_UNITS)) };
  return fillShared(sharedStandingRaid, values);
}
