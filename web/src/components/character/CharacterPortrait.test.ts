// web/src/components/character/CharacterPortrait.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterPortrait from './CharacterPortrait.svelte';

describe('CharacterPortrait', () => {
  it('draws the circular ringed class crest, never an avatar or a letter square', () => {
    const { body } = render(CharacterPortrait, {
      props: {
        character: { class: 'warrior', avatar_url: '/a.jpg' } as { class: string },
        size: 'md',
        testid: 't',
      },
    });
    expect(body).toContain('data-testid="t-avatar"');
    expect(body).toContain('/icons/hd/crests/warrior.webp');
    expect(body).toContain('rounded-full');
    expect(body).not.toContain('/a.jpg');
    expect(body).not.toContain('rounded-[3px]');
    expect(body).not.toContain('data-testid="t-avatar-fallback"');
  });

  it('shows a neutral ringed disc for a character with no class on file', () => {
    const { body } = render(CharacterPortrait, {
      props: { character: {}, size: 'md', testid: 't' },
    });
    expect(body).toContain('data-testid="t-avatar-fallback"');
    expect(body).toContain('rounded-full');
    expect(body).not.toContain('data-testid="t-avatar"');
  });

  it('sizes the crest 28 / 36 / 44px for sm, md and lg and responsively for xl', () => {
    const at = (size: 'sm' | 'md' | 'lg' | 'xl') =>
      render(CharacterPortrait, { props: { character: { class: 'priest' }, size, testid: 't' } }).body;
    expect(at('sm')).toContain('width: 28px');
    expect(at('md')).toContain('width: 36px');
    expect(at('lg')).toContain('width: 44px');
    expect(at('xl')).toContain('lg:h-[84px]');
  });
});
