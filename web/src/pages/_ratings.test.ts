// web/src/pages/_ratings.test.ts
// Underscore-prefixed like _setup.test.ts: everything else under src/pages/ is a route.
//
// Task L8.2 ruling 6 and fix round 1 item 7: /ratings explains both numbers. A page title and
// a two-item jump row; the Forever Score section (captioned, its five parts with the addon's
// own labels from the shared fixture, the tier strip with a visually hidden table, how logs
// get here); then the Performance Rating section, its headings one level down, with no
// "percentile" left in its words.
import { readFileSync } from 'node:fs';
import { beforeAll, describe, expect, it } from 'vitest';
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import { scoreCopy } from '../lib/score/copy';
import { tierRanges, tierStrip } from '../lib/score/score';
import { createPageContainer } from '../test-support/page-container';
import RatingsPage from './ratings.astro';

const source = readFileSync(new URL('./ratings.astro', import.meta.url), 'utf8');

let html: string;

beforeAll(async () => {
  const container = await createPageContainer();
  html = await container.renderToString(RatingsPage);
});

/** A top-level section's markup: from its id to the next one's (Panels are sections too, so
 *  the first `</section>` is not the end). */
function section(id: 'forever-score' | 'performance-rating'): string {
  const start = html.indexOf(`id="${id}"`);
  const end = id === 'forever-score' ? html.indexOf('id="performance-rating"') : html.indexOf('</main>');
  return start < 0 || end < 0 ? '' : html.slice(start, end);
}

function text(fragment: string): string {
  return fragment
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ');
}

describe('/ratings page frame', () => {
  it('is titled "Forever Score and ratings" with a two-item jump row', () => {
    expect(html).toMatch(/<h1[^>]*>Forever Score and ratings<\/h1>/);
    const jump = /<nav[^>]*data-testid="ratings-jump"[\s\S]*?<\/nav>/.exec(html)?.[0] ?? '';
    expect(jump).toContain('href="#forever-score"');
    expect(jump).toContain('href="#performance-rating"');
    expect(text(jump)).toContain('Forever Score · Performance Rating');
  });

  it('bridges the two systems with one caption each', () => {
    expect(section('forever-score')).toContain('Forever Score: your whole record, 0 to 2100.');
    expect(section('performance-rating')).toContain('Performance Rating: one fight, 0 to 100.');
  });

  it('puts the rating group’s headings one level down, under its own h2', () => {
    const rating = section('performance-rating');
    expect(rating).toMatch(/<h2[^>]*>Performance Rating<\/h2>/);
    expect(rating).toMatch(/<h3[^>]*>How ratings work<\/h3>/);
    expect(rating).toMatch(/<h3[^>]*>Weights, by role<\/h3>/);
    expect(rating).not.toMatch(/<h2[^>]*>Weights, by role/);
  });

  it('says "ranked against players of your spec", never "percentile", in the rating copy', () => {
    const rating = text(section('performance-rating'));
    expect(rating).not.toMatch(/percentile/i);
    expect(rating).toContain('ranked against players of your spec');
  });

  it('sits on the site column’s own gutter, not a centred narrower measure', () => {
    expect(html).not.toMatch(/<main[^>]*measure-wide/);
  });
});

