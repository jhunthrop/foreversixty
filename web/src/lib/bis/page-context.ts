// web/src/lib/bis/page-context.ts
// What the BiS page's state may take (page-state.ts's BisStateContext), read off the page data:
// the bands, the presets each faction and band publishes, the races per faction and the race
// each band was simmed as.
import { bisCopy } from './copy';
import { bandModelKey } from './model-key';
import type { BisPageData } from './page-model';
import type { BisStateContext } from './page-state';
import type { BisPresetId, Faction } from './types';

const PRESETS: readonly BisPresetId[] = ['raid', 'bare'];

export function bisContext(
  data: Pick<BisPageData, 'bands' | 'models'>,
  raceOptions: Readonly<Record<Faction, readonly { value: string }[]>>,
): BisStateContext {
  const presetsFor = (faction: Faction, band: number) => presetsOffered(data, faction, band);
  const racesFor = (faction: Faction) => raceOptions[faction].map((option) => option.value);
  return {
    bands: data.bands,
    presetsFor,
    racesFor,
    simRaceFor: (faction, band, preset) =>
      data.models[bandModelKey(faction, band, preset)]?.race ?? racesFor(faction)[0] ?? '',
  };
}

/** The presets a faction and band publish, raid-ready first. */
export function presetsOffered(data: Pick<BisPageData, 'models'>, faction: Faction, band: number): BisPresetId[] {
  return PRESETS.filter((preset) => data.models[bandModelKey(faction, band, preset)] !== undefined);
}

/** The Profile control's words. */
export function presetLabel(preset: BisPresetId): string {
  return preset === 'raid' ? bisCopy.presetOptionRaid : bisCopy.presetOptionBare;
}
