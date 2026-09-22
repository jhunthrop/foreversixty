<!-- web/src/components/ui/StatePanel.svelte -->
<!-- The design system's "State panel" (design/DESIGN-SYSTEM.md's Components section): a
     label row with the glowing gold dot and a right-aligned Updated/Sample stamp, over the
     panel box (`bg-raised`, `border-line`, `rounded-panel`). Svelte twin of
     StatePanel.astro (the Astro-only version, used on marketing/reference pages) -- that
     component takes a fixed rows array and cannot mount inside a Svelte island, and its own
     API is too narrow for /account's rail panels, which hold real controls (a Revoke
     button, a consent select, an anonymize checkbox), not just key/value text. This is the
     one exception to account-visual owning no other `components/ui/**` file (the brief's
     own carve-out): the rail needs the shared label-row treatment, and AccountPanel.svelte
     already duplicated the box half of it once. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    label,
    updated = '',
    testid,
    aside,
    children,
  }: {
    label: string;
    updated?: string;
    testid?: string;
    /** The label row's right-hand slot for something richer than plain text -- Characters'
     *  "Imported from Battle.net 4 hours ago · Refresh from Battle.net" link (brief
     *  2026-09-22 §B3). Takes priority over `updated` when both are given. */
    aside?: Snippet;
    children: Snippet;
  } = $props();
</script>

<div
  class="bg-raised border-line rounded-panel flex flex-col gap-4 border px-4 py-4 md:px-6 md:py-5"
  data-testid={testid}
>
  <div class="flex items-center gap-[10px]">
    <span class="bg-gold h-2 w-2 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]" aria-hidden="true"
    ></span>
    <span class="label text-gold">{label}</span>
    {#if aside !== undefined}
      <span class="text-muted ml-auto text-[12px]">{@render aside()}</span>
    {:else if updated !== ''}
      <span
        class="text-muted ml-auto text-[12px]"
        data-testid={testid !== undefined ? `${testid}-updated` : undefined}
      >
        Updated {updated}
      </span>
    {/if}
  </div>
  {@render children()}
</div>
