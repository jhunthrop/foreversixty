# Forever Sixty — Web

A fan reference site for World of Warcraft: Forever, built with [Astro](https://astro.build) (static output), [Svelte 5](https://svelte.dev) islands, and [Tailwind CSS 4](https://tailwindcss.com).

## Stack

- **Astro** — static site generation, `build.format: 'file'`
- **Svelte 5** — interactive islands
- **Tailwind CSS 4** — via the `@tailwindcss/vite` plugin
- **astro-pagefind** — static full-text search
- **@astrojs/sitemap** — sitemap generation

## Commands

All commands are run from this directory (`web/`):

| Command                | Action                                            |
| :--------------------- | :------------------------------------------------ |
| `npm install`          | Install dependencies                              |
| `npm run dev`          | Start the local dev server                        |
| `npm run build`        | Build the production site to `./dist/`            |
| `npm run preview`      | Preview the production build locally              |
| `npm run check`        | Type-check with `astro check` (TypeScript strict) |
| `npm run lint`         | Lint with ESLint (Astro, Svelte, TypeScript)      |
| `npm run lint:fix`     | Lint and apply the fixable rules                  |
| `npm run format`       | Format with Prettier                              |
| `npm run format:check` | Check formatting without writing                  |
| `npm run test`         | Run unit tests once with Vitest                   |
| `npm run test:watch`   | Run unit tests in watch mode                      |
| `npm run test:e2e`     | Run end-to-end tests with Playwright              |
| `npm run lhci`         | Run Lighthouse CI                                 |

## Project Structure

```text
web/
├── public/               # static assets served as-is (favicon, robots.txt)
├── src/
│   ├── components/       # Astro components and their co-located unit tests
│   ├── content/          # Markdown content collections (changelog, dungeons, guides, pages, zones)
│   ├── data/             # static JSON data (dates, tools, classes, community links, unknowns)
│   ├── directives/       # custom client directives (e.g. `client:interaction`)
│   ├── layouts/          # shared page layouts (Base, Content)
│   ├── lib/              # framework-agnostic helpers (dates, sources, OG image generation)
│   ├── pages/            # file-based routes, including dynamic `[slug]` and OG image routes
│   ├── styles/           # global CSS and design tokens
│   └── content.config.ts # content collection schemas (zod)
├── tests/e2e/             # Playwright end-to-end specs
├── lighthouserc.json      # Lighthouse CI budget
├── playwright.config.ts
├── astro.config.mjs
├── tsconfig.json
└── package.json
```

## Deploy (Cloudflare Workers, static assets)

The site is served by Cloudflare Workers as static assets with no Worker code; `wrangler.jsonc` holds the
config (asset directory `dist`, `auto-trailing-slash` so `dungeons.html` answers at `/dungeons`, the
`404.html` page for unknown paths, and the two custom domains `foreversixty.gg` and `www.foreversixty.gg`).
Cloudflare Pages was the original plan; it has since been folded into Workers and new projects cannot be
created through the API, so the site deploys with wrangler instead.

- Every push to `main` that touches `web/**` runs the verify job and then `wrangler deploy` (see
  `.github/workflows/web.yml`). It needs two repository secrets: `CLOUDFLARE_API_TOKEN` (Workers Scripts
  Edit, Zone DNS Edit for the custom domains) and `CLOUDFLARE_ACCOUNT_ID`.
- Manual deploy from this directory: `CLOUDFLARE_API_TOKEN=... npm run deploy` (`npm run build` then
  `wrangler deploy`). The workers.dev preview is https://foreversixty.jhunthrop.workers.dev.
- Node is pinned by `web/.node-version` (`22.12.0`); CI reads it with `node-version-file`.

Community links (Discord invite, GitHub repo, issue templates) live in `src/data/links.json`. The deploy
job builds with `CF_PAGES=1`, which makes the build fail and name the offending keys if any value still
contains `PLACEHOLDER`; local and CI verify builds are unaffected.

## Redirect foreversixty.com → foreversixty.gg

1. Add `foreversixty.com` as a zone in the same Cloudflare account; point the Namecheap nameservers at Cloudflare (done 2026-09-13: coby and lady .ns.cloudflare.com).
2. DNS: `A @ 192.0.2.1` proxied and `CNAME www foreversixty.com` proxied (placeholders so the proxy answers).
3. Rules → Redirect Rules → create: when hostname matches `foreversixty.com` or `www.foreversixty.com`, redirect 301 to `https://foreversixty.gg${uri}` preserving path and query.

## Verify

```
curl -sI https://foreversixty.com/dungeons | grep -i location   → https://foreversixty.gg/dungeons
curl -s https://foreversixty.gg/robots.txt                        → Sitemap line present
```
