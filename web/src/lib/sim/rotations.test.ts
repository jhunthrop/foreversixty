// web/src/lib/sim/rotations.test.ts
// rotations.ts reads src/data/generated/rotations.json, synced by
// scripts/sync-rotations.mjs from data/curated/apl/*.json -- `npm run sync:rotations`
// (chained into the `sync` script every pretest/precheck/predev/prebuild hook already runs)
// must have run before this file does. mage-frost is the curated fixture the brief itself
// names (data/curated/apl/mage-frost.json: one note, "Frostbolt is the whole rotation.").
import { describe, expect, it } from 'vitest';
import { rotationDrawerContent, rotationNotesFor } from './rotations';

describe('rotationNotesFor', () => {
  it("carries mage-frost's one curated step note", () => {
    expect(rotationNotesFor('mage-frost')).toEqual(['Frostbolt is the whole rotation.']);
  });

  it('answers an empty list for a spec sync-rotations.mjs never wrote an entry for', () => {
    expect(rotationNotesFor('not-a-real-spec')).toEqual([]);
  });
});

describe('rotationDrawerContent', () => {
  it('names the spec (its bare display name, the same form rotationCardBody uses) in the intro, and carries its ordered steps', () => {
    const content = rotationDrawerContent('mage-frost');
    expect(content.intro).toContain('Frost');
    expect(content.steps).toEqual(['Frostbolt is the whole rotation.']);
  });

  it('carries an empty step list, not an error, for an unknown spec', () => {
    const content = rotationDrawerContent('not-a-real-spec');
    expect(content.steps).toEqual([]);
  });
});
