// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { SESSION_GUILD_EVENT } from './config';
import { publishSessionGuild, readPublishedSessionGuild } from './session-guild-dom';

const guild = { name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' };

function mountHeader(): HTMLElement {
  document.body.innerHTML = '<header data-site-header></header>';
  return document.querySelector<HTMLElement>('[data-site-header]')!;
}

afterEach(() => {
  document.body.innerHTML = '';
});

describe('publishSessionGuild', () => {
  it('stamps the guild on the header and dispatches the event with it', () => {
    const header = mountHeader();
    const heard = vi.fn();
    window.addEventListener(SESSION_GUILD_EVENT, (event) => heard((event as CustomEvent).detail));
    publishSessionGuild(guild);
    expect(readPublishedSessionGuild(header)).toEqual(guild);
    expect(heard).toHaveBeenCalledWith(guild);
  });

  it('stamps null for a session without a guild, and still dispatches', () => {
    const header = mountHeader();
    publishSessionGuild(guild);
    const heard = vi.fn();
    window.addEventListener(SESSION_GUILD_EVENT, (event) => heard((event as CustomEvent).detail));
    publishSessionGuild(null);
    expect(readPublishedSessionGuild(header)).toBeNull();
    expect(heard).toHaveBeenCalledWith(null);
  });

  it('dispatches even when no header is in the page', () => {
    document.body.innerHTML = '';
    const heard = vi.fn();
    window.addEventListener(SESSION_GUILD_EVENT, heard);
    expect(() => publishSessionGuild(guild)).not.toThrow();
    expect(heard).toHaveBeenCalledTimes(1);
  });
});

describe('readPublishedSessionGuild', () => {
  it('is null before anything was published', () => {
    expect(readPublishedSessionGuild(mountHeader())).toBeNull();
  });

  it('is null for a stamp that is not a guild', () => {
    const header = mountHeader();
    header.dataset.sessionGuild = '{"name": 1}';
    expect(readPublishedSessionGuild(header)).toBeNull();
    header.dataset.sessionGuild = 'not json';
    expect(readPublishedSessionGuild(header)).toBeNull();
  });
});
