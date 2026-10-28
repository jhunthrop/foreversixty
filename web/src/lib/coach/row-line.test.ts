import { describe, expect, it } from 'vitest';
import {
  coachModelReview,
  coachNotReviewable,
  coachPreparing,
  coachRulesReview,
  coachTeaser,
} from '../../test-support/coach-fixtures';
import { rowLineFor } from './row-line';

describe('rowLineFor', () => {
  it('Premium: Fix first lead and the finding title with its spell', () => {
    expect(rowLineFor(coachModelReview())).toEqual({
      lead: 'Fix first · ',
      title: 'Bloodthirst: 9 casts lost',
      after: '',
      spellId: 23881,
      upgrade: '',
    });
  });
  it('free: count, the title then "first", and the Premium link as a separate target', () => {
    const line = rowLineFor(coachTeaser());
    expect(line).toEqual({
      lead: '3 things to fix · ',
      title: 'Bloodthirst: 9 casts lost',
      after: ' first',
      spellId: 23881,
      upgrade: 'Read the full review with Premium',
    });
  });
  it('free with one finding reads in the singular', () => {
    expect(rowLineFor(coachTeaser({ finding_count: 1 }))).toMatchObject({
      lead: '1 thing to fix · ',
      after: ' first',
    });
  });
  it('Premium with no Fix first but Then items', () => {
    const base = coachRulesReview();
    const answer = { ...base, review: { ...base.review, fixFirst: null } };
    expect(rowLineFor(answer)?.lead).toBe("1 thing to fix · Open your coach's review");
  });
  it('Premium with nothing at all is the clean line', () => {
    const base = coachRulesReview();
    const answer = { ...base, review: { ...base.review, fixFirst: null, thenFix: [] } };
    expect(rowLineFor(answer)?.lead).toBe('Nothing to fix on this pull that the log can show.');
  });
  it('free with a count and no first title keeps Premium as the gold link', () => {
    const line = rowLineFor(coachTeaser({ fix_first: null }));
    expect(line).toEqual({
      lead: '3 things to fix',
      title: '',
      after: '',
      spellId: null,
      upgrade: 'Read the full review with Premium',
    });
  });
  it('a free owner with no findings shows the no-findings line, not nothing', () => {
    expect(rowLineFor(coachTeaser({ finding_count: 0, fix_first: null }))?.lead).toBe(
      'Nothing to fix on this pull that the log can show.',
    );
  });
  it('preparing names the spec from the answer, else the roster fallback', () => {
    expect(rowLineFor(coachPreparing(), 'Warrior')?.lead).toBe(
      'Top Warrior logs for this boss are still loading.',
    );
    expect(rowLineFor({ ...coachPreparing(), spec: 'Fury' }, 'Warrior')?.lead).toBe(
      'Top Fury logs for this boss are still loading.',
    );
  });
  it('null and not reviewable render nothing', () => {
    expect(rowLineFor(null)).toBeNull();
    expect(rowLineFor(coachNotReviewable())).toBeNull();
  });
});
