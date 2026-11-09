// The planner round's talent primitives (planner spec 2026-11-08 section 3A): TalentEditor,
// TalentButton, TalentTip and the pure helpers in lib/ui/talent-editor.
import type { Component } from 'svelte';
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import {
  cellLabel,
  gridKeyAction,
  gridRows,
  moveCell,
  openTreeIndex,
  talentTipSide,
  talentTipSideLeft,
  sortedCells,
  tabStopId,
  type EditorCell,
  type EditorTree,
  type TalentEditorCopy,
} from '../../lib/ui/talent-editor';
import TalentButton from './TalentButton.svelte';
import TalentEditor from './TalentEditor.svelte';
import TalentTip from './TalentTip.svelte';

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

const COPY: TalentEditorCopy = {
  rankOf: (rank, max) => `Rank ${rank} of ${max}`,
  nextRank: 'Next rank',
  resetTree: 'Reset',
  resetTreeLabel: (name) => `Reset ${name}`,
};

function cell(id: number, tier: number, column: number, extra: Partial<EditorCell> = {}): EditorCell {
  return {
    id,
    tier,
    column,
    name: `Talent ${id}`,
    iconSrc: `/icons/${id}.webp`,
    rank: 0,
    max: 5,
    state: 'avail',
    now: '',
    next: 'Next text',
    ...extra,
  };
}

function tree(id: number, points: number, cells: EditorCell[]): EditorTree {
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
      { d: 'M 74 20 V 60', met: true },
      { d: 'M 126 20 V 60', met: false },
    ],
    cells,
  };
}

const CELLS = [cell(1, 0, 1), cell(2, 0, 2), cell(3, 1, 0), cell(4, 2, 3)];
const TREES = [tree(10, 0, CELLS), tree(20, 5, [cell(5, 0, 0)]), tree(30, 0, [cell(6, 0, 0)])];
const noop = () => {};

describe('talent-editor helpers', () => {
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
});

