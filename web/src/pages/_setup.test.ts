// web/src/pages/_setup.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts and _classes.test.ts:
// everything else under src/pages/ is a route, so an unprefixed setup.test.ts would build as
// the route /setup.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// Adapted from _addon.test.ts (Task 6, spec 2026-09-25 §3.4): /setup replaces /addon with
// three numbered panels. The GitHub release link is derived from links.json's githubRepo, so
// it is asserted off the rendered output rather than as a literal substring of the source,
// same as the page it replaces.
import { readFileSync } from 'node:fs';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { describe, expect, it, beforeAll } from 'vitest';
import { addonCopy } from '../lib/addon/copy';
import { SETUP_NAV_ITEM } from '../lib/nav';
import activeBuild from '../data/active-build.json';
import SetupPage, { CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL } from './setup.astro';

const source = readFileSync(new URL('./setup.astro', import.meta.url), 'utf8');

const STRING_COPY_KEYS = [
  'pageDescription',
  'pageNoNetwork',
  'installCurseForge',
  'installWago',
  'installGitHub',
  'flowOutTitle',
  'flowOutBody',
  'flowInTitle',
  'flowInBody',
  'setupStep1Title',
  'setupStep2Title',
  'setupStep3Title',
  'setupSignInBody',
  'setupSignInRealmNote',
  'setupSignInAction',
  'setupCompanionBody',
  'setupPageTitle',
  'setupStatusSignIn',
  'setupStatusSignedIn',
  'setupStatusAddon',
  'setupStatusAddonDone',
  'setupStatusCompanion',
  'setupSignedInLine',
  'setupSignedInLink',
  'setupAddonDoneLine',
  'setupHowItWorks',
  'setupCompanionDownloadLead',
  'setupCompanionPair',
  'disclosureOpen',
  'disclosureClose',
] as const;

let html: string;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  const container = await AstroContainer.create({ renderers });
  html = await container.renderToString(SetupPage);
});

describe('/setup', () => {
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

  it('is reachable from the primary nav as "Get set up"', () => {
    expect(SETUP_NAV_ITEM.href).toBe('/setup');
  });

  it('links all three install routes in the rendered output', () => {
    expect(html).toContain(CURSEFORGE_URL);
    expect(html).toContain(WAGO_URL);
    expect(html).toContain(GITHUB_RELEASES_URL);
  });

  it('shows the current data build from the data file in the rendered output', () => {
    expect(html).toContain(addonCopy.currentDataBuild(activeBuild.build));
  });

  it('names which realm types Blizzard serves today under step 1', () => {
    expect(html).toContain(addonCopy.setupSignInRealmNote);
  });

  it('every install link opts out of opener access', () => {
    for (const url of [CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL]) {
      const anchor = [...html.matchAll(/<a\s+[^>]*>/g)].map((m) => m[0]).find((tag) => tag.includes(url));
      expect(anchor, `no <a> found for ${url}`).toBeDefined();
      expect(anchor).toContain('rel="noopener"');
    }
  });

  it('keeps the paste box at the #paste anchor', () => {
    expect(html).toContain('id="paste"');
  });

  it('renders the three numbered steps in order', () => {
    const order = [addonCopy.setupStep1Title, addonCopy.setupStep2Title, addonCopy.setupStep3Title];
    let cursor = -1;
    for (const title of order) {
      const at = html.indexOf(title, cursor === -1 ? 0 : cursor);
      expect(at, `${title} not found after the previous step`).toBeGreaterThan(cursor);
      cursor = at;
    }
  });
});
