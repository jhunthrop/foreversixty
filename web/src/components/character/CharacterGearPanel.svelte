<!-- web/src/components/character/CharacterGearPanel.svelte -->
<!-- The profile's Gear section (Task L8.6 and its fix round 1, spec §4.5/§4.6), under the
     Forever Score card: the 17 slots of the character's newest addon export from GET
     /v1/characters/...'s `gear` (api/openapi.yaml `CharacterGear`). The section is the anchor the
     card's "Most to gain: Gear" opens (#gear, landing 96px under the top, its heading focused).

     The title line reads "Gear · +70 pts across 7 slots" (" · 1 not scored" when a pick has no
     gap); the dot after "Gear" is drawn as its ::after, glued by a no-break space, and at 480px
     and under the summary is its own line with no dot, so no line ever starts with "·" (fix
     round 2 ruling 4). 4px under it one muted 13px line sets the two totals side by side ("The
     card’s 88 counts the band’s best set. These 70 count the picks you can wear."), 16px above
     row 1. The footer holds only the BiS page link, 12px of padding above and below it.

     Each known name is marked the way /bis marks an item (`data-bis-item`, a focusable button,
     the quality colour and a 20px quality-bordered icon), and the shared BisTooltipHost opens
     the item tooltip (a mouse's hover, keyboard focus, a tap; fix round 2 ruling 2), its model
     fetched at runtime. Each icon sits in a 20px box filled #12151c that stays, empty, when the
     image fails or the item has no icon (ruling 1). Until the build's items file (quality, icon,
     two-handed) answers, the slot labels, gaps and sources draw at once and each name is held
     behind a 14px bar, so no name is ever painted white first and nothing moves when they come
     (ruling 8): the bar holds the real name, unpainted, in the name's own box, so every row
     wraps the same before and after (fix round 3 item 2); a failed file draws the names in the
     body colour.

     Layout. Below 1024px three lines a row: the slot label and the gap on one top, then worn →
     pick, then the source and chance; a row with no source ("Best for your band", no pick) has
     its note on the items' line, the dot drawn after the items. From 1024px one line: slot 104 |
     worn → pick 5fr | source 4fr | gap 88, a faint tint on hover and focus; every row is 44px at
     1440, 1280 and 1024 (ruling 3). The list spans the page column: site-column.test.ts bans a
     container width of 800px and up. Each row is named by its slot label (aria-labelledby), the
     list "Gear by slot" (ruling 9).

     `gear: []` (no export) draws the heading over one line that keeps the anchor; the caller
     draws nothing at all for an API that sends no `gear`. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { gearCopy } from '../../lib/gear/copy';
  import { SvelteSet } from 'svelte/reactivity';
  import {
    GEAR_ANCHOR,
    gearRows,
    gearTitleSummary,
    gearTotalsLine,
    type GearItemView,
    type GearRow,
  } from '../../lib/gear/panel';
  import { iconOf, itemFacts, nameClass, type GearItemTable } from '../../lib/gear/qualities';
  import type { CharacterGearSlot } from '../../lib/gear/types';
  import { fetchItemTooltipModel } from '../../lib/items/lookup';
  import { rarityBorderColorFor } from '../../lib/planner/items';
  import { dataUrl } from '../../lib/planner/load';
  import BisTooltipHost from '../BisTooltipHost.svelte';

  let {
    gear,
    total,
    cardGap = null,
    classSlug = null,
    bisHref = null,
    spec = null,
    band = null,
    preloaded = undefined,
  }: {
    /** `gear`: all 17 slots, or `[]` when the character has no export. */
    gear: readonly CharacterGearSlot[];
    /** `gear_gap_total`. */
    total: number | null;
    /** The score card's Gear gap ("88 below its top"), for the line under the title; null with
     *  no score. */
    cardGap?: number | null;
    /** The character's class ("priest"), whose item table gives the qualities and tooltips. */
    classSlug?: string | null;
    /** The spec's best-in-slot page at this band, for the footer link; none, no link. */
    bisHref?: string | null;
    /** The spec the link names ("Holy") and the band ("60"). */
    spec?: string | null;
    band?: number | null;
    /** The items file already read (a table) or failed (null); left out, the section reads it. */
    preloaded?: GearItemTable | null;
  } = $props();

  const build = activeBuild.build;

  /** The items file: undefined while it loads, null when it failed. */
  let facts = $state<GearItemTable | null | undefined>(undefined);
  $effect(() => {
    const slug = classSlug;
    if (preloaded !== undefined) {
      facts = preloaded;
      return;
    }
    if (slug === null) {
      facts = null;
      return;
    }
    if (gear.length === 0) return;
    void itemFacts(build, slug).then((loaded) => {
      if (slug === classSlug) facts = loaded;
    });
  });
  /** SSR and tests: the preloaded table draws at once. */
  const table = $derived(preloaded !== undefined ? preloaded : facts);
  const loading = $derived(table === undefined);

  /** A failed items file (null) cannot say whether the main hand is two-handed. While it loads,
   *  an empty off hand at 0 under a worn main hand is laid out as the swap (the API's 0 there is
   *  the two-hand unit's), its note held invisible: when the file answers nothing moves. */
  const rows = $derived(
    gearRows(
      gear,
      table === null ? null : table === undefined ? () => true : (id) => table.get(id)?.twoHand === true,
    ),
  );
  /** After "Gear": "+70 pts across 7 slots · 1 not scored", or nothing. */
  const summary = $derived(gearTitleSummary(gear, total));
  /** Icons whose image failed: their box stays, empty (fix round 2 ruling 1). */
  const brokenIcons = new SvelteSet<number>();

  const totalsLine = $derived(gearTotalsLine(cardGap, total));

  /** The landing (fix round 1 ruling 11): when the address is #gear, focus moves to the heading,
   *  whether the link was clicked, pressed or the page opened on it. */
  let heading = $state<HTMLHeadingElement | undefined>();
  $effect(() => {
    const target = heading;
    if (target === undefined) return;
    const focusHeading = (): void => {
      if (window.location.hash === `#${GEAR_ANCHOR}`) target.focus({ preventScroll: true });
    };
    const onClick = (event: MouseEvent): void => {
      if (event.target instanceof Element && event.target.closest(`a[href="#${GEAR_ANCHOR}"]`) !== null)
        requestAnimationFrame(focusHeading);
    };
    window.addEventListener('hashchange', focusHeading);
    document.addEventListener('click', onClick);
    if (window.location.hash === `#${GEAR_ANCHOR}`) {
      document.getElementById(GEAR_ANCHOR)?.scrollIntoView();
      focusHeading();
    }
    return () => {
      window.removeEventListener('hashchange', focusHeading);
      document.removeEventListener('click', onClick);
    };
  });

  /** Phone: [slot | gap] / [items, note] / [source]. From 1024px: [slot | items | source | gap]. */
  const ROW =
    'border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-start gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] transition-colors duration-[120ms] hover:bg-[rgba(255,255,255,0.04)] focus-within:bg-[rgba(255,255,255,0.04)] lg:grid-cols-[104px_minmax(0,5fr)_minmax(0,4fr)_88px] lg:items-center lg:py-0';
  const SLOT = 'text-muted col-start-1 row-start-1 text-[13px] leading-5';
  const LINE = 'col-span-2 row-start-2 flex min-w-0 flex-wrap items-baseline gap-x-1 gap-y-1 lg:contents';
  const ITEMS = 'min-w-0 break-words lg:col-start-2 lg:row-start-1';
  const META = 'text-muted min-w-0 text-[13px] lg:col-start-3 lg:row-start-1';
  const GAP =
    'tabular col-start-2 row-start-1 text-right font-mono text-[13px] leading-5 whitespace-nowrap lg:col-start-4 lg:row-start-1';
  /** A known name: the /bis item marking, its tooltip host, underlined on hover; the focus ring
   *  is the site's 2px gold `:focus-visible` outline. */
  const NAME =
    'relative inline-flex min-h-6 max-w-full items-center gap-1.5 align-middle hover:underline underline-offset-2';
  /** The icon's box: 20px, a 1px quality border, filled so a failed image shows no glyph. */
  const ICON_BOX = 'rounded-control inline-block h-5 w-5 shrink-0 overflow-hidden border bg-[#12151c]';
  /** A held name (fix round 3 item 2): the loaded name's own shape -- the same inline-flex box,
   *  a 20px icon place and the real name's text -- so it is exactly as wide and wraps exactly as
   *  the name will; the text is unpainted on a 14px bar. */
  const SKELETON_NAME = 'relative inline-flex min-h-6 max-w-full items-center gap-1.5 align-middle';
  const SKELETON_ICON = 'rounded-control inline-block h-5 w-5 shrink-0 bg-[#1a1d24]';
  const SKELETON = 'rounded-[2px] bg-[#1a1d24] leading-[14px] text-transparent select-none';
  /** An inline link in a sentence: its 44px phone target is vertical padding, which never opens
   *  the line (ForeverScoreCard's SENTENCE_LINK). */
  const SENTENCE_LINK = 'underline underline-offset-2 max-md:py-[13.5px]';
  /** The footer link: 12px of padding above and below, a 44px target. */
  const LINK = 'inline-flex items-center self-start py-3 text-[14px] underline underline-offset-2';
  const HEADING = 'flex flex-wrap items-baseline gap-x-2 focus:outline-none focus-visible:outline-2';
</script>

{#snippet item(view: GearItemView, testid: string)}
  {#if !view.known}
    <span class="text-muted" data-testid={testid}>{view.name}</span>
  {:else if loading}
    <span class={SKELETON_NAME} aria-hidden="true"
      ><span class={SKELETON_ICON}></span><span class={SKELETON} data-testid="character-gear-skeleton"
        >{view.name}</span
      ></span
    >
  {:else}
    {@const icon = brokenIcons.has(view.id) ? null : iconOf(table ?? null, view.id)}
    <span
      class={NAME}
      data-bis-item={view.id}
      tabindex="0"
      role="button"
      aria-haspopup="true"
      data-testid={`item-hover-${view.id}`}
      ><span
        class={ICON_BOX}
        style={`border-color: ${rarityBorderColorFor(table?.get(view.id)?.quality ?? 1)}`}
        aria-hidden="true"
        data-testid="character-gear-icon"
        >{#if icon !== null}<img
            src={dataUrl(build, `icons/${icon}.webp`)}
            alt=""
            width="20"
            height="20"
            loading="lazy"
            decoding="async"
            class="block h-full w-full object-cover"
            onerror={() => brokenIcons.add(view.id)}
          />{/if}</span
      ><span class={nameClass(table ?? null, view.id)} data-testid={testid}>{view.name}</span></span
    >
  {/if}
{/snippet}

{#snippet items(row: GearRow)}
  {#if row.worn === null}<span class="text-muted" data-testid="character-gear-empty">{gearCopy.empty}</span
    >{:else}{@render item(row.worn, 'character-gear-worn')}{/if}{#if row.pick !== null}<span
      class="text-muted"
      aria-hidden="true">{gearCopy.arrow}</span
    ><span class="sr-only">{gearCopy.arrowLabel}</span>{@render item(row.pick, 'character-gear-pick')}{/if}
{/snippet}

{#snippet title()}
  <h2 class={HEADING} tabindex="-1" bind:this={heading} data-testid="character-gear-title">
    {#if summary === null}
      <span class="section-title text-[18px]" data-testid="character-gear-heading">{gearCopy.heading}</span>
    {:else}
      <span
        class="section-title text-[18px] after:content-['\00a0·'] max-[480px]:after:content-none"
        data-testid="character-gear-heading">{gearCopy.heading}</span
      >
      <span
        class="text-strong font-body text-[15px] font-semibold max-[480px]:block max-[480px]:basis-full"
        data-testid="character-gear-summary">{summary}</span
      >
    {/if}
  </h2>
{/snippet}

{#if gear.length === 0}
  <section class="-order-2 flex flex-col gap-2" data-testid="character-gear-empty-section">
    {@render title()}
    <p class="text-muted scroll-mt-[132px] text-[14px]" id={GEAR_ANCHOR} data-testid="character-gear-none">
      {gearCopy.noGearParts.lead}<a class={SENTENCE_LINK} href={gearCopy.noGearParts.uploadHref}
        >{gearCopy.noGearParts.upload}</a
      >{gearCopy.noGearParts.or}<a class={SENTENCE_LINK} href={gearCopy.noGearParts.syncHref}
        >{gearCopy.noGearParts.sync}</a
      >{gearCopy.noGearParts.end}
    </p>
  </section>
{:else}
  <section class="-order-2 flex scroll-mt-24 flex-col" id={GEAR_ANCHOR} data-testid="character-gear">
    {@render title()}
    {#if totalsLine !== null}
      <p class="text-muted mt-1 text-[13px] leading-5" data-testid="character-gear-totals">{totalsLine}</p>
    {/if}
    <ul
      class="mt-4 flex flex-col"
      aria-label={gearCopy.listLabel}
      aria-busy={loading ? 'true' : undefined}
      data-testid="character-gear-list"
    >
      {#each rows as row (row.slot)}
        <li
          class={ROW}
          aria-labelledby={`gear-slot-${row.slot}`}
          data-testid={`character-gear-row-${row.slot}`}
        >
          <span class={SLOT} id={`gear-slot-${row.slot}`} data-testid="character-gear-slot">{row.label}</span>
          <span class={LINE} data-testid="character-gear-line">
            {#if row.inline}
              <!-- The note shares this line below 1024px: the dot trails the items. -->
              <span class="{ITEMS} max-lg:after:content-['\00a0·']" data-testid="character-gear-items"
                >{@render items(row)}</span
              >
              <span class={META} data-testid="character-gear-meta">{row.note}</span>
            {:else}
              <span class={ITEMS} data-testid="character-gear-items">{@render items(row)}</span>
            {/if}
            {#if !row.inline && row.source !== null}
              <span class={`${META} basis-full`} data-testid="character-gear-meta"
                >{row.source}{#if row.chance !== null}<span class="whitespace-nowrap"
                    >{gearCopy.chanceLead}{row.chance}</span
                  >{/if}{#if row.note !== null}<span class={loading ? 'invisible' : undefined}
                    >{gearCopy.noteLead}{row.note}</span
                  >{/if}</span
              >
            {/if}
          </span>
          <span
            class={`${GAP} ${
              row.gap === null || row.gap === gearCopy.gapUnscored || row.gap === gearCopy.gap(0)
                ? 'text-muted font-normal'
                : `text-kill ${row.strong ? 'font-bold' : 'font-normal'}`
            }`}
            aria-label={row.gapLabel ?? undefined}
            data-testid="character-gear-gap">{row.gap ?? ''}</span
          >
        </li>
      {/each}
    </ul>
    {#if bisHref !== null && spec !== null}
      <!-- Held (visibility, not display) while the items file loads, so nothing moves after. -->
      <footer data-testid="character-gear-footer">
        <a class={loading ? `${LINK} invisible` : LINK} href={bisHref} data-testid="character-gear-bis"
          >{gearCopy.bisLink(spec, band)}{gearCopy.bisGlue}<span aria-hidden="true">{gearCopy.bisArrow}</span
          ></a
        >
      </footer>
    {/if}
  </section>
  {#if classSlug !== null}
    <BisTooltipHost
      {build}
      modelsElementId="gear-tooltip-models"
      resolve={(id) => fetchItemTooltipModel(build, classSlug ?? '', id)}
    />
  {/if}
{/if}
