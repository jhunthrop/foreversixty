// web/src/lib/guild/refreshing-read.ts
// A read that can be told its answer is already out of date. A change that lands while the read is in
// flight marks it stale, and the read runs again when it settles, so the page never ends on a list
// that predates the change.

export interface RefreshingRead {
  /** Runs the read (and any re-read a change asked for); rejects when a read fails. */
  run: () => Promise<void>;
  /** The data changed: re-read now, or once the read in flight settles. A failed re-read is dropped. */
  invalidate: () => void;
}

export function createRefreshingRead(read: () => Promise<void>): RefreshingRead {
  let inFlight = false;
  let stale = false;

  async function run(): Promise<void> {
    inFlight = true;
    try {
      do {
        stale = false;
        await read();
      } while (stale);
    } finally {
      inFlight = false;
    }
  }

  function invalidate(): void {
    if (inFlight) stale = true;
    else void run().catch(() => {});
  }

  return { run, invalidate };
}
