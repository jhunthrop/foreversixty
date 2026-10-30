// web/src/pages/_setup.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts and _classes.test.ts:
// everything else under src/pages/ is a route, so an unprefixed setup.test.ts would build as
// the route /setup.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// Adapted from _addon.test.ts (Task 6, spec 2026-09-25 §3.4): /setup replaces /addon with
// three numbered panels. The GitHub release link is derived from links.json's githubRepo, so
// it is asserted off the rendered output rather than as a literal substring of the source,
// same as the page it replaces.
import { existsSync, readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { createPageContainer } from '../test-support/page-container';
import { describe, expect, it, beforeAll } from 'vitest';
import { addonCopy } from '../lib/addon/copy';
import { SETUP_NAV_ITEM } from '../lib/nav';
import type { Device, Me, MeCharacter } from '../lib/account/api';
import {
  addonLine,
  chipsDone,
  companionLine,
  companionPlatformsLine,
  newestAddonCharacter,
  signInLine,
} from '../lib/setup/progress';
import activeBuild from '../data/active-build.json';
import SetupPage, {
  CURSEFORGE_URL,
  DATA_GITHUB_RELEASES_URL,
  GITHUB_RELEASES_URL,
  WAGO_URL,
} from './setup.astro';

const source = readFileSync(new URL('./setup.astro', import.meta.url), 'utf8');

const STRING_COPY_KEYS = [
  'pageDescription',
  'pageNoNetwork',
  'installCurseForge',
  'installWago',
  'installGitHub',
  'installDataLead',
  'installDataGitHub',
  'flowOutTitle',
  'flowOutBody',
  'flowInTitle',
  'flowInBody',
  'setupStep1Title',
  'setupStep2Title',
  'setupStep3Title',
  'setupSignInBody',
  'setupSignInAction',
  'setupSubline',
  'setupHowItArrives',
  'setupCompanionNote',
  'installDataHint',
  'setupManualInstall',
  'setupTryExport',
  'setupCompanionBody',
  'setupPageTitle',
  'setupEverything',
  'setupAddonDoneLine',
  'setupHowItWorks',
  'setupCompanionPair',
] as const;

let html: string;

beforeAll(async () => {
  const container = await createPageContainer();
  html = await container.renderToString(SetupPage);
});

/** Copy shorter than this is too likely to occur in a comment to pin as absent from the page source. */
const MIN_LITERAL_PINNED = 14;

describe('/setup', () => {
  it('takes every visible string from the copy file, never a literal', () => {
    for (const key of STRING_COPY_KEYS) {
      expect(source).toContain(`addonCopy.${key}`);
      // A one or two word title ("Sign in") also appears in comments; only a sentence is pinned as absent.
      if (addonCopy[key].length > MIN_LITERAL_PINNED) expect(source).not.toContain(addonCopy[key]);
    }
    expect(source).toContain('addonCopy.currentDataBuild(');
  });

  it('tells the visitor the in-game setting by its exact path, in step 3', () => {
    expect(html).toContain(
      'Turn on Options &gt; Network &gt; Advanced Combat Logging once, then restart the game.',
    );
  });

  it('derives the GitHub release link from links.json, not a second copy of the repo URL', () => {
    expect(source).toContain('communityLinks.githubRepo');
    expect(source).not.toContain('jhunthrop/forever/releases');
  });

  it('is reachable from the primary nav as "Get set up"', () => {
    expect(SETUP_NAV_ITEM.href).toBe('/setup');
  });

  it('links all three install routes in the rendered output', () => {
    expect(html).toContain(CURSEFORGE_URL);
    expect(html).toContain(WAGO_URL);
    expect(html).toContain(GITHUB_RELEASES_URL);
  });

  it('keeps the data addon’s own links and the build number in the Manual install disclosure, not in the step', () => {
    expect(html).toContain('data-testid="setup-manual-install"');
    expect(html).toContain(DATA_GITHUB_RELEASES_URL);
    const open = html.indexOf('data-testid="setup-manual-install"');
    const close = html.indexOf('</details>', open);
    const inside = html.slice(open, close);
    for (const url of [GITHUB_RELEASES_URL, DATA_GITHUB_RELEASES_URL]) {
      expect(inside).toContain(url);
    }
    expect(inside).toContain(addonCopy.currentDataBuild(activeBuild.build));
    // Only the zip and the data release: the store buttons live in the step, not repeated here.
    expect(inside).not.toContain(CURSEFORGE_URL);
    expect(inside).not.toContain(WAGO_URL);
    expect(inside.match(/<a\s/g)).toHaveLength(2);
    expect(html.indexOf(addonCopy.installDataHint)).toBeLessThan(open);
  });

  it('addresses CurseForge by the project ids the TOCs carry, never a guessed slug', () => {
    expect(CURSEFORGE_URL).toBe('https://www.curseforge.com/projects/1706689');
  });

  it('every install link opts out of opener access', () => {
    for (const url of [CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL, DATA_GITHUB_RELEASES_URL]) {
      const anchor = [...html.matchAll(/<a\s+[^>]*>/g)].map((m) => m[0]).find((tag) => tag.includes(url));
      expect(anchor, `no <a> found for ${url}`).toBeDefined();
      expect(anchor).toContain('rel="noopener"');
    }
  });

  it('keeps the paste box at the #paste anchor', () => {
    expect(html).toContain('id="paste"');
  });

  it('renders the three numbered steps in order', () => {
    const order = [addonCopy.setupStep1Title, addonCopy.setupStep2Title, addonCopy.setupStep3Title];
    let cursor = -1;
    for (const title of order) {
      const at = html.indexOf(title, cursor === -1 ? 0 : cursor);
      expect(at, `${title} not found after the previous step`).toBeGreaterThan(cursor);
      cursor = at;
    }
  });
});

// ---------------------------------------------------------------------------------------
// Task L8.3: "What it does in the game" -- one row per surface in the addon-v2 spec's §4,
// in the spec's order, each with a capture from the client where one exists; the /fs
// commands the addon's /fs help prints; the companion handshake.
// ---------------------------------------------------------------------------------------

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
    // The spec's own word for 4.1 is "The chip"; a player is never taught it (W.10): the page says
    // "The Forever Sixty line", the name the addon's own settings use.
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
      const at = html.indexOf(`data-surface="${surface.section}"`, cursor);
      expect(at, `${surface.section} not after the previous surface`).toBeGreaterThan(cursor);
      cursor = at;
      // One sentence: a single full stop, at the end.
      expect(surface.body.trim().endsWith('.')).toBe(true);
      expect(surface.body.slice(0, -1)).not.toMatch(/[.!?]\s/);
    }
    const text = textOf(html.slice(html.indexOf('data-testid="setup-surfaces"')));
    for (const surface of addonCopy.surfaces) {
      expect(text).toContain(textOf(surface.title));
      expect(text).toContain(textOf(surface.body));
    }
  });

  it('sits between the paste box and the how-it-works flows', () => {
    const paste = html.indexOf('id="paste"');
    const surfaces = html.indexOf('data-testid="setup-surfaces"');
    const flows = html.indexOf(addonCopy.setupHowItWorks);
    expect(paste).toBeGreaterThan(-1);
    expect(surfaces).toBeGreaterThan(paste);
    expect(flows).toBeGreaterThan(surfaces);
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

  it('every capture has non-empty alt text, is lazy-loaded and sits in a fixed aspect box', () => {
    for (const image of addonCopy.surfaces.flatMap((s) => s.images)) {
      expect(image.alt.trim().length, image.src).toBeGreaterThan(20);
      const tag = [...html.matchAll(/<img\s[^>]*>/g)]
        .map((m) => m[0])
        .find((t) => t.includes(`src="${image.src}"`));
      expect(tag, `no <img> for ${image.src}`).toBeDefined();
      expect(tag).toContain('loading="lazy"');
      expect(tag).toContain(`width="${image.width}"`);
      expect(tag).toContain(`height="${image.height}"`);
      expect(tag).toMatch(/alt="[^"]+"/);
    }
    for (const surface of addonCopy.surfaces) {
      for (const image of surface.images) {
        expect(html).toContain(`aspect-ratio: ${image.width} / ${image.height}`);
      }
    }
  });

  it('shows no image and no placeholder for messaging (4.9), which has no capture yet', () => {
    const messaging = addonCopy.surfaces.find((s) => s.section === '4.9');
    expect(messaging?.images).toEqual([]);
    const block = html.slice(html.indexOf('data-surface="4.9"'), html.indexOf('data-surface="4.10"'));
    expect(block).not.toContain('<img');
    expect(textOf(block)).not.toMatch(/screenshot|capture|coming soon/i);
  });

  it('lists exactly the addon’s own settingsCommands, in its order, from the copy file', () => {
    expect(source).toContain('addonCopy.commands');
    const expected = localeSettingsCommands();
    expect(addonCopy.commands.map((c) => `${c.command}: ${c.description}`)).toEqual(expected);
    expect(addonCopy.commands).toHaveLength(13);
    expect(addonCopy.commands.map((c) => c.command).slice(7, 10)).toEqual([
      '/fs ready',
      '/fs loot',
      '/fs floor N',
    ]);
    // The tags between the command and its colon read as a space in textOf; a browser shows none.
    const list = textOf(html.slice(html.indexOf('data-testid="setup-commands"'))).replace(/ :/g, ':');
    for (const { command, description } of addonCopy.commands) {
      expect(description.length).toBeGreaterThan(0);
      expect(source).not.toContain(description);
      expect(list).toContain(`${command}: ${description}`);
    }
  });

  it('keeps the colon outside the bold monospace command, so the command is exactly what is typed', () => {
    const block = html.slice(html.indexOf('data-testid="setup-commands"'));
    for (const { command } of addonCopy.commands) {
      expect(block).toMatch(new RegExp(`<code class="[^"]*font-bold[^"]*">${command}</code>`));
      expect(block).not.toContain(`${command}:</code>`);
    }
  });

  it('keeps the testers’ commands off the public page', () => {
    expect(html).not.toMatch(/For testers/i);
    expect(html).not.toContain('/fs diag verbose');
    expect(html).not.toContain('/fs inbox');
    expect(Object.keys(addonCopy)).not.toContain('commandsTestersLine');
  });

  it('says what leaves the machine and sits the no-network sentence next to the handshake', () => {
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
    const block = html.slice(html.indexOf('data-testid="setup-handshake"'));
    expect(textOf(block)).toContain(textOf(addonCopy.pageNoNetwork));
    // Moved out of step 2, not repeated there.
    expect(html.split(textOf(addonCopy.pageNoNetwork).slice(0, 30)).length - 1).toBe(1);
  });

  it('gives every image box a surface colour, a per-image height cap (900 px portrait, 420 px landscape) and a link to the full-size file', () => {
    const images = addonCopy.surfaces.flatMap((s) => s.images);
    expect(
      images
        .filter((i) => i.tall)
        .map((i) => i.src)
        .sort(),
    ).toEqual([
      '/addon/build-dialog-copy.jpg',
      '/addon/build-dialog-paste.jpg',
      '/addon/character-panel-tab.jpg',
      '/addon/guild-tab.jpg',
      '/addon/settings-switches.jpg',
    ]);
    for (const image of images) {
      const at = html.indexOf(`src="${image.src}"`);
      const box = html.slice(html.lastIndexOf('<div', at), at);
      expect(box, image.src).toContain('bg-raised');
      expect(box).toContain(image.tall ? 'max-h-[900px]' : 'max-h-[420px]');
      expect(box).not.toContain(image.tall ? 'max-h-[420px]' : 'max-h-[900px]');
      expect(box).toContain(`max-width: ${image.width}px`);
      const tag = /<img\s[^>]*>/.exec(html.slice(at - 200))?.[0] ?? '';
      expect(tag).toContain('object-contain');
      // The <img> sits inside <a href="its file"> opening in the same tab.
      const link = html.slice(html.lastIndexOf('<a ', at), at);
      expect(link, image.src).toContain(`href="${image.src}"`);
      expect(link).not.toContain('target=');
      expect(html.slice(at, html.indexOf('</a>', at))).not.toContain('<a ');
    }
  });

  it('shows the hunter’s copy dialog with the export selected and the no-companion lines', () => {
    const copy = addonCopy.surfaces
      .flatMap((s) => s.images)
      .find((i) => i.src === '/addon/build-dialog-copy.jpg');
    expect(copy?.alt).toMatch(/hunter/i);
    expect(copy?.alt).toContain('No companion found.');
    expect(copy?.alt).toContain('Or paste it on the site');
    expect(copy?.alt).not.toMatch(/warrior|Deflection/i);
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
    const block = textOf(html.slice(html.indexOf('data-testid="setup-handshake"')));
    for (const sentence of addonCopy.handshake) {
      expect(sentence.trim().endsWith('.')).toBe(true);
      expect(block).toContain(textOf(sentence));
      expect(source).not.toContain(sentence);
    }
  });

  it('takes the section’s headings from the copy file', () => {
    for (const key of ['surfacesTitle', 'commandsTitle', 'handshakeTitle'] as const) {
      expect(source).toContain(`addonCopy.${key}`);
      expect(source).not.toContain(`>${addonCopy[key]}<`);
      expect(html).toContain(addonCopy[key]);
    }
  });

  it('no longer says the in-game window has no screenshot', () => {
    expect(addonCopy.inGameBody).not.toMatch(/no screenshot/i);
  });
});

