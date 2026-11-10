// web/src/lib/guild/dense/dense.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRaidView, GuildReadinessRow, GuildRosterRow } from '../api';
import type { RaidReportParse } from '../attendance-contract';
import type { LootNights } from '../loot-council-api';
import type { GuildRecentAttendance } from '../overview-api';
import { fairnessRows, guildItemsPerNight } from './fairness';
import { filterBoard, median, rosterBoard, seenText, sortBoard } from './roster-board';
import {
  compGaps,
  needsAWord,
  ratingMovers,
  recruitGaps,
  seatsOpen,
  shortReasons,
  slipping,
  weakestPart,
} from './week-board';

const NOW = new Date('2027-01-09T22:00:00Z');

function rosterRow(name: string, extra: Partial<GuildRosterRow> = {}): GuildRosterRow {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'pvp',
    name,
    class: 'mage',
    spec: 'Frost',
    role: 'dps',
    group: 'ranged',
    rank: 'member',
    verified: true,
    logged_recently: true,
    logged_at: '2027-01-09T19:30:00Z',
    item_level: 64,
    consent: 'gear_bags',
    account_key: `acct-${name.toLowerCase()}`,
    may_remove: false,
    ...extra,
  };
}

const grid: GuildRecentAttendance = {
  raids: [40, 41, 42, 43, 44].map((id) => ({
    id,
    starts_at: `2026-12-${10 + id - 40}T02:00:00Z`,
    zone_name: 'Onyxia',
  })),
  rows: [
    {
      name: 'Ana',
      class: 'mage',
      account_key: 'acct-ana',
      cells: ['present', 'present', 'present', 'present', 'present'],
      slipping: false,
    },
    {
      name: 'Bo',
      class: 'mage',
      account_key: 'acct-bo',
      cells: ['present', 'absent', 'absent', 'present', 'absent'],
      slipping: true,
    },
    {
      name: 'Cy',
      class: 'mage',
      account_key: 'acct-cy',
      cells: ['bench', 'present', 'bench', 'present', 'bench'],
      slipping: false,
    },
    {
      name: 'Di',
      class: 'mage',
      account_key: 'acct-di',
      cells: ['none', 'none', 'present', 'present', 'present'],
      slipping: false,
    },
  ],
};

function decision(id: number, name: string, item: string, at: string) {
  return {
    award_id: id,
    item_id: 100 + id,
    item_name: item,
    icon: '',
    quality: 4,
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'mage',
    ms_os: 'ms' as const,
    reason: null,
    votes_yes: 3,
    votes_no: 0,
    votes_officers: 3,
    awarded_by: 'Dunmaro',
    at,
  };
}

const loot: LootNights = {
  officers: 3,
  nights: [
    {
      night: 'raid:44',
      raid_id: 44,
      report_id: 'r44',
      date: '2027-01-06',
      at: '2027-01-07T02:00:00Z',
      zone: 'Onyxia',
      current: true,
      drops: [],
      decisions: [decision(2, 'Ana', 'Sapphiron Drape', '2027-01-07T03:00:00Z')],
    },
    {
      night: 'raid:41',
      raid_id: 41,
      report_id: 'r41',
      date: '2026-12-16',
      at: '2026-12-17T02:00:00Z',
      zone: 'Onyxia',
      current: false,
      drops: [],
      decisions: [decision(1, 'Cy', 'Ring of Binding', '2026-12-17T03:00:00Z')],
    },
  ],
};

