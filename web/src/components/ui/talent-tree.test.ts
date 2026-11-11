// The site's one talent tree (talent-tree lane 2026-11-10; planner spec 2026-11-08 section 3A):
// TalentTrees in its three modes, TalentTree, TalentButton, and the pure helpers in
// lib/ui/talent-tree.
import type { Component } from 'svelte';
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { talentTipSide, talentTipSideLeft } from '../../lib/items/tooltip-position';
import {
  cellLabel,
  cellState,
  cellTipKey,
  cellTipModel,
  gridKeyAction,
  gridRows,
  moveCell,
  openTreeIndex,
  sortedCells,
  tabStopId,
  treesAtTalents,
  viewState,
  withRanks,
  type TalentCellView,
  type TalentTreesCopy,
  type TalentTreeView,
} from '../../lib/ui/talent-tree';
import TalentButton from './TalentButton.svelte';
import TalentTrees from './TalentTrees.svelte';
import TalentTreesWithTips from './TalentTreesWithTips.svelte';

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  // Without comments and Svelte's scoping hashes, so the assertions read the authored markup.
  return render(component, { props })
    .body.replace(/<!--[^>]*-->/g, '')
    .replace(/ class="([^"]*)"/g, (_, names: string) => {
      const kept = names.split(/\s+/).filter((name) => name !== '' && !name.startsWith('svelte-'));
      return kept.length === 0 ? '' : ` class="${kept.join(' ')}"`;
    })
    .replace(/>\s+</g, '><');
}

const BUILD = '1.60.1.70790';
const COPY: TalentTreesCopy = {
  resetTree: 'Reset',
  resetTreeLabel: (name: string) => `Reset ${name}`,
};

function cell(id: number, tier: number, column: number, extra: Partial<TalentCellView> = {}): TalentCellView {
  return {
    id,
    tier,
    column,
    name: `Talent ${id}`,
    icon: `icon_${id}`,
    iconSrc: `/icons/${id}.webp`,
    rank: 0,
    max: 5,
    state: 'avail',
    ranks: ['One', 'Two', 'Three', 'Four', 'Five'],
    ...extra,
  };
}

function tree(id: number, points: number, cells: TalentCellView[]): TalentTreeView {
  return {
    id,
    name: `Tree ${id}`,
    iconSrc: `/spec/${id}.webp`,
    points,
    artSrc: `/trees/${id}.webp`,
    columns: 4,
    tiers: 3,
    viewBox: '0 0 200 150',
    connectors: [
      { d: 'M 74 20 V 60', met: true, from: 1, needs: 5 },
      { d: 'M 126 20 V 60', met: false, from: 2, needs: 5 },
    ],
    cells,
  };
}

const CELLS = [cell(1, 0, 1), cell(2, 0, 2), cell(3, 1, 0), cell(4, 2, 3)];
const TREES = [tree(10, 0, CELLS), tree(20, 5, [cell(5, 0, 0)]), tree(30, 0, [cell(6, 0, 0)])];
const noop = () => {};

