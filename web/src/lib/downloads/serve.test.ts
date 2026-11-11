import { describe, expect, it } from 'vitest';
import companion from '../../data/companion.json';
import type { Manifest } from './manifest';
import { parseManifest } from './manifest';
import { ADDON_DOWNLOAD_PATH, companionDownloadPath, DATA_DOWNLOAD_PATH, parseDownloadPath } from './paths';
import { serveDownload, type DownloadBucket, type DownloadEnv } from './serve';

const ORIGIN = 'https://foreversixty.gg';
const LEGACY = 'https://releases.example/acme/site/releases';
const [windows] = companion.downloads;

function manifest(family: Manifest['family'], tag: string, names: string[]): Manifest {
  return {
    schema: 1,
    family,
    tag_name: tag,
    published_at: '2026-11-06T10:00:00Z',
    assets: names.map((name) => ({
      name,
      browser_download_url: `${ORIGIN}/dl/${family}/${tag}/${name}`,
      size: 3,
      sha256: 'ab',
    })),
  };
}

/** A bucket holding these keys, each body its text. */
function bucket(objects: Record<string, string>): DownloadBucket {
  return {
    async get(key) {
      const text = objects[key];
      if (text === undefined) return null;
      return {
        body: new Response(text).body,
        httpEtag: `"${key}"`,
        size: text.length,
        text: async () => text,
      };
    },
  };
}

const FEED = {
  'downloads/companion/latest.json': JSON.stringify(
    manifest('companion', 'companion-v0.2.31', [windows.asset, `${windows.asset}.minisig`]),
  ),
  [`downloads/companion/companion-v0.2.31/${windows.asset}`]: 'exe',
  'downloads/addon/latest.json': JSON.stringify(
    manifest('addon', 'addon-v2.1.0', ['ForeverSixty-2.1.0.zip', 'ForeverSixty-2.1.0.zip.minisig']),
  ),
};

function get(path: string, env: DownloadEnv, method = 'GET'): Promise<Response> {
  return serveDownload(new Request(`${ORIGIN}${path}`, { method }), env);
}

describe('parseDownloadPath', () => {
  it('reads every shape the feed serves', () => {
    expect(parseDownloadPath(ADDON_DOWNLOAD_PATH)).toEqual({ kind: 'latest-zip', family: 'addon' });
    expect(parseDownloadPath(DATA_DOWNLOAD_PATH)).toEqual({ kind: 'latest-zip', family: 'data' });
    expect(parseDownloadPath('/dl/data/latest.json')).toEqual({ kind: 'manifest', family: 'data' });
    expect(parseDownloadPath(companionDownloadPath('a.exe'))).toEqual({
      kind: 'latest-file',
      family: 'companion',
      name: 'a.exe',
    });
    expect(parseDownloadPath('/dl/data/data-v2026.11.06/ForeverSixtyData-2026.11.06.zip')).toEqual({
      kind: 'file',
      family: 'data',
      tag: 'data-v2026.11.06',
      name: 'ForeverSixtyData-2026.11.06.zip',
    });
  });

  it('refuses unknown families, foreign tags, traversal and extra segments', () => {
    for (const path of [
      '/dl/',
      '/dl/companion',
      '/dl/engine/latest.json',
      '/dl/addon/data-v2026.11.06/x.zip',
      '/dl/addon/addon-v1..2/x.zip',
      '/dl/addon/..',
      '/dl/addon/.hidden',
      '/dl/addon/addon-v1.0.0/x.zip/more',
    ]) {
      expect(parseDownloadPath(path), path).toBeNull();
    }
  });
});

describe('parseManifest', () => {
  const good = manifest('data', 'data-v2026.11.06', ['ForeverSixtyData-2026.11.06.zip']);

  it('takes a manifest of its own family', () => {
    expect(parseManifest(JSON.stringify(good), 'data')?.tag_name).toBe('data-v2026.11.06');
  });

  it('treats anything else as no manifest', () => {
    expect(parseManifest('not json', 'data')).toBeNull();
    expect(parseManifest(JSON.stringify(good), 'addon')).toBeNull();
    expect(parseManifest(JSON.stringify({ ...good, schema: 2 }), 'data')).toBeNull();
    expect(parseManifest(JSON.stringify({ ...good, tag_name: 'addon-v1.0.0' }), 'data')).toBeNull();
    const badName = { ...good, assets: [{ ...good.assets[0], name: '../x' }] };
    expect(parseManifest(JSON.stringify(badName), 'data')).toBeNull();
  });
});

