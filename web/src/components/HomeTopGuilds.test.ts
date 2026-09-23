// web/src/components/HomeTopGuilds.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeTopGuilds from './HomeTopGuilds.svelte';

describe('HomeTopGuilds loading state', () => {
  it('reserves height with a five-row Skeleton, not the ready list markup', () => {
    const { body } = render(HomeTopGuilds, { props: {} });
    expect(body).toContain('data-testid="home-top-guilds-skeleton"');
    expect(body).not.toContain('data-testid="home-top-guilds"');
  });
});
