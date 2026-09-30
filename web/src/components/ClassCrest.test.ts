import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import ClassCrest from './ClassCrest.astro';

describe('ClassCrest', () => {
  it('renders the class icon at the given size, ringed in the class colour', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassCrest, { props: { slug: 'hunter', size: 56 } });
    expect(html).toContain('src="/icons/hd/crests/hunter.png"');
    expect(html).toContain('width="56"');
    expect(html).toContain('height="56"');
    expect(html).toContain('--c: #aad372');
    expect(html).toContain('class="crest rounded-full ');
  });

  it('falls back to white for an unknown slug rather than guessing a colour', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassCrest, { props: { slug: 'not-a-class', size: 36 } });
    expect(html).toContain('--c: #ffffff');
  });

  it('accepts every one of the five sizes this page uses', async () => {
    const c = await AstroContainer.create();
    for (const size of [36, 44, 56, 64, 84] as const) {
      const html = await c.renderToString(ClassCrest, { props: { slug: 'mage', size } });
      expect(html).toContain(`width="${size}"`);
    }
  });

  it('is lazy-loaded (home spec: crests lazy below the fold)', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassCrest, { props: { slug: 'druid', size: 64 } });
    expect(html).toContain('loading="lazy"');
  });
});
