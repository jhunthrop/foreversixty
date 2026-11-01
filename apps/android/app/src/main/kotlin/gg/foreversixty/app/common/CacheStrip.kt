package gg.foreversixty.app.common

import android.content.Context
import android.net.ConnectivityManager
import android.net.NetworkCapabilities
import gg.foreversixty.app.network.ApiError
import java.time.Duration
import java.time.Instant

/**
 * How long a saved answer counts as fresh (W3 §5.7 Offline mode and stale data). Past its limit a
 * failed refresh shows the stale strip; inside it the saved answer simply stays. Build data
 * (guides, tiers, BiS, talents, the search index) has no age limit: it is good until the active
 * build changes, so it never shows stale by age.
 */
object CacheLimits {
    val HOME: Duration = Duration.ofMinutes(5)
    val RAIDS: Duration = Duration.ofMinutes(5)
    val RANKINGS: Duration = Duration.ofMinutes(10)
    val CHARACTER: Duration = Duration.ofMinutes(10)
    val REPORTS: Duration = Duration.ofMinutes(10)

    /** Sims are not in the spec's table; they read like reports (ruling, lane F15). */
    val SIMS: Duration = Duration.ofMinutes(10)

    /** Roster and raid history are raid data (ruling, lane F15). */
    val GUILD: Duration = RAIDS
}

/** What a region draws after a refresh failed. */
sealed interface FailedRefresh {
    /** No connection: the saved answer under `Offline. Showing what you had at {time}.` */
    data class Offline(val since: Instant) : FailedRefresh

    /** Connected, but the server did not answer and the saved answer is past its limit. */
    data class Stale(val since: Instant) : FailedRefresh

    /** Connected, the request dropped or timed out, the saved answer is still fresh: it stays, no strip. */
    data object KeepSaved : FailedRefresh

    /** No saved answer, a 4xx, or a 5xx inside the limit: the region's own error state. */
    data object Error : FailedRefresh
}

/**
 * The one decision (W3 §5.7): offline when there is no connection and a saved answer exists; stale
 * when there is a connection, the request failed in transport, timed out or answered 5xx, a saved
 * answer exists and it is older than [limit]. A 4xx is always the region's error.
 */
object StripDecision {
    private const val SERVER_ERROR_FLOOR = 500

    fun afterFailure(error: ApiError, savedAt: Instant?, now: Instant, limit: Duration?, connected: Boolean): FailedRefresh {
        if (savedAt == null || !isTransient(error)) return FailedRefresh.Error
        if (error is ApiError.Network && !connected) return FailedRefresh.Offline(savedAt)
        val old = limit != null && Duration.between(savedAt, now) > limit
        return when {
            old -> FailedRefresh.Stale(savedAt)
            // Inside its limit a 5xx stays the region's error (v1 §5.0); a dropped request keeps the answer.
            error is ApiError.Network -> FailedRefresh.KeepSaved
            else -> FailedRefresh.Error
        }
    }

    /**
     * The saved [value] to draw after a failed refresh, with its strip ([SavedView.since] null: no
     * strip), or null when the region shows its error.
     */
    fun <T> savedView(error: ApiError, value: T?, savedAt: Instant?, now: Instant, limit: Duration?, connected: Boolean): SavedView<T>? {
        if (value == null) return null
        return when (val decision = afterFailure(error, savedAt, now, limit, connected)) {
            is FailedRefresh.Offline -> SavedView(value, decision.since, stale = false)
            is FailedRefresh.Stale -> SavedView(value, decision.since, stale = true)
            FailedRefresh.KeepSaved -> SavedView(value, null, stale = false)
            FailedRefresh.Error -> null
        }
    }

    /** Transport failures, timeouts and 5xx; never a 4xx, a decoding bug or the update gate. */
    fun isTransient(error: ApiError): Boolean = when (error) {
        is ApiError.Network -> true
        is ApiError.Status -> error.status >= SERVER_ERROR_FLOOR
        is ApiError.Server -> error.status >= SERVER_ERROR_FLOOR
        is ApiError.Decoding, is ApiError.UpdateRequired -> false
    }
}

/** A saved answer drawn after a failed refresh: under the offline strip, the stale strip ([stale]), or none. */
data class SavedView<T>(val value: T, val since: Instant?, val stale: Boolean)

/**
 * Whether the phone has a network that reaches the internet. A plain interface (not `fun`), so a
 * store's trailing lambda never turns into one by accident; tests use [fixed].
 */
interface Connectivity {
    fun isConnected(): Boolean

    companion object {
        /** No knowledge of the network: every transport failure reads as offline (the v1 rule). */
        val Unknown: Connectivity = fixed(false)

        fun fixed(connected: Boolean): Connectivity = from { connected }

        fun from(check: () -> Boolean): Connectivity = object : Connectivity {
            override fun isConnected(): Boolean = check()
        }

        fun of(context: Context): Connectivity {
            val manager = context.applicationContext.getSystemService(ConnectivityManager::class.java)
            return from {
                val capabilities = manager?.getNetworkCapabilities(manager.activeNetwork)
                capabilities?.hasCapability(NetworkCapabilities.NET_CAPABILITY_INTERNET) == true
            }
        }

        /** An unmetered network (Wi-Fi): the Offline data screen's `Download on Wi-Fi only`. */
        fun unmetered(context: Context): () -> Boolean {
            val manager = context.applicationContext.getSystemService(ConnectivityManager::class.java)
            return {
                val capabilities = manager?.getNetworkCapabilities(manager.activeNetwork)
                capabilities?.hasCapability(NetworkCapabilities.NET_CAPABILITY_NOT_METERED) == true
            }
        }
    }
}
