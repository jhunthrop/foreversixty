package gg.foreversixty.app.model

import kotlinx.serialization.KSerializer
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.json.jsonObject
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull
import org.junit.Test

/**
 * The models generated from api/openapi.yaml decode every success fixture of their routes. The
 * Go contract test holds the fixtures to the spec; this holds the generated Kotlin to both.
 */
class GeneratedModelsTest {
    private fun raw(name: String): String =
        requireNotNull(javaClass.classLoader?.getResource("fixtures/app/$name.json")) {
            "fixture $name missing from docs/contracts/fixtures/app"
        }.readText()

    private fun <T> data(serializer: KSerializer<T>, name: String): T =
        ApiJson.decodeFromJsonElement(serializer, ApiJson.parseToJsonElement(raw(name)).jsonObject.getValue("data"))

    private fun <T> bare(serializer: KSerializer<T>, name: String): T = ApiJson.decodeFromString(serializer, raw(name))

    @Test
    fun `account and character fixtures decode`() {
        listOf("me", "me.premium", "me.empty", "me.patch").forEach { data(Me.serializer(), it) }
        listOf("character.card", "character.card.empty").forEach { data(CharacterProfile.serializer(), it) }
        listOf("character.home", "character.home.empty", "character.home.no-gear").forEach {
            val home = data(CharacterHome.serializer(), it)
            assertEquals(it, home.upgradeCount, home.upgrades.size)
        }
        listOf("character.search", "character.search.empty").forEach { data(CharacterSearch.serializer(), it) }
        listOf("character.sim-input.addon", "character.sim-input.fight").forEach { data(SimInput.serializer(), it) }
    }

    @Test
    fun `sim fixtures decode`() {
        data(SimList.serializer(), "sims.list")
        data(SimList.serializer(), "sims.list.empty")
        data(SimResult.serializer(), "sims.result")
        data(SimProgress.serializer(), "sims.progress")
        data(SimProgress.serializer(), "sims.progress.queued")
        data(SimStarted.serializer(), "sims.rerun")
    }

    @Test
    fun `guild fixtures decode`() {
        data(GuildHome.serializer(), "guild.home")
        data(GuildRaids.serializer(), "guild.raids")
        data(GuildRaids.serializer(), "guild.raids.public")
        listOf("guild.raid", "guild.raid.done", "guild.raid.patch").forEach { data(RaidView.serializer(), it) }
        data(SignupAnswer.serializer(), "guild.signup")
        assertNotNull(data(SignupAnswer.serializer(), "guild.signup.moved").moved)
        data(GuildLoot.serializer(), "guild.loot")
    }

    @Test
    fun `report, build and app fixtures decode`() {
        data(ReportList.serializer(), "reports.list")
        data(Report.serializer(), "reports.get")
        data(BuildList.serializer(), "builds.list")
        data(Build.serializer(), "builds.get")
        data(BuildSaved.serializer(), "builds.save")
        data(AppSession.serializer(), "app.signin.exchange")
        data(SignedOut.serializer(), "app.logout")
        data(NotificationSettings.serializer(), "app.notifications")
        data(PushTokenRegistered.serializer(), "app.push-tokens.register")
        data(AppConfig.serializer(), "app.config")
    }

    @Test
    fun `static files decode`() {
        bare(ActiveBuild.serializer(), "cdn.active-build")
        bare(ListSerializer(KeyDate.serializer()), "cdn.dates")
        bare(ListSerializer(ClassInfo.serializer()), "cdn.classes")
        bare(BisFile.serializer(), "cdn.bis.warrior-fury")
        bare(ItemsFile.serializer(), "cdn.items.warrior")
        bare(TalentsFile.serializer(), "cdn.talents.warrior")
    }

    @Test
    fun `a request body encodes in the API's snake_case`() {
        val json = ApiJson.encodeToString(AppExchangeRequest.serializer(), AppExchangeRequest("c", "v", "Phone"))

        assertEquals("""{"code":"c","code_verifier":"v","device_name":"Phone"}""", json)
    }
}
