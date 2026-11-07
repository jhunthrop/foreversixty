<!-- web/src/components/premium/PremiumStateLine.svelte -->
<!-- /premium's one state line under the founder line (premium spec 2026-11-08 section 4.3): the last
     payment failed (an alert, Manage billing opens the portal), or a seat ended (G2: the way back
     at the account's own price). -->
<script lang="ts">
  import { openPortal } from '../../lib/billing/api';
  import { billingBlockCopy, checkoutCopy } from '../../lib/billing/copy';
  import type { PremiumStateLine } from '../../lib/billing/premium-page';
  import { premiumSourceCopy } from '../../lib/billing/premium-source-copy';
  import { LINK } from '../../lib/ui/primitives';
  import Tag from '../ui/Tag.svelte';

  let { line, portal = openPortal }: { line: Exclude<PremiumStateLine, null>; portal?: typeof openPortal } =
    $props();

  let busy = $state(false);
  let failed = $state(false);

  async function manage(): Promise<void> {
    busy = true;
    failed = false;
    try {
      window.location.assign((await portal(undefined)).portal_url);
    } catch {
      failed = true;
      busy = false;
    }
  }
</script>

{#if line.kind === 'pastDue'}
  <p class="premium-state" data-testid="premium-state-pastdue">
    <Tag tone="danger">{billingBlockCopy.pastDueTag}</Tag>
    <span class="text-strong" role="alert">{billingBlockCopy.pastDueBanner}</span>
    <button type="button" class={LINK} disabled={busy} aria-busy={busy} onclick={() => void manage()}
      >{premiumSourceCopy.buttons.manage} ›</button
    >
  </p>
  {#if failed}<p class="text-wipe m-0 text-[14px]" role="alert">{checkoutCopy.genericFailure}</p>{/if}
{:else}
  <p class="premium-state" data-testid="premium-state-ended">
    <span>{line.text}</span><a class={LINK} href={line.href}>{line.action} ›</a>
  </p>
{/if}

<style>
  .premium-state {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    gap: 6px 14px;
    margin: 0;
    padding-top: 14px;
    padding-bottom: 4px;
    font-size: 15px;
    line-height: 24px;
    color: var(--color-text);
  }
</style>
