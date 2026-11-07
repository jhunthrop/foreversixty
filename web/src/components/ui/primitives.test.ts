// The flat primitives (web layout spec 2026-11-05 §3A): what each renders and the rules it holds.
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Button from './Button.svelte';
import ControlBar from './ControlBar.svelte';
import Field from './Field.svelte';
import Figure from './Figure.svelte';
import FormRow from './FormRow.svelte';
import Inspector from './Inspector.svelte';
import ItemIcon from './ItemIcon.svelte';
import ItemRow from './ItemRow.svelte';
import Meter from './Meter.svelte';
import Notice from './Notice.svelte';
import Overlay from './Overlay.svelte';
import Page from './Page.svelte';
import PageHeader from './PageHeader.svelte';
import Plate from './Plate.svelte';
import Row from './Row.svelte';
import Section from './Section.svelte';

const text = (html: string) => createRawSnippet(() => ({ render: () => html }));

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const BUILD = '1.60.1.70790';
const HOOD = { id: 5404, name: "Serpent's Shoulders", quality: 3, icon: 'inv_shoulder_02' };

describe('Page', () => {
  it('composes header, bar and main, and adds the rail only for split and reading', () => {
    const body = html(Page, {
      layout: 'split',
      header: text('<header>H</header>'),
      bar: text('<div>B</div>'),
      rail: text('<div>R</div>'),
      stickyRail: true,
      children: text('<div>M</div>'),
    });
    expect(body.indexOf('H')).toBeLessThan(body.indexOf('B'));
    expect(body).toContain('data-layout="split"');
    expect(body).toContain('data-sticky="true"');
    expect(body).toMatch(/<aside[^>]*ui-page-rail/);
    const ledger = html(Page, {
      layout: 'ledger',
      rail: text('<div>R</div>'),
      children: text('<div>M</div>'),
    });
    expect(ledger).not.toContain('ui-page-rail');
  });

  it('form: one main column and no rail, whatever the caller passes', () => {
    const form = html(Page, {
      layout: 'form',
      header: text('<header>H</header>'),
      rail: text('<div>R</div>'),
      children: text('<div>M</div>'),
    });
    expect(form).toContain('data-layout="form"');
    expect(form).toMatch(/<div class="ui-page-main[^"]*"><div>M<\/div>/);
    expect(form).not.toContain('ui-page-rail');
  });
});

describe('PageHeader', () => {
  it('sets the one Marcellus title in the class colour, with the answer at the right', () => {
    const body = html(PageHeader, {
      title: 'Beast Mastery Hunter',
      tone: { classSlug: 'hunter' },
      answer: text('<span>110.9 DPS</span>'),
    });
    expect(body.match(/<h1/g)?.length).toBe(1);
    expect(body).toContain('class="ui-title');
    expect(body).toContain('color: var(--color-class-hunter)');
    expect(body).toContain('ui-page-header-answer');
  });

  it('takes the faction colour on a page with no class', () => {
    expect(html(PageHeader, { title: 'Olympus', tone: { faction: 'alliance' } })).toContain(
      'var(--color-alliance)',
    );
  });
});

describe('Plate', () => {
  it('is the one framed surface: raised ground, line edge, plate radius', () => {
    const body = html(Plate, { label: 'Gear', children: text('<p>x</p>') });
    expect(body).toContain('bg-raised');
    expect(body).toContain('border-line');
    expect(body).toContain('rounded-plate');
    expect(body).toContain('role="region"');
    expect(body).not.toContain('card-surface');
  });
});

describe('Section', () => {
  it('is an open heading over a hairline with no box, and a link action', () => {
    const body = html(Section, {
      title: 'Stat weights',
      level: 3,
      meta: 'Normalized to Agility = 1.00',
      action: { label: 'Use these weights', href: '/setup' },
      id: 'weights',
      children: text('<p>rows</p>'),
    });
    expect(body).toMatch(/<h3[^>]*id="weights-title"[^>]*>Stat weights<\/h3>/);
    expect(body).toContain('border-b');
    expect(body).toContain('href="/setup"');
    expect(body).not.toContain('bg-raised');
    expect(body).not.toContain('uppercase');
    // A sticky control bar or section nav never covers a section it scrolls to.
    expect(body).toContain('scroll-mt-[calc(var(--bar-h)+12px)]');
  });
});

