// web/vite.island.config.ts
// A second Vite build, writing dist/planner-island.js and dist/planner-island.css next to
// the Astro output (emptyOutDir: false). The Astro build cannot produce this: its island
// chunks are hashed and referenced only from its own HTML, while the API needs one stable,
// unhashed URL.
import { svelte } from '@sveltejs/vite-plugin-svelte';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';

export default defineConfig({
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
    modulePreload: false,
    // A plain entry with pinned output names rather than `build.lib`. Library mode inlines
    // every referenced asset as a base64 data URL whatever `assetsInlineLimit` says, and
    // src/styles/fonts.css references ten font files: it turned planner-island.css from 26 kB
    // into 307 kB of render-blocking base64 that no browser can cache separately. This emits
    // the faces as files under /assets/ instead, which is also what lets the `src:` URLs stay
    // root-absolute and resolve from https://foreversixty.gg/planner-island.css.
    rollupOptions: {
      input: 'src/planner-island.ts',
      output: {
        format: 'es',
        entryFileNames: 'planner-island.js',
        chunkFileNames: 'planner-island-[hash].js',
        // cssCodeSplit is off, so there is exactly one stylesheet and it takes the fixed
        // name. Everything else it pulls in -- the woff2/woff faces -- stays content-hashed.
        assetFileNames: (asset) =>
          asset.names.some((name) => name.endsWith('.css'))
            ? 'planner-island.css'
            : 'assets/[name]-[hash][extname]',
      },
    },
  },
});
