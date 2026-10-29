import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import AccountTitle from './AccountTitle.svelte';

describe('AccountTitle', () => {
  const body = render(AccountTitle).body;

  it('draws the Account eyebrow over the page h1', () => {
    expect(body).toMatch(/>Account<\/p>/);
    expect(body).toMatch(/<h1 class="heading-page">Your account<\/h1>/);
    expect(body.indexOf('Account</p>')).toBeLessThan(body.indexOf('<h1'));
  });

  it('keeps the eyebrow rule decorative', () => {
    expect(body).toContain('aria-hidden="true"');
  });
});
