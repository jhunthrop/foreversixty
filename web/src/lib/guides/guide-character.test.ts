// web/src/lib/guides/guide-character.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import type { GuideCharacter } from './guide-answer';
import {
  bandForLevel,
  MAX_LEVEL,
  onGuideCharacter,
  publishGuideCharacter,
  resetGuideCharacter,
} from './guide-character';

const BOW: GuideCharacter = {
  name: 'Bow',
  classSlug: 'hunter',
  specSlug: 'survival',
  className: 'Hunter',
  specName: 'Survival',
  level: 34,
  faction: 'alliance',
  levelLine: 'Level 34',
  identity: 'Dwarf Survival Hunter',
};

describe('bandForLevel', () => {
  it('reads the BiS page’s bands: tens from 20, 60 at the cap', () => {
    expect([1, 10, 19, 20, 29, 34, 49, 50, 59].map(bandForLevel)).toEqual([
      20, 20, 20, 20, 20, 30, 40, 50, 50,
    ]);
    expect(bandForLevel(MAX_LEVEL)).toBe(60);
  });

  it('is 60 for an unknown level', () => {
    expect(bandForLevel(undefined)).toBe(60);
  });
});

describe('the guide character channel', () => {
  afterEach(resetGuideCharacter);

  it('calls nobody at subscribe time before anything is published', () => {
    const listener = vi.fn();
    onGuideCharacter(listener);
    expect(listener).not.toHaveBeenCalled();
  });

  it('replays the last character to a late subscriber, then every new one', () => {
    publishGuideCharacter(null);
    publishGuideCharacter(BOW);
    const listener = vi.fn();
    onGuideCharacter(listener);
    expect(listener.mock.calls).toEqual([[BOW]]);
    publishGuideCharacter(null);
    expect(listener.mock.calls).toEqual([[BOW], [null]]);
  });

  it('replays a published null too (signed out is an answer)', () => {
    publishGuideCharacter(null);
    const listener = vi.fn();
    onGuideCharacter(listener);
    expect(listener.mock.calls).toEqual([[null]]);
  });

  it('stops calling an unsubscribed listener', () => {
    const listener = vi.fn();
    const stop = onGuideCharacter(listener);
    stop();
    publishGuideCharacter(BOW);
    expect(listener).not.toHaveBeenCalled();
  });
});
