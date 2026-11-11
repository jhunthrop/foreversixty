package gg.foreversixty.app.home.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertEquals
import java.time.Duration
import java.time.Instant
import org.junit.Test
import org.junit.runner.RunWith

/** E7 (owner, 2026-11-09): the answer names its character only with more than one character in the guild. */
@RunWith(AndroidJUnit4::class)
class RaidCardWordsTest {
    private val words = RaidCardWords(resources(), clock(), RaidTime.now)
    private val one = RaidFixture.raid()
    private val two = one.copy(viewer = one.viewer.copy(characters = one.viewer.characters + one.viewer.characters.first().copy(characterKey = "us/pvp/lightvow", name = "Lightvow")))

    private fun card(raid: RaidView, answer: SignupState?) = words.card(raid, "Olympus XXVII", "us/pvp/gorvash", answer, lateOpen = false, badge = null)

    @Test fun `one character in the guild - no answer line once answered, the segmented control shows it`() {
        assertEquals(null, card(one, SignupState.In).me)
        assertEquals(true, card(one, SignupState.In).sub?.endsWith("late"))
        assertEquals("You have not answered", card(one, null).me)
    }

    @Test fun `two characters in the guild name the one answering`() {
        assertEquals("As Gorvash, Damage", card(two, SignupState.In).me)
        assertEquals("You have not answered, as Gorvash, Damage", card(two, null).me)
    }

    @Test fun `the card titles the zone and puts the guild before the time`() {
        val model = card(one, null)
        assertEquals("Onyxia's Lair", model.zone)
        assertEquals(true, model.place.startsWith("Olympus XXVII · "))
    }

    @Test fun `the Raid tab head names the character only with two in the guild`() {
        assertEquals(null, card(one, SignupState.In).who)
        assertEquals("As Gorvash, Damage", card(two, SignupState.In).who)
    }

    @Test fun `a Late answer with minutes carries them, and the word reads them`() {
        val late = words.card(one, "Olympus XXVII", "us/pvp/gorvash", SignupState.Late, lateOpen = true, badge = null, lateMinutes = 30)
        assertEquals(30, late.lateMinutes)
        assertEquals("Late\u00a0·\u00a030\u00a0min", words.answerWord(SignupState.Late, late.lateMinutes))
    }

    @Test fun `minutes count only for a Late answer`() {
        val inWithMinutes = words.card(one, "Olympus XXVII", "us/pvp/gorvash", SignupState.In, lateOpen = true, badge = null, lateMinutes = 30)
        assertEquals(null, inWithMinutes.lateMinutes)
    }

    @Test fun `an officer's not-answered count is the raid's own list, the foot's population (row 6)`() {
        val badge = gg.foreversixty.app.model.RaidBadge(count = 6, youUnanswered = false, officer = true, unanswered = 99, applicants = 0)
        val card = words.card(one, "Olympus XXVII", "us/pvp/gorvash", null, lateOpen = false, badge = badge)
        assertEquals(true, card.officer?.unanswered?.startsWith("${one.notSigned.size} "))
    }

    @Test fun `within a day the head line gives the pull in hours and minutes`() {
        val pull = Instant.parse("2026-12-20T01:00:00Z")
        val soon = RaidCardWords(resources(), clock(), pull.minus(Duration.ofMinutes(168)))
        val line = soon.card(one, "Olympus XXVII", "us/pvp/gorvash", null, lateOpen = false, badge = null).headLine
        assertEquals(true, line.endsWith(" · pull in 2 h 48 min"))
        val minutes = RaidCardWords(resources(), clock(), pull.minus(Duration.ofMinutes(12)))
        assertEquals(true, minutes.card(one, "Olympus XXVII", null, null, lateOpen = false, badge = null).headLine.endsWith(" · pull in 12 min"))
    }

    @Test fun `a day or more out the head line gives the date and time`() {
        val model = card(one, null)
        assertEquals(model.aside, model.headLine)
    }
}
