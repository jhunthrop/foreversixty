// web/astro.config.mjs
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import pagefind from 'astro-pagefind';
import tailwindcss from '@tailwindcss/vite';

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

export default defineConfig({
  site: 'https://foreversixty.gg',
  trailingSlash: 'never',
  build: { format: 'file', inlineStylesheets: 'always' },
  integrations: [svelte(), sitemap(), pagefind(), interactionDirective],
  vite: { plugins: [tailwindcss()] },
});
