// web/src/lib/coach/privacy-section.ts
// The privacy page's "Log coach" section (log coach spec §9: what leaves our systems). It
// discloses; it promises nothing the provider does not (spec §19, owner decision 4).
import { COACH_MODEL_PROVIDER } from './provider';

export const COACH_PRIVACY_HEADING = 'Log coach';
/** The section's anchor (pages spec 2026-11-08 gap PG3): the contents rail links it like any heading. */
export const COACH_PRIVACY_ID = 'log-coach';
export const PRIVACY_SLUG = 'privacy';
// The section sits after "Stripe", so it goes in front of the next heading.
export const COACH_SECTION_BEFORE = '<h2 id="cookies">';
export const VERCEL_GATEWAY_URL = 'https://vercel.com/docs/ai-gateway/security-and-compliance/zdr';
export const PROVIDER_TERMS_URL =
  'https://cdn.deepseek.com/policies/en-US/deepseek-open-platform-terms-of-service.html';
export const PROVIDER_PRIVACY_URL = 'https://cdn.deepseek.com/policies/en-US/deepseek-privacy-policy.html';

export interface CoachPrivacyLink {
  readonly label: string;
  readonly href: string;
}

export interface CoachPrivacySection {
  readonly heading: string;
  readonly paragraphs: readonly string[];
  readonly links: readonly CoachPrivacyLink[];
}

export function coachPrivacySection(provider: string = COACH_MODEL_PROVIDER): CoachPrivacySection {
  return {
    heading: COACH_PRIVACY_HEADING,
    links: [
      { label: 'Vercel AI Gateway data retention', href: VERCEL_GATEWAY_URL },
      { label: `${provider} Open Platform terms`, href: PROVIDER_TERMS_URL },
      { label: `${provider} privacy policy`, href: PROVIDER_PRIVACY_URL },
    ],
    paragraphs: [
      `When a Premium player or an officer opens a coach review, one fight's numbers — spec, class, role, boss, raid size, kill or wipe, duration, level and the spells cast — are sent through the Vercel AI Gateway, which keeps nothing, to ${provider}, which writes the review. It carries no character name, account, guild, realm or Battle.net data. What ${provider} keeps, and whether it trains on it, is set by its Open Platform terms and privacy policy, and we do not control that. If the model is switched off, reviews fall back to the log-only text. Anonymising your account in its settings does not turn this off: a review is written when you, one of your guild's officers or a site moderator opens it, and anonymising stops only the moderator.`,
    ],
  };
}

export function splitBeforeHeading(html: string, heading: string): [string, string] {
  const at = html.indexOf(heading);
  if (at < 0)
    throw new Error(`the privacy page has no ${heading} heading to place the Log coach section before`);
  return [html.slice(0, at), html.slice(at)];
}

const escapeHtml = (text: string): string =>
  text.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');

/** The section as HTML, in the rendered markdown's own shape (an `h2` with its id, paragraphs, a list). */
export function coachPrivacyHtml(section: CoachPrivacySection): string {
  const paragraphs = section.paragraphs.map((text) => `<p>${escapeHtml(text)}</p>`).join('');
  const links = section.links
    .map((link) => `<li><a href="${escapeHtml(link.href)}">${escapeHtml(link.label)}</a></li>`)
    .join('');
  return `<h2 id="${COACH_PRIVACY_ID}">${escapeHtml(section.heading)}</h2>${paragraphs}<ul>${links}</ul>`;
}

/** A content page's body: the privacy page gains its Log coach section before Cookies. */
export function withCoachSection(slug: string, html: string, section: CoachPrivacySection): string {
  if (slug !== PRIVACY_SLUG) return html;
  const [head, tail] = splitBeforeHeading(html, COACH_SECTION_BEFORE);
  return head + coachPrivacyHtml(section) + tail;
}
