// web/src/lib/guild/guild-tabs-dense.test.ts
// The pure figures behind the dense guild tabs (h-guild-tabs): the Raids tab's nights, a boss's
// Progression nights, the planner's per-raider figures, an applicant's seat and wait, and the
// Settings tab's standing facts.
import { describe, expect, it } from 'vitest';
import type { GuildProgressionEncounter, GuildRaidRow, GuildReadinessRow, GuildRosterRow } from './api';
import type { CompGapRow } from './dense/week-board';
import { seatFill, waitedWords } from './applicant-figures';
import { medianItemLevel, plannerFacts, signed } from './planner-facts';
import { fasterSinceFirst, progressionNights } from './progression-nights';
import {
  absentByRaid,
  awardsByReport,
  bossTimeMs,
  deltaWords,
  raidNightFigures,
  raidNightsMeta,
} from './raid-nights-table';
import { settingsFacts, stepsLeft } from './settings-facts';
import { setupFixture } from '../../test-support/guild-setup-fixture';
import type { ApplicantCharacter } from './recruit-api';
import { clockDuration } from '../format/duration';

const ONYXIA = 1084;

function night(id: string, killMs: number, wipes = 0): GuildRaidRow {
  return {
    id,
    title: "Onyxia's Lair",
    zone: "Onyxia's Lair",
    created_at: '2027-01-07T02:00:00Z',
    duration_ms: 3_000_000,
    fight_count: 2 + wipes,
    kill_count: 2,
    wipe_count: wipes,
    raiders: 36,
    deaths: 2,
    fights: [
      {
        index: 0,
        name: 'Onyxian Whelp',
        encounter_id: null,
        kill: true,
        duration_ms: 60_000,
        deaths: 0,
        players: 36,
      },
      ...Array.from({ length: wipes }, (_, i) => ({
        index: 1 + i,
        name: 'Onyxia',
        encounter_id: ONYXIA,
        kill: false,
        duration_ms: 200_000,
        deaths: 5,
        players: 36,
      })),
      {
        index: 1 + wipes,
        name: 'Onyxia',
        encounter_id: ONYXIA,
        kill: true,
        duration_ms: killMs,
        deaths: 2,
        players: 36,
      },
    ],
    present: [],
  };
}

describe('the Raids tab nights', () => {
  const bosses = new Set([ONYXIA]);
  const rows = [night('b', 192_000), night('a', 198_000, 1)];

  it('times the bosses only and compares each night with the one before it', () => {
    expect(bossTimeMs(rows[0], bosses)).toBe(192_000);
    const [last, first] = raidNightFigures(rows, bosses);
    expect(last.deltaS).toBe(-6);
    expect(last.best).toBe(true);
    expect(first.deltaS).toBeNull();
    expect(first.wipes).toBe(1);
  });

  it('writes the clock, the change and the head line', () => {
    expect(clockDuration(192_000)).toBe('3:12');
    expect(deltaWords(-6)).toEqual({ value: '−6 s', tone: 'kill' });
    expect(deltaWords(4)).toEqual({ value: '+4 s', tone: 'wipe' });
    expect(raidNightsMeta(rows, bosses)).toBe('2 nights · 2 kills · best 3:12');
  });

  it('names who was absent and benched each guild raid', () => {
    const misses = absentByRaid({
      raids: [{ id: 44 }],
      rows: [
        { name: 'Voidcall', class: 'warlock', cells: ['absent'] },
        { name: 'Nightlatch', class: 'rogue', cells: ['bench'] },
        { name: 'Hollowmend', class: 'priest', cells: ['present'] },
      ],
    });
    expect(misses.get(44)).toEqual({
      absent: [{ name: 'Voidcall', class: 'warlock' }],
      bench: [{ name: 'Nightlatch', class: 'rogue' }],
    });
  });

  it('counts the council awards per report', () => {
    const map = awardsByReport([
      { report_id: 'a', decisions: [1, 2] },
      { report_id: null, decisions: [1] },
    ]);
    expect(map.get('a')).toBe(2);
    expect(map.size).toBe(1);
  });
});

