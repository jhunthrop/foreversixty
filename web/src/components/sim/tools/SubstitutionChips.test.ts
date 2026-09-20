// web/src/components/sim/tools/SubstitutionChips.test.ts
// Fix round 1, engine-lane rule 5: a two-hander replacing a main-plus-off-hand pair emits a
// SECOND substitution, `{kind: "item", slot: "off_hand", item_id: 0, name: "<item removed>"}`.
// This is a static-render check (no jsdom, no new dependency): `svelte/server`'s own `render`
// compiles the component's real markup into a string, which is enough to prove the sentinel
// never reaches the page as a literal item name.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Item } from '../../../lib/planner/types';
import type { Substitution } from '../../../lib/sim/bulk-types';
import { bulkCopy } from '../../../lib/sim/copy';
import SubstitutionChips from './SubstitutionChips.svelte';

function renderChips(substitutions: Substitution[], items: ReadonlyMap<number, Item> = new Map()): string {
  const { body } = render(SubstitutionChips, {
    props: { substitutions, items, treeVersion: 'test' },
  });
  return body;
}

describe('SubstitutionChips', () => {
  it('renders the emptied off-hand from a two-hander swap as an emptied slot, not a literal item name', () => {
    const body = renderChips([
      { kind: 'item', slot: 'main_hand', item_id: 19351, name: 'Sulfuras, Hand of Ragnaros' },
      { kind: 'item', slot: 'off_hand', item_id: 0, name: '<item removed>' },
    ]);
    expect(body).not.toContain('<item removed>');
    expect(body).toContain(bulkCopy.offHandEmptied);
    // The main-hand chip still renders normally beside it.
    expect(body).toContain('Sulfuras, Hand of Ragnaros');
  });

  it('does not treat an ordinary off-hand substitution as the emptied sentinel', () => {
    const body = renderChips([{ kind: 'item', slot: 'off_hand', item_id: 5, name: 'Real Off-hand' }]);
    expect(body).toContain('Real Off-hand');
    expect(body).not.toContain(bulkCopy.offHandEmptied);
  });

  it('renders a normal item substitution by its own name', () => {
    const body = renderChips([{ kind: 'item', slot: 'head', item_id: 16963, name: 'Helm of Wrath' }]);
    expect(body).toContain('Helm of Wrath');
  });
});
