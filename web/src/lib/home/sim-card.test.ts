import { describe, expect, it } from 'vitest';
import type { BisBand } from '../bis/types';
import type { SimListRow } from '../sim/types';
import { simCardFigures, simMatchesBand } from './sim-card';

const SPEC = 'hunter-marksmanship';
const BAND = {
  band: 20,
  faction: 'horde',
  preset: 'raid',
  role: 'dps',
  set_dps: 37.46,
} as unknown as BisBand;
const ROW: SimListRow = {
  sim_id: 's1',
  spec: SPEC,
  dps: 26.4,
  engine_version: '1',
  created_at: '2026-10-09T00:00:00Z',
  title: '',
  band: 20,
  faction: 'horde',
  preset: 'raid',
};

describe('simCardFigures', () => {
  it('shows both figures when the saved sim ran the band entry own setup', () => {
    expect(simCardFigures(ROW, { specKey: SPEC, band: BAND })).toEqual({ nowDps: 26.4, bandDps: 37.46 });
  });

  it.each([
    ['another spec', { spec: 'hunter-survival' }],
    ['another preset', { preset: 'bare' }],
    ['another band', { band: 30 }],
    ['another faction', { faction: 'alliance' }],
    ['no setup recorded', { preset: undefined, band: undefined, faction: undefined }],
  ])('withholds the "now" figure for %s and keeps the band figure', (_label, override) => {
    const row: SimListRow = { ...ROW, ...override };
    expect(simMatchesBand(row, SPEC, BAND)).toBe(false);
    expect(simCardFigures(row, { specKey: SPEC, band: BAND })).toEqual({ nowDps: null, bandDps: 37.46 });
  });

  it('shows no band figure for a tank band', () => {
    const tank = { ...BAND, role: 'tank' } as unknown as BisBand;
    expect(simCardFigures(ROW, { specKey: SPEC, band: tank }).bandDps).toBeNull();
  });

  it('with no band, shows the sim only when it is the hero spec', () => {
    expect(simCardFigures(ROW, { specKey: SPEC, band: undefined })).toEqual({ nowDps: 26.4, bandDps: null });
    expect(simCardFigures(ROW, { specKey: 'hunter-survival', band: undefined })).toEqual({
      nowDps: null,
      bandDps: null,
    });
  });
});
