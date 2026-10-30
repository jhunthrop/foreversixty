// web/src/components/RatingCard.test.ts
// Fix round 2 item 6: the Performance Rating in the Forever Score card's shell -- "65 of 100",
// the labelled sparkline with the data's own count, and the weakest part of a fight.
import { readdirSync, readFileSync, statSync } from 'node:fs';
import path from 'node:path';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../lib/rating/copy';
import RatingCard from './RatingCard.svelte';

const TREND = [56, 58, 60, 61, 63, 65].map((overall, index) => ({
  fought_at: `2026-12-0${index + 4}T20:00:00Z`,
  overall,
}));

/** The class tokens of the element carrying `data-testid="<testid>"`. */
function classOf(body: string, testid: string): string[] {
  return new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`).exec(body)?.[1].split(' ') ?? [];
}

function text(html: string): string {
  return html
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/\s+/g, ' ')
    .trim();
}

describe('RatingCard', () => {
  it('reads "65 of 100" in the card shell', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    const card = /<div class="([^"]*)" data-testid="rating-card"/.exec(body)?.[1] ?? '';
    expect(card.split(' ')).toEqual(expect.arrayContaining(['card-surface']));
    expect(body).toMatch(/data-testid="character-rating-overall">65</);
    expect(body).toContain(`>${ratingCopy.outOf100}<`);
  });

  it('names the weakest part of a fight in one line', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    const line = /data-testid="rating-card-weakest">([\s\S]*?)<\/p>/.exec(body)?.[1] ?? '';
    expect(text(line)).toBe('Weakest part of a fight: Activity.');
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

  // L8.4 fix round 5 item 2: the inline link takes its gold from the base `a` rule, so the
  // design system's hover and pressed --gold-hover apply; a text-gold utility would pin it.
  it('leaves the inline link’s colour to the base rule: no text-gold utility (fix round 5 item 2)', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    const links = [...body.matchAll(/<a class="([^"]*)"/g)].map((match) => match[1].split(' '));
    expect(links.length).toBeGreaterThan(0);
    for (const link of links) {
      expect(link).toContain('underline');
      expect(link).not.toContain('text-gold');
    }
  });

  // L8.5 ruling 1: the Forever Score card's treatment. From 1280 two columns: the number, its
  // segments and the verdict left; the sparkline filling the right column, captioned by the
  // same pure model as the score's ("Up 9 since 4 Dec", the newest date at the right).
  it('splits into two columns from 1280, the sparkline filling the right one (L8.5 ruling 1)', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    expect(classOf(body, 'rating-card-columns')).toEqual(
      expect.arrayContaining(['flex', 'flex-col', 'xl:grid', 'xl:grid-cols-2', 'xl:gap-x-8']),
    );
    const main = body.indexOf('data-testid="rating-card-main"');
    const side = body.indexOf('data-testid="rating-card-side"');
    expect(main).toBeGreaterThan(-1);
    expect(side).toBeGreaterThan(main);
    // The number and the verdict on the left, the sparkline on the right.
    expect(body.indexOf('character-rating-overall')).toBeGreaterThan(main);
    expect(body.indexOf('character-rating-overall')).toBeLessThan(side);
    expect(body.indexOf('rating-card-verdict')).toBeLessThan(side);
    expect(body.indexOf('data-testid="rating-trend"')).toBeGreaterThan(side);
    // The score card's plot: the whole column wide at a fixed height, stretched, an 8px end dot.
    expect(classOf(body, 'rating-card-history')).toEqual(expect.arrayContaining(['w-full']));
    expect(classOf(body, 'rating-card-history')).not.toContain('max-w-[460px]');
    expect(classOf(body, 'rating-card-plot')).toEqual(
      expect.arrayContaining(['relative', 'h-10', 'lg:h-12']),
    );
    expect(body).toMatch(
      /<svg class="absolute inset-0 block h-full w-full overflow-visible"[^>]*preserveAspectRatio="none"/,
    );
    expect(body).not.toMatch(/data-testid="rating-trend"[^>]*width="240"/);
    expect(classOf(body, 'rating-card-sparkline-last')).toEqual(
      expect.arrayContaining(['size-2', 'rounded-full']),
    );
  });

  it('captions the sparkline "Up 9 since 4 Dec" with the newest date at the right (L8.5 ruling 1)', () => {
    const { body } = render(RatingCard, { props: { overall: 65, trend: TREND, sampleSize: 6 } });
    expect(body).toMatch(/data-testid="rating-card-history-caption">Up 9 since 4 Dec</);
    expect(body).toMatch(/data-testid="rating-card-history-last">9 Dec</);
    expect(body).toContain('aria-label="Performance Rating, last 6 fights: 56 on 4 Dec, 65 on 9 Dec"');
    expect(body).toMatch(/<table class="sr-only" data-testid="rating-card-history-table">/);
    // The verdict carries the count now; the old "Last 6 fights" label is gone.
    expect(body).not.toContain('rating-card-trend-label');
  });

  // L8.5 fix round 1 items 1 and 2: the line says what the number is -- the latest fight's
  // rating -- and how many fights are rated, never an average over them; and it makes no claim
  // about the Forever Score's Parses, which comes from parse percentiles, not this rating.
  it('says the latest fight’s rating and how many fights are rated (fix round 1 items 1 and 2)', () => {
    expect(ratingCopy.verdict(65, 6)).toBe('Latest fight rated 65 of 100 · 6 fights rated');
    expect(ratingCopy.verdict(65, 1)).toBe('Latest fight rated 65 of 100 · 1 fight rated');
    expect('countsToward' in ratingCopy).toBe(false);
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    });
    const verdict = /<p class="([^"]*)" data-testid="rating-card-verdict">([\s\S]*?)<\/p>/.exec(body);
    expect(text(verdict?.[2] ?? '')).toBe('Latest fight rated 65 of 100 · 6 fights rated');
    expect(verdict?.[1].split(' ')).toEqual(expect.arrayContaining(['text-[14px]']));
    expect(body).not.toMatch(/counts toward|Parses/);
    // No overall, no line.
    const none = render(RatingCard, { props: { overall: null, trend: TREND, sampleSize: 6 } }).body;
    expect(none).not.toContain('rating-card-verdict');
  });

  it('keeps the line beside the too-few line when there is no sparkline yet', () => {
    const few = render(RatingCard, {
      props: { overall: 50, trend: TREND.slice(0, 2), sampleSize: 2 },
    }).body;
    expect(text(/data-testid="rating-card-verdict">([\s\S]*?)<\/p>/.exec(few)?.[1] ?? '')).toBe(
      'Latest fight rated 50 of 100 · 2 fights rated',
    );
    expect(few.indexOf('character-rating-too-few')).toBeGreaterThan(few.indexOf('rating-card-side'));
  });

  // Fix round 2 item 1: the line names the fight -- its encounter and its day-month date --
  // "Latest fight (Warden Kelthas, 9 Dec) rated 65 of 100 · 6 fights rated"; the date alone when
  // the encounter is unknown.
  it('names the latest fight: encounter and date (fix round 2 item 1)', () => {
    expect(ratingCopy.verdict(65, 6, { encounter: 'Warden Kelthas', date: '9 Dec' })).toBe(
      'Latest fight (Warden Kelthas, 9 Dec) rated 65 of 100 · 6 fights rated',
    );
    expect(ratingCopy.verdict(65, 6, { encounter: null, date: '9 Dec' })).toBe(
      'Latest fight (9 Dec) rated 65 of 100 · 6 fights rated',
    );
    expect(ratingCopy.verdict(65, 6, null)).toBe('Latest fight rated 65 of 100 · 6 fights rated');
    const { body } = render(RatingCard, {
      props: {
        overall: 65,
        trend: TREND,
        sampleSize: 6,
        latestFight: { encounter: 'Warden Kelthas', date: '9 Dec' },
      },
    });
    expect(text(/data-testid="rating-card-verdict">([\s\S]*?)<\/p>/.exec(body)?.[1] ?? '')).toBe(
      'Latest fight (Warden Kelthas, 9 Dec) rated 65 of 100 · 6 fights rated',
    );
  });

  // Fix round 2 item 7: under 360px the line breaks at the dot, "6 fights rated" on a line of
  // its own; "rated 65 of 100" never breaks inside, so a long fight name breaks before "rated".
  it('breaks at the dot under 360px, and before "rated" when the name needs it (fix round 2 item 7)', () => {
    const { body } = render(RatingCard, {
      props: {
        overall: 65,
        trend: TREND,
        sampleSize: 6,
        latestFight: { encounter: 'Warden Kelthas', date: '9 Dec' },
      },
    });
    const line = /data-testid="rating-card-verdict">([\s\S]*?)<\/p>/.exec(body)?.[1] ?? '';
    const spans = [...line.matchAll(/<span class="([^"]*)"[^>]*>([^<]*)<\/span>/g)].map((m) => ({
      classes: m[1].split(' '),
      text: m[2],
    }));
    expect(spans.map((span) => span.text)).toEqual([
      'Latest fight (Warden Kelthas, 9 Dec)',
      'rated 65 of 100',
      ' · ',
      '6 fights rated',
    ]);
    expect(spans[1].classes).toContain('whitespace-nowrap');
    expect(spans[2].classes).toContain('max-[360px]:hidden');
    expect(spans[3].classes).toEqual(expect.arrayContaining(['whitespace-nowrap', 'max-[360px]:block']));
    // The words between "…9 Dec)" and "rated" are a plain space, the only break before "rated".
    expect(line).toMatch(/9 Dec\)<\/span> <span/);
  });

  // Fix round 2 item 2: what a parse is next to what the rating is, one muted line under it.
  it('says how a parse differs from the rating, in one muted line under it (fix round 2 item 2)', () => {
    expect(ratingCopy.parseVsRating).toBe(
      'A parse is your rank against other players; the rating is how well the fight was played.',
    );
    const { body } = render(RatingCard, { props: { overall: 65, trend: TREND, sampleSize: 6 } });
    const note = /<p class="([^"]*)" data-testid="rating-card-parse-vs-rating">([^<]*)<\/p>/.exec(body);
    expect(note?.[2]).toBe(ratingCopy.parseVsRating);
    expect(note?.[1].split(' ')).toEqual(expect.arrayContaining(['text-muted', 'text-[12px]']));
    expect(body.indexOf('rating-card-verdict')).toBeLessThan(body.indexOf('rating-card-parse-vs-rating'));
    expect(body.indexOf('rating-card-parse-vs-rating')).toBeLessThan(body.indexOf('rating-card-side'));
  });

  it('leaves no "counts toward" claim in any string the site shows (fix round 1 item 2)', () => {
    const root = path.resolve(import.meta.dirname, '..');
    const sources = (dir: string): string[] =>
      readdirSync(dir).flatMap((name) => {
        const full = path.join(dir, name);
        if (statSync(full).isDirectory()) return sources(full);
        return /\.(ts|svelte|astro|md|mdx|html)$/.test(name) && !/\.test\.ts$/.test(name) ? [full] : [];
      });
    // A string literal or markup text that says it; a code comment about scrollWidth does not
    // (SimTabs.astro: "trailing padding counts toward scrollWidth").
    const claim = /(['"`>][^'"`<\n]*counts toward|counts toward (Parses|the Forever Score))/i;
    const offenders = sources(root).filter((file) => claim.test(readFileSync(file, 'utf8')));
    expect(offenders).toEqual([]);
  });

  // Fix round 1 item 5: the split keys off the card's own room, not the window: /ratings mounts
  // the sample in a 560px column, where two columns from a 1280px window would be cramped.
  it('stays one column when its container is narrow: columns={false} (fix round 1 item 5)', () => {
    const { body } = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, sample: true, columns: false },
    });
    const columns = classOf(body, 'rating-card-columns');
    expect(columns).toEqual(expect.arrayContaining(['flex', 'flex-col']));
    expect(columns.filter((token) => token.startsWith('xl:'))).toEqual([]);
  });

  it('links the weakest part to its fight’s row in "Every ranked fight" (L8.5 ruling 1)', () => {
    const { body } = render(RatingCard, {
      props: {
        overall: 65,
        trend: TREND,
        sampleSize: 6,
        worst: 'activity',
        weakestHref: '#fight-fixture2abcd-3',
      },
    });
    const line = /<p class="[^"]*" data-testid="rating-card-weakest">([\s\S]*?)<\/p>/.exec(body)?.[1] ?? '';
    expect(text(line)).toBe('Weakest part of a fight: Activity.');
    const link = /<a class="([^"]*)" href="#fight-fixture2abcd-3"[^>]*>Activity<\/a>/.exec(line);
    expect(link).not.toBeNull();
    // The score card's sentence link: a 44px phone target that stacks above the line under it.
    expect(link?.[1].split(' ')).toEqual(
      expect.arrayContaining(['relative', 'z-10', 'underline', 'max-md:py-[13.5px]']),
    );
    // No row to go to, no link.
    const plain = render(RatingCard, {
      props: { overall: 65, trend: TREND, sampleSize: 6, worst: 'activity' },
    }).body;
    expect(/data-testid="rating-card-weakest">([\s\S]*?)<\/p>/.exec(plain)?.[1]).not.toContain('<a');
  });
});
