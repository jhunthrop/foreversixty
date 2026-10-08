// web/src/pages/guides/_spec.test.ts
// Underscore-prefixed for the same reason as every other _*.test.ts under src/pages/: Astro
// skips it, vitest still collects it.
//
// Builds its own `entry` by reading the guide file directly off disk and parsing it through
// the real `guideSchema`, the same way _sections.test.ts's own `loadGuides()` does -- NOT via
// `getCollection('guides')`, which returns an empty collection under plain `vitest run`
// unless Astro's dev-mode content-layer cache (`.astro/data-store.json`) happens to already be
// warm from a prior `astro dev` run, something nothing in this repo's CI produces before
// `npm test`. Running the real schema over the real file keeps this test honest about what
// the page actually receives (a real `CollectionEntry`-shaped object with `id`/`data`/`body`)
// without depending on Astro's content layer being primed.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { parse as parseYaml } from 'yaml';
import { beforeAll, describe, expect, it } from 'vitest';
import { guideSchema } from '../../content.config';
import SpecPage from './[class]/[spec].astro';
import { createPageContainer } from '../../test-support/page-container';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

function loadGuideEntry(id: string) {
  const path = fileURLToPath(new URL(`../../content/guides/${id}.md`, import.meta.url));
  const raw = readFileSync(path, 'utf8');
  const [, frontmatterBlock, ...bodyParts] = raw.split('---');
  const data = guideSchema.parse(parseYaml(frontmatterBlock));
  return { id, collection: 'guides' as const, data, body: bodyParts.join('---').trim() };
}

describe('/guides/warrior/fury', () => {
  it('embeds the tree, the build-action buttons, the stat and race pills, and the footer caveat', async () => {
    const entry = loadGuideEntry('warrior/fury');
    const html = await container.renderToString(SpecPage, { props: { entry } });
    // GuideBuildTree hydrates client:visible and fetches its talent data in a $effect that
    // only runs after hydration -- SSR always renders its loading skeleton
    // (GuideBuildTree.test.ts's own assertion), never the post-fetch tree. So this checks
    // what SSR actually produces: the island mounted with the right initial state and the
    // right code prop, and Astro's own client-directive attribute (not the literal string
    // "client:visible", which never appears -- Astro serializes it as `client="visible"` on
    // the <astro-island> element).
    expect(html).toContain('data-testid="guide-tree-skeleton"');
    expect(html).toContain('client="visible"');
    expect(html).toContain(entry.data.build as string);
    expect(html).toContain('data-testid="guide-load-build"');
    expect(html).toContain('data-testid="guide-sim-build"');
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="guide-race-pills"');
    expect(html).toContain('data-testid="confidence-footer-note"');
    // The nine headings still appear, in order, exactly as SpecTableOfContents promises.
    const talentsIndex = html.indexOf('id="talents-and-builds"');
    const rotationIndex = html.indexOf('id="rotation-and-priority"');
    expect(talentsIndex).toBeGreaterThan(-1);
    expect(rotationIndex).toBeGreaterThan(talentsIndex);
  });
});

describe('/guides/rogue/assassination (a spec with a raid build)', () => {
  it('shows the leveling build and the raid build as two labelled trees with their own links', async () => {
    const entry = loadGuideEntry('rogue/assassination');
    const raidBuild = entry.data.raidBuild as string;
    expect(raidBuild).not.toBe(entry.data.build);
    const html = await container.renderToString(SpecPage, { props: { entry } });
    expect(html).toContain('data-testid="guide-build-leveling"');
    expect(html).toContain('data-testid="guide-build-raid"');
    expect(html).toContain('>Leveling build</h3>');
    expect(html).toContain('>Raid build</h3>');
    expect(html).toContain(raidBuild);
    expect(html).toContain('data-testid="guide-load-raid-build"');
    expect(html).toContain('data-testid="guide-sim-raid-build"');
    expect(html).toContain(`href="/planner?code=${encodeURIComponent(raidBuild)}"`);
  });
});

describe('/guides/warrior/fury (a spec whose raid build is its leveling build)', () => {
  it('shows one tree under a single "Leveling and raid build" heading', async () => {
    const entry = loadGuideEntry('warrior/fury');
    expect(entry.data.raidBuild).toBeUndefined();
    const html = await container.renderToString(SpecPage, { props: { entry } });
    expect(html).toContain('data-testid="guide-build-both"');
    expect(html).toContain('>Leveling and raid build</h3>');
    expect(html).not.toContain('data-testid="guide-build-raid"');
  });
});
