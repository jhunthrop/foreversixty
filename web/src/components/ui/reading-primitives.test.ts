// web/src/components/ui/reading-primitives.test.ts
// The guides round's primitives (guides spec 2026-11-08 §3A): ReadingSection, LaneBlock, Toc,
// PlateColumns and PlateColumn, Prose, LevelSlider; PageHeader's lead and the
// reading Page's top and contents rail. TalentTrees: talent-tree.test.ts.
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import LaneBlock from './LaneBlock.svelte';
import LevelSlider from './LevelSlider.svelte';
import Page from './Page.svelte';
import PageHeader from './PageHeader.svelte';
import PlateColumn from './PlateColumn.svelte';
import PlateColumns from './PlateColumns.svelte';
import Prose from './Prose.svelte';
import ReadingSection from './ReadingSection.svelte';
import Toc from './Toc.svelte';

const text = (html: string) => createRawSnippet(() => ({ render: () => html }));

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

describe('ReadingSection', () => {
  it('is a section region named for its title, anchored, with the lane before the prose in the DOM', () => {
    const body = html(ReadingSection, {
      title: 'Rotation and priority',
      anchor: 'rotation-and-priority',
      lane: text('<div>LANE</div>'),
      children: text('<p>PROSE</p>'),
    });
    expect(body).toMatch(
      /<section id="rotation-and-priority"[^>]*data-region="section-rotation-and-priority"/,
    );
    expect(body.match(/<h2[^>]*>Rotation and priority<\/h2>/g)).toHaveLength(1);
    expect(body).toMatch(/class="ui-rd-lane[^"]*"[^>]*><div>LANE<\/div>/);
    expect(body.indexOf('LANE')).toBeLessThan(body.indexOf('PROSE'));
    expect(body).not.toContain('data-solo');
  });

  it('stays solo at the measure with no lane', () => {
    const body = html(ReadingSection, { title: 'Professions', children: text('<p>PROSE</p>') });
    expect(body).toContain('data-solo=""');
    expect(body).not.toContain('ui-rd-lane');
    expect(body).toContain('data-region="section-professions"');
  });

  it('draws no surface', () => {
    const body = html(ReadingSection, { title: 'X', lane: text('<i></i>'), children: text('<p></p>') });
    expect(body).not.toMatch(/bg-raised|rounded-plate|card-surface/);
  });
});

