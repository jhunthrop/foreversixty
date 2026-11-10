import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Tag from './Tag.svelte';

const WORDS = createRawSnippet(() => ({ render: () => 'Founder' }));

function tag(props: Record<string, unknown> = {}): string {
  return render(Tag, { props: { children: WORDS, ...props } }).body;
}

describe('Tag', () => {
  it('is 22 px, 12 px 600 in sentence case with a 1 px border', () => {
    const body = tag();
    expect(body).toContain('h-[22px]');
    expect(body).toContain('text-[12px] font-semibold tracking-[.01em]');
    expect(body).not.toContain('uppercase');
    expect(body).toMatch(/ border /);
    expect(body).toContain('rounded-field');
  });

  it.each([
    ['gold', 'text-gold'],
    ['success', 'text-gain'],
    ['muted', 'text-muted'],
    ['danger', 'text-wipe'],
    ['warn', 'text-warn'],
    ['premium', 'bg-teal-edge'],
  ])('%s tone uses %s', (tone, ink) => {
    const body = tag({ tone });
    expect(body).toContain(ink);
    expect(body).toContain(`data-tone="${tone}"`);
  });

  it('warn is a 12% fill with a 40% border', () => {
    const body = tag({ tone: 'warn' });
    expect(body).toContain('bg-warn/12');
    expect(body).toContain('border-warn/40');
  });

  it('defaults to gold', () => {
    expect(tag()).toContain('data-tone="gold"');
  });

  it('draws a glowing 7 px dot only when asked', () => {
    expect(tag({ dot: true })).toContain('h-[7px] w-[7px]');
    expect(tag()).not.toContain('tag-dot');
  });

  it('carries a title and a testid', () => {
    const body = tag({ title: 'Founder price', testid: 'founder' });
    expect(body).toContain('title="Founder price"');
    expect(body).toContain('data-testid="founder"');
  });
});