describe('talent tree helpers', () => {
  it('lays a sparse tree out as tiers × columns with holes', () => {
    const rows = gridRows(TREES[0]);
    expect(rows).toHaveLength(3);
    expect(rows.every((row) => row.length === 4)).toBe(true);
    expect(rows[0].map((c) => c?.id)).toEqual([undefined, 1, 2, undefined]);
    expect(rows[2][3]?.id).toBe(4);
  });

  it('sorts cells in reading order', () => {
    expect(sortedCells([CELLS[3], CELLS[0], CELLS[2], CELLS[1]]).map((c) => c.id)).toEqual([1, 2, 3, 4]);
  });

  it('moves focus along the reading order and stops at the edges', () => {
    expect(moveCell(CELLS, 2, 1, 0)).toBe(3);
    expect(moveCell(CELLS, 1, -1, 0)).toBe(1);
    expect(moveCell(CELLS, 4, 1, 0)).toBe(4);
  });

  it('moves vertically to the nearest column in the next occupied tier', () => {
    expect(moveCell(CELLS, 1, 0, 1)).toBe(3);
    expect(moveCell(CELLS, 3, 0, 1)).toBe(4);
    expect(moveCell(CELLS, 4, 0, -1)).toBe(3);
    expect(moveCell(CELLS, 1, 0, -1)).toBe(1);
    expect(moveCell(CELLS, 99, 0, 1)).toBe(1);
  });

  it('maps keys to moves and removals; Enter stays the button click', () => {
    expect(gridKeyAction('ArrowDown')).toEqual({ kind: 'move', dx: 0, dy: 1 });
    expect(gridKeyAction('ArrowLeft')).toEqual({ kind: 'move', dx: -1, dy: 0 });
    expect(gridKeyAction('Backspace')).toEqual({ kind: 'remove' });
    expect(gridKeyAction('Delete')).toEqual({ kind: 'remove' });
    expect(gridKeyAction('Enter')).toBeNull();
  });

  it('opens the phone on the tree with the most points, the first when none', () => {
    expect(openTreeIndex(TREES)).toBe(1);
    expect(openTreeIndex([tree(1, 0, []), tree(2, 0, [])])).toBe(0);
    expect(openTreeIndex([tree(1, 3, []), tree(2, 3, [])])).toBe(0);
  });

  it('keeps the tab stop on the last-focused cell, else the first', () => {
    expect(tabStopId(TREES[0], undefined)).toBe(1);
    expect(tabStopId(TREES[0], 3)).toBe(3);
    expect(tabStopId(TREES[0], 99)).toBe(1);
    expect(tabStopId(tree(1, 0, []), undefined)).toBeUndefined();
  });

  it('names a cell with its rank, and its planned mark when planned', () => {
    expect(cellLabel(cell(1, 0, 0, { rank: 2 }))).toBe('Talent 1, rank 2 of 5');
    expect(cellLabel(cell(1, 0, 0, { planned: true, plannedLabel: 'planned past level 34' }))).toBe(
      'Talent 1, rank 0 of 5, planned past level 34',
    );
    expect(cellLabel(cell(1, 0, 0, { plannedLabel: 'planned past level 34' }))).toBe('Talent 1, rank 0 of 5');
  });

  it('colours a cell as the client: maxed, partial, can spend, locked; a read-only build never spends', () => {
    expect(cellState(5, 5, true)).toBe('max');
    expect(cellState(2, 5, false)).toBe('part');
    expect(cellState(0, 5, true)).toBe('avail');
    expect(cellState(0, 5, false)).toBe('lock');
    expect(viewState(0, 3)).toBe('lock');
    expect(viewState(1, 3)).toBe('part');
    expect(viewState(3, 3)).toBe('max');
  });
});

describe('the talent tooltip', () => {
  it('is the game tooltip: name, Rank n/m, this rank, the next, the refusal and the band line', () => {
    const model = cellTipModel(cell(7, 0, 0, { rank: 2, req: 'Needs 5 points', band: 'The 60 build: 5/5' }));
    expect(model).toMatchObject({
      name: 'Talent 7',
      icon: 'icon_7',
      subtitle: 'Rank 2/5',
      description: 'Two',
      nextRank: 'Three',
      requirement: 'Needs 5 points',
      note: 'The 60 build: 5/5',
    });
  });

  it('reads an untaken talent’s first rank, as the client does, with no next rank line', () => {
    const model = cellTipModel(cell(7, 0, 0));
    expect(model?.subtitle).toBe('Rank 0/5');
    expect(model?.description).toBe('One');
    expect(model?.nextRank).toBeUndefined();
    expect(model?.requirement).toBeUndefined();
  });

  it('has no model without text: the page host serves it by key', () => {
    expect(cellTipModel(cell(7, 0, 0, { ranks: undefined }))).toBeNull();
    expect(cellTipKey(cell(7, 0, 0, { rank: 3 }))).toBe('talent:7:3');
  });

  it('opens beside the cell, right first, else left, else neither (a phone)', () => {
    expect(talentTipSide(500, 548, 300, 1440)).toBe('right');
    expect(talentTipSide(1100, 1148, 300, 1440)).toBe('left');
    expect(talentTipSide(40, 88, 300, 390)).toBeNull();
    expect(talentTipSideLeft('right', 48, 300)).toBe(56);
    expect(talentTipSideLeft('left', 48, 300)).toBe(-308);
  });
});

describe('withRanks and treesAtTalents', () => {
  it('redraws the same layout at other ranks: states, points and arrows', () => {
    const [first] = withRanks(TREES, (c) => (c.id === 1 ? 5 : c.id === 2 ? 9 : 0));
    expect(first.cells.map((c) => [c.rank, c.state])).toEqual([
      [5, 'max'],
      [5, 'max'],
      [0, 'lock'],
      [0, 'lock'],
    ]);
    expect(first.points).toBe(10);
    expect(first.connectors.map((c) => c.met)).toEqual([true, true]);
  });

  it('reads a ladder string, one digit per talent in reading order, a missing tree as nothing', () => {
    const [first, second, third] = treesAtTalents(TREES, '1203-4');
    expect(first.cells.map((c) => c.rank)).toEqual([1, 2, 0, 3]);
    expect(first.connectors.map((c) => c.met)).toEqual([false, false]);
    expect(second.points).toBe(4);
    expect(third.points).toBe(0);
  });

  it('never changes the trees it was handed', () => {
    treesAtTalents(TREES, '5555');
    expect(TREES[0].cells[0].rank).toBe(0);
  });
});

