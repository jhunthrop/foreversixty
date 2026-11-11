<!-- web/src/components/ui/CopyButton.svelte -->
<!-- The one copy button (component audit 2026-11-10 cluster 13): a Button whose label reads
     "Copied" for about two seconds after the write (owner ruling 2026-11-10) and "Copy failed"
     when the clipboard refuses, through lib/ui/copy-feedback.ts. `text` is read at press time, so
     a caller can hand a function that builds the CSV or the post on demand. -->
<script lang="ts">
  import type { ButtonSize, ButtonTone } from '../../lib/ui/button';
  import { CopyState } from '../../lib/ui/copy-state.svelte';
  import Button from './Button.svelte';

  let {
    text,
    label,
    tone = 'secondary',
    size = 'sm',
    disabled = false,
    ariaLabel,
    oncopied,
    testid,
    class: extra = '',
  }: {
    text: string | (() => string);
    /** The idle words ("Copy link", "Copy for Discord"). */
    label: string;
    tone?: ButtonTone;
    size?: ButtonSize;
    disabled?: boolean;
    /** The accessible name when the label needs more ("Copy the roster as CSV"). */
    ariaLabel?: string;
    /** After a successful write (a page that also records the share). */
    oncopied?: () => void;
    testid?: string;
    class?: string;
  } = $props();

  const copier = new CopyState();
  $effect(() => () => copier.dispose());

  async function onclick(): Promise<void> {
    await copier.copy(typeof text === 'function' ? text() : text);
    if (copier.state === 'copied') oncopied?.();
  }
</script>

<Button {tone} {size} {disabled} label={ariaLabel} {testid} class={extra} onclick={() => void onclick()}
  ><span aria-live="polite">{copier.label(label)}</span></Button
>
