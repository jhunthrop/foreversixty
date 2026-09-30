// web/src/lib/bis/rotation-view.test.ts
import { describe, expect, it } from 'vitest';
import { bandTopLevel, rotationEntryFor, type RotationEntry } from './rotation-view';

describe('bandTopLevel', () => {
  it('is band + 9 for every band but the last', () => {
    expect(bandTopLevel(20)).toBe(29);
    expect(bandTopLevel(50)).toBe(59);
  });

  it('is the band itself at the plain-60 band', () => {
    expect(bandTopLevel(60)).toBe(60);
  });
});

describe('rotationEntryFor', () => {
  const entries: RotationEntry[] = [
    { level: 10, lines: [] },
    { level: 20, lines: [] },
    { level: 30, lines: [] },
  ];

  it('picks the highest level entry at or under the band top (spec §4.C.3 worked example)', () => {
    expect(rotationEntryFor(entries, 29)?.level).toBe(20);
  });

  it('never picks an entry above the band top', () => {
    expect(rotationEntryFor(entries, 25)?.level).toBe(20);
  });

  it('is undefined when nothing qualifies yet', () => {
    expect(rotationEntryFor(entries, 5)).toBeUndefined();
  });
});
