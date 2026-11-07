// web/src/lib/not-found/answer.test.ts
// The 404's signed-in answer (pages spec 2026-11-08 §4.4, ruling Q6): your guide, your best in
// slot at your level band, then the smaller Your character; every other state the guides'.
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { answerCopy } from '../account/answer-copy';
import { CHARACTER_LINK_INDEX, notFoundAnswerFrom } from './answer';

const felsnap: MeCharacter = {
  key: 'us/normal/felsnap',
  region: 'us',
  ruleset: 'normal',
  name: 'Felsnap',
  class: 'Hunter',
  spec: 'Beast Mastery',
  race: 'Dwarf',
  level: 34,
  faction: 'alliance',
};

describe('the 404 answer', () => {
  it('links your guide, your best in slot at your band, then your character', () => {
    const answer = notFoundAnswerFrom({ characters: [felsnap], main_character_key: felsnap.key }, null);
    expect(answer.state).toBe('ready');
    if (answer.state !== 'ready') return;
    expect(answer.links).toEqual([
      { label: 'Your guide: Beast Mastery Hunter ›', href: '/guides/hunter/beast-mastery' },
      { label: `${answerCopy.bis} ›`, href: '/bis/hunter/beast-mastery?band=30' },
      { label: 'Your character ›', href: '/character/us/normal/felsnap' },
    ]);
    expect(answer.links[CHARACTER_LINK_INDEX]?.label).toBe('Your character ›');
  });

  it('keeps the guides answer for every other state', () => {
    expect(notFoundAnswerFrom(null, null)).toEqual({ state: 'hidden' });
    expect(notFoundAnswerFrom(null, null, true).state).toBe('error');
    expect(notFoundAnswerFrom({ characters: [], main_character_key: null }, null).state).toBe('none');
    const unknown = { ...felsnap, class: undefined, spec: undefined };
    expect(notFoundAnswerFrom({ characters: [unknown], main_character_key: unknown.key }, null).state).toBe(
      'unknown',
    );
  });
});
