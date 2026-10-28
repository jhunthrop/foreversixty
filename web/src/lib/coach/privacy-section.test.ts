// web/src/lib/coach/privacy-section.test.ts
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../billing/plans-page-copy';
import { coachIsShipped, coachPrivacySection } from './privacy-section';
import { COACH_MODEL_PROVIDER } from './provider';

const HELD = { featuresHeld: { coach: 'held' } };
const SHIPPED = {};

function sectionText(premium: object, provider?: string): string {
  const section = coachPrivacySection(premium, provider);
  return section ? [section.heading, ...section.paragraphs].join('\n') : '';
}

describe('the privacy page Log coach section', () => {
  it('names Vercel and the provider constant, and says what is never sent', () => {
    const text = sectionText(SHIPPED);
    expect(text).toContain('Log coach');
    expect(text).toContain('Vercel');
    expect(text).toContain(COACH_MODEL_PROVIDER);
    expect(text).toMatch(/no character name, account, guild, realm or Battle\.net data/);
    expect(text).toMatch(/the numbers of one fight \(your spec, the boss, the raid size/);
    expect(text).toMatch(/does not train on it/);
  });

  it('follows the constant: swapping the provider changes every mention', () => {
    const text = sectionText(SHIPPED, 'Kimi');
    expect(text).toContain('Kimi');
    expect(text).not.toContain(COACH_MODEL_PROVIDER);
  });

  it('renders nothing while the plans page holds the coach line', () => {
    expect(coachPrivacySection(HELD)).toBeNull();
    expect(coachIsShipped(HELD)).toBe(false);
  });

  it('renders once the held line is gone', () => {
    expect(coachIsShipped(SHIPPED)).toBe(true);
    expect(coachPrivacySection(SHIPPED)).not.toBeNull();
  });

  it('is held today, on the same flag the plans page uses', () => {
    expect('featuresHeld' in plansPageCopy.premium).toBe(true);
    expect(coachPrivacySection()).toBeNull();
  });
});
