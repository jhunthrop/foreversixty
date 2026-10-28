// web/src/components/premium/PlanCard.test.ts
// Container renders of the plan card (spec 2026-10-27-launch-pricing §B3.2, §B3.3).
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../../lib/billing/plans-page-copy';
import { createPageContainer } from '../../test-support/page-container';
import PlanCard from './PlanCard.astro';

type CardPlan = 'free' | 'premium' | 'guild' | 'raid_guild';

async function renderCard(plan: CardPlan): Promise<Document> {
  const container = await createPageContainer();
  const html = await container.renderToString(PlanCard, { props: { plan } });
  return new JSDOM(html).window.document;
}

const slotOf = (doc: Document): Element => {
  const slot = doc.querySelector('[data-part=slot]');
  if (!slot) throw new Error('no slot');
  return slot;
};

describe('PlanCard, Premium', () => {
  it('premium slot holds the support line, not gold', async () => {
    const doc = await renderCard('premium');
    const support = slotOf(doc).querySelector('[data-testid=premium-support]');
    expect(support?.textContent?.trim()).toBe(plansPageCopy.premium.support);
    expect(support?.className).not.toMatch(/gold/);
    expect(slotOf(doc).className).not.toMatch(/gold/);
  });

  it('six benefits then the scope note, no tick on the note', async () => {
    const doc = await renderCard('premium');
    const items = [...doc.querySelectorAll('ul li')];
    expect(items).toHaveLength(6);
    expect(items.map((li) => li.textContent?.trim())).toEqual([...plansPageCopy.premium.features]);
    const scope = doc.querySelector('[data-testid=premium-scope]');
    expect(scope?.textContent?.trim()).toBe(plansPageCopy.premium.scope);
    expect(scope?.querySelector('svg')).toBeNull();
    const list = doc.querySelector('ul');
    expect(list && scope && list.compareDocumentPosition(scope) & 4).toBeTruthy();
  });

  it('keeps the gold top highlight', async () => {
    const card = (await renderCard('premium')).querySelector('[data-plan=premium]');
    expect(card?.hasAttribute('data-highlight')).toBe(true);
  });
});

describe('PlanCard, Free', () => {
  it('free slot stays blank', async () => {
    const doc = await renderCard('free');
    expect(slotOf(doc).textContent?.trim()).toBe('');
    expect(slotOf(doc).children).toHaveLength(0);
    expect(doc.querySelector('[data-testid=premium-scope]')).toBeNull();
  });
});
