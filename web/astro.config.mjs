// web/astro.config.mjs
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import pagefind from 'astro-pagefind';
import tailwindcss from '@tailwindcss/vite';
import links from './src/data/links.json';
import { assertLinksAreReal } from './src/lib/links';

// Registers `client:interaction` (see src/directives/interaction.ts): hydrates the homepage
// search island on the user's first focus/`/`-press instead of during initial page load, so
// its JS never competes with the mobile LCP budget in lighthouserc.json.
const interactionDirective = {
  name: 'interaction-directive',
  hooks: {
    'astro:config:setup': ({ addClientDirective }) => {
      addClientDirective({ name: 'interaction', entrypoint: './src/directives/interaction.ts' });
    },
  },
};

// CF_PAGES is set only inside a Cloudflare Pages build, so local and CI builds keep working
// against the placeholder community links while a deploy fails fast and names the file.
const placeholderGuard = {
  name: 'placeholder-guard',
  hooks: {
    'astro:config:setup': () => {
      if (process.env.CF_PAGES) assertLinksAreReal(links);
    },
  },
};

export default defineConfig({
  site: 'https://foreversixty.gg',
  trailingSlash: 'never',
  build: { format: 'file', inlineStylesheets: 'always' },
  integrations: [
    svelte(),
    // /b-unavailable is the 503 fallback src/worker.ts serves when the API cannot answer a
    // /b/:id request, and /reports is the shell the Worker clones for every report id.
    // Both are plumbing rather than destinations, so they stay out of the sitemap; the
    // report pages people actually link to are /reports/<id>, which cannot be enumerated
    // at build time.
    sitemap({ filter: (page) => !page.endsWith('/b-unavailable') && !page.endsWith('/reports') }),
    pagefind(),
    interactionDirective,
    placeholderGuard,
  ],
  vite: {
    plugins: [tailwindcss()],
    // astro-pagefind writes /pagefind/pagefind.js into dist/ *after* the bundle is generated,
    // so the search island's dynamic import of it can never be resolved at build time. Marking
    // the path external tells Rollup to emit the import untouched in both the client and the
    // SSR bundle instead of failing with UNRESOLVED_IMPORT.
    build: { rollupOptions: { external: [/^\/pagefind\//] } },
  },
});
