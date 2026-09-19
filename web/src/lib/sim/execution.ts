// web/src/lib/sim/execution.ts
// "92% of what your gear can do", beside the parse percentile on every ranked fight.
//
// The href is composed here rather than through lib/sim/url.ts's simSearch on purpose: this
// module ships with the rankings pages, which are a different group of this plan from the
// sim page and must be buildable without it. Two query parameters are not worth a
// dependency that would order two independent groups.
//
// The column is deliberately not coloured. A percentile already carries a colour token, and
// a second coloured number in the same row would read as a competing ranking rather than a
// different question. The design system also forbids colour-only signalling, so the state
// this column most often shows -- no score at all, which is every unvalidated spec -- is a
// dash with a sentence, not a grey cell.
// The two sentences live in copy.ts with every other user-visible string on this lane; the
// arithmetic lives here. Step 1a appends them.
import { simCopy } from './copy';

export function executionLabel(score: number | null): string {
  return score === null ? '—' : `${Math.round(score * 100)}%`;
}

export function executionTitle(score: number | null): string {
  return score === null ? simCopy.executionUnscored : simCopy.executionScored(executionLabel(score));
}

/** Compare mode on that fight: the design's "compare mode one click away". */
export function executionHref(reportId: string, fightIndex: number): string {
  const params = new URLSearchParams({
    source: 'fight',
    ref: `${reportId}:${fightIndex}`,
    mode: 'compare',
  });
  return `/sim?${params.toString()}`;
}
