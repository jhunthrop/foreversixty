import { describe, expect, it } from 'vitest';
import { plannerCopy } from '../planner/copy';
import { simCopy } from '../sim/copy';
import { gearCopy } from '../gear/copy';
import { homeUpgradesCopy } from '../home-panel-copy';
import { itemNotInDataLine, itemNotInDataPhrase } from './unknown-item-copy';

describe('the worn-item-not-in-our-data phrase', () => {
  it('is one phrase on every surface that names worn gear', () => {
    expect(gearCopy.unknownItem).toBe(itemNotInDataPhrase);
    expect(homeUpgradesCopy.unknownItem).toBe(itemNotInDataPhrase);
    expect(plannerCopy.unknownItem(288887)).toBe(itemNotInDataLine(288887));
    expect(simCopy.unknownItem(288887)).toBe('Item 288887 · Not in our item data yet');
  });
});
