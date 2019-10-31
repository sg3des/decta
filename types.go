package decta

type Product struct {
	ID          string  `json:"id,omitempty"`
	Type        string  `json:"type,omitempty"`
	Brand       Brand   `json:"brand"`
	Title       string  `json:"title"`
	Currency    string  `json:"currency"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
	Tax         *string `json:"tax,omitempty"`
	Images      []Image `json:"images,omitempty"`
}

type Brand struct {
	Type                      string    `json:"type,omitempty"`
	ID                        string    `json:"id,omitempty"`
	BadgeCode                 string    `json:"badge_code,omitempty"`
	BrandName                 string    `json:"brand_name"`
	SupportEmail              string    `json:"support_email,omitempty"`
	SupportPhoneCode          string    `json:"support_phone_code,omitempty"`
	SupportPhone              string    `json:"support_phone,omitempty"`
	MerchantUrl               string    `json:"merchant_url,omitempty"`
	SmsSenderName             string    `json:"sms_sender_name,omitempty"`
	DefaultCurrency           string    `json:"default_currency"`
	DefaultFeePaidBy          string    `json:"default_fee_paid_by"`
	DefaultInvoiceLanguage    string    `json:"default_invoice_language"`
	DefaultSkipCapture        bool      `json:"default_skip_capture"`
	DefaultDenyOverduePayment bool      `json:"default_deny_overdue_payment"`
	UseDies                   bool      `json:"use_dies,omitempty"`
	DefaultRequestClientInfo  []string  `json:"default_request_client_info"`
	DefaultRequiredClientInfo []string  `json:"default_required_client_info,omitempty"`
	DefaultTimezone           string    `json:"default_timezone"`
	LogoText                  string    `json:"logo_text,omitempty"`
	Logo                      *Image    `json:"logo,omitempty"`
	Websites                  []Website `json:"websites,omitempty"`
	BrandedCardDeclineReason  string    `json:"branded_card_decline_reason,omitempty"`
	BrandedCardLink           string    `json:"branded_card_link,omitempty"`
	BrandedCardDesign         string    `json:"branded_card_design,omitempty"`
	BrandedCardTypes          []string  `json:"branded_card_types,omitempty"`
	BrandedCardStatus         string    `json:"branded_card_status,omitempty"`
	CardCustomDesign          *Image    `json:"card_custom_design,omitempty"`
	BrandedCardBrandName      string    `json:"branded_card_brand_name,omitempty"`
	BrandedCardCurrencies     []string  `json:"branded_card_currencies,omitempty"`
}

type Website struct {
	Type                  string   `json:"type,omitempty"`
	ID                    int      `json:"id,omitempty"`
	UID                   string   `json:"uid,omitempty"`
	Brand                 *Website `json:"brand"`
	WebPage               string   `json:"web_page"`
	Logo                  *Image   `json:"logo,omitempty"`
	SupportPhoneCode      string   `json:"support_phone_code,omitempty"`
	SupportPhone          string   `json:"support_phone,omitempty"`
	SupportEmail          string   `json:"support_email,omitempty"`
	DefaultAccount        string   `json:"default_account,omitempty"`
	DefaultLanguage       string   `json:"default_language,omitempty"`
	DefaultCurrency       string   `json:"default_currency,omitempty"`
	Timezone              string   `json:"timezone,omitempty"`
	SmsSenderName         string   `json:"sms_sender_name,omitempty"`
	RequireTermsOnPayment bool     `json:"require_terms_on_payment,omitempty"`
}

type Image struct {
	ID               string `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	ImageType        string `json:"image_type,omitempty"`
	OriginalFilename string `json:"original_filename,omitempty"`
	Width            int    `json:"width,omitempty"`
	Height           int    `json:"height,omitempty"`
	ImageURL         string `json:"image"`
}

