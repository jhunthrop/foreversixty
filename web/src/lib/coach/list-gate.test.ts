import { describe, expect, it } from 'vitest';
import { LIST_GATE_MS, listGate, listGateOpen } from './list-gate';

describe('listGate', () => {
  it('holds the list while /v1/me is pending inside the window', () => {
    expect(listGate(false, false)).toBe('waiting');
    expect(listGateOpen('waiting')).toBe(false);
  });
  it('opens as soon as /v1/me resolves, signed in or out', () => {
    expect(listGate(true, false)).toBe('resolved');
    expect(listGateOpen('resolved')).toBe(true);
  });
  it('opens when the 150 ms are spent', () => {
    expect(LIST_GATE_MS).toBe(150);
    expect(listGate(false, true)).toBe('timed_out');
    expect(listGateOpen('timed_out')).toBe(true);
  });
  it('resolved wins over timed out', () => {
    expect(listGate(true, true)).toBe('resolved');
  });
});
