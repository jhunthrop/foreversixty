// web/src/lib/guides/guide-tips.ts
// Everything a spec guide can show a tooltip for, built at build time from the client data:
// the class's talents (every rank), the class's spells (`spelltips/<class>.json`, every rank
// of every trainable), the racials, the enchants and consumables the guide picks, and the
// guide's gear. The page marks each name with its key; this bundle is published beside the
// page (`/data/<build>/guide-tips/<class>/<spec>.json`) and fetched the first time a reader
// points at, focuses or taps a name, so the page itself carries none of it.
//
// The same reads give the prose its names (`proseNames`): a name the guide's text mentions is
// marked when the client data knows it as one of these, never from a list typed per guide.
import { readItemTooltipModel } from '../items/lookup-node';
import type { ItemTooltipModel } from '../items/tooltip';
import {
  displayDescription,
  enchantTipKey,
  racialTipKey,
  spellTipKey,
  spellTipModel,
  talentTipKey,
  type GameTipModel,
  type GuideTipBundle,
  type SpellTipRow,
  type TalentTipEntry,
} from '../tooltips/game-tip';
import { buildIconExists, spellIcons } from '../data/build-files';
import { readBuildJson } from './build-files';
import { loadTalentFile, type DecodedTreeView } from './build-tree';
import { guideGearExtras, type GearExtras } from './gear-extras';
import { loadGuideBis, type GuideBis } from './guide-bis';
import type { ProseTarget } from './prose-tips';
import type { RotationItem, RotationLists } from './rotation-lists';
import { racialIconsFor } from './race-table';

interface ClassSpellTips {
  spells: Record<string, SpellTipRow>;
}

interface SharedSpellTips {
  racials: Record<string, SpellTipRow>;
  enchants: Record<string, SpellTipRow>;
  consumables: Record<string, { item_id: number; spell_id: number; use: string }>;
}

interface CatalogItem {
  id: number;
  name: string;
  quality: number;
  item_level: number;
  required_level: number;
}

const RACIAL_SUBTITLE = 'Racial';

function classSpellTips(build: string, classSlug: string): Record<string, SpellTipRow> {
  return readBuildJson<ClassSpellTips>(build, `spelltips/${classSlug}.json`)?.spells ?? {};
}

const sharedCache = new Map<string, SharedSpellTips>();
function sharedSpellTips(build: string): SharedSpellTips {
  const cached = sharedCache.get(build);
  if (cached !== undefined) return cached;
  const read = readBuildJson<SharedSpellTips>(build, 'spelltips/shared.json') ?? {
    racials: {},
    enchants: {},
    consumables: {},
  };
  sharedCache.set(build, read);
  return read;
}

const catalogCache = new Map<string, Map<number, CatalogItem>>();
function itemCatalog(build: string): Map<number, CatalogItem> {
  const cached = catalogCache.get(build);
  if (cached !== undefined) return cached;
  const rows = readBuildJson<CatalogItem[]>(build, 'items.json') ?? [];
  const indexed = new Map(rows.map((row) => [row.id, row]));
  catalogCache.set(build, indexed);
  return indexed;
}

function shownIcon(build: string, icon: string | undefined): string | undefined {
  return icon !== undefined && buildIconExists(build, icon) ? icon : undefined;
}

/** The class's talents, every rank's text, by talent id. */
export function talentTipEntries(build: string, classSlug: string): Record<string, TalentTipEntry> {
  const file = loadTalentFile(build, classSlug);
  if (file === undefined) return {};
  return Object.fromEntries(
    file.trees.flatMap((tree) =>
      tree.talents.map((talent): [string, TalentTipEntry] => {
        const icon = shownIcon(build, talent.icon);
        return [
          String(talent.id),
          {
            name: talent.name,
            ...(icon === undefined ? {} : { icon }),
            maxRank: talent.max_rank,
            ranks: talent.ranks.map((rank) => rank.description),
          },
        ];
      }),
    ),
  );
}

/** Every rank of every class spell, as its tooltip. */
function spellTips(build: string, classSlug: string): Record<string, GameTipModel> {
  const icons = spellIcons(build);
  return Object.fromEntries(
    Object.entries(classSpellTips(build, classSlug)).map(([id, row]) => [
      spellTipKey(Number(id)),
      spellTipModel(row, shownIcon(build, row.icon ?? icons[id])),
    ]),
  );
}

