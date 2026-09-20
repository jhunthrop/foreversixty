import { describe, expect, it } from 'vitest';
import stylesJson from '../../fixtures/sim/styles.json';
import { simCopy } from './copy';
import { DEFAULT_STYLE_ID, FIGHT_STYLES, applyFightStyle, fightStyle, targetsSummary } from './styles';
import { DEFAULT_ENCOUNTER } from './types';

interface StyleRow {
  id: string;
  targets: number;
  execute_ratio: number;
  movement: { interval_sec: number; duration_sec: number; kind: string } | null;
  targets_over_time: { at_sec: number; count: number }[] | null;
  dummy: boolean;
}

// The fixture is the shared artefact: the sim module lane's own table test reads this same
// file, so a style added on one side and not the other fails on both.
const fixture = (stylesJson as { styles: StyleRow[] }).styles;

describe('FIGHT_STYLES', () => {
  it('is exactly the fixture, in the fixture’s order', () => {
    expect(FIGHT_STYLES.map((style) => style.id)).toEqual(fixture.map((row) => row.id));
  });

  it('expands each style to the fixture’s encounter fields', () => {
    for (const row of fixture) {
      const style = fightStyle(row.id);
      expect(style, row.id).not.toBeNull();
      expect({
        id: style!.id,
        targets: style!.targets,
        execute_ratio: style!.execute_ratio,
        movement: style!.movement,
        targets_over_time: style!.targets_over_time,
        dummy: style!.dummy,
      }).toEqual(row);
    }
  });

  it('is the contract’s nine ids', () => {
    expect(FIGHT_STYLES.map((style) => style.id)).toEqual([
      'patchwerk',
      'execute',
      'light-movement',
      'heavy-movement',
      'cleave-2',
      'cleave-3',
      'cleave-5',
      'dungeon',
      'dummy',
    ]);
  });

  it('opens on Patchwerk', () => {
    expect(DEFAULT_STYLE_ID).toBe('patchwerk');
    expect(fightStyle(DEFAULT_STYLE_ID)?.execute_ratio).toBe(DEFAULT_ENCOUNTER.execute_ratio);
  });

  it('names every style, and notes only the two that need one', () => {
    for (const style of FIGHT_STYLES) {
      expect(simCopy.styleLabel[style.id], style.id).toBeTruthy();
    }
    expect(Object.keys(simCopy.styleNote).sort()).toEqual(['heavy-movement', 'light-movement']);
  });

  it('answers null for an id nothing defines rather than guessing', () => {
    expect(fightStyle('raidbots-patchwerk')).toBeNull();
    expect(fightStyle('')).toBeNull();
  });
});

describe('applyFightStyle', () => {
  it('writes the style’s fields and the label, and keeps duration and variation', () => {
    const next = applyFightStyle({ ...DEFAULT_ENCOUNTER, duration_sec: 300, variation: 0.1 }, 'cleave-3');
    expect(next.style).toBe('cleave-3');
    expect(next.targets).toBe(3);
    expect(next.execute_ratio).toBe(0.25);
    expect(next.duration_sec).toBe(300);
    expect(next.variation).toBe(0.1);
    expect(next.dummy).toBe(false);
  });

  it('clears a previous style’s movement and timeline rather than leaving them behind', () => {
    const moving = applyFightStyle(DEFAULT_ENCOUNTER, 'heavy-movement');
    expect(moving.movement).toEqual({ interval_sec: 20, duration_sec: 5, kind: 'away' });
    const back = applyFightStyle(moving, 'patchwerk');
    expect(back.movement).toBeUndefined();
    expect(back.targets_over_time).toBeUndefined();
  });

  it('omits movement and targets_over_time from the wire shape when a style clears them', () => {
    // The contract's fields are Go `omitempty` pointers/slices: the wire shape is the key
    // absent, never `"movement":null`. This is the property that actually matters --
    // `undefined` vs `null` is only how JS gets there.
    const back = applyFightStyle(applyFightStyle(DEFAULT_ENCOUNTER, 'heavy-movement'), 'patchwerk');
    const wire = JSON.parse(JSON.stringify(back)) as Record<string, unknown>;
    expect('movement' in wire).toBe(false);
    expect('targets_over_time' in wire).toBe(false);
  });

  it('gives the dungeon pull its target-count timeline and no execute window', () => {
    const dungeon = applyFightStyle(DEFAULT_ENCOUNTER, 'dungeon');
    expect(dungeon.execute_ratio).toBe(0);
    expect(dungeon.targets_over_time).toEqual([
      { at_sec: 0, count: 1 },
      { at_sec: 40, count: 3 },
      { at_sec: 80, count: 5 },
      { at_sec: 130, count: 3 },
      { at_sec: 160, count: 1 },
    ]);
  });

  it('turns the dummy flag on for the dummy and off for everything else', () => {
    expect(applyFightStyle(DEFAULT_ENCOUNTER, 'dummy').dummy).toBe(true);
    expect(applyFightStyle(applyFightStyle(DEFAULT_ENCOUNTER, 'dummy'), 'execute').dummy).toBe(false);
  });

  it('never mutates the encounter it was given', () => {
    const base = { ...DEFAULT_ENCOUNTER };
    applyFightStyle(base, 'cleave-5');
    expect(base.targets).toBe(1);
    expect(base.style).toBe('patchwerk');
  });
});

// tank MAJOR, review.md:325-327: the on-screen TARGETS control has to describe the run
// even under a timeline style, where `encounter.targets` (1, the dungeon pull's opening
// count) is not what the engine actually simulates. `targetsSummary` is the pure decision
// SettingsBar.svelte renders off of -- styles.ts holds no words, so this is numbers and a
// flag, never a sentence.
describe('targetsSummary', () => {
  it('reads the dungeon pull as its ramp’s first and maximum counts, not its settled targets field', () => {
    const dungeon = applyFightStyle(DEFAULT_ENCOUNTER, 'dungeon');
    expect(targetsSummary(dungeon)).toEqual({ timeline: true, first: 1, max: 5, attached: true });
  });

  it('reads a style with no timeline as its own target count', () => {
    expect(targetsSummary(applyFightStyle(DEFAULT_ENCOUNTER, 'cleave-3'))).toEqual({
      timeline: false,
      count: 3,
    });
    expect(targetsSummary(applyFightStyle(DEFAULT_ENCOUNTER, 'patchwerk'))).toEqual({
      timeline: false,
      count: 1,
    });
  });

  it('reads an encounter detached from its style by its own targets field', () => {
    const detached = { ...DEFAULT_ENCOUNTER, style: '', targets: 4, targets_over_time: undefined };
    expect(targetsSummary(detached)).toEqual({ timeline: false, count: 4 });
  });

  // Fix round 1 (reviewer Important): `detached()` (settings.ts) must NOT clear
  // `targets_over_time` when a style-owned field like the dummy checkbox is toggled --
  // clearing it would silently drop the ramp from the run. So a ramp can still be present
  // on an encounter whose `style` is now "": still a timeline (the control must never show
  // a number the run will not use), but distinguishable from one still attached to its
  // style, because the note beside it has to stop crediting a style that is no longer set.
  it('still reads a ramp as a timeline once detached from its style, but marks it unattached', () => {
    const dungeon = applyFightStyle(DEFAULT_ENCOUNTER, 'dungeon');
    const detachedWithRamp = { ...dungeon, style: '' };
    expect(targetsSummary(detachedWithRamp)).toEqual({ timeline: true, first: 1, max: 5, attached: false });
  });
});
