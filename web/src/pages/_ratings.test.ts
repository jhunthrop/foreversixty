// web/src/pages/_ratings.test.ts
// Underscore-prefixed like _setup.test.ts: everything else under src/pages/ is a route.
//
// Task L8.2 ruling 6: /ratings opens with the Forever Score -- what it is, what moves each of
// its five parts, and the tier thresholds -- above "How ratings work". The five labels are the
// addon's own (Score.lines), read from the shared fixture both sides pin.
import { readFileSync } from 'node:fs';
import { beforeAll, describe, expect, it } from 'vitest';
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import { scoreCopy } from '../lib/score/copy';
import { tierRanges } from '../lib/score/score';
import { createPageContainer } from '../test-support/page-container';
import RatingsPage from './ratings.astro';

const source = readFileSync(new URL('./ratings.astro', import.meta.url), 'utf8');

let html: string;

beforeAll(async () => {
  const container = await createPageContainer();
  html = await container.renderToString(RatingsPage);
});

describe('/ratings Forever Score section', () => {
  it('sits above "How ratings work"', () => {
    const score = html.indexOf('data-testid="ratings-forever-score"');
    expect(score).toBeGreaterThan(-1);
    expect(score).toBeLessThan(html.indexOf('How ratings work'));
  });

  it('says what the score is in one paragraph', () => {
    expect(html).toContain(scoreCopy.explainerWhat);
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

  it('says how parses rank in plain words, never "percentile"', () => {
    const parses = scoreCopy.components.find((c) => c.key === 'parses');
    expect(parses?.moves).not.toMatch(/percentile/i);
    expect(parses?.moves).toMatch(/how your parses rank/i);
  });

  it('shows the tier thresholds as a table in the tier colours', () => {
    const table = /<table[^>]*data-testid="ratings-forever-tiers"[\s\S]*?<\/table>/.exec(html)?.[0] ?? '';
    expect(table).not.toBe('');
    for (const tier of Object.values(labels.tiers)) {
      expect(table).toContain(tier.label);
      expect(table).toContain(`color: ${tier.color}`);
    }
    for (const { range } of tierRanges()) expect(table).toContain(range);
  });

  it('shows an example with the same badge the profile uses', () => {
    expect(html).toMatch(/data-testid="score-badge"[^>]*>1830</);
  });

  it('takes its labels from the fixture and its words from the copy file, never a literal', () => {
    for (const component of labels.components) expect(source).not.toContain(`'${component.label}'`);
    expect(source).not.toContain(scoreCopy.explainerWhat);
  });
});