type Client struct {
	Type           string         `json:"type,omitempty"`
	ID             string         `json:"id,omitempty"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	FirstName      string         `json:"first_name,omitempty"`
	LastName       string         `json:"last_name,omitempty"`
	BirthDate      string         `json:"birth_date,omitempty"`
	PersonalCode   string         `json:"personal_code,omitempty"`
	BrandName      string         `json:"brand_name,omitempty"`
	LegalName      string         `json:"legal_name,omitempty"`
	RegistrationNr string         `json:"registration_nr,omitempty"`
	VatPayerNr     string         `json:"vat_payer_nr,omitempty"`
	LegalAddress   string         `json:"legal_address,omitempty"`
	Country        string         `json:"country,omitempty"`
	State          string         `json:"state,omitempty"`
	City           string         `json:"city,omitempty"`
	ZipCode        string         `json:"zip_code,omitempty"`
	BankAccount    string         `json:"bank_account,omitempty"`
	BankCode       string         `json:"bank_code,omitempty"`
	Cc             []string       `json:"cc,omitempty"`
	Bcc            []string       `json:"bcc,omitempty"`
	DisplayName    string         `json:"display_name,omitempty"`
	OrderCount     int            `json:"order_count,omitempty"`
	CreatedBy      *CreatedBy     `json:"created_by,omitempty"`
	SavedCardData  *SavedCardData `json:"saved_card_data,omitempty"`
}

type CreatedBy struct {
	UserID string `json:"user_id,omitempty"`
	Title  string `json:"title,omitempty"`
}

type SavedCardData struct {
	Type        string `json:"type,omitempty"`
	Last4Digits string `json:"last_4_digits,omitempty"`
}

type Order struct {
	Type                      string               `json:"type,omitempty"`
	ID                        string               `json:"id,omitempty"`
	Products                  []OrderProduct       `json:"products"`
	Client                    *OrderClient         `json:"client"`
	RequestClientInfo         []string             `json:"request_client_info,omitempty"`
	Brand                     *Brand               `json:"brand,omitempty"`
	Currency                  string               `json:"currency,omitempty"`
	Number                    string               `json:"number,omitempty"`
	Due                       int64                `json:"due,omitempty"`
	DenyOverduePayment        bool                 `json:"deny_overdue_payment,omitempty"`
	FeePaidBy                 string               `json:"fee_paid_by,omitempty"`
	SkipCapture               bool                 `json:"skip_capture,omitempty"`
	Language                  string               `json:"language,omitempty"`
	Notes                     string               `json:"notes,omitempty"`
	IsTest                    bool                 `json:"is_test,omitempty"`
	SaveCard                  bool                 `json:"save_card,omitempty"`
	PanFirst                  bool                 `json:"pan_first,omitempty"`
	IsMoto                    bool                 `json:"is_moto,omitempty"`
	Recurring3d               bool                 `json:"recurring_3d,omitempty"`
	VerifyCard                bool                 `json:"verify_card,omitempty"`
	UseVerifiedCard           bool                 `json:"use_verified_card,omitempty"`
	SdwoMerchantId            string               `json:"sdwo_merchant_id,omitempty"`
	PaymentSystem             string               `json:"payment_system,omitempty"`
	DynamicDescriptor         string               `json:"dynamic_descriptor,omitempty"`
	IsPayable                 bool                 `json:"is_payable,omitempty"`
	TerminalProcessingId      string               `json:"terminal_processing_id,omitempty"`
	Terminal                  string               `json:"terminal,omitempty"`
	SuccessRedirect           string               `json:"success_redirect,omitempty"`
	FailureRedirect           string               `json:"failure_redirect,omitempty"`
	CancelRedirect            string               `json:"cancel_redirect,omitempty"`
	CustomInvoiceUrl          string               `json:"custom_invoice_url,omitempty"`
	Link                      string               `json:"link,omitempty"`
	Shortlink                 string               `json:"shortlink,omitempty"`
	DownloadLink              string               `json:"download_link,omitempty"`
	PrintLink                 string               `json:"print_link,omitempty"`
	FullPageCheckout          string               `json:"full_page_checkout,omitempty"`
	IframeCheckout            string               `json:"iframe_checkout,omitempty"`
	DirectPost                string               `json:"direct_post,omitempty"`
	IframeCheckoutSendInvoice bool                 `json:"iframe_checkout_send_invoice,omitempty"`
	Subtotal                  float64              `json:"subtotal,omitempty"`
	TotalTax                  float64              `json:"total_tax,omitempty"`
	TotalDiscount             float64              `json:"total_discount,omitempty"`
	Total                     float64              `json:"total,omitempty"`
	SubtotalOverride          float64              `json:"subtotal_override,omitempty"`
	TotalTaxOverride          float64              `json:"total_tax_override,omitempty"`
	TotalDiscountOverride     float64              `json:"total_discount_override,omitempty"`
	TotalOverride             float64              `json:"total_override,omitempty"`
	RefundAmount              float64              `json:"refund_amount,omitempty"`
	FeeAmount                 float64              `json:"fee_amount,omitempty"`
	FeeDiscount               float64              `json:"fee_discount,omitempty"`
	WithdrawalParts           []WithdrawalParts    `json:"withdrawal_parts,omitempty"`
	CreatedBy                 *CreatedBy           `json:"created_by,omitempty"`
	Issuer                    *OrderIssuer         `json:"issuer,omitempty"`
	Status                    string               `json:"status,omitempty"`
	StatusChanges             []OrderStatusChanges `json:"status_changes,omitempty"`
	TransactionDetails        *TransactionDetails  `json:"transaction_details,omitempty"`
	Issued                    int                  `json:"issued,omitempty"`
	Modified                  int                  `json:"modified,omitempty"`
	Viewed                    int                  `json:"viewed,omitempty"`
	Captured                  int                  `json:"captured,omitempty"`
	Paid                      int                  `json:"paid,omitempty"`
	IssuedOverride            int                  `json:"issued_override,omitempty"`
	ClientDisplayName         string               `json:"client_display_name,omitempty"`
	Timezone                  string               `json:"timezone,omitempty"`
	FromApi                   bool                 `json:"from_api,omitempty"`
	FromSubscription          bool                 `json:"from_subscription,omitempty"`
	IssuedByClient            bool                 `json:"issued_by_client,omitempty"`
	Referrer                  string               `json:"referrer,omitempty"`
	ReferrerDisplayName       string               `json:"referrer_display_name,omitempty"`
	PermittedActions          []string             `json:"permitted_actions,omitempty"`
}

type OrderProduct struct {
	Type            string       `json:"type,omitempty"`
	ID              string       `json:"id,omitempty"`
	Title           string       `json:"title"`
	Price           float64      `json:"price"`
	Description     string       `json:"description,omitempty"`
	Quantity        float64      `json:"quantity,omitempty"`
	Tax             string       `json:"tax,omitempty"`
	TaxPercent      float64      `json:"tax_percent,omitempty"`
	DiscountPercent float64      `json:"discount_percent,omitempty"`
	DiscountAmount  float64      `json:"discount_amount,omitempty"`
	Image           *interface{} `json:"image,omitempty"`
	Currency        string       `json:"currency,omitempty"`
	Total           float64      `json:"total,omitempty"`
	TaxAmount       float64      `json:"tax_amount,omitempty"`
}

type OrderIssuer struct {
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	BrandName      string `json:"brand_name,omitempty"`
	LegalName      string `json:"legal_name,omitempty"`
	RegistrationNr string `json:"registration_nr,omitempty"`
	VatPayerNr     string `json:"vat_payer_nr,omitempty"`
	LegalAddress   string `json:"legal_address,omitempty"`
	BankAccount    string `json:"bank_account,omitempty"`
	BankCode       string `json:"bank_code,omitempty"`
}

type OrderStatusChanges struct {
	NewStatus *interface{} `json:"new_status,omitempty"`
	On        int          `json:"on,omitempty"`
}

const (
	OrderStatusDraft      = "draft"
	OrderStatusIssued     = "issued"
	OrderStatusViewed     = "viewed"
	OrderStatusCancelled  = "cancelled"
	OrderStatusOverdue    = "overdue"
	OrderStatusExpired    = "expired"
	OrderStatusRejected   = "rejected"
	OrderStatusHold       = "hold"
	OrderStatusPaid       = "paid"
	OrderStatusWithdrawn  = "withdrawn"
	OrderStatusChargeback = "chargeback"
	OrderStatusRefunded   = "refunded"
	OrderStatusReceived   = "received"
)

type OrderClient struct {
	Type                      string         `json:"type,omitempty"`
	ID                        string         `json:"id,omitempty"`
	Email                     string         `json:"email"`
	Phone                     string         `json:"phone"`
	FirstName                 string         `json:"first_name,omitempty"`
	LastName                  string         `json:"last_name,omitempty"`
	BirthDate                 string         `json:"birth_date,omitempty"`
	PersonalCode              string         `json:"personal_code,omitempty"`
	BrandName                 string         `json:"brand_name,omitempty"`
	LegalName                 string         `json:"legal_name,omitempty"`
	RegistrationNr            string         `json:"registration_nr,omitempty"`
	VatPayerNr                string         `json:"vat_payer_nr,omitempty"`
	LegalAddress              string         `json:"legal_address,omitempty"`
	Country                   string         `json:"country,omitempty"`
	State                     string         `json:"state,omitempty"`
	City                      string         `json:"city,omitempty"`
	ZipCode                   string         `json:"zip_code,omitempty"`
	BankAccount               string         `json:"bank_account,omitempty"`
	BankCode                  string         `json:"bank_code,omitempty"`
	Cc                        []string       `json:"cc,omitempty"`
	Bcc                       []string       `json:"bcc,omitempty"`
	SavedCardData             *SavedCardData `json:"saved_card_data,omitempty"`
	RecurringPaused           bool           `json:"recurring_paused,omitempty"`
	SendToEmail               bool           `json:"send_to_email,omitempty"`
	SendToPhone               bool           `json:"send_to_phone,omitempty"`
	LastOrderId               string         `json:"last_order_id,omitempty"`
	LastOrderStatus           *interface{}   `json:"last_order_status,omitempty"`
	LastOrderPermittedActions *interface{}   `json:"last_order_permitted_actions,omitempty"`
	LastOrderIssued           int            `json:"last_order_issued,omitempty"`
	LastOrderLink             *interface{}   `json:"last_order_link,omitempty"`
	LastOrderShortlink        *interface{}   `json:"last_order_shortlink,omitempty"`
	LastOrderDownloadLink     *interface{}   `json:"last_order_download_link,omitempty"`
	LastOrderPrintLink        *interface{}   `json:"last_order_print_link,omitempty"`
	NextOrderIssued           string         `json:"next_order_issued,omitempty"`
	OriginalClient            *Client        `json:"original_client,omitempty"`
	OriginalClientID          string         `json:"original_client,omitempty"`
	IsAdded                   bool           `json:"is_added,omitempty"`
}

type WithdrawalParts struct {
	Type           string  `json:"type,omitempty"`
	DelayRemaining int     `json:"delay_remaining,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	TransferNumber string  `json:"transfer_number,omitempty"`
}

