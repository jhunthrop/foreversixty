<!-- web/src/components/ui/PaperdollFrame.svelte -->
<!-- The in-game paperdoll on the owner's frame (round 5, design/specs/2026-11-09-directions.md
     Paperdoll; the mock is design/mocks/directions/paperdoll.py). Every socket, the medallion, the
     title bar and the centre are boxes in percent of the frame (lib/ui/paperdoll-frame.ts), so it
     scales with its width.
     - Slots: the item's real icon inset 6 % in a 2 px border of the client's quality colour; an
       empty slot draws nothing, so the frame's own glyph shows; a two-hander shows greyed in the
       off hand, as the client draws it. Each filled slot is a button; hover, focus or a tap
       selects it (`onselect`).
     - An icon that fails to load (a 404) draws the dark socket in its place, never a broken image.
     - Medallion: the character's Battle.net head render (`avatar`), else the painted seal; a render
       that fails falls back to the seal in the same box.
     - Title bar: one line, the title (the heading, alone), its meta, the figure and its unit at the right.
     - Centre: the caller's `centre` snippet for the selected slot (its client tooltip, the sim link).
       With `render` (signed in: the character's Battle.net render, `characters.render_url`) the
       model stands in the centre as the client's pane shows it, the snippet at its foot; a render
       that fails to load leaves the plain centre. The frame shows each item's tooltip in place, so
       it carries `data-inline-tip` and the page's floating tooltip host leaves its slots alone.
     The frame stays a frame while a slot is at least MIN_SLOT_PX (36 px) wide, a frame of 423 px:
     narrower (every phone), a container query swaps in the list form, the header strip and the
     slots in two columns on the panel, and the frame image is never fetched there. Cells are
     DollCells (lib/bis/doll-view.ts for BiS, lib/gear/doll-cells.ts for a worn set).
     - `form="list"` draws the list form at every width (the sim's saved-runs strip, a compact strip
       and not a gear panel; every gear panel is the frame); `head={false}` drops its strip where the page already
       names the character, and the slots sit straight on the page's panel; `facts` gives each list
       row the cell's line 2 after the slot and its enchant in the client's green ("Head · Item
       level 61 · Crusader") and its gain at the right. With no `centre`, or `listCentre={false}`
       (the centre in the frame only), the page's own tooltip host opens each listed item by the
       row it was tapped on. -->
<script lang="ts">
  import { COSMETIC_SLOTS, FRAME_SLOTS } from '../../lib/gear/slots';
  import { iconUrl } from '../../lib/ui/assets';
  import type { Snippet } from 'svelte';
  import { PAPERDOLL_FRAME } from '../../lib/ui/assets';
  import {
    CENTRE_BOX,
    FRAME_SIZE,
    MEDALLION_BOX,
    TITLE_BOX,
    boxStyle,
    medallionSrc,
    slotBox,
  } from '../../lib/ui/paperdoll-frame';
  import type { DollCell } from '../../lib/ui/paperdoll';
  import { rarityBorderColorFor, rarityClassFor } from '../../lib/planner/items';
  import { FACT_SEPARATOR, type ItemRef } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import { SLOT_DISPLAY_LABELS } from '../../lib/bis/slot-display-labels';
  import { gearCopy } from '../../lib/gear/copy';
  import type { Slot } from '../../lib/planner/types';

  let {
    cells,
    build,
    title,
    meta,
    figure,
    avatar = null,
    render = null,
    heading = 'h2',
    selected = $bindable(null),
    onselect,
    centre,
    shortTitle,
    label,
    form = 'auto',
    head = true,
    facts = false,
    listCentre = true,
    testid = 'paperdoll-frame',
  }: {
    cells: readonly DollCell[];
    build: string;
    /** The title bar: "Beast Mastery Hunter". */
    title: string;
    /** After the title, small: "20 to 29 · Alliance". */
    meta?: string;
    /** At the bar's right: the set's figure and its unit alone ("91.0", "DPS"). */
    figure?: { value: string; unit: string };
    /** The character's Battle.net head render URL (`characters.avatar_url`); null for the seal. */
    avatar?: string | null;
    /** The character's Battle.net full render URL (`characters.render_url`); null for none. */
    render?: string | null;
    /** The title bar's heading level: `h1` where the bar is the page's title (BiS). */
    heading?: 'h1' | 'h2';
    selected?: string | null;
    onselect?: (slot: string) => void;
    /** The centre panel for the selected cell (undefined when nothing is selected). */
    /** Its second argument says where it renders: the frame's centre, or under the phone's list. */
    centre?: Snippet<[DollCell | undefined, 'frame' | 'list']>;
    /** The phone strip's one-line title when `title` is too long for it ("Beast Mastery Hunter");
     *  the rest of `title` stays in the heading for assistive tech and crawlers. */
    shortTitle?: string;
    /** The group's accessible name; defaults to the title. */
    label?: string;
    /** `list`: the list form at every width, the frame never drawn. */
    form?: 'auto' | 'list';
    /** The list form's header strip (medallion, title, figure). */
    head?: boolean;
    /** List rows carry the cell's line 2, its enchant and its gain. */
    facts?: boolean;
    /** False: the centre shows in the frame only; the list form's items open the page's tooltip. */
    listCentre?: boolean;
    testid?: string;
  } = $props();

  let avatarFailed = $state(false);
  let renderFailed = $state(false);
  const stripShown = $derived(shortTitle !== undefined && title.startsWith(shortTitle) ? shortTitle : title);
  const stripRest = $derived(title.slice(stripShown.length));
  /** Icons that failed to load draw the dark socket, never a broken image. */
  let failedIcons = $state<Record<string, true>>({});

  function markFailed(name: string | undefined): void {
    if (name !== undefined) failedIcons = { ...failedIcons, [name]: true };
  }
  const model = $derived(render && !renderFailed ? render : null);
  const medallion = $derived(medallionSrc(avatar, avatarFailed));
  const bySlot = $derived(new Map(cells.map((cell) => [cell.slot, cell])));
  const selectedCell = $derived(selected === null ? undefined : bySlot.get(selected));
  const frameVar = `--pd-frame:image-set(url('${PAPERDOLL_FRAME.x1}') 1x, url('${PAPERDOLL_FRAME.x2}') 2x)`;

  /** What a socket shows: its own item, or the two-hander it is covered by (greyed). */
  function shown(cell: DollCell | undefined): { item: ItemRef; covered: boolean } | null {
    if (cell?.item) return { item: cell.item, covered: false };
    if (cell?.coveredBy) return { item: cell.coveredBy, covered: true };
    return null;
  }

  function choose(slot: string): void {
    if (selected === slot) return;
    selected = slot;
    onselect?.(slot);
  }

  function name(cell: DollCell, face: { item: ItemRef; covered: boolean }): string {
    return face.covered ? `${cell.label}: ${face.item.name}, two-handed` : `${cell.label}: ${face.item.name}`;
  }
</script>

{#snippet icon(face: { item: ItemRef; covered: boolean })}
  {@const src = iconUrl(build, face.item.icon)}
  {#if src !== null && !failedIcons[face.item.icon ?? '']}
    <img
      class="pdf-icon"
      class:pdf-covered={face.covered}
      {src}
      alt=""
      loading="lazy"
      decoding="async"
      style:--qb={rarityBorderColorFor(face.item.quality)}
      data-icon={face.item.icon}
      onerror={() => markFailed(face.item.icon)}
    />
  {:else}
    <span
      class="pdf-icon pdf-socket"
      data-icon-socket
      class:pdf-covered={face.covered}
      style:--qb={rarityBorderColorFor(face.item.quality)}
    ></span>
  {/if}
{/snippet}

{#snippet medal()}
  <img
    class="pdf-medal-img"
    src={medallion}
    alt=""
    width="128"
    height="128"
    onerror={() => (avatarFailed = true)}
  />
{/snippet}

{#snippet figureView()}
  {#if figure}<b class="pdf-figure">{figure.value} <small>{figure.unit}</small></b>{/if}
{/snippet}

{#snippet rowFacts(cell: DollCell)}
  <span class="pdf-sub"
    >{#if cell.source}<span class="pdf-fact-slot">{cell.label}{FACT_SEPARATOR}</span><span
        data-testid="paperdoll-list-line2-{cell.slot}">{cell.source}</span
      >{:else}{cell.label}{/if}</span
  >
  {#if cell.enchant}<span class="pdf-sub pdf-enchant" data-testid="paperdoll-list-enchant-{cell.slot}"
      >{cell.enchant}</span
    >{/if}
{/snippet}

{#snippet rowGain(cell: DollCell)}
  {#if cell.gain}
    <span class="pdf-gain" data-tone={cell.gain.tone} data-testid="paperdoll-list-gain-{cell.slot}">
      {#if cell.gain.basis?.item}{@render icon({ item: cell.gain.basis.item, covered: false })}{/if}
      <b>{cell.gain.value}{cell.gain.unit ? ` ${cell.gain.unit}` : ''}</b>
    </span>
  {/if}
{/snippet}

<div
  class="pdf"
  class:pdf-listonly={form === 'list'}
  data-testid={testid}
  role="group"
  aria-label={label ?? title}
  data-inline-tip={centre !== undefined && listCentre ? '' : undefined}
>
  {#if form !== 'list'}
    <!-- The frame (423 px and wider). -->
    <div
      class="pdf-frame"
      data-inline-tip={centre !== undefined && !listCentre ? '' : undefined}
      style="aspect-ratio:{FRAME_SIZE.width}/{FRAME_SIZE.height};{frameVar}"
      data-region={REGION.slots}
    >
      {#each FRAME_SLOTS as slot (slot)}
        {@const cell = bySlot.get(slot)}
        {@const face = shown(cell)}
        {#if cell && face}
          <button
            type="button"
            class="pdf-slot"
            class:pdf-sel={slot === selected}
            style={boxStyle(slotBox(slot))}
            aria-label={name(cell, face)}
            aria-pressed={slot === selected}
            data-bis-item={face.item.id}
            data-testid="paperdoll-frame-slot-{slot}"
            onmouseenter={() => choose(slot)}
            onfocus={() => choose(slot)}
            onclick={() => choose(slot)}>{@render icon(face)}</button
          >
        {/if}
      {/each}
      <span class="pdf-medal" style={boxStyle(MEDALLION_BOX)} data-region={REGION.paperdollMedallion}
        >{@render medal()}</span
      >
      <div class="pdf-title" style={boxStyle(TITLE_BOX)}>
        <svelte:element this={heading} class="pdf-title-text">{title}</svelte:element>
        {#if meta}<small class="pdf-meta">{meta}</small>{/if}
        {@render figureView()}
      </div>
      <div
        class="pdf-centre"
        class:pdf-modelled={model !== null}
        style={boxStyle(CENTRE_BOX)}
        data-region={REGION.paperdollCentre}
      >
        {#if model !== null}<img
            class="pdf-model"
            src={model}
            alt=""
            decoding="async"
            onerror={() => (renderFailed = true)}
          />{/if}
        {#if centre}{@render centre(selectedCell, 'frame')}{/if}
      </div>
    </div>
  {/if}

  <!-- The list form (narrower than 423 px): the header strip and two columns of slots. -->
  <div class="pdf-list">
    {#if head}<div class="pdf-strip">
        <span class="pdf-strip-medal">{@render medal()}</span>
        <!-- The strip repeats the title bar's heading for the phone: one heading element per page in
           the served HTML, so this one is a heading by role (only one of the two ever shows). -->
        <div class="pdf-strip-title" role="heading" aria-level={heading === 'h1' ? 1 : 2}>
          {stripShown}{#if stripRest}<span class="sr-only">{stripRest}</span>{/if}
        </div>
        {@render figureView()}
      </div>{/if}
    <ul class="pdf-items" class:pdf-headless={!head} class:pdf-facts={facts}>
      {#each FRAME_SLOTS.filter((slot) => !COSMETIC_SLOTS.has(slot)) as slot (slot)}
        {@const cell = bySlot.get(slot)}
        {@const face = shown(cell)}
        <li>
          {#if cell && face}
            <button
              type="button"
              class="pdf-item"
              class:pdf-sel={slot === selected && centre !== undefined && listCentre}
              aria-pressed={slot === selected}
              data-bis-item={face.item.id}
              data-testid="paperdoll-list-slot-{slot}"
              onclick={() => choose(slot)}
            >
              <span class="pdf-cell">{@render icon(face)}</span>
              <span class="min-w-0">
                {#if face.covered}
                  <span class="pdf-sub">{cell.label}</span><span class="pdf-sub"
                    >{gearCopy.paperdoll.twoHanded}</span
                  >
                {:else}
                  <span class="pdf-name {rarityClassFor(face.item.quality)}">{face.item.name}</span>
                  {#if facts}{@render rowFacts(cell)}{:else}<span class="pdf-sub">{cell.label}</span>{/if}
                {/if}
              </span>
              {#if facts && !face.covered}{@render rowGain(cell)}{/if}
            </button>
          {:else}
            <span class="pdf-item pdf-empty">
              <span class="pdf-cell"></span>
              <span class="pdf-sub">{cell?.label ?? SLOT_DISPLAY_LABELS[slot as Slot]}</span>
            </span>
          {/if}
        </li>
      {/each}
    </ul>
    {#if centre && listCentre && selectedCell}<div class="pdf-list-centre">
        {@render centre(selectedCell, 'list')}
      </div>{/if}
  </div>
</div>

<style>
  .pdf {
    container-type: inline-size;
    width: 100%;
  }
  .pdf-frame,
  .pdf-list {
    display: none;
  }
  .pdf-listonly .pdf-list {
    display: block;
  }
  /* MIN_FRAME_PX (lib/ui/paperdoll-frame.ts): a slot is 36 px at 423 px. */
  @container (min-width: 423px) {
    .pdf-frame {
      position: relative;
      display: block;
      width: 100%;
      background: var(--pd-frame) 0 0 / 100% 100% no-repeat;
    }
  }
  @container (max-width: 422.98px) {
    .pdf-list {
      display: block;
    }
  }
  .pdf-slot {
    position: absolute;
    display: block;
    padding: 0;
    border: 0;
    background: none;
    cursor: pointer;
  }
  .pdf-icon {
    position: absolute;
    inset: 6%;
    width: 88%;
    height: 88%;
    box-sizing: border-box;
    border: 2px solid var(--qb);
    border-radius: var(--radius-icon);
    box-shadow: 0 0 0 1px black;
    background: var(--color-bg);
  }
  .pdf-covered {
    filter: grayscale(1) brightness(0.55);
    border-color: var(--color-muted-card);
  }
  /* The selected slot takes the selected state's accent ring, held off the item's quality border
     by a black gap and lifted off the teal frame by a white hairline, so a rare's blue never reads
     as the ring. Keyboard focus keeps the cream ring, so a focused slot reads apart from the one the
     centre shows. */
  .pdf-slot.pdf-sel::after,
  .pdf-slot:focus-visible::after {
    content: '';
    position: absolute;
    inset: -9%;
    border: 3px solid var(--color-selected-mark);
    border-radius: var(--radius-field);
    box-shadow:
      inset 0 0 0 2px black,
      0 0 0 1px var(--color-strong),
      0 0 12px 2px color-mix(in srgb, var(--color-selected-mark) 60%, transparent);
  }
  .pdf-slot:focus-visible::after {
    border-color: var(--color-cream);
  }
  .pdf-slot:focus-visible {
    outline: none;
  }
  .pdf-medal {
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .pdf-medal-img {
    width: 92%;
    height: 92%;
    border-radius: 50%;
    object-fit: cover;
  }
  .pdf-title {
    position: absolute;
    display: flex;
    align-items: center;
    gap: 2cqw;
    margin: 0;
    padding: 0 2.2cqw;
    color: var(--color-cream);
    font: 400 3.4cqw/1 var(--font-title);
    letter-spacing: var(--tracking-heading);
    text-shadow: 0 1px 2px black;
    white-space: nowrap;
  }
  .pdf-title-text {
    margin: 0;
    font: inherit;
  }
  .pdf-meta {
    font: 600 2.3cqw var(--font-body);
    color: var(--color-text);
  }
  .pdf-figure {
    margin-left: auto;
    font: 600 3.2cqw var(--font-body);
    color: var(--color-strong);
    text-shadow: 0 1px 2px black;
    white-space: nowrap;
  }
  .pdf-figure small {
    font-size: 0.66em;
    color: var(--color-text);
  }
  .pdf-centre {
    position: absolute;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 3cqw;
    padding: 4cqw;
  }
  /* Signed in: the render fills the centre behind the snippet, its feet near the foot (the figure is
     about 57 % of the render's height, centred at 51 % across), and the snippet stands over it. */
  .pdf-modelled {
    justify-content: flex-end;
    overflow: hidden;
  }
  .pdf-model {
    position: absolute;
    left: 50%;
    top: 54%;
    height: 150%;
    max-width: none;
    transform: translate(-51%, -50%);
    pointer-events: none;
  }
  .pdf-modelled > :global(:not(.pdf-model)) {
    position: relative;
  }
  /* The list form: a header strip on the panel, then the slots in two columns. */
  .pdf-strip {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 8px 12px;
    border: 1px solid black;
    border-radius: var(--radius-plate) var(--radius-plate) 0 0;
    background: var(--color-nav-bar);
    box-shadow: 0 0 0 1px var(--color-panel-edge);
  }
  .pdf-strip-medal {
    display: flex;
    flex: none;
    width: 40px;
    height: 40px;
  }
  .pdf-strip-medal .pdf-medal-img {
    width: 100%;
    height: 100%;
  }
  /* The title and an item's name wrap rather than truncate, as the list form's mock draws them. */
  .pdf-strip-title {
    min-width: 0;
    margin: 0;
    font: 400 19px/1.2 var(--font-title);
    color: var(--color-cream);
  }
  .pdf-strip .pdf-figure {
    font-size: 17px;
  }
  .pdf-items {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 6px 10px;
    margin: 0;
    padding: 12px;
    list-style: none;
    border: 1px solid black;
    border-top: 0;
    border-radius: 0 0 var(--radius-plate) var(--radius-plate);
    background: var(--color-raised);
    box-shadow: 0 0 0 1px var(--color-panel-edge);
  }
  /* Headless, the list is the page panel's own content: no second box inside it. */
  .pdf-headless {
    padding: 0;
    border: 0;
    border-radius: 0;
    background: none;
    box-shadow: none;
  }
  .pdf-item {
    display: grid;
    grid-template-columns: 44px minmax(0, 1fr);
    gap: 8px;
    align-items: center;
    width: 100%;
    min-height: 44px;
    padding: 0;
    border: 0;
    background: none;
    color: inherit;
    font: inherit;
    text-align: left;
  }
  .pdf-cell {
    position: relative;
    width: 44px;
    height: 44px;
    border-radius: var(--radius-icon);
    background: var(--color-bg);
    box-shadow: inset 0 0 0 1px var(--color-line);
  }
  .pdf-cell .pdf-icon {
    inset: 3px;
    width: 38px;
    height: 38px;
  }
  .pdf-item.pdf-sel .pdf-cell {
    outline: 2px solid var(--color-selected-mark);
    outline-offset: 2px;
  }
  .pdf-name {
    display: block;
    font-size: 13px;
    font-weight: 600;
    line-height: 16px;
    overflow-wrap: anywhere;
  }
  .pdf-sub {
    display: block;
    overflow: hidden;
    font-size: 12px;
    line-height: 16px;
    color: var(--color-muted);
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  /* The enchant has its own green line under the item level, as the client's tooltip prints it,
     and wraps rather than lose a word. */
  .pdf-facts .pdf-enchant {
    white-space: normal;
    color: var(--color-rarity-uncommon);
  }
  /* A phone's half column keeps the facts and drops the slot word before them (the name says it). */
  @container (min-width: 360px) and (max-width: 559.98px) {
    .pdf-fact-slot {
      display: none;
    }
  }
  /* Too narrow for two whole names a row (a strip inside a phone's section): one column. */
  @container (max-width: 359.98px) {
    .pdf-facts {
      grid-template-columns: minmax(0, 1fr);
    }
  }
  .pdf-facts .pdf-item:has(.pdf-gain) {
    grid-template-columns: 44px minmax(0, 1fr) auto;
  }
  .pdf-gain {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    white-space: nowrap;
    color: var(--color-gold);
  }
  .pdf-gain[data-tone='kill'] {
    color: var(--color-kill);
  }
  .pdf-gain[data-tone='muted'] {
    color: var(--color-muted);
  }
  .pdf-gain .pdf-icon {
    position: relative;
    inset: auto;
    width: 22px;
    height: 22px;
    border-width: 1px;
  }
  .pdf-list-centre {
    margin-top: 10px;
  }
</style>
