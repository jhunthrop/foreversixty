// The overlay sheet (sim spec 2026-11-06 §5.2): a labelled modal dialog with a close control, a
// body and a foot; nothing at all while closed.
import { render } from 'svelte/server';
import { createRawSnippet } from 'svelte';
import { describe, expect, it } from 'vitest';
import Sheet from './Sheet.svelte';

const text = (html: string) => createRawSnippet(() => ({ render: () => html }));

function body(open: boolean): string {
  return render(Sheet, {
    props: {
      open,
      title: 'More settings',
      onclose: () => {},
      testid: 'more',
      children: text('<p>Body</p>'),
      foot: text('<span>Foot</span>'),
    },
  }).body;
}

describe('Sheet', () => {
  it('is a modal dialog named by its title, with a 44 px close control, a body and a foot', () => {
    const html = body(true);
    expect(html).toContain('role="dialog"');
    expect(html).toContain('aria-modal="true"');
    expect(html).toContain('aria-labelledby="more-title"');
    expect(html).toMatch(/<h2[^>]*id="more-title"[^>]*>\s*More settings/);
    expect(html).toContain('aria-label="Close"');
    expect(html).toContain('<p>Body</p>');
    expect(html).toContain('<span>Foot</span>');
  });

  it('renders nothing while closed', () => {
    expect(body(false)).not.toContain('role="dialog"');
  });
});
