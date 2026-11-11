// web/src/lib/gear/doll-cells.ts
// The worn and report paperdolls' cells (web layout spec 2026-11-05 §4.2, option C of §5.13): the
// one mapping from `PaperdollSlot` (every surface's slots, lib/gear/paperdoll-slots.ts) to the
// shared doll's `DollCell`. Line 1 is the worn item; line 2 its item level (the log's own on a
// report), then its enchant in the client's green: one format on every worn surface ("Item level
// 61 · Crusader", a report's "Item level 66 · Enchanted" where the id has no name). In `worn` mode the gain column
// is the upgrade to the band's pick in the band's unit (`+1.8 DPS` over `vs <pick>`), `Best in
// slot`, `Not scored`, `Same` for a tie (lib/bis/doll-view.ts `isTie`, the BiS page's rule) and
// `Not sim-checked` for a gain the rule cannot price; never 0. `report` mode has no gain.
import { isTie } from '../bis/doll-view';
import { scoreUnitWord, type SlotScoreUnit } from '../bis/tank-view';
import { itemNotInDataPhrase } from '../items/unknown-item-copy';
import type { DollCell } from '../ui/paperdoll';
import { signedFigure, type GainFigure, type ItemRef } from '../ui/primitives';
import { gearCopy } from './copy';
import {
  slotLabel,
  twoHandCover,
  unknownName,
  type PaperdollItemRef,
  type PaperdollSlot,
  type PaperdollUpgrade,
} from './paperdoll';
import { EMPTY_FIGURE } from '../format/number';

export type WornDollMode = 'worn' | 'report';

/** The common quality: a name the items file has not coloured (it failed) reads plain. */
const PLAIN_QUALITY = 1;
const POOR_QUALITY = 0;

function itemRef(ref: PaperdollItemRef): ItemRef {
  return {
    id: ref.id,
    name: ref.name,
    quality: ref.quality ?? PLAIN_QUALITY,
    icon: ref.icon ?? undefined,
  };
}

/** The gain column for a slot's upgrade (§4.2 `worn`). */
export function wornGain(upgrade: PaperdollUpgrade): Pick<DollCell, 'gain' | 'badge'> {
  if (upgrade.kind === 'best')
    return { gain: { value: gearCopy.paperdoll.best, tone: 'kill' }, badge: 'worn' };
  const basis = { prefix: gearCopy.paperdoll.versus, item: itemRef(upgrade.pick) };
  if (upgrade.kind === 'unscored')
    return { gain: { value: gearCopy.paperdoll.notScored, basis, tone: 'muted' } };
  if (upgrade.value === null)
    return {
      gain: { value: EMPTY_FIGURE, basis: { prefix: gearCopy.paperdoll.notSimChecked }, tone: 'muted' },
      badge: 'upgrade',
    };
  if (isTie(upgrade.value)) return { gain: { value: gearCopy.paperdoll.same, basis, tone: 'muted' } };
  const gain: GainFigure = {
    value: signedFigure(upgrade.value),
    unit: scoreUnitWord(upgrade.unit),
    basis,
    tone: 'gold',
  };
  return { gain, badge: 'upgrade' };
}

/** Line 2 of a worn or report cell: the item level (the caller's detail first), then the enchant. */
function line2(slot: PaperdollSlot): Pick<DollCell, 'source' | 'enchant'> {
  const level = slot.item?.itemLevel;
  const source =
    slot.detail ??
    (level === null || level === undefined ? undefined : gearCopy.paperdoll.itemLevelLine(level));
  return { source, ...(slot.enchant === undefined ? {} : { enchant: slot.enchant }) };
}

function cellFor(
  slot: PaperdollSlot,
  mode: WornDollMode,
  classSlug: string,
  cover: PaperdollItemRef | null,
): DollCell {
  const base = { slot: slot.slot, label: slotLabel(slot.slot, classSlug) };
  const gain = mode === 'worn' && slot.upgrade !== undefined ? wornGain(slot.upgrade) : {};
  if (slot.state === 'unknown') {
    const id = slot.unknownId ?? 0;
    return {
      ...base,
      ...gain,
      item: { id, name: unknownName(slot.slot, id, classSlug), quality: POOR_QUALITY },
      unknown: true,
      source: itemNotInDataPhrase,
    };
  }
  if (slot.state === 'empty' || slot.item === undefined) {
    const covered = slot.slot === 'off_hand' ? cover : null;
    return {
      ...base,
      ...gain,
      item: null,
      emptyTitle: covered ? gearCopy.paperdoll.twoHand : gearCopy.paperdoll.emptyTitle,
      coveredBy: covered ? itemRef(covered) : undefined,
    };
  }
  return { ...base, ...gain, item: itemRef(slot.item), ...line2(slot) };
}

/** Every slot as the shared doll's cell, in the slots' own order (the doll lays out the columns). */
export function dollCellsFromSlots(
  slots: readonly PaperdollSlot[],
  mode: WornDollMode,
  classSlug = '',
): DollCell[] {
  const cover = twoHandCover(slots);
  return slots.map((slot) => cellFor(slot, mode, classSlug, cover));
}

/** The head's gold line: "+12.4 DPS across 9 slots", the priced gains that are not ties summed. */
export function wornGainSummary(slots: readonly PaperdollSlot[]): string | null {
  const gains = slots.flatMap((slot) =>
    slot.upgrade?.kind === 'gain' && slot.upgrade.value !== null && !isTie(slot.upgrade.value)
      ? [{ value: slot.upgrade.value, unit: slot.upgrade.unit }]
      : [],
  );
  if (gains.length === 0) return null;
  const unit: SlotScoreUnit = gains[0].unit;
  const total = gains.reduce((sum, gain) => sum + gain.value, 0);
  return gearCopy.gainSummary(signedFigure(total), scoreUnitWord(unit), gains.length);
}
