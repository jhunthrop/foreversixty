// web/tests/e2e/support/parity.ts
// The mock parity gate (design/parity.md): a built page against its committed mock, region by
// region. Both sides carry the same `data-region` names (src/lib/ui/regions.ts and the mock
// library's ws_lib.REGION). For every region the gate compares
//   - the inventory: the region's accessibility tree as role chains ("listbox>option",
//     "paragraph>link", "heading[1]"), names and text ignored, plus a census of visible
//     disclosures (details, summary, [aria-expanded]). A block on one side and not the other
//     fails, whatever its text;
//   - the box: x, y (y from the top of <main>, so the site header's own height is not every
//     region's error) and width, height, within PARITY_TOLERANCE.
// The mock and the build show the same data (design/mocks/bis_facts.mjs), so a difference is a
// difference in design, never in data.
import { readFileSync, mkdirSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import type { Browser, Page } from '@playwright/test';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../..');
export const REPO_ROOT = path.resolve(WEB_ROOT, '..');
export const MOCKS_DIR = path.join(REPO_ROOT, 'design/mocks/web');
export const PARITY_OUT = path.join(WEB_ROOT, 'test-results/parity');

/**
 * Position within 8 px: the two sides share tokens and type, and one text line is 18 to 24 px,
 * so 8 px absorbs sub-pixel and font-hinting drift (Linux CI against macOS) and a 1 to 4 px
 * padding nudge, never a line, a control or a block. Size within 8 px or 5%, whichever is larger:
 * a region's height grows with its rows, and 5% of the BiS plate (about 600 px) is 30 px, one row
 * that wraps to a third line, while a new row (65 px) or a second bar line (46 px) fails.
 */
export const PARITY_TOLERANCE = { position: 8, sizeAbs: 8, sizeRel: 0.05 } as const;

export interface ParityState {
  id: string;
  path: string;
  mock: string;
}

export interface ParityConfig {
  viewport: { width: number; height: number };
  states: ParityState[];
}

export interface RegionBox {
  x: number;
  y: number;
  width: number;
  height: number;
}

export interface RegionSnapshot {
  name: string;
  box: RegionBox;
  inventory: Record<string, number>;
}

export interface ParityFailure {
  region: string;
  kind: 'missing' | 'extra' | 'inventory' | 'box';
  detail: string;
}

export function readParityConfig(file: string): ParityConfig {
  return JSON.parse(readFileSync(path.join(MOCKS_DIR, file), 'utf8')) as ParityConfig;
}

const FONT_FILES: readonly [family: string, weight: number, file: string][] = [
  ['Barlow', 400, 'barlow-400-subset.woff2'],
  ['Barlow', 600, 'barlow-600-subset.woff2'],
  ['Barlow', 700, 'barlow-700-subset.woff2'],
  ['Marcellus', 400, 'marcellus-400-subset.woff2'],
];

/** The site's own font files as @font-face rules, for the mock's Google Fonts request: both
 *  sides then lay text out with the same faces, offline. */
function siteFontCss(): string {
  return FONT_FILES.map(([family, weight, file]) => {
    const data = readFileSync(path.join(WEB_ROOT, 'src/assets/fonts', file)).toString('base64');
    return `@font-face{font-family:'${family}';font-weight:${weight};font-style:normal;src:url(data:font/woff2;base64,${data}) format('woff2')}`;
  }).join('\n');
}

const FREEZE_CSS =
  '*,*::before,*::after{animation:none!important;transition:none!important;caret-color:transparent!important}';

/** Before navigation: fonts served locally, motion off. */
export async function prepareParityPage(page: Page): Promise<void> {
  const css = siteFontCss();
  await page.route('https://fonts.googleapis.com/**', (route) =>
    route.fulfill({ status: 200, contentType: 'text/css', body: css }),
  );
  await page.route('https://fonts.gstatic.com/**', (route) => route.abort());
  await page.emulateMedia({ reducedMotion: 'reduce' });
}

/** After the page is ready: animations frozen, fonts loaded, scrolled to the top. */
export async function settleParityPage(page: Page): Promise<void> {
  await page.addStyleTag({ content: FREEZE_CSS });
  await page.evaluate(async () => {
    await document.fonts.ready;
    window.scrollTo(0, 0);
  });
}

export async function openMock(page: Page, file: string): Promise<void> {
  await prepareParityPage(page);
  await page.goto(pathToFileURL(path.join(MOCKS_DIR, file)).href);
  await settleParityPage(page);
}

/** One line of an aria snapshot: its depth and role ("heading[1]"), or nothing for text and
 *  props (`- text: …`, `- /url: …`). */
function snapshotLine(line: string): { depth: number; role: string } | undefined {
  const match = /^(\s*)- ['"]?([a-z]+)\b(.*)$/i.exec(line);
  if (match === null) return undefined;
  const [, indent, role, rest] = match;
  if (role === 'text') return undefined;
  const level = /\[level=(\d)\]/.exec(rest);
  return { depth: indent.length / 2, role: level ? `${role}[${level[1]}]` : role };
}

/** A region's role chains, counted. A combobox's options are its data, not blocks: skipped. */
export function inventoryOf(ariaSnapshot: string): Record<string, number> {
  const counts: Record<string, number> = {};
  const stack: string[] = [];
  for (const line of ariaSnapshot.split('\n')) {
    const parsed = snapshotLine(line);
    if (parsed === undefined) continue;
    stack.length = parsed.depth;
    const inSelect = stack.includes('combobox');
    stack.push(parsed.role);
    if (inSelect) continue;
    const key = stack.join('>');
    counts[key] = (counts[key] ?? 0) + 1;
  }
  return counts;
}

/** Every visible region on the page: its box from the top of <main> and its inventory. */
export async function captureRegions(page: Page): Promise<RegionSnapshot[]> {
  const raw = await page.evaluate(() => {
    const top = (document.querySelector('main')?.getBoundingClientRect().top ?? 0) + window.scrollY;
    const seen = new Map<string, number>();
    const out: { name: string; index: number; box: RegionBox; disclosures: Record<string, number> }[] = [];
    document.querySelectorAll<HTMLElement>('[data-region]').forEach((el, index) => {
      const rect = el.getBoundingClientRect();
      if (!el.checkVisibility({ visibilityProperty: true }) || rect.width === 0 || rect.height === 0) return;
      const base = el.dataset.region ?? '';
      const n = (seen.get(base) ?? 0) + 1;
      seen.set(base, n);
      const disclosures: Record<string, number> = {};
      // A select's own popup is not a disclosure: comboboxes are left to the role inventory.
      el.querySelectorAll<HTMLElement>('details, summary, [aria-expanded]:not([role="combobox"])').forEach(
        (d) => {
          if (!d.checkVisibility({ visibilityProperty: true })) return;
          const key = `disclosure:${d.tagName.toLowerCase()}`;
          disclosures[key] = (disclosures[key] ?? 0) + 1;
        },
      );
      out.push({
        name: n === 1 ? base : `${base}#${n}`,
        index,
        box: {
          x: Math.round(rect.left + window.scrollX),
          y: Math.round(rect.top + window.scrollY - top),
          width: Math.round(rect.width),
          height: Math.round(rect.height),
        },
        disclosures,
      });
    });
    return out;
  });
  const regions = page.locator('[data-region]');
  const snapshots: RegionSnapshot[] = [];
  for (const region of raw) {
    const aria = await regions.nth(region.index).ariaSnapshot();
    snapshots.push({
      name: region.name,
      box: region.box,
      inventory: { ...inventoryOf(aria), ...region.disclosures },
    });
  }
  return snapshots;
}

function sizeOff(mock: number, build: number): boolean {
  return Math.abs(mock - build) > Math.max(PARITY_TOLERANCE.sizeAbs, mock * PARITY_TOLERANCE.sizeRel);
}

function inventoryFailures(name: string, mock: RegionSnapshot, build: RegionSnapshot): ParityFailure[] {
  const keys = new Set([...Object.keys(mock.inventory), ...Object.keys(build.inventory)]);
  const failures: ParityFailure[] = [];
  for (const key of [...keys].sort()) {
    const m = mock.inventory[key] ?? 0;
    const b = build.inventory[key] ?? 0;
    if (m !== b) failures.push({ region: name, kind: 'inventory', detail: `${key}: mock ${m}, build ${b}` });
  }
  return failures;
}

function boxFailures(name: string, mock: RegionBox, build: RegionBox): ParityFailure[] {
  const moved =
    Math.abs(mock.x - build.x) > PARITY_TOLERANCE.position ||
    Math.abs(mock.y - build.y) > PARITY_TOLERANCE.position;
  const resized = sizeOff(mock.width, build.width) || sizeOff(mock.height, build.height);
  if (!moved && !resized) return [];
  const fmt = (b: RegionBox) => `${b.x},${b.y} ${b.width}x${b.height}`;
  return [{ region: name, kind: 'box', detail: `mock ${fmt(mock)}, build ${fmt(build)}` }];
}

/** Every difference between the two sides, region by region. */
export function compareRegions(mock: RegionSnapshot[], build: RegionSnapshot[]): ParityFailure[] {
  const byName = (list: RegionSnapshot[]) => new Map(list.map((r) => [r.name, r]));
  const m = byName(mock);
  const b = byName(build);
  const failures: ParityFailure[] = [];
  for (const name of new Set([...m.keys(), ...b.keys()])) {
    const mockRegion = m.get(name);
    const buildRegion = b.get(name);
    if (mockRegion === undefined) {
      failures.push({ region: name, kind: 'extra', detail: 'the build has this region, the mock does not' });
    } else if (buildRegion === undefined) {
      failures.push({
        region: name,
        kind: 'missing',
        detail: 'the mock has this region, the build does not',
      });
    } else {
      failures.push(
        ...inventoryFailures(name, mockRegion, buildRegion),
        ...boxFailures(name, mockRegion.box, buildRegion.box),
      );
    }
  }
  return failures;
}

export function formatFailures(failures: readonly ParityFailure[]): string {
  return failures.map((f) => `  [${f.region}] ${f.kind}: ${f.detail}`).join('\n');
}

const OUTLINE_LAYER = 'parity-outline-layer';

/** A full-page capture with every region outlined and named (an overlay; the page is not restyled). */
export async function captureOutlined(page: Page, regions: readonly RegionSnapshot[]): Promise<Buffer> {
  await page.evaluate(
    ({ regions: list, layerId }) => {
      const top = (document.querySelector('main')?.getBoundingClientRect().top ?? 0) + window.scrollY;
      const layer = document.createElement('div');
      layer.id = layerId;
      layer.style.cssText =
        'position:absolute;left:0;top:0;width:0;height:0;z-index:2147483647;pointer-events:none';
      for (const r of list) {
        const box = document.createElement('div');
        box.style.cssText = `position:absolute;left:${r.box.x}px;top:${r.box.y + top}px;width:${r.box.width}px;height:${r.box.height}px;outline:2px solid #ff3df0;outline-offset:-1px`;
        const tag = document.createElement('span');
        tag.textContent = r.name;
        tag.style.cssText =
          'position:absolute;right:0;top:0;font:700 11px/14px sans-serif;background:#ff3df0;color:#000;padding:0 3px';
        box.append(tag);
        layer.append(box);
      }
      document.body.append(layer);
    },
    { regions: regions.map((r) => ({ name: r.name, box: r.box })), layerId: OUTLINE_LAYER },
  );
  const png = await page.screenshot({ fullPage: true, animations: 'disabled' });
  await page.evaluate((id) => document.getElementById(id)?.remove(), OUTLINE_LAYER);
  return png;
}

/** mock | build, the same scale, top aligned: test-results/parity/<name>.png. */
export async function writeSideBySide(
  browser: Browser,
  name: string,
  mock: Buffer,
  build: Buffer,
): Promise<string> {
  mkdirSync(PARITY_OUT, { recursive: true });
  const page = await browser.newPage({ viewport: { width: 2 * 1440 + 36, height: 800 } });
  const img = (png: Buffer, label: string) =>
    `<figure style="margin:0"><figcaption style="font:700 16px/28px sans-serif;color:#e5b955">${label}</figcaption><img style="display:block;width:1440px" src="data:image/png;base64,${png.toString('base64')}"></figure>`;
  await page.setContent(
    `<body style="margin:0;padding:12px;background:#1c1c1e;display:flex;gap:12px;align-items:flex-start">${img(mock, `MOCK ${name}`)}${img(build, `BUILD ${name}`)}</body>`,
  );
  const out = path.join(PARITY_OUT, `${name}.png`);
  writeFileSync(out, await page.screenshot({ fullPage: true }));
  await page.close();
  return out;
}
