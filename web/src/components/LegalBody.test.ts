// web/src/components/LegalBody.test.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { coachPrivacySection } from '../lib/coach/privacy-section';
import LegalBody from './LegalBody.astro';

const HTML = '<h2 id="stripe">Stripe</h2><p>pay</p><h2 id="cookies">Cookies</h2><p>c</p>';
const SHIPPED = coachPrivacySection({});
const HELD = coachPrivacySection({ featuresHeld: {} });

async function render(slug: string, section: typeof SHIPPED): Promise<string> {
  const c = await AstroContainer.create();
  return c.renderToString(LegalBody, { props: { slug, html: HTML, section } });
}

describe('the legal page body', () => {
  it('/privacy shows Log coach between Stripe and Cookies once the coach ships', async () => {
    const out = await render('privacy', SHIPPED);
    expect(out).toContain('<h2>Log coach</h2>');
    expect(out.indexOf('Stripe')).toBeLessThan(out.indexOf('Log coach'));
    expect(out.indexOf('Log coach')).toBeLessThan(out.indexOf('Cookies'));
  });

  it('/privacy omits it while the coach is held', async () => {
    const out = await render('privacy', HELD);
    expect(out).not.toContain('Log coach');
    expect(out).toContain('Cookies');
  });

  it.each(['terms', 'refunds', 'about'])('/%s never carries the section', async (slug) => {
    const out = await render(slug, SHIPPED);
    expect(out).not.toContain('Log coach');
    expect(out).toBe(HTML);
  });

  it('privacy.md has the headings the section is placed between', () => {
    const md = readFileSync(path.resolve(import.meta.dirname, '../content/pages/privacy.md'), 'utf8');
    expect(md).toMatch(/^## Stripe$/m);
    expect(md).toMatch(/^## Cookies$/m);
  });
});
