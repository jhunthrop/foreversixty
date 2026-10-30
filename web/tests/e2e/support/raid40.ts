// web/tests/e2e/support/raid40.ts
// A synthetic 40-player pull for the chart specs and captures (controls and charts spec §D
// captures 9 to 11), served over the fixture report's fight 3 by route mocking. It is built from
// the fixture's own shape at test time and never written anywhere the site serves: the checked-in
// fixture stays a 5-player log. The pull is shaped like the boards' Ragnaros: 5:40 long, the boss
// submerged and untargetable (no damage at all) from 3:00 to 4:30, one tank death at 2:06 with
// its killing blow, a taunt and a target change a second later, and a Rebirth and a Power
// Infusion for the cast lane.
import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import type { Page } from '@playwright/test';

const FIXTURE = join(import.meta.dirname, '../../../src/fixtures/report');
export const RAID40_FIGHT = 3;
export const RAID40_URL = `/reports/fixture2abcd?fight=${RAID40_FIGHT}`;
export const BOSS_GUID = 'Creature-0-2085-2284-7855-11502-0000AA0099';
const BOSS = 'Ragnaros';
const DURATION_S = 340;
const SUBMERGE = { start: 180, end: 270 };
const CLASSES = ['Warrior', 'Rogue', 'Mage', 'Warlock', 'Hunter', 'Priest', 'Druid', 'Shaman'] as const;
const NAMES = [
  'Garrosk',
  'Durgan',
  'Ashka',
  'Kargath',
  'Velyn',
  'Zaxis',
  'Morrigan',
  'Ysolde',
  'Thandril',
  'Brokk',
  'Rukh',
  'Hrothgar',
  'Bek',
  'Lightningbolltz',
  'Orla',
  'Tessane',
  'Mirev',
  'Quill',
  'Sabra',
  'Torvin',
  'Ulfric',
  'Vessa',
  'Wren',
  'Yarrow',
  'Zephra',
  'Aldric',
  'Brynn',
  'Caddis',
  'Dorran',
  'Elske',
  'Fenwick',
  'Galdra',
  'Halvar',
  'Iska',
  'Jorund',
  'Kesra',
  'Lothar',
  'Maelis',
  'Nyx',
  'Osric',
];

export const guidOf = (index: number): string => `Player-4184-0000B${String(index).padStart(3, '0')}`;
/** Thandril and Brokk tank; the next eight heal; the rest deal damage. */
export const TANKS = [8, 9];
export const BROKK = 9;
/** Lightningbolltz, a damage-dealing shaman well outside the top 8: the viewer's own character. */
export const YOU = 13;
const roleOf = (index: number): 'tank' | 'healer' | 'dps' =>
  TANKS.includes(index) ? 'tank' : index >= 30 ? 'healer' : 'dps';

/** A seeded wobble, so every run draws the same lines. */
function wobble(index: number, second: number): number {
  const x = Math.sin(index * 12.9898 + second * 78.233) * 43758.5453;
  return x - Math.floor(x);
}

function damageSeries(index: number): number[] {
  const role = roleOf(index);
  const base = role === 'tank' ? 180 : role === 'healer' ? 40 : 600 - index * 6;
  return Array.from({ length: DURATION_S }, (_, second) => {
    if (second >= SUBMERGE.start && second < SUBMERGE.end) return 0;
    if (index === BROKK && second > 126) return 0;
    return Math.round(base * (0.6 + wobble(index, second) * 0.8));
  });
}

const sum = (values: number[]): number => values.reduce((total, value) => total + value, 0);

