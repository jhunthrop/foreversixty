<!-- web/src/components/ui/Button.svelte -->
<!-- The button (web layout spec 2026-11-05 §3.5) over lib/ui/button.ts's recipes, the single
     source: a link when it navigates (`href`), else a <button>. One primary per view. `busy` is the
     one busy look (owner ruling 2026-11-10): the spinner and the busy word ("Saving…") in place of
     the label, disabled and aria-busy, the button keeping its tone. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { button, type ButtonSize, type ButtonTone } from '../../lib/ui/button';
  import { busyCopy } from '../../lib/ui/copy';
  import Spinner from './Spinner.svelte';

  let {
    tone = 'secondary',
    size = 'md',
    href,
    rel,
    download,
    type = 'button',
    disabled = false,
    onclick,
    pressed,
    label,
    testid,
    busy = false,
    busyLabel = busyCopy.working,
    class: extra = '',
    children,
  }: {
    tone?: ButtonTone;
    size?: ButtonSize;
    href?: string;
    /** A link's rel (`noopener` on every install and download link). */
    rel?: string;
    /** A link's saved file name (the companion's ForeverSixty.dmg, ForeverSixtySetup.exe). */
    download?: string;
    type?: 'button' | 'submit';
    disabled?: boolean;
    onclick?: (event: MouseEvent) => void;
    /** A toggle's state (aria-pressed): the current choice of a set ("Late" among I'm in, Late, Can't come). */
    pressed?: boolean;
    /** The accessible name when the visible words need more ("Retry" that reloads the talents). */
    label?: string;
    testid?: string;
    /** The request this button started is in flight: spinner plus `busyLabel`, disabled. */
    busy?: boolean;
    /** The busy word ("Saving…", "Sending…"); lib/ui/copy.ts busyCopy has the common ones. */
    busyLabel?: string;
    /** Layout only (width, alignment): the look is the tone's. */
    class?: string;
    children: Snippet;
  } = $props();

  const classes = $derived(extra === '' ? button(tone, size) : `${button(tone, size)} ${extra}`);
</script>

{#if href !== undefined && !disabled && !busy}
  <a class={classes} {href} {rel} {download} aria-label={label} data-testid={testid}>{@render children()}</a>
{:else}
  <button
    class={classes}
    {type}
    disabled={disabled || busy}
    aria-busy={busy ? 'true' : undefined}
    {onclick}
    aria-pressed={pressed}
    aria-label={label}
    data-testid={testid}
    >{#if busy}<Spinner />{busyLabel}{:else}{@render children()}{/if}</button
  >
{/if}
