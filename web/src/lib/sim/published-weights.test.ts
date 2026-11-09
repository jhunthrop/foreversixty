// web/src/lib/sim/published-weights.test.ts
import { describe, expect, it } from 'vitest';
import type { BisFile } from '../bis/types';
import { bulkCopy } from './copy';
import { publishedWeights } from './published-weights';

const weight = (stat: string, value: number, scale: number) => ({
  stat,
  weight: value,
  error: 0.01,
  insignificant: false,
  scale_factor: scale,
  dps_per_point: value * 0.3,
  scale_error: 0.001,
});
const band = (faction: string, preset: string) => ({
  band: 60,
  faction,
  preset,
  spec: 'warrior-fury',
  talents: '',
  weights: [weight('strength', 2.2, 1), weight('attack_power', 1, 0.45)],
  scale_reference_stat: 'strength',
  reference_dps_per_point: 0.3,
});
const file = {
  spec: 'warrior-fury',
  bands: [band('horde', 'raid'), band('horde', 'bare')],
} as unknown as BisFile;

describe('publishedWeights', () => {
  it('reads the character’s band, faction and preset, top stat first, linking to that set', () => {
    const got = publishedWeights(file, { spec: 'warrior-fury', race: 'orc', level: 60, preset: 'raid' });
    expect(got?.view.rows.map((row) => row.stat)).toEqual(['strength', 'attack_power']);
    expect(got?.view.rows[0]?.scaleFactor).toBe(1);
    expect(got?.label).toBe(bulkCopy.weightsPublishedBand(60, 'raid'));
    expect(got?.href).toContain('faction=horde');
  });

  it('is null with no band for the faction or no known race', () => {
    expect(
      publishedWeights(file, { spec: 'warrior-fury', race: 'human', level: 60, preset: 'raid' }),
    ).toBeNull();
    expect(
      publishedWeights(file, { spec: 'warrior-fury', race: 'murloc', level: 60, preset: 'raid' }),
    ).toBeNull();
  });
});
