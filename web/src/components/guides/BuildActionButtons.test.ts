// web/src/components/guides/BuildActionButtons.test.ts
//
// The task-7 brief guessed a single web/tests/astro/guides-components.test.ts might be
// needed for the three new .astro presentational components, but only if no existing
// pattern for directly rendering .astro components existed. It does: ClassTile.test.ts,
// FeedRow.test.ts and ScopeNote.test.ts already render .astro components through Astro's
// own experimental container API, co-located next to the component they test. This file
// applies that same existing harness to BuildActionButtons.astro rather than inventing a
// new location or skipping coverage.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { guidesCopy } from '../../lib/guides/copy';
import { loadBuildHref, simBuildHref } from '../../lib/guides/build-links';
import BuildActionButtons from './BuildActionButtons.astro';

const CODE = 'FS1:1.60.1.69893:warrior:human:1/2/3:';

describe('BuildActionButtons', () => {
  it('links "Load this build" to the planner href for the code', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(BuildActionButtons, { props: { code: CODE } });
    expect(html).toContain(`href="${loadBuildHref(CODE)}"`);
    expect(html).toContain(guidesCopy.loadThisBuild);
    expect(html).toContain('data-testid="guide-load-build"');
  });

  it('links "Sim this build" to the quick-sim href for the code', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(BuildActionButtons, { props: { code: CODE } });
    expect(html).toContain(`href="${simBuildHref(CODE)}"`);
    expect(html).toContain(guidesCopy.simThisBuild);
    expect(html).toContain('data-testid="guide-sim-build"');
  });

  it('gives both links a 44px hit target at every breakpoint', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(BuildActionButtons, { props: { code: CODE } });
    const anchors = html.match(/<a[^>]*>/g) ?? [];
    expect(anchors).toHaveLength(2);
    // SECONDARY_BUTTON_FIXED (lib/planner/styles.ts): h-11 on every breakpoint, never the
    // plain SECONDARY_BUTTON's md:h-9 desktop shrink -- round-1 fix (polish) swapped the
    // old sentence-case `min-h-11` class string for the design system's shared secondary
    // button recipe (uppercase 12px 700, 0.06em tracking), which keeps this guarantee.
    for (const anchor of anchors) {
      expect(anchor).toContain('h-11');
      expect(anchor).not.toContain('md:h-9');
      expect(anchor).toContain('uppercase');
    }
  });
});
