// web/src/pages/_classes.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts: everything else under
// src/pages/ is a route, so an unprefixed classes.test.ts would build as the route
// /classes.test and run its top-level beforeAll during the build. Astro skips
// `_`-prefixed files; vitest still collects it.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { beforeAll, describe, expect, it } from 'vitest';
import Classes from './classes.astro';

let html: string;

beforeAll(async () => {
  const container = await AstroContainer.create();
  html = await container.renderToString(Classes);
});

describe('classes.astro', () => {
  it('lists every class and every race, Skyborne included', () => {
    for (const name of ['Warrior', 'Paladin', 'Druid', 'Undead', 'Tauren', 'Skyborne']) {
      expect(html).toContain(name);
    }
  });

  it('deep links every class into the planner', () => {
    expect(html).toContain('href="/planner?class=paladin"');
    expect(html).toContain('href="/planner?class=warrior"');
  });

  it('deep links each legal combination with both the class and the race', () => {
    // The ampersand may come back raw or entity-escaped depending on how Astro serializes
    // the attribute; the test cares about the link, not the encoding.
    const link = (cls: string, race: string) =>
      new RegExp(`href="/planner\\?class=${cls}(&|&#38;|&amp;)race=${race}"`);
    expect(html).toMatch(link('paladin', 'undead'));
    expect(html).toMatch(link('warrior', 'human'));
  });

  it('marks the combinations Forever adds', () => {
    expect(html).toContain('New in Forever');
  });

  it('renders a source pill for every documented change', () => {
    expect(html).toMatch(/pill-(blizzard|datamined|community|site)/);
  });

  it('says which rows are still placeholders', () => {
    expect(html).toContain('Placeholder');
  });

  it('exposes the Forever marking in the accessible name, not just a sighted-only pill', () => {
    // The pill's text sits inside the link, but the aria-label overrides it, so the marking
    // has to be folded into the label itself or a screen reader never hears it.
    expect(html).toContain('aria-label="Plan a Undead Paladin, new in Forever"');
    expect(html).toContain('aria-label="Plan a Human Warrior"');
    expect(html).not.toContain('aria-label="Plan a Human Warrior, new in Forever"');
  });

  it('ships no island', () => {
    expect(html).not.toContain('<astro-island');
  });
});
