<!-- web/src/components/guild/setup/SetupPlanRail.svelte -->
<!-- The setup page's Plan panel (guild onboarding spec §3, step 9): the guild's plan, what every
     claimed guild gets free, what Guild and Raid guild add, and one link to compare. No subscribe
     button: the checklist never sells. -->
<script lang="ts">
  import type { SetupPlan } from '../../../lib/guild/setup-api';
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import { planLists, planName } from '../../../lib/guild/setup-model';
  import Section from '../../ui/Section.svelte';

  let { plan }: { plan: SetupPlan } = $props();

  const lists = planLists();
</script>

<Section
  title={setupCopy.plan.heading}
  panel
  action={{ label: setupCopy.plan.compare, href: '/premium#guild-leaders' }}
  testid="setup-plan"
>
  <p class="text-strong text-[14px] font-semibold" data-testid="setup-plan-current">{planName(plan)}</p>
  {#each lists as list (list.heading)}
    <div class="pt-3">
      <h3 class="text-strong text-[14px] font-semibold">{list.heading}</h3>
      <ul class="mt-1 flex list-disc flex-col gap-0.5 pl-5 text-[13px]">
        {#each list.items as item (item)}<li>{item}</li>{/each}
      </ul>
    </div>
  {/each}
</Section>
