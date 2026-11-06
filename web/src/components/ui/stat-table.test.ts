// The sim round's primitives (sim spec 2026-11-06 §4.5, §4.7, ruling S8): StatTable, StatRow,
// Facts, IconList and Notice's lead. What each renders and the rules each holds.
import { readFileSync } from 'node:fs';
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ariaSortOf, sortMark } from '../../lib/ui/stat-table';
import Facts from './Facts.svelte';
import IconList from './IconList.svelte';
import Notice from './Notice.svelte';
import StatRow from './StatRow.svelte';
import StatTable from './StatTable.svelte';

const text = (html: string) => createRawSnippet(() => ({ render: () => html }));

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const COLUMNS = [
  { id: 'dps', label: 'DPS', sortable: true },
  { id: 'share', label: 'Share', sortable: true },
  { id: 'casts', label: 'Casts' },
];

describe('StatTable', () => {
  it('is an ARIA table whose sortable heads are buttons under a columnheader with aria-sort', () => {
    const body = html(StatTable, {
      columns: COLUMNS,
      nameHead: 'Ability',
      barHead: 'Share of damage',
      sort: { column: 'dps', direction: 'desc' },
      onsort: () => {},
      children: text('<div></div>'),
    });
    expect(body).toContain('role="table"');
    expect(body.match(/role="columnheader"/g)).toHaveLength(5);
    expect(body).toMatch(/role="columnheader"[^>]*aria-sort="descending"[^>]*>\s*<button[^>]*>DPS/);
    expect(body).toMatch(/aria-sort="none"[^>]*>\s*<button[^>]*>Share/);
    // A head that does not sort says nothing about sorting and holds no button.
    expect(body).toMatch(/role="columnheader" data-align="end">Casts/);
    expect(body).toContain('▾');
  });

  it('sets the figure count on the grid and drops the bar track when there is no bar', () => {
    const body = html(StatTable, { columns: COLUMNS, bar: false, children: text('<div></div>') });
    expect(body).toContain('--st-n: 3');
    expect(body).toContain('data-bar="false"');
  });

  it('draws no head row without a name head: the open lists', () => {
    const body = html(StatTable, { columns: COLUMNS.slice(0, 1), open: true, children: text('<div></div>') });
    expect(body).not.toContain('role="columnheader"');
    expect(body).toContain('data-open="true"');
  });

  it('holds no track floor above its content (the buffs overlap fix, ruling M1)', () => {
    const css = readFileSync(new URL('./stat-table.css', import.meta.url), 'utf8');
    expect(css).toContain('--st-name: minmax(0, 2fr)');
    expect(css).toContain('--st-bar: minmax(0, 1fr)');
    expect(css).toContain('--st-fig: minmax(max-content, 62px)');
    expect(css).not.toMatch(/--st-(?:name|bar):\s*minmax\([1-9]/);
  });
});

describe('StatRow', () => {
  const base = {
    name: text('<b>Bloodthirst</b>'),
    figures: [{ value: '125.4' }, { value: '–', muted: true }],
  };

  it('keeps the icon and the bar decorative and makes the name and figures cells', () => {
    const body = html(StatRow, { ...base, icon: text('<img alt="">'), meter: text('<i></i>') });
    expect(body).toContain('role="row"');
    expect(body.match(/role="cell"/g)).toHaveLength(3);
    expect(body.match(/aria-hidden="true"/g)).toHaveLength(2);
  });

  it('marks the primary figure, mutes a dash and carries the phone line under the primary', () => {
    const body = html(StatRow, { ...base, phone: '13.9%' });
    expect(body).toMatch(/data-primary="true"[^>]*>125\.4<span class="ui-st-phone">13\.9%<\/span>/);
    expect(body).toMatch(/data-muted="true"[^>]*>–/);
  });

  it('a Total row keeps its empty bar cell so the figures stay under their heads', () => {
    const body = html(StatRow, { ...base, total: true });
    expect(body).toContain('data-total="true"');
    expect(body).toContain('class="ui-st-bar"');
    expect(html(StatRow, { ...base, bar: false })).not.toContain('ui-st-bar');
  });
});

describe('stat-table rules', () => {
  it('aria-sort reads the direction on the sorted head, none on another sortable one, nothing otherwise', () => {
    const sort = { column: 'dps', direction: 'asc' } as const;
    expect(ariaSortOf(COLUMNS[0]!, sort)).toBe('ascending');
    expect(ariaSortOf(COLUMNS[1]!, sort)).toBe('none');
    expect(ariaSortOf(COLUMNS[2]!, sort)).toBeUndefined();
    expect(sortMark('desc')).toBe('▾');
    expect(sortMark('asc')).toBe('▴');
  });
});

describe('Facts', () => {
  it('is a dl of terms and definitions with the sub line under the value', () => {
    const body = html(Facts, {
      items: [
        { id: 'fight', label: 'Fight length', value: '3:00, ± 20%', sub: '2:24 to 3:36' },
        { id: 'engine', label: 'Engine', value: 'bf72769', muted: true },
      ],
    });
    expect(body.match(/<dt/g)).toHaveLength(2);
    expect(body).toMatch(/<dd[^>]*>\s*3:00, ± 20%<small[^>]*>2:24 to 3:36<\/small>/);
    expect(body).toMatch(/<dd class="[^"]*text-muted[^"]*"[^>]*>\s*bf72769/);
  });
});

describe('IconList', () => {
  it('heads the list with its label, dims the no-effect entries and keeps a socket for a missing icon', () => {
    const body = html(IconList, {
      label: 'On you · 2',
      items: [
        { id: 'motw', name: 'Gift of the Wild', icon: '/i/motw.webp' },
        { id: 'ai', name: 'Arcane Brilliance', icon: null, dim: true },
      ],
    });
    expect(body).toMatch(/<h3[^>]*>On you · 2<\/h3>/);
    expect(body).toMatch(/data-dim="true"[^>]*>[\s\S]*?opacity-\[\.62\][\s\S]*?Arcane Brilliance/);
    expect(body.match(/<img/g)).toHaveLength(1);
  });

  it('opens the game tooltip from an entry that names one', () => {
    const body = html(IconList, {
      items: [{ id: 'a', name: 'Battle Shout', icon: null, tip: 'spell:25289' }],
    });
    expect(body).toMatch(/data-game-tip="spell:25289" tabindex="0"/);
  });
});

describe('Notice', () => {
  it('leads with its mark before the sentence', () => {
    const body = html(Notice, {
      text: 'Death Wish has a 3 minute cooldown.',
      tone: 'neutral',
      lead: text('<img alt="">'),
    });
    expect(body).toMatch(/<img alt="">[\s\S]*Death Wish/);
  });
});
