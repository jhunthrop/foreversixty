import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import FactionMark from './FactionMark.astro';

describe('FactionMark', () => {
  it('renders the alliance emblem at the given size, unboxed', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(FactionMark, { props: { faction: 'alliance', size: 20 } });
    expect(html).toContain('src="/icons/hd/faction/alliance.webp"');
    expect(html).toContain('width="20"');
    expect(html).toContain('height="20"');
    expect(html).not.toContain('border');
    expect(html).not.toContain('box-shadow');
  });

  it('renders the horde emblem', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(FactionMark, { props: { faction: 'horde', size: 16 } });
    expect(html).toContain('src="/icons/hd/faction/horde.webp"');
  });

  it('accepts every one of the three documented sizes', async () => {
    const c = await AstroContainer.create();
    for (const size of [16, 20, 36] as const) {
      const html = await c.renderToString(FactionMark, { props: { faction: 'horde', size } });
      expect(html).toContain(`width="${size}"`);
    }
  });

  it('carries no visible alt text (decorative; the word beside it says the faction)', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(FactionMark, { props: { faction: 'alliance', size: 36 } });
    expect(html).toContain('alt=""');
  });
});
