// web/src/lib/app-beta/copy.ts
// The words on foreversixty.gg/app, the page the apps' QR codes open: the apps are coming, what
// they do (only what is built), the tester signup, the QR for a desktop visitor, and, once
// src/data/apps.json has a store link, the store buttons. design/mocks/ws_app_beta.py reads them
// through design/mocks/app_beta_facts.mjs.
import type { Platform } from './platform';
import { uiCopy } from '../ui/copy';

/** The page's own address, the one the QR encodes. */
export const APP_PAGE_URL = 'https://foreversixty.gg/app';

export const PLATFORM_NAMES: Readonly<Record<Platform, string>> = { ios: 'iPhone', android: 'Android' };

export const appBetaCopy = {
  pageTitle: 'Apps for iPhone and Android',
  pageDescription:
    'The Forever Sixty apps for iPhone and Android are coming soon: raid sign-ups from the notification, raid night on your lock screen, loot and sim pushes. Sign up to test them.',
  heading: 'Forever Sixty for iPhone and Android',
  soon: 'Coming soon',
  featuresTitle: 'What the apps do',
  features: [
    'Answer a raid sign-up straight from the notification: In, Late, Bench or Out.',
    'Raid night on your lock screen.',
    'A push when you get loot, or when your sim or report is ready.',
    'Your next upgrade, and where it drops.',
  ],
  formTitle: 'Test the apps',
  formJob: 'The beta opens before the stores do. Leave your email and pick your phone.',
  emailLabel: 'Email',
  emailPlaceholder: 'you@example.com',
  platformLabel: 'Your phone',
  submit: 'Sign up to test',
  submitting: 'Signing up',
  consent: "We'll only email you about the app beta.",
  invalidEmail: 'Enter one email address, like you@example.com.',
  failed: `That didn't go through. ${uiCopy.retry} in a moment.`,
  tooMany: `Too many sign-ups from here. ${uiCopy.retry} in an hour.`,
  /** The one line after a send, a new address or a repeat alike (the API tells them apart to no one). */
  received: (platform: Platform): string =>
    `You're on the list for ${PLATFORM_NAMES[platform]}. We'll email you when the beta opens.`,
  qrTitle: 'On your phone',
  qrLine: 'Scan with your phone to open this page.',
  qrAlt: 'QR code for foreversixty.gg/app',
  storesTitle: 'Get the apps',
  stores: { ios: 'Download on the App Store', android: 'Get it on Google Play' } as Readonly<
    Record<Platform, string>
  >,
} as const;
