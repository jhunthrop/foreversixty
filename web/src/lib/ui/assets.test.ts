import { describe, expect, it } from 'vitest';
import { SEAL, sealFor } from './assets';

describe('sealFor', () => {
  it('is the flat seal only at favicon size', () => {
    expect(sealFor(16)).toBe(SEAL.flat32);
  });

  it('is the painted seal everywhere else, the nav included: one logo (owner 2026-11-09)', () => {
    expect(sealFor(40)).toBe(SEAL.painted88);
    expect(sealFor(64)).toBe(SEAL.painted88);
    expect(sealFor(88)).toBe(SEAL.painted88);
    expect(sealFor(128)).toBe(SEAL.painted128);
  });

  it('holds no flat seal above favicon size', () => {
    expect(Object.keys(SEAL)).toEqual(['flat32', 'painted88', 'painted128']);
  });
});
