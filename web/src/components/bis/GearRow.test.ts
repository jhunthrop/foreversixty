// web/src/components/bis/GearRow.test.ts
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import type { RowView } from '../../lib/bis/panel-view';
import GearRow from './GearRow.astro';

/**
 * The exact substring of one `<span>...</span>` in a rendered row's HTML, found by its own
 * opening tag containing `startNeedle` and balanced forward against every nested `<span>`/
 * `</span>` pair until its matching close -- a hand-rolled balance walk rather than a real
 * DOM parser, since this file's `AstroContainer.renderToString` calls must stay in the
 * default `node` test environment.
 */
function extractSpan(html: string, startNeedle: string): string {
  const startIdx = html.indexOf(startNeedle);
  expect(startIdx).toBeGreaterThanOrEqual(0);
  const tagOpen = html.lastIndexOf('<span', startIdx);
  expect(tagOpen).toBeGreaterThanOrEqual(0);
  const openTagEnd = html.indexOf('>', tagOpen) + 1;
  const tagPattern = /<span\b|<\/span>/g;
  tagPattern.lastIndex = openTagEnd;
  let depth = 1;
  let match: RegExpExecArray | null;
  while ((match = tagPattern.exec(html)) !== null) {
    depth += match[0] === '</span>' ? -1 : 1;
    if (depth === 0) return html.slice(tagOpen, tagPattern.lastIndex);
  }
  throw new Error(`unbalanced <span> starting at "${startNeedle}"`);
}

function repRow(overrides: Partial<RowView> = {}): RowView {
  return {
    slot: 'neck',
    empty: false,
    itemId: 1,
    itemName: 'Test Amulet',
    sourceKind: 'rep',
    sourceDetail: 'Silverwing Sentinels (honored)',
    alternatives: [],
    ...overrides,
  };
}

describe('GearRow', () => {
  it.each(['alliance', 'horde'] as const)(
    'shows the 16px FactionMark before a reputation source line on the pick, for %s',
    async (faction) => {
      const c = await AstroContainer.create();
      const html = await c.renderToString(GearRow, {
        props: { row: repRow(), isNew: false, build: 'test-build', faction, slotTestId: 'test-slot' },
      });
      expect(html).toContain(`src="/icons/hd/faction/${faction}.webp"`);
      expect(html).toContain('data-testid="faction-mark-' + faction + '"');
    },
  );

  it('shows the FactionMark on a reputation alternative too, not just the main pick', async () => {
    const c = await AstroContainer.create();
    const row = repRow({
      sourceKind: 'vendor',
      sourceDetail: 'Vendor: Someone',
      alternatives: [
        {
          itemId: 2,
          itemName: 'Alt Amulet',
          sourceKind: 'rep',
          sourceDetail: 'The Defilers',
          dpsDelta: -0.4,
        },
      ],
    });
    const html = await c.renderToString(GearRow, {
      props: { row, isNew: false, build: 'test-build', faction: 'horde', slotTestId: 'test-slot' },
    });
    expect(html).toContain('src="/icons/hd/faction/horde.webp"');
  });

  it('never shows a FactionMark for a non-reputation source', async () => {
    const c = await AstroContainer.create();
    const row = repRow({ sourceKind: 'vendor', sourceDetail: 'Vendor: Someone' });
    const html = await c.renderToString(GearRow, {
      props: { row, isNew: false, build: 'test-build', faction: 'alliance', slotTestId: 'test-slot' },
    });
    expect(html).not.toContain('faction-mark-');
  });

  describe('the hover trigger (owner report: the tooltip used to open from anywhere in the cell)', () => {
    it('is the icon and the name only -- the stats and source lines are outside it', async () => {
      const c = await AstroContainer.create();
      const row = repRow({ keyStatsLine: '12 AP · 5% Crit' });
      const html = await c.renderToString(GearRow, {
        props: { row, isNew: false, build: 'test-build', faction: 'alliance', slotTestId: 'test-slot' },
      });
      const trigger = extractSpan(html, 'data-testid="item-hover-1"');
      expect(trigger).toContain('data-bis-item="1"');
      expect(trigger).toContain('role="button"');
      expect(trigger).toContain('tabindex="0"');
      expect(trigger).toContain('gear-row-name');
      // The actual hover/focus trigger -- the name-line -- never contains the stats or
      // source line: those stayed `.gear-row-content`'s own children, outside this
      // element, which is the whole point of the fix. Never a nested `[role="button"]`
      // either (the a11y fix tests/e2e/bis.spec.ts also guards).
      expect(trigger).not.toContain('gear-row-stats');
      expect(trigger).not.toContain('gear-row-source');
      const buttonCount = trigger.split('role="button"').length - 1;
      expect(buttonCount).toBe(1);
      // The full row DOES render the stats and source lines -- just outside the trigger.
      expect(html).toContain('gear-row-stats');
      expect(html).toContain('gear-row-source');
    });

    it("gives the icon its own hover host for the same item, separate from the name's trigger", async () => {
      const c = await AstroContainer.create();
      const html = await c.renderToString(GearRow, {
        props: {
          row: repRow(),
          isNew: false,
          build: 'test-build',
          faction: 'alliance',
          slotTestId: 'test-slot',
        },
      });
      const iconHost = extractSpan(html, 'gear-row-icon-trigger');
      expect(iconHost).toContain('data-bis-item="1"');
      expect(iconHost).not.toContain('data-testid="item-hover-1"');
      expect(iconHost).toMatch(/<img\b|gear-row-icon-placeholder/);
      const nameTrigger = extractSpan(html, 'data-testid="item-hover-1"');
      expect(nameTrigger).toContain('data-bis-item="1"');
      expect(nameTrigger).not.toContain('gear-row-icon-trigger');
      // A single `item-hover-1` testid on the page -- the icon's own host is never a
      // second one, so tests/e2e/bis.spec.ts's "one hover host per pick" count stays true.
      expect(html.split('data-testid="item-hover-1"').length - 1).toBe(1);
    });
  });
});
