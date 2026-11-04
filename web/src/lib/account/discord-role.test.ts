// @vitest-environment jsdom
import { afterEach, describe, expect, it } from 'vitest';
import { discordRoleCopy } from './discord-role-copy';
import { discordRoleReturn, discordRoleStartUrl, takeDiscordRoleFlag } from './discord-role';

afterEach(() => window.history.replaceState({}, '', '/'));

describe('the Discord Premium role row', () => {
  it('starts at the API verification route', () => {
    expect(discordRoleStartUrl('https://api.foreversixty.gg')).toBe(
      'https://api.foreversixty.gg/v1/discord/linked-role',
    );
  });

  it('maps each return flag', () => {
    expect(discordRoleReturn('linked')).toEqual({
      message: discordRoleCopy.returnLinked,
      role: 'status',
      navigate: false,
    });
    expect(discordRoleReturn('failed').role).toBe('alert');
    expect(discordRoleReturn('start').navigate).toBe(true);
  });

  it('reads the flag once and keeps the rest of the address', () => {
    window.history.replaceState({}, '', '/account?x=1&discord_role=denied#you');
    expect(takeDiscordRoleFlag()).toBe('denied');
    expect(`${window.location.pathname}${window.location.search}${window.location.hash}`).toBe(
      '/account?x=1#you',
    );
    expect(takeDiscordRoleFlag()).toBeNull();
  });

  it('reads an unknown value as no flag', () => {
    window.history.replaceState({}, '', '/account?discord_role=https://evil.example');
    expect(takeDiscordRoleFlag()).toBeNull();
    expect(window.location.search).toBe('');
  });
});
