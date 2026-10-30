import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import HomeDatesLine from './HomeDatesLine.astro';

const DATES = [
  { key: 'Oct 27', value: 'Name reservation', iso: '2026-10-27', endIso: '2026-11-03' },
  { key: 'Nov 4', value: 'Launch', iso: '2026-11-04', time: '15:00 PST', at: '2026-11-04T23:00:00Z' },
  { key: 'Dec 9', value: 'First raids', iso: '2026-12-09' },
];

async function render(now: string): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(HomeDatesLine, { props: { dates: DATES, now: new Date(now) } });
}

describe('HomeDatesLine', () => {
  it('renders one line of the dates still to come, the next in gold, and carries the entries for the clock script', async () => {
    const html = await render('2026-10-29T12:00:00Z');
    expect(html).toContain('data-testid="home-dates-line"');
    expect(html).toContain('data-entries=');
    expect((html.match(/dates-item/g) ?? []).length).toBeGreaterThanOrEqual(2);
    expect(html).toMatch(/text-gold[^>]*data-next[^>]*>\s*Launch Nov 4 · 15:00 PST/);
    expect(html).toContain('First raids Dec 9');
    expect(html).not.toMatch(/dates-item[^>]*>\s*Name reservation/);
  });

  it('is body-face tabular, never mono', async () => {
    const html = await render('2026-10-29T12:00:00Z');
    expect(html).toContain('numeric');
    expect(html).not.toContain('mono');
  });

  it('renders nothing once every date has passed', async () => {
    const html = await render('2026-12-10T12:00:00Z');
    expect(html).not.toContain('home-dates-line');
  });
});
