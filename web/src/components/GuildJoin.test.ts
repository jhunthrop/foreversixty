// web/src/components/GuildJoin.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildJoin from './GuildJoin.svelte';
import { guildJoinCopy } from '../lib/guild/copy';

describe('GuildJoin', () => {
  it('renders the guild-join testid with a loading state', () => {
    const { body } = render(GuildJoin, { props: { token: 'abc-123' } });
    expect(body).toContain('data-testid="guild-join"');
    expect(body).toContain(guildJoinCopy.loading);
  });
});
