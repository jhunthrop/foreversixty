// web/src/lib/sim/buff-names.ts
// A display name and an icon for each id in the engine's vocabulary, from the build.
//
// The same shape and the same rule as action-names.ts: the build publishes a table, an id
// the table does not carry keeps a legible form of itself, and a build with no table at
// all renders humanised ids -- which is honest and is not worth an error banner on a page
// whose numbers are all correct. The table is `data/builds/<build>/simbuffs.json`
// (contract 10.4), owned by the data lane; until a build publishes one, every row falls
// back.
import { dataUrl, fetchJson } from '../planner/load';
import { humanise } from './humanise';

export interface BuffNames {
  /** Engine id to its display name and the build's icon key (no extension). */
  entries: Record<string, { name: string; icon: string }>;
}

export const EMPTY_BUFF_NAMES: BuffNames = { entries: {} };

export function loadBuffNames(build: string): Promise<BuffNames> {
  return fetchJson<BuffNames>(dataUrl(build, 'simbuffs.json'));
}

/**
 * An id's name. A qualified id ("main_hand_imbue:shadow_oil") humanises both halves and
 * keeps the qualifier, because the qualifier is the part that says which hand. A client
 * item id ("item:13452") is left exactly as it is: there is no word in it to raise, and
 * inventing "Item: 13452" would read as a name rather than as the raw id it is.
 */
export function buffLabel(id: string, names: BuffNames | null): string {
  const known = names?.entries[id]?.name;
  if (known !== undefined && known !== '') return known;
  const colon = id.indexOf(':');
  if (colon === -1) return humanise(id);
  const qualifier = id.slice(0, colon);
  const rest = id.slice(colon + 1);
  if (qualifier === 'item') return id;
  return `${humanise(qualifier)}: ${humanise(rest)}`;
}

/** The build's icon for an id, or null when there is none to draw. */
export function buffIcon(build: string, id: string, names: BuffNames | null): string | null {
  const icon = names?.entries[id]?.icon;
  return icon === undefined || icon === '' ? null : dataUrl(build, `icons/${icon}.webp`);
}
