// web/src/components/premium/PremiumCoachSample.test.ts
// Polish spec 2026-10-28 section 3.7 checks 2 and 3: the sample reads as a sample, shows each review
// string once, and leaves out the drill, the link sentences and the unchecked lines.
import { readFileSync } from 'node:fs';
import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import { coachSampleView, loadCoachSampleAnswer } from '../../lib/billing/coach-sample';
import { plansPageCopy } from '../../lib/billing/plans-page-copy';
import { createPageContainer } from '../../test-support/page-container';
import PremiumCoachSample from './PremiumCoachSample.astro';

const answer = loadCoachSampleAnswer();
const view = coachSampleView(answer, (id) => (id === null ? null : `/data/b/icons/${id}.webp`));
const source = readFileSync(new URL('./PremiumCoachSample.astro', import.meta.url), 'utf8');

async function render(): Promise<Document> {
  const container = await createPageContainer();
  const html = await container.renderToString(PremiumCoachSample, { props: { view } });
  return new JSDOM(html).window.document;
}

const occurrences = (text: string, needle: string): number => text.split(needle).length - 1;

describe('PremiumCoachSample', () => {
  it('opens with the Sample tag and ends with the pinned caption', async () => {
    const doc = await render();
    const card = doc.querySelector('[data-testid=coach-sample]');
    expect(card?.getAttribute('aria-label')).toBe(plansPageCopy.coachSample.ariaLabel);
    expect(card?.textContent?.trim().startsWith(plansPageCopy.coachSample.tag)).toBe(true);
    expect(doc.querySelector('[data-testid=coach-sample-tag]')?.textContent).toBe('Sample');
    expect(doc.querySelector('[data-testid=coach-sample-caption]')?.textContent).toBe(
      plansPageCopy.coachSample.caption,
    );
    expect(plansPageCopy.coachSample.caption).toMatch(/^A sample review\./);
  });

  it('shows each review string once', async () => {
    const text = (await render()).body.textContent ?? '';
    const { review, persona } = answer;
    const strings = [
      persona.label,
      review.opening,
      review.fixFirst?.title ?? '',
      review.fixFirst?.why ?? '',
      ...review.thenFix.flatMap((item) => [item.title, item.why]),
      ...review.doingWell.map((item) => item.text),
    ];
    for (const needle of strings) {
      expect(needle).not.toBe('');
      expect(occurrences(text, needle), needle).toBe(1);
    }
  });

  it('leaves out the drill, the link sentences and the unchecked lines', async () => {
    const text = (await render()).body.textContent ?? '';
    const { review } = answer;
    const left = [
      review.fixFirst?.drillText ?? '',
      ...(review.fixFirst?.links ?? []).map((link) => link.sentence ?? ''),
      ...review.thenFix.flatMap((item) => item.links.map((link) => link.sentence ?? '')),
      ...review.couldntCheck,
    ];
    for (const needle of left) {
      expect(needle).not.toBe('');
      expect(text).not.toContain(needle);
    }
  });

  it('draws the spell icons at their sizes and none for a spell without one', async () => {
    const doc = await render();
    const sizes = [...doc.querySelectorAll('img[src^="/data/"]')].map((img) => img.getAttribute('width'));
    expect(sizes).toEqual(['36', '20']);
    for (const img of doc.querySelectorAll('img[src^="/data/"]')) {
      expect(img.getAttribute('onerror')).toBe("this.style.visibility='hidden'");
    }
  });

  it('is static: no script and no client directive', () => {
    expect(source).not.toMatch(/<script|client:/);
  });
});
