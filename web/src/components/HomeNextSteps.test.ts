import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeNextSteps from './HomeNextSteps.svelte';

describe('HomeNextSteps', () => {
  it('renders only the empty, occlusion-ready root before the session resolves (SSR)', () => {
    const { body } = render(HomeNextSteps, { props: {} });
    expect(body).toContain('[grid-area:1/1]');
    expect(body).not.toContain('data-testid="home-next-steps"');
  });
});
