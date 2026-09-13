// web/astro.config.mjs
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import pagefind from 'astro-pagefind';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  site: 'https://foreversixty.gg',
  trailingSlash: 'never',
  build: { format: 'file' },
  integrations: [svelte(), sitemap(), pagefind()],
  vite: { plugins: [tailwindcss()] },
});
