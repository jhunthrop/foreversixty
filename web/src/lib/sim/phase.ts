// web/src/lib/sim/phase.ts
// The content phases, and whether a loot source has opened yet.
//
// Two sources, in this order (contract 10.4 and 10.6):
//   * GET /v1/phases at runtime, so a date that moves after a deploy is picked up without
//     a rebuild;
//   * web/src/data/phases.json at build time, which the pipeline emits from
//     data/curated/phases.json -- the file api/internal/phase.Boundaries is tested
//     against -- as the fallback when the API is unreachable.
//
// The data file is `[{ "name", "start" }]` and carries no label: a label is display copy
// and lives here, in DISPLAY_LABELS, beside the rest of this lane's words.
import builtIn from '../../data/phases.json';
import { fetchPhases as fetchPhasesFromApi } from './api';

export interface PhaseRow {
  /** The stored, URL-safe name: pre-beta | beta | launch | raids-1. */
  name: string;
  /** RFC3339. api/internal/phase's zero `time.Time` marshals as "0001-01-01T00:00:00Z" for
   *  the phase with no opening instant (pre-beta) -- see ZERO_TIME below, which this module
   *  treats the same as "no start" everywhere a date would otherwise be displayed. */
  start: string;
}

/** The build-time table, and the fallback when the API cannot be reached. */
export const BUILT_IN_PHASES: readonly PhaseRow[] = builtIn as PhaseRow[];

/**
 * Contract 10.4: "a source whose date is unknown carries `opens: "later"`, which the page
 * shows as unreleased without a date." It is never a row in the phase table, and it never
 * opens.
 */
export const PHASE_LATER = 'later';

/** api/internal/phase's zero time, the wire value for "this phase has no opening instant". */
const ZERO_TIME = '0001-01-01T00:00:00Z';

/** True for a start that is a real instant -- neither "" nor the Go zero time. */
function hasStart(start: string): boolean {
  return start !== '' && start !== ZERO_TIME;
}

/** A start's millis, or -Infinity for "no start" -- always the earliest possible boundary. */
function startMillis(start: string): number {
  return hasStart(start) ? Date.parse(start) : Number.NEGATIVE_INFINITY;
}

const DISPLAY_LABELS: Record<string, string> = {
  'pre-beta': 'Before beta',
  beta: 'Beta',
  launch: 'Launch',
  'raids-1': 'First raids',
  [PHASE_LATER]: 'Later',
};

/**
 * The live table, or the build-time one. Every failure -- an unreachable API, a 500, a
 * body that is not the expected shape -- falls back rather than throwing: a phase gate
 * that cannot answer must not take the Droptimizer down with it. The single call site for
 * `api.ts`'s `fetchPhases`; every other module reads the phase table through this one.
 */
export async function fetchPhases(apiBase?: string): Promise<readonly PhaseRow[]> {
  try {
    const rows = await fetchPhasesFromApi(apiBase);
    return Array.isArray(rows) && rows.length > 0 ? rows : BUILT_IN_PHASES;
  } catch {
    return BUILT_IN_PHASES;
  }
}

/** The phase a moment falls in. Later boundaries win, so the last match is the answer. */
export function phaseAt(phases: readonly PhaseRow[], when: Date): string {
  let name = phases[0]?.name ?? '';
  for (const phase of phases) {
    if (when.getTime() >= startMillis(phase.start)) name = phase.name;
  }
  return name;
}

function rowOf(phases: readonly PhaseRow[], name: string): PhaseRow | undefined {
  return phases.find((phase) => phase.name === name);
}

/** What a player calls it. Display copy, not data -- the table carries no label column. */
export function phaseLabel(name: string): string {
  return DISPLAY_LABELS[name] ?? name;
}

/** When the phase opens, or null for the one with no start, for "later", and for an unknown. */
export function phaseStart(phases: readonly PhaseRow[], name: string): Date | null {
  const row = rowOf(phases, name);
  if (row === undefined || !hasStart(row.start)) return null;
  return new Date(row.start);
}

/**
 * Whether content gated on this phase is available at `when`.
 *
 * An absent phase is open: contract 6.1 says "a source without it is open from launch".
 * The literal "later" is never open (contract 10.4). A phase this build has never heard of
 * is NOT open either -- a loot file naming a phase the site does not carry is ahead of the
 * site, and showing its raid as live would be a lie in the one direction that matters.
 */
export function hasOpened(phases: readonly PhaseRow[], name: string | undefined, when: Date): boolean {
  if (name === undefined || name === '') return true;
  if (name === PHASE_LATER) return false;
  const row = rowOf(phases, name);
  if (row === undefined) return false;
  if (!hasStart(row.start)) return true;
  return when.getTime() >= Date.parse(row.start);
}

const DATE_FORMAT = new Intl.DateTimeFormat('en-GB', {
  day: 'numeric',
  month: 'long',
  year: 'numeric',
  timeZone: 'UTC',
});

/** "9 December 2026", or "" for a phase with no start, for "later", and for an unknown. */
export function openDateLabel(phases: readonly PhaseRow[], name: string): string {
  const start = phaseStart(phases, name);
  return start === null ? '' : DATE_FORMAT.format(start);
}
