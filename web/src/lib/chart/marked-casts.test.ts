// web/src/lib/chart/marked-casts.test.ts
// Tenet 8: every marked cast is checked by spell id against the ACTIVE client build's spell
// table, the same file the builds test reads, so a build that drops a spell fails here.
import { existsSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { describe, expect, it } from 'vitest';
import { MARKED_CASTS, markedCast } from './marked-casts';

const repoRoot = join(__dirname, '../../../..');
const build = (
  JSON.parse(readFileSync(join(repoRoot, 'web/src/data/active-build.json'), 'utf8')) as { build: string }
).build;
const spellsFile = join(repoRoot, 'data/builds', build, 'spells.json');
const spells = new Map(
  (JSON.parse(readFileSync(spellsFile, 'utf8')) as { id: number; name: string | null }[]).map((row) => [
    row.id,
    row.name,
  ]),
);

/** The client names a few effects differently from the cast a raider says. */
const CLIENT_NAMES: Record<string, readonly string[]> = {
  Rebirth: ['Rebirth'],
  'Soulstone Resurrection': ['Soulstone Resurrection'],
  'Power Infusion': ['Power Infusion'],
  'Mana Tide Totem': ['Mana Tide Totem'],
  'Major Healing Potion': ['Healing Potion', 'Major Healing Potion'],
  Healthstone: [
    'Minor Healthstone',
    'Lesser Healthstone',
    'Healthstone',
    'Greater Healthstone',
    'Major Healthstone',
  ],
};

describe(`marked casts against the active build (${build})`, () => {
  it('reads the build’s spell table', () => {
    expect(existsSync(spellsFile)).toBe(true);
  });

  for (const cast of MARKED_CASTS) {
    it(`${cast.label}: every id exists in Forever under an expected name`, () => {
      for (const id of cast.ids) {
        expect(spells.has(id), `${cast.label} ${id}`).toBe(true);
        expect(CLIENT_NAMES[cast.label], `${cast.label} ${id}`).toContain(spells.get(id));
      }
    });
  }

  it('icons named here ship with the build', () => {
    for (const cast of MARKED_CASTS) {
      if (cast.icon === undefined) continue;
      expect(existsSync(join(repoRoot, 'data/builds', build, 'icons', `${cast.icon}.webp`)), cast.icon).toBe(
        true,
      );
    }
  });

  it('finds a cast by id, not by name', () => {
    expect(markedCast(20484)?.label).toBe('Rebirth');
    expect(markedCast(17534)?.label).toBe('Major Healing Potion');
    expect(markedCast(1)).toBeUndefined();
  });
});
