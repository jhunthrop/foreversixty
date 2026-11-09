// web/src/lib/bis/doll-view.ts
// A BiS band as the shared paperdoll's cells (web layout spec 2026-11-05 §4, §5.10, owner decision
// D10): every row two lines, the gain in plain units and against a named basis, the level flag,
// the full source, the short stats and the next best. One tie rule for the whole page
// (`isTie`), so the signed-out and signed-in numbers never disagree without a label. Reads the
// page's compact model (page-model.ts), so the island and the server render share every word.
import type { DollCell, DollFlag, DollNote } from '../ui/paperdoll';
import type { GainFigure, ItemRef } from '../ui/primitives';
import { signedFigure } from '../ui/primitives';
import { gearCopy } from '../gear/copy';
import type { BisBandModel, BisCandidate, BisItem, BisRowModel, GainUnit } from './page-model';
import { filterKindOfSource, filteredBand, pickKinds, type FilteredPick } from './source-filter';
import type { BisByKindEntry, SourceFilterKind } from './types';

/** The paperdoll cell words the site and the apps share (apps/shared/copy/shared.json). */
const DOLL = gearCopy.paperdoll;

export type { GainUnit } from './page-model';

/** A gain under this is the same DPS: the ranker's own noise floor, and the old page's rule. */
export const TIE_DPS = 0.05;

export function isTie(delta: number): boolean {
  return Math.abs(delta) < TIE_DPS;
}

/** Stats on a row's line 2: two keep a desktop row to two lines next to the full source; the detail
 *  lists them all. */
export const ROW_STATS = 2;

/** The game's stat words, shortened the way the mock does (`+8 Agi · +10 Sta`). */
const STAT_ABBR: Readonly<Record<string, string>> = {
  agility: 'Agi',
  stamina: 'Sta',
  strength: 'Str',
  intellect: 'Int',
  spirit: 'Spi',
  attack_power: 'AP',
  ranged_attack_power: 'RAP',
  hit: 'Hit',
  crit: 'Crit',
  spell_power: 'SP',
  spell_hit: 'Spell hit',
  spell_crit: 'Spell crit',
  healing_power: 'Heal',
  mp5: 'Mp5',
  defense: 'Def',
  dodge: 'Dodge',
  parry: 'Parry',
  block: 'Block',
  block_value: 'Block value',
  armor: 'Armor',
  bonus_armor: 'Armor',
};

function statWord(stat: string): string {
  return STAT_ABBR[stat] ?? stat.replace(/_/g, ' ').replace(/^\w/, (c) => c.toUpperCase());
}

function statText(stat: string, value: number): string {
  return `${value > 0 ? '+' : '−'}${Math.abs(value)} ${statWord(stat)}`;
}

/** The item's first `limit` stats in its own order, abbreviated. */
export function shortStats(
  stats: Readonly<Record<string, number>> | undefined,
  limit = 3,
): string | undefined {
  const entries = Object.entries(stats ?? {}).filter(([, value]) => value !== 0);
  if (entries.length === 0) return undefined;
  return entries
    .slice(0, limit)
    .map(([stat, value]) => statText(stat, value))
    .join(' · ');
}

/** The pick minus the worn item, stat by stat, largest changes first (`+5 Agi, −3 Mp5`). */
export function statDelta(
  pick: Readonly<Record<string, number>> | undefined,
  worn: Readonly<Record<string, number>> | undefined,
  limit = 3,
): string | undefined {
  const stats = new Set([...Object.keys(pick ?? {}), ...Object.keys(worn ?? {})]);
  const deltas = [...stats]
    .map((stat) => [stat, (pick?.[stat] ?? 0) - (worn?.[stat] ?? 0)] as const)
    .filter(([, value]) => value !== 0)
    .sort((a, b) => Math.abs(b[1]) - Math.abs(a[1]));
  if (deltas.length === 0) return undefined;
  return deltas
    .slice(0, limit)
    .map(([stat, value]) => statText(stat, value))
    .join(', ');
}

/** The unit as the page says it after a figure. */
export function unitWord(unit: GainUnit): string {
  return unit;
}

/** The signed-in character the doll compares against (R3: what the header already loaded). */
export interface DollCharacter {
  name: string;
  level: number;
  /** Slot -> worn item id. */
  worn: Readonly<Record<string, number | undefined>>;
  /** Slot -> the pick's gain over the worn item (lib/home/upgrades.ts), null when unscored. */
  gains: Readonly<Record<string, number | null | undefined>>;
  /** The worn items' own data, for names, item levels and stat deltas. */
  items?: Readonly<Record<number, BisItem>>;
  /** "Night Elf Hunter": the line under the name in the paperdoll's centre. */
  raceClass?: string;
  /** Battle.net media (referenced by URL, never copied): the head render for the medallion and
   *  the full render the centre stands the model on. */
  avatar?: string;
  render?: string;
}

