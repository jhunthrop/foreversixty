// web/src/lib/items/lookup.test.ts
import { describe, expect, it } from 'vitest';
import { readItemTooltipModel } from './lookup';

const BUILD = '1.60.1.70291';

describe('readItemTooltipModel', () => {
  it('reads a plain armor row straight off the real build files', () => {
    const model = readItemTooltipModel(BUILD, 'warrior', 20143);
    expect(model?.name).toBe('90 Epic Warrior Neck');
    expect(model?.slotLabel).toBe('Neck');
    expect(model?.stats).toContain('+18 Strength');
  });

  it('joins the item to its set name via set_id', () => {
    const model = readItemTooltipModel(BUILD, 'warrior', 21995);
    expect(model?.setName).toBe('Battlegear of Heroism');
  });

  it('is null for an item id the class file does not carry', () => {
    expect(readItemTooltipModel(BUILD, 'warrior', 999999999)).toBeNull();
  });

  it('is null for a build with no item file for the class', () => {
    expect(readItemTooltipModel('no-such-build', 'warrior', 20143)).toBeNull();
  });

  it('reads the same item twice without the second call producing a different answer -- the in-process cache stays correct across repeat calls (a real /bis page calls this once per slot per band per faction)', () => {
    const first = readItemTooltipModel(BUILD, 'warrior', 20143);
    const second = readItemTooltipModel(BUILD, 'warrior', 20143);
    expect(second).toEqual(first);
    // A different item in the same class file still resolves correctly off the cached array.
    expect(readItemTooltipModel(BUILD, 'warrior', 21995)?.name).toBe('Boots of Heroism');
  });
});
