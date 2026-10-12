<!-- web/src/components/character/CharacterGearPanel.svelte -->
<!-- The profile's Gear section (Task L8.6, spec §4.5/§4.6), under the Forever Score card: the
     17 slots of the character's newest addon export from GET /v1/characters/...'s `gear` (api/
     openapi.yaml `CharacterGear`), each "Head · worn → pick · source · chance · +N pts", the
     names in their quality colour and each opening the item tooltip (tenet 2). The title line
     reads "Gear · +88 pts across 6 slots" from `gear_gap_total` (ruling 1). The section is the
     anchor the card's "Most to gain: Gear" opens (#gear, landing 96px under the top), and its
     footer carries the BiS page link and one muted sentence on how the slot gaps differ from the
     card's "below its top" (rulings 2 and 3).

     Layout (ruling 4): from md one line a row -- slot, the items (at most 60% of the row), the
     source and chance after them, the gap in a right-aligned mono column at the row's end (the
     page column's width, as the fight tables below). Below md two lines --
     the items, then the source and chance with the gap at the line's end -- in 44px rows that
     wrap a long name rather than widen the page.

     `gear: []` (no export) draws no section, only one line in its place that keeps the anchor;
     the caller draws nothing at all for an API that sends no `gear`. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { gearCopy } from '../../lib/gear/copy';
  import { GEAR_ANCHOR, gearRows, gearTitle, type GearItemView } from '../../lib/gear/panel';
  import { itemQualities } from '../../lib/gear/qualities';
  import type { CharacterGearSlot } from '../../lib/gear/types';
  import { rarityClassFor } from '../../lib/planner/items';
  import ItemHover from '../ItemHover.svelte';

  let {
    gear,
    total,
    classSlug = null,
    bisHref = null,
    spec = null,
    band = null,
  }: {
    /** `gear`: all 17 slots, or `[]` when the character has no export. */
    gear: readonly CharacterGearSlot[];
    /** `gear_gap_total`. */
    total: number | null;
    /** The character's class ("priest"), whose item table gives the qualities and tooltips. */
    classSlug?: string | null;
    /** The spec's best-in-slot page at this band, for the footer link; none, no link. */
    bisHref?: string | null;
    /** The spec the link names ("Holy") and the band ("60"). */
    spec?: string | null;
    band?: number | null;
  } = $props();

  /** Item id -> quality, once the class's item table has loaded; until then (and for an item
   *  the table does not have) a name takes the common colour. */
  let qualities = $state<ReadonlyMap<number, number>>(new Map());
  $effect(() => {
    const slug = classSlug;
    if (slug === null || gear.length === 0) return;
    void itemQualities(activeBuild.build, slug).then((loaded) => {
      if (slug === classSlug) qualities = loaded;
    });
  });

  const rows = $derived(gearRows(gear));
  /** "Gear" in the section-title face, then " · +70 pts across 7 slots" in body text. */
  const titleRest = $derived(gearTitle(gear, total).slice(gearCopy.heading.length));

  /** Phone: [slot | items items] then [ | meta | gap]. From md: [slot | items | meta | gap]. */
  const ROW =
    'border-line-soft grid min-h-11 grid-cols-[72px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-0.5 border-b px-2 py-2 text-[14px] md:grid-cols-[112px_fit-content(60%)_minmax(0,1fr)_88px]';
  const SLOT = 'text-muted col-start-1 row-start-1 text-[13px]';
  const ITEMS = 'col-span-2 col-start-2 row-start-1 min-w-0 break-words md:col-span-1';
  const META = 'text-muted col-start-2 row-start-2 min-w-0 text-[13px] md:col-start-3 md:row-start-1';
  const GAP =
    'tabular col-start-3 row-start-2 text-right font-mono text-[13px] whitespace-nowrap md:col-start-4 md:row-start-1';
  /** An inline link in a sentence: its 44px phone target is vertical padding, which never opens
   *  the line (ForeverScoreCard's SENTENCE_LINK). */
  const SENTENCE_LINK = 'underline underline-offset-2 max-md:py-[13.5px]';
  const LINK =
    'inline-flex min-h-11 items-center self-start text-[14px] underline underline-offset-2 md:min-h-0';