/** The source choice a band is resolved under. */
export interface SourceChoice {
  /** The kinds that are on. */
  on: readonly SourceFilterKind[];
  /** Below 60 an unchecked Reputation still allows standings up to Friendly (§5.11). */
  repCap?: 'friendly';
}

export interface DollViewDeps {
  items: Readonly<Record<number, BisItem>>;
  unit: GainUnit;
  character?: DollCharacter;
}

export function itemRef(id: number, deps: Pick<DollViewDeps, 'items' | 'character'>): ItemRef {
  const item = deps.items[id] ?? deps.character?.items?.[id];
  return { id, name: item?.name ?? `Item ${id}`, quality: item?.quality ?? 1, icon: item?.icon };
}

function itemOf(id: number, deps: Pick<DollViewDeps, 'items' | 'character'>): BisItem | undefined {
  return deps.items[id] ?? deps.character?.items?.[id];
}

function levelFlag(required: number | undefined, character: DollCharacter | undefined): DollFlag | undefined {
  if (required === undefined || required <= 0) return undefined;
  if (character === undefined) return { text: `Level ${required}`, tone: 'muted' };
  return required <= character.level
    ? { text: 'Usable now', tone: 'kill' }
    : { text: `At ${required}`, tone: 'gold' };
}

/**
 * Signed out: the pick's lead over its best runner-up, with the runner-up named under it (§5.10
 * items 1 and 5, option C). A tied runner-up names the nearest alternative that does not tie.
 */
function leadGain(row: BisRowModel, deps: DollViewDeps): GainFigure {
  const unit = unitWord(deps.unit);
  const next = row.alts[0];
  if (next === undefined) return { value: '\u2014', basis: { prefix: 'Only option' }, tone: 'muted' };
  if (!isTie(next.delta)) {
    return {
      value: signedFigure(-next.delta),
      unit,
      basis: { prefix: DOLL.versus, item: itemRef(next.id, deps) },
      tone: 'strong',
    };
  }
  const different = row.alts
    .filter((alt) => !isTie(alt.delta))
    .reduce<BisCandidate | undefined>(
      (best, alt) => (best === undefined || alt.delta > best.delta ? alt : best),
      undefined,
    );
  if (different === undefined)
    return { value: DOLL.same, basis: { prefix: 'No close alternative' }, tone: 'muted' };
  return {
    value: DOLL.same,

    basis: { prefix: `next: ${signedFigure(different.delta)} vs`, item: itemRef(different.id, deps) },
    tone: 'muted',
  };
}

interface Compare {
  gain: GainFigure;
  badge: 'upgrade' | 'worn' | undefined;
  worn: ItemRef | undefined;
  note: DollNote | undefined;
}

/** Signed in: the pick against the worn item (§5.10 item 8). Undefined when nothing is worn there. */
function compareWith(
  pickId: number,
  slot: string,
  character: DollCharacter,
  deps: DollViewDeps,
): Compare | undefined {
  const wornId = character.worn[slot];
  if (wornId === undefined) return undefined;
  const gain = character.gains[slot];
  const worn = itemRef(wornId, deps);
  if (wornId === pickId || (gain !== null && gain !== undefined && isTie(gain))) {
    return {
      gain: { value: DOLL.worn, basis: { prefix: DOLL.youHaveThis }, tone: 'kill' },
      badge: 'worn',
      worn: undefined,
      note: undefined,
    };
  }
  const wornItem = itemOf(wornId, deps);
  const delta = statDelta(itemOf(pickId, deps)?.stats, wornItem?.stats);
  const facts = [
    wornItem?.ilvl === undefined ? undefined : `item level ${wornItem.ilvl}`,
    shortStats(wornItem?.stats, 2),
  ]
    .filter(Boolean)
    .join(', ');
  const pickLine =
    gain === null || gain === undefined
      ? delta && `The pick: ${delta}`
      : `The pick is ${signedFigure(gain)} ${unitWord(deps.unit)}${delta ? `, ${delta}` : ''}.`;
  return {
    gain:
      gain === null || gain === undefined
        ? { value: '\u2014', basis: { prefix: DOLL.notSimChecked }, tone: 'muted' }
        : {
            value: signedFigure(gain),
            unit: unitWord(deps.unit),
            basis: { prefix: DOLL.versusYour, item: worn },
            tone: 'gold',
          },
    badge: 'upgrade',
    worn,
    note: {
      label: 'You wear',
      lead: 'You wear',
      item: worn,
      tail: [facts && `(${facts}).`, pickLine].filter(Boolean).join(' ') || undefined,
      tone: 'muted',
    },
  };
}

