// web/src/lib/billing/example-image.test.ts
import { describe, expect, it } from 'vitest';
import { exampleImage } from './example-image';
import { serverRunCaption } from './premium-example-copy';
import type { ServerRunExample } from './premium-examples';

const VALID: ServerRunExample = {
  simId: 'fixture-run',
  preset: false,
  characterName: 'Thrall',
  level: 58,
  specLabel: 'Arms Warrior',
  placeholder: false,
  combinations: 1234,
  ranAt: '2027-02-03T10:00:00Z',
};

describe('exampleImage', () => {
  it('has the image source, size and loading', () => {
    expect(exampleImage(VALID)).toMatchObject({
      avif: '/premium/example-server-run.avif',
      webp: '/premium/example-server-run.webp',
      width: 440,
      height: 330,
      loading: 'lazy',
      decoding: 'async',
    });
  });

  it('reads the data file and its caption by default', () =>
    expect(exampleImage().alt).toBe(serverRunCaption()));

  it.each([true, false])('alt equals the caption (preset=%s)', (preset) =>
    expect(exampleImage({ ...VALID, preset }).alt).toBe(serverRunCaption({ ...VALID, preset })),
  );
});
