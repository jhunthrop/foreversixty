// web/src/lib/gear/paperdoll.ts
// The paperdoll's pure model (design/specs/2026-10-31-paperdoll.md): where each slot sits in the
// character frame, what its upgrade line says, the item-level ruling and the panel's words.
// Paperdoll.svelte draws what this returns; the sim card, the character page and the report
// each map their own data to `PaperdollSlot`s and nothing else.
import { wearsListedBest } from '../bis/best-in-slot';
import { itemIdLabel, itemNotInDataPhrase } from '../items/unknown-item-copy';
import { SLOT_LABELS, type Slot } from '../planner/types';
import { gearCopy } from './copy';
import type { CharacterGearSlot } from './types';

export type PaperdollSlotKey = Slot | 'shirt' | 'tabard';

/** An item as the frame draws it. `icon` and `quality` are null when the items file has not
 *  said (failed): the plain socket, the name in the body colour. */
export interface PaperdollItem {
  id: number;
  name: string;
  icon: string | null;
  quality: number | null;
  itemLevel: number | null;
  twoHand: boolean;
}

/** A pick, for its icon, colour and tooltip. */
export interface PaperdollItemRef {
  id: number;
  name: string;
  icon: string | null;
  quality: number | null;
}

export type PaperdollUpgrade =
  | { kind: 'gain'; points: number; pick: PaperdollItemRef }
  | { kind: 'best' }
  | { kind: 'unscored'; pick: PaperdollItemRef };

export interface PaperdollSlot {
  slot: PaperdollSlotKey;
  state: 'item' | 'empty' | 'unknown';
  /** With `state: 'item'`. */
  item?: PaperdollItem;
  /** With `state: 'unknown'`: the id the data does not have. */
  unknownId?: number;
  upgrade?: PaperdollUpgrade;
  /** The report's "Item level 66 · Enchanted". */
  detail?: string;
}

export interface PaperdollIdentity {
  name: string;
  classSlug: string;
  /** "Human"; empty when unknown (a combat log names no race). */
  race: string;
  /** "Arms". */
  spec: string;
  level: number | null;
  faction: 'alliance' | 'horde' | null;
  /** The source pill's words ("Addon export, 12 min ago", "At pull"). */
  source: string;
}

export type PaperdollSide = 'left' | 'right' | 'centre';

/** One element of the frame in DOM order: a slot cell or its phone upgrade row. */
export interface PaperdollCell {
  kind: 'slot' | 'upgrade';
  slot: PaperdollSlotKey;
  side: PaperdollSide;
  /** The desktop grid row (1 to 10). */
  row: number;
  /** The phone grid row; row 1 is the identity. */
  phoneRow: number;
  /** DOM order. */
  order: number;
  /** The empty off hand under a two-handed main hand: "Two-hand in main hand". */
  twoHandOff: boolean;
  data: PaperdollSlot;
}

const LEFT: readonly PaperdollSlotKey[] = [
  'head',
  'neck',
  'shoulder',
  'back',
  'chest',
  'shirt',
  'tabard',
  'wrist',
];
const RIGHT: readonly PaperdollSlotKey[] = [
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
];
/** Shirt and tabard are drawn only when the data carries them (§3.1). */
const OPTIONAL: ReadonlySet<PaperdollSlotKey> = new Set(['shirt', 'tabard']);
const WEAPON_ROW = 9;
const RANGED_ROW = 10;
/** The client's own classes with a relic slot in place of ranged. */
const RELIC_CLASSES: ReadonlySet<string> = new Set(['paladin', 'shaman', 'druid']);

/** The phone frame draws an upgrade row for a gain or "Not scored"; never for "Best in slot". */
export function phoneUpgradeShown(upgrade: PaperdollUpgrade | undefined): boolean {
  return upgrade !== undefined && upgrade.kind !== 'best';
}

/** The off hand is empty and the main hand is two-handed. */
function twoHandOffHand(bySlot: ReadonlyMap<PaperdollSlotKey, PaperdollSlot>): boolean {
  return bySlot.get('off_hand')?.state !== 'item' && bySlot.get('main_hand')?.item?.twoHand === true;
}

/** Every slot the frame draws, in its column order; a missing required slot is drawn empty. */
function columns(bySlot: ReadonlyMap<PaperdollSlotKey, PaperdollSlot>): {
  left: PaperdollSlot[];
  right: PaperdollSlot[];
} {
  const pick = (keys: readonly PaperdollSlotKey[]): PaperdollSlot[] =>
    keys
      .filter((key) => !OPTIONAL.has(key) || bySlot.has(key))
      .map((key) => bySlot.get(key) ?? { slot: key, state: 'empty' });
  return { left: pick(LEFT), right: pick(RIGHT) };
}

