import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeGetSetUp from './HomeGetSetUp.svelte';

describe('HomeGetSetUp', () => {
  it('renders nothing before the session resolves (SSR): the static "Get set up" card stays the only visible thing', () => {
    const { body } = render(HomeGetSetUp, { props: {} });
    expect(body).not.toContain('data-testid="home-get-set-up"');
  });
});
