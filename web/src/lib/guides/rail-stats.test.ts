// web/src/lib/guides/rail-stats.test.ts
import { describe, expect, it } from 'vitest';
import { railStatRows } from './rail-stats';
import { bandEntry, loadBisFile } from '../bis/load';

const BUILD = '1.60.1.70009';
const STAT_PRIORITY = ['Attack power', 'Strength', 'Agility', 'Critical strike', 'Hit', 'Melee haste'];

describe('railStatRows', () => {
  it('normalizes to the top significant non-haste stat and reports haste per 1%, outside the scale', () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined) throw new Error('band 60 alliance missing');

    const rows = railStatRows(STAT_PRIORITY, band.weights, 'warrior-fury', band.haste_scale_factor ?? null);
    const nonHaste = rows
      .filter((row) => !row.perPercent && row.value !== undefined)
      .map((row) => row.value!);
    expect(Math.max(...nonHaste)).toBeCloseTo(1.0, 5);
    const haste = rows.find((row) => row.label === 'Melee haste')!;
    expect(haste.perPercent).toBe(true);
    if (band.haste_scale_factor != null) expect(haste.value).toBeCloseTo(band.haste_scale_factor, 5);
  });

  it('never drops a row, even one with no match in weights', () => {
    const rows = railStatRows(['Not a real stat'], [], 'warrior-fury');
    expect(rows).toHaveLength(1);
    expect(rows[0]!.value).toBeUndefined();
  });
});
