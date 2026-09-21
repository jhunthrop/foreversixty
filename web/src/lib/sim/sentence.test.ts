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

  it('names the main hand, not off-hand, for a two-handed build’s only attack row', () => {
    // dps-minmaxer review round 1, D2: a two-handed build's only auto-attack row is the
    // main hand. The engine tags it other:attack/1 (sim/core/attack.go: tagMainhand = 1);
    // this pins that the sentence reads it as the main hand and never as off-hand.
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const twoHander: Summary = {
      ...withoutAuras(base),
      damage_done: [
        {
          ...actor,
          total: 610,
          abilities: [
            { ...actor.abilities[0], spell_id: 20022000007, name: 'other:attack/1', total: 400 },
            { ...actor.abilities[0], spell_id: 25286, name: 'spell:25286', total: 210 },
          ],
        },
      ],
    };
    expect(summarySentence(twoHander, names)).toBe(
      'main-hand white hits and Heroic Strike are 100% of your damage.',
    );
  });

  it('falls back to prose rather than the raw key or the id itself when the build has no name', () => {
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const one: Summary = {
      ...withoutAuras(base),
      damage_done: [{ ...actor, total: 400, abilities: [actor.abilities[0]] }],
    };
    expect(summarySentence(one, null)).toBe('An unnamed spell is 100% of your damage.');
  });

  it('never emits a raw spell:/item:/dungeon: key in the sentence, resolved or not, and never the bare id either', () => {
    // dps-minmaxer review round 2, D48: "off-hand white hits and spell:20662 are 78% of
    // your damage" -- the exact defect this pins down, both resolved and unresolved.
    // 2026-09-21 result-page review round 2: D48's own fix, "Spell 20662", turned out to
    // still be an id reaching a player -- this also pins that the unresolved fallback
    // never contains a bare number where a name belongs.
    const resolved = summarySentence(twoAbilities(), names);
    expect(resolved).not.toMatch(/\b(spell|item|dungeon):/);

    const unresolved = summarySentence(twoAbilities(), null);
    expect(unresolved).not.toMatch(/\b(spell|item|dungeon):/);
    expect(unresolved).not.toMatch(/\bSpell \d+\b/);
    expect(unresolved).toBe(
      'An unnamed spell and white hits are 61% of your damage; An unnamed spell is up 78% of the fight.',
    );
  });

  it('never names an aura the BUFFS tab’s own sanitizer would drop (final whole-branch review, Finding 1)', () => {
    // Before Task 4, the sentence and the tabs read the same raw array; this regression
    // reintroduced the split. other:move is the engine's own movement bookkeeping (never a
    // player-facing aura) -- sanitizeAuraTracks drops it, and a sentence reading raw auras
    // can still pick it as the biggest "buff" and say "Move is up N% of the fight".
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const withMove: Summary = {
      ...base,
      duration_ms: 10_000,
      auras: [
        { ...summary.auras[0], target_guid: actor.guid, type: 'BUFF', name: 'other:move', uptime_ms: 4_100 },
      ],
    };
    expect(summarySentence(withMove, names)).toBe('Heroic Strike and white hits are 61% of your damage.');
  });

  it('names the sanitizer’s folded aura, not a raw tag-variant row, and reports the tab’s own summed uptime', () => {
    // The same fold-by-identity BUFFS/DEBUFFS reads (aura-rows.ts): a spell metered under
    // two tags is one aura, not two, and its uptime is the sum. A sentence reading raw auras
    // can pick the bigger of the two rows alone -- naming a tag suffix ("Flurry (2)") the
    // tab never shows, and understating the uptime the tab reports for the same aura.
    const withFlurry: ActionNames = { spell: { ...names.spell, '12974': 'Flurry' }, item: {} };
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const tagged: Summary = {
      ...base,
      duration_ms: 10_000,
      auras: [
        { ...summary.auras[0], target_guid: actor.guid, type: 'BUFF', name: 'spell:12974', uptime_ms: 3_000 },
        {
          ...summary.auras[0],
          target_guid: actor.guid,
          type: 'BUFF',
          name: 'spell:12974/1',
          uptime_ms: 5_000,
        },
      ],
    };
    expect(summarySentence(tagged, withFlurry)).toBe(
      'Heroic Strike and white hits are 61% of your damage; Flurry is up 80% of the fight.',
    );
  });

  it('says so rather than dividing by zero when nothing landed', () => {
    expect(summarySentence({ ...summary, damage_done: [] }, names)).toBe(simCopy.noDamage);
    expect(
      summarySentence({ ...summary, damage_done: [{ ...summary.damage_done[0], total: 0 }] }, names),
    ).toBe(simCopy.noDamage);
  });

  // 2026-09-21 result-page review, Defect 2: "Frostbolt and Cold Snap are 100% of your
  // damage; Spell 9910 is up 100% of the fight" on a raid-buffed Frost Mage -- 9910 is
  // Thorns, a raid buff nothing in the mage's own kit ever casts. It won the uptime sort
  // because a raid buff applied once at pull outlasts almost anything the player's own spec
  // does. These two tests are the fix: prefer an own-spec aura, and say nothing about
  // uptime -- never name the raid buff -- when there is not one.
  it('prefers the player’s own aura over a higher-uptime external raid buff', () => {
    const withThorns: ActionNames = {
      spell: { ...names.spell, '12974': 'Flurry', '9910': 'Thorns' },
      item: {},
      // Flurry (12974) is the player's own; Thorns (9910) is not -- the shape
      // loadActionNames produces for a raid-buffed sim (action-names.ts).
      ownSpell: new Set(['25286', '12974']),
    };
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const withRaidBuff: Summary = {
      ...base,
      duration_ms: 10_000,
      auras: [
        { ...summary.auras[0], target_guid: actor.guid, type: 'BUFF', name: 'spell:12974', uptime_ms: 3_000 },
        // Up the whole fight -- the highest uptime in this list by far, and exactly the
        // aura the old "highest uptime wins" sort would have named instead.
        { ...summary.auras[0], target_guid: actor.guid, type: 'BUFF', name: 'spell:9910', uptime_ms: 10_000 },
      ],
    };
    expect(summarySentence(withRaidBuff, withThorns)).toBe(
      'Heroic Strike and white hits are 61% of your damage; Flurry is up 30% of the fight.',
    );
  });

  it('says nothing about uptime, rather than naming a raid buff, when no own-spec aura is up', () => {
    const withThorns: ActionNames = {
      spell: { ...names.spell, '9910': 'Thorns' },
      item: {},
      ownSpell: new Set(['25286']),
    };
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const onlyRaidBuff: Summary = {
      ...base,
      duration_ms: 10_000,
      auras: [
        { ...summary.auras[0], target_guid: actor.guid, type: 'BUFF', name: 'spell:9910', uptime_ms: 10_000 },
      ],
    };
    expect(summarySentence(onlyRaidBuff, withThorns)).toBe(
      'Heroic Strike and white hits are 61% of your damage.',
    );
  });

  // 2026-09-21 result-page review round 3, E8: "Battle Shout is up 101% of the fight" --
  // sim/combine's own bug (fixed at the source), pinned here as a last line of defence:
  // an aura whose uptime_ms somehow exceeds duration_ms must still read exactly 100%.
  it('never reads an uptime over 100%, even when uptime_ms exceeds duration_ms', () => {
    const base = twoAbilities();
    const actor = base.damage_done[0];
    const impossible: Summary = {
      ...base,
      duration_ms: 178_700,
      auras: [
        {
          ...summary.auras[0],
          target_guid: actor.guid,
          type: 'BUFF',
          uptime_ms: 180_200,
          name: 'spell:12974',
        },
      ],
    };
    const withFlurry: ActionNames = { spell: { ...names.spell, '12974': 'Flurry' }, item: {} };
    expect(summarySentence(impossible, withFlurry)).toBe(
      'Heroic Strike and white hits are 61% of your damage; Flurry is up 100% of the fight.',
    );
  });
});

