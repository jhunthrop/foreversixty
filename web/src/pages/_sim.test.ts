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
describe('the shells', () => {
  // Spec 2026-11-06-sim §4.2: the static shell is the signed-out landing itself (the head, the
  // paste Plate, the best-in-slot rail), the island's first render of the same component.
  it('sim.astro renders the landing: the head, the paste box and the best-in-slot rail, with no spine bar', async () => {
    const html = await container.renderToString(Sim);
    expect(html).toContain('data-region="head"');
    expect(html).toContain('id="sim-cold-paste-input"');
    expect(html).toContain('data-region="section-sim-a-best-in-slot-set"');
    expect(html).toContain('data-bis-sets=');
    expect(html).not.toContain('current-character-bar');
  });

  it('sim/specs.astro wraps the grids in the view gap, with no spine bar slot', async () => {
    const html = await container.renderToString(Specs);
    expect(html).toContain(`<div class="flex flex-col ${VIEW_GAP}">`);
    expect(html).not.toContain('current-character-bar');
  });

  // The cast is only for `astro check`: the route types its props from `getStaticPaths`, which
  // it never reads (it reads `Astro.params`).
  it('sim/[id].astro renders a prerendered run itself, and the skeleton for any other id', async () => {
    const render = (id: string) =>
      container.renderToString(SimById as unknown as AstroComponentFactory, { params: { id } });
    const fixture = await render(FIXTURE_SIM_ID);
    expect(fixture).toContain('data-testid="sim-saved"');
    expect(fixture).toContain('data-sim-notes=');
    expect(fixture).not.toContain('data-sim-tab');
    expect(fixture).not.toContain('current-character-bar');
    const other = await render('abcdefghijkl');
    expect(other).toContain('data-testid="sim-saved-skeleton"');
  });
});

// Experience spec 2026-10-24 section 3.7: a payer's "Your server runs" box is reserved before
// the island paints, and only for a hinted payer.
describe('the server runs reserve', () => {
  it('sim.astro reserves the hidden box under the head, ahead of the paste box', async () => {
    const html = await container.renderToString(Sim);
    const head = html.indexOf('data-region="head"');
    const reserve = html.indexOf('server-runs-shell-reserve');
    expect(reserve).toBeGreaterThan(head);
    expect(reserve).toBeLessThan(html.indexOf('id="sim-cold-paste-input"'));
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
