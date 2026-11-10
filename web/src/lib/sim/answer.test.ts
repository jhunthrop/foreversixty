import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/result-full.json';
import { simAnswer, type AnswerInput } from './answer';
import type { SimResult } from './types';

const result = fixture as unknown as SimResult;
const EMPTY = { mean: 0, error: 0, stddev: 0, min: 0, max: 0 };

function input(over: Partial<AnswerInput>): AnswerInput {
  return {
    phase: 'idle',
    simulated: true,
    specName: 'Fury Warrior',
    estimate: EMPTY,
    iterationsDone: 0,
    iterationsTotal: 0,
    serverRunning: false,
    result: null,
    message: null,
    detail: '',
    stopped: null,
    lastLane: 'browser',
    ...over,
  };
}

describe('simAnswer', () => {
  it('has no sim yet before a run, with Run sim the primary', () => {
    const a = simAnswer(input({}));
    expect(a).toMatchObject({ state: 'none', label: 'No sim yet', primary: 'run' });
    expect(a.figure).toBeUndefined();
  });

  it('a spec the sim does not run disables the primary and says so', () => {
    const a = simAnswer(input({ simulated: false, specName: 'Holy Priest' }));
    expect(a.primary).toBe('disabled');
    expect(a.caption?.rest).toBe('Holy Priest is not simulated yet; there is nothing to run.');
  });

  it('running: a whole number in gold, the live 95% band, runs done of total, Stop', () => {
    const a = simAnswer(
      input({
        phase: 'running',
        estimate: { ...EMPTY, mean: 905.3, error: 1.34 },
        iterationsDone: 1357,
        iterationsTotal: 10000,
      }),
    );
    expect(a.figure).toEqual({ value: '905', tone: 'gold', unit: true });
    expect(a.caption).toEqual({ bold: '± 2.6 DPS', rest: ' · 1,357 of 10,000 runs' });
    expect(a.meter?.value).toBeCloseTo(0.1357);
    expect(a.primary).toBe('stop');
  });

  it('the result reads its band in DPS, with no lane label', () => {
    const a = simAnswer(input({ phase: 'done', result }));
    expect(a.state).toBe('result');
    expect(a.label).toBe('');
    expect(a.caption?.bold).toMatch(/^± [\d.]+ DPS$/);
    expect(a.caption?.rest).toBeUndefined();
    expect(a.primary).toBe('run-again');
  });

  it('stopped keeps the partial number, greyed', () => {
    const a = simAnswer(
      input({
        message: 'Stopped.',
        stopped: { estimate: { ...EMPTY, mean: 905.3, error: 1.34 }, done: 1357, total: 10000 },
      }),
    );
    expect(a).toMatchObject({
      state: 'stopped',
      label: 'Stopped at 1,357 of 10,000 runs',
      primary: 'run-again',
    });
    expect(a.figure).toEqual({ value: 'about 905', tone: 'muted', unit: true });
    expect(a.caption?.rest).toBe('Stopped at 1,357 runs: about 905 ± 2.6. A stopped run is not saved.');
  });

  it('a browser failure names the engine reason and the one step to take', () => {
    const a = simAnswer(input({ phase: 'error', message: 'x', detail: 'Finger 1 holds item 999999.' }));
    expect(a.caption).toEqual({
      rest: 'The engine could not run this character: Finger 1 holds item 999999.',
      tone: 'wipe',
    });
    expect(a.fix).toMatch(/^Paste a fresh addon export/);
  });

  it('server lane: starting, running, failed', () => {
    expect(simAnswer(input({ serverRunning: true, iterationsTotal: 10000 })).state).toBe('server-queued');
    expect(
      simAnswer(input({ serverRunning: true, iterationsDone: 4000, iterationsTotal: 10000 })).caption?.rest,
    ).toBe('Running · 4,000 of 10,000 runs');
    expect(simAnswer(input({ phase: 'error', message: 'x', lastLane: 'server' })).label).toBe(
      'On our servers · failed',
    );
  });
});
