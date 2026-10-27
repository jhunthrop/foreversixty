// web/tests/e2e/support/short-controls.ts
// The 44 px touch floor at 360 (guild raids spec §9, attendance spec §9): the visible controls in a
// scope that are shorter than 44 px, counting a `::after` hit area, by test id or text.
import type { Locator } from '@playwright/test';

export const MIN_TARGET_PX = 44;

const CONTROLS =
  'button, select, input:not([type="checkbox"]), textarea, a:not(.prose-forever a), label:has(input[type="checkbox"]), [role="radio"]';

export async function shortControls(scope: Locator): Promise<string[]> {
  return scope.locator(CONTROLS).evaluateAll(
    (nodes, min) =>
      nodes
        .filter((node) => {
          const box = node.getBoundingClientRect();
          return box.width > 0 && box.height > 0 && getComputedStyle(node).visibility !== 'hidden';
        })
        .filter((node) => {
          const own = node.getBoundingClientRect().height;
          const after = parseFloat(getComputedStyle(node, '::after').height) || 0;
          return Math.max(own, after) < min;
        })
        .map(
          (node) =>
            `${node.getAttribute('data-testid') ?? node.textContent?.trim()}: ${node.getBoundingClientRect().height}`,
        ),
    MIN_TARGET_PX,
  );
}
