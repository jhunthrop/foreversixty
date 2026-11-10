<!-- web/src/components/coach/CoachReviewBody.svelte -->
<!-- The review's fixed structure (spec 3.1 and 5.1; logs-reports 6.2 Coach): the opening sentence
     (17/26, strong), Fix first (gold Label, H3 title with its spell icon, why, The drill), Then fix,
     Doing well (green checks), held apart by type and space, no rule or box. It renders only these fields; nothing is free prose outside them. While a
     writing review is held at its minimums (`reserve` set), every text slot keeps its reserved
     lines and the empty Then and Doing well slots stay, so the model's text swaps in with no
     shift. Evidence anchors become chips through coachAnchorHref. -->
<script lang="ts">
  import Icon from '../ui/Icon.svelte';
  import { ICON_ARROW_RIGHT } from '../../lib/ui/ui-icons';
  import { chipsFor, firstSpellId } from '../../lib/coach/chips';
  import { coachCopy } from '../../lib/coach/copy';
  import { withAnchorNames, type CoachNames } from '../../lib/coach/names';
  import { slotClass, type Reserve } from '../../lib/coach/reserves';
  import type { CoachFull, CoachLink } from '../../lib/coach/types';
  import CoachChips from './CoachChips.svelte';
  import CoachText from './CoachText.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    answer,
    names: tables,
    build,
    classSlug,
    firstFight,
    playerGuid,
    reserve = null,
  }: {
    answer: CoachFull;
    names: CoachNames;
    build: string;
    classSlug: string;
    firstFight: number;
    playerGuid: string;
    reserve?: Reserve | null;
  } = $props();

  const names = $derived(withAnchorNames(tables, answer.anchors));
  const review = $derived(answer.review);
  const held = $derived(reserve !== null);
  const thenSlots = $derived(Math.max(review.thenFix.length, reserve?.thenCount ?? 0));
  const wellSlots = $derived(Math.max(review.doingWell.length, reserve?.wellCount ?? 0));
  const slots = (count: number): number[] => Array.from({ length: count }, (_, i) => i);

  const chipsOf = (ids: readonly string[]) =>
    chipsFor(ids, {
      anchors: answer.anchors,
      fightIndex: answer.fight_index,
      firstFight,
      playerGuid,
      spellName: names.spellName,
    });
  // Link lines belong to the rule-only rendering (spec 3.1 item 5, 10.4): the model's prose
  // carries the connection itself, so its review prints none.
  const sentences = (links: readonly CoachLink[]): string[] =>
    answer.source !== 'rules' ? [] : links.flatMap((link) => (link.sentence === null ? [] : [link.sentence]));
  const linkSlot = (index: number): boolean => reserve?.links[index] === true;
</script>

{#snippet linkLines(links: readonly CoachLink[], index: number)}
  {@const lines = sentences(links)}
  {#if lines.length > 0 || (held && linkSlot(index))}
    <ul class="mt-2 flex flex-col gap-1 {slotClass('links', held)}" data-testid="coach-link-lines">
      {#each lines as line (line)}
        <li class="text-muted flex gap-2 text-[13px] leading-5">
          <Icon icon={ICON_ARROW_RIGHT} size={12} class="text-gold mt-1" /><span
            ><CoachText text={line} {names} {build} {classSlug} /></span
          >
        </li>
      {/each}
    </ul>
  {/if}
{/snippet}

<div data-testid="coach-review-body">
  <p
    class="text-strong max-w-[720px] text-[17px] leading-[26px] {slotClass('opening', held)}"
    data-testid="coach-opening"
  >
    <CoachText text={review.opening} {names} {build} {classSlug} />
  </p>

  {#if review.fixFirst !== null}
    {@const fix = review.fixFirst}
    <section class="mt-6" data-testid="coach-fix-first">
      <h3 class="text-label text-gold font-bold">{coachCopy.labelFixFirst}</h3>
      <div class="mt-2 flex items-center gap-3">
        <SpellIcon url={names.iconUrl(firstSpellId(fix.anchors, answer.anchors))} size={36} />
        <p class="text-strong text-[20px] leading-7 font-bold" data-testid="coach-fix-title">
          <CoachText text={fix.title} {names} {build} {classSlug} plainSpells />
        </p>
      </div>
      <p class="text-text mt-2 max-w-[720px] text-[15px] leading-6 {slotClass('why', held)}">
        <CoachText text={fix.why} {names} {build} {classSlug} />
      </p>
      {@render linkLines(fix.links, 0)}
      <div class="mt-4">
        <h4 class="text-label text-muted font-bold">{coachCopy.labelPractice}</h4>
        <p
          class="text-text mt-1 max-w-[720px] text-[15px] leading-6 {slotClass('drill', held)}"
          data-testid="coach-drill"
        >
          <CoachText text={fix.drillText} {names} {build} {classSlug} />
        </p>
      </div>
      <CoachChips chips={chipsOf(fix.anchors)} {names} reserved={held} />
    </section>
  {/if}

  {#if thenSlots > 0}
    <section class="mt-8" data-testid="coach-then">
      <h3 class="text-label text-muted font-bold">{coachCopy.labelThen}</h3>
      <ul class="mt-2 flex flex-col gap-5">
        {#each slots(thenSlots) as index (index)}
          {@const item = review.thenFix[index]}
          <li data-testid="coach-then-item">
            {#if item !== undefined}
              <div class="flex items-center gap-2">
                <SpellIcon url={names.iconUrl(firstSpellId(item.anchors, answer.anchors))} size={24} />
                <p class="text-strong text-[16px] leading-6 font-semibold">
                  <CoachText text={item.title} {names} {build} {classSlug} plainSpells />
                </p>
              </div>
              <p class="text-text mt-1 max-w-[720px] text-[15px] leading-6 {slotClass('thenWhy', held)}">
                <CoachText text={item.why} {names} {build} {classSlug} />
              </p>
              {@render linkLines(item.links, index + 1)}
              <CoachChips chips={chipsOf(item.anchors)} {names} reserved={held} />
            {:else}
              <!-- The replaced item's height: title row, why, links when it had any, chips. -->
              <div aria-hidden="true" data-testid="coach-then-empty">
                <div class="h-6"></div>
                <div class="mt-1 {slotClass('thenWhy', held)}"></div>
                {#if held && linkSlot(index + 1)}<div class="mt-2 {slotClass('links', held)}"></div>{/if}
                <div class="mt-3 {slotClass('chips', held)}"></div>
              </div>
            {/if}
          </li>
        {/each}
      </ul>
    </section>
  {/if}

  {#if wellSlots > 0}
    <section class="mt-8" data-testid="coach-well">
      <h3 class="text-label text-kill font-bold">{coachCopy.labelWell}</h3>
      <ul class="mt-2 flex flex-col gap-1">
        {#each slots(wellSlots) as index (index)}
          {@const item = review.doingWell[index]}
          <li class="flex gap-2 text-[15px] leading-6 {slotClass('well', held)}">
            {#if item !== undefined}
              <span class="text-kill shrink-0" aria-hidden="true">✓</span>
              <span><CoachText text={item.text} {names} {build} {classSlug} /></span>
            {/if}
          </li>
        {/each}
      </ul>
    </section>
  {/if}
</div>
