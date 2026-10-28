// web/src/lib/coach/states.ts
// Which panel state (experience spec section 4.2) an answer puts the Coach tab in. One pure
// function so the panel, the row strip and the tests agree. The spec treats Officer, Starter and
// the healer persona as overlays on a review, not states of their own: an officer's review that
// is still writing must keep polling. So a view is a base `state` (what the panel does: poll,
// wait, show) plus the overlay flags (what the header and notes say).
import { CoachApiError } from './api';
import type { CoachFull, CoachTeaser } from './types';

export type CoachPanelState =
  | 'loading'
  | 'error'
  | 'not_available'
  | 'signed_out'
  | 'whole_night'
  | 'teaser'
  | 'teaser_empty'
  | 'preparing'
  | 'not_reviewable'
  | 'writing'
  | 'busy'
  | 'no_findings'
  | 'raid_buffs_only'
  | 'ready_model'
  | 'ready_rules';

export interface CoachPanelView {
  readonly state: CoachPanelState;
  /** The review is the officer's view of another player's: title, pill and note change. */
  readonly officer: boolean;
  /** The persona is the early version: title and hint change. */
  readonly starter: boolean;
  /** The reviewed player heals: the healer rows (overheal) apply. */
  readonly healer: boolean;
}

const SIGNED_OUT_STATUSES: readonly number[] = [401, 403];
const NO_OVERLAYS = { officer: false, starter: false, healer: false } as const;

/** A view with no overlays: the states a request or the mode decides. */
export const viewOf = (state: CoachPanelState): CoachPanelView => ({ state, ...NO_OVERLAYS });

/** The states that poll (writing) or wait (busy) keep their base state under every overlay. */
export const isPolling = (view: CoachPanelView): boolean => view.state === 'writing';

function teaserState(answer: CoachTeaser): CoachPanelState {
  if (answer.status === 'preparing') return 'preparing';
  if (answer.status === 'not_reviewable') return 'not_reviewable';
  return answer.finding_count === 0 ? 'teaser_empty' : 'teaser';
}

function reviewState(answer: CoachFull): CoachPanelState {
  if (answer.status === 'preparing') return 'preparing';
  if (answer.status === 'not_reviewable') return 'not_reviewable';
  if (answer.busy) return 'busy';
  if (answer.status === 'writing') return 'writing';
  if (answer.review.fixFirst === null) {
    return answer.review.thenFix.length === 0 ? 'no_findings' : 'raid_buffs_only';
  }
  return answer.source === 'model' ? 'ready_model' : 'ready_rules';
}

/** The view for a settled answer: `null` is a 404. */
export function viewForAnswer(answer: CoachFull | CoachTeaser | null): CoachPanelView {
  if (answer === null) return viewOf('not_available');
  if (answer.access === 'teaser') return viewOf(teaserState(answer));
  return {
    state: reviewState(answer),
    officer: answer.access === 'officer',
    starter: answer.persona.status === 'starter',
    healer: answer.role === 'healer',
  };
}

/** The view for a failed request: a missing session reads as signed out, the rest as an error. */
export function viewForFailure(error: unknown): CoachPanelView {
  const signedOut =
    error instanceof CoachApiError && SIGNED_OUT_STATUSES.includes(error.status) && error.code !== 'no_plan';
  return viewOf(signedOut ? 'signed_out' : 'error');
}
