<!-- web/src/components/coach/CoachChips.svelte -->
<!-- The evidence chips under a finding (spec 3.1): one link each, 36 px tall on desktop and
     44 px on touch, built by lib/coach/chips.ts. A chip opens its place in the report in the
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

  const CHIP =
    'border-line text-text hover:border-gold hover:text-strong focus-visible:outline-gold active:bg-card-top inline-flex min-h-11 items-center gap-1.5 rounded-[3px] border px-2.5 text-[12px] focus-visible:outline-2 md:min-h-9';
</script>

<ul class="mt-3 flex flex-wrap gap-2 {slotClass('chips', reserved)}" data-testid="coach-chips">
  {#each chips as chip (chip.id)}
    <li>
      <a class={CHIP} href={chip.href} onclick={(event) => followReportLink(event, chip.href)}>
        <SpellIcon url={names.iconUrl(chip.spellId)} size={16} />{chip.label}
      </a>
    </li>
  {/each}
</ul>