describe('TalentButton', () => {
  const props = (over: Partial<TalentCellView> = {}, editable = true) => ({
    cell: cell(7, 0, 0, over),
    build: BUILD,
    editable,
    focusable: false,
    onadd: noop,
    onremove: noop,
  });

  it('is one button with the icon, the hidden rank pill, its state and its label', () => {
    const body = html(TalentButton, props({ rank: 2, state: 'part' }));
    expect(body.match(/<button/g)).toHaveLength(1);
    expect(body).toContain('type="button"');
    expect(body).toMatch(/class="tt-cell part[^"]*"/);
    expect(body).toContain('aria-label="Talent 7, rank 2 of 5"');
    expect(body).toContain('data-testid="talent-7"');
    expect(body).toContain('data-rank="2"');
    expect(body).toContain('data-state="part"');
    expect(body).toMatch(/<img[^>]*width="48"[^>]*height="48"[^>]*loading="eager"/);
    expect(body).toMatch(/<span class="tt-rk" aria-hidden="true">2\/5<\/span>/);
    expect(body).toContain('tabindex="-1"');
    expect(body).not.toContain('role="tooltip"');
  });

  it.each(['avail', 'part', 'max', 'lock'] as const)('carries the %s state class', (state) => {
    expect(html(TalentButton, props({ state }))).toMatch(new RegExp(`class="tt-cell ${state}`));
  });

  it('editable: wraps the button in its tooltip host, beside the cell, with no page key', () => {
    const body = html(TalentButton, props());
    expect(body).toMatch(/<span class="tt-wrap" data-tip-beside="">/);
    expect(body).not.toContain('data-game-tip');
  });

  it('read-only: the same button names its tooltip for the page host, beside the cell', () => {
    const body = html(TalentButton, props({ rank: 3, state: 'part' }, false));
    expect(body).not.toContain('tt-wrap');
    expect(body).toContain('data-game-tip="talent:7:3"');
    expect(body).toContain('data-tip-beside=""');
    expect(body).not.toContain('aria-disabled');
  });

  it('marks a locked editable cell aria-disabled, and only a locked one that cannot plan ahead', () => {
    expect(html(TalentButton, props({ state: 'lock' }))).toContain('aria-disabled="true"');
    expect(html(TalentButton, props({ state: 'lock', plannable: true }))).not.toContain('aria-disabled');
    expect(html(TalentButton, props({ state: 'max' }))).not.toContain('aria-disabled');
  });

  it('draws the differs dot and the planned pill, and names the planned mark', () => {
    const body = html(
      TalentButton,
      props({ differs: true, planned: true, plannedLabel: 'planned past level 34' }),
    );
    expect(body).toMatch(/class="tt-cell avail diff plan/);
    expect(body).toContain('aria-label="Talent 7, rank 0 of 5, planned past level 34"');
  });

  it('holds the tab stop when focusable', () => {
    expect(html(TalentButton, { ...props(), focusable: true })).toContain('tabindex="0"');
  });

  it('falls back to the dark socket without an icon', () => {
    const body = html(TalentButton, props({ iconSrc: '' }));
    expect(body).not.toContain('<img');
    expect(body).toMatch(/<span class="tt-ic" aria-hidden="true"><\/span>/);
  });
});

describe('TalentTrees', () => {
  const edit = { trees: TREES, build: BUILD, mode: 'edit' as const, copy: COPY, onadd: noop, onremove: noop };

  it('edit: sits on the Plate as a region named Talent trees, at the lg cell', () => {
    const body = html(TalentTrees, edit);
    expect(body).toContain('data-region="plate"');
    expect(body).toContain('role="region" aria-label="Talent trees"');
    expect(body).toMatch(/class="tt-trees" data-size="lg" data-mode="edit"/);
  });

  it('edit with a title: a Section panel with its heading inside and its frame', () => {
    const body = html(TalentTrees, { ...edit, title: 'Talents', frame: 'framed', testid: 'trees' });
    expect(body).toMatch(/<section class="ui-section[^"]*ui-section-panel[^"]*" data-testid="trees"/);
    expect(body).toMatch(/<h2 class="ui-section-title[^"]*">Talents<\/h2>/);
    expect(body).not.toContain('role="region"');
  });

  it('view and compact: bare, at md and sm, with no Plate and no Reset', () => {
    const view = html(TalentTrees, {
      trees: TREES,
      build: BUILD,
      mode: 'view',
      onresettree: noop,
      testid: 't',
    });
    expect(view).not.toContain('data-region="plate"');
    expect(view).toMatch(
      /data-testid="t"><div class="tt-trees" data-size="md" data-mode="view" data-bare=""/,
    );
    expect(view).not.toContain('tt-reset');
    const compact = html(TalentTrees, { trees: TREES, build: BUILD, mode: 'compact' });
    expect(compact).toMatch(/data-size="sm" data-mode="compact"/);
    expect(compact).not.toContain('role="tablist"');
  });

  it('every mode draws the same tree: head, art, arrows, grid and cells', () => {
    for (const mode of ['edit', 'view', 'compact'] as const) {
      const body = html(TalentTrees, { trees: TREES, build: BUILD, mode });
      expect(body.match(/<h3[^>]*>Tree \d0<\/h3>/g)).toHaveLength(3);
      expect(body).toContain('<b data-testid="tree-points-10" class="zero">0</b>');
      expect(body).toMatch(/<img class="tt-art" src="\/trees\/10.webp" alt="" aria-hidden="true"/);
      expect(body).toContain('viewBox="0 0 200 150"');
      expect(body.match(/class="halo"/g)).toHaveLength(6);
      expect(body).toMatch(/class="on" d="M 74 20 V 60"/);
      expect(body).toMatch(/class="off" d="M 126 20 V 60"/);
      expect(body.match(/role="row"/g)).toHaveLength(9);
      expect(body.match(/role="gridcell"/g)).toHaveLength(36);
      expect(body.match(/<button type="button" class="tt-cell/g)).toHaveLength(6);
    }
  });

  it('the server render makes every talent a tab stop (a static page has no script to rove)', () => {
    const body = html(TalentTrees, { trees: TREES, build: BUILD, mode: 'compact' });
    expect(body.match(/class="tt-cell[^"]*"[^>]*tabindex="0"/g)).toHaveLength(6);
  });

  it('lg and md: a tablist of the trees with the active one selected, a tabpanel per tree', () => {
    const body = html(TalentTrees, { ...edit, active: 1 });
    expect(body).toContain('role="tablist"');
    expect(body.match(/role="tab"/g)).toHaveLength(3);
    expect(body).toMatch(/aria-selected="true"[^>]*tabindex="0"[^>]*data-testid="tree-tab-1"/);
    expect(body.match(/role="tabpanel"/g)).toHaveLength(3);
    expect(body).toMatch(/class="tt-col on" role="tabpanel"[^>]*aria-label="Tree 20"/);
    expect(body).toMatch(/--n: 3/);
  });

  it('opens a phone on the tree with the most points until one is picked', () => {
    const body = html(TalentTrees, { trees: TREES, build: BUILD, mode: 'view' });
    expect(body).toMatch(/aria-selected="true"[^>]*data-testid="tree-tab-1"/);
    expect(body).toMatch(/class="tt-col on" role="tabpanel"[^>]*aria-label="Tree 20"/);
  });

  it('edit: the per-tree Reset only with points and a callback', () => {
    const reset = /<button type="button" class="tt-reset" aria-label="Reset Tree 20"[^>]*>Reset<\/button>/;
    expect(html(TalentTrees, { ...edit, onresettree: noop })).toMatch(reset);
    expect(html(TalentTrees, { ...edit, onresettree: noop }).match(/class="tt-reset/g)).toHaveLength(1);
    expect(html(TalentTrees, edit)).not.toContain('tt-reset');
  });

  it('edit: the foot only when given', () => {
    expect(html(TalentTrees, edit)).not.toContain('tt-f');
    const foot = createRawSnippet(() => ({ render: () => '<p class="st2" role="status"></p>' }));
    expect(html(TalentTrees, { ...edit, foot })).toMatch(/class="tt-f"><p class="st2" role="status">/);
  });

  it('scopes its tooltips to a host of its own only when asked (TalentTreesWithTips)', () => {
    expect(html(TalentTrees, { trees: TREES, build: BUILD, mode: 'view', tipScope: true })).toContain(
      'data-tip-scope=""',
    );
    expect(html(TalentTrees, { trees: TREES, build: BUILD, mode: 'compact' })).not.toContain(
      'data-tip-scope',
    );
  });

  it('with tips: the view trees, scoped, and no host in the server render', () => {
    const body = html(TalentTreesWithTips, { trees: TREES, build: BUILD, testid: 'w' });
    expect(body).toMatch(
      /data-testid="w"><div class="tt-trees" data-size="md" data-mode="view"[^>]*data-tip-scope=""/,
    );
  });
});
