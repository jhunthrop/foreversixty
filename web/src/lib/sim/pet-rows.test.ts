import { describe, expect, it } from 'vitest';
import demonologyJson from '../../fixtures/sim/demonology-result.json';
import furyJson from '../../fixtures/sim/result.json';
import type { Summary } from '../report/types';
import { withoutAbsentPets } from './pet-rows';

// A real Demonology run through the engine and sim/adapter: the warlock registers all four
// demons as pets, the Imp is summoned and sacrificed before the pull, and only the
// Succubus is out for the fight.
const demonology = demonologyJson.summary as unknown as Summary;

function names(summary: Summary): string[] {
  return summary.damage_done.map((actor) => actor.name);
}

describe('withoutAbsentPets', () => {
  it('keeps the player and the Succubus and drops the three demons never out in the fight', () => {
    expect(names(demonology)).toEqual(['Demo', 'Felhunter', 'Imp', 'Succubus', 'Voidwalker']);
    expect(names(withoutAbsentPets(demonology))).toEqual(['Demo', 'Succubus']);
  });

  it('keeps each kept row as the engine measured it, Active included', () => {
    const succubus = withoutAbsentPets(demonology).damage_done.find((actor) => actor.name === 'Succubus');
    expect(succubus?.active_ms).toBe(demonology.duration_ms);
    expect(succubus?.total).toBeGreaterThan(0);
  });

  it('keeps a zero-damage pet that really was out, with its own uptime', () => {
    const idle = demonology.damage_done.map((actor) =>
      actor.name === 'Voidwalker' ? { ...actor, active_ms: 40_000 } : actor,
    );
    const kept = withoutAbsentPets({ ...demonology, damage_done: idle }).damage_done;
    expect(kept.map((actor) => actor.name)).toEqual(['Demo', 'Succubus', 'Voidwalker']);
    expect(kept.find((actor) => actor.name === 'Voidwalker')?.active_ms).toBe(40_000);
  });

  it('never drops the player, even with no damage and no active time', () => {
    const idle = demonology.damage_done.map((actor) =>
      actor.guid === 'sim-player' ? { ...actor, total: 0, effective: 0, active_ms: 0 } : actor,
    );
    expect(names(withoutAbsentPets({ ...demonology, damage_done: idle }))).toContain('Demo');
  });

  it('drops an absent pet from every other unit list too: casts, auras and resources', () => {
    const phantom = demonology.damage_done.find((actor) => actor.name === 'Imp');
    expect(phantom).toBeDefined();
    const guid = phantom?.guid ?? '';
    const seeded: Summary = {
      ...demonology,
      casts: [...demonology.casts, { ...demonology.casts[0], guid, name: 'Imp' }],
      auras: [...demonology.auras, { ...demonology.auras[0], target_guid: guid, target_name: 'Imp' }],
      resources: [...demonology.resources, { ...demonology.resources[0], guid, name: 'Imp' }],
    };
    const kept = withoutAbsentPets(seeded);
    expect(kept.casts.some((row) => row.guid === guid)).toBe(false);
    expect(kept.auras.some((track) => track.target_guid === guid)).toBe(false);
    expect(kept.resources.some((track) => track.guid === guid)).toBe(false);
    expect(kept.casts).toHaveLength(demonology.casts.length);
  });

  describe('a result saved before the engine measured pet uptime', () => {
    // What sim/adapter wrote before pet_active_seconds_avg: no pet_uptime mark, and every
    // row, every demon included, stamped active for the whole fight.
    const legacy: Summary = {
      ...demonology,
      damage_done: demonology.damage_done.map(({ pet_uptime: _mark, ...actor }) => ({
        ...actor,
        active_ms: demonology.duration_ms,
      })),
    };

    it('drops a pet with no damage, whatever Active it was stamped with', () => {
      expect(names(withoutAbsentPets(legacy))).toEqual(['Demo', 'Succubus']);
    });

    it('marks a pet that did damage as having no measured Active', () => {
      const rows = withoutAbsentPets(legacy).damage_done;
      expect(rows.find((actor) => actor.name === 'Succubus')?.active_unknown).toBe(true);
      expect(rows.find((actor) => actor.name === 'Demo')?.active_unknown).toBeUndefined();
    });
  });

  it('never marks a measured pet row as unknown', () => {
    const rows = withoutAbsentPets(demonology).damage_done;
    expect(rows.some((actor) => actor.active_unknown)).toBe(false);
  });

  it('leaves a result with no absent pet as it found it', () => {
    const fury = furyJson.summary as unknown as Summary;
    expect(withoutAbsentPets(fury)).toBe(fury);
  });
});
