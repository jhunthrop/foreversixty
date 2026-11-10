// web/src/components/guild/setup/SetupStepBody.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildSettingsData } from '../../../lib/guild/api';
import type { GuildSetup, SetupStepKey } from '../../../lib/guild/setup-api';
import { BUTTON_PRIMARY } from '../../../lib/ui/button';
import { setupFixture } from '../../../test-support/guild-setup-fixture';
import SetupStepBody from './SetupStepBody.svelte';

const SETTINGS = {
  default_visibility: 'public',
  officer_max_rank_index: 1,
  claimed_by: null,
  claim_pending: null,
  claim: { state: 'claimed', yours: true },
  invite: { rotated_at: '2026-11-09T00:00:00Z', url: '/guild/invite/abc123', made_by: null },
} as unknown as GuildSettingsData;

function body(key: SetupStepKey, setup: GuildSetup = setupFixture(), primary = false): string {
  const step = setup.steps.find((s) => s.key === key);
  if (step === undefined) throw new Error(key);
  return render(SetupStepBody, {
    props: {
      step,
      setup,
      settings: SETTINGS,
      guildId: 7,
      guildName: 'Olympus XXVII',
      guildPath: '/guild/us/pvp/olympus-xxvii',
      primary,
      busy: false,
      when: () => 'Tue 8 Dec, 8:00 pm',
      run: () => {},
      onConfirm: () => {},
      onSettings: () => {},
    },
  }).body;
}

describe('SetupStepBody', () => {
  it('officer ranks: the ranks, the counted officers, the input, Save and Keep', () => {
    const html = body('officer_ranks', setupFixture(), true);
    expect(html).toContain('Rank 0 · 1: Dunmaro');
    expect(html).toContain('Rank 3 · 12: Dawnhollow, Embersong, Frostquill +9');
    expect(html).toContain('Rank 3 · 12');
    expect(html).toContain('Ranks 0 to 1 count as officers: 3 members');
    expect(html).toContain('data-testid="setup-ranks-input"');
    expect(html).toContain('Keep 0 to 1');
  });

  it('the Discord step installs the bot and comes back to the checklist, the one fill when first', () => {
    const html = body('discord_bot', setupFixture({ officer_ranks: 'done', invite_link: 'done' }), true);
    expect(html).toMatch(/href="[^"]*\/v1\/guilds\/7\/discord\/bot\/install\?return=setup"/);
    expect(html.split(BUTTON_PRIMARY).length - 1).toBe(1);
    expect(html).toContain('/raid, /signup, /bis and /char');
  });

  it('a done step shows its facts and no controls', () => {
    const done = setupFixture({ officer_ranks: 'done', invite_link: 'done' });
    const html = body('officer_ranks', done);
    expect(html).toContain('count as officers');
    expect(html).not.toContain('setup-ranks-input');
  });

  it('the invite step pastes the exact guild chat line and offers I posted it', () => {
    const html = body('invite_link', setupFixture({}));
    expect(html).toMatch(
      /Join Olympus XXVII on Forever Sixty and sign up for raids: \S+\/guild\/invite\/abc123/,
    );
    expect(html).toContain('data-testid="setup-invite-posted"');
  });

  it('the addon step shows the synced count, the bar and the paste line', () => {
    const html = body('addon_sync');
    expect(html).toContain('9 of 24 members synced in the last 14 days · 12 needed');
    expect(html).toContain('data-testid="setup-sync-meter"');
    expect(html).toContain('https://foreversixty.gg/setup');
  });

  it('recruiting asks for the bot first when it is not installed, and links the pitch', () => {
    const html = body('recruiting');
    expect(html).toContain('Add the Discord bot first');
    expect(html).toContain('href="/guild/us/pvp/olympus-xxvii/recruit"');
  });
});
