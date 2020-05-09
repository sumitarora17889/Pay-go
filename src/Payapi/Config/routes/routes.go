package routes

import (
	"Payapi/controllers/audit/auditchecklist"
	"Payapi/controllers/audit/auditrequest"
	"Payapi/controllers/config/configgateway"
	"Payapi/controllers/config/configgatewayaccount"
	"Payapi/controllers/config/configgatewaytooption"
	"Payapi/controllers/config/configgatewaytoplan"
	"Payapi/controllers/config/configmode"
	"Payapi/controllers/config/configoptions"
	"Payapi/controllers/config/configoptiontoplatform"
	"Payapi/controllers/config/configpaymentplans"
	"Payapi/controllers/config/configplantoplatform"
	"Payapi/controllers/config/configreconciliationtype"
	"Payapi/controllers/config/configsettlementaccount"
	"Payapi/controllers/config/configtax"
	"Payapi/controllers/config/configtransactioncharges"
	"Payapi/controllers/credit/creditapplication"
	"Payapi/controllers/credit/crediteligibility"
	"Payapi/controllers/gateway/gatewaytransaction"
	"Payapi/controllers/payment/paymentadvice"
	"Payapi/controllers/payment/paymentcreditnote"
	"Payapi/controllers/payment/paymentdisposition"
	"Payapi/controllers/payment/paymentgateway"
	"Payapi/controllers/payment/paymentinvoice"
	"Payapi/controllers/payment/paymentoptions"
	"Payapi/controllers/payment/paymentorder"
	"Payapi/controllers/payment/paymentplans"
	"Payapi/controllers/payment/paymentstatus"
	"Payapi/controllers/payment/paymenttransaction"
	"Payapi/controllers/payout/payoutconsent"
	"Payapi/controllers/payout/payoutrequest"
	"Payapi/controllers/refund/refundrequest"
	"Payapi/controllers/users/userslimit"
	"Payapi/controllers/users/usersoptions"
	"net/http"
)

/* List of all the modules along with their Router. The router will contain the list of controllers. The controller will list all the actions.*/
var Routers map[string]map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]map[string]func(http.ResponseWriter, *http.Request){
	"credit":  CreditRouter,
	"config":  ConfigRouter,
	"audit":   AuditRouter,
	"gateway": GatewayRouter,
	"payment": PaymentRouter,
	"payout":  PayoutRouter,
	"refund":  RefundRouter,
	"users":   UsersRouter,
}

var CreditRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"application": creditapplication.Actions,
	"eligibility": crediteligibility.Actions,
}

var ConfigRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"gateway":            configgateway.Actions,
	"gatewayaccount":     configgatewayaccount.Actions,
	"gatewaytooption":    configgatewaytooption.Actions,
	"gatewaytoplan":      configgatewaytoplan.Actions,
	"mode":               configmode.Actions,
	"options":            configoptions.Actions,
	"optiontoplatform":   configoptiontoplatform.Actions,
	"paymentplans":       configpaymentplans.Actions,
	"plantoplatform":     configplantoplatform.Actions,
	"reconciliationtype": configreconciliationtype.Actions,
	"settlementaccount":  configsettlementaccount.Actions,
	"tax":                configtax.Actions,
	"transactioncharges": configtransactioncharges.Actions,
}

var AuditRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"checklist": auditchecklist.Actions,
	"request":   auditrequest.Actions,
}

var GatewayRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"transaction": gatewaytransaction.Actions,
}

var PaymentRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"advice":      paymentadvice.Actions,
	"creditnote":  paymentcreditnote.Actions,
	"disposition": paymentdisposition.Actions,
	"gateway":     paymentgateway.Actions,
	"invoice":     paymentinvoice.Actions,
	"options":     paymentoptions.Actions,
	"order":       paymentorder.Actions,
	"plans":       paymentplans.Actions,
	"status":      paymentstatus.Actions,
	"transaction": paymenttransaction.Actions,
}

var PayoutRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"consent": payoutconsent.Actions,
	"request": payoutrequest.Actions,
}

var RefundRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"request": refundrequest.Actions,
}

var UsersRouter map[string]map[string]func(http.ResponseWriter, *http.Request) = map[string]map[string]func(http.ResponseWriter, *http.Request){
	"limit":          userslimit.Actions,
	"optionsdetails": usersoptions.Actions,
}