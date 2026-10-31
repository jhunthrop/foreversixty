// web/src/lib/sim/bulk-skeleton.test.ts
import { describe, expect, it } from 'vitest';
import { VIEW_GAP } from '../current-character-layout';
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

  // 2026-10-31: the "Jump to" spine bar is gone, so no skeleton reserves a band for it and
  // the character strip opens every skeleton.
  it('reserves no band for the retired spine bar; the character strip comes first', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html).not.toContain('current-character-bar');
      const visible = html.slice(html.indexOf('aria-hidden="true"'));
      expect(visible.indexOf('rounded-panel')).toBeLessThan(visible.indexOf('<ul'));
    }
  });

  // Fix round 1, Important #2: the visible wrapper holding the strip and everything
  // after it must carry the identical gap ToolsView.svelte's own root uses between the
  // same elements once hydrated -- otherwise the strip/grid/run-bar block paints 22px/32px
  // higher before hydration than after it, at the exact moment this skeleton is on screen
  // (the Astro shell's own set:html, and ToolsView.svelte's lazy fallback).
  it('wraps the body in ToolsView.svelte’s own view gap', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      const visible = html.slice(html.indexOf('aria-hidden="true"'));
      const wrapperTag = visible.slice(0, visible.indexOf('>') + 1);
      expect(wrapperTag).toContain(VIEW_GAP);
    }
  });
});