const KIND_WORD: Readonly<Record<string, string>> = {
  dungeon: 'dungeon',
  raid: 'raid',
  world: 'world drop',
  world_drop: 'world drop',
  zone: 'world drop',
  quest: 'quest',
  crafted: 'crafted',
  rep: 'reputation',
  vendor: 'vendor',
  pvp: 'PvP',
};

export function kindWord(kind: string | undefined): string {
  return (kind && KIND_WORD[kind]) ?? 'other';
}

/** The row's own cell, as the unfiltered pick. */
function pickCell(row: BisRowModel, pickId: number, deps: DollViewDeps): DollCell {
  const item = itemOf(pickId, deps);
  const compare = deps.character ? compareWith(pickId, row.slot, deps.character, deps) : undefined;
  return {
    slot: row.slot,
    label: row.label,
    flag: levelFlag(item?.req, deps.character),
    item: itemRef(pickId, deps),
    badge: compare?.badge,
    stats: shortStats(item?.stats, ROW_STATS),
    source: row.sources[pickId]?.line,
    sourceKind: row.sources[pickId]?.kind,
    note: compare?.note,
    gain: compare?.gain ?? leadGain(row, deps),
    worn: compare?.worn,
  };
}

function emptyCell(row: BisRowModel, deps: DollViewDeps, reason: string | undefined): DollCell {
  return {
    slot: row.slot,
    label: row.label,
    item: null,
    empty: reason,
    coveredBy: row.coveredBy === undefined ? undefined : itemRef(row.coveredBy, deps),
  };
}

/** A slot re-picked under the source filter (§5.11): its own gap to the best pick and why. */
function repickedCell(
  row: BisRowModel,
  pick: FilteredPick,
  choiceLabel: string,
  deps: DollViewDeps,
): DollCell {
  const item = itemOf(pick.item_id, deps);
  const best = row.id === undefined ? undefined : itemRef(row.id, deps);
  const bestKind = row.id === undefined ? undefined : row.sources[row.id]?.kind;
  const gap = -pick.dps_delta;
  return {
    slot: row.slot,
    label: row.label,
    flag: levelFlag(item?.req, deps.character),
    item: itemRef(pick.item_id, deps),
    stats: shortStats(item?.stats, ROW_STATS),
    source: row.sources[pick.item_id]?.line,
    sourceKind: row.sources[pick.item_id]?.kind ?? pick.source_kind,
    note:
      best === undefined
        ? undefined
        : {
            label: 'Re-picked',
            lead: `Re-picked (${choiceLabel}). Best overall:`,
            item: best,
            tail: `(${kindWord(bestKind)})${isTie(gap) ? ', same ' + unitWord(deps.unit) : `, ${signedFigure(gap)} ${unitWord(deps.unit)}`}.`,
            tone: 'muted',
          },
    gain: isTie(gap)
      ? { value: DOLL.same, basis: best && { prefix: 'as', item: best }, tone: 'muted' }
      : {
          value: signedFigure(-gap),
          unit: unitWord(deps.unit),
          basis: best && { prefix: 'vs best:', item: best },
          tone: 'wipe',
        },
  };
}

function noOptionCopy(row: BisRowModel, deps: DollViewDeps): string {
  if (row.id === undefined) return row.empty ?? 'No option from the sources you chose.';
  return `No option from the sources you chose. Best overall: ${itemRef(row.id, deps).name} (${kindWord(row.sources[row.id]?.kind)}).`;
}

/**
 * One row as a doll cell. `pick` is the row's answer under the source filter: `undefined` when
 * no filter applies (the published pick), `null` when nothing the chosen sources offer fits.
 */
export function dollCellFor(
  row: BisRowModel,
  deps: DollViewDeps,
  pick?: FilteredPick | null,
  choiceLabel = 'filtered',
): DollCell {
  if (row.id === undefined) return emptyCell(row, deps, row.empty);
  if (pick === undefined || (pick !== null && pick.item_id === row.id)) return pickCell(row, row.id, deps);
  if (pick === null) return emptyCell(row, deps, noOptionCopy(row, deps));
  return repickedCell(row, pick, choiceLabel, deps);
}

