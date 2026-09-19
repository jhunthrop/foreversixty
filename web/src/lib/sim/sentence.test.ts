import { describe, expect, it } from 'vitest';
import type { ActionNames } from './action-names';
import type { Summary } from '../report/types';
import { fixtureResult } from '../../test-support/sim-api';
import { simCopy } from './copy';
import { namedSummary, playerActor, summarySentence } from './sentence';

const summary = fixtureResult.summary;

// The engine names nothing (Task 23), so a sentence about a real result is a sentence over
// resolved names. These are the two the golden's top abilities actually are.
const names: ActionNames = { spell: { '25286': 'Heroic Strike' }, item: {} };

function withoutAuras(base: Summary): Summary {
  return { ...base, auras: [] };
}

/** A two-ability actor with round numbers, so the share is exact and the test is readable. */
function twoAbilities(): Summary {
  const actor = summary.damage_done[0];
  return {
    ...summary,
    damage_done: [
      {
        ...actor,
        total: 1000,
        // 400 + 210 = 610 of 1000, so the top two are 61% -- the design's own example
        // figure, and arithmetic a reader can check without running anything. The third
        // ability is deliberately the smallest so it never enters the sentence.
        abilities: [
          { ...actor.abilities[0], spell_id: 25286, name: 'spell:25286', total: 400 },
          { ...actor.abilities[0], spell_id: 20022000007, name: 'other:attack', total: 210 },
          { ...actor.abilities[0], spell_id: 99999, name: 'spell:99999', total: 190 },
        ],
      },
    ],
    auras: [
      {
        ...summary.auras[0],
        target_guid: actor.guid,
        name: 'spell:12974',
        uptime_ms: summary.duration_ms * 0.78,
      },
    ],
  };
}

describe('summarySentence', () => {
  it('names the two biggest damage sources, their share, and the biggest buff', () => {
    const withFlurry: ActionNames = { spell: { ...names.spell, '12974': 'Flurry' }, item: {} };
    expect(summarySentence(twoAbilities(), withFlurry)).toBe(
      'Heroic Strike and white hits are 61% of your damage; Flurry is up 78% of the fight.',
    );
  });

  it('drops the buff clause when there are no buffs to name', () => {
    expect(summarySentence(withoutAuras(twoAbilities()), names)).toBe(
      'Heroic Strike and white hits are 61% of your damage.',
    );
  });

  it('reads as one ability when there is only one', () => {
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const one: Summary = {
      ...withoutAuras(base),
      damage_done: [{ ...actor, total: 400, abilities: [actor.abilities[0]] }],
    };
    expect(summarySentence(one, names)).toBe('Heroic Strike is 100% of your damage.');
  });

  it('falls back to the action key rather than a blank when the build has no name', () => {
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const one: Summary = {
      ...withoutAuras(base),
      damage_done: [{ ...actor, total: 400, abilities: [actor.abilities[0]] }],
    };
    expect(summarySentence(one, null)).toBe('spell:25286 is 100% of your damage.');
  });

  it('says so rather than dividing by zero when nothing landed', () => {
    expect(summarySentence({ ...summary, damage_done: [] }, names)).toBe(simCopy.noDamage);
    expect(
      summarySentence({ ...summary, damage_done: [{ ...summary.damage_done[0], total: 0 }] }, names),
    ).toBe(simCopy.noDamage);
  });
});

describe('namedSummary', () => {
  it('rewrites ability, aura and cast names and touches nothing else', () => {
    const named = namedSummary(summary, names);
    expect(named.damage_done[0].abilities[0].name).toBe('Heroic Strike');
    expect(named.casts.every((row) => !row.spell_name.startsWith('other:'))).toBe(true);
    // The row identity is untouched: the report keys its {#each} blocks on it.
    expect(named.casts.map((row) => row.spell_id)).toEqual(summary.casts.map((row) => row.spell_id));
    expect(named.duration_ms).toBe(summary.duration_ms);
    expect(named.damage_done[0].total).toBe(summary.damage_done[0].total);
  });

  it('leaves the original untouched, because the stored result keeps the engine keys', () => {
    const before = summary.damage_done[0].abilities[0].name;
    namedSummary(summary, names);
    expect(summary.damage_done[0].abilities[0].name).toBe(before);
  });
});

describe('playerActor', () => {
  it('is the actor with the most damage, since the sim has one player and its pets', () => {
    const pet = { ...summary.damage_done[0], guid: 'sim-pet', name: 'Pet', total: 10 };
    expect(playerActor({ ...summary, damage_done: [pet, summary.damage_done[0]] })?.guid).toBe('sim-player');
  });

  it('is null for an empty table', () => {
    expect(playerActor({ ...summary, damage_done: [] })).toBeNull();
  });
});
