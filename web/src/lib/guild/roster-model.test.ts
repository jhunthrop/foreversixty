// web/src/lib/guild/roster-model.test.ts
// The roster as everyone the addon saw: the count line, the row-less entries and how the
// filter strip treats them, the capability test and the approve-all question.
import { describe, expect, it } from 'vitest';
import type { GuildRosterCounts, GuildRosterEntry, GuildRosterRow } from './api';
import {
  afterApprove,
  afterRemove,
  approveAllQuestion,
  can,
  countLine,
  derivedCounts,
  entriesWithoutRow,
  filterEntries,
  officersAmong,
  rankWord,
  stateWord,
} from './roster-model';

function row(name: string, rank: GuildRosterRow['rank'], verified: boolean): GuildRosterRow {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'pvp',
    name,
    rank,
    verified,
    logged_recently: false,
    consent: 'gear',
    may_remove: true,
  };
}

const NO_ACCOUNT: GuildRosterEntry = {
  name: 'Strayhound',
  state: 'no_account',
  rank: 'member',
  class: 'warrior',
};
const NOT_SEEN: GuildRosterEntry = { name: 'Driftwood', state: 'not_seen', rank: 'member', class: 'rogue' };
const APPROVED: GuildRosterEntry = {
  name: 'Bow',
  character_key: 'us/pvp/bow',
  state: 'approved',
  rank: 'leader',
};
const WAITING: GuildRosterEntry = {
  name: 'Ashenveil',
  character_key: 'us/pvp/ashenveil',
  state: 'waiting',
  rank: 'officer',
};

describe('countLine', () => {
  const counts = (over: Partial<GuildRosterCounts>): GuildRosterCounts => ({
    on_roster: 7,
    in_game_no_account: 0,
    waiting: 0,
    not_seen: 0,
    ...over,
  });
  it.each([
    [
      counts({ in_game_no_account: 2, waiting: 1, not_seen: 1 }),
      '7 on the roster · 2 in game without an account · 1 waiting for approval · 1 not seen in game yet',
    ],
    [counts({}), '7 on the roster'],
    [counts({ waiting: 4 }), '7 on the roster · 4 waiting for approval'],
    [
      counts({ in_game_no_account: 2, not_seen: 1 }),
      '7 on the roster · 2 in game without an account · 1 not seen in game yet',
    ],
    [counts({ on_roster: 0 }), '0 on the roster'],
  ])('%j reads %s', (input, want) => {
    expect(countLine(input)).toBe(want);
  });
});

describe('derivedCounts', () => {
  it('counts the account rows alone: nobody in game without an account or unseen', () => {
    expect(
      derivedCounts([row('A', 'leader', true), row('B', 'member', false), row('C', 'member', true)]),
    ).toEqual({ on_roster: 3, in_game_no_account: 0, waiting: 1, not_seen: 0 });
  });
});

describe('words', () => {
  it('names every state and every rank, Member for a plain member', () => {
    expect(stateWord('approved')).toBe('Approved');
    expect(stateWord('waiting')).toBe('Waiting for approval');
    expect(stateWord('no_account')).toBe('No account yet');
    expect(stateWord('not_seen')).toBe('Not seen in game yet');
    expect(rankWord('leader')).toBe('Leader');
    expect(rankWord('officer')).toBe('Officer');
    expect(rankWord('member')).toBe('Member');
    expect(rankWord(undefined)).toBe('Member');
  });
});

describe('can', () => {
  it('is true exactly when the word is in the list', () => {
    expect(can(['approve', 'remove'], 'approve')).toBe(true);
    expect(can(['approve', 'remove'], 'billing')).toBe(false);
    expect(can([], 'approve')).toBe(false);
    expect(can(undefined, 'approve')).toBe(false);
  });
});

describe('entriesWithoutRow and filterEntries', () => {
  const all = [APPROVED, WAITING, NO_ACCOUNT, NOT_SEEN];
  const open = { role: 'all', classFilter: 'all', verifiedOnly: false, noAccountOnly: false } as const;
  it('lists only the entries with no character row', () => {
    expect(entriesWithoutRow(all)).toEqual([NO_ACCOUNT, NOT_SEEN]);
    expect(entriesWithoutRow(undefined)).toEqual([]);
  });
  it('keeps both with no filter', () => {
    expect(filterEntries([NO_ACCOUNT, NOT_SEEN], open)).toEqual([NO_ACCOUNT, NOT_SEEN]);
  });
  it('"Verified only" drops the names with no account but keeps an invited member', () => {
    expect(filterEntries([NO_ACCOUNT, NOT_SEEN], { ...open, verifiedOnly: true })).toEqual([NOT_SEEN]);
  });
  it('"No account" keeps those alone', () => {
    expect(filterEntries([NO_ACCOUNT, NOT_SEEN], { ...open, noAccountOnly: true })).toEqual([NO_ACCOUNT]);
  });
  it('a role filter drops them all (an entry has no role); a class filter keeps its class', () => {
    expect(filterEntries([NO_ACCOUNT, NOT_SEEN], { ...open, role: 'dps' })).toEqual([]);
    expect(filterEntries([NO_ACCOUNT, NOT_SEEN], { ...open, classFilter: 'rogue' })).toEqual([NOT_SEEN]);
  });
});

describe('approveAllQuestion', () => {
  it('asks first when an officer or the leader is waiting', () => {
    const rows = [
      row('A', 'member', false),
      row('B', 'officer', false),
      row('C', 'member', false),
      row('D', 'member', false),
    ];
    expect(officersAmong(rows)).toBe(1);
    expect(approveAllQuestion(rows)).toBe('Approve 4, including 1 officer?');
    expect(approveAllQuestion([...rows, row('E', 'leader', false)])).toBe('Approve 5, including 2 officers?');
  });
  it('does not ask when nobody waiting is an officer', () => {
    expect(approveAllQuestion([row('A', 'member', false), row('B', 'member', false)])).toBeNull();
    expect(approveAllQuestion([])).toBeNull();
  });
});

describe('afterApprove and afterRemove keep the list and the counts in step', () => {
  const model = {
    members: [
      APPROVED,
      WAITING,
      { ...WAITING, name: 'Cindergale', character_key: 'us/pvp/cindergale', rank: 'member' as const },
      NO_ACCOUNT,
    ],
    counts: { on_roster: 4, in_game_no_account: 1, waiting: 2, not_seen: 0 },
  };
  it('approving turns the waiting entries into approved and drops the waiting count with them', () => {
    const next = afterApprove(model, new Set(['us/pvp/ashenveil']));
    expect(next.members?.map((m) => m.state)).toEqual(['approved', 'approved', 'waiting', 'no_account']);
    expect(next.counts).toEqual({ on_roster: 4, in_game_no_account: 1, waiting: 1, not_seen: 0 });
    const all = afterApprove(model, new Set(['us/pvp/ashenveil', 'us/pvp/cindergale']));
    expect(all.counts?.waiting).toBe(0);
  });
  it('removing takes the entry off the list and the roster count, and the waiting count if it waited', () => {
    const next = afterRemove(model, 'us/pvp/cindergale');
    expect(next.members?.map((m) => m.name)).toEqual(['Bow', 'Ashenveil', 'Strayhound']);
    expect(next.counts).toEqual({ on_roster: 3, in_game_no_account: 1, waiting: 1, not_seen: 0 });
    expect(afterRemove(model, 'us/pvp/bow').counts?.waiting).toBe(2);
  });
  it('a home without the roster model passes through untouched', () => {
    expect(afterApprove({}, new Set(['x']))).toEqual({});
    expect(afterRemove({}, 'x')).toEqual({});
  });
});
