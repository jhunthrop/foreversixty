<!-- web/src/components/character/CharacterGearPanel.svelte -->
<!-- The profile's Gear section (Task L8.6 and its fix round 1, spec §4.5/§4.6), under the
     Forever Score card: the 17 slots of the character's newest addon export from GET
     /v1/characters/...'s `gear` (api/openapi.yaml `CharacterGear`). The section is the anchor the
     card's "Most to gain: Gear" opens (#gear, landing 96px under the top, its heading focused).

     The title line reads "Gear · +70 pts across 7 slots" (" · 1 not scored" when a pick has no
     gap); the dot after "Gear" is drawn as its ::after, glued by a no-break space, and at 480px
     and under the summary is its own line with no dot, so no line ever starts with "·" (fix
     round 2 ruling 4). Every separator the section draws is a pseudo-element (the <style> at
     the end), never a text node, with alt text "" so no screen reader reads it (fix round 5
     item 2). 4px under the title one muted 13px line sets the two totals side by side ("The
     card’s 88 counts the band’s best set. These 70 count the picks you can wear."), 16px above
     row 1. The footer holds the BiS page link, 12px of padding above and below it, and, when the
     items file failed, one muted line under it (fix round 6 item 2).

     Each known name is marked the way /bis marks an item (`data-bis-item`, a focusable button,
     the quality colour and a 20px quality-bordered icon), and the shared BisTooltipHost opens
     the item tooltip (a mouse's hover, keyboard focus, a tap; fix round 2 ruling 2), its model
     fetched at runtime. Each icon sits in a 20px box filled #14171f: with an image, a 1px
     border in its quality colour; with none (no icon, a failed image, a failed items file) the
     box stays, empty, in the muted line border (ruling 1, fix round 4 item A). Until the
     build's items file (quality, icon, two-handed) answers, the slot labels, gaps and sources
     draw at once and each name is held behind a 14px bar, so no name is ever painted white
     first and nothing moves when they come (ruling 8): the bar holds the real name, unpainted,
     in the name's own box, so every row wraps the same before and after (fix round 3 item 2);
     a failed file draws every name in the body colour, none held (fix round 6 item 1), and says
     so in the footer, under the link, where nothing moves (fix round 6 item 2).

     Layout. Below 1024px three lines a row: the slot label and the gap on one top, then worn →
     pick, then the source and chance; a row with no source ("Best for your band", no pick) has
     its note as that third line, under the items with no dot, at every width from 320px to
     1024px, so no dot can lead or trail a line (fix round 5 item 1). From 1024px one line: slot
     104 | worn → pick 5fr | source or note 4fr | gap 88, a faint tint on a fine pointer's hover
     and on keyboard focus; every row is 44px at 1440, 1280 and 1024 (ruling 3). The list spans
     the page column: site-column.test.ts bans a container width of 800px and up. Each row is
     named by its slot label (aria-labelledby), the list "Gear by slot" (ruling 9).

     `gear: []` (no export) draws the heading over one line that keeps the anchor; the caller
     draws nothing at all for an API that sends no `gear`. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { gearCopy, notReadCopy } from '../../lib/gear/copy';
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
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import ExportNotRead from './ExportNotRead.svelte';

  let {
    gear,
    total,
    cardGap = null,
    classSlug = null,
    bisHref = null,
    spec = null,
    band = null,
    specLine = null,
    preloaded = undefined,
    unread = false,
    capturedAt = null,
    owned = true,
    first = false,
    planner = false,
  }: {
    /** No gear and a score: the one primary prompt is the planner (lib/character-prompt.ts). */
    planner?: boolean;
    /** The panel goes above the score card (a character with no score yet: the gear picks are the
     *  best thing on the page, W.9). */
    first?: boolean;
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
    /** `gearSpecLine`: the spec the picks are for and where it came from, muted after the heading. */
    specLine?: string | null;
    /** The items file already read (a table) or failed (null); left out, the section reads it. */
    preloaded?: GearItemTable | null;
    /** The export carried no gear and no talents (W.3): the slot list gives way to the "not read
     *  yet" panel, and the band's best in slot stays below it, picks only. */
    unread?: boolean;
    /** The sync's ISO time, when the page knows it, for the panel's "(19 Oct, 21:40)". */
    capturedAt?: string | null;
    /** The viewer's own character; the panel's steps are the owner's. */
    owned?: boolean;
  } = $props();

  const build = activeBuild.build;
  /** The section's place in the profile's flex column: with the score card (-2), after it by DOM
   *  order, or ahead of it (-3, after the header which is -3 and earlier in the DOM). */
  const ORDER = $derived(first ? '-order-3' : '-order-2');

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

  /** Until the items file answers (or when it failed) the main hand's two-handedness is unknown:
   *  the empty off hand at 0 holds its swap note's line (`noteUnknown`, fix round 4 item 4). */
  const rows = $derived(
    gearRows(gear, table === undefined || table === null ? null : (id) => table.get(id)?.twoHand === true),
  );
  /** After "Gear": "+70 pts across 7 slots · 1 not scored", or nothing. */
  const summary = $derived(gearTitleSummary(gear, total));
  /** Icons whose image failed: their box stays, empty (fix round 2 ruling 1). */
  const brokenIcons = new SvelteSet<number>();

  const totalsLine = $derived(gearTotalsLine(cardGap, total));
  /** The unread panel's list: the band's picks only, a row for each slot that has one. */
  const pickRows = $derived(rows.filter((row) => row.pick !== null));

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

  /** Below 1024px: [slot | gap] / [items] / [source, or the note on its own line]. From 1024px:
   *  [slot | items | source or note | gap]. */
  const ROW =
    'border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-start gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] transition-colors duration-[120ms] [@media(hover:hover)_and_(pointer:fine)]:hover:bg-[rgba(255,255,255,0.04)] has-[:focus-visible]:bg-[rgba(255,255,255,0.04)] lg:grid-cols-[104px_minmax(0,5fr)_minmax(0,4fr)_88px] lg:items-center lg:py-0';
  const ROW_PICKS =
    'border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)] items-start gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] transition-colors duration-[120ms] [@media(hover:hover)_and_(pointer:fine)]:hover:bg-[rgba(255,255,255,0.04)] has-[:focus-visible]:bg-[rgba(255,255,255,0.04)] lg:grid-cols-[104px_minmax(0,5fr)_minmax(0,4fr)] lg:items-center lg:py-0';
  const SLOT = 'text-muted col-start-1 row-start-1 text-[13px] leading-5';
  const LINE = 'col-span-2 row-start-2 flex min-w-0 flex-wrap items-baseline gap-x-1 gap-y-1 lg:contents';
  const ITEMS = 'min-w-0 break-words lg:col-start-2 lg:row-start-1';
  const META = 'text-muted min-w-0 text-[13px] lg:col-start-3 lg:row-start-1';
  const GAP =
    'col-start-2 row-start-1 text-right numeric text-[13px] leading-5 whitespace-nowrap lg:col-start-4 lg:row-start-1';
  /** A known name: the /bis item marking, its tooltip host, underlined on hover where there is a
   *  fine pointer that hovers (fix round 4 item C: a tap never leaves it underlined); the focus
   *  ring is the site's 2px gold `:focus-visible` outline. */
  const NAME =
    'relative inline-flex min-h-6 max-w-full items-center gap-1.5 align-middle [@media(hover:hover)_and_(pointer:fine)]:hover:underline underline-offset-2';
  /** The icon's box: 20px, filled #14171f so a failed image shows no glyph. With an image, a 1px
   *  quality border; without one (no icon, a failed image, a failed items file), the skeleton's
   *  treatment, the muted line border (fix round 4 item A). */
  const ICON_BOX = 'rounded-control inline-block h-5 w-5 shrink-0 overflow-hidden border bg-[#14171f]';
  /** A held name (fix round 3 item 2): the loaded name's own shape -- the same inline-flex box,
   *  a 20px icon place and the real name's text -- so it is exactly as wide and wraps exactly as
   *  the name will; the text is unpainted on a 14px bar. */
  const SKELETON_NAME = 'relative inline-flex min-h-6 max-w-full items-center gap-1.5 align-middle';
  const SKELETON_ICON = 'rounded-control inline-block h-5 w-5 shrink-0 bg-[#1a1d24]';
  /** The bar is the wrapper's background; the name inside is visibility-hidden, its width kept
   *  and never found by a page search (fix round 4 item 7). */
  const SKELETON = 'rounded-[2px] bg-[#1a1d24] leading-[14px] select-none';
  /** An inline link in a sentence: its 44px phone target is vertical padding, which never opens
   *  the line (ForeverScoreCard's SENTENCE_LINK). */
  const SENTENCE_LINK = 'underline underline-offset-2 max-md:py-[13.5px]';
  /** The footer link: 12px of padding above and below, a 44px target. */
  const LINK = 'inline-flex items-center self-start py-3 text-[14px] underline underline-offset-2';
  /** The failed file's line: muted 13px under the link; with no link, 12px under the list. */
  const FAILED = 'text-muted text-[13px] leading-5 first:pt-3';
  const HEADING = 'flex flex-wrap items-baseline gap-x-2 focus:outline-none focus-visible:outline-2';
</script>

{#snippet item(view: GearItemView, testid: string)}
  {#if !view.known}
    <span class="text-muted" data-testid={testid}>{view.name}</span>
  {:else if loading}
    <span class={SKELETON_NAME} aria-hidden="true"
      ><span class={SKELETON_ICON}></span><span class={SKELETON} data-testid="character-gear-skeleton"
        ><span class="invisible">{view.name}</span></span
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
        class={icon === null ? `${ICON_BOX} border-line` : ICON_BOX}
        style={icon === null
          ? undefined
          : `border-color: ${rarityBorderColorFor(table?.get(view.id)?.quality ?? 1)}`}
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

<!-- "main-hand swap" is one unbreakable run, so neither "main-" nor "swap" is left alone at a
     line's end or start (fix round 6 item 3); the words are the copy's. -->
{#snippet noteText(note: string)}{#if note === gearCopy.swap}{gearCopy.swapLead}<span
      class="whitespace-nowrap">{gearCopy.swapKept}</span
    >{:else}{note}{/if}{/snippet}

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
      <span class="section-title" data-testid="character-gear-heading">{gearCopy.heading}</span>
    {:else}
      <span class="section-title gear-dot-after" data-testid="character-gear-heading">{gearCopy.heading}</span
      >
      <span
        class="text-strong font-body text-[15px] font-semibold max-[480px]:block max-[480px]:basis-full"
        data-testid="character-gear-summary">{summary}</span
      >
    {/if}
    {#if specLine !== null}
      <span class="text-muted font-body text-[13px] font-normal" data-testid="character-gear-spec"
        >{specLine}</span
      >
    {/if}
  </h2>
{/snippet}

{#if unread && gear.length > 0}
  <section
    class="-order-2 flex scroll-mt-24 flex-col gap-4"
    id={GEAR_ANCHOR}
    data-testid="character-gear-unread"
  >
    <ExportNotRead {capturedAt} {owned} />
    <div class="flex flex-col">
      <h2 class={HEADING} tabindex="-1" bind:this={heading} data-testid="character-gear-title">
        <span class="section-title" data-testid="character-gear-heading">{notReadCopy.bestInSlot(band)}</span>
      </h2>
      <ul
        class="mt-4 flex flex-col"
        aria-label={gearCopy.listLabel}
        aria-busy={loading ? 'true' : undefined}
        data-testid="character-gear-list"
      >
        {#each pickRows as row (row.slot)}
          <li
            class={ROW_PICKS}
            aria-labelledby={`gear-slot-${row.slot}`}
            data-testid={`character-gear-row-${row.slot}`}
          >
            <span class={SLOT} id={`gear-slot-${row.slot}`} data-testid="character-gear-slot"
              >{row.label}</span
            >
            <span class={LINE} data-testid="character-gear-line">
              <span class={ITEMS} data-testid="character-gear-items"
                >{#if row.pick !== null}{@render item(row.pick, 'character-gear-pick')}{/if}</span
              >
              <span class={`${META} basis-full`} data-testid="character-gear-meta"
                >{row.source}{#if row.chance !== null}<span aria-hidden="true"
                    >{gearCopy.sourceSeparator}</span
                  ><span class="whitespace-nowrap" data-testid="character-gear-chance">{row.chance}</span
                  >{/if}</span
              >
            </span>
          </li>
        {/each}
      </ul>
      {#if bisHref !== null && spec !== null}
        <footer data-testid="character-gear-footer">
          <a class={loading ? `${LINK} invisible` : LINK} href={bisHref} data-testid="character-gear-bis"
            >{gearCopy.bisLink(spec, band)}{gearCopy.bisGlue}<span aria-hidden="true"
              >{gearCopy.bisArrow}</span
            ></a
          >
        </footer>
      {/if}
    </div>
  </section>
  {#if classSlug !== null}
    <BisTooltipHost
      {build}
      modelsElementId="gear-tooltip-models"
      resolve={(id) => fetchItemTooltipModel(build, classSlug ?? '', id)}
    />
  {/if}
{:else if gear.length === 0}
  <section class="-order-2 flex flex-col gap-2" data-testid="character-gear-empty-section">
    {@render title()}
    <p class="text-muted scroll-mt-[132px] text-[14px]" id={GEAR_ANCHOR} data-testid="character-gear-none">
      {gearCopy.noGearParts.lead}<a class={SENTENCE_LINK} href={gearCopy.noGearParts.syncHref}
        >{gearCopy.noGearParts.sync}</a
      >{#if owned}{gearCopy.noGearParts.end}{:else}{gearCopy.noGearParts.or}<a
          class={SENTENCE_LINK}
          href={gearCopy.noGearParts.signInHref}>{gearCopy.noGearParts.signIn}</a
        >{gearCopy.noGearParts.claimEnd}{/if}
    </p>
    {#if planner}
      <a
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4`}
        href={gearCopy.noGearPlannerHref}
        data-testid="character-gear-planner">{gearCopy.noGearPlanner}</a
      >
    {/if}
  </section>
{:else}
  <section class="{ORDER} flex scroll-mt-24 flex-col" id={GEAR_ANCHOR} data-testid="character-gear">
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
            <span class={ITEMS} data-testid="character-gear-items">{@render items(row)}</span>
            {#if row.inline}
              <!-- A row with no source: its note is the third line, under the items with no dot,
                   at every width below 1024px; from 1024px the source column (fix round 5 item 1). -->
              <span class={`${META} basis-full`} data-testid="character-gear-meta">{row.note}</span>
            {:else if row.source !== null}
              <span class={`${META} basis-full`} data-testid="character-gear-meta"
                >{row.source}{#if row.chance !== null}<span aria-hidden="true"
                    >{gearCopy.sourceSeparator}</span
                  ><span class="whitespace-nowrap" data-testid="character-gear-chance">{row.chance}</span
                  >{/if}{#if row.note !== null && row.noteUnknown && loading}<span
                    class="rounded-[2px] bg-[#1a1d24]"
                    data-testid="character-gear-note-bar"
                    ><span class="gear-dot-before invisible" data-testid="character-gear-note"
                      >{@render noteText(row.note)}</span
                    ></span
                  >{:else if row.note !== null}<span
                    class="gear-dot-before{row.noteUnknown ? ' invisible' : ''}"
                    data-testid="character-gear-note">{@render noteText(row.note)}</span
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
    {#if (bisHref !== null && spec !== null) || table === null}
      <!-- The link is held (visibility, not display) while the items file loads, so nothing moves
           after. A failed file is said here, under the link, so no row and not the link moves
           (fix round 6 item 2). -->
      <footer data-testid="character-gear-footer">
        {#if bisHref !== null && spec !== null}
          <a class={loading ? `${LINK} invisible` : LINK} href={bisHref} data-testid="character-gear-bis"
            >{gearCopy.bisLink(spec, band)}{gearCopy.bisGlue}<span aria-hidden="true"
              >{gearCopy.bisArrow}</span
            ></a
          >
        {/if}
        {#if table === null}
          <p class={FAILED} data-testid="character-gear-failed">{gearCopy.failed}</p>
        {/if}
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

<style>
  /* The drawn separators (fix round 2 ruling 4, fix round 4 item B): each dot is a pseudo-element
     glued by no-break spaces to the texts either side of it, never a text node. Its alt text is ""
     so that no screen reader reads it (fix round 5 item 2); a browser without content's alt text
     drops that declaration and draws the plain one before it, the same glyph. Global, so that
     the markup carries no scoping class: the names are this section's own. */
  :global(.gear-dot-before)::before {
    content: '\00a0·\00a0';
    content: '\00a0·\00a0' / '';
  }

  :global(.gear-dot-after)::after {
    content: '\00a0·';
    content: '\00a0·' / '';
  }

  /* At 480px and under the title's summary is a line of its own: no dot after "Gear". */
  @media (max-width: 480px) {
    :global(.gear-dot-after)::after {
      content: none;
    }
  }
</style>
