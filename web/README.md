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

| Command              | Action                                            |
| :------------------- | :------------------------------------------------ |
| `npm install`        | Install dependencies                              |
| `npm run dev`        | Start the local dev server                        |
| `npm run build`      | Build the production site to `./dist/`            |
| `npm run preview`    | Preview the production build locally              |
| `npm run check`      | Type-check with `astro check` (TypeScript strict) |
| `npm run test`       | Run unit tests once with Vitest                   |
| `npm run test:watch` | Run unit tests in watch mode                      |
| `npm run test:e2e`   | Run end-to-end tests with Playwright              |
| `npm run lhci`       | Run Lighthouse CI                                 |

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

## Deploy (Cloudflare Pages)

1. Cloudflare dashboard → Workers & Pages → Create → Pages → Connect to Git → this repo.
2. Build settings: framework preset Astro; root directory `web`; build command `npm run build`; output directory `dist`; Node version env `NODE_VERSION=22`. A `.node-version` file (`22.12.0`) is also checked in at `web/.node-version` — Cloudflare Pages reads it automatically, and the `NODE_VERSION` env var is the fallback if it doesn't.
3. Custom domains: add `foreversixty.gg` (apex) and `www.foreversixty.gg`.
4. Enable Web Analytics on the Pages project (dashboard toggle; no script is added to the repo).

Before the first deploy, replace the `PLACEHOLDER` values in `src/data/links.json`. A build with `CF_PAGES` set (that is, a Cloudflare Pages build) fails and names the offending keys if any are left; local and CI builds are unaffected.

## Redirect foreversixty.com → foreversixty.gg

1. Add `foreversixty.com` as a zone in the same Cloudflare account; point registrar nameservers at Cloudflare.
2. DNS: `A @ 192.0.2.1` proxied and `CNAME www foreversixty.com` proxied (placeholders so the proxy answers).
3. Rules → Redirect Rules → create: when hostname matches `foreversixty.com` or `www.foreversixty.com`, redirect 301 to `https://foreversixty.gg${uri}` preserving path and query.

## Verify

```
curl -sI https://foreversixty.com/dungeons | grep -i location   → https://foreversixty.gg/dungeons
curl -s https://foreversixty.gg/robots.txt                        → Sitemap line present
```
