// web/src/components/SupporterChip.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SupporterChip from './SupporterChip.svelte';

describe('SupporterChip', () => {
  it('renders the one word with its testid', () => {
    const { body } = render(SupporterChip);
    expect(body).toContain('data-testid="supporter-chip"');
    expect(body).toContain('Supporter');
  });

  it('explains itself on hover, sets the warm border explicitly and links nowhere', () => {
    const { body } = render(SupporterChip);
    expect(body).toContain('title="Supports Forever Sixty (Premium or guild plan)"');
    expect(body).toContain('border-color: var(--color-line-warm)');
    expect(body).not.toContain('<a');
  });
});
