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

  it('tells a guild master how a guild that is not on the site yet gets here (day one)', () => {
    const { body } = render(GuildNotFound, {
      props: { path: { region: 'us', ruleset: 'normal', slug: 'knights-of-the-old-forest' } },
    });
    expect(body).toContain('Not on Forever Sixty yet.');
    expect(body).toContain('signs in and imports a character, or syncs the addon.');
    expect(body).toMatch(/href="\/setup"[^>]*>Get set up</);
  });

  it('offers a guild it knows whose name starts with what was typed (guild loop round 3)', () => {
    const { body } = render(GuildNotFound, {
      props: {
        path: { region: 'us', ruleset: 'pvp', slug: 'lifecycle' },
        candidates: [{ name: 'Lifecycle Test', region: 'us', ruleset: 'pvp' }],
      },
    });
    const line = body.slice(body.indexOf('data-testid="guild-not-found-suggestion"'));
    expect(line.replace(/<!--[\s\S]*?-->/g, '').replace(/<[^>]+>/g, '')).toMatch(
      /^[^>]*>\s*Did you mean Lifecycle Test \(US PvP\)\?/,
    );
    expect(line).toContain('href="/guild/us/pvp/lifecycle-test"');
  });

  it("offers a stranger the guild search's rows, up to three, joined with commas and or", () => {
    const { body } = render(GuildNotFound, {
      props: {
        path: { region: 'us', ruleset: 'normal', slug: 'lifecyle-tset' },
        candidates: [],
        searched: [
          { name: 'Lifecycle Test', region: 'us', ruleset: 'pvp' },
          { name: 'Lifecycle Testers', region: 'us', ruleset: 'normal' },
          { name: 'Lifecycle Tests', region: 'us', ruleset: 'hardcore' },
          { name: 'Lifecycle Fourth', region: 'us', ruleset: 'rp' },
        ],
      },
    });
    const line = body.slice(body.indexOf('data-testid="guild-not-found-suggestion"'));
    expect(line.replace(/<!--[\s\S]*?-->/g, '').replace(/<[^>]+>/g, '')).toMatch(
      /^[^>]*>\s*Did you mean Lifecycle Test \(US PvP\), Lifecycle Testers \(US Normal\) or Lifecycle Tests \(US Hardcore\)\?/,
    );
    expect(line).toContain('href="/guild/us/normal/lifecycle-testers"');
    expect(line).not.toContain('Lifecycle Fourth');
  });

  it('offers nothing when no known guild is close', () => {
    const { body } = render(GuildNotFound, {
      props: { path: { region: 'us', ruleset: 'pvp', slug: 'zzz' }, candidates: [], searched: [] },
    });
    expect(body).not.toContain('guild-not-found-suggestion');
  });

  it('writes the ruleset as the game does', () => {
    const { body } = render(GuildNotFound, {
      props: { path: { region: 'eu', ruleset: 'pvp', slug: 'x' } },
    });
    expect(body).toContain('on EU PvP.');
  });
});
