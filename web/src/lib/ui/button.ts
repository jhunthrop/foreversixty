// web/src/lib/ui/button.ts
// The button recipes (round 5, design/specs/2026-11-09-directions.md): primary is their flat teal
// with the teal edge (Buy Now), secondary their blue (Upgrade Now), ghost the gold text link,
// danger the wipe outline; each in the two sizes, on the 6 px control radius, a white semibold
// label in sentence case. A caller composes `${BUTTON_PRIMARY} ${BUTTON_MD}`, or takes
// `button('primary', 'md')`. Disabled reads at 45 % with a not-allowed cursor, except a busy button,
// which keeps its look (aria-busy="true"; the caller swaps the glyph for ui/Spinner and the label for
// the busy word). Premium in the nav is its own recipe (nav.css: their Subscribe, exactly).

const BASE =
  'rounded-field inline-flex items-center justify-center gap-2 border px-5 button-label whitespace-nowrap ' +
  'transition-[filter,background-color,border-color] duration-[120ms] ' +
  'focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold ' +
  'disabled:cursor-not-allowed disabled:not-aria-busy:opacity-45 aria-disabled:cursor-not-allowed aria-disabled:not-aria-busy:opacity-45';

// Hover applies only where a pointer can hover, so a tap never leaves a stuck hover look. Each
// hover class is written out whole, `[@media(hover:hover)_and_(pointer:fine)]:hover:<utility>`:
// Tailwind finds classes by scanning source text, so a class joined from pieces in a template
// string is never emitted.

export const BUTTON_PRIMARY =
  `${BASE} border-teal-edge text-strong bg-teal ` +
  '[@media(hover:hover)_and_(pointer:fine)]:hover:bg-teal-edge active:translate-y-px active:brightness-[.92]';

export const BUTTON_SECONDARY =
  `${BASE} border-blue-edge text-strong bg-blue ` +
  '[@media(hover:hover)_and_(pointer:fine)]:hover:brightness-[1.15] active:translate-y-px active:brightness-[.92]';

export const BUTTON_GHOST =
  `${BASE} text-gold border-transparent bg-transparent px-[10px] ` +
  '[@media(hover:hover)_and_(pointer:fine)]:hover:text-gold-hover active:text-gold-hover';

/** A destructive confirm (premium spec 2026-11-08 section 3A): wipe ink and a 40 % wipe hairline,
 *  never filled; hover takes the full wipe border. */
export const BUTTON_DANGER =
  `${BASE} text-wipe border-wipe/40 bg-transparent ` +
  '[@media(hover:hover)_and_(pointer:fine)]:hover:border-wipe active:translate-y-px';

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
