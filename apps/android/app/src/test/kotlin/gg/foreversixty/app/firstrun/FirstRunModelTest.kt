package gg.foreversixty.app.firstrun

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Me
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.json.jsonObject
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

/** F8's steps, the same list as iOS `FirstRunTests`. */
class FirstRunModelTest {
    private class Asked(var asked: Boolean = false) : FirstRunAskStore {
        override fun wasAsked() = asked
        override fun markAsked() {
            asked = true
        }
    }

    private val thandrel = FirstRunFixtures.characters.first()
    private val kaelor = FirstRunFixtures.characters[2]

    private fun me(main: String? = null): Me {
        val raw = requireNotNull(javaClass.classLoader?.getResource("fixtures/app/me.json")) { "no me.json" }.readText()
        val data = requireNotNull(ApiJson.parseToJsonElement(raw).jsonObject["data"])
        // Highest level not first: the model sorts, keeping the account's order between equals.
        val characters = FirstRunFixtures.characters.drop(2) + FirstRunFixtures.characters.take(2)
        return ApiJson.decodeFromJsonElement(Me.serializer(), data).copy(characters = characters, mainCharacterKey = main)
    }

    private fun model(saves: Boolean = true, undetermined: Boolean = true, asked: Asked = Asked(), saved: MutableList<String> = mutableListOf()) =
        FirstRunModel(FirstRunModel.Dependencies({ key -> saved += key; saves }, { undetermined }, asked))

    @Test
    fun `no main asks for one with the highest level preselected`() {
        val model = model()
        model.update(me())
        assertEquals(FirstRunModel.Step.PickMain, model.step)
        assertEquals(thandrel.key, model.selection)
        assertEquals(6, FirstRunModel.candidates(me()).size)
    }

    @Test
    fun `an account with a main is left alone`() {
        val model = model()
        model.update(me(main = thandrel.key))
        assertNull(model.step)
    }

    @Test
    fun `a guild main goes on to the ask`() = runTest {
        val saved = mutableListOf<String>()
        val model = model(saved = saved)
        model.update(me())
        model.continueWithMain(me())
        assertEquals(listOf(thandrel.key), saved)
        assertEquals(FirstRunModel.Step.Permission("Olympus XXVII"), model.step)
    }

    @Test
    fun `a main with no guild, an earlier ask or a decided system skips the ask`() = runTest {
        val noGuild = model().apply { update(me()); selection = kaelor.key }
        noGuild.continueWithMain(me())
        assertNull(noGuild.step)
        val askedBefore = model(asked = Asked(asked = true)).apply { update(me()) }
        askedBefore.continueWithMain(me())
        assertNull(askedBefore.step)
        val decided = model(undetermined = false).apply { update(me()) }
        decided.continueWithMain(me())
        assertNull(decided.step)
    }

    @Test
    fun `a failed save stays on Pick your main and says so`() = runTest {
        val model = model(saves = false)
        model.update(me())
        model.continueWithMain(me())
        assertEquals(FirstRunModel.Step.PickMain, model.step)
        assertTrue(model.failed)
        assertFalse(model.saving)
    }

    @Test
    fun `either answer is the only ask`() {
        var requested = 0
        val asked = Asked()
        model(asked = asked).turnOnNotifications { requested++ }
        assertEquals(1, requested)
        assertTrue(asked.asked)
        val notNow = Asked()
        model(asked = notNow).notNow()
        assertTrue(notNow.asked)
    }
}