describe('a boss on Progression', () => {
  const encounter = {
    encounter_id: ONYXIA,
    name: 'Onyxia',
    zone: "Onyxia's Lair",
    pulls: 3,
    kills: 2,
    deaths_per_pull: 3,
    best_by_role: {},
    pulls_by_night: [
      { report_id: 'a', date: '2026-12-09', pulls: 2, killed: true, kill_ms: 225_000, deaths: 12 },
      { report_id: 'b', date: '2026-12-16', pulls: 1, killed: true, kill_ms: 214_000, deaths: 4 },
    ],
  } satisfies GuildProgressionEncounter;

  it('lists the nights newest first with the change and the wipes', () => {
    const [newest, oldest] = progressionNights(encounter);
    expect(newest).toMatchObject({ reportId: 'b', deltaS: -11, wipes: 0, best: true, first: false });
    expect(oldest).toMatchObject({ reportId: 'a', deltaS: null, wipes: 1, first: true });
    expect(fasterSinceFirst(encounter)).toBe(11);
  });
});

function readinessRow(key: string): GuildReadinessRow {
  return {
    character_key: key,
    no_export: false,
    gear_gap: { upgrades: 0, gain_dps: null, not_sim_checked: 0, unit: null },
    enchants: { missing_slots: [], checked: true },
    flask_food: { flask_or_elixir: false, food: false },
    consumables: { state: 'short' },
    talent_points_unspent: 0,
    attuned: true,
  } as unknown as GuildReadinessRow;
}

describe('the planner figures', () => {
  const row = (key: string, ilvl: number): GuildRosterRow =>
    ({
      character_key: key,
      item_level: ilvl,
      attendance_30: { present: 4, nights: 5 },
      rating: { overall: 77 },
    }) as unknown as GuildRosterRow;

  it('reads item level against the median, the readiness shortfalls and attendance', () => {
    const roster = [row('a', 60), row('b', 62), row('c', 64)];
    expect(medianItemLevel(roster)).toBe(62);
    const reads = { readiness: [readinessRow('c')], attendance: null, checklist: null };
    const facts = plannerFacts(roster, reads);
    expect(facts.get('c')).toMatchObject({
      itemLevel: 64,
      itemLevelDelta: 2,
      short: 2,
      attendancePct: 80,
      rating: 77,
    });
    expect(facts.get('c')?.why).toEqual(['no flask or elixir', 'no food']);
    expect(facts.get('a')?.short).toBeNull();
    expect(plannerFacts(roster, { ...reads, readiness: null }).get('c')?.short).toBeNull();
    expect(signed(-1)).toBe('−1');
    expect(signed(0)).toBe('+0');
  });
});

describe('an applicant', () => {
  const gaps = [
    { group: 'healer', need: 9, in: 6, late: 1, gap: 2, fillers: [], stillShort: 2 },
    { group: 'melee', need: 12, in: 12, late: 0, gap: 0, fillers: [], stillShort: 0 },
  ] satisfies CompGapRow[];
  const who = (group: ApplicantCharacter['group']) => ({ group }) as ApplicantCharacter;

  it('fills the open seat of their group, or none', () => {
    expect(seatFill(who('healer'), gaps)).toEqual({ text: 'Healers, 2 short', short: true });
    expect(seatFill(who('melee'), gaps)).toEqual({ text: 'Melee, full', short: false });
    expect(seatFill(who(null), gaps)).toBeNull();
    expect(seatFill(who('healer'), null)).toBeNull();
  });

  it('says how long the application has waited', () => {
    const now = new Date('2027-01-09T22:00:00Z');
    expect(waitedWords('2027-01-07T18:00:00Z', now)).toBe('2 d');
    expect(waitedWords('2027-01-09T20:00:00Z', now)).toBe('today');
  });
});

describe('the Settings facts', () => {
  it('names the officers, the invite and who has not synced, and the steps left', () => {
    const setup = setupFixture({ invite_link: 'done', officer_ranks: 'done' });
    const facts = settingsFacts(setup);
    expect(facts.map((f) => f.title)).toEqual([
      '3 officers',
      '14 joined by the invite link',
      '9 of 24 synced in 14 days',
    ]);
    expect(facts[0].sub).toBe('Dunmaro (GM), Kraggor, Mirelle');
    expect(facts[2].sub).toBe('Not synced: Ashvane, Brightwell, Corvane');
    expect(stepsLeft(setup)).toContain('Discord bot');
    expect(stepsLeft(setup)).not.toContain('Officer ranks');
  });
});
