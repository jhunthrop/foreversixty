// web/src/components/premium/GuildRules.test.ts
// How the guild plans work (spec 2026-10-27-launch-pricing §B3.3, §B5, §B8 check 8).
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../../lib/billing/plans-page-copy';
import { createPageContainer } from '../../test-support/page-container';
import GuildRules from './GuildRules.astro';

async function renderRules(): Promise<Document> {
  const container = await createPageContainer();
  return new JSDOM(await container.renderToString(GuildRules)).window.document;
}

describe('GuildRules', () => {
  it('h3, Verified members first, Paying together after Size, no Battle.net', async () => {
    const doc = await renderRules();
    const heading = doc.querySelector('h3');
    expect(heading?.textContent?.trim()).toBe(plansPageCopy.howItWorks);
    expect(heading?.className.split(/\s+/)).toContain('heading-plan');
    expect(doc.querySelector('h2')).toBeNull();
    const terms = [...doc.querySelectorAll('[data-rule-term]')].map((n) => n.textContent?.trim());
    expect(terms).toEqual(plansPageCopy.rules.map((rule) => rule.term));
    expect(terms[0]).toBe('Verified members');
    expect(terms[terms.indexOf('Size') + 1]).toBe('Paying together');
    expect(doc.body.textContent).not.toContain('Battle.net');
  });
});