describe('/setup first-run copy', () => {
  it('says the companion connects itself, never that it pairs', () => {
    expect(addonCopy.setupCompanionPair).toBe(
      'Sign in with Battle.net inside the companion; it connects itself.',
    );
    expect(html).toContain(addonCopy.setupCompanionPair);
    expect(html).not.toMatch(/pairs itself|pair the companion/i);
  });

  it('makes "Try an export" in that line a link that opens its disclosure', () => {
    const line = html.slice(html.indexOf('data-testid="setup-how-it-arrives"'));
    expect(line.slice(0, line.indexOf('</p>'))).toContain(
      `<a href="#paste" class="whitespace-nowrap" data-open-disclosure="setup-try-export">${addonCopy.setupTryExport}</a>`,
    );
    // The phrase around the link never breaks: non-breaking spaces on both sides of it.
    expect(line.slice(0, line.indexOf('</p>'))).toMatch(
      /\(see\u00A0<a [^>]*>Try an export<\/a>\u00A0below\)\./,
    );
    expect(html).toContain('data-testid="setup-try-export"');
  });

  it('says once, the way the addon welcome does, how a character reaches the site', () => {
    expect(addonCopy.setupHowItArrives).toBe(
      'The addon saves your character when you log out; the companion sends it to the site. No companion? Paste an export instead (see Try an export below).',
    );
    expect(html.split('No companion? Paste an export instead (see\u00A0').length - 1).toBe(1);
  });

  it('drops the Blizzard status sentence from step 1', () => {
    expect(html).not.toMatch(/waits on Blizzard|Character import/i);
    expect(addonCopy.setupSignInBody).toBe('Saves your builds, sims and logs to your account.');
  });
});

