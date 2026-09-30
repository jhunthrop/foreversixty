// web/src/lib/bis/slot-display-labels.test.ts
import { describe, expect, it } from 'vitest';
import { SLOT_LABELS } from '../planner/types';
import { SLOT_DISPLAY_LABELS } from './slot-display-labels';

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
