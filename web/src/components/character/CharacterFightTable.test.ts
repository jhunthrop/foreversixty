// web/src/components/character/CharacterFightTable.test.ts
// L8.4 fix round 3 items 4, 6 and 7: the character page's fight tables.
//   4. On a phone, "Every ranked fight" hides its Execution column; a scored cell (a link) used
//      to leak onto a grid row of its own because rowLink's `inline-flex` beat `hidden`. Both
//      states now take one class set, hidden alike.
//   6. On a desktop the scored cell sat ~12px above its row (`md:inline` blockified a 44px box
//      with its text at the top). Both states are one 44px, vertically centred flex box.
//   7. The column is explained: a muted caption under the table's title from md, the header's
//      title and screen-reader text on a phone.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { characterTableCopy, executionCellClass } from '../../lib/character-profile';
import type { CharacterFight } from '../../lib/rankings/api';
import CharacterFightTable from './CharacterFightTable.svelte';

const row = (fightIndex: number, execution: number | null): CharacterFight => ({
  encounter: 'Warden Kelthas',
  encounter_id: 9001,
  difficulty: 8,
  metric: 'hps',
  value: 1840,
  percentile: 96.2,
  spec: 'Holy',
  fought_at: '2026-10-12T01:04:53Z',
  report_id: 'r1',
  fight_index: fightIndex,
  execution_score: execution,
});

function table(executionOnPhone: boolean): string {
  return render(CharacterFightTable, {
    props: {
      title: 'Every ranked fight',
      rows: [row(1, 0.94), row(2, null)],
      testid: 'character-history',
      bossHref: () => '/reports/r1',
      rowKey: (fight: CharacterFight) => String(fight.fight_index),
      executionOnPhone,
      executionTestid: 'character-execution',
    },
  }).body.replace(/<!--[^>]*-->/g, '');
}

/** Every element with that testid's class list, in document order. */
const cells = (html: string, testid: string): string[][] =>
  [...html.matchAll(new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`, 'g'))].map((m) =>
    m[1].split(/\s+/).filter((token) => token !== ''),
  );
const sorted = (tokens: string[]): string[] => [...tokens].sort();
const escaped = (text: string): string => text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');

describe('CharacterFightTable', () => {
  it('gives a scored and an unscored Execution cell the one class set, hidden on a phone where the column is', () => {
    const [scored, unscored] = cells(table(false), 'character-execution');
    expect(sorted(scored)).toEqual(sorted(executionCellClass(false, true).split(' ')));
    expect(sorted(unscored)).toEqual(sorted(executionCellClass(false, false).split(' ')));
    for (const cell of [scored, unscored]) {
      expect(cell).toEqual(expect.arrayContaining(['hidden', 'md:flex', 'min-h-11', 'items-center']));
      expect(cell).not.toContain('inline-flex');
    }
  });

  it('keeps both states on a phone where the column stays, in the same row-aligned box', () => {
    const [scored, unscored] = cells(table(true), 'character-execution');
    expect(sorted(scored)).toEqual(sorted(executionCellClass(true, true).split(' ')));
    expect(sorted(unscored)).toEqual(sorted(executionCellClass(true, false).split(' ')));
    expect(scored).not.toContain('hidden');
  });

  it('explains Execution under the title from md, and in the header’s title and sr text on a phone', () => {
    const html = table(false);
    const definition = escaped(characterTableCopy.executionDefinition);
    const caption = new RegExp(
      `<p class="([^"]*)" data-testid="character-execution-definition">${definition}</p>`,
    ).exec(html);
    expect(caption).not.toBeNull();
    expect(caption?.[1].split(' ')).toEqual(
      expect.arrayContaining(['text-muted', 'text-[12px]', 'max-md:sr-only']),
    );
    // Under the title, before the table.
    expect(html.indexOf('Every ranked fight')).toBeLessThan(html.indexOf('character-execution-definition'));
    expect(html.indexOf('character-execution-definition')).toBeLessThan(
      html.indexOf('data-testid="character-history"'),
    );
    expect(html).toMatch(
      new RegExp(`<span[^>]*title="${definition}"[^>]*>${characterTableCopy.execution}</span>`),
    );
  });
});
