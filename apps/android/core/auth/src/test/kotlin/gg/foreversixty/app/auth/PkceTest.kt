package gg.foreversixty.app.auth

import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class PkceTest {
    @Test
    fun `challenge matches the RFC 7636 appendix B vector`() {
        val verifier = "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"

        assertEquals("E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM", Pkce.challenge(verifier))
    }

    @Test
    fun `a new verifier is 43 url-safe characters`() {
        val verifier = Pkce.newVerifier()

        assertEquals(43, verifier.length)
        assertTrue(verifier.all { it.isLetterOrDigit() || it == '-' || it == '_' })
    }

    @Test
    fun `verifiers and states are fresh each time`() {
        assertNotEquals(Pkce.newVerifier(), Pkce.newVerifier())
        assertNotEquals(Pkce.newState(), Pkce.newState())
    }

    @Test(expected = IllegalArgumentException::class)
    fun `a verifier shorter than the RFC minimum is refused`() {
        Pkce.challenge("too-short")
    }
}
