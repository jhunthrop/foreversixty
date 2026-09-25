import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Header from './Header.astro';

async function renderHeader(path: string): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(Header, { props: { path } });
}

describe('Header', () => {
  it('renders the five doors and Get set up, in that order', async () => {
    const html = await renderHeader('/');
    const order = ['Planner', 'Simulator', 'Logs', 'Rankings', 'Guides', 'Get set up'];
    let cursor = -1;
    for (const label of order) {
      const at = html.indexOf(`>${label}<`, cursor === -1 ? 0 : cursor);
      expect(at, `${label} not found after the previous item`).toBeGreaterThan(cursor);
      cursor = at;
    }
  });

  it('renders no Reference disclosure and no Premium link', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('<details');
    expect(html).not.toContain('>Premium<');
    expect(html).not.toContain('>Classes<');
    expect(html).not.toContain('>The addon<');
  });

  it('marks Planner aria-current when the path is /planner', async () => {
    const html = await renderHeader('/planner');
    expect(html).toMatch(/<a href="\/planner" aria-current="page"[^>]*>[\s\S]{0,20}Planner/);
  });

  it('does not mark Planner aria-current on an unrelated path', async () => {
    const html = await renderHeader('/sim');
    const plannerLink = html.match(/<a href="\/planner"[^>]*>/)?.[0] ?? '';
    expect(plannerLink).not.toContain('aria-current');
  });

  it('marks Guides aria-current on a guide sub-path', async () => {
    const html = await renderHeader('/guides/warrior');
    expect(html).toMatch(/<a href="\/guides" aria-current="page"[^>]*>[\s\S]{0,20}Guides/);
  });

  it('renders the session slot next to Discord', async () => {
    const html = await renderHeader('/');
    const discordAt = html.indexOf('Discord');
    const headerCloseAt = html.lastIndexOf('</header>');
    expect(discordAt).toBeGreaterThan(-1);
    expect(discordAt).toBeLessThan(headerCloseAt);
  });

  it('renders the phone nav as a fixed-height wrapping grid, not a horizontally scrolling row', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('overflow-x-auto');
    expect(html).not.toContain('nav-scroll-fade');
  });
});
