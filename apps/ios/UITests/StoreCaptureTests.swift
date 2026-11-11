import XCTest

/// The App Store screenshots' raw captures (lane BA2): the fixture app in dark appearance on Home,
/// Gear, and a raid's page and its roster. design/listing/capture-apps.sh runs it on the
/// shared simulator with the status bar set to 9:41 and full bars, then
/// design/listing/build-store.mjs draws each capture onto its store plate. The PNGs land in
/// design/listing/app-captures/ios. Not part of `make test`.
@MainActor
final class StoreCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/listing/app-captures/ios")
    /// The fixtures' Onyxia's Lair raid, the one the raid captures use.
    private static let raidLink = "https://foreversixty.gg/app/raid/2/31"
    /// How long a screen takes to draw its fixture answers and load its icons from the CDN.
    private static let settles: UInt32 = 6

    override func setUp() async throws {
        continueAfterFailure = false
        XCUIDevice.shared.appearance = .dark
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name).png"))
    }

    func testStoreScreens() throws {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        try save("home")
        app.tabBars.buttons["Gear"].tap()
        sleep(Self.settles)
        try save("gear")
        app.terminate()

        let raid = XCUIApplication()
        raid.launchArguments = ["-FSOpenLink", Self.raidLink]
        raid.launch()
        sleep(Self.settles)
        try save("raid")
        raid.swipeUp()
        sleep(Self.settles)
        try save("raid-roster")
        raid.terminate()
    }
}
