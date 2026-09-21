// web/src/components/GuildSettings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildSettings from './GuildSettings.svelte';
import { guildSettingsCopy } from '../lib/guild/copy';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'the-last-watch' };

describe('GuildSettings', () => {
  it('renders the guild-settings testid with a loading state', () => {
    const { body } = render(GuildSettings, { props: { path: PATH } });
    expect(body).toContain('data-testid="guild-settings"');
    expect(body).toContain(guildSettingsCopy.loading);
  });

  it('renders no claim-state or frozen testids before settings resolve', () => {
    const { body } = render(GuildSettings, { props: { path: PATH } });
    expect(body).not.toContain('data-testid="guild-settings-claim-state"');
    expect(body).not.toContain('data-testid="guild-settings-frozen"');
  });
});
