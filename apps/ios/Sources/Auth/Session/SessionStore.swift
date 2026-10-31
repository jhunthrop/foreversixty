import Foundation
import Models
import Networking
import Observation

/// What an account action came to. `needsBrowser` is `403 browser_session_required`: the act
/// is done on foreversixty.gg, so the view opens the site and never retries.
public enum AccountActionOutcome: Equatable, Sendable {
    case done
    case failed
    case needsBrowser

    init(_ error: APIError) {
        self = error.needsBrowserSession ? .needsBrowser : .failed
    }
}

public extension AccountDeletionRequest {
    /// The body `DELETE /v1/me` takes after the person confirmed (api/openapi.yaml).
    static let confirmed = AccountDeletionRequest(confirm: "delete my account")
}

/// The account for the whole app: the phase (signed out, signing in, importing, signed in), the
/// current character, and the account acts of spec §5.1 (sign in, sign out, revoke, delete).
/// Any 401 on any request ends the session through `SessionTokenSource.discarded`.
@MainActor
@Observable
public final class SessionStore {
    public private(set) var phase: SessionPhase
    public private(set) var pickedKey: String?

    @ObservationIgnored private let client: APIClient
    @ObservationIgnored private let tokens: SessionTokenSource
    @ObservationIgnored private let signInFlow: AppSignIn
    @ObservationIgnored private let picked: any StoredValue
    @ObservationIgnored private let deviceID: any StoredValue
    @ObservationIgnored private let clearCaches: @MainActor () -> Void
    @ObservationIgnored private var endingOnPurpose = false

    public init(
        client: APIClient,
        tokens: SessionTokenSource,
        web: any WebAuthenticating,
        picked: any StoredValue,
        deviceID: any StoredValue,
        clearCaches: @escaping @MainActor () -> Void = {}
    ) {
        self.client = client
        self.tokens = tokens
        self.signInFlow = AppSignIn(client: client, web: web, store: tokens)
        self.picked = picked
        self.deviceID = deviceID
        self.clearCaches = clearCaches
        self.pickedKey = picked.load()
        self.phase = tokens.bearerToken() == nil ? .signedOut(nil) : .signedIn(nil)
    }

    /// This install's device id, from the sign-in exchange.
    public var thisDeviceID: String? { deviceID.load() }

    /// The character Home, Raids and BiS follow (spec §3.4).
    public var currentCharacter: AccountCharacter? {
        phase.me.flatMap { CharacterSelection.current(in: $0, picked: pickedKey) }
    }

    /// Ends the session whenever a request is answered 401. Run once from the app root.
    public func watchForDeadTokens() async {
        for await _ in tokens.discarded {
            tokenDied()
        }
    }

    /// Reads `GET /v1/me`. A 401 ends the session; any other failure keeps what is on screen.
    public func refresh() async {
        guard phase.isSignedIn else { return }
        do throws(APIError) {
            let me = try await client.send(.getMe)
            if phase.isSignedIn { phase = .signedIn(me) }
        } catch {
            if case .unauthorized = error { tokenDied() }
        }
    }

    /// A request learned the token is dead: signed out, with the short line, without asking.
    public func tokenDied() {
        guard !endingOnPurpose, phase.isSignedIn else { return }
        end(with: .sessionEnded)
    }

    /// F1.2 to F1.4: the system browser, the exchange, then the import.
    public func signIn(deviceName: String) async {
        phase = .signingIn
        do throws(AuthError) {
            let session = try await signInFlow.signIn(deviceName: deviceName)
            deviceID.save(session.device.id)
            // F1.3 ends when GET /v1/me answers; the exchange's own copy stands in if it fails.
            phase = .importing
            let me = (try? await client.send(.getMe)) ?? session.me
            phase = .signedIn(me)
        } catch {
            phase = .signedOut(SignInFailure(error).map(SignedOutNotice.signInFailed))
        }
    }

    /// The tap on F1.1 clears its line (spec F1.4: held until the next tap).
    public func dismissNotice() {
        if case .signedOut(.some) = phase { phase = .signedOut(nil) }
    }

    public func pick(_ key: String) {
        picked.save(key)
        pickedKey = key
    }

    /// You › Sign out: revoke this device on the server, then forget everything local.
    public func signOut() async {
        endingOnPurpose = true
        defer { endingOnPurpose = false }
        try? await signInFlow.signOut()
        end(with: nil)
    }

    public func devices(now: Date = .now) async throws(APIError) -> [DeviceRow] {
        let devices = try await client.send(.listDevices)
        return DeviceList.rows(devices, thisDevice: thisDeviceID, now: now)
    }

    public func revoke(_ id: String) async -> AccountActionOutcome {
        do throws(APIError) {
            _ = try await client.send(.revokeDevice(id: id))
            return .done
        } catch {
            return AccountActionOutcome(error)
        }
    }

    /// `DELETE /v1/me` after the confirmation. On 202 the server has signed out every device,
    /// so the app forgets its token without calling logout and shows the date the data goes.
    public func deleteAccount() async -> AccountActionOutcome {
        do throws(APIError) {
            let deletion = try await client.send(.deleteAccount(.confirmed))
            endingOnPurpose = true
            defer { endingOnPurpose = false }
            try? tokens.clear()
            end(with: .accountDeleted(completeBy: deletion.completeBy))
            return .done
        } catch {
            return AccountActionOutcome(error)
        }
    }

    private func end(with notice: SignedOutNotice?) {
        try? tokens.clear()
        deviceID.save(nil)
        picked.save(nil)
        pickedKey = nil
        clearCaches()
        // TODO(lane A2, F7): unregister this device's push token before the token goes.
        phase = .signedOut(notice)
    }
}
