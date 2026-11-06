// web/src/lib/bis/page-registry.ts
// How the BiS island gets its data without carrying it as island props (which Astro would
// serialize a second time into the island's attribute): on the server the page registers the
// data before render and the island takes it from here; in the browser the island parses the
// one `<script type="application/json">` the page rendered beside it.
import type { BisPageData } from './page-model';

export const BIS_PAGE_DATA_ID = 'bis-page-data';

const registry = new Map<string, BisPageData>();

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
