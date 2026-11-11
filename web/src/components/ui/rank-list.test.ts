// The ranked-list primitives (tiers spec 2026-11-06 §4.4, owner decision T6): RankList, ColumnHeads,
// RankGroup, RankRow and Meter's class tone. What each renders and the rules each holds.
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { rankListStyle } from '../../lib/ui/rank-list';
import ColumnHeads from './ColumnHeads.svelte';
import Meter from './Meter.svelte';
import RankGroup from './RankGroup.svelte';
import RankList from './RankList.svelte';
import RankRow from './RankRow.svelte';

const text = (html: string) => createRawSnippet(() => ({ render: () => html }));

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const ROW = {
  rank: '4',
  name: 'Arms',
  href: '/bis/warrior/arms',
  classSlug: 'warrior',
  sub: 'Warrior · Night Elf',
  fraction: 0.886,
  figures: [{ value: '942.2', unit: 'DPS', label: 'Damage per second', testid: 'fig' }],
  gap: { text: '−11.4%', label: 'Gap to the top' },
  links: [
    { label: 'BiS', href: '/bis/warrior/arms', primary: true, ariaLabel: 'BiS: Arms Warrior' },
    { label: 'Guide', href: '/guides/warrior/arms' },
  ],
};

describe('rankListStyle', () => {
  it('sets the figure count, the figure width and the link column from the shape', () => {
    expect(rankListStyle({ figures: 1, links: 4 })).toBe('--rl-n: 1; --rl-fig: 112px; --rl-links: 206px');
    expect(rankListStyle({ figures: 3, links: 3 })).toBe('--rl-n: 3; --rl-fig: 112px; --rl-links: 156px');
    expect(rankListStyle({ figures: 1, links: 0 })).toContain('--rl-links: 0px');
    expect(rankListStyle({ figures: 4, links: 2, figureWidth: 96 })).toBe(
      '--rl-n: 4; --rl-fig: 96px; --rl-links: 106px',
    );
  });

  it('fails fast on a shape the grid cannot hold', () => {
    expect(() => rankListStyle({ figures: 0, links: 4 })).toThrow(/figures must be 1 to 4/);
    expect(() => rankListStyle({ figures: 5, links: 4 })).toThrow(/figures/);
    expect(() => rankListStyle({ figures: 1, links: 5 })).toThrow(/links must be 0 to 4/);
    expect(() => rankListStyle({ figures: 1.5, links: 2 })).toThrow(/figures/);
  });
});

describe('RankList', () => {
  it('puts the heads over an ordered list on one grid, and marks a grouped list', () => {
    const body = html(RankList, {
      heads: [{ label: 'Spec' }, { label: 'DPS', align: 'end' }],
      figures: 1,
      links: 4,
      grouped: true,
      label: 'DPS tier list',
      children: text('<li>row</li>'),
    });
    expect(body).toContain('data-grouped="true"');
    expect(body).toContain('--rl-links: 206px');
    expect(body.indexOf('ui-rank-heads')).toBeLessThan(body.indexOf('<ol'));
    expect(body).toMatch(/<ol class="ui-rank-body" aria-label="DPS tier list"><li>row<\/li><\/ol>/);
  });

  it('draws no heads when it has none', () => {
    expect(html(RankList, { children: text('<li>r</li>') })).not.toContain('ui-rank-heads');
  });
});

describe('ColumnHeads', () => {
  it('is one decorative row of sentence-case labels, numeric heads at the end, each with its full text as a title', () => {
    const body = html(ColumnHeads, { cells: [{ label: 'Spec' }, { label: 'Taken / s', align: 'end' }] });
    expect(body).toMatch(/^<div class="ui-rank-heads ui-rank-cols" aria-hidden="true"/);
    expect(body).not.toContain('uppercase');
    expect(body).toMatch(/data-align="end" title="Taken \/ s">Taken \/ s</);
    expect(body).not.toMatch(/data-align="end"[^>]*>Spec</);
  });
});

describe('RankGroup', () => {
  it('names the group for assistive tech, shows the letter in its tone, and lists its rows', () => {
    const body = html(RankGroup, {
      letter: 'B',
      tone: 'var(--color-rarity-rare-text)',
      band: '10 to 20% behind',
      count: '8 specs',
      label: 'Tier B',
      testid: 'tier-group',
      children: text('<li>row</li>'),
    });
    expect(body).toContain('--rg-tone: var(--color-rarity-rare-text)');
    expect(body).toMatch(/<span class="sr-only">Tier B: <\/span><span aria-hidden="true">B<\/span>/);
    expect(body).toContain('10 to 20% behind');
    expect(body).toContain('8 specs');
    expect(body).toMatch(/<ol class="ui-rank-group-rows" aria-label="Tier B"><li>row<\/li><\/ol>/);
  });
});