/**
 * The frame's cells in DOM order (§3.3): pairs of left and right slots, each pair followed by its
 * phone upgrade rows (the left slot's first), then the weapons pair and its rows, then ranged and
 * its row. The DOM order is the phone's visual order; the desktop grid places the same elements
 * by `row` and `side`.
 */
export function paperdollLayout(slots: readonly PaperdollSlot[]): PaperdollCell[] {
  const bySlot = new Map(slots.map((slot) => [slot.slot, slot]));
  const offUnderTwoHand = twoHandOffHand(bySlot);
  const { left, right } = columns(bySlot);
  const empty = (key: PaperdollSlotKey): PaperdollSlot => bySlot.get(key) ?? { slot: key, state: 'empty' };
  const pairs: { left?: PaperdollSlot; right?: PaperdollSlot; row: number }[] = [];
  for (let index = 0; index < Math.max(left.length, right.length); index++)
    pairs.push({ left: left[index], right: right[index], row: index + 1 });
  pairs.push({ left: empty('main_hand'), right: empty('off_hand'), row: WEAPON_ROW });

  const cells: PaperdollCell[] = [];
  let phoneRow = 2;
  const push = (
    kind: PaperdollCell['kind'],
    data: PaperdollSlot,
    side: PaperdollSide,
    row: number,
    at: number,
  ) =>
    cells.push({
      kind,
      slot: data.slot,
      side,
      row,
      phoneRow: at,
      order: cells.length,
      twoHandOff: kind === 'slot' && data.slot === 'off_hand' && offUnderTwoHand,
      data,
    });
  for (const pair of pairs) {
    const pairRow = phoneRow++;
    if (pair.left) push('slot', pair.left, 'left', pair.row, pairRow);
    if (pair.right) push('slot', pair.right, 'right', pair.row, pairRow);
    for (const [data, side] of [
      [pair.left, 'left'],
      [pair.right, 'right'],
    ] as const)
      if (data && phoneUpgradeShown(data.upgrade)) push('upgrade', data, side, pair.row, phoneRow++);
  }
  const ranged = empty('ranged');
  push('slot', ranged, 'centre', RANGED_ROW, phoneRow++);
  if (phoneUpgradeShown(ranged.upgrade)) push('upgrade', ranged, 'centre', RANGED_ROW, phoneRow);
  return cells;
}

/**
 * The centre's item level (§3.5): the mean item level of every equipped slot but shirt and tabard,
 * one decimal. An empty slot (the off hand under a two-hander included) is not equipped. Null when
 * nothing is equipped, or when any equipped item is not in our data or has no known level: never a
 * partial mean.
 */
export function averageItemLevel(slots: readonly PaperdollSlot[]): string | null {
  const counted = slots.filter((slot) => !OPTIONAL.has(slot.slot) && slot.state !== 'empty');
  if (counted.length === 0) return null;
  const levels = counted.map((slot) => (slot.state === 'item' ? (slot.item?.itemLevel ?? null) : null));
  if (levels.some((level) => level === null)) return null;
  const total = (levels as number[]).reduce((sum, level) => sum + level, 0);
  return (total / levels.length).toFixed(1);
}

/** How many equipped slots hold an item our data does not have. */
export function unknownCount(slots: readonly PaperdollSlot[]): number {
  return slots.filter((slot) => slot.state === 'unknown').length;
}

/**
 * A character-page slot's upgrade (§3.4) from the API's `CharacterGearSlot`. `wornBySlot` is the
 * worn item id per slot, for the rings and trinkets matched either way round; `ref` resolves a
 * pick's icon and quality. Never a gain when the pick is worn.
 */
export function upgradeFor(
  gear: CharacterGearSlot,
  wornBySlot: Readonly<Record<string, number | undefined>>,
  ref: (id: number, name: string) => PaperdollItemRef,
): PaperdollUpgrade | undefined {
  const { worn, pick, gap } = gear;
  if (pick === null) return undefined;
  if (worn !== null && worn.name === null) return { kind: 'unscored', pick: ref(pick.id, pick.name) };
  if (worn !== null && wearsListedBest(gear.slot, pick.id, { ...wornBySlot, [gear.slot]: worn.id }))
    return { kind: 'best' };
  if (gap === null) return worn === null ? { kind: 'unscored', pick: ref(pick.id, pick.name) } : undefined;
  return gap > 0 ? { kind: 'gain', points: gap, pick: ref(pick.id, pick.name) } : undefined;
}

/** The slot's word: `SLOT_LABELS`, plus shirt, tabard and the relic classes' "Relic". */
export function slotLabel(slot: PaperdollSlotKey, classSlug = ''): string {
  if (slot === 'shirt') return gearCopy.paperdoll.shirt;
  if (slot === 'tabard') return gearCopy.paperdoll.tabard;
  if (slot === 'ranged' && RELIC_CLASSES.has(classSlug)) return gearCopy.paperdoll.relic;
  return SLOT_LABELS[slot];
}

