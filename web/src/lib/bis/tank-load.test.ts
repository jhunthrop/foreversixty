// web/src/lib/bis/tank-load.test.ts
import { describe, expect, it } from 'vitest';
import { loadBisFile, normaliseBisFile } from './load';
import { bandsOf, depsWithSpec, fileOf, tankBand } from './tank-test-support';
import { bandInfosFor } from './panel-view';
import { tankCopy } from './copy';
import type { BisBand } from './types';
import { isTankMetrics } from './tank-view';

function effectiveHealthOf(band: BisBand | undefined): number | undefined {
  return isTankMetrics(band?.metrics) ? band.metrics.effective_health : undefined;
}

describe('normaliseBisFile: role and metrics', () => {
  it('defaults an older file to role dps and metrics null', () => {
    const [band] = normaliseBisFile(fileOf([tankBand({ role: undefined, metrics: undefined })])).bands;
    expect(band?.role).toBe('dps');
    expect(band?.metrics).toBeNull();
  });

  it('keeps a tank band role and metrics as published', () => {
    const [band] = bandsOf(normaliseBisFile(fileOf([tankBand()])));
    expect(band?.role).toBe('tank');
    expect(effectiveHealthOf(band)).toBe(11235);
  });

  it('fails fast on a tank band that carries no metrics', () => {
    expect(() => normaliseBisFile(fileOf([tankBand({ metrics: null })]))).toThrow(/role tank but no metrics/);
  });
});

describe('the warrior-protection fixture', () => {
  it('is a tank file whose raid and bare bands carry different metrics', () => {
    const file = loadBisFile('warrior-protection', 'no-such-build');
    const bands = file?.bands.filter((band) => band.faction === 'alliance') ?? [];
    expect(bands.map((band) => band.role)).toEqual(['tank', 'tank']);
    expect(new Set(bands.map(effectiveHealthOf)).size).toBe(2);
    expect(file?.bands.every((band) => band.score_unit === 'tank_score')).toBe(true);
  });
});

describe('bandInfosFor on a tank band', () => {
  const file = normaliseBisFile(fileOf([tankBand()]));
  const [info] = bandInfosFor(file, [60], 'alliance', depsWithSpec('warrior-protection'));

  it('carries the tank headline and the score unit', () => {
    expect(info?.tank?.figures).toHaveLength(4);
    expect(info?.scoreUnit).toBe('tank_score');
  });

  it('labels per-point weights and the scale note as score, never DPS', () => {
    expect(info?.scaleRows.every((row) => row.perPointUnit === tankCopy.scoreWord)).toBe(true);
    expect(info?.scaleNoteLine).toContain('score per point');
    expect(info?.scaleNoteLine).not.toContain('DPS');
  });

  it('renders rating-family stats with the rating suffix and normalises to the top stat at 1.00', () => {
    const labels = info?.scaleRows.map((row) => row.label) ?? [];
    expect(labels).toContain('Defense rating');
    expect(info?.scaleRows[0]).toMatchObject({ stat: 'defense', scaleFactor: 1 });
  });

  it('words slot figures in score points', () => {
    const head = info?.rows.find((row) => row.slot === 'head');
    expect(head?.alternatives?.[0]?.unit).toBe('tank_score');
    expect(head?.evidenceLine).toBe(tankCopy.evidenceLineDelta('Spare Helm', 1.3));
    expect(head?.verifiedGlyphTitle).toBeUndefined();
  });
});
