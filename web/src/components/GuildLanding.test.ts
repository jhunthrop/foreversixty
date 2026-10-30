// web/src/components/GuildLanding.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { guildLandingCopy } from '../lib/guild/copy';
import GuildLanding from './GuildLanding.svelte';

describe('GuildLanding', () => {
  const { body } = render(GuildLanding, { props: {} });

  it('has the Guilds heading, the lead and a Top guilds section', () => {
    expect(body).toMatch(/<h1[^>]*>Guilds<\/h1>/);
    expect(body).toContain(guildLandingCopy.lead);
    expect(body).toMatch(/<h2[^>]*>(?:<!--.*?-->|\s)*Top guilds(?:<!--.*?-->|\s)*<\/h2>/);
    expect(body).toContain('home-top-guilds');
  });

  it('renders a labelled lookup form with region and ruleset defaulted', () => {
    expect(body).toContain('role="search"');
    expect(body).toContain('data-testid="guild-lookup"');
    expect(body).toContain('Guild name');
    expect(body).toContain('Open guild');
    expect(body).toMatch(/<option[^>]*value="us"[^>]*selected/);
    expect(body).toMatch(/<option[^>]*value="normal"[^>]*selected/);
    expect(body).toContain('>Hardcore<');
  });

  it('has a live status line and the officer line', () => {
    expect(body).toContain('role="status"');
    expect(body).toContain(guildLandingCopy.officerLine);
  });

  it('puts the #claim anchor on the claim sentence, where the header’s Claim your guild lands', () => {
    expect(body).toMatch(/<p id="claim"[^>]*>\s*Run a guild\?/);
  });
});
