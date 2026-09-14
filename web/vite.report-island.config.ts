// web/vite.report-island.config.ts
// The report island's bundle: dist/report-island.js and dist/report-island.css, written
// next to the Astro output. The report shells are static HTML the Worker clones for every
// report id, so they link one fixed url rather than a hashed Astro chunk.
import { defineConfig } from 'vite';
import { islandConfig } from './vite.island.config';

export default defineConfig(islandConfig('report-island'));
