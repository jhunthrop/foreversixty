// web/src/lib/setup/downloads.test.ts
import { describe, expect, it } from 'vitest';
import { COMPANION_DOWNLOADS, downloadFor, isMacPlatform, type CompanionDownload } from './downloads';

const WINDOWS: CompanionDownload = {
  platform: 'Windows',
  asset: 'foreversixty-companion_windows_amd64.exe',
  firstRun: 'w',
};
const MAC: CompanionDownload = {
  platform: 'Mac',
  asset: 'foreversixty-companion_darwin_universal.dmg',
  firstRun: 'm',
};

describe('downloadFor', () => {
  it('gives Windows the Windows build whatever the list order', () => {
    expect(downloadFor('windows', [MAC, WINDOWS])).toBe(WINDOWS);
    expect(downloadFor('windows', [WINDOWS])).toBe(WINDOWS);
  });

  it('gives a Mac the Mac build, or nothing while the feed carries none', () => {
    expect(downloadFor('mac', [WINDOWS, MAC])).toBe(MAC);
    expect(downloadFor('mac', [WINDOWS])).toBeNull();
  });

  it('reads companion.json by default, and its Windows entry is first', () => {
    expect(downloadFor('windows')).toBe(COMPANION_DOWNLOADS[0]);
    expect(COMPANION_DOWNLOADS[0].platform).toBe('Windows');
  });

  it('names a Mac build by its platform word', () => {
    expect(isMacPlatform('Mac')).toBe(true);
    expect(isMacPlatform('macOS')).toBe(true);
    expect(isMacPlatform('Windows')).toBe(false);
  });
});
