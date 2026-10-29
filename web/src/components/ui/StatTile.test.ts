import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import StatTile from './StatTile.svelte';

describe('StatTile', () => {
  const body = render(StatTile, { props: { value: '88', label: 'Forever Score', testid: 'tile' } }).body;

  it('sets the figure in the heading face at the stat size over a quiet label (type pass ruling: grey sentence case)', () => {
    expect(body).toMatch(/figure-stat[^>]*>\s*(<!--.*?-->)?\s*<span[^>]*>88</s);
    expect(body.indexOf('88')).toBeLessThan(body.indexOf('Forever Score'));
    expect(body).toContain('<span class="quiet-label">Forever Score</span>');
  });

  it('is a quiet tile: 12 14 padding, control radius, soft line', () => {
    expect(body).toContain('rounded-control');
    expect(body).toContain('px-[14px] py-3');
    expect(body).toContain('border-line-soft');
    expect(body).toContain('bg-white/[.03]');
  });

  it('carries its testid', () => {
    expect(body).toContain('data-testid="tile"');
  });

  it('puts a figure test id on the figure alone and renders a note under the label', () => {
    const withNote = render(StatTile, {
      props: {
        value: '51',
        label: 'Points left',
        figureTestid: 'figure',
        note: createRawSnippet(() => ({ render: () => '<span>a note</span>' })),
      },
    }).body;
    expect(withNote).toMatch(/<span data-testid="figure">51</);
    expect(withNote.indexOf('Points left')).toBeLessThan(withNote.indexOf('a note'));
  });

  it('renders a figure snippet in place of the value', () => {
    const custom = render(StatTile, {
      props: {
        label: 'DPS',
        figure: createRawSnippet(() => ({ render: () => '<button>Show DPS</button>' })),
      },
    }).body;
    expect(custom).toContain('<button>Show DPS</button>');
  });
});
