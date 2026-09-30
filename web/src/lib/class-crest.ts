// web/src/lib/class-crest.ts
// The one place a class crest's image path is built (home rebuild spec §5/§6):
// `ClassCrest.astro` renders it for every Astro page, and a Svelte island (which cannot
// import a `.astro` file the way an Astro page can -- Vite's Astro integration only
// compiles `.astro` files reached from another `.astro` entry point, never from a
// framework component) inlines the identical `<img>` itself, built from this one path
// function plus `classColorVar` for the ring colour, rather than a second, drifting copy.
export function classCrestSrc(slug: string): string {
  return `/icons/hd/crests/${slug}.png`;
}
