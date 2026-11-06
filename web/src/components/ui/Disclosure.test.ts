// Disclosure (character spec 2026-11-06 C5, C8): a summary that opens its body in place, no box.
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Disclosure from './Disclosure.svelte';

const body = createRawSnippet(() => ({ render: () => '<p data-testid="body">Inside</p>' }));

function html(props: Parameters<typeof render<typeof Disclosure>>[1]['props']): string {
  return render(Disclosure, { props }).body.replace(/<!--[^>]*-->/g, '');
}

describe('Disclosure', () => {
  it('shows the summary and keeps the body out until it opens', () => {
    const out = html({ summary: 'Parts and trend', meta: 'five parts', children: body });
    expect(out).toContain('Parts and trend');
    expect(out).toContain('five parts');
    expect(out).not.toContain('data-testid="body"');
    expect(out).not.toMatch(/bg-raised|border-line|rounded-panel/);
  });

  it('renders the body when open', () => {
    const out = html({ summary: 'History', open: true, children: body, testid: 'd' });
    expect(out).toContain('data-testid="body"');
    expect(out).toMatch(/<details[^>]*open/);
  });
});
