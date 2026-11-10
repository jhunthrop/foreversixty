<!-- web/src/components/sim/ActionMark.svelte -->
<!-- One action as the saved run's lower cards draw it: its icon in a fixed box and its name, opening
     the shared tooltip when the build has one. A spell reads the game tooltip, an item the item
     tooltip. A key with no icon draws the dark socket every missing icon shows (component audit
     2026-11-10 cluster 1). The
     engine's rage refund has no client icon: in the 24 px icon column it takes the design system's
     muted refresh glyph, centred in the column, so it reads as deliberate, never as a missing icon. -->
<script lang="ts">
  import type { SpellIconSize } from '../../lib/ui/spell-icon';
  import { spellTipKey } from '../../lib/tooltips/game-tip';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import type { Snippet } from 'svelte';
  import type { CardTables } from '../../lib/sim/card-tables.svelte';
  import { actionTip, resolveIconUrl, REFUND_KEY } from '../../lib/sim/action-mark';
  import Icon from '../ui/Icon.svelte';
  import { ICON_REFRESH } from '../../lib/ui/ui-icons';

  /** The refund row's tile size: the 24 px icon column of the Resource card. */
  const REFUND_TILE = 24;

  let {
    action,
    tables,
    size = 24,
    fallbackIcon = null,
    children,
  }: {
    /** Null for a row that is no engine action: a raid buff the request named. */
    action: string | null;
    tables: CardTables;
    size?: SpellIconSize;
    /** The simbuffs icon of the buff this row is, when the action has none of its own. */
    fallbackIcon?: string | null;
    children: Snippet;
  } = $props();

  const icon = $derived(resolveIconUrl(action, tables.visuals, fallbackIcon));
  const opens = $derived(actionTip(action, (id) => tables.spells.fact(id) !== undefined));
  const tipKey = $derived(opens.spell === undefined ? undefined : spellTipKey(opens.spell));
  const itemId = $derived(opens.item);
  const openable = $derived(tipKey !== undefined || itemId !== undefined);
</script>

<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
<span
  class="flex min-w-0 items-center gap-2.5"
  data-game-tip={tipKey}
  data-bis-item={itemId}
  tabindex={openable ? 0 : undefined}
>
  <span class="grid shrink-0 place-items-center" style:width={`${size}px`} style:height={`${size}px`}>
    {#if action !== REFUND_KEY}
      <SpellIcon url={icon} {size} />
    {:else if size === REFUND_TILE}
      <span
        class="text-muted inline-flex shrink-0 items-center justify-center"
        style:width="{REFUND_TILE}px"
        style:height="{REFUND_TILE}px"
        aria-hidden="true"
        data-testid="refund-glyph"><Icon icon={ICON_REFRESH} size={14} /></span
      >
    {:else}
      <Icon icon={ICON_REFRESH} size={Math.round(size * 0.7)} class="text-muted" />
    {/if}
  </span>
  <span class="min-w-0">{@render children()}</span>
</span>
