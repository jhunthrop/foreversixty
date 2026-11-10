// web/src/lib/coach/privacy-section.test.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import {
  COACH_PRIVACY_ID,
  coachPrivacyHtml,
  coachPrivacySection,
  withCoachSection,
  PROVIDER_PRIVACY_URL,
  PROVIDER_TERMS_URL,
  VERCEL_GATEWAY_URL,
} from './privacy-section';
import { COACH_MODEL_PROVIDER } from './provider';

const text = (provider?: string): string => coachPrivacySection(provider).paragraphs.join('\n');

describe('the privacy page Log coach section', () => {
  it('discloses what is sent, to whom, and what is never sent', () => {
    expect(text()).toContain('Vercel AI Gateway, which keeps nothing');
    expect(text()).toContain(COACH_MODEL_PROVIDER);
    expect(text()).toContain(
      "one fight's numbers (spec, class, role, boss, raid size, kill or wipe, duration, level and the spells cast)",
    );
    expect(text().match(/numbers/g)).toHaveLength(1);
    expect(text()).toMatch(/no character name, account, guild, realm or Battle\.net data/);
    expect(text()).toMatch(/we do not control that/);
    expect(text()).toMatch(/fall back to the log-only text/);
  });

  it('says plainly what anonymising does and does not stop (api/internal/coach/access.go)', () => {
    expect(text()).toContain(
      "Anonymising your account in its settings does not turn this off: a review is written when you, one of your guild's officers or a site moderator opens it, and anonymising stops only the moderator.",
    );
  });

  it('promises nothing about the provider training or retention', () => {
    expect(text()).not.toMatch(/does not train|never train|deleted|not keep/i);
  });

  it('links the gateway page, the provider terms and the provider privacy policy', () => {
    expect(coachPrivacySection().links.map((l) => l.href)).toEqual([
      VERCEL_GATEWAY_URL,
      PROVIDER_TERMS_URL,
      PROVIDER_PRIVACY_URL,
    ]);
  });

  it('follows the provider constant in the text and link labels', () => {
    const section = coachPrivacySection('Kimi');
    expect(section.paragraphs.join('')).not.toContain(COACH_MODEL_PROVIDER);
    expect(section.links[1]?.label).toContain('Kimi');
  });
});

describe('the privacy page body with its Log coach section', () => {
  const HTML = '<h2 id="stripe">Stripe</h2><p>pay</p><h2 id="cookies">Cookies</h2><p>c</p>';

  it('/privacy shows Log coach, anchored, between Stripe and Cookies', () => {
    const out = withCoachSection('privacy', HTML, coachPrivacySection());
    expect(out).toContain(`<h2 id="${COACH_PRIVACY_ID}">Log coach</h2>`);
    expect(out.indexOf('Stripe')).toBeLessThan(out.indexOf('Log coach'));
    expect(out.indexOf('Log coach')).toBeLessThan(out.indexOf('Cookies'));
  });

  it.each(['terms', 'refunds', 'about'])('/%s never carries the section', (slug) => {
    expect(withCoachSection(slug, HTML, coachPrivacySection())).toBe(HTML);
  });

  it('escapes the section into HTML', () => {
    const out = coachPrivacyHtml({
      heading: 'A & B',
      paragraphs: ['<x>'],
      links: [{ label: 'L', href: 'https://a.b/?q="1"' }],
    });
    expect(out).toBe(
      '<h2 id="log-coach">A &amp; B</h2><p>&lt;x&gt;</p><ul><li><a href="https://a.b/?q=&quot;1&quot;">L</a></li></ul>',
    );
  });

  it('privacy.md has the headings the section is placed between', () => {
    const md = readFileSync(path.resolve(import.meta.dirname, '../../content/pages/privacy.md'), 'utf8');
    expect(md).toMatch(/^## Stripe$/m);
    expect(md).toMatch(/^## Cookies$/m);
  });
});
