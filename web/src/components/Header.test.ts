import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Header from './Header.astro';

async function renderHeader(path: string): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(Header, { props: { path } });
}

describe('Header', () => {
  it('renders the doors and Get set up, in that order', async () => {
    const html = await renderHeader('/');
    const order = ['Planner', 'BiS', 'Simulator', 'Logs', 'Rankings', 'Tier List', 'Guides', 'Get set up'];
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

  it('renders a collapsed Menu button that controls the menu panel holding the doors', async () => {
    const html = await renderHeader('/');
    expect(html).toMatch(/<button[^>]*aria-expanded="false"[^>]*aria-controls="site-menu"/);
    expect(html).toContain('id="site-menu"');
    expect(html.indexOf('id="site-menu"')).toBeLessThan(html.indexOf('id="primary-nav"'));
  });

  it('never says Leveling BiS and labels the Discord link for screen readers', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('Leveling BiS');
    expect(html).toMatch(/<a[^>]*aria-label="Discord"/);
  });

  it('has a slot for the selector in the bar, before Discord', async () => {
    const html = await renderHeader('/');
    expect(html.indexOf('site-right')).toBeLessThan(html.indexOf('site-discord'));
    expect(html.indexOf('site-menu"')).toBeLessThan(html.indexOf('primary-nav'));
  });

  it('renders the phone nav as a fixed-height wrapping grid, not a horizontally scrolling row', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('overflow-x-auto');
    expect(html).not.toContain('nav-scroll-fade');
  });
});
