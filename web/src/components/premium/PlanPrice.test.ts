// web/src/components/premium/PlanPrice.test.ts
// A plan's price cells (premium spec 2026-11-08 section 4.4): every interval and founder state in
// the DOM, each cell the card model's own figure and lines, the founder line only in a founder cell.
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { INTERVALS, PLAN_ORDER } from '../../lib/billing/plans';
import { createPageContainer } from '../../test-support/page-container';
import PlanPrice from './PlanPrice.astro';
import { FOUNDER_STATES, costLine, founderKey, priceCell } from './model';

async function cellsOf(plan: (typeof PLAN_ORDER)[number]): Promise<Element[]> {
  const container = await createPageContainer();
  const html = await container.renderToString(PlanPrice, { props: { plan } });
  return [...new JSDOM(html).window.document.querySelectorAll('.plan-price-cell')];
}

describe('PlanPrice', () => {
  it.each(PLAN_ORDER)('%s draws one cell per interval and founder state from the model', async (plan) => {
    const cells = await cellsOf(plan);
    expect(cells).toHaveLength(INTERVALS.length * FOUNDER_STATES.length);
    for (const interval of INTERVALS) {
      for (const founder of FOUNDER_STATES) {
        const cell = cells.find(
          (c) => c.getAttribute('data-iv') === interval && c.getAttribute('data-fd') === founderKey(founder),
        );
        const want = priceCell(plan, interval, founder);
        const text = cell?.textContent ?? '';
        expect(text).toContain(want.amount);
        expect(text).toContain(want.secondLine);
        if (want.founderLine)
          expect(cell?.querySelector('[data-testid=founder-line]')?.textContent).toBe(want.founderLine);
        else expect(cell?.querySelector('[data-testid=founder-line]')).toBeNull();
        const cost = costLine(plan, interval, founder);
        if (cost) expect(text).toContain(cost);
      }
    }
  });
});
