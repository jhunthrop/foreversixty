<!-- web/src/components/GuildRosterHandoff.svelte -->
<!-- Open in simulator / Open in planner for one guild roster row at gear/gear_bags
     consent. The guild home's own version of CharacterHandoffLinks.svelte: same
     lookupAddonExport check, but through lib/guild/roster-links.ts (spec section 4.1's
     "one helper") rather than handoff-links.ts directly, and with no "needs the addon"
     fallback line -- a roster row simply omits the links when there is no export, since
     "you have not run the companion" is already covered by this row's own "logged
     recently" state, and repeating it here would be the same fact said twice. -->
<script lang="ts">
  import { lookupAddonExport } from '../lib/addon-export';
  import type { CharacterPath } from '../lib/characters';
  import { rosterPlannerHref, rosterSimHref } from '../lib/guild/roster-links';
  import { guildHomeCopy } from '../lib/guild/copy';

  let { path }: { path: CharacterPath } = $props();

  let code = $state<string | null>(null);
  let status = $state<'loading' | 'ready'>('loading');

  $effect(() => {
    const requested = path;
    status = 'loading';
    code = null;
    void lookupAddonExport(requested).then((result) => {
      if (result.path !== requested) return;
      code = result.code;
      status = 'ready';
    });
  });

  const LINK = 'inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0';
</script>

<div class="flex min-h-11 flex-wrap items-center gap-3 md:min-h-0" data-testid="guild-roster-handoff">
  {#if status === 'loading'}
    <span class="invisible text-[13px]" aria-hidden="true">{guildHomeCopy.openSim}</span>
  {:else if code !== null}
    <a class={LINK} href={rosterSimHref(code)} data-testid="guild-roster-open-sim">
      {guildHomeCopy.openSim}
    </a>
    <a class={LINK} href={rosterPlannerHref(code)} data-testid="guild-roster-open-planner">
      {guildHomeCopy.openPlanner}
    </a>
  {/if}
</div>
