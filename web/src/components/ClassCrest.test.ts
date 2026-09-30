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

  it('is lazy-loaded by default (home spec: crests lazy below the fold)', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassCrest, { props: { slug: 'druid', size: 64 } });
    expect(html).toContain('loading="lazy"');
  });

  it('loads eagerly, but never at fetchpriority=high, when the caller marks it above the fold (review round 1 item 6)', async () => {
    // fetchpriority="high" measurably regressed the index page's LCP from ~1.1s to ~6.9s
    // against web/lighthouserc.json's own throttled mobile profile (see ClassCrest.astro's
    // own doc) -- eager alone is the fix; this pins fetchpriority never coming back.
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassCrest, { props: { slug: 'druid', size: 56, priority: true } });
    expect(html).toContain('loading="eager"');
    expect(html).not.toContain('fetchpriority');
  });
});
