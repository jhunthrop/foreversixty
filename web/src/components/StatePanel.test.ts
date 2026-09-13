import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import StatePanel from './StatePanel.astro';

describe('StatePanel', () => {
  it('renders rows and an updated stamp', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(StatePanel, {
      props: {
        title: 'Right now',
        updated: new Date('2026-09-12T00:00:00Z'),
        rows: [{ key: 'Sept 17', value: 'Beta opens', note: 'level cap 30', highlight: true }],
      },
    });
    expect(html).toContain('Right now');
    expect(html).toContain('Beta opens');
    expect(html).toContain('level cap 30');
    expect(html).toContain('Updated Sept 12');
  });
  it('shows a Sample pill when sample is true', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(StatePanel, {
      props: { title: 'This week', sample: true, rows: [] },
    });
    expect(html).toContain('pill-sample');
  });
});
