// web/src/pages/_sim.test.ts
// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: sim.astro mounts the Account island through
// Base's `session` prop, and Astro's integrations are not loaded in a unit test, so
// without this the container throws NoMatchingRenderer rather than rendering an empty
// shell. sim/specs.astro takes no `session` prop and would pass either way; it shares the
// container for one setup rather than two.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import type { AstroComponentFactory } from 'astro/runtime/server/index.js';
import { createPageContainer } from '../test-support/page-container';
import { readFileSync } from 'node:fs';
import { beforeAll, describe, expect, it } from 'vitest';
import { VIEW_GAP } from '../lib/current-character-layout';
import { RESERVE_STORAGE_KEY } from '../lib/sim/server-runs-poll';
import { FIXTURE_SIM_ID } from '../lib/report/shell-paths';
import Sim from './sim.astro';
import SimById from './sim/[id].astro';
import Specs from './sim/specs.astro';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
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
    // Loaded by IslandScript.astro after the first paint, so no plain module tag.
    expect(html).toContain('data-island="/sim-island.js"');
    expect(html).not.toContain('src="/sim-island.js"');
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

// Fix round 1, Important #2: SimView.svelte's hydrated root is `flex flex-col gap-[22px]
// md:gap-8`, so each shell wraps its blocks in that identical gap or everything paints higher
// before hydration than after. 2026-10-31: the "Jump to" spine bar is gone, so no shell
// reserves a band for it any more.
describe('the shells’ gap wrapper', () => {
  const wrapper = `<div class="flex flex-col ${VIEW_GAP}">`;

  it('sim.astro wraps its blocks in the view’s own flex gap, with no spine bar slot', async () => {
    const html = await container.renderToString(Sim);
    expect(html).toContain(wrapper);
    expect(html).not.toContain('current-character-bar');
  });

  it('sim/specs.astro wraps the grids in the identical gap, with no spine bar slot', async () => {
    const html = await container.renderToString(Specs);
    expect(html).toContain(wrapper);
    expect(html).not.toContain('current-character-bar');
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
  it('sim/[id].astro wraps the saved-sim skeleton in the identical gap, with no spine bar slot', async () => {
    const html = await container.renderToString(SimById as unknown as AstroComponentFactory, {
      params: { id: FIXTURE_SIM_ID },
    });
    expect(html).toContain(wrapper);
    expect(html).not.toContain('current-character-bar');
  });
});

// Experience spec 2026-10-24 section 3.7: a payer's "Your server runs" box is reserved before
// the island paints, and only for a hinted payer.
describe('the server runs reserve', () => {
  it('sim.astro opens its view with a hidden reserve box, ahead of the paste hero', async () => {
    const html = await container.renderToString(Sim);
    const wrapper = html.indexOf(`<div class="flex flex-col ${VIEW_GAP}">`);
    const reserve = html.indexOf('server-runs-shell-reserve');
    expect(reserve).toBeGreaterThan(wrapper);
    expect(reserve).toBeLessThan(wrapper + 400);
  });

  it('the stylesheet shows the shell box only under the session and server-sims hints', () => {
    const css = readFileSync(new URL('../styles/global.css', import.meta.url), 'utf8');
    expect(css).toMatch(/\.server-runs-shell-reserve\s*\{\s*display: none;/);
    expect(css).toMatch(
      /html\[data-session='1'\]\[data-server-sims='1'\] \.server-runs-shell-reserve\s*\{\s*display: block;/,
    );
  });

  it('the pre-paint script stamps data-server-sims from the cached /v1/me and the stored list height', async () => {
    const html = await container.renderToString(Sim);
    expect(html).toContain("root.dataset.serverSims = '1'");
    expect(html).toContain('entitlements.server_sims === true');
    expect(html).toContain(`data-server-runs-reserve-key="${RESERVE_STORAGE_KEY}"`);
  });
});