function racialTips(build: string): Record<string, GameTipModel> {
  const icons = racialIconsFor(build);
  const key = (name: string): string => name.toLowerCase().replace(/[^a-z0-9]+/g, '');
  return Object.fromEntries(
    Object.entries(sharedSpellTips(build).racials).map(([name, row]) => [
      racialTipKey(name),
      spellTipModel(row, icons.get(key(name)), RACIAL_SUBTITLE),
    ]),
  );
}

/** The enchant's own spell tooltip, by enchant id; undefined when the client has none. */
export function enchantTip(
  build: string,
  enchantId: number,
  icon: string | undefined,
): GameTipModel | undefined {
  const row = sharedSpellTips(build).enchants[String(enchantId)];
  return row === undefined ? undefined : spellTipModel(row, icon ?? shownIcon(build, row.icon));
}

/** A consumable's item, as the shared item tooltip shows it: its name in its quality, the
 *  level it needs and its "Use:" line. `undefined` when the client names no one item for it. */
export function consumableItemModel(
  build: string,
  consumableId: string,
  icon: string | undefined,
): ItemTooltipModel | undefined {
  const tip = sharedSpellTips(build).consumables[consumableId];
  if (tip === undefined) return undefined;
  const item = itemCatalog(build).get(tip.item_id);
  if (item === undefined) return undefined;
  return {
    id: item.id,
    name: item.name,
    quality: item.quality,
    icon: icon ?? '',
    slotLabel: '',
    typeLabel: undefined,
    bindLabel: null,
    itemLevel: 0,
    requiredLevel: item.required_level,
    armor: null,
    weapon: null,
    stats: [],
    effectText: `Use: ${displayDescription(tip.use)}`,
    setName: null,
    sourceLines: [],
    unique: false,
    clientUnconfirmed: false,
  };
}

/** The consumable's item id, when its tooltip can be shown. */
export function consumableItemId(build: string, consumableId: string): number | undefined {
  return sharedSpellTips(build).consumables[consumableId]?.item_id;
}

export interface GuideTipInputs {
  build: string;
  classSlug: string;
  bis: GuideBis | undefined;
  extras: GearExtras | undefined;
}

function bisItemModels(
  build: string,
  classSlug: string,
  bis: GuideBis | undefined,
): Record<string, ItemTooltipModel> {
  const ids = new Set(
    (bis?.panels ?? []).flatMap((panel) =>
      panel.rows.flatMap((row) => (row.itemId === undefined ? [] : [row.itemId])),
    ),
  );
  return Object.fromEntries(
    [...ids].flatMap((id) => {
      const model = readItemTooltipModel(build, classSlug, id);
      return model === null ? [] : [[String(id), model]];
    }),
  );
}

function extrasTips(
  build: string,
  extras: GearExtras | undefined,
): {
  tips: Record<string, GameTipModel>;
  items: Record<string, ItemTooltipModel>;
} {
  const tips: Record<string, GameTipModel> = {};
  const items: Record<string, ItemTooltipModel> = {};
  for (const row of extras?.enchants ?? []) {
    const tip = enchantTip(build, row.id, row.icon);
    if (tip !== undefined) tips[enchantTipKey(row.id)] = tip;
  }
  for (const row of extras?.consumables ?? []) {
    const model = consumableItemModel(build, row.id, row.icon);
    if (model !== undefined) items[String(model.id)] = model;
  }
  return { tips, items };
}

/** The guide's whole tooltip bundle. */
export function guideTipBundle(inputs: GuideTipInputs): GuideTipBundle {
  const { build, classSlug } = inputs;
  const extras = extrasTips(build, inputs.extras);
  return {
    tips: { ...spellTips(build, classSlug), ...racialTips(build), ...extras.tips },
    talents: talentTipEntries(build, classSlug),
    items: { ...bisItemModels(build, classSlug, inputs.bis), ...extras.items },
  };
}

export interface SpecGuideData {
  bis: GuideBis | undefined;
  extras: GearExtras | undefined;
  bundle: GuideTipBundle;
}

const specCache = new Map<string, SpecGuideData>();

/** The guide's gear tables and its tooltip bundle, read once per spec: the page renders the
 *  tables and marks their names, and the bundle endpoint publishes the same reads. */
