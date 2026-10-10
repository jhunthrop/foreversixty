// web/src/lib/home/upgrades-loader.ts
// The browser-side half of the "Best in slot"/"Talents"/"Your upgrades" comparison: finds a
// character's spec key, fetches its BiS file and item table (both already-published, already
// runtime-fetchable assets -- `lib/bis/hover.ts`'s own `fetchBisFile` for the planner's hover
// popover, `lib/planner/load.ts`'s `loadItems` for the planner island itself), and resolves
// the one band entry that applies to this character's level and faction. `lib/home/
// upgrades.ts` and `lib/home/talent-delta.ts` stay pure and synchronous; every network/async
// concern lives here instead, once, so `HomeHeroCards.svelte` and `HomeUpgradesPanel.svelte`
// share the identical fetch-and-cache path rather than each growing its own.
import activeBuild from '../../data/active-build.json';
import type { MeCharacter } from '../account/api';
import { bandEntryFor, bandForLevel, fetchBisFile } from '../bis/hover';
import type { BisBand, BisFile } from '../bis/types';
import { query } from '../data/query';
import { loadItems } from '../planner/load';
import type { Item } from '../planner/types';
import { upgradesFor, type UpgradesResult } from './upgrades';
import { classSlugFromName } from '../report/tree-sizes';
import { SPECS } from '../sim/specs';

/** `character.class` ("Hunter") + `character.spec` ("Marksmanship") -> the canonical spec
 *  key (`"hunter-marksmanship"`) `SPECS`/BiS files key by. `undefined` when either half is
 *  missing (an addon export with no learned spec yet) or names a class/spec pair `SPECS`
 *  does not recognise -- never guessed at. */
export function specKeyForCharacter(character: Pick<MeCharacter, 'class' | 'spec'>): string | undefined {
  if (character.class === undefined || character.spec === undefined) return undefined;
  const classSlug = classSlugFromName(character.class);
  return SPECS.find((row) => row.class_slug === classSlug && row.name === character.spec)?.spec;
}

/** One item-file fetch per class slug for the life of the page -- every character of the
 *  same class shares the one promise. Resolves to
 *  an empty map, never a rejected promise, when this build ships no item file for the class
 *  (`loadItems` throws `DataLoadError` on a 404 the same way `loadTalents` does) -- an empty
 *  map degrades every item lookup to "unknown" (`upgradesFor`'s own `wornUnknown`/no-icon
 *  fallback already handles that), which is the honest answer rather than an unhandled
 *  rejection breaking every card that awaits this. */
const itemsCache = new Map<string, Promise<Map<number, Item>>>();

function itemsFor(classSlug: string): Promise<Map<number, Item>> {
  const cached = itemsCache.get(classSlug);
  if (cached !== undefined) return cached;
  const promise = loadItems(activeBuild.build, classSlug)
    .then((file) => new Map(file.items.map((item) => [item.id, item])))
    .catch(() => new Map<number, Item>());
  itemsCache.set(classSlug, promise);
  return promise;
}

export interface BisContext {
  specKey: string;
  bisFile: BisFile;
  band: BisBand;
  items: ReadonlyMap<number, Item>;
}

/**
 * Everything a character's own band comparison needs, fetched once and shared by every
 * card/panel that reads it. `null` when any one fact it depends on is missing or this
 * build/spec has nothing published yet -- every honest reason this comparison cannot run:
 * no spec (`specKeyForCharacter`), no level or faction to pick a band with, no BiS file
 * published for the spec, or no band entry for this exact level/faction pair in that file.
 * Callers read `null` as "show the not-available state," never retry it into a guess.
 */
export async function loadBisContextFor(
  character: Pick<MeCharacter, 'class' | 'spec' | 'level' | 'faction'>,
): Promise<BisContext | null> {
  const specKey = specKeyForCharacter(character);
  if (specKey === undefined || character.level === undefined || character.faction === undefined) return null;
  const classSlug = classSlugFromName(character.class ?? '');
  const [bisFile, items] = await Promise.all([fetchBisFile(activeBuild.build, specKey), itemsFor(classSlug)]);
  if (bisFile === null) return null;
  const band = bandEntryFor(bisFile, bandForLevel(character.level), character.faction);
  if (band === undefined) return null;
  return { specKey, bisFile, band, items };
}

/** The home page's own answer for one character, small enough to persist: the comparison
 *  result plus only the item rows it names (worn items, picks, already-BiS items), so the
 *  panel paints from localStorage on a repeat visit instead of
 *  waiting on a 4-6MB class item table to download and parse first (owner 2026-10-04, "are we
 *  caching the upgrades?" -- we were not). `band` is the band number; `specKey` and it are
 *  all the panel's "Full list for …" link needs. */
export interface CachedUpgrades {
  specKey: string;
  band: number;
  result: UpgradesResult;
  items: Item[];
}

/** One hour, the same shelf life as the data files it is computed from (public/_headers). */
export const UPGRADES_TTL_MS = 60 * 60 * 1000;

/** The cache key names everything the answer depends on: data build, the character, their
 *  level and faction (band choice) and the exact worn set, so a new addon export or a level-up
 *  is a different key, never a stale hit. */
export function upgradesCacheKey(
  character: Pick<MeCharacter, 'key' | 'level' | 'faction' | 'build'>,
  build: string = activeBuild.build,
): string {
  const gear = character.build?.gear ?? {};
  const worn = Object.keys(gear)
    .sort()
    .map((slot) => `${slot}=${gear[slot]}`)
    .join(',');
  return `fs.upgrades:${build}:${character.key}:${character.level ?? ''}:${character.faction ?? ''}:${worn}`;
}

function itemsNamedBy(result: UpgradesResult, items: ReadonlyMap<number, Item>): Item[] {
  const ids = new Set<number>();
  for (const upgrade of result.upgrades) {
    if (upgrade.wornItemId !== undefined) ids.add(upgrade.wornItemId);
    ids.add(upgrade.pick.item_id);
  }
  for (const entry of result.alreadyBis) ids.add(entry.itemId);
  return [...ids].flatMap((id) => {
    const item = items.get(id);
    return item === undefined ? [] : [item];
  });
}

/**
 * `loadBisContextFor` + `upgradesFor` for one character, remembered per visitor (private
 * scope: cleared with the session) for an hour. A fresh persisted answer resolves instantly
 * and revalidates in the background through `query`'s stale-while-revalidate; `null` means
 * the same "not available" `loadBisContextFor` means, and is cached too so a character with
 * no published list does not re-fetch on every visit.
 */
export function cachedUpgradesFor(
  character: Pick<MeCharacter, 'key' | 'class' | 'spec' | 'level' | 'faction' | 'build'>,
): Promise<CachedUpgrades | null> {
  return query<CachedUpgrades | null>(
    upgradesCacheKey(character),
    async () => {
      const ctx = await loadBisContextFor(character);
      if (ctx === null) return null;
      const result = upgradesFor(character, ctx.band, ctx.items);
      return { specKey: ctx.specKey, band: ctx.band.band, result, items: itemsNamedBy(result, ctx.items) };
    },
    { scope: 'private', ttlMs: UPGRADES_TTL_MS },
  );
}