// W.9: one gold button per step, a done line per step, the manual behind disclosures.
describe('/setup as three steps (W.9)', () => {
  const GOLD = 'bg-gold';
  const anchorWith = (needle: string): string =>
    [...html.matchAll(/<a\s[^>]*>/g)].map((m) => m[0]).find((tag) => tag.includes(needle)) ?? '';

  it('titles the page "Get set up" with the one-line subline, not the old lede', () => {
    expect(addonCopy.setupPageTitle).toBe('Get set up');
    expect(addonCopy.setupSubline).toBe('Three steps. The first two take two minutes.');
    expect(html).toContain(addonCopy.setupSubline);
  });

  it('draws three chips that are anchors to their steps, numbered', () => {
    const start = html.indexOf('data-testid="setup-status"');
    const chips = html.slice(start, start + 4000);
    for (const [index, chip] of addonCopy.setupChips.entries()) {
      expect(chip.step).toBe(index + 1);
      expect(chips).toContain(`href="${chip.href}"`);
      expect(chips).toContain(chip.label);
      expect(html).toContain(`id="${chip.href.slice(1)}"`);
    }
    // One label per chip whatever its state, so a done row never reflows (W.10): the tick says done.
    expect(addonCopy.setupChips.map((chip) => chip.label)).toEqual(['Sign in', 'Addon', 'Companion']);
    expect(addonCopy.setupChips[0]?.labelDone).toBe('Signed in');
    expect(chips).not.toMatch(/Addon seen|Companion running/);
  });

  it('gives each step one button, gold only on the first undone step (global.css .setup-btn): Sign in with Battle.net, Install from CurseForge, Download for Windows', () => {
    for (const [needle, step] of [
      ['battlenet', '1'],
      [CURSEFORGE_URL, '2'],
      ['data-testid="setup-companion-download"', '3'],
    ] as const) {
      const tag = anchorWith(needle);
      expect(tag, needle).toContain('setup-btn');
      expect(tag, needle).toContain(`data-step="${step}"`);
      // Gold is the page's rule for the first undone step, never a fixed class on the button.
      expect(tag, needle).not.toContain(GOLD);
    }
    const css = readFileSync(new URL('../styles/global.css', import.meta.url), 'utf8');
    expect(css).toContain("html[data-setup-first='2'] .setup-btn[data-step='2']");
    expect(css).toContain("html:not([data-setup-first]):not([data-session='1']) .setup-btn[data-step='1']");
    expect(html).toContain('Download for Windows');
    // Wago and the zip are quiet links, not buttons.
    expect(anchorWith(WAGO_URL)).not.toContain(GOLD);
    expect(anchorWith(GITHUB_RELEASES_URL)).not.toContain(GOLD);
    expect(addonCopy.installGitHub).toBe('Download the zip');
  });

  it('gives each step a line component, hidden until a session exists', () => {
    for (const step of ['signin', 'addon', 'companion']) {
      const tag = new RegExp(`<p class="([^"]*)"[^>]*data-testid="setup-line-${step}"`).exec(html);
      expect(tag?.[1], step).toContain('setup-when-in');
    }
  });

  it('puts the export box, the manual install and the manual behind three disclosures, the manual install open', () => {
    expect(html.match(/<details/g)).toHaveLength(3);
    for (const testid of ['setup-try-export', 'setup-manual-install', 'setup-everything']) {
      const open = new RegExp(`<details[^>]*data-testid="${testid}"[^>]*>`).exec(html)?.[0] ?? '';
      expect(open, testid).not.toBe('');
      expect(/ open[ >=]/.test(open), testid).toBe(testid === 'setup-manual-install');
    }
    const tryExport = html.indexOf('data-testid="setup-try-export"');
    expect(html.indexOf('id="paste"')).toBeGreaterThan(tryExport);
    expect(html.indexOf('id="paste"')).toBeLessThan(html.indexOf('</details>', tryExport));
    const disclosure = html.indexOf('data-testid="setup-everything"');
    const end = html.indexOf('</details>', disclosure);
    const inside = html.slice(disclosure, end);
    expect(addonCopy.setupEverything).toBe('Everything the addon does');
    for (const marker of ['setup-surfaces', 'setup-commands', 'setup-handshake', 'setup-flows']) {
      expect(inside, marker).toContain(`data-testid="${marker}"`);
    }
    expect(html.indexOf('id="companion"')).toBeGreaterThan(end);
  });

  it('shows the 0.2.11 companion Status page, two-word name, in the third step', () => {
    const capture = addonCopy.setupCompanionCapture;
    const step3 = html.slice(html.indexOf('id="companion"'));
    expect(step3).toContain(`src="${capture.src}"`);
    expect(capture.alt).toContain('Status');
    expect(capture.alt).toContain('Bow Jackzon');
    expect(capture.alt).not.toContain('Bow-Jackzon');
    const bytes = readFileSync(new URL(`../../public${capture.src}`, import.meta.url));
    expect(jpegSize(bytes)).toEqual({ width: capture.width, height: capture.height });
    // A 2x capture: at least 1600 px of source for the 800 px it is shown at.
    expect(capture.width).toBeGreaterThanOrEqual(1600);
  });

  it('draws the capture as a picture with a phone source: the headline card alone, sized by attributes (W.10)', () => {
    const capture = addonCopy.setupCompanionCapture;
    const step3 = html.slice(html.indexOf('id="companion"'));
    expect(step3).toMatch(/<picture[^>]*>/);
    const source = /<source[^>]*>/.exec(step3)?.[0] ?? '';
    expect(source).toContain('media="(max-width: 639px)"');
    expect(source).toContain(`srcset="${capture.phoneSrc}"`);
    expect(source).toContain(`width="${capture.phoneWidth}"`);
    expect(source).toContain(`height="${capture.phoneHeight}"`);
    // Not lazy: a lazy image caught early was the blank box at 390.
    const img =
      [...step3.matchAll(/<img[^>]*>/g)].map((m) => m[0]).find((tag) => tag.includes(capture.src)) ?? '';
    expect(img).not.toBe('');
    expect(img).not.toContain('loading="lazy"');
    const bytes = readFileSync(new URL(`../../public${capture.phoneSrc}`, import.meta.url));
    expect(jpegSize(bytes)).toEqual({ width: capture.phoneWidth, height: capture.phoneHeight });
    // A crop of the headline card at 2x: about 700 x 300 or more of source for the 350 px it is shown at.
    expect(capture.phoneWidth).toBeGreaterThanOrEqual(700);
    expect(capture.phoneHeight).toBeLessThanOrEqual(320);
  });

  it('says what an export is inside its disclosure', () => {
    const tryExport = html.slice(html.indexOf('data-testid="setup-try-export"'));
    expect(tryExport.slice(0, tryExport.indexOf('</details>'))).toContain(addonCopy.setupTryExportBody);
  });

  it('says Mac is not supported while the release carries no Mac bundle', () => {
    const step3 = html.slice(html.indexOf('id="companion"'));
    expect(step3).toContain('Windows only; Mac is not supported yet.');
    expect(step3).not.toMatch(/macOS next/i);
    expect(companionPlatformsLine(['Windows', 'macOS'])).toBeNull();
    expect(companionPlatformsLine(['Windows'])).toBe('Windows only; Mac is not supported yet.');
  });

  it('says which of the steps is optional', () => {
    expect(addonCopy.setupCompanionNote).toBe(
      'Sends your character automatically and logs live. Without it, paste an export after each change.',
    );
  });
});

