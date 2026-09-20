import { describe, expect, it } from 'vitest';
// A plain .mjs so scripts/sync-rotations.mjs can import the same parser this test drives:
// the script runs under node before vitest ever starts, so the parser cannot be a .ts --
// the same reason ids-md.test.ts drives scripts/ids-md.mjs this way.
import { rotationNotesOf } from '../../../scripts/rotation-apl.mjs';

describe('rotationNotesOf', () => {
  it('carries only steps with a non-empty notes string, in priority-list order', () => {
    const apl = {
      rotation: {
        priorityList: [
          { action: { autocastOtherCooldowns: {} } },
          { notes: 'Frostbolt is the whole rotation.', action: { castSpell: {} } },
          { notes: '', action: { castSpell: {} } },
          { notes: '   ', action: { castSpell: {} } },
          { action: { castSpell: {} } },
          { notes: 'Refresh the debuff before it falls off.', action: { castSpell: {} } },
        ],
      },
    };

    expect(rotationNotesOf(apl)).toEqual([
      'Frostbolt is the whole rotation.',
      'Refresh the debuff before it falls off.',
    ]);
  });

  it('trims surrounding whitespace from a note', () => {
    const apl = {
      rotation: { priorityList: [{ notes: '  Spend every global.  ', action: {} }] },
    };
    expect(rotationNotesOf(apl)).toEqual(['Spend every global.']);
  });

  it('never reads the file-level notes field -- that is developer prose about measurement, not a step', () => {
    const apl = {
      notes: 'the engine lane measured a 2.0 percent DPS loss versus Frostbolt alone',
      rotation: {
        priorityList: [{ notes: 'Frostbolt is the whole rotation.', action: {} }],
      },
    };
    expect(rotationNotesOf(apl)).toEqual(['Frostbolt is the whole rotation.']);
  });

  it('answers an empty list for a rotation with no steps, or a spec (e.g. a tank/healer) with none at all', () => {
    expect(rotationNotesOf({ rotation: { priorityList: [] } })).toEqual([]);
    expect(rotationNotesOf({})).toEqual([]);
  });
});
