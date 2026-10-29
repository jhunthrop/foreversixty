import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ActionLink from './ActionLink.svelte';

const WORDS = createRawSnippet(() => ({ render: () => 'See plans' }));

describe('ActionLink', () => {
  const body = render(ActionLink, { props: { href: '/premium', children: WORDS, 'data-testid': 'al' } }).body;

  it('is a real link with a trailing chevron', () => {
    expect(body).toMatch(/<a href="\/premium"/);
    expect(body.indexOf('See plans')).toBeLessThan(body.indexOf('data-glyph="chevron"'));
  });

  it('is 14 px 600 gold, 44 px tall and not underlined', () => {
    expect(body).toContain('text-[14px] font-semibold');
    expect(body).toContain('text-gold');
    expect(body).toContain('min-h-11');
    expect(body).not.toContain('underline');
  });

  it('lightens on hover and active, rings on focus, and moves the chevron 2 px', () => {
    expect(body).toContain('hover:text-gold-hover');
    expect(body).toContain('active:text-gold-hover');
    expect(body).toContain('focus-visible:outline-2');
    expect(body).toContain('group-hover:translate-x-0.5');
  });

  it('passes extra attributes through', () => {
    expect(body).toContain('data-testid="al"');
  });
});
