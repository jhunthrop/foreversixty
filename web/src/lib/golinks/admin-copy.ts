// web/src/lib/golinks/admin-copy.ts
// Every word on /admin/links, the owner's tracked links page.
import { countLabel } from './funnel';
import { uiCopy } from '../ui/copy';

export const adminLinksCopy = {
  pageTitle: 'Links',
  job: 'Where people come from: clicks counted without cookies, then the accounts each link brought in.',
  pageDescription: 'The owner’s tracked links and the accounts each one brought in.',
  signedOut: 'Sign in with the owner’s account to see the links.',
  notOwner: 'This page is for the site’s owner. Your account can’t see the links.',
  loadFailed: 'The links didn’t load.',
  tableTitle: 'Your links',
  tableMeta: (count: number): string => `${count} ${count === 1 ? 'link' : 'links'}`,
  empty: 'No links yet. Make the first one below.',
  columns: {
    link: 'Link',
    clicks: 'Clicks',
    signIns: 'Signed up',
    linked: 'Character linked',
    premium: 'Premium',
  },
  /** What each step's percent is of: the step before it. */
  shareOf: {
    signIns: 'of clicks',
    linked: 'of signed up',
    premium: 'of linked',
  },
  /** Under a link's clicks: bots and link previews, which the click count leaves out. */
  bots: (count: number): string =>
    count === 0 ? 'no bots' : `+${countLabel(count)} ${count === 1 ? 'bot' : 'bots'}, not counted`,
  noClicks: 'No clicks yet',
  totals: 'All links',
  copy: 'Copy',
  copied: uiCopy.copied,
  copyLabel: (slug: string): string => `Copy the link for ${slug}`,
  formTitle: 'New link',
  formJob: 'The link counts every click and sends people to the page you pick, on this site only.',
  slugLabel: 'Slug',
  slugHint: 'Lowercase letters, digits and dashes. It becomes foreversixty.gg/go/your-slug.',
  targetLabel: 'Sends people to',
  labelLabel: 'Label',
  create: 'Make link',
  creating: 'Saving',
  created: (slug: string): string => `foreversixty.gg/go/${slug} is live.`,
  createFailed: 'That link wasn’t saved; try again.',
} as const;
