import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { inlineArt } from './art-inline';
import { BAND_ARTS } from './assets';

const publicFile = (path: string): Buffer => readFileSync(new URL(`../../../public${path}`, import.meta.url));

describe('inlineArt', () => {
  it.each(Object.entries(BAND_ARTS))(
    'is the %s phone file, byte for byte, as a WebP data URL',
    (_key, art) => {
      const [head, body] = inlineArt(art.phone).split(',');
      expect(head).toBe('data:image/webp;base64');
      expect(Buffer.from(body ?? '', 'base64').equals(publicFile(art.phone))).toBe(true);
    },
  );

  it('fails loudly on a file it has no data URL type for, or one that is missing', () => {
    expect(() => inlineArt('/art/og-art.jpg')).toThrow(/no data URL type/);
    expect(() => inlineArt('/art/missing.webp')).toThrow(/ENOENT/);
  });
});
