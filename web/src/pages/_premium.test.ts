// web/src/pages/_premium.test.ts
// Container renders of the plans page in the flat system (design/specs/2026-11-08-premium-account.md
// sections 4, 5, 14), through the shared `renderPremium` in test-support.
import { readFileSync } from 'node:fs';
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../lib/billing/plans-page-copy';
import { premiumSourceCopy } from '../lib/billing/premium-source-copy';
import { renderPremium } from '../test-support/render-premium';

const regions = (doc: Document): string[] =>
  [...doc.querySelectorAll('main [data-region]')].map((el) => el.getAttribute('data-region') ?? '');

describe('the plans page in the flat system', () => {
  it('names itself Plans once, in the one page title, with the lead and no slogan or eyebrow', async () => {
    const doc = await renderPremium();
    const titles = doc.querySelectorAll('main h1');
    expect(titles).toHaveLength(1);
    expect(titles[0]?.textContent?.trim()).toBe(plansPageCopy.title);
    const text = doc.querySelector('main')?.textContent ?? '';
    expect(text).toContain(plansPageCopy.band.lede);
    expect(text).not.toContain(plansPageCopy.band.heading);
  });

  it('carries the regions of spec 14 in order, with one Plate', async () => {
    const doc = await renderPremium();
    const order = regions(doc).filter((r) => !r.startsWith('head-') || r === 'head-answer');
    expect(order).toEqual([
      'head',
      'head-answer',
      'controls',
      'notes',
      'plate',
      'section-for-guild-leaders',
      'section-how-the-guild-plans-work',
      'section-questions',
      'section-what-the-coach-says-after-a-raid',
    ]);
    expect(doc.querySelectorAll('main [data-region=plate]')).toHaveLength(1);
  });

  it('the interval bar links to ?interval=, and the bar end sends a guild leader down the page', async () => {
    const doc = await renderPremium();
    const radios = [...doc.querySelectorAll('[data-param=interval] [role=radio]')];
    expect(radios.map((r) => r.textContent?.trim())).toEqual(['Yearly', 'Monthly']);
    const link = doc.querySelector('[data-testid=band-guild-link]');
    expect(link?.getAttribute('href')).toBe('#guild-leaders');
    expect(link?.textContent?.trim()).toBe(`${plansPageCopy.band.guildLink} ›`);
  });

  it('says founder once: the notice line, no Founder tag on a plan column', async () => {
    const doc = await renderPremium();
    expect(doc.querySelector('[data-testid=founder-row]')?.textContent).toContain(plansPageCopy.founder.open);
    expect(doc.querySelectorAll('[data-testid=plan-premium] .ui-plh-name')[0]?.textContent).not.toContain(
      plansPageCopy.founder.label,
    );
  });

  it('the Plate holds Free then Premium (Premium stacks first on a phone)', async () => {
    const plate = (await renderPremium()).querySelector('[data-region=plate]');
    const cols = plate?.querySelector('.ui-cols');
    expect(cols?.getAttribute('data-swap')).toBe('true');
    expect([...(cols?.children ?? [])].map((c) => c.getAttribute('data-testid'))).toEqual([
      'plan-free',
      'plan-premium',
    ]);
  });

  it('holds the three owned lines and See your plan beside the subscribe buttons', async () => {
    const card = (await renderPremium()).querySelector('[data-testid=plan-premium]');
    for (const [source, line] of Object.entries(premiumSourceCopy.premiumCard)) {
      expect(card?.querySelector(`.owned-line[data-source=${source}]`)?.textContent?.trim()).toBe(line);
    }
    const see = [...(card?.querySelectorAll('a') ?? [])].find(
      (a) => a.textContent?.trim() === premiumSourceCopy.buttons.seePlan,
    );
    expect(see?.getAttribute('href')).toBe('/account#plan');
    expect(card?.querySelectorAll('.plan-btn a[href^="/premium/checkout?plan=premium"]').length).toBe(4);
  });

  it('opens the rules and the questions (no disclosures)', async () => {
    const doc = await renderPremium();
    expect(doc.querySelectorAll('main details, main summary')).toHaveLength(0);
    const text = doc.querySelector('main')?.textContent ?? '';
    for (const rule of plansPageCopy.rules) expect(text).toContain(rule.sentence);
    for (const entry of plansPageCopy.faq) expect(text).toContain(entry.answer);
  });

  it('promises nothing unshipped: no chip-in, no trial email, the coach line first', async () => {
    const text = (await renderPremium()).body.textContent ?? '';
    expect(text).toContain(plansPageCopy.premium.features[0]?.text ?? 'missing');
    expect(text).not.toMatch(/chip in/i);
    expect(text).not.toMatch(/email you/i);
  });

  it('keeps one island, the layout character selector: the account read is a module imported only signed in', async () => {
    const doc = await renderPremium();
    expect(doc.querySelectorAll('astro-island')).toHaveLength(1);
  });
});

