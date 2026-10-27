// @vitest-environment jsdom
import { afterEach, describe, expect, it } from 'vitest';
import type { Me } from './api';
import { discordLinkCopy } from './discord-link-copy';
import { discordLinkReturn, takeDiscordLinkFlag } from './discord-link-return';
import { discordLinkHint } from './layout';

afterEach(() => window.history.replaceState({}, '', '/'));

describe('discordLinkReturn', () => {
  it('maps each flag', () => {
    expect(discordLinkReturn('linked')).toMatchObject({
      message: discordLinkCopy.returnLinked,
      role: 'status',
      focus: '[data-testid="discord-unlink"]',
      navigate: false,
    });
    expect(discordLinkReturn('denied')).toMatchObject({
      role: 'status',
      message: discordLinkCopy.returnDenied,
    });
    expect(discordLinkReturn('failed')).toMatchObject({
      role: 'alert',
      message: discordLinkCopy.returnFailed,
    });
    expect(discordLinkReturn('conflict')).toMatchObject({
      role: 'alert',
      message: discordLinkCopy.returnConflict,
      help: 'Sign in to that account and unlink it on its account page, then try again.',
      focus: '[data-testid="discord-link-start"]',
    });
    expect(discordLinkReturn('start').navigate).toBe(true);
  });

  it('only a linked return holds for the fresh me', () => {
    expect(discordLinkReturn('linked').refresh).toBe(true);
    expect(discordLinkReturn('conflict').refresh).toBe(false);
  });
});

describe('takeDiscordLinkFlag', () => {
  it('reads the flag once, keeps the rest of the address, and drops it', () => {
    window.history.replaceState({}, '', '/account?discord_link=conflict&x=1#you');
    expect(takeDiscordLinkFlag()).toBe('conflict');
    expect(window.location.pathname + window.location.search + window.location.hash).toBe('/account?x=1#you');
    expect(takeDiscordLinkFlag()).toBeNull();
  });

  it('an unknown value is dropped and reads as none', () => {
    window.history.replaceState({}, '', '/account?discord_link=nope#you');
    expect(takeDiscordLinkFlag()).toBeNull();
    expect(window.location.search).toBe('');
  });
});

describe('discordLinkHint', () => {
  const me = (extra: Partial<Me>): Me => ({
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    ...extra,
  });

  it('reads the hint from me', () => {
    expect(discordLinkHint(null)).toBe('none');
    expect(discordLinkHint(me({}))).toBe('none');
    expect(discordLinkHint(me({ discord_link_available: false }))).toBe('none');
    expect(discordLinkHint(me({ discord_link_available: true }))).toBe('unlinked');
    expect(discordLinkHint(me({ discord_link_available: true, discord_link: null }))).toBe('unlinked');
    expect(discordLinkHint(me({ discord_link_available: true, discord_link_suggested: 'Olympus' }))).toBe(
      'suggested',
    );
    expect(
      discordLinkHint(
        me({
          discord_link_available: true,
          discord_link: { display_name: 'zulmara' },
          discord_link_suggested: 'X',
        }),
      ),
    ).toBe('linked');
  });
});