describe('namedSummary', () => {
  it('rewrites ability, aura and cast names and touches nothing else', () => {
    const named = namedSummary(summary, names);
    // abilities[0] is the fixture's own auto-attack row (a tagged "other:" key, resolved
    // through attackHandName rather than `names`); spell 25286 -- the one id `names`
    // carries a resolved name for -- is abilities[3] in this fixture.
    expect(named.damage_done[0].abilities[3].name).toBe('Heroic Strike');
    expect(named.casts.every((row) => !row.spell_name.startsWith('other:'))).toBe(true);
    // The row identity is untouched: the report keys its {#each} blocks on it.
    expect(named.casts.map((row) => row.spell_id)).toEqual(summary.casts.map((row) => row.spell_id));
    expect(named.duration_ms).toBe(summary.duration_ms);
    expect(named.damage_done[0].total).toBe(summary.damage_done[0].total);
  });

  it('drops engine-internal aura rows and resolves names on what survives (Task 4)', () => {
    // The fixture's own 13 raw aura rows: 5 are inert (never applied, never up -- 2457, 71,
    // 18499, 25288 and spell 20569's lone tag row) and other:move is the engine's own
    // movement bookkeeping, never a player-facing aura. 13 - 5 - 1 = 7. This fixture's own
    // rows happen not to collide under fold-by-identity (no two rows share a normalized
    // spell id); `aura-rows.test.ts`'s own "aggregates tag/rank variants" test covers that
    // merge directly, with synthetic rows built to collide on purpose, so it does not need
    // rediscovering here against whatever a fixture regeneration happens to produce.
    const named = namedSummary(summary, names);
    expect(named.auras).toHaveLength(7);
    expect(named.auras.some((track) => track.spell_id === 2457)).toBe(false);
    expect(named.auras.some((track) => track.name === 'other:move')).toBe(false);

    // spell:25286's own tag row (/1) is renamed to its base id by normalizeSpellIdentity.
    const renamed = named.auras.find((track) => track.spell_id === 25286);
    expect(renamed?.applications).toBe(3);
    expect(renamed?.uptime_ms).toBe(11_327);

    // sim/adapter/adapter.go writes the player's own display name into every aura's
    // appliers, not a GUID; every surviving row is normalized to the target's guid
    // instead, so AuraTable renders a self-buff with no "from an unnamed source" line.
    expect(named.auras.every((track) => !track.appliers.includes(track.target_name))).toBe(true);
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
