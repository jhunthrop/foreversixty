// web/src/pages/guides/_guides.test.ts
// Underscore-prefixed for the same reason as _classes.test.ts: Astro skips `_`-prefixed
// files when building routes, vitest still collects it.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { getCollection } from 'astro:content';
import { beforeAll, describe, expect, it } from 'vitest';
import GuidesIndex from './index.astro';
import ClassGuide from './[class]/index.astro';
import SpecGuide from './[class]/[spec].astro';
import { slugify, SPEC_SECTIONS } from '../../lib/guides/sections';

let container: AstroContainer;
let indexHtml: string;
let classHtml: string;
let specHtml: string;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
  indexHtml = await container.renderToString(GuidesIndex);

  const guides = await getCollection('guides');
  const warriorLanding = guides.find((g) => g.id === 'warrior/index')!;
  const warriorSpecGuides = guides.filter((g) => g.data.classSlug === 'warrior' && g.data.spec !== undefined);
  classHtml = await container.renderToString(ClassGuide, {
    props: { entry: warriorLanding, specGuides: warriorSpecGuides },
  });

  const specEntry = warriorSpecGuides[0]!;
  specHtml = await container.renderToString(SpecGuide, { props: { entry: specEntry } });
});

describe('guides/index.astro', () => {
  it('lists every class and links its spec guides', () => {
    for (const name of ['Warrior', 'Paladin', 'Druid']) {
      expect(indexHtml).toContain(name);
    }
    expect(indexHtml).toContain('href="/guides/warrior/fury"');
  });

  it('ships no island', () => {
    expect(indexHtml).not.toContain('<astro-island');
  });
});

describe('guides/[class]/index.astro', () => {
  it('renders the class landing page and links out to its spec guides', () => {
    expect(classHtml).toContain('Warrior');
    expect(classHtml).toContain('href="/guides/warrior/fury"');
    expect(classHtml).toContain('href="/planner?class=warrior"');
  });

  it('renders the sources footer', () => {
    expect(classHtml).toMatch(/pill-(blizzard|datamined|community|site)/);
  });
});

describe('guides/[class]/[spec].astro', () => {
  it('renders a table of contents with all nine sections, in order', () => {
    for (const section of SPEC_SECTIONS) {
      expect(specHtml).toContain(`href="#${slugify(section)}"`);
      expect(specHtml).toContain(`>${section}<`);
    }
  });

  it('renders the sources footer and a planner deep link', () => {
    expect(specHtml).toMatch(/pill-(blizzard|datamined|community|site)/);
    expect(specHtml).toContain('href="/planner?class=warrior"');
  });
});
