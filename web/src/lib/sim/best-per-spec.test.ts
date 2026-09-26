// web/src/lib/sim/best-per-spec.test.ts
import { describe, expect, it } from 'vitest';
import { bestPerSpec } from './best-per-spec';
import type { SimListRow } from './types';

function row(overrides: Partial<SimListRow> & Pick<SimListRow, 'sim_id' | 'spec' | 'dps'>): SimListRow {
  return {
    engine_version: 'test-engine',
    created_at: '2026-09-20T00:00:00Z',
    title: '',
    ...overrides,
  };
}

describe('bestPerSpec', () => {
  it('returns nothing for an empty list', () => {
    expect(bestPerSpec([])).toEqual([]);
  });

  it('keeps one row per spec, the highest DPS one', () => {
    const rows = [
      row({ sim_id: 'a', spec: 'warrior-fury', dps: 1000, created_at: '2026-09-20T00:00:00Z' }),
      row({ sim_id: 'b', spec: 'warrior-fury', dps: 1200, created_at: '2026-09-21T00:00:00Z' }),
      row({ sim_id: 'c', spec: 'warrior-fury', dps: 900, created_at: '2026-09-19T00:00:00Z' }),
    ];
    expect(bestPerSpec(rows)).toEqual([
      { spec: 'warrior-fury', dps: 1200, createdAt: '2026-09-21T00:00:00Z' },
    ]);
  });

  it('sorts multiple specs best DPS first', () => {
    const rows = [
      row({ sim_id: 'a', spec: 'warrior-fury', dps: 1000 }),
      row({ sim_id: 'b', spec: 'warrior-arms', dps: 1400 }),
      row({ sim_id: 'c', spec: 'mage-fire', dps: 1200 }),
    ];
    expect(bestPerSpec(rows).map((r) => r.spec)).toEqual(['warrior-arms', 'mage-fire', 'warrior-fury']);
  });
});
