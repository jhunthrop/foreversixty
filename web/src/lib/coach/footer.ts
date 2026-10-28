// web/src/lib/coach/footer.ts
// The review's evidence footer (experience spec 3.1 and 8.2): `Reviewed {date} · Based on: {labels}.`,
// or `Reviewed {date}.` when the pack names no basis (a pull with nothing to fix).
import { coachCopy, fill } from './copy';
import { reviewedDate } from './dates';

/** The basis labels are sentences: one list, each without its own full stop, `; ` between. */
const basisList = (labels: readonly string[]): string =>
  labels.map((label) => label.replace(/\.$/, '')).join('; ');

export function footerLine(reviewedAt: string, labels: readonly string[]): string {
  const date = reviewedDate(reviewedAt);
  const named = labels.filter((label) => label.trim() !== '');
  if (named.length === 0) return fill(coachCopy.footerNoBasis, { date });
  return fill(coachCopy.footer, { date, labels: basisList(named) });
}
