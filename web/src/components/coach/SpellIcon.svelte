<!-- web/src/components/coach/SpellIcon.svelte -->
<!-- A spell's HD icon from the build's icon files at a fixed size; `width` and `height` are set
     so the box is reserved before the image arrives (spec 4.1 rule 8). A spell with no icon in
     the build renders nothing, so no gap is left. -->
<script lang="ts">
  let { url, size, name = '' }: { url: string | null; size: number; name?: string } = $props();
</script>

{#if url !== null}
  <img
    src={url}
    width={size}
    height={size}
    alt=""
    title={name === '' ? undefined : name}
    class="border-line-soft inline-block shrink-0 rounded-[3px] border"
    style={`width: ${size}px; height: ${size}px`}
    data-testid="coach-spell-icon"
  />
{/if}
