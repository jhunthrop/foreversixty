import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import SourcePill from './SourcePill.astro';

describe('SourcePill', () => {
  it('renders kind class and default label', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(SourcePill, { props: { kind: 'datamined' } });
    expect(html).toContain('pill-datamined');
    expect(html).toContain('Datamined');
  });
  it('accepts a label override', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(SourcePill, { props: { kind: 'blizzard', label: 'Hotfix' } });
    expect(html).toContain('Hotfix');
  });
});
