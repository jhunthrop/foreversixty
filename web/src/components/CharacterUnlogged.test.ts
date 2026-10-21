// web/src/components/CharacterUnlogged.test.ts
// L8.4 fix round 3 item 1 (D1): the 404 no-logs page sorted the Forever Score card (its
// section is `-order-2` at every width since fix round 2 item 7) above the character header,
// which carried no order. The unlogged header now carries the logged header's own `-order-3`.
// The page is a flex column, so what a reader sees is the children sorted by their `order`
// at that width (then source order): this works that out from the classes for a phone and a
// desktop and pins it. The owner's no-logs card still says "Upload logs" (ruling 4).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { scoreCopy } from '../lib/score/copy';
import CharacterUnlogged from './CharacterUnlogged.svelte';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'elyra-duskvale' };

/** The CSS `order` a class list gives below md (a phone) or from md (a desktop). */
function orderAt(classes: string, width: 'phone' | 'desktop'): number {
  let order = 0;
  for (const token of classes.split(/\s+/)) {
    const match = /^(?:(max-md|md|lg|xl):)?(-?)order-(\d+|first|last|none)$/.exec(token);
    if (match === null) continue;
    const [, prefix, negative, value] = match;
    const applies = prefix === undefined || (width === 'phone' ? prefix === 'max-md' : prefix !== 'max-md');
    if (!applies) continue;
    const n = value === 'first' ? -9999 : value === 'last' ? 9999 : value === 'none' ? 0 : Number(value);
    order = negative === '-' ? -n : n;
  }
  return order;
}

/** The page's own blocks (its flex column's children), each by the opening tag that starts it:
 *  the header, the Forever Score section, the rating panel's section and the In game section. */
const BLOCKS: Record<string, RegExp> = {
  header: /<header class="([^"]*)"/,
  'forever-score': /<section class="([^"]*)"[^>]*data-testid="forever-score"/,
  'character-rating-loading': /<section class="([^"]*)"[^>]*data-testid="character-rating-loading"/,
  'character-in-game': /<section class="([^"]*)"[^>]*data-testid="character-in-game"/,
};

/** The blocks in the order they are drawn at that width: by `order`, then source order. */
function drawnOrder(body: string, width: 'phone' | 'desktop'): string[] {
  return Object.entries(BLOCKS)
    .map(([name, pattern]) => {
      const match = pattern.exec(body);
      if (match === null) throw new Error(`no ${name} block`);
      return { name, order: orderAt(match[1], width), index: match.index };
    })
    .sort((a, b) => a.order - b.order || a.index - b.index)
    .map((block) => block.name);
}

describe('CharacterUnlogged (the 404 no-logs page)', () => {
  it('draws the character header first, then the Forever Score card, on a phone and on a desktop', () => {
    const { body } = render(CharacterUnlogged, {
      props: { path: PATH, name: 'Elyra Duskvale', owned: false },
    });
    for (const width of ['phone', 'desktop'] as const) {
      const drawn = drawnOrder(body, width);
      expect(drawn.slice(0, 2), width).toEqual(['header', 'forever-score']);
      expect(drawn, width).toContain('character-in-game');
    }
  });

  it('still gives its owner "Upload logs", and a visitor the lookup', () => {
    const owner = render(CharacterUnlogged, {
      props: { path: PATH, name: 'Elyra Duskvale', owned: true },
    }).body;
    expect(owner).toContain(`>${scoreCopy.noLogsUpload}</a>`);
    expect(owner).toContain(`href="${scoreCopy.noLogsUploadHref}"`);
    const visitor = render(CharacterUnlogged, {
      props: { path: PATH, name: 'Elyra Duskvale', owned: false },
    }).body;
    // A visitor's "Upload logs" is a link in the next-step line (fix round 4 item 5), never the
    // owner's button.
    expect(visitor).not.toContain(`data-testid="forever-score-empty-action">${scoreCopy.noLogsUpload}<`);
    expect(visitor).toContain(scoreCopy.noLogsCompanion);
    expect(visitor).not.toContain('Check the name');
    expect(visitor).toContain('data-testid="character-ruleset-chip"');
  });

  // L8.4 fix round 4 item 5: the 404 page is not a dead end for a visitor.
  it('gives a visitor a next step under the no-logs message, and leaves the owner its button alone', () => {
    const visitor = render(CharacterUnlogged, {
      props: { path: PATH, name: 'Elyra Duskvale', owned: false },
    }).body;
    const next = /data-testid="forever-score-empty-next">([\s\S]*?)<\/p>/.exec(visitor)?.[1] ?? '';
    expect(next).toContain(`href="${scoreCopy.noLogsNext.addonHref}"`);
    expect(next).toContain(`>${scoreCopy.noLogsNext.addon}</a>`);
    expect(next).toContain(`href="${scoreCopy.noLogsNext.uploadHref}"`);
    expect(next).toContain(`>${scoreCopy.noLogsNext.upload}</a>`);
    const owner = render(CharacterUnlogged, {
      props: { path: PATH, name: 'Elyra Duskvale', owned: true },
    }).body;
    expect(owner).not.toContain('forever-score-empty-next');
    expect(owner).toMatch(/href="\/setup"[^>]*data-testid="forever-score-empty-action">Upload logs</);
  });
});
