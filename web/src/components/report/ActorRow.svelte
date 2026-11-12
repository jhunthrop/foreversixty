<!-- web/src/components/report/ActorRow.svelte -->
<!-- One row of an actor table (logs spec 2026-11-07 6.2): a StatRow with the crest, the name in
     its class colour (the link that opens the player's own view, decision R5), the sub line, a
     class-colour Meter (share of the top row) and the figures; the coach's line for the viewer's
     own characters follows it (components/coach/CoachRowLine, placed as shipped). -->
<script lang="ts">
  import { unitLabel } from '../../lib/report/names.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { Actor } from '../../lib/report/types';
  import type { StatFigure } from '../../lib/ui/stat-table';
  import Meter from '../ui/Meter.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import CoachRowLine from '../coach/CoachRowLine.svelte';
  import Crest from '../ui/Crest.svelte';

  let {
    actor,
    subLine,
    figures,
    ofTop,
    phone,
    href,
    selected = false,
    onSelect,
  }: {
    actor: Actor;
    /** "{spec} {class} · ilvl {n}", plus overheal or "dead since" where they apply. */
    subLine: string;
    figures: readonly StatFigure[];
    /** This row's amount over the top row's, 0 to 1. */
    ofTop: number;
    /** The primary figure's phone line ("3.2 DPS"). */
    phone: string;
    /** The player's own view (Source = this player); undefined for a row that is not a player. */
    href: string | undefined;
    /** The bar's Source is this player. */
    selected?: boolean;
    onSelect: (guid: string) => void;
  } = $props();

  const label = $derived(unitLabel(actor.name));
  const colour = $derived(classColorVar(actor.class));

  function follow(event: MouseEvent): void {
    if (event.metaKey || event.ctrlKey || event.shiftKey || event.button !== 0) return;
    event.preventDefault();
    onSelect(actor.guid);
  }
</script>

<StatRow {figures} {phone} testid={`actor-${actor.guid}`}>
  {#snippet icon()}<Crest cls={actor.class} size={28} />{/snippet}
  {#snippet name()}{#if href !== undefined}<a
        {href}
        class="font-bold hover:underline [@media(pointer:coarse)]:-my-[13px] [@media(pointer:coarse)]:inline-block [@media(pointer:coarse)]:py-[13px]"
        style:color={colour}
        title={label.title}
        aria-current={selected ? 'true' : undefined}
        data-testid="row-name"
        onclick={follow}>{label.name}<span class="text-gold ml-1" aria-hidden="true">›</span></a
      >{:else}<b style:color={colour} title={label.title} data-testid="row-name">{label.name}</b
      >{/if}{#if label.realm !== ''}<span class="text-muted ml-1 text-[12px]">· {label.realm}</span
      >{/if}{/snippet}
  {#snippet sub()}{subLine || ' '}{/snippet}
  {#snippet meter()}<Meter value={ofTop} tone="class" color={colour} />{/snippet}
</StatRow>
<CoachRowLine {actor} rowOpen={selected} />
