// web/src/lib/bis/slot-display-labels.test.ts
import { describe, expect, it } from 'vitest';
import { SLOT_LABELS } from '../planner/types';
import { SLOT_DISPLAY_LABELS, isRelicSlot, relicNoun, slotDisplayLabel, slotLabelForClass } from './slot-display-labels';

describe('SLOT_DISPLAY_LABELS', () => {
  it('collapses both ring sockets and both trinket sockets to one word each', () => {
    expect(SLOT_DISPLAY_LABELS.finger1).toBe('Ring');
    expect(SLOT_DISPLAY_LABELS.finger2).toBe('Ring');
    expect(SLOT_DISPLAY_LABELS.trinket1).toBe('Trinket');
    expect(SLOT_DISPLAY_LABELS.trinket2).toBe('Trinket');
  });

  it('keeps every other slot the same as the disambiguating SLOT_LABELS', () => {
    for (const slot of Object.keys(SLOT_LABELS) as (keyof typeof SLOT_LABELS)[]) {
      if (['finger1', 'finger2', 'trinket1', 'trinket2'].includes(slot)) continue;
      expect(SLOT_DISPLAY_LABELS[slot]).toBe(SLOT_LABELS[slot]);
    }
  });
});

// Round 19 (wow-player MUST): the paladin, shaman and druid paperdoll calls the ranged slot
// "Relic" (librams, totems, idols); every other class keeps "Ranged" for its bow, gun,
// thrown weapon or wand.
describe('the relic slot', () => {
  it('reads Relic for the relic classes and Ranged for everyone else', () => {
    for (const classSlug of ['paladin', 'shaman', 'druid']) {
      expect(slotDisplayLabel('ranged', classSlug)).toBe('Relic');
      expect(slotLabelForClass('ranged', classSlug)).toBe('Relic');
      expect(isRelicSlot('ranged', classSlug)).toBe(true);
    }
    for (const classSlug of ['hunter', 'rogue', 'warrior', 'mage', 'priest', 'warlock', undefined]) {
      expect(slotDisplayLabel('ranged', classSlug)).toBe('Ranged');
      expect(slotLabelForClass('ranged', classSlug)).toBe('Ranged');
    }
  });

  it('leaves every other slot as it was', () => {
    expect(slotDisplayLabel('finger1', 'paladin')).toBe('Ring');
    expect(slotLabelForClass('finger1', 'paladin')).toBe('Finger 1');
    expect(slotDisplayLabel('head', 'druid')).toBe('Head');
    expect(slotDisplayLabel('not_a_slot', 'druid')).toBe('not_a_slot');
  });
});

describe('relicNoun', () => {
  it('names each relic class its own relic, and any other class a plain relic (round 20)', () => {
    expect(relicNoun('paladin')).toBe('libram');
    expect(relicNoun('druid')).toBe('idol');
    expect(relicNoun('shaman')).toBe('totem');
    expect(relicNoun('hunter')).toBe('relic');
    expect(relicNoun(undefined)).toBe('relic');
  });
});
