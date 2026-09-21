import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { existsSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { resolveDataSource, SYNC_ENTRIES, syncData, writeSimNames } from '../../../scripts/sync-data.mjs';

const BUILD = '1.15.9.69722';
const OTHER_BUILD = '1.16.0.70000';
const silent = { log: () => {}, warn: () => {} };

let root: string;
let repoRoot: string;
let webRoot: string;

function writeJson(file: string, value: unknown): void {
  mkdirSync(path.dirname(file), { recursive: true });
  writeFileSync(file, `${JSON.stringify(value, null, 2)}\n`, 'utf8');
}

/** A web/ tree with active-build.json and a fixture build, matching the real layout. */
function scaffoldWeb(): void {
  writeJson(path.join(webRoot, 'src/data/active-build.json'), { build: BUILD });
  const fixture = path.join(webRoot, 'src/fixtures/planner');
  writeJson(path.join(fixture, 'talents/warrior.json'), { build: BUILD, class_slug: 'warrior', trees: [] });
  writeJson(path.join(fixture, 'items/warrior.json'), { build: BUILD, class_slug: 'warrior', items: [] });
  writeJson(path.join(fixture, 'sets.json'), []);
  writeJson(path.join(fixture, 'classes.json'), [{ id: 1, slug: 'warrior' }]);
  writeJson(path.join(fixture, 'races.json'), [{ id: 1, slug: 'human' }]);
  writeJson(path.join(fixture, 'combos.json'), [{ race_id: 1, class_id: 1, new_in_forever: false }]);
  mkdirSync(path.join(fixture, 'icons'), { recursive: true });
  writeFileSync(path.join(fixture, 'icons/fixture_a.webp'), 'x');
  writeJson(path.join(fixture, 'manifest.json'), { build: BUILD, fixture: true, files: {} });
}

/** A data/builds/<build>/ tree carrying real Phase 1 files, one class per build. */
function scaffoldRealBuild(build = BUILD, slug = 'paladin'): string {
  const dir = path.join(repoRoot, 'data/builds', build);
  writeJson(path.join(dir, `talents/${slug}.json`), { build, class_slug: slug, trees: [] });
  writeJson(path.join(dir, 'classes.json'), [{ id: 2, slug }]);
  writeJson(path.join(dir, 'races.json'), [{ id: 5, slug: 'undead' }]);
  writeJson(path.join(dir, 'combos.json'), [{ race_id: 5, class_id: 2, new_in_forever: true }]);
  writeJson(path.join(dir, 'manifest.json'), {
    build,
    files: {
      [`talents/${slug}.json`]: 'aa',
      'classes.json': 'bb',
      'races.json': 'cc',
      'combos.json': 'dd',
    },
  });
  return dir;
}

/** A data/builds/<build>/ tree with only the Phase 0 flat files, as Era ships today. */
function scaffoldPhaseZeroBuild(build: string): void {
  const dir = path.join(repoRoot, 'data/builds', build);
  writeJson(path.join(dir, 'talents.json'), []);
  writeJson(path.join(dir, 'manifest.json'), { build, files: { 'talents.json': 'aa' } });
}

beforeEach(() => {
  root = mkdtempSync(path.join(tmpdir(), 'sync-data-'));
  repoRoot = path.join(root, 'repo');
  webRoot = path.join(repoRoot, 'web');
  scaffoldWeb();
});

afterEach(() => rmSync(root, { recursive: true, force: true }));

describe('syncData', () => {
  it('falls back to the fixture when the build has no Phase 1 files', async () => {
    const result = await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(result.usedFixture).toBe(true);
    expect(result.build).toBe(BUILD);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/warrior.json'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'icons/fixture_a.webp'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'combos.json'))).toBe(true);
  });

  it('writes the three page-imported files into src/data/generated', async () => {
    await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    for (const name of ['classes.json', 'races.json', 'combos.json']) {
      const file = path.join(webRoot, 'src/data/generated', name);
      expect(existsSync(file)).toBe(true);
      expect(() => JSON.parse(readFileSync(file, 'utf8'))).not.toThrow();
    }
  });

  it('refuses the fixture fallback when fixtures are not allowed', async () => {
    await expect(syncData({ repoRoot, webRoot, allowFixture: false, log: silent })).rejects.toThrow(
      /no Phase 1 talent data/,
    );
  });

  it('copies the real build when the manifest lists Phase 1 files', async () => {
    scaffoldRealBuild();
    const result = await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(result.usedFixture).toBe(false);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/paladin.json'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/warrior.json'))).toBe(false);
  });

  it('fails loudly and names every manifest file missing from disk', async () => {
    const dir = scaffoldRealBuild();
    rmSync(path.join(dir, 'races.json'));
    await expect(syncData({ repoRoot, webRoot, allowFixture: true, log: silent })).rejects.toThrow(
      /manifest\.json lists files that are missing: races\.json/,
    );
  });

  it('fails when a required file is absent from the source directory', async () => {
    const dir = scaffoldRealBuild();
    rmSync(path.join(dir, 'combos.json'));
    writeJson(path.join(dir, 'manifest.json'), {
      build: BUILD,
      files: { 'talents/paladin.json': 'aa', 'classes.json': 'bb', 'races.json': 'cc' },
    });
    await expect(syncData({ repoRoot, webRoot, allowFixture: true, log: silent })).rejects.toThrow(
      /required entry combos\.json is missing/,
    );
  });

  it('replaces a stale public/data directory instead of merging into it', async () => {
    const stale = path.join(webRoot, 'public/data', BUILD, 'talents/stale.json');
    writeJson(stale, { gone: true });
    await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(existsSync(stale)).toBe(false);
  });

  it('publishes every build carrying Phase 1 data, not only the active one', async () => {
    scaffoldRealBuild();
    scaffoldRealBuild(OTHER_BUILD, 'mage');
    const result = await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(result.published).toEqual([BUILD, OTHER_BUILD]);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/paladin.json'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', OTHER_BUILD, 'talents/mage.json'))).toBe(true);
    // A retained build has to carry the reference files too: loadReference() fetches
    // classes.json, races.json and combos.json from the shared build's own directory.
    for (const name of ['classes.json', 'races.json', 'combos.json']) {
      expect(existsSync(path.join(webRoot, 'public/data', OTHER_BUILD, name))).toBe(true);
    }
  });

  it('skips a build with no Phase 1 data and warns naming it', async () => {
    scaffoldRealBuild();
    scaffoldPhaseZeroBuild(OTHER_BUILD);
    const warnings: string[] = [];
    const log = { log: () => {}, warn: (message: unknown) => void warnings.push(String(message)) };
    const result = await syncData({ repoRoot, webRoot, allowFixture: true, log });
    expect(result.published).toEqual([BUILD]);
    expect(existsSync(path.join(webRoot, 'public/data', OTHER_BUILD))).toBe(false);
    expect(warnings.join('\n')).toContain(OTHER_BUILD);
  });

  it('mirrors the active build alone into src/data/generated', async () => {
    scaffoldRealBuild();
    scaffoldRealBuild(OTHER_BUILD, 'mage');
    await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    const generated = readFileSync(path.join(webRoot, 'src/data/generated/classes.json'), 'utf8');
    expect(JSON.parse(generated)).toEqual([{ id: 2, slug: 'paladin' }]);
  });

  it('publishes the fixture when asked, whatever data/builds holds', async () => {
    scaffoldRealBuild();
    const result = await syncData({ repoRoot, webRoot, source: 'fixture', log: silent });
    expect(result.usedFixture).toBe(true);
    expect(result.published).toEqual([BUILD]);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/warrior.json'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/paladin.json'))).toBe(false);
    const generated = readFileSync(path.join(webRoot, 'src/data/generated/classes.json'), 'utf8');
    expect(JSON.parse(generated)).toEqual([{ id: 1, slug: 'warrior' }]);
  });

  it('does not retain other real builds when publishing the fixture', async () => {
    scaffoldRealBuild();
    scaffoldRealBuild(OTHER_BUILD, 'mage');
    const result = await syncData({ repoRoot, webRoot, source: 'fixture', log: silent });
    expect(result.published).toEqual([BUILD]);
    expect(existsSync(path.join(webRoot, 'public/data', OTHER_BUILD))).toBe(false);
  });

  it('refuses the fixture source where fixtures are not allowed', async () => {
    scaffoldRealBuild();
    await expect(
      syncData({ repoRoot, webRoot, source: 'fixture', allowFixture: false, log: silent }),
    ).rejects.toThrow(/FOREVER_DATA=fixture is refused here/);
  });

  it('removes a build public/data still holds from an earlier run', async () => {
    scaffoldRealBuild();
    scaffoldRealBuild(OTHER_BUILD, 'mage');
    await syncData({ repoRoot, webRoot, log: silent });
    rmSync(path.join(repoRoot, 'data/builds', OTHER_BUILD), { recursive: true, force: true });
    const result = await syncData({ repoRoot, webRoot, log: silent });
    expect(result.pruned).toEqual([OTHER_BUILD]);
    expect(existsSync(path.join(webRoot, 'public/data', OTHER_BUILD))).toBe(false);
  });

  it('names the published builds on stdout', async () => {
    scaffoldRealBuild();
    scaffoldRealBuild(OTHER_BUILD, 'mage');
    const lines: string[] = [];
    const log = { log: (message: unknown) => void lines.push(String(message)), warn: () => {} };
    await syncData({ repoRoot, webRoot, allowFixture: true, log });
    expect(lines.join('\n')).toContain(OTHER_BUILD);
  });
});

