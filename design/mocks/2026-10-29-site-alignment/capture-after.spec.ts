// design/mocks/2026-10-29-site-alignment/capture-after.spec.ts
// The after-mocks for the four highest-traffic routes (spec design/specs/2026-10-29-site-alignment.md
// section 6): the built fixture page, signed in, restyled in the browser by after-kit.js to the
// spec's rules. Writes after-<route>-<width>.png beside this file. A mock, not a test.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { isolate, signIn } from './stubs';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const SETTLE_MS = 1500;

type Transform = () => void;

/** /bis/warrior/fury: the three panels, the list and the footer panels become Cards. */
const bisSpec: Transform = () => {
  const { byText, boxOf, cardify, band, surfaceOf } = (window as any).__align;
  band(document.querySelector('.art-panel'));
  surfaceOf(document.querySelector('[data-testid="character-header"]'));
  const thisSet = document.querySelector('[data-testid="bis-this-set"]');
  cardify(thisSet, { glyph: 'bolt', title: 'This set', aside: 'Simmed 29 Oct', drop: [byText('This set', thisSet)] });
  const weights = document.querySelector('[data-testid="bis-weights-alliance-20"]');
  cardify(weights, { glyph: 'scale', title: 'Stat weights', aside: '20 to 29', drop: [byText('Stat weights at 20 to 29', weights)] });
  const play = document.querySelector('[data-testid="bis-play-it"]');
  cardify(play, { glyph: 'sword', title: 'Play it', aside: '20 to 29', drop: [byText('Play it · 20 to 29', play)] });
  const list = document.querySelector('.the-list-panel');
  cardify(list, { glyph: 'list', title: 'The list', aside: 'Hover or tap an item for its stats', inset: true,
    drop: [byText('The list'), byText('Hover or tap an item for its stats')] });
  const groups = document.querySelector('[data-testid="bis-source-groups-alliance-20"]');
  cardify(groups, { glyph: 'pin', title: 'Where to get it', drop: [byText('Where to get it')] });
  const fresh = document.querySelector('[data-testid="bis-new-at-band-alliance-20"]');
  cardify(fresh, { glyph: 'star', title: 'New at this range', drop: [byText('New at this level range')] });
};

/** /guides/warrior/fury: the top trio become Cards; the body stays prose; tables go quiet. */
const guide: Transform = () => {
  const { byText, boxOf, cardify, band } = (window as any).__align;
  band(document.querySelector('.art-panel'));
  const q = (id: string) => document.querySelector(`[data-testid="${id}"]`);
  const label = (el: Element | null) => el?.querySelector(':scope > .label') ?? null;
  const build = q('guide-rail-build');
  cardify(build, { glyph: 'tree', title: 'Build', aside: 'Arms 17 · Fury 32 · Prot 2', drop: [label(build)] });
  const rot = q('guide-rail-rotation');
  cardify(rot, { glyph: 'sword', title: 'Rotation', aside: 'Level 60', drop: [label(rot)] });
  const stat = q('guide-rail-stat-priority');
  cardify(stat, { glyph: 'target', title: 'Stat priority', aside: 'Hit = 1.00', drop: [label(stat)] });
  document.querySelectorAll('.guide-table-box').forEach((el) => {
    (el as HTMLElement).style.setProperty('background', 'var(--surface-card)', 'important');
    (el as HTMLElement).style.setProperty('border-color', 'var(--color-line-card)', 'important');
    (el as HTMLElement).style.setProperty('border-radius', '12px', 'important');
  });
  document.querySelectorAll('.overview-card, .guide-status-note').forEach((el) => {
    (el as HTMLElement).style.setProperty('background', 'var(--surface-card)', 'important');
    (el as HTMLElement).style.setProperty('border-radius', '12px', 'important');
  });
};

/** /planner: the band, the character card, and the rail panels as Cards; the trees keep their canvas. */
const planner: Transform = () => {
  const { byText, cardify, band, surfaceOf } = (window as any).__align;
  band(document.querySelector('.planner-header-band-gradient'));
  surfaceOf(document.querySelector('[data-testid="character-header"]'));
  surfaceOf(document.querySelector('[data-testid="planner-export-not-read"]'), true);
  document.querySelectorAll('[data-testid^="tree-panel-"]').forEach((el) => surfaceOf(el, true));
  const imp = document.querySelector('[data-testid="import-box"]');
  cardify(imp, { glyph: 'download', title: 'Import from addon', dropFirst: true });
  const band60 = document.querySelector('[data-testid="band-compare"]');
  cardify(band60, { glyph: 'layers', title: 'Level 60 build', aside: 'Your level’s range',
    drop: [byText('Level 60 build', band60), byText('your level’s range', band60)] });
  const order = document.querySelector('[data-testid="order-strip"]');
  cardify(order, { glyph: 'list', title: 'Point order', aside: '0 spent', dropFirst: true });
  const share = document.querySelector('[data-testid="planner-share-section"]');
  const reset = share?.firstElementChild?.querySelector('button');
  cardify(share, { glyph: 'share', title: 'Share this build', dropFirst: true });
  if (reset) share!.querySelector('.x-head')!.appendChild(Object.assign(reset, { style: 'margin-left:auto' }));
};

