// web/src/components/premium/PriceBlock.test.ts
// Container renders of the price block (spec 2026-10-27-launch-pricing §B3.3): every interval and
// founder state in the DOM, the guild blocks fixed at 128 px with the cost line in a 54 px box.
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import type { PlanKey } from '../../lib/billing/plans';
import { createPageContainer } from '../../test-support/page-container';
import PriceBlock from './PriceBlock.astro';
import { costLine, founderKey } from './model';

const GUILD_PLANS = ['guild', 'raid_guild'] as const;

async function renderBlock(plan: PlanKey): Promise<Element> {
  const container = await createPageContainer();
  const html = await container.renderToString(PriceBlock, { props: { plan } });
  const root = new JSDOM(html).window.document.querySelector(`[data-testid=price-${plan}]`);
  if (!root) throw new Error(`no price block for ${plan}`);
  return root;
}

const classesOf = (el: Element | null | undefined): string[] => (el?.className ?? '').split(/\s+/);

describe('PriceBlock', () => {
  it.each([
    ['yearly', true],
    ['yearly', false],
    ['monthly', true],
    ['monthly', false],
  ] as const)(
    'cost line sits in a 54 px box in all four forms (%s, founder=%s)',
    async (interval, founder) => {
      for (const plan of GUILD_PLANS) {
        const root = await renderBlock(plan);
        const cell = root.querySelector(`[data-iv=${interval}][data-fd=${founderKey(founder)}]`);
        const line = cell?.querySelector('[data-testid=cost-line]');
        expect(classesOf(line)).toContain('h-[54px]');
        expect(line?.textContent?.trim()).toBe(costLine(plan, interval, founder));
      }
    },
  );

  it('guild price block is fixed 128 px at every width', async () => {
    for (const plan of GUILD_PLANS) {
      expect(classesOf(await renderBlock(plan))).toContain('h-[128px]');
    }
  });

  it('premium has no cost line', async () => {
    expect((await renderBlock('premium')).querySelector('[data-testid=cost-line]')).toBeNull();
  });

  it('guild price figure 22/24 px, premium 44/52 (polish spec 3.3)', async () => {
    const guild = classesOf((await renderBlock('raid_guild')).querySelector('.amount'));
    expect(guild).toEqual(expect.arrayContaining(['text-[22px]', 'md:text-[24px]']));
    const premium = classesOf((await renderBlock('premium')).querySelector('.amount'));
    expect(premium).toEqual(expect.arrayContaining(['text-[44px]', 'md:text-[52px]']));
  });
});
