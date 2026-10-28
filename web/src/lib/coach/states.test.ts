import { describe, expect, it } from 'vitest';
import {
  coachModelReview,
  coachNotReviewable,
  coachOfficerReview,
  coachPreparing,
  coachRulesReview,
  coachTeaser,
} from '../../test-support/coach-fixtures';
import { CoachApiError } from './api';
import { isPolling, viewForAnswer, viewForFailure, viewOf } from './states';
import type { CoachReview } from './types';

const reviewWith = (patch: Partial<CoachReview>): CoachReview => ({ ...coachModelReview().review, ...patch });

describe('viewForAnswer', () => {
  it('a 404 is not available', () => {
    expect(viewForAnswer(null)).toEqual(viewOf('not_available'));
  });

  it('teasers: findings, none, preparing, not reviewable', () => {
    expect(viewForAnswer(coachTeaser()).state).toBe('teaser');
    expect(viewForAnswer(coachTeaser({ finding_count: 0, fix_first: null })).state).toBe('teaser_empty');
    expect(viewForAnswer(coachPreparing()).state).toBe('preparing');
    expect(viewForAnswer(coachNotReviewable()).state).toBe('not_reviewable');
  });

  it('ready reviews by source', () => {
    expect(viewForAnswer(coachModelReview()).state).toBe('ready_model');
    expect(viewForAnswer(coachRulesReview()).state).toBe('ready_rules');
  });

  it('no findings and raid buffs only come from the review contents', () => {
    expect(
      viewForAnswer(coachModelReview({ review: reviewWith({ fixFirst: null, thenFix: [] }) })).state,
    ).toBe('no_findings');
    expect(viewForAnswer(coachModelReview({ review: reviewWith({ fixFirst: null }) })).state).toBe(
      'raid_buffs_only',
    );
  });

  it('writing and busy are base states', () => {
    expect(viewForAnswer(coachRulesReview({ status: 'writing' })).state).toBe('writing');
    expect(viewForAnswer(coachRulesReview({ busy: true })).state).toBe('busy');
    expect(isPolling(viewForAnswer(coachRulesReview({ status: 'writing' })))).toBe(true);
    expect(isPolling(viewForAnswer(coachRulesReview({ busy: true })))).toBe(false);
  });

  it('officer is an overlay: an officer review that is writing or busy keeps that state', () => {
    const writing = viewForAnswer({ ...coachOfficerReview(), source: 'rules', status: 'writing' });
    expect(writing).toMatchObject({ state: 'writing', officer: true });
    expect(isPolling(writing)).toBe(true);
    expect(viewForAnswer({ ...coachOfficerReview(), busy: true })).toMatchObject({
      state: 'busy',
      officer: true,
    });
    expect(viewForAnswer(coachOfficerReview())).toMatchObject({ state: 'ready_model', officer: true });
  });

  it('starter and healer are overlays too', () => {
    const view = viewForAnswer(
      coachModelReview({ persona: { label: 'x', status: 'starter' }, role: 'healer', status: 'writing' }),
    );
    expect(view).toEqual({ state: 'writing', officer: false, starter: true, healer: true });
  });
});

describe('viewForFailure', () => {
  it('a 401 is signed out', () => {
    expect(viewForFailure(new CoachApiError('x', 401)).state).toBe('signed_out');
  });
  it('a no_plan 403, another 403 code, a 500 and a stray error are handled', () => {
    expect(viewForFailure(new CoachApiError('x', 403, 'no_plan')).state).toBe('error');
    expect(viewForFailure(new CoachApiError('x', 403, 'forbidden')).state).toBe('signed_out');
    expect(viewForFailure(new CoachApiError('x', 500)).state).toBe('error');
    expect(viewForFailure(new Error('x')).state).toBe('error');
  });
});

describe('request and mode states', () => {
  it('loading and whole night carry no overlays', () => {
    expect(viewOf('loading')).toEqual({ state: 'loading', officer: false, starter: false, healer: false });
    expect(viewOf('whole_night').state).toBe('whole_night');
  });
});
