package gg.foreversixty.app.common

import gg.foreversixty.app.network.ApiError
import java.io.IOException
import java.time.Duration
import java.time.Instant
import org.junit.Assert.assertEquals
import org.junit.Test

/** The strip decision (W3 §5.7): offline, stale, keep the saved answer, or the region's error. */
class StripDecisionTest {
    private val now = Instant.parse("2026-11-05T21:41:00Z")
    private val fresh = now.minusSeconds(60)
    private val old = now.minus(Duration.ofMinutes(6))
    private val network = ApiError.Network(IOException("timeout"))
    private val outage = ApiError.Server(503, "unavailable", "x", emptyMap(), "r")
    private val missing = ApiError.Server(404, "not_found", "x", emptyMap(), "r")

    private fun decide(error: ApiError, savedAt: Instant?, connected: Boolean, limit: Duration? = CacheLimits.HOME) =
        StripDecision.afterFailure(error, savedAt, now, limit, connected)

    @Test fun `no connection with a saved answer is offline, whatever its age`() {
        assertEquals(FailedRefresh.Offline(fresh), decide(network, fresh, connected = false))
        assertEquals(FailedRefresh.Offline(old), decide(network, old, connected = false))
    }

    @Test fun `connected, a timeout or a 5xx past the limit is stale`() {
        assertEquals(FailedRefresh.Stale(old), decide(network, old, connected = true))
        assertEquals(FailedRefresh.Stale(old), decide(outage, old, connected = true))
        assertEquals(FailedRefresh.Stale(old), decide(ApiError.Status(502), old, connected = true))
    }

    @Test fun `inside the limit a timeout keeps the saved answer and a 5xx is the error`() {
        assertEquals(FailedRefresh.KeepSaved, decide(network, fresh, connected = true))
        assertEquals(FailedRefresh.Error, decide(outage, fresh, connected = true))
    }

    @Test fun `a 4xx is always the region's error`() {
        assertEquals(FailedRefresh.Error, decide(missing, old, connected = true))
        assertEquals(FailedRefresh.Error, decide(ApiError.Status(404), old, connected = false))
    }

    @Test fun `nothing saved is the region's error`() {
        assertEquals(FailedRefresh.Error, decide(network, null, connected = false))
        assertEquals(FailedRefresh.Error, decide(outage, null, connected = true))
    }

    @Test fun `build data never goes stale by age`() {
        assertEquals(FailedRefresh.KeepSaved, decide(network, now.minus(Duration.ofDays(9)), connected = true, limit = null))
    }

    @Test fun `the limits are five minutes for home and raids, ten for rankings and reports`() {
        assertEquals(Duration.ofMinutes(5), CacheLimits.HOME)
        assertEquals(Duration.ofMinutes(5), CacheLimits.RAIDS)
        assertEquals(Duration.ofMinutes(10), CacheLimits.RANKINGS)
        assertEquals(Duration.ofMinutes(10), CacheLimits.CHARACTER)
        assertEquals(Duration.ofMinutes(10), CacheLimits.REPORTS)
    }
}
