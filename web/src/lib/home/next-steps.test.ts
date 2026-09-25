import { describe, expect, it } from 'vitest';
import { logsCardLine, plannerPointsLabel, ratingCardValue, simCardLine } from './next-steps';
import type { SimListRow } from '../sim/types';
import type { MyReport } from '../account/api';
import type { CharacterRating, RatingCardPlayer } from '../rating/types';

describe('simCardLine', () => {
  it('reads the kind label and the API headline', () => {
    const row: SimListRow = {
      sim_id: 's1',
      spec: 'fury-warrior',
      dps: 842,
      engine_version: '1',
      created_at: '2026-09-20T00:00:00Z',
      title: '',
      kind: 'gear',
      headline: "+41 DPS from Vis'kag",
    };
    expect(simCardLine(row)).toBe("Top Gear · +41 DPS from Vis'kag");
  });

  it('falls back to the DPS figure when the row has no headline', () => {
    const row: SimListRow = {
      sim_id: 's2',
      spec: 'fury-warrior',
      dps: 1204,
      engine_version: '1',
      created_at: '2026-09-20T00:00:00Z',
      title: '',
      kind: 'run',
    };
    expect(simCardLine(row)).toBe('Sim · 1,204 DPS');
  });
});

describe('plannerPointsLabel', () => {
  it('sums a talent split string against MAX_POINTS', () => {
    // "31/0/20" is SimInput.talents' own shape (fight_metrics.talent_split): points per
    // tree, slash-separated. 31 + 0 + 20 = 51.
    expect(plannerPointsLabel('31/0/20')).toBe('51 of 51 points');
  });
  it('sums a split with only one non-zero tree', () => {
    // Distinguishes the '/'-separated tree-total parse from the unrelated dash-separated,
    // per-character digit parse `lib/sim/character.ts`'s `talentPointsFromString` uses for
    // an FS1-decoded build's own talent order: a value like '10' would sum to 1 under that
    // parse (each character read as a single-digit rank) but sums to 10 under this one.
    expect(plannerPointsLabel('10/0/0')).toBe('10 of 51 points');
  });
  it('returns empty for no recorded talents', () => {
    expect(plannerPointsLabel('')).toBe('');
  });
});

describe('logsCardLine', () => {
  it('joins the title and a formatted date', () => {
    const report: MyReport = {
      id: 'r1',
      title: 'Barrow Deeps 9/20',
      zone: 'Barrow Deeps',
      status: 'done',
      visibility: 'public',
      created_at: '2026-09-20T00:00:00Z',
      fight_count: 6,
      kill_count: 5,
    };
    expect(logsCardLine(report)).toBe('Barrow Deeps 9/20, Sept 20');
  });
});

describe('ratingCardValue', () => {
  it('is empty with no rating', () => {
    expect(ratingCardValue(null)).toBe('');
  });

  it('is empty with a zero sample size', () => {
    const rating: CharacterRating = {
      player_key: 'k',
      sample_size: 0,
      trend: [],
      best_component: '',
      worst_component: '',
      latest: null,
    };
    expect(ratingCardValue(rating)).toBe('');
  });

  it('formats the latest overall to two decimals', () => {
    const latest: RatingCardPlayer = {
      player_key: 'k',
      player_name: 'Grommash',
      class: 'warrior',
      spec: 'fury-warrior',
      role: 'melee',
      // 1.076, not 1.075: (1.075).toFixed(2) is '1.07' in JS -- 1.075 is not exactly
      // representable in IEEE 754 double and rounds down -- so this picks a value that
      // rounds up unambiguously and still exercises the two-decimal formatting.
      overall: 1.076,
      overall_uncapped: 1.076,
      overall_capped: false,
      coverage: 1,
      insufficient: false,
      insufficient_reason: '',
      basis: 'percentile',
      components: [],
    };
    const rating: CharacterRating = {
      player_key: 'k',
      sample_size: 4,
      trend: [],
      best_component: '',
      worst_component: '',
      latest,
    };
    expect(ratingCardValue(rating)).toBe('1.08');
  });
});
