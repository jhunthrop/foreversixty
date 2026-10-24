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
});
