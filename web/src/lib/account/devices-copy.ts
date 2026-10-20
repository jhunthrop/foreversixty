// web/src/lib/account/devices-copy.ts
// The Devices panel's strings and the rules that choose between them (lifecycle round 3,
// finding 9): which pairing entry a panel shows, and what Revoke asks before it signs a
// companion out.
export const devicesCopy = {
  revoke: 'Revoke',
  cancel: 'Cancel',
  revokeQuestion: (deviceName: string): string => `Revoke ${deviceName}? The companion signs out.`,
  /** The quiet link beside a paired device: a second computer, not the next step. */
  pairAnother: 'Pair another computer',
  /** The line above the code: the code is the fallback for a companion that cannot open a browser. */
  codeFallbackLine:
    'Only for a companion that cannot open a browser; the companion signs in with Battle.net by itself.',
} as const;

/** What the panel's pairing entry is called: with no device it stays the plain fallback entry. */
export function pairingEntryLabel(deviceCount: number, getCodeLabel: string): string {
  return deviceCount > 0 ? devicesCopy.pairAnother : getCodeLabel;
}
