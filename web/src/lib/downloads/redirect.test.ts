import { describe, expect, it } from 'vitest';
import companion from '../../data/companion.json';
import { ADDON_DOWNLOAD_PATH, companionDownloadPath, DATA_DOWNLOAD_PATH, parseDownloadPath } from './paths';
import { releaseLocation, serveDownload } from './redirect';

const RELEASES = 'https://releases.example/acme/site/releases';
const ORIGIN = 'https://foreversixty.gg';
const [windows] = companion.downloads;

function get(path: string, method = 'GET'): Request {
  return new Request(`${ORIGIN}${path}`, { method });
}

describe('parseDownloadPath', () => {
  const assets = ['app.exe'];

  it('reads the addon, the data pack and a listed companion build', () => {
    expect(parseDownloadPath(ADDON_DOWNLOAD_PATH, assets)).toEqual({ kind: 'addon' });
    expect(parseDownloadPath(DATA_DOWNLOAD_PATH, assets)).toEqual({ kind: 'data' });
    expect(parseDownloadPath(companionDownloadPath('app.exe'), assets)).toEqual({
      kind: 'companion',
      asset: 'app.exe',
    });
  });

  it('names nothing the site does not offer', () => {
    expect(parseDownloadPath('/dl/', assets)).toBeNull();
    expect(parseDownloadPath('/dl/addon/extra', assets)).toBeNull();
    expect(parseDownloadPath(companionDownloadPath('other.exe'), assets)).toBeNull();
    expect(parseDownloadPath(companionDownloadPath('../../x'), assets)).toBeNull();
  });
});

describe('releaseLocation', () => {
  it('sends the addon and data pack to their tagged release lists, a companion to the latest asset', () => {
    expect(releaseLocation(RELEASES, { kind: 'addon' })).toBe(`${RELEASES}?q=addon-v`);
    expect(releaseLocation(RELEASES, { kind: 'data' })).toBe(`${RELEASES}?q=data-v`);
    expect(releaseLocation(RELEASES, { kind: 'companion', asset: 'a.exe' })).toBe(
      `${RELEASES}/latest/download/a.exe`,
    );
  });
});

describe('serveDownload', () => {
  const env = { RELEASES_URL: RELEASES };

  it('answers a listed download with a short-cached 302 to the release host', () => {
    const response = serveDownload(get(companionDownloadPath(windows.asset)), env);
    expect(response.status).toBe(302);
    expect(response.headers.get('location')).toBe(`${RELEASES}/latest/download/${windows.asset}`);
    expect(response.headers.get('cache-control')).toBe('public, max-age=60');
    expect(serveDownload(get(ADDON_DOWNLOAD_PATH, 'HEAD'), env).status).toBe(302);
  });

  it('404s an unlisted name instead of redirecting it', () => {
    const response = serveDownload(get('/dl/companion/evil.exe'), env);
    expect(response.status).toBe(404);
    expect(response.headers.get('location')).toBeNull();
  });

  it('refuses a write and says so when the host is not configured', () => {
    expect(serveDownload(get(ADDON_DOWNLOAD_PATH, 'POST'), env).status).toBe(405);
    expect(serveDownload(get(ADDON_DOWNLOAD_PATH), {}).status).toBe(503);
  });
});
