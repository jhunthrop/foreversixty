<!-- web/src/components/sim/MoreSettingsSheet.svelte -->
<!-- More settings (spec 2026-11-06-sim §5.2, ruling S7): the run line stays one line and the rest
     of the request opens in the overlay sheet: the execute phase, the target, the rotation the
     engine follows, the consumables, the custom buffs (Buffs: Custom), the whole request as JSON,
     and for a payer the server-run notification. Changes apply as they are made; the result is not
     re-run until the player presses Run again. -->
<script lang="ts">
  import { buffLabel, type BuffNames } from '../../lib/sim/buff-names';
  import { MAX_LEVEL } from '../../lib/sim/character';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { LANE_ITERATION_CEILING } from '../../lib/sim/precision';
  import type { RequestValidation } from '../../lib/sim/engine';
  import {
    TARGET_LEVELS,
    TARGET_ARMOR_BY_LEVEL,
    executePhaseOn,
    styleIdOf,
    targetLevelFor,
    withExecutePhase,
    withStyle,
    withTargetLevel,
    type SimSettings,
  } from '../../lib/sim/settings';
  import { simCopy } from '../../lib/sim/copy';
  import { isSimulatedSpec, specDisplayName } from '../../lib/sim/spec-label';
  import { fightStyle } from '../../lib/sim/styles';
  import type { SimRequest } from '../../lib/sim/types';
  import Button from '../ui/Button.svelte';
  import Facts from '../ui/Facts.svelte';
  import Segmented from '../ui/Segmented.svelte';
  import Select from '../ui/Select.svelte';
  import Sheet from '../ui/Sheet.svelte';
  import Switch from '../ui/Switch.svelte';
  import BuffPanel from './BuffPanel.svelte';
  import RequestDrawer from './RequestDrawer.svelte';

  let {
    open,
    settings,
    spec,
    level,
    build,
    classSlug,
    names,
    request,
    disabled,
    notifier = false,
    notifyWanted = false,
    premium = false,
    onserver,
    onclose,
    onchange,
    onnotify,
    onvalidate,
    onapply,
    onrun,
    onshare,
  }: {
    open: boolean;
    settings: SimSettings;
    spec: string;
    level: number;
    build: string;
    classSlug: string;
    names: BuffNames | null;
    request: SimRequest | null;
    disabled: boolean;
    /** The browser can notify (a payer's server runs finish while the tab is away). */
    notifier?: boolean;
    notifyWanted?: boolean;
    /** A payer: this request on our servers (the server lane), from here. */
    premium?: boolean;
    onserver?: () => void;
    onclose: () => void;
    onchange: (next: SimSettings) => void;
    onnotify?: (wanted: boolean) => void;
    onvalidate: (json: string) => Promise<RequestValidation>;
    onapply: (request: SimRequest) => void;
    onrun: (request: SimRequest) => void;
    onshare: (request: SimRequest) => string | null;
  } = $props();

  let showRequest = $state(false);
  // A spec the sim does not run has no rotation to name and nothing to send to our servers
  // (healer review, MAJOR: "Default for Restoration" was a false statement).
  const simulated = $derived(isSimulatedSpec(spec));
  const rotationValue = $derived(
    simulated
      ? copy.more.rotationValue(specDisplayName(spec))
      : simCopy.rotationNotSimulated(specDisplayName(spec)),
  );
  const executeOn = $derived(executePhaseOn(settings));
  const executePct = $derived(Math.round(settings.encounter.execute_ratio * 100));
  // The target resolved for this character (a level-33 export fights a level-36 target), and the
  // stored offset form the select carries (60 = same level, 63 = three above; settings.ts).
  const targetLevel = $derived(targetLevelFor(settings.encounter, level));
  const targetOffsetLevel = $derived(MAX_LEVEL + (targetLevel - level));
  const consumables = $derived(
    settings.consumables.map((id) => buffLabel(id, names)).filter((name) => name !== ''),
  );

  /** Back to the preset: the fight style's own settings, the buffs and consumables untouched. */
  function backToPreset(): void {
    const style = fightStyle(styleIdOf(settings)) ?? fightStyle('patchwerk');
    if (style !== null) onchange(withStyle(settings, style.id));
  }
