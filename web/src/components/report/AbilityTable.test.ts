// web/src/components/report/AbilityTable.test.ts
// Spec 2026-11-04-sim-results 3.6: the drawn table, on a static render.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import { abilityTable } from '../../lib/report/ability-table';
import type { SimResult } from '../../lib/sim/types';
import AbilityTable, { type AbilityDecorate } from './AbilityTable.svelte';

const { summary } = full as unknown as SimResult;
const table = abilityTable({
  actors: summary.damage_done,
  casts: summary.casts,
  auras: summary.auras,
  durationMs: summary.duration_ms,
  names: { spell: { '20662': 'Execute' }, item: {} },
});
const decorate: AbilityDecorate = (row) => ({
  iconUrl: row.spellId === 20662 ? '/icons/execute.webp' : null,
  rank: row.spellId === 20662 ? 5 : 0,
  weapon: row.hand === 'main' ? { id: 1, name: 'Ironwood Blade', quality: 4, iconUrl: '/icons/blade.webp' } : null,
  sources: row.hand === 'extra' ? [{ name: 'Windfury Totem', count: null }] : [],
  swingsBecame: row.hand === 'main' ? 32 : 0,
  tip: row.spellId === 20662,
});
const { body } = render(AbilityTable, { props: { table, decorate } });
const rowHtml = (key: string): string => {
  const match = new RegExp(`data-testid="ability-row-${key}"[\\s\\S]*?</tr>`).exec(body);
  if (match === null) throw new Error(`no row ${key}`);
  return match[0];
};
const text = (html: string): string => html.replace(/<!--[\s\S]*?-->/g, '').replace(/<[^>]+>/g, ' ').replace(/\s+/g, ' ');

describe('AbilityTable', () => {
  it('draws a spell row with its icon, rank, figures and tooltip key', () => {
    const row = rowHtml('spell:20662');
    expect(row).toContain('/icons/execute.webp');
    expect(row).toContain('data-game-tip="spell:20662"');
    const words = text(row);
    for (const figure of ['Execute', 'Rank 5', 'Dodge 8.0%', '190.8', '21.1%', '24', '1,493', '52.2%']) {
      expect(words).toContain(figure);
    }
  });

  it('draws the main-hand row with its weapon and the swings Heroic Strike replaced', () => {
    const row = rowHtml('other:attack/1');
    expect(row).toContain('/icons/blade.webp');
    expect(row).toContain('data-bis-item="1"');
    const words = text(row);
    expect(words).toContain('Main hand');
    expect(words).toContain('Ironwood Blade');
    expect(words).toContain('Miss 19.0%');
    expect(words).toContain('Glance 39.7%');
    expect(words).toContain('32 swings became Heroic Strike');
    expect(words).toContain('63 swings');
  });

  it('names the Extra attacks sources and draws a periodic row in ticks', () => {
    expect(text(rowHtml('other:attack/3'))).toContain('From Windfury Totem');
    const wounds = text(rowHtml('spell:12867'));
    expect(wounds).toContain('59 ticks');
    expect(wounds).toContain('tick crit 37.3%');
    expect(wounds).toContain('100.0%');
  });

  it('never draws a broken image: a row with no icon keeps the empty box', () => {
    expect(rowHtml('spell:23894')).not.toContain('<img');
  });

  it('marks the sorted column and totals the table', () => {
    expect(body).toMatch(/aria-sort="descending"[\s\S]*?ability-sort-dps/);
    const total = /data-testid="ability-total"[\s\S]*?<\/tr>/.exec(body)?.[0] ?? '';
    expect(text(total)).toContain('Total, 9 sources');
    expect(text(total)).toContain('904.4');
    expect(text(total)).toContain('162,715 damage a fight');
  });

  it('carries the Crit header tooltip key and draws no Biggest hit column at max 0', () => {
    expect(body).toContain('data-game-tip="note:crit"');
    expect(body).not.toContain('Biggest hit');
  });
});
