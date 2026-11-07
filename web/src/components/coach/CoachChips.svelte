<!-- web/src/components/coach/CoachChips.svelte -->
<!-- The evidence chips under a finding (spec 3.1): one quiet link each (its spell icon and
     place, no box), 36 px tall on desktop and 44 px on touch, built by lib/coach/chips.ts. A chip opens its place in the report in the
     same tab. The row keeps one line of height while the review is held at its minimums. -->
<script lang="ts">
  import { slotClass } from '../../lib/coach/reserves';
  import type { Chip } from '../../lib/coach/chips';
  import type { CoachNames } from '../../lib/coach/names';
  import { followReportLink } from '../../lib/coach/navigate';
  import SpellIcon from './SpellIcon.svelte';

  let {
    chips,
    names,
    reserved = false,
  }: { chips: readonly Chip[]; names: CoachNames; reserved?: boolean } = $props();

  /** An anchor: the spell's icon and its place (`Execute at 2:41`) as a quiet link, no box: the
   *  text colour at rest, gold on hover (the flat system's link, spec 2026-11-05 3.5). */
  const CHIP =
    'text-text hover:text-gold-hover active:text-gold-hover rounded-field focus-visible:outline-gold inline-flex min-h-11 items-center gap-2 text-[13px] font-semibold focus-visible:outline-2 focus-visible:outline-offset-2 md:min-h-9';
</script>

<ul class="mt-2 flex flex-wrap gap-x-5 gap-y-0 {slotClass('chips', reserved)}" data-testid="coach-chips">
  {#each chips as chip (chip.id)}
    <li>
      <a class={CHIP} href={chip.href} onclick={(event) => followReportLink(event, chip.href)}>
        <SpellIcon url={names.iconUrl(chip.spellId)} size={20} />{chip.label}
      </a>
    </li>
  {/each}
</ul>
