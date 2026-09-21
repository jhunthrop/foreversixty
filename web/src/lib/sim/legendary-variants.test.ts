// web/src/lib/sim/legendary-variants.test.ts
import { describe, expect, it } from 'vitest';
import { legendaryVariantLabel } from './legendary-variants';

describe('legendaryVariantLabel', () => {
  it('labels each of the four Atiesh ids by the class its quest chain grants it to', () => {
    expect(legendaryVariantLabel(22589)).toBe('Mage');
    expect(legendaryVariantLabel(22630)).toBe('Warlock');
    expect(legendaryVariantLabel(22631)).toBe('Priest');
    expect(legendaryVariantLabel(22632)).toBe('Druid');
  });

  it('labels every other item with nothing', () => {
    expect(legendaryVariantLabel(12784)).toBe('');
  });
});
