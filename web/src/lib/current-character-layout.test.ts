// web/src/lib/current-character-layout.test.ts
import { describe, expect, it } from 'vitest';
import { CHIP_HEIGHT, VIEW_GAP } from './current-character-layout';

describe('CHIP_HEIGHT', () => {
  it('is the two-row-phone, one-row-desktop height every chip and its reserved slot share', () => {
    expect(CHIP_HEIGHT).toBe('h-[88px] md:h-11');
  });
});

describe('VIEW_GAP', () => {
  it('is the flex gap SimView.svelte and ToolsView.svelte both use after the chip', () => {
    expect(VIEW_GAP).toBe('gap-[22px] md:gap-8');
  });
});
