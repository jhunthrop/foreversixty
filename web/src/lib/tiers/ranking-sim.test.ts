// web/src/lib/tiers/ranking-sim.test.ts
// "View ranking sim" reproduces the tier number's request. Over the real published files of the
// active build, every DPS row's link is decoded the way /sim's `?code=` load does
// (`characterFromFs1`, then the store's `withSpecForPreset` defaults) and built into a request
// by the store's own `buildRequest`; that request must be the one the ranker ran for the band:
// the same gear and enchants, talents, race, level, raid buffs and consumables, and fight. The
// caster stand-off the ranker uses is not in the request; sim/request applies it for every
// caster spec (TestCasterSpecsStandAtCasterDistance).
import { render } from 'svelte/server';
import { beforeAll, describe, expect, it } from 'vitest';
import RankRow from '../../components/ui/RankRow.svelte';
import activeBuild from '../../data/active-build.json';
import { bandEntry, loadBisFile, readSpecCatalog } from '../bis/load';
import type { BisBand, BisFile, Faction } from '../bis/types';
import { loadTalentFile } from '../guides/build-tree';
import { classRows, raceRows } from '../planner/reference';
import { SLOTS, type Slot, type TalentFile } from '../planner/types';
import { characterFromFs1, MAX_LEVEL, type SimCharacter } from '../sim/character';
import { loadPresetSets, raidSetAt } from '../sim/preset-sets';
import { defaultSettings } from '../sim/settings';
import { isSimulatedSpec } from '../sim/spec-label';
import { createRequestMethods, type StoreRequestDeps } from '../sim/store-request';
import type { SimRequest } from '../sim/types';
import { parseSimState } from '../sim/url';
import { loadTierData } from './tier-load';
import { hrefsFor, type TierRow } from './tier-list';
import { rankRowProps } from './tier-row-view';

const BUILD = activeBuild.build;
const FACTIONS: readonly Faction[] = ['alliance', 'horde'];
/** The ranker's fight (sim/leveling.RankerEncounter == api.DefaultEncounter). */
const RANKER_ENCOUNTER = {
  duration_sec: 180,
  variation: 0.2,
  targets: 1,
  execute_ratio: 0.25,
  target_level: 63,
};

function talentsFor(classSlug: string): TalentFile {
  const file = loadTalentFile(BUILD, classSlug);
  if (file === undefined) throw new Error(`no talents for ${classSlug} at ${BUILD}`);
  return file;
}

/** The store's `buildRequest` over one decoded character, with the settings a `?code=` load
 *  adopts; anything else the store's deps carry is never reached. */
function storeRequest(character: SimCharacter, talents: TalentFile): SimRequest {
  const used: Partial<StoreRequestDeps> = {
    getCharacter: () => character,
    getTalents: () => talents,
    getSettings: () => defaultSettings(character.spec, character.level),
    getPrecisionId: () => 'normal',
  };
  const deps = new Proxy(used, {
    get(target, key) {
      if (key in target) return target[key as keyof StoreRequestDeps];
      throw new Error(`buildRequest reached deps.${String(key)}`);
    },
  }) as StoreRequestDeps;
  const request = createRequestMethods(deps).buildRequest();
  if (request === null) throw new Error(`${character.spec}: the store built no request`);
  return request;
}

function decode(row: TierRow): { character: SimCharacter; talents: TalentFile } {
  expect(row.simHref, `${row.spec}: a DPS row with no sim link`).toBeDefined();
  const href = row.simHref!;
  expect(href.startsWith('/sim?')).toBe(true);
  const state = parseSimState(href.slice('/sim'.length));
  expect(state.run, `${row.spec}: the link does not start the run`).toBe(true);
  const talents = talentsFor(row.classSlug);
  const result = characterFromFs1(state.code, talents, classRows, raceRows, {
    kind: 'manual',
    ref: '',
    captured_at: '',
  });
  if (!result.ok) throw new Error(`${row.spec}: /sim refuses the code: ${result.message}`);
  return { character: result.character, talents };
}

function isSlot(slot: string): slot is Slot {
  return (SLOTS as readonly string[]).includes(slot);
}

