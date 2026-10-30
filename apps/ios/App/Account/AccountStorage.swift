import Auth
import Foundation
import Models
import Networking
import UIKit

/// Where the account keeps its two small values, and which browser signs in, per API mode.
/// Live: UserDefaults and ASWebAuthenticationSession. Fixture: memory and a browser that answers
/// at once, so the whole sign-in runs on the contract fixtures; the device id is seeded from the
/// exchange fixture, as if this install had signed in, so Devices marks `This iPhone`.
struct AccountStorage {
    let picked: any StoredValue
    let deviceID: any StoredValue
    private let makeWeb: @MainActor @Sendable () -> any WebAuthenticating

    @MainActor
    func web() -> any WebAuthenticating { makeWeb() }

    static var device: AccountStorage {
        AccountStorage(
            picked: DefaultsValue.pickedCharacter,
            deviceID: DefaultsValue.deviceID,
            makeWeb: { SystemWebAuthenticator(anchor: keyWindow) }
        )
    }

    static func fixture(bundle: Bundle) -> AccountStorage {
        AccountStorage(
            picked: MemoryValue(),
            deviceID: MemoryValue(fixtureDeviceID(bundle)),
            makeWeb: { FixtureWebAuthenticator() }
        )
    }

    private static let exchangeFixture = "app.signin.exchange"

    private static func fixtureDeviceID(_ bundle: Bundle) -> String? {
        guard let url = bundle.url(forResource: exchangeFixture, withExtension: "json", subdirectory: "Fixtures"),
              let data = try? Data(contentsOf: url)
        else { return nil }
        return (try? WireFormat.decoder().decode(Envelope<AppSession>.self, from: data))?.data?.device.id
    }

    @MainActor
    private static func keyWindow() -> UIWindow {
        let scene = UIApplication.shared.connectedScenes.compactMap { $0 as? UIWindowScene }.first
        if let window = scene?.keyWindow ?? scene?.windows.first { return window }
        guard let scene else { preconditionFailure("sign-in needs a window scene") }
        return UIWindow(windowScene: scene)
    }
}
