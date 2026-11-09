import { describe, expect, it } from 'vitest';
import { OG_BRAND, ogLayout, renderOg } from './og';

function size(png: Uint8Array): number[] {
  const view = new DataView(png.buffer, png.byteOffset);
  return [view.getUint32(16), view.getUint32(20)];
}

describe('renderOg', () => {
  it('returns a PNG of the expected size', async () => {
    const png = await renderOg({ title: 'Hall of Thanes', kicker: 'Dungeon · Dun Morogh' });
    expect(png.slice(0, 8)).toEqual(new Uint8Array([0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a]));
    expect(size(png)).toEqual([1200, 630]);
  });

  it("draws home's card as the brand card, the name never twice", async () => {
    expect(ogLayout({ title: OG_BRAND, kicker: 'World of Warcraft: Forever' })).toBe('brand');
    expect(ogLayout({ title: 'Best in slot', kicker: 'Gear' })).toBe('page');
    expect(size(await renderOg({ title: OG_BRAND, kicker: 'World of Warcraft: Forever' }))).toEqual([
      1200, 630,
    ]);
  });
});
