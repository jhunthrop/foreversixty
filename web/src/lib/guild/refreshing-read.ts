// web/src/lib/guild/refreshing-read.ts
// A read that can be told its answer is already out of date. A change that lands while the read is in
// flight marks it stale, and the read runs again when it settles, so the page never ends on a list
// that predates the change.

export interface RefreshingRead {
  /** Runs the read (and any re-read a change asked for); a call while one runs joins it. Rejects when a
   *  read fails: a read that fails after being marked stale reports that failure (the tab shows
   *  `missing`) although a further re-read might have succeeded. */
  run: () => Promise<void>;
  /** The data changed: re-read now, or once the read in flight settles. A failed re-read is dropped. */
  invalidate: () => void;
}

export function createRefreshingRead(read: () => Promise<void>): RefreshingRead {
  let inFlight: Promise<void> | null = null;
  let stale = false;

  async function loop(): Promise<void> {
    try {
      do {
        stale = false;
        await read();
      } while (stale);
    } finally {
      inFlight = null;
    }
  }

  function run(): Promise<void> {
    inFlight ??= loop();
    return inFlight;
  }

  function invalidate(): void {
    if (inFlight !== null) stale = true;
    else void run().catch(() => {});
  }

  return { run, invalidate };
}
