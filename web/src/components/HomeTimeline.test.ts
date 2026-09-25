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
      props: { dates: DATES, now: new Date('2026-09-24T00:00:00Z'), updated: new Date('2026-09-20T00:00:00Z') },
    });
    expect(html).toContain('data-testid="home-timeline"');
    expect((html.match(/data-testid="home-timeline-row"/g) ?? []).length).toBe(4);
    expect(html).toContain('data-status="past"');
    expect(html).toContain('data-status="next"');
    expect(html).toContain('data-status="future"');
    expect((html.match(/data-testid="home-timeline-dot"/g) ?? []).length).toBe(1);
    expect(html).toContain('Beta opens');
    expect(html).toContain('level cap 30');
    expect(html).toContain('Updated Sept 20');
  });
});
