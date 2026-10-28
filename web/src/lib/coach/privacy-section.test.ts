// web/src/lib/coach/privacy-section.test.ts
import { describe, expect, it } from 'vitest';
import {
  coachPrivacySection,
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
    expect(text()).toMatch(/spec, class, role, boss, raid size, kill or wipe, duration, level, spell names/);
    expect(text()).toMatch(/no character name, account, guild, realm or Battle\.net data/);
    expect(text()).toMatch(/we do not control that/);
    expect(text()).toMatch(/fall back to the log-only text/);
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
