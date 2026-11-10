<!-- web/src/components/ui/RaceIcon.svelte -->
<!-- A race's portrait (the client's square Achievement_Character_<Race>_Male icon, read from the
     build's raceicons.json): the HD icon in SpellIcon's 1 px warm frame, 5 px radius, width and
     height always set. The race's name is its title on hover; the name beside it says the rest,
     so screen readers skip it. A race with no portrait draws nothing, never an empty socket, so
     the caller renders it only with a url; an image that fails to load removes itself too. In a
     static render (an Astro page, no hydration) the error handler never runs, so a broken image is
     covered by a faction-tinted fill instead of the browser's broken-image glyph. -->
<script lang="ts">
  let {
    url,
    name,
    size = 28,
    testid,
  }: {
    /** The portrait's url (`raceIconsFor`). */
    url: string;
    /** The race's name ("Night Elf"), the icon's title. */
    name: string;
    size?: 20 | 24 | 28 | 40;
    testid?: string;
  } = $props();

  /** The url whose image failed; a new url gets its own try. */
  let failed = $state<string | null>(null);
</script>

{#if url !== failed}
  <span
    class="ui-race-icon border-line-warm-strong inline-flex shrink-0 overflow-hidden rounded-[5px] border bg-[var(--color-bg)]"
    style:width="{size}px"
    style:height="{size}px"
    title={name}
    aria-hidden="true"
    data-race-icon={name}
    data-testid={testid}
  >
    <img
      src={url}
      alt=""
      width={size}
      height={size}
      loading="lazy"
      decoding="async"
      class="ui-race-icon-img block h-full w-full object-cover"
      onerror={() => (failed = url)}
    />
  </span>
{/if}

<style>
  .ui-race-icon {
    position: relative;
  }
  /* Only a broken image draws pseudo-elements: this covers the browser's glyph. */
  .ui-race-icon-img::after {
    content: '';
    position: absolute;
    inset: 0;
    background: color-mix(in srgb, currentColor 30%, var(--color-bg));
  }
</style>
