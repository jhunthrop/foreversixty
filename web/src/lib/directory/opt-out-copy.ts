// web/src/lib/directory/opt-out-copy.ts
// The "Remove my character" page's words. Blizzard is named as the source (API terms 2(m)).
export const optOutCopy = {
  title: 'Remove my character',
  lead: 'Forever Sixty reads public character data from Blizzard so guild rosters and Discord raid sign-ups know who is who. Remove a character and we delete what we hold for it and never store it again.',
  findHeading: 'Your character',
  regionLabel: 'Region',
  nameLabel: 'Full name',
  namePlaceholder: 'First name and surname',
  nameInvalid: 'Enter the first name and surname as they show in game.',
  battlenetHeading: 'Sign in with Battle.net',
  battlenetLine: 'Quickest. The character has to be on the Battle.net account you sign in with.',
  battlenetSignIn: 'Sign in with Battle.net',
  battlenetRemove: (name: string): string => `Remove ${name}`,
  noteHeading: 'Use your guild note',
  noteLine:
    "No account needed. Put a code in the character's public guild note; when a guildmate who runs the Forever Sixty addon logs out, we remove it. For a character in a guild.",
  getCode: 'Get a code',
  codeLead: (name: string): string => `Put this in ${name}'s public guild note:`,
  codeExpires: (date: string): string =>
    `The code works until ${date}. You can take it out of the note once this page says Removed.`,
  check: 'Check now',
  checking: 'Checking',
  waiting: 'Not yet. A guildmate running the addon has to log out after the note is saved.',
  removedTitle: 'Removed',
  removed: (name: string): string => `We hold nothing for ${name} and will not store it again.`,
  notFound: 'No public character has that name in that region, so we hold nothing for it.',
  notYours: 'That character is not on the Battle.net account you signed in with. Use the guild note instead.',
  failed: 'That did not go through. Try again in a minute.',
  regions: [
    { value: 'us', label: 'Americas' },
    { value: 'eu', label: 'Europe' },
  ] as const,
} as const;
