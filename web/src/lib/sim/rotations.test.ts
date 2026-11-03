// web/src/lib/sim/rotations.test.ts
// rotations.ts reads src/data/generated/rotations.json, synced by
// scripts/sync-rotations.mjs from data/curated/apl/*.json -- `npm run sync:rotations`
// (chained into the `sync` script every pretest/precheck/predev/prebuild hook already runs)
// must have run before this file does. mage-frost is the curated fixture the brief itself
// names (data/curated/apl/mage-frost.json: the Frostbolt note and the wand Shoot note).
import { beforeAll, describe, expect, it } from 'vitest';
import { loadRotationNotes, rotationDrawerContent, rotationNotesFor, type RotationNotes } from './rotations';

let notes: RotationNotes;
beforeAll(async () => {
  notes = await loadRotationNotes();
});

const FROSTBOLT_NOTE =
  'Frostbolt is the filler, and every cast can roll Fingers of Frost or Missile Barrage. Never stop casting to wait for a proc.';
const WAND_NOTE =
  "Shoot only when you are out of mana and waiting on a gem, a rune or Evocation. A wand shot holds you for the wand's full speed, so never weave it between spells.";

describe('rotationNotesFor', () => {
  // Pins the two notes the brief names and their order, not the whole list: the curated
  // rotation gains steps as the model improves (data contracts are floors).
  it("carries mage-frost's curated step notes, the Frostbolt filler before the wand", () => {
    const frost = rotationNotesFor('mage-frost', notes);
    const filler = frost.indexOf(FROSTBOLT_NOTE);
    expect(filler).toBeGreaterThanOrEqual(0);
    expect(frost.indexOf(WAND_NOTE)).toBeGreaterThan(filler);
  });

  it('answers an empty list for a spec sync-rotations.mjs never wrote an entry for', () => {
    expect(rotationNotesFor('not-a-real-spec', notes)).toEqual([]);
  });
});

describe('loadRotationNotes', () => {
  it('loads the notes once: every caller shares the first load', () => {
    expect(loadRotationNotes()).toBe(loadRotationNotes());
  });
});

describe('rotationDrawerContent', () => {
  it('has no steps yet while the notes load, so the drawer says nothing rather than "none"', () => {
    expect(rotationDrawerContent('mage-frost', null).steps).toBeNull();
  });

  it('names the spec (its bare display name, the same form rotationCardBody uses) in the intro, and carries its ordered steps', () => {
    const content = rotationDrawerContent('mage-frost', notes);
    expect(content.intro).toContain('Frost');
    expect(content.steps).toEqual(rotationNotesFor('mage-frost', notes));
    expect(content.steps).toContain(FROSTBOLT_NOTE);
  });

  it('carries an empty step list, not an error, for an unknown spec', () => {
    const content = rotationDrawerContent('not-a-real-spec', notes);
    expect(content.steps).toEqual([]);
  });
});
