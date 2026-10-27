<!-- web/src/components/guild/billing/GuildUpgradeConfirm.svelte -->
<!-- Move to Raid guild, confirmed in place (spec 3.2.4): the prorated sentence from the preview, a
     Skeleton in the amount's place and the commit disabled until it lands, focus on Keep Guild,
     Escape acts as Keep Guild. -->
<script lang="ts">
  import type { UpgradePreview } from '../../../lib/billing/contract';
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import { BILLING_TEXT_BUTTON } from '../../../lib/guild/billing-styles';
  import { focusOnMount } from '../../../lib/guild/focus-now';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import Skeleton from '../../ui/Skeleton.svelte';

  let {
    preview,
    committing,
    error,
    onCommit,
    onKeep,
  }: {
    /** Null while the preview loads. */
    preview: UpgradePreview | null;
    committing: boolean;
    error: string;
    onCommit: () => void;
    onKeep: () => void;
  } = $props();

  const parts = $derived(
    preview === null
      ? null
      : copy.upgrade.parts(
          preview.amount_due_cents,
          preview.interval,
          preview.renews_at,
          preview.next_amount_cents,
        ),
  );

  const lead = copy.upgrade.parts(0, 'monthly', '2026-01-01', 0).before;

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    onKeep();
  }
</script>

<div
  class="border-line-soft flex flex-col gap-3 border p-4"
  role="group"
  aria-labelledby="guild-upgrade-sentence"
  aria-busy={preview === null || committing}
  onkeydown={onKeydown}
  data-testid="guild-upgrade-confirm"
>
  <div id="guild-upgrade-sentence" class="flex flex-wrap items-center text-[14px]">
    {#if parts === null}
      <span>{lead}</span>
      <span class="w-16"
        ><Skeleton
          lines={1}
          rowHeight="h-3.5"
          label="Loading the amount"
          testid="guild-upgrade-skeleton"
        /></span
      >
    {:else}
      <span>{parts.before}{parts.amount}{parts.after}</span>
    {/if}
  </div>
  {#if error !== ''}
    <p class="text-[13px]" role="alert" data-testid="guild-upgrade-error">{error}</p>
  {/if}
  <div class="flex flex-wrap items-center gap-3">
    <button
      type="button"
      class={PRIMARY_BUTTON_FIXED + ' px-4'}
      disabled={preview === null || committing}
      onclick={onCommit}
      data-testid="guild-upgrade-commit"
    >
      {committing ? copy.upgrade.committing : copy.upgrade.commit}
    </button>
    <button
      type="button"
      class={BILLING_TEXT_BUTTON}
      use:focusOnMount
      onclick={onKeep}
      data-testid="guild-upgrade-keep"
    >
      {copy.upgrade.keep}
    </button>
  </div>
</div>
