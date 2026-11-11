<!-- web/src/components/guild/GateActionInline.svelte -->
<!-- Part 2 of the gate sentence (spec 2026-11-05 §3.2), inline after part 1: the link, the note,
     the granted plan's `Ask us on Discord`, or `Ask {crest}{name} to turn it on.`. The gate
     sentence and the band's none-planned line both draw it, so the guild master meets the same
     link on every tab (guild round 3). -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import { guildGateCopy } from '../../lib/guild/copy';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import type { GateAction } from '../../lib/guild/plan-gate';
  import Crest from '../ui/Crest.svelte';

  let { action }: { action: GateAction } = $props();

  /** "Ask"'s word space, drawn inside the holder's flex box so it takes no width (guild round 2). */
  const WORD_SPACE = ' ';
</script>

{#if action.kind === 'link'}<a class={INLINE_LINK} href={action.href}>{action.label}</a
  >.{:else if action.kind === 'contact'}{action.before}<a
    class={INLINE_LINK}
    href={action.href}
    target="_blank"
    rel="noopener"
    data-testid="guild-gate-contact">{action.link}</a
  >{action.after}{:else if action.kind === 'note'}{action.text}{:else if action.holder === null}{guildGateCopy.askMaster}{:else}{guildGateCopy.ask.trimEnd()}<span
    class="ml-1 inline-flex items-baseline gap-1 font-semibold whitespace-nowrap"
    style:color={classColorVar(action.holder.class ?? undefined)}
    data-testid="guild-gate-holder"
    ><!-- The word space lives inside the flex box, which draws no whitespace-only run: the text
           still reads "Ask Thornhide" while the crest sits exactly 4 px (ml-1) from "Ask". -->{WORD_SPACE}<Crest
      cls={action.holder.class ?? ''}
      size={16}
      class="size-4 self-center"
    />{action.holder.name}</span
  >{guildGateCopy.askTail}{/if}
