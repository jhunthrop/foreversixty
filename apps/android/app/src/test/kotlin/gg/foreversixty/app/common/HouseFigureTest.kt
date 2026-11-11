package gg.foreversixty.app.common

import org.junit.Assert.assertEquals
import org.junit.Test

/** The house figure (apps/shared/formats.json `numbers`): grouped and whole from 100, one decimal below. */
class HouseFigureTest {
    @Test fun `from 100 a figure is grouped and whole`() {
        assertEquals("182", HouseFigure.of(182.0))
        assertEquals("204", HouseFigure.of(203.6))
        assertEquals("1,196", HouseFigure.of(1196.2))
        assertEquals("1,235", HouseFigure.of(1234.6))
        assertEquals("1,234,567", HouseFigure.of(1234567.0))
    }

    @Test fun `a gain is signed with a plus, a loss with a true minus`() {
        assertEquals("+158", HouseFigure.signed(157.6))
        assertEquals("+2.8", HouseFigure.signed(2.8))
        assertEquals("−0.8", HouseFigure.signed(-0.8))
    }

    @Test fun `below 100 a figure keeps one decimal`() {
        assertEquals("18.0", HouseFigure.of(18.0))
        assertEquals("12.3", HouseFigure.of(12.34))
        assertEquals("3.0", HouseFigure.of(3.0))
        assertEquals("-12.0", HouseFigure.of(-12.0))
    }
}
