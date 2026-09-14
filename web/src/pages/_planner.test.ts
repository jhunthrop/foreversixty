// web/src/pages/_planner.test.ts
// Underscore-prefixed because everything else under src/pages/ is a route: Astro would
// otherwise build this file as the endpoint /planner.test and fail the build on its
// top-level vitest calls. Astro skips `_`-prefixed files; vitest still collects it.
//
// The Svelte renderer has to be handed to the container explicitly: Astro's integrations
// are not loaded in a unit test, so without this the island renders as an empty shell.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { beforeAll, describe, expect, it } from 'vitest';
import Planner from './planner.astro';
import activeBuild from '../data/active-build.json';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

describe('planner.astro', () => {
  it('renders the shell with the active build id and the Era-data notice', async () => {
    const html = await container.renderToString(Planner);
    expect(html).toContain('Build planner');
    expect(html).toContain(activeBuild.build);
    expect(html).toContain('Classic Era trees shown until the beta client exports');
    expect(html).toContain('href="https://foreversixty.gg/planner"');
  });

  it('mounts exactly one island', async () => {
    const html = await container.renderToString(Planner);
    expect(html.match(/<astro-island/g) ?? []).toHaveLength(1);
  });
});
