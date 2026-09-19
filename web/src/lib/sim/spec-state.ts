// web/src/lib/sim/spec-state.ts
// How a spec's fidelity reads, and the rule that the support page always shows every spec.
//
// The grid is always the canonical list, never the API's answer: one card per dps spec, in
// rankings-population order, so a spec with no row reads "Not yet" rather than being absent.
// An absent card is indistinguishable from a page that failed to load half its content, and
// "can I trust this" is the question this page exists to answer.
import { simCopy } from './copy';
import { dpsSpecs } from './spec-label';
import type { SpecFidelity, SpecState } from './types';

// The six sentences live in copy.ts with every other user-visible string on this lane; the
// mapping from state to sentence lives here, because it is logic and not copy.
const LABELS: Record<SpecState, string> = {
  validated: simCopy.specValidated,
  in_progress: simCopy.specInProgress,
  unsupported: simCopy.specNotYet,
};

const NOTES: Record<SpecState, string> = {
  validated: simCopy.specValidatedNote,
  in_progress: simCopy.specInProgressNote,
  unsupported: simCopy.specNotYetNote,
};

export function specStateLabel(state: SpecState): string {
  return LABELS[state];
}

export function specStateNote(state: SpecState): string {
  return NOTES[state];
}

export function specPillClass(state: SpecState): string {
  return state === 'validated' ? 'pill pill-site' : 'pill pill-sample';
}

export function mergeSpecRows(rows: readonly SpecFidelity[]): SpecFidelity[] {
  const bySpec = new Map(rows.map((row) => [row.spec, row]));
  return dpsSpecs().map(
    (entry) =>
      bySpec.get(entry.spec) ?? {
        spec: entry.spec,
        state: 'unsupported' as const,
        median_gap: null,
        parses: 0,
        worst_actions: [],
        engine_version: '',
        updated_at: null,
      },
  );
}
