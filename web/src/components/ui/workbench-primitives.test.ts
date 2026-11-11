// The planner round's workbench primitives (planner spec 2026-11-08 section 3A): PointOrder,
// CodeField, Page layout="workbench" with Ordered, and the ControlBar's compact class picker.
import type { Component } from 'svelte';
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import {
  drawsTalent,
  holdsPoint,
  ORDER_MARKERS,
  orderIndices,
  type PointOrderRow,
} from '../../lib/ui/point-order';
import CodeField from './CodeField.svelte';
import ControlBar from './ControlBar.svelte';
import Ordered from './Ordered.svelte';
import Page from './Page.svelte';
import PointOrder from './PointOrder.svelte';

const text = (markup: string) => createRawSnippet(() => ({ render: () => markup }));

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

const ROWS: PointOrderRow[] = [
  {
    label: 'Levels 10 to 19',
    slots: [
      { kind: 'spent', name: 'Level 10: Endurance Training 1', iconSrc: '/i/et.webp', rank: 1 },
      { kind: 'planned', name: 'Level 11: Endurance Training 2', iconSrc: '/i/et.webp', rank: 2 },
      { kind: 'next', name: 'Next point: level 12' },
      ...Array.from({ length: 7 }, (_, i) => ({ kind: 'empty' as const, name: `Level ${13 + i}` })),
    ],
  },
  {
    label: 'Levels 20 to 29',
    slots: [
      { kind: 'spent', name: 'Level 20: Pathfinding 1', rank: 1 },
      ...Array.from({ length: 9 }, (_, i) => ({ kind: 'empty' as const, name: `Level ${21 + i}` })),
    ],
  },
];

describe('point-order helpers', () => {
  it('marks the last digit of each socket', () => {
    expect(ORDER_MARKERS).toEqual([0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);
  });

  it('counts only the sockets that hold a point, in reading order', () => {
    expect(holdsPoint({ kind: 'planned', name: '' })).toBe(true);
    expect(holdsPoint({ kind: 'next', name: '' })).toBe(false);
    const indices = orderIndices(ROWS);
    expect(indices[0].slice(0, 3)).toEqual([0, 1, null]);
    expect(indices[1][0]).toBe(2);
  });
});

describe('PointOrder', () => {
  it('leads with the hidden 0 to 9 marker row', () => {
    const body = html(PointOrder, { build: '1.60.1.70790', rows: ROWS });
    expect(body).toMatch(/^<div class="porder[^"]*"><div class="porder-m[^"]*" aria-hidden="true">/);
    expect(body).toMatch(/(<span[^>]*>\d<\/span>){10}/);
  });

  it('renders each row as its label and a list of ten listitem sockets', () => {
    const body = html(PointOrder, { build: '1.60.1.70790', rows: ROWS });
    expect(body.match(/role="list"/g)).toHaveLength(2);
    expect(body.match(/role="listitem"/g)).toHaveLength(20);
    expect(body).toContain('Levels 10 to 19');
  });

  it('draws each slot kind: spent and planned icons as tab stops, empty and next sockets', () => {
    const body = html(PointOrder, { build: '1.60.1.70790', rows: ROWS });
    expect(body).toMatch(
      /role="listitem" tabindex="0" aria-label="Level 10: Endurance Training 1" title="Level 10: Endurance Training 1" data-kind="spent" data-tip-beside="" data-testid="order-point-0"><img[^>]*width="36"[^>]*><b[^>]*>1<\/b>/,
    );
    expect(body).toMatch(/role="listitem" class="pl[^"]*" tabindex="0"[^>]*data-testid="order-point-1"/);
    expect(body).toMatch(
      /aria-label="Next point: level 12" title="Next point: level 12" data-kind="next"><span class="em nx/,
    );
    expect(body).toMatch(
      /aria-label="Level 13" title="Level 13" data-kind="empty"><span class="em[^"]*"><\/span>/,
    );
    expect(body.match(/tabindex="0"/g)).toHaveLength(3);
    expect(body).toMatch(/data-testid="order-point-2"><span class="ic[^"]*" aria-hidden="true"><\/span>/);
  });

  it('draws a guide point ahead dimmed on its own test id, and a tipped point without a title', () => {
    const rows = [
      {
        label: 'Levels 10 to 19',
        slots: [
          {
            kind: 'spent' as const,
            name: 'Level 10: Deflection 1',
            iconSrc: '/i/d.webp',
            rank: 1,
            tip: { name: 'Deflection', subtitle: 'Rank 1/5', lines: [], description: 'Parry 1%' },
          },
          { kind: 'ahead' as const, name: 'Level 11: Deflection 2', iconSrc: '/i/d.webp', rank: 2 },
        ],
      },
    ];
    const body = html(PointOrder, { build: '1.60.1.70790', rows });
    expect(body).toMatch(
      /aria-label="Level 10: Deflection 1" data-kind="spent" data-tip-beside="" data-testid="order-point-0"/,
    );
    expect(body).toMatch(
      /class="ah[^"]*" tabindex="0"[^>]*data-kind="ahead" data-tip-beside="" data-testid="order-ahead-1"/,
    );
    expect(drawsTalent({ kind: 'ahead', name: '' })).toBe(true);
    expect(holdsPoint({ kind: 'ahead', name: '' })).toBe(false);
  });

  it('puts the gold cap line after row capAfter only when it has text', () => {
    const body = html(PointOrder, {
      build: '1.60.1.70790',
      rows: ROWS,
      capAfter: 0,
      capText: 'Felsnap is level 19.',
    });
    expect(body).toMatch(
      /Levels 10 to 19<\/span>.*<\/div><\/div><p class="porder-cap[^"]*">Felsnap is level 19.<\/p>/,
    );
    expect(html(PointOrder, { build: '1.60.1.70790', rows: ROWS, capAfter: 0 })).not.toContain('porder-cap');
    expect(html(PointOrder, { build: '1.60.1.70790', rows: ROWS, capText: 'x' })).not.toContain('porder-cap');
  });
});

