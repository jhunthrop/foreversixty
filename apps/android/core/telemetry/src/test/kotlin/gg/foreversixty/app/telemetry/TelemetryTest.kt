package gg.foreversixty.app.telemetry

import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Test

class ScrubberTest {
    @Test
    fun `a URL loses its query, fragment and user info`() {
        assertEquals("https://foreversixty.gg/app/auth", Scrubber.url("https://foreversixty.gg/app/auth?code=abc&state=xyz"))
        assertEquals("https://api.foreversixty.gg/v1/me", Scrubber.url("https://api.foreversixty.gg/v1/me#frag"))
        assertEquals("https://api.foreversixty.gg/v1/me", Scrubber.url("https://user:pw@api.foreversixty.gg/v1/me"))
        assertEquals("https://api.foreversixty.gg/v1/sims/12", Scrubber.url("https://api.foreversixty.gg/v1/sims/12"))
    }

    @Test
    fun `free text loses tokens and credential parameters`() {
        assertEquals("token fsd_[redacted] expired", Scrubber.text("token fsd_AbC-123_x expired"))
        assertEquals("Authorization: Bearer [redacted]", Scrubber.text("Authorization: Bearer abc.def-ghi"))
        assertEquals(
            "callback /app/auth?code=[redacted]&state=[redacted] failed",
            Scrubber.text("callback /app/auth?code=s3cr3t&state=q failed"),
        )
        assertEquals(
            "body code_verifier=[redacted] and access_token=[redacted]",
            Scrubber.text("body code_verifier=v1 and access_token=t2"),
        )
        assertEquals("nothing secret here", Scrubber.text("nothing secret here"))
    }

    @Test
    fun `text that is not a URL is scrubbed as text`() {
        assertEquals("/app/auth?code=[redacted]", Scrubber.url("/app/auth?code=abc"))
        assertEquals("/app/auth?code=[redacted]", Scrubber.value("/app/auth?code=abc"))
    }
}

class CrashReportingPolicyTest {
    private val dsn = "https://public@o0.ingest.sentry.io/0"

    @Test
    fun `on only for a release build on the live API with a DSN`() {
        assertNotNull(CrashReportingPolicy.of(dsn, isDebugBuild = false, isLiveApi = true, version = "1.2.0", build = 42))
        assertNull(CrashReportingPolicy.of(dsn, isDebugBuild = true, isLiveApi = true, version = "1.2.0", build = 42))
        assertNull(CrashReportingPolicy.of(dsn, isDebugBuild = false, isLiveApi = false, version = "1.2.0", build = 42))
        assertNull(CrashReportingPolicy.of("", isDebugBuild = false, isLiveApi = true, version = "1.2.0", build = 42))
        assertNull(CrashReportingPolicy.of(null, isDebugBuild = false, isLiveApi = true, version = "1.2.0", build = 42))
    }

    @Test
    fun `the release names the app, version and build`() {
        val policy = CrashReportingPolicy.of(dsn, isDebugBuild = false, isLiveApi = true, version = "1.2.0", build = 42)
        assertEquals("gg.foreversixty.app@1.2.0+42", policy?.release)
    }
}
