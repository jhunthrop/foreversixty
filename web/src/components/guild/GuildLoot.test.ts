// web/src/components/guild/GuildLoot.test.ts
// The Loot tab's first paint (svelte/server, no effects): the skeleton while the nights load, and
// the shipped plan gate sentence leading for an officer without the Raid guild plan. The council
// itself is drawn by loot/*.svelte (loot.test.ts) and driven by loot-council-state.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { gateFactsFixture } from '../../test-support/gate-facts';
import GuildLoot from './GuildLoot.svelte';

function html(officer: boolean, tools = { officerTools: true, raidGuildTools: true }): string {
  return render(GuildLoot, {
    props: {
      guildId: 7,
      officer,
      ...tools,
      gateFacts: gateFactsFixture({
        guildTier: tools.officerTools && !tools.raidGuildTools,
        settingsHref: '/g/settings',
      }),
      onRefused: () => {},
    },
  }).body;
}

describe('GuildLoot', () => {
  it('reserves the tab with a skeleton while the nights load', () => {
    const body = html(true);
    expect(body).toContain('data-testid="guild-loot-tab"');
    expect(body).toContain('data-testid="guild-loot-skeleton"');
  });

  it('leads with the plan gate for an officer whose guild lacks the tools, never for a member', () => {
    expect(html(true, { officerTools: true, raidGuildTools: false })).toContain(
      'data-testid="guild-loot-plan-gate"',
    );
    expect(html(true, { officerTools: false, raidGuildTools: true })).toContain(
      'id="officer-tools-gate-loot"',
    );
    expect(html(true)).not.toContain('guild-loot-plan-gate');
    expect(html(false, { officerTools: false, raidGuildTools: false })).not.toContain('guild-loot-plan-gate');
  });
});
