// web/vite.sim-island.config.ts
import { defineConfig } from 'vite';
import { islandConfig } from './vite.island.config';

export default defineConfig(islandConfig('sim-island'));