export function specGuideData(
  build: string,
  classSlug: string,
  spec: string,
  recommendedRaces: readonly string[],
): SpecGuideData {
  const cacheKey = `${build}::${spec}::${recommendedRaces.join(',')}`;
  const cached = specCache.get(cacheKey);
  if (cached !== undefined) return cached;
  const bis = loadGuideBis(build, classSlug, spec, recommendedRaces);
  const extras = guideGearExtras(build, classSlug, spec, recommendedRaces, bis);
  const data = { bis, extras, bundle: guideTipBundle({ build, classSlug, bis, extras }) };
  specCache.set(cacheKey, data);
  return data;
}

/** The tooltip a rotation or cooldown line opens: its spell's, else (a racial cooldown, which
 *  is not a class spell) the racial's by name; undefined when the bundle has neither. */
export function lineTipKey(
  bundle: GuideTipBundle,
  spellId: number | undefined,
  name: string,
): string | undefined {
  if (spellId !== undefined && bundle.tips[spellTipKey(spellId)] !== undefined) return spellTipKey(spellId);
  return bundle.tips[racialTipKey(name)] !== undefined ? racialTipKey(name) : undefined;
}

/** A racial's tooltip key by the name the race data uses, which can spell it apart from the
 *  client ("Read Ley Line", "Read Leyline"); undefined when the bundle has no such racial. */
export function racialKeyFor(bundle: GuideTipBundle, name: string): string | undefined {
  const fold = (text: string): string => text.toLowerCase().replace(/[^a-z0-9]+/g, '');
  const wanted = fold(name);
  const prefix = racialTipKey('');
  return Object.keys(bundle.tips).find(
    (key) => key.startsWith(prefix) && fold(key.slice(prefix.length)) === wanted,
  );
}

/** The lists with each line's tooltip key. */
export function withTipKeys(lists: RotationLists, bundle: GuideTipBundle): RotationLists {
  const keyed = (items: readonly RotationItem[]): RotationItem[] =>
    items.map((item) => {
      const tipKey = lineTipKey(bundle, item.spellId, item.name);
      return tipKey === undefined ? item : { ...item, tipKey };
    });
  return { opener: keyed(lists.opener), priority: keyed(lists.priority), cooldowns: keyed(lists.cooldowns) };
}

/** Where the page fetches its bundle from. */
export function guideTipsUrl(build: string, classSlug: string, specSlug: string): string {
  return `/data/${build}/guide-tips/${classSlug}/${specSlug}.json`;
}

/** The highest rank of each class spell a player casts, by name: the spell a guide means
 *  when its prose names the ability. Passive skills (Parry, Dual Wield) are not linked. */
function castSpellsByName(build: string, classSlug: string): Map<string, number> {
  const best = new Map<string, { id: number; rank: number }>();
  for (const [id, row] of Object.entries(classSpellTips(build, classSlug))) {
    if (row.passive || row.description === '') continue;
    const current = best.get(row.name);
    if (current === undefined || row.rank > current.rank)
      best.set(row.name, { id: Number(id), rank: row.rank });
  }
  return new Map([...best].map(([name, { id }]) => [name, id]));
}

/**
 * The names the guide's prose may mark, each with what it opens: the class's cast spells (at
 * their top rank), its talents (at the guide build's rank), every racial, and the items and
 * consumables the guide itself shows. A name that is both a talent and the spell it teaches
 * opens the spell, the way an action button's tooltip does.
 */
export function proseNames(
  build: string,
  classSlug: string,
  trees: readonly DecodedTreeView[],
  bundle: GuideTipBundle,
): Map<string, ProseTarget> {
  const names = new Map<string, ProseTarget>();
  for (const [name] of Object.entries(sharedSpellTips(build).racials))
    names.set(name, { key: racialTipKey(name) });
  const ranks = new Map(
    trees.flatMap((tree) => tree.talents.map((talent) => [talent.id, talent.rank] as const)),
  );
  for (const [id, entry] of Object.entries(bundle.talents)) {
    names.set(entry.name, { key: talentTipKey(Number(id), ranks.get(Number(id)) ?? 0) });
  }
  for (const [name, id] of castSpellsByName(build, classSlug)) names.set(name, { key: spellTipKey(id) });
  for (const model of Object.values(bundle.items)) {
    names.set(model.name, { itemId: model.id, quality: model.quality });
  }
  return names;
}
