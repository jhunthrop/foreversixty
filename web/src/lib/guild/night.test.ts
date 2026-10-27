import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import type { RaidAttendanceRow, RaidNight, RaidReport } from './attendance-contract';
import {
  attendanceGroupOf,
  canExcuse,
  daysBeforeRaid,
  discordReportText,
  groupAttendance,
  visibleRows,
  yourNightFacts,
  yourNightRows,
  yourNightSentence,
} from './night';

const CHICAGO = 'America/Chicago';
const SPEC = new URL('../../../../design/specs/2026-10-27-guild-attendance.md', import.meta.url);
const HEADING = '### 6.1 On the page';

function row(extra: Partial<RaidAttendanceRow> = {}): RaidAttendanceRow {
  return {
    character_key: 'us/classic/pyrewisp',
    name: 'Pyrewisp',
    class: 'mage',
    spec: 'Fire',
    signed: 'in',
    role: 'dps',
    outcome: 'present',
    came_as: null,
    first_seen_offset_s: 0,
    pulls_present: 11,
    pulls_total: 11,
    excused: false,
    own: false,
    ...extra,
  };
}

describe('groupAttendance', () => {
  it('puts each outcome in its group, in the spec order, and drops empty groups', () => {
    const rows = [
      row({ name: 'P', outcome: 'present' }),
      row({ name: 'B', signed: 'bench', role: null, outcome: 'absent' }),
      row({ name: 'L', outcome: 'late', first_seen_offset_s: 1320 }),
      row({ name: 'J', signed: null, role: null, outcome: 'present' }),
      row({ name: 'N', outcome: 'no_show' }),
    ];
    const groups = groupAttendance(rows);
    expect(groups.map((g) => g.key)).toEqual(['noShow', 'late', 'joined', 'present', 'benchOut']);
    expect(groups.map((g) => g.label)).toEqual([
      'Signed up, not in the log',
      'Late',
      'Joined from the bench or without a sign-up',
      'Present',
      'Bench and out',
    ]);
    expect(groups.map((g) => g.compact)).toEqual([false, false, false, true, true]);
    expect(groupAttendance([row()]).map((g) => g.key)).toEqual(['present']);
  });

  it('never calls a bench or unsigned character late', () => {
    expect(attendanceGroupOf(row({ signed: 'bench', outcome: 'late' }))).toBe('joined');
    expect(attendanceGroupOf(row({ signed: null, outcome: 'late' }))).toBe('joined');
    expect(attendanceGroupOf(row({ signed: 'late', outcome: 'late' }))).toBe('late');
  });

  it('keeps a came-as row in Present', () => {
    expect(attendanceGroupOf(row({ came_as: 'Grimtotem' }))).toBe('present');
  });

  it('sorts a row with no known role last', () => {
    const [group] = groupAttendance([
      row({ name: 'Mystery', role: null, spec: 'Unknown', class: 'warrior' }),
      row({ name: 'Zed', class: 'mage', role: 'dps' }),
    ]);
    expect(group?.rows.map((r) => r.name)).toEqual(['Zed', 'Mystery']);
  });

  it('orders a group by role, class, then name', () => {
    const [group] = groupAttendance([
      row({ name: 'Zed', class: 'mage', role: 'dps' }),
      row({ name: 'Amy', class: 'rogue', role: 'dps' }),
      row({ name: 'Tank', class: 'warrior', role: 'tank' }),
      row({ name: 'Bob', class: 'rogue', role: 'dps' }),
    ]);
    expect(group?.rows.map((r) => r.name)).toEqual(['Tank', 'Amy', 'Bob', 'Zed']);
  });

  it('offers Excused on not-in-log and late rows only', () => {
    expect(canExcuse(row({ outcome: 'no_show' }))).toBe(true);
    expect(canExcuse(row({ outcome: 'late' }))).toBe(true);
    expect(canExcuse(row({ outcome: 'present' }))).toBe(false);
    expect(canExcuse(row({ outcome: 'absent' }))).toBe(false);
  });

  it('caps a list at ten until expanded', () => {
    const items = Array.from({ length: 14 }, (_, i) => i);
    expect(visibleRows(items, false)).toHaveLength(10);
    expect(visibleRows(items, true)).toHaveLength(14);
  });
});

