// web/src/lib/setup/release.test.ts
import { describe, expect, it } from 'vitest';
import {
  ADDON_MANIFEST_PATH,
  COMPANION_MANIFEST_PATH,
  downloadHref,
  fetchAddonVersion,
  fetchCompanionRelease,
  formatSize,
  parseRelease,
  releaseMeta,
  versionBehind,
} from './release';

const EXE = 'foreversixty-companion_windows_amd64.exe';
const DMG = 'foreversixty-companion_darwin_universal.dmg';
const MANIFEST = {
  schema: 1,
  family: 'companion',
  tag_name: 'companion-v0.2.37',
  assets: [
    {
      name: EXE,
      browser_download_url: `https://foreversixty.gg/dl/companion/companion-v0.2.37/${EXE}`,
      size: 16425752,
    },
    {
      name: DMG,
      browser_download_url: `https://foreversixty.gg/dl/companion/companion-v0.2.37/${DMG}`,
      size: 16762780,
    },
    { name: 'stray', browser_download_url: 'https://elsewhere.example/x', size: 1 },
    { name: 'broken', browser_download_url: 'not a url', size: 1 },
  ],
};

describe('parseRelease', () => {
  it('reads the version and each file as a path on this site', () => {
    const release = parseRelease(MANIFEST);
    expect(release?.version).toBe('0.2.37');
    expect(release?.files[DMG]).toEqual({ path: `/dl/companion/companion-v0.2.37/${DMG}`, size: 16762780 });
    expect(Object.keys(release?.files ?? {})).toEqual([EXE, DMG]);
  });

  it('refuses anything that is not a companion release', () => {
    expect(parseRelease(null)).toBeNull();
    expect(parseRelease({ ...MANIFEST, tag_name: 'addon-v2.1.0' })).toBeNull();
    expect(parseRelease({ tag_name: 'companion-v1.0.0' })).toBeNull();
  });
});

describe('downloadHref and releaseMeta', () => {
  const release = parseRelease(MANIFEST);
  it('link the release file once the manifest answered, the latest-file redirect before', () => {
    expect(downloadHref(DMG, release)).toBe(`/dl/companion/companion-v0.2.37/${DMG}`);
    expect(downloadHref(DMG, null)).toBe(`/dl/companion/${DMG}`);
  });

  it('say the version and the size in decimal megabytes', () => {
    expect(releaseMeta(DMG, release)).toBe('Version 0.2.37 · 16.8 MB');
    expect(releaseMeta(DMG, null)).toBeNull();
    expect(releaseMeta('missing', release)).toBeNull();
    expect(formatSize(16425752)).toBe('16.4 MB');
  });
});

describe('fetchCompanionRelease', () => {
  it('reads the feed manifest, and gives null when it does not answer', async () => {
    const asked: string[] = [];
    const ok = (async (url: string) => {
      asked.push(url);
      return new Response(JSON.stringify(MANIFEST));
    }) as typeof fetch;
    expect((await fetchCompanionRelease(ok))?.version).toBe('0.2.37');
    expect(asked).toEqual([COMPANION_MANIFEST_PATH]);
    expect(
      await fetchCompanionRelease((async () => new Response('', { status: 404 })) as typeof fetch),
    ).toBeNull();
    expect(
      await fetchCompanionRelease((async () => {
        throw new TypeError('offline');
      }) as typeof fetch),
    ).toBeNull();
  });
});

describe('the addon feed', () => {
  it('reads the newest addon version off its manifest, null when it answers anything else', async () => {
    const asked: string[] = [];
    const answer = (body: unknown) =>
      (async (url: string) => {
        asked.push(url);
        return new Response(JSON.stringify(body));
      }) as typeof fetch;
    expect(await fetchAddonVersion(answer({ tag_name: 'addon-v2.0.37' }))).toBe('2.0.37');
    expect(asked).toEqual([ADDON_MANIFEST_PATH]);
    expect(await fetchAddonVersion(answer({ tag_name: 'data-v2026.11.09' }))).toBeNull();
    expect(await fetchAddonVersion(answer({ tag_name: 'addon-v' }))).toBeNull();
    expect(
      await fetchAddonVersion((async () => new Response('', { status: 404 })) as typeof fetch),
    ).toBeNull();
  });

  it('calls an addon behind only on a later dotted version', () => {
    expect(versionBehind('2.0.9', '2.0.37')).toBe(true);
    expect(versionBehind('2.0.37', '2.0.37')).toBe(false);
    expect(versionBehind('2.1', '2.0.37')).toBe(false);
    expect(versionBehind('2.0', '2.0.1')).toBe(true);
    expect(versionBehind('', '2.0.37')).toBe(false);
    expect(versionBehind('2.0.37-dev', '2.0.38')).toBe(false);
  });
});