describe('LaneBlock', () => {
  it('labels its rows with an h3 heading and an optional meta at its right', () => {
    const body = html(LaneBlock, {
      label: 'Priority',
      meta: 'Level 60',
      testid: 'lane',
      children: text('<ol><li>Row</li></ol>'),
    });
    expect(body).toMatch(/data-testid="lane"/);
    expect(body).toMatch(/<h3 class="[^"]*ui-lane-label[^"]*">Priority<\/h3>/);
    expect(body).not.toContain('uppercase');
    expect(body).toMatch(/ml-auto[^"]*">Level 60<\/span>/);
    expect(body.indexOf('Priority')).toBeLessThan(body.indexOf('<ol>'));
    expect(html(LaneBlock, { label: 'L', children: text('<i></i>') })).not.toContain('ml-auto');
  });
});

describe('Toc', () => {
  const items = [
    { title: 'Overview', id: 'overview' },
    { title: 'Stat priority', id: 'stat-priority' },
  ];

  it('links each section by its anchor and marks the current one', () => {
    const body = html(Toc, { label: 'On this page', items, current: 'stat-priority', testid: 'toc' });
    expect(body).toContain('data-testid="toc"');
    expect(body).toContain('>On this page<');
    expect(body.match(/<li>/g)).toHaveLength(2);
    expect(body).toMatch(/href="#overview" data-toc-link="overview"(?: class="[^"]*")?>Overview</);
    expect(body).toMatch(
      /href="#stat-priority" data-toc-link="stat-priority" aria-current="location"[^>]*>Stat priority</,
    );
    expect(body.match(/aria-current/g)).toHaveLength(1);
  });

  it('marks none without a current', () => {
    expect(html(Toc, { label: 'L', items })).not.toContain('aria-current');
  });
});

describe('PlateColumns, PlateColumn', () => {
  it('is one plate region with its label, holding labelled columns with their foot', () => {
    const column = html(PlateColumn, {
      label: 'Build',
      meta: 'Level 60',
      foot: text('<a>Load</a>'),
      testid: 'col',
      children: text('<div>TREES</div>'),
    });
    expect(column).toMatch(/data-testid="col"/);
    expect(column).toMatch(/<h2 class="[^"]*text-label[^"]*">Build<\/h2>/);
    expect(column).not.toContain('uppercase');
    expect(column).toMatch(/ml-auto[^"]*">Level 60<\/span>/);
    expect(column).toMatch(/TREES<\/div>\s*<div class="ui-pcol-foot[^"]*"><a>Load<\/a>/);
    expect(html(PlateColumn, { label: 'X', children: text('<i></i>') })).not.toContain('ui-pcol-foot');

    const plate = html(PlateColumns, {
      count: 3,
      label: 'At a glance',
      testid: 'glance',
      children: text(column),
    });
    expect(plate.match(/role="region"/g)).toHaveLength(1);
    expect(plate).toContain('aria-label="At a glance"');
    expect(plate.match(/data-region="plate"/g)).toHaveLength(1);
    expect(plate).toContain('data-testid="glance"');
    expect(plate).toContain('--pcols: 3');
    expect(plate).toContain('data-testid="col"');
  });
});

describe('Prose', () => {
  it('takes the build’s own html in, with its modifier classes', () => {
    const body = html(Prose, {
      html: '<ul><li><strong>A</strong></li></ul>',
      mods: ['ruled', 'cols'],
      testid: 'p',
    });
    expect(body).toMatch(
      /<div class="ui-prose ui-prose--ruled ui-prose--cols[^"]*" data-testid="p"><ul><li><strong>A<\/strong>/,
    );
  });

  it('is the bare typography with no modifier', () => {
    expect(html(Prose, { html: '<p>x</p>' })).toMatch(
      /<div class="ui-prose(?: svelte-[\w]+)?"><p>x<\/p><\/div>/,
    );
  });
});

describe('LevelSlider', () => {
  const base = {
    min: 10,
    max: 60,
    label: 'Level',
    previousLabel: 'Previous level',
    nextLabel: 'Next level',
    onchange: () => undefined,
  };
  const disabledOf = (body: string, part: 'previous' | 'next'): boolean =>
    new RegExp(`aria-label="${part === 'previous' ? 'Previous' : 'Next'} level" disabled`).test(body);

  it('is a range with its value and spoken value text between two named steps', () => {
    const body = html(LevelSlider, { ...base, level: 34, valueText: 'Level 34', testid: 'lvl' });
    expect(body).toMatch(
      /<input type="range"[^>]*min="10" max="60" step="1" value="34" aria-label="Level" aria-valuetext="Level 34"[^>]*data-testid="lvl-range"/,
    );
    expect(body).toContain('data-testid="lvl-previous"');
    expect(body).toContain('data-testid="lvl-next"');
    expect(disabledOf(body, 'previous')).toBe(false);
    expect(disabledOf(body, 'next')).toBe(false);
    expect(body).toContain(`--fill: ${((34 - 10) / 50) * 100}%`);
  });

  it('turns previous off at the first level and next off at the last', () => {
    const first = html(LevelSlider, { ...base, level: 10, valueText: 'Level 10' });
    expect(disabledOf(first, 'previous')).toBe(true);
    expect(disabledOf(first, 'next')).toBe(false);
    const last = html(LevelSlider, { ...base, level: 60, valueText: 'Level 60' });
    expect(disabledOf(last, 'previous')).toBe(false);
    expect(disabledOf(last, 'next')).toBe(true);
  });

  it('turns every control off when disabled', () => {
    const body = html(LevelSlider, { ...base, level: 30, valueText: 'Level 30', disabled: true });
    expect(disabledOf(body, 'previous')).toBe(true);
    expect(disabledOf(body, 'next')).toBe(true);
    expect(body).toMatch(/<input type="range"[^>]*disabled/);
  });
});

describe('PageHeader lead', () => {
  it('sets the one sentence directly under the title, the meta line under it, and nothing without it', () => {
    const body = html(PageHeader, {
      title: 'Frost Mage',
      tone: { classSlug: 'mage' },
      meta: text('<span>DPS</span>'),
      lead: 'Frost casts Frostbolt behind Ice Barrier.',
    });
    expect(body).toMatch(
      /<p class="ui-page-header-lead[^"]*">Frost casts Frostbolt behind Ice Barrier\.<\/p>/,
    );
    // seo-fix 2026-11-10: the lead is the page's first sentence, never under a caveat line.
    expect(body.indexOf('Frost casts')).toBeLessThan(body.indexOf('DPS'));
    expect(html(PageHeader, { title: 'Frost Mage' })).not.toContain('ui-page-header-lead');
  });
});

describe('Page reading', () => {
  it('renders the top full width, then the contents nav, then main, and never a rail', () => {
    const body = html(Page, {
      layout: 'reading',
      header: text('<header>H</header>'),
      top: text('<div>PLATE</div>'),
      toc: text('<div>TOC</div>'),
      rail: text('<div>RAIL</div>'),
      children: text('<div>MAIN</div>'),
    });
    expect(body).toContain('data-layout="reading"');
    expect(body).toMatch(/<div class="ui-page-top[^"]*"><div>PLATE<\/div><\/div>/);
    expect(body).toMatch(/<nav class="ui-page-toc[^"]*" aria-label="Contents"><div>TOC<\/div><\/nav>/);
    expect(body).toMatch(/<div class="ui-page-main[^"]*"><div>MAIN<\/div>/);
    expect(body.indexOf('PLATE')).toBeLessThan(body.indexOf('TOC'));
    expect(body.indexOf('TOC')).toBeLessThan(body.indexOf('MAIN'));
    expect(body).not.toContain('RAIL');
    expect(body).not.toContain('ui-page-rail');
  });

  it('leaves out the top and the nav it is not handed', () => {
    const body = html(Page, { layout: 'reading', children: text('<div>MAIN</div>') });
    expect(body).not.toContain('ui-page-top');
    expect(body).not.toContain('<nav');
  });
});
