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
  it('runs best first through the six steps, each in the colour its rank numeral takes', () => {
    expect(RANK_LEGEND.map((step) => step.label)).toEqual(['99+', '95+', '75+', '50+', '25+', 'under 25']);
    expect(RANK_LEGEND.map((step) => step.color)).toEqual([
      'var(--color-parse-pink)',
      'var(--color-parse-orange)',
      'var(--color-parse-purple)',
      'var(--color-parse-blue)',
      'var(--color-parse-green)',
      'var(--color-parse-grey-text)',
    ]);
  });

  it('colours a rank from its percentile in the board', () => {
    expect(rankPercentile(1, 200)).toBe(100);
    expect(rankPercentile(101, 200)).toBe(50);
    expect(rankColor(14, 212)).toBe('var(--color-parse-purple)');
    // The first rank is a 100th percentile but the legend ends at 99+, so it is pink, not gold.
    expect(rankColor(1, 212)).toBe('var(--color-parse-pink)');
    expect(rankColor(204, 212)).toBe('var(--color-parse-grey-text)');
  });
});

describe('the row date', () => {
  const now = new Date('2026-12-13T12:00:00Z');

  it('reads relative inside two weeks', () => {
    expect(rankingDate('2026-12-10T12:00:00Z', now)).toBe('3 days ago');
    expect(rankingDate('2026-12-13T09:00:00Z', now)).toBe('3 hours ago');
  });

  it('prints the date from two weeks on', () => {
    expect(rankingDate('2026-11-25T12:00:00Z', now)).toBe('Nov 25');
  });
});

describe('the moderation line', () => {
  it('is one plain line for a touched kill and nothing for a normal one', () => {
    expect(moderationLine('removed')).toBe('Removed by a moderator');
    expect(moderationLine('at_risk')).toBe('Under review by a moderator');
    expect(moderationLine('ok')).toBeNull();
  });
});
