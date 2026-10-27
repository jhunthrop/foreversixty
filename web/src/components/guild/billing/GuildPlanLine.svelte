<!-- web/src/components/guild/billing/GuildPlanLine.svelte -->
<!-- The plan line, line 2, the past-due alert and the payer's buttons (spec 3.2.1): Manage billing
     (the Stripe portal) and, on Guild only, Move to Raid guild unless the size line carries it. -->
<script lang="ts">
  import { BillingApiError, openPortal } from '../../../lib/billing/api';
  import type { GuildPlanView } from '../../../lib/billing/contract';
  import { billedLine, guildBillingCopy, guildPlanLine } from '../../../lib/billing/guild-billing-copy';
  import { billingBlockCopy } from '../../../lib/billing/copy';
  import { lineTwoState, type BillingViewState } from '../../../lib/guild/billing-state';
  import {
    BILLING_MOVE_BUTTON,
    BILLING_OUTLINE_BUTTON,
    OPENER_PRESSED,
  } from '../../../lib/guild/billing-styles';
  import { canManageBilling, officerToolsExplanation, settingsWords } from '../../../lib/guild/settings-view';
  import { BUSY_CLASS } from '../../../lib/ui/busy';

  let {
    guildId,
    plan,
    lineState,
    payer,
    showMove,
    expanded = false,
    onMove,
  }: {
    guildId: number;
    plan: GuildPlanView;
    lineState: BillingViewState;
    /** The reader holds the claim: only then do the buttons and the past-due alert draw. */
    payer: boolean;
    /** Move to Raid guild draws here (the size line is not showing). */
    showMove: boolean;
    /** The upgrade confirm is showing: Move to Raid guild draws pressed. */
    expanded?: boolean;
    onMove: () => void;
  } = $props();

  let busy = $state(false);
  /** `retry`: Stripe or the network failed, Try again may work. `refused`: a 403/404 never will. */
  let failure = $state<'retry' | 'refused' | null>(null);

  const planText = $derived.by(() => {
    if (lineState.kind === 'grant' || lineState.kind === 'none') return guildPlanLine(plan);
    return guildBillingCopy.planLine(lineState);
  });
  const lineTwo = $derived.by(() => {
    if (lineState.kind === 'grant') {
      return plan.scope === 'officer_tools' ? officerToolsExplanation(plan) : billedLine(plan);
    }
    const two = lineTwoState(lineState, plan);
    return two === null ? '' : guildBillingCopy.line2(two);
  });
  const manageable = $derived(payer && lineState.kind !== 'lapsed' && canManageBilling(plan));

  function isRetryable(thrown: unknown): boolean {
    return !(thrown instanceof BillingApiError) || thrown.status === 0 || thrown.status >= 500;
  }

  async function manage(): Promise<void> {
    busy = true;
    failure = null;
    try {
      const result = await openPortal(guildId);
      window.location.assign(result.portal_url);
    } catch (thrown) {
      // Retry helps for Stripe being down or the network, not for a refusal (403, 404).
      failure = isRetryable(thrown) ? 'retry' : 'refused';
    } finally {
      busy = false;
    }
  }
</script>

<div
  class="flex flex-col gap-1 md:flex-row md:items-start md:justify-between md:gap-4"
  data-testid="guild-billing-planline"
>
  <div class="flex flex-col gap-1">
    <p class="text-[14px]" data-testid="guild-billing-plan">
      {planText}{#if plan.founder && lineState.kind !== 'lapsed'}&#32;<span
          class="pill-sample"
          data-testid="guild-billing-founder">{guildBillingCopy.founderPill}</span
        >{/if}
    </p>
    <p class="text-muted text-[13px]" data-testid="guild-billing-contact">{lineTwo}</p>
    {#if payer && plan.status === 'past_due'}
      <p class="text-strong text-[13px]" role="alert">{billingBlockCopy.pastDueBanner}</p>
    {/if}
    {#if failure !== null}
      <p class="text-[14px]" role="alert" data-testid="guild-billing-unavailable">
        {settingsWords.billingUnavailable}
      </p>
    {/if}
  </div>
  {#if manageable || (payer && showMove)}
    <div class="mt-2 flex flex-wrap gap-2 md:mt-0">
      {#if manageable}
        <button
          type="button"
          class="{BILLING_OUTLINE_BUTTON} {busy ? BUSY_CLASS : ''}"
          disabled={busy}
          aria-busy={busy}
          onclick={() => void manage()}
          data-testid={failure === 'retry' ? 'guild-billing-retry' : 'guild-manage-billing'}
        >
          {failure === 'retry' ? settingsWords.tryAgain : billingBlockCopy.manageBilling}
        </button>
      {/if}
      {#if payer && showMove}
        <button
          type="button"
          class="{BILLING_MOVE_BUTTON} {expanded ? OPENER_PRESSED : ''}"
          aria-expanded={expanded}
          onclick={onMove}
          data-testid="guild-move-raid-guild"
        >
          {guildBillingCopy.size.move}
        </button>
      {/if}
    </div>
  {/if}
</div>
