// web/src/lib/guides/rail-stats.test.ts
import { describe, expect, it } from 'vitest';
import { railStatRows } from './rail-stats';
import { bandEntry, loadBisFile } from '../bis/load';

const BUILD = '1.60.1.70009';
const STAT_PRIORITY = ['Attack power', 'Strength', 'Agility', 'Critical strike', 'Hit', 'Melee haste'];

describe('railStatRows', () => {
  it('normalizes Fury’s band-60 weights so the top significant stat reads 1.00 (mock board worked example)', () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined) throw new Error('band 60 alliance missing');

    const rows = railStatRows(STAT_PRIORITY, band.weights, 'warrior-fury');
    const byLabel = new Map(rows.map((row) => [row.label, row.value]));

    expect(byLabel.get('Melee haste')).toBeCloseTo(1.0, 2);
    expect(byLabel.get('Attack power')).toBeCloseTo(0.15, 2);
    expect(byLabel.get('Critical strike')).toBeCloseTo(0.14, 2);
    expect(byLabel.get('Hit')).toBeCloseTo(0.03, 2);
    expect(byLabel.get('Strength')).toBeUndefined();
    expect(byLabel.get('Agility')).toBeUndefined();
  });

  it('never drops a row, even one with no match in weights', () => {
    const rows = railStatRows(['Not a real stat'], [], 'warrior-fury');
    expect(rows).toHaveLength(1);
    expect(rows[0]!.value).toBeUndefined();
  });
});
