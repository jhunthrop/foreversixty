<!-- web/src/components/guild/billing/GuildUpgradeConfirm.svelte -->
<!-- Move to Raid guild, confirmed in place (spec 3.2.4): the prorated sentence from the preview, a
     Skeleton in only the amount's place and the commit disabled until it lands, focus on Keep Guild
     (its ring shows on focus moved in code), Escape acts as Keep Guild. The renewal date is the plan
     line's own, so the two sentences never disagree. Below 768 the buttons are full width. -->
<script lang="ts">
  import type { GuildPlanView, UpgradePreview } from '../../../lib/billing/contract';
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import { CONFIRM_ACTION_WIDTH, CONFIRM_ACTIONS, CONFIRM_BLOCK } from '../../../lib/guild/billing-styles';
  import { upgradeFacts } from '../../../lib/guild/upgrade-facts';
  import ConfirmButtons from '../../ui/ConfirmButtons.svelte';
  import Skeleton from '../../ui/Skeleton.svelte';

  let {
    plan = null,
    preview,
    committing,
    error,
    onCommit,
    onKeep,
  }: {
    /** The plan being moved: its interval, founder price and renewal date frame the sentence. */
    plan?: Pick<GuildPlanView, 'interval' | 'founder' | 'current_period_end'> | null;
    /** Null while the preview loads. */
    preview: UpgradePreview | null;
    committing: boolean;
    error: string;
    onCommit: () => void;
    onKeep: () => void;
  } = $props();

  const facts = $derived(upgradeFacts(plan, preview));
  const parts = $derived(
    facts === null
      ? null
      : copy.upgrade.parts(facts.proratedCents ?? 0, facts.interval, facts.renewsAt, facts.nextCents),
  );
  const loading = $derived(preview === null);
</script>

<div
  class={CONFIRM_BLOCK}
  role="group"
  aria-labelledby="guild-upgrade-sentence"
  aria-busy={preview === null || committing}
  data-testid="guild-upgrade-confirm"
>
  <div id="guild-upgrade-sentence" class="flex flex-wrap items-center text-[14px]">
    {#if loading}
      <span>{copy.upgrade.leadIn}</span>
      <span class="w-16"
        ><Skeleton
          lines={1}
          rowHeight="h-3.5"
          label="Loading the amount"
          testid="guild-upgrade-skeleton"
        /></span
      >
      {#if parts !== null}<span data-testid="guild-upgrade-after">{parts.after}</span>{/if}
    {:else if parts !== null}
      <span>{parts.before}{parts.amount}{parts.after}</span>
    {/if}
  </div>
  {#if error !== ''}
    <p class="text-[13px]" role="alert" data-testid="guild-upgrade-error">{error}</p>
  {/if}
  <div class={CONFIRM_ACTIONS}>
    <ConfirmButtons
      tone="primary"
      size="md"
      width={CONFIRM_ACTION_WIDTH}
      confirm={copy.upgrade.commit}
      cancel={copy.upgrade.keep}
      busy={committing}
      busyLabel={copy.upgrade.committing}
      disabled={preview === null}
      onconfirm={onCommit}
      oncancel={onKeep}
      testid="guild-upgrade"
      confirmTestid="guild-upgrade-commit"
      cancelTestid="guild-upgrade-keep"
    />
  </div>
</div>
