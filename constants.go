package ceffu

// Api Base Consts

const CeffuVersionPath = "/open-api/v1/"
const CeffuVersion2Path = "/open-api/v2/"
const CeffuApiBaseUrl = "https://open-api.ceffu.com"

// Error Codes: see https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/doc-338174

const (
	ErrorBadRequest              = "G20002" // Server cannot or will not process the request due to something that is perceived to be a client error
	ErrorExceededPaginationSize  = "G20003" // The request params exceeded the pagination size
	ErrorExceededPaginationLimit = "G20004" // The request params exceeded the pagination limit
	ErrorInvalidParameterValue   = "G20005" // Invalid input params value
	ErrorTimeStampEmpty          = "G20006" // Request params timestamp is missing
	ErrorTimeStampExpired        = "G20007" // Request params timestamp is expired
	ErrorMissingKey              = "G20008" // Missisng API Key or Signature in the request
	ErrorInvalidSignature        = "G20009" // Invalid signature
	ErrorInvalidApiKey           = "G20010" // Invalid API Key
	ErrorInvalidIP               = "G20011" // Invalid IP
	ErrorRateLimitExceeded       = "G20012" // Rate limit exceeded
	ErrorNoWalletPermission      = "G20013" // No wallet permission
	ErrorInvalidReqIDFormat      = "G20014" // Invalid request ID format
	ErrorDuplicateReqID          = "G20015" // Duplicate request ID
	ErrorEndpointPermission      = "G20016" // No endpoint permission
	ErrorSubWithdrawNotSupported = "G20017" // Withdraw not supported for sub wallet
	ErrorSubWalletIDRequired     = "G20018" // Sub wallet ID is required
	ErrorPrimeWalletIDRequired   = "G20019" // Prime wallet ID is required
	ErrorWalletIDNotFound        = "G20020" // Wallet with the given ID not found
	ErrorWalletRelationship      = "G20021" // Invalid wallet relationship
	ErrorInvalidAmount           = "G20022" // Invalid amount
	ErrorInvalidRequestFormat    = "G20023" // Invalid request format
	ErrorWalletTypeNotSupported  = "G20024" // Wallet type not supported
	ErrorAddressNotActivated     = "G20025" // Address is not activated. Please contact the account manager to activate your deposit address.
	ErrorSearchableTimeRange     = "G20026" // The time range is too large to search. Please narrow down the time range.
	ErrorPrimeOrSubIDRequired    = "G20027" // Prime or sub wallet ID is required
	ErrorApiKeyExpired           = "G20028" // API Key is expired
	ErrorMirrorLink              = "G20029" // Invalid Mirror linkage relationship
	ErrorSubWalletIDNotSupported = "G20030" // Sub wallet ID is not supported for this endpoint
)

// Error code to message map
var ErrorMap = map[string]string{
	ErrorBadRequest:              "bad request",
	ErrorExceededPaginationSize:  "the request params exceeded the pagination size",
	ErrorExceededPaginationLimit: "the request params exceeded the pagination limit",
	ErrorInvalidParameterValue:   "invalid input params value",
	ErrorTimeStampEmpty:          "request params timestamp is missing",
	ErrorTimeStampExpired:        "request params timestamp is expired",
	ErrorMissingKey:              "missing api key or signature in the request",
	ErrorInvalidSignature:        "invalid signature",
	ErrorInvalidApiKey:           "invalid api key",
	ErrorInvalidIP:               "invalid ip",
	ErrorRateLimitExceeded:       "rate limit exceeded",
	ErrorNoWalletPermission:      "no wallet permission",
	ErrorInvalidReqIDFormat:      "invalid request id format",
	ErrorDuplicateReqID:          "duplicate request id",
	ErrorEndpointPermission:      "no endpoint permission",
	ErrorSubWithdrawNotSupported: "withdraw not supported for sub wallet",
	ErrorSubWalletIDRequired:     "sub wallet id is required",
	ErrorPrimeWalletIDRequired:   "prime wallet id is required",
	ErrorWalletIDNotFound:        "wallet with the given id not found",
	ErrorWalletRelationship:      "invalid wallet relationship",
	ErrorInvalidAmount:           "invalid amount",
	ErrorInvalidRequestFormat:    "invalid request format",
	ErrorWalletTypeNotSupported:  "wallet type not supported",
	ErrorAddressNotActivated:     "address is not activated. please contact the account manager to activate your deposit address",
	ErrorSearchableTimeRange:     "the time range is too large to search. please narrow down the time range",
	ErrorPrimeOrSubIDRequired:    "prime or sub wallet id is required",
	ErrorApiKeyExpired:           "api key is expired",
	ErrorMirrorLink:              "invalid mirror linkage relationship",
	ErrorSubWalletIDNotSupported: "sub wallet id is not supported for this endpoint",
}

const WalletTypeQualified = "10"
const WalletTypePrime = "20"

type WalletType int

const (
	WalletTypeIntQualified WalletType = 10
	WalletTypeIntPrime     WalletType = 20
)

type TransferDirection int

const (
	TransferDirectionIntDeposit  TransferDirection = 10
	TransferDirectionIntWithdraw TransferDirection = 20
)

type TransferType int

const (
	TransferTypeOnChain  TransferType = 10
	TransferTypeInternal TransferType = 20
)

type WithdrawStatus int

const (
	WithdrawStatusPending    WithdrawStatus = 10
	WithdrawStatusProcessing WithdrawStatus = 20
	WithdrawStatusSuccess    WithdrawStatus = 30
	WithdrawStatusConfirmed  WithdrawStatus = 40
	WithdrawStatusFailed     WithdrawStatus = 99
)

type SubWalletTransferType int

const (
	SubWalletNotFiltered SubWalletTransferType = 0
	SubWalletParentToSub SubWalletTransferType = 10
	SubWalletSubToParent SubWalletTransferType = 20
	SubWalletSubToSub    SubWalletTransferType = 30
)

type SubWalletTransferStatus int

const (
	SubWalletTransferStatusPending    SubWalletTransferStatus = 10
	SubWalletTransferStatusProcessing SubWalletTransferStatus = 20
	SubWalletTransferStatusSuccess    SubWalletTransferStatus = 30
	SubWalletTransferStatusFailed     SubWalletTransferStatus = 99
)

type MirrorXOrderType int

const (
	MirrorXOrderTypeAll      MirrorXOrderType = 0
	MirrorXOrderTypeDeposit  MirrorXOrderType = 10
	MirrorXOrderTypeWithdraw MirrorXOrderType = 20
)
