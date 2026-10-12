// web/src/components/RatingCard.test.ts
// Fix round 2 item 6: the Performance Rating in the Forever Score card's shell -- "65 of 100",
// the labelled sparkline with the data's own count, and the weakest part of a fight.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../lib/rating/copy';
import RatingCard from './RatingCard.svelte';

const TREND = [40, 45, 50, 55, 60, 65].map((overall, index) => ({
  fought_at: `2026-12-0${index + 1}T00:00:00Z`,
  overall,
}));

describe('RatingCard', () => {
  it('reads "65 of 100" in the card shell', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    const card = /<div class="([^"]*)" data-testid="rating-card"/.exec(body)?.[1] ?? '';
    expect(card.split(' ')).toEqual(
      expect.arrayContaining(['bg-raised', 'border-line', 'rounded-panel', 'border']),
    );
    expect(body).toMatch(/data-testid="character-rating-overall">65</);
    expect(body).toContain(`>${ratingCopy.outOf100}<`);
  });

  it('labels the sparkline with the count from the data: "Last 6 fights"', () => {
    const { body } = render(RatingCard, { props: { overall: 65, trend: TREND, sampleSize: 6 } });
    expect(body).toMatch(/data-testid="rating-card-trend-label">\s*Last 6 fights\s*</);
    expect(body).toContain('data-testid="rating-trend"');
    const three = render(RatingCard, {
      props: { overall: 65, trend: TREND.slice(0, 3), sampleSize: 9 },
    }).body;
    expect(three).toMatch(/>\s*Last 3 fights\s*</);
  });

  it('names the weakest part of a fight in one line', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    expect(body).toMatch(/data-testid="rating-card-weakest">\s*Weakest part of a fight: Activity\.\s*</);
    expect(body).not.toContain('Best:');
  });

  // L8.4 fix round 4 item 3: "Activity" is defined the way Execution is in the fight tables, a
  // one-line muted 12px caption from the copy module, visible at every width (item 6: no
  // sr-only, no title).
  it('defines the part it names in one muted 12px line under it, at every width (fix round 4 item 3)', () => {
    expect(ratingCopy.componentDefinition('activity')).toBe(
      'Activity: the share of the fight you spent doing something rather than idle, leaving out the times the fight expects you to wait.',
    );
    for (const part of ['output', 'survival', 'mechanics', 'utility', 'preparation', 'activity']) {
      expect(ratingCopy.componentDefinition(part), part).toMatch(/^[A-Z][a-z]+: .+\.$/);
    }
    expect(ratingCopy.componentDefinition('threat')).toBeNull();
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    const definition = /<p class="([^"]*)" data-testid="rating-card-weakest-definition">([^<]*)<\/p>/.exec(
      body,
    );
    expect(definition).not.toBeNull();
    expect(definition?.[2].trim()).toBe(ratingCopy.componentDefinition('activity'));
    const tokens = definition?.[1].split(' ') ?? [];
    expect(tokens).toEqual(expect.arrayContaining(['text-muted', 'text-[12px]']));
    expect(tokens.filter((token) => token.includes('sr-only') || token.includes('hidden'))).toEqual([]);
    expect(body.indexOf('rating-card-weakest"')).toBeLessThan(body.indexOf('rating-card-weakest-definition'));
    expect(body).not.toContain(`title="${ratingCopy.componentDefinition('activity')}"`);
    // No weakest part, no definition.
    const none = render(RatingCard, { props: { overall: 65, trend: TREND, sampleSize: 6 } }).body;
    expect(none).not.toContain('rating-card-weakest-definition');
  });

  it('says how many fights it needs below the trend threshold, and carries the Sample pill as a sample', () => {
    const few = render(RatingCard, { props: { overall: 50, trend: TREND.slice(0, 2), sampleSize: 2 } }).body;
    expect(few).toContain('data-testid="character-rating-too-few"');
    expect(few).not.toContain('rating-card-trend');
    const sample = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, sample: true },
    }).body;
    expect(sample).toMatch(/class="pill pill-sample[^"]*">Sample</);
    expect(sample).not.toContain('How is this calculated?');
  });
});