function bandGear(band: BisBand): { slot: string; item_id: number; enchant?: number }[] {
  return band.slots
    .filter((row) => isSlot(row.slot) && row.item_id > 0)
    .map((row) => ({
      slot: row.slot,
      item_id: row.item_id,
      ...(row.enchant_id ? { enchant: row.enchant_id } : {}),
    }));
}

function bySlot<T extends { slot: string }>(rows: readonly T[]): T[] {
  return [...rows].sort((a, b) => a.slot.localeCompare(b.slot));
}

/** A talent string with each tree's trailing zeros dropped: the engine reads "5-…" and
 *  "5000000000000000-…" as the same points, and the encoder writes the short form. */
function spentPoints(talents: string): string {
  return talents
    .split('-')
    .map((tree) => tree.replace(/0+$/, ''))
    .join('-');
}

function escapeRegExp(text: string): string {
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

function raidBuffIds(file: BisFile): string[] {
  const raid = file.presets?.raid;
  if (raid === undefined) throw new Error(`${file.spec}: no raid preset in the file`);
  // The files write the request ids ("arcane_brilliance"); BisPresetEffect types them as numbers.
  return [...raid.buffs, ...raid.debuffs].map((row) => String(row.id)).sort();
}

describe('View ranking sim reproduces the tier number', () => {
  beforeAll(() => loadPresetSets());

  const data = loadTierData();
  const dpsSpecs = readSpecCatalog().filter((entry) => isSimulatedSpec(entry.spec));

  it('gives every published DPS spec a row with a sim link, and no tank or healer row one', () => {
    for (const faction of FACTIONS) {
      const lists = data.lists[faction];
      expect(lists.dps.length).toBeGreaterThan(0);
      const published = dpsSpecs.filter((entry) => {
        const file = loadBisFile(entry.spec, BUILD);
        return file !== null && bandEntry(file, MAX_LEVEL, faction) !== undefined;
      });
      expect(lists.dps.map((row) => row.spec).sort()).toEqual(published.map((entry) => entry.spec).sort());
      expect(lists.dps.every((row) => row.simHref !== undefined)).toBe(true);
      expect([...lists.tank, ...lists.healer].every((row) => row.simHref === undefined)).toBe(true);
    }
  });

  it('renders every DPS row with its figure as an anchor to the sim run, and its name to BiS', () => {
    const escape = (href: string) => href.replaceAll('&', '&amp;');
    for (const faction of FACTIONS) {
      const rows = data.lists[faction].dps;
      for (const row of rows) {
        const body = render(RankRow, { props: rankRowProps(row, rows) }).body;
        const simAnchor = new RegExp(
          `<a[^>]*href="${escapeRegExp(escape(row.simHref!))}"[^>]*data-testid="tier-number"`,
        );
        expect(body, `${faction} ${row.spec}`).toMatch(simAnchor);
        expect(body).toContain(`href="${hrefsFor(row).bis}"`);
      }
    }
  });

  for (const faction of FACTIONS) {
    it(`sends the ranker's request for every ${faction} DPS row`, () => {
      for (const row of data.lists[faction].dps) {
        const file = loadBisFile(row.spec, BUILD)!;
        const band = bandEntry(file, MAX_LEVEL, faction)!;
        const label = `${row.spec} ${faction}`;
        expect(band.preset, `${label}: the tier row is not the raid band`).toBe('raid');
        const { character, talents } = decode(row);
        const request = storeRequest(character, talents);
        expect(request.spec, `${label}: /sim names another spec`).toBe(band.spec);
        expect(request.character.level, label).toBe(MAX_LEVEL);
        expect(request.character.race, label).toBe(band.race);
        // bisSimHref keeps only the points the trees reach; a cut here means the link runs a
        // build the ranker did not, so it fails rather than passing on the cut string.
        expect(
          spentPoints(request.character.talents),
          `${label}: the band's talents were cut to reach the trees`,
        ).toBe(spentPoints(band.talents));
        expect(bySlot(request.character.gear), `${label}: gear`).toEqual(bySlot(bandGear(band)));
        expect([...request.character.buffs].sort(), `${label}: buffs`).toEqual(raidBuffIds(file));
        expect(request.character.consumes, `${label}: consumables`).toEqual([
          ...(raidSetAt(row.spec, MAX_LEVEL)?.consumes ?? []),
        ]);
        expect(request.encounter, `${label}: fight`).toMatchObject(RANKER_ENCOUNTER);
      }
    });
  }
});
