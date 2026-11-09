<!-- web/src/components/ui/SpellIcon.svelte -->
<!-- A spell's icon in a row (logs spec 2026-11-07 section 9, "Spell / item icon"; the mock's `sicon`
     in ws_report_lib.py over ws_lib's `.ic`): the HD icon in a 1 px warm frame, 5 px radius, width
     and height always set. A spell the build's icon table does not carry (Melee, most of Forever's
     new ids: gap G15) draws the dark socket at the same size, never a broken image; an image that
     fails to load becomes the socket too. `tip` opens the shared spell tooltip on hover and focus
     of the row's name; the icon carries the same key for the pointer. Decorative: the name beside
     it says what it is. -->
<script lang="ts">
  import { spellTipKey } from '../../lib/tooltips/game-tip';

  let {
    url,
    spellId,
    size = 28,
    tip = false,
    testid,
  }: {
    /** The icon's url, from the build's spell icon table; null draws the socket. */
    url: string | null;
    /** The spell, for the tooltip. */
    spellId?: number;
    size?: 20 | 24 | 28;
    /** Whether hovering the icon opens the spell's tooltip. */
    tip?: boolean;
    testid?: string;
  } = $props();

  /** The url whose image failed; a new url gets its own try. */
  let failed = $state<string | null>(null);
  const shown = $derived(url !== null && url !== failed ? url : null);
</script>

<span
  class="ui-spell-icon border-line-warm-strong inline-flex shrink-0 overflow-hidden rounded-[5px] border bg-[var(--color-bg)]"
  style:width="{size}px"
  style:height="{size}px"
  aria-hidden="true"
  data-game-tip={tip && spellId !== undefined && spellId > 0 ? spellTipKey(spellId) : undefined}
  data-icon-socket={shown === null ? '' : undefined}
  data-testid={testid}
>
  {#if shown !== null}
    <img
      src={shown}
      alt=""
      width={size}
      height={size}
      loading="lazy"
      decoding="async"
      class="block h-full w-full object-cover"
      onerror={() => (failed = shown)}
    />
  {/if}
</span>
