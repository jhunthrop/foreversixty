// web/src/lib/guides/rail-stats.test.ts
import { describe, expect, it } from 'vitest';
import { railStatRows, railStatValueText, rankedStatLabels } from './rail-stats';
import { bandEntry, loadBisFile } from '../bis/load';
import type { BisStatWeight } from '../bis/types';
import { statLabelForSpec } from '../sim/weights-display';

const BUILD = '1.60.1.70374';
const STAT_PRIORITY = ['Attack power', 'Strength', 'Agility', 'Critical strike', 'Hit', 'Melee haste'];

describe('railStatRows', () => {
  it('normalizes to the top significant non-haste stat and reports haste per 1%, outside the scale', () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined) throw new Error('band 60 alliance missing');

    const rows = railStatRows(STAT_PRIORITY, band.weights, 'warrior-fury', {
      scaleFactor: band.haste_scale_factor ?? null,
      onItems: true,
    });
    const nonHaste = rows
      .filter((row) => !row.perPercent && row.value !== undefined)
      .map((row) => row.value!);
    expect(Math.max(...nonHaste)).toBeCloseTo(1.0, 5);
    const haste = rows.find((row) => row.label === 'Melee haste')!;
    expect(haste.perPercent).toBe('haste');
    if (band.haste_scale_factor != null) expect(haste.value).toBeCloseTo(band.haste_scale_factor, 5);
  });

  it('orders rows by the sim, not by the guide written statPriority order (real Fury band 60)', () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined) throw new Error('band 60 alliance missing');

    const rows = railStatRows(STAT_PRIORITY, band.weights, 'warrior-fury', {
      scaleFactor: band.haste_scale_factor ?? null,
      onItems: true,
    });
    // Every written stat is kept; the order is a property of the data, never pinned.
    expect(rows.map((row) => row.label).sort()).toEqual([...STAT_PRIORITY].sort());
    const perPoint = rows.filter((row) => !row.perPercent && row.value !== undefined);
    const hasteIndex = rows.findIndex((row) => row.perPercent);
    const firstInsignificant = rows.findIndex((row) => !row.perPercent && row.value === undefined);
    // Per-point tier first, sorted descending, its top at exactly 1.00.
    expect(perPoint[0]!.value).toBeCloseTo(1.0, 5);
    for (let i = 1; i < perPoint.length; i += 1) {
      expect(perPoint[i]!.value!).toBeLessThanOrEqual(perPoint[i - 1]!.value!);
    }
    // Haste after every per-point row, and insignificant rows after haste.
    expect(hasteIndex).toBeGreaterThanOrEqual(perPoint.length);
    if (firstInsignificant !== -1) expect(firstInsignificant).toBeGreaterThan(hasteIndex);
  });

  it('reads expertise per 1%, outside the normalization, so it never tops a tank table', () => {
    const weights: BisStatWeight[] = [
      { stat: 'stamina', weight: 1, error: 0.001, scale_factor: 6, dps_per_point: 158 },
      { stat: 'defense', weight: 1.9, error: 0.25, unit: 'rating', scale_factor: 11.4, dps_per_point: 298 },
      { stat: 'expertise', weight: 9.8, error: 0.7, scale_factor: 59.4, dps_per_point: 1547 },
    ];
    const rows = railStatRows(['Expertise', 'Stamina', 'Defense'], weights, 'warrior-protection');
    expect(rows.map((row) => row.label)).toEqual(['Defense', 'Stamina', 'Expertise']);
    expect(rows[0]!.value).toBeCloseTo(1, 5);
    expect(rows[2]!.perPercent).toBe('expertise');
    expect(rows[2]!.value).toBeCloseTo(59.4 / 11.4, 5);
    expect(railStatValueText(rows[2]!)).toBe(`${(59.4 / 11.4).toFixed(1)} per 1% expertise`);
  });

  it('reads a weight inside its error as Not significant and one below zero beyond it as No gain', () => {
    const weights: BisStatWeight[] = [
      { stat: 'spell_power', weight: 1, error: 0.01, scale_factor: 1, dps_per_point: 0.6 },
      { stat: 'spirit', weight: -0.07, error: 0.28, scale_factor: -0.11 },
      { stat: 'mp5', weight: 0.05, error: 0.3, insignificant: false, scale_factor: 0.05 },
      { stat: 'spell_haste', weight: -3.5, error: 0.64, insignificant: false, scale_factor: -6.4 },
    ];
    const rows = railStatRows(['Spell haste', 'Spirit', 'MP5', 'Spell power'], weights, 'mage-fire', {
      scaleFactor: -6.4,
      onItems: true,
    });
    const byLabel = new Map(rows.map((row) => [row.label, row]));
    expect(byLabel.get('Spirit')!.reading).toBe('not-significant');
    expect(byLabel.get('MP5')!.reading).toBe('not-significant');
    expect(byLabel.get('Spell haste')!.reading).toBe('no-gain');
    for (const row of rows.filter((r) => r.reading !== 'value')) {
      expect(row.value).toBeUndefined();
      expect(row.perPoint).toBeUndefined();
    }
    expect(railStatValueText(byLabel.get('Spell haste')!)).toBe('No gain');
    expect(railStatValueText(byLabel.get('Spirit')!)).toBe('Not significant');
    // Value first, then No gain, then Not significant.
    expect(rows.map((row) => row.label)).toEqual(['Spell power', 'Spell haste', 'Spirit', 'MP5']);
  });

  it('never prints a negative number for any spec on the real files', () => {
    for (const spec of [
      'mage-fire',
      'priest-holy',
      'priest-discipline',
      'shaman-elemental',
      'druid-restoration',
    ]) {
      const file = loadBisFile(spec, BUILD);
      if (file === null) continue;
      for (const band of file.bands.filter((entry) => entry.band === 60)) {
        const labels = band.weights.map((w) => statLabelForSpec(w.stat, spec));
        const rows = railStatRows(labels, band.weights, spec, {
          scaleFactor: band.haste_scale_factor ?? null,
          onItems: band.haste_on_items ?? true,
        });
        for (const row of rows) {
          if (row.value !== undefined) expect(row.value).toBeGreaterThanOrEqual(0);
          expect(railStatValueText(row)).not.toMatch(/^-/);
        }
      }
    }
  });

  it('keeps the guide written order as a stable tiebreak within a tier', () => {
    const rows = railStatRows(['Hit', 'Attack power'], [], 'warrior-fury');
    expect(rows.map((row) => row.label)).toEqual(['Hit', 'Attack power']);
  });

  it('never drops a row, even one with no match in weights', () => {
    const rows = railStatRows(['Not a real stat'], [], 'warrior-fury');
    expect(rows).toHaveLength(1);
    expect(rows[0]!.value).toBeUndefined();
  });

  it('reads haste no item of the band carries as Not on any item, after the ranked stats', () => {
    const weights: BisStatWeight[] = [
      { stat: 'spell_power', weight: 1, error: 0.01, scale_factor: 1, dps_per_point: 0.5 },
      { stat: 'spell_haste', weight: 4.9, error: 0.4, scale_factor: 16 },
      { stat: 'spirit', weight: 0.5, error: 0.9 },
    ];
    const rows = railStatRows(['Spell haste', 'Spirit', 'Spell power'], weights, 'priest-shadow', {
      scaleFactor: 8.1,
      onItems: false,
    });
    expect(rows.map((row) => [row.label, row.reading])).toEqual([
      ['Spell power', 'value'],
      ['Spell haste', 'not-on-items'],
      ['Spirit', 'not-significant'],
    ]);
    expect(rows[1]!.value).toBeCloseTo(8.1, 5);
    expect(railStatValueText(rows[1]!)).toBe('Not on any item');
  });
});

