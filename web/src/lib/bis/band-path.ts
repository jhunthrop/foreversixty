// web/src/lib/bis/band-path.ts
// The BiS level bands' URLs. /bis/<class>/<spec> is the spec's page: level 60 on the doll (owner
// 2026-11-17: open at 60 when there is no character) and the leveling bands (20-29 to 50-59)
// folded into it as a server-rendered Levelling section.
// Since 2026-11-15 those bands have no page of their own (the SEO report of 2026-11-10: fewer,
// stronger pages): a link opens one as the spec page with a band hash (`#levels-40-49`), which the
// island reads to show that band on the doll, and their old URLs 301 there (public/_redirects). A
// fragment is no new URL to a crawler, so every such link counts toward the spec page. The hash is
// no element's id, so the browser never scrolls past the doll. Level 60 keeps its own page,
// /bis/<class>/<spec>/60. The one source for those paths, hashes, slugs and band words.
import { MAX_BAND } from './band-levels';

/** The band's path segment: `60`, or `40-49` below it. */
export function bandSlug(band: number): string {
  return band >= MAX_BAND ? String(MAX_BAND) : `${band}-${band + 9}`;
}

/** The band a path segment names, or null for anything that is not a band slug. */
export function bandFromSlug(slug: string): number | null {
  if (slug === String(MAX_BAND)) return MAX_BAND;
  const match = /^(\d{2})-(\d{2})$/.exec(slug);
  if (match === null) return null;
  const band = Number(match[1]);
  return band % 10 === 0 && band < MAX_BAND && Number(match[2]) === band + 9 ? band : null;
}

/** `/bis/<class>/<spec>`: the spec's own page, opening on its home band (specPageBand). */
export function bisSpecPath(classSlug: string, specSlug: string): string {
  return `/bis/${classSlug}/${specSlug}`;
}

/** The band /bis/<class>/<spec> opens on with nothing in the URL: level 60, else (a spec ranked
 *  short of 60) its first band. */
export function specPageBand(bands: readonly number[]): number | undefined {
  return bands.includes(MAX_BAND) ? MAX_BAND : bands[0];
}

/** Whether the band lives on the spec page's Levelling section: every band below 60. */
export function isFoldedBand(band: number): boolean {
  return band < MAX_BAND;
}

/** A folded band's hash on the spec page, `#levels-40-49`. */
export function bandHash(band: number): string {
  return `#levels-${bandSlug(band)}`;
}

const BAND_HASH = /^#levels-(\d{2})-(\d{2})$/;

/** The band a `#levels-40-49` hash names, or null for any other hash. */
export function bandFromHash(hash: string): number | null {
  const match = BAND_HASH.exec(hash);
  if (match === null) return null;
  const band = bandFromSlug(`${match[1]}-${match[2]}`);
  return band !== null && band < MAX_BAND ? band : null;
}

/** The page that holds a band: `/bis/<class>/<spec>/60` for level 60 (unless the spec's first band
 *  is 60, when there is no /60 page), else the spec's own path. */
export function bisBandPath(classSlug: string, specSlug: string, band: number, firstBand: number): string {
  const base = bisSpecPath(classSlug, specSlug);
  return band >= MAX_BAND && band !== firstBand ? `${base}/${bandSlug(band)}` : base;
}

/** A link that opens a band: its page, the query (`?faction=horde`) and a folded band's hash. */
export function bisBandHref(
  classSlug: string,
  specSlug: string,
  band: number,
  firstBand: number,
  query = '',
): string {
  const hash = isFoldedBand(band) ? bandHash(band) : '';
  return `${bisBandPath(classSlug, specSlug, band, firstBand)}${query}${hash}`;
}

/** "Levels 40 to 49", or "Level 60". */
export function bandWords(band: number): string {
  return band >= MAX_BAND ? `Level ${MAX_BAND}` : `Levels ${band} to ${band + 9}`;
}
