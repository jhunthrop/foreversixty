// web/vite.island.config.ts
// A second Vite build, writing dist/planner-island.js and dist/planner-island.css next to
// the Astro output (emptyOutDir: false). The Astro build cannot produce this: its island
// chunks are hashed and referenced only from its own HTML, while the API needs one stable,
// unhashed URL.
import { svelte } from '@sveltejs/vite-plugin-svelte';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';

/**
 * Both standalone islands are built the same way and differ only in their entry name, so
 * the config is a function. The planner's bundle is loaded by the Go API's /b/:id page and
 * the report's by the Worker-served report shells; both need one stable, unhashed URL,
 * which the Astro build cannot produce because its island chunks are hashed and referenced
 * only from its own HTML.
 */
export function islandConfig(name: 'planner-island' | 'report-island' | 'sim-island' | 'sim-tools-island') {
  return {
    // Astro exposes PUBLIC_*; plain Vite exposes VITE_* only, so PUBLIC_API_BASE_URL has to be
    // allow-listed or lib/planner/config.ts would silently fall back to the default origin.
    envPrefix: ['VITE_', 'PUBLIC_'],
    plugins: [svelte(), tailwindcss()],
    build: {
      outDir: 'dist',
      emptyOutDir: false,
      cssCodeSplit: false,
      sourcemap: false,
      // Nothing here is dynamically imported, so the entry needs no preload polyfill.
      // sim-tools-island is the exception: it lazily imports one view chunk per tool page,
      // so /sim/weights never downloads the slot grid. The chunks take
      // `${name}-[hash].js`, which is already the chunkFileNames rule below.
      modulePreload: false,
      // A plain entry with pinned output names rather than `build.lib`. Library mode inlines
      // every referenced asset as a base64 data URL whatever `assetsInlineLimit` says, and
      // src/styles/fonts.css references ten font files: it turned planner-island.css from 26 kB
      // into 307 kB of render-blocking base64 that no browser can cache separately. This emits
      // the faces as files under /assets/ instead, which is also what lets the `src:` URLs stay
      // root-absolute and resolve from https://foreversixty.gg/planner-island.css.
      rollupOptions: {
        input: `src/${name}.ts`,
        output: {
          format: 'es' as const,
          entryFileNames: `${name}.js`,
          chunkFileNames: `${name}-[hash].js`,
          // cssCodeSplit is off, so there is exactly one stylesheet and it takes the fixed
          // name. Everything else it pulls in -- the woff2/woff faces -- stays content-hashed.
          assetFileNames: (asset: { names: string[] }) =>
            asset.names.some((n) => n.endsWith('.css')) ? `${name}.css` : 'assets/[name]-[hash][extname]',
        },
      },
    },
  };
}

export default defineConfig(islandConfig('planner-island'));
