<!-- web/src/components/sim/MoreSettingsSheet.svelte -->
<!-- More settings (spec 2026-11-06-sim §5.2, ruling S7): the run line stays one line and the rest
     of the request opens in the overlay sheet: the execute phase, the target, the rotation the
     engine follows, the consumables, the custom buffs (Buffs: Custom), the whole request as JSON,
     and for a payer the server-run notification. Changes apply as they are made; the result is not
     re-run until the player presses Run again. -->
<script lang="ts">
  import type { BuffNames } from '../../lib/sim/buff-names';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
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
  import { specDisplayName } from '../../lib/sim/spec-label';
  import { fightStyle } from '../../lib/sim/styles';
  import type { SimRequest } from '../../lib/sim/types';
  import Button from '../ui/Button.svelte';
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
    onclose: () => void;
    onchange: (next: SimSettings) => void;
    onnotify?: (wanted: boolean) => void;
    onvalidate: (json: string) => Promise<RequestValidation>;
    onapply: (request: SimRequest) => void;
    onrun: (request: SimRequest) => void;
    onshare: (request: SimRequest) => string | null;
  } = $props();

  let showRequest = $state(false);
  const executeOn = $derived(executePhaseOn(settings));
  const executePct = $derived(Math.round(settings.encounter.execute_ratio * 100));
  const targetLevel = $derived(targetLevelFor(settings.encounter, level));
  const consumables = $derived(
    settings.consumables.map((id) => names?.entries[id]?.name ?? id).filter((name) => name !== ''),
  );

  /** Back to the preset: the fight style's own settings, the buffs and consumables untouched. */
  function backToPreset(): void {
    const style = fightStyle(styleIdOf(settings)) ?? fightStyle('patchwerk');
    if (style !== null) onchange(withStyle(settings, style.id));
  }
</script>

<Sheet {open} title={copy.more.title} closeLabel={copy.more.close} {onclose} testid="sim-more-settings">
  <dl class="sim-more-facts">
    <dt>{copy.more.execute}</dt>
    <dd>
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
      {#if executeOn}<small>{copy.more.executeSub(executePct)}</small>{/if}
    </dd>
    <dt>{copy.more.target}</dt>
    <dd>
      <Select
        id="sim-target-level"
        label={copy.more.target}
        size="dense"
        fit
        {disabled}
        value={String(targetLevel)}
        options={TARGET_LEVELS.map((value) => ({
          value: String(value),
          label: copy.more.targetLevel(value),
        }))}
        testid="sim-target-level"
        onchange={(event) =>
          onchange(withTargetLevel(settings, Number((event.currentTarget as HTMLSelectElement).value)))}
      />
      {#if TARGET_ARMOR_BY_LEVEL[targetLevel] !== undefined}<small
          >{copy.more.armorSub(TARGET_ARMOR_BY_LEVEL[targetLevel].toLocaleString('en-US'))}</small
        >{/if}
    </dd>
    <dt>Rotation</dt>
    <dd>The {specDisplayName(spec)} rotation</dd>
  </dl>

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

  {#snippet foot()}
    <Button {disabled} onclick={backToPreset} testid="sim-more-reset">{copy.more.backToPreset}</Button>
    <Button tone="primary" onclick={onclose} testid="sim-more-done">{copy.more.done}</Button>
  {/snippet}
</Sheet>

<style>
  .sim-more-facts {
    display: grid;
    grid-template-columns: minmax(0, auto) minmax(0, 1fr);
    margin: 0;
  }
  .sim-more-facts dt,
  .sim-more-facts dd {
    padding: 10px 0;
    border-bottom: 1px solid var(--color-line-soft);
    font-size: 14px;
    line-height: 20px;
    margin: 0;
  }
  .sim-more-facts dt {
    color: var(--color-muted);
    padding-right: 16px;
    align-self: center;
  }
  .sim-more-facts dd {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 4px;
    color: var(--color-strong);
    font-weight: 600;
  }
  .sim-more-facts dd small {
    font-weight: 400;
    color: var(--color-muted);
    font-size: 13px;
  }
</style>
