package gg.foreversixty.app.auth

import gg.foreversixty.app.auth.session.DeletionConfirmation
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Test

/** The server's rule for the typed confirmation (iOS DeletionConfirmationTests mirrors it). */
class DeletionConfirmationTest {
    private val tag = "Yell#1604"

    @Test
    fun `the BattleTag matches trimmed and in any case`() {
        assertTrue(DeletionConfirmation.matches("Yell#1604", tag))
        assertTrue(DeletionConfirmation.matches("  yell#1604\n", tag))
        assertTrue(DeletionConfirmation.matches("YELL#1604", tag))
    }

    @Test
    fun `the word delete matches with or without a BattleTag`() {
        assertTrue(DeletionConfirmation.matches("delete", tag))
        assertTrue(DeletionConfirmation.matches(" DELETE ", null))
    }

    @Test
    fun `anything else does not match`() {
        assertFalse(DeletionConfirmation.matches("", tag))
        assertFalse(DeletionConfirmation.matches("   ", null))
        assertFalse(DeletionConfirmation.matches("Yell", tag))
        assertFalse(DeletionConfirmation.matches("Yell#16045", tag))
        assertFalse(DeletionConfirmation.matches("delete my account", tag))
        assertFalse(DeletionConfirmation.matches("Yell#1604", null))
    }

    @Test
    fun `the field asks for the BattleTag, or delete when there is none`() {
        assertEquals(tag, DeletionConfirmation.expected(tag))
        assertEquals(DeletionConfirmation.WORD, DeletionConfirmation.expected(null))
        assertEquals(DeletionConfirmation.WORD, DeletionConfirmation.expected(" "))
    }
}
