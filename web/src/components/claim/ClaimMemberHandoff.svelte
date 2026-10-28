<!-- web/src/components/claim/ClaimMemberHandoff.svelte -->
<!-- Spec 2026-10-27 §6.2: a member cannot claim, so the page hands them the link to send the
     guild master: the sentence and a copy field holding this page's URL with its plan. -->
<script lang="ts">
  import { checkoutGuildCopy } from '../../lib/billing/checkout-guild-copy';
  import type { GuildTier } from '../../lib/billing/contract';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import CopyLinkField from '../ui/CopyLinkField.svelte';

  let { guildName, url, tier = undefined }: { guildName: string; url: string; tier?: GuildTier } = $props();
</script>

<div class="flex flex-col gap-3" data-testid="guild-claim-not-eligible">
  <p class="text-[14px]">{guildClaimCopy.memberHandoff(guildName)}</p>
  <CopyLinkField
    {url}
    label={checkoutGuildCopy.copyField.labelGuildMaster(guildName)}
    shareTitle={checkoutGuildCopy.share.title(guildName)}
    shareText={tier === undefined ? undefined : checkoutGuildCopy.share.claimText(guildName, tier)}
    testid="guild-claim-handoff"
  />
</div>
