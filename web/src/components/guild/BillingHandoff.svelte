<!-- web/src/components/guild/BillingHandoff.svelte -->
<!-- Spec 2026-10-27 §6.4: under the read-only sentence, an officer who does not hold the claim gets
     a link to send the claim holder, naming the plan that fits the verified count. -->
<script lang="ts">
  import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
  import type { CharacterPath } from '../../lib/characters';
  import { billingHandoffUrl } from '../../lib/guild/billing-handoff';
  import { fitsTier } from '../../lib/guild/billing-state';
  import { guildSettingsCopy } from '../../lib/guild/copy';
  import CopyLinkField from '../ui/CopyLinkField.svelte';

  let { path, guildName, verifiedCount }: { path: CharacterPath; guildName: string; verifiedCount: number } =
    $props();

  const tier = $derived(fitsTier(verifiedCount));
</script>

<div class="flex max-w-[720px] flex-col gap-2" data-testid="guild-billing-handoff-block">
  <p class="text-muted text-[13px]">{guildSettingsCopy.sendThemThisLink}</p>
  <CopyLinkField
    url={billingHandoffUrl(path, verifiedCount)}
    label={copy.copyField.labelGuildMaster(guildName)}
    shareTitle={copy.share.title(guildName)}
    shareText={copy.share.settingsText(guildName, tier)}
    testid="guild-billing-handoff"
  />
</div>
