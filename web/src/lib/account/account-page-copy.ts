// web/src/lib/account/account-page-copy.ts
// Every visible string /account's page header and rail panels use (brief 2026-09-22 §B1,
// §B2) that character-list-copy.ts and signin-copy.ts do not already carry.

export const accountPageCopy = {
  title: 'Your account',
  signedInWithBattlenet: 'Signed in with Battle.net',
  signedInByEmail: 'Signed in by email',
  signOut: 'Sign out',
  devicesLabel: 'Devices',
  youLabel: 'You',
  guildsAndPlanLabel: 'Guilds and plan',
  noDevices: 'No devices paired.',
  pairADevice: 'Pair a device',
  pseudonymLabel: 'Show a pseudonym instead of my character names',
  pseudonymNote:
    'Applies everywhere your characters appear, on reports and rankings alike. Reports themselves are never deleted or rewritten.',
  planKey: 'Plan',
} as const;
