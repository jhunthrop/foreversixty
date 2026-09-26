// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Rankings from './Rankings.svelte';

describe('Rankings', () => {
  it('mounts the spine bar as the first element', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    const rankingsIndex = body.indexOf('data-testid="rankings"');
    const barIndex = body.indexOf('data-testid="current-character-bar"');
    expect(barIndex).toBeGreaterThan(-1);
    expect(barIndex).toBeGreaterThan(rankingsIndex);
  });
});