describe('yourNightSentence', () => {
  const sentence = (index: number, fill: Record<string, string>): string =>
    specString(SPEC, HEADING, 'Your night sentences', index, fill);

  it('reads each case of spec 6.1, through the spec table', () => {
    expect(yourNightSentence(row())).toBe(sentence(0, { State: 'In', Role: 'DPS', n: '11', t: '11' }));
    expect(
      yourNightSentence(
        row({ outcome: 'late', signed: 'late', role: 'healer', first_seen_offset_s: 1320, pulls_present: 9 }),
      ),
    ).toBe(sentence(1, { State: 'Late', Role: 'Healer', m: '22', n: '9', t: '11' }));
    expect(yourNightSentence(row({ came_as: 'Grimtotem', pulls_present: 8 }))).toBe(
      sentence(2, { Name: 'Pyrewisp', Alt: 'Grimtotem', n: '8', t: '11' }),
    );
    expect(yourNightSentence(row({ outcome: 'no_show', pulls_present: 0 }))).toBe(
      sentence(3, { Role: 'DPS' }),
    );
    expect(yourNightSentence(row({ signed: null, role: null, pulls_present: 4 }))).toBe(
      sentence(4, { n: '4', t: '11' }),
    );
    expect(yourNightSentence(row({ signed: 'bench', role: null, pulls_present: 6 }))).toBe(
      sentence(5, { n: '6', t: '11' }),
    );
  });

  it('falls back to the spec role, and drops the role clause when none is known', () => {
    expect(yourNightSentence(row({ role: null, spec: 'Fury' }))).toContain(' as DPS.');
    expect(yourNightSentence(row({ role: null, spec: 'Unknown' }))).toMatch(/^Signed In\. Present/);
  });

  it('says when a raider signed Late came, though they count as present', () => {
    const came = row({ signed: 'late', outcome: 'present', first_seen_offset_s: 540, pulls_present: 11 });
    expect(yourNightSentence(came)).toBe(
      'Signed Late as DPS. First seen 9 min after the first pull, 11 of 11 pulls.',
    );
    const onTime = row({ signed: 'late', outcome: 'present', first_seen_offset_s: 0 });
    expect(yourNightSentence(onTime)).toContain('Present from the first pull');
  });

  it('reads a late row with no offset without the minutes', () => {
    const late = row({ outcome: 'late', signed: 'late', first_seen_offset_s: null, pulls_present: 3 });
    expect(yourNightSentence(late)).toBe('Signed Late as DPS, 3 of 11 pulls.');
  });

  it('never prints the word no-show', () => {
    const outcomes = ['present', 'late', 'absent', 'no_show'] as const;
    for (const outcome of outcomes) {
      expect(yourNightSentence(row({ outcome })).toLowerCase()).not.toContain('no-show');
    }
  });
});

