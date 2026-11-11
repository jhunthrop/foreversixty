// web/src/pages/_seo-text.test.ts
// The text-floor guard (seo-fix 2026-11-11, lane G2): Search Console sent home's and /rankings'
// impressions to other pages because their server render held about 90 and 17 words. Home now draws
// its own words (HomeAbout) and the rankings pages the build's snapshot rows; this renders each page
// as the build does and fails when the words a crawler reads in <main> drop under the floor, so a
// later change cannot quietly take them away again.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import type { AstroComponentFactory } from 'astro/runtime/server/index.js';
import { beforeAll, describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import { staticWords } from '../test-support/static-text';
import Home from './index.astro';
import Rankings from './rankings.astro';
import RankingsEncounter from './rankings/[slug].astro';

/** The floors, a little under what each page draws today: home 300 (the brief's floor), the
 *  rankings landing and an encounter's page 150 (the fixture snapshot's rows). */
export const TEXT_FLOORS = { home: 300, rankings: 150, encounter: 150 } as const;

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

describe('the counter itself', () => {
  it('counts only what a crawler reads in <main>', () => {
    const thin = `<header>Rankings Tier list Guides BiS</header><main><h1>Rankings</h1>
      <script>const words = "one two three four five";</script><style>.a{}</style>
      <noscript>Turn on JavaScript to see the board</noscript><svg><text>crest</text></svg>
      <p>Loading&hellip;</p></main><footer>Privacy Terms</footer>`;
    expect(staticWords(thin)).toEqual(['Rankings', 'Loading']);
  });

  it('fails a page as thin as /rankings was before the snapshot', () => {
    const before = '<main><h1>Rankings</h1><p>Tier list</p><p>Encounters</p><p>Loading</p></main>';
    expect(staticWords(before).length).toBeLessThan(TEXT_FLOORS.rankings);
  });
});

describe('server-rendered text floors', () => {
  it(`home reads at least ${TEXT_FLOORS.home} words signed out`, async () => {
    const words = staticWords(await container.renderToString(Home));
    expect(words.length).toBeGreaterThanOrEqual(TEXT_FLOORS.home);
  });

  it(`the rankings landing reads at least ${TEXT_FLOORS.rankings} words`, async () => {
    const words = staticWords(await container.renderToString(Rankings));
    expect(words.length).toBeGreaterThanOrEqual(TEXT_FLOORS.rankings);
  });

  it(`an encounter's rankings page reads at least ${TEXT_FLOORS.encounter} words`, async () => {
    const html = await container.renderToString(RankingsEncounter as unknown as AstroComponentFactory, {
      params: { slug: 'onyxia' },
    });
    expect(staticWords(html).length).toBeGreaterThanOrEqual(TEXT_FLOORS.encounter);
  });
});
