// web/src/pages/_premium.test.ts
// Container renders of the plans page on the round 5 system (directions mock b4-premium; parity twin
// design/mocks/ws_premium.py; design/specs/2026-11-08-premium-account.md sections 5, 14), through the
// shared `renderPremium` in test-support.
import { readFileSync } from 'node:fs';
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../lib/billing/plans-page-copy';
import { premiumSourceCopy } from '../lib/billing/premium-source-copy';
import { renderPremium } from '../test-support/render-premium';

const regions = (doc: Document): string[] =>
  [...doc.querySelectorAll('main [data-region]')].map((el) => el.getAttribute('data-region') ?? '');

describe('the plans page on the round 5 system', () => {
  it('names itself Plans once, in the compact title band, with no lead, slogan or eyebrow', async () => {
    const doc = await renderPremium();
    const titles = doc.querySelectorAll('main h1');
    expect(titles).toHaveLength(1);
    expect(titles[0]?.textContent?.trim()).toBe(plansPageCopy.title);
    const text = doc.querySelector('main')?.textContent ?? '';
    expect(text).not.toContain(plansPageCopy.band.lede);
    expect(text).not.toContain(plansPageCopy.band.heading);
  });

  it('carries its regions in order: four plan panels, then Questions with the rules, and the coach sample', async () => {
    const doc = await renderPremium();
    const order = regions(doc).filter((r) => !r.startsWith('head-') || r === 'head-answer');
    expect(order).toEqual([
      'head',
      'head-answer',
      'controls',
      'notes',
      'plate',
      'plate',
      'plate',
      'plate',
      'section-questions',
      'plate',
      'section-the-coach-after-a-raid',
      'plate',
    ]);
  });

  it('the interval bar links to ?interval= and carries nothing else', async () => {
    const doc = await renderPremium();
    const radios = [...doc.querySelectorAll('[data-param=interval] [role=radio]')];
    expect(radios.map((r) => r.textContent?.trim())).toEqual(['Yearly', 'Monthly']);
    expect(doc.querySelector('[data-testid=band-guild-link]')).toBeNull();
  });

  it('says founder once: the notice line, no Founder tag on a plan column', async () => {
    const doc = await renderPremium();
    expect(doc.querySelector('[data-testid=founder-row]')?.textContent).toContain(plansPageCopy.founder.open);
    expect(doc.querySelectorAll('[data-testid=plan-premium] .ui-plh-name')[0]?.textContent).not.toContain(
      plansPageCopy.founder.label,
    );
  });

  it('the four plans in order, Premium framed with the teal Subscribe and the guild plans in the secondary tone', async () => {
    const grid = (await renderPremium()).querySelector('[data-testid=plans-plate]');
    const plans = [...(grid?.querySelectorAll('[data-region=plate] > .plan') ?? [])];
    expect(plans.map((p) => p.getAttribute('data-testid'))).toEqual([
      'plan-free',
      'plan-premium',
      'plan-guild',
      'plan-raid_guild',
    ]);
    expect(plans[1]?.parentElement?.className).toContain('framed');
    for (const link of plans[1]?.querySelectorAll('.plan-btn a') ?? [])
      expect(link.className).toContain('bg-teal-edge');
    for (const link of plans[2]?.querySelectorAll('.plan-btn a') ?? [])
      expect(link.className).toContain('bg-blue');
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
      const lazyStorage = {
        getItem: (key: string) => storageOf().getItem(key),
        setItem: (key: string, value: string) => storageOf().setItem?.(key, value),
      };
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

    it('signed in, no hint and no server sims: the tag shows', async () => {
      expect((await runOn(storing(null), '<html data-session="1">')).premiumNav).toBeUndefined();
    });
    const storing = (value: string | null) => (): Storage => ({ getItem: () => value }) as unknown as Storage;

    describe('signed out: the tag waits for a returning visitor (2026-11-14)', () => {
      const VISIT_KEY = 'fs.visitDays';
      const now = new Date();
      const today = `${now.getFullYear()}-${now.getMonth() + 1}-${now.getDate()}`;
      const withVisits = (days: unknown) => {
        const store = new Map<string, string>();
        if (days !== undefined) store.set(VISIT_KEY, typeof days === 'string' ? days : JSON.stringify(days));
        const storage = {
          getItem: (key: string) => store.get(key) ?? null,
          setItem: (key: string, value: string) => void store.set(key, value),
        } as unknown as Storage;
        return { store, storageOf: () => storage };
      };
      const visitsOf = (store: Map<string, string>): string[] => JSON.parse(store.get(VISIT_KEY) ?? '[]');

      it('a first visit hides the tag and records today', async () => {
        const { store, storageOf } = withVisits(undefined);
        expect((await runOn(storageOf)).premiumNav).toBe('hide');
        expect(visitsOf(store)).toEqual([today]);
      });

      it('more page loads on the same day count once', async () => {
        const { store, storageOf } = withVisits(['2026-11-12', today]);
        expect((await runOn(storageOf)).premiumNav).toBe('hide');
        expect(visitsOf(store)).toEqual(['2026-11-12', today]);
      });

      it('the third distinct day shows the tag', async () => {
        const { store, storageOf } = withVisits(['2026-11-10', '2026-11-12']);
        expect((await runOn(storageOf)).premiumNav).toBeUndefined();
        expect(visitsOf(store)).toEqual(['2026-11-10', '2026-11-12', today]);
      });

      it('keeps only the last three days, so the tag keeps showing', async () => {
        const { store, storageOf } = withVisits(['2026-11-01', '2026-11-05', '2026-11-09']);
        expect((await runOn(storageOf)).premiumNav).toBeUndefined();
        expect(visitsOf(store)).toEqual(['2026-11-05', '2026-11-09', today]);
      });

      it('a malformed record counts as a first visit', async () => {
        const { storageOf } = withVisits('{not json');
        expect((await runOn(storageOf)).premiumNav).toBe('hide');
      });

      it('blocked storage hides the tag', async () => {
        const blocked = (): Storage =>
          ({
            getItem: () => {
              throw new Error('blocked');
            },
          }) as unknown as Storage;
        expect((await runOn(blocked)).premiumNav).toBe('hide');
      });

      it('the session cookie alone shows the tag (/b/{id} has no data-session)', async () => {
        const { storageOf } = withVisits(undefined);
        const text = sourceScripts(await renderPremium())[0]?.textContent ?? '';
        const doc = new JSDOM('<!doctype html><html></html>', { url: 'https://foreversixty.gg/' }).window
          .document;
        doc.cookie = 'fs_csrf=abc';
        new Function('localStorage', 'document', text)(storageOf(), doc);
        expect(doc.documentElement.dataset.premiumNav).toBeUndefined();
      });

      it('Premium still hides the tag for a returning visitor', async () => {
        const { storageOf } = withVisits(['2026-11-10', '2026-11-12']);
        expect((await runOn(storageOf, '<html data-server-sims="1">')).premiumNav).toBe('hide');
      });
    });

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
