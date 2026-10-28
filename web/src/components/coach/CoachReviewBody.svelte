<!-- web/src/components/coach/CoachReviewBody.svelte -->
<!-- The review's fixed structure (spec 3.1 and 5.1): opening, Fix first with its drill, Then,
     Doing well. It renders only these fields; nothing is free prose outside them. While a
     writing review is held at its minimums (`reserve` set), every text slot keeps its reserved
     lines and the empty Then and Doing well slots stay, so the model's text swaps in with no
     shift. Evidence anchors become chips through coachAnchorHref. -->
<script lang="ts">
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
          <svg class="text-gold mt-1 h-3 w-3 shrink-0" viewBox="0 0 12 12" aria-hidden="true">
            <path d="M2 6h8M7 3l3 3-3 3" fill="none" stroke="currentColor" stroke-width="1.5" />
          </svg>{line}
        </li>
      {/each}
    </ul>
  {/if}
{/snippet}

<div data-testid="coach-review-body">
  <p
    class="text-text max-w-[720px] text-[15px] leading-[22px] {slotClass('opening', held)}"
    data-testid="coach-opening"
  >
    {review.opening}
  </p>

  {#if review.fixFirst !== null}
    {@const fix = review.fixFirst}
    <section class="border-gold mt-5 border-l-2 pl-4" data-testid="coach-fix-first">
      <h3 class="label text-gold">{coachCopy.labelFixFirst}</h3>
      <div class="mt-1 flex items-center gap-2">
        <SpellIcon url={names.iconUrl(firstSpellId(fix.anchors, answer.anchors))} size={32} />
        <p class="text-strong text-[20px] leading-7 font-bold" data-testid="coach-fix-title">
          <CoachText text={fix.title} {names} {build} {classSlug} plainSpells />
        </p>
      </div>
      <p class="text-text mt-1 text-[15px] leading-[22px] {slotClass('why', held)}">
        <CoachText text={fix.why} {names} {build} {classSlug} />
      </p>
      {@render linkLines(fix.links, 0)}
      <div class="border-line-soft mt-3 border-t pt-3">
        <h4 class="label text-muted">{coachCopy.labelPractice}</h4>
        <p
          class="text-text mt-1 text-[14px] leading-[22px] {slotClass('drill', held)}"
          data-testid="coach-drill"
        >
          {fix.drillText}
        </p>
      </div>
      <CoachChips chips={chipsOf(fix.anchors)} {names} reserved={held} />
    </section>
  {/if}

  {#if thenSlots > 0}
    <section class="mt-6" data-testid="coach-then">
      <h3 class="label text-muted">{coachCopy.labelThen}</h3>
      <ul class="mt-1 flex flex-col">
        {#each slots(thenSlots) as index (index)}
          {@const item = review.thenFix[index]}
          <li class="border-line-soft border-t py-3.5 first:border-t-0" data-testid="coach-then-item">
            {#if item !== undefined}
              <div class="flex items-center gap-2">
                <SpellIcon url={names.iconUrl(firstSpellId(item.anchors, answer.anchors))} size={24} />
                <p class="text-strong text-[16px] leading-6 font-semibold">
                  <CoachText text={item.title} {names} {build} {classSlug} plainSpells />
                </p>
              </div>
              <p class="text-text mt-1 text-[14px] leading-[22px] {slotClass('thenWhy', held)}">
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
    <section class="mt-5" data-testid="coach-well">
      <h3 class="label text-kill">{coachCopy.labelWell}</h3>
      <ul class="mt-1 flex flex-col gap-1">
        {#each slots(wellSlots) as index (index)}
          {@const item = review.doingWell[index]}
          <li class="flex gap-2 text-[14px] leading-[22px] {slotClass('well', held)}">
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