describe('yourNightRows and facts', () => {
  it('lists own rows with the header character first', () => {
    const rows = [
      row({ character_key: 'a', own: true }),
      row({ character_key: 'b', own: false }),
      row({ character_key: 'c', own: true }),
    ];
    expect(yourNightRows(rows, 'c').map((r) => r.character_key)).toEqual(['c', 'a']);
    expect(yourNightRows(rows, null).map((r) => r.character_key)).toEqual(['a', 'c']);
  });

  const report: RaidReport = {
    kill_order: [],
    unnamed: null,
    deaths: [
      { character_key: 'us/classic/pyrewisp', name: 'Pyrewisp', class: 'mage', on_wipes: 1, on_kills: 1 },
    ],
    consumables: { log_has_buffs: false, rows: [] },
    loot: [],
    moves: {
      previous: { id: 3, title: 'Onyxia A', starts_at: '2026-12-04T02:00:00Z' },
      rows: [
        {
          character_key: 'us/classic/pyrewisp',
          name: 'Pyrewisp',
          class: 'mage',
          metric: 'dps',
          encounter: 'Onyxia',
          before: 238,
          after: 291,
        },
        {
          character_key: 'x',
          name: 'X',
          class: 'mage',
          metric: 'dps',
          encounter: 'Onyxia',
          before: 1,
          after: 2,
        },
      ],
    },
    posted_at: null,
  };

  it('gives deaths, the rise, and the excused flag', () => {
    expect(yourNightFacts(row({ excused: true }), report, CHICAGO)).toEqual([
      { kind: 'deaths', text: '2 deaths' },
      { kind: 'move', lead: 'DPS 291 on Onyxia, ', change: 'up 53', tail: ' since Thu 3 Dec', rising: true },
      { kind: 'excused', text: 'Excused' },
    ]);
  });

  it('keeps a fall muted and says down', () => {
    const fall = {
      ...report,
      moves: { ...report.moves, rows: [{ ...report.moves.rows[0]!, before: 300, after: 288 }] },
    };
    const [, move] = yourNightFacts(row(), fall, CHICAGO);
    expect(move).toMatchObject({ change: 'down 12', rising: false });
  });

  it('skips a flat move', () => {
    const flat = {
      ...report,
      moves: { ...report.moves, rows: [{ ...report.moves.rows[0]!, before: 291, after: 291 }] },
    };
    expect(yourNightFacts(row(), flat, CHICAGO).map((f) => f.kind)).toEqual(['deaths']);
  });

  it('gives 0 deaths for a character in the log, even without a report', () => {
    expect(yourNightFacts(row(), null, CHICAGO)).toEqual([{ kind: 'deaths', text: '0 deaths' }]);
  });

  it('gives a character not in the log nothing from the log, only Excused when set', () => {
    expect(yourNightFacts(row({ outcome: 'no_show' }), report, CHICAGO)).toEqual([]);
    expect(yourNightFacts(row({ outcome: 'no_show', excused: true }), report, CHICAGO)).toEqual([
      { kind: 'excused', text: 'Excused' },
    ]);
  });
});

