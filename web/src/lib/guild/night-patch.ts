// web/src/lib/guild/night-patch.ts
// The two optimistic edits a raid view takes before its write answers: an Excused tick (the row and
// the night's `excused` count, so the summary sentence moves with the checkbox) and the posted stamp
// (so `Posted to Discord at {time}.` replaces the button at once). Both return a new raid.
import type { GuildRaidView } from './api';

/** `excused` set on one attendance row, `night.excused` moved by one when the row changed. */
export function withExcused(raid: GuildRaidView, characterKey: string, excused: boolean): GuildRaidView {
  const rows = raid.attendance;
  const row = rows?.find((r) => r.character_key === characterKey);
  if (rows === null || row === undefined || row.excused === excused) return raid;
  const night =
    raid.night === null ? null : { ...raid.night, excused: raid.night.excused + (excused ? 1 : -1) };
  return {
    ...raid,
    night,
    attendance: rows.map((r) => (r.character_key === characterKey ? { ...r, excused } : r)),
  };
}

/** The report stamped as posted. */
export function withPosted(raid: GuildRaidView, postedAt: string): GuildRaidView {
  return raid.report === null ? raid : { ...raid, report: { ...raid.report, posted_at: postedAt } };
}