</script>

<Sheet {open} title={copy.more.title} closeLabel={copy.more.close} {onclose} testid="sim-more-settings">
  {#snippet executeControl()}
    <span class="inline-flex flex-col items-end gap-1">
      <Segmented
        options={[
          { value: 'on', label: copy.more.on },
          { value: 'off', label: copy.more.off },
        ]}
        value={executeOn ? 'on' : 'off'}
        label={copy.more.execute}
        size="sm"
        {disabled}
        testid="sim-execute"
        onChange={(value) => onchange(withExecutePhase(settings, value === 'on'))}
      />
    </span>
  {/snippet}
  {#snippet targetControl()}
    <span class="inline-flex justify-end">
      <Select
        id="sim-target-level"
        label={copy.more.target}
        size="dense"
        fit
        {disabled}
        value={String(targetOffsetLevel)}
        options={TARGET_LEVELS.map((value) => ({
          value: String(value),
          label: copy.more.targetLevel(level + (value - MAX_LEVEL)),
        }))}
        testid="sim-target-level"
        onchange={(event) =>
          onchange(withTargetLevel(settings, Number((event.currentTarget as HTMLSelectElement).value)))}
      />
    </span>
  {/snippet}
  <Facts
    items={[
      {
        id: 'execute',
        label: copy.more.execute,
        content: executeControl,
        sub: executeOn ? copy.more.executeSub(executePct) : undefined,
      },
      {
        id: 'target',
        label: copy.more.target,
        content: targetControl,
        sub:
          TARGET_ARMOR_BY_LEVEL[targetLevel] === undefined
            ? undefined
            : copy.more.armorSub(TARGET_ARMOR_BY_LEVEL[targetLevel].toLocaleString('en-US')),
      },
      { id: 'rotation', label: copy.more.rotation, value: rotationValue, testid: 'sim-rotation' },
    ]}
  />

  {#if settings.preset === 'custom'}
    <div class="pt-4">
      <BuffPanel {settings} {build} {classSlug} {names} {disabled} {onchange} />
    </div>
  {:else if consumables.length > 0}
    <h3 class="text-label tracking-label text-muted mt-4 font-bold uppercase">
      Consumables · {consumables.length}
    </h3>
    <p class="text-text mt-2 text-[14px] leading-5" data-testid="sim-consumables">{consumables.join(', ')}</p>
  {/if}

  <h3 class="text-label tracking-label text-muted mt-5 font-bold uppercase">{copy.more.request}</h3>
  <p class="text-small text-muted mt-1">
    {copy.more.requestLine}
    <button
      type="button"
      class="text-gold font-semibold"
      onclick={() => (showRequest = !showRequest)}
      data-testid="sim-request-toggle">{showRequest ? copy.more.hideIt : copy.more.showIt}</button
    >
  </p>
  {#if showRequest}
    <div class="pt-3">
      <RequestDrawer {request} {disabled} {onvalidate} {onapply} {onrun} {onshare} />
    </div>
  {/if}

  {#if notifier && onnotify}
    <div class="flex items-center justify-between gap-4 pt-5">
      <span class="text-text text-[14px]">{copy.more.notify}</span>
      <Switch checked={notifyWanted} label={copy.more.notify} onchange={onnotify} testid="sim-notify" />
    </div>
  {/if}

  {#if premium && onserver && simulated}
    <div class="flex items-center justify-between gap-4 pt-5">
      <span class="text-text text-[14px]"
        >{copy.more.serverLine(LANE_ITERATION_CEILING.server.toLocaleString('en-US'))}</span
      >
      <Button size="sm" {disabled} onclick={onserver} testid="sim-server-run">{copy.bar.runOnServers}</Button>
    </div>
  {/if}

  {#snippet foot()}
    <Button {disabled} onclick={backToPreset} testid="sim-more-reset">{copy.more.backToPreset}</Button>
    <Button tone="primary" onclick={onclose} testid="sim-more-done">{copy.more.done}</Button>
  {/snippet}
</Sheet>
