// web/src/lib/character/standing.ts
// The owner's standing line and the character's activity (design/specs/2026-11-06-character.md
// section 6, rulings C14 to C22): the character's place in its guild's sim ladder, its spec's
// grade this week as a link to the tier list (never home's tile), the next guild raid as a
// short clock, and its own rows of the site feed (GET /v1/feed?character=).
import { requestEnvelope } from '../account/api';
import { query } from '../data/query';
import type { NextRaidView } from '../guild/api';
import type { FeedPage, FeedSegment } from '../home/feed';
import type { GuildSims, SpecWeek } from '../home/home-answer';
import { API_BASE_URL } from '../planner/config';
import { clockLabel, clockParts } from '../ui/countdown';
import { sharedCharacterActivity, sharedCharacterActivityFailed } from '../copy/shared/character.generated';
import {
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
  /** The rail panel holding the guild place, the tier grade and the raid clock. */
  standing: 'Standing',
  /** The tier list row's label after the spec: "Holy tier". */
  tierWord: 'tier',
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

/** One Standing row: what it is on the left, the figure on the right (h-audit 2026-11-09: three
 *  whole sentences wrapped and read as unlabelled lines). */
export interface StandingRow {
  label: string;
  value: string;
}

/** C16: "Olympus XXVII · 60 raid" | "#3 of 14 DPS"; null with no own row on the ladder. */
export function guildRankRow(sims: GuildSims | null): StandingRow | null {
  if (sims === null) return null;
  const own = sims.you ?? sims.rows.find((row) => row.you) ?? null;
  if (own === null) return null;
  const band = sims.band >= MAX_LEVEL ? String(MAX_LEVEL) : `${sims.band} to ${sims.band + 9}`;
  const preset = sims.preset === 'raid' ? ' raid' : '';
  return {
    label: `${sims.guild_name} · ${band}${preset}`,
    value: `#${own.rank} of ${sims.of} ${ROLE_WORD[sims.role]}`,
  };
}

/** C19: "Holy tier" | "A", linked to the tier list; a tank or healer list has no letter: "#4 of 9". */
export function gradeRow(week: SpecWeek | null, specName: string | null): StandingRow | null {
  if (week === null || specName === null) return null;
  return {
    label: `${specName} ${standingCopy.tierWord}`,
    value: week.grade === null ? `#${week.place} of ${week.of}` : week.grade,
  };
}

/** The raid clock keeps its two largest units. */
const CLOCK_UNITS = 2;

/** C21: "Onyxia's Lair" | "33 days, 8 hours", the clock to its two largest units; null once it has
 *  started. The sign-up count stays on the raid's own page (review 2026-11-09). */
export function raidRow(raid: NextRaidView | null, now: Date): StandingRow | null {
  if (raid === null) return null;
  const parts = clockParts(new Date(raid.starts_at), now);
  if (parts === null) return null;
  return { label: raid.title, value: clockLabel(parts.slice(0, CLOCK_UNITS)) };
}

/** A feed sentence on the character's own page: the character's own name drops out (the page is
 *  theirs; "Elyra Duskvale uploaded Molten Core" reads "Uploaded Molten Core"), and the text that
 *  followed it opens the sentence with a capital. Another character's name (a guild row) stays. */
export function ownActivitySegments(segments: readonly FeedSegment[], name: string): FeedSegment[] {
  const at = segments.findIndex((seg) => seg.style === 'name' && seg.text === name);
  if (at === -1) return [...segments];
  const rest = segments.slice(at + 1);
  const next = rest[0];
  if (next === undefined || next.style !== undefined) return [...segments.slice(0, at), ...rest];
  const text = next.text.trimStart();
  const opened = at === 0 ? `${text.charAt(0).toUpperCase()}${text.slice(1)}` : text;
  return [...segments.slice(0, at), { ...next, text: opened }, ...rest.slice(1)];
}
