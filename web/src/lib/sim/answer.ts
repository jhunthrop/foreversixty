// web/src/lib/sim/answer.ts
// The head's answer on /sim in every state (spec 2026-11-06-sim §4.3 to §4.5, §5, §5.1), and the
// run line's one primary Button for the same state. One pure function, so every state is a unit
// test and the page only draws the model.
//
// One ± everywhere (ruling S2): the 95% band, 1.96 × the standard error, labelled "(95%)" once,
// in the answer. Data gaps stay hidden (G1, G15): no "since your last sim" line and no last sim
// on the landing until a saved-sim row carries a character key and its error. The server lane
// reports no failure reason (G8 and the job reason are not on the progress read), so a server
// failure reads the one "failed" variant of §5.1.

import { confidenceBand, formatMargin } from './estimate';
import { simPageCopy as copy } from './page-copy';
import { grouped, oneDecimal } from './result-copy';
import type { SimPhase } from './store.svelte';
import type { Estimate, SimResult } from './types';

export type AnswerTone = 'gold' | 'strong' | 'muted';

export type PrimaryAction =
  'run' | 'run-again' | 'try-again' | 'stop' | 'loading-engine' | 'server' | 'disabled';

export interface AnswerModel {
  /** Which state, for the page's test hooks and its actions. */
  state:
    | 'none'
    | 'not-simulated'
    | 'loading-engine'
    | 'server-queued'
    | 'server-running'
    | 'running'
    | 'stopped'
    | 'failed'
    | 'result';
  label: string;
  figure?: { value: string; tone: AnswerTone; unit: boolean };
  /** The first caption; `bold` is its lead (the band), drawn in --color-text. */
  caption?: { bold?: string; rest?: string; tone?: 'wipe' };
  /** The second caption: the step the player can take after a failure. */
  fix?: string;
  /** done / total for the Meter, or null for none. */
  meter?: { value: number; tone: 'gold' | 'muted' } | null;
  primary: PrimaryAction;
}

/** A snapshot taken when the player pressed Stop: the partial number stays, greyed (§5.1). */
export interface StoppedSnapshot {
  estimate: Estimate;
  done: number;
  total: number;
}

export interface AnswerInput {
  phase: SimPhase;
  simulated: boolean;
  specName: string;
  estimate: Estimate;
  iterationsDone: number;
  iterationsTotal: number;
  serverRunning: boolean;
  result: SimResult | null;
  message: string | null;
  detail: string;
  stopped: StoppedSnapshot | null;
  /** The lane the last run went to, for a failure's label. */
  lastLane: 'browser' | 'server';
}

const band = (estimate: Estimate): string => formatMargin(confidenceBand(estimate));
const fraction = (done: number, total: number): number => (total > 0 ? Math.min(1, done / total) : 0);

function runningAnswer(input: AnswerInput): AnswerModel {
  const { estimate, iterationsDone: done, iterationsTotal: total } = input;
  const hasFigure = estimate.mean > 0;
  return {
    state: 'running',
    label: copy.answer.runningBrowser,
    figure: hasFigure ? { value: grouped(estimate.mean), tone: 'gold', unit: true } : undefined,
    caption: hasFigure
      ? { bold: `± ${band(estimate)} (95%)`, rest: ` · ${grouped(done)} of ${grouped(total)} runs` }
      : { rest: `${grouped(done)} of ${grouped(total)} runs` },
    meter: { value: fraction(done, total), tone: 'gold' },
    primary: 'stop',
  };
}

function serverAnswer(input: AnswerInput): AnswerModel {
  const { iterationsDone: done, iterationsTotal: total } = input;
  if (done === 0)
    return {
      state: 'server-queued',
      label: copy.answer.onServers,
      caption: { bold: copy.answer.starting },
      meter: { value: 0, tone: 'muted' },
      primary: 'server',
    };
  return {
    state: 'server-running',
    label: copy.answer.onServers,
    caption: { rest: copy.answer.serverRunning(grouped(done), grouped(total)) },
    meter: { value: fraction(done, total), tone: 'gold' },
    primary: 'server',
  };
}

function stoppedAnswer(snapshot: StoppedSnapshot): AnswerModel {
  const estimate = grouped(snapshot.estimate.mean);
  const done = grouped(snapshot.done);
  return {
    state: 'stopped',
    label: copy.answer.stoppedLabel(done, grouped(snapshot.total)),
    figure: { value: copy.answer.stoppedFigure(estimate), tone: 'muted', unit: true },
    caption: { rest: copy.answer.stoppedCaption(done, estimate, band(snapshot.estimate)) },
    primary: 'run-again',
  };
}

function failedAnswer(input: AnswerInput): AnswerModel {
  if (input.lastLane === 'server')
    return {
      state: 'failed',
      label: copy.answer.serverFailed,
      caption: { rest: copy.answer.serverFailedReason, tone: 'wipe' },
      fix: copy.answer.serverFailedFix,
      primary: 'run-again',
    };
  const reason = input.detail !== '' ? input.detail : (input.message ?? '');
  return {
    state: 'failed',
    label: copy.answer.browserFailed,
    caption: { rest: copy.answer.browserRefused(reason.replace(/[.\s]+$/, '')), tone: 'wipe' },
    fix: copy.answer.browserRefusedFix,
    primary: 'run-again',
  };
}

/** The live result (§4.5): where it ran, the figure and its band labelled once (the spec is the head's). */
export function resultAnswer(result: SimResult): AnswerModel {
  return {
    state: 'result',
    label: copy.answer.resultLabel(result.lane === 'server' ? 'server' : 'browser'),
    figure: { value: oneDecimal(result.dps.mean), tone: 'gold', unit: true },
    caption: { bold: `± ${band(result.dps)} (95%)` },
    primary: 'run-again',
  };
}

/** The answer for a character on /sim, and the bar's primary for the same state. */
export function simAnswer(input: AnswerInput): AnswerModel {
  if (!input.simulated)
    return {
      state: 'not-simulated',
      label: copy.answer.notSimulated,
      caption: { rest: copy.answer.notSimulatedCaption(input.specName) },
      primary: 'disabled',
    };
  if (input.serverRunning) return serverAnswer(input);
  if (input.phase === 'loading-engine')
    return {
      state: 'loading-engine',
      label: copy.answer.loadingEngine,
      caption: { rest: copy.answer.loadingEngineCaption },
      primary: 'loading-engine',
    };
  if (input.phase === 'running') return runningAnswer(input);
  if (input.stopped !== null) return stoppedAnswer(input.stopped);
  if (input.phase === 'error' || (input.message !== null && input.result === null && input.phase !== 'idle'))
    return failedAnswer(input);
  if (input.result !== null) return resultAnswer(input.result);
  return {
    state: 'none',
    label: copy.answer.noSim,
    caption: { rest: copy.answer.noSimCaption },
    primary: 'run',
  };
}
