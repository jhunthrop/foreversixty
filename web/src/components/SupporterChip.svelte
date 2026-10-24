<!-- web/src/components/SupporterChip.svelte -->
<!-- The Supporter mark: one gold text chip beside a name. Public surfaces show this or nothing;
     it is never drawn on a crest. -->
<script lang="ts">
  import { premiumBlockCopy } from '../lib/billing/copy';
</script>

<span
  class="label text-gold border-line-warm rounded-control shrink-0 border px-2 py-0.5"
  data-testid="supporter-chip">{premiumBlockCopy.supporterChip}</span
>
