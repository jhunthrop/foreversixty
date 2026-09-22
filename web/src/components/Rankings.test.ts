// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Rankings from './Rankings.svelte';

describe('Rankings loading state', () => {
  it('reserves height with a Skeleton instead of a bare "Loading rankings." line', () => {
    const { body } = render(Rankings, { props: { slug: 'warden-kelthas' } });
    expect(body).toContain('data-testid="rankings-skeleton"');
    expect(body).not.toContain('Loading rankings.');
  });
});
