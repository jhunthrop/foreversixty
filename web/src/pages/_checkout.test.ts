// web/src/pages/_checkout.test.ts
// The checkout page's pre-paint frame (spec 2026-10-27-subscribe-and-join §7, §11): the head
// script marks the guild step from the URL before first paint, with readMeGuildCount's card
// count, so the CSS holds the step's height and hydration moves nothing. The script text is
// run against fake location, storage and document globals, as _premium.test.ts does.
import { readFileSync } from 'node:fs';
import { JSDOM } from 'jsdom';
import { afterEach, describe, expect, it, vi } from 'vitest';
import { ME_GUILD_COUNT_KEY, readMeGuildCount } from '../lib/account/premium-hint';
import { createPageContainer } from '../test-support/page-container';
import Checkout from './premium/checkout.astro';

async function renderCheckout(): Promise<Document> {
  const container = await createPageContainer();
  return new JSDOM(await container.renderToString(Checkout)).window.document;
}

function frameScript(doc: Document): string {
  const scripts = [...doc.head.querySelectorAll('script:not([src])')].filter((script) =>
    script.textContent?.includes('checkoutStep'),
  );
  expect(scripts).toHaveLength(1);
  return scripts[0]?.textContent ?? '';
}

interface Run {
  step: string | undefined;
  cards: string;
}

async function run(search: string, stored: string | null | (() => never) = null): Promise<Run> {
  const text = frameScript(await renderCheckout());
  const doc = new JSDOM('<!doctype html><html></html>').window.document;
  const storage = {
    getItem: (): string | null => (typeof stored === 'function' ? stored() : stored),
  };
  new Function('location', 'localStorage', 'document', text)({ search }, storage, doc);
  return {
    step: doc.documentElement.dataset.checkoutStep,
    cards: doc.documentElement.style.getPropertyValue('--checkout-guild-cards'),
  };
}

afterEach(() => vi.unstubAllGlobals());

describe('checkout page frame: the head script', () => {
  it.each(['?plan=guild&interval=yearly', '?plan=raid_guild&interval=monthly&status=x'])(
    '%s marks the guild step before paint',
    async (search) => {
      expect((await run(search)).step).toBe('guild');
    },
  );

  it.each([
    '?plan=premium&interval=yearly',
    '?plan=guild&interval=yearly&guild_id=4',
    '?plan=guild&interval=yearly&guild_id=4&status=success',
    '?plan=guild&interval=weekly',
    '?plan=guild&interval=yearly&comeback=9',
    '?plan=guild&interval=yearly&status=success',
    '',
  ])('%s leaves the page unmarked (the launcher draws the narrow body)', async (search) => {
    expect((await run(search)).step).toBeUndefined();
  });

  it.each([null, '', '0', '-3', '1', '3', '9', 'x', '2.6'])(
    'stored %s: the reserve counts exactly the cards readMeGuildCount gives the skeleton',
    async (stored) => {
      const items = new Map(stored === null ? [] : [[ME_GUILD_COUNT_KEY, stored]]);
      vi.stubGlobal('localStorage', { getItem: (key: string) => items.get(key) ?? null });
      const expected = String(readMeGuildCount());
      expect((await run('?plan=guild&interval=yearly', stored)).cards).toBe(expected);
    },
  );

  it('a throwing localStorage still marks the step, with the default count', async () => {
    const blocked = (): never => {
      throw new Error('blocked');
    };
    expect(await run('?plan=guild&interval=yearly', blocked)).toEqual({ step: 'guild', cards: '2' });
  });
});

describe('checkout page frame: the CSS', () => {
  const source = readFileSync(new URL('./premium/checkout.astro', import.meta.url), 'utf8');

  it('holds the step height at each breakpoint from guildFrameReserve and the card count', () => {
    for (const [fixed, card] of [
      ['phoneFixed', 'phoneCard'],
      ['mdFixed', 'mdCard'],
      ['lgFixed', 'lgCard'],
    ]) {
      expect(source).toContain(
        `min-height: calc((var(--${fixed}) + var(--checkout-guild-cards) * var(--${card})) * 1px);`,
      );
    }
    expect(source).toMatch(/@media \(min-width: 768px\)/);
    expect(source).toMatch(/@media \(min-width: 1024px\)/);
  });

  it('hides the static narrow body under the guild mark, never removes it', () => {
    expect(source).toMatch(/data-checkout-step='guild'.*checkout-body.*\{\s*visibility: hidden;/s);
  });

  it('sets the site measures to the page columns: 640 message and panel, 880 guild step', () => {
    expect(source).toMatch(/\.checkout-main \{\s*--measure-form: 640px;\s*--measure-wide: 880px;\s*\}/);
  });

  it('main carries the class the CSS keys on and no measure', async () => {
    const main = (await renderCheckout()).querySelector('main#main');
    expect(main?.className.split(/\s+/)).toEqual(expect.arrayContaining(['checkout-main', 'w-full']));
    expect(main?.className).not.toContain('measure-form');
  });
});
