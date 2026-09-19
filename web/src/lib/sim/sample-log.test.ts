import { describe, expect, it } from 'vitest';
import { RESOURCE_ORDER, sampleLog, sampleTime } from './sample-log';
import type { SampleCast } from './types';

const names = { spell: { '25286': 'Heroic Strike', '11305': 'Bloodrage' }, item: {} };

const sample: SampleCast[] = [
  { at_ms: -1500, action: 'spell:11305', resources: { rage: 0 } },
  { at_ms: 320, action: 'spell:25286', target: 'Target', resources: { rage: 42 } },
  { at_ms: 1900, action: 'spell:25286', target: 'Target', resources: { rage: 12, mana: 300 } },
];

describe('sampleTime', () => {
  it('is a signed count of seconds, one decimal', () => {
    expect(sampleTime(-1500)).toBe('-1.5 s');
    expect(sampleTime(0)).toBe('0.0 s');
    expect(sampleTime(1900)).toBe('1.9 s');
  });
});

describe('sampleLog', () => {
  const log = sampleLog(sample, names);

  it('resolves every action key through the build’s table, never rendering the key', () => {
    expect(log.rows.map((row) => row.name)).toEqual(['Bloodrage', 'Heroic Strike', 'Heroic Strike']);
  });

  it('keeps an “other” key readable without a table, the way every cast row does', () => {
    expect(sampleLog([{ at_ms: 0, action: 'other:melee' }], null).rows[0].name).toBe('Melee');
  });

  it('marks the pre-pull casts and counts them, so the table can rule a line under them', () => {
    expect(log.rows.map((row) => row.prePull)).toEqual([true, false, false]);
    expect(log.prePullCount).toBe(1);
  });

  it('keeps the engine’s order rather than sorting, because the order is the point', () => {
    expect(log.rows.map((row) => row.atMs)).toEqual([-1500, 320, 1900]);
  });

  it('gives every row a key of its own, so two casts of one spell at one instant still render', () => {
    const doubled = sampleLog([sample[1], sample[1]], names);
    expect(new Set(doubled.rows.map((row) => row.key)).size).toBe(2);
  });

  it('columns are the resources that appear, in the engine’s own order', () => {
    expect(log.columns).toEqual(['mana', 'rage']);
    expect(RESOURCE_ORDER.indexOf('mana')).toBeLessThan(RESOURCE_ORDER.indexOf('rage'));
  });

  it('puts a resource nothing anticipated after the known ones rather than dropping it', () => {
    const odd = sampleLog([{ at_ms: 0, action: 'spell:1', resources: { rage: 1, zeal: 2 } }], names);
    expect(odd.columns).toEqual(['rage', 'zeal']);
  });

  it('is empty, not a throw, for a result with no sample at all', () => {
    expect(sampleLog(undefined, names)).toEqual({ rows: [], columns: [], prePullCount: 0 });
  });
});
