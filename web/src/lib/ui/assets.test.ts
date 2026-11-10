import { describe, expect, it } from 'vitest';
import { SEAL, sealFor } from './assets';

describe('sealFor', () => {
  it('is the one painted seal (v4) at every size: one logo, nav and footer alike (owner 2026-11-09)', () => {
    expect(sealFor(16)).toBe(SEAL.painted88);
    expect(sealFor(40)).toBe(SEAL.painted88);
    expect(sealFor(88)).toBe(SEAL.painted88);
    expect(sealFor(128)).toBe(SEAL.painted128);
  });

  it('holds no flat seal', () => {
    expect(Object.keys(SEAL)).toEqual(['painted88', 'painted128']);
  });
});
