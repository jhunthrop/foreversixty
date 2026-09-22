<!-- web/src/components/CharacterHandoffLinks.svelte -->
<!-- Open in simulator / Open in planner for one character, or what is needed instead.
     Shared by Account.svelte's Characters list and Character.svelte's header: both need
     the exact same addon-export check (lib/addon-export.ts), so it lives once. -->
<script lang="ts">
  import { lookupAddonExport } from '../lib/addon-export';
  import { handoffCopy } from '../lib/handoff-copy';
  import { plannerCodeHref, simCodeHref } from '../lib/handoff-links';
  import type { CharacterPath } from '../lib/characters';

  let { path, apiBase = undefined }: { path: CharacterPath; apiBase?: string } = $props();

  let code = $state<string | null>(null);
  let status = $state<'loading' | 'ready'>('loading');

  // One load per `path`; the reference-equality guard is the same "requested vs resolved"
  // pattern Character.svelte's own effect uses, so a prop change mid-flight cannot land a
  // stale result over a newer one.
  $effect(() => {
    const requested = path;
    status = 'loading';
    code = null;
    void lookupAddonExport(requested, apiBase).then((result) => {
      if (result.path !== requested) return;
      code = result.code;
      status = 'ready';
    });
  });

  /** A text link that still clears the 44px hit target on a phone. */
  const LINK = 'inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0';
</script>

<div class="flex min-h-11 flex-wrap items-center gap-3 md:min-h-0" data-testid="character-handoff">
  {#if status === 'loading'}
    <span class="invisible text-[13px]" aria-hidden="true">{handoffCopy.openInSimulator}</span>
  {:else if code !== null}
    <a class={LINK} href={simCodeHref(code)} data-testid="character-open-sim">{handoffCopy.openInSimulator}</a
    >
    <a class={LINK} href={plannerCodeHref(code)} data-testid="character-open-planner"
      >{handoffCopy.openInPlanner}</a
    >
  {:else}
    <span class="text-muted text-[13px]" data-testid="character-needs-addon">
      {handoffCopy.needsExportLead}
      <a class="text-text inline-flex min-h-11 items-center underline md:min-h-0" href={handoffCopy.pasteHref}
        >{handoffCopy.needsExportPasteLink}</a
      >
    </span>
  {/if}
</div>
