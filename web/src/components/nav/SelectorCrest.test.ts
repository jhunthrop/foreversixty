// web/src/components/nav/SelectorCrest.test.ts
// The selector's crest (owner bug 2026-11-09: the header's character image went missing for a
// moment on every page change). The bar's crest replaces a pre-painted crest of the same
// image, so it must decode with the paint, never a frame later; the list's rows stay lazy.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SelectorCrest from './SelectorCrest.svelte';

const html = (props: { classSlug: string; size: number; lazy?: boolean }): string =>
  render(SelectorCrest, { props }).body;

describe('SelectorCrest', () => {
  it('paints the bar crest with the page, so the swap from the pre-painted crest never shows empty', () => {
    const body = html({ classSlug: 'warrior', size: 32 });
    expect(body).toContain('src="/icons/hd/crests/warrior.webp"');
    expect(body).toContain('decoding="sync"');
    expect(body).not.toContain('loading="lazy"');
  });

  it('keeps the list rows lazy and off the paint', () => {
    const body = html({ classSlug: 'mage', size: 36, lazy: true });
    expect(body).toContain('loading="lazy"');
    expect(body).toContain('decoding="async"');
  });

  it('draws the neutral disc, never a broken request, for a character with no class', () => {
    const body = html({ classSlug: '', size: 32 });
    expect(body).not.toContain('<img');
    expect(body).not.toContain('/crests/.webp');
  });
});
