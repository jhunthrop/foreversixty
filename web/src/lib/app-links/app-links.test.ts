// The universal links and App Links files open only /app/* and the site pages links.json names
// in the native apps, and the committed copies under public/.well-known match web/app-links.json.

const SITE_PAGES = [
  '/b/*',
  '/bis/*/*',
  '/guides/*',
  '/guides/*/*',
  '/sim/*',
  '/reports/*',
  '/character/*/*/*',
  '/guild/*/*/*/raids/*',
  '/tiers',
];
const SIM_TOOLS = [
  '/sim/specs',
  '/sim/gear',
  '/sim/drops',
  '/sim/talents',
  '/sim/weights',
  '/reports/sample',
];
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import {
  AASA_FILE,
  ASSETLINKS_FILE,
  appleAppSiteAssociation,
  assetLinks,
  readAppLinksConfig,
  readSitePages,
  renderAppLinks,
  wellKnownDir,
} from '../../../scripts/app-links.mjs';

const cfg = { ...readAppLinksConfig() };

describe('app links', () => {
  it('opens only /app/* and the site pages on iOS, the sim tools excluded, for the Team ID and bundle id', () => {
    const aasa = appleAppSiteAssociation({ ...cfg, appleTeamId: 'ABCDE12345' });
    expect(aasa.applinks.details).toEqual([
      {
        appIDs: ['ABCDE12345.gg.foreversixty.app'],
        components: [
          ...SIM_TOOLS.map((p) => ({ '/': p, exclude: true })),
          ...['/app/*', ...SITE_PAGES].map((p) => ({ '/': p })),
        ],
      },
    ]);
    expect(aasa.webcredentials.apps).toEqual(['ABCDE12345.gg.foreversixty.app']);
  });

  it('delegates URL handling to the Android package with its signing fingerprints', () => {
    const [statement] = assetLinks({ ...cfg, androidSha256Fingerprints: ['AA:BB'] });
    expect(statement.relation).toEqual(['delegate_permission/common.handle_all_urls']);
    expect(statement.target).toEqual({
      namespace: 'android_app',
      package_name: 'gg.foreversixty.app',
      sha256_cert_fingerprints: ['AA:BB'],
    });
  });

  it('refuses a path outside /app/', () => {
    const file = path.join(import.meta.dirname, 'bad-app-links.fixture.json');
    expect(() => readAppLinksConfig(file)).toThrow(/under \/app\//);
  });

  it('admits a page outside /app/ only when links.json names it a site page', () => {
    expect(readSitePages()).toEqual(SITE_PAGES);
    expect(cfg.paths).toEqual(['/app/*', ...SITE_PAGES]);
    const file = path.join(import.meta.dirname, 'bad-app-links.fixture.json');
    expect(() => readAppLinksConfig(file, ['/account'])).not.toThrow();
  });

  it('keeps the committed .well-known files in step with app-links.json', () => {
    for (const [name, body] of Object.entries(renderAppLinks(cfg))) {
      expect(readFileSync(path.join(wellKnownDir, name), 'utf8'), name).toBe(body);
    }
    expect([AASA_FILE, ASSETLINKS_FILE]).toEqual(['apple-app-site-association', 'assetlinks.json']);
  });
});
