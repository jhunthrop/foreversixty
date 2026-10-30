package gg.foreversixty.app.model

import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonObject
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

    /** A fixture is either the whole envelope or the data alone; both are accepted. */
    private fun dataOf(raw: String): JsonElement {
        val root = ApiJson.parseToJsonElement(raw).jsonObject
        return if ("ok" in root) root.getValue("data") else JsonObject(root)
    }

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
        val character = ApiJson.decodeFromString(Character.serializer(), json)

        assertNull(character.level)
        assertNull(character.build)
        assertEquals("Alt", character.name)
    }

    @Test
    fun `a null slice from Go decodes as an empty list`() {
        val json = """{"id":1,"region":"eu","ruleset":"forever","name":"G","verified":true,"verified_count":1,"can":null}"""
        val guild = ApiJson.decodeFromString(Guild.serializer(), json)

        assertEquals(emptyList<String>(), guild.can)
    }

    @Test
    fun `unknown fields added by the API are ignored`() {
        val json = """{"id":1,"role":"user","brand_new_field":{"x":1}}"""
        val user = ApiJson.decodeFromString(User.serializer(), json)

        assertEquals(1L, user.id)
    }

    @Test
    fun `every app fixture whose name starts with me decodes`() {
        val dir = requireNotNull(javaClass.classLoader?.getResource("fixtures/app")) { "fixtures dir missing" }
        val names = java.io.File(dir.toURI()).listFiles().orEmpty()
            .map { it.name }
            .filter { it.startsWith("me") && it.endsWith(".json") }

        assertTrue("no me*.json fixture found", names.isNotEmpty())
        names.forEach { ApiJson.decodeFromJsonElement(Me.serializer(), dataOf(fixture(it))) }
    }
}
