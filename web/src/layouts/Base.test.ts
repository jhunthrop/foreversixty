// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: Base.astro mounts the character selector (client:idle) on
// every page, and Astro's integrations are not loaded in a unit test, so without this the
// container throws NoMatchingRenderer rather than rendering an empty shell.
import { createPageContainer } from '../test-support/page-container';
import { describe, expect, it } from 'vitest';
import Base from './Base.astro';

describe('Base layout', () => {
  it('sets title, canonical, and og image from props', async () => {
    const container = await createPageContainer();
    const html = await container.renderToString(Base, {
      props: { title: 'Dungeons', description: 'All nine.', path: '/dungeons' },
      slots: { default: '<main>body</main>' },
    });
    expect(html).toContain('<title data-og="title">Dungeons · Forever Sixty · WoW Forever</title>');
    expect(html).toContain('href="https://foreversixty.gg/dungeons"');
    expect(html).toContain('content="https://foreversixty.gg/og/dungeons.png"');
    expect(html).toContain('<main>body</main>');
  });

  it('keeps og:title in step with the title and names the site', async () => {
    const container = await createPageContainer();
    const html = await container.renderToString(Base, {
      props: { title: 'Tier list', description: 'x', path: '/tiers' },
    });
    expect(html).toContain('property="og:title" content="Tier list · Forever Sixty · WoW Forever"');
    expect(html).toContain('property="og:site_name" content="Forever Sixty"');
  });

  it('gives the homepage the game name and the WebSite node', async () => {
    const container = await createPageContainer();
    const html = await container.renderToString(Base, {
      props: { title: 'Forever Sixty', description: 'x', path: '/' },
    });
    expect(html).toContain(
      '<title data-og="title">Forever Sixty · Play your class better in World of Warcraft: Forever</title>',
    );
    expect(html).toContain('"alternateName":"foreversixty.gg"');
  });
});
