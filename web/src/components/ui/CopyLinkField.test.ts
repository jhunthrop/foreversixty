// The vitest config hands mount() Svelte's server entry, so clicks cannot be simulated here:
// the copy/Copied timing is covered by lib/ui/copy-feedback.test.ts, and this file proves
// the rendered markup the component owns (the server render has no navigator.share).
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { uiCopy } from '../../lib/ui/copy';
import CopyLinkField from './CopyLinkField.svelte';

const URL_VALUE = 'https://foreversixty.gg/join/abc123';

function renderField(withNote = false): string {
  const note = createRawSnippet(() => ({ render: () => '<span data-testid="note">Shown once.</span>' }));
  const { body } = render(CopyLinkField, {
    props: {
      url: URL_VALUE,
      label: 'Invite link',
      shareTitle: 'Join',
      shareText: 'Join us',
      testid: 'guild-invite',
      ...(withNote ? { note } : {}),
    },
  });
  return body;
}

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
