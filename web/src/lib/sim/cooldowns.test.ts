import { describe, expect, it } from 'vitest';
import { COOLDOWN_MODES, executeStartSec, modeOf, rowsFor, specFor, withCooldown } from './cooldowns';
import { simCopy } from './copy';
import { DEFAULT_ENCOUNTER } from './types';

const encounter = { ...DEFAULT_ENCOUNTER, duration_sec: 200, execute_ratio: 0.25 };

describe('executeStartSec', () => {
  it('is where the execute window opens: the fight’s length less its execute share', () => {
    expect(executeStartSec(encounter)).toBe(150);
  });

  it('is the end of the fight when there is no execute window at all', () => {
    expect(executeStartSec({ ...encounter, execute_ratio: 0 })).toBe(200);
  });
});

describe('specFor and modeOf', () => {
  it('turns each mode into the times the engine takes', () => {
    expect(specFor('major_mana_potion', 'on-cooldown', 0, encounter).at_sec).toEqual([]);
    expect(specFor('major_mana_potion', 'on-pull', 0, encounter).at_sec).toEqual([0]);
    expect(specFor('major_mana_potion', 'at-time', 42, encounter).at_sec).toEqual([42]);
    expect(specFor('major_mana_potion', 'at-execute', 0, encounter).at_sec).toEqual([150]);
  });

  it('clamps a hand-typed time into the fight', () => {
    expect(specFor('x', 'at-time', -5, encounter).at_sec).toEqual([0]);
    expect(specFor('x', 'at-time', 9999, encounter).at_sec).toEqual([200]);
  });

  it('reads a stored spec back as the mode that wrote it', () => {
    expect(modeOf({ id: 'x', at_sec: [] }, encounter)).toBe('on-cooldown');
    expect(modeOf({ id: 'x', at_sec: [0] }, encounter)).toBe('on-pull');
    expect(modeOf({ id: 'x', at_sec: [150] }, encounter)).toBe('at-execute');
    expect(modeOf({ id: 'x', at_sec: [42] }, encounter)).toBe('at-time');
  });

  it('reads a multi-time spec as a plain time rather than claiming a mode it is not', () => {
    expect(modeOf({ id: 'x', at_sec: [0, 90] }, encounter)).toBe('at-time');
  });

  it('offers the four modes the design names', () => {
    expect(COOLDOWN_MODES).toEqual(['on-cooldown', 'on-pull', 'at-time', 'at-execute']);
  });
});

describe('rowsFor', () => {
  it('is one row per id, defaulting to on cooldown', () => {
    expect(rowsFor(['a', 'b'], [], encounter)).toEqual([
      { id: 'a', mode: 'on-cooldown', atSec: 0 },
      { id: 'b', mode: 'on-cooldown', atSec: 0 },
    ]);
  });

  it('takes a row’s mode and time from the stored spec', () => {
    expect(rowsFor(['a'], [{ id: 'a', at_sec: [42] }], encounter)).toEqual([
      { id: 'a', mode: 'at-time', atSec: 42 },
    ]);
  });

  it('keeps a stored spec whose id is not offered, so a pasted request round-trips', () => {
    expect(rowsFor(['a'], [{ id: 'spell:11305', at_sec: [0] }], encounter)).toEqual([
      { id: 'a', mode: 'on-cooldown', atSec: 0 },
      { id: 'spell:11305', mode: 'on-pull', atSec: 0 },
    ]);
  });
});

describe('withCooldown', () => {
  it('drops the spec entirely for “on cooldown”, which is the engine’s own default', () => {
    expect(withCooldown([{ id: 'a', at_sec: [0] }], 'a', 'on-cooldown', 0, encounter)).toEqual([]);
  });

  it('replaces a spec rather than appending a second one for the same id', () => {
    const once = withCooldown([], 'a', 'on-pull', 0, encounter);
    const twice = withCooldown(once, 'a', 'at-time', 60, encounter);
    expect(twice).toEqual([{ id: 'a', at_sec: [60] }]);
  });

  it('does not mutate the list it was given', () => {
    const specs = [{ id: 'a', at_sec: [0] }];
    withCooldown(specs, 'a', 'at-time', 10, encounter);
    expect(specs).toEqual([{ id: 'a', at_sec: [0] }]);
  });
});

describe('the row controls’ accessible names', () => {
  // CooldownRows.svelte gives the mode select and the seconds input each a per-row name
  // built from these two helpers, mirroring gradeFor in BuffPanel.svelte. A shared,
  // row-independent string here is exactly the defect the review caught: several rows
  // would read as identical controls to a screen reader.
  it('names each control by the row it belongs to, not a string every row shares', () => {
    expect(simCopy.cooldownModeFor('Major mana potion')).toContain('Major mana potion');
    expect(simCopy.cooldownAtFor('Major mana potion')).toContain('Major mana potion');
    expect(simCopy.cooldownModeFor('Major mana potion')).not.toBe(simCopy.cooldownModeFor('Thistle tea'));
    expect(simCopy.cooldownAtFor('Major mana potion')).not.toBe(simCopy.cooldownAtFor('Thistle tea'));
  });
});