const STANDING_RANK: Readonly<Record<string, number>> = {
  neutral: 0,
  friendly: 1,
  honored: 2,
  revered: 3,
  exalted: 4,
};

function aboveCap(standing: string | undefined, cap: 'friendly'): boolean {
  // An unknown standing reads as above the cap: never offered as reachable when the data is silent.
  return (STANDING_RANK[standing ?? ''] ?? Number.POSITIVE_INFINITY) > STANDING_RANK[cap]!;
}

/** The first of an entry and its next that the cap allows. A rep entry with no standing reads
 *  as above the cap; any other kind's entry is capped only by a standing it names. */
function cappedEntry(
  entry: BisByKindEntry | undefined,
  cap: 'friendly',
  rep: boolean,
): BisByKindEntry | undefined {
  for (let node = entry; node !== undefined; node = node.next) {
    if (rep ? !aboveCap(node.standing, cap) : node.standing === undefined || !aboveCap(node.standing, cap)) {
      return node;
    }
  }
  return undefined;
}

/** A slot's filter fields with reputation above the cap taken out (the leveling default). */
function cappedSlot(row: BisRowModel, cap: 'friendly'): BisRowModel['filter'] {
  const filter = row.filter;
  const kinds = pickKinds(filter);
  // Every own kind a standing above the cap gates leaves: a rep item its quartermaster also
  // sells is Honored through its vendor kind too (P2 review, 2026-11-09). A file without
  // kind_standings knows only the rep kind's standing.
  const capped = (kind: SourceFilterKind): boolean => {
    const standing = filter.kind_standings?.[kind] ?? (kind === 'rep' ? filter.standing : undefined);
    return kind === 'rep' ? aboveCap(standing, cap) : standing !== undefined && aboveCap(standing, cap);
  };
  const pickCapped = kinds.some(capped);
  const rest = kinds.filter((kind) => !capped(kind));
  const byKind = { ...filter.by_kind };
  for (const kind of Object.keys(byKind) as SourceFilterKind[]) {
    const kept = cappedEntry(byKind[kind], cap, kind === 'rep');
    if (kept === undefined) delete byKind[kind];
    else byKind[kind] = kept;
  }
  if (!pickCapped) return { ...filter, by_kind: byKind };
  return rest.length > 0
    ? { ...filter, kinds: rest, by_kind: byKind }
    : { ...filter, item_id: 0, by_kind: byKind };
}

/** Every row's answer under a source choice; empty when every kind is on (the published set). */
export function filteredPicks(model: BisBandModel, choice: SourceChoice): Map<string, FilteredPick | null> {
  const allOn = choice.on.length === 8 && choice.repCap === undefined;
  if (allOn) return new Map();
  const allowed = new Set<SourceFilterKind>(choice.on);
  if (choice.repCap !== undefined) allowed.add('rep');
  const slots = model.rows.map((row) =>
    choice.repCap === undefined ? row.filter : cappedSlot(row, choice.repCap),
  );
  return filteredBand(slots, allowed);
}

/** Every row of a band, as cells, under a source choice. */
export function dollCellsFor(
  model: BisBandModel,
  deps: DollViewDeps,
  picks: ReadonlyMap<string, FilteredPick | null> = new Map(),
  choiceLabel?: string,
): DollCell[] {
  return model.rows.map((row) =>
    dollCellFor(row, deps, picks.size === 0 ? undefined : picks.get(row.slot), choiceLabel),
  );
}

/** How many slots a choice re-picks or leaves without an option. */
export function changedCount(model: BisBandModel, picks: ReadonlyMap<string, FilteredPick | null>): number {
  return model.rows.filter((row) => {
    if (row.id === undefined || picks.size === 0) return false;
    const pick = picks.get(row.slot);
    return pick === null || (pick !== undefined && pick.item_id !== row.id);
  }).length;
}

/** Slots with a ranked option from each kind (§5.11's counts). */
export function sourceCounts(model: BisBandModel): Record<SourceFilterKind, number> {
  const counts = { dungeon: 0, raid: 0, world: 0, quest: 0, crafted: 0, rep: 0, vendor: 0, pvp: 0 };
  for (const row of model.rows) {
    if (row.id === undefined) continue;
    const kinds = new Set<SourceFilterKind>(pickKinds(row.filter));
    for (const kind of Object.keys(row.filter.by_kind ?? {}) as SourceFilterKind[]) kinds.add(kind);
    for (const alt of row.alts) {
      const kind = filterKindOfSource(row.sources[alt.id]?.kind ?? '');
      if (kind !== undefined) kinds.add(kind);
    }
    for (const kind of kinds) counts[kind] += 1;
  }
  return counts;
}

