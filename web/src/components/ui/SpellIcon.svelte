<!-- web/src/components/ui/SpellIcon.svelte -->
<!-- The one spell icon (component audit 2026-11-10 cluster 1; logs spec 2026-11-07 section 9,
     "Spell / item icon"): every spell, ability, buff, enchant, consumable and spec icon on the
     site, the HD art in a 1 px warm frame at the icon radius, width and height always set so
     nothing shifts. No url (an id the build's icon table does not carry, Melee) draws the dark
     socket at the same size, never a broken image; an image that fails to load becomes the socket
     too. `loading` holds the skeleton while the caller's icon table loads. `tip` names the spell
     whose shared tooltip the icon opens on hover, `item` the item's. `label` names an icon that stands alone; beside a
     printed name it stays empty and the icon is decoration. Its static twin is SpellIcon.astro. -->
<script lang="ts">
  import { spellTipKey } from '../../lib/tooltips/game-tip';
  import {
    SPELL_ICON_FRAME,
    SPELL_ICON_IMAGE,
    spellIconA11y,
    type SpellIconSize,
  } from '../../lib/ui/spell-icon';

  let {
    url,
    size = 28,
    tip,
    item,
    label = '',
    loading = false,
    eager = false,
    onfail,
    testid,
  }: {
    /** The icon's url (lib/ui/assets iconUrl, or a build's icon table); null draws the socket. */
    url: string | null;
    size?: SpellIconSize;
    /** The spell whose tooltip opens on hover. */
    tip?: number;
    /** Or the item whose tooltip opens on hover (a trinket or potion the action is). */
    item?: number;
    /** The name, for an icon with no printed name beside it. */
    label?: string;
    /** The icon table is still loading: the skeleton, not yet the socket. */
    loading?: boolean;
    /** Above the fold: load at once instead of lazily. */
    eager?: boolean;
    /** Called once when the image fails to load, for a caller that draws its own fallback. */
    onfail?: () => void;
    testid?: string;
  } = $props();

  /** The url whose image failed; a new url gets its own try. */
  let failed = $state<string | null>(null);
  const shown = $derived(url !== null && url !== failed ? url : null);
</script>

{#if loading}
  <span
    class="skeleton-block rounded-icon inline-block shrink-0"
    style:width="{size}px"
    style:height="{size}px"
    aria-hidden="true"
    data-testid={testid}
  ></span>
{:else}
  <span
    class={SPELL_ICON_FRAME}
    style:width="{size}px"
    style:height="{size}px"
    {...spellIconA11y(label)}
    data-game-tip={tip !== undefined && tip > 0 ? spellTipKey(tip) : undefined}
    data-bis-item={item}
    data-icon-socket={shown === null ? '' : undefined}
    data-testid={testid}
  >
    {#if shown !== null}
      <img
        src={shown}
        alt=""
        width={size}
        height={size}
        loading={eager ? 'eager' : 'lazy'}
        decoding="async"
        class={SPELL_ICON_IMAGE}
        onerror={() => {
          failed = shown;
          onfail?.();
        }}
      />
    {/if}
  </span>
{/if}
