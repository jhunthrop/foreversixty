// web/src/lib/bis/band-path.ts
// Every BiS level band has its own server-rendered URL, so a search engine (and a shared link)
// reaches each band's list without running the page: /bis/<class>/<spec> is the spec's first band
// (its canonical), /bis/<class>/<spec>/<slug> each other one (`30-39`, `40-49`, `50-59`, `60`).
// The one source for those paths, their slugs and the words that name a band in a title or h1.
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

/** `/bis/<class>/<spec>`: the spec's own page, its first band. */
export function bisSpecPath(classSlug: string, specSlug: string): string {
  return `/bis/${classSlug}/${specSlug}`;
}

/** A band's page: the spec's own path for its first band, else `/bis/<class>/<spec>/<slug>`. */
export function bisBandPath(classSlug: string, specSlug: string, band: number, firstBand: number): string {
  const base = bisSpecPath(classSlug, specSlug);
  return band === firstBand ? base : `${base}/${bandSlug(band)}`;
}

/** "Levels 40 to 49", or "Level 60". */
export function bandWords(band: number): string {
  return band >= MAX_BAND ? `Level ${MAX_BAND}` : `Levels ${band} to ${band + 9}`;
}
