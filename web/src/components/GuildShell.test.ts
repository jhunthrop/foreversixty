import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildShell from './GuildShell.svelte';

// Guild.svelte only reaches `status === 'ready'`/`'missing'` (and their testids) inside an
// $effect, which svelte/server's render() never runs (see Guild.test.ts's own comment) --
// so the two cases that delegate to Guild.svelte are asserted against its real, documented
// pre-effect SSR output (the "guild-skeleton" testid, same as Guild.test.ts) plus the
// absence of the other three views' testids, rather than a testid Guild.svelte cannot
// produce in this render path.
describe('GuildShell', () => {
  it('renders the plain guild view (Guild.svelte) when given a plain guild path', () => {
    const { body } = render(GuildShell, {
      props: { path: '/guild/us/hardcore/the-last-watch' },
    });
    expect(body).toContain('data-testid="guild-skeleton"');
    expect(body).not.toContain('data-testid="guild-claim"');
    expect(body).not.toContain('data-testid="guild-settings"');
    expect(body).not.toContain('data-testid="guild-join"');
  });

  it('renders the claim view for a /claim path', () => {
    const { body } = render(GuildShell, {
      props: { path: '/guild/us/hardcore/the-last-watch/claim' },
    });
    expect(body).toContain('data-testid="guild-claim"');
  });

  it('renders the settings view for a /settings path', () => {
    const { body } = render(GuildShell, {
      props: { path: '/guild/us/hardcore/the-last-watch/settings' },
    });
    expect(body).toContain('data-testid="guild-settings"');
  });

  it('renders the join view for an invite path', () => {
    const { body } = render(GuildShell, { props: { path: '/guild/invite/abc-123' } });
    expect(body).toContain('data-testid="guild-join"');
  });

  it('falls through to Guild.svelte (which resolves the missing state itself) for anything else', () => {
    const { body } = render(GuildShell, { props: { path: '/guild/not-enough-segments' } });
    expect(body).toContain('data-testid="guild-skeleton"');
    expect(body).not.toContain('data-testid="guild-claim"');
    expect(body).not.toContain('data-testid="guild-settings"');
    expect(body).not.toContain('data-testid="guild-join"');
  });
});
