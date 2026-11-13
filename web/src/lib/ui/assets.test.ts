import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { ART_CSS_VARS, BAND_ART, BAND_ART_INLINE, FAVICON, catalogueIconUrl, iconUrl } from './assets';
import { SEAL, sealFor } from './seal';
import { PLACEHOLDER_ICON } from '../icons/placeholder';

const publicFile = (path: string): Buffer => readFileSync(new URL(`../../../public${path}`, import.meta.url));

describe('iconUrl', () => {
  it("is the build's HD icon file for a real stem", () => {
    expect(iconUrl('1.15.8.65000', 'spell_holy_heal')).toBe('/data/1.15.8.65000/icons/spell_holy_heal.webp');
  });

  it('is null for no stem, an empty one and the red question mark placeholder', () => {
    expect(iconUrl('b', undefined)).toBeNull();
    expect(iconUrl('b', null)).toBeNull();
    expect(iconUrl('b', '')).toBeNull();
    expect(iconUrl('b', PLACEHOLDER_ICON)).toBeNull();
  });
});

describe('catalogueIconUrl', () => {
  it('is the icon url for a real stem', () => {
    expect(catalogueIconUrl('b', 'ability_hunter_swiftstrike')).toBe(
      '/data/b/icons/ability_hunter_swiftstrike.webp',
    );
  });

  it('throws on the placeholder or an empty stem', () => {
    expect(() => catalogueIconUrl('b', PLACEHOLDER_ICON)).toThrow(/not a real icon/);
    expect(() => catalogueIconUrl('b', '')).toThrow(/not a real icon/);
  });
});

describe('sealFor', () => {
  it('is the tiny bold flat seal at 24 px and below (owner 2026-11-10)', () => {
    expect(sealFor(16)).toBe(SEAL.flatTiny);
    expect(sealFor(24)).toBe(SEAL.flatTiny);
  });

  it('is the small flat seal from 25 to 44 px: the nav', () => {
    expect(sealFor(25)).toBe(SEAL.flatSmall);
    expect(sealFor(34)).toBe(SEAL.flatSmall);
    expect(sealFor(40)).toBe(SEAL.flatSmall);
    expect(sealFor(44)).toBe(SEAL.flatSmall);
  });

  it('is the painted seal v4 above 44 px: the footer and the medallion', () => {
    expect(sealFor(45)).toBe(SEAL.painted88);
    expect(sealFor(88)).toBe(SEAL.painted88);
    expect(sealFor(128)).toBe(SEAL.painted128);
  });

  it('serves the flat seals as SVG, crisp at every density', () => {
    expect(SEAL.flatTiny).toMatch(/\.svg$/);
    expect(SEAL.flatSmall).toMatch(/\.svg$/);
  });
});

describe('FAVICON', () => {
  it('is the tiny bold flat seal, byte for byte the nav set (tools/art/build_art.py copies both)', () => {
    expect(publicFile(FAVICON.svg).equals(publicFile(SEAL.flatTiny))).toBe(true);
  });
});

describe('BAND_ART_INLINE', () => {
  it('is the phone band file, byte for byte, as a WebP data URL', () => {
    const [head, body] = BAND_ART_INLINE.split(',');
    expect(head).toBe('data:image/webp;base64');
    expect(Buffer.from(body ?? '', 'base64').equals(publicFile(BAND_ART.phone))).toBe(true);
  });

  it('is what the pages paint the band with, so it needs no request of its own', () => {
    expect(ART_CSS_VARS['--art-band-phone']).toBe(`url('${BAND_ART_INLINE}')`);
  });
});
