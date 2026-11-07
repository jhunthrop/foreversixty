<!-- web/src/components/report/SpellIcon.svelte -->
<!-- A spell's HD icon in a table row or a rail row (logs spec 2026-11-07 section 9, "Spell / item
     icon"): width and height always set, so nothing shifts; a spell the build has no icon for, a
     melee swing, or an image that fails draws the dark socket at size, never a broken image. -->
<script lang="ts">
  let { url, size = 24 }: { url: string | null; size?: number } = $props();
  let failed = $state<string | null>(null);
</script>

{#if url !== null && url !== failed}
  <img
    src={url}
    width={size}
    height={size}
    alt=""
    loading="lazy"
    decoding="async"
    class="inline-block shrink-0 rounded-[4px]"
    style:width="{size}px"
    style:height="{size}px"
    onerror={() => (failed = url)}
  />
{:else}
  <span
    class="inline-block shrink-0 rounded-[4px] bg-white/[.04] shadow-[inset_0_0_0_1px_rgba(255,255,255,.06)]"
    style:width="{size}px"
    style:height="{size}px"
    aria-hidden="true"
  ></span>
{/if}
