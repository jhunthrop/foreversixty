// web/src/components/guides/RacePillRow.test.ts
// See BuildActionButtons.test.ts's header comment for why this file exists at all: the
// codebase already has a co-located AstroContainer render-test convention for .astro
// components (ClassTile.test.ts, FeedRow.test.ts, ScopeNote.test.ts), so this applies it
// here rather than deferring coverage to a later whole-page test.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { guidesCopy } from '../../lib/guides/copy';
import RacePillRow from './RacePillRow.astro';

describe('RacePillRow', () => {
  it('lists every race legal for the class, marking recommended vs not', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(RacePillRow, {
      props: { classSlug: 'warrior', recommendedRaces: ['human', 'troll'] },
    });
    expect(html).toContain('data-testid="guide-race-pills"');
    expect(html).toContain('data-testid="race-pill-human"');
    expect(html).toContain('data-testid="race-pill-orc"');
  });

  it('gives a recommended race a filled pill and the screen-reader label', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(RacePillRow, {
      props: { classSlug: 'warrior', recommendedRaces: ['human'] },
    });
    const human = /<span[^>]*data-testid="race-pill-human"[^>]*>[\s\S]*?<\/span>/.exec(html)?.[0] ?? '';
    expect(human).toContain('data-recommended="true"');
    expect(human).toContain('color-mix(');
    expect(human).toContain(guidesCopy.recommendedRaceLabel);
  });

  it('gives a non-recommended race an outline pill and no screen-reader label', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(RacePillRow, {
      props: { classSlug: 'warrior', recommendedRaces: ['human'] },
    });
    const orc = /<span[^>]*data-testid="race-pill-orc"[^>]*>[\s\S]*?<\/span>/.exec(html)?.[0] ?? '';
    expect(orc).toContain('data-recommended="false"');
    expect(orc).not.toContain('color-mix(');
    expect(orc).not.toContain(guidesCopy.recommendedRaceLabel);
  });

  it('returns no pills for an unknown class slug', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(RacePillRow, {
      props: { classSlug: 'not-a-class', recommendedRaces: [] },
    });
    expect(html).toContain('data-testid="guide-race-pills"');
    expect(html).not.toContain('data-testid="race-pill-');
  });
});
