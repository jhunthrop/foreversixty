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
import { LOOT_KINDS, SOURCE_KIND_LABELS, type LootKind } from '../sim/loot';
import { bisCopy } from './copy';
import type { BisBand, BisFile, BisSlot, Faction, SpecCatalogEntry } from './types';

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
  if (existsSync(real)) return readJson<BisFile>(real);
  const fixture = fixtureBisPath(spec);
  if (existsSync(fixture)) return readJson<BisFile>(fixture);
  return null;
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

/** Every one of the planner's 17 slots, in its canonical order, filled from the band where
 *  it has a BiS pick and marked missing where it does not -- the honesty the design doc
 *  asks for: a slot with no known source is a visible row, not a silent gap. */
export function filledSlots(band: BisBand): SlotRow[] {
  const bySlot = new Map(band.slots.map((slot) => [slot.slot, slot]));
  return SLOTS.map((slot): SlotRow => bySlot.get(slot) ?? { slot, missing: true });
}

const isLootKind = (kind: string): kind is LootKind => (LOOT_KINDS as readonly string[]).includes(kind);

/**
 * item id -> quality, from this build's items/<classSlug>.json, for the rarity-coloured
 * item name GearList.svelte and CandidateRows.svelte both already use (rarityClassFor,
 * planner/items.ts). The BiS contract's own slot rows carry no quality field, so this joins
 * against the same per-class item file the rest of the site already ships; an id neither
 * the real build nor (for the fixture spec) any real item file names simply gets no entry,
 * and the page falls back to plain text the same way GearList does for an unknown id.
 */
export function itemQualities(build: string, classSlug: string): Map<number, number> {
  const file = path.join(REPO_ROOT, 'data/builds', build, 'items', `${classSlug}.json`);
  if (!existsSync(file)) return new Map();
  const { items } = readJson<{ items: { id: number; quality: number }[] }>(file);
  return new Map(items.map((item) => [item.id, item.quality]));
}

/** The picker's own badge word for a source kind (loot.ts's SOURCE_KIND_LABELS), except a
 *  quest -- design 6.1's isolation-by-faction is the whole point of this page, so a quest
 *  row names the band's own faction instead of the generic "Quests" label every kind here
 *  otherwise shares with the Droptimizer's picker. */
export function sourceBadgeLabel(slot: BisSlot, faction: Faction): string {
  if (slot.source_kind === 'quest') return bisCopy.questFactionBadge(faction);
  return isLootKind(slot.source_kind) ? SOURCE_KIND_LABELS[slot.source_kind] : slot.source_kind;
}
