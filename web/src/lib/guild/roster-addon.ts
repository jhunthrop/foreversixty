// web/src/lib/guild/roster-addon.ts
// Task L8.3: the Roster tab's "In game" column, from the guild's newest addon roster snapshot
// (Task L6.14: `roster[].addon` and `addon_snapshot` on the guild home). Pure, so the
// component holds no formatting of its own.
import { relativeTime } from '../dates';
import { SPEC_BY_KEY } from '../sim/specs';
import type { GuildAddonSnapshot, GuildRosterRow } from './api';
import { guildRosterAddonCopy } from './copy';

/** The spec's own name ("Beast Mastery") for the snapshot's slug; an unknown slug is printed
 *  unchanged rather than guessed at. Read from the generated spec list directly, not through
 *  sim/spec-label.ts, so the lazy roster chunk does not pull in the planner's class data. */
export function addonSpecName(slug: string): string {
  return SPEC_BY_KEY[slug]?.name ?? slug;
}

/**
 * What the snapshot says about who the member is, where it differs from the site's row.
 * The spec shows only when it differs from (or the site lacks) the row's own spec. The level
 * always shows: the site's roster row carries no level of its own to agree with (the
 * guild-centre contract's `RosterRow`), so the addon's is the only one there is. `null` for a
 * row the snapshot did not carry.
 */
export function addonIdentity(row: GuildRosterRow): { level: number | null; spec: string | null } | null {
  const addon = row.addon;
  if (addon === undefined || addon === null) return null;
  if (addon.spec === null) return { level: addon.level, spec: null };
  const name = addonSpecName(addon.spec);
  const same = row.spec !== undefined && row.spec.toLowerCase() === name.toLowerCase();
  return { level: addon.level, spec: same ? null : name };
}

/** "gap 12 DPS": the addon's integer gap (it may be negative), under the addon's own heading. */
export function addonGapLabel(gap: number | null): string | null {
  return gap === null ? null : guildRosterAddonCopy.gap(gap);
}

/** "21:40" from the member's local "YYYY-MM-DD HH:MM"; null for anything else. */
export function addonSeenClock(seenAt: string): string | null {
  return /^\d{4}-\d{2}-\d{2} (\d{2}:\d{2})$/.exec(seenAt)?.[1] ?? null;
}

/** The member's `seen_at` (their own local "YYYY-MM-DD HH:MM") as an instant, read in the
 *  viewer's zone: the addon announces no zone, so the clock is taken as the viewer's own, the
 *  same reading the rest of the page makes of a local stamp. Null for anything else. */
export function addonSeenAt(seenAt: string): Date | null {
  const m = /^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2})$/.exec(seenAt);
  if (m === null) return null;
  const [, y, mo, d, h, mi] = m.map(Number);
  const at = new Date(y, mo - 1, d, h, mi);
  return Number.isNaN(at.getTime()) ? null : at;
}

/**
 * "seen 2 hours ago", relative to `now` (a stamp ahead of the viewer's clock reads "just now",
 * never a negative age). `withheld` is a member who shares nothing else: the hover then says so
 * instead of naming the clock. Null for a `seen_at` that is not "YYYY-MM-DD HH:MM".
 */
export function addonSeen(
  seenAt: string,
  now: Date,
  withheld = false,
): { label: string; hover: string } | null {
  const clock = addonSeenClock(seenAt);
  const at = addonSeenAt(seenAt);
  if (clock === null || at === null) return null;
  return {
    label: guildRosterAddonCopy.seen(relativeTime(at, now)),
    hover: withheld ? guildRosterAddonCopy.withheldHover : guildRosterAddonCopy.seenHover(clock),
  };
}

/** The sender's name: `by_character` is a character key (`us/pvp/thorgrim`), so the name is
 *  that row's; a sender on no row reads as the key's last segment, capitalised. */
export function snapshotSenderName(byCharacter: string, rows: readonly GuildRosterRow[]): string {
  const row = rows.find((r) => r.character_key === byCharacter);
  if (row !== undefined) return row.name;
  const last = byCharacter.split('/').pop() ?? byCharacter;
  return last.charAt(0).toUpperCase() + last.slice(1);
}

/** The line above the table: who sent the newest snapshot and when. A home without the field
 *  (before the addendum) reads as none. How many people are on the roster, and in what state,
 *  is the count line (roster-model.ts). */
export function addonSnapshotLine(
  snapshot: GuildAddonSnapshot | null | undefined,
  rows: readonly GuildRosterRow[],
  now: Date = new Date(),
): string {
  if (snapshot === undefined || snapshot === null) return guildRosterAddonCopy.noSnapshot;
  return guildRosterAddonCopy.snapshotLine(
    snapshotSenderName(snapshot.by_character, rows),
    relativeTime(new Date(snapshot.received_at), now),
  );
}
