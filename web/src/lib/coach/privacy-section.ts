// web/src/lib/coach/privacy-section.ts
// The privacy page's "Log coach" section (log coach spec §9: what leaves our systems). It is
// written now and renders only once the coach line does: the plans page holds that line in
// `premium.featuresHeld` until the coach ships, and Task 39 deletes the key.
import { plansPageCopy } from '../billing/plans-page-copy';
import { COACH_MODEL_PROVIDER } from './provider';

export const COACH_PRIVACY_HEADING = 'Log coach';
export const VERCEL_PRIVACY_URL = 'https://vercel.com/legal/privacy-policy';

export interface CoachPrivacySection {
  readonly heading: string;
  readonly paragraphs: readonly string[];
  readonly link: { readonly label: string; readonly href: string };
}

interface PremiumCopyShape {
  readonly featuresHeld?: unknown;
}

export function coachIsShipped(premium: PremiumCopyShape = plansPageCopy.premium): boolean {
  return premium.featuresHeld === undefined;
}

export function coachPrivacySection(
  premium: PremiumCopyShape = plansPageCopy.premium,
  provider: string = COACH_MODEL_PROVIDER,
): CoachPrivacySection | null {
  if (!coachIsShipped(premium)) return null;
  return {
    heading: COACH_PRIVACY_HEADING,
    link: { label: "Vercel's privacy policy", href: VERCEL_PRIVACY_URL },
    paragraphs: [
      `When a Premium player or an officer opens a coach review, the numbers of one fight for your spec and boss (spec, boss, raid size and what happened in that fight) are sent through Vercel's AI Gateway to ${provider}, which writes the review.`,
      'It carries no character name, account, guild, realm or Battle.net data.',
      `Vercel and ${provider} process that data for this purpose only. ${provider} does not train on it and keeps it only for the period its terms state.`,
    ],
  };
}
