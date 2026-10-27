<!-- web/src/components/account/discord/DiscordNotice.svelte -->
<!-- The 2 px left-bar notice: failure colour for a switched-off webhook, gold for a lapsed plan. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    tone,
    title,
    detail = null,
    strong = null,
    children,
  }: {
    tone: 'wipe' | 'gold';
    title: string;
    /** A part of `title` drawn `--strong` 600. */
    strong?: string | null;
    detail?: string | null;
    children?: Snippet;
  } = $props();

  const at = $derived(strong === null ? -1 : title.indexOf(strong));
  const head = $derived(at < 0 ? title : title.slice(0, at));
  const tail = $derived(strong === null || at < 0 ? '' : title.slice(at + strong.length));
</script>

<div
  class="flex flex-col gap-1 border-l-2 pl-3 {tone === 'wipe' ? 'border-wipe' : 'border-gold'}"
  data-testid="discord-notice"
>
  <p class="text-strong text-[14px]">
    {head}{#if at >= 0}<span class="font-semibold">{strong}</span>{tail}{/if}
  </p>
  {#if detail !== null}<p class="text-muted text-[13px]">{detail}</p>{/if}
  {@render children?.()}
</div>
