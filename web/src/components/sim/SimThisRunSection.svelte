<!-- web/src/components/sim/SimThisRunSection.svelte -->
<!-- This run (sim spec 2026-11-06 §4.5, §4.9; spec 2026-11-04-sim-results 3.12), its own chunk:
     the fight, the target and the engine as Facts, then On you, On the target and Consumables as
     IconLists (each count is the entries listed; the buffs with no effect on this spec dimmed and
     last), then the three small actions: Copy for the game, Copy JSON and Download JSON. It loads
     what only it reads: the buff table, the spec's weights and the talent layout Copy for the game
     needs. DetailsCard stays for bulk and weights results (SavedKind). -->
<script lang="ts">
  import { iconUrl } from '../../lib/ui/assets';
  import { onDestroy } from 'svelte';
  import activeBuild from '../../data/active-build.json';
  import { addonCopy } from '../../lib/addon/copy';
  import { copyForGameLabel, createCopyForGame, type CopyForGameState } from '../../lib/addon/copy-for-game';
  import { loadItems, loadTalents, loadWeights } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import type { Item } from '../../lib/planner/types';
  import { damageSchools, type SpecWeights } from '../../lib/sim/buff-effect';
  import { loadBuffNames } from '../../lib/sim/buff-names';
  import { gameCodeForSimResult } from '../../lib/sim/game-code';
  import { settingsCopy as copy } from '../../lib/sim/settings-copy';
  import {
    settingsChipGroups,
    settingsRows,
    type SettingsBuffs,
    type SettingsChip,
  } from '../../lib/sim/settings-rows';
  import type { SimResult } from '../../lib/sim/types';
  import type { Fact, IconListItem } from '../../lib/ui/facts';
  import Button from '../ui/Button.svelte';
  import CopyButton from '../ui/CopyButton.svelte';
  import Facts from '../ui/Facts.svelte';
  import IconList from '../ui/IconList.svelte';
  import Section from '../ui/Section.svelte';

  let { result, simId = '' }: { result: SimResult; simId?: string } = $props();

  const ENGINE_SHA_LENGTH = 7;
  const build = activeBuild.build;
  const classSlug = $derived(result.request.character.class);

  let buffs = $state<SettingsBuffs>({});
  let weights = $state<SpecWeights>({});
  let items = $state<ReadonlyMap<number, Item>>(new Map());
  let talentIndex = $state<TalentIndex | null>(null);

  // One load each, for the one result this section shows. A failed file leaves it standing: the
  // lists fall back to legible names and Copy for the game stays disabled.
  void (async () => {
    try {
      buffs = (await loadBuffNames(build)).entries;
    } catch (error) {
      console.error('sim this run: the buff table did not load', error);
    }
  })();
  void (async () => {
    try {
      const rows = await loadWeights(build);
      weights = rows.find((row) => row.spec === result.request.spec)?.weights ?? {};
    } catch (error) {
      console.error('sim this run: the stat weights did not load', error);
    }
  })();
  void (async () => {
    try {
      const [file, talents] = await Promise.all([loadItems(build, classSlug), loadTalents(build, classSlug)]);
      items = new Map(file.items.map((item) => [item.id, item]));
      talentIndex = indexTalents(talents);
    } catch (error) {
      console.error('sim this run: the character files did not load', error);
    }
  })();

  const facts = $derived<Fact[]>([
    ...settingsRows(result).map((row) => ({ ...row, testid: `sim-settings-${row.id}` })),
    {
      id: 'engine',
      label: copy.keys.engine,
      value: result.engine_version.slice(0, ENGINE_SHA_LENGTH),
      muted: true,
      testid: 'sim-settings-engine',
    },
  ]);
  const groups = $derived(
    settingsChipGroups(result, buffs, weights, damageSchools(result.summary.damage_done), (stem) =>
      iconUrl(build, stem),
    ),
  );

  const entry = (chip: SettingsChip, dim: boolean): IconListItem => ({
    id: chip.id,
    name: chip.reason === null ? chip.name : `${chip.name} · ${chip.reason}`,
    icon: chip.icon,
    dim,
  });

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

  let gameState = $state<CopyForGameState>('idle');
  const gameCopier = createCopyForGame((next) => (gameState = next));

  onDestroy(() => gameCopier.dispose());
</script>

<Section panel title={copy.title} testid="sim-settings-card">
  <Facts items={facts} />
  {#each groups as group (group.id)}
    <IconList
      label={copy.group(group.label, group.count)}
      items={[
        ...group.chips.map((chip) => entry(chip, false)),
        ...group.noEffect.map((chip) => entry(chip, true)),
      ]}
      testid={`sim-settings-group-${group.id}`}
    />
  {/each}
  <div class="flex flex-wrap gap-2 pt-5">
    <Button
      size="sm"
      disabled={gameCode === ''}
      onclick={() => void gameCopier.copy(gameCode)}
      testid="sim-saved-copy-for-game">{copyForGameLabel(gameState)}</Button
    >
    <CopyButton text={resultJson} label={copy.copyJson} testid="sim-copy-json" />
    <Button size="sm" onclick={downloadJson} testid="sim-download-json">{copy.downloadJson}</Button>
  </div>
  {#if gameState === 'failed'}
    <p class="text-muted pt-2 text-[13px]" role="status" data-testid="sim-saved-copy-for-game-failed">
      {addonCopy.copyFailed} <code class="text-text break-all select-all">{gameCode}</code>
    </p>
  {/if}
</Section>
