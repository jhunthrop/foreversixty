<!-- web/src/components/premium/GuildOnPlan.svelte -->
<!-- A guild plan column when one of the account's guilds is on it (premium spec 2026-11-08 section
     4.5): the sentence, and Manage billing (the guild's portal) when the account pays. -->
<script lang="ts">
  import { openPortal } from '../../lib/billing/api';
  import { checkoutCopy } from '../../lib/billing/copy';
  import { checkoutGuildCopy } from '../../lib/billing/checkout-guild-copy';
  import type { GuildOnPlanView } from '../../lib/billing/premium-page';
  import Button from '../ui/Button.svelte';

  let { view, portal = openPortal }: { view: GuildOnPlanView; portal?: typeof openPortal } = $props();

  let busy = $state(false);
  let failed = $state(false);

  async function manage(): Promise<void> {
    busy = true;
    failed = false;
    try {
      window.location.assign((await portal(view.guildId)).portal_url);
    } catch {
      failed = true;
      busy = false;
    }
  }
</script>

<p class="text-strong m-0 text-[15px] leading-6" data-testid="guild-on-plan">{view.line}</p>
{#if view.manage}
  <span class="self-start">
    <Button disabled={busy} onclick={() => void manage()} testid="guild-on-plan-manage"
      >{checkoutGuildCopy.onPlan.manageBilling}</Button
    >
  </span>
{/if}
{#if failed}<p class="text-wipe m-0 text-[14px]" role="alert">{checkoutCopy.genericFailure}</p>{/if}