describe('the setup done lines (progress.ts)', () => {
  const NOW = new Date('2026-10-20T12:00:00Z');
  const user = { id: 1, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false };
  const character = (
    name: string,
    capturedAt: string | null,
    source: 'addon' | 'blizzard' = 'addon',
  ): MeCharacter => ({
    key: `us/hardcore/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'hardcore',
    name,
    ...(capturedAt === null ? {} : { build: { source, captured_at: capturedAt } }),
  });
  const me = (characters: MeCharacter[], battletag: string | null = 'hunthrop#1894'): Me => ({
    user: { ...user, battletag },
    characters,
    guilds: [],
  });
  const device = (name: string, lastSeen: string | null, created = '2026-10-01T00:00:00Z'): Device => ({
    id: name,
    name,
    platform: 'windows',
    created_at: created,
    last_seen_at: lastSeen,
  });

  it('says nothing signed out', () => {
    expect(signInLine(null)).toBeNull();
    expect(addonLine(null, NOW)).toBeNull();
    expect(companionLine(null, [], NOW)).toBeNull();
    expect(chipsDone({ me: null, devices: [] })).toEqual([false, false, false]);
  });

  it('names the signed-in account: the battletag, else the email', () => {
    expect(signInLine(me([]))).toEqual({ done: true, text: 'Signed in as hunthrop#1894' });
    const emailOnly = { ...me([], null), user: { ...user, battletag: null, email: 'a@b.gg' } };
    expect(signInLine(emailOnly)?.text).toBe('Signed in as a@b.gg');
  });

  it('names the newest character the addon synced and when, ignoring Battle.net-only builds', () => {
    const account = me([
      character('Old Hunter', '2026-10-10T00:00:00Z'),
      character('Bow Jackzon', '2026-10-20T11:00:00Z'),
      character('Blizzard Only', '2026-10-20T11:59:00Z', 'blizzard'),
      character('No Build', null),
    ]);
    expect(newestAddonCharacter(account.characters)?.character.name).toBe('Bow Jackzon');
    expect(addonLine(account, NOW)).toEqual({
      done: true,
      text: 'Bow Jackzon arrived from the addon 1 hour ago',
    });
  });

  it('says what to do when no addon character has arrived', () => {
    expect(addonLine(me([character('Blizzard Only', '2026-10-20T11:59:00Z', 'blizzard')]), NOW)).toEqual({
      done: false,
      text: 'No character yet. Log out once with the addon on and the companion running, or paste an export.',
    });
  });

  it('names the most recently seen companion device and when', () => {
    const devices = [
      device('OLD-LAPTOP', '2026-10-02T00:00:00Z'),
      device('JH-DESKTOP', '2026-10-20T11:57:00Z'),
      device('NEVER-SEEN', null, '2026-10-03T00:00:00Z'),
    ];
    expect(companionLine(me([]), devices, NOW)).toEqual({
      done: true,
      text: 'Companion signed in on JH-DESKTOP · last seen 3 minutes ago',
    });
  });

  it('falls back to the creation time for a device never heard from', () => {
    const line = companionLine(me([]), [device('JH-DESKTOP', null, '2026-10-20T11:58:00Z')], NOW);
    expect(line?.text).toBe('Companion signed in on JH-DESKTOP · last seen 2 minutes ago');
  });

  it('says no companion has signed in, and claims nothing when the device list did not load', () => {
    expect(companionLine(me([]), [], NOW)).toEqual({ done: false, text: 'No companion signed in yet' });
    expect(companionLine(me([]), null, NOW)).toBeNull();
  });

  it('marks a chip done only when the account shows its step', () => {
    const account = me([character('Bow Jackzon', '2026-10-20T11:00:00Z')]);
    expect(chipsDone({ me: account, devices: [] })).toEqual([true, true, false]);
    expect(chipsDone({ me: me([]), devices: [device('JH-DESKTOP', null)] })).toEqual([true, false, true]);
    expect(chipsDone({ me: me([]), devices: null })).toEqual([true, false, false]);
  });
});
