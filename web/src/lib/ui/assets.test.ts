import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { FAVICON, SEAL, sealFor } from './assets';

const publicFile = (path: string): Buffer => readFileSync(new URL(`../../../public${path}`, import.meta.url));

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
