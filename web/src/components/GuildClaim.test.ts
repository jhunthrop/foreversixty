// web/src/components/GuildClaim.test.ts
// Static-render smoke test: svelte/server's render only ever sees pre-$effect state
// (Guild.test.ts documents the same limitation), so this pins the loading placeholder the
// common, common case -- data not resolved yet -- renders on first paint.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { guildClaimCopy } from '../lib/guild/copy';
import GuildClaim from './GuildClaim.svelte';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'the-last-watch' };

describe('GuildClaim', () => {
  it('renders the guild-claim testid with a loading state before data resolves', () => {
    const { body } = render(GuildClaim, { props: { path: PATH } });
    expect(body).toContain('data-testid="guild-claim"');
    expect(body).toContain(guildClaimCopy.loading);
  });
});
