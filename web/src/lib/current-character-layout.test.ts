// web/src/lib/current-character-layout.test.ts
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { CHIP_HEIGHT, VIEW_GAP } from './current-character-layout';

describe('the retired spine bar', () => {
  it('leaves no pre-paint slot rule behind', () => {
    const css = readFileSync(new URL('../styles/global.css', import.meta.url), 'utf-8');
    expect(css).not.toContain('.chip-slot');
  });
});

describe('CHIP_HEIGHT', () => {
  it('is the two-row-phone, one-row-desktop height every chip and its reserved slot share', () => {
    expect(CHIP_HEIGHT).toBe('h-[88px] md:h-11');
  });
});

describe('VIEW_GAP', () => {
  it('is the flex gap SimView.svelte and ToolsView.svelte both use between their blocks', () => {
    expect(VIEW_GAP).toBe('gap-[22px] md:gap-8');
  });
});
