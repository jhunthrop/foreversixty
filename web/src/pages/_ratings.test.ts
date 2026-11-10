// web/src/pages/_ratings.test.ts
// Underscore-prefixed like _setup.test.ts: everything else under src/pages/ is a route.
//
// /ratings explains both numbers (design/specs/2026-11-07-logs-reports.md section 8, T1): the
// PageHeader with the sample Forever Score as its answer in its tier's colour and the two-tab
// jump row; the Forever Score Section with its five parts as the one Plate (the addon's labels,
// from the shared fixture); the Performance Rating Section with three figures; the reference
// Sections; the tiers in the item quality colours with the sample tinted (MUST 6); how logs get
// here. No "percentile" in the rating's words.
import { readFileSync } from 'node:fs';
import { beforeAll, describe, expect, it } from 'vitest';
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import { ratingCopy } from '../lib/rating/copy';
import { scoreCopy } from '../lib/score/copy';
import { SAMPLE_SCORE, WEIGHTS, ratingsPageCopy } from '../lib/score/ratings-page';
import { COMPONENT_MAX, tierRanges } from '../lib/score/score';
import { createPageContainer } from '../test-support/page-container';
import RatingsPage from './ratings.astro';

const source = readFileSync(new URL('../components/score/RatingsPage.svelte', import.meta.url), 'utf8');

let html: string;

beforeAll(async () => {
  const container = await createPageContainer();
  html = await container.renderToString(RatingsPage);
});

/** A Section's markup, from its region to the next Section's. */
function section(slug: string): string {
  const start = html.indexOf(`data-region="section-${slug}"`);
  const end = html.indexOf('data-region="section-', start + 1);
  return start < 0 ? '' : html.slice(start, end < 0 ? undefined : end);
}

function text(fragment: string): string {
  return fragment
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ');
}

