// The apps' parity region names (ParityRegion.swift, ParityRegion.kt) are the web's REGION values,
// so a region of an app mock pairs with the same region of the native build (design/parity.md
// "Native"). The site header and footer have no app counterpart, nor the paperdoll frame's
// medallion and centre (the apps' Gear tab draws the frame's list form, design/specs/2026-11-09).
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { REGION, sectionRegion } from '../ui/regions';

const SWIFT = '../../../../apps/ios/Sources/DesignSystem/ParityRegion.swift';
const KOTLIN =
  '../../../../apps/android/core/designsystem/src/main/kotlin/gg/foreversixty/app/designsystem/ParityRegion.kt';
const WEB_ONLY = new Set<string>([
  REGION.siteHeader,
  REGION.siteFooter,
  REGION.paperdollMedallion,
  REGION.paperdollCentre,
]);

const read = (path: string) => readFileSync(new URL(path, import.meta.url), 'utf8');
const sorted = (names: Iterable<string>) => [...names].sort();
const appRegions = sorted(Object.values(REGION).filter((name) => !WEB_ONLY.has(name)));

describe('native parity region names', () => {
  it('iOS names every app region the web has, and no other', () => {
    const names = [...read(SWIFT).matchAll(/public static let \w+ = Self\("([a-z-]+)"\)/g)].map((m) => m[1]);
    expect(sorted(names)).toEqual(appRegions);
  });

  it('Android names every app region the web has, and no other', () => {
    const names = [...read(KOTLIN).matchAll(/const val [A-Z_]+ = "([a-z-]+)"/g)].map((m) => m[1]);
    expect(sorted(names)).toEqual(appRegions);
  });

  it('section slugs follow the web rule on both (the native self-tests use these titles)', () => {
    expect(sectionRegion('What to stack')).toBe('section-what-to-stack');
    expect(sectionRegion('Top 3 for Fury Warrior, 60')).toBe('section-top-3-for-fury-warrior-60');
  });
});
