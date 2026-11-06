// web/src/lib/bis/source-line.ts
// The full source line a BiS row shows (web layout spec 2026-11-05 §5.10 item 3): where the item
// comes from in words a leveler can act on, wrapping rather than truncating. Built from the same
// resolved `SourceCell` the rest of the page reads (`source-cell.ts`), plus the facts the build
// data carries beside it: an instance's level range (the client's dungeon finder rows in
// data/curated/raids/events.json) and the item's bind rule. A fact the data does not carry is
// left out, never guessed: the craft skill and quest zone show only when loot.json carries them
// (loot-fields lane), and a source's own level range wins over the curated instance table.
import { describeSourceCell, type SourceCell } from './source-cell';

/** An instance's level range, by its display name. */
export type InstanceLevels = ReadonlyMap<string, { min: number; max: number }>;

export interface SourceLineContext {
  instanceLevels: InstanceLevels;
  /** The item's bind rule (`ItemDetail.bonding`): `equip` adds the auction-house note. */
  bonding?: string;
}

/** A drop chance as the row prints it: a whole percent from 1, one decimal below. The row's
 *  own words are kept short so a source fits the paperdoll's second line (owner option C). */
export function dropChanceText(percent: number): string {
  return `${percent >= 1 ? Math.round(percent) : Number(percent.toFixed(1))}%`;
}

function titleCase(word: string): string {
  return word.replace(/\b\w/g, (letter) => letter.toUpperCase());
}

/** ` (15-24)`, or ` (60)` for a one-level range; empty when unknown. */
export function levelsSuffix(levels: { min: number; max: number } | undefined): string {
  if (levels === undefined) return '';
  return levels.min === levels.max ? ` (${levels.min})` : ` (${levels.min}-${levels.max})`;
}

/** Bind on equip: tradeable, so also on the auction house (the game's own short words). */
const BOE_NOTE = 'BoE · on the AH';

/** The one line, joined with middle dots. */
export function fullSourceLine(cell: SourceCell, context: SourceLineContext): string {
  const boe = context.bonding === 'equip';
  switch (cell.kind) {
    case 'dungeon':
    case 'raid': {
      const own =
        cell.levelMin !== undefined && cell.levelMax !== undefined
          ? { min: cell.levelMin, max: cell.levelMax }
          : undefined;
      const place = `${cell.instance}${levelsSuffix(own ?? context.instanceLevels.get(cell.instance))}`;
      return [cell.boss, place, cell.dropChance === undefined ? undefined : dropChanceText(cell.dropChance)]
        .filter(Boolean)
        .join(' · ');
    }
    case 'crafted':
      return cell.skill === undefined
        ? `Crafted · ${titleCase(cell.name)}`
        : `Crafted · ${titleCase(cell.name)} (skill ${cell.skill})`;
    case 'rep':
      return cell.standing === undefined
        ? `${cell.faction} reputation`
        : `${cell.faction} reputation · ${titleCase(cell.standing)}`;
    case 'quest':
      return [
        `Quest: ${cell.questName}`,
        cell.zone,
        cell.minLevel > 0 ? `requires level ${cell.minLevel}` : undefined,
      ]
        .filter(Boolean)
        .join(' · ');
    case 'vendor':
      return `Vendor: ${cell.npc}`;
    case 'world_drop': {
      const range =
        cell.levelMin !== undefined && cell.levelMax !== undefined
          ? levelsSuffix({ min: cell.levelMin, max: cell.levelMax })
          : '';
      return boe ? `World drop${range} · ${BOE_NOTE}` : `World drop${range}`;
    }
    case 'zone':
    case 'world': {
      const parts = [cell.place, cell.dropChance === undefined ? undefined : dropChanceText(cell.dropChance)];
      if (boe) parts.push(BOE_NOTE);
      return parts.filter(Boolean).join(' · ');
    }
    case 'pvp':
      return describeSourceCell(cell);
    case 'unknown':
      return cell.label;
  }
}

/** The reputation standing a source needs, when it is a reputation source that names one. */
export function standingOf(cell: SourceCell): string | undefined {
  return cell.kind === 'rep' ? cell.standing?.toLowerCase() : undefined;
}
