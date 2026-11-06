// web/src/lib/home/example.ts
// The signed-out home's example (home spec 2026-11-06 H9): one class at every band, read at build
// time from the same BiS file /bis serves. One row per band for the Alliance list (the bare set
// first, then raid-ready at 60): the band's pick for one slot, its source and the set's DPS.
// Node only (lib/bis/load reads the disk).
import { itemDetails, loadBisFile } from '../bis/load';
import { presetOf } from '../bis/presets';
import type { BisBand } from '../bis/types';
import type { ItemRef } from '../ui/primitives';
import { slotSourceLine } from './source-line';

/** The fixture-backed spec, so the parity gate runs on CI's data (spec H9). */
export const HOME_EXAMPLE = {
  spec: 'hunter-marksmanship',
  classSlug: 'hunter',
  name: 'Marksmanship Hunter',
  slot: 'ranged',
  faction: 'alliance',
  bisHref: '/bis/hunter/marksmanship',
} as const;

const MAX_BAND = 60;
const BAND_SPAN = 9;

export interface ExampleRow {
  key: string;
  label: string;
  item: ItemRef;
  source: string;
  setDps: string;
}

export interface HomeExample {
  title: string;
  rows: ExampleRow[];
}

function bandLabel(band: BisBand): string {
  if (band.band < MAX_BAND) return `${band.band} to ${band.band + BAND_SPAN}`;
  return presetOf(band) === 'raid' ? `${MAX_BAND} raid` : String(MAX_BAND);
}

/** Null when the build has no list for the example spec. */
export function homeExample(
  build: string,
  title: (first: number, last: number) => string,
): HomeExample | null {
  const file = loadBisFile(HOME_EXAMPLE.spec, build);
  if (file === null) return null;
  const items = itemDetails(build, HOME_EXAMPLE.classSlug);
  const bands = file.bands
    .filter((band) => band.faction === HOME_EXAMPLE.faction)
    .sort((a, b) => a.band - b.band || Number(presetOf(a) !== 'bare') - Number(presetOf(b) !== 'bare'));
  const rows = bands.flatMap((band) => {
    const pick = band.slots.find((slot) => slot.slot === HOME_EXAMPLE.slot && slot.item_id !== undefined);
    if (pick === undefined) return [];
    const detail = items.get(pick.item_id);
    return [
      {
        key: `${band.band}-${presetOf(band)}`,
        label: bandLabel(band),
        item: { id: pick.item_id, name: pick.item_name, quality: detail?.quality ?? 1, icon: detail?.icon },
        source: slotSourceLine(pick),
        setDps: band.set_dps.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 }),
      },
    ];
  });
  if (rows.length === 0) return null;
  const levels = bands.map((band) => band.band);
  return { title: title(Math.min(...levels), Math.max(...levels)), rows };
}
