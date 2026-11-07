// web/src/components/ui/premium-primitives.test.ts
// The premium and account round's primitives (premium spec 2026-11-08 section 3A): Columns, PlanHead,
// and the Button's danger tone (lib/ui/button.test.ts holds the recipe itself).
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Button from './Button.svelte';
import Columns from './Columns.svelte';
import PlanHead from './PlanHead.svelte';

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const raw = (value: string) => createRawSnippet(() => ({ render: () => value }));

describe('Columns', () => {
  it('lays its children out as columns: open by default, the count on the grid', () => {
    const body = html(Columns, { count: 2, children: raw('<div><p>a</p><p>b</p></div>'), testid: 'c' } as never);
    expect(body).toContain('class="ui-cols');
    expect(body).toContain('data-open="true"');
    expect(body).toContain('--cols: 2');
    expect(body).not.toContain('data-swap');
    expect(body).not.toContain('role="group"');
  });

  it('is padded when closed, swaps the second column first when asked, and names a group', () => {
    const body = html(Columns, {
      count: 2,
      open: false,
      secondFirst: true,
      label: 'Plans for a player',
      children: raw('<div><p>a</p><p>b</p></div>'),
    } as never);
    expect(body).toContain('data-open="false"');
    expect(body).toContain('data-swap="true"');
    expect(body).toMatch(/role="group"[^>]*aria-label="Plans for a player"/);
  });
});

describe('PlanHead', () => {
  it('draws the name, audience, the price Figure and its lines; gold on the plan the page sells', () => {
    const body = html(PlanHead, {
      name: 'Premium',
      audience: 'For one player',
      value: '$40',
      unit: '/year',
      gold: true,
      lines: [{ text: '$3.33 a month, billed yearly' }, { text: 'Founder price', tone: 'gold' }],
    } as never);
    expect(body).toMatch(/<h3 class="ui-plh-name[^"]*">Premium<\/h3>/);
    expect(body).toContain('For one player');
    expect(body).toMatch(/text-gold[^"]*">\$40<span[^>]*>\/year/);
    expect(body).toMatch(/class="ui-plh-line[^"]*">\$3\.33 a month, billed yearly</);
    expect(body).toMatch(/class="ui-plh-line is-gold[^"]*">Founder price</);
  });

  it('is strong on a plan the page does not sell, takes a level, a Tag and a price snippet', () => {
    const body = html(PlanHead, {
      name: 'Free',
      audience: 'For every player',
      value: '$0',
      level: 2,
      tag: raw('<span class="tag">Your plan</span>'),
    } as never);
    expect(body).toMatch(/<h2 class="ui-plh-name[^"]*">Free<span class="tag">Your plan<\/span><\/h2>/);
    expect(body).toMatch(/text-strong[^"]*">\$0</);
    const custom = html(PlanHead, { name: 'Guild', audience: 'a', value: '', price: raw('<b>both</b>') } as never);
    expect(custom).toContain('<b>both</b>');
    expect(custom).not.toContain('ui-figure');
  });
});

describe('Button danger', () => {
  it('draws the destructive confirm as an outline in wipe ink, never filled', () => {
    const body = html(Button, { tone: 'danger', size: 'sm', children: raw('Leave guild') } as never);
    expect(body).toContain('text-wipe');
    expect(body).toContain('border-wipe/40');
    expect(body).not.toMatch(/bg-wipe/);
    expect(body).toContain('>Leave guild</button>');
  });
});
