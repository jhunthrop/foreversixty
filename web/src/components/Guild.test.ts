// web/src/components/Guild.test.ts
// Static-render smoke test: svelte/server's render only ever sees pre-$effect state, so
// this pins the "signed-in section renders nothing extra before the session/home queries
// resolve" contract -- the CLS-relevant claim from the plan's Global Constraints (nothing
// shifts for the common, signed-out visitor, because there is nothing there to begin
// with). The full signed-in home is covered by e2e (guild-home.spec.ts).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Guild from './Guild.svelte';
import { guildHomeCopy } from '../lib/guild/copy';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'the-last-watch' };

describe('Guild', () => {
  it('renders no guild-home testid before the session resolves', () => {
    const { body } = render(Guild, { props: { path: PATH } });
    expect(body).not.toContain('data-testid="guild-home"');
  });

  it('still renders the loading state as its initial static render (existing public-page behavior)', () => {
    const { body } = render(Guild, { props: { path: PATH } });
    expect(body).toContain('data-testid="guild-skeleton"');
  });

  it('the empty-roster copy names one action for an officer and none for a member', () => {
    expect(guildHomeCopy.manageInvite).toBe('Guild settings');
    expect(guildHomeCopy.emptyRosterMember).not.toMatch(/officer|invite/i);
  });
});
