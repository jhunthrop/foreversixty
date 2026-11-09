import { describe, expect, it } from 'vitest';
import { DEFAULT_PLATFORM, detectPlatform, isEmailShaped, phoneOS, storeFor, storeLinks } from './platform';

const IPHONE =
  'Mozilla/5.0 (iPhone; CPU iPhone OS 18_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Mobile/15E148 Safari/604.1';
const IPAD_DESKTOP_UA =
  'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1 Safari/605.1.15';
const ANDROID =
  'Mozilla/5.0 (Linux; Android 15; Pixel 9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Mobile Safari/537.36';
const WINDOWS =
  'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36';

describe('the phone a visitor holds', () => {
  it('reads an iPhone, an iPad passing as a Mac, and Android', () => {
    expect(phoneOS(IPHONE, 5)).toBe('ios');
    expect(phoneOS(IPAD_DESKTOP_UA, 5)).toBe('ios');
    expect(phoneOS(ANDROID, 5)).toBe('android');
  });

  it('reads no phone on a desktop, a Mac with no touch included', () => {
    expect(phoneOS(WINDOWS, 0)).toBeNull();
    expect(phoneOS(IPAD_DESKTOP_UA, 0)).toBeNull();
  });

  it('defaults the form to the phone, else to iPhone', () => {
    expect(detectPlatform(ANDROID, 5)).toBe('android');
    expect(detectPlatform(IPHONE, 5)).toBe('ios');
    expect(detectPlatform(WINDOWS, 0)).toBe(DEFAULT_PLATFORM);
    expect(DEFAULT_PLATFORM).toBe('ios');
  });
});

describe('the store switch', () => {
  it('is off while both links are null', () => {
    expect(storeLinks({ appStoreUrl: null, playStoreUrl: null })).toEqual([]);
  });

  it('lists only the stores that have a link, App Store first', () => {
    const both = {
      appStoreUrl: 'https://apps.apple.com/app/id1',
      playStoreUrl: 'https://play.google.com/store/apps/details?id=gg',
    };
    expect(storeLinks(both).map((s) => s.platform)).toEqual(['ios', 'android']);
    expect(storeLinks({ ...both, appStoreUrl: null }).map((s) => s.platform)).toEqual(['android']);
  });

  it('refuses a link that is not https', () => {
    expect(() => storeLinks({ appStoreUrl: 'http://apps.apple.com/x', playStoreUrl: null })).toThrow(/https/);
  });

  it('sends a phone to its own store only when that store is live', () => {
    const stores = storeLinks({ appStoreUrl: 'https://apps.apple.com/app/id1', playStoreUrl: null });
    expect(storeFor(stores, 'ios')?.url).toBe('https://apps.apple.com/app/id1');
    expect(storeFor(stores, 'android')).toBeNull();
    expect(storeFor(stores, null)).toBeNull();
  });
});

describe('the email check before sending', () => {
  it('takes one plain address', () => {
    expect(isEmailShaped(' you@example.com ')).toBe(true);
  });

  it('refuses what the API would refuse anyway', () => {
    for (const bad of [
      '',
      'you',
      'you@',
      '@example.com',
      'you@localhost',
      'a b@example.com',
      `${'a'.repeat(250)}@x.co`,
    ]) {
      expect(isEmailShaped(bad), bad).toBe(false);
    }
  });
});
