// web/scripts/app-links.mjs
// The two files that let https://foreversixty.gg/app/* open the native apps instead of the
// browser: Apple's apple-app-site-association (universal links) and Android's
// assetlinks.json (App Links). Both are built from web/app-links.json, the one place the
// Team ID, the bundle id, the package name and the signing fingerprints live, and written
// into public/.well-known/ so the static build serves them at /.well-known/.
//
// Only the configured paths (/app/*) open the apps; every other URL stays on the site. The
// API's app sign-in hands its one-time code to /app/auth (api/internal/auth/app_auth.go).
//
// `npm run sync:app-links` rewrites the files; prebuild runs it so a deploy always carries
// the config's current values.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
export const configPath = path.join(webRoot, 'app-links.json');
export const wellKnownDir = path.join(webRoot, 'public', '.well-known');
export const AASA_FILE = 'apple-app-site-association';
export const ASSETLINKS_FILE = 'assetlinks.json';

/** Reads and checks web/app-links.json; a missing or empty field fails loudly. */
export function readAppLinksConfig(file = configPath) {
  const cfg = JSON.parse(readFileSync(file, 'utf8'));
  for (const key of ['appleTeamId', 'iosBundleId', 'androidPackage']) {
    if (typeof cfg[key] !== 'string' || cfg[key] === '')
      throw new Error(`app-links.json: ${key} is required`);
  }
  for (const key of ['androidSha256Fingerprints', 'paths']) {
    if (!Array.isArray(cfg[key]) || cfg[key].length === 0)
      throw new Error(`app-links.json: ${key} must be a non-empty list`);
  }
  if (!cfg.paths.every((p) => typeof p === 'string' && p.startsWith('/app/'))) {
    throw new Error('app-links.json: every path must sit under /app/, so no other page opens the apps');
  }
  return cfg;
}

/** Apple's universal links file: the app id is "<Team ID>.<bundle id>". */
export function appleAppSiteAssociation(cfg) {
  const appID = `${cfg.appleTeamId}.${cfg.iosBundleId}`;
  return {
    applinks: {
      details: [{ appIDs: [appID], components: cfg.paths.map((p) => ({ '/': p })) }],
    },
    webcredentials: { apps: [appID] },
  };
}

/** Android's Digital Asset Links statement for App Links. */
export function assetLinks(cfg) {
  return [
    {
      relation: ['delegate_permission/common.handle_all_urls'],
      target: {
        namespace: 'android_app',
        package_name: cfg.androidPackage,
        sha256_cert_fingerprints: cfg.androidSha256Fingerprints,
      },
    },
  ];
}

/** The two files' bytes, keyed by file name. */
export function renderAppLinks(cfg) {
  const render = (value) => `${JSON.stringify(value, null, 2)}\n`;
  return { [AASA_FILE]: render(appleAppSiteAssociation(cfg)), [ASSETLINKS_FILE]: render(assetLinks(cfg)) };
}

export function writeAppLinks(cfg = readAppLinksConfig(), dir = wellKnownDir) {
  mkdirSync(dir, { recursive: true });
  for (const [name, body] of Object.entries(renderAppLinks(cfg))) writeFileSync(path.join(dir, name), body);
}

if (process.argv[1] === fileURLToPath(import.meta.url)) writeAppLinks();
