package ofx

type StatusCode int
type Severity string

const (
	Info  Severity = "INFO"
	Warn           = "WARN"
	Error          = "ERROR"
)

const (
	InfoOK                               StatusCode = 0
	InfoClientUpToDate                              = 1
	ErrorGeneral                                    = 2000
	ErrorInvalidAccount                             = 2001
	ErrorGeneralAccount                             = 2002
	ErrorAccountNotFound                            = 2003
	ErrorAccountClosed                              = 2004
	ErrorAccountNotAuthorized                       = 2005
	ErrorSourceAccountNotFound                      = 2006
	ErrorSourceAccountClosed                        = 2007
	ErrorSourceAccountNotAuthorized                 = 2008
	ErrorDestinationAccountNotFound                 = 2009
	ErrorDestinationAccountClosed                   = 2010
	ErrorDestinationAccountNotAuthorized            = 2011
	ErrorInvalidAmount                              = 2012
	ErrorDateTooSoon                                = 2014
	ErrorDateTooFarInTheFuture                      = 2015
	ErrorAlreadyCommitted                           = 2016
)

/*
TODO
2017
Already canceled (ERROR)
The transaction cannot be canceled or modified because it has already been canceled.
2018
Unknown server ID (ERROR)
The specified server ID does not exist or no longer exists.
2019
Duplicate request (ERROR)
A request with this <TRNUID> has already been received and processed.
2020
Invalid date (ERROR)
The specified datetime stamp cannot be parsed; for instance, the datetime stamp specifies 25:00 hours.
2021
Unsupported version (ERROR)
The server does not support the requested version.
2022
Invalid TAN (ERROR)
The server was unable to validate the TAN sent in the request.
3000
Authentication Required (ERROR)
User credentials are correct, but further authentication required.  Client should send <MFACHALLENGERQ>.
3001
Authentication error (ERROR)
<MFACHALLENGEA> contains invalid information.
10000
Stop check in process (INFO)
Stop check is already in process.
10500
Too many checks to process (ERROR)
The stop-payment request <STPCHKRQ> specifies too many checks.
10501
Invalid payee (ERROR)
Payee error not specified by the remaining error codes.
10502
Invalid payee address (ERROR)
Some portion of the payee’s address is incorrect or unknown.
10503
Invalid payee account number (ERROR)
The account number <PAYACCT> of the requested payee is invalid.
10504
Insufficient funds (ERROR)
The server cannot process the request because the specified account does not have enough funds.
10505
Cannot modify element (ERROR)
The server does not allow modifications to one or more values in a modification request.
10506
Cannot modify source account (ERROR)
Reserved for future use.
10507
Cannot modify destination account (ERROR)
Reserved for future use.
10508
Invalid frequency (ERROR)
The specified frequency <FREQ> does not match one of the accepted frequencies for recurring transactions.
10509
Model already canceled (ERROR)
The server has already canceled the specified recurring model.
10510
Invalid payee ID (ERROR)
The specified payee ID does not exist or no longer exists.
10511
Invalid payee city (ERROR)
The specified city is incorrect or unknown.
10512
Invalid payee state (ERROR)
The specified state is incorrect or unknown.
10513
Invalid payee postal code (ERROR)
The specified postal code is incorrect or unknown.
10514
Bank payment already processed (ERROR)
The server has already processed the bank payment.
10515
Payee not modifiable by client (ERROR)
The server does not allow clients to change payee information.
10516
Wire beneficiary invalid (ERROR)
The specified wire beneficiary does not exist or no longer exists.
10517
Invalid payee name (ERROR)
The server does not recognize the specified payee name.
10518
Unknown model ID (ERROR)
The specified model ID does not exist or no longer exists.
10519
Invalid payee list ID (ERROR)
The specified payee list ID does not exist or no longer exists.
12250
Investment transaction download not supported (WARN)
The server does not support investment transaction download.
12251
Investment position download not supported (WARN)
The server does not support investment position download.
12252
Investment positions for specified date not available (WARN)
The server does not support investment positions for the specified date.
12253
Investment open order download not supported (WARN)
The server does not support open order download.
12254
Investment balances download not supported (WARN)
The server does not support investment balances download.
12500
One or more securities not found (ERROR)
The server could not find the requested securities.
13000
User ID & password will be sent out-of-band (INFO)
The server will send the user ID and password via postal mail, e-mail, or another means. The accompanying message will provide details.
13500
Unable to enroll user (ERROR)
The server could not enroll the user.
13501
User already enrolled (ERROR)
The server has already enrolled the user.
13502
Invalid service (ERROR)
The server does not support the service <SVC> specified in the service-activation request.
13503
Cannot change user information (ERROR)
The server does not support the <CHGUSERINFORQ> request.
15000
Must change USERPASS (INFO)
The user must change his or her USERPASS number as part of the next OFX request.
15500
Signon (for example, user ID or password) invalid (ERROR)
The user cannot signon because he or she entered an invalid user ID or password.
15501
Customer account already in use (ERROR)
The server allows only one connection at a time, and another user is already signed on. Please try again later.
15502
USERPASS lockout (ERROR)
The server has received too many failed signon attempts for this user. Please call the FI’s technical support number.
15503
Could not change USERPASS (ERROR)
The server does not support the <PINCHRQ> request.
15504
Could not provide random data (ERROR)
The server could not generate random data as requested by the CHALLENGERQ.
15510
User must register CLIENTUID (ERROR)
CLIENTUID error.  User must register CLIENTUID.
15511
User contact required (ERROR)
User should contact financial institution
15512
OFX Server requires AUTHTOKEN in signon during the next session. (ERROR)
User needs to contact financial institution to obtain AUTHTOKEN. Client should send it in the next request.
15513
AUTHTOKEN error (ERROR)
AUTHTOKEN invalid.
16500
HTML not allowed (ERROR)
The server does not accept HTML formatting in the request.
16501
Unknown mail To: (ERROR)
The server was unable to send mail to the specified Internet address.
16502
Invalid URL (ERROR)
The server could not parse the URL.
16503
Unable to get URL (ERROR)
The server was unable to retrieve the information at this URL (e.g., an HTTP 400 or 500 series error).
*/
