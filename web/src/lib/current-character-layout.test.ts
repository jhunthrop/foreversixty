// web/src/lib/current-character-layout.test.ts
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { CHIP_HEIGHT } from './current-character-layout';

describe('chip-slot pre-paint rule', () => {
  it('reserves the slot while a pointer is stored or a session hint is set', () => {
    const css = readFileSync(new URL('../styles/global.css', import.meta.url), 'utf-8');
    expect(css).toContain(
      "html:not([data-pointer='1']):not([data-session='1']) .chip-slot {\n  display: none;\n}",
    );
  });

  it('CHIP_HEIGHT is the two-tier phone/desktop shape the spine bar and the chip both reuse', () => {
    expect(CHIP_HEIGHT).toBe('h-[88px] md:h-11');
  });
});
