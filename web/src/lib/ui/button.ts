// web/src/lib/ui/button.ts
// The button recipes (polish spec 4.2): primary, secondary, ghost and danger, each in the two
// sizes. A caller composes `${BUTTON_PRIMARY} ${BUTTON_MD}`, or takes `button('primary', 'md')`.
// Disabled reads at 45 % with a not-allowed cursor, except a busy button, which keeps its look
// (aria-busy="true"; the caller swaps the glyph for ui/Spinner and the label for the busy word).

const BASE =
  'rounded-control button-label inline-flex items-center justify-center gap-2 border px-[18px] whitespace-nowrap ' +
  'transition-[filter,background-color,border-color] duration-[120ms] ' +
  'focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold ' +
  'disabled:cursor-not-allowed disabled:not-aria-busy:opacity-45 aria-disabled:cursor-not-allowed aria-disabled:not-aria-busy:opacity-45';

/** Hover only where a pointer can hover, so a tap never leaves a stuck hover look. */
const HOVER = '[@media(hover:hover)_and_(pointer:fine)]:hover:';

export const BUTTON_PRIMARY =
  `${BASE} border-transparent text-bg bg-[linear-gradient(180deg,#f0cc6c,#c99a3a)] ` +
  'shadow-[inset_0_0_0_1px_var(--color-gold-deep),0_6px_20px_rgba(229,185,85,.22)] ' +
  `${HOVER}hover:brightness-[1.08] active:translate-y-px active:brightness-[.92] disabled:shadow-none aria-disabled:shadow-none`;

export const BUTTON_SECONDARY =
  `${BASE} border-line-card text-strong bg-white/[.045] ` +
  `${HOVER}hover:border-gold/30 ${HOVER}hover:bg-white/[.08] active:bg-card-top`;

export const BUTTON_GHOST =
  `${BASE} text-gold border-transparent bg-transparent px-[10px] ` +
  `${HOVER}hover:text-gold-hover ${HOVER}hover:bg-gold/[.06] active:text-gold-hover`;

export const BUTTON_DANGER =
  `${BASE} text-wipe border-wipe/30 bg-wipe/[.06] ` + `${HOVER}hover:bg-wipe/[.12] active:bg-wipe/[.16]`;

/** 44 px everywhere. */
export const BUTTON_MD = 'h-11';

/** 36 px from 1024 with a fine pointer; 44 px below, where the hit target has to clear the minimum. */
export const BUTTON_SM = 'h-11 lg:[@media(pointer:fine)]:h-9';

export type ButtonTone = 'primary' | 'secondary' | 'ghost' | 'danger';
export type ButtonSize = 'md' | 'sm';

const TONES: Readonly<Record<ButtonTone, string>> = {
  primary: BUTTON_PRIMARY,
  secondary: BUTTON_SECONDARY,
  ghost: BUTTON_GHOST,
  danger: BUTTON_DANGER,
};

const SIZES: Readonly<Record<ButtonSize, string>> = { md: BUTTON_MD, sm: BUTTON_SM };

/** One button class string: `button('danger', 'md')`. */
export function button(tone: ButtonTone, size: ButtonSize = 'md'): string {
  return `${TONES[tone]} ${SIZES[size]}`;
}
