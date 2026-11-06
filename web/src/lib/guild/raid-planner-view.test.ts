// web/src/lib/guild/raid-planner-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRaidSignup } from './api';
import { plannedRaid } from '../../test-support/raid-fixtures';
import { pingAllowed } from './overview-role-view';
import {
  benchRows,
  classPlural,
  compDraft,
  defaultGrouping,
  groupingFrom,
  groupingOptions,
  groupingSearch,
  groupingSummary,
  notSignedRows,
  parseCompDraft,
  pingBlockedLine,
  plannerGroups,
  plannerStatusTag,
  rowSub,
  stateCountsLine,
  stateTag,
} from './raid-planner-view';

function signup(name: string, extra: Partial<GuildRaidSignup> = {}): GuildRaidSignup {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'mage',
    spec: 'Frost',
    spec_key: 'mage-frost',
    group: 'ranged',
    role: 'dps',
    state: 'in',
    note: '',
    own: false,
    ...extra,
  };
}

const SIGNUPS = [
  signup('Thornwicke', {
    class: 'warrior',
    spec: 'Protection',
    spec_key: 'warrior-protection',
    group: 'tank',
    role: 'tank',
  }),
  signup('Ironbrand', {
    class: 'paladin',
    spec: 'Holy',
    spec_key: 'paladin-holy',
    group: 'healer',
    role: 'healer',
  }),
  signup('Frostquill'),
  signup('Cinderglass', { state: 'bench', spec: 'Fire', spec_key: 'mage-fire' }),
  signup('Pinefeather', {
    class: 'hunter',
    spec: 'Marksmanship',
    spec_key: 'hunter-marksmanship',
    state: 'late',
  }),
  signup('Gone', { state: 'out' }),
];
const RAID = plannedRaid({ signups: SIGNUPS });
const EVENT = plannedRaid({ zone_kind: 'event', signups: SIGNUPS });

const names = (rows: { name: string }[]): string[] => rows.map((r) => r.name);

describe('grouping', () => {
  it('opens a raid on Role and an event on Status, which has no Role choice', () => {
    expect(defaultGrouping(RAID)).toBe('role');
    expect(defaultGrouping(EVENT)).toBe('status');
    expect(groupingOptions(RAID).map((o) => o.label)).toEqual(['Status', 'Role', 'Class']);
    expect(groupingOptions(EVENT).map((o) => o.value)).toEqual(['status', 'class']);
  });

  it('reads the URL, falling back on an unknown or unoffered value', () => {
    expect(groupingFrom('?group=class', RAID)).toBe('class');
    expect(groupingFrom('?group=nope', RAID)).toBe('role');
    expect(groupingFrom('?group=role', EVENT)).toBe('status');
    expect(groupingFrom('', RAID)).toBe('role');
  });

  it('writes the param, keeping others, and drops it for the default', () => {
    expect(groupingSearch('?x=1', 'class', RAID)).toBe('?x=1&group=class');
    expect(groupingSearch('?group=class', 'role', RAID)).toBe('');
    expect(groupingSummary('status')).toBe('Group by status');
  });
});

describe('plannerGroups', () => {
  it('Role: the four In groups always, then Late, Bench and Out when they have anyone', () => {
    const groups = plannerGroups(RAID, 'role');
    expect(groups.map((g) => g.id)).toEqual(['tank', 'healer', 'melee', 'ranged', 'late', 'bench', 'out']);
    expect(groups.find((g) => g.id === 'melee')).toMatchObject({ rows: [], empty: 'No melee yet.' });
    expect(names(groups.find((g) => g.id === 'ranged')!.rows)).toEqual(['Frostquill']);
  });

  it('Status: In, Late, Bench, Out, empty ones left out, role then class then name', () => {
    const groups = plannerGroups(RAID, 'status');
    expect(groups.map((g) => g.label)).toEqual(['In', 'Late', 'Bench', 'Out']);
    expect(names(groups[0].rows)).toEqual(['Thornwicke', 'Ironbrand', 'Frostquill']);
    expect(plannerGroups(plannedRaid({ signups: [SIGNUPS[0]] }), 'status').map((g) => g.id)).toEqual(['in']);
  });

  it('Class: the game order, each class In first then Late, Bench, Out', () => {
    const groups = plannerGroups(RAID, 'class');
    expect(groups.map((g) => g.label)).toEqual(['Warriors', 'Paladins', 'Hunters', 'Mages']);
    expect(names(groups[3].rows)).toEqual(['Frostquill', 'Cinderglass', 'Gone']);
  });

  it('an event asked for Role reads by Status', () => {
    expect(plannerGroups(EVENT, 'role').map((g) => g.id)).toEqual(['in', 'late', 'bench', 'out']);
  });
});

