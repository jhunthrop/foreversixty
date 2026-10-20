// web/src/lib/account/connect-copy.ts
// Every visible string on /connect, the page the companion's "Sign in with Battle.net"
// opens (companion review round 1, UX "Site connect page" and the player's Battle.net
// sign-in copy). One state per key; the layout never changes between them.

export const connectCopy = {
  title: 'Connect your companion',
  /** The signed-out line. The device name is not known until a session can look the code up. */
  signedOutLead: 'Sign in to connect your companion.',
  signedOutLeadNamed: (device: string): string => `Sign in to connect ${device}.`,
  codeLabel: (code: string): string => `Code ${code}`,
  codeCheck: 'Check that this code matches the one in the companion.',
  /** Under the Code row on the pending page, where the check matters most. */
  pendingCodeCheck: 'Only press Connect if this code matches the one in the companion window.',
  consent:
    'Forever Sixty will see your BattleTag and character list. It cannot see your password or play for you.',
  battlenet: 'Continue with Battle.net',
  emailToggle: 'Email me a sign-in link instead',
  emailLabel: 'Email address',
  emailSend: 'Send link',
  checking: 'Checking the link.',
  pendingTitle: (device: string, who: string): string => `Connect ${device} to ${who}?`,
  pendingBody:
    'The companion will upload your combat logs and install your chosen builds in the addon. You can remove it any time in Account > Devices.',
  notYou: 'Not you?',
  signOut: 'Sign out',
  connect: 'Connect',
  connecting: 'Connecting...',
  cancel: 'Cancel',
  successTitle: 'Connected.',
  successBody:
    'Connected. Go back to the companion; it is ready. Now log out once in game, and your characters appear on the site within a minute.',
  yourAccount: 'Your account',
  getSetUp: 'Get set up',
  openReports: 'Open my reports',
  cancelledTitle: 'Not connected.',
  cancelledBody: 'Nothing was connected.',
  close: 'Close',
  expiredTitle: 'This link has expired.',
  expiredBody: 'Go back to the companion and press Sign in with Battle.net again.',
  backHome: 'Back to Forever Sixty',
  usedTitle: 'This device is already connected.',
  usedBody: 'If this was not you, remove it in Account > Devices.',
  failedTitle: 'Could not reach the site.',
  failedBody: 'Your device was not connected. Check your connection and try again.',
  tryAgain: 'Try again',
  deviceLabel: 'Device',
  platformLabel: 'Platform',
  codeRowLabel: 'Code',
} as const;
