// web/src/lib/sim/bulk-skeleton.test.ts
import { describe, expect, it } from 'vitest';
import { CHIP_HEIGHT } from '../current-character-layout';
import { TOOL_SKELETONS } from './bulk-skeleton';
import { TOOLS } from './bulk-store.svelte';

describe('TOOL_SKELETONS', () => {
  it('has one skeleton per tool page', () => {
    expect(Object.keys(TOOL_SKELETONS).sort()).toEqual([...TOOLS].sort());
  });

  it('announces itself once to a screen reader and hides its blocks from it', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html.match(/role="status"/g) ?? []).toHaveLength(1);
      expect(html).toContain('aria-hidden="true"');
      expect(html).toContain('aria-busy="true"');
    }
  });

  it('puts no data in the markup, so the shell and the island render the same bytes', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html).not.toMatch(/\{|\$\{/);
    }
  });

  // Fix round 1, Task 4's review (Critical): every skeleton reserves the current-character
  // chip's own band, first, using the identical CHIP_HEIGHT constant the chip and
  // ToolsView.svelte's own always-present slot render with -- one shared reservation, not
  // four copies that could drift apart.
  it('opens every skeleton with the chip slot, reserving CHIP_HEIGHT before the strip', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html).toContain(`data-testid="sim-chip-slot"`);
      expect(html).toContain(CHIP_HEIGHT);
      // The chip slot is the very first thing inside the skeleton's aria-hidden body --
      // the strip band that used to open it now comes strictly after.
      const body = html.slice(html.indexOf('aria-hidden="true"'));
      expect(body.indexOf('sim-chip-slot')).toBeLessThan(body.indexOf('skeleton-block'));
    }
  });
});