/** The 40-player summary, on the fixture's fight-3 summary's shape. */
export function raid40Summary(): Record<string, unknown> {
  const base = JSON.parse(readFileSync(join(FIXTURE, 'fights/3/summary.json'), 'utf8')) as Record<
    string,
    unknown
  >;
  const players = NAMES.map((name, index) => ({
    guid: guidOf(index),
    name: `${name}-Nightslayer`,
    class: TANKS.includes(index)
      ? 'Warrior'
      : index === YOU
        ? 'Shaman'
        : roleOf(index) === 'healer'
          ? 'Priest'
          : CLASSES[index % CLASSES.length],
    role: roleOf(index),
    series: damageSeries(index),
  }));
  const actor = (player: (typeof players)[number], series: number[]): Record<string, unknown> => ({
    guid: player.guid,
    name: player.name,
    class: player.class,
    total: sum(series),
    effective: sum(series),
    active_ms: series.filter((value) => value > 0).length * 1000,
    abilities: [],
    targets: [{ guid: BOSS_GUID, name: BOSS, total: sum(series) }],
    series,
  });
  const healing = players.map((player, index) => {
    const series = Array.from({ length: DURATION_S }, (_, second) =>
      roleOf(index) === 'healer' ? Math.round(400 * (0.5 + wobble(index + 99, second))) : 0,
    );
    return actor(player, series);
  });
  const taken = players.map((player, index) => {
    const series = Array.from({ length: DURATION_S }, (_, second) =>
      TANKS.includes(index) && !(second >= SUBMERGE.start && second < SUBMERGE.end) ? 300 : 0,
    );
    return actor(player, series);
  });
  const threat = players.map((player, index) => {
    const series = damageSeries(index);
    return {
      guid: player.guid,
      name: player.name,
      target_guid: BOSS_GUID,
      target_name: BOSS,
      threat: sum(series),
      series,
    };
  });
  const brokk = players[BROKK];
  const thandril = players[TANKS[0]];
  return {
    ...base,
    duration_ms: DURATION_S * 1000,
    roster: players.map((player, index) => ({
      guid: player.guid,
      name: player.name,
      class: player.class,
      class_source: 'combatant_info',
      role: roleOf(index),
      active_ms: 0,
      activity_pct: 0,
      deaths: index === BROKK ? 1 : 0,
      damage_done: sum(player.series),
      healing_done: 0,
      damage_taken: 0,
      dps: sum(player.series) / DURATION_S,
      hps: 0,
      dtps: 0,
    })),
    damage_done: players.map((player) => actor(player, player.series)),
    healing,
    healing_taken: [],
    damage_taken: taken,
    deaths: [
      {
        guid: brokk.guid,
        name: brokk.name,
        class: 'Warrior',
        at_ms: 126_000,
        killing_blow: {
          at_ms: 126_000,
          source_guid: BOSS_GUID,
          source_name: BOSS,
          spell_id: 0,
          spell_name: 'Melee',
          amount: 3862,
          overkill: 1104,
        },
        last: [],
        auras_held: [],
        auras_lost: [],
      },
    ],
    casts: [
      {
        guid: players[31].guid,
        name: players[31].name,
        spell_id: 10060,
        spell_name: 'Power Infusion',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [62_000],
      },
      {
        guid: players[32].guid,
        name: players[32].name,
        spell_id: 20484,
        spell_name: 'Rebirth',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [140_000],
      },
    ],
    auras: [],
    interrupts: [],
    dispels: [],
    resources: [],
    combatants: [],
    threat: players.map((player) => ({
      guid: player.guid,
      name: player.name,
      threat: sum(player.series),
      model_version: 'base-1',
      complete: false,
    })),
    threat_by_target: threat,
    taunts: [
      {
        at_ms: 127_000,
        source_guid: thandril.guid,
        source_name: thandril.name,
        target_guid: BOSS_GUID,
        target_name: BOSS,
        spell_id: 355,
        spell_name: 'Taunt',
      },
    ],
    target_changes: [
      {
        at_ms: 2000,
        enemy_guid: BOSS_GUID,
        enemy_name: BOSS,
        to_guid: brokk.guid,
        to_name: brokk.name,
        basis: 'melee',
      },
      {
        at_ms: 127_000,
        enemy_guid: BOSS_GUID,
        enemy_name: BOSS,
        to_guid: thandril.guid,
        to_name: thandril.name,
        basis: 'melee',
      },
    ],
    phases: [
      { name: 'Ragnaros', start_ms: 0, end_ms: SUBMERGE.start * 1000 },
      { name: 'Submerge', start_ms: SUBMERGE.start * 1000, end_ms: SUBMERGE.end * 1000 },
      { name: 'Ragnaros', start_ms: SUBMERGE.end * 1000, end_ms: DURATION_S * 1000 },
    ],
  };
}

/** The fixture's report.json with the 40 players and Ragnaros as units, and fight 3 at 5:40. */
function raid40Report(): Record<string, unknown> {
  const report = JSON.parse(readFileSync(join(FIXTURE, 'report.json'), 'utf8')) as {
    fights: Record<string, unknown>[];
    units: Record<string, unknown>[];
  };
  const seen = { first_seen: '2026-09-26T20:12:00Z', last_seen: '2026-09-26T20:17:40Z' };
  return {
    ...report,
    fights: report.fights.map((fight) =>
      fight.index === RAID40_FIGHT
        ? {
            ...fight,
            duration_ms: DURATION_S * 1000,
            players: NAMES.map((_, index) => guidOf(index)),
            deaths: 1,
          }
        : fight,
    ),
    units: [
      ...report.units,
      ...NAMES.map((name, index) => ({
        guid: guidOf(index),
        name: `${name}-Nightslayer`,
        kind: 'player',
        flags: 1297,
        ...seen,
      })),
      { guid: BOSS_GUID, name: BOSS, kind: 'creature', npc_id: 11502, flags: 2632, ...seen },
    ],
  };
}

/** Serves the 40-player pull over fight 3. Call before `page.goto`. */
export async function serveRaid40(page: Page): Promise<void> {
  const summary = JSON.stringify(raid40Summary());
  const report = JSON.stringify(raid40Report());
  await page.route(/\/reports\/fixture2abcd\/fights\/3\/summary\.json/, (route) =>
    route.fulfill({ contentType: 'application/json', body: summary }),
  );
  await page.route(/\/reports\/fixture2abcd\/report\.json/, (route) =>
    route.fulfill({ contentType: 'application/json', body: report }),
  );
}

/** Every player's DPS over the whole pull, rounded, for the head's reconciliation. */
export function raid40DpsSum(): number {
  return NAMES.reduce((total, _, index) => total + Math.round(sum(damageSeries(index)) / DURATION_S), 0);
}

/** Makes Lightningbolltz the viewer's current character, so the chart pins them as you. */
export async function beLightningbolltz(page: Page): Promise<void> {
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/nightslayer/lightningbolltz',
        label: 'Lightningbolltz · Elemental Shaman',
        classSlug: 'shaman',
        savedAt: '2026-10-30T00:00:00Z',
      }),
    );
  });
}
