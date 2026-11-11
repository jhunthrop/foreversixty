// web/src/lib/class-crest.ts
// The one place a class crest's image is chosen (component audit 2026-11-10 cluster 2, brand audit
// 2026-11-10 hand-off D): every class or character image on the web is the circular ringed HD crest,
// drawn by ui/Crest (Svelte), ui/Crest.astro (static pages) or, for DOM-built rows, `crestImage`
// here with styles/crest.css's `.crest` class. A character with no known class draws the LX seal
// inside a muted ring, never a person glyph or a broken request.
//
// The crests are 128 px lossy WebP (4-7 KB each); the PNGs are gone from public/.
import { classSlugFromName } from './report/format';
import { sealFor } from './ui/assets';
import { classColor, isClassSlug, type ClassSlug } from './ui/primitives';

/** The answer boxes' crest (ui/CharacterAnswer `.ui-cans-crest`), painted by the BiS and guide mounts. */
export const ANSWER_CREST_PX = 40;

/** The one class every crest element carries (styles/crest.css). */
export const CREST_CLASS = 'crest';

export function classCrestSrc(slug: ClassSlug): string {
  return `/icons/hd/crests/${slug}.webp`;
}

/** A crest is drawn for every playable class (public/icons/hd/crests/, class-crest.test.ts). */
export function hasClassCrest(slug: string): slug is ClassSlug {
  return isClassSlug(slug);
}

/** A class name ("Death Knight", "Warrior") or slug ("warrior") as a crest slug; null for none. */
export function crestSlug(cls: string | null | undefined): ClassSlug | null {
  const slug = classSlugFromName(cls ?? '');
  return hasClassCrest(slug) ? slug : null;
}

/** What one crest draws: the image, its ring colour, and whether it is a class or the LX fallback. */
export interface CrestImage {
  readonly src: string;
  readonly ring: string;
  readonly known: boolean;
}

/** The neutral ring around the LX fallback: visible on the row, never a class colour. */
const NEUTRAL_RING = 'var(--color-muted)';

/** The crest for a class name or slug at a drawn size; an unknown class the LX seal in a muted ring. */
export function crestImage(cls: string | null | undefined, size: number): CrestImage {
  const slug = crestSlug(cls);
  return slug === null
    ? { src: sealFor(size), ring: NEUTRAL_RING, known: false }
    : { src: classCrestSrc(slug), ring: classColor(slug), known: true };
}

/** Paints a DOM-built crest `<img>` (the typeahead, the answer boxes): src, ring and fallback flag. */
export function paintCrest(img: HTMLImageElement, cls: string | null | undefined, size: number): void {
  const image = crestImage(cls, size);
  img.src = image.src;
  img.style.setProperty('--c', image.ring);
  img.classList.add(CREST_CLASS);
  img.classList.toggle('crest-none', !image.known);
}
