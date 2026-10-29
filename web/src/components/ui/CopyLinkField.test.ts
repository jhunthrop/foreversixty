// The vitest config hands mount() Svelte's server entry, so clicks cannot be simulated here:
// the copy/Copied timing is covered by lib/ui/copy-feedback.test.ts, and this file proves
// the rendered markup the component owns (the server render has no navigator.share).
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { uiCopy } from '../../lib/ui/copy';
import CopyLinkField from './CopyLinkField.svelte';

const URL_VALUE = 'https://foreversixty.gg/join/abc123';

interface Variant {
  labelVisible?: boolean;
  onCard?: boolean;
  phoneFullWidth?: boolean;
}

function renderField(withNote = false, variant: Variant = {}): string {
  const note = createRawSnippet(() => ({ render: () => '<span data-testid="note">Shown once.</span>' }));
  const { body } = render(CopyLinkField, {
    props: {
      url: URL_VALUE,
      label: 'Invite link',
      shareTitle: 'Join',
      shareText: 'Join us',
      testid: 'guild-invite',
      ...(withNote ? { note } : {}),
      ...variant,
    },
  });
  return body;
}

describe('CopyLinkField on a card (spec §5.5, gate fixes G3 to G5)', () => {
  it('the default is the invite render: visible label, raised ground, fitted buttons', () => {
    const body = renderField();
    expect(body).toContain('<span class="quiet-label">Invite link</span>');
    expect(body).toMatch(/<input[^>]*bg-raised/);
    expect(body).toMatch(/<div class="flex flex-wrap items-center gap-3">/);
    expect(body).toMatch(/w-fit px-4"[^>]*data-testid="guild-invite-copy"/);
  });

  it('labelVisible false keeps the label for screen readers only', () => {
    expect(renderField(false, { labelVisible: false })).toContain('<span class="sr-only">Invite link</span>');
  });

  it('onCard puts the field on the page ground', () => {
    const body = renderField(false, { onCard: true });
    expect(body).toMatch(/<input[^>]*bg-bg/);
    expect(body).not.toMatch(/<input[^>]*bg-raised/);
  });

  it('phoneFullWidth fills the row below 768 and fits the buttons from 768', () => {
    const body = renderField(false, { phoneFullWidth: true });
    expect(body).toContain('grid gap-3 md:flex md:flex-wrap md:items-center grid-cols-1');
    expect(body).toMatch(/w-full justify-center md:w-fit px-4"[^>]*data-testid="guild-invite-copy"/);
  });
});

describe('CopyLinkField', () => {
  it('renders the label and a read-only 44 px mono field holding the whole url', () => {
    const body = renderField();
    expect(body).toContain('>Invite link<');
    expect(body).toMatch(/<input[^>]*readonly/);
    expect(body).toContain(`value="${URL_VALUE}"`);
    expect(body).toContain('h-11');
    expect(body).toContain('font-mono');
    expect(body).toContain('data-testid="guild-invite-url"');
  });

  it('renders the Copy link button with the shared words', () => {
    const body = renderField();
    expect(body).toContain('data-testid="guild-invite-copy"');
    expect(body).toContain(uiCopy.copyLink);
    expect(uiCopy.copyLink).toBe('Copy link');
    expect(uiCopy.copiedLink).toBe('Copied');
    expect(uiCopy.copyLinkFailed).toBe('Could not copy. Select the link and copy it.');
  });

  it('shows no Share button and no failure line without navigator.share or a failed copy', () => {
    const body = renderField();
    expect(body).not.toContain('guild-invite-share');
    expect(body).not.toContain('guild-invite-copy-error');
    expect(body).not.toContain(uiCopy.copyLinkFailed);
  });

  it('wraps in the -token block and renders the optional note in the button row', () => {
    expect(renderField()).toContain('data-testid="guild-invite-token"');
    expect(renderField()).not.toContain('data-testid="note"');
    expect(renderField(true)).toContain('data-testid="note"');
  });
});
