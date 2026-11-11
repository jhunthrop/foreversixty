<!-- web/src/components/guild/GuildCharacterName.svelte -->
<!-- A guild character named the way the game names one (spec 2026-11-05 §3.3, §3.4): its ringed
     crest, then the name in class colour, 600. The claim holder, the invite link's maker and the
     inviter on the join card all read this way. crestSize 0 drops the crest. -->
<script lang="ts">
  import type { ClaimHolder } from '../../lib/guild/api';
  import { classColorVar } from '../../lib/report/format';
  import Crest from '../ui/Crest.svelte';

  let { holder, crestSize = 20 }: { holder: ClaimHolder; crestSize?: number } = $props();
</script>

{#if crestSize > 0}<Crest cls={holder.class ?? ''} size={crestSize} />{/if}<span
  class="font-semibold"
  style:color={classColorVar(holder.class ?? undefined)}>{holder.name}</span
>
