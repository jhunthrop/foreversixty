package gg.foreversixty.app.guides.logic

import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticResource
import java.time.Instant

/**
 * The static files F9 reads (W3-C1, plus active-build and dates), each kept for offline use. Live
 * is [LiveGuideFiles]; tests answer from the contract fixtures.
 */
interface GuideFiles {
    /** The file now, saved for offline; offline, the saved copy; else the failure. */
    suspend fun <T> load(resource: StaticResource<T>): ApiResult<CachedFile<T>>

    /** The saved copy alone, with no request: what a screen draws at once. */
    fun <T> cached(resource: StaticResource<T>): CachedFile<T>?

    /** The saved copy's size in bytes, or null when the file is not on the phone. */
    fun savedBytes(resource: StaticResource<*>): Long?
}

/** The CDN through the app's one static data store (the same store BiS, Talents and Search use). */
class LiveGuideFiles(private val files: StaticDataStore) : GuideFiles {
    override suspend fun <T> load(resource: StaticResource<T>) = files.load(resource)
    override fun <T> cached(resource: StaticResource<T>) = files.cached(resource)
    override fun savedBytes(resource: StaticResource<*>): Long? = files.savedBytes(resource)
}

/** One region's data (spec v1 §5.0): loading, drawn (with when it was saved, offline), not saved, or failed. */
sealed interface Loadable<out T> {
    data object Loading : Loadable<Nothing>

    /** [offlineSince] is set while a saved copy shows because the read could not connect. */
    data class Loaded<T>(val value: T, val offlineSince: Instant?) : Loadable<T>

    /** Offline, and the file was never saved. */
    data object NotCached : Loadable<Nothing>

    /** The CDN answered with an error. */
    data object Failed : Loadable<Nothing>

    companion object {
        /** The saved copy as drawn before any request: never marked offline, since nothing has failed yet. */
        fun <T> saved(file: CachedFile<T>?): Loadable<T>? = file?.let { Loaded(it.value, offlineSince = null) }

        /** A read's answer over what is drawn: a failed refresh keeps the drawn file (no error over content). */
        fun <T> of(answer: ApiResult<CachedFile<T>>, current: Loadable<T>): Loadable<T> = when (answer) {
            is ApiResult.Success -> Loaded(answer.data.value, answer.data.offlineSince())
            is ApiResult.Failure -> when {
                current is Loaded -> current
                answer.error is ApiError.Network -> NotCached
                else -> Failed
            }
        }

        private fun CachedFile<*>.offlineSince(): Instant? = if (isFromCache) Instant.ofEpochMilli(savedAtMillis) else null
    }
}

/** The drawn value, or null while the region is not loaded. */
val <T> Loadable<T>.valueOrNull: T? get() = (this as? Loadable.Loaded)?.value
