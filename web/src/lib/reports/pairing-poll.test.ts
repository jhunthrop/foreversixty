// web/src/lib/reports/pairing-poll.test.ts
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { Device } from '../account/api';
import { createPairingPoll, findNewDevice, PAIRING_POLL_INTERVAL_MS } from './pairing-poll';

function device(overrides: Partial<Device> = {}): Device {
  return {
    id: 'd1',
    name: 'MacBook Pro',
    platform: 'macOS',
    created_at: '2026-10-04T00:00:00Z',
    last_seen_at: null,
    ...overrides,
  };
}

describe('findNewDevice', () => {
  it('is null when every device was already known', () => {
    const before = new Set(['d1']);
    expect(findNewDevice(before, [device({ id: 'd1' })])).toBeNull();
  });

  it('finds the one id not in the before set', () => {
    const before = new Set(['d1']);
    const after = [device({ id: 'd1' }), device({ id: 'd2', name: 'Companion' })];
    expect(findNewDevice(before, after)?.id).toBe('d2');
  });
});

describe('createPairingPoll', () => {
  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.useRealTimers();
  });

  it('ends the poll and fires onPaired once a new device id appears', async () => {
    const newDevice = device({ id: 'd2', name: 'Companion' });
    let call = 0;
    const listDevices = vi.fn(async () => {
      call += 1;
      return call === 1 ? [] : [newDevice];
    });
    const onPaired = vi.fn();
    const poll = createPairingPoll(listDevices, new Set(), onPaired);
    poll.start();

    await vi.advanceTimersByTimeAsync(PAIRING_POLL_INTERVAL_MS);
    expect(onPaired).not.toHaveBeenCalled();

    await vi.advanceTimersByTimeAsync(PAIRING_POLL_INTERVAL_MS);
    expect(onPaired).toHaveBeenCalledWith(newDevice);

    // The poll stopped itself: no further listDevices calls past the one that found it.
    const callsAtSuccess = listDevices.mock.calls.length;
    await vi.advanceTimersByTimeAsync(PAIRING_POLL_INTERVAL_MS * 3);
    expect(listDevices.mock.calls.length).toBe(callsAtSuccess);
  });

  it('stops polling once the code expires, with no success', async () => {
    const listDevices = vi.fn(async () => []);
    const onPaired = vi.fn();
    const poll = createPairingPoll(listDevices, new Set(), onPaired, {
      expiresInMs: PAIRING_POLL_INTERVAL_MS * 2.5,
    });
    poll.start();

    await vi.advanceTimersByTimeAsync(PAIRING_POLL_INTERVAL_MS * 10);
    expect(onPaired).not.toHaveBeenCalled();
    const callsAtExpiry = listDevices.mock.calls.length;
    // Expired at 2.5 intervals -- no more than 2-3 ticks should have fired before the
    // expiry timer stopped the interval.
    expect(callsAtExpiry).toBeLessThanOrEqual(3);
  });

  it('stop() ends the poll immediately', async () => {
    const listDevices = vi.fn(async () => []);
    const poll = createPairingPoll(listDevices, new Set(), vi.fn());
    poll.start();
    poll.stop();
    await vi.advanceTimersByTimeAsync(PAIRING_POLL_INTERVAL_MS * 5);
    expect(listDevices).not.toHaveBeenCalled();
  });
});
