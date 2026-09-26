// web/src/pages/_sim.test.ts
// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: sim.astro mounts the Account island through
// Base's `session` prop, and Astro's integrations are not loaded in a unit test, so
// without this the container throws NoMatchingRenderer rather than rendering an empty
// shell. sim/specs.astro takes no `session` prop and would pass either way; it shares the
// container for one setup rather than two.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import type { AstroComponentFactory } from 'astro/runtime/server/index.js';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { beforeAll, describe, expect, it } from 'vitest';
import { VIEW_GAP } from '../lib/current-character-layout';
import { FIXTURE_SIM_ID } from '../lib/report/shell-paths';
import Sim from './sim.astro';
import SimById from './sim/[id].astro';
import Specs from './sim/specs.astro';

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

describe('the sim shell', () => {
  it('carries every data-og hook src/worker.ts rewrites', async () => {
    const html = await container.renderToString(Sim);
    for (const selector of OG_HOOKS) expect(html, selector).toContain(selector);
  });

  it('mounts the island and links its one stylesheet and one script', async () => {
    const html = await container.renderToString(Sim);
    expect(html).toContain('id="sim"');
    expect(html).toContain('data-sim-mount');
    expect(html).toContain('href="/sim-island.css"');
    expect(html).toContain('src="/sim-island.js"');
  });

  it('says what the page needs and offers the planner when there is no JavaScript', async () => {
    const html = await container.renderToString(Sim);
    expect(html).toContain('runs the engine in your browser');
    expect(html).toContain('href="/planner"');
  });

  it('ships no Astro island for the simulator itself, so nothing competes with the standalone bundle', async () => {
    const html = await container.renderToString(Sim);
    // The one <astro-island> present is Base's `session` slot mounting Account
    // client:load, the same header account widget every other `session` page (e.g.
    // logs.astro) carries -- unrelated to the simulator, which is mounted only by
    // sim-island.js below. A second island here would mean SimView got mounted twice.
    expect(html.match(/<astro-island/g) ?? []).toHaveLength(1);
  });
});

describe('the spec support shell', () => {
  it('carries the hooks, the mount and its own canonical', async () => {
    const html = await container.renderToString(Specs);
    for (const selector of OG_HOOKS) expect(html, selector).toContain(selector);
    expect(html).toContain('href="https://foreversixty.gg/sim/specs"');
    expect(html).toContain('data-sim-view="specs"');
  });
});

// Fix round 1, Important #2: the static shell used to stack the chip slot directly
// against the next element (0 gap, a plain non-flex #sim), while SimView.svelte's own
// hydrated root is `flex flex-col gap-[22px] md:gap-8` -- so everything below the chip
// painted 22px/32px higher before hydration than after. Both shells now wrap the slot and
// what follows it in that identical gap, ahead of the chip slot in the rendered markup.
describe('the chip slot’s gap wrapper', () => {
  it('sim.astro wraps the chip slot and the LCP card in the view’s own flex gap', async () => {
    const html = await container.renderToString(Sim);
    const wrapper = `<div class="flex flex-col ${VIEW_GAP}">`;
    expect(html).toContain(wrapper);
    expect(html.indexOf(wrapper)).toBeLessThan(html.indexOf('data-testid="current-character-bar"'));
  });

  it('sim/specs.astro wraps the chip slot and the grids in the identical gap', async () => {
    const html = await container.renderToString(Specs);
    const wrapper = `<div class="flex flex-col ${VIEW_GAP}">`;
    expect(html).toContain(wrapper);
    expect(html.indexOf(wrapper)).toBeLessThan(html.indexOf('data-testid="current-character-bar"'));
  });

  // sim/[id].astro is the page this repo actually measures with Lighthouse (the Round 3
  // comment on that file), so the same invariant is asserted for it too. Its own
  // `getStaticPaths` only yields a path under `FOREVER_DATA=fixture`, but the container
  // API renders the component directly from `params` -- the same object `getStaticPaths`
  // would have supplied -- without going through routing or that env var at all. The cast
  // below is only for `astro check`: a dynamic route's default export types its `props`
  // parameter from its own `getStaticPaths` return, which this route never uses (it reads
  // `Astro.params`, not `Astro.props`) and which `astro check` cannot narrow from the
  // conditional `FOREVER_DATA` check inside `fixtureSimPaths()` -- so it falls back to
  // `never`. `renderToString` never passes a `props` value for this route, so the cast
  // changes nothing about what actually renders.
  it('sim/[id].astro wraps the chip slot and the saved-sim skeleton in the identical gap', async () => {
    const html = await container.renderToString(SimById as unknown as AstroComponentFactory, {
      params: { id: FIXTURE_SIM_ID },
    });
    const wrapper = `<div class="flex flex-col ${VIEW_GAP}">`;
    expect(html).toContain(wrapper);
    expect(html.indexOf(wrapper)).toBeLessThan(html.indexOf('data-testid="current-character-bar"'));
  });
});
