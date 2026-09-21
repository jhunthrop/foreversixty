// web/src/components/CurrentCharacterChip.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { currentCharacterCopy } from '../lib/current-character-copy';
import type { CurrentCharacter } from '../lib/current-character';
import { classColorVar } from '../lib/report/format';
import CurrentCharacterChip from './CurrentCharacterChip.svelte';

const current: CurrentCharacter = {
  source: 'addon',
  ref: 'FS1:1:warrior:orc:0/0/0:',
  label: 'Simfury · Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-21T00:00:00.000Z',
};

describe('CurrentCharacterChip', () => {
  it('shows the label and the three links when a character is loaded', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).toContain('Simfury · Fury Warrior');
    expect(body).toContain(currentCharacterCopy.openInPlanner);
    expect(body).toContain(currentCharacterCopy.openInSimulator);
    expect(body).toContain(currentCharacterCopy.forget);
    expect(body).toContain('/planner?code=');
    expect(body).toContain('/sim?code=');
  });

  it('renders the no-character line on a page without its own paste box', () => {
    const { body } = render(CurrentCharacterChip, { props: { current: null, onforget: () => {} } });
    expect(body).toContain(currentCharacterCopy.noCharacterLine);
  });

  it('renders nothing on a page that already has its own paste box', () => {
    const { body } = render(CurrentCharacterChip, {
      props: { current: null, onforget: () => {}, hasOwnPasteBox: true },
    });
    // Svelte 5's SSR still emits its own anchor comments (`<!--[-->...<!--]-->`) around an
    // empty conditional region -- internal hydration markers, never visible content -- so
    // the assertion strips HTML comments before checking that nothing else was rendered.
    expect(body.replace(/<!--.*?-->/gs, '').trim()).toBe('');
  });

  it('fixes its own height so resolving does not move content (data-testid anchor present)', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).toContain('data-testid="current-character-chip"');
  });

  it('shows the label in class colour', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).toContain(classColorVar('warrior'));
  });

  it('renders no copy-addon-code button when addonCode is empty', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).not.toContain('data-testid="current-character-copy-addon"');
  });

  it('renders the copy-addon-code button when addonCode is given, with the copy label', () => {
    const { body } = render(CurrentCharacterChip, {
      props: { current, onforget: () => {}, addonCode: 'FS1:1:warrior:orc:0/0/0:' },
    });
    expect(body).toContain('data-testid="current-character-copy-addon"');
    expect(body).toContain(currentCharacterCopy.copyAddonCode);
    expect(body).not.toContain(currentCharacterCopy.copiedAddonCode);
  });

  // Fix round 1 (Critical): the loaded chip and the no-character line must share the exact
  // same fixed height, in every rendered state, or resolving between them moves content
  // below the chip. Both classes of CHIP_HEIGHT are asserted directly rather than via a
  // single combined substring, so a change to either one alone still fails this test.
  it('gives the loaded chip and the no-character line the identical fixed height', () => {
    const loaded = render(CurrentCharacterChip, { props: { current, onforget: () => {} } }).body;
    const empty = render(CurrentCharacterChip, { props: { current: null, onforget: () => {} } }).body;
    for (const body of [loaded, empty]) {
      expect(body).toContain('h-[88px]');
      expect(body).toContain('md:h-11');
    }
  });

  it('never wraps the loaded chip (no flex-wrap anywhere in its markup)', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).not.toContain('flex-wrap');
  });
});
