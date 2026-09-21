// web/src/lib/planner/live-gate.ts
// Whether the planner should be simming at all right now.
//
// A sim behind every click was three kinds of waste. The first click of a visit downloaded a
// 3.7 MB engine for someone who may only have come to plan talents; a build with twelve of
// its fifty-one points spent produced a number that means nothing and moved on every click;
// and a phone spun up workers for a figure nobody had asked for. So the live estimate waits
// for a build worth simming, and on a device that pays for it in battery or data it waits to
// be asked.
import { MAX_POINTS } from './types';

/**
 * `unfinished`  points are still unspent: nothing is downloaded or run.
 * `ask`         the build is complete, but this device should be asked first.
 * `run`         sim it, and keep simming it on every change.
 */
export type LiveGate = 'unfinished' | 'ask' | 'run';

export interface LiveGateInput {
  spent: number;
  /** A touch-first device or a data-saver connection: see `isConstrainedDevice`. */
  constrained: boolean;
  /** The visitor pressed "Show DPS" on this page. It holds for the rest of the visit. */
  optedIn: boolean;
}

export function liveGate({ spent, constrained, optedIn }: LiveGateInput): LiveGate {
  if (spent < MAX_POINTS) return 'unfinished';
  return constrained && !optedIn ? 'ask' : 'run';
}

/** The slice of the browser this module reads, so a test can hand it a fake. */
export interface DeviceProbe {
  matchMedia?: (query: string) => { matches: boolean };
  navigator?: { connection?: { saveData?: boolean } };
}

const TOUCH_FIRST = '(pointer: coarse)';

/**
 * A phone or tablet, or anyone who has asked their browser to save data. Neither is proof
 * of a weak device; both are people for whom an unrequested 3.7 MB download and a burst of
 * worker threads is a cost they should choose.
 */
export function isConstrainedDevice(probe: DeviceProbe = globalThis as DeviceProbe): boolean {
  if (probe.navigator?.connection?.saveData === true) return true;
  return probe.matchMedia?.(TOUCH_FIRST).matches ?? false;
}
