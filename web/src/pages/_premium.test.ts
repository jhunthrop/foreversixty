// web/src/pages/_premium.test.ts
// Container renders of the plans page (spec 2026-10-27-launch-pricing §B3), through the shared
// `renderPremium` in test-support.
import { readFileSync } from 'node:fs';
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../lib/billing/plans-page-copy';
import { PLAN_CARD_ROWS } from '../components/premium/model';
import { renderPremium } from '../test-support/render-premium';
import { specCell, specString } from '../test-support/spec-table';

describe('the plans page band', () => {
  it('band guild link from 768 only', async () => {
    const doc = await renderPremium();
    const link = doc.querySelector('[data-testid=band-guild-link]');
    expect(link?.getAttribute('href')).toBe('#guild-leaders');
    expect(link?.textContent?.trim()).toBe(plansPageCopy.band.guildLink);
    const row = link?.closest('[data-band-guild-row]');
    const classes = (row?.className ?? '').split(/\s+/);
    expect(classes).toEqual(expect.arrayContaining(['hidden', 'md:flex']));
    expect(row?.textContent).toContain(plansPageCopy.band.guildLead);
    expect(classesOf(link)).toEqual(expect.arrayContaining(['inline-flex', 'min-h-11', 'items-center']));
  });
});

describe('the plans page promises nothing unshipped (§B8.9)', () => {
  it('renders no chip-in and no trial email, and renders the coach line', async () => {
    const text = (await renderPremium()).body.textContent ?? '';
    expect(text).toContain(plansPageCopy.premium.features[0]?.text ?? 'missing');
    expect(text).not.toMatch(/chip in/i);
    expect(text).not.toMatch(/email you/i);
  });
});

const classesOf = (el: Element | null | undefined): string[] => (el?.className ?? '').split(/\s+/);

