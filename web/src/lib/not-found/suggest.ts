// web/src/lib/not-found/suggest.ts
// The 404's "Did you mean" (pages spec 2026-11-08 decision N3): the failed path against a
// build-time list of the site's pages (the footer's pages, every guide, class landing and BiS
// spec page), the nearest one within an edit distance of 3. `SUGGEST_SCRIPT` is the same
// search as an inline script, run before the class picker paints so the line never shifts it.
import { FOOTER_COLUMNS, FOOTER_LEGAL_LINKS } from '../footer-copy';
import { notFoundCopy } from './copy';

/** A known page: its path and the name the suggestion shows. */
export type KnownPage = readonly [path: string, title: string];

/** The furthest a mistyped path may be from a page and still be suggested. */
export const MAX_SUGGEST_DISTANCE = 3;

export interface SpecPage {
  classSlug: string;
  className: string;
  specSlug: string;
  specName: string;
}

/** The footer's pages (no fragments), then each class landing, spec guide and BiS spec page. */
export function knownPages(
  classes: readonly { slug: string; name: string }[],
  specs: readonly SpecPage[],
): KnownPage[] {
  const footer = [...FOOTER_COLUMNS.flatMap((column) => column.links), ...FOOTER_LEGAL_LINKS]
    .filter((link) => !link.href.includes('#'))
    .map((link): KnownPage => [link.href, link.label]);
  return [
    ...footer,
    ...classes.map((c): KnownPage => [`/guides/${c.slug}`, notFoundCopy.classGuideTitle(c.name)]),
    ...specs.flatMap((s): KnownPage[] => [
      [`/guides/${s.classSlug}/${s.specSlug}`, notFoundCopy.guideTitle(s.specName, s.className)],
      [`/bis/${s.classSlug}/${s.specSlug}`, notFoundCopy.bisTitle(s.specName, s.className)],
    ]),
  ];
}

/** Levenshtein distance, stopping early past `limit` (returns limit + 1). */
export function editDistance(a: string, b: string, limit: number = MAX_SUGGEST_DISTANCE): number {
  if (Math.abs(a.length - b.length) > limit) return limit + 1;
  let previous = Array.from({ length: b.length + 1 }, (_, i) => i);
  for (let i = 1; i <= a.length; i += 1) {
    const current = [i];
    for (let j = 1; j <= b.length; j += 1) {
      current[j] = Math.min(
        previous[j]! + 1,
        current[j - 1]! + 1,
        previous[j - 1]! + (a[i - 1] === b[j - 1] ? 0 : 1),
      );
    }
    if (Math.min(...current) > limit) return limit + 1;
    previous = current;
  }
  return previous[b.length]!;
}

/** The page nearest `path` (case and a trailing slash ignored) within the limit, or null. */
export function suggestPage(path: string, pages: readonly KnownPage[]): KnownPage | null {
  const wanted = path.toLowerCase().replace(/\/+$/, '') || '/';
  let best: KnownPage | null = null;
  let bestDistance = MAX_SUGGEST_DISTANCE + 1;
  for (const page of pages) {
    const distance = editDistance(wanted, page[0], MAX_SUGGEST_DISTANCE);
    if (distance < bestDistance && page[0] !== path) {
      best = page;
      bestDistance = distance;
    }
  }
  return best;
}

/**
 * The inline script: `editDistance` and `suggestPage` serialised from this module, run on the
 * pages JSON in `#nf-pages`; it fills `#nf-path` with the failed address and, on a match, fills
 * and shows `#nf-suggest` (its link and name).
 */
export const SUGGEST_SCRIPT = `(function(){var MAX_SUGGEST_DISTANCE=${MAX_SUGGEST_DISTANCE};var editDistance=${editDistance.toString()};var suggestPage=${suggestPage.toString()};
var p=location.pathname,m=document.getElementById('nf-path');if(m)m.textContent=p;
if(p.indexOf('/character/')===0||p==='/character')document.documentElement.setAttribute('data-nf-character','');
var d=document.getElementById('nf-pages'),s=document.getElementById('nf-suggest');if(!d||!s)return;
var hit=suggestPage(p,JSON.parse(d.textContent||'[]'));if(!hit)return;var a=s.querySelector('a');if(a){a.href=hit[0];a.textContent=hit[1]+' \\u203a';}s.hidden=false;})();`;
