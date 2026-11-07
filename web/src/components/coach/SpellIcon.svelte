<!-- web/src/components/coach/SpellIcon.svelte -->
<!-- A spell's HD icon from the build's icon files at a fixed size, framed the way the best addons
     frame an action button: the art cropped 8 % to drop Blizzard's baked bevel, a 1 px black edge,
     a 4 px radius (the best-addon tenet; ItemIcon is its item twin). `width` and `height` are set so
     the box is reserved before the image arrives (spec 4.1 rule 8). A spell with no icon in the
     build renders nothing, so no gap is left; an image that fails to load hides itself. -->
<script lang="ts">
  let { url, size, name = '' }: { url: string | null; size: number; name?: string } = $props();

  let failed = $state<string | null>(null);
</script>

{#if url !== null && failed !== url}
  <span
    class="inline-flex shrink-0 overflow-hidden rounded-[4px] border border-black/80 shadow-[0_0_0_1px_rgba(255,255,255,0.04)]"
    style:width="{size}px"
    style:height="{size}px"
  >
    <img
      src={url}
      width={size}
      height={size}
      alt=""
      title={name === '' ? undefined : name}
      class="h-full w-full scale-[1.08] object-cover"
      data-testid="coach-spell-icon"
      onerror={() => (failed = url)}
    />
  </span>
{/if}
