// web/vite.island.config.ts
// A second Vite build, in library mode, writing dist/planner-island.js and
// dist/planner-island.css next to the Astro output (emptyOutDir: false). The Astro build
// cannot produce this: its island chunks are hashed and referenced only from its own HTML,
// while the API needs one stable, unhashed URL.
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
    lib: {
      entry: 'src/planner-island.ts',
      formats: ['es'],
      fileName: () => 'planner-island.js',
      cssFileName: 'planner-island',
    },
  },
});
