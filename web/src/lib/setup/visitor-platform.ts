// web/src/lib/setup/visitor-platform.ts
// Which companion download a visitor is offered (lib/setup/downloads.ts): Windows and Mac get their
// own build as the one primary button, a phone or tablet gets "runs on your PC or Mac" and a link to
// copy, and anything else (Linux, ChromeOS, a browser that will not say) gets both builds side by
// side. The server renders `unknown`, so the page never guesses wrong before hydration; the island
// reads the browser and swaps inside the same reserved row.

export type VisitorPlatform = 'windows' | 'mac' | 'mobile' | 'unknown';

export interface NavigatorLike {
  userAgentData?: { platform?: string; mobile?: boolean };
  platform?: string;
  userAgent?: string;
  maxTouchPoints?: number;
}

/** A phone or tablet in the user agent: Android, iPhone, iPad, iPod, or a "Mobile" browser. Base.astro's
 *  pre-paint script reads its source to set `data-platform="mobile"` on <html> (/setup's head hint). */
export const MOBILE_UA = /android|iphone|ipad|ipod|mobile/i;
const MOBILE_PLATFORM = /android|ios|iphone|ipad/i;

/** The visitor's system, from User-Agent Client Hints where the browser offers them, else the user agent. */
export function visitorPlatform(nav: NavigatorLike | undefined): VisitorPlatform {
  if (nav === undefined) return 'unknown';
  const hints = nav.userAgentData;
  if (hints?.mobile === true) return 'mobile';
  const ua = nav.userAgent ?? '';
  const reported = hints?.platform || nav.platform || ua;
  if (MOBILE_PLATFORM.test(reported) || MOBILE_UA.test(ua)) return 'mobile';
  // iPadOS asks for the desktop site and reports "MacIntel"; a Mac has no touch screen.
  if (/mac/i.test(reported)) return (nav.maxTouchPoints ?? 0) > 1 ? 'mobile' : 'mac';
  if (/win/i.test(reported)) return 'windows';
  return 'unknown';
}
