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
