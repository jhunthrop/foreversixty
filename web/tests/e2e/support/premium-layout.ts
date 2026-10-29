// web/tests/e2e/support/premium-layout.ts
// Measurements of the /premium plan cards that the e2e spec asserts on and the launch-pricing
// capture prints (design/specs/2026-10-27-launch-pricing.md B8 check 4, B9 rows 16 and 18). Each
// reads the box a visitor sees: price states hidden by the interval or founder CSS have no
// client rects and are skipped.
import type { Page } from '@playwright/test';

/** The guild price blocks are one height at every width: price, founder line, per-month line, 54 px cost box. */
const GUILD_BLOCK_PX = 128;
/** The cost line's box: three 18 px lines (PriceBlock.astro). */
export const COST_BOX_PX = 54;
export const GUILD_PLANS = ['guild', 'raid_guild'] as const;

export const guildBlockPx = (): number => GUILD_BLOCK_PX;

export interface GuildPriceBoxes {
  plan: string;
  block: number;
  costBox: number;
  costOverflows: boolean;
}

/** Every `[data-plan]` box's height by plan: free, premium, example, guild, raid_guild, history. */
export function cardHeights(page: Page): Promise<Record<string, number>> {
  return page
    .locator('[data-plan]')
    .evaluateAll((nodes) =>
      Object.fromEntries(
        nodes.map((node) => [node.getAttribute('data-plan') ?? '', node.getBoundingClientRect().height]),
      ),
    );
}

/** Each guild plan's price block, its visible cost line's box, and whether the text overflows it. */
export function guildPriceBoxes(page: Page): Promise<GuildPriceBoxes[]> {
  return Promise.all(
    GUILD_PLANS.map((plan) =>
      page.getByTestId(`price-${plan}`).evaluate((block, name): GuildPriceBoxes => {
        const cost = [...block.querySelectorAll<HTMLElement>('[data-testid="cost-line"]')].find(
          (node) => node.getClientRects().length > 0,
        );
        if (cost === undefined) throw new Error(`${name}: no visible cost line`);
        return {
          plan: name,
          block: block.getBoundingClientRect().height,
          costBox: cost.getBoundingClientRect().height,
          costOverflows: cost.scrollHeight > cost.clientHeight,
        };
      }, plan),
    ),
  );
}

export interface FounderLineBox {
  text: string;
  lines: number;
  textWidth: number;
  boxWidth: number;
}

/** Each visible founder line: how many lines its text sets in, and its text and box widths. */
export function founderLineBoxes(page: Page): Promise<FounderLineBox[]> {
  return page.locator('.founder-line').evaluateAll((nodes) =>
    nodes
      .filter((node) => node.getClientRects().length > 0)
      .map((node): FounderLineBox => {
        const range = document.createRange();
        range.selectNodeContents(node);
        const tops = new Set([...range.getClientRects()].map((rect) => Math.round(rect.top)));
        return {
          text: node.textContent ?? '',
          lines: tops.size,
          textWidth: range.getBoundingClientRect().width,
          boxWidth: node.getBoundingClientRect().width,
        };
      }),
  );
}
