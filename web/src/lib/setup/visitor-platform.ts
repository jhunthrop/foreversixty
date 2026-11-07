// web/src/lib/setup/visitor-platform.ts
// Which companion row a visitor sees (setup spec U10, G4): a Mac gets the Mac download when
// companion.json lists one (lib/setup/downloads.ts) and the export route when not, every other
// visitor the Windows download. Server renders `windows`; the island swaps after hydration
// inside the same reserved row.

export type VisitorPlatform = 'windows' | 'mac';

interface NavigatorLike {
  userAgentData?: { platform?: string };
  platform?: string;
  userAgent?: string;
}

const MAC = /mac/i;
/** iPads report "MacIntel" with touch; they cannot run the companion either way, so they read as a Mac. */
export function visitorPlatform(nav: NavigatorLike | undefined): VisitorPlatform {
  if (nav === undefined) return 'windows';
  const reported = nav.userAgentData?.platform ?? nav.platform ?? nav.userAgent ?? '';
  return MAC.test(reported) ? 'mac' : 'windows';
}
