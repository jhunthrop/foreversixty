<!-- web/src/components/sim/WorthKnowing.svelte -->
<!-- The notes that qualify a saved run's number (spec 2026-11-04-sim-results 3.5): each one is
     true, so none is dismissible; with none, the caller draws nothing. A row holds the spell's
     icon or the alert tile, one sentence and an optional link. -->
<script lang="ts">
  import { ACTION_LINK } from '../../lib/ui/link';
  import { warningsCopy } from '../../lib/sim/result-copy';
  import type { SpellInfo } from '../../lib/sim/spell-info';
  import type { SimWarning } from '../../lib/sim/warnings';
  import { CARD_QUIET } from '../../lib/ui/card';
  import CardHead from '../ui/CardHead.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import IconTile from '../ui/IconTile.svelte';

  let {
    notes,
    spells,
    plannerHref,
    onrerun,
  }: {
    notes: readonly SimWarning[];
    spells: SpellInfo;
    plannerHref: string;
    onrerun: () => void;
  } = $props();
</script>

<section
  class={`${CARD_QUIET} border-l-ember mx-[18px] overflow-hidden border-l-[3px] md:mx-0`}
  data-testid="sim-worth-knowing"
>
  <CardHead glyph="alert" label={warningsCopy.label} title={warningsCopy.title} meta={warningsCopy.notes(notes.length)} />
  <ul class="border-line-soft border-t">
    {#each notes as note (`${note.id}:${note.actionKey ?? ''}`)}
      {@const icon = note.actionKey === undefined ? null : spells.iconUrl(note.actionKey)}
      <li
        class="border-line-soft flex min-h-14 items-center gap-3.5 border-t px-5 py-3 first:border-t-0"
        data-testid={`sim-note-${note.id}`}
      >
        {#if icon !== null}
          <img
            src={icon}
            alt=""
            width="28"
            height="28"
            class="h-7 w-7 shrink-0 rounded-[6px] border border-black shadow-[0_0_0_1px_rgba(229,185,85,.18)]"
          />
        {:else if note.actionKey !== undefined}
          <span
            class="h-7 w-7 shrink-0 rounded-[6px] border border-black bg-white/[.05] shadow-[0_0_0_1px_rgba(229,185,85,.18)]"
          ></span>
        {:else}
          <IconTile glyph="alert" size={28} tone="muted" />
        {/if}
        <p class="text-text min-w-0 flex-1 text-[15px] leading-[22px]">{note.text}</p>
        {#if note.action === 'planner'}
          <a
            class={`${ACTION_LINK} shrink-0 whitespace-nowrap`}
            href={plannerHref}>{warningsCopy.openInPlanner}<Glyph name="chevron" /></a
          >
        {:else if note.action === 'sim'}
          <button
            type="button"
            class={`${ACTION_LINK} shrink-0 whitespace-nowrap`}
            onclick={onrerun}>{warningsCopy.openInSim}<Glyph name="chevron" /></button
          >
        {/if}
      </li>
    {/each}
  </ul>
</section>
