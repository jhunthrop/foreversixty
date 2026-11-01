// The W3-C8 fixtures (docs/contracts/fixtures/app/cdn.search.json, cdn.item-uses.json) are held
// to what this generator writes for the active build: every field a fixture element names is in
// every generated element of the same list. The Go contract tests hold the fixtures to
// api/openapi.yaml; this holds the generator to the fixtures. Skipped on a checkout with no
// build data. Floors only: no count or id is pinned.
import { readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { buildSearchIndex, type SearchIndex } from './search-index';
import { buildItemUses, type ItemUses, type ItemUsesShard } from './item-uses';
import {
  ACTIVE_BUILD,
  hasBuildData,
  loadLootFile,
  readBisFiles,
  readClassItems,
  readSpecCatalog,
  readTalentFiles,
} from './read-build';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const FIXTURES = path.resolve(HERE, '../../../../docs/contracts/fixtures/app');
const GUIDES = path.resolve(HERE, '../../content/guides');

function fixture<T>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, `${name}.json`), 'utf8')) as T;
}

/** Fields every fixture element names; optional ones (a quest's min_level) are left out. */
function requiredKeys(rows: readonly object[], optional: readonly string[] = []): string[] {
  return [...new Set(rows.flatMap((row) => Object.keys(row)))].filter((key) => !optional.includes(key));
}

function expectEveryRowHas(rows: readonly object[], keys: readonly string[]): void {
  for (const row of rows) expect(Object.keys(row)).toEqual(expect.arrayContaining([...keys]));
}

function guidePaths(): Set<string> {
  return new Set(
    readdirSync(GUIDES, { withFileTypes: true })
      .filter((entry) => entry.isDirectory())
      .flatMap((entry) =>
        readdirSync(path.join(GUIDES, entry.name))
          .filter((name) => name.endsWith('.md') && name !== 'index.md')
          .map((name) => `${entry.name}/${name.replace(/\.md$/, '')}`),
      ),
  );
}

describe.skipIf(!hasBuildData(ACTIVE_BUILD))('W3-C8 generator against its fixtures', () => {
  const index = buildSearchIndex({
    build: ACTIVE_BUILD,
    classItems: readClassItems(ACTIVE_BUILD),
    talents: readTalentFiles(ACTIVE_BUILD),
    specs: readSpecCatalog(),
    guidePaths: guidePaths(),
  });

  it('search.json carries every field cdn.search names', () => {
    const want = fixture<SearchIndex>('cdn.search');
    expect(Object.keys(index).sort()).toEqual(Object.keys(want).sort());
    expectEveryRowHas(index.items, requiredKeys(want.items));
    expectEveryRowHas(index.talents, requiredKeys(want.talents));
    expectEveryRowHas(index.guides, requiredKeys(want.guides));
    expect(index.guides.length).toBeGreaterThan(0);
  });

  it('item-uses carries every field cdn.item-uses names', () => {
    const want = fixture<ItemUsesShard>('cdn.item-uses');
    const shards = buildItemUses({
      build: ACTIVE_BUILD,
      itemIds: new Set(index.items.map((item) => item.id)),
      loot: loadLootFile(ACTIVE_BUILD),
      bisFiles: readBisFiles(ACTIVE_BUILD),
      specs: readSpecCatalog(),
    });
    const wantUses: ItemUses[] = Object.values(want.items);
    const gotUses = shards.flatMap((shard) => Object.values(shard.items));
    expect(Object.keys(shards[0]).sort()).toEqual(Object.keys(want).sort());
    expectEveryRowHas(gotUses, requiredKeys(wantUses));
    expectEveryRowHas(
      gotUses.flatMap((uses) => uses.sources),
      requiredKeys(
        wantUses.flatMap((uses) => uses.sources),
        ['place', 'chance', 'min_level'],
      ),
    );
    expectEveryRowHas(
      gotUses.flatMap((uses) => uses.specs),
      requiredKeys(wantUses.flatMap((uses) => uses.specs)),
    );
  });
});
