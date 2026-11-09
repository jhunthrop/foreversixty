// web/src/lib/bis/compare-loader.ts
// Compare mode for the BiS page (web layout spec 2026-11-05 §5.10 item 8): the signed-in
// character's worn gear against the shown band. Loaded on demand only once a session exists:
// it fetches the spec's BiS file and the class item table (both cached, the same files the home
// page's upgrade card reads) and scores each pick over the worn item with the one gain rule
// (lib/home/upgrades.ts), so this page and the home card never disagree.
import type { MeCharacter } from '../account/api';
import { itemsFor } from '../home/upgrades-loader';
import { upgradesFor } from '../home/upgrades';
import type { Item } from '../planner/types';
import { classSlugFromName } from '../report/tree-sizes';
import type { DollCharacter } from './doll-view';
import { fetchBisFile } from './hover';
import type { BisItem } from './page-model';
import { selectBand } from './presets';
import type { BisPresetId, Faction } from './types';

export interface CompareTarget {
  build: string;
  spec: string;
  classSlug: string;
  band: number;
  faction: Faction;
  preset: BisPresetId;
}

function bisItem(item: Item): BisItem {
  const stats = Object.fromEntries(
    Object.entries(item.stats).filter((entry): entry is [string, number] => typeof entry[1] === 'number'),
  );
  return {
    name: item.name,
    quality: item.quality,
    icon: item.icon,
    ilvl: item.item_level,
    req: item.required_level,
    ...(Object.keys(stats).length > 0 ? { stats } : {}),
  };
}

/** Whether the character can be compared on this page: same class, a known level. */
export function comparable(character: MeCharacter | null, classSlug: string): character is MeCharacter {
  return (
    character !== null &&
    character.level !== undefined &&
    classSlugFromName(character.class ?? '') === classSlug
  );
}

/** The doll's character for the shown band, or null when the band or file is unavailable. */
export async function loadDollCharacter(
  character: MeCharacter,
  target: CompareTarget,
): Promise<DollCharacter | null> {
  const [file, items] = await Promise.all([
    fetchBisFile(target.build, target.spec),
    itemsFor(target.classSlug),
  ]);
  if (file === null) return null;
  const band = selectBand(file, target.band, target.faction, target.preset);
  if (band === undefined) return null;
  const result = upgradesFor(character, band, items);
  const gear: Readonly<Record<string, number | undefined>> = character.build?.gear ?? {};
  const gains: Record<string, number | null> = {};
  for (const upgrade of result.upgrades) gains[upgrade.slot] = upgrade.gainDps;
  for (const worn of result.alreadyBis) gains[worn.slot] = 0;
  const wornItems: Record<number, BisItem> = {};
  for (const id of Object.values(gear)) {
    const item = id === undefined ? undefined : items.get(id);
    if (item !== undefined) wornItems[item.id] = bisItem(item);
  }
  const raceClass = [character.race, character.class].filter(Boolean).join(' ');
  return {
    name: character.name,
    level: character.level ?? 0,
    worn: gear,
    gains,
    items: wornItems,
    ...(raceClass === '' ? {} : { raceClass }),
    ...(character.avatar_url ? { avatar: character.avatar_url } : {}),
    ...(character.render_url ? { render: character.render_url } : {}),
  };
}