describe('rankedStatLabels', () => {
  it('ranks only the rows that read as a value, so the pills match the table', () => {
    const weights: BisStatWeight[] = [
      { stat: 'hit', weight: 1.2, error: 0.06, scale_factor: 1.2 },
      { stat: 'spell_power', weight: 1, error: 0.002, scale_factor: 1 },
      { stat: 'spell_haste', weight: 0.9, error: 0.94, insignificant: true },
      { stat: 'spell_penetration', weight: 0, error: 0, insignificant: true },
    ];
    const rows = railStatRows(
      ['Hit', 'Spell power', 'Spell haste', 'Spell penetration'],
      weights,
      'mage-arcane',
    );
    expect(rankedStatLabels(rows, [])).toEqual(['Hit', 'Spell power']);
  });

  it('drops haste no item carries from the pills', () => {
    const weights: BisStatWeight[] = [
      { stat: 'strength', weight: 2, error: 0.01, scale_factor: 2 },
      { stat: 'melee_haste', weight: 16.6, error: 2 },
    ];
    const rows = railStatRows(['Strength', 'Melee haste'], weights, 'warrior-fury', {
      scaleFactor: 8.3,
      onItems: false,
    });
    expect(rankedStatLabels(rows, [])).toEqual(['Strength']);
  });

  it('falls back to the guide’s own list for a spec with no weights yet', () => {
    expect(rankedStatLabels([], ['Hit', 'Spell power'])).toEqual(['Hit', 'Spell power']);
  });
});
