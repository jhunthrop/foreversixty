// web/src/lib/sim/rotations.test.ts
// rotations.ts reads src/data/generated/rotations.json, synced by
// scripts/sync-rotations.mjs from data/curated/apl/*.json -- `npm run sync:rotations`
// (chained into the `sync` script every pretest/precheck/predev/prebuild hook already runs)
// must have run before this file does. mage-frost is the curated fixture the brief itself
// names (data/curated/apl/mage-frost.json: the Frostbolt note and the wand Shoot note).
import { describe, expect, it } from 'vitest';
import { rotationDrawerContent, rotationNotesFor } from './rotations';

const FROSTBOLT_NOTE = 'Frostbolt is the filler between Fingers of Frost charges.';
const WAND_NOTE =
  "Shoot only when nothing else can be cast, which in practice means out of mana: a wand shot cannot start during the global cooldown and holds you for the wand's full speed, so it never fits between spells.";

describe('rotationNotesFor', () => {
  // Pins the two notes the brief names and their order, not the whole list: the curated
  // rotation gains steps as the model improves (data contracts are floors).
  it("carries mage-frost's curated step notes, the Frostbolt filler before the wand", () => {
    const notes = rotationNotesFor('mage-frost');
    const filler = notes.indexOf(FROSTBOLT_NOTE);
    expect(filler).toBeGreaterThanOrEqual(0);
    expect(notes.indexOf(WAND_NOTE)).toBeGreaterThan(filler);
  });

  it('answers an empty list for a spec sync-rotations.mjs never wrote an entry for', () => {
    expect(rotationNotesFor('not-a-real-spec')).toEqual([]);
  });
});

describe('rotationDrawerContent', () => {
  it('names the spec (its bare display name, the same form rotationCardBody uses) in the intro, and carries its ordered steps', () => {
    const content = rotationDrawerContent('mage-frost');
    expect(content.intro).toContain('Frost');
    expect(content.steps).toEqual(rotationNotesFor('mage-frost'));
    expect(content.steps).toContain(FROSTBOLT_NOTE);
  });

  it('carries an empty step list, not an error, for an unknown spec', () => {
    const content = rotationDrawerContent('not-a-real-spec');
    expect(content.steps).toEqual([]);
  });
});
