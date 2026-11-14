import { describe, expect, it } from 'vitest';
import type { ItemTooltipModel } from '../items/tooltip';
import { addTooltipModels, lateTooltipModel } from './page-registry';
import { sliceTooltips } from './page-server';
import type { BisPageData } from './page-model';

const tip = (id: number): ItemTooltipModel => ({ id, name: `Item ${id}` }) as unknown as ItemTooltipModel;

describe('sliceTooltips', () => {
  it("keeps the tooltips of the slice's items alone", () => {
    const slice = { items: { 1: {}, 3: {} } } as unknown as BisPageData;
    expect(sliceTooltips({ 1: tip(1), 2: tip(2), 3: tip(3) }, slice)).toEqual({ 1: tip(1), 3: tip(3) });
  });
});

describe('addTooltipModels', () => {
  it('makes the models it is given visible to every reader on the page', () => {
    expect(lateTooltipModel(42)).toBeUndefined();
    const all = addTooltipModels({ 42: tip(42) });
    expect(lateTooltipModel(42)).toEqual(tip(42));
    expect(all.get(42)).toEqual(tip(42));
  });
});