describe('/ratings', () => {
  it('is titled "Forever Score and ratings" with a two-tab jump row to its two Sections', () => {
    expect(html).toMatch(/<h1[^>]*>Forever Score and ratings<\/h1>/);
    const jump = /data-testid="ratings-jump"[\s\S]*?<\/nav>/.exec(html)?.[0] ?? '';
    expect(jump).toContain('href="#forever-score"');
    expect(jump).toContain('href="#performance-rating"');
    expect(html).toContain('id="forever-score"');
    expect(html).toContain('id="performance-rating"');
  });

  it('bridges the two systems with one caption each, in the head, said once', () => {
    expect(text(html).split(scoreCopy.explainerScoreCaption).length - 1).toBe(1);
    expect(text(html).split(scoreCopy.explainerRatingCaption).length - 1).toBe(1);
  });

  it('answers with the sample score in its tier colour, out of the five ceilings', () => {
    const answer = /data-region="head-answer"[\s\S]*?<\/header>/.exec(html)?.[0] ?? '';
    expect(text(answer)).toContain(`${scoreCopy.samplePill} · ${scoreCopy.explainerHeading}`);
    expect(answer).toMatch(/data-testid="ratings-sample-score"[\s\S]*?style="color: var\(--color-rarity-epic\)/);
    const max = Object.values(COMPONENT_MAX).reduce((sum, n) => sum + n, 0);
    expect(text(answer)).toContain(`of ${max.toLocaleString('en-US')} · Epic`);
    expect(text(answer)).toContain(scoreCopy.mostToGainLead);
  });

  it('opens the Forever Score on its parts, then says when raid parts count', () => {
    const score = text(section('forever-score'));
    expect(score).toContain(scoreCopy.explainerLaunch);
  });

  it('sets every section on its panel and lists the five parts in the addon’s order, each with what moves it', () => {
    expect(html.split('data-region="plate"').length - 1).toBe(4);
    let last = -1;
    for (const component of labels.components) {
      const at = html.indexOf(`data-testid="ratings-forever-score-${component.key}"`);
      expect(at, component.key).toBeGreaterThan(last);
      last = at;
      expect(html).toContain(`>${component.label}<`);
      expect(html).toContain(`>${COMPONENT_MAX[component.key as keyof typeof COMPONENT_MAX]}<`);
    }
    for (const component of scoreCopy.components) expect(html).toContain(component.moves);
  });

  it('says how parses rank in plain words, never "percentile"', () => {
    const parses = scoreCopy.components.find((c) => c.key === 'parses');
    expect(parses?.moves).toBe(
      'How your best damage or healing on each boss ranks against others of your spec, averaged over the bosses you have killed.',
    );
    expect(parses?.moves).not.toMatch(/percentile/i);
  });

  it('reads the Gear blurb by level band', () => {
    expect(scoreCopy.components[0].moves).toContain(
      'How close your gear is to your spec’s best set for your level (20, 30, 40, 50 or 60)',
    );
  });

  it('lists every tier in its item quality colour, the sample’s tinted, then the note', () => {
    const tiers = section('tiers');
    for (const { tier, range } of tierRanges()) {
      expect(tiers).toContain(`color: var(--color-rarity-${tier})`);
      expect(tiers).toContain(range);
    }
    expect(tiers).toMatch(/data-current="true"[^>]*style="--tier: var\(--color-rarity-epic\)/);
    expect(tiers).toContain('The sample');
    expect(tiers).toContain(scoreCopy.tiersNote);
  });

  it('says how logs get here, with a "Get the companion" button', () => {
    const logs = section('how-to-get-your-logs-here');
    expect(logs).toContain(scoreCopy.explainerLogsBody);
    expect(logs).toMatch(/<a[^>]*href="\/setup"[^>]*data-testid="ratings-forever-logs-button"/);
    expect(text(logs)).toContain('Get the companion');
  });

  it('shows the rating sample as three figures: 65 of 100, the last six fights, the weakest part', () => {
    const sample = text(
      /data-testid="ratings-rating-sample"[\s\S]*?<\/svg>[\s\S]*?Weakest part<\/span>[\s\S]*?<\/span>/.exec(
        html,
      )?.[0] ?? '',
    );
    expect(sample).toContain('65');
    expect(sample).toContain(ratingCopy.outOf100);
    expect(sample).toContain('Last 6 fights');
    expect(sample).toContain('Up 25 since the first');
    expect(sample).toMatch(/Weakest part\s*Activity/);
  });

  it('says "ranked against players of your spec", never "percentile", in the rating copy', () => {
    const rating = text(
      [section('performance-rating'), section('ranked-against-your-spec-or-absolute')].join(' '),
    );
    expect(rating).not.toMatch(/percentile/i);
    expect(rating).toContain('ranked against players of your spec');
    expect(rating).toContain(ratingsPageCopy.visibility);
  });

  it('defines each of the six components, the weights by role, the basis and the cap', () => {
    for (const { part, body } of ratingsPageCopy.components) {
      expect(section('what-each-component-measures')).toContain(body);
      expect(text(section('what-each-component-measures'))).toContain(part);
    }
    const weights = text(section('weights-by-role'));
    for (const row of WEIGHTS) {
      // the phone line repeats the three after the primary figure
      expect(weights).toContain(
        `${row.component} ${row.dps} ${row.dps} / ${row.healer} / ${row.tank} ${row.healer} ${row.tank}`,
      );
    }
    expect(weights).toContain(ratingsPageCopy.roles);
    expect(weights).toContain(ratingsPageCopy.adjustable);
    for (const paragraph of ratingsPageCopy.rankedOrAbsolute) {
      expect(section('ranked-against-your-spec-or-absolute')).toContain(paragraph);
    }
    expect(section('the-cap')).toContain(ratingCopy.cappedRule);
  });

  it('draws no card: the two systems on panels, the reference open', () => {
    expect(html).not.toMatch(/card-surface|rounded-panel/);
    expect(SAMPLE_SCORE.tier).toBe('epic');
  });

  it('takes its labels from the fixture and its words from the copy modules, never a literal', () => {
    for (const component of labels.components.filter((c) => c.key !== 'discipline')) {
      expect(source).not.toContain(`'${component.label}'`);
    }
    expect(source).not.toContain(scoreCopy.explainerLaunch);
    expect(source).not.toContain(scoreCopy.tiersNote);
  });
});
