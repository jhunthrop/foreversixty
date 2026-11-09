<!-- web/src/components/creators/CreatorApply.svelte -->
<!-- /creators/apply (design/specs/2026-11-08-creator-program.md 3f): the application form (channel,
     main platform, a rough size, what they make, an optional contact email only the owner sees), then
     its status: Pending, Approved (Set up payouts) or Not this time (the form comes back). A creator
     already in the program sees their stats link instead. -->
<script lang="ts">
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import { fetchApplication, sendApplication, type MyApplicationRead } from '../../lib/creators/api';
  import { PLATFORMS, applyCopy as copy, type Platform } from '../../lib/creators/copy';
  import { ABOUT_MAX, dayLabel } from '../../lib/creators/model';
  import { button } from '../../lib/ui/button';
  import { FIELD_LABEL } from '../../lib/ui/field';
  import { LINK } from '../../lib/ui/primitives';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Field from '../ui/Field.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Tag from '../ui/Tag.svelte';

  type Phase = 'loading' | 'signed-out' | 'failed' | 'ready';

  let phase = $state<Phase>('loading');
  let mine = $state<MyApplicationRead | null>(null);
  let values = $state({ channel_url: '', audience: '', about: '', contact_email: '' });
  let platform = $state<Platform>('twitch');
  let sending = $state(false);
  let disclosed = $state(false);
  let fieldErrors = $state<Readonly<Record<string, string>>>({});
  let formError = $state('');

  async function load(): Promise<void> {
    phase = 'loading';
    try {
      if ((await fetchMeOnce()) === null) {
        phase = 'signed-out';
        return;
      }
      mine = await fetchApplication();
      phase = 'ready';
    } catch (error) {
      phase = error instanceof AccountError && error.status === 401 ? 'signed-out' : 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  /** The size as the API takes it: a whole number, null when left empty, NaN-proofed to -1 for a 400. */
  function audience(raw: string): number | null {
    const t = raw.replace(/[,\s]/g, '');
    if (t === '') return null;
    const n = Number(t);
    return Number.isInteger(n) ? n : -1;
  }

  async function onSubmit(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    fieldErrors = {};
    formError = '';
    if (!disclosed) {
      formError = copy.discloseNeeded;
      return;
    }
    sending = true;
    try {
      await sendApplication({
        channel_url: values.channel_url.trim(),
        platform,
        audience: audience(values.audience),
        about: values.about.trim(),
        contact_email: values.contact_email.trim(),
      });
      mine = await fetchApplication();
    } catch (error) {
      if (error instanceof AccountError && error.fields !== null) fieldErrors = error.fields;
      else
        formError = error instanceof AccountError && error.message !== '' ? error.message : copy.sendFailed;
    } finally {
      sending = false;
    }
  }

  const showForm = $derived(
    mine !== null &&
      !mine.creator &&
      (mine.status === null || (mine.status === 'declined' && !mine.reapply_on)),
  );
</script>

{#if phase === 'loading'}
  <div data-testid="apply-loading"><Skeleton lines={6} rowHeight="h-6" /></div>
{:else if phase === 'signed-out'}
  <SignInPrompt line={copy.signedOut} next="/creators/apply" testid="apply-signin" />
{:else if phase === 'failed'}
  <LoadError message={copy.sendFailed} onRetry={() => void load()} testid="apply-load-error" />
{:else if mine !== null}
  <div class="flex flex-col gap-8">
    {#if mine.creator}
      <Section title={copy.statusTitle} testid="apply-creator">
        <div class="flex flex-col gap-3">
          <p>{copy.alreadyCreator}</p>
          <p><a class={LINK} href="/creators/stats" data-testid="apply-stats">{copy.statsLink} ›</a></p>
        </div>
      </Section>
    {:else if mine.status !== null}
      <Section title={copy.statusTitle} testid={`apply-status-${mine.status}`}>
        <div class="flex flex-col gap-3">
          {#if mine.status === 'pending'}
            <div><Tag tone="muted">{copy.pendingTag}</Tag></div>
            <p>{copy.pending(dayLabel(mine.created_at ?? ''))}</p>
          {:else if mine.status === 'approved'}
            <div><Tag tone="success">{copy.approvedTag}</Tag></div>
            <p>{copy.approved}</p>
            <div>
              <a class={button('primary')} href="/creators/onboarding" data-testid="apply-setup"
                >{copy.setup}</a
              >
            </div>
          {:else}
            <div><Tag tone="muted">{copy.declinedTag}</Tag></div>
            <p>{copy.declined}</p>
            {#if mine.reapply_on}<p class="text-muted text-[14px]" data-testid="apply-reapply">
                {copy.reapplyOn(dayLabel(mine.reapply_on))}
              </p>{/if}
          {/if}
        </div>
      </Section>
    {/if}
    {#if showForm}
      <Section title={copy.formTitle} testid="apply-form-section">
        <form class="flex flex-col gap-3" onsubmit={onSubmit} novalidate data-testid="apply-form">
          <Field
            id="apply-channel"
            type="url"
            label={copy.fields.channel.label}
            hint={copy.fields.channel.hint}
            placeholder={copy.fields.channel.placeholder}
            error={fieldErrors.channel_url}
            autocomplete="url"
            required
            bind:value={values.channel_url}
            testid="apply-channel"
          />
          <fieldset class="m-0 flex flex-col gap-1.5 border-0 p-0" data-testid="apply-platform">
            <legend class={FIELD_LABEL}>{copy.fields.platform.label}</legend>
            <div class="flex flex-wrap items-center gap-x-6 gap-y-2">
              {#each PLATFORMS as p (p)}
                <label class="flex min-h-11 items-center gap-2 text-[15px]">
                  <input
                    type="radio"
                    name="platform"
                    value={p}
                    bind:group={platform}
                    class="accent-gold size-[18px]"
                  />
                  {copy.platforms[p]}
                </label>
              {/each}
            </div>
            {#if fieldErrors.platform}<span class="text-wipe text-[13px]">{fieldErrors.platform}</span>{/if}
          </fieldset>
          <Field
            id="apply-audience"
            label={copy.fields.audience.label}
            hint={copy.fields.audience.hint}
            placeholder={copy.fields.audience.placeholder}
            inputmode="numeric"
            error={fieldErrors.audience}
            bind:value={values.audience}
            testid="apply-audience"
          />
          <Field
            id="apply-about"
            label={copy.fields.about.label}
            hint={copy.counter([...values.about].length, ABOUT_MAX)}
            rows={5}
            maxlength={ABOUT_MAX}
            required
            error={fieldErrors.about}
            bind:value={values.about}
            testid="apply-about"
          />
          <Field
            id="apply-email"
            type="email"
            label={copy.fields.email.label}
            hint={copy.fields.email.hint}
            placeholder={copy.fields.email.placeholder}
            autocomplete="email"
            error={fieldErrors.contact_email}
            bind:value={values.contact_email}
            testid="apply-email"
          />
          <label class="flex min-h-11 items-start gap-2 text-[15px]" data-testid="apply-disclose">
            <input
              type="checkbox"
              bind:checked={disclosed}
              class="accent-gold mt-[3px] size-[18px] shrink-0"
            />
            <span>{copy.disclose}</span>
          </label>
          <p class="text-muted text-[14px]">{copy.consent}</p>
          <div>
            <button
              type="submit"
              class={button('primary')}
              disabled={sending}
              aria-busy={sending}
              data-testid="apply-submit">{sending ? copy.sending : copy.submit}</button
            >
          </div>
          {#if formError !== ''}<p class="text-wipe text-[14px]" role="alert" data-testid="apply-error">
              {formError}
            </p>{/if}
        </form>
      </Section>
    {/if}
  </div>
{/if}
