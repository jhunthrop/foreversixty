import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GUILD_LOADING } from '../lib/guild/layout';
import { GUILD_VIEW_MIN, RAID_BAND_MIN } from '../lib/guild/raid-layout';
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

  it("routes a /claim path to the claim chunk, holding the claim view's loading height", () => {
    const { body } = render(GuildShell, {
      props: { path: '/guild/us/hardcore/the-last-watch/claim' },
    });
    expect(body).toContain('data-testid="guild-claim-chunk-skeleton"');
    expect(body).toContain(GUILD_LOADING.claim.minHeight);
    expect(body).not.toContain('data-testid="guild-skeleton"');
  });

  it("routes a /settings path to the settings chunk, holding the settings view's loading height", () => {
    const { body } = render(GuildShell, {
      props: { path: '/guild/us/hardcore/the-last-watch/settings' },
    });
    expect(body).toContain('data-testid="guild-settings-chunk-skeleton"');
    expect(body).toContain(GUILD_LOADING.settings.minHeight);
  });

  it('routes a raid path to the raid chunk, holding the band height while it loads', () => {
    const { body } = render(GuildShell, { props: { path: '/guild/us/pvp/olympus-xxvii/raids/7' } });
    expect(body).toContain('data-testid="raid-chunk-skeleton"');
    expect(body).toContain(RAID_BAND_MIN);
    expect(body).not.toContain('data-testid="guild-skeleton"');
  });

  it("routes an invite path to the join chunk, holding the join view's loading height", () => {
    const { body } = render(GuildShell, { props: { path: '/guild/invite/abc-123' } });
    expect(body).toContain('data-testid="guild-join-chunk-skeleton"');
    expect(body).toContain(GUILD_LOADING.join.minHeight);
  });

  it('falls through to Guild.svelte (which resolves the missing state itself) for anything else', () => {
    const { body } = render(GuildShell, { props: { path: '/guild/not-enough-segments' } });
    expect(body).toContain('data-testid="guild-skeleton"');
    expect(body).not.toContain('data-testid="guild-claim"');
    expect(body).not.toContain('data-testid="guild-settings"');
    expect(body).not.toContain('data-testid="guild-join"');
  });

  it('holds every view at least one screen tall from the first paint, so the footer starts below the fold', () => {
    for (const path of ['/guild/us/pvp/olympus-xxvii', '/guild/us/pvp/olympus-xxvii/raids/7', null]) {
      const { body } = render(GuildShell, { props: { path } });
      expect(body).toContain(GUILD_VIEW_MIN);
    }
  });

  it('renders the landing for the bare /guild path and /guild/', () => {
    for (const path of ['/guild', '/guild/']) {
      const { body } = render(GuildShell, { props: { path } });
      expect(body).toContain('data-testid="guild-landing"');
      expect(body).not.toContain('data-testid="guild-skeleton"');
    }
  });

  it('keeps the missing branch for a malformed address such as /guild/eu', () => {
    const { body } = render(GuildShell, { props: { path: '/guild/eu' } });
    expect(body).not.toContain('data-testid="guild-landing"');
    expect(body).toContain('data-testid="guild-skeleton"');
  });
});
