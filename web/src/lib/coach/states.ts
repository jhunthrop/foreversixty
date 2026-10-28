// web/src/lib/coach/states.ts
// Which panel state (experience spec section 4.2) an answer puts the Coach tab in. One pure
// function so the panel, the row strip and the tests agree. The states that come from the
// request itself (loading, error, not available, signed out, whole night) are named here too.
import { CoachApiError } from './api';
import type { CoachFull, CoachTeaser } from './types';

export type CoachPanelState =
  | 'loading'
  | 'error'
  | 'not_available'
  | 'signed_out'
  | 'teaser'
  | 'preparing'
  | 'not_reviewable'
  | 'writing'
  | 'busy'
  | 'officer'
  | 'starter'
  | 'ready_model'
  | 'ready_rules';

const SIGNED_OUT_STATUSES: readonly number[] = [401, 403];

/** The state for a settled answer: `null` is a 404. */
export function stateForAnswer(answer: CoachFull | CoachTeaser | null): CoachPanelState {
  if (answer === null) return 'not_available';
  if (answer.access === 'teaser') return answer.status === 'ready' ? 'teaser' : statusState(answer.status);
  if (answer.status === 'preparing' || answer.status === 'not_reviewable') return statusState(answer.status);
  if (answer.access === 'officer') return 'officer';
  if (answer.busy) return 'busy';
  if (answer.status === 'writing') return 'writing';
  if (answer.persona.status === 'starter') return 'starter';
  return answer.source === 'model' ? 'ready_model' : 'ready_rules';
}

function statusState(status: CoachFull['status']): CoachPanelState {
  if (status === 'preparing') return 'preparing';
  if (status === 'not_reviewable') return 'not_reviewable';
  return 'teaser';
}

/** The state for a failed request: a missing session reads as signed out, the rest as an error. */
export function stateForFailure(error: unknown): CoachPanelState {
  if (
    error instanceof CoachApiError &&
    SIGNED_OUT_STATUSES.includes(error.status) &&
    error.code !== 'no_plan'
  ) {
    return 'signed_out';
  }
  return 'error';
}
