// web/src/test-support/report-fixtures.ts
// The raid-night report's fixtures: an attendance row and a full report (one kill, deaths, buffs,
// loot and a previous raid to compare with), for the Your night, Report and Attendance tests.
import type { RaidAttendanceRow, RaidReport } from '../lib/guild/api';

/** One attendance row: a present DPS signed In. */
export function attendanceRow(extra: Partial<RaidAttendanceRow> = {}): RaidAttendanceRow {
  return {
    character_key: 'us-pvp-kraggor',
    name: 'Kraggor',
    class: 'warrior',
    spec: 'Fury',
    signed: 'in',
    role: 'dps',
    outcome: 'present',
    came_as: null,
    first_seen_offset_s: null,
    pulls_present: 4,
    pulls_total: 4,
    excused: false,
    own: false,
    ...extra,
  };
}

export function raidReport(extra: Partial<RaidReport> = {}): RaidReport {
  return {
    kill_order: [
      {
        encounter_id: 1084,
        name: 'Onyxia',
        fight_index: 4,
        killed: true,
        kill_ms: 102_000,
        wipes_before: 3,
        deaths: 2,
      },
    ],
    unnamed: null,
    deaths: [
      { character_key: 'us-pvp-pyrewisp', name: 'Pyrewisp', class: 'mage', on_wipes: 3, on_kills: 1 },
      { character_key: 'us-pvp-kraggor', name: 'Kraggor', class: 'warrior', on_wipes: 1, on_kills: 1 },
    ],
    consumables: {
      log_has_buffs: true,
      rows: [
        {
          character_key: 'us-pvp-pyrewisp',
          name: 'Pyrewisp',
          class: 'mage',
          before: 'short',
          flask: false,
          food: true,
          potions: 3,
        },
      ],
    },
    loot: [
      {
        encounter_id: 1084,
        boss: 'Onyxia',
        awards: [
          {
            item_id: 18423,
            item_name: 'Head of Onyxia',
            quality: 4,
            character_key: 'us-pvp-kraggor',
            name: 'Kraggor',
            class: 'warrior',
            by_name: 'Bow Jackzon',
            at: '2026-12-11T05:30:00Z',
          },
        ],
      },
    ],
    moves: {
      previous: { id: 6, title: "Onyxia's Lair", starts_at: '2026-12-04T02:00:00Z' },
      rows: [
        {
          character_key: 'us-pvp-kraggor',
          name: 'Kraggor',
          class: 'warrior',
          metric: 'dps',
          encounter: 'Onyxia',
          before: 238,
          after: 291,
        },
      ],
    },
    posted_at: null,
    ...extra,
  };
}