describe('/ratings Forever Score section', () => {
  it('sits above the Performance Rating', () => {
    const score = html.indexOf('data-testid="ratings-forever-score"');
    expect(score).toBeGreaterThan(-1);
    expect(score).toBeLessThan(html.indexOf('data-testid="ratings-performance-rating"'));
    expect(score).toBeLessThan(html.indexOf('How ratings work'));
  });

  it('opens with what to use it for, then says what the score is, and when raid parts count', () => {
    const forever = section('forever-score');
    expect(scoreCopy.explainerLead).toBe(
      'Use it to see how your character compares and which part to work on.',
    );
    expect(forever.indexOf(scoreCopy.explainerLead)).toBeLessThan(forever.indexOf(scoreCopy.explainerWhat));
    expect(forever).toContain('Parses, attendance and raid bosses count once raids open.');
  });

  it('lists the five components with the addon’s exact labels, in its order, each with what moves it', () => {
    let last = -1;
    for (const component of labels.components) {
      const at = html.indexOf(`data-testid="ratings-forever-score-${component.key}"`);
      expect(at, component.key).toBeGreaterThan(last);
      last = at;
      expect(html).toContain(`>${component.label}<`);
    }
    for (const component of scoreCopy.components) expect(html).toContain(component.moves);
  });

  it('holds each component line to about 70 characters of measure', () => {
    for (const component of labels.components) {
      expect(html).toMatch(
        new RegExp(`class="[^"]*max-w-\\[70ch\\][^"]*" data-testid="ratings-forever-score-${component.key}"`),
      );
    }
  });

  it('says how parses rank in plain words, never "percentile"', () => {
    const parses = scoreCopy.components.find((c) => c.key === 'parses');
    expect(parses?.moves).toBe(
      'How your best damage or healing on each boss ranks against others of your spec, averaged over the bosses you have killed.',
    );
    expect(parses?.moves).not.toMatch(/percentile/i);
  });

  it('draws the tiers as a strip in the tier colours, with the table kept for assistive tech', () => {
    const strip = /<ol[^>]*data-testid="tier-strip"[\s\S]*?<\/ol>/.exec(html)?.[0] ?? '';
    for (const segment of tierStrip()) {
      expect(strip).toContain(`background-color: ${segment.color}`);
      expect(strip).toContain(`color: ${segment.textColor}`);
      expect(strip).toContain(segment.label);
      expect(strip).toContain(segment.range);
    }
    expect(labels.tiers.common.color).toBe('#ffffff');
    const table = /<table[^>]*data-testid="ratings-forever-tiers"[\s\S]*?<\/table>/.exec(html)?.[0] ?? '';
    expect(table).toMatch(/<table class="sr-only"/);
    for (const tier of Object.values(labels.tiers)) expect(table).toContain(tier.label);
    for (const { range } of tierRanges()) expect(table).toContain(range);
  });

  it('notes that the tiers are item quality colours and the ranges may change', () => {
    expect(html).toContain(
      'The tiers take the colours of item quality in game. The ranges may change as more players are logged.',
    );
  });

  it('says how logs get here, with a "Get the companion" button beside it', () => {
    const logs = /<p[^>]*data-testid="ratings-forever-logs"[\s\S]*?<\/p>/.exec(html)?.[0] ?? '';
    expect(logs).toContain(scoreCopy.explainerLogsBody);
    expect(html).toMatch(/href="\/setup" data-testid="ratings-forever-logs-button">Get the companion</);
    expect(html).toContain(`>${scoreCopy.explainerLogsHeading}<`);
  });

  it('leads with one sentence and then the profile’s own card as a Sample (round 2 item 5)', () => {
    const forever = section('forever-score');
    const lead = forever.indexOf(scoreCopy.explainerLead);
    const sample = forever.indexOf('data-testid="ratings-forever-sample"');
    expect(lead).toBeGreaterThan(-1);
    expect(sample).toBeGreaterThan(lead);
    expect(sample).toBeLessThan(forever.indexOf(scoreCopy.explainerWhat));
    const card = forever.slice(sample);
    expect(card).toContain('data-testid="forever-score-card"');
    expect(card).toMatch(/class="pill pill-sample[^"]*">Sample</);
    expect(card).toMatch(/data-testid="score-badge"[^>]*>1830</);
  });

  it('says "0 to 2100" once', () => {
    expect(section('forever-score').match(/0 to 2100/g)).toHaveLength(1);
  });

  it('draws each part’s bar in proportion to its ceiling, with the tier strip in the same panel', () => {
    const widths = { gear: 71.4, experience: 100, parses: 85.7, attendance: 28.6, discipline: 14.3 };
    for (const [key, width] of Object.entries(widths)) {
      expect(html).toMatch(new RegExp(`style="width: ${width}%" data-testid="ratings-forever-bar-${key}"`));
    }
    const panel =
      /What moves each part[\s\S]*?data-testid="ratings-forever-tiers-note"/.exec(html)?.[0] ?? '';
    expect(panel).toContain('data-testid="ratings-forever-score-gear"');
    expect(panel).toContain('data-testid="tier-strip"');
    expect(html).not.toMatch(/>Tiers<\/h3>/);
  });

  it('reads the Gear blurb by level band', () => {
    expect(scoreCopy.components[0].moves).toContain(
      'How close your gear is to your spec’s best set for your level (20, 30, 40, 50 or 60)',
    );
  });

  it('bridges to the rating in one sentence, and says the two Preparations are the same thing (item 11)', () => {
    const rating = section('performance-rating');
    expect(rating).toContain('A rating scores one fight. The Forever Score counts all of them.');
    expect(rating).toContain(scoreCopy.explainerPreparationShared);
  });

  it('renders the bridge as running text with a space between its sentences: "them. Both" (fix round 3 item 1)', () => {
    const bridge = /<p[^>]*data-testid="ratings-rating-bridge"[^>]*>([\s\S]*?)<\/p>/.exec(html)?.[1] ?? '';
    expect(text(bridge)).toContain('them. Both');
    expect(bridge).not.toContain('them.Both');
  });

  // L8.5 ruling 1: the verdict carries the count the old "Last 6 fights" label did; the sample
  // has no Forever Score, so it has no Parses half.
  it('shows a Performance Rating sample, "65 of 100" with its verdict and sparkline', () => {
    const rating = section('performance-rating');
    const sample = rating.slice(rating.indexOf('data-testid="ratings-rating-sample"'));
    expect(sample).toMatch(/data-testid="character-rating-overall">65</);
    expect(sample).toContain('>of 100<');
    expect(sample).toMatch(/data-testid="rating-card-verdict">\s*Rated 65 of 100 over the last 6 fights\s*</);
    expect(sample).toContain('data-testid="rating-trend"');
    expect(sample).toMatch(/class="pill pill-sample[^"]*">Sample</);
  });

  it('leaves 24px between the rating’s opening and "How ratings work"', () => {
    expect(html).toMatch(/<section[^>]*id="performance-rating"[^>]*class="[^"]*gap-6[^"]*"/);
  });

  it('takes its labels from the fixture and its words from the copy file, never a literal', () => {
    // "Preparation" is also the rating's own component in its weights table, on purpose (item 11).
    for (const component of labels.components.filter((c) => c.key !== 'discipline')) {
      expect(source).not.toContain(`'${component.label}'`);
    }
    expect(source).not.toContain(scoreCopy.explainerWhat);
    expect(source).not.toContain(scoreCopy.explainerLead);
  });
});
