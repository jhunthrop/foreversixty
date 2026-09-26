// web/src/lib/current-character-layout.test.ts
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { CHIP_HEIGHT, VIEW_GAP } from './current-character-layout';

describe('chip-slot pre-paint rule', () => {
  it('reserves the slot while a pointer is stored or a session hint is set', () => {
    const css = readFileSync(new URL('../styles/global.css', import.meta.url), 'utf-8');
    expect(css).toContain(
      "html:not([data-pointer='1']):not([data-session='1']) .chip-slot {\n  display: none;\n}",
    );
  });
});

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
