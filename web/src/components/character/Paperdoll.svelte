<!-- web/src/components/character/Paperdoll.svelte -->
<!-- The character-screen gear panel (design/specs/2026-10-31-paperdoll.md): the client's
     character frame drawn with the site's parts. Head to wrist down the left, hands to trinket 2
     down the right, sockets on the outer edge and names toward the ringed crest in the centre;
     main hand and off hand under the columns and ranged centred below, as the Discord sheet draws
     them. Under a 640 px container it folds to the identity on top and two mirrored columns, each
     pair followed by its full-width upgrade rows.

     It draws the frame only (the stage, the slots, the centre and the note); each call site wraps
     it in its own Card. Every known item and every pick is an ItemHover: hover, focus or tap
     opens the shared tooltip, the right column's toward the centre (`align: 'end'`). The layout,
     the upgrade rules, the item-level ruling and every word come from lib/gear/paperdoll.ts. -->
<script lang="ts">
  // The frame's styles: a plain stylesheet, not a scoped <style>, so the character page's island
  // ships them though its SSR pass renders no frame (src/styles/paperdoll.css).
  import '../../styles/paperdoll.css';
  import { factionLogoSrc } from '../../lib/faction-mark';
  import {
    averageItemLevel,
    cellLabel,
    hyphenate,
    paperdollLayout,
    paperdollNote,
    pickLabel,
    slotLabel,
    unknownCount,
    unknownName,
    type PaperdollCell,
    type PaperdollIdentity,
    type PaperdollItemRef,
    type PaperdollSlot,
    type PaperdollSlotKey,
    type PaperdollUpgrade,
  } from '../../lib/gear/paperdoll';
  import { gearCopy } from '../../lib/gear/copy';
  import { rarityBorderColorFor, rarityClassFor } from '../../lib/planner/items';
  import { dataUrl } from '../../lib/planner/load';
  import { classColorVar } from '../../lib/report/format';
  import ItemHover from '../ItemHover.svelte';
  import ClassCrestRing from './ClassCrestRing.svelte';

  let {
    identity,
    slots,
    build,
    classSlug,
    loading = false,
    failed = false,
    testid = 'paperdoll',
  }: {
    identity: PaperdollIdentity;
    slots: readonly PaperdollSlot[];
    /** The data build the icons and tooltips are read from. */
    build: string;
    /** The class whose item file the tooltips read. */
    classSlug: string;
    /** The items file is pending: shimmer sockets and bars, the identity drawn at once. */
    loading?: boolean;
    /** The items file failed: plain sockets, names in the body colour, the note says so. */
    failed?: boolean;
    testid?: string;
  } = $props();

  const copy = gearCopy.paperdoll;
  const cells = $derived(paperdollLayout(slots));
  const itemLevel = $derived(loading || failed ? null : averageItemLevel(slots));
  const note = $derived(loading ? null : paperdollNote(unknownCount(slots), failed));
  /** Rows grow to 64 px when any slot carries a second line (the character page, the report). */
  const lined = $derived(slots.some((slot) => slot.upgrade !== undefined || slot.detail !== undefined));
  const colour = $derived(classColorVar(identity.classSlug));
  const descriptor = $derived(
    [identity.race, identity.spec, classWord(identity.classSlug)].filter((word) => word !== '').join(' '),
  );

  /** "warrior" -> "Warrior". */
  function classWord(slug: string): string {
    return slug === '' ? '' : slug.charAt(0).toUpperCase() + slug.slice(1);
  }

  /** The client's empty-slot texture for a slot (§2): pairs share one, ranged is the relic's for
   *  the relic classes. */
  function textureFor(slot: PaperdollSlotKey): string {
    const key =
      slot === 'finger1' || slot === 'finger2'
        ? 'finger'
        : slot === 'trinket1' || slot === 'trinket2'
          ? 'trinket'
          : slot === 'ranged' && slotLabel('ranged', classSlug) === copy.relic
            ? 'relic'
            : slot;
    return `/icons/hd/slots/${key}.webp`;
  }

  /** A failed image keeps the plain socket (§8), never the broken-image glyph. */
  function hideBroken(event: Event): void {
    (event.currentTarget as HTMLImageElement).style.visibility = 'hidden';
  }

  function iconSrc(icon: string): string {
    return dataUrl(build, `icons/${icon}.webp`);
  }

  function quality(ref: { quality: number | null }): string {
    return ref.quality === null ? 'pd-plain' : rarityClassFor(ref.quality);
  }

  function ring(ref: { quality: number | null }): string {
    return ref.quality === null ? '' : `--q: ${rarityBorderColorFor(ref.quality)}`;
  }

  function place(cell: PaperdollCell): string {
    return `--r: ${cell.row}; --rp: ${cell.phoneRow}`;
  }

  function sideClass(cell: PaperdollCell): string {
    const weapon = cell.slot === 'main_hand' || cell.slot === 'off_hand' ? ' pd-weapon' : '';
    return `${cell.side === 'centre' ? 'pd-mid' : `pd-${cell.side}`}${weapon}`;
  }

  function align(cell: PaperdollCell): 'start' | 'end' {
    return cell.side === 'right' ? 'end' : 'start';
  }

  function pickOf(upgrade: PaperdollUpgrade): PaperdollItemRef | null {
    return upgrade.kind === 'best' ? null : upgrade.pick;
  }
