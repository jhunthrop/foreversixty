package gg.foreversixty.app.home.flows

import org.junit.Assert.assertEquals
import org.junit.Test

/** Round 4 row 8: the hero line breaks at the dot at large text, never stranding `· 2`. */
class HomeHeroTest {
    private val facts = listOf("Level 58 Fury Warrior", "2 levels to the 60 list")

    @Test fun `at the default size the facts share one line on the dot`() {
        assertEquals(listOf("Level 58 Fury Warrior · 2 levels to the 60 list"), heroFactLines(facts, stacked = false))
    }

    @Test fun `at large text each fact takes its own line`() {
        assertEquals(facts, heroFactLines(facts, stacked = true))
    }
}
