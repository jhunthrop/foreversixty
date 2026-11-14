// web/src/lib/bis/page-registry.ts
// How the BiS island gets its data without carrying it as island props (which Astro would
// serialize a second time into the island's attribute): on the server the page registers the
// data before render and the island takes it from here; in the browser the island parses the
// one `<script type="application/json">` the page rendered beside it.
import type { ItemTooltipModel } from '../items/tooltip';
import type { BisPageData } from './page-model';

export const BIS_PAGE_DATA_ID = 'bis-page-data';
/** The page's item tooltips (BisTooltipHost and the paperdoll's centre read the same script). */
export const BIS_TOOLTIP_MODELS_ID = 'bis-tooltip-models';

export type TooltipModels = ReadonlyMap<number, ItemTooltipModel>;

const registry = new Map<string, BisPageData>();
const tooltipRegistry = new Map<string, TooltipModels>();
/** In the browser: the tooltips the BiS island took with the full data after mount. One module
 *  instance serves every island on the page, so BisTooltipHost sees them too. */
const lateTooltips = new Map<number, ItemTooltipModel>();

/** Adds tooltip models a page learned after its first paint (the BiS island's full data). */
export function addTooltipModels(models: Readonly<Record<string, ItemTooltipModel>>): TooltipModels {
  for (const [id, model] of toModelMap(models)) lateTooltips.set(id, model);
  return lateTooltips;
}

/** A tooltip model added after the first paint, or undefined. */
export function lateTooltipModel(id: number): ItemTooltipModel | undefined {
  return lateTooltips.get(id);
}

export function registerBisPageData(data: BisPageData): void {
  registry.set(data.spec, data);
}

/** The page's data, on the server from the registry, in the browser from the page's JSON. */
export function readBisPageData(spec: string, doc: Document | undefined = globalThis.document): BisPageData {
  if (doc === undefined) {
    const data = registry.get(spec);
    if (data === undefined) throw new Error(`BiS page data for ${spec} was not registered before render`);
    return data;
  }
  const node = doc.getElementById(BIS_PAGE_DATA_ID);
  if (node?.textContent == null) throw new Error(`#${BIS_PAGE_DATA_ID} is missing from the page`);
  return JSON.parse(node.textContent) as BisPageData;
}

/** The JSON the page inlines, safe inside a script element. */
export function bisPageDataJson(data: BisPageData): string {
  return JSON.stringify(data).replace(/</g, '\\u003c');
}

/** Registers the page's tooltip models for the server render (the centre shows the selected slot's). */
export function registerBisTooltips(spec: string, models: Readonly<Record<string, ItemTooltipModel>>): void {
  tooltipRegistry.set(spec, toModelMap(models));
}

function toModelMap(models: Readonly<Record<string, ItemTooltipModel>>): TooltipModels {
  return new Map(Object.entries(models).map(([id, model]) => [Number(id), model]));
}

/** The tooltip models: on the server from the registry, in the browser from the page's JSON
 *  script (`id`); empty when there is none or it does not parse (a tooltip never breaks a page). */
export function readTooltipModels(
  spec: string | undefined,
  id: string = BIS_TOOLTIP_MODELS_ID,
  doc: Document | undefined = globalThis.document,
): TooltipModels {
  if (doc === undefined) return (spec === undefined ? undefined : tooltipRegistry.get(spec)) ?? new Map();
  const text = doc.getElementById(id)?.textContent;
  if (text == null || text === '') return new Map();
  try {
    return toModelMap(JSON.parse(text) as Record<string, ItemTooltipModel>);
  } catch {
    return new Map();
  }
}
