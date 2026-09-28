// web/src/components/HomeGuildCard.test.ts
// svelte/server's render only ever sees pre-$effect state, so these pin what the first
// paint shows: nothing at all without the session cookie's readable half (Lighthouse's
// signed-out run saw a skeleton appear and vanish, a layout shift over budget), and the
// height-reserving skeleton when there is one. The cookie arrives as a prop only because
// this runs under node, where there is no document to set one on.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeGuildCard from './HomeGuildCard.svelte';

describe('HomeGuildCard first paint', () => {
  it('draws nothing when no session cookie hints at a signed-in visitor', () => {
    const { body } = render(HomeGuildCard, { props: { sessionCookie: '' } });
    expect(body).not.toContain('data-testid="home-guild-card-skeleton"');
    expect(body).not.toContain('data-testid="home-guild-card"');
  });

  it('reserves the card height with a Skeleton when the session cookie is present', () => {
    const { body } = render(HomeGuildCard, { props: { sessionCookie: 'fs_csrf=hint' } });
    expect(body).toContain('data-testid="home-guild-card-skeleton"');
    expect(body).toContain('min-h-[168px]');
    expect(body).not.toContain('data-testid="home-guild-card"');
  });
});
