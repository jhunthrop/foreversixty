// Card, CardHead and the static twin (site alignment spec 2026-10-29 §3.2 to §3.4, S1).
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { JSDOM } from 'jsdom';
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { CARD_BODY, CARD_FOOT, CARD_HEAD } from '../../lib/ui/card';
import { createPageContainer } from '../../test-support/page-container';
import CardAstro from './Card.astro';
import Card from './Card.svelte';

const snippet = (html: string) => createRawSnippet(() => ({ render: () => html }));
const children = snippet('<p data-testid="content">Body</p>');

function card(props: Record<string, unknown> = {}): Document {
  const body = render(Card, { props: { glyph: 'scale', title: 'Stat weights', children, ...props } }).body;
  return new JSDOM(body).window.document;
}

async function astroCard(props: Record<string, unknown> = {}, slots: Record<string, string> = {}) {
  const container: AstroContainer = await createPageContainer();
  const html = await container.renderToString(CardAstro, {
    props: { glyph: 'list', title: 'The list', ...props },
    slots: { default: '<p data-testid="content">Body</p>', ...slots },
  });
  return new JSDOM(html).window.document;
}

const section = (doc: Document): Element => {
  const el = doc.querySelector('section');
  if (el === null) throw new Error('no section');
  return el;
};

describe('Card (Svelte)', () => {
  it('is a section on the primary card surface, clipped, with a 36 px tile and an uppercase h2', () => {
    const doc = card({ testid: 'weights', id: 'weights' });
    const root = section(doc);
    expect(root.className).toBe('card-surface overflow-hidden');
    expect(root.id).toBe('weights');
    expect(root.getAttribute('data-testid')).toBe('weights');
    const head = root.firstElementChild!;
    expect(head.className).toBe(CARD_HEAD.root);
    expect(head.querySelector('[data-testid=icon-tile]')?.getAttribute('style')).toContain('width: 36px');
    const h2 = root.querySelector('h2');
    expect(h2?.className).toBe('heading-plan uppercase');
    expect(h2?.textContent).toBe('Stat weights');
  });

  it('drops the gold top edge when quiet', () => {
    expect(section(card({ quiet: true })).className).toBe('card-surface-quiet overflow-hidden');
  });

  it.each(['text', 'table', 'rows'] as const)('wraps a %s body in its padding', (body) => {
    const wrapper = card({ body }).querySelector('[data-testid=content]')?.parentElement;
    expect(wrapper?.className).toBe(CARD_BODY[body]);
    expect(wrapper?.getAttribute('data-card-body')).toBe(body);
  });

  it('renders a none body straight into the card, full-bleed', () => {
    const doc = card({ body: 'none' });
    expect(doc.querySelector('[data-testid=content]')?.parentElement?.tagName).toBe('SECTION');
  });

  it('draws the compact head: 48 px, a 28 px tile, the 18 px card title, no micro-label', () => {
    const doc = card({ compact: true, label: 'Ignored', meta: '20 to 29' });
    const head = section(doc).firstElementChild!;
    expect(head.className).toBe(CARD_HEAD.compact);
    expect(head.getAttribute('data-compact')).toBe('true');
    expect(head.querySelector('[data-testid=icon-tile]')?.getAttribute('style')).toContain('width: 28px');
    expect(doc.querySelector('h2')?.className).toBe('section-title leading-6');
    expect(head.textContent).not.toContain('Ignored');
    expect(doc.querySelector('[data-testid=card-head-meta]')?.textContent).toBe('20 to 29');
  });

  it('keeps the title in its own case under a micro-label', () => {
    const doc = card({ glyph: 'crown', label: 'Plan', title: 'Premium' });
    expect(doc.querySelector('h2')?.className).toBe('heading-plan');
    expect(doc.body.textContent).toContain('Plan');
  });

  it('heads at h3 inside a section with its own h2', () => {
    const doc = card({ headingLevel: 3 });
    expect(doc.querySelector('h2')).toBeNull();
    expect(doc.querySelector('h3')?.textContent).toBe('Stat weights');
  });

  it('puts the ActionLink in the right slot, a real link with the chevron', () => {
    const doc = card({ action: { label: 'See all', href: '/reports' } });
    const link = doc.querySelector('a');
    expect(link?.getAttribute('href')).toBe('/reports');
    expect(link?.textContent).toBe('See all');
    expect(link?.parentElement?.className).toBe(CARD_HEAD.aside);
    expect(link?.querySelector('[data-glyph=chevron]')).not.toBeNull();
  });

  it('replaces the tile with the lead (an item icon) and renders aside and foot', () => {
    const doc = card({
      lead: snippet('<img data-testid="item-icon" alt="" width="36" height="36">'),
      aside: snippet('<span data-testid="tag">Connected</span>'),
      foot: snippet('<button>Manage billing</button>'),
    });
    expect(doc.querySelector('[data-testid=icon-tile]')).toBeNull();
    expect(doc.querySelector('[data-testid=item-icon]')).not.toBeNull();
    expect(doc.querySelector('[data-testid=tag]')?.parentElement?.className).toBe(CARD_HEAD.aside);
    const foot = doc.querySelector('[data-card-foot]');
    expect(foot?.className).toBe(CARD_FOOT);
    expect(foot?.querySelector('button')?.textContent).toBe('Manage billing');
  });

  it('has no tile and no right slot when given neither (a prose aside)', () => {
    const doc = card({ glyph: undefined, quiet: true });
    expect(doc.querySelector('[data-testid=icon-tile]')).toBeNull();
    expect(doc.querySelector(`[class="${CARD_HEAD.aside}"]`)).toBeNull();
  });

  it('keeps the tile decorative, so the heading is the only name', () => {
    expect(card().querySelector('[data-testid=icon-tile]')?.getAttribute('aria-hidden')).toBe('true');
  });
});

