// web/src/components/HomeGuildCard.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeGuildCard from './HomeGuildCard.svelte';

describe('HomeGuildCard loading state', () => {
  it('reserves height with a Skeleton before the session resolves (SSR)', () => {
    const { body } = render(HomeGuildCard, { props: {} });
    expect(body).toContain('data-testid="home-guild-card-skeleton"');
    expect(body).not.toContain('data-testid="home-guild-card"');
  });
});
