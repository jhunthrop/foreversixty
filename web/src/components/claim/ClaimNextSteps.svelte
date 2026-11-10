<!-- web/src/components/claim/ClaimNextSteps.svelte -->
<!-- What the claim page shows once the viewer holds the claim (guild onboarding spec 2026-11-09
     §3.2): "Set up <Guild>", the checklist's steps in two columns with their live states, and
     Start setup, the page's one fill, to the setup page. A claim link made for a plan keeps its
     subscribe button as the fill, Start setup beside it; Approve your members stays a link. The steps read
     GET …/setup; until they answer the list holds its height with the step names alone. -->
<script lang="ts">
  import type { ClaimPlan } from '../../lib/guild/claim-plan';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import { fetchGuildSetup, type GuildSetup, type SetupStepKey } from '../../lib/guild/setup-api';
  import { setupCopy } from '../../lib/guild/setup-copy';
  import { nextStep, stepMark } from '../../lib/guild/setup-model';
  import { ROW_ITEM } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';
  import StepMark from '../ui/StepMark.svelte';
  import ClaimSubscribeButton from './ClaimSubscribeButton.svelte';

  const STEP_KEYS = Object.keys(setupCopy.titles) as SetupStepKey[];

  let {
    guildId,
    guildName,
    subscribe,
    claimPlan,
    guildBase,
    setupHref,
    initial = null,
  }: {
    guildId: number;
    guildName: string;
    /** Set when the guild is ready to subscribe: the id to check out and the button's label. */
    subscribe: { id: number; label: string } | null;
    claimPlan: ClaimPlan | null;
    guildBase: string;
    setupHref: string;
    /** A checklist already read (tests and the server render); the island reads its own. */
    initial?: GuildSetup | null;
  } = $props();

  let setup = $state<GuildSetup | null>(initial);

  $effect(() => {
    if (setup !== null) return;
    fetchGuildSetup(guildId)
      .then((fresh) => (setup = fresh))
      .catch(() => (setup = null));
  });

  /** A claim link made for a plan keeps its subscribe button as the page's one fill. */
  const subscribing = $derived(claimPlan !== null && subscribe !== null);
  const next = $derived(setup === null ? null : nextStep(setup.steps));
  const rows = $derived(
    STEP_KEYS.map((key, i) => {
      const state = setup?.steps.find((s) => s.key === key)?.state ?? 'todo';
      return { key, n: i + 1, state, mark: stepMark(state, key === next) };
    }),
  );
</script>

<Section
  title={setupCopy.claim.heading(guildName)}
  meta={setup === null ? undefined : `${setup.done} ${setupCopy.doneFigure(setup.done, setup.total)}`}
  panel
  testid="guild-claim-next"
>
  <ol class="grid gap-x-8 md:grid-cols-2" data-testid="guild-claim-setup-steps">
    {#each rows as row (row.key)}
      <li class="{ROW_ITEM} flex min-h-10 items-center gap-3 py-1.5 text-[14px]">
        <StepMark n={row.n} mark={row.mark} />
        <span class="text-strong flex-1">{setupCopy.titles[row.key]}</span>
        <span class="text-muted text-[13px]">{setupCopy.states[row.state]}</span>
      </li>
    {/each}
  </ol>
  <div class="flex flex-col gap-3 pt-4 md:flex-row md:flex-wrap [&>*]:max-md:w-full">
    {#if claimPlan !== null && subscribe !== null}
      <ClaimSubscribeButton guildId={subscribe.id} {claimPlan} label={subscribe.label} />
    {/if}
    <Button tone={subscribing ? 'secondary' : 'primary'} href={setupHref} testid="guild-claim-next-setup"
      >{setupCopy.claim.start}</Button
    >
    <Button href={`${guildBase}#roster`} testid="guild-claim-next-roster">{guildClaimCopy.nextRoster}</Button>
  </div>
</Section>
