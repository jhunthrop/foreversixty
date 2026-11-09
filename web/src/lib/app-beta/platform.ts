// web/src/lib/app-beta/platform.ts
// Which phone a visitor holds (the tester form's default, and later the store a phone is sent
// to), the store switch read from src/data/apps.json, and the email check the form runs before
// it sends. The API (api/internal/apptesters) is the authority on the address; this only saves a
// round trip for the obvious mistakes.

export type Platform = 'ios' | 'android';

/** The form's choice on a desktop, where no phone can be read. */
export const DEFAULT_PLATFORM: Platform = 'ios';

/** RFC 5321's limit on a whole address, as the API holds it. */
export const MAX_EMAIL_LENGTH = 254;

/** The phone behind a user agent. iPadOS asks for desktop pages as a Mac, so a "Macintosh" with
 *  touch points is an iPad; a Mac has none. */
export function phoneOS(userAgent: string, maxTouchPoints: number): Platform | null {
  if (/android/i.test(userAgent)) return 'android';
  if (/iphone|ipad|ipod/i.test(userAgent)) return 'ios';
  if (/macintosh/i.test(userAgent) && maxTouchPoints > 1) return 'ios';
  return null;
}

/** The form's default: the visitor's phone, else iPhone. */
export function detectPlatform(userAgent: string, maxTouchPoints: number): Platform {
  return phoneOS(userAgent, maxTouchPoints) ?? DEFAULT_PLATFORM;
}

/** src/data/apps.json: each store's listing, null until the app is live there. */
export interface AppsConfig {
  appStoreUrl: string | null;
  playStoreUrl: string | null;
}

export interface StoreLink {
  platform: Platform;
  url: string;
}

function checkedUrl(url: string): string {
  if (!url.startsWith('https://')) throw new Error(`apps.json: a store link must be https: ${url}`);
  return url;
}

/** The live stores, App Store first; empty while the switch is off (both null). */
export function storeLinks(config: AppsConfig): StoreLink[] {
  const links: StoreLink[] = [];
  if (config.appStoreUrl !== null) links.push({ platform: 'ios', url: checkedUrl(config.appStoreUrl) });
  if (config.playStoreUrl !== null) links.push({ platform: 'android', url: checkedUrl(config.playStoreUrl) });
  return links;
}

/** The store a phone goes straight to: its own, when live. */
export function storeFor(stores: readonly StoreLink[], phone: Platform | null): StoreLink | null {
  return stores.find((store) => store.platform === phone) ?? null;
}

/** One plain address with a dotted domain, within the API's length. */
export function isEmailShaped(raw: string): boolean {
  const email = raw.trim();
  return email.length <= MAX_EMAIL_LENGTH && /^[^\s@<>,]+@[^\s@<>,.]+(\.[^\s@<>,.]+)+$/.test(email);
}
