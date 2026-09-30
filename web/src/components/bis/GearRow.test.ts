// web/src/components/bis/GearRow.test.ts
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import type { RowView } from '../../lib/bis/panel-view';
import GearRow from './GearRow.astro';

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
      expect(html).toContain(`src="/icons/hd/faction/${faction}.png"`);
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
    expect(html).toContain('src="/icons/hd/faction/horde.png"');
  });

  it('never shows a FactionMark for a non-reputation source', async () => {
    const c = await AstroContainer.create();
    const row = repRow({ sourceKind: 'vendor', sourceDetail: 'Vendor: Someone' });
    const html = await c.renderToString(GearRow, {
      props: { row, isNew: false, build: 'test-build', faction: 'alliance', slotTestId: 'test-slot' },
    });
    expect(html).not.toContain('faction-mark-');
  });
});
