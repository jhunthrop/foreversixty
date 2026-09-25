// web/astro.config.mjs
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import { createReadStream, readFileSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import pagefind from 'astro-pagefind';
import tailwindcss from '@tailwindcss/vite';
import links from './src/data/links.json';
import { assertLinksAreReal } from './src/lib/links';
import { DUCKDB_WASM_VERSION, RUNTIME_MODULES, stagingDir } from './scripts/duckdb-runtime.mjs';

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

// The two DuckDB engine modules are over Cloudflare's static-asset limit, so they are not
// in public/ and `astro dev` has nothing to serve them with: in production src/worker.ts
// answers /duckdb-runtime/<version>/<file> out of the LOGS bucket. This serves the same
// two paths straight off the staging directory scripts/sync-duckdb.mjs fills, so the
// Queries view works in `npm run dev`. `astro preview` cannot be extended this way -- it
// strips user Vite plugins -- so the Playwright suite stubs the same route instead
// (tests/e2e/support/duckdb-runtime.ts).
const duckdbRuntime = {
  name: 'duckdb-runtime',
  hooks: {
    'astro:server:setup': ({ server }) => {
      server.middlewares.use((request, response, next) => {
        const pathname = (request.url ?? '').split('?')[0];
        const file = RUNTIME_MODULES.find(
          (name) => pathname === `/duckdb-runtime/${DUCKDB_WASM_VERSION}/${name}`,
        );
        if (file === undefined) {
          next();
          return;
        }
        response.setHeader('content-type', 'application/wasm');
        createReadStream(join(stagingDir, file)).pipe(response);
      });
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
  const routeOf = { pages: '', zones: '/zones', dungeons: '/dungeons' };
  const byPath = new Map();
  let newest = '';
  const record = (routePath, fileContents) => {
    const match = fileContents.match(/^updated:\s*'?(\d{4}-\d{2}-\d{2})/m);
    if (!match) return;
    byPath.set(routePath, match[1]);
    if (match[1] > newest) newest = match[1];
  };
  for (const [dir, prefix] of Object.entries(routeOf)) {
    for (const file of readdirSync(join(root, dir))) {
      if (!file.endsWith('.md')) continue;
      record(`${prefix}/${file.slice(0, -3)}`, readFileSync(join(root, dir, file), 'utf8'));
    }
  }
  // guides/ is one directory per class (src/content/guides/<class>/{index,<spec>}.md):
  // index.md is the class landing page at /guides/<class>, every other file is a spec guide
  // at /guides/<class>/<spec>.
  for (const classSlug of readdirSync(join(root, 'guides'), { withFileTypes: true })) {
    if (!classSlug.isDirectory()) continue;
    for (const file of readdirSync(join(root, 'guides', classSlug.name))) {
      if (!file.endsWith('.md')) continue;
      const stem = file.slice(0, -3);
      const routePath = stem === 'index' ? `/guides/${classSlug.name}` : `/guides/${classSlug.name}/${stem}`;
      record(routePath, readFileSync(join(root, 'guides', classSlug.name, file), 'utf8'));
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
    // /b/:id request, and /reports is the shell the Worker clones for every report id.
    // Both are plumbing rather than destinations, so they stay out of the sitemap; the
    // report pages people actually link to are /reports/<id>, which cannot be enumerated
    // at build time.
    sitemap({
      filter: (page) => !page.endsWith('/b-unavailable') && !page.endsWith('/reports'),
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
    duckdbRuntime,
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
