// web/src/lib/home/upgrades-loader.ts
// The browser-side half of the "Best in slot"/"Talents"/"Your upgrades" comparison: finds a
// character's spec key, fetches its BiS file and item table (both already-published, already
// runtime-fetchable assets -- `lib/bis/hover.ts`'s own `fetchBisFile` for the planner's hover
// popover, `lib/planner/load.ts`'s `loadItems` for the planner island itself), and resolves
// the one band entry that applies to this character's level and faction. `lib/home/
// upgrades.ts` and `lib/home/talent-delta.ts` stay pure and synchronous; every network/async
// concern lives here instead, once, so `HomeHeroCards.svelte`, `HomeUpgradesPanel.svelte` and
// `HomeSwitchCharacterPanel.svelte` share the identical fetch-and-cache path rather than each
// growing its own.
import activeBuild from '../../data/active-build.json';
import type { MeCharacter } from '../account/api';
import { bandEntryFor, bandForLevel, fetchBisFile } from '../bis/hover';
import type { BisBand, BisFile } from '../bis/types';
import { loadItems } from '../planner/load';
import type { Item } from '../planner/types';
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
 *  same class (the hero, every row in Switch character) shares the one promise. Resolves to
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
