// web/src/lib/logs/panel-order.ts
// The order of /logs' right column, "ways in" (design/specs/2026-10-11-logs.md section 2.5): the
// companion first when a device is paired (a player who logs live goes there), Upload first for
// everyone else, signed out included.
export type WaysPanel = 'companion' | 'upload';

export interface WaysState {
  signedIn: boolean;
  /** At least one device is paired with the account. */
  paired: boolean;
}

export function waysOrder({ signedIn, paired }: WaysState): readonly [WaysPanel, WaysPanel] {
  return signedIn && paired ? ['companion', 'upload'] : ['upload', 'companion'];
}
