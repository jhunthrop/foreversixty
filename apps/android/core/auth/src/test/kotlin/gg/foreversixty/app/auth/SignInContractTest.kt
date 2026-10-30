package gg.foreversixty.app.auth

import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test

class SignInContractTest {
    private val redirect = SignInContract.REDIRECT_URI

    @Test
    fun `start url carries the challenge, method, state and platform`() {
        val url = SignInContract.startUrl("https://api.example.test/", "chal", "st8")

        assertTrue(url.startsWith("https://api.example.test/v1/auth/app/start?"))
        assertTrue(url.contains("code_challenge=chal"))
        assertTrue(url.contains("code_challenge_method=S256"))
        assertTrue(url.contains("state=st8"))
        assertTrue(url.contains("platform=android"))
    }

    @Test
    fun `start url sends no redirect_uri, which the server ignores`() {
        val url = SignInContract.startUrl("https://api.example.test", "chal", "st8")

        assertTrue(!url.contains("redirect_uri"))
    }

    @Test
    fun `the redirect is exactly foreversixty gg app auth`() {
        assertEquals("https://foreversixty.gg/app/auth", SignInContract.REDIRECT_URI)
    }

    @Test
    fun `a redirect with the issued state yields its code`() {
        assertEquals(RedirectResult.Code("abc"), SignInContract.parseRedirect("$redirect?code=abc&state=s1", "s1"))
    }

    @Test
    fun `a redirect with another state is a mismatch`() {
        assertEquals(RedirectResult.StateMismatch, SignInContract.parseRedirect("$redirect?code=abc&state=evil", "s1"))
    }

    @Test
    fun `an error redirect is denied with its reason`() {
        assertEquals(RedirectResult.Denied("access_denied"), SignInContract.parseRedirect("$redirect?error=access_denied", "s1"))
    }

    @Test
    fun `a redirect to a lookalike path is malformed`() {
        for (path in listOf("/app/authx", "/app/auth/callback")) {
            val result = SignInContract.parseRedirect("https://foreversixty.gg$path?code=a&state=s1", "s1")

            assertEquals(path, RedirectResult.Malformed, result)
        }
    }

    @Test
    fun `a redirect to another host is malformed`() {
        val result = SignInContract.parseRedirect("https://evil.example/app/auth?code=a&state=s1", "s1")

        assertEquals(RedirectResult.Malformed, result)
    }

    @Test
    fun `a redirect without a code is malformed`() {
        assertEquals(RedirectResult.Malformed, SignInContract.parseRedirect("$redirect?state=s1", "s1"))
    }
}
