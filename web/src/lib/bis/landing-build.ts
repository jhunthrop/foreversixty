// web/src/lib/bis/landing-build.ts
// Build-time data for the /bis landing (node only: it reads the BiS files off disk, like load.ts):
// the class cards in game order with their spec rows, the band strip's captions, the raid preset's
// label for the strip's note, the newest `generated_at` for the one Updated stamp, and the config
// the page's browser script reads. A spec with no ranked list yet keeps its row -- its own page
// carries the empty state -- rather than vanishing from the index.
import classes from '../../data/classes.json';
import { dataUrl } from '../planner/load';
import { SPECS } from '../sim/specs';
import { latest } from '../dates';
import { LANDING_BANDS, bandCaptions, type BandCaption } from './landing-bands';
import type { LandingConfig } from './landing-config';
import { bandEntry, loadBisFile } from './load';
import type { BisFile, BisSlot } from './types';

export interface LandingSpecRow {
  key: string;
  classSlug: string;
  specSlug: string;
  name: string;
  role: (typeof SPECS)[number]['role'];
  iconSrc: string;
}

export interface LandingClassCard {
  slug: string;
  name: string;
  color: string;
  specs: LandingSpecRow[];
}

export interface LandingData {
  cards: LandingClassCard[];
  captions: BandCaption[];
  /** The raid preset's own label as a sentence continues it ("raid-ready, Phase 1"); empty when no
   *  file carries one. */
  raidLabel: string;
  /** The newest file's `generated_at`, or null when no spec has a file. */
  updated: Date | null;
  config: LandingConfig;
}

function lowerFirst(text: string): string {
  return text.charAt(0).toLowerCase() + text.slice(1);
}

function cardsFor(build: string): LandingClassCard[] {
  return classes
    .map((cls) => ({
      slug: cls.slug,
      name: cls.name,
      color: cls.color,
      specs: SPECS.filter((spec) => spec.class_slug === cls.slug).map((spec): LandingSpecRow => ({
        key: spec.spec,
        classSlug: spec.class_slug,
        specSlug: spec.spec_slug,
        name: spec.name,
        role: spec.role,
        iconSrc: dataUrl(build, `icons/${spec.icon}.webp`),
      })),
    }))
    .filter((card) => card.specs.length > 0);
}

function slotsByBand(files: readonly BisFile[]): Map<number, BisSlot[]> {
  return new Map(
    LANDING_BANDS.map((band) => [
      band,
      files.flatMap((file) => bandEntry(file, band, 'alliance')?.slots ?? []),
    ]),
  );
}

function configFor(build: string, cards: readonly LandingClassCard[]): LandingConfig {
  return {
    build,
    classes: Object.fromEntries(
      cards.map((card) => [
        card.slug,
        { slug: card.slug, name: card.name, color: card.color, defaultSpecKey: card.specs[0]!.key },
      ]),
    ),
    specs: Object.fromEntries(
      cards.flatMap((card) =>
        card.specs.map((spec) => [
          spec.key,
          { classSlug: spec.classSlug, specSlug: spec.specSlug, name: spec.name },
        ]),
      ),
    ),
  };
}

export function loadLandingData(build: string): LandingData {
  const cards = cardsFor(build);
  const files = cards
    .flatMap((card) => card.specs)
    .map((spec) => loadBisFile(spec.key, build))
    .filter((file): file is BisFile => file !== null);
  return {
    cards,
    captions: bandCaptions(slotsByBand(files)),
    raidLabel: lowerFirst(
      files.map((file) => file.presets?.raid?.label).find((label) => label !== undefined) ?? '',
    ),
    updated: files.length === 0 ? null : latest(files.map((file) => new Date(file.generated_at))),
    config: configFor(build, cards),
  };
}