/** The summary line (§5.5, §5.10 item 8). */
export function dollSummary(
  cells: readonly DollCell[],
  character: DollCharacter | undefined,
  unit: GainUnit,
): string {
  const filled = cells.filter((cell) => cell.item !== null);
  if (character === undefined) {
    return `${filled.length} of ${cells.length} slots have a pick at this level.`;
  }
  const worn = filled.filter((cell) => cell.badge === 'worn').length;
  const upgrades = filled.filter((cell) => cell.badge === 'upgrade');
  // No slot compared means no worn gear read for this character, not a full set.
  if (worn === 0 && upgrades.length === 0) return `${character.name}'s gear isn't imported.`;
  if (upgrades.length === 0) return `${character.name} wears every pick for this level.`;
  // Only sim-checked gains add up; an upgrade with no measured gain is counted, never priced at 0.
  const measured = upgrades.filter((cell) => typeof character.gains[cell.slot] === 'number');
  const total = measured.reduce((sum, cell) => sum + (character.gains[cell.slot] ?? 0), 0);
  const unchecked = upgrades.length - measured.length;
  const noun = `${upgrades.length} ${upgrades.length === 1 ? 'upgrade' : 'upgrades'}`;
  const worth =
    measured.length === 0
      ? `${noun}, not sim-checked yet.`
      : `${noun}, ${signedFigure(total)} ${unitWord(unit)} available${unchecked > 0 ? ` from ${measured.length} sim-checked` : ''}.`;
  return `${character.name} (level ${character.level}) has ${worn} of ${filled.length} picks. ${worth}`;
}

/** A Where to go group (§4.7): an instance, a profession, a faction, or a kind. */
export interface WhereGroup {
  key: string;
  title: string;
  count: string;
  slots: string[];
}

const GROUP_ORDER: readonly string[] = [
  'dungeon',
  'raid',
  'crafted',
  'rep',
  'quest',
  'vendor',
  'world',
  'pvp',
  'other',
];
const GROUP_KIND: Readonly<Record<string, string>> = {
  dungeon: 'dungeon',
  raid: 'raid',
  crafted: 'crafted',
  rep: 'rep',
  quest: 'quest',
  vendor: 'vendor',
  world: 'world',
  world_drop: 'world',
  zone: 'world',
  pvp: 'pvp',
};
const KIND_TITLE: Readonly<Record<string, string>> = {
  quest: 'Quest',
  vendor: 'Vendor',
  world: 'World / BoE',
  pvp: 'PvP',
  other: 'Other',
};

/** The cells grouped by where to go: instances by pieces (most first), then crafted, reputation,
 *  quest, vendor, world drops and PvP. Empty slots are left out. */
export function whereGroups(model: BisBandModel, cells: readonly DollCell[]): WhereGroup[] {
  const rows = new Map(model.rows.map((row) => [row.slot, row]));
  const groups = new Map<string, { kind: string; title: string; slots: string[] }>();
  for (const cell of cells) {
    if (cell.item === null) continue;
    const info = rows.get(cell.slot)?.sources[cell.item.id];
    const kind = GROUP_KIND[info?.kind ?? ''] ?? 'other';
    const place =
      kind === 'dungeon' || kind === 'raid' || kind === 'crafted' || kind === 'rep' ? info?.place : undefined;
    const key = place === undefined ? kind : `${kind}:${place}`;
    const title =
      place === undefined
        ? (KIND_TITLE[kind] ?? 'Other')
        : kind === 'crafted'
          ? `Crafted · ${place}`
          : kind === 'rep'
            ? `${place} reputation`
            : place;
    const group = groups.get(key) ?? { kind, title, slots: [] };
    group.slots.push(cell.slot);
    groups.set(key, group);
  }
  return [...groups.entries()]
    .sort(([, a], [, b]) => {
      const order = GROUP_ORDER.indexOf(a.kind) - GROUP_ORDER.indexOf(b.kind);
      const instance = (kind: string) => kind === 'dungeon' || kind === 'raid';
      if (instance(a.kind) && instance(b.kind)) return b.slots.length - a.slots.length || order;
      return order || b.slots.length - a.slots.length;
    })
    .map(([key, group]) => ({
      key,
      title: group.title,
      count: `${group.slots.length} ${group.slots.length === 1 ? 'piece' : 'pieces'}`,
      slots: group.slots,
    }));
}
