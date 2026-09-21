// web/src/components/report/ExchangeTable.test.ts
// Anchor id coverage for ExchangeTable's rows: id={exchange-<source_guid>-<spell_id>} is the
// convention src/lib/rating/moments.ts already builds hrefs against (Ruling 4 of
// docs/superpowers/plans/2026-09-21-rating-web.md), so a rating moment's link actually lands
// on the row it names. Static-render check (svelte/server, no jsdom), the pattern
// AuraTable.test.ts already uses for a report component whose whole surface is a few props.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { ExchangeRow } from '../../lib/report/types';
import ExchangeTable from './ExchangeTable.svelte';

function row(overrides: Partial<ExchangeRow> = {}): ExchangeRow {
  return {
    kind: 'interrupt',
    source_guid: 'Player-4184-000000A1',
    source_name: 'Elyra Duskvale',
    target_guid: 'Creature-1',
    target_name: 'Shazzrah',
    spell_id: 20066,
    spell_name: 'Repentance',
    extra_spell_id: 19712,
    extra_spell_name: 'Arcane Explosion',
    count: 1,
    ...overrides,
  };
}

describe('ExchangeTable', () => {
  it('every row carries an exchange-<guid>-<spellId> anchor id for a rating moment to jump to', () => {
    const { body } = render(ExchangeTable, {
      props: { rows: [row()], emptyText: 'Nothing interrupted.' },
    });
    expect(body).toContain('id="exchange-Player-4184-000000A1-20066"');
  });
});
