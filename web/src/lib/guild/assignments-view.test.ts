import { describe, expect, it } from 'vitest';
import { MAGE, OWN_HEALER, TANK, onyxia, picked, slot } from '../../test-support/assignment-fixtures';
import {
  anyPicked,
  assignmentsCopy,
  encountersOf,
  filledLine,
  holdsSlot,
  isFull,
  nameTone,
  notInClientOf,
  ownJobs,
  pickerOptions,
  raidWarningText,
  slotSub,
  withAdded,
  withRemoved,
} from './assignments-view';

describe('assignments view', () => {
  it('reads the block, empty for a run without one', () => {
    expect(encountersOf(onyxia())).toHaveLength(1);
    expect(encountersOf(null)).toEqual([]);
    expect(notInClientOf(undefined)).toEqual([]);
    expect(notInClientOf(onyxia({ encounters: [], not_in_client: ['Barrow Deeps'] }))).toEqual([
      'Barrow Deeps',
    ]);
  });

  it('tones the viewer gold, an off-role pick amber, the rest muted', () => {
    expect(nameTone(picked(OWN_HEALER))).toBe('gold');
    expect(nameTone(picked(MAGE, true))).toBe('warn');
    expect(nameTone(picked(TANK))).toBe('muted');
  });

  it('says how full a slot is and what it does', () => {
    const s = slot({ names: [picked(TANK)] });
    expect(filledLine(s)).toBe('1 of 1');
    expect(slotSub(s)).toBe('1 of 1 · Picks her up and holds her side-on to the raid in phases 1 and 3.');
    expect(isFull(s)).toBe(true);
    expect(isFull(slot())).toBe(false);
  });

  it('lists the raiders who fit the slot first, never one already in it', () => {
    const options = pickerOptions(slot({ count: 2, names: [picked(TANK)] }), [TANK, OWN_HEALER, MAGE]);
    expect(options.map((o) => [o.label, o.group])).toEqual([
      ['Mendara', assignmentsCopy.others],
      ['Frostquill', assignmentsCopy.others],
    ]);
    expect(options[0]).toMatchObject({ you: true, colour: '--color-class-priest' });
    expect(options[1]?.meta).toBe('Late');
    const healers = pickerOptions(slot({ role: 'healer', count: 4 }), [TANK, OWN_HEALER]);
    expect(healers.map((o) => [o.label, o.group])).toEqual([
      ['Mendara', assignmentsCopy.fits],
      ['Bulwark', assignmentsCopy.others],
    ]);
    expect(pickerOptions(slot({ role: 'any', count: 3 }), [TANK]).map((o) => o.group)).toEqual([undefined]);
  });

  it('adds at the end and removes by key, never changing the slot', () => {
    const s = slot({ count: 3, names: [picked(TANK)] });
    expect(withAdded(s, MAGE.character_key)).toEqual([TANK.character_key, MAGE.character_key]);
    expect(withAdded(s, TANK.character_key)).toEqual([TANK.character_key]);
    expect(withRemoved(s, TANK.character_key)).toEqual([]);
    expect(s.names).toHaveLength(1);
  });

  it('joins the /rw lines and names the viewer’s jobs', () => {
    const [ony] = encountersOf(onyxia());
    expect(raidWarningText(ony!)).toBe('MT: Bulwark\nMT heals: Mendara');
    expect(ownJobs(ony!)).toEqual(['Main tank healers']);
    expect(anyPicked(ony!)).toBe(true);
    expect(holdsSlot(ony!.slots[0]!)).toBe(false);
    expect(anyPicked({ ...ony!, slots: [slot()] })).toBe(false);
  });

  it('writes no em dash', () => {
    for (const value of Object.values(assignmentsCopy)) {
      const line = typeof value === 'function' ? value('X') : value;
      expect(line).not.toContain('—');
    }
  });
});
