<!-- web/src/components/sim/ActionMark.svelte -->
<!-- One action as the saved run's lower cards draw it: its icon in a fixed box and its name, opening
     the shared tooltip when the build has one. A spell reads the game tooltip, an item the item
     tooltip. A key with no icon keeps the blank cell and never an empty box (build ruling R1). The
     engine's rage refund has no client icon and takes the page's one glyph. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { parseActionKey } from '../../lib/sim/action-names';
  import type { CardTables } from '../../lib/sim/card-tables.svelte';
  import { resolveIconUrl, REFUND_KEY } from '../../lib/sim/action-mark';
  import Glyph from '../ui/Glyph.svelte';

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
    size?: number;
    /** The simbuffs icon of the buff this row is, when the action has none of its own. */
    fallbackIcon?: string | null;
    children: Snippet;
  } = $props();

  const parsed = $derived(action === null ? null : parseActionKey(action));
  const icon = $derived(resolveIconUrl(action, tables.visuals, fallbackIcon));
  const tipKey = $derived(
    parsed?.kind === 'spell' && tables.spells.fact(parsed.id) !== undefined
      ? `spell:${parsed.id}`
      : undefined,
  );
  const itemId = $derived(parsed?.kind === 'item' ? parsed.id : undefined);
  const openable = $derived(tipKey !== undefined || itemId !== undefined);
</script>

<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
<span
  class="flex min-w-0 items-center gap-2.5"
  data-game-tip={tipKey}
  data-bis-item={itemId}
  tabindex={openable ? 0 : undefined}
>
  <span
    class="grid shrink-0 place-items-center overflow-hidden rounded-[4px]"
    style:width={`${size}px`}
    style:height={`${size}px`}
    data-testid="action-icon"
  >
    {#if icon !== null}
      <img
        src={icon}
        alt=""
        width={size}
        height={size}
        loading="lazy"
        decoding="async"
        class="h-full w-full"
      />
    {:else if action === REFUND_KEY}
      <Glyph name="refresh" size={Math.round(size * 0.7)} class="text-muted" />
    {/if}
  </span>
  <span class="min-w-0">{@render children()}</span>
</span>
