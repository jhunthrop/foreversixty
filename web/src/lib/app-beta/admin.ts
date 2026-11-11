// web/src/lib/app-beta/admin.ts
// The owner's /admin/testers: who asked to test the phone apps (GET /v1/admin/app-testers), as a
// Grid of address, phone and day, and each platform's addresses ready to paste into TestFlight or
// the Play Console's tester list. design/mocks/ws_app_beta.py reads the copy through
// design/mocks/app_beta_facts.mjs.
import { utcDayMonthYear } from '../dates';
import type { GridRowModel } from '../ui/cells';
import type { AppTester } from './api';
import { PLATFORM_NAMES } from './copy';
import type { Platform } from './platform';
import { uiCopy } from '../ui/copy';

export const PLATFORM_ORDER: readonly Platform[] = ['ios', 'android'];

export const adminTestersCopy = {
  pageTitle: 'App testers',
  job: 'Who asked to test the phone apps, newest first. Copy a platform’s addresses into TestFlight or the Play Console.',
  pageDescription: 'The owner’s list of phone app beta sign-ups.',
  signedOut: 'Sign in with the owner’s account to see the testers.',
  notOwner: 'This page is for the site’s owner. Your account can’t see the testers.',
  loadFailed: 'The testers didn’t load.',
  tableTitle: 'Sign-ups',
  caption: 'App beta sign-ups',
  emailHeader: 'Email',
  columns: ['Phone', 'Signed up'],
  empty: 'No sign-ups yet. foreversixty.gg/app is where they come from.',
  copyEmails: (platform: Platform, count: number): string =>
    `Copy ${count} ${PLATFORM_NAMES[platform]} ${count === 1 ? 'email' : 'emails'}`,
  copied: uiCopy.copied,
  copyFailed: 'Couldn’t copy; select the addresses by hand.',
} as const;

/** One platform's addresses in list order, comma separated (what both stores' tester lists take). */
export function emailsFor(testers: readonly AppTester[], platform: Platform): string {
  return testers
    .filter((t) => t.platform === platform)
    .map((t) => t.email)
    .join(', ');
}

export function countsByPlatform(testers: readonly AppTester[]): Record<Platform, number> {
  return {
    ios: testers.filter((t) => t.platform === 'ios').length,
    android: testers.filter((t) => t.platform === 'android').length,
  };
}

/** The Grid's rows: the address, then the phone and the UTC day it signed up. */
export function testerRows(testers: readonly AppTester[]): GridRowModel[] {
  return testers.map((t) => ({
    key: `${t.email} ${t.platform}`,
    label: t.email,
    cells: [
      { tone: 'none', text: PLATFORM_NAMES[t.platform] },
      { tone: 'none', text: utcDayMonthYear(t.created_at) },
    ],
  }));
}
