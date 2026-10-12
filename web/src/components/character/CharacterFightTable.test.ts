// web/src/components/character/CharacterFightTable.test.ts
// L8.4 fix round 3 items 4, 6 and 7, and fix round 4 items 3, 6, 7 and 8: the character page's
// fight tables.
//   Round 3, 4 and 6. A scored Execution cell (a link) used to leak onto a grid row of its own on
//      a phone and to sit ~12px above its row on a desktop. Both states take one class set: a
//      44px, vertically centred flex box.
//   Round 3, 7. The column is explained by a muted caption under the table's title.
//   Round 4, 3 and 7. Both tables keep the Execution column on a phone (Spec and Date still go),
//      the cell at least 44px wide and 16px after the amount; the caption shows on a phone too
//      (item 6), and the header's title and sr-only duplicate are gone.
//   Round 4, 8. 8px more between the amount and Date from md.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { characterTableCopy, executionCellClass, fightRowId } from '../../lib/character-profile';
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

function table(title = 'Every ranked fight'): string {
  return render(CharacterFightTable, {
    props: {
      title,
      rows: [row(1, 0.94), row(2, null)],
      testid: 'character-history',
      bossHref: () => '/reports/r1',
      rowKey: (fight: CharacterFight) => String(fight.fight_index),
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
/** The header row's cells, in column order. */
const headerCells = (html: string): { classes: string[]; attrs: string; text: string }[] => {
  const header = /<li[^>]*data-testid="character-table-header"[^>]*>([\s\S]*?)<\/li>/.exec(html)?.[1] ?? '';
  return [...header.matchAll(/<span( class="([^"]*)")?([^>]*)>([^<]*)<\/span>/g)].map((m) => ({
    classes: (m[2] ?? '').split(/\s+/).filter((token) => token !== ''),
    attrs: m[3],
    text: m[4],
  }));
};

describe('CharacterFightTable', () => {
  it('gives a scored and an unscored Execution cell the one class set, shown on a phone too (fix round 4 item 3)', () => {
    const [scored, unscored] = cells(table(), 'character-execution');
    expect(sorted(scored)).toEqual(sorted(executionCellClass(true).split(' ')));
    expect(sorted(unscored)).toEqual(sorted(executionCellClass(false).split(' ')));
    for (const cell of [scored, unscored]) {
      expect(cell).toEqual(expect.arrayContaining(['flex', 'min-h-11', 'min-w-11', 'items-center']));
      expect(cell.filter((token) => /(^|:)hidden$/.test(token))).toEqual([]);
      expect(cell).not.toContain('inline-flex');
    }
  });

  it('keeps the Execution header on a phone, 16px after the amount, and drops Spec and Date there (fix round 4 items 3 and 7)', () => {
    const [boss, spec, parse, amount, date, execution] = headerCells(table());
    expect(boss.text).toBe(characterTableCopy.boss);
    expect(spec.classes).toEqual(expect.arrayContaining(['hidden', 'md:inline']));
    expect(parse.text).toBe(characterTableCopy.parse);
    expect(amount.text).toBe('HPS');
    expect(date.classes).toEqual(expect.arrayContaining(['hidden', 'md:inline']));
    expect(execution.text).toBe(characterTableCopy.execution);
    expect(execution.classes.filter((token) => /(^|:)hidden$/.test(token))).toEqual([]);
    expect(execution.classes).toContain('max-md:ml-1');
    // Boss, Parse %, the amount and Execution on a phone.
    const list = /<ul[^>]*data-testid="character-history"/.exec(table())?.[0] ?? '';
    expect(list).toBeTruthy();
    const grid = /<li class="([^"]*)"[^>]*data-testid="character-table-header"/.exec(table())?.[1] ?? '';
    expect(grid.split(' ')).toContain('grid-cols-[minmax(0,1fr)_auto_auto_auto]');
  });

  it('puts 8px more between the amount and Date from md (fix round 4 item 8)', () => {
    const html = table();
    const date = headerCells(html)[4];
    expect(date.classes).toContain('md:ml-2');
    for (const cell of cells(html, 'character-fight-date')) expect(cell).toContain('md:ml-2');
  });

  it('explains Execution under the title in a muted caption at every width, with no title or sr-only copy (fix round 4 items 6 and 7)', () => {
    for (const title of ['Every ranked fight', 'Best per encounter']) {
      const html = table(title);
      const definition = escaped(characterTableCopy.executionDefinition);
      const caption = new RegExp(
        `<p class="([^"]*)" data-testid="character-execution-definition">${definition}</p>`,
      ).exec(html);
      expect(caption, title).not.toBeNull();
      const tokens = caption?.[1].split(' ') ?? [];
      expect(tokens).toEqual(expect.arrayContaining(['text-muted', 'text-[12px]']));
      expect(tokens.filter((token) => token.includes('sr-only') || token.includes('hidden'))).toEqual([]);
      // Under the title, before the table.
      expect(html.indexOf(title)).toBeLessThan(html.indexOf('character-execution-definition'));
      expect(html.indexOf('character-execution-definition')).toBeLessThan(
        html.indexOf('data-testid="character-history"'),
      );
      // Said once: no title attribute repeating it on the header.
      expect(html).not.toContain(`title="${characterTableCopy.executionDefinition}"`);
    }
  });

  // Fix round 5 item 7: the dash is named "Not scored yet", on hover and read aloud; a scored
  // cell keeps its own title.
  it('names the "—" cell "Not scored yet" in its title and aria-label (fix round 5 item 7)', () => {
    const html = table();
    const unscored =
      /<span class="[^"]*" title="([^"]*)" aria-label="([^"]*)" data-testid="character-execution">—<\/span>/.exec(
        html,
      );
    expect(unscored?.slice(1)).toEqual([
      characterTableCopy.executionUnscored,
      characterTableCopy.executionUnscored,
    ]);
    expect(characterTableCopy.executionUnscored).toBe('Not scored yet');
    expect(html).toMatch(/<a [^>]*title="94% of what this gear can do, simulated"[^>]*>94%<\/a>/);
  });

  // L8.5 ruling 1: the rating card's "Weakest part of a fight" anchors to a row of "Every ranked
  // fight", so that table's rows carry an id; "Best per encounter" repeats fights and gets none.
  it('gives each row the id the rating card links to, when asked (L8.5 ruling 1)', () => {
    const withIds = render(CharacterFightTable, {
      props: {
        title: 'Every ranked fight',
        rows: [row(1, 0.94), row(2, null)],
        testid: 'character-history',
        bossHref: () => '/reports/r1',
        rowKey: (fight: CharacterFight) => String(fight.fight_index),
        executionTestid: 'character-execution',
        rowId: fightRowId,
      },
    }).body;
    const ids = [...withIds.matchAll(/<li class="([^"]*)" id="([^"]*)"/g)];
    expect(ids.map((match) => match[2])).toEqual(['fight-r1-1', 'fight-r1-2']);
    // Scrolled to with room above it, and marked while it is the address's target.
    for (const match of ids) {
      expect(match[1].split(' ')).toEqual(expect.arrayContaining(['scroll-mt-4', 'target:bg-raised']));
    }
    expect(table()).not.toMatch(/<li[^>]* id="/);
  });
});