type TransactionDetails struct {
	Errors                []TransactionError  `json:"errors,omitempty"`
	CardMaskedPan         string              `json:"card_masked_pan,omitempty"`
	Method                string              `json:"method,omitempty"`
	CardIssuerCountry     *CardIssuerCountry  `json:"card_issuer_country,omitempty"`
	CardholderName        string              `json:"cardholder_name,omitempty"`
	Arn                   string              `json:"arn,omitempty"`
	Cleared               int                 `json:"cleared,omitempty"`
	Settled               int                 `json:"settled,omitempty"`
	PaymentSystem         string              `json:"payment_system,omitempty"`
	ProcessingPaymentId   string              `json:"processing_payment_id,omitempty"`
	Enable3dSecure        bool                `json:"enable_3d_secure,omitempty"`
	Mcc                   string              `json:"mcc,omitempty"`
	ProcessingMerchantId  string              `json:"processing_merchant_id,omitempty"`
	Descriptor            string              `json:"descriptor,omitempty"`
	PaidFromIp            string              `json:"paid_from_ip,omitempty"`
	CreatedFromIp         string              `json:"created_from_ip,omitempty"`
	ProcessingStatus      *ProcessingStatus   `json:"processing_status,omitempty"`
	ThreeDSecureStatus    *ThreeDSecureStatus `json:"three_d_secure_status,omitempty"`
	FraudCheckStatus      *FraudCheckStatus   `json:"fraud_check_status,omitempty"`
	AmountRefunded        float64             `json:"amount_refunded,omitempty"`
	AmountReversed        float64             `json:"amount_reversed,omitempty"`
	Chargebacks           *interface{}        `json:"chargebacks,omitempty"`
	ChargebacksReversed   *interface{}        `json:"chargebacks_reversed,omitempty"`
	ChargebacksReasonCode float64             `json:"chargebacks_reason_code,omitempty"`
	FeeDetails            *FeeDetails         `json:"fee_details,omitempty"`
	RollingAmount         float64             `json:"rolling_amount,omitempty"`
}

type TransactionError struct {
	On          int64  `json:"on"`
	Description string `json:"description"`
}

type CardIssuerCountry struct {
	Type        string `json:"type,omitempty"`
	ISOcode     int    `json:"iso_code,omitempty"`
	Code        string `json:"code,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type ProcessingStatus struct {
	Description string `json:"description,omitempty"`
	On          string `json:"on,omitempty"`
	Code        string `json:"code,omitempty"`
}

type ThreeDSecureStatus struct {
	Description   string `json:"description,omitempty"`
	On            string `json:"on,omitempty"`
	MPIStatusCode string `json:"mpi_status_code,omitempty"`
}

type FraudCheckStatus struct {
	Description string `json:"description,omitempty"`
	Code        string `json:"code,omitempty"`
}

type FeeDetails struct {
	FeeSell    float64 `json:"fee_sell,omitempty"`
	FeeFixSell float64 `json:"fee_fix_sell,omitempty"`
	MinFeeFix  float64 `json:"min_fee_fix,omitempty"`
	MaxFeeFix  float64 `json:"max_fee_fix,omitempty"`
}
