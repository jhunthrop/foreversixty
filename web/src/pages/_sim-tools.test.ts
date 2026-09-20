// web/src/pages/_sim-tools.test.ts
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { beforeAll, describe, expect, it } from 'vitest';
import Gear from './sim/gear.astro';
import Talents from './sim/talents.astro';
import Drops from './sim/drops.astro';
import Weights from './sim/weights.astro';
import { TOOL_SKELETONS } from '../lib/sim/bulk-skeleton';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

const OG_HOOKS = [
  'data-og="title"',
  'data-og="description"',
  'data-og="canonical"',
  'data-og="og-title"',
  'data-og="og-description"',
  'data-og="og-url"',
  'data-og="og-image"',
];

const PAGES = [
  { name: 'gear', component: Gear, path: '/sim/gear', heading: 'Top Gear' },
  { name: 'talents', component: Talents, path: '/sim/talents', heading: 'Talent compare' },
  { name: 'drops', component: Drops, path: '/sim/drops', heading: 'Droptimizer' },
  { name: 'weights', component: Weights, path: '/sim/weights', heading: 'Stat weights' },
] as const;

describe.each(PAGES)('the $name shell', ({ component, path, heading, name }) => {
  it('carries every data-og hook and its own canonical', async () => {
    const html = await container.renderToString(component);
    for (const selector of OG_HOOKS) expect(html, selector).toContain(selector);
    expect(html).toContain(`href="https://foreversixty.gg${path}"`);
  });

  it('mounts the tools island, stamps its tool, and links one stylesheet and one script', async () => {
    const html = await container.renderToString(component);
    expect(html).toContain('id="sim-tools"');
    expect(html).toContain(`data-sim-tool="${name}"`);
    expect(html).toContain('href="/sim-tools-island.css"');
    expect(html).toContain('src="/sim-tools-island.js"');
  });

  it('renders its heading and its own skeleton before the island, so nothing shifts', async () => {
    const html = await container.renderToString(component);
    expect(html).toContain(heading);
    expect(html).toContain(TOOL_SKELETONS[name].slice(0, 80));
  });

  it('says what the page needs when there is no JavaScript', async () => {
    const html = await container.renderToString(component);
    expect(html).toContain('needs JavaScript');
    expect(html).toContain('href="/sim"');
  });
});