describe('row words', () => {
  it('reads spec, class and role on a raid, the class alone at an event', () => {
    const row = { class: 'paladin', spec_key: 'paladin-holy', role: 'healer' as const };
    expect(rowSub(row, true)).toBe('Holy Paladin · Healer');
    expect(rowSub(row, false)).toBe('Paladin');
    expect(classPlural('shaman')).toBe('Shamans');
  });

  it('tags each state in its tone, and the raid status', () => {
    expect(stateTag('in')).toEqual({ text: 'In', tone: 'success' });
    expect(stateTag('late').tone).toBe('warn');
    expect(stateTag('bench').tone).toBe('muted');
    expect(plannerStatusTag('cancelled')).toEqual({ text: 'Cancelled', tone: 'danger' });
    expect(plannerStatusTag('ended').text).toBe('Ended');
  });

  it('counts In, Late and Bench', () => {
    expect(stateCountsLine(plannedRaid().counts)).toBe('21 In · 2 Late · 1 Bench');
  });
});

describe('bench and not signed', () => {
  it('lists the benched sign-ups', () => {
    expect(names(benchRows(RAID))).toEqual(['Cinderglass']);
  });

  it('lists the raiders with no sign-up, class order, naming who has no Discord', () => {
    const raid = plannedRaid({
      not_signed: [
        {
          character_key: 'us/pvp/voidcall',
          name: 'Voidcall',
          class: 'priest',
          spec: 'Shadow',
          spec_key: 'priest-shadow',
          group: 'ranged',
        },
      ],
      not_signed_raiders: [
        {
          character_key: 'us/pvp/voidcall',
          name: 'Voidcall',
          class: 'priest',
          account_key: 'a',
          discord_linked: false,
        },
        {
          character_key: 'us/pvp/axe',
          name: 'Axe',
          class: 'warrior',
          account_key: 'b',
          discord_linked: true,
        },
      ],
    });
    const rows = notSignedRows(raid);
    expect(names(rows)).toEqual(['Axe', 'Voidcall']);
    expect(rows[1].sub).toBe('Shadow Priest · DPS · No Discord linked');
    expect(rows[0].sub).toBe('Warrior · DPS');
    expect(rows[1].roster).toMatchObject({
      state: null,
      role: 'dps',
      group: 'ranged',
      spec_key: 'priest-shadow',
    });
    expect(notSignedRows(plannedRaid())).toEqual([]);
  });
});

describe('ping', () => {
  const NOW = new Date('2026-12-08T18:00:00Z');
  const post = { posted_at: '2026-12-07T18:00:00Z', post_at: null, reminded_at: null, remind: false };
  const clock = (at: Date): string => at.toISOString().slice(11, 16);

  it('needs the raid post, then waits out the throttle', () => {
    expect(pingAllowed({ discord_post: null }, NOW)).toBe(false);
    expect(pingBlockedLine({ discord_post: null }, NOW, clock)).toBe(
      'Post the raid to Discord first, then ping under it.',
    );
    expect(pingAllowed({ discord_post: post }, NOW)).toBe(true);
    const throttled = {
      discord_post: post,
      ping: { last_pinged_at: null, next_allowed_at: '2026-12-08T20:10:00Z' },
    };
    expect(pingAllowed(throttled, NOW)).toBe(false);
    expect(pingBlockedLine(throttled, NOW, clock)).toBe('Pinged already. The next ping can go out at 20:10.');
    expect(pingBlockedLine({ discord_post: post }, NOW, clock)).toBeNull();
  });
});

describe('comp draft', () => {
  it('round-trips the seats and refuses anything but a whole number', () => {
    const draft = compDraft({ tanks: 2, healers: 8, melee: 10, ranged: 20 });
    expect(draft).toEqual({ tanks: '2', healers: '8', melee: '10', ranged: '20' });
    expect(parseCompDraft({ ...draft, melee: ' 9 ' })).toEqual({
      ok: true,
      target: { tanks: 2, healers: 8, melee: 9, ranged: 20 },
    });
    expect(parseCompDraft({ ...draft, tanks: '', ranged: '-1' })).toEqual({
      ok: false,
      fields: { tanks: 'Use a whole number, 0 or more.', ranged: 'Use a whole number, 0 or more.' },
    });
  });
});
