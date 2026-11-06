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
    auto_outcome: extra.outcome ?? 'present',
    auto_excused: extra.excused ?? false,
    override: null,
    ...extra,
  };
}

export function raidReport(extra: Partial<RaidReport> = {}): RaidReport {
  return {
    headline: { kills: 1, wipes: 3, attempts: 4, deaths: 14 },
    pulls: [
      {
        fight_index: 1,
        encounter_id: null,
        name: 'Onyxian Whelp pack',
        attempt: null,
        result: 'clear',
        duration_ms: 60_000,
        deaths: 0,
        death_causes: [],
      },
      {
        fight_index: 2,
        encounter_id: 1084,
        name: 'Onyxia',
        attempt: 1,
        result: 'wipe',
        duration_ms: 240_000,
        deaths: 2,
        death_causes: [
          {
            character_key: 'us-pvp-pyrewisp',
            name: 'Pyrewisp',
            class: 'mage',
            at_ms: 102_000,
            killing_blow: {
              at_ms: 102_000,
              spell_id: 18435,
              ability: 'Flame Breath',
              source: 'Onyxia',
              amount: 4210,
              overkill: 1050,
            },
            last_hits: [
              {
                at_ms: 99_000,
                spell_id: 1,
                ability: 'Melee',
                source: 'Onyxian Whelp',
                amount: 310,
                overkill: 0,
              },
              {
                at_ms: 102_000,
                spell_id: 18435,
                ability: 'Flame Breath',
                source: 'Onyxia',
                amount: 4210,
                overkill: 1050,
              },
            ],
          },
          {
            character_key: 'us-pvp-kraggor',
            name: 'Kraggor',
            class: 'warrior',
            at_ms: 180_000,
            killing_blow: null,
            last_hits: [],
          },
        ],
      },
      {
        fight_index: 4,
        encounter_id: 1084,
        name: 'Onyxia',
        attempt: 2,
        result: 'kill',
        duration_ms: 102_000,
        deaths: 1,
        death_causes: null,
      },
    ],
    parses: [
      {
        character_key: 'us-pvp-pyrewisp',
        name: 'Pyrewisp',
        class: 'mage',
        role: 'dps',
        encounter_id: 1084,
        encounter: 'Onyxia',
        fight_index: 4,
        metric: 'dps',
        value: 812.4,
        rating: 96,
      },
      {
        character_key: 'us-pvp-menda',
        name: 'Menda',
        class: 'priest',
        role: 'healer',
        encounter_id: 1084,
        encounter: 'Onyxia',
        fight_index: 4,
        metric: 'hps',
        value: 1204,
        rating: null,
      },
    ],
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
      by_pull: [
        { fight_index: 2, name: 'Onyxia', attempt: 1, reason: 'no_character_sheet', rows: null },
        {
          fight_index: 4,
          name: 'Onyxia',
          attempt: 2,
          reason: null,
          rows: [
            {
              character_key: 'us-pvp-pyrewisp',
              name: 'Pyrewisp',
              class: 'mage',
              flask: true,
              food: true,
              weapon_enchant: null,
              potions: 1,
              brought: ['Flask of Supreme Power', 'Runn Tum Tuber Surprise'],
            },
          ],
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
            icon: 'inv_misc_head_dragon_01',
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