describe('serveDownload from the feed', () => {
  const env = { LOGS: bucket(FEED), RELEASES_URL: LEGACY };

  it('serves a manifest, short-cached', async () => {
    const response = await get('/dl/companion/latest.json', env);
    expect(response.status).toBe(200);
    expect(response.headers.get('content-type')).toContain('application/json');
    expect(response.headers.get('cache-control')).toBe('public, max-age=60');
    expect(JSON.parse(await response.text()).tag_name).toBe('companion-v0.2.31');
  });

  it('serves a published file as an immutable attachment, HEAD without a body', async () => {
    const path = `/dl/companion/companion-v0.2.31/${windows.asset}`;
    const response = await get(path, env);
    expect(response.status).toBe(200);
    expect(await response.text()).toBe('exe');
    expect(response.headers.get('cache-control')).toBe('public, max-age=31536000, immutable');
    expect(response.headers.get('content-disposition')).toBe(`attachment; filename="${windows.saveAs}"`);
    const head = await get(path, env, 'HEAD');
    expect(head.status).toBe(200);
    expect(head.body).toBeNull();
  });

  it('saves each /setup build under its friendly name, and every updater file under its own', async () => {
    const tag = 'companion-v0.2.46';
    const names = [
      'foreversixty-companion-setup.exe',
      'foreversixty-companion_darwin_universal.dmg',
      'foreversixty-companion_windows_amd64.exe',
      'foreversixty-companion_windows_amd64.exe.minisig',
      'foreversixty-companion_darwin_universal.app.zip',
      'foreversixty-companion_darwin_universal.app.zip.minisig',
    ];
    const files = Object.fromEntries(names.map((name) => [`downloads/companion/${tag}/${name}`, name]));
    const feed = { LOGS: bucket(files) };
    const savedAs = async (name: string): Promise<string | null> =>
      (await get(`/dl/companion/${tag}/${name}`, feed)).headers.get('content-disposition');
    expect(await savedAs('foreversixty-companion-setup.exe')).toBe(
      'attachment; filename="ForeverSixtySetup.exe"',
    );
    expect(await savedAs('foreversixty-companion_darwin_universal.dmg')).toBe(
      'attachment; filename="ForeverSixty.dmg"',
    );
    for (const name of names.slice(2)) {
      expect(await savedAs(name)).toBe(`attachment; filename="${name}"`);
    }
    // The manifest is read, not saved: no attachment at all.
    expect((await get('/dl/companion/latest.json', env)).headers.get('content-disposition')).toBeNull();
  });

  it('redirects the latest companion build and the latest addon zip to their own paths, never the old host', async () => {
    const exe = await get(companionDownloadPath(windows.asset), env);
    expect(exe.status).toBe(302);
    expect(exe.headers.get('location')).toBe(`/dl/companion/companion-v0.2.31/${windows.asset}`);
    const zip = await get(ADDON_DOWNLOAD_PATH, env);
    expect(zip.headers.get('location')).toBe('/dl/addon/addon-v2.1.0/ForeverSixty-2.1.0.zip');
  });

  it('404s a file the latest release does not carry and a file not in the bucket', async () => {
    expect((await get(companionDownloadPath('other.exe'), env)).status).toBe(404);
    expect((await get('/dl/addon/addon-v9.9.9/ForeverSixty-9.9.9.zip', env)).status).toBe(404);
  });

  it('refuses a write', async () => {
    expect((await get(ADDON_DOWNLOAD_PATH, env, 'POST')).status).toBe(405);
  });
});

describe('serveDownload before a family has published to the feed', () => {
  const env = { LOGS: bucket({}), RELEASES_URL: LEGACY };

  it('falls back to the old host for the latest paths it had', async () => {
    expect((await get(DATA_DOWNLOAD_PATH, env)).headers.get('location')).toBe(`${LEGACY}?q=data-v`);
    expect((await get(companionDownloadPath(windows.asset), env)).headers.get('location')).toBe(
      `${LEGACY}/latest/download/${windows.asset}`,
    );
  });

  it('has no fallback for a name the site never offered, nor once the old host is unset', async () => {
    expect((await get(companionDownloadPath('evil.exe'), env)).status).toBe(404);
    expect((await get(DATA_DOWNLOAD_PATH, { LOGS: bucket({}) })).status).toBe(404);
    expect((await get('/dl/data/latest.json', {})).status).toBe(404);
  });
});