</script>

{#snippet socketIcon(icon: string | null)}
  {#if icon !== null}<img
      src={iconSrc(icon)}
      alt=""
      width="44"
      height="44"
      loading="lazy"
      decoding="async"
      onerror={hideBroken}
    />{/if}
{/snippet}

{#snippet line2(cell: PaperdollCell)}
  {@const upgrade = cell.data.upgrade}
  {#if !loading && upgrade !== undefined}
    <span class="pd-det" data-testid="paperdoll-line">
      {#if upgrade.kind === 'best'}
        <span class="pd-muted" data-testid="paperdoll-best">{copy.best}</span>
      {:else}
        {@const pick = upgrade.pick}
        <ItemHover
          itemId={pick.id}
          {classSlug}
          {build}
          align={align(cell)}
          class="pd-pick"
          label={pickLabel(cell.slot, upgrade, classSlug)}
          testid={`paperdoll-pick-${cell.slot}`}
        >
          <span class="pd-pick-sock" style={ring(pick)}>{@render socketIcon(pick.icon)}</span>
          <span class="pd-pick-name {quality(pick)}">{pick.name}</span>
        </ItemHover>
        {#if upgrade.kind === 'gain'}
          <b class="pd-gain" data-testid="paperdoll-gain">{copy.gain(upgrade.points)}</b>
        {:else}
          <span class="pd-muted" data-testid="paperdoll-unscored">{copy.notScored}</span>
        {/if}
      {/if}
    </span>
  {:else if !loading && cell.data.detail !== undefined}
    <span class="pd-det pd-muted" data-testid="paperdoll-detail">{cell.data.detail}</span>
  {/if}
{/snippet}

{#snippet slotCell(cell: PaperdollCell)}
  {@const data = cell.data}
  {@const item = data.item}
  <div
    class="pd-cell {sideClass(cell)}"
    class:pd-has-line={!loading && (data.upgrade !== undefined || data.detail !== undefined)}
    style={place(cell)}
    role="listitem"
    data-testid={`paperdoll-slot-${cell.slot}`}
    data-state={cell.twoHandOff ? 'two-hand' : loading ? 'loading' : data.state}
  >
    {#if loading}
      <div class="pd-slot" aria-busy="true">
        <span class="pd-sock pd-shimmer"></span>
        <span class="pd-txt"><span class="pd-bar pd-shimmer"></span></span>
      </div>
    {:else if data.state === 'item' && item !== undefined}
      <ItemHover
        itemId={item.id}
        {classSlug}
        {build}
        hostClass="pd-host"
        align={align(cell)}
        class="pd-slot pd-known"
        label={cellLabel(data, classSlug)}
        testid={`paperdoll-item-${cell.slot}`}
      >
        <span class="pd-sock" class:pd-bare={item.icon === null} style={ring(item)}
          >{@render socketIcon(item.icon)}</span
        >
        <span class="pd-txt"><span class="pd-name {quality(item)}">{hyphenate(item.name)}</span></span>
      </ItemHover>
    {:else if data.state === 'unknown'}
      <div class="pd-slot" aria-label={cellLabel(data, classSlug)} role="img">
        <span class="pd-sock pd-unknown"
          ><img
            src={iconSrc('inv_misc_questionmark')}
            alt=""
            width="44"
            height="44"
            loading="lazy"
            decoding="async"
            onerror={hideBroken}
          /></span
        >
        <span class="pd-txt"
          ><span class="pd-name pd-id" data-testid="paperdoll-unknown"
            >{unknownName(cell.slot, data.unknownId ?? 0, classSlug)}</span
          ></span
        >
      </div>
    {:else}
      <div class="pd-slot" aria-label={cellLabel(data, classSlug, cell.twoHandOff)} role="img">
        <span class="pd-sock pd-empty"
          ><img
            src={textureFor(cell.slot)}
            alt=""
            width="44"
            height="44"
            loading="lazy"
            decoding="async"
            onerror={hideBroken}
          /></span
        >
        <span class="pd-txt"
          ><span class="pd-name pd-word" data-testid="paperdoll-empty"
            >{cell.twoHandOff ? copy.twoHand : slotLabel(cell.slot, classSlug)}</span
          ></span
        >
      </div>
    {/if}
    {@render line2(cell)}
  </div>
{/snippet}

{#snippet upgradeRow(cell: PaperdollCell)}
  {@const upgrade = cell.data.upgrade}
  {#if upgrade !== undefined && !loading}
    {@const pick = pickOf(upgrade)}
    {#if pick !== null}
      <div class="pd-up-wrap" style={place(cell)} role="listitem" data-testid={`paperdoll-up-${cell.slot}`}>
        <ItemHover
          itemId={pick.id}
          {classSlug}
          {build}
          hostClass="pd-host"
          class="pd-up"
          label={pickLabel(cell.slot, upgrade, classSlug)}
          testid={`paperdoll-uprow-${cell.slot}`}
        >
          <span class="pd-up-slot">{slotLabel(cell.slot, classSlug)}</span>
          <span class="pd-up-sock" style={ring(pick)}>{@render socketIcon(pick.icon)}</span>
          <span class="pd-up-name {quality(pick)}">{hyphenate(pick.name)}</span>
          {#if upgrade.kind === 'gain'}
            <b class="pd-gain">{copy.gain(upgrade.points)}</b>
          {:else}
            <span class="pd-muted">{copy.notScored}</span>
          {/if}
        </ItemHover>
      </div>
    {/if}
  {/if}
{/snippet}

<div class="pd" data-testid={testid}>
  <div class="pd-stage" class:pd-lined={lined}>
    <div class="pd-grid" role="list" aria-label={gearCopy.listLabel} aria-busy={loading ? 'true' : undefined}>
      <div class="pd-centre" data-testid="paperdoll-centre" role="presentation">
        <ClassCrestRing characterClass={identity.classSlug} size={96} class="pd-crest" />
        <div class="pd-who">
          {#if identity.name !== ''}
            <span class="pd-cname" style={`color: ${colour}`} data-testid="paperdoll-name"
              >{identity.name}</span
            >
          {/if}
          {#if descriptor !== ''}
            <span class="pd-desc" data-testid="paperdoll-descriptor"
              >{#if identity.faction !== null}<img
                  class="pd-fac"
                  src={factionLogoSrc(identity.faction)}
                  alt=""
                  width="16"
                  height="16"
                />{/if}{descriptor}</span
            >
          {/if}
          <span class="pd-meta">
            {#if identity.level !== null}<span class="pd-lvl" data-testid="paperdoll-level"
                >{copy.level(identity.level)}</span
              >{/if}
            {#if loading}
              <span class="pd-ilvl"
                ><span class="pd-micro">{copy.itemLevel}</span><span class="pd-figure pd-shimmer"
                ></span></span
              >
            {:else if itemLevel !== null}
              <span class="pd-ilvl" data-testid="paperdoll-ilvl"
                ><span class="pd-micro">{copy.itemLevel}</span><span class="pd-figure">{itemLevel}</span
                ></span
              >
            {/if}
          </span>
          {#if identity.source !== ''}
            <span class="pill pill-site pd-pill" data-testid="paperdoll-source">{identity.source}</span>
          {/if}
        </div>
      </div>
      {#each cells as cell (`${cell.kind}-${cell.slot}`)}
        {#if cell.kind === 'slot'}
          {@render slotCell(cell)}
        {:else}
          {@render upgradeRow(cell)}
        {/if}
      {/each}
    </div>
  </div>
  {#if note !== null}
    <p class="pd-note" data-testid="paperdoll-note">{note}</p>
  {/if}
</div>
