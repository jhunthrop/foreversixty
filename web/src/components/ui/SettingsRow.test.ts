import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SettingsRow from './SettingsRow.svelte';

function row(props: Record<string, unknown>): string {
  return render(SettingsRow, { props: { label: 'Server sims', ...props } }).body;
}

const first = (body: string, tagName: string): string =>
  new RegExp(`<${tagName}\\b[^>]*>`).exec(body)?.[0] ?? '';

describe('SettingsRow', () => {
  it('is a plain div with no hover or chevron when it neither links nor acts', () => {
    const body = row({ value: 'On' });
    expect(body).toMatch(/^<!--\[-->?.*?<div class="border-line-soft/s);
    expect(body).not.toContain('hover:bg-row-hover');
    expect(body).not.toContain('data-glyph="chevron"');
  });

  it('is an anchor with a chevron and the row hover when it has an href', () => {
    const body = row({ value: 'On', href: '/sim' });
    expect(first(body, 'a')).toContain('href="/sim"');
    expect(body).toContain('hover:bg-row-hover');
    expect(body).toContain('focus-visible:-outline-offset-2');
    expect(body).toContain('data-glyph="chevron"');
  });

  it('opens an external link in a new tab with noopener', () => {
    const tagText = first(row({ href: 'https://example.com', external: true }), 'a');
    expect(tagText).toContain('target="_blank"');
    expect(tagText).toContain('rel="noopener"');
  });

  it('is a button when it has onclick', () => {
    const tagText = first(row({ onclick: () => {} }), 'button');
    expect(tagText).toContain('type="button"');
  });

  it('draws the label at 15 px 600 and the sub at 13 px muted', () => {
    const body = row({ sub: 'Off, never posts.' });
    expect(body).toContain('text-[15px] font-semibold');
    expect(body).toMatch(/text-muted text-\[13px\][^>]*>Off, never posts\./);
  });

  it('draws an icon tile at 32 px when it has a glyph', () => {
    const body = row({ glyph: 'server' });
    expect(body).toContain('data-glyph="server"');
    expect(body).toContain('width: 32px');
  });

  it('keeps a 56 px minimum height with 20 px side padding and a top rule', () => {
    const body = row({});
    expect(body).toContain('min-h-14');
    expect(body).toContain('px-5');
    expect(body).toContain('border-t');
  });

  it('moves the value under the label below 480 px', () => {
    const body = row({ value: 'Open on your main character' });
    expect(body).toContain('max-[479px]:order-last');
    expect(body).toContain('max-[479px]:basis-full');
  });

  it('indents a nested row to 70 px', () => {
    expect(row({ nested: true })).toContain('padding-left: 70px');
    expect(row({})).not.toContain('padding-left');
  });

  it('a nested row with a tile indents the tile to 46 px', () => {
    expect(row({ nested: true, glyph: 'lock' })).toContain('padding-left: 46px');
  });

  it('muted value: --muted ink and no chevron even with a link', () => {
    const body = row({ value: 'Premium only', muted: true, href: '/premium' });
    expect(body).toMatch(/text-muted"[^>]*data-testid="settings-row-value">Premium only/);
    expect(body).not.toContain('data-glyph="chevron"');
  });

  it('disabled: muted label, not a link, no chevron, aria-disabled', () => {
    const body = row({ disabled: true, href: '/sim', value: 'On' });
    expect(body).not.toContain('href=');
    expect(body).toContain('aria-disabled="true"');
    expect(body).toContain('text-muted');
    expect(body).not.toContain('data-glyph="chevron"');
  });

  it('loading: a 96 x 12 skeleton bar instead of the value, marked busy', () => {
    const body = row({ loading: true, value: 'On' });
    expect(body).toContain('data-testid="settings-row-loading"');
    expect(body).toContain('h-3 w-24');
    expect(body).toContain('aria-busy="true"');
    expect(body).not.toContain('>On<');
  });

  it('empty: a muted dash', () => {
    expect(row({ empty: true })).toMatch(/text-muted[^>]*data-testid="settings-row-value">—</);
  });

  it('error: the failure colour and the caller words, no chevron', () => {
    const body = row({ error: 'Could not load', href: '/x' });
    expect(body).toMatch(/text-wipe[^>]*data-testid="settings-row-value">Could not load/);
    expect(body).not.toContain('data-glyph="chevron"');
  });

  it('hands its label id to a trailing control so a Switch can name itself by it', () => {
    const control = createRawSnippet((id: () => string) => ({
      render: () => `<button data-labelled="${id()}">c</button>`,
    }));
    const body = row({ control });
    const labelId = /<span id="([^"]+)"/.exec(body)?.[1];
    expect(labelId).toBeDefined();
    expect(body).toContain(`data-labelled="${labelId}"`);
  });

  it('can move a wide control under the text on a phone', () => {
    const control = createRawSnippet(() => ({ render: () => '<button>Open</button>' }));
    expect(row({ control, controlBelow: true })).toContain('max-[479px]:basis-full');
  });
});