describe('discordReportText', () => {
  const night: RaidNight = {
    signed: 23,
    came: 22,
    late: 1,
    not_in_log: 1,
    joined: 0,
    excused: 0,
    kills: 1,
    wipes: 3,
    deaths: 9,
    duration_ms: 161 * 60_000,
  };
  const award = (item: string, name: string) => ({
    item_id: 1,
    item_name: item,
    quality: 4,
    icon: '',
    character_key: name,
    name,
    class: 'mage',
    by_name: 'Boss',
    at: '',
  });
  const full: RaidReport = {
    kill_order: [
      {
        encounter_id: 1,
        name: 'Onyxia',
        fight_index: 4,
        killed: true,
        kill_ms: 102_000,
        wipes_before: 3,
        deaths: 2,
      },
      {
        encounter_id: 2,
        name: 'Nefarian',
        fight_index: 9,
        killed: false,
        kill_ms: null,
        wipes_before: 2,
        deaths: 0,
      },
    ],
    unnamed: null,
    deaths: [
      { character_key: 'a', name: 'Vexlash', class: 'rogue', on_wipes: 1, on_kills: 1 },
      { character_key: 'b', name: 'Pyrewisp', class: 'mage', on_wipes: 3, on_kills: 1 },
      { character_key: 'c', name: 'Windtalon', class: 'hunter', on_wipes: 2, on_kills: 1 },
      { character_key: 'd', name: 'Quiet', class: 'hunter', on_wipes: 0, on_kills: 0 },
      { character_key: 'e', name: 'Fifth', class: 'hunter', on_wipes: 1, on_kills: 0 },
    ],
    consumables: { log_has_buffs: false, rows: [] },
    loot: [
      {
        encounter_id: 1,
        boss: 'Onyxia',
        awards: [award('Drake Fang Talisman', 'Pyrewisp'), award('Head of Onyxia', 'Vexlash')],
      },
    ],
    moves: {
      previous: { id: 3, title: 'Onyxia A', starts_at: '2026-12-04T02:00:00Z' },
      rows: [
        {
          character_key: 'a',
          name: 'Pyrewisp',
          class: 'mage',
          metric: 'dps',
          encounter: 'Onyxia',
          before: 238,
          after: 291,
        },
        {
          character_key: 'b',
          name: 'Lightbrand',
          class: 'paladin',
          metric: 'hps',
          encounter: 'Onyxia',
          before: 100,
          after: 141,
        },
        {
          character_key: 'c',
          name: 'Faller',
          class: 'mage',
          metric: 'dps',
          encounter: 'Onyxia',
          before: 300,
          after: 250,
        },
      ],
    },
    posted_at: null,
  };
  const input = {
    title: "Onyxia's Lair",
    startsAt: '2026-12-11T02:00:00Z',
    guildName: 'Olympus XXVII',
    tz: CHICAGO,
    night,
    report: full,
    reportUrl: 'https://foreversixty.gg/guild/olympus/raids/12#report',
  };

  it('writes the spec lines in order, rises only, no absentees', () => {
    expect(discordReportText(input).split('\n')).toEqual([
      "Onyxia's Lair · Thu 10 Dec · Olympus XXVII",
      '1 kill, 3 wipes, 2 h 41 min. 22 of 23 came.',
      'Kills: Onyxia 1:42 (after 3 wipes)',
      'Most deaths: Pyrewisp 4, Windtalon 3, Vexlash 2',
      'Loot: Onyxia: Drake Fang Talisman to Pyrewisp; Head of Onyxia to Vexlash',
      'Biggest changes: Pyrewisp DPS +53, Lightbrand HPS +41',
      'Report: https://foreversixty.gg/guild/olympus/raids/12#report',
    ]);
  });

  it('leaves out lines with nothing to say', () => {
    const bare = { ...full, kill_order: [], deaths: [], loot: [], moves: { previous: null, rows: [] } };
    expect(discordReportText({ ...input, report: bare }).split('\n')).toHaveLength(3);
  });

  it('never passes 12 lines with a full report: loot gives way to a +N more line', () => {
    const bosses = Array.from({ length: 20 }, (_, i) => ({
      encounter_id: i,
      boss: `Boss ${i}`,
      awards: [award('Item', 'Pyrewisp')],
    }));
    const lines = discordReportText({ ...input, report: { ...full, loot: bosses } }).split('\n');
    expect(lines).toHaveLength(12);
    expect(lines[2]).toMatch(/^Kills:/);
    expect(lines[3]).toMatch(/^Most deaths:/);
    expect(lines.filter((l) => l.startsWith('Loot:'))).toHaveLength(5);
    expect(lines[9]).toBe('+15 more');
    expect(lines[10]).toMatch(/^Biggest changes:/);
    expect(lines[11]).toMatch(/^Report: /);
  });

  it('writes the labels the spec shows', () => {
    const spec = readFileSync(SPEC, 'utf8');
    for (const label of ['Kills:', 'Most deaths:', 'Loot:', 'Biggest changes:', 'Report:']) {
      expect(spec).toContain(`\`${label}`);
    }
  });
});

describe('daysBeforeRaid', () => {
  it('counts whole days and never goes negative', () => {
    expect(daysBeforeRaid('2026-12-01T18:00:00Z', '2026-12-10T02:00:00Z')).toBe(8);
    expect(daysBeforeRaid('2026-12-11T00:00:00Z', '2026-12-10T02:00:00Z')).toBe(0);
  });
});