/** /: every section under the hero becomes a Card with a CardHead and an ActionLink. */
const home: Transform = () => {
  const { byText, boxOf, cardify, wrap, rows, surfaceOf } = (window as any).__align;
  surfaceOf(document.querySelector('[data-testid="character-header"]'));
  const talents = byText('Talents');
  cardify(boxOf(talents), { glyph: 'tree', title: 'Talents', drop: [talents] });
  const sim = byText('Simulator');
  cardify(boxOf(sim), { glyph: 'bolt', title: 'Simulator', drop: [sim] });
  const up = byText('Your upgrades');
  const upBox = boxOf(byText('Not available yet: no gear export for this character. Open the addon once to send it.')!);
  if (up && upBox) cardify(upBox, { glyph: 'list', title: 'Your upgrades', aside: 'Level 60', drop: [up] });
  const look = byText('Look up a character');
  const lookSection = look?.closest('section');
  if (lookSection) {
    const sub = byText('See any character’s Forever Score.', lookSection);
    cardify(lookSection, { glyph: 'search', title: 'Look up a character', aside: 'See any character’s Forever Score.', drop: [look, sub] });
  }
  const started = byText('Get started');
  const list = started?.nextElementSibling;
  if (started && list) {
    cardify(list, { glyph: 'star', title: 'Get started', drop: [started] });
    rows(list.querySelector('.x-body'), [
      { glyph: 'tree', label: 'Plan your talents', sub: 'Pick a build and share it by link.' },
      { glyph: 'bolt', label: 'Sim your gear', sub: 'See what each upgrade is worth to your damage.' },
      { glyph: 'upload', label: 'Upload tonight’s raid', sub: 'Read every pull, and who did what.' },
    ]);
    list.querySelector('.x-body')!.classList.add('x-bleed');
  }
  const another = byText('Another class');
  const classBox = another ? boxOf(another.parentElement!.parentElement!.querySelector('img') ?? another) : null;
  if (another && classBox) cardify(classBox, { glyph: 'user', title: 'Another class', action: 'All 28 specs', drop: [another.parentElement] });
  const recent = byText('Recent reports');
  const recentBox = recent ? boxOf(byText('A sample report')!) : null;
  if (recent && recentBox) cardify(recentBox, { glyph: 'chart', title: 'Recent reports', action: 'See all', drop: [recent.parentElement] });
  const top = byText('Top guilds');
  const topBox = top ? boxOf(byText('Rankings did not load.')!) : null;
  if (top && topBox) cardify(topBox, { glyph: 'trophy', title: 'Top guilds', action: 'See all', drop: [top.parentElement] });
  const latest = byText('Latest');
  const latestBox = latest ? boxOf(byText('The log coach is here.')!) : null;
  if (latest && latestBox) cardify(latestBox, { glyph: 'news', title: 'Latest', action: 'Everything we know', bleed: true, drop: [latest.parentElement] });
};

const ROUTES: { name: string; url: string; run: Transform }[] = [
  { name: 'bis-warrior-fury', url: '/bis/warrior/fury', run: bisSpec },
  { name: 'guide-warrior-fury', url: '/guides/warrior/fury', run: guide },
  { name: 'planner', url: '/planner', run: planner },
  { name: 'home-signed-in', url: '/', run: home },
];

async function render(page: Page, route: (typeof ROUTES)[number], width: number): Promise<void> {
  await page.setViewportSize({ width, height: 900 });
  await isolate(page);
  await signIn(page);
  await page.goto(route.url);
  await page.waitForLoadState('networkidle').catch(() => undefined);
  await page.waitForTimeout(SETTLE_MS);
  await page.addScriptTag({ path: path.join(HERE, 'after-kit.js') });
  await page.evaluate(route.run);
  await page.waitForTimeout(200);
  await page.screenshot({ path: path.join(HERE, `after-${route.name}-${width}.png`), fullPage: true });
}

for (const route of ROUTES) {
  for (const width of [1440, 390]) {
    test(`after ${route.name} ${width}`, async ({ page }) => render(page, route, width));
  }
}
