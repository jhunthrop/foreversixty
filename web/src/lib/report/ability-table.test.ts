// web/src/lib/report/ability-table.test.ts
// Spec 2026-11-04-sim-results section 10, check 1: the table's figures pinned against the fixture
// by hand calculation. Attempts are hits + MISS + DODGE + PARRY.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import type { SimResult } from '../sim/types';
import { abilityTable, attemptsOf, DEFAULT_SORT, nextSort, sortAbilityRows, type AbilityRow } from './ability-table';

const result = full as unknown as SimResult;
const { summary } = result;
const table = abilityTable({
  actors: summary.damage_done,
  casts: summary.casts,
  auras: summary.auras,
  durationMs: summary.duration_ms,
  names: null,
});
const row = (key: string): AbilityRow => {
  const found = table.rows.find((entry) => entry.key === key);
  if (found === undefined) throw new Error(`no row ${key}`);
  return found;
};
const share = (entry: AbilityRow, kind: string): number | undefined =>
  entry.outcomes.find((outcome) => outcome.kind === kind)?.pct;

describe('abilityTable on the Fury fixture', () => {
  it('Execute: 190.8 DPS, 21.1%, 24 casts, 1,493 avg hit, 52.2% crit', () => {
    const execute = row('spell:20662');
    expect(execute.dps).toBeCloseTo(34337 / 179.92, 5);
    expect(execute.dps.toFixed(1)).toBe('190.8');
    expect(execute.share.toFixed(1)).toBe('21.1');
    expect(execute.casts).toBe(24);
    expect(Math.round(execute.avgHit)).toBe(1493);
    expect(execute.crit?.toFixed(1)).toBe('52.2');
  });

  it('Main hand: 63 swings, 506 avg hit, 42.6% crit, miss 19.0, dodge 6.3, glance 39.7', () => {
    const main = row('other:attack/1');
    expect(main.name).toBe('Main hand');
    expect(main.casts).toBe(63);
    expect(Math.round(main.avgHit)).toBe(506);
    expect(main.crit?.toFixed(1)).toBe('42.6');
    expect(share(main, 'MISS')?.toFixed(1)).toBe('19.0');
    expect(share(main, 'DODGE')?.toFixed(1)).toBe('6.3');
    expect(share(main, 'GLANCING')?.toFixed(1)).toBe('39.7');
  });

  it('Extra attacks are their own row: 21 swings, 557 avg hit', () => {
    const extra = row('other:attack/3');
    expect(extra.name).toBe('Extra attacks');
    expect(extra.casts).toBe(21);
    expect(Math.round(extra.avgHit)).toBe(557);
  });

  it('Deep Wounds: 59 ticks, 141 per tick, tick crit 37.3%, no casts, uptime from its debuff', () => {
    const wounds = row('spell:12867');
    expect(wounds.periodic).toBe(true);
    expect(wounds.casts).toBeNull();
    expect(wounds.ticks).toBe(59);
    expect(Math.round(wounds.avgHit)).toBe(141);
    expect(wounds.crit?.toFixed(1)).toBe('37.3');
    expect(wounds.uptime?.toFixed(1)).toBe('100.0');
  });

  it('counts a glance as landed, never as a second attempt', () => {
    expect(attemptsOf({ hits: 47, misses: { MISS: 12, DODGE: 4, GLANCING: 25 } } as never)).toBe(63);
  });

  it('has nine sources, sums to the run, and draws no column it has no figures for', () => {
    expect(table.rows).toHaveLength(9);
    expect(table.total).toBe(162715);
    expect(table.dps.toFixed(1)).toBe('904.4');
    expect(table.hasBiggestHit).toBe(false);
    expect(table.hasUptime).toBe(true);
  });

  it('an older result has no debuff tracks, so no Uptime column', () => {
    const older = abilityTable({
      actors: summary.damage_done,
      casts: summary.casts,
      auras: summary.auras.filter((track) => track.type !== 'DEBUFF'),
      durationMs: summary.duration_ms,
      names: null,
    });
    expect(older.hasUptime).toBe(false);
  });
});

describe('sorting', () => {
  it('a first press sorts descending, a second ascending', () => {
    const crit = nextSort(DEFAULT_SORT, 'crit');
    expect(crit).toEqual({ column: 'crit', direction: 'desc' });
    expect(nextSort(crit, 'crit')).toEqual({ column: 'crit', direction: 'asc' });
  });

  it('sorts by the column, and DPS descending by default', () => {
    expect(sortAbilityRows(table.rows, DEFAULT_SORT)[0]?.key).toBe('spell:20662');
    const lowestCrit = sortAbilityRows(table.rows, { column: 'crit', direction: 'asc' })[0];
    expect(lowestCrit?.key).toBe('spell:23894');
  });
});
