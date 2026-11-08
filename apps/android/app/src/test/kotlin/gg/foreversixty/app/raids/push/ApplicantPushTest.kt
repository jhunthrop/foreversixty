package gg.foreversixty.app.raids.push

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.DiscordOutcome
import gg.foreversixty.app.model.GuildApplicationDecision
import gg.foreversixty.app.model.GuildOfficerApplications
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.push.PushMessage
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidsDestination
import gg.foreversixty.app.raids.RaidsNavigator
import gg.foreversixty.app.raids.flows.RaidFlowService
import gg.foreversixty.app.raids.resources
import java.io.IOException
import kotlinx.coroutines.test.runTest
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** guild.applicant (spec 2026-11-08 F3): the tap opens the applicant sheet, Decline answers in the shade. */
@RunWith(AndroidJUnit4::class)
class ApplicantPushTest {
    private val data = mapOf(
        "kind" to ApplicantPushPayload.KIND,
        "group" to "officer",
        "title" to "New applicant: Test Applicant, Holy Paladin 60",
        "body" to "Applied to Olympus XXVII.",
        "link" to "https://foreversixty.gg/app/guild/2/applicants/41",
        "category" to RaidPushCategory.SIGNUP,
        "guild_id" to "2",
        "application_id" to "41",
    )
    private val payload = ApplicantPushPayload(2, 41)
    private val application = RaidFixture.load(GuildOfficerApplications.serializer(), "guild.applications").rows.first { it.id == 41L }

    private fun decision(dm: Boolean) = GuildApplicationDecision(
        application = application.copy(status = "declined"),
        discordRole = DiscordOutcome(done = false, reason = "not_accepted"),
        discordDm = DiscordOutcome(done = dm, reason = if (dm) null else "not_opted_in"),
    )

    private fun server(status: Int, code: String) = ApiResult.Failure(ApiError.Server(status, code, code, emptyMap(), "test"))

    private suspend fun decline(result: ApiResult<GuildApplicationDecision>): Pair<ApplicantDeclineOutcome, FakeFlows> {
        val flows = FakeFlows(result)
        return ApplicantPushDecliner(flows, resources()).decline(payload) to flows
    }

    @Test
    fun `a tap on guild_applicant lands on the applicant sheet, one notification per application`() {
        val message = PushMessage.parse(data, null, null, collapse = null)
        assertEquals(DeepLink.GuildApplicant(2, 41).url(), message?.landing())
        assertEquals("guild.applicant-41", message?.collapse)
        // The ids decide, even when the link names another host.
        val staging = PushMessage.parse(data + ("link" to "https://staging.example/app/guild/2/applicants/41"), null, null, null)
        assertEquals(DeepLink.GuildApplicant(2, 41).url(), staging?.landing())
    }

    @Test
    fun `the applicant link opens the Raids tab on the applicant`() {
        val navigator = RaidsNavigator()
        navigator.push(RaidsDestination.Night(9))
        assertTrue(navigator.open(requireNotNull(DeepLink.parse("https://foreversixty.gg/app/guild/2/applicants/41"))))
        assertEquals(listOf<RaidsDestination>(RaidsDestination.Applicant(2, 41)), navigator.stack.toList())
        assertTrue(navigator.wantsTab)
    }

    @Test
    fun `the payload needs both ids`() {
        assertEquals(payload, ApplicantPushPayload.parse(data))
        assertNull(ApplicantPushPayload.parse(data - "application_id"))
        assertNull(ApplicantPushPayload.parse(data + ("guild_id" to "x")))
    }

    @Test
    fun `Decline posts the decline and names the applicant`() = runTest {
        val (outcome, flows) = decline(ApiResult.Success(decision(dm = false), "test"))
        assertEquals(listOf(2L to 41L), flows.declined)
        assertEquals(ApplicantDeclineOutcome.Answered("Declined Test Applicant", ""), outcome)
    }

    @Test
    fun `a Discord message the bot sent is said`() = runTest {
        val (outcome) = decline(ApiResult.Success(decision(dm = true), "test"))
        assertEquals(ApplicantDeclineOutcome.Answered("Declined Test Applicant", "They got a Discord message saying so."), outcome)
    }

    @Test
    fun `an application already decided says so, without Decline`() = runTest {
        val (outcome) = decline(server(409, "application_closed"))
        assertEquals(
            ApplicantDeclineOutcome.Answered("Already decided", "Another officer answered this application first."),
            outcome,
        )
    }

    @Test
    fun `a refusal the shade cannot word opens the sheet`() = runTest {
        assertEquals(ApplicantDeclineOutcome.Open(DeepLink.GuildApplicant(2, 41)), decline(server(403, "forbidden")).first)
        assertEquals(ApplicantDeclineOutcome.Open(DeepLink.GuildApplicant(2, 41)), decline(server(404, "not_found")).first)
    }

    @Test
    fun `a lost connection keeps Decline to try again`() = runTest {
        val expected = ApplicantDeclineOutcome.Answered("The decline did not go through. Try again.", "", retry = true)
        assertEquals(expected, decline(ApiResult.Failure(ApiError.Network(IOException("offline")))).first)
        assertEquals(expected, decline(ApiResult.Failure(ApiError.Status(503))).first)
    }

    private class FakeFlows(private val result: ApiResult<GuildApplicationDecision>) : RaidFlowService {
        val declined = mutableListOf<Pair<Long, Long>>()

        override suspend fun decline(guild: Long, application: Long): ApiResult<GuildApplicationDecision> {
            declined += guild to application
            return result
        }

        override suspend fun applications(guild: Long) = error("unused")
        override suspend fun accept(guild: Long, application: Long) = error("unused")
        override suspend fun recruiting(spec: String, faction: String, region: String?, ruleset: String?) = error("unused")
        override suspend fun ping(guild: Long, raid: Long, role: String?) = error("unused")
        override suspend fun ready(characterKey: String, raid: Long) = error("unused")
    }
}
