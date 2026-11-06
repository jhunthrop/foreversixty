// The guild round's primitives (web layout spec 2026-11-05 §7.3): Segmented, Grid, Tally, Checklist.
import { render } from 'svelte/server';
import type { Component } from 'svelte';
import { describe, expect, it } from 'vitest';
import Checklist from './Checklist.svelte';
import Grid from './Grid.svelte';
import Segmented from './Segmented.svelte';
import Tally from './Tally.svelte';

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

describe('Segmented', () => {
  const options = [
    { value: 'in', label: 'In' },
    { value: 'late', label: 'Late' },
    { value: 'bench', label: 'Bench' },
  ] as const;

  it('is a radio group with one tab stop on the checked option', () => {
    const body = html(Segmented<string>, { options, value: 'late', label: 'Sign-up', onChange: () => {} });
    expect(body).toContain('role="radiogroup"');
    expect(body.match(/tabindex="0"/g)).toHaveLength(1);
    expect(body).toMatch(/aria-checked="true"[^>]*tabindex="0"[^>]*>Late/);
  });

  it('marks a pressed-but-unsaved option busy, never checked', () => {
    const body = html(Segmented<string>, {
      options,
      value: 'in',
      pending: 'bench',
      label: 'S',
      onChange: () => {},
    });
    expect(body).toMatch(/aria-checked="false"[^>]*aria-busy="true"[^>]*>Bench/);
  });

  it('takes the in-row size', () => {
    expect(
      html(Segmented<string>, { options, value: null, label: 'S', size: 'sm', onChange: () => {} }),
    ).toContain('ui-seg--sm');
  });
});

describe('Grid', () => {
  it('is a table whose cells carry their tone and spoken words', () => {
    const body = html(Grid, {
      caption: 'Attendance, last 4 raids',
      columns: ['Mon', 'Tue'],
      rows: [
        {
          key: 'a',
          label: 'Pinefeather',
          sub: 'slipping',
          cells: [
            { tone: 'ok', text: '✓', label: 'There' },
            { tone: 'miss', text: '×', label: 'Missed' },
          ],
        },
      ],
    });
    expect(body).toContain('<table');
    expect(body).toContain('scope="col"');
    expect(body).toMatch(/text-kill[^>]*><span aria-hidden="true">✓<\/span><span class="sr-only">There/);
    expect(body).toContain('text-wipe');
    expect(body).toContain('slipping');
  });
});

describe('Tally', () => {
  it('reads each group against its target, toned at, near and short of it', () => {
    const body = html(Tally, {
      items: [
        { key: 'tanks', label: 'Tanks', have: 2, want: 2 },
        { key: 'healers', label: 'Healers', have: 6, want: 8 },
        { key: 'melee', label: 'Melee', have: 5, want: 10 },
      ],
    });
    expect(body).toMatch(/text-kill[^>]*>2\/2/);
    expect(body).toMatch(/text-gold[^>]*>6\/8/);
    expect(body).toMatch(/text-wipe[^>]*>5\/10/);
  });
});

describe('Checklist', () => {
  it('marks each check and links only the fix', () => {
    const body = html(Checklist, {
      items: [
        { key: 'level', tone: 'ok', title: 'Level 60', sub: 'You are level 60.', fix: null },
        {
          key: 'gear',
          tone: 'warn',
          title: '2 slots are not best in slot',
          sub: 'Hands and Finger 1.',
          fix: { label: 'Fix in the BiS list', href: '/bis/priest/holy' },
        },
      ],
    });
    expect(body).toContain('aria-label="Done"');
    expect(body).toContain('aria-label="To do"');
    expect(body.match(/<a /g)).toHaveLength(1);
    expect(body).toContain('href="/bis/priest/holy"');
  });
});