describe('SYNC_ENTRIES', () => {
  it('publishes the per-tree background art, and does not require it', () => {
    const trees = SYNC_ENTRIES.find((entry) => entry.name === 'trees');
    expect(trees).toEqual({ name: 'trees', kind: 'dir', required: false });
  });

  it('publishes contract 10.4’s buff name and icon table, and does not require it', () => {
    const simbuffs = SYNC_ENTRIES.find((entry) => entry.name === 'simbuffs.json');
    expect(simbuffs).toEqual({ name: 'simbuffs.json', kind: 'file', required: false });
  });

  it('publishes the engine’s known item ids, and does not require them', () => {
    const simitems = SYNC_ENTRIES.find((entry) => entry.name === 'simitems.json');
    expect(simitems).toEqual({ name: 'simitems.json', kind: 'file', required: false });
  });

  it('publishes the build’s stat weights, and does not require them', () => {
    const weights = SYNC_ENTRIES.find((entry) => entry.name === 'stat-weights.json');
    expect(weights).toEqual({ name: 'stat-weights.json', kind: 'file', required: false });
  });
});

describe('resolveDataSource', () => {
  it('defaults to real and passes the two known values through', () => {
    expect(resolveDataSource(undefined)).toBe('real');
    expect(resolveDataSource('')).toBe('real');
    expect(resolveDataSource('real')).toBe('real');
    expect(resolveDataSource('fixture')).toBe('fixture');
  });

  it('refuses an unknown value rather than guessing', () => {
    expect(() => resolveDataSource('fixtures')).toThrow(/FOREVER_DATA must be one of real, fixture/);
  });
});

