<!-- web/src/components/guild/GuildBillingSection.svelte -->
<!-- The Billing section of guild settings (launch pricing spec 3.2): it reads the guild billing view
     itself, reserves its ready height with Skeleton lines while that loads, and draws the blocks in
     order: plan line, size line, chooser, the seats and Who pays mount points, Records. -->
<script lang="ts">
  import { getGuildBilling, previewUpgrade, upgrade } from '../../lib/billing/guild-billing-api';
  import type { GuildBillingSettings, UpgradePreview } from '../../lib/billing/contract';
  import { guildBillingCopy as copy } from '../../lib/billing/guild-billing-copy';
  import {
    billingState,
    canBill,
    movesFromPlanLine,
    showSizeLine,
    showsChooser,
  } from '../../lib/guild/billing-state';
  import { guildRosterModelCopy } from '../../lib/guild/copy';
  import { GUILD_BILLING_LOADING } from '../../lib/guild/layout';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import GuildPlanChooser from './billing/GuildPlanChooser.svelte';
  import GuildPlanLine from './billing/GuildPlanLine.svelte';
  import GuildRecords from './billing/GuildRecords.svelte';
  import GuildSizeLine from './billing/GuildSizeLine.svelte';
  import GuildUpgradeConfirm from './billing/GuildUpgradeConfirm.svelte';

  let { guildId, now = new Date() }: { guildId: number; now?: Date } = $props();

  let view = $state<GuildBillingSettings | null>(null);
  let failed = $state(false);
  let confirming = $state(false);
  let preview = $state<UpgradePreview | null>(null);
  let committing = $state(false);
  let upgradeError = $state('');
  let moved = $state(false);

  const viewState = $derived(billingState(view, now));
  const payer = $derived(view !== null && canBill(view));

  async function load(): Promise<void> {
    failed = false;
    try {
      view = await getGuildBilling(guildId);
    } catch {
      failed = true;
    }
  }

  $effect(() => {
    void load();
  });

  async function openConfirm(): Promise<void> {
    confirming = true;
    preview = null;
    upgradeError = '';
    try {
      preview = await previewUpgrade(guildId);
    } catch {
      upgradeError = copy.upgrade.failed;
    }
  }

  function closeConfirm(): void {
    confirming = false;
    upgradeError = '';
  }

  async function commit(): Promise<void> {
    committing = true;
    upgradeError = '';
    try {
      view = await upgrade(guildId);
      confirming = false;
      moved = true;
    } catch {
      upgradeError = copy.upgrade.failed;
    } finally {
      committing = false;
    }
  }
</script>

<section id="billing" class="flex scroll-mt-20 flex-col gap-5" data-testid="guild-settings-billing">
  <h2 class="section-title text-[18px]">Billing</h2>
  {#if failed}
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="guild-billing-load-error" />
  {:else if view === null}
    <div class={GUILD_BILLING_LOADING.minHeight}>
      <Skeleton lines={GUILD_BILLING_LOADING.lines} testid="guild-billing-skeleton" />
    </div>
  {:else}
    {#if view.billing !== null && viewState.kind !== 'none'}
      <GuildPlanLine
        {guildId}
        plan={view.billing}
        lineState={viewState}
        {payer}
        showMove={movesFromPlanLine(view, viewState)}
        onMove={() => void openConfirm()}
      />
    {/if}
    {#if moved}
      <p class="text-muted text-[13px]" role="status" data-testid="guild-upgrade-done">{copy.upgrade.done}</p>
    {/if}
    {#if showSizeLine(view)}
      <GuildSizeLine count={view.verified_count} {payer} onMove={() => void openConfirm()} />
    {/if}
    {#if confirming}
      <GuildUpgradeConfirm
        {preview}
        {committing}
        error={upgradeError}
        onCommit={() => void commit()}
        onKeep={closeConfirm}
      />
    {/if}
    {#if showsChooser(viewState) && payer}
      <GuildPlanChooser
        {guildId}
        verifiedCount={view.verified_count}
        trialUsed={view.trial_used}
        lapsed={viewState.kind === 'lapsed'}
        {now}
      />
    {/if}
    <div data-testid="guild-billing-seats-slot"></div>
    <div data-testid="guild-billing-whopays-slot"></div>
    <GuildRecords {guildId} />
    {#if !payer}
      <p class="text-muted text-[13px]" data-testid="guild-settings-master-only">
        {guildRosterModelCopy.guildMasterOnly}
      </p>
    {/if}
  {/if}
</section>
