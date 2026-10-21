// web/src/lib/guild/roster-facts.ts
// What one roster row says under the name, in the three blocks the phone layout stacks (lifecycle
// round 7, F7-4): stat chips (units named), the professions line, and the in-game facts in one
// fixed order on every row -- Level, Gear gap, spec, gear set, seen. Pure, so the order and the
// words are pinned without a page.
import { guildRosterAddonCopy } from './copy';
import type { GuildRosterEntry, GuildRosterRow } from './api';
import { addonGapLabel, addonIdentity, addonSeen } from './roster-addon';

/** "Item level 68", "5 of 8 nights", "Best DPS 238", "Rating 65": every figure with its unit
 *  named first, never a bare "65 rating". */
export function rosterStatChips(row: GuildRosterRow): string[] {
  const chips: string[] = [];
  if (row.item_level !== undefined) chips.push(`Item level ${row.item_level}`);
  if (row.attendance !== undefined)
    chips.push(`${row.attendance.present} of ${row.attendance.nights} nights`);
  if (row.best_parse !== undefined && row.best_parse !== null) {
    chips.push(`Best ${row.best_parse.metric.toUpperCase()} ${Math.round(row.best_parse.value)}`);
  }
  if (row.rating !== undefined && row.rating !== null) chips.push(`Rating ${row.rating.overall}`);
  return chips;
}

/** "Protection raid" -> "Gear set: Protection (raid)": a trailing one-word lowercase tag is the
 *  set's kind and goes in brackets; a name without one is printed as it is. */
export function gearSetLabel(build: string): string {
  const words = build.trim().split(/\s+/);
  const last = words[words.length - 1];
  const named =
    words.length > 1 && /^[a-z]+$/.test(last) ? `${words.slice(0, -1).join(' ')} (${last})` : build.trim();
  return `Gear set: ${named}`;
}

/** "Survival in game": the addon's spec, where it differs from the site's row. */
export function inGameSpecLabel(spec: string): string {
  return `${spec} in game`;
}

export type InGameFactKey = 'level' | 'gap' | 'spec' | 'build' | 'seen';
export interface InGameFact {
  key: InGameFactKey;
  text: string;
  /** The hover on the fact, when it has one (the seen time). */
  hover?: string;
}

/**
 * The in-game facts of a row the addon saw, always in this order: Level, Gear gap, spec (only
 * where it differs from the site's), gear set, seen. A fact the member withholds is absent;
 * the order of the rest never changes. Empty for a row the snapshot did not carry.
 */
export function inGameFacts(row: GuildRosterRow, now: Date): InGameFact[] {
  const identity = addonIdentity(row);
  if (identity === null || row.addon === undefined || row.addon === null) return [];
  const addon = row.addon;
  const withheld =
    addon.level === null && addon.spec === null && addon.gear_gap === null && addon.build === null;
  const facts: InGameFact[] = [];
  if (identity.level !== null) facts.push({ key: 'level', text: guildRosterAddonCopy.level(identity.level) });
  const gap = addonGapLabel(addon.gear_gap);
  if (gap !== null) facts.push({ key: 'gap', text: gap });
  if (identity.spec !== null) facts.push({ key: 'spec', text: inGameSpecLabel(identity.spec) });
  if (addon.build !== null && addon.build !== '')
    facts.push({ key: 'build', text: gearSetLabel(addon.build) });
  const seen = addonSeen(addon.seen_at, now, withheld);
  if (seen !== null) facts.push({ key: 'seen', text: seen.label, hover: seen.hover });
  return facts;
}

/** The in-game facts of a name only the addon saw: Level, then seen. */
export function entryFacts(entry: GuildRosterEntry, now: Date): InGameFact[] {
  const facts: InGameFact[] = [];
  if (entry.level !== undefined) facts.push({ key: 'level', text: guildRosterAddonCopy.level(entry.level) });
  const seen = entry.seen_at === undefined ? null : addonSeen(entry.seen_at, now);
  if (seen !== null) facts.push({ key: 'seen', text: seen.label, hover: seen.hover });
  return facts;
}

/** Shown once above the roster when some rows carry no rank (names only the addon saw). */
export const RANKS_PENDING_LINE = 'Ranks show once the addon sends them.';
