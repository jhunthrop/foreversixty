<!-- web/src/components/ui/Button.svelte -->
<!-- The button (web layout spec 2026-11-05 §3.5) over lib/ui/button.ts's recipes, the single
     source: a link when it navigates (`href`), else a <button>. One primary per view. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { button, type ButtonSize, type ButtonTone } from '../../lib/ui/button';

  let {
    tone = 'secondary',
    size = 'md',
    sentence = false,
    href,
    rel,
    type = 'button',
    disabled = false,
    onclick,
    testid,
    children,
  }: {
    tone?: ButtonTone;
    size?: ButtonSize;
    sentence?: boolean;
    href?: string;
    /** A link's rel (`noopener` on every install and download link). */
    rel?: string;
    type?: 'button' | 'submit';
    disabled?: boolean;
    onclick?: (event: MouseEvent) => void;
    testid?: string;
    children: Snippet;
  } = $props();

  const classes = $derived(button(tone, size, sentence));
</script>

{#if href !== undefined && !disabled}
  <a class={classes} {href} {rel} data-testid={testid}>{@render children()}</a>
{:else}
  <button class={classes} {type} {disabled} {onclick} data-testid={testid}>{@render children()}</button>
{/if}
