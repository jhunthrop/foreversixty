// web/src/lib/sim/warnings.test.ts
// Spec 2026-11-04-sim-results section 3.5 and section 10, check 5.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import type { GameTipModel } from '../tooltips/game-tip';
import type { ActionNames } from './action-names';
import type { SimSpellFact } from './spell-info';
import type { SimResult } from './types';
import { capstoneSpellIds, cooldownWords, MAX_WARNINGS, simWarnings } from './warnings';

const result = full as unknown as SimResult;
const names: ActionNames = {
  spell: { '1719': 'Recklessness', '23894': 'Bloodthirst', '1329555': 'Reckless Swings' },
  item: {},
};
const tip: GameTipModel = { name: '', lines: [], description: '' };
const COOLDOWNS: Readonly<Record<number, number>> = { 1719: 1_800_000, 12328: 180_000, 18499: 30_000 };
const fact = (id: number): SimSpellFact | undefined =>
  COOLDOWNS[id] === undefined ? undefined : { name: '', rank: 0, cooldown_ms: COOLDOWNS[id], tip };
const ids = (from: SimResult) => simWarnings({ result: from, names, fact }).map((note) => note.id);

const withoutBloodthirst: SimResult = {
  ...result,
  summary: {
    ...result.summary,
    casts: result.summary.casts.filter((row) => row.spell_name !== 'spell:23894'),
    damage_done: result.summary.damage_done.map((actor) => ({
      ...actor,
      abilities: actor.abilities.filter((ability) => ability.name !== 'spell:23894'),
    })),
  },
};

describe('simWarnings', () => {
  it('on the fixture fires W4 for Recklessness and nothing else', () => {
    const notes = simWarnings({ result, names, fact });
    expect(notes).toHaveLength(1);
    expect(notes[0]).toMatchObject({ id: 'long-cooldown', actionKey: 'spell:1719' });
    expect(notes[0]?.text).toBe(
      'Recklessness has a 30 minute cooldown, longer than this fight, so the result counts one use every pull. In a raid it is back only every 30 minutes.',
    );
  });

  it('draws no cooldown note until the spell facts are here', () => {
    expect(simWarnings({ result, names, fact: () => undefined })).toEqual([]);
  });

  it('W2: fires with the Bloodthirst rows removed, whether or not the talents decode', () => {
    const notes = simWarnings({ result: withoutBloodthirst, names, fact });
    expect(notes[0]).toMatchObject({ id: 'capstone', action: 'sim', actionKey: 'spell:23894' });
    expect(notes[0]?.text).toBe('Bloodthirst was never pressed in this run. Check the talents in Open in the sim.');
  });

  it('W2: stays quiet below the level a capstone can be held', () => {
    const low: SimResult = {
      ...withoutBloodthirst,
      request: {
        ...withoutBloodthirst.request,
        character: { ...withoutBloodthirst.request.character, level: 30 },
      },
    };
    expect(ids(low)).not.toContain('capstone');
  });

  it('W1: an older engine comes first and links to the sim', () => {
    const notes = simWarnings({ result: { ...result, engine_version: 'abc1234' }, names, fact });
    expect(notes[0]?.id).toBe('stale');
    expect(notes[0]?.action).toBe('sim');
    expect(notes[0]?.text).toBe('This run used an older engine. A run today may give a different number.');
  });

  it('W3: only when W2 did not fire and the talents are known to fit', () => {
    const gaps = [{ actionKey: 'spell:1329555' }];
    const note = simWarnings({ result, names, fact, talentGaps: gaps })[0];
    expect(note).toMatchObject({ id: 'missing-talent', action: 'planner' });
    expect(note?.text).toBe('Reckless Swings is in the Fury rotation, but these talents do not take it.');
    const without = simWarnings({ result: withoutBloodthirst, names, fact, talentGaps: gaps });
    expect(without.map((entry) => entry.id)).not.toContain('missing-talent');
    expect(simWarnings({ result, names, fact, talentGaps: null }).map((entry) => entry.id)).toEqual([
      'long-cooldown',
    ]);
  });

  it('W5: fires at 5% lost at the cap, not at 2.7%, with the rage advice', () => {
    const [track] = result.summary.resources;
    if (track === undefined) throw new Error('no resource');
    const capped: SimResult = {
      ...result,
      summary: { ...result.summary, resources: [{ ...track, wasted: 120 }] },
    };
    const note = simWarnings({ result: capped, names, fact }).find((entry) => entry.id === 'at-cap');
    expect(note?.text).toBe(
      '120 rage (5.6%) was lost at the cap. Queue Heroic Strike sooner so swings do not land at full rage.',
    );
    expect(ids(result)).not.toContain('at-cap');
  });

  it('W6: fires above 0.5% of the mean', () => {
    const wide: SimResult = { ...result, dps: { ...result.dps, error: 9 } };
    const note = simWarnings({ result: wide, names, fact }).find((entry) => entry.id === 'wide-error');
    expect(note?.text).toBe(
      "This run's error is ± 9.0 DPS. A difference smaller than that between two runs is noise.",
    );
  });

  it('never shows more than three, highest priority first', () => {
    const [track] = result.summary.resources;
    if (track === undefined) throw new Error('no resource');
    const all: SimResult = {
      ...withoutBloodthirst,
      engine_version: 'abc1234',
      dps: { ...result.dps, error: 9 },
      summary: { ...withoutBloodthirst.summary, resources: [{ ...track, wasted: 120 }] },
    };
    expect(ids(all)).toEqual(['stale', 'capstone', 'long-cooldown']);
    expect(ids(all)).toHaveLength(MAX_WARNINGS);
  });
});

describe('capstoneSpellIds', () => {
  it('pins Fury and Arms to every rank, and names no capstone for a spec it does not know', () => {
    expect(capstoneSpellIds('warrior-fury')).toEqual([23881, 23892, 23893, 23894]);
    expect(capstoneSpellIds('warrior-arms')).toEqual([12294, 21551, 21552, 21553]);
    expect(capstoneSpellIds('mage-arcane')).toEqual([]);
  });
});

describe('cooldownWords', () => {
  it('reads an adjective and a noun', () => {
    expect(cooldownWords(1_800_000)).toEqual({ length: '30 minute', every: '30 minutes' });
    expect(cooldownWords(7_200_000)).toEqual({ length: '2 hour', every: '2 hours' });
  });
});
