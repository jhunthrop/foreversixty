import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { HeaderAnswer } from '../character-header/types';
import {
  answerBodyOf,
  answerCharacterOf,
  emptyAnswer,
  landingAnswerCharacter,
  readyAnswer,
} from './landing-answer-view';

const FELSNAP: MeCharacter = {
  key: 'us/normal/felsnap',
  region: 'us',
  ruleset: 'normal',
  name: 'Felsnap',
  class: 'Hunter',
  spec: 'Beast Mastery',
  race: 'Dwarf',
  level: 24,
  faction: 'alliance',
} as MeCharacter;

const COUNT: HeaderAnswer = {
  kind: 'figures',
  label: 'Best-in-slot pieces you wear',
  staleLabel: 'Pieces you wore at the last sync',
  figures: [{ value: '2', unit: 'of 14 slots' }],
  bar: 2 / 14,
  note: { text: 'Trinkets and Off hand have no pick at this level range.' },
};

describe('the /bis landing answer', () => {
  it('names the character with its level and identity on one line', () => {
    expect(answerCharacterOf(FELSNAP)).toEqual({
      name: 'Felsnap',
      classSlug: 'hunter',
      className: 'Hunter',
      line: 'Level 24 · Dwarf Beast Mastery Hunter',
    });
  });

  it('a count is the Figure, its caption, the Meter and the note', () => {
    expect(answerBodyOf('load', COUNT, false)).toEqual({
      kind: 'count',
      worn: 2,
      unit: 'of 14 slots',
      caption: 'Best-in-slot pieces you wear',
      bar: 2 / 14,
      note: 'Trinkets and Off hand have no pick at this level range.',
    });
  });

  it('a stale addon sync takes the stale caption', () => {
    expect(answerBodyOf('load', COUNT, true)).toMatchObject({ caption: 'Pieces you wore at the last sync' });
  });

  it('under 20 and no gear are sentences; a failed count offers a retry', () => {
    expect(answerBodyOf('below-floor', undefined, false)).toEqual({
      kind: 'sentence',
      text: 'Best in slot starts at 20.',
    });
    expect(answerBodyOf('no-gear', undefined, false)).toMatchObject({
      kind: 'sentence',
      link: { href: '/setup' },
    });
    expect(answerBodyOf('load', 'failed', false)).toEqual({
      kind: 'failed',
      text: 'Your gear count did not load.',
      retry: 'Try again',
    });
  });

  it('the ready answer carries the one sentence-case button', () => {
    const answer = readyAnswer({
      character: FELSNAP,
      plan: 'load',
      fetched: COUNT,
      action: { label: 'Beast Mastery gear for 20 to 29', href: '/bis/hunter/beast-mastery' },
      now: new Date('2026-11-07T12:00:00Z'),
    });
    expect(answer).toMatchObject({ state: 'ready', action: { label: 'Beast Mastery gear for 20 to 29' } });
  });

  it('signed out is hidden, an empty account offers setup, a failed read offers a retry', () => {
    expect(emptyAnswer(null, false)).toEqual({ state: 'hidden' });
    expect(emptyAnswer({ characters: [] }, false)).toMatchObject({ state: 'none', link: { href: '/setup' } });
    expect(emptyAnswer(null, true)).toMatchObject({ state: 'error' });
    expect(landingAnswerCharacter(null, null)).toBeNull();
    expect(
      landingAnswerCharacter({ characters: [FELSNAP], main_character_key: FELSNAP.key }, null)?.name,
    ).toBe('Felsnap');
  });
});
