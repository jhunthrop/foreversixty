// web/src/components/sim/ScopeNote.test.ts
//
// Task 9's own brief said Astro partials under src/components/sim aren't unit-tested here
// -- true of this directory specifically (no sibling .test.ts existed for ScopeNote.astro
// or SimTabs.astro before this file), but not true of the codebase as a whole: several
// top-level components (Panel.astro, SourcesList.astro, Footer.astro, StatePanel.astro,
// SourcePill.astro, FeedRow.astro) already have one, rendered through
// Astro's own experimental container API. That is a real, existing harness, not an
// invented one, so this file applies it here for the first time rather than leaving the
// new below-60 line covered only by a later e2e task.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { simCopy } from '../../lib/sim/copy';
import ScopeNote from './ScopeNote.astro';

describe('ScopeNote', () => {
  it('still renders the damage-specs-only line', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(ScopeNote);
    expect(html).toContain(simCopy.scopeNote);
  });

  it('renders one note and no below-60 framing line', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(ScopeNote);
    expect(html).toContain('data-testid="sim-scope-note"');
    expect(html).not.toContain('sim-below-sixty-note');
  });
});
