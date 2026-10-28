// web/src/lib/coach/provider.ts
// The processor that hosts the coach's model slug on the Vercel AI Gateway, named once: the
// privacy page renders through it. It is the provider pinned in `LLM_PROVIDERS`, not the model's
// maker; Task 40 sets it from the pinned provider after the evaluation. "DeepSeek" is the
// launch default until then.
export const COACH_MODEL_PROVIDER = 'DeepSeek';
