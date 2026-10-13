// web/src/lib/gear/qualities.ts
// What the Gear section needs from the build's items/<class>.json (Task L8.6, fix round 1
// rulings 2, 9 and 10): each item's quality (the name's colour and the icon's border), its icon,
// and whether it is two-handed. The API's gear panel carries ids and names only. The file is the
// one the planner, the home upgrades panel and the item tooltip read (planner/load.ts's
// loadItems). One read per class while it succeeds; a failed read is not remembered, so the next
// open of the page retries it (ruling 2). Until the read settles the section draws no names, so
// nothing is ever painted white and recoloured.
import { loadItems } from '../planner/load';
import { rarityClassFor } from '../planner/items';

export interface GearItemFacts {
  /** Client value: 0 poor ... 5 legendary. */
  quality: number;
  /** The icon's file name, or null for the pipeline's questionmark placeholder (the /bis rule:
   *  never a red "?"). */
  icon: string | null;
  twoHand: boolean;
}

export type GearItemTable = ReadonlyMap<number, GearItemFacts>;

/** The pipeline's placeholder for an id it has no real icon for (bis/GearRow.astro). */
const PLACEHOLDER_ICON = 'inv_misc_questionmark';

const cache = new Map<string, Promise<GearItemTable>>();

/** The class's item facts, or null when the file did not load (the names then take the body
 *  colour, and the failure is forgotten so the next call reads the file again). */
export async function itemFacts(build: string, classSlug: string): Promise<GearItemTable | null> {
  const key = `${build}/${classSlug}`;
  let promise = cache.get(key);
  if (promise === undefined) {
    promise = loadItems(build, classSlug).then(
      (file): GearItemTable =>
        new Map(
          file.items.map((item) => [
            item.id,
            {
              quality: item.quality,
              icon: item.icon === undefined || item.icon === PLACEHOLDER_ICON ? null : item.icon,
              twoHand: item.two_hand === true,
            },
          ]),
        ),
    );
    cache.set(key, promise);
  }
  const settled = promise;
  try {
    return await settled;
  } catch {
    if (cache.get(key) === settled) cache.delete(key);
    return null;
  }
}

/** A name's colour class: its quality's, or the body colour for an id the table lacks or when
 *  the file failed (never the common white by default). */
export function nameClass(table: GearItemTable | null, id: number): string {
  const facts = table?.get(id);
  return facts === undefined ? 'text-text' : rarityClassFor(facts.quality);
}

/** The item's icon file name, or null (no table, an id it lacks, or the placeholder). */
export function iconOf(table: GearItemTable | null, id: number): string | null {
  return table?.get(id)?.icon ?? null;
}

/** Tests only: forget every read. */
export function resetItemFacts(): void {
  cache.clear();
}
