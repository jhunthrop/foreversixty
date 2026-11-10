// web/src/lib/sitemap-canonical.test.ts
import { mkdirSync, mkdtempSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { describe, expect, it } from 'vitest';
import { builtFileOf, canonicalOf, isSelfCanonical } from './sitemap-canonical.mjs';

const SITE = 'https://foreversixty.gg';
const page = (canonical: string | null): string =>
  `<html><head>${canonical === null ? '' : `<link rel="canonical" href="${canonical}" data-og="canonical">`}</head></html>`;

function distWith(files: Record<string, string>): string {
  const dir = mkdtempSync(join(tmpdir(), 'sitemap-canonical-'));
  for (const [path, html] of Object.entries(files)) {
    const file = join(dir, path);
    mkdirSync(join(file, '..'), { recursive: true });
    writeFileSync(file, html);
  }
  return dir;
}

describe('canonicalOf', () => {
  it('reads the href whatever the attribute order', () => {
    expect(canonicalOf('<link href="https://x/a" rel="canonical">')).toBe('https://x/a');
    expect(canonicalOf(page(`${SITE}/sim`))).toBe(`${SITE}/sim`);
  });

  it('is null without a canonical tag', () => {
    expect(canonicalOf(page(null))).toBeNull();
  });
});

describe('builtFileOf', () => {
  it('maps the build.format file layout', () => {
    expect(builtFileOf('/d', `${SITE}/`)).toBe('/d/index.html');
    expect(builtFileOf('/d', `${SITE}/bis/mage/fire/60/pre-raid`)).toBe('/d/bis/mage/fire/60/pre-raid.html');
  });
});

describe('isSelfCanonical', () => {
  const dist = distWith({
    'index.html': page(`${SITE}/`),
    'bis/mage/fire/60.html': page(`${SITE}/bis/mage/fire/60`),
    'bis/mage/fire/60/pre-raid.html': page(`${SITE}/bis/mage/fire/60`),
    'plain.html': page(null),
  });

  it('keeps a page that is its own canonical', () => {
    expect(isSelfCanonical(dist, `${SITE}/`)).toBe(true);
    expect(isSelfCanonical(dist, `${SITE}/bis/mage/fire/60`)).toBe(true);
  });

  it('drops an alias whose canonical is another page', () => {
    expect(isSelfCanonical(dist, `${SITE}/bis/mage/fire/60/pre-raid`)).toBe(false);
  });

  it('keeps a page with no canonical, and one it cannot find', () => {
    expect(isSelfCanonical(dist, `${SITE}/plain`)).toBe(true);
    expect(isSelfCanonical(dist, `${SITE}/not-built`)).toBe(true);
  });
});
