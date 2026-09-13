import { describe, expect, it } from 'vitest';
import { renderOg } from './og';

describe('renderOg', () => {
  it('returns a PNG of the expected size', async () => {
    const png = await renderOg({ title: 'Hall of Thanes', kicker: 'Dungeon · Dun Morogh' });
    expect(png.slice(0, 8)).toEqual(new Uint8Array([0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a]));
    const width = new DataView(png.buffer, png.byteOffset).getUint32(16);
    const height = new DataView(png.buffer, png.byteOffset).getUint32(20);
    expect([width, height]).toEqual([1200, 630]);
  });
});
