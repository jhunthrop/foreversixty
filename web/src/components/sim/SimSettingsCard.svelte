<!-- web/src/components/sim/SimSettingsCard.svelte -->
<!-- The saved run's Settings (spec 2026-11-04-sim-results 3.12), its own chunk: the fight, the
     target and the engine on the left, the buffs, debuffs and consumables as chips on the right
     with what changes nothing for this spec filed apart, and the copy and download actions at the
     foot. It loads what only it reads: the buff table, the spec's weights and the talent layout
     Copy for the game needs. DetailsCard stays for the live /sim page. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import activeBuild from '../../data/active-build.json';
  import { loadItems, loadTalents, loadWeights, dataUrl } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import type { Item } from '../../lib/planner/types';
  import { damageSchools, type SpecWeights } from '../../lib/sim/buff-effect';
  import { loadBuffNames } from '../../lib/sim/buff-names';
  import { gameCodeForSimResult } from '../../lib/sim/game-code';
  import { bandCopy } from '../../lib/sim/result-copy';
  import { settingsCopy as copy } from '../../lib/sim/settings-copy';
  import {
    settingsChipGroups,
    settingsRows,
    type SettingsBuffs,
    type SettingsChip,
  } from '../../lib/sim/settings-rows';
  import type { SimResult } from '../../lib/sim/types';
  import { isStale } from '../../lib/sim/version';
  import { button } from '../../lib/ui/button';
  import { createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import CopyForGame from '../CopyForGame.svelte';
  import Card from '../ui/Card.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import Tag from '../ui/Tag.svelte';

  let { result, simId = '' }: { result: SimResult; simId?: string } = $props();

  /** Chips a phone shows of one group before "Show all" (the page's one disclosure). */
  const PHONE_CHIPS = 6;
  /** The site's phone breakpoint, in step with the `md:` classes. */
  const PHONE_QUERY = '(max-width: 767px)';
  const COPY_FAILED_MS = 3000;
  const ENGINE_SHA_LENGTH = 7;

  const build = activeBuild.build;
  const classSlug = $derived(result.request.character.class);

  let buffs = $state<SettingsBuffs>({});
  let weights = $state<SpecWeights>({});
  let items = $state<ReadonlyMap<number, Item>>(new Map());
  let talentIndex = $state<TalentIndex | null>(null);

  // One load each, for the one result this card shows. A failed file leaves the card standing:
  // chips fall back to legible names and Copy for the game stays disabled.
  void (async () => {
    try {
      buffs = (await loadBuffNames(build)).entries;
    } catch (error) {
      console.error('sim settings: the buff table did not load', error);
    }
  })();
  void (async () => {
    try {
      const rows = await loadWeights(build);
      weights = rows.find((row) => row.spec === result.request.spec)?.weights ?? {};
    } catch (error) {
      console.error('sim settings: the stat weights did not load', error);
    }
  })();
  void (async () => {
    try {
      const [file, talents] = await Promise.all([loadItems(build, classSlug), loadTalents(build, classSlug)]);
      items = new Map(file.items.map((item) => [item.id, item]));
      talentIndex = indexTalents(talents);
    } catch (error) {
      console.error('sim settings: the character files did not load', error);
    }
  })();

  const rows = $derived(settingsRows(result));
  const stale = $derived(isStale(result.engine_version));
  const groups = $derived(
    settingsChipGroups(result, buffs, weights, damageSchools(result.summary.damage_done), (stem) =>
      dataUrl(build, `icons/${stem}.webp`),
    ),
  );
  const gameCode = $derived(gameCodeForSimResult(result, build, talentIndex, items));
  const fileName = $derived(copy.fileName(simId !== '' ? simId : (result.sim_id ?? 'result')));
  const resultJson = (): string => JSON.stringify(result, null, 2);

  function downloadJson(): void {
    const url = URL.createObjectURL(new Blob([resultJson()], { type: 'application/json' }));
    const anchor = document.createElement('a');
    anchor.href = url;
    anchor.download = fileName;
    anchor.click();
    URL.revokeObjectURL(url);
  }

  let phone = $state(false);
  let expanded = $state<ReadonlySet<string>>(new Set());
  $effect(() => {
    const query = window.matchMedia(PHONE_QUERY);
    const read = (): void => {
      phone = query.matches;
    };
    read();
    query.addEventListener('change', read);
    return () => query.removeEventListener('change', read);
  });
  const capped = (group: string): boolean => phone && !expanded.has(group);

  let copyState = $state<CopyFeedbackState>('idle');
  let failedTimer: ReturnType<typeof setTimeout> | undefined;
  const copier = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (next) => {
      copyState = next;
      clearTimeout(failedTimer);
      if (next === 'failed') failedTimer = setTimeout(() => (copyState = 'idle'), COPY_FAILED_MS);
    },
  );
  onDestroy(() => {
    copier.dispose();
    clearTimeout(failedTimer);
  });
  const jsonLabel = $derived.by(() => {
    if (copyState === 'copied') return bandCopy.copied;
    return copyState === 'failed' ? bandCopy.copyFailed : copy.copyJson;
  });
