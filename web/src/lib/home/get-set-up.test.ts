import { describe, expect, it } from 'vitest';
import { addonLinked, getSetUpLine } from './get-set-up';
import type { Me, MeCharacter } from '../account/api';

const USER: Me['user'] = { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false };

function meWith(characters: MeCharacter[]): Me {
  return { user: USER, characters, guilds: [] };
}

const WITHOUT_BUILD: MeCharacter = { key: 'us/normal/kiloz', region: 'us', ruleset: 'normal', name: 'Kiloz' };
const WITH_ADDON_BUILD: MeCharacter = {
  ...WITHOUT_BUILD,
  build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
};
const WITH_BLIZZARD_BUILD: MeCharacter = {
  ...WITHOUT_BUILD,
  key: 'us/normal/dottzz',
  build: { source: 'blizzard', captured_at: '2026-09-20T00:00:00Z' },
};

describe('addonLinked', () => {
  it('is false with no characters', () => {
    expect(addonLinked(meWith([]))).toBe(false);
  });

  it('is false when every character has no build, or a Blizzard-sourced one', () => {
    expect(addonLinked(meWith([WITHOUT_BUILD]))).toBe(false);
    expect(addonLinked(meWith([WITH_BLIZZARD_BUILD]))).toBe(false);
  });

  it('is true once any character carries an addon-sourced build', () => {
    expect(addonLinked(meWith([WITH_BLIZZARD_BUILD, WITH_ADDON_BUILD]))).toBe(true);
  });
});

describe('getSetUpLine', () => {
  it('names the remaining addon and companion steps when the addon is not linked yet', () => {
    expect(getSetUpLine(meWith([WITHOUT_BUILD]))).toBe('Install the addon · Set up the companion →');
  });

  it('collapses to the exact one-line sentence once the addon is linked', () => {
    expect(getSetUpLine(meWith([WITH_ADDON_BUILD]))).toBe(
      'Signed in and addon linked · Set up the companion →',
    );
  });
});
