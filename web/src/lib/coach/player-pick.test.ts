import { describe, expect, it } from 'vitest';
import type { Me } from '../account/me-types';
import { pickPlayer } from './player-pick';
import { viewerFromMe } from './viewer';

const roster = [
  { guid: 'P-1', name: 'Simfury-X-', class: 'Warrior' },
  { guid: 'P-2', name: 'Bow-Jackzon-X-', class: 'Hunter' },
  { guid: 'P-3', name: 'Thrall-X-', class: 'Shaman' },
];
const meWith = (officer: boolean): Me =>
  ({
    user: { id: 1 },
    characters: [
      { key: 'us/normal/bow-jackzon', name: 'Bow Jackzon' },
      { key: 'us/normal/thrall', name: 'Thrall' },
    ],
    guilds: [],
    entitlements: { officer_views: officer },
  }) as unknown as Me;

describe('pickPlayer', () => {
  const owner = viewerFromMe(meWith(false));
  it('friendlies chooses the first owned character and lists all of them', () => {
    const pick = pickPlayer(owner, roster, 'friendlies');
    expect(pick.unit?.guid).toBe('P-2');
    expect(pick.ownKey).toBe('us/normal/bow-jackzon');
    expect(pick.owned.map((u) => u.guid)).toEqual(['P-2', 'P-3']);
  });
  it('a named owned unit wins over the default', () => {
    expect(pickPlayer(owner, roster, 'P-3').unit?.guid).toBe('P-3');
  });
  it('a stranger is reviewable by an officer only', () => {
    expect(pickPlayer(owner, roster, 'P-1').unit).toBeNull();
    const officer = pickPlayer(viewerFromMe(meWith(true)), roster, 'P-1');
    expect(officer.unit?.guid).toBe('P-1');
    expect(officer.ownKey).toBeNull();
  });
  it('enemies and unknown sources choose nobody', () => {
    expect(pickPlayer(owner, roster, 'enemies').unit).toBeNull();
    expect(pickPlayer(owner, roster, 'Gone').unit).toBeNull();
  });
});
