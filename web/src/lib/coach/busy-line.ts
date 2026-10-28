// web/src/lib/coach/busy-line.ts
// The busy line (spec 3.1, 8.2): the account's own cap when the answer names one, else the
// site's busy line.
import { coachCopy, fill } from './copy';
import { utcDayDate } from './dates';
import type { CoachFull } from './types';

/** The line a busy answer shows above the rule-only review. */
export function busyLine(answer: Pick<CoachFull, 'cap'>): string {
  if (answer.cap?.period === 'day') return coachCopy.capAccountDay;
  if (answer.cap?.period === 'month') {
    return fill(coachCopy.capAccountMonth, { date: utcDayDate(answer.cap.until) });
  }
  return coachCopy.busy;
}
