// web/src/lib/reports/pairing-poll.ts
// Logs landing spec (2026-10-04) §4.D.2: once a pairing code is shown, poll `listDevices()`
// every 5 seconds and compare the returned ids against the set captured the moment the code
// was requested. A new id ends the poll with a success line; the code's own `expires_in`
// elapsing with no new device ends it with nothing (the existing expiry display is
// unchanged -- this module only has to stop polling by then, not render anything itself).
import type { Device } from '../account/api';

export const PAIRING_POLL_INTERVAL_MS = 5000;

/** The first device in `after` that was not in `before` -- null when none is new. One
 *  pairing code pairs one device, so "first" is a defensive tie-break, not an expected
 *  multi-match case. */
export function findNewDevice(before: ReadonlySet<string>, after: readonly Device[]): Device | null {
  return after.find((device) => !before.has(device.id)) ?? null;
}

export interface PairingPoll {
  start(): void;
  stop(): void;
}

export interface PairingPollOptions {
  intervalMs?: number;
  /** Stops the poll with no success once this elapses -- the code's own `expires_in`, in ms. */
  expiresInMs?: number;
}

/** Wraps `findNewDevice` in an interval timer. `onPaired` fires at most once and ends the
 *  poll itself; `stop()` is for the caller's own unmount/cleanup. */
export function createPairingPoll(
  listDevices: () => Promise<readonly Device[]>,
  beforeIds: ReadonlySet<string>,
  onPaired: (device: Device) => void,
  options: PairingPollOptions = {},
): PairingPoll {
  const intervalMs = options.intervalMs ?? PAIRING_POLL_INTERVAL_MS;
  let timer: ReturnType<typeof setInterval> | null = null;
  let expiry: ReturnType<typeof setTimeout> | null = null;
  let stopped = false;

  function clear(): void {
    if (timer !== null) clearInterval(timer);
    if (expiry !== null) clearTimeout(expiry);
    timer = null;
    expiry = null;
  }

  return {
    start(): void {
      if (stopped) return;
      timer = setInterval(() => {
        void listDevices().then((devices) => {
          if (stopped) return;
          const found = findNewDevice(beforeIds, devices);
          if (found !== null) {
            stopped = true;
            clear();
            onPaired(found);
          }
        });
      }, intervalMs);
      if (options.expiresInMs !== undefined) {
        expiry = setTimeout(() => {
          stopped = true;
          clear();
        }, options.expiresInMs);
      }
    },
    stop(): void {
      stopped = true;
      clear();
    },
  };
}
