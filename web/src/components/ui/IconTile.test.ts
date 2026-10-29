import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import IconTile from './IconTile.svelte';

describe('IconTile', () => {
  it('is 36 px with an 18 px glyph by default, decorative', () => {
    const body = render(IconTile, { props: { glyph: 'crown' } }).body;
    expect(body).toContain('width: 36px');
    expect(body).toContain('height: 36px');
    expect(body).toContain('width="18"');
    expect(body).toContain('aria-hidden="true"');
    expect(body).toContain('rounded-control');
  });

  it('is 32 px with a 16 px glyph at the small size', () => {
    const body = render(IconTile, { props: { glyph: 'bell', size: 32 } }).body;
    expect(body).toContain('width: 32px');
    expect(body).toContain('width="16"');
  });

  it.each([
    ['gold', 'bg-tile text-tile-ink'],
    ['discord', 'bg-[rgba(88,101,242,.16)] text-[#9aa4ff]'],
    ['muted', 'text-muted'],
  ])('%s tone', (tone, classes) => {
    const body = render(IconTile, {
      props: { glyph: 'hash', tone: tone as 'gold' | 'discord' | 'muted' },
    }).body;
    expect(body).toContain(classes);
    expect(body).toContain(`data-tone="${tone}"`);
  });

  it('draws the Discord mark as a solid shape and the rest as 2 px strokes', () => {
    const solid = render(IconTile, { props: { glyph: 'discord' } }).body;
    expect(solid).toContain('fill="currentColor"');
    const stroked = render(IconTile, { props: { glyph: 'chart' } }).body;
    expect(stroked).toContain('stroke-width="2"');
    expect(stroked).toContain('fill="none"');
  });
});
