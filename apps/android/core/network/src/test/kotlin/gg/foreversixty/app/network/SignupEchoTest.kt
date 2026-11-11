package gg.foreversixty.app.network

import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.model.SignupRequest
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test

/** Fixture mode answers a sign-up write with the tapped answer in place (the fixture app's answers move, as the server's do). */
class SignupEchoTest {
    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
    private val api = ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(source)))
    private val key = "us/pvp/gorvash"

    private suspend fun put(body: SignupRequest): SignupAnswer {
        val answer = api.send(ApiEndpoints.putRaidSignup(2, 31, "us", "pvp", "gorvash", body))
        return (answer as ApiResult.Success).data
    }

    @Test
    fun `Late with its minutes lands on the character's sign-up and viewer row`() = runTest {
        val raid = put(SignupRequest(state = "late", lateMinutes = 30)).raid
        assertEquals("late", raid.signups.first { it.characterKey == key }.state)
        assertEquals(30, raid.signups.first { it.characterKey == key }.lateMinutes)
        assertEquals("late", raid.viewer.characters.first { it.characterKey == key }.state)
    }

    @Test
    fun `Out lands and clears the minutes`() = runTest {
        val raid = put(SignupRequest(state = "out")).raid
        assertEquals("out", raid.viewer.characters.first { it.characterKey == key }.state)
        assertNull(raid.signups.first { it.characterKey == key }.lateMinutes)
    }

    @Test
    fun `anything that is not a sign-up is left alone`() {
        assertNull(SignupEcho.apply("{}", "/v1/me", """{"state":"in"}"""))
    }
}
