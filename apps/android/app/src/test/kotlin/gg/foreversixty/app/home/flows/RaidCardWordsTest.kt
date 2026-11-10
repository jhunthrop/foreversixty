package gg.foreversixty.app.home.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import org.junit.Test
import org.junit.runner.RunWith

/** E7 (owner, 2026-11-09): the answer names its character only with more than one character in the guild. */
@RunWith(AndroidJUnit4::class)
class RaidCardWordsTest {
    private val words = RaidCardWords(resources(), clock(), RaidTime.now)
    private val one = RaidFixture.raid()
    private val two = one.copy(viewer = one.viewer.copy(characters = one.viewer.characters + one.viewer.characters.first().copy(characterKey = "us/pvp/lightvow", name = "Lightvow")))

    private fun card(raid: RaidView, answer: SignupState?) = words.card(raid, "Olympus XXVII", "us/pvp/obnoxious-yell", answer, lateOpen = false, badge = null)

    @Test fun `one character in the guild gives the answer alone`() {
        assertEquals("You are In", card(one, SignupState.In).me)
        assertEquals("You have not answered", card(one, null).me)
    }

    @Test fun `two characters in the guild name the one answering`() {
        assertEquals("You are In, as Obnoxious Yell, Damage", card(two, SignupState.In).me)
        assertEquals("You have not answered, as Obnoxious Yell, Damage", card(two, null).me)
    }

    @Test fun `the card titles the zone and puts the guild before the time`() {
        val model = card(one, null)
        assertEquals("Onyxia's Lair", model.zone)
        assertEquals(true, model.place.startsWith("Olympus XXVII · "))
    }
}