describe('RankRow', () => {
  it('links the name in the class colour, fills the bar in it, and is no link itself', () => {
    const body = html(RankRow, ROW);
    expect(body).toMatch(/^<li class="ui-rank-row ui-rank-cols"/);
    expect(body).toContain('--rr-tone: var(--color-class-warrior)');
    expect(body).toMatch(
      /<a class="ui-rr-link" href="\/bis\/warrior\/arms" data-testid="rank-name">Arms<\/a>/,
    );
    expect(body).toContain('width: 88.6%');
    expect(body).toContain('background: var(--color-class-warrior)');
    expect(body).not.toContain('aria-current');
  });

  it('labels each figure and the gap for assistive tech, and marks the best of a column with a title', () => {
    const body = html(RankRow, {
      ...ROW,
      figures: [
        { value: '245', unit: 'taken / s', label: 'Damage taken per second' },
        {
          value: '49,794',
          unit: 'HP',
          label: 'Effective health',
          best: true,
          bestTitle: 'Best of the 3 tanks',
        },
      ],
    });
    expect(body).toContain('<span class="sr-only">Damage taken per second: </span>');
    expect(body).toMatch(/data-best="true" title="Best of the 3 tanks"/);
    expect(body).toMatch(/<span class="ui-rr-extra">.*49,794/);
    expect(body).toContain(
      '<span class="sr-only">Gap to the top: </span><span data-testid="rank-gap">−11.4%</span>',
    );
  });

  it('shows the tie under the bar and the note under the name, each with its reason', () => {
    const body = html(RankRow, {
      ...ROW,
      tie: { text: '≈ tie with Survival', title: 'Within 1% of Survival Hunter' },
      note: { text: 'Stat weights less certain', title: 'Why' },
    });
    expect(body).toMatch(
      /class="ui-rr-bar">.*title="Within 1% of Survival Hunter" data-testid="rank-tie">≈ tie with Survival</,
    );
    expect(body).toMatch(/title="Why" data-testid="rank-note">Stat weights less certain/);
  });

  it('makes the first link primary, carries each aria label, and marks a selected row current', () => {
    const body = html(RankRow, {
      ...ROW,
      selected: true,
      key: 'warrior/arms',
      gap: { ...ROW.gap, leader: true },
    });
    expect(body).toMatch(/data-primary="true" aria-label="BiS: Arms Warrior"/);
    expect(body).toContain('aria-current="true"');
    expect(body).toContain('data-rank-key="warrior/arms"');
    expect(body).toContain('data-leader="true"');
  });

  it('draws the ringed crest from the class', () => {
    const body = html(RankRow, { ...ROW });
    expect(body).toMatch(/<span class="ui-rr-lead"><img src="\/icons\/hd\/crests\/warrior.webp"[^>]*class="crest ui-rr-crest/);
  });

  it('falls back to gold and no crest without a class, and renders a tag', () => {
    const body = html(RankRow, {
      ...ROW,
      classSlug: undefined,
      links: [],
      tag: text('<span>Your spec</span>'),
    });
    expect(body).toContain('--rr-tone: var(--color-gold)');
    expect(body).toContain('<span class="ui-rr-lead"></span>');
    expect(body).toContain('Arms</a><span>Your spec</span>');
    expect(body).not.toContain('ui-rr-links');
  });
});

describe('phoneFigures (the log rankings, round 5)', () => {
  it('marks a compact list that keeps its first extra figures beside the sub line on a phone', () => {
    const kept = html(RankList, { compact: true, phoneFigures: 1, children: text('<li>r</li>') });
    expect(kept).toContain('data-phone-figures="1"');
  });

  it('leaves every list without it, and every list that is not compact, as it was', () => {
    expect(html(RankList, { compact: true, children: text('<li>r</li>') })).not.toContain(
      'data-phone-figures',
    );
    expect(html(RankList, { phoneFigures: 1, children: text('<li>r</li>') })).not.toContain(
      'data-phone-figures',
    );
  });
});

describe('the compact ranked list (round 5)', () => {
  it('marks a compact list, and leaves every other list as it was', () => {
    const compact = html(RankList, { compact: true, grouped: true, children: text('<li>r</li>') });
    expect(compact).toContain('data-compact="true"');
    expect(html(RankList, { children: text('<li>r</li>') })).not.toContain('data-compact');
  });

  it('draws a letter column with the letter alone when the group has no band or count', () => {
    const body = html(RankGroup, { letter: 'B', label: 'Tier B', children: text('<li>r</li>') });
    expect(body).toContain('<span aria-hidden="true">B</span>');
    expect(body).not.toContain('ui-rg-band');
    expect(body).not.toContain('ui-rg-count');
    expect(body).not.toContain('--rg-tone');
  });

  it("fills the bar in the accent on request, and drops a figure's unit when the head names it", () => {
    const body = html(RankRow, {
      ...ROW,
      sub: undefined,
      links: [],
      barTone: 'accent',
      figures: [{ value: '942.2', label: 'Damage per second' }],
    });
    expect(body).toContain('bg-accent');
    expect(body).not.toContain('background: var(--color-class-warrior)');
    expect(body).not.toContain('<small>');
    expect(body).not.toContain('ui-rr-sub');
    // The class tone stays the default.
    expect(html(RankRow, ROW)).toContain('background: var(--color-class-warrior)');
  });
});

describe('Meter class tone', () => {
  it('fills in the colour it is given, and refuses the class tone without one', () => {
    expect(html(Meter, { value: 0.5, tone: 'class', color: 'var(--color-class-mage)' })).toContain(
      'background: var(--color-class-mage)',
    );
    expect(() => html(Meter, { value: 0.5, tone: 'class' })).toThrow(/needs a color/);
  });
});
