// web/src/components/ui/Crest.test.ts
// The one class crest, Svelte and Astro: the ringed HD crest in the class's ring, the LX seal in a
// muted ring for no class, lazy and low priority unless eager, sized inline unless the caller sizes it.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CrestAstro from './Crest.astro';
import Crest from './Crest.svelte';

const svelte = (props: {
  cls: string | null;
  size: number;
  eager?: boolean;
  class?: string;
  testid?: string;
}) => render(Crest, { props }).body;

describe('Crest', () => {
  it('draws the class crest ringed in its class colour, from a name or a slug', () => {
    for (const cls of ['hunter', 'Hunter']) {
      const body = svelte({ cls, size: 36, testid: 't' });
      expect(body).toContain('src="/icons/hd/crests/hunter.webp"');
      expect(body).toContain('class="crest');
      expect(body).toContain('--c: var(--color-class-hunter); width: 36px; height: 36px;');
      expect(body).toContain('data-testid="t"');
      expect(body).not.toContain('crest-none');
    }
  });

  it('draws the LX seal in a muted ring for no class, never a person glyph or a broken request', () => {
    const body = svelte({ cls: '', size: 36, testid: 't' });
    expect(body).toMatch(/src="\/art\/seal-[^"]+"/);
    expect(body).toContain('crest-none');
    expect(body).toContain('--c: var(--color-muted)');
    expect(body).toContain('data-testid="t-fallback"');
    expect(body).not.toContain('<svg');
    expect(body).not.toContain('/crests/.webp');
  });

  it('is lazy and low priority by default, eager and decoded with the paint when asked', () => {
    expect(svelte({ cls: 'mage', size: 28 })).toContain(
      'loading="lazy" fetchpriority="low" decoding="async"',
    );
    const eager = svelte({ cls: 'mage', size: 28, eager: true });
    expect(eager).toContain('loading="eager"');
    expect(eager).toContain('decoding="sync"');
    expect(eager).not.toContain('fetchpriority');
  });

  it("leaves the drawn size to the caller's class", () => {
    const body = svelte({ cls: 'mage', size: 32, class: 'csel-crest' });
    expect(body).toContain('class="crest csel-crest"');
    expect(body).not.toContain('width: 32px');
  });

  it('the Astro twin draws the same crest', async () => {
    const container = await AstroContainer.create();
    const known = await container.renderToString(CrestAstro, {
      props: { cls: 'hunter', size: 56, eager: true },
    });
    expect(known).toContain('src="/icons/hd/crests/hunter.webp"');
    expect(known).toContain('class="crest"');
    expect(known).toContain('--c: var(--color-class-hunter); width: 56px; height: 56px;');
    expect(known).toContain('loading="eager"');
    const unknown = await container.renderToString(CrestAstro, { props: { cls: 'not-a-class', size: 36 } });
    expect(unknown).toContain('crest-none');
    expect(unknown).toContain('fetchpriority="low"');
  });
});
