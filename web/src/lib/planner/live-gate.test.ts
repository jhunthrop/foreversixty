// web/src/lib/planner/live-gate.test.ts
import { describe, expect, it } from 'vitest';
import { isConstrainedDevice, liveGate } from './live-gate';
import { MAX_POINTS } from './types';

describe('liveGate', () => {
  it('runs nothing while points are still unspent, whatever the device', () => {
    expect(liveGate({ spent: 0, constrained: false, optedIn: false })).toBe('unfinished');
    expect(liveGate({ spent: MAX_POINTS - 1, constrained: false, optedIn: true })).toBe('unfinished');
  });

  it("counts a character build finished once its level's points are spent (gap G7)", () => {
    expect(liveGate({ spent: 24, needed: 25, constrained: false, optedIn: false })).toBe('unfinished');
    expect(liveGate({ spent: 25, needed: 25, constrained: false, optedIn: false })).toBe('run');
    expect(liveGate({ spent: 27, needed: 25, constrained: false, optedIn: false })).toBe('run');
    expect(liveGate({ spent: 0, needed: 0, constrained: false, optedIn: false })).toBe('unfinished');
  });

  it('sims a finished build on a desktop without being asked', () => {
    expect(liveGate({ spent: MAX_POINTS, constrained: false, optedIn: false })).toBe('run');
  });

  it('asks first on a phone or a data-saver connection, and remembers the answer', () => {
    expect(liveGate({ spent: MAX_POINTS, constrained: true, optedIn: false })).toBe('ask');
    expect(liveGate({ spent: MAX_POINTS, constrained: true, optedIn: true })).toBe('run');
  });
});

describe('isConstrainedDevice', () => {
  const pointer = (coarse: boolean) => ({ matchMedia: () => ({ matches: coarse }) });

  it('is a touch-first device', () => {
    expect(isConstrainedDevice(pointer(true))).toBe(true);
    expect(isConstrainedDevice(pointer(false))).toBe(false);
  });

  it('is anyone who asked their browser to save data, even with a mouse', () => {
    expect(isConstrainedDevice({ ...pointer(false), navigator: { connection: { saveData: true } } })).toBe(
      true,
    );
  });

  it('is not constrained where the browser can say neither, such as during the build', () => {
    expect(isConstrainedDevice({})).toBe(false);
  });
});
