// web/src/lib/bis/landing-answer.ts
// What the /bis landing's head answer counts (landing-answer-view.ts paints it): how many of the
// character's band's best-in-slot slots it already wears and which slots the band has no pick for,
// as the shared AnswerSpec contract (lib/character-header/types.ts), and the one action (the
// character's own spec page at its own band).
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
import { levelingSpecKey } from '../leveling/leveling-plan';
import { specLinkHref } from './landing-links';
import { slotDisplayLabel } from './slot-display-labels';
import { hasKnownSource } from './source-cell';
import type { BisBand, BisFile, Faction } from './types';

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

/** The note's names: both of a pair together (`Trinkets`), a lone half by its own label, a
 *  relic class's ranged slot as `Relic` (`slotDisplayLabel`). */
export function excludedSlotNames(slots: readonly Slot[], classSlug?: string): string[] {
  const names: string[] = [];
  for (const slot of SLOTS.filter((candidate) => slots.includes(candidate))) {
    const group = slot.replace(PAIR_SUFFIX, '');
    const pairName = PAIRED_SLOT_NAMES[group];
    const bothMissing =
      pairName !== undefined && slots.includes(`${group}1` as Slot) && slots.includes(`${group}2` as Slot);
    const name = bothMissing ? pairName : slotDisplayLabel(slot, classSlug);
    if (!names.includes(name)) names.push(name);
  }
  return names;
}

/**
 * Of the slots this band has a sourced pick for, how many the character already wears the pick
 * in. The band's other slots are named, never counted: an unknown is not a zero.
 */
export function slotCount(
  band: Pick<BisBand, 'slots'>,
  gear: Readonly<Record<string, number>>,
  classSlug?: string,
): SlotCount {
  const sourced = band.slots.filter(hasKnownSource);
  const worn = sourced.filter((pick) => wearsListedBest(pick.slot, pick.item_id, gear)).length;
  const unsourced = SLOTS.filter((slot) => !sourced.some((pick) => pick.slot === slot));
  return { worn, total: sourced.length, excluded: excludedSlotNames(unsourced, classSlug) };
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
  classSlug?: string,
): HeaderAnswer {
  const { worn, total, excluded } = slotCount(band, gear, classSlug);
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
  /** Reads a spec's published file (default: the browser fetch); the mock facts read it off disk. */
  fetchFile?: (build: string, specKey: string) => Promise<BisFile | null>;
}

/** The character's band entry in the spec's published file; undefined when there is no file or no entry. */
async function bandEntryOf(
  options: Pick<BisAnswerOptions, 'build' | 'fetchFile'>,
  specKey: string,
  character: MeCharacter,
): Promise<BisBand | undefined> {
  const file = await (options.fetchFile ?? fetchBisFile)(options.build, specKey);
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
  const entry = await bandEntryOf(options, specKey, character);
  if (entry === undefined) return withoutList;
  return countAnswer(entry, character.build?.gear ?? {}, specKey.split('-')[0]);
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
export function landingAnswerSpec(
  config: LandingConfig,
  fetchFile?: BisAnswerOptions['fetchFile'],
): AnswerSpec {
  return bisAnswerSpec({
    build: config.build,
    ...(fetchFile === undefined ? {} : { fetchFile }),
    specKeyOf: (character) =>
      specKeyFor(config, character.class, characterSpecName(character)) ?? levelingSpecKey(character),
  });
}

/** The character's own spec page at its own band and faction; the pointer to the grid for an unknown class. */
export function landingAction(character: MeCharacter, config: LandingConfig): HeaderAction {
  const cls = classFor(config, character.class);
  // No spec yet (a fresh import): the class's leveling spec, never a "pick a spec" prompt.
  const specKey =
    specKeyFor(config, character.class, characterSpecName(character)) ?? levelingSpecKey(character);
  const spec = cls === undefined ? undefined : config.specs[specKey ?? cls.defaultSpecKey];
  if (cls === undefined || spec === undefined) {
    return { label: landingCopy.fallbackAction, href: landingCopy.fallbackHref };
  }
  const band = characterBand(character.level) ?? DEFAULT_LANDING_BAND;
  return {
    label: landingCopy.leadButton(spec.name, band),
    href: specLinkHref({
      classSlug: spec.classSlug,
      specSlug: spec.specSlug,
      band,
      faction: character.faction,
    }),
  };
}
