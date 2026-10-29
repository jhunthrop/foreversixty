import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import HomeTimeline from './HomeTimeline.astro';

const DATES = [
  { key: 'Sept 13', value: 'Deep Dive panel', iso: '2026-09-13' },
  { key: 'Sept 17', value: 'Beta opens', note: 'level cap 30', iso: '2026-09-17' },
  { key: 'Oct 27', value: 'Name reservation', iso: '2026-10-27' },
  { key: 'Nov 4', value: 'Launch', iso: '2026-11-04' },
];

describe('HomeTimeline', () => {
  it('renders every row with its own status, and the updated stamp', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(HomeTimeline, {
      props: {
        dates: DATES,
        now: new Date('2026-09-24T00:00:00Z'),
        updated: new Date('2026-09-20T00:00:00Z'),
      },
    });
    expect(html).toContain('data-testid="home-timeline"');
    expect((html.match(/data-testid="home-timeline-row"/g) ?? []).length).toBe(2);
    expect(html).not.toContain('data-status="past"');
    expect(html).toContain('data-status="next"');
    expect(html).toContain('data-status="future"');
    expect((html.match(/data-testid="home-timeline-dot"/g) ?? []).length).toBe(1);
    expect(html).toContain('Name reservation');
    expect(html).toContain('Launch');
    expect(html).toContain('Updated Sept 20');
  });

  it('is focusable and named, so a keyboard user can reach and scroll it once it overflows (axe scrollable-region-focusable)', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(HomeTimeline, {
      props: { dates: DATES, now: new Date('2026-09-24T00:00:00Z') },
    });
    expect(html).toContain('tabindex="0"');
    expect(html).toContain('aria-label="Key dates timeline"');
  });

  it('drops a passed date: no past row, no dimming', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(HomeTimeline, {
      props: { dates: DATES, now: new Date('2026-09-24T00:00:00Z') },
    });
    expect(html).not.toContain('>Deep Dive panel<');
    expect(html).not.toContain('>Beta opens<');
    expect(html).not.toContain('opacity-60');
  });

  it('lists what is still to come in date order and carries the entries for the clock script', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(HomeTimeline, {
      props: { dates: DATES, now: new Date('2026-09-24T00:00:00Z') },
    });
    const statuses = [...html.matchAll(/data-status="([a-z]+)"/g)].map((m) => m[1]);
    expect(statuses).toEqual(['next', 'future']);
    expect(html.indexOf('>Name reservation<')).toBeLessThan(html.indexOf('>Launch<'));
    expect(html).toContain('data-entries=');
  });

  it('shows no stamp unless one is passed', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(HomeTimeline, { props: { dates: DATES } });
    expect(html).not.toContain('Updated');
  });
});