describe('Row', () => {
  it('is a link when it has an href, and a plain row otherwise', () => {
    expect(html(Row, { title: 'Olympus', href: '/g/1' })).toMatch(/^<a[^>]*href="\/g\/1"/);
    expect(html(Row, { title: 'Olympus' })).toMatch(/^<div/);
  });

  it('marks the selected row with the gold ground and bar', () => {
    const body = html(Row, { title: 'Head', selected: true, rank: '1' });
    expect(body).toContain('aria-current="true"');
    expect(body).toContain('bg-selected');
    expect(body).toContain('grid-cols-[auto_minmax(0,1fr)_auto]');
  });

  it('colours the title from a class prop, never from the page', () => {
    expect(html(Row, { title: 'Felsnap', titleTone: { classSlug: 'hunter' } })).toContain(
      'var(--color-class-hunter)',
    );
  });

  it('puts a foot of controls on its own full-width line under the body', () => {
    const body = html(Row, { title: 'Seraphine', foot: text('<button>Award</button>') });
    expect(body).toMatch(/<span class="col-span-full[^"]*"><button>Award<\/button><\/span>/);
    expect(html(Row, { title: 'Seraphine' })).not.toContain('col-span-full');
  });
});

describe('ItemIcon', () => {
  it('frames the icon in its quality hue with fixed geometry', () => {
    const body = html(ItemIcon, { icon: 'inv_x', quality: 4, build: BUILD, size: 40 });
    expect(body).toContain(`/data/${BUILD}/icons/inv_x.webp`);
    expect(body).toContain('width="40"');
    expect(body).toContain('var(--color-rarity-epic)');
  });

  it('draws the dark socket with no icon, never a broken image', () => {
    const body = html(ItemIcon, { build: BUILD });
    expect(body).not.toContain('<img');
    expect(body).toContain('border-dashed');
  });

  it('carries the upgrade and worn badges', () => {
    expect(html(ItemIcon, { icon: 'a', build: BUILD, badge: 'upgrade' })).toContain('title="Upgrade"');
    expect(html(ItemIcon, { icon: 'a', build: BUILD, badge: 'worn' })).toContain('title="You wear the pick"');
  });
});

describe('ItemRow', () => {
  it('names the item in its quality colour with the tooltip hook, its source and its gain', () => {
    const body = html(ItemRow, {
      item: HOOD,
      build: BUILD,
      slotLabel: 'Shoulder',
      source: 'Wailing Caverns · Lady Anacondra · 40%',
      gain: { value: '+1.8', caption: 'DPS vs next best' },
    });
    expect(body).toContain('Shoulder');
    expect(body).toContain('data-bis-item="5404"');
    expect(body).toContain('data-testid="item-hover-5404"');
    expect(body).toContain('text-rarity-rare');
    expect(body).toContain('Wailing Caverns · Lady Anacondra · 40%');
    expect(body).toContain('+1.8');
    expect(body).toContain('DPS vs next best');
  });

  it("reads the empty slot's reason", () => {
    const body = html(ItemRow, { item: null, build: BUILD, emptyDetail: 'First trinket comes at 40' });
    expect(body).toContain('Nothing better yet');
    expect(body).toContain('First trinket comes at 40');
  });

  it('shows the worn item beside the gain, with its own tooltip', () => {
    const body = html(ItemRow, {
      item: HOOD,
      build: BUILD,
      worn: { id: 1, name: 'Old', quality: 2, icon: 'b' },
    });
    expect(body).toContain('data-testid="item-row-worn"');
    expect(body).toContain('title="You wear Old"');
  });

  it('keeps the name out of the tab order inside a roving container', () => {
    expect(html(ItemRow, { item: HOOD, build: BUILD, focusableName: false })).not.toContain('tabindex="0"');
  });

  it('is one pressed button when it picks its item, its name no tab stop of its own', () => {
    const body = html(ItemRow, { item: HOOD, build: BUILD, onselect: () => {}, selected: true });
    expect(body).toMatch(/<button[^>]*type="button"[^>]*aria-pressed="true"/);
    expect(body).toContain('bg-selected');
    expect(body).not.toContain('tabindex="0"');
    expect(html(ItemRow, { item: HOOD, build: BUILD })).not.toContain('<button');
  });

  it("prints the gain's unit after the figure", () => {
    const body = html(ItemRow, {
      item: HOOD,
      build: BUILD,
      gain: { value: '+1.8', unit: 'DPS', caption: 'vs your Old' },
    });
    expect(body).toMatch(/\+1\.8<span[^>]*>DPS<\/span>/);
  });

  it('takes a trail in place of the gain', () => {
    expect(html(ItemRow, { item: HOOD, build: BUILD, trail: text('<b>Open</b>') })).toContain('<b>Open</b>');
  });
});

