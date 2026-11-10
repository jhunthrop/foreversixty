<!-- web/src/components/guild/setup/SetupPlanRail.svelte -->
<!-- The setup page's Plan panel (guild onboarding spec §3): the guild's plan, then one line each
     for what every claimed guild gets free and what Guild and Raid guild add, and Compare plans.
     No subscribe button: the checklist never sells. -->
<script lang="ts">
  import type { GuildSetup, SetupPlan } from '../../../lib/guild/setup-api';
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import { planLines, planName } from '../../../lib/guild/setup-model';
  import { LINK } from '../../../lib/ui/primitives';
  import Section from '../../ui/Section.svelte';

  let { plan, readiness = null }: { plan: SetupPlan; readiness?: GuildSetup['readiness'] } = $props();

  const lines = planLines();
</script>

<Section title={setupCopy.plan.heading} meta={planName(plan)} panel testid="setup-plan">
  <dl class="flex flex-col gap-2 text-[13px]" data-testid="setup-plan-lines">
    {#each lines as line (line.label)}
      <div>
        <dt class="text-strong inline font-semibold">{line.label}:</dt>
        <dd class="text-text inline">{line.text}</dd>
      </div>
    {/each}
  </dl>
  {#if readiness !== null && readiness.signed > 0}
    <p class="text-muted pt-3 text-[13px]" data-testid="setup-plan-readiness">
      {setupCopy.plan.readiness(readiness.raid_title, readiness.failing, readiness.signed)}
    </p>
  {/if}
  <p class="pt-3 text-[14px]"><a class={LINK} href="/premium#guild-leaders">{setupCopy.plan.compare} ›</a></p>
</Section>
