// web/src/lib/bis/landing-answer.ts
// What the /bis landing hands the shared character header (components/character/
// CharacterHeader.svelte): the answer ("Best-in-slot slots you wear", how many of the character's
// band's best-in-slot slots it already wears, and which slots the band has no pick for) and the
// one action (the character's own spec page at its own band). The identity, the sync line and
// every state are the header's own. The spec page's card reuses the answer through
// `bisAnswerSpec` with its own spec lookup.
import type { MeCharacter } from '../account/api';
import type { AnswerSpec, HeaderAction, HeaderAnswer } from '../character-header/types';
import { characterSpecName } from '../character-spec';
import { homeHeroCardsCopy } from '../home-panel-copy';
import { SLOTS, type Slot } from '../planner/types';
import { wearsListedBest } from './best-in-slot';
import { bandEntryFor, fetchBisFile } from './hover';
import { DEFAULT_LANDING_BAND, FIRST_BAND_LEVEL, characterBand } from './landing-bands';
import { classFor, specKeyFor, type LandingConfig } from './landing-config';
import { landingCopy } from './landing-copy';
import { specLinkHref } from './landing-links';
import { SLOT_DISPLAY_LABELS } from './slot-display-labels';
import { hasKnownSource } from './source-cell';
import type { BisBand, Faction } from './types';

const DEFAULT_LOOKUP_FACTION: Faction = 'alliance';

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
  const worn = sourced.filter((pick) => wearsListedBest(pick.slot, pick.item_id, gear)).length;
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

/** The sentence for a character with gear on file but no list to compare it with. */
const NO_LIST: HeaderAnswer = { kind: 'sentence', text: homeHeroCardsCopy.bestInSlotNoListYet };

export function countAnswer(
  band: Pick<BisBand, 'slots'>,
  gear: Readonly<Record<string, number>>,
): HeaderAnswer {
  const { worn, total, excluded } = slotCount(band, gear);
  const note = excludedNote(excluded);
  return {
    kind: 'figures',
    label: landingCopy.countLabel,
    staleLabel: landingCopy.countLabelStale,
    figures: [{ value: String(worn), unit: landingCopy.countUnit(total) }],
    bar: total === 0 ? 0 : worn / total,
    ...(note === null ? {} : { note: { text: note } }),
  };
}

export interface BisAnswerOptions {
  build: string;
  /** The spec whose list the answer counts against; undefined when the character has none the site knows. */
  specKeyOf(character: MeCharacter): string | undefined;
  /** What the slot shows when `specKeyOf` finds no spec, or the spec has no list at the band. */
  withoutList?: HeaderAnswer;
}

/** The character's band entry in the spec's published file; undefined when there is no file or no entry. */
async function bandEntryOf(
  build: string,
  specKey: string,
  character: MeCharacter,
): Promise<BisBand | undefined> {
  const file = await fetchBisFile(build, specKey);
  if (file === null) return undefined;
  const band = characterBand(character.level) ?? DEFAULT_LANDING_BAND;
  return bandEntryFor(file, band, character.faction ?? DEFAULT_LOOKUP_FACTION);
}

/** The BiS count answer for a character that is at or above level 20 and has gear on file. */
export async function loadBisAnswer(
  character: MeCharacter,
  options: BisAnswerOptions,
): Promise<HeaderAnswer> {
  const withoutList = options.withoutList ?? NO_LIST;
  const specKey = options.specKeyOf(character);
  if (specKey === undefined) return withoutList;
  const entry = await bandEntryOf(options.build, specKey, character);
  if (entry === undefined) return withoutList;
  return countAnswer(entry, character.build?.gear ?? {});
}

/** The header's answer spec for any page that counts worn best-in-slot picks. */
export function bisAnswerSpec(options: BisAnswerOptions): AnswerSpec {
  return {
    levelFloor: FIRST_BAND_LEVEL,
    needsGear: true,
    placeholderBar: true,
    strings: {
      belowLevelFloor: landingCopy.under20,
      noGear: {
        sentence: landingCopy.noSync,
        link: { label: landingCopy.noSyncLink, href: landingCopy.noSyncHref },
      },
      failed: landingCopy.countFailed,
    },
    load: (character) => loadBisAnswer(character, options),
  };
}

/** The landing's answer: counts against the spec the character reports. */
export function landingAnswerSpec(config: LandingConfig): AnswerSpec {
  return bisAnswerSpec({
    build: config.build,
    specKeyOf: (character) => specKeyFor(config, character.class, characterSpecName(character)),
  });
}

/** The character's own spec page at its own band and faction; the pointer to the grid for an unknown class. */
export function landingAction(character: MeCharacter, config: LandingConfig): HeaderAction {
  const cls = classFor(config, character.class);
  const specKey = specKeyFor(config, character.class, characterSpecName(character));
  const spec = cls === undefined ? undefined : config.specs[specKey ?? cls.defaultSpecKey];
  if (cls === undefined || spec === undefined) {
    return { label: landingCopy.fallbackAction, href: landingCopy.fallbackHref };
  }
  const band = characterBand(character.level) ?? DEFAULT_LANDING_BAND;
  return {
    label: landingCopy.leadButton(specKey === undefined ? cls.name : spec.name, band),
    href: specLinkHref({
      classSlug: spec.classSlug,
      specSlug: spec.specSlug,
      band,
      faction: character.faction,
    }),
  };
}
