package gg.foreversixty.app

import gg.foreversixty.app.network.FixtureVariant
import io.ktor.http.HttpStatusCode
import org.junit.Assert.assertEquals
import org.junit.Assert.assertThrows
import org.junit.Test

/** Round 4 a2: the debug fixture states read from the walk's intent extra. */
class DebugFixturesTest {
    @Test fun `stem=variant pairs become named variants`() {
        assertEquals(
            mapOf("guild.raid-badge" to FixtureVariant.Named("officer"), "me" to FixtureVariant.Named("empty")),
            DebugFixtures.parse("guild.raid-badge=officer, me=empty"),
        )
    }

    @Test fun `fail makes the route answer a 503 error (round 6 row 3)`() {
        assertEquals(
            mapOf("guild.raid.ping" to FixtureVariant.Error("unavailable", HttpStatusCode.ServiceUnavailable)),
            DebugFixtures.parse("guild.raid.ping=${DebugFixtures.FAIL}"),
        )
    }

    @Test fun `a pair without a variant fails fast`() {
        assertThrows(IllegalArgumentException::class.java) { DebugFixtures.parse("guild.raid-badge") }
    }
}
