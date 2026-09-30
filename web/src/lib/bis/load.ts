// web/src/lib/bis/load.ts
// Build-time data for the /bis pages. Astro's static pages render in Node at build time
// (getStaticPaths and page frontmatter both run there, never in the browser), so this reads
// straight off disk with node:fs the same way scripts/sync-data.mjs and
// src/lib/sim/no-raw-ids.test.ts's REPO_ROOT reach the repo's data/ directory -- a plain fs
// read is outside Vite's module graph and so is never subject to its dev-server fs.allow
// restriction the way an `import` of a repo-root file would be.
//
// A spec's real file (lane bis-all, data/builds/<build>/bis/<spec>.json) is not on main yet.
// Until it lands, every written spec falls back to this lane's own fixture
// (src/data/fixtures/bis/<spec>.json) when one is committed, and to the explicit "no list
// yet" empty state otherwise -- never a guess. The moment a real file exists for a build,
// it wins over the fixture with no code change here.
import { existsSync, readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { SLOTS, type Slot } from '../planner/types';
import { LOOT_KINDS, SOURCE_KIND_LABELS, type LootFile, type LootKind } from '../sim/loot';
import { bisCopy } from './copy';
import { hasKnownSource } from './source-cell';
import type {
  BisBand,
  BisFile,
  BisSlot,
  ChangedSlot,
  Faction,
  ItemDetail,
  ItemHoverModel,
  LootQuestsFile,
  SpecCatalogEntry,
} from './types';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const WEB_ROOT = path.resolve(REPO_ROOT, 'web');

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

/** The master list of written specs (data/curated/specs.json) -- the route list for /bis
 *  comes from here, not from which specs happen to have a BiS file, so a spec with nothing
 *  ranked yet still gets a page with the "no list yet" empty state instead of a 404. */
export function readSpecCatalog(): SpecCatalogEntry[] {
  return readJson<SpecCatalogEntry[]>(path.join(REPO_ROOT, 'data/curated/specs.json'));
}

export interface ClassSpecs {
  classSlug: string;
  specs: SpecCatalogEntry[];
}

/** The catalog grouped by class, class order preserved as data/curated/specs.json lists
 *  them (first-seen order), specs within a class in file order. */
export function groupByClass(catalog: readonly SpecCatalogEntry[]): ClassSpecs[] {
  const order: string[] = [];
  const byClass = new Map<string, SpecCatalogEntry[]>();
  for (const entry of catalog) {
    if (!byClass.has(entry.class_slug)) {
      byClass.set(entry.class_slug, []);
      order.push(entry.class_slug);
    }
    byClass.get(entry.class_slug)!.push(entry);
  }
  return order.map((classSlug) => ({ classSlug, specs: byClass.get(classSlug)! }));
}

function realBisPath(build: string, spec: string): string {
  return path.join(REPO_ROOT, 'data/builds', build, 'bis', `${spec}.json`);
}

function fixtureBisPath(spec: string): string {
  return path.join(WEB_ROOT, 'src/data/fixtures/bis', `${spec}.json`);
}

/**
 * A spec's BiS file: the real, pipeline-published one when data/builds/<build>/bis/<spec>.
 * json exists, else this lane's committed fixture, else null (no list published for this
 * spec yet -- the page's own empty state, not an error).
 */
export function loadBisFile(spec: string, build: string): BisFile | null {
  const real = realBisPath(build, spec);
  if (existsSync(real)) return normaliseBisFile(readJson<BisFile>(real));
  const fixture = fixtureBisPath(spec);
  if (existsSync(fixture)) return normaliseBisFile(readJson<BisFile>(fixture));
  return null;
}

/** The nightly's first files wrote `null` for an empty "new at this band" list (a Go nil
 *  slice); the contract is an array, so every band is read as one. Also defaults `coverage`
 *  (lane `rank-guardrails`' guardrail A, `report.go`'s own `Coverage` field) to `{}`: a file
 *  published before that field existed carries no `coverage` key at all (`undefined` here),
 *  and a Go nil map also encodes as JSON `null` -- either way the contract is an object, the
 *  same reasoning `new_at_band` already applies to its own array.
 *
 *  Also defaults each slot's `alternatives` (`sim/cmd/leveling-bis/report.go`'s own
 *  `slotRow.Alternatives`) to `[]`, each slot's `dps_delta` to `null`, each band's
 *  `reference_dps_per_point` to `null`, each band's `weights_reason` to `null`, each band's
 *  `set_dps_partial` to `false`, each band's `scale_reference_stat` to `null`, each band's
 *  `haste_scale_factor` to `null`, and each band's `haste_on_items` to `true` (an unknown
 *  band should not silently start showing a claim -- "no item has it" -- an older file never
 *  made) -- all new, optional fields a file published before the lane that added them landed
 *  carries none of, and a Go nil slice/omitted float/omitted string/omitted bool likewise
 *  reach here as `undefined` rather than `[]`/`null`/`null`/`false`/`null`/`null`/`true`. */
export function normaliseBisFile(file: BisFile): BisFile {
  return {
    ...file,
    bands: file.bands.map((band) => ({
      ...band,
      new_at_band: band.new_at_band ?? [],
      coverage: band.coverage ?? {},
      reference_dps_per_point: band.reference_dps_per_point ?? null,
      weights_reason: band.weights_reason ?? null,
      set_dps_partial: band.set_dps_partial ?? false,
      scale_reference_stat: band.scale_reference_stat ?? null,
      haste_scale_factor: band.haste_scale_factor ?? null,
      haste_on_items: band.haste_on_items ?? true,
      slots: band.slots.map((slot) => ({
        ...slot,
        alternatives: slot.alternatives ?? [],
        dps_delta: slot.dps_delta ?? null,
      })),
    })),
  };
}

/** Every spec slug this lane ships a fixture for, read off disk rather than hardcoded, so a
 *  later fixture just works. */
export function fixtureSpecs(): string[] {
  const dir = path.join(WEB_ROOT, 'src/data/fixtures/bis');
  if (!existsSync(dir)) return [];
  return readdirSync(dir)
    .filter((name) => name.endsWith('.json'))
    .map((name) => name.replace(/\.json$/, ''))
    .sort();
}

/** The bands a file covers, ascending, deduplicated across factions. */
export function bandLevels(file: BisFile): number[] {
  return [...new Set(file.bands.map((band) => band.band))].sort((a, b) => a - b);
}

export function bandEntry(file: BisFile, band: number, faction: Faction): BisBand | undefined {
  return file.bands.find((entry) => entry.band === band && entry.faction === faction);
}

export type SlotRow = BisSlot | { slot: Slot; missing: true };

export function isMissingSlot(row: SlotRow): row is { slot: Slot; missing: true } {
  return 'missing' in row;
}

/**
 * A row with nothing to show for it, either way the pipeline can mean that: absent from the
 * band's own `slots` array (`isMissingSlot`) or present but carrying no pick (`{ slot,
 * verified: false }`, `hasKnownSource`'s own doc comment). The page's one empty-slot branch
 * reads this, not `isMissingSlot` alone -- the design brief's "an empty slot says why, not a
 * dash" covers both shapes the real files actually use.
 */
export function isEmptySlotRow(row: SlotRow): boolean {
  return isMissingSlot(row) || !hasKnownSource(row);
}

/** Every one of the planner's 17 slots, in its canonical order, filled from the band where
 *  it has a BiS pick and marked missing where it does not -- the honesty the design doc
 *  asks for: a slot with no known source is a visible row, not a silent gap. */
export function filledSlots(band: BisBand): SlotRow[] {
  const bySlot = new Map(band.slots.map((slot) => [slot.slot, slot]));
  return SLOTS.map((slot): SlotRow => bySlot.get(slot) ?? { slot, missing: true });
}

const isLootKind = (kind: string): kind is LootKind => (LOOT_KINDS as readonly string[]).includes(kind);

function itemsFile<T>(build: string, classSlug: string): T | undefined {
  const file = path.join(REPO_ROOT, 'data/builds', build, 'items', `${classSlug}.json`);
  return existsSync(file) ? readJson<T>(file) : undefined;
}

/**
 * item id -> the fields this lane's item-level column and rarity colouring need, from this
 * build's items/<classSlug>.json. An id neither the real build nor (for the fixture spec)
 * any real item file names simply gets no entry, and the page falls back to plain text the
 * same way GearList does for an unknown id.
 */
export function itemDetails(build: string, classSlug: string): Map<number, ItemDetail> {
  const parsed = itemsFile<{ items: ({ id: number } & ItemDetail)[] }>(build, classSlug);
  if (parsed === undefined) return new Map();
  return new Map(
    parsed.items.map(({ id, name, quality, item_level, required_level, icon, stats }) => [
      id,
      { name, quality, item_level, required_level, icon, stats },
    ]),
  );
}

/**
 * item id -> quality alone, from `itemDetails` -- the shape the rarity-coloured item name
 * GearList.svelte and CandidateRows.svelte both already use (rarityClassFor, planner/
 * items.ts) needs, kept as its own export so this page's existing callers and tests do not
 * have to widen to the full `ItemDetail` just to colour a name.
 */
export function itemQualities(build: string, classSlug: string): Map<number, number> {
  return new Map([...itemDetails(build, classSlug)].map(([id, detail]) => [id, detail.quality]));
}

/**
 * `ItemHover`'s stub `model` prop for one item: the name always comes from the caller (the
 * BiS file's own `item_name` for a picked slot, or a runner-up's name parsed off
 * `swap_note`) rather than `itemDetails`, so the pill still shows the real name even for an
 * id `itemDetails` does not know (a data gap `sim/loot.ts`'s own header names). Quality,
 * item level, required level, icon and stats come from `itemDetails` when it does know the
 * id, and fall back to an uncoloured, level-less pill when it does not.
 */
export function itemHoverModel(
  itemId: number,
  itemName: string,
  details: ReadonlyMap<number, ItemDetail>,
): ItemHoverModel {
  const detail = details.get(itemId);
  return detail === undefined
    ? { name: itemName, quality: 1, itemLevel: 0, requiredLevel: 0 }
    : {
        name: itemName,
        quality: detail.quality,
        itemLevel: detail.item_level,
        requiredLevel: detail.required_level,
        icon: detail.icon,
        stats: detail.stats,
      };
}

/** The picker's own badge word for a source kind (loot.ts's SOURCE_KIND_LABELS), except a
 *  quest -- design 6.1's isolation-by-faction is the whole point of this page, so a quest
 *  row names the band's own faction instead of the generic "Quests" label every kind here
 *  otherwise shares with the Droptimizer's picker. `slot` only needs `source_kind` -- a
 *  `Pick`, not the full `BisSlot`, so an alternative (`item_id`/`source_kind` alone) reads
 *  its own badge word through this exact function too. */
export function sourceBadgeLabel(slot: Pick<BisSlot, 'source_kind'>, faction: Faction): string {
  if (slot.source_kind === 'quest') return bisCopy.questFactionBadge(faction);
  return isLootKind(slot.source_kind) ? SOURCE_KIND_LABELS[slot.source_kind] : slot.source_kind;
}

const EMPTY_LOOT: LootFile & LootQuestsFile = { sources: [], quests: {} };

/**
 * `loot.json`, off disk at build time (see this file's own header) rather than `sim/
 * loot.ts`'s `loadLoot`, which fetches -- Astro's static pages never run in a browser. A
 * build the data lane has not shipped a loot table for yet reads as empty, the same
 * `loadLoot` contract, so a missing file degrades every source cell to its generic badge
 * instead of failing the page.
 */
export function loadLootFile(build: string): LootFile & LootQuestsFile {
  const file = path.join(REPO_ROOT, 'data/builds', build, 'loot.json');
  return existsSync(file) ? readJson<LootFile & LootQuestsFile>(file) : EMPTY_LOOT;
}

/** The band immediately before this one in `bandLevels(file)`'s own order (the contract's
 *  bands are dense, 20..60 step 10, so this is always `band - 10`, but reading it off the
 *  file's own band list rather than hardcoding the step means a file that ever changed its
 *  cadence would still diff correctly). Undefined at the file's first band. */
export function previousBandLevel(file: BisFile, band: number): number | undefined {
  const levels = bandLevels(file);
  const index = levels.indexOf(band);
  return index <= 0 ? undefined : levels[index - 1];
}

/**
 * Every slot's pick this band and the previous band, faction held constant -- the "what
 * changed since <band>" panel's rows, and (via `slot` membership) the set of slots this
 * band's own table marks "new". Undefined at the file's first band -- nothing to have
 * changed since.
 */
export function changedSinceBand(file: BisFile, band: number, faction: Faction): ChangedSlot[] | undefined {
  const previousBand = previousBandLevel(file, band);
  if (previousBand === undefined) return undefined;
  const current = bandEntry(file, band, faction);
  const previous = bandEntry(file, previousBand, faction);
  if (current === undefined) return [];
  const currentBySlot = new Map(current.slots.map((slot) => [slot.slot, slot]));
  const previousBySlot = new Map((previous?.slots ?? []).map((slot) => [slot.slot, slot]));
  const changed: ChangedSlot[] = [];
  for (const slotName of SLOTS) {
    const afterRow = currentBySlot.get(slotName);
    const beforeRow = previousBySlot.get(slotName);
    const after = afterRow !== undefined && hasKnownSource(afterRow) ? afterRow : undefined;
    const before = beforeRow !== undefined && hasKnownSource(beforeRow) ? beforeRow : undefined;
    if (after?.item_id !== before?.item_id) changed.push({ slot: slotName, before, after });
  }
  return changed;
}
