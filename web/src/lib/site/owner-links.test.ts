// web/src/lib/site/owner-links.test.ts
// Nothing the site renders may link to the owner's code host account: the repository goes
// private after launch and must not lead back to its owner. The built site is checked by
// scripts/check-no-owner-links.mjs in postbuild; this drives the same check over a fixture and
// over every source the pages render from, so a stray citation fails before any build.
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { OWNER_LINK, ownerLinkFailures } from '../../../scripts/owner-links.mjs';

const WEB = path.resolve(import.meta.dirname, '../../..');
const REPO = path.resolve(WEB, '..');

/** What the pages render from: the site's data and content, and the curated data it syncs. */
const RENDERED_SOURCES = [
  path.join(WEB, 'src/data'),
  path.join(WEB, 'src/content'),
  path.join(WEB, 'public'),
  path.join(REPO, 'data/curated'),
];

let root: string;

beforeEach(() => {
  root = mkdtempSync(path.join(tmpdir(), 'owner-links-'));
});

afterEach(() => {
  rmSync(root, { recursive: true, force: true });
});

describe('ownerLinkFailures', () => {
  it('names each text file that links to the owner, with its line', () => {
    mkdirSync(path.join(root, 'guides'));
    writeFileSync(
      path.join(root, 'guides', 'fire.html'),
      `<p>ok</p>\n<a href="https://${OWNER_LINK}/x">src</a>\n`,
    );
    writeFileSync(path.join(root, 'data.json'), `{"url":"https://GitHub.com/JHunthrop/x"}`);
    writeFileSync(
      path.join(root, 'clean.html'),
      '<a href="https://worldofwarcraft.blizzard.com/">Blizzard</a>',
    );
    expect(ownerLinkFailures([root], root).sort()).toEqual([
      'data.json:1',
      path.join('guides', 'fire.html:2'),
    ]);
  });

  it('skips binary assets', () => {
    writeFileSync(path.join(root, 'sim.wasm'), OWNER_LINK);
    expect(ownerLinkFailures([root], root)).toEqual([]);
  });
});

describe('the sources the site renders', () => {
  it(`contain no ${OWNER_LINK} link`, () => {
    expect(ownerLinkFailures(RENDERED_SOURCES, REPO)).toEqual([]);
  });
});
