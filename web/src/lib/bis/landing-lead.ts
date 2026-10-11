// web/src/lib/bis/landing-lead.ts
// The /bis landing's "your character" lead as a pure view model: which band the character is in,
// how many of that band's best-in-slot slots it wears (and which slots the band has no pick for),
// what the sync line says and when it turns stale, and the one button. The page script only
// writes these fields into the lead's markup.
import type { MeCharacter } from '../account/api';
import { isStale } from '../character-selector/stale';
import { relativeTime } from '../dates';
import { SLOTS, type Slot } from '../planner/types';
import { DEFAULT_LANDING_BAND, FIRST_BAND_LEVEL, characterBand } from './landing-bands';
import { classFor, specKeyFor, type LandingConfig } from './landing-config';
import { landingCopy } from './landing-copy';
import { specLinkHref } from './landing-links';
import { SLOT_DISPLAY_LABELS } from './slot-display-labels';
import { hasKnownSource } from './source-cell';
import type { BisBand, Faction } from './types';

/** Slot pairs the note names together: `Trinkets`, `Rings`. */
const PAIRED_SLOT_NAMES: Readonly<Record<string, string>> = { trinket: 'Trinkets', finger: 'Rings' };
const PAIR_SUFFIX = /[12]$/;

export interface SlotCount {
  worn: number;
  total: number;
  /** Player-facing names of the slots this band has no pick for, in paper-doll order. */
  excluded: string[];
}

/** The note's names: both of a pair together (`Trinkets`), a lone half by its own label. */
export function excludedSlotNames(slots: readonly Slot[]): string[] {
  const names: string[] = [];
  for (const slot of SLOTS.filter((candidate) => slots.includes(candidate))) {
    const group = slot.replace(PAIR_SUFFIX, '');
    const pairName = PAIRED_SLOT_NAMES[group];
    const bothMissing =
      pairName !== undefined && slots.includes(`${group}1` as Slot) && slots.includes(`${group}2` as Slot);
    const name = bothMissing ? pairName : SLOT_DISPLAY_LABELS[slot];
    if (!names.includes(name)) names.push(name);
  }
  return names;
}

/**
 * Of the slots this band has a sourced pick for, how many the character already wears the pick
 * in. The band's other slots are named, never counted: an unknown is not a zero.
 */
export function slotCount(band: Pick<BisBand, 'slots'>, gear: Readonly<Record<string, number>>): SlotCount {
  const sourced = band.slots.filter(hasKnownSource);
  const worn = sourced.filter((pick) => gear[pick.slot] === pick.item_id).length;
  const unsourced = SLOTS.filter((slot) => !sourced.some((pick) => pick.slot === slot));
  return { worn, total: sourced.length, excluded: excludedSlotNames(unsourced) };
}

/** `Trinkets have no pick at this band.`; null when every slot has a pick. */
export function excludedNote(excluded: readonly string[]): string | null {
  if (excluded.length === 0) return null;
  const last = excluded[excluded.length - 1]!;
  const slots = excluded.length === 1 ? last : `${excluded.slice(0, -1).join(', ')} and ${last}`;
  const plural = excluded.length > 1 || Object.values(PAIRED_SLOT_NAMES).includes(last);
  return landingCopy.excludedNote(slots, plural);
}

export interface SyncLine {
  text: string;
  /** The words past the stale limit, rendered in ember; null while the sync is fresh. */
  staleTail: string | null;
}

/** The addon sync line, or null when the site holds no addon build for the character. */
export function syncLine(build: MeCharacter['build'], now: Date): SyncLine | null {
  if (build?.source !== 'addon') return null;
  return {
    text: landingCopy.syncLine(relativeTime(new Date(build.captured_at), now)),
    staleTail: isStale(build, now) ? landingCopy.staleTail : null,
  };
}

export type CountView =
  | {
      kind: 'count';
      worn: number;
      total: number;
      percent: number;
      label: string;
      note: string | null;
    }
  | { kind: 'nosync' }
  | { kind: 'under20' }
  /** Gear on file, but no list to compare it with (no file for the spec, or none for the band). */
  | { kind: 'none' };

export interface LeadView {
  name: string;
  classSlug: string;
  classColor: string;
  /** `Fury Warrior`, or the bare class when the character has no spec yet. */
  specClass: string;
  level: number | undefined;
  faction: Faction | undefined;
  /** The character's spec page at the character's own band; never follows the strip. */
  href: string;
  buttonLabel: string;
  sync: SyncLine | null;
  count: CountView;
}

export interface LeadInput {
  character: MeCharacter;
  config: LandingConfig;
  /** The character's band entry in its spec's file, when the file and the band exist. */
  bandEntry: BisBand | undefined;
  now: Date;
}

function gearOf(character: MeCharacter): Readonly<Record<string, number>> | undefined {
  const gear = character.build?.gear;
  return gear === undefined || Object.keys(gear).length === 0 ? undefined : gear;
}

function countViewFor(input: LeadInput, stale: boolean): CountView {
  const { character, bandEntry } = input;
  if (character.level !== undefined && character.level < FIRST_BAND_LEVEL) return { kind: 'under20' };
  const gear = gearOf(character);
  if (gear === undefined) return { kind: 'nosync' };
  if (bandEntry === undefined) return { kind: 'none' };
  const { worn, total, excluded } = slotCount(bandEntry, gear);
  return {
    kind: 'count',
    worn,
    total,
    percent: total === 0 ? 0 : (worn / total) * 100,
    label: stale ? landingCopy.countLabelStale : landingCopy.countLabel,
    note: excludedNote(excluded),
  };
}

/** The lead for a character, or null when its class is unknown (nothing to open a page for). */
export function buildLeadView(input: LeadInput): LeadView | null {
  const { character, config, now } = input;
  const cls = classFor(config, character.class);
  if (cls === undefined) return null;
  const specKey = specKeyFor(config, character.class, character.spec);
  const spec = config.specs[specKey ?? cls.defaultSpecKey];
  if (spec === undefined) return null;
  const band = characterBand(character.level) ?? DEFAULT_LANDING_BAND;
  const sync = syncLine(character.build, now);
  const stale = sync !== null && sync.staleTail !== null;
  return {
    name: character.name,
    classSlug: cls.slug,
    classColor: cls.color,
    specClass: specKey === undefined ? cls.name : `${spec.name} ${cls.name}`,
    level: character.level,
    faction: character.faction,
    href: specLinkHref({
      classSlug: spec.classSlug,
      specSlug: spec.specSlug,
      band,
      faction: character.faction,
    }),
    buttonLabel: landingCopy.leadButton(specKey === undefined ? cls.name : spec.name, band),
    sync,
    count: countViewFor(input, stale),
  };
}
