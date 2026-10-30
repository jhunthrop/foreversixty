// PageHead and its static twin (site alignment spec 2026-10-29 §3.1, S2).
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { JSDOM } from 'jsdom';
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { PAGE_HEAD } from '../../lib/ui/page-head';
import PageHeadAstro from './PageHead.astro';
import PageHead from './PageHead.svelte';

/** Node.DOCUMENT_POSITION_FOLLOWING (no DOM globals in this environment). */
const FOLLOWING = 4;

const snippet = (html: string) => createRawSnippet(() => ({ render: () => html }));

function head(props: Record<string, unknown> = {}): Document {
  const body = render(PageHead, { props: { eyebrow: 'Account', title: 'Your account', ...props } }).body;
  return new JSDOM(body).window.document;
}

async function astroHead(props: Record<string, unknown> = {}, slots: Record<string, string> = {}) {
  const container = await AstroContainer.create();
  const html = await container.renderToString(PageHeadAstro, {
    props: { eyebrow: 'Guides', title: 'Class guides', ...props },
    slots,
  });
  return new JSDOM(html).window.document;
}

describe('PageHead (Svelte)', () => {
  it('draws the gold-ruled eyebrow over the one h1, the account head exactly', () => {
    const doc = head();
    const eyebrow = doc.querySelector('p');
    expect(eyebrow?.className).toBe(PAGE_HEAD.eyebrow);
    expect(eyebrow?.textContent).toBe('Account');
    expect(eyebrow?.querySelector('span')?.getAttribute('aria-hidden')).toBe('true');
    const h1s = doc.querySelectorAll('h1');
    expect(h1s).toHaveLength(1);
    expect(h1s[0].className).toBe('heading-page');
    expect(h1s[0].textContent).toBe('Your account');
  });

  it('draws no band, no card and no optional line when given none', () => {
    const doc = head();
    const html = doc.body.innerHTML;
    expect(html).not.toMatch(/surface-band|card-surface|#26405c/);
    expect(doc.querySelector('[data-testid=page-head-job]')).toBeNull();
    expect(doc.querySelector('[data-testid=page-head-meta]')).toBeNull();
    expect(doc.querySelector('[data-testid=page-head-aside]')).toBeNull();
  });

  it('adds the job line and the meta line under the h1, in that order', () => {
    const doc = head({ job: 'Play your class better.', meta: snippet('<span>Updated 29 Oct</span>') });
    const job = doc.querySelector('[data-testid=page-head-job]');
    const meta = doc.querySelector('[data-testid=page-head-meta]');
    expect(job?.className).toBe(PAGE_HEAD.job);
    expect(meta?.textContent).toBe('Updated 29 Oct');
    const h1 = doc.querySelector('h1')!;
    expect(h1.compareDocumentPosition(job!) & FOLLOWING).toBeTruthy();
    expect(job!.compareDocumentPosition(meta!) & FOLLOWING).toBeTruthy();
  });

  it('puts the crest at the title block left and colours the h1', () => {
    const doc = head({
      lead: snippet('<img data-testid="crest" alt="" width="64" height="64">'),
      titleColor: 'var(--color-class-warrior)',
      titleClass: 'class-title',
    });
    const crest = doc.querySelector('[data-testid=crest]')!;
    expect(crest.parentElement?.className).toBe(PAGE_HEAD.withLead);
    const h1 = doc.querySelector('h1')!;
    expect(h1.className).toBe('heading-page class-title');
    expect(h1.getAttribute('style')).toContain('color: var(--color-class-warrior)');
  });

  it('renders the aside in the head row and the below row under it', () => {
    const doc = head({
      aside: snippet('<button>Sign out</button>'),
      below: snippet('<nav data-testid="tabs">Tabs</nav>'),
      testid: 'head',
    });
    const root = doc.querySelector('[data-testid=head]')!;
    expect(root.className).toBe(PAGE_HEAD.root);
    const row = root.firstElementChild!;
    expect(row.className).toBe(PAGE_HEAD.row);
    expect(row.querySelector('[data-testid=page-head-aside] button')?.textContent).toBe('Sign out');
    expect(root.lastElementChild?.getAttribute('data-testid')).toBe('tabs');
  });
});

describe('PageHead (Astro twin)', () => {
  it('draws the same eyebrow, h1 and row as the Svelte head', async () => {
    const doc = await astroHead({ job: 'Every spec, written by players.' });
    expect(doc.querySelector('p')?.className).toBe(PAGE_HEAD.eyebrow);
    expect(doc.querySelector('h1')?.className).toBe('heading-page');
    expect(doc.querySelector('h1')?.textContent?.trim()).toBe('Class guides');
    expect(doc.querySelector('[data-testid=page-head-job]')?.textContent?.trim()).toBe(
      'Every spec, written by players.',
    );
    expect(doc.querySelector('[data-testid=page-head-aside]')).toBeNull();
    expect(doc.body.innerHTML).not.toMatch(/surface-band|card-surface|<script/);
  });

  it('fills the lead, meta, aside and below slots', async () => {
    const doc = await astroHead(
      { titleColor: 'var(--color-class-mage)' },
      {
        lead: '<img data-testid="crest" alt="">',
        meta: 'Updated 29 Oct',
        aside: '<button>Yearly</button>',
        below: '<nav data-testid="tabs"></nav>',
      },
    );
    expect(doc.querySelector('[data-testid=crest]')?.parentElement?.className).toBe(PAGE_HEAD.withLead);
    expect(doc.querySelector('[data-testid=page-head-meta]')?.textContent?.trim()).toBe('Updated 29 Oct');
    expect(doc.querySelector('[data-testid=page-head-aside] button')).not.toBeNull();
    expect(doc.querySelector('[data-testid=tabs]')).not.toBeNull();
    expect(doc.querySelector('h1')?.getAttribute('style')).toBe('color:var(--color-class-mage)');
  });
});
