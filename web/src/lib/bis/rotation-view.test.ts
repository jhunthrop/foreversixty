// web/src/lib/bis/rotation-view.test.ts
import { describe, expect, it } from 'vitest';
import {
  alignNoteRank,
  bandTopLevel,
  loadRotations,
  rotationEntryFor,
  rotationLinesFor,
  type RotationEntry,
} from './rotation-view';

const BUILD = '1.60.1.70291';

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
    const entries = loadRotations(BUILD, 'hunter-beast-mastery');
    const entry = entries?.[0];
    expect(entry).toBeDefined();
    const views = rotationLinesFor(entry!, BUILD, 'hunter', 'hunter-beast-mastery');
    const serpentSting = views.find((view) => view.name === 'Serpent Sting');
    expect(serpentSting?.icon).toBe('ability_hunter_quickshot');
  });

  it('omits icon when the line names no icon at all', () => {
    const entry: RotationEntry = {
      level: 10,
      lines: [{ spell_id: 999999, name: 'No Icon Line', condition: 'x' }],
    };
    const views = rotationLinesFor(entry, BUILD, 'hunter', 'hunter-beast-mastery');
    expect(views[0].icon).toBeUndefined();
  });

  it("omits icon when the build's icon tree has no file for the named stem", () => {
    const entry: RotationEntry = {
      level: 10,
      lines: [
        { spell_id: 999999, name: 'Missing File Line', condition: 'x', icon: 'does_not_exist_anywhere' },
      ],
    };
    const views = rotationLinesFor(entry, BUILD, 'hunter', 'hunter-beast-mastery');
    expect(views[0].icon).toBeUndefined();
  });
});

describe('rotationLinesFor on the holy priest band 20 list', () => {
  const entry = rotationEntryFor(loadRotations(BUILD, 'priest-holy') ?? [], bandTopLevel(20));
  const rows = () => rotationLinesFor(entry!, BUILD, 'priest', 'priest-holy');

  it('has one row per run of the same spell', () => {
    expect(entry).toBeDefined();
    expect(rows().map((row) => row.name)).toEqual([
      'Shadowfiend',
      'Dark Sacrifice',
      'Renew',
      'Flash Heal',
      'Heal',
    ]);
    expect(rows().map((row) => row.steps)).toEqual([1, 1, 1, 5, 10]);
  });

  it('says every condition of the collapsed steps in words', () => {
    const heal = rows().find((row) => row.name === 'Heal');
    expect(heal?.condition).toMatch(
      /Cast on the tank below \d+%, then a party member below \d+%, when your mana is at least \d+% of the share of the fight left; or on the tank below \d+%, then a party member below \d+%\./,
    );
    const flash = rows().find((row) => row.name === 'Flash Heal');
    expect(flash?.condition).toMatch(/Cast on the tank below \d+%, then a party member below \d+%\./);
  });

  it("shows the rank the band casts and corrects the note's rank to it", () => {
    const heal = rows().find((row) => row.name === 'Heal');
    expect(heal?.rank).toBe(1);
    expect(heal?.condition).toContain('Heal rank 1 is the mana-efficient filler');
    expect(heal?.condition).not.toContain('Heal rank 4');
    expect(rows().find((row) => row.name === 'Renew')?.condition).toContain('at rank 3');
  });
});

describe('alignNoteRank', () => {
  it('rewrites the spell own rank and nothing else', () => {
    const note = 'Heal rank 4 beats Greater Heal rank 5, and Flash Heal (rank 7) is dear.';
    expect(alignNoteRank(note, 'Heal', 1)).toBe(
      'Heal rank 1 beats Greater Heal rank 5, and Flash Heal (rank 7) is dear.',
    );
    expect(alignNoteRank(note, 'Flash Heal', 2)).toContain('Flash Heal (rank 2)');
  });

  it('leaves a single-rank ability alone', () => {
    expect(alignNoteRank('Shadowfiend rank 2', 'Shadowfiend', undefined)).toBe('Shadowfiend rank 2');
  });
});

describe('every published rotation entry lines up with its curated steps', () => {
  it.each(['priest-holy', 'hunter-beast-mastery', 'warrior-arms', 'mage-fire'])('%s', (spec) => {
    const [classSlug] = spec.split('-');
    for (const entry of loadRotations(BUILD, spec) ?? []) {
      const views = rotationLinesFor(entry, BUILD, classSlug, spec);
      expect(views.reduce((sum, view) => sum + view.steps, 0)).toBe(entry.lines.length);
      views.slice(1).forEach((view, index) => {
        const previous = views[index];
        expect(`${view.name}|${view.rank}`).not.toBe(`${previous.name}|${previous.rank}`);
      });
    }
  });
});
