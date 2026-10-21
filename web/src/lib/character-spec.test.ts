// web/src/lib/character-spec.test.ts
// W.12 ruling 2: one spec rule for every surface -- the talents when the record carries ranks,
// else the spec the addon reported, always as a display name.
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from './account/api';
import { characterDescriptor } from './account/character-descriptor';
import { pointerForCharacter } from './account/main-character';
import { descriptorLines } from './character-header/descriptor';
import {
  characterIsSimulated,
  characterSpecClass,
  characterSpecName,
  specFromPoints,
} from './character-spec';
import { resolvedProfileSpec } from './character-profile';

const BOW = {
  key: 'us/pvp/bow-jackzon',
  region: 'us',
  ruleset: 'pvp',
  name: 'Bow Jackzon',
  class: 'hunter',
  race: 'High Order Skyborne',
  level: 60,
  spec: 'BeastMastery',
  build: { source: 'addon', captured_at: 'x', talents: { trees: ['', '', ''], points: [31, 16, 4] } },
} as MeCharacter;

const PAGE = { character: { class: 'hunter' }, history: [], best: [] };

describe('specFromPoints', () => {
  it('names the tree with the most points, ties to the first', () => {
    expect(specFromPoints('hunter', [31, 16, 4])).toBe('Beast Mastery');
    expect(specFromPoints('hunter', [0, 5, 46])).toBe('Survival');
    expect(specFromPoints('priest', [0, 46, 5])).toBe('Holy');
    expect(specFromPoints('priest', [20, 20, 0])).toBe('Discipline');
  });
  it('names nothing when nothing is spent', () => {
    expect(specFromPoints('hunter', [0, 0, 0])).toBeNull();
  });
});

describe('characterSpecName', () => {
  it('reads the talents first, whatever spec the addon reported', () => {
    expect(characterSpecName({ ...BOW, spec: 'Survival' })).toBe('Beast Mastery');
  });
  it('falls back to the reported spec as a display name, never "BeastMastery"', () => {
    expect(characterSpecName({ ...BOW, build: undefined })).toBe('Beast Mastery');
    expect(characterSpecName({ class: 'priest', spec: 'Holy' })).toBe('Holy');
  });
  it('is undefined with neither', () => {
    expect(characterSpecName({ class: 'hunter' })).toBeUndefined();
  });
  it('says the spec then the class', () => {
    expect(characterSpecClass(BOW)).toBe('Beast Mastery Hunter');
    expect(characterSpecClass({ class: 'hunter' })).toBe('Hunter');
  });
});

describe('every surface speaks the one spec', () => {
  it('the header card, the account line, the pointer and the character page agree', () => {
    const record = { ...BOW, spec: 'Survival' };
    expect(descriptorLines(record).identity).toBe('High Order Skyborne Beast Mastery Hunter');
    expect(characterDescriptor(record)).toContain('High Order Skyborne Beast Mastery Hunter · Level 60');
    expect(pointerForCharacter(record).label).toBe('Bow Jackzon · Beast Mastery Hunter');
    expect(resolvedProfileSpec(PAGE as never, record)).toBe('Beast Mastery');
  });

  it('the character page reads the public split when the viewer does not own the character', () => {
    const page = {
      ...PAGE,
      builds_seen: [{ talent_split: '0/5/46', spec: 'BeastMastery', first_seen: 'x' }],
    };
    expect(resolvedProfileSpec(page as never)).toBe('Survival');
  });

  it('the character page never prints the run-together spelling', () => {
    const page = { ...PAGE, best: [{ spec: 'BeastMastery' }], history: [] };
    expect(resolvedProfileSpec(page as never)).toBe('Beast Mastery');
  });
});

describe('characterIsSimulated', () => {
  it('is false for a healing spec and true for a damage spec or an unknown one', () => {
    expect(characterIsSimulated({ class: 'priest', spec: 'Holy' })).toBe(false);
    expect(characterIsSimulated(BOW)).toBe(true);
    expect(characterIsSimulated({ class: 'hunter' })).toBe(true);
  });
});