</script>

{#snippet chip(entry: SettingsChip, quiet: boolean)}
  <li
    class={`border-line rounded-control flex min-h-[30px] items-center gap-1.5 border px-2.5 py-1 text-[13px] ${quiet ? 'opacity-[.62]' : ''}`}
    data-testid="sim-settings-chip"
  >
    {#if entry.icon !== null}
      <img
        src={entry.icon}
        alt=""
        width="22"
        height="22"
        loading="lazy"
        class="h-[22px] w-[22px] rounded-[4px]"
      />
    {:else}
      <span class="h-[22px] w-[22px] shrink-0 rounded-[4px] bg-white/[.05]"></span>
    {/if}
    <span class="text-text"
      >{entry.name}{#if entry.reason !== null}<span class="text-muted"> · {entry.reason}</span>{/if}</span
    >
  </li>
{/snippet}

<Card
  glyph="gear"
  label={copy.label}
  title={copy.title}
  body="none"
  testid="sim-settings-card"
  class="mx-[18px] md:mx-0"
>
  <div class="border-line-soft grid grid-cols-1 border-t lg:grid-cols-2">
    <dl class="border-line-soft flex flex-col px-5 pt-1 pb-4 lg:border-r">
      {#each rows as row (row.id)}
        <div
          class="border-line-soft flex min-h-11 items-center justify-between gap-6 border-b py-2.5"
          data-testid={`sim-settings-${row.id}`}
        >
          <dt class="text-muted text-[14px]">{row.label}</dt>
          <dd class="text-strong numeric m-0 text-right text-[14px] font-semibold">{row.value}</dd>
        </div>
      {/each}
      <div class="flex min-h-11 items-center justify-between gap-6 py-2.5" data-testid="sim-settings-engine">
        <dt class="text-muted text-[14px]">{copy.keys.engine}</dt>
        <dd class="text-muted m-0 flex items-center gap-2 font-mono text-[13px]">
          {#if stale}<Tag tone="muted" testid="sim-stale-tag">{copy.olderEngine}</Tag>{/if}
          {result.engine_version.slice(0, ENGINE_SHA_LENGTH)}
        </dd>
      </div>
    </dl>
    <div class="flex flex-col gap-4 px-5 py-4">
      {#each groups as group (group.id)}
        {@const all = [...group.chips, ...group.noEffect]}
        {@const shownChips = capped(group.id) ? group.chips.slice(0, PHONE_CHIPS) : group.chips}
        {@const left = capped(group.id)
          ? Math.max(0, PHONE_CHIPS - shownChips.length)
          : group.noEffect.length}
        <section class="flex flex-col gap-2" data-testid={`sim-settings-group-${group.id}`}>
          <h3
            class="border-line-soft text-muted flex justify-between border-b pb-1.5 text-[13px] font-semibold"
          >
            <span>{group.label}</span><span class="numeric">{group.count}</span>
          </h3>
          {#if group.count === 0}
            <p class="text-muted text-[13px]">{copy.empty}</p>
          {:else}
            {#if shownChips.length > 0}
              <ul class="flex flex-wrap gap-2">
                {#each shownChips as entry (entry.id)}{@render chip(entry, false)}{/each}
              </ul>
            {/if}
            {#if left > 0}
              <p class="text-muted text-[13px]">{copy.noEffect}</p>
              <ul class="flex flex-wrap gap-2" data-testid="sim-settings-no-effect">
                {#each group.noEffect.slice(0, left) as entry (entry.id)}{@render chip(entry, true)}{/each}
              </ul>
            {/if}
            {#if capped(group.id) && all.length > PHONE_CHIPS}
              <button
                type="button"
                class={`${button('secondary', 'sm')} w-fit`}
                onclick={() => (expanded = new Set([...expanded, group.id]))}
                data-testid={`sim-settings-show-all-${group.id}`}
              >
                {copy.showAll(all.length)}
              </button>
            {/if}
          {/if}
        </section>
      {/each}
    </div>
  </div>
  <div class="border-line-soft grid grid-cols-2 items-center gap-3 border-t px-5 py-3 md:flex md:flex-wrap">
    <div class="col-span-2 md:col-span-1">
      <CopyForGame code={gameCode} testid="sim-saved-copy-for-game" />
    </div>
    <button
      type="button"
      class={button('secondary')}
      onclick={() => void copier.copy(resultJson())}
      data-testid="sim-copy-json"
    >
      <Glyph name={copyState === 'copied' ? 'check' : 'copy'} />
      <span aria-live="polite">{jsonLabel}</span>
    </button>
    <button type="button" class={button('secondary')} onclick={downloadJson} data-testid="sim-download-json">
      <Glyph name="download" />
      {copy.downloadJson}
    </button>
    <p class="text-muted col-span-2 text-[13px] md:ml-1">{copy.jsonNote}</p>
  </div>
</Card>
