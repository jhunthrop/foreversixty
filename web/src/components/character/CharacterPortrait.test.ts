// web/src/components/character/CharacterPortrait.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterPortrait from './CharacterPortrait.svelte';

describe('CharacterPortrait', () => {
  it('shows the avatar image when avatar_url is set', () => {
    const { body } = render(CharacterPortrait, {
      props: {
        character: { name: 'Thoradin', class: 'warrior', avatar_url: '/a.jpg' },
        size: 'md',
        testid: 't',
      },
    });
    expect(body).toContain('data-testid="t-avatar"');
    expect(body).toContain('src="/a.jpg"');
    expect(body).toContain('alt=""');
    expect(body).toContain('loading="lazy"');
    expect(body).not.toContain('data-testid="t-avatar-fallback"');
  });

  it('shows the class icon over the letter square when there is no avatar', () => {
    const { body } = render(CharacterPortrait, {
      props: { character: { name: 'Thoradin', class: 'warrior' }, size: 'md', testid: 't' },
    });
    expect(body).toContain('data-testid="t-avatar-fallback"');
    expect(body).toContain('t-avatar-fallback">W');
    expect(body).toContain('data-testid="t-class-icon"');
    expect(body).toContain('classicon_warrior.jpg');
    expect(body).not.toContain('data-testid="t-avatar"');
  });

  it('shows the letter square alone for a class with no known icon', () => {
    const { body } = render(CharacterPortrait, {
      props: { character: { name: 'Elyra' }, size: 'md', testid: 't' },
    });
    expect(body).toContain('data-testid="t-avatar-fallback"');
    expect(body).toContain('t-avatar-fallback">E');
    expect(body).not.toContain('data-testid="t-class-icon"');
  });

  it('sizes the box for sm, md and lg', () => {
    const sm = render(CharacterPortrait, { props: { character: { name: 'E' }, size: 'sm', testid: 't' } });
    expect(sm.body).toContain('h-7 w-7');
    const md = render(CharacterPortrait, { props: { character: { name: 'E' }, size: 'md', testid: 't' } });
    expect(md.body).toContain('h-9 w-9');
    const lg = render(CharacterPortrait, { props: { character: { name: 'E' }, size: 'lg', testid: 't' } });
    expect(lg.body).toContain('h-11 w-11');
  });
});
