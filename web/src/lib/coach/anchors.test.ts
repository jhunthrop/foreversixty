import { describe, expect, it } from 'vitest';
import { defaultState, reportSearch, withState, type ReportState } from '../report/url';
import { coachAnchorHref } from './anchors';
import type { CoachAnchor } from './types';

const search = (patch: Partial<ReportState>): string => reportSearch(withState(defaultState(1), patch), 1);

describe('coachAnchorHref', () => {
  it('cast opens the Casts tab on the ability', () =>
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'cast', spell_id: 23881 })).toBe(
      search({ fight: 3, tab: 'casts', source: 'Player-1', ability: 23881 }),
    ));
  it('healing opens the Healing tab on the ability', () =>
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'healing', spell_id: 5 })).toBe(
      search({ fight: 3, tab: 'healing', source: 'Player-1', ability: 5 }),
    ));
  it('aura opens Buffs, or Debuffs when flagged, at the aura row', () => {
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'aura', spell_id: 9 })).toBe(
      '?fight=3&tab=buffs#aura-Player-1-9',
    );
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'aura', spell_id: 9, debuff: true })).toBe(
      '?fight=3&tab=debuffs#aura-Player-1-9',
    );
  });
  it('buffs opens the player Buffs tab', () =>
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'buffs' })).toBe(
      search({ fight: 3, tab: 'buffs', source: 'Player-1' }),
    ));
  it('death keeps the DeathsTab convention', () =>
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'death', death: 'death-Player-1-200000' })).toBe(
      '?fight=3&tab=deaths#death-Player-1-200000',
    ));
  it('resource opens the timelines window', () =>
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'resource', at_ms: 161000, to_ms: 170000 })).toBe(
      search({ fight: 3, view: 'timelines', source: 'Player-1', start: 161000, end: 170000 }),
    ));
  it('a single time gets a five second window', () =>
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'time', at_ms: 1000 })).toBe(
      search({ fight: 3, view: 'timelines', source: 'Player-1', start: 1000, end: 6000 }),
    ));
  it('the first fight leaves fight= out, as the encoder does', () =>
    expect(coachAnchorHref(1, 1, 'Player-1', { kind: 'cast', spell_id: 2 })).not.toContain('fight='));
  it('unknown kind, a missing spell, time or death, or no guid is null', () => {
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'cast' })).toBeNull();
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'nope' } as unknown as CoachAnchor)).toBeNull();
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'time' })).toBeNull();
    expect(coachAnchorHref(3, 1, 'Player-1', { kind: 'death' })).toBeNull();
    expect(coachAnchorHref(3, 1, '', { kind: 'buffs' })).toBeNull();
  });
});
