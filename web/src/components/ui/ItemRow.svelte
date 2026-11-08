<!-- web/src/components/ui/ItemRow.svelte -->
<!-- An item as a row (web layout spec 2026-11-05 §3A.2, §4.3), the Row preset every item list
     uses: an optional slot label, the ItemIcon in its quality frame, the name (body face 700 in the
     AA-safe quality colour) over a muted source line, and at the right the worn item's 28 px icon
     and a gain Figure. One level of emphasis: the name; then the gain; the source is muted. The icon
     and name carry `data-bis-item`, so the page's one tooltip host (BisTooltipHost) opens the real
     item on hover, focus and long-press. With no item it reads the empty slot's reason. With
     `onselect` the whole row is one button (aria-pressed, the gold selected ground), its name no
     longer a tab stop of its own. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { sharedPaperdollNothingBetter } from '../../lib/copy/shared/paperdoll.generated';
  import { rarityClassFor } from '../../lib/planner/items';
  import type { GainFigure, ItemRef } from '../../lib/ui/primitives';
  import Figure from './Figure.svelte';
  import ItemIcon from './ItemIcon.svelte';

  let {
    item,
    build,
    slotLabel,
    source,
    sourceTitle,
    sourceContent,
    detail,
    gain,
    worn,
    badge,
    iconSize = 40,
    emptyText = sharedPaperdollNothingBetter,
    emptyDetail,
    unknown = false,
    eager = false,
    focusableName = true,
    bare = false,
    onselect,
    selected = false,
    trail,
    testid,
  }: {
    item: ItemRef | null;
    build: string;
    /** The slot's name, a 60 px column at the left (above the name below 400 px). */
    slotLabel?: string;
    /** One muted line under the name: where it comes from. */
    source?: string;
    /** The full source line when the visible one gives way (a long boss name). */
    sourceTitle?: string;
    /** Replaces `source` (a line with a source-kind glyph). */
    sourceContent?: Snippet;
    /** A second muted line (short stats, "usable now"). */
    detail?: Snippet;
    gain?: GainFigure;
    worn?: ItemRef;
    badge?: 'upgrade' | 'worn';
    iconSize?: 28 | 32 | 36 | 40;
    emptyText?: string;
    emptyDetail?: string;
    /** The worn id is missing from our item data. */
    unknown?: boolean;
    eager?: boolean;
    /** The name is its own tab stop for the tooltip; off inside a roving-focus container. */
    focusableName?: boolean;
    /** No row padding or hairline: the caller (a paperdoll cell, an inspector list) draws its own. */
    bare?: boolean;
    /** The row picks its item (a list beside an inspector): the whole row is one button. */
    onselect?: () => void;
    selected?: boolean;
    /** Replaces the worn icon and gain at the right (a state Tag). */
    trail?: Snippet;
    testid?: string;
  } = $props();

  /** The name is a tab stop of its own only when the row is not itself the button. */
  const nameStop = $derived(focusableName && !unknown && onselect === undefined);
</script>

<div class="ui-item-row-box min-w-0">
  <svelte:element
    this={onselect ? 'button' : 'div'}
    type={onselect ? 'button' : undefined}
    aria-pressed={onselect ? (selected ? 'true' : 'false') : undefined}
    onclick={onselect}
    class={[
      'ui-item-row grid w-full min-w-0 items-center gap-x-3 text-left',
      !bare && 'border-line-soft min-h-[var(--row-h)] border-b px-4 py-2 last:border-b-0',
      slotLabel ? 'ui-item-row--slot' : 'grid-cols-[auto_minmax(0,1fr)_auto]',
      onselect &&
        'hover:bg-hover active:bg-pressed focus-visible:outline-gold cursor-pointer focus-visible:outline-2 focus-visible:-outline-offset-2',
      selected && 'bg-selected shadow-[inset_3px_0_0_var(--color-gold)]',
    ]}
    data-testid={testid}
    data-item-id={item?.id}
  >
    {#if slotLabel}<span class="ui-item-row-label text-muted text-[13px] leading-[18px] font-semibold"
        >{slotLabel}</span
      >{/if}
    <span data-bis-item={item && !unknown ? item.id : undefined}>
      <ItemIcon icon={item?.icon} quality={item?.quality} {build} size={iconSize} {badge} {unknown} {eager} />
    </span>
    <span class="min-w-0">
      {#if item}
        <span
          class="block text-[15px] leading-5 font-bold {unknown
            ? 'text-muted'
            : rarityClassFor(item.quality)}"
          data-bis-item={unknown ? undefined : item.id}
          data-testid={`item-hover-${item.id}`}
          tabindex={nameStop ? 0 : undefined}
          role={nameStop ? 'button' : undefined}
          aria-haspopup={nameStop ? 'true' : undefined}>{item.name}</span
        >
      {:else}
        <span class="text-muted block text-[15px] leading-5">{emptyText}</span>
      {/if}
      {#if sourceContent}<span
          class="ui-item-row-src text-muted block text-[13px] leading-[18px]"
          title={sourceTitle}>{@render sourceContent()}</span
        >
      {:else if source}<span
          class="ui-item-row-src text-muted block text-[13px] leading-[18px]"
          title={sourceTitle}>{source}</span
        >
      {:else if !item && emptyDetail}<span class="text-muted block text-[13px] leading-[18px]"
          >{emptyDetail}</span
        >{/if}
      {#if detail}<span class="text-muted block text-[13px] leading-[18px]">{@render detail()}</span>{/if}
    </span>
    <span class="flex items-center gap-3 justify-self-end">
      {#if trail}{@render trail()}{/if}
      {#if worn}
        <span data-bis-item={worn.id} title={`You wear ${worn.name}`} data-testid="item-row-worn">
          <ItemIcon icon={worn.icon} quality={worn.quality} {build} size={28} />
        </span>
      {/if}
      {#if gain}<Figure
          value={gain.value}
          unit={gain.unit}
          caption={gain.caption}
          tone={gain.tone ?? 'strong'}
          size="sm"
          align="end"
        />{/if}
    </span>
  </svelte:element>
</div>

<style>
  .ui-item-row-box {
    container-type: inline-size;
  }
  .ui-item-row--slot {
    grid-template-columns: 60px auto minmax(0, 1fr) auto;
  }
  /* A narrow column (a phone, a 1024 doll column): the label sits above the name, 12 px caps. */
  @container (max-width: 399px) {
    .ui-item-row--slot {
      grid-template-columns: auto minmax(0, 1fr) auto;
    }
    .ui-item-row-label {
      grid-column: 2 / -1;
      grid-row: 1;
      font-size: 12px;
      line-height: 16px;
      font-weight: 700;
      letter-spacing: var(--tracking-label);
      text-transform: uppercase;
    }
    .ui-item-row--slot > :global(:nth-child(2)) {
      grid-row: 1 / span 2;
      grid-column: 1;
    }
  }
</style>
