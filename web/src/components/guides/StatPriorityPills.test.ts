// web/src/components/guides/StatPriorityPills.test.ts
// See BuildActionButtons.test.ts's header comment for why this file exists at all: the
// codebase already has a co-located AstroContainer render-test convention for .astro
// components (ClassTile.test.ts, FeedRow.test.ts, ScopeNote.test.ts), so this applies it
// here rather than deferring coverage to a later whole-page test.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import StatPriorityPills from './StatPriorityPills.astro';

describe('StatPriorityPills', () => {
  it('renders one numbered pill per stat, in the given order', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(StatPriorityPills, {
      props: { stats: ['Strength', 'Haste', 'Crit'] },
    });
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="stat-pill-0"');
    expect(html).toContain('data-testid="stat-pill-1"');
    expect(html).toContain('data-testid="stat-pill-2"');
    const pill0 = /data-testid="stat-pill-0"[\s\S]*?<\/li>/.exec(html)?.[0] ?? '';
    expect(pill0).toContain('>1<');
    expect(pill0).toContain('Strength');
  });

  it('renders no pills for an empty stat list', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(StatPriorityPills, { props: { stats: [] } });
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).not.toContain('data-testid="stat-pill-0"');
  });
});
