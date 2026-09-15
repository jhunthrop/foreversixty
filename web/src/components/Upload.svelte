<!-- web/src/components/Upload.svelte -->
<!-- Whole-file upload for a night that was already logged. The report page fills in as the
     parse job closes each fight, so this hands off to /reports/<id> the moment the API
     accepts the upload rather than waiting for a parse it cannot see. -->
<script lang="ts">
  import { formatAmount } from '../lib/report/format';
  import {
    MAX_UPLOAD_BYTES,
    UPLOAD_CANCELLED,
    UPLOAD_TOO_LARGE,
    completeUpload,
    createUpload,
    uploadParts,
  } from '../lib/upload/multipart';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';

  const VISIBILITIES = [
    { id: 'public', label: 'Public', note: 'Listed, ranked, anyone can open it.' },
    { id: 'unlisted', label: 'Unlisted', note: 'Ranked, but only people with the link can open it.' },
    { id: 'private', label: 'Private', note: 'Only you. Never ranked.' },
    { id: 'guild', label: 'Guild', note: 'Your guild page only. Ranked.' },
  ] as const;

  let file = $state<File | null>(null);
  let title = $state('');
  let visibility = $state<string>('public');
  let phase = $state<'idle' | 'uploading' | 'finishing' | 'failed'>('idle');
  let uploaded = $state(0);
  let error = $state('');
  /** Non-null only while an upload is in flight; `cancel` is the only thing that fires it. */
  let inFlight: AbortController | null = null;

  const percent = $derived(file === null || uploaded === 0 ? 0 : Math.round((uploaded / file.size) * 100));
  const tooLarge = $derived(file !== null && file.size > MAX_UPLOAD_BYTES);

  function pick(event: Event): void {
    const input = event.currentTarget as HTMLInputElement;
    file = input.files?.[0] ?? null;
    uploaded = 0;
    error = tooLarge ? UPLOAD_TOO_LARGE : '';
    phase = 'idle';
  }

  function onDrop(event: DragEvent): void {
    event.preventDefault();
    file = event.dataTransfer?.files?.[0] ?? null;
    uploaded = 0;
    error = tooLarge ? UPLOAD_TOO_LARGE : '';
  }

  /**
   * Cancels the part in flight, not just the queue: multipart.ts hands the signal to the
   * XMLHttpRequest, so a 64 MiB part that is halfway up stops there rather than finishing
   * first. Nothing is uploaded afterwards, and the multipart upload is left for the API to
   * clean up -- the browser has no endpoint to abort it with.
   */
  function cancel(): void {
    inFlight?.abort();
  }

  async function start(): Promise<void> {
    if (file === null || tooLarge) return;
    const chosen = file;
    const controller = new AbortController();
    inFlight = controller;
    phase = 'uploading';
    error = '';
    uploaded = 0;
    try {
      const created = await createUpload(chosen);
      // `restart` is how an upload that outlives its own signed urls -- an hour, and a
      // 4 GiB night on a home uplink is longer -- recovers instead of throwing away every
      // part it already sent. It is a fresh upload, not a re-signing of this one, so the
      // id to complete is whichever one uploadParts actually finished against.
      const { uploadId, etags } = await uploadParts(
        chosen,
        created,
        (progress) => {
          uploaded = progress.uploadedBytes;
        },
        { signal: controller.signal, restart: () => createUpload(chosen) },
      );
      phase = 'finishing';
      const reportId = await completeUpload(uploadId, etags, { title: title.trim(), visibility });
      window.location.assign(`/reports/${reportId}`);
    } catch (thrown) {
      const message = thrown instanceof Error ? thrown.message : UPLOAD_TOO_LARGE;
      // Cancelling is not a failure to report back: the visitor asked for it and already
      // knows. The form goes back to where it was, with the same file still chosen.
      if (message === UPLOAD_CANCELLED) {
        phase = 'idle';
        uploaded = 0;
        return;
      }
      phase = 'failed';
      error = message;
    } finally {
      if (inFlight === controller) inFlight = null;
    }
  }
</script>

<section class="flex flex-col gap-4" data-testid="upload">
  <h2 class="section-title text-[18px]">Upload a log</h2>
  <p class="text-muted text-[14px]">
    A whole <code class="font-mono">WoWCombatLog.txt</code>. The first fight is readable within ten seconds of
    the upload finishing, and the rest appear as they are parsed.
  </p>

  <div
    class="border-line rounded-panel flex flex-col gap-3 border border-dashed p-4"
    ondrop={onDrop}
    ondragover={(event) => event.preventDefault()}
    role="group"
    aria-label="Choose a combat log"
  >
    <label class="label text-muted" for="upload-file">Combat log</label>
    <input
      id="upload-file"
      class="text-text min-h-11 text-[14px]"
      type="file"
      accept=".txt,text/plain"
      onchange={pick}
      disabled={phase === 'uploading' || phase === 'finishing'}
      data-testid="upload-file"
    />
    {#if file !== null}
      <p class="text-muted tabular font-mono text-[13px]" data-testid="upload-size">
        {file.name} · {formatAmount(file.size)} bytes
      </p>
    {/if}
  </div>

  <label class="label text-muted" for="upload-title">Title</label>
  <input
    id="upload-title"
    class="border-line-warm bg-raised rounded-control text-text h-11 px-3 text-[15px]"
    type="text"
    maxlength="60"
    bind:value={title}
    disabled={phase === 'uploading' || phase === 'finishing'}
  />

  <fieldset class="flex flex-col gap-2">
    <legend class="label text-muted">Who can see it</legend>
    {#each VISIBILITIES as option (option.id)}
      <label class="flex min-h-11 items-start gap-3 text-[14px]">
        <input type="radio" name="visibility" value={option.id} bind:group={visibility} class="mt-1" />
        <span>
          {option.label}
          <span class="text-muted block text-[13px]">{option.note}</span>
        </span>
      </label>
    {/each}
  </fieldset>

  <div class="flex flex-wrap items-center gap-3">
    <button
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
      onclick={() => void start()}
      disabled={file === null || tooLarge || phase === 'uploading' || phase === 'finishing'}
      data-testid="upload-start"
    >
      Upload
    </button>
    {#if phase === 'uploading'}
      <button
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-muted w-fit px-4"
        onclick={cancel}
        data-testid="upload-cancel"
      >
        Cancel
      </button>
    {/if}
  </div>

  {#if phase === 'uploading' || phase === 'finishing'}
    <div class="flex flex-col gap-2" data-testid="upload-progress">
      <div
        class="bg-line-soft h-[6px] w-full overflow-hidden"
        role="progressbar"
        aria-valuenow={percent}
        aria-valuemin="0"
        aria-valuemax="100"
      >
        <div class="bg-gold h-full" style={`width: ${percent}%`}></div>
      </div>
      <p class="text-muted tabular font-mono text-[13px]">
        {phase === 'finishing' ? 'Finishing' : `${percent}%`}
      </p>
    </div>
  {/if}

  {#if error !== ''}
    <p class="text-[14px]" role="alert" data-testid="upload-error">{error}</p>
  {/if}
</section>
