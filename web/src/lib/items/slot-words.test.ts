// web/src/lib/items/slot-words.test.ts
import { describe, expect, it } from 'vitest';
import { slotLine } from './slot-words';

describe('slotLine', () => {
  it.each([
    ['Stormrage Cover', 1, 4, 2, 'Head · Leather'],
    ['Netherwind Crown', 1, 4, 1, 'Head · Cloth'],
    ["Dragonstalker's Helm", 1, 4, 3, 'Head · Mail'],
    ['Helm of Wrath', 1, 4, 4, 'Head · Plate'],
    ['Shard of the Scale', 12, 4, 0, 'Trinket'],
    ['Ancient Cornerstone Grimoire', 23, 4, 0, 'Held In Off-hand'],
    ['Deathbringer', 13, 2, 0, 'One-Hand · Axe'],
    ["Vis'kag the Bloodletter", 13, 2, 7, 'One-Hand · Sword'],
    ['Sapphiron Drape', 16, 4, 1, 'Back'],
    ["Eskhandar's Collar", 2, 4, 0, 'Neck'],
    ['Ring of Binding', 11, 4, 0, 'Finger'],
  ])('reads %s the way its tooltip does', (_name, inventoryType, itemClass, itemSubclass, want) => {
    expect(slotLine(inventoryType, itemClass, itemSubclass)).toBe(want);
  });

  it('keeps the client casing on hands', () => {
    expect(slotLine(21, 2, 7)).toBe('Main Hand · Sword');
    expect(slotLine(14, 4, 6)).toBe('Off Hand · Shield');
  });

  it('is empty for an item with no slot', () => {
    expect(slotLine(0, 12, 0)).toBe('');
  });
});