describe('writeSimNames', () => {
  it('prunes spells.json to the class’s own constants and its items', async () => {
    const dir = mkdtempSync(path.join(tmpdir(), 'simnames-'));
    const build = path.join(dir, 'build');
    const out = path.join(dir, 'out');
    mkdirSync(path.join(build, 'spellconst'), { recursive: true });
    mkdirSync(path.join(build, 'items'), { recursive: true });
    mkdirSync(out, { recursive: true });
    writeFileSync(
      path.join(build, 'spells.json'),
      JSON.stringify([
        { id: 25286, name: 'Heroic Strike' },
        { id: 99999, name: 'Something Else' },
      ]),
    );
    writeFileSync(
      path.join(build, 'spellconst', 'warrior.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'warrior',
        family: 4,
        spells: { '25286': { name: 'Heroic Strike' } },
      }),
    );
    writeFileSync(
      path.join(build, 'items', 'warrior.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'warrior',
        items: [{ id: 14554, name: 'Cloudkeeper Legplates' }],
      }),
    );

    expect(await writeSimNames(build, out)).toEqual(['simnames/warrior.json', 'simnames/_shared.json']);

    const table = JSON.parse(readFileSync(path.join(out, 'simnames', 'warrior.json'), 'utf8'));
    expect(table.spell['25286']).toBe('Heroic Strike');
    expect(table.spell['99999']).toBeUndefined();
    expect(table.item['14554']).toBe('Cloudkeeper Legplates');
    rmSync(dir, { recursive: true, force: true });
  });

  // Defect 2 (2026-09-21 result-page review): a raid-buffed sim can put an aura on the
  // player from any class's spellbook, not only their own, and the class's own
  // simnames/<class>.json (above) has nowhere to resolve a spell it does not own. This is
  // the id-union table the web falls back to for exactly that case.
  it("writes simnames/_shared.json as every class's spell table unioned, plus the classless raid buffs", async () => {
    const dir = mkdtempSync(path.join(tmpdir(), 'simnames-shared-'));
    const build = path.join(dir, 'build');
    const out = path.join(dir, 'out');
    mkdirSync(path.join(build, 'spellconst'), { recursive: true });
    mkdirSync(out, { recursive: true });
    writeFileSync(path.join(build, 'spells.json'), JSON.stringify([]));
    writeFileSync(
      path.join(build, 'spellconst', 'mage.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'mage',
        family: 1,
        spells: { '133': { name: 'Fireball' } },
      }),
    );
    writeFileSync(
      path.join(build, 'spellconst', 'paladin.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'paladin',
        family: 2,
        // A mage running the raid-buffed preset can be Blessing-of-Kings'd by a paladin
        // who is not in the sim at all; the mage's own simnames/mage.json never carries
        // paladin spell ids, so this is exactly the row _shared.json exists to cover.
        spells: { '20217': { name: 'Blessing of Kings' } },
      }),
    );

    const written = await writeSimNames(build, out);
    expect(written).toContain('simnames/_shared.json');

    const shared = JSON.parse(readFileSync(path.join(out, 'simnames', '_shared.json'), 'utf8'));
    expect(shared.spell['133']).toBe('Fireball');
    expect(shared.spell['20217']).toBe('Blessing of Kings');
    // Thorns: no player class's spellbook is in this fixture at all, so only the
    // classless table (CLASSLESS_BUFF_SPELLS' Darkmoon Faire buff, the one entry small
    // enough to hand-check here) proves the union ran rather than one class's own file
    // being copied verbatim.
    expect(shared.spell['23735']).toBe("Sayge's Dark Fortune of Strength");
    rmSync(dir, { recursive: true, force: true });
  });

  // 2026-09-21 result-page review round 2: three more kinds of id reached production as
  // "Spell <n>"/"Item <n>" even after the first shared-table pass -- a talent-granted
  // passive proc (no cast bar entry, so spellconst never carries it), a racial, and a
  // consumable item (nobody's gear, so no class's own items/<class>.json has it either).
  it('unions a class’s own talent spells and a build-wide consumable/reagent item table into the shared file', async () => {
    const dir = mkdtempSync(path.join(tmpdir(), 'simnames-round2-'));
    const build = path.join(dir, 'build');
    const out = path.join(dir, 'out');
    mkdirSync(path.join(build, 'spellconst'), { recursive: true });
    mkdirSync(path.join(build, 'talents'), { recursive: true });
    mkdirSync(out, { recursive: true });
    writeFileSync(path.join(build, 'spells.json'), JSON.stringify([]));
    writeFileSync(
      path.join(build, 'spellconst', 'warrior.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'warrior',
        family: 4,
        // Flurry (a talent proc) is deliberately NOT here: spellconst is the CASTABLE
        // spellbook, and Flurry has no cast bar entry -- the talents union below is what
        // has to carry it.
        spells: {},
      }),
    );
    writeFileSync(
      path.join(build, 'talents', 'warrior.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'warrior',
        trees: [
          {
            id: 1,
            name: 'Fury',
            position: 0,
            talents: [
              {
                id: 1,
                name: 'Flurry',
                max_rank: 5,
                tier: 0,
                column: 0,
                ranks: [
                  { spell_id: 12319, description: 'r1' },
                  { spell_id: 12319, description: 'r2' },
                ],
                spell_id: 12319,
              },
            ],
          },
        ],
      }),
    );
    writeFileSync(
      path.join(build, 'items.json'),
      JSON.stringify([
        { id: 13442, name: 'Mighty Rage Potion', class_id: 0 },
        { id: 5514, name: 'Mana Agate', class_id: 4 }, // misclassified as Trade Goods
        { id: 12345, name: 'Some Random Crafting Mat', class_id: 4 }, // must NOT leak in
        { id: 20560, name: 'Thunderfury, Blessed Blade of the Windseeker', class_id: 2 }, // gear, must NOT leak in
      ]),
    );

    await writeSimNames(build, out);
    const shared = JSON.parse(readFileSync(path.join(out, 'simnames', '_shared.json'), 'utf8'));
    expect(shared.spell['12319']).toBe('Flurry');
    expect(shared.spell['20572']).toBe('Blood Fury'); // RACIAL_SPELLS
    expect(shared.item['13442']).toBe('Mighty Rage Potion');
    expect(shared.item['5514']).toBe('Mana Agate'); // ITEM_NAME_OVERRIDES
    expect(shared.item['12345']).toBeUndefined();
    expect(shared.item['20560']).toBeUndefined();
    rmSync(dir, { recursive: true, force: true });
  });

  it('publishes nothing rather than failing when the build has no spellconst', async () => {
    const dir = mkdtempSync(path.join(tmpdir(), 'simnames-none-'));
    mkdirSync(path.join(dir, 'build'), { recursive: true });
    mkdirSync(path.join(dir, 'out'), { recursive: true });
    expect(await writeSimNames(path.join(dir, 'build'), path.join(dir, 'out'))).toEqual([]);
    rmSync(dir, { recursive: true, force: true });
  });

  it('never adds a spell id from the item list -- items/<class>.json carries no proc or on-use field', async () => {
    // Pins the doc comment's corrected claim: `spell` membership comes only from
    // spellconst, never from `item`, because the normalized item rows this lane reads
    // (id, name, icon, slot, quality, required_level, item_level, armor, stats, set_id,
    // unique) have nowhere to carry a triggered spell id even if this function wanted to
    // read one. An id that is both a trinket's item id and, coincidentally, a real spell
    // id in spells.json must not leak into `spell` just because it showed up in `items`.
    const dir = mkdtempSync(path.join(tmpdir(), 'simnames-item-spell-'));
    const build = path.join(dir, 'build');
    const out = path.join(dir, 'out');
    mkdirSync(path.join(build, 'spellconst'), { recursive: true });
    mkdirSync(path.join(build, 'items'), { recursive: true });
    mkdirSync(out, { recursive: true });
    writeFileSync(
      path.join(build, 'spells.json'),
      JSON.stringify([
        { id: 25286, name: 'Heroic Strike' },
        { id: 14554, name: 'Some Proc Effect' },
      ]),
    );
    writeFileSync(
      path.join(build, 'spellconst', 'warrior.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'warrior',
        family: 4,
        spells: { '25286': { name: 'Heroic Strike' } },
      }),
    );
    writeFileSync(
      path.join(build, 'items', 'warrior.json'),
      JSON.stringify({
        build: '1.60.1.69893',
        class_slug: 'warrior',
        // id 14554 collides with the spells.json row above on purpose.
        items: [{ id: 14554, name: 'Cloudkeeper Legplates' }],
      }),
    );

    await writeSimNames(build, out);

    const table = JSON.parse(readFileSync(path.join(out, 'simnames', 'warrior.json'), 'utf8'));
    expect(table.item['14554']).toBe('Cloudkeeper Legplates');
    expect(table.spell['14554']).toBeUndefined();
    rmSync(dir, { recursive: true, force: true });
  });
});
