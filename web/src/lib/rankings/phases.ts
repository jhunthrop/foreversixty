// web/src/lib/rankings/phases.ts
// The content phases rankings are bracketed by. The API derives a row's phase from when
// the fight happened (api/internal/phase/phase.go) and stores it, so a percentile lookup
// has to name the same phase the API filed the row under: `phaseAt` mirrors the API's
// boundaries for that one purpose. The list is also what the rankings page offers as a
// filter, and an id the API returns that is not here is shown as-is (phaseLabel echoes
// it) rather than hidden.
//
// Launch is 15:00 PST on Nov 4, which is 23:00 UTC; the other boundaries open at midnight
// UTC on their day. These are the API's values, copied, and the api's phase_test.go and
// phases.test.ts pin the same moments so the two cannot drift silently.
export const PHASES = [
  { id: 'pre-beta', label: 'Pre-beta', start: '0001-01-01T00:00:00Z' },
  { id: 'beta', label: 'Beta', start: '2026-09-17T00:00:00Z' },
  { id: 'launch', label: 'Launch', start: '2026-11-04T23:00:00Z' },
  { id: 'raids-1', label: 'Raids, phase 1', start: '2026-12-09T00:00:00Z' },
] as const;

export type Phase = (typeof PHASES)[number]['id'];

export function phaseLabel(id: string): string {
  return PHASES.find((phase) => phase.id === id)?.label ?? id;
}

/** The phase a moment falls in: the last phase whose start is not after it. */
export function phaseAt(iso: string): Phase {
  const at = Date.parse(iso);
  let found: Phase = PHASES[0].id;
  if (Number.isNaN(at)) return found;
  for (const phase of PHASES) {
    if (Date.parse(phase.start) <= at) found = phase.id;
  }
  return found;
}
