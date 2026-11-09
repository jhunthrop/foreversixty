<!-- web/src/components/coach/SpellIcon.svelte -->
<!-- A spell's HD icon from the build's icon files at a fixed size, framed the way the best addons
     frame an action button: the art cropped 8 % to drop Blizzard's baked bevel, a 1 px black edge,
     a 4 px radius (the best-addon tenet; ItemIcon is its item twin). `width` and `height` are set so
     the box is reserved before the image arrives (spec 4.1 rule 8). A spell with no icon in the
     build renders nothing, so no gap is left; an image that fails to load leaves the dark socket at
     its size (spec 2026-11-07 logs section 9: the socket when the build has no icon), never a gap. -->
<script lang="ts">
  let { url, size, name = '' }: { url: string | null; size: number; name?: string } = $props();

  let failed = $state<string | null>(null);
</script>

{#if url !== null}
  <span
    class="inline-flex shrink-0 overflow-hidden rounded-[4px] border border-black/80 bg-black/60 ring-1 ring-white/4"
    style:width="{size}px"
    style:height="{size}px"
  >
    {#if failed === url}
      <!-- With no name it is decoration: hidden, so it is never an image with no text. -->
      <span
        role={name === '' ? undefined : 'img'}
        aria-hidden={name === '' ? 'true' : undefined}
        aria-label={name === '' ? undefined : name}
        title={name === '' ? undefined : name}
        class="h-full w-full"
        data-testid="coach-spell-icon"
      ></span>
    {:else}
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
    {/if}
  </span>
{/if}
