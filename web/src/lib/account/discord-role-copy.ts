// web/src/lib/account/discord-role-copy.ts
// Every visible string of the account page's Discord Premium role row.
import { withRetry } from '../ui/copy';

export const discordRoleCopy = {
  label: 'Premium role in Discord',
  lead: 'Gives you the Premium role in the Forever Sixty Discord while you have Premium. Link now and it turns on when you subscribe.',
  /** An account that already has Premium (ruling A4). */
  leadOwned: 'Link it and the role turns on.',
  link: 'Link Premium role',
  returnLinked: 'Linked. Your Discord Premium role now follows your Forever Sixty Premium.',
  returnDenied: 'Cancelled in Discord. Nothing was linked.',
  returnFailed: withRetry('Could not link the Premium role.'),
  stuck: withRetry('Discord did not open.'),
} as const;
