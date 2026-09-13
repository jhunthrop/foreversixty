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

| Command             | Action                                        |
| :------------------- | :--------------------------------------------- |
| `npm install`         | Install dependencies                          |
| `npm run dev`         | Start the local dev server                    |
| `npm run build`       | Build the production site to `./dist/`        |
| `npm run preview`     | Preview the production build locally          |
| `npm run check`       | Type-check with `astro check` (TypeScript strict) |
| `npm run test`        | Run unit tests once with Vitest               |
| `npm run test:watch`  | Run unit tests in watch mode                  |
| `npm run test:e2e`    | Run end-to-end tests with Playwright          |
| `npm run lhci`        | Run Lighthouse CI                             |

## Project Structure

```text
web/
├── public/            # static assets served as-is (favicon, etc.)
├── src/
│   └── pages/         # file-based routes
├── astro.config.mjs
├── tsconfig.json
└── package.json
```
