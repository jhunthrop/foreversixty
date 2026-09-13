<!-- web/src/components/planner/SharePanel.svelte -->
<!-- Saving a build and showing its link. The build stays in the store whatever the API
     says, so a failed save costs nothing but a retry. The preview card is the same PNG
     Discord and Twitter unfurl. -->
<script lang="ts">
  import { cardUrlFor, saveBuild, type SaveOutcome } from '../../lib/planner/share';
  import { MAX_TITLE_LENGTH, type PlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';

  // Retry and Copy link are both the neutral (non-gold) secondary button; only the layout
  // around them differs.
  const NEUTRAL_BUTTON = `${SECONDARY_BUTTON} border-line-warm text-text px-4`;

  let { store }: { store: PlannerStore } = $props();

  let saving = $state(false);
  let outcome = $state<SaveOutcome | null>(null);
  let copied = $state(false);
  let cardBroken = $state(false);

  const saved = $derived(outcome !== null && outcome.ok ? outcome.build : null);
  const failure = $derived(outcome !== null && !outcome.ok ? outcome : null);

  // Any edit invalidates the link that was shown: builds are immutable, so a changed build
  // is a different build.
  $effect(() => {
    void store.order.length;
    void store.gear;
    outcome = null;
    copied = false;
  });

  async function share(): Promise<void> {
    saving = true;
    cardBroken = false;
    copied = false;
    const draft = store.toDraft();
    const result = await saveBuild(draft);
    saving = false;
    // The grid stays editable while a save is in flight -- a slow API should not lock the
    // planner -- so the build this response describes may no longer be the current one.
    // Comparing the drafts rather than disabling input catches every kind of edit (points,
    // gear, title, even a class switch) and discards a response that arrives stale.
    if (JSON.stringify(store.toDraft()) !== JSON.stringify(draft)) return;
    outcome = result;
  }

  async function copy(url: string): Promise<void> {
    try {
      await navigator.clipboard.writeText(url);
      copied = true;
    } catch {
      copied = false;
    }
  }
</script>

<section class="flex w-full flex-col gap-3" data-testid="share-panel">
  <div class="flex flex-wrap items-end gap-3">
    <label class="flex flex-col gap-1">
      <span class="label text-muted">Title</span>
      <input
        type="text"
        maxlength={MAX_TITLE_LENGTH}
        value={store.title}
        oninput={(event) => store.setTitle(event.currentTarget.value)}
        placeholder="Optional"
        class="border-line-warm rounded-control bg-raised text-text placeholder:text-muted h-11 w-[260px] border px-3 text-[14px]"
      />
    </label>
    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
      disabled={saving || store.spent === 0}
      onclick={share}
    >
      {saving ? 'Saving' : 'Share'}
    </button>
  </div>

  {#if failure}
    <div class="border-line bg-raised rounded-panel flex flex-col gap-2 border p-3" role="alert">
      <p class="text-strong text-[14px]">{failure.message}</p>
      {#each Object.entries(failure.fields) as [field, message] (field)}
        <p class="text-muted text-[13px]"><span class="font-mono">{field}</span>: {message}</p>
      {/each}
      <button type="button" class="{NEUTRAL_BUTTON} w-fit" onclick={share}> Retry </button>
    </div>
  {/if}

  {#if saved}
    <div class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-3">
      <div class="flex flex-wrap items-center gap-3">
        <a href={saved.url} class="font-mono text-[14px]" data-testid="share-link">{saved.url}</a>
        <button type="button" class={NEUTRAL_BUTTON} onclick={() => copy(saved.url)}>
          {copied ? 'Copied' : 'Copy link'}
        </button>
      </div>
      {#if !cardBroken}
        <img
          src={cardUrlFor(saved.id)}
          alt={`Preview card for build ${saved.id}`}
          width="600"
          height="315"
          loading="lazy"
          decoding="async"
          class="rounded-panel border-line w-full max-w-[600px] border"
          onerror={() => (cardBroken = true)}
        />
      {/if}
    </div>
  {/if}
</section>