describe('TalentButton', () => {
  const props = (over: Partial<EditorCell> = {}) => ({
    cell: cell(7, 0, 0, over),
    copy: COPY,
    onadd: noop,
    onremove: noop,
  });

  it('is one button with the icon, the hidden rank pill and its state', () => {
    const body = html(TalentButton, props({ rank: 2, state: 'part' }));
    expect(body.match(/<button/g)).toHaveLength(1);
    expect(body).toContain('type="button"');
    expect(body).toMatch(/class="tc2 part[^"]*"/);
    expect(body).toContain('aria-label="Talent 7, rank 2 of 5"');
    expect(body).toContain('data-testid="talent-7"');
    expect(body).toContain('data-rank="2"');
    expect(body).toContain('data-state="part"');
    expect(body).toMatch(/<img[^>]*width="48"[^>]*height="48"[^>]*loading="eager"/);
    expect(body).toMatch(/<span class="rk[^"]*" aria-hidden="true">2\/5<\/span>/);
    expect(body).toContain('tabindex="-1"');
    expect(body).not.toContain('role="tooltip"');
  });

  it.each(['avail', 'part', 'max', 'lock'] as const)('carries the %s state class', (state) => {
    expect(html(TalentButton, props({ state }))).toMatch(new RegExp(`class="tc2 ${state}`));
  });

  it('marks a locked cell aria-disabled, and only a locked one', () => {
    expect(html(TalentButton, props({ state: 'lock' }))).toContain('aria-disabled="true"');
    expect(html(TalentButton, props({ state: 'max' }))).not.toContain('aria-disabled');
  });

  it('draws the differs dot and the planned pill, and names the planned mark', () => {
    const body = html(
      TalentButton,
      props({ differs: true, planned: true, plannedLabel: 'planned past level 34' }),
    );
    expect(body).toMatch(/class="tc2 avail diff plan/);
    expect(body).toContain('aria-label="Talent 7, rank 0 of 5, planned past level 34"');
  });

  it('holds the tab stop when focused', () => {
    expect(html(TalentButton, { ...props(), focused: true })).toContain('tabindex="0"');
  });

  it('falls back to the dark socket without an icon', () => {
    const body = html(TalentButton, props({ iconSrc: '' }));
    expect(body).not.toContain('<img');
    expect(body).toMatch(/<span class="ic[^"]*" aria-hidden="true"><\/span>/);
  });
});

describe('TalentTip', () => {
  it('is a tooltip with the name, rank line, this rank and the next under its label', () => {
    const body = html(TalentTip, {
      id: 'tip-1',
      talentId: 1,
      name: 'Pathfinding',
      rankLine: 'Rank 1 of 2',
      now: 'Now text',
      next: 'Next text',
      nextLabel: 'Next rank',
    });
    expect(body).toContain('role="tooltip"');
    expect(body).toContain('id="tip-1"');
    expect(body).toContain('data-testid="talent-tip-1"');
    expect(body).toMatch(/<b[^>]*>Pathfinding<\/b>/);
    expect(body).toContain('Rank 1 of 2');
    expect(body).toMatch(/<p[^>]*>Now text<\/p>/);
    expect(body).toMatch(/class="nx[^"]*"><span class="t-label[^"]*">Next rank<\/span>Next text/);
    expect(body).not.toContain('class="req');
    expect(body).not.toContain('talent-band-line');
  });

  it('adds the refusal and the band line only when given', () => {
    const body = html(TalentTip, {
      talentId: 4,
      name: 'X',
      rankLine: 'Rank 0 of 1',
      nextLabel: 'Next rank',
      next: 'n',
      req: 'Needs 5 points in Beast Mastery',
      band: 'The level 60 build takes 1/1',
    });
    expect(body).toMatch(/class="req[^"]*">Needs 5 points in Beast Mastery/);
    expect(body).toMatch(/data-testid="talent-band-line-4"[^>]*>The level 60 build takes 1\/1/);
    expect(body).not.toMatch(/<p>/);
  });
});

describe('talent tooltip side', () => {
  it('opens beside the cell, right first, else left, else neither (a phone)', () => {
    expect(talentTipSide(500, 548, 300, 1440)).toBe('right');
    expect(talentTipSide(1100, 1148, 300, 1440)).toBe('left');
    expect(talentTipSide(40, 88, 300, 390)).toBeNull();
    expect(talentTipSideLeft('right', 48, 300)).toBe(56);
    expect(talentTipSideLeft('left', 48, 300)).toBe(-308);
  });
});

describe('TalentEditor', () => {
  const base = { trees: TREES, copy: COPY, onadd: noop, onremove: noop };

  it('sits on the Plate as a region named Talent trees', () => {
    const body = html(TalentEditor, base);
    expect(body).toContain('data-region="plate"');
    expect(body).toContain('role="region" aria-label="Talent trees"');
  });

  it('with a title, sits in a Section panel with its heading inside and its frame', () => {
    const body = html(TalentEditor, { ...base, title: 'Talents', frame: 'framed', testid: 'trees' });
    expect(body).toMatch(/<section class="ui-section[^"]*ui-section-panel[^"]*" data-testid="trees"/);
    expect(body).toMatch(/<h2 class="ui-section-title[^"]*">Talents<\/h2>/);
    expect(body).toMatch(/data-region="plate"[\s\S]*class="tedit"/);
    expect(body).not.toContain('role="region"');
  });

  it('renders a tablist of the trees with the active one selected', () => {
    const body = html(TalentEditor, { ...base, active: 1 });
    expect(body).toContain('role="tablist"');
    expect(body.match(/role="tab"/g)).toHaveLength(3);
    expect(body).toMatch(/aria-selected="true"[^>]*tabindex="0"[^>]*data-testid="tree-tab-1"/);
    expect(body).toMatch(/aria-selected="false"[^>]*tabindex="-1"[^>]*data-testid="tree-tab-0"/);
    expect(body).toMatch(/<b[^>]*>5<\/b><\/span>Tree 20/);
  });

  it('renders a tabpanel per tree, the active one marked on', () => {
    const body = html(TalentEditor, { ...base, active: 2 });
    expect(body.match(/role="tabpanel"/g)).toHaveLength(3);
    expect(body).toMatch(/class="tedit-tree[^"]*on[^"]*" role="tabpanel"[^>]*aria-label="Tree 30"/);
    expect(body).toMatch(/--n: 3/);
  });

  it('heads each tree with its icon, h2 and points, muted at zero', () => {
    const body = html(TalentEditor, base);
    expect(body.match(/<h2[^>]*>Tree \d0<\/h2>/g)).toHaveLength(3);
    expect(body).toContain('<b data-testid="tree-points-10" class="zero">0</b>');
    expect(body).toContain('<b data-testid="tree-points-20">5</b>');
  });

  it('lays each grid out as tiers × columns of gridcells, empty where no talent sits', () => {
    const body = html(TalentEditor, base);
    expect(body).toContain('role="grid"');
    expect(body).toContain('aria-label="Tree 10 talents"');
    expect(body.match(/role="row"/g)).toHaveLength(9);
    expect(body.match(/role="gridcell"/g)).toHaveLength(36);
    expect(body.match(/<button type="button" class="tc2/g)).toHaveLength(6);
    expect(body).toMatch(/<span role="gridcell"[^>]*><\/span>/);
  });

  it('draws the art and each connector as a halo and an on or off path', () => {
    const body = html(TalentEditor, base);
    expect(body).toMatch(/<img class="art[^"]*" src="\/trees\/10.webp" alt="" aria-hidden="true"/);
    expect(body).toContain('viewBox="0 0 200 150"');
    expect(body).toContain('preserveAspectRatio="none"');
    expect(body.match(/class="halo[^"]*"/g)).toHaveLength(6);
    expect(body).toMatch(/class="on[^"]*" d="M 74 20 V 60"/);
    expect(body).toMatch(/class="off[^"]*" d="M 126 20 V 60"/);
  });

  it('gives each tree one tab stop on its first cell', () => {
    const body = html(TalentEditor, base);
    expect(body.match(/class="tc2[^"]*"[^>]*tabindex="0"/g)).toHaveLength(3);
    expect(body).toMatch(/tabindex="0"[^>]*data-testid="talent-1"/);
  });

  it('shows the per-tree Reset only with points, a callback and not read-only', () => {
    const reset = /<button type="button" class="trs[^"]*" aria-label="Reset Tree 20"[^>]*>Reset<\/button>/;
    expect(html(TalentEditor, { ...base, onresettree: noop })).toMatch(reset);
    expect(html(TalentEditor, { ...base, onresettree: noop }).match(/class="trs/g)).toHaveLength(1);
    expect(html(TalentEditor, base)).not.toContain('class="trs');
    expect(html(TalentEditor, { ...base, onresettree: noop, readOnly: true })).not.toContain('class="trs');
  });

  it('renders the foot only when given', () => {
    expect(html(TalentEditor, base)).not.toContain('tedit-f');
    const foot = createRawSnippet(() => ({ render: () => '<p class="st2" role="status"></p>' }));
    expect(html(TalentEditor, { ...base, foot })).toMatch(
      /class="tedit-f[^"]*"><p class="st2" role="status">/,
    );
  });
});
