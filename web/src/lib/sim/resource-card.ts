// web/src/lib/sim/resource-card.ts
// The Resource card's figures (spec 2026-11-04-sim-results 3.10): the three tiles from the
// summary, the two lists from `resource_sources`. Every figure is the engine's average per fight;
// none is recomputed beyond a share and a rounding.
import { attackHand, parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { damageCopy, powerCopy } from './result-copy';
import { resourceCardCopy as copy } from './input-cards-copy';
import { REFUND_KEY } from './action-mark';
import { primaryPowerTrack, type PrimaryPower } from './primary-power';
import type { ResourceSourceRow, SimResult } from './types';
import { grouped, oneDecimal } from '../format/number';

export interface ResourceRow {
  action: string;
  name: string;
  amount: string;
  /** Against the list's top row, 0 to 100. */
  barPct: number;
  /** "−22.8 lost"; null when nothing overflowed. */
  lost: string | null;
}

export interface ResourceTile {
  label: string;
  value: string;
}

export interface ResourceCard {
  /** "Rage", "Energy", "Mana". */
  title: string;
  tiles: ResourceTile[];
  /** Null on a run saved before the sources were kept. */
  lists: { gained: ResourceRow[]; gainedTotal: string; spent: ResourceRow[]; spentTotal: string } | null;
}

const WHOLE_FROM = 10;
/** Under this an overflow rounds to 0.0 and says nothing. */
const LOST_FLOOR = 0.05;

/** Whole numbers with separators; one decimal under 10. */
export function resourceAmount(value: number): string {
  return value < WHOLE_FROM ? oneDecimal(value) : grouped(value);
}

/** A white swing reads as the damage table names it (spec section 6): "Main hand". */
function sourceName(action: string, names: ActionNames | null): string {
  if (action === REFUND_KEY) return copy.refunded;
  const parsed = parseActionKey(action);
  const hand = parsed?.kind === 'other' && parsed.label === 'attack' ? attackHand(parsed.tag) : null;
  return hand === null ? resolveActionName(action, names) : damageCopy.hands[hand];
}

function rowsOf(rows: readonly ResourceSourceRow[], names: ActionNames | null): ResourceRow[] {
  const sorted = [...rows].sort((a, b) => b.amount - a.amount);
  const top = sorted[0]?.amount ?? 0;
  return sorted.map((row) => ({
    action: row.action,
    name: sourceName(row.action, names),
    amount: resourceAmount(row.amount),
    barPct: top === 0 ? 0 : (row.amount / top) * 100,
    lost: (row.wasted ?? 0) >= LOST_FLOOR ? copy.lost(oneDecimal(row.wasted ?? 0)) : null,
  }));
}

const total = (rows: readonly ResourceSourceRow[]): number => rows.reduce((sum, row) => sum + row.amount, 0);

/** Mana has no cap to lose against: what the sample run ended with, when it recorded any. */
function manaLeft(result: SimResult): number | undefined {
  const rows = result.sample ?? [];
  for (let at = rows.length - 1; at >= 0; at -= 1) {
    const mana = rows[at]!.resources?.mana;
    if (mana !== undefined) return mana;
  }
  return undefined;
}

function thirdTile(
  result: SimResult,
  power: PrimaryPower,
  gained: number,
  wasted: number | undefined,
): ResourceTile[] {
  if (power === 'mana') {
    const left = manaLeft(result);
    return left === undefined ? [] : [{ label: copy.endedWith, value: copy.amount(grouped(left)) }];
  }
  // A run that never recorded the overflow says nothing about it, rather than a 0 it did not measure.
  if (wasted === undefined) return [];
  const share = gained === 0 ? 0 : (wasted / gained) * 100;
  return [{ label: copy.lostAtCap, value: copy.capValue(grouped(wasted), oneDecimal(share)) }];
}

/** Null when the run reports no resource at all. */
export function resourceCard(result: SimResult, names: ActionNames | null): ResourceCard | null {
  const found = primaryPowerTrack(result);
  if (found === null) return null;
  const { power, track: resource } = found;
  const title = powerCopy[power];
  const sources = result.resource_sources?.find((row) => row.resource === power);
  return {
    title,
    tiles: [
      { label: copy.gained, value: copy.amount(grouped(resource.gained)) },
      { label: copy.spent, value: copy.amount(grouped(resource.spent)) },
      ...thirdTile(result, power, resource.gained, resource.wasted),
    ],
    lists:
      sources === undefined
        ? null
        : {
            gained: rowsOf(sources.gained, names),
            gainedTotal: grouped(total(sources.gained)),
            spent: rowsOf(sources.spent, names),
            spentTotal: grouped(total(sources.spent)),
          },
  };
}
