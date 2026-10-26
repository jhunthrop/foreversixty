// web/src/lib/guild/match-picker.svelte.ts
// The match picker's state (design/specs/2026-10-27-guild-attendance.md 3.2, 4): the candidate
// list's load, the one selected log (the `Closest` one when it is free), the rows another raid
// holds (R-A7), the `Match this log` write and the 409 `report_matched` race. Network calls come
// in through `MatchPickerDeps`, so the tests drive it with fakes.
import type { MatchCandidate, MatchCandidates } from './attendance-contract';
import { PICKER_MAX_ROWS } from './attendance-copy';
import { segmentKeyAction, type SegmentKeyAction } from './segments-keys';

/** The API's refusal when another raid already holds the log. */
export const REPORT_MATCHED = 'report_matched';
/** The API's refusal once the guild's officer tools have ended (officer-tools gate spec 4). */
export const PLAN_REQUIRED = 'guild_plan_required';

/** The answer of an officer write on the match: done, or refused with the line to show. */
export type MatchOutcome = { ok: true } | { ok: false; code: string; message: string };

export interface MatchPickerDeps {
  load: () => Promise<MatchCandidates>;
  /** Writes the match; answers instead of throwing, the page owns the error copy. */
  match: (reportId: string) => Promise<MatchOutcome>;
}

export type MatchPickerLoad = 'loading' | 'ready' | 'failed';

export interface MatchPickerSeed {
  load: MatchPickerLoad;
  rows: MatchCandidate[];
  zoneKnown: boolean;
  selected: string | null;
  busy: boolean;
  line: string | null;
  taken: string[];
}

/** Arrows move focus across the free rows (wrapping), Space and Enter select: the sign-up
 *  segments' model over the rows that can be chosen. Anything else is not the group's key. */
export type PickerKeyAction = { kind: SegmentKeyAction['kind']; reportId: string };

export function pickerKeyAction(
  key: string,
  reportId: string,
  freeIds: readonly string[],
): PickerKeyAction | null {
  const position = freeIds.indexOf(reportId);
  if (position < 0) return null;
  const action = segmentKeyAction(key, position, freeIds.length);
  return action === null ? null : { kind: action.kind, reportId: freeIds[action.index] };
}

export class MatchPicker {
  load = $state<MatchPickerLoad>('loading');
  rows = $state<MatchCandidate[]>([]);
  zoneKnown = $state(true);
  selected = $state<string | null>(null);
  busy = $state(false);
  /** The 20 px status line under the actions. */
  line = $state<string | null>(null);
  /** Logs another raid took while the picker was open. */
  private taken = $state<string[]>([]);

  private readonly deps: MatchPickerDeps;

  constructor(deps: MatchPickerDeps, seed: Partial<MatchPickerSeed> = {}) {
    this.deps = deps;
    Object.assign(this, seed);
  }

  isTaken(row: MatchCandidate): boolean {
    return row.matched_raid !== null || this.taken.includes(row.report_id);
  }

  /** The rows that can be chosen, in list order. */
  get freeIds(): string[] {
    return this.rows.filter((row) => !this.isTaken(row)).map((row) => row.report_id);
  }

  get canMatch(): boolean {
    return !this.busy && this.selected !== null && this.freeIds.includes(this.selected);
  }

  /** Reads the candidates; `quiet` keeps the rows on screen while it refreshes them. */
  async start(quiet = false): Promise<void> {
    if (!quiet) this.load = 'loading';
    try {
      const page = await this.deps.load();
      this.zoneKnown = page.zone_known;
      this.rows = page.rows.slice(0, PICKER_MAX_ROWS);
      this.load = 'ready';
      if (!quiet) this.selected = this.preselected();
    } catch {
      // A quiet re-read that fails leaves the rows as they were; the taken row is already disabled.
      if (!quiet) this.load = 'failed';
    }
  }

  /** `Closest` is chosen up front; when another raid holds it, the first free row is. */
  private preselected(): string | null {
    const closest = this.rows.find((row) => row.closest && !this.isTaken(row));
    return closest?.report_id ?? this.freeIds[0] ?? null;
  }

  select(reportId: string): void {
    if (this.freeIds.includes(reportId)) this.selected = reportId;
  }

  /** Presses `Match this log`; true once the log is matched. */
  async submit(): Promise<boolean> {
    const reportId = this.selected;
    if (reportId === null || !this.canMatch) return false;
    this.busy = true;
    this.line = null;
    try {
      const outcome = await this.deps.match(reportId);
      if (outcome.ok) return true;
      this.line = outcome.message;
      if (outcome.code === REPORT_MATCHED) await this.lostRace(reportId);
      return false;
    } finally {
      this.busy = false;
    }
  }

  /** The row turns matched-elsewhere; a quiet re-read brings the raid it went to. */
  private async lostRace(reportId: string): Promise<void> {
    this.taken = [...this.taken, reportId];
    this.selected = null;
    await this.start(true);
  }
}
