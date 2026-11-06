// web/src/lib/bis/page-data.ts
// Server side of the BiS page (web layout spec 2026-11-05 §5): reads the spec's BiS file and the
// build's item, loot, instance and enchant tables once, and builds the serializable page data the
// island renders (`BisPageData`: one compact model per faction × band × preset, plus the item
// dictionary those models name). The raw BiS file never ships to the browser.
import { readItemTooltipModel } from '../items/lookup-node';
import type { ItemTooltipModel } from '../items/tooltip';
import { readBuildJson } from '../guides/build-files';
import type { TalentIndex } from '../planner/rules';
import { bandLevels, instanceLevels, itemDetails, loadLootFile } from './load';
import { bandModelFor, bandModelKey, itemDictionary, modelItemIds, type BisBandModel, type BisPageData, type EnchantNames, type GainUnit } from './page-model';
import { bandInfosFor } from './panel-view';
import { bandPresets, selectBand } from './presets';
import { bisSimHref, rankingSimHrefFor } from './sim-link';
import { slotScoreUnitFor } from './tank-view';
import type { BisBand, BisFile, Faction } from './types';

export const BIS_FACTIONS: readonly Faction[] = ['alliance', 'horde'];

interface EnchantRecord {
  id: number;
  name: string;
  slots: string[];
}

export function enchantNamesFor(build: string): EnchantNames {
  const records = readBuildJson<EnchantRecord[]>(build, 'enchants.json') ?? [];
  const byKey = new Map<string, string>();
  for (const record of records) {
    for (const slot of record.slots) byKey.set(`${record.id}:${slot}`, record.name);
  }
  return (id, slot) => byKey.get(`${id}:${slot}`) ?? byKey.get(`${id}:${slot.replace(/[12]$/, '')}`);
}

function unitOf(file: BisFile): GainUnit {
  const first = file.bands[0];
  if (first === undefined) return 'DPS';
  const unit = slotScoreUnitFor(first);
  return unit === 'tank_score' ? 'score' : unit === 'hps' ? 'HPS' : 'DPS';
}

export interface BisPageDataInput {
  spec: string;
  classSlug: string;
  build: string;
  file: BisFile;
  talentIndex: TalentIndex;
}

export interface BisServerPage {
  data: BisPageData;
  tooltips: Record<number, ItemTooltipModel>;
}

/** Every model the page offers, the item dictionary, and the tooltip models for those items. */
export function buildBisPageData(input: BisPageDataInput): BisServerPage {
  const { spec, classSlug, build, file, talentIndex } = input;
  const details = itemDetails(build, classSlug);
  const loot = loadLootFile(build);
  const tooltipFor = (id: number) => readItemTooltipModel(build, classSlug, id) ?? undefined;
  const bands = bandLevels(file);
  const unit = unitOf(file);
  const deps = {
    file,
    details,
    loot,
    instanceLevels: instanceLevels(),
    enchantName: enchantNamesFor(build),
    unit,
    simHrefFor: (band: BisBand) => bisSimHref(band, talentIndex, build),
    rankingSimHrefFor: (band: BisBand) => rankingSimHrefFor(band, talentIndex, build),
    plannerHrefFor: (band: BisBand) => `/planner?spec=${spec}&talents=${encodeURIComponent(band.talents ?? '')}`,
  };
  const models: Record<string, BisBandModel> = {};
  for (const faction of BIS_FACTIONS) {
    const presets = new Set(bands.flatMap((band) => bandPresets(file, band, faction)));
    for (const preset of presets) {
      const infos = bandInfosFor(file, bands, faction, { itemDetails: details, loot, tooltipFor, spec, preset });
      for (const info of infos) {
        const raw = selectBand(file, info.band, faction, preset);
        if (raw === undefined || (raw.preset ?? 'bare') !== preset) continue;
        models[bandModelKey(faction, info.band, preset)] = bandModelFor(info, raw, deps);
      }
    }
  }
  const names = new Map<number, string>();
  for (const band of file.bands) for (const slot of band.slots) if (typeof slot.item_id === 'number') names.set(slot.item_id, slot.item_name);
  const ids = modelItemIds(Object.values(models));
  const tooltips: Record<number, ItemTooltipModel> = {};
  for (const id of ids) {
    const model = tooltipFor(id);
    if (model !== undefined) tooltips[id] = model;
  }
  return {
    data: { spec, build, bands, models, items: itemDictionary(ids, details, (id) => names.get(id)) },
    tooltips,
  };
}