describe('CodeField', () => {
  it('is a named textarea, 68 px by default', () => {
    const body = html(CodeField, { label: 'Import from addon', placeholder: 'FS1:…', testid: 'code' });
    expect(body).toMatch(/^<textarea class="codef[^"]*"/);
    expect(body).toContain('aria-label="Import from addon"');
    expect(body).toContain('placeholder="FS1:…"');
    expect(body).toContain('rows="2"');
    expect(body).not.toContain('aria-invalid');
    expect(body).not.toContain('readonly');
  });

  it('marks an invalid code and points at its message', () => {
    const body = html(CodeField, { label: 'x', invalid: true, describedby: 'msg' });
    expect(body).toMatch(/class="codef[^"]*bad/);
    expect(body).toContain('aria-invalid="true"');
    expect(body).toContain('aria-describedby="msg"');
  });

  it('is read-only and one line when asked', () => {
    const body = html(CodeField, { label: 'x', readonly: true, short: true, value: 'FS1:abc' });
    expect(body).toContain('readonly');
    expect(body).toMatch(/class="codef[^"]*short/);
    expect(body).toContain('rows="1"');
    expect(body).toContain('FS1:abc');
  });
});

describe('Page workbench', () => {
  it('renders main and the rail, each region in an Ordered slot', () => {
    const body = html(Page, {
      layout: 'workbench',
      rail: text('<section>R</section>'),
      children: text('<section>M</section>'),
    });
    expect(body).toContain('data-layout="workbench"');
    expect(body).toMatch(/<div class="ui-page-main[^"]*"><section>M<\/section><\/div>/);
    expect(body).toMatch(/<aside class="ui-page-rail[^"]*"><section>R<\/section><\/aside>/);
  });

  it('Ordered carries its order and span', () => {
    expect(html(Ordered, { order: 6, children: text('<p>x</p>') })).toBe(
      '<div data-order="" style="--o: 6;"><p>x</p></div>',
    );
    expect(html(Ordered, { order: 1, span: true, children: text('<p>x</p>') })).toMatch(
      /<div data-order="" data-span="" style="--o: 1;">/,
    );
  });
});

describe('ControlBar compact class picker', () => {
  const groups = [
    {
      label: 'Class',
      param: 'class',
      kind: 'segmented' as const,
      compact: true,
      value: 'hunter',
      options: [
        { value: 'warrior', label: 'Warrior', crest: 'warrior' },
        { value: 'hunter', label: 'Hunter', crest: 'hunter' },
      ],
    },
  ];

  it('keeps each radio named by its class while the label hides', () => {
    const body = html(ControlBar, { groups, label: 'Planner choices' });
    expect(body).toMatch(/class="ui-seg ui-seg--compact"/);
    expect(body).toMatch(/role="radio"[^>]*aria-checked="false" aria-label="Warrior"/);
    expect(body).toMatch(/role="radio"[^>]*aria-checked="true" aria-label="Hunter"/);
    expect(body).toMatch(/<span class="ui-seg-lbl">Hunter<\/span>/);
  });

  it('draws the ringed crest at 22 px', () => {
    const body = html(ControlBar, { groups, label: 'x' });
    expect(body).toMatch(
      /<img src="\/icons\/hd\/crests\/hunter.webp" alt="" width="22" height="22"[^>]*class="crest"[^>]*style="--c: var\(--color-class-hunter\); width: 22px; height: 22px;"/,
    );
  });

  it('leaves a plain group as it was', () => {
    const plain = [
      { ...groups[0], compact: undefined, options: [{ value: 'a', label: 'A', icon: '/a.webp' }] },
    ];
    const body = html(ControlBar, { groups: plain, label: 'x' });
    expect(body).not.toContain('ui-seg--compact');
    expect(body).not.toContain('aria-label="A"');
    expect(body).toContain('ui-spell-icon');
    expect(body).toContain('<span>A</span>');
  });
});
