// web/src/pages/_addon.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts and _classes.test.ts:
// everything else under src/pages/ is a route, so an unprefixed addon.test.ts would build
// as the route /addon.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// The GitHub release link is derived from links.json's githubRepo (constraints.md,
// data-sync footgun section, and the controller's correction to this task's brief), so it
// can no longer be asserted as a literal substring of the source -- it has to be read back
// off the rendered output. The three install links and the current data build are asserted
// that way; the copy-file rule (every visible string comes from addonCopy, never a
// literal) genuinely is a property of the source text, so that stays a source-text check.
//
// The Svelte renderer has to be handed to the container explicitly since addon.astro now
// mounts AddonPasteBox.svelte -- Astro's integrations are not loaded in a unit test,
// without this the island renders as an empty shell (see _planner.test.ts for the same
// pattern).
import { readFileSync } from 'node:fs';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { describe, expect, it, beforeAll } from 'vitest';
import { addonCopy } from '../lib/addon/copy';
import activeBuild from '../data/active-build.json';
import AddonPage, { CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL } from './addon.astro';

const source = readFileSync(new URL('./addon.astro', import.meta.url), 'utf8');

// String-valued copy keys this page uses. currentDataBuild is a function and is checked
// separately below -- toContain against a function value doesn't express "the literal
// text appears nowhere in the source" the way it does for a string.
const STRING_COPY_KEYS = [
  'pageTitle',
  'pageDescription',
  'pageNoNetwork',
  'installCurseForge',
  'installWago',
  'installGitHub',
  'flowOutTitle',
  'flowOutBody',
  'flowInTitle',
  'flowInBody',
] as const;

let html: string;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  const container = await AstroContainer.create({ renderers });
  html = await container.renderToString(AddonPage);
});

describe('/addon', () => {
  it('takes every visible string from the copy file, never a literal', () => {
    for (const key of STRING_COPY_KEYS) {
      expect(source).toContain(`addonCopy.${key}`);
      expect(source).not.toContain(addonCopy[key]);
    }
    expect(source).toContain('addonCopy.currentDataBuild(');
  });

  it('derives the GitHub release link from links.json, not a second copy of the repo URL', () => {
    expect(source).toContain('communityLinks.githubRepo');
    expect(source).not.toContain('jhunthrop/forever/releases');
  });

  it('is reachable from the footer', () => {
    const footer = readFileSync(new URL('../components/Footer.astro', import.meta.url), 'utf8');
    expect(footer).toContain('/addon');
  });

  it('links all three install routes in the rendered output', () => {
    expect(html).toContain(CURSEFORGE_URL);
    expect(html).toContain(WAGO_URL);
    expect(html).toContain(GITHUB_RELEASES_URL);
  });

  it('shows the current data build from the data file in the rendered output', () => {
    expect(html).toContain(addonCopy.currentDataBuild(activeBuild.build));
  });

  it('every install link opts out of opener access', () => {
    // Scoped to the page's own three install links -- Base pulls in Header, which links
    // Discord without rel="noopener" and is out of this task's scope.
    for (const url of [CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL]) {
      const anchor = [...html.matchAll(/<a\s+[^>]*>/g)].map((m) => m[0]).find((tag) => tag.includes(url));
      expect(anchor, `no <a> found for ${url}`).toBeDefined();
      expect(anchor).toContain('rel="noopener"');
    }
  });
});
