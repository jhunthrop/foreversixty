// web/src/lib/current-character-layout.test.ts
import { describe, expect, it } from 'vitest';
import { CHIP_HEIGHT } from './current-character-layout';

describe('CHIP_HEIGHT', () => {
  it('is the two-row-phone, one-row-desktop height every chip and its reserved slot share', () => {
    expect(CHIP_HEIGHT).toBe('h-[88px] md:h-11');
  });
});
