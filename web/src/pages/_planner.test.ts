// web/src/pages/_planner.test.ts
// Underscore-prefixed because everything else under src/pages/ is a route: Astro would
// otherwise build this file as the endpoint /planner.test and fail the build on its
// top-level vitest calls. Astro skips `_`-prefixed files; vitest still collects it.
//
// The Svelte renderer has to be handed to the container explicitly: Astro's integrations
// are not loaded in a unit test, so without this the island renders as an empty shell.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { createPageContainer } from '../test-support/page-container';
import { beforeAll, describe, expect, it } from 'vitest';
import Planner from './planner.astro';
import activeBuild from '../data/active-build.json';
import { loadTalentFile } from '../lib/guides/build-tree';
import { DEFAULT_CLASS_SLUG } from '../lib/planner/config';
import { plannerCopy } from '../lib/planner/copy';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

describe('planner.astro', () => {
  it('renders the shell with the active build id and where its trees came from', async () => {
    const html = await container.renderToString(Planner);
    expect(html).toContain('WoW Forever Talent Calculator');
    expect(html).toContain(activeBuild.build);
    // No treeSourceNotice assertion here: Task 5 moved that paragraph inside the
    // `{:else if store.talentIndex}` ready branch, which only exists once Planner.svelte's
    // own `$effect` has populated `store.talentIndex` -- and, same as Guild.test.ts and
    // CharacterRatingPanel.test.ts document for their own components, svelte/server's
    // render() (which AstroContainer.renderToString uses under the hood) never runs
    // `$effect` at all, so this render never leaves the pre-`$effect` loading state.
    expect(html).toContain('href="https://foreversixty.gg/planner"');
  });

  it('names the opening class’s trees at 0 points and says what the planner does, before the island', async () => {
    const html = await container.renderToString(Planner);
    expect(html).toContain(plannerCopy.lead);
    const trees = loadTalentFile(activeBuild.build, DEFAULT_CLASS_SLUG)?.trees ?? [];
    expect(trees.length).toBeGreaterThan(0);
    for (const tree of trees) expect(html).toMatch(new RegExp(`<h2[^>]*>${tree.name}</h2>`));
  });

  it('mounts exactly two islands', async () => {
    // Planner.svelte, and the header's session link: planner.astro passes Base's `session`
    // prop, because a page that already runs the planner can afford to show who is signed in.
    const html = await container.renderToString(Planner);
    expect(html.match(/<astro-island/g) ?? []).toHaveLength(2);
  });
});
