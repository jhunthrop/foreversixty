// web/src/pages/_premium.test.ts
// Container renders of the plans page (spec 2026-10-27-launch-pricing §B3). `renderPremium` is the
// shared entry point: later lanes render the same page through it and add their own cases.
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../lib/billing/plans-page-copy';
import { createPageContainer } from '../test-support/page-container';
import Premium from './premium.astro';

/** Renders /premium through Base.astro and returns its parsed document. */
export async function renderPremium(): Promise<Document> {
  const container = await createPageContainer();
  const html = await container.renderToString(Premium);
  return new JSDOM(html).window.document;
}

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
  });
});

const classesOf = (el: Element | null | undefined): string[] => (el?.className ?? '').split(/\s+/);

describe('the plans page region A', () => {
  it('region A holds Free, Premium, example in that order', async () => {
    const region = (await renderPremium()).querySelector('[data-region=a]');
    const order = [...(region?.children ?? [])].map((n) => n.getAttribute('data-plan'));
    expect(order).toEqual(['free', 'premium', 'example']);
    expect(region?.querySelector('[data-plan=example] [data-testid=premium-example]')).not.toBeNull();
  });

  it('Free renders once', async () => {
    expect((await renderPremium()).querySelectorAll('[data-plan=free]')).toHaveLength(1);
  });

  it('region A grid classes', async () => {
    const doc = await renderPremium();
    expect(classesOf(doc.querySelector('[data-region=a]'))).toEqual(
      expect.arrayContaining(['grid', 'md:grid-cols-2', 'lg:grid-cols-3', 'md:gap-4']),
    );
    expect(classesOf(doc.querySelector('[data-region=a] > [data-plan=example]'))).toEqual(
      expect.arrayContaining(['md:col-span-2', 'md:max-w-[480px]', 'lg:col-span-1', 'lg:max-w-none']),
    );
  });

  it('no four-column grid', async () => {
    expect((await renderPremium()).body.innerHTML).not.toMatch(/grid-cols-4|repeat\(4/);
  });

  it('keeps one island, the layout character selector', async () => {
    expect((await renderPremium()).querySelectorAll('astro-island')).toHaveLength(1);
  });
});
