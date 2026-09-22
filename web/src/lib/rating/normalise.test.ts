import { describe, expect, it } from 'vitest';
import { normaliseReportRatings } from './normalise';
import type { ReportRatings } from './types';

// The API writes `percentile`, `bracket_n` and `reason` with omitempty, so a component that
// is absolute-based, or excluded for a reason the API left blank, arrives without those keys.
// The page's checks are against null and 0, and `undefined !== null` is true: an absolute
// component would have read "undefined percentile". The fetch boundary fills the gaps.
describe('normaliseReportRatings', () => {
  it('fills omitted per-component fields with the values the page checks against', () => {
    const wire = {
      fight_index: 1,
      kill: true,
      kill_time_band: 'fast',
      model_version: 'v1',
      players: [
        {
          player_key: 'us:normal:bob',
          player_name: 'Bob',
          class: 'warrior',
          spec: 'fury',
          role: 'dps',
          overall: 80,
          overall_uncapped: 80,
          overall_capped: false,
          basis: 'mixed',
          components: [
            { name: 'output', score: 90, weight: 40, basis: 'absolute', excluded: false, moments: [] },
          ],
        },
      ],
    } as unknown as ReportRatings;
    const part = normaliseReportRatings(wire).players[0].components[0];
    expect(part.percentile).toBeNull();
    expect(part.bracket_n).toBe(0);
    expect(part.reason).toBe('');
    expect(part.moments).toEqual([]);
  });

  it('leaves present fields alone and tolerates a missing players list', () => {
    const wire = {
      fight_index: 1,
      kill: false,
      kill_time_band: '',
      model_version: 'v1',
    } as unknown as ReportRatings;
    expect(normaliseReportRatings(wire).players).toEqual([]);
  });
});

describe('insufficient cards', () => {
  it('reads the coverage fields and defaults them for an older API', () => {
    const wire = {
      fight_index: 1,
      kill: false,
      kill_time_band: '',
      model_version: 'v1',
      players: [
        {
          player_key: 'k',
          player_name: 'Bob',
          class: 'priest',
          spec: 'shadow',
          role: 'dps',
          overall: 0,
          overall_uncapped: 0,
          overall_capped: false,
          basis: '',
          coverage: 0.3,
          insufficient: true,
          insufficient_reason: 'wipe; no mechanics table',
          components: [],
        },
        {
          player_key: 'k2',
          player_name: 'Ann',
          class: 'mage',
          spec: 'frost',
          role: 'dps',
          overall: 80,
          overall_uncapped: 80,
          overall_capped: false,
          basis: 'mixed',
          components: [],
        },
      ],
    } as unknown as ReportRatings;
    const [bob, ann] = normaliseReportRatings(wire).players;
    expect(bob.insufficient).toBe(true);
    expect(bob.coverage).toBe(0.3);
    expect(ann.insufficient).toBe(false);
    expect(ann.coverage).toBe(1);
    expect(ann.insufficient_reason).toBe('');
  });
});