describe('the Premium card when the account has Premium', () => {
  const SOURCES = ['own', 'seat', 'grant'] as const;
  const sourceScripts = (doc: Document): Element[] =>
    [...doc.querySelectorAll('script:not([src])')].filter((s) => s.textContent?.includes('fs.premiumSource'));

  describe('the hint script (Header.astro owns it, so /b/{id} has it too; spec 2026-10-29-premium-entry section 4)', () => {
    const runOn = async (storageOf: () => Storage, html = '<html>'): Promise<DOMStringMap> => {
      const text = sourceScripts(await renderPremium())[0]?.textContent ?? '';
      const doc = new JSDOM(`<!doctype html>${html}</html>`).window.document;
      // The script reads the `localStorage` and `document` globals: hand it fakes of both.
      const lazyStorage = { getItem: (key: string) => storageOf().getItem(key) };
      new Function('localStorage', 'document', text)(lazyStorage, doc);
      return doc.documentElement.dataset;
    };
    const run = async (storageOf: () => Storage): Promise<string | undefined> =>
      (await runOn(storageOf)).premium;

    it.each(SOURCES)('%s hides the header Premium tag before paint', async (word) => {
      expect((await runOn(storing(word))).premiumNav).toBe('hide');
    });

    it('server sims from the cached /v1/me (the fallback) hide the tag too', async () => {
      expect((await runOn(storing(null), '<html data-server-sims="1">')).premiumNav).toBe('hide');
    });

    it('no hint and no server sims: the tag shows', async () => {
      expect((await runOn(storing(null))).premiumNav).toBeUndefined();
    });
    const storing = (value: string | null) => (): Storage => ({ getItem: () => value }) as unknown as Storage;

    it.each(SOURCES)('%s sets the attribute', async (word) => {
      expect(await run(storing(word))).toBe(word);
    });

    it.each([null, '', 'none', 'fallback', 'GRANT', 'own_then_seat'])('%s is ignored', async (junk) => {
      expect(await run(storing(junk))).toBeUndefined();
    });

    it('a throwing localStorage does not throw', async () => {
      const throwing = (): Storage => {
        throw new Error('blocked');
      };
      expect(await run(throwing)).toBeUndefined();
    });
  });
});

describe('the /premium anchors (spec 2026-10-29-premium-entry section 8)', () => {
  it('every anchor a link lands on exists once', async () => {
    const doc = await renderPremium();
    for (const id of ['premium', 'coach', 'server-sims', 'history', 'notifications', 'guild-leaders']) {
      expect(doc.querySelectorAll(`[id="${id}"]`)).toHaveLength(1);
    }
    expect(doc.getElementById('premium')?.getAttribute('data-testid')).toBe('plan-premium');
  });

  it('each anchor is a Premium feature line, and server-sims marks lines 2 to 4', async () => {
    const card = (await renderPremium()).querySelector('[data-testid=premium-features]')!;
    const lines = [...card.querySelectorAll('li')];
    expect(lines.map((li) => li.getAttribute('data-anchor'))).toEqual([
      'coach',
      'server-sims',
      'server-sims',
      'server-sims',
      'history',
      'notifications',
      null,
    ]);
    expect(lines.map((li) => li.id)).toEqual([
      'coach',
      'server-sims',
      '',
      '',
      'history',
      'notifications',
      '',
    ]);
  });

  it('the landing highlight is CSS on :target', () => {
    const source = readFileSync(new URL('../components/premium/PlanFeatures.astro', import.meta.url), 'utf8');
    for (const id of ['coach', 'server-sims', 'history', 'notifications']) {
      expect(source).toContain(`.plan-features:has(#${id}:target) :global([data-anchor='${id}'])`);
    }
  });
});
