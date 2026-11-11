// web/src/lib/sim/action-visuals.ts
// What an action key looks like: its icon and its rank (spec 2026-11-04-sim-results 0: icons from
// spellicons.json, then the class's talent file, then simbuffs.json for a buff or consumable of
// the same name; a white swing shows the weapon it swings). A key with no icon resolves to null,
// and the page draws no box for it (spec 15, R1). Every file is optional: a build without one
// gives plain names, never a broken row.
import activeBuild from '../../data/active-build.json';
import { dataUrl, loadItems, loadOptional, loadTalents } from '../planner/load';
import { parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { spellTipKey } from '../tooltips/game-tip';
import type { SimResult } from './types';
import { iconUrl } from '../ui/assets';

export interface ActionVisual {
  /** The icon's url; null when the build carries none for this key. */
  icon: string | null;
  /** The spell's client rank; 0 when it has none. */
  rank: number;
  /** The shared tooltip's key; null for a swing, an item or anything else the host cannot open. */
  tipKey: string | null;
}

export type ActionVisuals = (key: string) => ActionVisual;

/** Before the files arrive: no icon, no rank, and the host's key where one can be read off the action. */
export const pendingVisual = (key: string): ActionVisual => ({ icon: null, rank: 0, tipKey: tipKeyOf(key) });

function tipKeyOf(key: string): string | null {
  const parsed = parseActionKey(key);
  return parsed?.kind === 'spell' ? spellTipKey(parsed.id) : null;
}

interface SimBuffFile {
  entries?: Record<string, { name?: string; icon?: string }>;
}

interface SpellConstFile {
  spells?: Record<string, { rank?: number }>;
}

/** The files a visual reads, as plain lookups, so the rule can be tested without a fetch. */
export interface VisualTables {
  spellIcons: Readonly<Record<string, string>>;
  /** Spell id to icon, from the class's talent file. */
  talentIcons: ReadonlyMap<number, string>;
  /** A buff or consumable's display name to icon. */
  buffIcons: ReadonlyMap<string, string>;
  ranks: Readonly<Record<string, { rank?: number }>>;
  /** The hand tag (1 main, 2 off, 3 extra) to the weapon's icon. */
  weaponIcons: ReadonlyMap<number, string>;
}

function iconStem(key: string, tables: VisualTables, names: ActionNames | null): string | undefined {
  const parsed = parseActionKey(key);
  if (parsed === null) return undefined;
  if (parsed.kind === 'other')
    return parsed.label === 'attack' ? tables.weaponIcons.get(parsed.tag) : undefined;
  const byName = tables.buffIcons.get(resolveActionName(`${parsed.kind}:${parsed.label}`, names));
  if (parsed.kind !== 'spell') return byName;
  const id = String(parsed.id);
  return tables.spellIcons[id] ?? tables.talentIcons.get(parsed.id) ?? byName;
}

export function actionVisuals(
  tables: VisualTables,
  names: ActionNames | null,
  build: string = activeBuild.build,
): ActionVisuals {
  return (key) => {
    const stem = iconStem(key, tables, names);
    const parsed = parseActionKey(key);
    return {
      icon: iconUrl(build, stem),
      rank: parsed?.kind === 'spell' ? (tables.ranks[String(parsed.id)]?.rank ?? 0) : 0,
      tipKey: tipKeyOf(key),
    };
  };
}

const HAND_SLOTS: readonly (readonly [number, string])[] = [
  [1, 'main_hand'],
  [2, 'off_hand'],
  [3, 'main_hand'],
];

/** One file's failure leaves its lookups empty: an icon is never worth failing the card over. */
async function orEmpty<T>(load: Promise<T>, empty: T): Promise<T> {
  try {
    return await load;
  } catch {
    return empty;
  }
}

/** The build's tables for this result's class. */
export async function loadVisualTables(
  /** Only the request is read (the class and the worn weapons), so a page before its run can ask. */
  result: Pick<SimResult, 'request'>,
  build: string = activeBuild.build,
): Promise<VisualTables> {
  const classSlug = result.request.character.class;
  const [spellIcons, buffs, spellConst, talents, items] = await Promise.all([
    orEmpty(loadOptional<Record<string, string>>(dataUrl(build, 'spellicons.json'), {}), {}),
    orEmpty(loadOptional<SimBuffFile>(dataUrl(build, 'simbuffs.json'), {}), {} as SimBuffFile),
    orEmpty(
      loadOptional<SpellConstFile>(dataUrl(build, `spellconst/${classSlug}.json`), {}),
      {} as SpellConstFile,
    ),
    orEmpty(
      loadTalents(build, classSlug).then((file) => file.trees),
      [],
    ),
    orEmpty(
      loadItems(build, classSlug).then((file) => file.items),
      [],
    ),
  ]);

  const buffIcons = new Map<string, string>();
  for (const entry of Object.values(buffs.entries ?? {})) {
    if (entry.name !== undefined && entry.icon !== undefined) buffIcons.set(entry.name, entry.icon);
  }
  const talentIcons = new Map<number, string>();
  for (const tree of talents) {
    for (const talent of tree.talents) {
      for (const rank of talent.ranks) talentIcons.set(rank.spell_id, talent.icon);
    }
  }
  const itemIcons = new Map(items.map((item) => [item.id, item.icon]));
  const weaponIcons = new Map<number, string>();
  for (const [tag, slot] of HAND_SLOTS) {
    const worn = result.request.character.gear.find((row) => row.slot === slot);
    const icon = worn === undefined ? undefined : itemIcons.get(worn.item_id);
    if (icon !== undefined) weaponIcons.set(tag, icon);
  }
  return { spellIcons, talentIcons, buffIcons, ranks: spellConst.spells ?? {}, weaponIcons };
}
