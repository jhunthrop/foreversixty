// web/src/lib/items/lookup.test.ts
import { describe, expect, it } from 'vitest';
import { readItemTooltipModel } from './lookup';

const BUILD = '1.60.1.70009';

describe('readItemTooltipModel', () => {
  it('reads a plain armor row straight off the real build files', () => {
    const model = readItemTooltipModel(BUILD, 'warrior', 20143);
    expect(model?.name).toBe('90 Epic Warrior Neck');
    expect(model?.slotLabel).toBe('Neck');
    expect(model?.stats).toContain('+18 Strength');
  });

  it('joins the item to its set name via set_id', () => {
    const model = readItemTooltipModel(BUILD, 'warrior', 226857);
    expect(model?.setName).toBe('Battlegear of Heroism');
  });

  it('is null for an item id the class file does not carry', () => {
    expect(readItemTooltipModel(BUILD, 'warrior', 999999999)).toBeNull();
  });

  it('is null for a build with no item file for the class', () => {
    expect(readItemTooltipModel('no-such-build', 'warrior', 20143)).toBeNull();
  });
});
