// web/src/components/EmailLinkAlternative.test.ts
// W.8: "or use an email link", one component for every sign-in surface; the words after "or"
// are the link, and the old dangling sentence is gone from all of them.
import { readFileSync } from 'node:fs';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SignInPrompt from './SignInPrompt.svelte';
import EmailLinkAlternative from './EmailLinkAlternative.svelte';

describe('EmailLinkAlternative', () => {
  it('reads "or use an email link" with the link on the last three words', () => {
    const { body } = render(EmailLinkAlternative, { props: {} });
    expect(body).toMatch(/or\s*<a class="[^"]*" href="\/login">use an email link<\/a>/);
    expect(body).not.toContain('Use an email link instead.');
  });

  it('is what the sign-in prompt shows under its Battle.net button', () => {
    const { body } = render(SignInPrompt, { props: { line: 'Sign in to see this.', testid: 'prompt' } });
    expect(body).toContain('data-testid="prompt-email-link"');
    expect(body.indexOf('Sign in with Battle.net')).toBeLessThan(body.indexOf('use an email link'));
    expect(body).not.toContain('Use an email link instead.');
  });

  it('is what the sim’s account card uses, not a sentence of its own', () => {
    const source = readFileSync(new URL('./sim/SourceSwitcher.svelte', import.meta.url), 'utf8');
    expect(source).toContain('<EmailLinkAlternative');
    expect(source).not.toContain('Use an email link instead');
  });
});
