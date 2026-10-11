// web/src/lib/bis/landing-bands.ts
// The /bis landing's band strip as pure data: the five bands, their labels, what the picks at
// each band are mostly made of (the pill captions, computed at build time from the BiS files so
// prose never quotes a nightly number), and which band the strip has selected.
import { BAND_STEP, MAX_BAND, MIN_BAND } from './hover';
import type { BisSlot } from './types';

/** Signed out, or with no character: launch players are leveling, so the first band. */
export const DEFAULT_LANDING_BAND = MIN_BAND;

export const LANDING_BANDS: readonly number[] = Array.from(
  { length: (MAX_BAND - MIN_BAND) / BAND_STEP + 1 },
  (_, index) => MIN_BAND + index * BAND_STEP,
);

/** The player-facing word for each `source_kind` the BiS files carry. */
const SOURCE_KIND_WORDS: Readonly<Record<string, string>> = {
  dungeon: 'dungeons',
  crafted: 'crafting',
  vendor: 'vendors',
  quest: 'quests',
  world_drop: 'world drops',
  world: 'world drops',
  rep: 'reputation rewards',
  pvp: 'PvP rewards',
};

const TOP_SOURCE_COUNT = 2;

/** `20 to 29`; the top band is plain `60`, never "60 to 60". */
export function bandLabel(band: number): string {
  return band >= MAX_BAND ? String(MAX_BAND) : `${band} to ${band + BAND_STEP - 1}`;
}

/** A band number from untrusted text (`?band=`, storage); null unless it is one of the strip's. */
export function parseBandParam(raw: string | null | undefined): number | null {
  if (raw === null || raw === undefined || !/^\d+$/.test(raw)) return null;
  const band = Number(raw);
  return LANDING_BANDS.includes(band) ? band : null;
}

/**
 * The selected band: the viewer's own choice (the URL, else the stored one) wins; otherwise the
 * character's band; otherwise the default. The character's band is undefined under level 20.
 */
export function selectedLandingBand(explicit: number | null, characterBand: number | undefined): number {
  return explicit ?? characterBand ?? DEFAULT_LANDING_BAND;
}

/** The two source kinds most picks come from, most common first (ties keep first-seen order). */
export function topSourceWords(slots: readonly Pick<BisSlot, 'source_kind'>[]): string[] {
  const counts = new Map<string, number>();
  for (const slot of slots) {
    const word = SOURCE_KIND_WORDS[slot.source_kind ?? ''];
    if (word !== undefined) counts.set(word, (counts.get(word) ?? 0) + 1);
  }
  return [...counts.entries()]
    .sort((a, b) => b[1] - a[1])
    .slice(0, TOP_SOURCE_COUNT)
    .map(([word]) => word);
}

export interface BandCaption {
  band: number;
  label: string;
  /** `Top sources: dungeons, crafting.` where the sources differ from the band before (the first
   *  band always); `Same sources as 20 to 29` where they match. */
  text: string;
  /** True for the `Same sources as` form, which renders muted. */
  same: boolean;
}

/** One caption per band; `slotsByBand` is every written spec's picks at that band, pooled. */
export function bandCaptions(
  slotsByBand: ReadonlyMap<number, readonly Pick<BisSlot, 'source_kind'>[]>,
): BandCaption[] {
  const captions: BandCaption[] = [];
  let previous: { band: number; words: string } | undefined;
  for (const band of LANDING_BANDS) {
    const words = topSourceWords(slotsByBand.get(band) ?? []).join(', ');
    const same = previous !== undefined && previous.words === words;
    captions.push({
      band,
      label: bandLabel(band),
      same,
      text: same ? `Same sources as ${bandLabel(previous.band)}` : `Top sources: ${words}.`,
    });
    previous = { band, words };
  }
  return captions;
}
