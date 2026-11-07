// web/src/lib/account/platform-label.ts
// A device's platform as a player reads it (setup spec G5, tenet 7): the companion reports Go's
// `GOOS/GOARCH` ("windows/amd64"); the site says "Windows". One mapping for /setup's device
// lines, /logs and /connect.

const LABELS: readonly (readonly [prefix: string, label: string])[] = [
  ['windows', 'Windows'],
  ['darwin', 'macOS'],
  ['macos', 'macOS'],
  ['linux', 'Linux'],
  ['ios', 'iOS'],
  ['android', 'Android'],
];

/** "windows/amd64" -> "Windows"; an already friendly or unknown value is returned as sent. */
export function platformLabel(platform: string): string {
  const head = platform.trim().toLowerCase().split('/')[0];
  return LABELS.find(([prefix]) => head === prefix)?.[1] ?? platform;
}