describe('rosterBoard', () => {
  const roster = [
    rosterRow('Ana', {
      rating: {
        overall: 80,
        output: 70,
        survival: 60,
        mechanics: 90,
        utility: 85,
        preparation: 75,
        activity: 88,
      },
      item_level: 66,
    }),
    rosterRow('Bo', { group: 'healer', role: 'healer', item_level: 60, logged_at: '2026-12-30T19:30:00Z' }),
    rosterRow('Cy', {
      rating: { overall: 90, output: 1, survival: 1, mechanics: 1, utility: 1, preparation: 1, activity: 1 },
    }),
  ];
  const readiness = [
    {
      character_key: 'us/pvp/ana',
      failing: 2,
      no_export: false,
      gear_gap: { upgrades: 2, gain_dps: 12.5, not_sim_checked: 0, unit: 'dps' },
      enchants: { missing_slots: ['hands'], checked: true },
      consumables: { state: 'short' },
      flask_food: { flask_or_elixir: false, food: true },
      talent_points_unspent: 0,
    },
  ] as unknown as GuildReadinessRow[];
  const groups = rosterBoard({
    roster,
    attendance: grid,
    loot,
    readiness,
    medianItemLevel: 64,
    target: { tanks: 3, healers: 9, melee: 12, ranged: 16, source: 'raid' },
    now: NOW,
  });

  it('groups by the raid group in role order, best rating first, with the seats', () => {
    expect(groups.map((g) => g.group)).toEqual(['healer', 'ranged']);
    expect(groups[1].rows.map((r) => r.name)).toEqual(['Cy', 'Ana']);
    expect(groups[1].need).toBe(16);
  });

  it('joins the attendance cells, items this tier, readiness and the median', () => {
    const ana = groups[1].rows[1];
    expect(ana.cells).toHaveLength(5);
    expect(ana.items).toBe(1);
    expect(ana.short).toBe(2);
    expect(ana.gap).toBe(12.5);
    expect(ana.itemLevelDelta).toBe(2);
    expect(groups[0].rows[0].seenDays).toBe(10);
  });

  it('sorts a missing figure last and filters by group and name', () => {
    const rows = sortBoard(
      groups.flatMap((g) => g.rows),
      'rating',
    );
    expect(rows.at(-1)?.name).toBe('Bo');
    expect(filterBoard(groups, 'healer', '').flatMap((g) => g.rows.map((r) => r.name))).toEqual(['Bo']);
    expect(filterBoard(groups, 'all', 'an').flatMap((g) => g.rows.map((r) => r.name))).toEqual(['Ana']);
  });

  it('reads the small figures', () => {
    expect(median([3, 1, 2, 10])).toBe(2.5);
    expect(median([])).toBeNull();
    expect(seenText(1)).toBe('');
    expect(seenText(9)).toBe('9 d');
  });
});

describe('fairness', () => {
  const roster = ['Ana', 'Bo', 'Cy', 'Di'].map((n) => rosterRow(n));

  it('ranks by nights came since the last item, then fewest items', () => {
    const rows = fairnessRows(roster, loot, grid);
    expect(rows.map((r) => [r.raiders.map((x) => x.name).join(','), r.since, r.items])).toEqual([
      ['Di', 3, 0],
      ['Bo', 2, 0],
      ['Cy', 1, 1],
      ['Ana', 0, 1],
    ]);
    expect(rows[2].raiders[0].last?.item).toBe('Ring of Binding');
  });

  it('puts raiders level on both figures in one row from three', () => {
    const many = ['Ana', 'Bo', 'Cy'].map((n) => rosterRow(n));
    const flat = {
      ...grid,
      rows: grid.rows
        .slice(0, 3)
        .map((r) => ({ ...r, cells: ['present', 'present', 'present', 'present', 'present'] as const })),
    };
    const rows = fairnessRows(
      many,
      { officers: 3, nights: [] },
      { ...flat, rows: flat.rows.map((r) => ({ ...r, cells: [...r.cells] })) },
    );
    expect(rows).toHaveLength(1);
    expect(rows[0].raiders).toHaveLength(3);
  });

  it('counts the guild items per night attended', () => {
    expect(guildItemsPerNight(roster, loot, grid)).toBeCloseTo(2 / 12);
  });
});

function view(extra: Partial<GuildRaidView>): GuildRaidView {
  return {
    signups: [],
    not_signed: [],
    comp_target: { tanks: 2, healers: 3, melee: 1, ranged: 1, source: 'raid' },
    ...extra,
  } as unknown as GuildRaidView;
}

