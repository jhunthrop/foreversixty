// web/src/lib/sim/aura-rows.test.ts
import { describe, expect, it } from 'vitest';
import { sanitizeAuraTracks } from './aura-rows';
import type { AuraTrack } from '../report/types';

/** A minimal, valid AuraTrack, overridden per test. */
function track(overrides: Partial<AuraTrack>): AuraTrack {
  return {
    target_guid: 'sim-player',
    target_name: 'Mage',
    spell_id: 1,
    name: 'spell:1',
    type: 'BUFF',
    applications: 1,
    max_stacks: 0,
    uptime_ms: 1000,
    segments: [],
    appliers: ['Mage'],
    ...overrides,
  };
}

describe('sanitizeAuraTracks', () => {
  it('drops engine-seeded rows that never applied and were never up, real id or not', () => {
    // "Blizzard#10" / "Blizzard#6141": two different, genuinely unrelated client spell
    // ids the engine seeds one row for per rank, whether or not the rank ever fired. They
    // are dropped by their own uptime/applications, never by the name they would later
    // resolve to (they are not even given a resolvable name here).
    const rows = [
      track({ spell_id: 10, name: 'spell:10', applications: 0, uptime_ms: 0 }),
      track({ spell_id: 6141, name: 'spell:6141', applications: 0, uptime_ms: 0 }),
    ];
    expect(sanitizeAuraTracks(rows)).toEqual([]);
  });

  it('keeps a proc that fired but rounds to 0ms of uptime -- applications alone is a real signal', () => {
    // sim/adapter/adapter.go's UptimeMS and Applications are two independently-rounded
    // averages (UptimeSecondsAvg/ProcsAvg): a proc consumed almost instantly, or one that
    // only fires in a small fraction of iterations, can have applications > 0 while its
    // average duration rounds to 0ms. isInert must require BOTH fields to be zero, not
    // just uptime_ms, or this row -- the only place its "fired N times" signal lives --
    // disappears with no recovery path (the Casts tab tracks casts, not passive procs).
    const shortProc = track({ spell_id: 42, name: 'spell:42', applications: 7, uptime_ms: 0 });
    const [surviving] = sanitizeAuraTracks([shortProc]);
    expect(surviving.applications).toBe(7);
    expect(surviving.uptime_ms).toBe(0);
  });

  it('keeps a real buff untouched, apart from its already-self-applied source', () => {
    const buff = track({
      spell_id: 9910,
      name: 'spell:9910',
      type: 'BUFF',
      uptime_ms: 179_856,
      applications: 1,
    });
    expect(sanitizeAuraTracks([buff])).toEqual([{ ...buff, appliers: ['sim-player'] }]);
  });

  it('keeps a real debuff as a debuff -- type is structural, never rewritten', () => {
    // Rend: a fixture debuff, so the classification pipeline is proven to respect a real
    // DEBUFF row rather than assuming every sim aura is a BUFF.
    const rend = track({
      target_guid: 'enemy-1',
      target_name: 'Training Dummy',
      spell_id: 772,
      name: 'spell:772',
      type: 'DEBUFF',
      uptime_ms: 45_000,
      applications: 56,
      appliers: ['Mage'],
    });
    const [surviving] = sanitizeAuraTracks([rend]);
    expect(surviving.type).toBe('DEBUFF');
    expect(surviving.applications).toBe(56);
  });

  it('drops the movement pseudo-aura by its raw key structure, not by uptime', () => {
    // Given a non-zero uptime deliberately: this proves other:move is filtered because of
    // what it IS (kind "other", label "move"), not because it happened to be at 0%.
    const move = track({ spell_id: 22_000_014, name: 'other:move', applications: 4, uptime_ms: 5000 });
    expect(sanitizeAuraTracks([move])).toEqual([]);
  });

  it('aggregates tag/rank variants of one real spell id into a single row', () => {
    const first = track({
      spell_id: 10_002_020_007,
      name: 'spell:20007/1',
      applications: 5,
      uptime_ms: 63_636,
    });
    const second = track({
      spell_id: 20_002_020_007,
      name: 'spell:20007/2',
      applications: 3,
      uptime_ms: 43_712,
    });
    const merged = sanitizeAuraTracks([first, second]);
    expect(merged).toHaveLength(1);
    expect(merged[0]).toMatchObject({
      name: 'spell:20007',
      spell_id: 20007,
      applications: 8,
      uptime_ms: 107_348,
    });
  });

  it('never merges two different spell ids, even if they would resolve to the same display name', () => {
    // sanitizeAuraTracks never calls resolveActionName and is given no name table here:
    // if it merged by name it would have nothing to merge by, so two distinct ids with
    // real uptime must both survive as two rows.
    const rankOne = track({ spell_id: 10, name: 'spell:10', uptime_ms: 2000, applications: 2 });
    const rankSix = track({ spell_id: 6141, name: 'spell:6141', uptime_ms: 3000, applications: 3 });
    const result = sanitizeAuraTracks([rankOne, rankSix]);
    expect(result.map((row) => row.spell_id).sort()).toEqual([10, 6141]);
  });

  it('replaces a self-applied source written as the target’s own name with its guid', () => {
    // sim/adapter/adapter.go writes Appliers: []string{u.Name} unconditionally -- the
    // player's display name, not a GUID, for every sim aura row. AuraTable.svelte's "from
    // X" line only recognises a GUID equal to target_guid as self-applied; without this
    // fix the applier lookup fails and the row reads "from an unnamed source".
    const selfBuff = track({ target_guid: 'sim-player', target_name: 'Mage', appliers: ['Mage'] });
    const [fixed] = sanitizeAuraTracks([selfBuff]);
    expect(fixed.appliers).toEqual(['sim-player']);
  });

  it('leaves an applier alone when it does not match the target name (a real, non-self source)', () => {
    const raidBuff = track({ target_guid: 'sim-player', target_name: 'Mage', appliers: ['priest-guid'] });
    const [unchanged] = sanitizeAuraTracks([raidBuff]);
    expect(unchanged.appliers).toEqual(['priest-guid']);
  });

  it('does not mutate its input', () => {
    const rows = [track({})];
    const before = JSON.parse(JSON.stringify(rows)) as AuraTrack[];
    sanitizeAuraTracks(rows);
    expect(rows).toEqual(before);
  });
});
