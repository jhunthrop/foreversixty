// web/astro.config.mjs
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import { readFileSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
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

// Reads `updated: YYYY-MM-DD` out of the content collections without loading Astro's
// content layer, which is not available inside the config. Path → date, for the sitemap.
const contentLastmod = (() => {
  const root = fileURLToPath(new URL('./src/content', import.meta.url));
  const routeOf = { pages: '', guides: '/guides', zones: '/zones', dungeons: '/dungeons' };
  const byPath = new Map();
  let newest = '';
  for (const [dir, prefix] of Object.entries(routeOf)) {
    for (const file of readdirSync(join(root, dir))) {
      if (!file.endsWith('.md')) continue;
      const match = readFileSync(join(root, dir, file), 'utf8').match(/^updated:\s*'?(\d{4}-\d{2}-\d{2})/m);
      if (!match) continue;
      byPath.set(`${prefix}/${file.slice(0, -3)}`, match[1]);
      if (match[1] > newest) newest = match[1];
    }
  }
  for (const file of readdirSync(join(root, 'changelog'))) {
    const match = readFileSync(join(root, 'changelog', file), 'utf8').match(
      /^updated:\s*'?(\d{4}-\d{2}-\d{2})/m,
    );
    if (match && match[1] > newest) newest = match[1];
  }
  byPath.set('/', newest);
  byPath.set('/changelog', newest);
  byPath.set('/everything-we-know', byPath.get('/everything-we-know') ?? newest);
  return (url) => byPath.get(new URL(url).pathname.replace(/\/$/, '') || '/');
})();

export default defineConfig({
  site: 'https://foreversixty.gg',
  trailingSlash: 'never',
  build: { format: 'file', inlineStylesheets: 'always' },
  integrations: [
    svelte(),
    // /b-unavailable is the 503 fallback src/worker.ts serves when the API cannot answer a
    // /b/:id request. It is internal plumbing rather than a destination, so it stays out of
    // the sitemap; the page also carries its own noindex for a crawler that finds it anyway.
    sitemap({
      filter: (page) => !page.endsWith('/b-unavailable'),
      // lastmod comes from each content file's `updated` field, so a page's sitemap entry
      // moves only when its facts do; pages without a content file carry no lastmod.
      serialize: (item) => {
        const lastmod = contentLastmod(item.url);
        return lastmod ? { ...item, lastmod } : item;
      },
    }),
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
