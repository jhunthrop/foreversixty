// web/src/components/premium/PremiumExampleCard.test.ts
import { readFileSync } from 'node:fs';
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { serverRunCaption } from '../../lib/billing/plans-page-copy';
import type { ServerRunExample } from '../../lib/billing/premium-examples';
import { EXAMPLE_IMAGE_ON_ERROR } from '../../lib/billing/example-image';
import { createPageContainer } from '../../test-support/page-container';
import PremiumExampleCard from './PremiumExampleCard.astro';

const source = readFileSync(new URL('./PremiumExampleCard.astro', import.meta.url), 'utf8');

const VALID: ServerRunExample = {
  simId: 'abc123',
  characterName: 'Kraggor',
  level: 60,
  specLabel: 'Fury Warrior',
  preset: true,
  placeholder: false,
  combinations: 4812,
  ranAt: '2026-10-27T18:00:00Z',
};

async function render(example: ServerRunExample): Promise<{ html: string; doc: Document }> {
  const container = await createPageContainer();
  const html = await container.renderToString(PremiumExampleCard, { props: { example } });
  return { html, doc: new JSDOM(html).window.document };
}

describe('PremiumExampleCard', () => {
  it.each([[true], [false]])('renders label, picture, caption, link in order (preset=%s)', async (preset) => {
    const ex = { ...VALID, preset };
    const { doc } = await render(ex);
    const order = [...doc.querySelectorAll('[data-part]')].map((n) => n.getAttribute('data-part'));
    expect(order).toEqual(['label', 'image', 'caption', 'link']);
    const img = doc.querySelector('img');
    expect(img?.getAttribute('width')).toBe('440');
    expect(img?.getAttribute('height')).toBe('330');
    expect(img?.getAttribute('loading')).toBe('lazy');
    expect(img?.getAttribute('decoding')).toBe('async');
    expect(img?.getAttribute('alt')).toBe(serverRunCaption(ex));
    const types = [...doc.querySelectorAll('picture source')].map((s) => s.getAttribute('type'));
    expect(types).toEqual(['image/avif', 'image/webp']);
    const caption = doc.querySelector('[data-testid=premium-example-caption]');
    expect(caption?.textContent?.trim()).toBe(serverRunCaption(ex));
    expect(doc.querySelector('[data-testid=premium-example-link]')?.getAttribute('href')).toBe('/sim/abc123');
  });

  it('names no character for a preset run and names the character otherwise', async () => {
    const preset = (await render({ ...VALID, preset: true })).doc.body.textContent ?? '';
    expect(preset).toContain("A level-60 Fury Warrior in the site's preset gear");
    expect(preset).not.toContain('Kraggor');
    const own = (await render({ ...VALID, preset: false })).doc.body.textContent ?? '';
    expect(own).toContain('Kraggor, Fury Warrior');
  });

  it('sets the numbers in the caption in mono', async () => {
    const { doc } = await render(VALID);
    const mono = [...doc.querySelectorAll('[data-testid=premium-example-caption] .mono')].map(
      (n) => n.textContent,
    );
    expect(mono).toContain('4,812');
    // The whole date is one run in the number font, never `27` mono beside a plain `Oct`.
    expect(mono).toContain('27 Oct 2026');
  });

  it('hides the img on a failed load, so no glyph or alt text draws in the box (B4)', async () => {
    const { doc } = await render(VALID);
    expect(doc.querySelector('img')?.getAttribute('onerror')).toBe(EXAMPLE_IMAGE_ON_ERROR);
    expect(EXAMPLE_IMAGE_ON_ERROR).toBe("this.style.visibility='hidden'");
  });

  it('keeps the reserved box, caption and link when the image is absent', async () => {
    const { doc } = await render(VALID);
    doc.querySelector('picture')?.remove();
    const box = doc.querySelector('[data-part=image]');
    expect(box).not.toBeNull();
    expect(box?.className).toContain('image-box');
    expect(doc.querySelector('[data-testid=premium-example-caption]')).not.toBeNull();
    expect(doc.querySelector('[data-testid=premium-example-link]')).not.toBeNull();
  });

  it('reserves a 4 / 3 box and has no gold top highlight', () => {
    // Scoped CSS is not part of the container's HTML, so read the component source.
    expect(source.match(/aspect-ratio:\s*4\s*\/\s*3/g)).toHaveLength(1);
    expect(source).not.toMatch(/229,\s*185,\s*85|border-top-color|gold-top/);
  });
});
