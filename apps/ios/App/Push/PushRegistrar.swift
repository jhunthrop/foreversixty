import Foundation
import Models
import Networking
import OSLog
import UIKit
import UserNotifications

/// The API's push-token route (`POST /v1/me/push-tokens`), behind a protocol for tests.
protocol PushTokenRegistering: Sendable {
    func register(token: String) async throws
}

struct PushTokenCalls: PushTokenRegistering {
    let client: APIClient
    let appVersion: String

    func register(token: String) async throws {
        let request = PushTokenRequest(platform: ClientIdentity.platform, token: token, appVersion: appVersion)
        _ = try await client.send(.registerPushToken(request))
    }
}

/// Keeps this install's APNs device token on the account (flows spec 5.1). It never asks for
/// permission: the ask comes at the moment of value (E6), and once it is granted, or on any launch
/// where it already was, the app registers with APNs and the token goes to the API. Signing out
/// revokes the device, and the API drops its push tokens with it (migration 0062).
@MainActor
final class PushRegistrar {
    static let shared = PushRegistrar()

    /// The live route; set once at launch (RootView). Tests build their own registrar.
    var calls: (any PushTokenRegistering)?
    private let isAuthorized: @Sendable () async -> Bool
    private let registerWithAPNs: @MainActor () -> Void
    private var signedIn = false
    private var token: String?
    /// The token the API holds for this signed-in session, so a launch posts it once.
    private var sent: String?

    init(
        isAuthorized: @escaping @Sendable () async -> Bool = PushRegistrar.systemAuthorized,
        registerWithAPNs: @escaping @MainActor () -> Void = { UIApplication.shared.registerForRemoteNotifications() }
    ) {
        self.isAuthorized = isAuthorized
        self.registerWithAPNs = registerWithAPNs
    }

    /// The account changed (or the app launched): registers when signed in and allowed.
    func sync(signedIn: Bool) async {
        if signedIn != self.signedIn { sent = nil }
        self.signedIn = signedIn
        await registerIfAllowed()
    }

    /// The permission was just asked; a yes registers at once.
    func permissionAnswered() async {
        await registerIfAllowed()
    }

    /// APNs gave the app its device token (AppDelegate).
    func received(deviceToken: Data) async {
        token = deviceToken.map { String(format: "%02x", $0) }.joined()
        await send()
    }

    private func registerIfAllowed() async {
        guard signedIn, await isAuthorized() else { return }
        registerWithAPNs()
        await send()
    }

    private func send() async {
        guard signedIn, let token, token != sent, let calls else { return }
        do {
            try await calls.register(token: token)
            sent = token
        } catch {
            // Not fatal: the next launch or foreground registers again.
            Self.log.notice("push token not registered: \(String(describing: error), privacy: .public)")
        }
    }

    private static let log = Logger(subsystem: "gg.foreversixty.app", category: "push")

    nonisolated static let systemAuthorized: @Sendable () async -> Bool = {
        switch await UNUserNotificationCenter.current().notificationSettings().authorizationStatus {
        case .authorized, .provisional, .ephemeral: true
        default: false
        }
    }
}
