// web/src/pages/_setup.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts and _classes.test.ts:
// everything else under src/pages/ is a route, so an unprefixed setup.test.ts would build as
// the route /setup.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// /setup in the flat web system (design/specs/2026-11-07-setup.md): one island
// (components/setup/SetupPage.svelte) whose server markup is the signed-out page. Every download
// links to the site's own /dl/ paths (lib/downloads/paths.ts), asserted off the rendered output.
// The step rows' states are lib/setup/steps.test.ts; the account reads are lib/setup/*.test.ts.
import { existsSync, readFileSync, readdirSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { createPageContainer } from '../test-support/page-container';
import { describe, expect, it, beforeAll } from 'vitest';
import { addonCopy } from '../lib/addon/copy';
import { SETUP_NAV_ITEM } from '../lib/nav';
import { COMPANION_DOWNLOADS } from '../lib/setup/downloads';
import { CURSEFORGE_URL, WAGO_URL } from '../lib/setup/links';
import activeBuild from '../data/active-build.json';
import companion from '../data/companion.json';
import { ADDON_DOWNLOAD_PATH, DATA_DOWNLOAD_PATH, companionDownloadPath } from '../lib/downloads/paths';
import SetupPageRoute from './setup.astro';

const SETUP_DIR = new URL('../components/setup/', import.meta.url);
/** The page's own source: the route and every component of its island. */
const source = [
  readFileSync(new URL('./setup.astro', import.meta.url), 'utf8'),
  ...readdirSync(SETUP_DIR)
    .filter((name) => name.endsWith('.svelte'))
    .map((name) => readFileSync(new URL(name, SETUP_DIR), 'utf8')),
].join('\n');

const STRING_COPY_KEYS = [
  'pageDescription',
  'installCurseForge',
  'installWago',
  'installZip',
  'installDataRelease',
  'flowOutTitle',
  'flowOutBody',
  'flowInTitle',
  'flowInBody',
  'setupAddonTitle',
  'setupCompanionTitle',
  'setupCompanionNote',
  'installDataHint',
  'setupTryExport',
  'setupTryExportBody',
  'setupCompanionBody',
  'setupPageTitle',
  'setupHowItWorks',
  'setupCompanionPair',
  'setupHandInstall',
  'setupCompanionMobile',
  'surfacesShow',
] as const;

let html: string;

beforeAll(async () => {
  const container = await createPageContainer();
  html = await container.renderToString(SetupPageRoute);
});

/** Copy shorter than this is too likely to occur in a comment to pin as absent from the page source. */
const MIN_LITERAL_PINNED = 14;

/** The markup of one Section, from its testid to the next Section. */
function section(testid: string): string {
  const at = html.indexOf(`data-testid="${testid}"`);
  expect(at, testid).toBeGreaterThan(-1);
  const next = html.indexOf('<section', at);
  return html.slice(html.lastIndexOf('<section', at), next === -1 ? undefined : next);
}

describe('/setup', () => {
  it('takes every visible string from the copy file, never a literal', () => {
    for (const key of STRING_COPY_KEYS) {
      expect(source).toContain(`addonCopy.${key}`);
      if (addonCopy[key].length > MIN_LITERAL_PINNED) expect(source).not.toContain(addonCopy[key]);
    }
    expect(source).toContain('addonCopy.currentDataBuild(');
    expect(source).toContain('addonCopy.setupHandInstallBody(');
  });

  it('is one island over the Base layout, the server markup the signed-out page', () => {
    expect(html.match(/<astro-island/g)?.length ?? 0).toBeGreaterThan(0);
    expect(html).toContain('data-testid="setup-page"');
    expect(html).toContain('data-region="head"');
    expect(html).toContain(`>${addonCopy.setupPageTitle}</h1>`);
    // No eyebrow, no progress chips, no step cards, no character bar.
    expect(html).not.toMatch(/setup-chip|setup-step-tag|setup-btn|data-setup-first|current-character-bar/);
  });

  it('holds the three steps in one framed Plate named Your setup, in order', () => {
    expect(html.match(/aria-label="Your setup"/g)).toHaveLength(1);
    expect(html).toMatch(/class="[^"]*ui-framed[^"]*"[^>]*aria-label="Your setup"/);
    const plate = html.slice(html.indexOf('aria-label="Your setup"'));
    let cursor = -1;
    for (const title of [
      addonCopy.setupSignInTitle,
      addonCopy.setupCompanionTitle,
      addonCopy.setupAddonTitle,
    ]) {
      const at = plate.indexOf(title, cursor + 1);
      expect(at, title).toBeGreaterThan(cursor);
      cursor = at;
    }
    expect(plate).toContain('aria-label="Step 1: Next"');
  });

  it('tells the visitor the in-game setting by its exact path, in The companion', () => {
    expect(section('setup-companion')).toContain(
      'Turn on Options > System > Network > Advanced Combat Logging once.',
    );
  });

  it('links every download through the site’s own /dl/ paths, never the release host', () => {
    expect(source).not.toMatch(/github\.com/i);
    // The server offers every build; the visitor's own is picked after hydration, through the same paths.
    for (const build of COMPANION_DOWNLOADS)
      expect(html).toContain(`href="${companionDownloadPath(build.asset)}"`);
    for (const entry of companion.downloads)
      expect(companionDownloadPath(entry.asset)).toBe(`/dl/companion/${entry.asset}`);
    expect(html).not.toMatch(/github\.com/i);
  });

  it('offers the companion at the top and right under The companion heading, never gold (owner 2026-11-08)', () => {
    // The download is the head's actions, beside its answer (ui/PageHeader).
    const head = html.slice(html.indexOf('data-region="head"'), html.indexOf('data-region="plate"'));
    const companionSection = section('setup-companion');
    for (const markup of [head, companionSection]) {
      for (const build of COMPANION_DOWNLOADS)
        expect(markup).toContain(`href="${companionDownloadPath(build.asset)}"`);
    }
    expect(companionSection.indexOf('setup-companion-download')).toBeLessThan(
      companionSection.indexOf('setup-companion-lead'),
    );
    expect(head).not.toContain('#f0cc6c');
  });

  it('is reachable as "Get set up" from the account menu and the footer', () => {
    expect(SETUP_NAV_ITEM.href).toBe('/setup');
  });

  it('links all three install routes in the rendered output', () => {
    expect(html).toContain(CURSEFORGE_URL);
    expect(html).toContain(WAGO_URL);
    expect(html).toContain(ADDON_DOWNLOAD_PATH);
  });

  it('keeps the hand install and the build number in The addon, CurseForge and Wago on the plate', () => {
    const addon = section('setup-addon');
    expect(addon).toContain('id="addon"');
    // The plate's addon step carries both stores while it is to do; the Section repeats them once ticked.
    const plate = html.slice(
      html.indexOf('aria-label="Your setup"'),
      html.indexOf('data-testid="setup-companion"'),
    );
    for (const url of [CURSEFORGE_URL, WAGO_URL]) expect(plate).toContain(url);
    expect(addon).not.toContain(WAGO_URL);
    // The companion's Section comes before the addon's, in the plate's order.
    expect(html.indexOf('data-testid="setup-companion"')).toBeLessThan(
      html.indexOf('data-testid="setup-addon"'),
    );
    expect(addon.indexOf(addonCopy.installDataHint)).toBeLessThan(addon.indexOf(addonCopy.setupHandInstall));
    const hand = addon.slice(addon.indexOf('data-testid="setup-hand-install"'));
    for (const url of [ADDON_DOWNLOAD_PATH, DATA_DOWNLOAD_PATH]) expect(hand).toContain(url);
    expect(hand).toContain(addonCopy.currentDataBuild(activeBuild.build));
    // Server-rendered as an unknown system: the Windows path, which most players use.
    expect(hand).toContain(addonCopy.setupHandInstallBody(false));
    expect(hand).toContain('World of Warcraft\\_forever_\\Interface\\AddOns');
    // A Mac reads the same folders with slashes.
    expect(addonCopy.setupHandInstallBody(true)).toContain('World of Warcraft/_forever_/Interface/AddOns');
    // Signed out, the addon step is not ticked: CurseForge is the plate row's action, not repeated here (U4).
    expect(addon).not.toContain(CURSEFORGE_URL);
  });

  it('addresses CurseForge by the project ids the TOCs carry, never a guessed slug', () => {
    expect(CURSEFORGE_URL).toBe('https://www.curseforge.com/projects/1706689');
  });

  it('every install link opts out of opener access', () => {
    const urls = [CURSEFORGE_URL, WAGO_URL, ADDON_DOWNLOAD_PATH, DATA_DOWNLOAD_PATH];
    for (const url of [...urls, companionDownloadPath(companion.downloads[0].asset)]) {
      const anchors = [...html.matchAll(/<a\s+[^>]*>/g)]
        .map((m) => m[0])
        .filter((tag) => tag.includes(`"${url}"`));
      expect(anchors.length, `no <a> found for ${url}`).toBeGreaterThan(0);
      for (const anchor of anchors) expect(anchor).toContain('rel="noopener"');
    }
  });

  it('keeps the paste box open in the rail at the #paste anchor', () => {
    const paste = section('setup-try-export');
    expect(paste).toContain('id="paste"');
    expect(paste).toContain(addonCopy.setupTryExportBody);
    expect(paste).toContain('data-testid="addon-paste-box"');
    expect(html).not.toContain('<details class="ui-disclosure');
  });
});

const SPEC_PATH = new URL('../../../docs/superpowers/specs/2026-10-11-addon-v2-design.md', import.meta.url);
const PUBLIC_DIR = new URL('../../public/', import.meta.url);
const MAX_IMAGE_BYTES = 200 * 1024;

/** The addon's own `settingsCommands` block, "/fs x: what it does" per line, read from the
 *  addon's Locale.lua so this page cannot drift from it again. FS_LOCALE_LUA overrides the
 *  path; by default the main checkout's working tree beside this one, else this repo's copy.
 */
function localeSettingsCommands(): string[] {
  const candidates = [
    process.env.FS_LOCALE_LUA,
    fileURLToPath(new URL('../../../../foreversixty/addon/ForeverSixty/Locale.lua', import.meta.url)),
    fileURLToPath(new URL('../../../addon/ForeverSixty/Locale.lua', import.meta.url)),
  ].filter((p): p is string => p !== undefined && existsSync(p));
  const lua = readFileSync(candidates[0], 'utf8');
  const block = /settingsCommands\s*=\s*\{([\s\S]*?)\n\t\},/.exec(lua)?.[1] ?? '';
  const lines = [...block.matchAll(/"([^"]+)"/g)].map((m) => m[1]);
  expect(lines.length, 'settingsCommands block not found').toBeGreaterThan(9);
  expect(lines, 'the file lists /fs loot').toContain('/fs loot: what an item is worth to you');
  return lines;
}

/** The `### 4.N Title` headings of the spec's §4, in file order. */
function specSurfaceHeadings(): { section: string; title: string }[] {
  const spec = readFileSync(SPEC_PATH, 'utf8');
  return [...spec.matchAll(/^### (4\.\d+) (.+)$/gm)].map((m) => ({ section: m[1], title: m[2].trim() }));
}

/** A JPEG's pixel size, read from its first start-of-frame marker. */
function jpegSize(bytes: Buffer): { width: number; height: number } {
  expect(bytes[0] === 0xff && bytes[1] === 0xd8, 'not a JPEG').toBe(true);
  let at = 2;
  while (at < bytes.length) {
    const marker = bytes[at + 1];
    const length = bytes.readUInt16BE(at + 2);
    if (marker >= 0xc0 && marker <= 0xcf && marker !== 0xc4 && marker !== 0xc8 && marker !== 0xcc) {
      return { height: bytes.readUInt16BE(at + 5), width: bytes.readUInt16BE(at + 7) };
    }
    at += 2 + length;
  }
  throw new Error('no start-of-frame marker');
}

function textOf(fragment: string): string {
  return fragment
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&#39;/g, "'")
    .replace(/&quot;/g, '"')
    .replace(/&amp;/g, '&')
    .replace(/\s+/g, ' ')
    .trim();
}

describe('/setup: what it does in the game (L8.3)', () => {
  it('lists the eleven §4 surfaces in the spec’s order, each title the spec’s own heading', () => {
    const headings = specSurfaceHeadings();
    expect(headings).toHaveLength(11);
    expect(addonCopy.surfaces.map((s) => s.section)).toEqual(headings.map((h) => h.section));
    // The spec's own word for 4.1 is "The chip"; a player is never taught it (W.10).
    expect(addonCopy.surfaces.map((s) => s.title)).toEqual(
      headings.map((h) => (h.title === 'The chip' ? 'The Forever Sixty line' : h.title)),
    );
  });

  it('renders each surface’s title and one-sentence description, in order, from the copy file', () => {
    expect(source).toContain('addonCopy.surfaces');
    let cursor = html.indexOf('data-testid="setup-surfaces"');
    expect(cursor).toBeGreaterThan(-1);
    for (const surface of addonCopy.surfaces) {
      expect(source).not.toContain(surface.body);
      const at = html.indexOf(`data-testid="setup-surface-${surface.section}"`, cursor);
      expect(at, `${surface.section} not after the previous surface`).toBeGreaterThan(cursor);
      cursor = at;
      expect(surface.body.trim().endsWith('.')).toBe(true);
      expect(surface.body.slice(0, -1)).not.toMatch(/[.!?]\s/);
    }
    const text = textOf(section('setup-surfaces'));
    for (const surface of addonCopy.surfaces) {
      expect(text).toContain(textOf(surface.title));
      expect(text).toContain(textOf(surface.body));
    }
  });

  it('folds on a phone behind one summary row, open from 768 (U5)', () => {
    const surfaces = section('setup-surfaces');
    expect(surfaces).toMatch(/<button[^>]*setup-surf-fold[^>]*aria-expanded="false"[^>]*>/);
    expect(surfaces).toContain(addonCopy.surfacesShow);
  });

  it('sits in main after the plate, before the rail', () => {
    const plate = html.indexOf('data-region="plate"');
    const surfaces = html.indexOf('data-testid="setup-surfaces"');
    const rail = html.indexOf('<aside');
    expect(surfaces).toBeGreaterThan(plate);
    expect(rail).toBeGreaterThan(surfaces);
    expect(html.indexOf('id="paste"')).toBeGreaterThan(rail);
    expect(html.indexOf(addonCopy.setupHowItWorks)).toBeGreaterThan(rail);
  });

  it('every capture exists under public/addon/, is a JPEG of at most 200 KB and 1600 px wide, at its declared size', () => {
    const images = addonCopy.surfaces.flatMap((s) => s.images);
    expect(images.length).toBeGreaterThanOrEqual(9);
    for (const image of images) {
      expect(image.src).toMatch(/^\/addon\/[a-z0-9-]+\.jpg$/);
      const bytes = readFileSync(new URL(`.${image.src}`, PUBLIC_DIR));
      expect(bytes.length, `${image.src} is ${bytes.length} bytes`).toBeLessThanOrEqual(MAX_IMAGE_BYTES);
      const size = jpegSize(bytes);
      expect(size, image.src).toEqual({ width: image.width, height: image.height });
      expect(size.width).toBeLessThanOrEqual(1600);
    }
  });

  it('every capture is a lazy Capture: alt text, its size, a 240 px cap and a link to its file', () => {
    for (const image of addonCopy.surfaces.flatMap((s) => s.images)) {
      expect(image.alt.trim().length, image.src).toBeGreaterThan(20);
      const at = html.indexOf(`src="${image.src}"`);
      expect(at, `no <img> for ${image.src}`).toBeGreaterThan(-1);
      const tag = /<img\s[^>]*>/.exec(html.slice(html.lastIndexOf('<img', at)))?.[0] ?? '';
      expect(tag).toContain('loading="lazy"');
      expect(tag).toContain(`width="${image.width}"`);
      expect(tag).toContain(`height="${image.height}"`);
      expect(tag).toMatch(/alt="[^"]+"/);
      const frame = html.slice(html.lastIndexOf('<a ', at), at);
      expect(frame, image.src).toContain(`href="${image.src}"`);
      expect(frame).toContain('max-height: 240px');
      expect(frame).toContain(`aspect-ratio: ${image.width} / ${image.height}`);
      expect(frame).not.toContain('target=');
    }
  });

  it('shows no image and no placeholder for messaging (4.9), which has no capture yet', () => {
    const messaging = addonCopy.surfaces.find((s) => s.section === '4.9');
    expect(messaging?.images).toEqual([]);
    const block = html.slice(
      html.indexOf('data-testid="setup-surface-4.9"'),
      html.indexOf('data-testid="setup-surface-4.10"'),
    );
    expect(block).not.toContain('<img');
    expect(textOf(block)).not.toMatch(/screenshot|capture|coming soon/i);
  });

  it('lists exactly the addon’s own settingsCommands, in its order, from the copy file', () => {
    expect(source).toContain('addonCopy.commands');
    const expected = localeSettingsCommands();
    expect(addonCopy.commands.map((c) => `${c.command}: ${c.description}`)).toEqual(expected);
    // Fourteen since the guild redesign added /fs job (raid assignments) after /fs ready;
    // fifteen since /fs tour (replay the tour) joined after /fs chip.
    expect(addonCopy.commands).toHaveLength(15);
    expect(addonCopy.commands.map((c) => c.command).slice(8, 12)).toEqual([
      '/fs ready',
      '/fs job',
      '/fs loot',
      '/fs floor N',
    ]);
  });

  it('draws the commands as a key-binding list: the description, then the command as Code', () => {
    const list = section('setup-commands');
    expect(list).toContain('<dl');
    let cursor = 0;
    for (const { command, description } of addonCopy.commands) {
      expect(source).not.toContain(description);
      const dt = list.indexOf(`>${description}</dt>`, cursor);
      expect(dt, description).toBeGreaterThan(-1);
      const code = list.indexOf(`>${command}</code>`, dt);
      expect(code, command).toBeGreaterThan(dt);
      cursor = code;
    }
  });

  it('keeps the testers’ commands off the public page', () => {
    expect(html).not.toMatch(/For testers/i);
    expect(html).not.toContain('/fs diag verbose');
    expect(html).not.toContain('/fs inbox');
    expect(Object.keys(addonCopy)).not.toContain('commandsTestersLine');
  });

  it('says what leaves the machine', () => {
    const text = addonCopy.handshake.join(' ');
    for (const part of [
      'level, spec, talents, equipped gear',
      'bags when the bag switch is on',
      'any sim request',
      'live roster',
      'Share my build with my guild',
      'seen HH:MM',
      'Site has your save from',
    ]) {
      expect(text).toContain(part);
    }
  });

  it('marks the four portrait captures tall', () => {
    const images = addonCopy.surfaces.flatMap((s) => s.images);
    expect(
      images
        .filter((i) => i.tall)
        .map((i) => i.src)
        .sort(),
    ).toEqual([
      '/addon/build-dialog-paste.jpg',
      '/addon/character-panel-tab.jpg',
      '/addon/guild-tab.jpg',
      '/addon/settings-switches.jpg',
    ]);
  });

  it('crops the settings capture to the switch list, the Commands list out', () => {
    const settings = addonCopy.surfaces
      .flatMap((s) => s.images)
      .find((i) => i.src === '/addon/settings-switches.jpg');
    expect(settings?.height).toBeLessThanOrEqual(800);
    expect(settings?.alt).not.toMatch(/Commands list/);
  });

  it('shows the hunter’s own captures, never a state that is not in the picture', () => {
    const alt = (src: string) =>
      addonCopy.surfaces.flatMap((s) => s.images).find((i) => i.src === src)?.alt ?? '';
    expect(alt('/addon/bag-arrow.jpg')).toContain(
      'small green arrow on the top-right corner of an item in the second row',
    );
    expect(alt('/addon/bag-arrow.jpg')).not.toMatch(/gloves/i);
    expect(alt('/addon/guild-tab.jpg')).not.toMatch(/too old|3 of 97/);
    expect(alt('/addon/talent-marks-trees.jpg')).toMatch(/hunter/i);
    expect(alt('/addon/build-dialog-paste.jpg')).toMatch(/hunter/i);
    expect(alt('/addon/settings-switches.jpg')).toMatch(/hunter|eight switches/i);
  });

  it('states the companion handshake in three sentences, naming the “Site has your save from” line', () => {
    expect(addonCopy.handshake).toHaveLength(3);
    expect(addonCopy.handshake.join(' ')).toContain('Site has your save from');
    const block = textOf(section('setup-handshake'));
    for (const sentence of addonCopy.handshake) {
      expect(sentence.trim().endsWith('.')).toBe(true);
      expect(block).toContain(textOf(sentence));
      expect(source).not.toContain(sentence);
    }
  });

  it('takes the sections’ headings from the copy file', () => {
    for (const key of ['surfacesTitle', 'commandsTitle', 'handshakeTitle'] as const) {
      expect(source).toContain(`addonCopy.${key}`);
      expect(source).not.toContain(`>${addonCopy[key]}<`);
      expect(html).toContain(addonCopy[key]);
    }
  });

  it('keeps two flows in How the addon works (U9)', () => {
    const flows = section('setup-flows');
    expect(flows).toContain(addonCopy.flowOutTitle);
    expect(flows).toContain(addonCopy.flowInTitle);
    expect(flows.match(/class="ui-row/g)).toHaveLength(2);
  });
});

describe('/setup: The companion', () => {
  it('says the companion connects itself, never that it pairs', () => {
    expect(addonCopy.setupCompanionPair).toBe(
      'Sign in with Battle.net inside the companion; it connects itself.',
    );
    expect(section('setup-companion')).toContain(addonCopy.setupCompanionPair);
    expect(html).not.toMatch(/pairs itself|pair the companion/i);
  });

  it('ranks the three how-to rows, /combatlog as Code', () => {
    const companionHtml = section('setup-companion');
    const howTo = companionHtml.slice(companionHtml.indexOf('data-testid="setup-companion-howto"'));
    for (const rank of ['1', '2', '3']) expect(howTo).toMatch(new RegExp(`>${rank}</span>`));
    expect(howTo).toMatch(/<code class="ui-code[^"]*"[^>]*>\/combatlog<\/code>/);
  });

  it('shows the 0.2.44 companion Status page as an eager Capture with its Example tag', () => {
    const capture = addonCopy.setupCompanionCapture;
    const companionHtml = section('setup-companion');
    expect(companionHtml).toContain(`src="${capture.src}"`);
    const img =
      [...companionHtml.matchAll(/<img[^>]*>/g)].map((m) => m[0]).find((tag) => tag.includes(capture.src)) ??
      '';
    expect(img).toContain('loading="eager"');
    expect(companionHtml).toContain(`>${capture.exampleLabel}<`);
    expect(capture.alt).toContain('Bow Jackzon');
    const bytes = readFileSync(new URL(`../../public${capture.src}`, import.meta.url));
    expect(jpegSize(bytes)).toEqual({ width: capture.width, height: capture.height });
    expect(capture.width).toBeGreaterThanOrEqual(1600);
    expect(companionHtml).toContain(`srcset="${capture.narrow.src}"`);
    const narrowBytes = readFileSync(new URL(`../../public${capture.narrow.src}`, import.meta.url));
    expect(jpegSize(narrowBytes)).toEqual({ width: capture.narrow.width, height: capture.narrow.height });
  });

  it('server-renders every build and its signing line, with no word of Windows only (U7)', () => {
    const companionHtml = section('setup-companion');
    for (const build of COMPANION_DOWNLOADS) {
      expect(html).toContain(`href="${companionDownloadPath(build.asset)}"`);
      expect(companionHtml).toContain(`${build.platform}: ${build.firstRun}`);
    }
    expect(html).not.toMatch(/Windows[ -]only|not supported yet/i);
  });

  it('draws no pairing block signed out', () => {
    expect(html).not.toContain('data-testid="setup-pairing"');
  });

  it('says what the companion sends, never calling the step the plate points at optional', () => {
    expect(addonCopy.setupCompanionNote).toBe('Sends bags, bank and live logs after every session.');
    expect(section('setup-companion')).toContain(addonCopy.setupCompanionNote);
  });
});
