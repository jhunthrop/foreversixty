// web/src/lib/bis/spec-header.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { characterIsOfSpec, specPageAction, specPageAnswerSpec } from './spec-header';

const PAGE = { specKey: 'warrior-fury', classSlug: 'warrior', specName: 'Fury', build: 'b1' };

const character = (overrides: Partial<MeCharacter> = {}): MeCharacter => ({
  key: 'us/normal/tester',
  region: 'us',
  ruleset: 'normal',
  name: 'Tester',
  class: 'Warrior',
  spec: 'Fury',
  level: 47,
  ...overrides,
});

describe('characterIsOfSpec', () => {
  it('matches the page’s class and spec, whatever the class name’s case', () => {
    expect(characterIsOfSpec(character(), PAGE)).toBe(true);
    expect(characterIsOfSpec(character({ class: 'warrior' }), PAGE)).toBe(true);
  });

  it('does not match another spec, another class, or a character with neither', () => {
    expect(characterIsOfSpec(character({ spec: 'Arms' }), PAGE)).toBe(false);
    expect(characterIsOfSpec(character({ class: 'Mage', spec: 'Fury' }), PAGE)).toBe(false);
    expect(characterIsOfSpec(character({ class: undefined, spec: undefined }), PAGE)).toBe(false);
  });
});

describe('the spec page’s answer and action', () => {
  it('starts at level 20 and reads gear, like the landing', () => {
    const spec = specPageAnswerSpec(PAGE);
    expect(spec.levelFloor).toBe(20);
    expect(spec.needsGear).toBe(true);
  });

  it('leaves the slot out for a character of another spec', async () => {
    const spec = specPageAnswerSpec(PAGE);
    expect(await spec.load(character({ spec: 'Arms' }))).toEqual({ kind: 'none' });
  });

  it('sends the list to the addon', () => {
    expect(specPageAction('/addon')).toEqual({ label: 'Send this list to the addon', href: '/addon' });
  });
});