describe('Card (Astro twin)', () => {
  it('draws the same surface, head and text body as the Svelte card', async () => {
    const doc = await astroCard({ testid: 'list', meta: 'Hover or tap an item for its stats' });
    const root = section(doc);
    expect(root.className).toBe('card-surface overflow-hidden');
    expect(root.firstElementChild?.className).toBe(CARD_HEAD.root);
    expect(doc.querySelector('h2')?.className).toBe('heading-plan uppercase');
    expect(doc.querySelector('[data-testid=card-head-meta]')?.textContent?.trim()).toBe(
      'Hover or tap an item for its stats',
    );
    expect(doc.querySelector('[data-testid=content]')?.parentElement?.className).toBe(CARD_BODY.text);
    expect(doc.body.innerHTML).not.toMatch(/<script|astro-island/);
  });

  it('draws the compact head, a quiet surface, an h3 and a table body', async () => {
    const doc = await astroCard({
      compact: true,
      quiet: true,
      headingLevel: 3,
      body: 'table',
      class: 'col-span-6',
    });
    const root = section(doc);
    expect(root.className).toBe('card-surface-quiet overflow-hidden col-span-6');
    expect(root.firstElementChild?.className).toBe(CARD_HEAD.compact);
    expect(doc.querySelector('h3')?.className).toBe('section-title leading-6');
    expect(doc.querySelector('[data-testid=icon-tile]')?.getAttribute('style')).toContain('width: 28px');
    expect(doc.querySelector('[data-testid=content]')?.parentElement?.className).toBe(CARD_BODY.table);
  });

  it('fills the lead, aside and foot slots and renders the ActionLink', async () => {
    const doc = await astroCard(
      { action: { label: 'See all', href: '/guild' } },
      {
        lead: '<img data-testid="item-icon" alt="">',
        aside: '<span data-testid="tag">3</span>',
        foot: '<button>Go</button>',
      },
    );
    expect(doc.querySelector('[data-testid=icon-tile]')).toBeNull();
    expect(doc.querySelector('[data-testid=item-icon]')).not.toBeNull();
    expect(doc.querySelector('a[href="/guild"]')?.textContent?.trim()).toBe('See all');
    expect(doc.querySelector('[data-testid=tag]')).not.toBeNull();
    expect(doc.querySelector('[data-card-foot] button')).not.toBeNull();
  });

  it('renders a none body straight into the section', async () => {
    const doc = await astroCard({ body: 'none' });
    expect(doc.querySelector('[data-testid=content]')?.parentElement?.tagName).toBe('SECTION');
  });
});
