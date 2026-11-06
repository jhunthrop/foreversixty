// web/src/lib/bis/doll-view.ts
// A BiS band as the shared paperdoll's cells (web layout spec 2026-11-05 §4, §5.10, owner decision
// D10): every row two lines, the gain in plain units and against a named basis, the level flag,
// the full source, the short stats and the next best. One tie rule for the whole page
// (`isTie`), so the signed-out and signed-in numbers never disagree without a label.
import type { DollCell, DollFlag, DollNote } from '../ui/paperdoll';
import type { GainFigure, ItemRef } from '../ui/primitives';
import { signedFigure } from '../ui/primitives';
import type { AlternativeView, BandInfo, RowView } from './panel-view';
import type { ItemDetail } from './types';

/** A gain under this is the same DPS: the ranker's own noise floor, and the old page's rule. */
export const TIE_DPS = 0.05;

export function isTie(delta: number): boolean {
  return Math.abs(delta) < TIE_DPS;
}

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
  healing_power: 'Heal',
  mp5: 'Mp5',
  defense: 'Def',
  armor: 'Armor',
};

function statWord(stat: string): string {
  return STAT_ABBR[stat] ?? stat.replace(/_/g, ' ').replace(/^\w/, (c) => c.toUpperCase());
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
    .map(([stat, value]) => `${value > 0 ? '+' : '−'}${Math.abs(value)} ${statWord(stat)}`)
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
    .map(([stat, value]) => `${value > 0 ? '+' : '−'}${Math.abs(value)} ${statWord(stat)}`)
    .join(', ');
}

/** The unit a band's figures are in, as the gain caption says it. */
export type GainUnit = 'DPS' | 'HPS' | 'score';

function unitWord(unit: GainUnit): string {
  return unit === 'score' ? 'points' : unit;
}

/** The signed-in character the doll compares against (R3: what the header already loaded). */
export interface DollCharacter {
  name: string;
  level: number;
  /** Slot -> worn item id. */
  worn: Readonly<Record<string, number | undefined>>;
  /** Slot -> the pick's gain over the worn item (lib/home/upgrades.ts), null when unscored. */
  gains: Readonly<Record<string, number | null | undefined>>;
}

export interface DollViewDeps {
  details: ReadonlyMap<number, ItemDetail>;
  unit: GainUnit;
  character?: DollCharacter;
}

function itemRef(
  id: number,
  name: string,
  details: ReadonlyMap<number, ItemDetail>,
  quality?: number,
): ItemRef {
  const detail = details.get(id);
  return { id, name: detail?.name ?? name, quality: detail?.quality ?? quality ?? 1, icon: detail?.icon };
}

function levelFlag(required: number | undefined, character: DollCharacter | undefined): DollFlag | undefined {
  if (required === undefined || required <= 0) return undefined;
  if (character === undefined) return { text: `Level ${required}`, tone: 'muted' };
  return required <= character.level
    ? { text: 'Usable now', tone: 'kill' }
    : { text: `At ${required}`, tone: 'gold' };
}

/** Signed out: the pick's lead over its best runner-up (§5.10 item 1). */
function leadGain(next: AlternativeView | undefined, unit: GainUnit): GainFigure {
  if (next === undefined) return { value: '—', caption: 'Only option', tone: 'muted' };
  const lead = -next.dpsDelta;
  if (isTie(lead)) return { value: '=', caption: 'Same as next best', tone: 'muted' };
  return { value: signedFigure(lead), caption: `${unitWord(unit)} vs next best`, tone: 'strong' };
}

function nextBestNote(next: AlternativeView | undefined, deps: DollViewDeps): DollNote | undefined {
  if (next === undefined) return undefined;
  const tie = isTie(next.dpsDelta);
  return {
    lead: 'Next best:',
    item: itemRef(next.itemId, next.itemName, deps.details, next.model?.quality),
    tail: tie ? `(same ${unitWord(deps.unit)})` : `${signedFigure(next.dpsDelta)} ${unitWord(deps.unit)}`,
    tone: tie ? 'muted' : 'wipe',
  };
}

interface Compare {
  gain: GainFigure;
  badge: 'upgrade' | 'worn' | undefined;
  worn: ItemRef | undefined;
  note: DollNote | undefined;
}

/** Signed in: the pick against the worn item (§5.10 item 8). Undefined when nothing is worn there. */
function compareWith(row: RowView, character: DollCharacter, deps: DollViewDeps): Compare | undefined {
  const wornId = character.worn[row.slot];
  if (wornId === undefined || row.itemId === undefined) return undefined;
  const gain = character.gains[row.slot];
  const wornDetail = deps.details.get(wornId);
  const worn = itemRef(wornId, wornDetail?.name ?? `Item ${wornId}`, deps.details);
  if (wornId === row.itemId || (gain !== null && gain !== undefined && isTie(gain))) {
    return {
      gain: { value: 'Worn', caption: 'You have this', tone: 'kill' },
      badge: 'worn',
      worn: undefined,
      note: undefined,
    };
  }
  const pickDetail = deps.details.get(row.itemId);
  const delta = statDelta(pickDetail?.stats, wornDetail?.stats);
  const level = wornDetail ? `item level ${wornDetail.item_level}` : undefined;
  return {
    gain:
      gain === null || gain === undefined
        ? { value: '—', caption: 'Not sim-checked', tone: 'muted' }
        : { value: signedFigure(gain), caption: `${unitWord(deps.unit)} vs your ${worn.name}`, tone: 'gold' },
    badge: 'upgrade',
    worn,
    note: {
      lead: 'You wear',
      item: worn,
      tail: [level && `(${level})`, delta && `The pick: ${delta}`].filter(Boolean).join(' · ') || undefined,
      tone: 'muted',
    },
  };
}

/** One RowView as a doll cell. */
export function dollCellFor(row: RowView, deps: DollViewDeps): DollCell {
  if (row.empty || row.itemId === undefined) {
    return { slot: row.slot, label: row.slotLabel, item: null, empty: row.emptyCopy };
  }
  const detail = deps.details.get(row.itemId);
  const item = itemRef(row.itemId, row.itemName ?? `Item ${row.itemId}`, deps.details, row.model?.quality);
  const next = row.alternatives?.[0];
  const compare = deps.character ? compareWith(row, deps.character, deps) : undefined;
  return {
    slot: row.slot,
    label: row.slotLabel,
    flag: levelFlag(detail?.required_level ?? row.model?.requiredLevel, deps.character),
    item,
    badge: compare?.badge,
    stats: shortStats(detail?.stats),
    source: row.sourceDetail,
    sourceKind: row.sourceKind,
    note: compare?.note ?? nextBestNote(next, deps),
    gain: compare?.gain ?? leadGain(next, deps.unit),
    worn: compare?.worn,
  };
}

/** Every row of a band, as cells. */
export function dollCellsFor(info: BandInfo, deps: DollViewDeps): DollCell[] {
  return info.rows.map((row) => dollCellFor(row, deps));
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
  if (upgrades.length === 0) return `${character.name} wears every pick for this level.`;
  const total = upgrades.reduce((sum, cell) => sum + (character.gains[cell.slot] ?? 0), 0);
  return (
    `${character.name} (level ${character.level}) has ${worn} of ${filled.length} picks. ` +
    `${upgrades.length} ${upgrades.length === 1 ? 'upgrade' : 'upgrades'}, ${signedFigure(total)} ${unitWord(unit)} available.`
  );
}
