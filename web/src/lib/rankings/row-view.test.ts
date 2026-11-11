// web/src/lib/rankings/row-view.test.ts
import { describe, expect, it } from 'vitest';
import type { RankingRow } from './api';
import {
  RANK_LEGEND,
  moderationLine,
  ownRowPlacement,
  rankColor,
  rankPercentile,
  rankingDate,
  showRowCrests,
  withoutIdleRows,
} from './row-view';

const row = (key: string, cls = 'Warrior'): Pick<RankingRow, 'player'> => ({
  player: { key, name: key, class: cls, spec: 'Fury' },
});
const placed = (key: string, rank = 14) => ({ rank, row: row(key) as RankingRow });

describe('the pin rule', () => {
  const page = [row('us/normal/a'), row('us/normal/me'), row('us/normal/b')];

  it('marks the row in place when it is on the visible page', () => {
    expect(ownRowPlacement(page, placed('us/normal/me'))).toBe('in-place');
  });

  it('pins the row only when it is off the visible page', () => {
    expect(ownRowPlacement(page, placed('us/normal/far-away', 204))).toBe('pinned');
  });

  it('places nothing without a rank, or when the character has none on this board', () => {
    expect(ownRowPlacement(page, undefined)).toBe('none');
    expect(ownRowPlacement(page, null)).toBe('none');
  });
});

describe('the crest rule', () => {
  it('shows no crest on a board of one class', () => {
    expect(showRowCrests([row('a'), row('b'), row('c')])).toBe(false);
  });

  it('shows a crest on every row once the board spans more than one class', () => {
    expect(showRowCrests([row('a'), row('b', 'Priest')])).toBe(true);
  });

  it('shows none on an empty board', () => {
    expect(showRowCrests([])).toBe(false);
  });
});

describe('the colour legend', () => {
  it('runs best first through the seven steps, each in the colour its rank numeral takes', () => {
    expect(RANK_LEGEND.map((step) => step.label)).toEqual([
      '100',
      '99+',
      '95+',
      '75+',
      '50+',
      '25+',
      'under 25',
    ]);
    expect(RANK_LEGEND.map((step) => step.color)).toEqual([
      'var(--color-parse-gold)',
      'var(--color-parse-pink)',
      'var(--color-parse-orange)',
      'var(--color-rarity-epic-text)',
      'var(--color-rarity-rare-text)',
      'var(--color-parse-green)',
      'var(--color-parse-grey-text)',
    ]);
  });

  it('colours a rank from its percentile in the board', () => {
    expect(rankPercentile(1, 200)).toBe(100);
    expect(rankPercentile(101, 200)).toBe(50);
    expect(rankColor(14, 212)).toBe('var(--color-rarity-epic-text)');
    // The WCL scale, floored (spec 2026-11-07 MUST 3, bug B4): rank 1 of 212 is the 100th
    // percentile and gold; ranks 2 and 3 (99.53, 99.06) are pink, never rounded up to gold;
    // rank 4 (98.58) is orange.
    expect(rankColor(1, 212)).toBe('var(--color-parse-gold)');
    expect(rankColor(2, 212)).toBe('var(--color-parse-pink)');
    expect(rankColor(3, 212)).toBe('var(--color-parse-pink)');
    expect(rankColor(4, 212)).toBe('var(--color-parse-orange)');
    expect(rankColor(204, 212)).toBe('var(--color-parse-grey-text)');
  });
});

describe('the row date', () => {
  it('prints the kill day', () => {
    expect(rankingDate('2026-12-10T12:00:00Z')).toBe('10 Dec');
    expect(rankingDate('2026-11-25T12:00:00Z')).toBe('25 Nov');
  });
});

describe('the moderation line', () => {
  it('is one plain line for a touched kill and nothing for a normal one', () => {
    expect(moderationLine('removed')).toBe('Removed by a moderator');
    expect(moderationLine('at_risk')).toBe('Under review by a moderator');
    expect(moderationLine('ok')).toBeNull();
  });
});

describe('withoutIdleRows', () => {
  const rows = [{ value: 120 }, { value: 80 }, { value: 0 }] as RankingRow[];
  it('drops a 0 DPS or HPS row and counts it out of the total', () => {
    expect(withoutIdleRows({ rows, total: 3 }, 'dps')).toEqual({ rows: rows.slice(0, 2), total: 2 });
    expect(withoutIdleRows({ rows, total: 3 }, 'hps').total).toBe(2);
  });
  it('leaves the other metrics alone', () => {
    expect(withoutIdleRows({ rows, total: 3 }, 'damage_taken').rows).toHaveLength(3);
  });
});
