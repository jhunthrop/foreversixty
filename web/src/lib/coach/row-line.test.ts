import { describe, expect, it } from 'vitest';
import {
  coachModelReview,
  coachNotReviewable,
  coachRulesReview,
  coachTeaser,
} from '../../test-support/coach-fixtures';
import { rowLineFor } from './row-line';

describe('rowLineFor', () => {
  it('Premium: Fix first lead and the finding title with its spell', () => {
    expect(rowLineFor(coachModelReview())).toEqual({
      lead: 'Fix first · ',
      title: 'Bloodthirst: 9 casts lost',
      spellId: 23881,
      upgrade: '',
    });
  });
  it('free: count, first title, and the Premium link as a separate target', () => {
    const line = rowLineFor(coachTeaser());
    expect(line).toEqual({
      lead: '3 things to fix · First: ',
      title: 'Bloodthirst: 9 casts lost',
      spellId: 23881,
      upgrade: 'Read the full review with Premium',
    });
  });
  it('free with one finding reads in the singular', () => {
    expect(rowLineFor(coachTeaser({ finding_count: 1 }))?.lead).toBe('1 thing to fix · First: ');
  });
  it('free with a count and no first title', () => {
    expect(rowLineFor(coachTeaser({ fix_first: null }))?.lead).toBe(
      '3 things to fix · Read the full review with Premium',
    );
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
  it('null, zero findings and not reviewable render nothing', () => {
    expect(rowLineFor(null)).toBeNull();
    expect(rowLineFor(coachTeaser({ finding_count: 0, fix_first: null }))).toBeNull();
    expect(rowLineFor(coachNotReviewable())).toBeNull();
  });
});