</script>

{#snippet item(view: GearItemView, testid: string)}
  {#if !view.known}
    <span class="text-muted" data-testid={testid}>{view.name}</span>
  {:else if classSlug === null}
    <span class={rarityClassFor(qualities.get(view.id) ?? 1)} data-testid={testid}>{view.name}</span>
  {:else}
    <ItemHover itemId={view.id} {classSlug} build={activeBuild.build}
      ><span class={rarityClassFor(qualities.get(view.id) ?? 1)} data-testid={testid}>{view.name}</span
      ></ItemHover
    >
  {/if}
{/snippet}

{#if gear.length === 0}
  <p class="text-muted -order-2 scroll-mt-24 text-[14px]" id={GEAR_ANCHOR} data-testid="character-gear-none">
    {gearCopy.noGearParts.lead}<a class={SENTENCE_LINK} href={gearCopy.noGearParts.uploadHref}
      >{gearCopy.noGearParts.upload}</a
    >{gearCopy.noGearParts.or}<a class={SENTENCE_LINK} href={gearCopy.noGearParts.syncHref}
      >{gearCopy.noGearParts.sync}</a
    >{gearCopy.noGearParts.end}
  </p>
{:else}
  <section class="-order-2 flex scroll-mt-24 flex-col gap-2" id={GEAR_ANCHOR} data-testid="character-gear">
    <h2 class="flex flex-wrap items-baseline gap-x-2" data-testid="character-gear-title">
      <span class="section-title text-[18px]">{gearCopy.heading}</span><span
        class="text-strong font-body text-[15px] font-semibold">{titleRest}</span
      >
    </h2>
    <ul class="flex flex-col" data-testid="character-gear-list">
      {#each rows as row (row.slot)}
        <li class={ROW} data-testid={`character-gear-row-${row.slot}`}>
          <span class={SLOT} data-testid="character-gear-slot">{row.label}</span>
          <span class={ITEMS} data-testid="character-gear-items"
            >{#if row.worn === null}<span class="text-muted" data-testid="character-gear-empty"
                >{gearCopy.empty}</span
              >{:else}{@render item(row.worn, 'character-gear-worn')}{/if}{#if row.pick !== null}<span
                class="text-muted"
                aria-hidden="true">{gearCopy.arrow}</span
              ><span class="sr-only">{gearCopy.arrowLabel}</span>{@render item(
                row.pick,
                'character-gear-pick',
              )}{/if}</span
          >
          {#if row.state === 'best'}
            <span class={META} data-testid="character-gear-meta"
              ><span class="max-md:hidden" aria-hidden="true">{gearCopy.metaLead}</span>{gearCopy.best}</span
            >
          {:else if row.source !== null}
            <span class={META} data-testid="character-gear-meta"
              ><span class="max-md:hidden" aria-hidden="true">{gearCopy.metaLead}</span
              >{row.source}{#if row.chance !== null}{gearCopy.chanceLead}{row.chance}{/if}</span
            >
          {/if}
          <span
            class={`${GAP} ${row.gap === null || row.gap === gearCopy.gapUnscored || row.gap === gearCopy.gap(0) ? 'text-muted' : 'text-kill font-semibold'}`}
            aria-label={row.gapLabel ?? undefined}
            data-testid="character-gear-gap">{row.gap ?? ''}</span
          >
        </li>
      {/each}
    </ul>
    <footer class="flex flex-col gap-1 pt-1" data-testid="character-gear-footer">
      <p class="text-muted text-[13px] leading-[20px]" data-testid="character-gear-difference">
        {gearCopy.difference}
      </p>
      {#if bisHref !== null && spec !== null}
        <a class={LINK} href={bisHref} data-testid="character-gear-bis"
          >{gearCopy.bisLink(spec, band)}{gearCopy.bisGlue}<span aria-hidden="true">{gearCopy.bisArrow}</span
          ></a
        >
      {/if}
    </footer>
  </section>
{/if}
