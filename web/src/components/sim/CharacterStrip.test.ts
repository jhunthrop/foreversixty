// web/src/components/sim/CharacterStrip.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { CharacterSource } from '../../lib/sim/types';
import { handoffCopy } from '../../lib/sim/handoff-copy';
import type { SimCharacter } from '../../lib/sim/character';
import CharacterStrip from './CharacterStrip.svelte';

const source: CharacterSource = { kind: 'addon', ref: '', captured_at: '2026-09-21T10:00:00Z' };

// Minimal filler; each fixture below overrides only the fields the assertion cares about --
// the same pattern character.test.ts's own `plannerGearFor` describe block uses.
const base: SimCharacter = {
  name: 'Thrallgar',
  spec: 'warrior-fury',
  class_slug: 'warrior',
  race_slug: 'orc',
  talent_level: 60,
  tree_version: '1.15.9.69722',
  point_order: [],
  gear: {},
  gear_slots: [],
  buffs: [],
  consumables: [],
  source,
  professions: [],
  bags: [],
  bank: [],
  sets: [],
  loadouts: [],
};

const requiredProps = { items: new Map(), onchange: () => {}, onrace: () => {} };

// 30 points is well under the 51 a full build spends -- talent_level 39 (BASE_LEVEL 9 + 30)
// is what characterFromFs1/fromBuildDraft would actually stamp for this many points, but
// this test only needs it under SIM_LEVEL, which is the component's own gate.
const underLeveled: SimCharacter = { ...base, talent_level: 39, point_order: new Array(30).fill(0) };
const fullBuild: SimCharacter = { ...base, talent_level: 60, point_order: new Array(51).fill(0) };

describe('CharacterStrip’s honest "simmed at 60" note', () => {
  it('says the character is simmed at 60 when it has fewer than 51 points', () => {
    const { body } = render(CharacterStrip, { props: { character: underLeveled, ...requiredProps } });
    expect(body).toContain(handoffCopy.simmedAtSixty);
    expect(body).toContain('30 talent points');
  });

  it('says nothing extra for a full 51-point build', () => {
    const { body } = render(CharacterStrip, { props: { character: fullBuild, ...requiredProps } });
    expect(body).not.toContain(handoffCopy.simmedAtSixty);
  });
});
