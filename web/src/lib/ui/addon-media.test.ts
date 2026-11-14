// The addon's textures (addon/ForeverSixty/media) against the site sources they are written from by
// scripts/gen-addon-icons.mjs (design/specs/2026-11-14-addon-brand.md section 3). The check hashes the
// sources, the generator and the committed files against media/MANIFEST.json, so it holds on any
// platform; the renderers' bytes are not compared.
import { describe, expect, it } from 'vitest';
import { ADDON_ICONS, checkAddonMedia, tga } from '../../../scripts/gen-addon-icons.mjs';

describe('the addon media', () => {
  it('is current with the site sources (node scripts/gen-addon-icons.mjs)', () => {
    expect(checkAddonMedia()).toEqual([]);
  });

  it('draws every hub tab with a site glyph', () => {
    expect(ADDON_ICONS.slice(0, 6)).toEqual(['character', 'planner', 'bis', 'sim', 'guild', 'manage']);
  });

  it('writes 32-bit uncompressed top-left TGAs and refuses a side that is not a power of two', () => {
    const out = tga(2, 1, Buffer.from([1, 2, 3, 4, 5, 6, 7, 8]));
    expect([...out.subarray(0, 18)]).toEqual([0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 1, 0, 32, 0x20]);
    expect([...out.subarray(18)]).toEqual([3, 2, 1, 4, 7, 6, 5, 8]);
    expect(() => tga(3, 2, Buffer.alloc(24))).toThrow(/power-of-two/);
  });
});
