// web/src/lib/guild/night-patch.ts
// The optimistic edit a raid view takes before a re-read: the posted stamp, so `Posted to Discord
// at {time}.` shows at once. Returns a new raid.
import type { GuildRaidView } from './api';

/** The report stamped as posted. */
export function withPosted(raid: GuildRaidView, postedAt: string): GuildRaidView {
  return raid.report === null ? raid : { ...raid, report: { ...raid.report, posted_at: postedAt } };
}
