// web/src/lib/class-crest.ts
// The one place a class crest's image path is built (home rebuild spec §5/§6):
// `ClassCrest.astro` renders it for every Astro page, and a Svelte island (which cannot
// import a `.astro` file the way an Astro page can -- Vite's Astro integration only
// compiles `.astro` files reached from another `.astro` entry point, never from a
// framework component) inlines the identical `<img>` itself, built from this one path
// function plus `classColorVar` for the ring colour, rather than a second, drifting copy.
//
// Review round 2 (index LCP still over web/lighthouserc.json's 2200ms budget after the
// round 1 256px -> 128px PNG resize): WebP, lossy quality 90 with alpha, resampled to the
// same 128px from the original scratchpad 256px sources, cuts every crest to 4-7KB (was
// 17-30KB as PNG at the same 128px). The PNGs are gone from public/ entirely -- this is
// the only path any caller reads.
export function classCrestSrc(slug: string): string {
  return `/icons/hd/crests/${slug}.webp`;
}
