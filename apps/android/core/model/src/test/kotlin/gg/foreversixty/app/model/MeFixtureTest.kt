package gg.foreversixty.app.model

import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.jsonObject
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

/** The models decode lane C's contract fixtures, copied onto the classpath by the build. */
class MeFixtureTest {
    private fun fixture(name: String): String =
        requireNotNull(javaClass.classLoader?.getResource("fixtures/app/$name")) {
            "fixture $name missing from docs/contracts/fixtures/app"
        }.readText()

    /** Lane C's API fixtures are whole envelopes; the models read their data. */
    private fun dataOf(raw: String): JsonElement =
        ApiJson.parseToJsonElement(raw).jsonObject.getValue("data")

    @Test
    fun `me fixture decodes into Me`() {
        val me = ApiJson.decodeFromJsonElement(Me.serializer(), dataOf(fixture("me.json")))

        assertTrue(me.user.id > 0)
        assertTrue(me.characters.isNotEmpty())
        assertNotNull(me.mainCharacter)
    }

    @Test
    fun `main character falls back to the first when none is chosen`() {
        val me = ApiJson.decodeFromJsonElement(Me.serializer(), dataOf(fixture("me.json")))
        val unchosen = me.copy(mainCharacterKey = null)

        assertEquals(me.characters.first(), unchosen.mainCharacter)
    }

    @Test
    fun `a character with only the required fields decodes with nulls`() {
        val json = """{"key":"k","region":"eu","ruleset":"forever","name":"Alt","source":"bnet"}"""
        val character = ApiJson.decodeFromString(AccountCharacter.serializer(), json)

        assertNull(character.level)
        assertNull(character.build)
        assertEquals("Alt", character.name)
    }

    @Test
    fun `a null slice from Go decodes as an empty list`() {
        val json = """{"id":1,"region":"eu","ruleset":"forever","name":"G","verified":true,"verified_count":1,"can":null}"""
        val guild = ApiJson.decodeFromString(AccountGuild.serializer(), json)

        assertEquals(emptyList<String>(), guild.can)
    }

    @Test
    fun `unknown fields added by the API are ignored`() {
        val json = """{"id":1,"role":"user","anonymize":false,"brand_new_field":{"x":1}}"""
        val user = ApiJson.decodeFromString(AppUser.serializer(), json)

        assertEquals(1L, user.id)
    }

    @Test
    fun `every me success fixture decodes`() {
        for (name in listOf("me.json", "me.premium.json", "me.empty.json", "me.patch.json")) {
            val me = ApiJson.decodeFromJsonElement(Me.serializer(), dataOf(fixture(name)))

            assertTrue(name, me.user.id > 0)
        }
    }

    @Test
    fun `the premium fixture carries billing and the empty one no characters`() {
        val premium = ApiJson.decodeFromJsonElement(Me.serializer(), dataOf(fixture("me.premium.json")))
        val empty = ApiJson.decodeFromJsonElement(Me.serializer(), dataOf(fixture("me.empty.json")))

        assertNotNull(premium.entitlements.billing)
        assertTrue(empty.characters.isEmpty())
        assertNull(empty.mainCharacter)
    }

    @Test
    fun `the sign-in exchange fixture carries the token, device and me`() {
        val answer = ApiJson.decodeFromJsonElement(AppSession.serializer(), dataOf(fixture("app.signin.exchange.json")))

        assertTrue(answer.token.startsWith("fsd_"))
        assertTrue(answer.device.id.isNotEmpty())
        assertTrue(answer.me.user.id > 0)
    }

    @Test
    fun `a rated character card decodes its score and gear panel`() {
        val card = ApiJson.decodeFromJsonElement(CharacterProfile.serializer(), dataOf(fixture("character.card.json")))

        assertTrue(card.character.key.isNotEmpty())
        assertNotNull(card.score)
        assertNotNull(card.components)
        assertTrue(card.gear.all { it.slot.isNotEmpty() })
    }

    @Test
    fun `a character with no rated fight has every score field null`() {
        val card = ApiJson.decodeFromJsonElement(CharacterProfile.serializer(), dataOf(fixture("character.card.empty.json")))

        assertNull(card.score)
        assertNull(card.tier)
        assertNull(card.components)
        assertNull(card.gearGapTotal)
    }
}
