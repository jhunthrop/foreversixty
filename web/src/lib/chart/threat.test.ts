// web/src/lib/chart/threat.test.ts
import { describe, expect, it } from 'vitest';
import { castEvent, targetChangeEvent } from './events';
import type { TargetChange, ThreatPair } from '../report/types';
import { attackedPlayers, buildThreatModel, shareAt, type ThreatPlayer } from './threat';

const BOSS = 'Creature-0-1-1-11502-1';

/** Twelve players building threat on Ragnaros; two tanks low, as the base-1 model draws them. */
function fixture(): { pairs: ThreatPair[]; players: Map<string, ThreatPlayer> } {
  const players = new Map<string, ThreatPlayer>();
  const pairs: ThreatPair[] = [];
  for (let index = 0; index < 12; index += 1) {
    const guid = `Player-${index}`;
    const role = index >= 10 ? 'tank' : 'dps';
    players.set(guid, {
      name: `Raider${index}-Realm`,
      className: role === 'tank' ? 'Warrior' : 'Mage',
      role,
    });
    const perSecond = role === 'tank' ? 100 : 1000 - index * 10;
    const series = Array.from({ length: 160 }, () => perSecond);
    pairs.push({
      guid,
      name: `Raider${index}`,
      target_guid: BOSS,
      target_name: 'Ragnaros',
      threat: perSecond * 160,
      series,
    });
  }
  return { pairs, players };
}

describe('buildThreatModel', () => {
  it('draws the top 8 by standing plus the tanks, whatever their threat', () => {
    const { pairs, players } = fixture();
    const model = buildThreatModel({
      pairs,
      enemyGuids: [BOSS],
      players,
      window: { startMs: 0, endMs: 160_000 },
      taunts: [],
      chosen: null,
    });
    const ids = model.series.map((line) => line.id);
    expect(ids).toHaveLength(10);
    expect(ids).toContain('Player-10');
    expect(ids).toContain('Player-11');
    expect(model.series.find((line) => line.id === 'Player-10')!.pinned).toBe(true);
    // Cumulative: the standing at the window's end is the last second's running total.
    expect(model.series[0].figure).toBe(1000 * 160);
  });

  it('always draws whoever taunted the enemy in the window', () => {
    const { pairs, players } = fixture();
    const model = buildThreatModel({
      pairs,
      enemyGuids: [BOSS],
      players,
      window: { startMs: 90_000, endMs: 160_000 },
      taunts: [
        {
          at_ms: 127_000,
          source_guid: 'Player-9',
          source_name: 'Raider9',
          target_guid: BOSS,
          target_name: 'Ragnaros',
          spell_id: 355,
          spell_name: 'Taunt',
        },
      ],
      chosen: null,
    });
    expect(model.series.map((line) => line.id)).toContain('Player-9');
  });

  it('draws one Source player and the always-drawn players only', () => {
    const { pairs, players } = fixture();
    const model = buildThreatModel({
      pairs,
      enemyGuids: [BOSS],
      players,
      window: { startMs: 0, endMs: 160_000 },
      taunts: [],
      sourcePlayer: 'Player-5',
      chosen: null,
    });
    expect(model.series.map((line) => line.id).sort()).toEqual(['Player-10', 'Player-11', 'Player-5']);
  });

  it('shares against every player’s estimated threat on the enemy at that second', () => {
    const { pairs, players } = fixture();
    const model = buildThreatModel({
      pairs,
      enemyGuids: [BOSS],
      players,
      window: { startMs: 0, endMs: 160_000 },
      taunts: [],
      chosen: null,
    });
    const top = model.series[0];
    const at = 127_000;
    const value = top.values[127]!;
    const share = shareAt(model, value, at);
    const everyone = [...players.keys()].length;
    expect(share).toBeGreaterThan(0);
    expect(share).toBeLessThan(1 / everyone + 0.05);
    const sum = model.entries.reduce((total, entry) => total + shareAt(model, entry.figure, 160_000 - 1), 0);
    expect(sum).toBeCloseTo(1, 6);
  });
});

describe('target changes (spec §B12: drawn only where the summary carries them)', () => {
  const changes: TargetChange[] = [
    {
      at_ms: 20_000,
      enemy_guid: BOSS,
      enemy_name: 'Ragnaros',
      to_guid: 'Player-10',
      to_name: 'Raider10',
      basis: 'melee',
    },
    {
      at_ms: 127_000,
      enemy_guid: BOSS,
      enemy_name: 'Ragnaros',
      to_guid: 'Player-11',
      to_name: 'Raider11',
      basis: 'melee',
    },
    {
      at_ms: 140_000,
      enemy_guid: 'other',
      enemy_name: 'Son of Flame',
      to_guid: 'Player-3',
      to_name: 'Raider3',
      basis: 'melee',
    },
  ];

  it('counts the target as the window opened and every change inside it, on this enemy only', () => {
    expect([...attackedPlayers([BOSS], { startMs: 90_000, endMs: 160_000 }, changes)].sort()).toEqual([
      'Player-10',
      'Player-11',
    ]);
  });

  it('pins the attacked players on the chart', () => {
    const { pairs, players } = fixture();
    const model = buildThreatModel({
      pairs,
      enemyGuids: [BOSS],
      players: new Map([...players].map(([guid, player]) => [guid, { ...player, role: 'dps' as const }])),
      window: { startMs: 90_000, endMs: 160_000 },
      taunts: [],
      targetChanges: changes,
      chosen: null,
    });
    expect(model.series.find((line) => line.id === 'Player-11')!.pinned).toBe(true);
  });

  it('turns each change into a marker line', () => {
    expect(targetChangeEvent(changes[1]).kind).toBe('target-change');
    expect(targetChangeEvent(changes[1]).atMs).toBe(127_000);
  });

  it('falls back to damage taken when the summary has no target changes', () => {
    const hit = {
      guid: 'Player-4',
      name: 'Raider4',
      total: 1,
      effective: 1,
      active_ms: 0,
      abilities: [],
      targets: [{ guid: BOSS, name: 'Ragnaros', total: 1 }],
      series: Array.from({ length: 160 }, (_, second) => (second === 100 ? 500 : 0)),
    };
    expect([...attackedPlayers([BOSS], { startMs: 90_000, endMs: 160_000 }, undefined, [hit])]).toEqual([
      'Player-4',
    ]);
    expect([...attackedPlayers([BOSS], { startMs: 0, endMs: 50_000 }, undefined, [hit])]).toEqual([]);
  });

  it('a cast event still reads with no target changes at all', () => {
    expect(castEvent(1000, 'Power Infusion', 'Sunwick-Nightslayer', 10060).text).toBe(
      'Power Infusion · Sunwick',
    );
  });
});
