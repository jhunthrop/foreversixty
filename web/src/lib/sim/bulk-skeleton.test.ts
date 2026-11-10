// web/src/lib/sim/bulk-skeleton.test.ts
import { describe, expect, it } from 'vitest';
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

  // The redesign (mock design/mocks/ws_simtools.py): the run bar's band, then the open page
  // column. No surface is drawn: every section on the page is open.
  it('opens on the run bar band and draws no surface', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html).toContain('h-[var(--bar-h)]');
      expect(html).not.toMatch(/rounded-plate|border-line|bg-raised|current-character-bar/);
    }
  });
});
