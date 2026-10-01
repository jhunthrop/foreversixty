// web/src/lib/bis/rotation-view.test.ts
import { describe, expect, it } from 'vitest';
import {
  bandTopLevel,
  loadRotations,
  rotationEntryFor,
  rotationLinesFor,
  type RotationEntry,
} from './rotation-view';

const BUILD = '1.60.1.70009';

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

describe('rotationLinesFor', () => {
  it("threads a line's own icon stem through to the view when the build ships the file", () => {
    const entries = loadRotations(BUILD, 'hunter-marksmanship');
    const entry = entries?.[0];
    expect(entry).toBeDefined();
    const views = rotationLinesFor(entry!, BUILD, 'hunter');
    const serpentSting = views.find((view) => view.name === 'Serpent Sting');
    expect(serpentSting?.icon).toBe('ability_hunter_quickshot');
  });

  it('omits icon when the line names no icon at all', () => {
    const entry: RotationEntry = {
      level: 10,
      lines: [{ spell_id: 999999, name: 'No Icon Line', condition: 'x' }],
    };
    const views = rotationLinesFor(entry, BUILD, 'hunter');
    expect(views[0].icon).toBeUndefined();
  });

  it("omits icon when the build's icon tree has no file for the named stem", () => {
    const entry: RotationEntry = {
      level: 10,
      lines: [
        { spell_id: 999999, name: 'Missing File Line', condition: 'x', icon: 'does_not_exist_anywhere' },
      ],
    };
    const views = rotationLinesFor(entry, BUILD, 'hunter');
    expect(views[0].icon).toBeUndefined();
  });
});