describe('the plans page region A', () => {
  it('region A holds Free, Premium, the sample coach review in that order', async () => {
    const region = (await renderPremium()).querySelector('[data-region=a]');
    const order = [...(region?.children ?? [])].map((n) => n.getAttribute('data-plan'));
    expect(order).toEqual(['free', 'premium', 'coach-sample']);
    expect(region?.querySelector('[data-plan=coach-sample] [data-testid=coach-sample]')).not.toBeNull();
  });

  it('Free renders once', async () => {
    expect((await renderPremium()).querySelectorAll('[data-plan=free]')).toHaveLength(1);
  });

  it('region A grid classes', async () => {
    const doc = await renderPremium();
    expect(classesOf(doc.querySelector('[data-region=a]'))).toEqual(
      expect.arrayContaining(['grid', 'md:grid-cols-2', 'lg:grid-cols-3', 'md:gap-4']),
    );
    expect(classesOf(doc.querySelector('[data-region=a] > [data-plan=coach-sample]'))).toEqual(
      expect.arrayContaining(['md:col-span-2', 'md:max-w-[480px]', 'lg:col-span-1', 'lg:max-w-none']),
    );
  });

  it('no four-column grid', async () => {
    // The page's own main: the footer's link columns are not this page's plan grid.
    expect((await renderPremium()).querySelector('main')?.innerHTML).not.toMatch(/grid-cols-4|repeat\(4/);
  });

  it('keeps one island, the layout character selector', async () => {
    expect((await renderPremium()).querySelectorAll('astro-island')).toHaveLength(1);
  });
});

function guildSection(doc: Document): HTMLElement {
  const section = doc.getElementById('guild-leaders');
  if (!section) throw new Error('no #guild-leaders');
  return section;
}

describe('the plans page region B', () => {
  it('#guild-leaders holds heading, lede, members, Guild, Raid guild, history promise', async () => {
    const s = guildSection(await renderPremium());
    expect(s.tagName).toBe('SECTION');
    expect(s.querySelector('h2')?.textContent?.trim()).toBe(plansPageCopy.guildRegion.heading);
    expect(classesOf(s.querySelector('h2'))).toContain('heading-section');
    expect(s.textContent).toContain(plansPageCopy.guildRegion.lede);
    expect(s.textContent).toContain(plansPageCopy.guildRegion.members);
    const row = [...s.querySelectorAll('[data-guild-row] > *')].map((n) => n.getAttribute('data-plan'));
    expect(row).toEqual(['guild', 'raid_guild', 'history']);
    expect(s.querySelector('[data-plan=history]')?.textContent?.trim()).toBe(plansPageCopy.historyPromise);
  });

  it('scroll margin and gaps', async () => {
    const s = guildSection(await renderPremium());
    expect(classesOf(s)).toEqual(expect.arrayContaining(['scroll-mt-4', 'mt-7', 'md:mt-10']));
    expect(classesOf(s.querySelector('[data-guild-row]'))).toEqual(
      expect.arrayContaining(['mt-4', 'md:grid-cols-2', 'lg:grid-cols-3', 'md:gap-4']),
    );
    expect(classesOf(s.querySelector('[data-plan=history]'))).toEqual(
      expect.arrayContaining(['md:col-span-2', 'lg:col-span-1']),
    );
  });

  // Each plan card is a subgrid of PLAN_CARD_ROWS rows, so whatever shares its row set from 1024
  // (region A's sample, region B's history promise) must span the same rows, or auto-placement
  // drops it into the card's first row and stretches the name row.
  it("row-set neighbours span the plan cards' rows", async () => {
    const doc = await renderPremium();
    for (const plan of ['free', 'premium', 'guild', 'raid_guild']) {
      expect(doc.querySelector(`article[data-plan=${plan}]`)?.children, plan).toHaveLength(PLAN_CARD_ROWS);
    }
    for (const cell of ['coach-sample', 'history']) {
      expect(classesOf(doc.querySelector(`[data-plan=${cell}]`)), cell).toContain(
        `lg:row-span-${PLAN_CARD_ROWS}`,
      );
    }
  });

  it('holds How the guild plans work', async () => {
    const s = guildSection(await renderPremium());
    expect(s.querySelector('[data-testid=guild-rules] h3')?.textContent?.trim()).toBe(
      plansPageCopy.howItWorks,
    );
    expect(s.querySelectorAll('[data-rule-term]')).toHaveLength(plansPageCopy.rules.length);
  });

  it('names the guild cards as h3 under the section heading', async () => {
    const s = guildSection(await renderPremium());
    const names = [...s.querySelectorAll('[data-guild-row] h3')].map((n) => n.textContent?.trim());
    expect(names).toEqual([plansPageCopy.guild.name, plansPageCopy.raidGuild.name]);
    expect(s.querySelectorAll('h2')).toHaveLength(1);
  });

  it('no loot-council example', async () => {
    expect((await renderPremium()).body.textContent).not.toContain('EXAMPLE · A REAL LOOT COUNCIL');
  });
});

describe('the Premium card when the account has Premium', () => {
  const SPEC = new URL('../../../design/specs/2026-10-27-subscribe-and-join.md', import.meta.url);
  const HEADING = '### 9.2';
  const SOURCES = ['own', 'seat', 'grant'] as const;
  const premiumCard = async (): Promise<Element> => {
    const card = (await renderPremium()).querySelector('[data-region=a] [data-plan=premium]');
    if (!card) throw new Error('no Premium card');
    return card;
  };
  const sourceScripts = (doc: Document): Element[] =>
    [...doc.querySelectorAll('script:not([src])')].filter((s) => s.textContent?.includes('fs.premiumSource'));

  it('one inline script (in the header) copies the three stored words onto data-premium', async () => {
    const doc = await renderPremium();
    const scripts = sourceScripts(doc);
    expect(scripts).toHaveLength(1);
    const text = scripts[0]?.textContent ?? '';
    expect(text).toContain('dataset.premium');
    for (const word of SOURCES) expect(text).toContain(word);
    expect(doc.querySelectorAll('script[src*="_astro"]')).toHaveLength(0);
  });

  it('holds all three source lines beside the default audience, each in the reserved cell', async () => {
    const card = await premiumCard();
    const cell = card.querySelector('[data-part=audience]');
    expect(cell?.querySelector('[data-default]')?.textContent?.trim()).toBe(plansPageCopy.premium.audience);
    const lines = specCell(SPEC, HEADING, '`/premium` card line');
    SOURCES.forEach((source, index) => {
      expect(cell?.querySelector(`[data-source=${source}]`)?.textContent?.trim()).toBe(lines[index]);
    });
  });

  it('holds See your plan beside the default buttons, in the fixed action slot', async () => {
    const card = await premiumCard();
    const slot = card.querySelector('.action-slot');
    expect(classesOf(slot)).toContain('min-h-11');
    const link = slot?.querySelector('a[data-owned]');
    expect(link?.textContent?.trim()).toBe(specString(SPEC, HEADING, 'Buttons', 2));
    expect(link?.getAttribute('href')).toBe('/account#plan');
    expect(classesOf(link)).toEqual(expect.arrayContaining(['border-line-warm-strong', 'text-strong']));
    expect(classesOf(link)).not.toContain('text-gold');
    expect(slot?.querySelectorAll('a[data-iv]').length).toBeGreaterThan(0);
  });

  it('the other cards carry no owned versions', async () => {
    const doc = await renderPremium();
    expect(
      doc.querySelectorAll('[data-plan=guild] [data-owned], [data-plan=free] [data-source]'),
    ).toHaveLength(0);
  });

  it('the audience cell takes the same responsive gap as the plain audience line', () => {
    const source = readFileSync(new URL('../components/premium/PlanCard.astro', import.meta.url), 'utf8');
    expect(source).toMatch(/\.audience,\s*\.audience-cell\s*\{\s*margin-bottom: 8px;/);
  });

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
    expect(doc.getElementById('premium')?.getAttribute('data-plan')).toBe('premium');
  });

  it('each anchor is a Premium feature line, and server-sims marks lines 2 to 4', async () => {
    const card = (await renderPremium()).querySelector('[data-region=a] [data-plan=premium]')!;
    const lines = [...card.querySelectorAll('li.feature')];
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

  it('the landing highlight is CSS on :target, a shadow rule, and the page ships no script file', async () => {
    const source = readFileSync(new URL('../components/premium/PlanCard.astro', import.meta.url), 'utf8');
    for (const id of ['coach', 'server-sims', 'history', 'notifications']) {
      expect(source).toContain(`.features:has(#${id}:target) [data-anchor='${id}']`);
    }
    expect(source).toContain('box-shadow: inset 2px 0 0 var(--color-gold)');
    expect(source).toMatch(/\.feature \{[^}]*padding: 6px 8px;[^}]*border-radius: var\(--radius-control\);/);
    expect((await renderPremium()).querySelectorAll('script[src*="_astro"]')).toHaveLength(0);
  });
});
