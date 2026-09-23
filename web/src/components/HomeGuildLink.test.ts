// web/src/components/HomeGuildLink.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeGuildLink from './HomeGuildLink.svelte';

describe('HomeGuildLink loading state', () => {
  it('reserves height with a Skeleton before the session resolves (SSR)', () => {
    const { body } = render(HomeGuildLink, { props: {} });
    expect(body).toContain('data-testid="home-guild-skeleton"');
    expect(body).not.toContain('data-testid="home-guild-claim"');
    expect(body).not.toContain('data-testid="home-guild-ready"');
  });
});
