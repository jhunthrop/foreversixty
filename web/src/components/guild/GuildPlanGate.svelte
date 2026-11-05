<!-- web/src/components/guild/GuildPlanGate.svelte -->
<!-- The gate sentence (spec 2026-11-05 §3.2): officers and the guild master only, above the gated
     control. Part 1 names the plan the feature is part of, part 2 what this viewer can do about it
     (lib/guild/plan-gate.ts). Its id is what each disabled control names in aria-describedby. The
     link is the inline-link recipe (DESIGN-SYSTEM.md, Accent). It never wraps; below md a 44px
     ::after centred on the link (an inline box is only its font's height, about 16px, so a
     fixed height, not 12px insets, is what reaches 44px) is its hit area on a phone,
     while the focus ring stays on the link and hugs the text. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import { guildGateCopy } from '../../lib/guild/copy';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { gateAction, type GateFacts, type GateFeature } from '../../lib/guild/plan-gate';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  let {
    tab,
    facts,
    id = undefined,
    testid = undefined,
  }: {
    tab: GateFeature;
    facts: GateFacts;
    /** The element id the disabled controls name; `null` for a second copy on one page. */
    id?: string | null;
    testid?: string;
  } = $props();

  const action = $derived(gateAction(tab, facts));
  /** "Ask"'s word space, drawn inside the holder's flex box so it takes no width (guild round 2). */
  const WORD_SPACE = ' ';
</script>

<p
  id={id === null ? undefined : (id ?? `officer-tools-gate-${tab}`)}
  class="text-text text-[13px] leading-5"
  data-testid={testid ?? `guild-${tab}-plan-gate`}
>
  {guildGateCopy.lead[tab]}
  {#if action.kind === 'link'}<a class={INLINE_LINK} href={action.href}>{action.label}</a
    >.{:else if action.kind === 'note'}{action.text}{:else if action.holder === null}{guildGateCopy.askMaster}{:else}{guildGateCopy.ask.trimEnd()}<span
      class="ml-1 inline-flex items-baseline gap-1 font-semibold whitespace-nowrap"
      style:color={classColorVar(action.holder.class ?? undefined)}
      data-testid="guild-gate-holder"
      ><!-- The word space lives inside the flex box, which draws no whitespace-only run: the text
           still reads "Ask Thornhide" while the crest sits exactly 4 px (ml-1) from "Ask". -->{WORD_SPACE}<ClassCrestRing
        characterClass={action.holder.class ?? ''}
        size={16}
        class="size-4 self-center"
      />{action.holder.name}</span
    >{guildGateCopy.askTail}{/if}
</p>
