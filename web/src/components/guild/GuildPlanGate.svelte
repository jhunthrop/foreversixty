<!-- web/src/components/guild/GuildPlanGate.svelte -->
<!-- The gate sentence (spec 2026-11-05 §3.2): officers and the guild master only, above the gated
     control. Part 1 names the plan the feature is part of, part 2 what this viewer can do about it
     (lib/guild/plan-gate.ts, drawn by GateActionInline). Its id is what each disabled control
     names in aria-describedby. The link is the inline-link recipe (DESIGN-SYSTEM.md, Accent). It
     never wraps; below md a 44px ::after centred on the link (an inline box is only its font's
     height, about 16px, so a fixed height, not 12px insets, is what reaches 44px) is its hit area
     on a phone, while the focus ring stays on the link and hugs the text. -->
<script lang="ts">
  import { guildGateCopy } from '../../lib/guild/copy';
  import { gateAction, type GateFacts, type GateFeature } from '../../lib/guild/plan-gate';
  import GateActionInline from './GateActionInline.svelte';

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
</script>

<p
  id={id === null ? undefined : (id ?? `officer-tools-gate-${tab}`)}
  class="text-text text-[13px] leading-5"
  data-testid={testid ?? `guild-${tab}-plan-gate`}
>
  {guildGateCopy.lead[tab]}
  <GateActionInline {action} />
</p>
