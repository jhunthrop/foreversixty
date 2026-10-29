// web/src/pages/_index-hero.test.ts
// The landing's example card (W.10, UX review r6 F6-4, player review P6-10): one label that says
// whose list it is, "levels 20, 30, 40, 50, 60" (never "bands"), and the dates strip's live dot on
// the tile that matches the card. Underscore-prefixed like the other page tests.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { createPageContainer } from '../test-support/page-container';
import { beforeAll, describe, expect, it } from 'vitest';
import dates from '../data/dates.json';
import { homeHeroCopy, homeHeroStatePanelCopy } from '../lib/home-landing-copy';
import { heroCardFor } from '../lib/home/hero-card';
import { dotIndex, timelineDisplayRows, timelineRows } from '../lib/home/timeline';
import Index from './index.astro';

let container: AstroContainer;
let html = '';

beforeAll(async () => {
  container = await createPageContainer();
  html = await container.renderToString(Index);
});

const textOf = (markup: string): string =>
  markup
    .replace(/<script[^]*?<[/]script>/g, ' ')
    .replace(/<style[^]*?<[/]style>/g, ' ')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();

describe('the landing launch card', () => {
  it('has no example label: the card links to the real Hunter list', () => {
    expect(html).not.toContain('home-hero-label');
    expect(textOf(html)).not.toMatch(/Example:/);
    const link = /<a[^>]*data-testid="home-hero-link"[^>]*>/.exec(html)?.[0] ?? '';
    expect(link).toContain(`href="${homeHeroStatePanelCopy.linkHref}"`);
  });

  it('shows the countdown as three stat tiles, Days, Hours and Minutes', () => {
    for (const unit of ['days', 'hours', 'minutes']) {
      expect(html).toContain(`data-testid="home-hero-countdown-${unit}"`);
    }
    expect(textOf(/data-testid="home-hero-countdown"[^]*?home-hero-link/.exec(html)?.[0] ?? '')).toMatch(
      /Days.*Hours.*Minutes/,
    );
  });

  it('says "levels" in both variants, one phrase shape, and never "bands"', () => {
    expect(homeHeroStatePanelCopy.betaBandsLine).toBe('Leveling best in slot, levels 10 to 30');
    expect(homeHeroStatePanelCopy.bandsLine).toBe('Leveling best in slot, levels 20, 30, 40, 50, 60');
    const after = heroCardFor(dates, new Date('2026-10-22T08:00:00Z'));
    expect(after.bandsLine).toBe('Leveling best in slot, levels 20, 30, 40, 50, 60');
    expect(textOf(html)).not.toMatch(/\bbands?\b/i);
  });

  it('puts the live dot on the Beta tile while the card says "Beta is live"', () => {
    const now = new Date('2026-10-20T12:00:00Z');
    expect(heroCardFor(dates, now).eyebrow).toBe('Beta is live');
    const shown = timelineDisplayRows(timelineRows(dates, now));
    const at = dotIndex(shown);
    expect(at).not.toBeNull();
    expect(shown[at ?? 0].role).toBe('beta');
    // After the beta only the next date carries it.
    const later = timelineDisplayRows(timelineRows(dates, new Date('2026-10-22T08:00:00Z')));
    expect(later[dotIndex(later) ?? 0].value).toBe('Name reservation');
    expect(dotIndex([])).toBeNull();
  });
});

describe('the hero headline', () => {
  it('is two-tone: the lead in --strong and the accent word in gold', () => {
    const h1 = /<h1[^>]*>([^]*?)<[/]h1>/.exec(html)?.[1] ?? '';
    expect(h1).toMatch(/<span class="text-strong">Play your class<[/]span>/);
    expect(h1).toMatch(/<span class="gold-text"[^>]*>better[.]<[/]span>/);
    expect(textOf(h1)).toBe(homeHeroCopy.headline);
  });

  it('keeps the heading recipe and no longer paints the whole line gold', () => {
    expect(/<h1[^>]*>/.exec(html)?.[0]).toContain('heading-hero');
    expect(/<h1[^>]*>/.exec(html)?.[0]).not.toContain('gold-text');
  });
});
