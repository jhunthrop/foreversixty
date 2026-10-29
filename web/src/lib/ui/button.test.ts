import { describe, expect, it } from 'vitest';
import {
  BUTTON_DANGER,
  BUTTON_GHOST,
  BUTTON_MD,
  BUTTON_PRIMARY,
  BUTTON_SECONDARY,
  BUTTON_SM,
  button,
} from './button';

const ALL = [BUTTON_PRIMARY, BUTTON_SECONDARY, BUTTON_GHOST, BUTTON_DANGER];

describe('button recipes', () => {
  it.each(ALL)(
    'every tone shares the label (13 px 700 uppercase), the ring and the 45 % disabled look',
    (recipe) => {
      expect(recipe).toContain('button-label');
      expect(recipe).toContain('focus-visible:outline-gold');
      expect(recipe).toContain('disabled:not-aria-busy:opacity-45');
      expect(recipe).toContain('disabled:cursor-not-allowed');
      expect(recipe).toContain('gap-2');
    },
  );

  it('primary is the gold gradient with dark ink, the deep inset and the soft shadow', () => {
    expect(BUTTON_PRIMARY).toContain('bg-[linear-gradient(180deg,#f0cc6c,#c99a3a)]');
    expect(BUTTON_PRIMARY).toContain('text-bg');
    expect(BUTTON_PRIMARY).toContain('var(--color-gold-deep)');
    expect(BUTTON_PRIMARY).toContain('rgba(229,185,85,.22)');
  });

  it('primary hovers brighter, presses down darker and drops its shadow when disabled', () => {
    expect(BUTTON_PRIMARY).toContain('hover:brightness-[1.08]');
    expect(BUTTON_PRIMARY).toContain('active:brightness-[.92]');
    expect(BUTTON_PRIMARY).toContain('active:translate-y-px');
    expect(BUTTON_PRIMARY).toContain('disabled:shadow-none');
  });

  it('secondary has the card border, the strong ink and the three states', () => {
    expect(BUTTON_SECONDARY).toContain('border-line-card');
    expect(BUTTON_SECONDARY).toContain('text-strong');
    expect(BUTTON_SECONDARY).toContain('hover:border-gold/30');
    expect(BUTTON_SECONDARY).toContain('hover:bg-white/[.08]');
    expect(BUTTON_SECONDARY).toContain('active:bg-card-top');
  });

  it('ghost is gold ink on no ground with 10 px padding', () => {
    expect(BUTTON_GHOST).toContain('text-gold');
    expect(BUTTON_GHOST).toContain('bg-transparent');
    expect(BUTTON_GHOST).toContain('px-[10px]');
    expect(BUTTON_GHOST).toContain('hover:bg-gold/[.06]');
    expect(BUTTON_GHOST).toContain('hover:text-gold-hover');
  });

  it('danger is wipe ink on a 6 % ground with a 32 % border, 12 % on hover, 16 % pressed', () => {
    expect(BUTTON_DANGER).toContain('text-wipe');
    expect(BUTTON_DANGER).toContain('bg-wipe/[.06]');
    expect(BUTTON_DANGER).toContain('border-wipe/30');
    expect(BUTTON_DANGER).toContain('hover:bg-wipe/[.12]');
    expect(BUTTON_DANGER).toContain('active:bg-wipe/[.16]');
  });

  it('hover applies only where a pointer can hover', () => {
    for (const recipe of ALL.filter((r) => r !== BUTTON_PRIMARY || true))
      expect(recipe).toContain('[@media(hover:hover)_and_(pointer:fine)]:hover:');
  });

  it('md is 44 px; sm is 44 px until a fine pointer at 1024 and up, then 36 px', () => {
    expect(BUTTON_MD).toBe('h-11');
    expect(BUTTON_SM).toContain('h-11');
    expect(BUTTON_SM).toContain('lg:[@media(pointer:fine)]:h-9');
  });

  it('button() composes a tone and a size, md by default', () => {
    expect(button('danger')).toBe(`${BUTTON_DANGER} ${BUTTON_MD}`);
    expect(button('primary', 'sm')).toBe(`${BUTTON_PRIMARY} ${BUTTON_SM}`);
  });
});
