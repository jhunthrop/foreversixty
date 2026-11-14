// web/src/components/RemoveCharacter.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { validCharacterName } from '../lib/directory/opt-out-api';
import { optOutCopy } from '../lib/directory/opt-out-copy';
import RemoveCharacter from './RemoveCharacter.svelte';

describe('RemoveCharacter', () => {
  it('renders the form and both ways to remove a character, signed out', () => {
    const { body } = render(RemoveCharacter);
    expect(body).toContain('data-testid="remove-character"');
    expect(body).toContain('data-state="form"');
    expect(body).toContain('data-testid="remove-name"');
    expect(body).toContain('data-testid="remove-battlenet"');
    expect(body).toContain('data-testid="remove-get-code"');
    expect(body).toContain(optOutCopy.title);
  });

  it('names Blizzard as the source', () => {
    expect(optOutCopy.lead).toContain('from Blizzard');
  });
});

describe('validCharacterName', () => {
  it('takes a first name and surname, or a first name alone', () => {
    expect(validCharacterName('Vela Thornbrook')).toBe(true);
    expect(validCharacterName('  Vela  ')).toBe(true);
    expect(validCharacterName('Vëla Thornbrook')).toBe(true);
  });

  it('refuses anything a character cannot be named', () => {
    expect(validCharacterName('')).toBe(false);
    expect(validCharacterName('V')).toBe(false);
    expect(validCharacterName('Vela Thornbrook Third')).toBe(false);
    expect(validCharacterName('Vela<b>')).toBe(false);
  });
});
