// web/src/lib/ui/preload-images.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { preloadImages } from './preload-images';

afterEach(() => vi.useRealTimers());

function fakeImage(decode: () => Promise<void>): HTMLImageElement {
  return { src: '', decode } as unknown as HTMLImageElement;
}

describe('preloadImages', () => {
  it('resolves at once for no urls', async () => {
    await expect(preloadImages([], 1000)).resolves.toBeUndefined();
  });

  it('resolves when every image decoded or failed', async () => {
    const decodes = [() => Promise.resolve(), () => Promise.reject(new Error('broken'))];
    let call = 0;
    const create = () => fakeImage(decodes[call++ % 2] as () => Promise<void>);
    await expect(preloadImages(['/a.webp', '/b.webp'], 1000, create)).resolves.toBeUndefined();
  });

  it('gives up after the cap when an image never settles', async () => {
    vi.useFakeTimers();
    const done = vi.fn();
    const never = () => fakeImage(() => new Promise<void>(() => {}));
    void preloadImages(['/slow.webp'], 1000, never).then(done);
    await vi.advanceTimersByTimeAsync(999);
    expect(done).not.toHaveBeenCalled();
    await vi.advanceTimersByTimeAsync(1);
    expect(done).toHaveBeenCalled();
  });
});
