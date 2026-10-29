<!-- web/src/components/coach/CoachTeaser.svelte -->
<!-- The free viewer's panel body (spec 4.2): the count in gold, the first title with its spell
     icon, and one outlined gold button to /premium. Zero findings says so and offers nothing. -->
<script lang="ts">
  import { coachCopy, fill } from '../../lib/coach/copy';
  import { PREMIUM_LINKS } from '../../lib/billing/premium-links';
  import type { CoachNames } from '../../lib/coach/names';
  import type { CoachTeaser } from '../../lib/coach/types';
  import SpellIcon from './SpellIcon.svelte';

  let { answer, names }: { answer: CoachTeaser; names: CoachNames } = $props();

  const count = $derived(
    answer.finding_count === 1
      ? coachCopy.teaserCountOne
      : fill(coachCopy.teaserCount, { n: answer.finding_count }),
  );
</script>

<div data-testid="coach-teaser">
  {#if answer.finding_count === 0}
    <p class="text-text text-[15px] leading-[22px]">{coachCopy.rulesNothing}</p>
  {:else}
    <h3 class="label text-gold" data-testid="coach-teaser-count">{count}</h3>
    {#if answer.fix_first !== null}
      <div class="border-gold mt-2 flex items-center gap-2 border-l-2 pl-4">
        <SpellIcon url={names.iconUrl(answer.fix_first.spell_id)} size={32} />
        <p class="text-strong text-[20px] leading-7 font-bold" data-testid="coach-teaser-first">
          {fill(coachCopy.teaserFirst, { title: answer.fix_first.title })}
        </p>
      </div>
    {/if}
    <a
      href={PREMIUM_LINKS.coach}
      class="rounded-control border-gold text-gold hover:border-gold-hover hover:text-gold-hover focus-visible:outline-gold mt-5 inline-flex h-11 items-center border px-4 text-[14px] font-semibold hover:bg-[#e5b95514] focus-visible:outline-2 focus-visible:outline-offset-2 active:bg-[#e5b95524]"
      data-testid="coach-upgrade">{coachCopy.upgrade}</a
    >
  {/if}
</div>