/** "Waist · Item 288887": an unknown item always carries its slot word (§3.3). */
export function unknownName(slot: PaperdollSlotKey, id: number, classSlug = ''): string {
  return `${slotLabel(slot, classSlug)} · ${itemIdLabel(id)}`;
}

/** The cell's accessible name (§6): "Head: Lionheart Helm". */
export function cellLabel(slot: PaperdollSlot, classSlug = '', twoHandOff = false): string {
  const label = slotLabel(slot.slot, classSlug);
  if (slot.state === 'item' && slot.item) return `${label}: ${slot.item.name}`;
  if (slot.state === 'unknown')
    return `${label}: ${itemIdLabel(slot.unknownId ?? 0)}, ${itemNotInDataPhrase.toLowerCase()}`;
  return `${label}: ${twoHandOff ? gearCopy.paperdoll.twoHand.toLowerCase() : gearCopy.paperdoll.empty}`;
}

/** The pick's accessible name (§6): "Head upgrade: Lionheart Helm, 60 points to gain". */
export function pickLabel(slot: PaperdollSlotKey, upgrade: PaperdollUpgrade, classSlug = ''): string {
  if (upgrade.kind === 'best') return `${slotLabel(slot, classSlug)}: ${gearCopy.paperdoll.best}`;
  const tail =
    upgrade.kind === 'gain' ? gearCopy.gapLabel(upgrade.points) : gearCopy.paperdoll.notScoredLabel;
  return `${slotLabel(slot, classSlug)}${gearCopy.paperdoll.upgradeLabel}${upgrade.pick.name}, ${tail}`;
}

/** The note under the stage (§3.3, §6): once, only when a slot is unknown or the file failed. */
export function paperdollNote(unknown: number, failed: boolean): string | null {
  if (failed) return gearCopy.paperdoll.noteFailed;
  if (unknown === 0) return null;
  return unknown === 1 ? gearCopy.paperdoll.noteOne : gearCopy.paperdoll.noteMany(unknown);
}

const SOFT_HYPHEN = '­';
const VOWEL = /[aeiouy]/i;
/** Consonant onsets English keeps together at a syllable's start. */
const ONSETS: ReadonlySet<string> = new Set([
  'bl',
  'br',
  'ch',
  'cl',
  'cr',
  'dr',
  'fl',
  'fr',
  'gl',
  'gr',
  'kn',
  'ph',
  'pl',
  'pr',
  'qu',
  'sc',
  'sh',
  'sk',
  'sl',
  'sm',
  'sn',
  'sp',
  'st',
  'sw',
  'th',
  'tr',
  'tw',
  'wh',
  'wr',
  'chr',
  'sch',
  'scr',
  'shr',
  'spl',
  'spr',
  'str',
  'thr',
]);
/** Only a word this long can be wider than the narrowest name line (a 360 px phone's ~100 px at
 *  12 px holds about 12 letters, so "Shoulderguard" needs a break): shorter words wrap whole, and a
 *  soft hyphen in one would make the browser hyphenate a word that fits its own line ("Mysta-ria"). */
const MIN_WORD = 12;
const MIN_PART = 3;
/** A syllable between two breaks may be two letters ("mag-ne-tic"). */
const MIN_INNER = 2;

/** The onset starting at `index`, when a vowel follows it: its length, or 0. */
function onsetAt(word: string, index: number): number {
  if (VOWEL.test(word[index] ?? '')) return 0;
  for (const length of [3, 2, 1]) {
    const onset = word.slice(index, index + length).toLowerCase();
    const known = length === 1 || ONSETS.has(onset);
    if (known && onset.length === length && VOWEL.test(word[index + length] ?? '')) return length;
  }
  return 0;
}

function hyphenateWord(word: string): string {
  if (word.length < MIN_WORD) return word;
  let out = '';
  let last = 0;
  for (let index = MIN_PART; index <= word.length - MIN_PART; index++) {
    if (onsetAt(word, index) === 0 || !VOWEL.test(word.slice(last, index))) continue;
    const across = word.slice(index - 1, index + 1).toLowerCase();
    if (!VOWEL.test(word[index - 1] ?? '') && ONSETS.has(across)) continue;
    if (index - last < (last === 0 ? MIN_PART : MIN_INNER)) continue;
    out += word.slice(last, index) + SOFT_HYPHEN;
    last = index;
  }
  return out + word.slice(last);
}

/**
 * Soft hyphens at the syllable breaks of every long word (§3.3, build requirement): a word wider
 * than its line breaks with a visible hyphen in every engine, without relying on `hyphens: auto`,
 * which headless Chromium and some WebKit builds do not apply. Only words too long for the
 * narrowest line get them, at maximal-onset breaks, never within three letters of either end.
 * Punctuation and spacing are untouched.
 */
export function hyphenate(text: string): string {
  return text.replace(/[A-Za-z]+/g, hyphenateWord);
}
