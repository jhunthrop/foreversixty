import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Header from './Header.astro';

async function renderHeader(path: string): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(Header, { props: { path } });
}

describe('Header', () => {
  it('renders the four tools, Reference, and The addon, in that order', async () => {
    const html = await renderHeader('/');
    const order = ['Planner', 'Simulator', 'Logs', 'Rankings', 'Reference', 'The addon'];
    let cursor = -1;
    for (const label of order) {
      const at = html.indexOf(`>${label}<`, cursor === -1 ? 0 : cursor);
      expect(at, `${label} not found after the previous item`).toBeGreaterThan(cursor);
      cursor = at;
    }
  });

  it('does not render Changelog (moved to the footer)', async () => {
    const html = await renderHeader('/');
    expect(html).not.toContain('>Changelog<');
  });

  it('renders the Reference disclosure as a closed, native details element', async () => {
    const html = await renderHeader('/');
    expect(html).toMatch(/<details[^>]*>[\s\S]*<summary[^>]*>[\s\S]*Reference/);
    expect(html).not.toContain('<details open');
  });

  it('lists Classes, Guides, Zones, Dungeons inside the Reference panel', async () => {
    const html = await renderHeader('/');
    for (const label of ['Classes', 'Guides', 'Zones', 'Dungeons']) {
      expect(html).toContain(`>${label}<`);
    }
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

  it('marks the Reference summary aria-current when a reference page is current', async () => {
    const html = await renderHeader('/classes');
    const summary = html.match(/<summary[^>]*>/)?.[0] ?? '';
    expect(summary).toContain('aria-current="page"');
  });

  it('does not mark the Reference summary aria-current elsewhere', async () => {
    const html = await renderHeader('/planner');
    const summary = html.match(/<summary[^>]*>/)?.[0] ?? '';
    expect(summary).not.toContain('aria-current');
  });

  it('renders the session slot next to Discord', async () => {
    const html = await renderHeader('/');
    const discordAt = html.indexOf('Discord');
    const headerCloseAt = html.lastIndexOf('</header>');
    expect(discordAt).toBeGreaterThan(-1);
    expect(discordAt).toBeLessThan(headerCloseAt);
  });
});
