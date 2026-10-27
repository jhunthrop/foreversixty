import { describe, expect, it } from 'vitest';
import { createRefreshingRead } from './refreshing-read';

function deferred(): { promise: Promise<void>; resolve: () => void } {
  let resolve: () => void = () => {};
  const promise = new Promise<void>((done) => {
    resolve = done;
  });
  return { promise, resolve };
}

describe('createRefreshingRead', () => {
  it('re-reads when a change lands while a read is in flight', async () => {
    const gates = [deferred(), deferred()];
    let calls = 0;
    const reading = createRefreshingRead(() => gates[calls++].promise);
    const done = reading.run();
    reading.invalidate();
    expect(calls).toBe(1);
    gates[0].resolve();
    await Promise.resolve();
    await Promise.resolve();
    expect(calls).toBe(2);
    gates[1].resolve();
    await done;
    expect(calls).toBe(2);
  });

  it('reads again at once when idle', async () => {
    let calls = 0;
    const reading = createRefreshingRead(() => {
      calls += 1;
      return Promise.resolve();
    });
    reading.invalidate();
    await Promise.resolve();
    expect(calls).toBe(1);
  });

  it('coalesces several changes into one re-read', async () => {
    const gate = deferred();
    let calls = 0;
    const reading = createRefreshingRead(() => {
      calls += 1;
      return calls === 1 ? gate.promise : Promise.resolve();
    });
    const done = reading.run();
    reading.invalidate();
    reading.invalidate();
    gate.resolve();
    await done;
    expect(calls).toBe(2);
  });

  it('rejects a failed read and is usable again', async () => {
    let fail = true;
    const reading = createRefreshingRead(() => (fail ? Promise.reject(new Error('x')) : Promise.resolve()));
    await expect(reading.run()).rejects.toThrow('x');
    fail = false;
    await expect(reading.run()).resolves.toBeUndefined();
  });

  it('joins overlapping runs and still re-reads once for an invalidate', async () => {
    const gates = [deferred(), deferred()];
    let calls = 0;
    const reading = createRefreshingRead(() => gates[calls++].promise);
    const first = reading.run();
    const second = reading.run();
    reading.invalidate();
    expect(second).toBe(first);
    expect(calls).toBe(1);
    gates[0].resolve();
    await Promise.resolve();
    await Promise.resolve();
    expect(calls).toBe(2);
    gates[1].resolve();
    await Promise.all([first, second]);
    expect(calls).toBe(2);
  });
});
