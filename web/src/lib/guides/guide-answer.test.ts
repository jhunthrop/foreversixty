// web/src/lib/guides/guide-answer.test.ts
// The signed-in answer's model (guides spec 2026-11-08 §4.4): its states, the header selector's
// character (bug B1: never silently the account's first), and the links a character reads on its
// own spec's guide below 60, at 60, and anywhere else.
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { answerCopy } from '../account/answer-copy';
import { pointerForCharacter } from '../account/main-character';
import { addonCopy } from '../addon/copy';
import { guidesCopy } from './copy';
import { guideAnswerFrom, guideCharacterOf, type GuideAnswer, type GuideAnswerPage } from './guide-answer';

function character(name: string, overrides: Partial<MeCharacter> = {}): MeCharacter {
  return { key: `us/normal/${name.toLowerCase()}`, region: 'us', ruleset: 'normal', name, ...overrides };
}

const armsy = character('Armsy', { class: 'Warrior', spec: 'Arms', level: 60, faction: 'alliance' });
const bow = character('Bow', {
  class: 'Hunter',
  spec: 'Survival',
  level: 34,
  race: 'Dwarf',
  faction: 'alliance',
});

const SURVIVAL: GuideAnswerPage = {
  classSlug: 'hunter',
  specSlug: 'survival',
  ladder: [
    { level: 33, name: 'Lacerate' },
    { level: 34, name: 'Lacerate' },
    { level: 35, name: null },
    { level: 36, name: 'Expose Prey' },
  ],
};

function ready(answer: GuideAnswer): Extract<GuideAnswer, { state: 'ready' }> {
  if (answer.state !== 'ready') throw new Error(`expected a ready answer, got ${answer.state}`);
  return answer;
}

describe('guideAnswerFrom states', () => {
  it('is hidden signed out', () => {
    expect(guideAnswerFrom(null, null, {})).toEqual({ state: 'hidden' });
  });

  it('says so, with the setup link, for an account with no character', () => {
    expect(guideAnswerFrom({ characters: [] }, null, {})).toEqual({
      state: 'none',
      message: guidesCopy.noCharacter,
      link: { label: addonCopy.setupPageTitle, href: '/setup' },
    });
  });

  it('names the character whose spec is not known yet, with the selector link', () => {
    const unknown = character('Nospec', { class: 'Mage' });
    expect(guideAnswerFrom({ characters: [unknown] }, null, {})).toEqual({
      state: 'unknown',
      message: answerCopy.unknown('Nospec'),
      linkLabel: answerCopy.unknownLink,
    });
  });

  it('reports a failed read with a retry, whatever else it was handed', () => {
    expect(guideAnswerFrom({ characters: [armsy] }, null, {}, true)).toEqual({
      state: 'error',
      message: answerCopy.error,
      retry: answerCopy.retry,
    });
  });

  it('is hidden when the pointer names no account character (a pasted export is its own identity)', () => {
    const pasted = { ...pointerForCharacter(character('Stranger', { class: 'Mage', spec: 'Frost' })) };
    expect(guideAnswerFrom({ characters: [armsy, bow] }, pasted, {})).toEqual({ state: 'hidden' });
  });
});

describe('the character the answer names (bug B1)', () => {
  it('is the header pointer’s character, not the account’s first', () => {
    const answer = ready(guideAnswerFrom({ characters: [armsy, bow] }, pointerForCharacter(bow), {}));
    expect(answer.character.name).toBe('Bow');
  });

  it('is the account main with no pointer at all', () => {
    const answer = ready(
      guideAnswerFrom({ characters: [armsy, bow], main_character_key: bow.key }, null, {}),
    );
    expect(answer.character.name).toBe('Bow');
  });

  it('carries the slugs, the level line and the identity line', () => {
    expect(guideCharacterOf(bow)).toEqual({
      name: 'Bow',
      classSlug: 'hunter',
      specSlug: 'survival',
      className: 'Hunter',
      specName: 'Survival',
      level: 34,
      faction: 'alliance',
      levelLine: 'Level 34',
      identity: 'Dwarf Survival Hunter',
    });
  });

  it('slugs a two-word spec for its guide link', () => {
    const bm = character('Pets', { class: 'Hunter', spec: 'Beast Mastery', level: 60 });
    expect(guideCharacterOf(bm)?.specSlug).toBe('beast-mastery');
  });

  it('is undefined until the class and spec are known', () => {
    expect(guideCharacterOf(character('X', { class: 'Mage' }))).toBeUndefined();
    expect(guideCharacterOf(character('X', { spec: 'Frost' }))).toBeUndefined();
  });
});

describe('the answer’s links', () => {
  const me = { characters: [armsy, bow] };

  it('on its own spec below 60: talents at its level, the ladder’s next point, BiS at its band', () => {
    const answer = ready(guideAnswerFrom(me, pointerForCharacter(bow), SURVIVAL));
    expect(answer.links).toEqual([
      { label: 'Talents at level 34 ›', href: '#leveling' },
      { label: `${answerCopy.bis} ›`, href: '/bis/hunter/survival?band=30' },
    ]);
    // The first ladder step above its level that names a talent (35 spends nothing).
    expect(answer.next).toBe('Next point at 36: Expose Prey');
  });

  it('on its own spec below 60 with no point left on the ladder: no next line', () => {
    const answer = ready(guideAnswerFrom(me, pointerForCharacter(bow), { ...SURVIVAL, ladder: [] }));
    expect(answer.next).toBeUndefined();
    expect(answer.links).toHaveLength(2);
  });

  it('on its own spec at 60: its best in slot and the sim', () => {
    const answer = ready(
      guideAnswerFrom(me, pointerForCharacter(armsy), { classSlug: 'warrior', specSlug: 'arms' }),
    );
    expect(answer.links).toEqual([
      { label: `${answerCopy.bis} ›`, href: '/bis/warrior/arms' },
      { label: `${answerCopy.sim} ›`, href: '/sim' },
    ]);
    expect(answer.next).toBeUndefined();
  });

  it('anywhere else (another spec, another class, /guides): its own guide', () => {
    for (const page of [{ classSlug: 'hunter', specSlug: 'marksmanship' }, { classSlug: 'mage' }, {}]) {
      const answer = ready(guideAnswerFrom(me, pointerForCharacter(bow), page));
      expect(answer.links).toEqual([
        { label: 'Your guide: Survival Hunter ›', href: '/guides/hunter/survival' },
      ]);
      expect(answer.next).toBeUndefined();
    }
  });
});
