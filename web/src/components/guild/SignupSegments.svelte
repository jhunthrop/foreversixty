<!-- web/src/components/guild/SignupSegments.svelte -->
<!-- In · Late · Bench · Out (spec 3.3, 3.4, 4): bis/PresetControl.astro's segmented recipe rebuilt
     in Tailwind. Buttons with role="radio", not native radios, so the 44 px boxes style freely;
     one roving tab stop, arrows move and select, Space and Enter select (segments-keys.ts). -->
<script lang="ts">
  import type { GuildSignupState } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { SIGNUP_STATES, segmentKeyAction, tabbableSegment } from '../../lib/guild/segments-keys';

  let {
    value,
    pending = null,
    name,
    size,
    disabled = false,
    describedBy,
    onChange,
    testid = 'signup-segments',
  }: {
    value: GuildSignupState | null;
    /** Pressed but not saved (the spec pick is owed): drawn as pending, never as checked. */
    pending?: GuildSignupState | null;
    name: string;
    /** `strip` is 44 px everywhere; `band` is 44 px on phone and 36 px from md. */
    size: 'band' | 'strip';
    disabled?: boolean;
    /** The state line's id, read while disabled. */
    describedBy?: string;
    onChange: (state: GuildSignupState) => void;
    testid?: string;
  } = $props();

  const buttons: HTMLButtonElement[] = [];
  const tabStop = $derived(tabbableSegment(value));

  const BASE =
    'min-w-[72px] flex-1 px-2 text-[13px] font-semibold transition-colors duration-[120ms] focus-visible:outline-gold focus-visible:outline-2 focus-visible:-outline-offset-2 active:bg-card-top disabled:cursor-not-allowed [@media(hover:hover)_and_(pointer:fine)]:enabled:hover:text-text';
  const REST = 'text-muted bg-raised';
  // The checked ground is visible against the rest's `--raised` (spec 3.3, amended 2026-10-27, review round).
  const CHECKED_IN = 'text-strong bg-kill/15 shadow-[inset_0_-2px_0_var(--color-kill)]';
  const CHECKED_OUT = 'text-strong bg-line-warm shadow-[inset_0_-2px_0_var(--color-muted)]';
  const PENDING = 'text-text outline-gold outline-2 -outline-offset-2 outline-dashed';
  const SIZE = { strip: 'h-11', band: 'h-11 md:h-9' } as const;

  function segmentClass(state: GuildSignupState, index: number): string {
    const divider = index > 0 ? 'border-line border-l' : '';
    const checked = value === state ? (state === 'in' || state === 'late' ? CHECKED_IN : CHECKED_OUT) : REST;
    const waiting = pending === state ? PENDING : '';
    return [BASE, SIZE[size], divider, checked, waiting].filter(Boolean).join(' ');
  }

  function select(state: GuildSignupState): void {
    if (state !== value) onChange(state);
  }

  /** Arrows move focus only (DOM focus; the tab stop stays on the checked segment), so stepping
   *  across the group never saves; Space and Enter select the focused segment. */
  function onKeydown(event: KeyboardEvent, index: number): void {
    const action = segmentKeyAction(event.key, index, SIGNUP_STATES.length);
    if (action === null) return;
    event.preventDefault();
    if (action.kind === 'move') buttons[action.index]?.focus();
    else select(SIGNUP_STATES[action.index]);
  }
</script>

<div
  role="radiogroup"
  aria-label={`${name}'s sign-up`}
  aria-describedby={disabled ? describedBy : undefined}
  class="border-line rounded-control flex w-full overflow-hidden border md:w-auto {disabled
    ? 'opacity-50'
    : ''}"
  data-testid={testid}
>
  {#each SIGNUP_STATES as state, index (state)}
    <button
      bind:this={buttons[index]}
      type="button"
      role="radio"
      aria-checked={value === state ? 'true' : 'false'}
      aria-busy={pending === state ? 'true' : undefined}
      tabindex={index === tabStop ? 0 : -1}
      {disabled}
      class={segmentClass(state, index)}
      data-testid={`${testid}-${state}`}
      onclick={() => select(state)}
      onkeydown={(event) => onKeydown(event, index)}
    >
      {raidCopy.segments[state]}
    </button>
  {/each}
</div>
