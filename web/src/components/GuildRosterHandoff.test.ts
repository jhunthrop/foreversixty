// web/src/components/GuildRosterHandoff.test.ts
// Static-render smoke test, same pattern as SessionNav.test.ts: svelte/server's render
// only ever sees the component's pre-$effect initial state, which is the loading
// placeholder — the interactive resolved states are covered by e2e (guild-home.spec.ts).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildRosterHandoff from './GuildRosterHandoff.svelte';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'thrallgar' };

describe('GuildRosterHandoff', () => {
  it('renders the guild-roster-handoff testid with an invisible loading placeholder', () => {
    const { body } = render(GuildRosterHandoff, { props: { path: PATH } });
    expect(body).toContain('data-testid="guild-roster-handoff"');
    expect(body).toContain('invisible');
    expect(body).not.toContain('href="/sim');
    expect(body).not.toContain('href="/planner');
  });
});
