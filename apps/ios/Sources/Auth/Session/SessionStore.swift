import Foundation
import Models
import Networking
import Observation

/// What an account action came to. `needsBrowser` is `403 browser_session_required`: the act
/// is done on foreversixty.gg, so the view opens the site and never retries. `confirmMismatch`
/// is `400 confirm_mismatch` (the typed confirmation was refused); `refused` carries the API's
/// own sentence for a 409, 429 or 503, which the view shows verbatim.
public enum AccountActionOutcome: Equatable, Sendable {
    case done
    case failed
    case needsBrowser
    case confirmMismatch
    case refused(message: String)

    /// The statuses whose message names the reason and the fix (api/openapi.yaml, deleteAccount).
    static let verbatimStatuses: Set<Int> = [409, 429, 503]
    static let confirmMismatchCode = "confirm_mismatch"

    init(_ error: APIError) {
        if error.needsBrowserSession {
            self = .needsBrowser
        } else if error.code == Self.confirmMismatchCode {
            self = .confirmMismatch
        } else if case let .server(status, body, _) = error, Self.verbatimStatuses.contains(status) {
            self = .refused(message: body.message)
        } else {
            self = .failed
        }
    }
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
    /// Reloads the account; false when the call failed (the drawer's Refresh says so).
    @discardableResult
    public func refresh() async -> Bool {
        guard phase.isSignedIn else { return false }
        do throws(APIError) {
            let me = try await client.send(.getMe)
            if phase.isSignedIn { phase = .signedIn(me) }
            return true
        } catch {
            if case .unauthorized = error { tokenDied() }
            return false
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

    /// Marks `key` the account's main (`PATCH /v1/me`, flows spec F8) and makes it the current
    /// character. Throws when the API refused or could not be reached; nothing changes then.
    public func setMain(_ key: String) async throws(APIError) {
        let me = try await client.send(.updateMe(MePatch(mainCharacterKey: key)))
        if phase.isSignedIn { phase = .signedIn(me) }
        pick(key)
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

    /// `DELETE /v1/me` with what the person typed (`DeletionConfirmation`). On 200 the account is
    /// gone and every device signed out, so the app forgets its token without calling logout.
    public func deleteAccount(confirm typed: String) async -> AccountActionOutcome {
        do throws(APIError) {
            _ = try await client.send(.deleteAccount(AccountDeletionRequest(confirm: typed)))
            endingOnPurpose = true
            defer { endingOnPurpose = false }
            try? tokens.clear()
            end(with: .accountDeleted)
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
