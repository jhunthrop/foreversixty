// web/src/lib/sim/preset-sets.test.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { loadPresetSets, raidSetAt } from './preset-sets';

const GENERATED = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../sim/request/raid-consumes.json',
);

// The table is a lazily loaded chunk (the islands' size budgets): asking for a spec's set
// before it has loaded is a programming error, said loudly, never an empty preset.
describe('raid sets, loaded on demand', () => {
  it('refuses a spec before the table has loaded, and answers "no character" with nothing', () => {
    expect(() => raidSetAt('rogue-combat', 60)).toThrow(/loadPresetSets/);
    expect(raidSetAt('', 60)).toBeUndefined();
  });

  it('answers every spec at every level from the generated rows once loaded', async () => {
    await loadPresetSets();
    const rows = (
      JSON.parse(readFileSync(GENERATED, 'utf8')) as {
        specs: Record<string, { from_level: number; buffs: string[]; consumes: string[] }[]>;
      }
    ).specs;
    for (const [spec, specRows] of Object.entries(rows)) {
      for (let level = 1; level <= 60; level += 1) {
        const want = specRows.filter((row) => row.from_level <= level).at(-1)!;
        expect(raidSetAt(spec, level), `${spec} ${level}`).toEqual(want);
      }
    }
    expect(raidSetAt('not-a-real-spec', 60)).toBeUndefined();
  });
});
