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
import activeBuild from '../data/active-build.json';
import SetupPage, { CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL } from './setup.astro';

const source = readFileSync(new URL('./setup.astro', import.meta.url), 'utf8');

const STRING_COPY_KEYS = [
  'pageDescription',
  'pageNoNetwork',
  'installCurseForge',
  'installWago',
  'installGitHub',
  'flowOutTitle',
  'flowOutBody',
  'flowInTitle',
  'flowInBody',
  'setupStep1Title',
  'setupStep2Title',
  'setupStep3Title',
  'setupSignInBody',
  'setupSignInRealmNote',
  'setupSignInAction',
  'setupCompanionBody',
  'setupPageTitle',
  'setupStatusSignIn',
  'setupStatusSignedIn',
  'setupStatusAddon',
  'setupStatusAddonDone',
  'setupStatusCompanion',
  'setupSignedInLine',
  'setupSignedInLink',
  'setupAddonDoneLine',
  'setupHowItWorks',
  'setupCompanionDownloadLead',
  'setupCompanionPair',
  'disclosureOpen',
  'disclosureClose',
] as const;

let html: string;

beforeAll(async () => {
  const container = await createPageContainer();
  html = await container.renderToString(SetupPage);
});

describe('/setup', () => {
  it('takes every visible string from the copy file, never a literal', () => {
    for (const key of STRING_COPY_KEYS) {
      expect(source).toContain(`addonCopy.${key}`);
      expect(source).not.toContain(addonCopy[key]);
    }
    expect(source).toContain('addonCopy.currentDataBuild(');
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

  it('shows the current data build from the data file in the rendered output', () => {
    expect(html).toContain(addonCopy.currentDataBuild(activeBuild.build));
  });

  it('names which realm types Blizzard serves today under step 1', () => {
    expect(html).toContain(addonCopy.setupSignInRealmNote);
  });

  it('every install link opts out of opener access', () => {
    for (const url of [CURSEFORGE_URL, WAGO_URL, GITHUB_RELEASES_URL]) {
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
    expect(addonCopy.surfaces.map((s) => s.title)).toEqual(headings.map((h) => h.title));
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

  it('puts the testers’ commands on one muted line under a small “For testers” label', () => {
    const block = html.slice(
      html.indexOf('data-testid="setup-commands"'),
      html.indexOf('data-testid="setup-handshake"'),
    );
    expect(addonCopy.commandsTestersLabel).toBe('For testers');
    expect(addonCopy.commandsTestersLine).toBe(
      '/fs diag verbose, /fs inbox, /fs hello, /fs seen, /fs record',
    );
    expect(source).toContain('addonCopy.commandsTestersLine');
    expect(block.match(/For testers/g)).toHaveLength(1);
    expect(block).toMatch(
      /<p class="[^"]*text-muted[^"]*">\/fs diag verbose, \/fs inbox, \/fs hello, \/fs seen, \/fs record<\/p>/,
    );
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
      'Seen by the site',
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

  it('states the companion handshake in three sentences, naming the “Seen by the site” line', () => {
    expect(addonCopy.handshake).toHaveLength(3);
    expect(addonCopy.handshake.join(' ')).toContain('Seen by the site');
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
