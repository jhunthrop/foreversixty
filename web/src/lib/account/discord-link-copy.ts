// web/src/lib/account/discord-link-copy.ts
// Every visible string of the account page's Discord link row (experience spec
// 2026-10-27-discord-bot §6.1, "Account lead", "Account controls" and "Account returns").
// discord-link-copy.test.ts reads each one out of the spec.

export const discordLinkCopy = {
  lead: "Link your Discord account so your guild's sign-up buttons and Forever Sixty commands know your characters.",
  leadSuggested: (guild: string): string =>
    `${guild} signs up in Discord. Link your Discord account to use the buttons.`,
  link: 'Link Discord',
  opening: 'Opening Discord…',
  linkedLead: 'Linked to Discord as',
  unlink: 'Unlink',
  unlinking: 'Unlinking…',
  unlinkQuestion: 'Unlink Discord? Sign-up buttons and commands will ask you to link again.',
  cancel: 'Cancel',
  unlinked: 'Unlinked.',
  unlinkFailed: 'Could not unlink. Try again.',
  returnLinked: 'Linked. In Discord, sign-up buttons and commands now know your characters.',
  returnDenied: 'Cancelled in Discord. Nothing was linked.',
  returnFailed: 'Could not link. Try again.',
  returnConflict: 'That Discord account is linked to another Forever Sixty account.',
  conflictHelp: 'Sign in to that account and unlink it on its account page, then try again.',
  stuck: 'Discord did not open. Try again.',
} as const;
