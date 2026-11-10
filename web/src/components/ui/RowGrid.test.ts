// RowGrid (home spec 2026-11-06 §2): open Rows in columns, no surface of its own.
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import RowGrid from './RowGrid.svelte';

const rows = createRawSnippet(() => ({ render: () => '<div><a class="ui-row" href="/a">A</a></div>' }));

function html(props: Parameters<typeof render<typeof RowGrid>>[1]['props']): string {
  return render(RowGrid, { props }).body.replace(/<!--[^>]*-->/g, '');
}

describe('RowGrid', () => {
  it('lays its rows out in the given columns', () => {
    const body = html({ cols: 3, children: rows, testid: 'grid' });
    expect(body).toContain('--row-grid-cols: 3');
    expect(body).toContain('data-testid="grid"');
    expect(body).toContain('href="/a"');
  });

  it('defaults to three columns and draws no surface', () => {
    const body = html({ children: rows });
    expect(body).toContain('--row-grid-cols: 3');
    expect(body).not.toMatch(/bg-raised|border-line|rounded-plate/);
  });
});