describe('Figure', () => {
  it('is the body face 700 with tabular figures, never mono or condensed', () => {
    const body = html(Figure, { value: '110.9', unit: 'DPS', label: 'This set', tone: 'gold', size: 'big' });
    expect(body).toContain('tabular-nums');
    expect(body).toContain('font-bold');
    expect(body).toContain('text-gold');
    expect(body).not.toMatch(/font-mono|figure-stat|heading/);
  });
});

describe('Meter', () => {
  it('clamps its fill and is decorative unless labelled', () => {
    expect(html(Meter, { value: 1.4 })).toContain('width: 100.0%');
    expect(html(Meter, { value: 0.5 })).toContain('aria-hidden="true"');
    expect(html(Meter, { value: 0.5, label: 'Agility' })).toContain('role="meter"');
  });

  it('lights the share of its segments', () => {
    const body = html(Meter, { value: 0.5, segments: 4, tone: 'kill' });
    expect(body.match(/bg-kill/g)?.length).toBe(2);
  });
});

describe('Notice, Overlay', () => {
  it('Notice is a sentence and a link with no box', () => {
    const body = html(Notice, { text: 'Raids open Thursday.', action: { label: 'Plan', href: '/p' } });
    expect(body).toContain('Raids open Thursday.');
    expect(body).not.toContain('border');
  });

  it('Overlay draws the one overlay surface', () => {
    const body = html(Overlay, { kind: 'menu', children: text('<p>x</p>') });
    expect(body).toContain('shadow-overlay');
    expect(body).toContain('role="menu"');
  });
});

describe('Inspector', () => {
  it('names the selection and labels each section', () => {
    const body = html(Inspector, {
      selection: 'Head',
      sections: [
        { label: 'Where to get it', body: text('<p>WC</p>') },
        { label: 'Runners-up · DPS against the pick', body: text('<p>r</p>') },
      ],
    });
    expect(body).toContain('aria-label="Head detail"');
    expect(body).toContain('data-placement="side"');
    expect(body).toContain('Where to get it');
    expect(body.match(/<section/g)?.length).toBe(2);
  });
});

describe('Field, FormRow, Button', () => {
  it('Field ties its error to the input', () => {
    const body = html(Field, { label: 'Guild', id: 'guild', error: 'No guild by that name.' });
    expect(body).toContain('aria-describedby="guild-error"');
    expect(body).toContain('aria-invalid="true"');
  });

  it('FormRow is a labelled form with one submit', () => {
    const body = html(FormRow, {
      label: 'Open a guild',
      children: text('<span>f</span>'),
      submit: text('<button>Go</button>'),
    });
    expect(body).toMatch(/<form[^>]*aria-label="Open a guild"/);
  });

  it('Button is a link when it navigates and a button otherwise', () => {
    expect(html(Button, { href: '/x', children: text('Go') })).toMatch(/^<a/);
    expect(html(Button, { children: text('Go') })).toMatch(/^<button/);
  });
});

describe('ControlBar', () => {
  const groups = [
    {
      label: 'Level',
      param: 'band',
      kind: 'segmented' as const,
      value: '20',
      options: [
        { value: '20', label: '20-29' },
        { value: '60', label: '60', disabled: true, disabledReason: 'No data yet at this level' },
      ],
    },
    {
      label: 'Spec',
      param: 'spec',
      kind: 'segmented' as const,
      value: 'bm',
      options: [{ value: 'bm', label: 'BM', href: '/bis/hunter/bm' }],
    },
  ];

  it('renders each segmented group as a radio group of links, the URL being the truth', () => {
    const body = html(ControlBar, { groups, label: 'Best in slot choices' });
    expect(body).toContain('role="radiogroup"');
    expect(body).toContain('href="?band=20"');
    expect(body).toContain('href="/bis/hunter/bm"');
    expect(body).toMatch(/data-value="20"[^>]*|aria-checked="true"/);
    expect(body).toContain('aria-disabled="true"');
    expect(body).toContain('title="No data yet at this level"');
  });

  it('carries the phone summary line and its toggle', () => {
    const body = html(ControlBar, { groups, label: 'x' });
    expect(body).toContain('20-29 · BM');
    expect(body).toContain('data-control-bar-toggle');
    expect(body).toContain('aria-expanded="false"');
  });
});
