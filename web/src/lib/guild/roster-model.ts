// web/src/lib/guild/roster-model.ts
// The roster as everyone the addon saw, with a word beside each (lifecycle review round 4,
// N4/N5/N6/N13): the count line, the entries with no account row, the capability test and the
// approve-all question. Pure, so the component holds no rule of its own.
import type { GuildCapability, GuildRosterCounts, GuildRosterEntry, GuildRosterRow, GuildRank } from './api';
import { guildRosterModelCopy } from './copy';

/** Whether the account's `can` list (the guild entry on /v1/me) names `word`. The page draws a
 *  control exactly when this is true and never infers one from the rank. */
export function can(list: readonly string[] | undefined, word: GuildCapability): boolean {
  return list !== undefined && list.includes(word);
}

/** The state word beside an entry: Approved, Waiting for approval, No account yet, Not seen in
 *  game yet. */
export function stateWord(state: GuildRosterEntry['state']): string {
  return guildRosterModelCopy.states[state];
}

/** The rank word for everyone: Leader, Officer, and Member for a plain member (never blank). */
export function rankWord(rank: GuildRank | string | undefined): string {
  if (rank === 'leader' || rank === 'officer') return guildRosterModelCopy.ranks[rank];
  return guildRosterModelCopy.ranks.member;
}

/**
 * The rank word for a roster entry. A name only the addon saw has no rank on the site and the
 * addon's snapshot carries none, so it reads "Rank not known" rather than a plain "Member"; an
 * invited member joined as one, and an account's character carries the addon's rank.
 */
export function entryRankWord(entry: Pick<GuildRosterEntry, 'state' | 'rank'>): string {
  return entry.state === 'no_account' ? guildRosterModelCopy.rankNotKnown : rankWord(entry.rank);
}

/** True for a rank the site makes an officer of the guild (Officer or Leader). */
export function isOfficerRank(rank: GuildRank | string): boolean {
  return rank === 'officer' || rank === 'leader';
}

/**
 * The counts for a home that carries none (a response from before the roster model): the
 * account rows alone, verified or waiting. Nobody is "in game without an account" or "not
 * seen" then, because nothing says so.
 */
export function derivedCounts(roster: readonly GuildRosterRow[]): GuildRosterCounts {
  const waiting = roster.filter((row) => !row.verified).length;
  return { on_roster: roster.length, in_game_no_account: 0, waiting, not_seen: 0 };
}

/** The count line: "7 on the roster · 2 in game without an account · 1 waiting for approval ·
 *  1 not seen in game yet", the parts that are zero left out. */
export function countLine(counts: GuildRosterCounts): string {
  return guildRosterModelCopy.countLine(counts);
}

/** The entries with no roster row of their own: a name the addon saw with no account, and an
 *  invited member. They list after the account rows. */
export function entriesWithoutRow(members: readonly GuildRosterEntry[] | undefined): GuildRosterEntry[] {
  return (members ?? []).filter((entry) => entry.character_key === undefined);
}

export interface EntryFilters {
  role: 'all' | 'tank' | 'healer' | 'dps';
  classFilter: string;
  verifiedOnly: boolean;
  noAccountOnly: boolean;
}

/**
 * Which row-less entries survive the filter strip. "Verified only" drops the names with no
 * account (an invited member is verified); "No account" keeps nothing else; a role filter drops
 * them all, since an entry carries no role; a class filter keeps those of that class.
 */
export function filterEntries(
  entries: readonly GuildRosterEntry[],
  filters: EntryFilters,
): GuildRosterEntry[] {
  return entries.filter((entry) => {
    if (filters.role !== 'all') return false;
    if (filters.noAccountOnly && entry.state !== 'no_account') return false;
    if (filters.verifiedOnly && entry.state === 'no_account') return false;
    if (filters.classFilter !== 'all' && entry.class !== filters.classFilter) return false;
    return true;
  });
}

/** How many of the waiting rows are an officer (Officer or Leader): approving them hands out
 *  the site's officer controls. */
export function officersAmong(rows: readonly GuildRosterRow[]): number {
  return rows.filter((row) => isOfficerRank(row.rank)).length;
}

/**
 * What "Approve all" asks first, or null when no waiting row is an officer and it just goes:
 * "Approve 4, including 1 officer?".
 */
export function approveAllQuestion(rows: readonly GuildRosterRow[]): string | null {
  const officers = officersAmong(rows);
  return officers === 0 ? null : guildRosterModelCopy.approveAllQuestion(rows.length, officers);
}

export interface RosterModel {
  members?: GuildRosterEntry[];
  counts?: GuildRosterCounts;
}

/**
 * The members and counts once the characters in `keys` are approved: each waiting entry for one
 * becomes approved and the waiting count drops with it, so the count line and the list agree in
 * the same frame. A home without the roster model passes through untouched.
 */
export function afterApprove(model: RosterModel, keys: ReadonlySet<string>): RosterModel {
  if (model.members === undefined || model.counts === undefined) return model;
  let changed = 0;
  const members = model.members.map((entry) => {
    if (entry.state !== 'waiting' || entry.character_key === undefined || !keys.has(entry.character_key)) {
      return entry;
    }
    changed += 1;
    return { ...entry, state: 'approved' as const };
  });
  return { members, counts: { ...model.counts, waiting: Math.max(0, model.counts.waiting - changed) } };
}

/**
 * The members and counts once the character `key` is removed (or declined): its entry leaves
 * the list and the roster count, and the waiting count with it when it was waiting. The next
 * read of the home says whether the addon still names it (as "no account yet").
 */
export function afterRemove(model: RosterModel, key: string): RosterModel {
  if (model.members === undefined || model.counts === undefined) return model;
  const gone = model.members.find((entry) => entry.character_key === key);
  if (gone === undefined) return model;
  return {
    members: model.members.filter((entry) => entry !== gone),
    counts: {
      ...model.counts,
      on_roster: Math.max(0, model.counts.on_roster - 1),
      waiting: Math.max(0, model.counts.waiting - (gone.state === 'waiting' ? 1 : 0)),
    },
  };
}