describe('compGaps', () => {
  const v = view({
    signups: [
      {
        character_key: 'a',
        name: 'Tanky',
        class: 'warrior',
        spec: 'Protection',
        role: 'tank',
        state: 'in',
        note: '',
        own: false,
        group: 'tank',
      },
      {
        character_key: 'b',
        name: 'Heals',
        class: 'priest',
        spec: 'Holy',
        role: 'healer',
        state: 'late',
        note: '',
        own: false,
        group: 'healer',
      },
      {
        character_key: 'c',
        name: 'Benchy',
        class: 'priest',
        spec: 'Holy',
        role: 'healer',
        state: 'bench',
        note: '',
        own: false,
        group: 'healer',
      },
      {
        character_key: 'd',
        name: 'Stab',
        class: 'rogue',
        spec: 'Combat',
        role: 'dps',
        state: 'in',
        note: '',
        own: false,
        group: 'melee',
      },
    ],
    not_signed: [{ character_key: 'e', name: 'Quiet', class: 'warrior', spec: 'Protection', group: 'tank' }],
  });
  const rows = compGaps(v);

  it('counts Late as coming and names not-answered then benched fillers', () => {
    expect(rows.map((r) => [r.group, r.in, r.late, r.gap])).toEqual([
      ['tank', 1, 0, 1],
      ['healer', 0, 1, 2],
      ['melee', 1, 0, 0],
      ['ranged', 0, 0, 1],
    ]);
    expect(rows[0].fillers).toEqual([{ key: 'e', name: 'Quiet', cls: 'warrior', kind: 'ping' }]);
    expect(rows[1].fillers.map((f) => f.kind)).toEqual(['bench']);
  });

  it('says how many seats stay open if everyone answers', () => {
    expect(rows.map((r) => r.stillShort)).toEqual([0, 1, 0, 1]);
    expect(seatsOpen(rows)).toBe(4);
    expect(recruitGaps(rows)).toEqual([
      { group: 'healer', short: 1 },
      { group: 'ranged', short: 1 },
    ]);
  });

  it('has no rows without a target', () => {
    expect(compGaps(view({ comp_target: null }))).toEqual([]);
  });
});

function parse(name: string, rating: number | null, value = 600): RaidReportParse {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'mage',
    role: 'dps',
    encounter_id: 1084,
    encounter: 'Onyxia',
    fight_index: 3,
    metric: 'dps',
    value,
    rating,
  };
}

describe('ratingMovers and needsAWord', () => {
  const roster = [
    rosterRow('Ana', {
      rating: {
        overall: 72,
        output: 80,
        survival: 40,
        mechanics: 70,
        utility: 60,
        preparation: 75,
        activity: 90,
      },
    }),
    rosterRow('Bo'),
    rosterRow('Cy'),
  ];
  const movers = ratingMovers(
    [parse('Ana', 72), parse('Bo', 90), parse('Cy', 60), parse('Di', 50)],
    [parse('Ana', 91), parse('Bo', 70), parse('Cy', null)],
    roster,
  );

  it('lists risers and fallers against the night before, with the weakest rating part', () => {
    expect(movers.up.map((m) => [m.name, m.change])).toEqual([['Bo', 20]]);
    expect(movers.down.map((m) => [m.name, m.change])).toEqual([['Ana', -19]]);
    expect(movers.down[0].weakest).toEqual({ part: 'survival', value: 40 });
    expect(weakestPart(null)).toBeNull();
  });

  it('names who to talk to, one action each', () => {
    const readiness = [
      {
        character_key: 'us/pvp/cy',
        name: 'Cy',
        class: 'mage',
        attuned: false,
        no_export: false,
        consumables: { state: 'short' },
        enchants: { missing_slots: ['hands'], checked: true },
        flask_food: { flask_or_elixir: false, food: true },
        talent_points_unspent: 0,
        talent_points_read: true,
      },
    ] as unknown as GuildReadinessRow[];
    const words = needsAWord({
      view: view({
        not_signed_raiders: [
          {
            character_key: 'us/pvp/bo',
            name: 'Bo',
            class: 'mage',
            account_key: 'acct-bo',
            discord_linked: false,
          },
        ],
      }),
      roster: roster.map((r) => (r.name === 'Bo' ? { ...r, logged_at: '2026-12-31T19:30:00Z' } : r)),
      grid,
      readiness,
      movers,
      now: NOW,
    });
    expect(words.map((w) => [w.name, w.action, w.reasons.join(', ')])).toEqual([
      ['Bo', 'profile', 'Not answered, missed 3 of the last 4, last log 9 d ago, no Discord'],
      ['Cy', 'nudge', '1 enchant missing, no flask or elixir, not attuned'],
      ['Ana', 'report', 'Rating 91 to 72 · survival 40'],
    ]);
    expect(shortReasons(readiness[0])).toHaveLength(3);
  });

  it('lists the slipping, never for a bench', () => {
    expect(slipping(grid).map((r) => [r.name, r.present, r.nights])).toEqual([['Bo', 2, 5]]);
    expect(slipping(null)).toEqual([]);
  });
});
