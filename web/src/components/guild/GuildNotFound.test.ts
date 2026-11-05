// web/src/components/guild/GuildNotFound.test.ts
// A guild address that names no guild reads a written line with where it looked and a way back to
// the lookup, never the API's bare "no such guild" and a retry (guild round 3).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildNotFound from './GuildNotFound.svelte';

describe('GuildNotFound', () => {
  it('names the region and ruleset it looked in and links back to the lookup', () => {
    const { body } = render(GuildNotFound, {
      props: { path: { region: 'us', ruleset: 'normal', slug: 'lifecycle-tset' } },
    });
    expect(body).toContain('No guild by that name on US Normal. Check the spelling and the ruleset.');
    expect(body).toContain('href="/guild"');
    expect(body).toContain('Look it up again');
    expect(body).not.toContain('Try again');
  });

  it('writes the ruleset as the game does', () => {
    const { body } = render(GuildNotFound, {
      props: { path: { region: 'eu', ruleset: 'pvp', slug: 'x' } },
    });
    expect(body).toContain('on EU PvP.');
  });
});
