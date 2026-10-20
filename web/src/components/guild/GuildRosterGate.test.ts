import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildRosterGate from './GuildRosterGate.svelte';

describe('GuildRosterGate', () => {
  it('is the #roster anchor and gives a visitor the members-only sentence, with no rows', () => {
    const { body } = render(GuildRosterGate, { props: { state: 'visitor' } });
    expect(body).toContain('id="roster"');
    expect(body).toContain('Roster is for members. Sign in with the account your addon syncs to.');
    expect(body).not.toContain('<table');
  });

  it('tells an unverified member what opens the roster', () => {
    const { body } = render(GuildRosterGate, { props: { state: 'unverified' } });
    expect(body).toContain('waiting for an officer to verify');
  });

  it('tells a signed-in stranger the roster is for the guild’s members', () => {
    const { body } = render(GuildRosterGate, { props: { state: 'stranger' } });
    expect(body).toContain('Roster is for members of this guild');
  });
});
