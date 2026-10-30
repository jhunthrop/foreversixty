import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import HomeDatesLine from './HomeDatesLine.astro';

const DATES = [
  { key: 'Sept 17', value: 'Beta opens', iso: '2026-09-17', endIso: '2026-10-21', role: 'beta' },
  {
    key: 'Oct 27',
    value: 'Name reservation',
    iso: '2026-10-27',
    endIso: '2026-11-03',
    running: 'Name reservation open through Nov 3',
  },
  {
    key: 'Nov 4',
    value: 'Launch',
    role: 'launch',
    iso: '2026-11-04',
    time: '15:00 PST',
    at: '2026-11-04T23:00:00Z',
  },
  { key: 'Dec 9', value: 'First raids', role: 'raids', iso: '2026-12-09' },
];

async function render(now: string): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(HomeDatesLine, { props: { dates: DATES, now: new Date(now) } });
}

describe('HomeDatesLine', () => {
  it('renders one line, the running event first and in gold, and carries the entries for the clock script', async () => {
    const html = await render('2026-10-29T12:00:00Z');
    expect(html).toContain('data-testid="home-dates-line"');
    expect(html).toContain('data-entries=');
    expect(html).toMatch(/text-gold[^>]*data-next[^>]*>\s*Name reservation open through Nov 3/);
    expect(html).toContain('Launch Nov 4 · 15:00 PST');
    expect(html).toContain('First raids Dec 9');
  });

  it('drops the running event once it has ended', async () => {
    const html = await render('2026-11-04T01:00:00Z');
    expect(html).not.toMatch(/dates-item[^>]*>\s*Name reservation/);
    expect(html).toMatch(/text-gold[^>]*data-next[^>]*>\s*Launch Nov 4/);
  });

  it('steps aside signed out while the launch card counts down, and not after launch', async () => {
    expect(await render('2026-10-29T12:00:00Z')).toContain('data-under-card');
    expect(await render('2026-11-05T12:00:00Z')).toContain('data-under-card');
    expect(await render('2026-12-09T12:00:00Z')).not.toContain('data-under-card');
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
