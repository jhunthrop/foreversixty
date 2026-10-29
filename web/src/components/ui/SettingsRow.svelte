<!-- web/src/components/ui/SettingsRow.svelte -->
<!-- One settings row (polish spec 4.2): optional IconTile, label, optional sub, a trailing value,
     a trailing control (a Switch or a button) and a chevron when it navigates. An `a` with an
     `href`, a `button` with `onclick`, else a `div`. Below 480 px the value moves under the
     label so the two never collide. A `nested` row indents its text to 70 px. States: muted
     value, `loading` (a 96 x 12 skeleton bar), `empty` (a muted dash), `error` (the wipe line),
     `disabled` (muted, no chevron, not a link). The label's id is exposed as `labelId` so a
     Switch in `control` can name itself by it. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import { ROW, ROW_INTERACTIVE, ROW_LABEL, ROW_SUB } from '../../lib/ui/row';
  import Glyph from './Glyph.svelte';
  import IconTile, { type TileTone } from './IconTile.svelte';

  const uid = $props.id();

  let {
    label,
    sub,
    value,
    muted = false,
    glyph,
    tone = 'gold',
    href,
    external = false,
    onclick,
    nested = false,
    loading = false,
    empty = false,
    error,
    disabled = false,
    control,
    controlBelow = false,
    below,
    testid,
    valueTestid,
  }: {
    label: string;
    sub?: string;
    value?: string;
    /** The value is a quiet state (off, not shipped): `--muted` and no chevron. */
    muted?: boolean;
    glyph?: GlyphName;
    tone?: TileTone;
    href?: string;
    /** A link that opens another tab: `target=_blank` with `rel=noopener`. */
    external?: boolean;
    onclick?: () => void;
    nested?: boolean;
    loading?: boolean;
    /** No value yet: a muted dash. */
    empty?: boolean;
    /** A short error replaces the value, in the failure colour. */
    error?: string;
    disabled?: boolean;
    /** The trailing control; receives the label's id. */
    control?: Snippet<[string]>;
    /** Below 480 px the control moves under the text, indented to it (a button too wide to sit beside it). */
    controlBelow?: boolean;
    /** Extra lines under the label and sub (a status line), inside the text column. */
    below?: Snippet;
    testid?: string;
    valueTestid?: string;
  } = $props();

  const labelId = $derived(`${uid}-label`);
  const interactive = $derived(!disabled && !loading && (href !== undefined || onclick !== undefined));
  const tag = $derived(interactive ? (href !== undefined ? 'a' : 'button') : 'div');
  const chevron = $derived(interactive && !muted && error === undefined);
  const shownValue = $derived(empty && value === undefined ? '—' : value);
  const valueColour = $derived(
    error !== undefined ? 'text-wipe' : muted || empty || disabled ? 'text-muted' : 'text-text',
  );
  /** Under the label on a phone: the tile's 32 px and the 14 px gap. */
  const phoneIndent = $derived(glyph === undefined ? '' : 'max-[479px]:pl-[46px]');
</script>

<!-- The element is a real button whenever it has a click handler (tag above); the compiler cannot see that. -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<svelte:element
  this={tag}
  href={tag === 'a' ? href : undefined}
  target={tag === 'a' && external ? '_blank' : undefined}
  rel={tag === 'a' && external ? 'noopener' : undefined}
  type={tag === 'button' ? 'button' : undefined}
  onclick={tag === 'button' ? onclick : undefined}
  aria-disabled={disabled ? 'true' : undefined}
  aria-busy={loading ? 'true' : undefined}
  data-testid={testid}
  class="{ROW} {interactive ? ROW_INTERACTIVE : ''}"
  style:padding-left={nested ? (glyph === undefined ? '70px' : '46px') : undefined}
>
  {#if glyph !== undefined}<IconTile {glyph} size={32} {tone} />{/if}
  <span class="flex min-w-0 flex-1 flex-col">
    <span
      id={labelId}
      class="{ROW_LABEL} {disabled ? 'text-muted' : 'text-strong'}"
      data-testid={testid === undefined ? undefined : `${testid}-label`}>{label}</span
    >
    {#if sub !== undefined}<span class={ROW_SUB} data-testid="settings-row-sub">{sub}</span>{/if}
    {@render below?.()}
  </span>
  {#if loading}
    <span
      class="bg-line h-3 w-24 animate-pulse rounded-full motion-reduce:animate-none max-[479px]:order-last max-[479px]:basis-full {phoneIndent}"
      data-testid="settings-row-loading"
    ></span>
  {:else if shownValue !== undefined || error !== undefined}
    <span
      class="text-[14px] max-[479px]:order-last max-[479px]:basis-full max-[479px]:text-[13px] min-[480px]:text-right {phoneIndent} {valueColour}"
      data-testid={valueTestid ?? 'settings-row-value'}>{error ?? shownValue}</span
    >
  {/if}
  {#if control !== undefined}
    <span class={controlBelow ? `max-[479px]:order-last max-[479px]:basis-full ${phoneIndent}` : ''}
      >{@render control(labelId)}</span
    >
  {/if}
  {#if chevron}<Glyph name="chevron" size={16} class="text-muted" />{/if}
</svelte:element>
