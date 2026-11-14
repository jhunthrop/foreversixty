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
  battlenetHeading: 'Remove with Battle.net',
  battlenetLine: 'Sign in with Battle.net to remove any character on your account, in a guild or not.',
  battlenetSignIn: 'Sign in with Battle.net',
  battlenetRemove: (name: string): string => `Remove ${name}`,
  noteHeading: 'Remove with your guild note',
  noteLine:
    'For a character in a guild, no account needed. Put a code in its public guild note, not the officer note. A guildmate who runs the Forever Sixty addon sends the note when they log out, and we remove the character.',
  getCode: 'Get a code',
  codeLead: (name: string): string => `Put this in ${name}'s public guild note:`,
  codeExpires: (date: string): string => `The code works until ${date}.`,
  check: 'Check now',
  checking: 'Checking',
  waiting: 'Not yet. A guildmate running the addon has to log out after saving the note.',
  removedTitle: 'Removed',
  removeAnother: 'Remove another character',
  removed: (name: string): string => `We hold nothing for ${name} and will not store it again.`,
  notYours: 'That character is not on the Battle.net account you signed in with. Use the guild note instead.',
  failed: 'That did not go through. Try again in a minute.',
  regions: [
    { value: 'us', label: 'Americas' },
    { value: 'eu', label: 'Europe' },
  ] as const,
} as const;
