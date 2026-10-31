// web/astro.config.mjs
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import { createReadStream, readFileSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import tailwindcss from '@tailwindcss/vite';
import links from './src/data/links.json';
import { assertLinksAreReal } from './src/lib/links';
import { LOCAL_GUILD_SHELL_PATH } from './src/lib/report/shell-paths';
import { DUCKDB_WASM_VERSION, RUNTIME_MODULES, stagingDir } from './scripts/duckdb-runtime.mjs';

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

// The local walkthrough stack (tools/dev-stack) has no Worker to serve the guild shell for a
// real guild's URL, so under FS_LOCAL_GUILD_SHELL=1 (set only by scripts/dev-local-api.mjs)
// every /guild/<anything> request is rewritten, server side, to the one prerendered shell
// path. The browser's address bar keeps the real URL, which is what the shell parses, as it
// does behind the Worker in production. Off by default, so no build or other dev run changes.
const localGuildShell = {
  name: 'local-guild-shell',
  hooks: {
    'astro:server:setup': ({ server }) => {
      if (process.env.FS_LOCAL_GUILD_SHELL !== '1') return;
      server.middlewares.use((request, _response, next) => {
        const url = request.url ?? '';
        const [pathname, query] = url.split('?');
        if (pathname.startsWith('/guild/') && pathname !== `/guild/${LOCAL_GUILD_SHELL_PATH}`) {
          request.url = `/guild/${LOCAL_GUILD_SHELL_PATH}${query === undefined ? '' : `?${query}`}`;
        }
        next();
      });
    },
  },
};

// A `client:idle-after-load` directive alongside Astro's built-in `client:idle` (see
// src/directives/idle-after-load.ts for why): it defers a component's fetch and hydration
// until after the window `load` event, then idle, instead of idle alone. Registered as its
// own directive rather than replacing `client:idle` everywhere, so only the components that
// were actually racing the home page's own paint (day-3 home-lcp lane report) opt in.
const idleAfterLoadDirective = {
  name: 'idle-after-load-directive',
  hooks: {
    'astro:config:setup': ({ addClientDirective }) => {
      addClientDirective({
        name: 'idle-after-load',
        entrypoint: './src/directives/idle-after-load.ts',
      });
    },
  },
};

// A `client:after-paint` directive (src/directives/after-paint.ts): hydrates once the first
// contentful paint has happened. The guild page opts in so its server-rendered lede is the
// largest contentful paint on every run.
const afterPaintDirective = {
  name: 'after-paint-directive',
  hooks: {
    'astro:config:setup': ({ addClientDirective }) => {
      addClientDirective({ name: 'after-paint', entrypoint: './src/directives/after-paint.ts' });
    },
  },
};

const webRoot = fileURLToPath(new URL('.', import.meta.url));
const contractFixtures = fileURLToPath(new URL('../docs/contracts/fixtures', import.meta.url));

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
  const routeOf = { pages: '' };
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
  return (url) => byPath.get(new URL(url).pathname.replace(/\/$/, '') || '/');
})();

export default defineConfig({
  site: 'https://foreversixty.gg',
  trailingSlash: 'never',
  build: { format: 'file', inlineStylesheets: 'always' },
  integrations: [
    svelte(),
    // /b-unavailable is the 503 fallback src/worker.ts serves when the API cannot answer a
    // /b/:id request, and /report-shell is the shell the Worker clones for every report id.
    // Both are plumbing rather than destinations, so they stay out of the sitemap; the
    // report pages people actually link to are /reports/<id>, which cannot be enumerated
    // at build time. /reports, the signed-in "Your reports" list, has nothing to index either.
    sitemap({
      filter: (page) =>
        !page.endsWith('/b-unavailable') &&
        !page.endsWith('/report-shell') &&
        !page.endsWith('/reports') &&
        !page.endsWith('/design-sheet'),
      // lastmod comes from each content file's `updated` field, so a page's sitemap entry
      // moves only when its facts do; pages without a content file carry no lastmod.
      serialize: (item) => {
        const lastmod = contentLastmod(item.url);
        return lastmod ? { ...item, lastmod } : item;
      },
    }),
    idleAfterLoadDirective,
    afterPaintDirective,
    placeholderGuard,
    duckdbRuntime,
    localGuildShell,
  ],
  vite: {
    plugins: [tailwindcss()],
    // The Forever Score's labels and tier colours are the addon's, read from the contract
    // fixture both sides pin (docs/contracts/fixtures/score-labels.json, lib/score/). A build
    // bundles it like any import; `astro dev` serves it to the browser only from an allowed
    // directory, and Vite's default allows nothing above web/.
    server: { fs: { allow: [webRoot, contractFixtures] } },
  },
});
