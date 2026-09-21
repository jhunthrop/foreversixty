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
import { treeSourceNotice } from '../lib/planner/tree-source';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

describe('planner.astro', () => {
  it('renders the shell with the active build id and where its trees came from', async () => {
    const html = await container.renderToString(Planner);
    expect(html).toContain('Build planner');
    expect(html).toContain(activeBuild.build);
    // Computed rather than a literal string, so this keeps passing once Task 12 flips
    // active-build.json off the pre-beta snapshot: treeSourceNotice's other branch renders
    // it a "read from the game client" build id, and the assertion follows it there too.
    expect(html).toContain(treeSourceNotice(activeBuild.build));
    expect(html).toContain('href="https://foreversixty.gg/planner"');
  });

  it('mounts exactly two islands', async () => {
    // Planner.svelte, and the header's session link: planner.astro passes Base's `session`
    // prop, because a page that already runs the planner can afford to show who is signed in.
    const html = await container.renderToString(Planner);
    expect(html.match(/<astro-island/g) ?? []).toHaveLength(2);
  });
});
