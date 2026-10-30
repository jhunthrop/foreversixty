// web/src/components/premium/PlanCard.test.ts
// Container renders of the plan card (spec 2026-10-27-launch-pricing §B3.2, §B3.3).
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../../lib/billing/plans-page-copy';
import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
import { CARD, CARD_QUIET } from '../../lib/ui/card';
import { createPageContainer } from '../../test-support/page-container';
import PlanCard from './PlanCard.astro';
import { trialParts } from './model';

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

  it('seven benefits then the scope note, no tick on the note', async () => {
    const doc = await renderCard('premium');
    const items = [...doc.querySelectorAll('ul li')];
    expect(items).toHaveLength(7);
    expect(items.map((li) => li.textContent?.trim())).toEqual(
      plansPageCopy.premium.features.map((line) => line.text),
    );
    const scope = doc.querySelector('[data-testid=premium-scope]');
    expect(scope?.textContent?.trim()).toBe(plansPageCopy.premium.scope);
    expect(scope?.querySelector('svg')).toBeNull();
    const list = doc.querySelector('ul');
    expect(list && scope && list.compareDocumentPosition(scope) & 4).toBeTruthy();
  });

  it.each(['guild', 'raid_guild'] as const)('%s card wears the quiet card surface', async (plan) => {
    const card = (await renderCard(plan)).querySelector(`[data-plan=${plan}]`);
    expect(card?.classList.contains(CARD_QUIET)).toBe(true);
    expect(card?.hasAttribute('data-highlight')).toBe(false);
  });

  it.each(['free', 'premium'] as const)('%s card keeps the gold top highlight', async (plan) => {
    const card = (await renderCard(plan)).querySelector(`[data-plan=${plan}]`);
    expect(card?.hasAttribute('data-highlight')).toBe(true);
    for (const cls of CARD.split(' ')) expect(card?.classList.contains(cls)).toBe(true);
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

describe.each(['guild', 'raid_guild'] as const)('PlanCard, %s', (plan) => {
  it('guild button is neutral', async () => {
    const buttons = [...(await renderCard(plan)).querySelectorAll('a')];
    expect(buttons.length).toBeGreaterThan(0);
    for (const button of buttons) {
      const classes = button.className.split(/\s+/);
      expect(classes).toEqual(expect.arrayContaining(SECONDARY_BUTTON_FIXED.split(/\s+/)));
      expect(classes).toEqual(
        expect.arrayContaining(['border-line-warm-strong', 'text-text', 'w-full', 'justify-center']),
      );
      expect(classes).not.toContain('border-gold');
      expect(classes).not.toContain('text-gold');
    }
  });

  it('guild cards have no gold top highlight', async () => {
    const card = (await renderCard(plan)).querySelector(`[data-plan=${plan}]`);
    expect(card?.hasAttribute('data-highlight')).toBe(false);
    expect(card?.className).not.toContain('shadow-[inset');
  });

  it('trial slot min 40 px with the strong lead', async () => {
    const slot = slotOf(await renderCard(plan));
    expect(slot.className.split(/\s+/)).toContain('min-h-10');
    expect(slot.querySelector('strong')?.textContent?.trim()).toBe(trialParts().lead);
    expect(slot.textContent?.replace(/\s+/g, ' ').trim()).toBe(plansPageCopy.trial);
  });

  it('five feature lines with no strong lead row', async () => {
    const doc = await renderCard(plan);
    const copy = plan === 'guild' ? plansPageCopy.guild : plansPageCopy.raidGuild;
    const items = [...doc.querySelectorAll('ul li')].map((li) => li.textContent?.trim());
    expect(items).toEqual([...copy.features]);
    expect(doc.querySelector('li.lead')).toBeNull();
  });
});
