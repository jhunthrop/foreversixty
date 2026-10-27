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
</script>

<div
  class="flex flex-col gap-1 border-l-2 pl-3 {tone === 'wipe' ? 'border-wipe' : 'border-gold'}"
  data-testid="discord-notice"
>
  <p class="text-strong text-[14px]">
    {#if strong !== null && title.includes(strong)}{title.slice(0, title.indexOf(strong))}<span
        class="font-semibold">{strong}</span
      >{title.slice(title.indexOf(strong) + strong.length)}{:else}{title}{/if}
  </p>
  {#if detail !== null}<p class="text-muted text-[13px]">{detail}</p>{/if}
  {@render children?.()}
</div>
