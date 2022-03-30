package scapi

const (
	APC_FAILED                                          = -9
	APDU_ERROR                                          = -1
	IRIS_API_ERROR                                      = -4
	PB_API_ERROR                                        = -6
	PB_E_FINALIZE_FAILED                                = 54
	PB_E_INTERNAL_MAX                                   = 21
	PB_E_INTERNAL_MIN                                   = 1
	PB_E_STORE_BIO_HEADER_FAILED                        = 52
	PB_E_STORE_REF_DATA_FAILED                          = 53
	PB_VALIDATE_SCORE_INSUFFICIENT                      = 51
	SAS_FAILED                                          = -5
	SAS_STATUS_ALGO_NOT_AVAILABLE                       = 4100
	SAS_STATUS_AUTHEN_FAILED                            = 4112
	SAS_STATUS_INVALID_MESSAGE                          = 4098
	SAS_STATUS_KEY_NOT_AVAILABLE                        = 4106
	SAS_STATUS_REQUEST_FAILED                           = 4097
	SAS_STATUS_VERIFY_FAILED                            = 4110
	SCAPI_BIOCARDAPI_FAILED                             = -3
	SCAPI_BRC100T_FAILED                                = -4
	SCAPI_FAILED                                        = -1
	SCAPI_SCARD_FAILED                                  = -2
	SCAPI_STATUS_APPLICATION_INVALIDATED                = 1102
	SCAPI_STATUS_APPLICATION_NOT_EXIST                  = 1101
	SCAPI_STATUS_AUTHENTICATION_FAILED                  = 1011
	SCAPI_STATUS_CARD_LOCKED                            = 1501
	SCAPI_STATUS_CID_NOT_FOUND                          = 3011
	SCAPI_STATUS_COMMUNICATION_ERROR                    = 3002
	SCAPI_STATUS_CONDITION_NOT_SATISFIED                = 1024
	SCAPI_STATUS_ENCODING_ERROR                         = 3012
	SCAPI_STATUS_ENCODING_UNKNOWN_ERROR                 = 3099
	SCAPI_STATUS_FUNCTION_NOT_SUPPORT                   = 1201
	SCAPI_STATUS_IDCARD_INTERNAL_ERROR                  = 1998
	SCAPI_STATUS_IDCARD_UNKNOWN_ERROR                   = 1999
	SCAPI_STATUS_INCORRECT_PIN                          = 1001
	SCAPI_STATUS_INPUT_INCORRECT                        = 1202
	SCAPI_STATUS_KEY_CURRENTLY_BLOCKED                  = 1012
	SCAPI_STATUS_NEW_PIN_NOT_MATCH                      = 1007
	SCAPI_STATUS_NO_CARD_PRESENT                        = 2002
	SCAPI_STATUS_NO_LICENSE_MANAGER                     = 9001
	SCAPI_STATUS_NO_PERMISSION                          = 1022
	SCAPI_STATUS_NO_PERMIT_FROM_CARD_HOLDER             = 1003
	SCAPI_STATUS_NOT_FP_AUTHORIZE                       = 1004
	SCAPI_STATUS_NOT_PIN_AUTHORIZE                      = 1005
	SCAPI_STATUS_PIN_CURRENTLY_BLOCKED                  = 1002
	SCAPI_STATUS_PINBOX_OBJ_ERROR                       = 1006
	SCAPI_STATUS_READER_NOT_OPEN_YET                    = 2003
	SCAPI_STATUS_REFERENCE_DATA_INVALID                 = 1211
	SCAPI_STATUS_REFERENCE_DATA_NOT_FOUND               = 1212
	SCAPI_STATUS_SAME_OR_UNKNOWN_OR_INAPPROPIATE_STATUS = 1103
	SCAPI_STATUS_SECURITY_STATUS_NOT_SATISFIED          = 1023
	SCAPI_STATUS_SYSTEM_CANCELLED                       = 3001
	SCAPI_STATUS_UNKNOWN_CARD_TYPE                      = 2004
	SCAPI_STATUS_UNKNOWN_READER                         = 2001
	SCAPI_STATUS_WRONG_OPTION                           = 1203
	SCAPI_SUCCESS                                       = 0
	SUCCESS                                             = 0
	SW_ALGO_NOT_AVAILABLE                               = 28435
	SW_APP_NOT_AVAILABLE                                = 28434
	SW_CARD_NOT_SUPPORTED                               = 28448
	SW_GEN_PKI_KEY_FAILED                               = 28428
	SW_INSTALL_NOT_AVAILABLE                            = 28433
	SW_KEY_INCORRECT                                    = 28427
	SW_KEY_NOT_AVAILABLE                                = 28426
	SW_LOADFILE_NOT_AVAILABLE                           = 28432
	SW_NOT_AVAILABLE                                    = 28424
	SW_NOT_ENABLED                                      = 28569
	SW_NOT_MATCH                                        = 28431
	SW_NOT_SUPPORTED                                    = 28425
	SW_PIN_NOT_AVAILABLE                                = 28442
	SW_RANDOM_NOT_AVAILABLE                             = 28423
	SW_SEED_LOCKED                                      = 28421
	SW_SEED_NOT_LOCKED                                  = 28422
	SW_SERVICE_NOT_AUTHENTICATED                        = 27380
	SW_SERVICE_NOT_AVAILABLE                            = 27377
	SW_SERVICE_NOT_ENABLED                              = 27379
	SW_SERVICE_NOT_SUPPORTED                            = 27378
	SW_SIGN_FAILED                                      = 28429
	SW_VERIFY_FAILED                                    = 28430
	SW_WRONG_ALGORITHM                                  = 28418
	SW_WRONG_RANDOM                                     = 28419
	SW_WRONG_SAM_TYPE                                   = 28420
)

func GetScapiStatus(status int32) (output string) {
	switch status {
	case APC_FAILED:
		output = "APC FAILED"
		break
	case PB_API_ERROR:
		output = "PB API ERROR"
		break
	case PB_E_FINALIZE_FAILED:
		output = "PB E FINALIZE FAILED"
		break
	case PB_E_INTERNAL_MAX:
		output = "PB E INTERNAL MAX"
		break
	case PB_E_INTERNAL_MIN:
		output = "PB E INTERNAL MIN"
		break
	case PB_E_STORE_BIO_HEADER_FAILED:
		output = "PB E STORE BIO HEADER FAILED"
		break
	case PB_E_STORE_REF_DATA_FAILED:
		output = "PB E STORE REF DATA FAILED"
		break
	case PB_VALIDATE_SCORE_INSUFFICIENT:
		output = "PB VALIDATE SCORE INSUFFICIENT"
		break
	case SAS_FAILED:
		output = "SAS FAILED"
		break
	case SAS_STATUS_ALGO_NOT_AVAILABLE:
		output = "SAS STATUS ALGO NOT AVAILABLE"
		break
	case SAS_STATUS_AUTHEN_FAILED:
		output = "SAS STATUS AUTHEN FAILED"
		break
	case SAS_STATUS_INVALID_MESSAGE:
		output = "SAS STATUS INVALID MESSAGE"
		break
	case SAS_STATUS_KEY_NOT_AVAILABLE:
		output = "SAS STATUS KEY NOT AVAILABLE"
		break
	case SAS_STATUS_REQUEST_FAILED:
		output = "SAS STATUS REQUEST FAILED"
		break
	case SAS_STATUS_VERIFY_FAILED:
		output = "SAS STATUS VERIFY FAILED"
		break
	case SCAPI_BIOCARDAPI_FAILED:
		output = "SCAPI BIOCARDAPI FAILED"
		break
	case SCAPI_BRC100T_FAILED:
		output = "SCAPI BRC100T FAILED"
		break
	case SCAPI_FAILED:
		output = "SCAPI FAILED"
		break
	case SCAPI_SCARD_FAILED:
		output = "SCAPI SCARD FAILED"
		break
	case SCAPI_STATUS_APPLICATION_INVALIDATED:
		output = "SCAPI STATUS APPLICATION INVALIDATED"
		break
	case SCAPI_STATUS_APPLICATION_NOT_EXIST:
		output = "SCAPI STATUS APPLICATION NOT EXIST"
		break
	case SCAPI_STATUS_AUTHENTICATION_FAILED:
		output = "SCAPI STATUS AUTHENTICATION FAILED"
		break
	case SCAPI_STATUS_CARD_LOCKED:
		output = "SCAPI STATUS CARD LOCKED"
		break
	case SCAPI_STATUS_CID_NOT_FOUND:
		output = "SCAPI STATUS CID NOT FOUND"
		break
	case SCAPI_STATUS_COMMUNICATION_ERROR:
		output = "SCAPI STATUS COMMUNICATION ERROR"
		break
	case SCAPI_STATUS_CONDITION_NOT_SATISFIED:
		output = "SCAPI STATUS CONDITION NOT SATISFIED"
		break
	case SCAPI_STATUS_ENCODING_ERROR:
		output = "SCAPI STATUS ENCODING ERROR"
		break
	case SCAPI_STATUS_ENCODING_UNKNOWN_ERROR:
		output = "SCAPI STATUS ENCODING UNKNOWN ERROR"
		break
	case SCAPI_STATUS_FUNCTION_NOT_SUPPORT:
		output = "SCAPI STATUS FUNCTION NOT SUPPORT"
		break
	case SCAPI_STATUS_IDCARD_INTERNAL_ERROR:
		output = "SCAPI STATUS IDCARD INTERNAL ERROR"
		break
	case SCAPI_STATUS_IDCARD_UNKNOWN_ERROR:
		output = "SCAPI STATUS IDCARD UNKNOWN ERROR"
		break
	case SCAPI_STATUS_INCORRECT_PIN:
		output = "PIN INCORRECT"
		break
	case SCAPI_STATUS_INPUT_INCORRECT:
		output = "INPUT INCORRECT"
		break
	case SCAPI_STATUS_KEY_CURRENTLY_BLOCKED:
		output = "SCAPI STATUS KEY CURRENTLY BLOCKED"
		break
	case SCAPI_STATUS_NEW_PIN_NOT_MATCH:
		output = "SCAPI STATUS NEW PIN NOT MATCH"
		break
	case SCAPI_STATUS_NO_CARD_PRESENT:
		output = "SCAPI STATUS NO CARD PRESENT"
		break
	case SCAPI_STATUS_NO_LICENSE_MANAGER:
		output = "SCAPI STATUS NO LICENSE MANAGER"
		break
	case SCAPI_STATUS_NO_PERMISSION:
		output = "SCAPI STATUS NO PERMISSION"
		break
	case SCAPI_STATUS_NO_PERMIT_FROM_CARD_HOLDER:
		output = "SCAPI STATUS NO PERMIT FROM CARD HOLDER"
		break
	case SCAPI_STATUS_NOT_FP_AUTHORIZE:
		output = "SCAPI STATUS NOT FP AUTHORIZE"
		break
	case SCAPI_STATUS_NOT_PIN_AUTHORIZE:
		output = "SCAPI STATUS NOT PIN AUTHORIZE"
		break
	case SCAPI_STATUS_PIN_CURRENTLY_BLOCKED:
		output = "SCAPI STATUS PIN CURRENTLY BLOCKED"
		break
	case SCAPI_STATUS_PINBOX_OBJ_ERROR:
		output = "SCAPI STATUS PINBOX OBJ ERROR"
		break
	case SCAPI_STATUS_READER_NOT_OPEN_YET:
		output = "SCAPI STATUS READER NOT OPEN YET"
		break
	case SCAPI_STATUS_REFERENCE_DATA_INVALID:
		output = "SCAPI STATUS REFERENCE DATA INVALID"
		break
	case SCAPI_STATUS_REFERENCE_DATA_NOT_FOUND:
		output = "SCAPI STATUS REFERENCE DATA NOT FOUND"
		break
	case SCAPI_STATUS_SAME_OR_UNKNOWN_OR_INAPPROPIATE_STATUS:
		output = "SCAPI STATUS SAME OR UNKNOWN OR INAPPROPIATE STATUS"
		break
	case SCAPI_STATUS_SECURITY_STATUS_NOT_SATISFIED:
		output = "SCAPI STATUS SECURITY STATUS NOT SATISFIED"
		break
	case SCAPI_STATUS_SYSTEM_CANCELLED:
		output = "SCAPI STATUS SYSTEM CANCELLED"
		break
	case SCAPI_STATUS_UNKNOWN_CARD_TYPE:
		output = "SCAPI STATUS UNKNOWN CARD TYPE"
		break
	case SCAPI_STATUS_UNKNOWN_READER:
		output = "SCAPI STATUS UNKNOWN READER"
		break
	case SCAPI_STATUS_WRONG_OPTION:
		output = "SCAPI STATUS WRONG OPTION"
		break
	case SCAPI_SUCCESS:
		output = "SCAPI SUCCESS"
		break
	case SW_ALGO_NOT_AVAILABLE:
		output = "SW ALGO NOT AVAILABLE"
		break
	case SW_APP_NOT_AVAILABLE:
		output = "SW APP NOT AVAILABLE"
		break
	case SW_CARD_NOT_SUPPORTED:
		output = "SW CARD NOT SUPPORTED"
		break
	case SW_GEN_PKI_KEY_FAILED:
		output = "SW GEN PKI KEY FAILED"
		break
	case SW_INSTALL_NOT_AVAILABLE:
		output = "SW INSTALL NOT AVAILABLE"
		break
	case SW_KEY_INCORRECT:
		output = "SW KEY INCORRECT"
		break
	case SW_KEY_NOT_AVAILABLE:
		output = "SW KEY NOT AVAILABLE"
		break
	case SW_LOADFILE_NOT_AVAILABLE:
		output = "SW LOADFILE NOT AVAILABLE"
		break
	case SW_NOT_AVAILABLE:
		output = "SW NOT AVAILABLE"
		break
	case SW_NOT_ENABLED:
		output = "SW NOT ENABLED"
		break
	case SW_NOT_MATCH:
		output = "SW NOT MATCH"
		break
	case SW_NOT_SUPPORTED:
		output = "SW NOT SUPPORTED"
		break
	case SW_PIN_NOT_AVAILABLE:
		output = "SW PIN NOT AVAILABLE"
		break
	case SW_RANDOM_NOT_AVAILABLE:
		output = "SW RANDOM NOT AVAILABLE"
		break
	case SW_SEED_LOCKED:
		output = "SW SEED LOCKED"
		break
	case SW_SEED_NOT_LOCKED:
		output = "SW SEED NOT LOCKED"
		break
	case SW_SERVICE_NOT_AUTHENTICATED:
		output = "SW SERVICE NOT AUTHENTICATED"
		break
	case SW_SERVICE_NOT_AVAILABLE:
		output = "SW SERVICE NOT AVAILABLE"
		break
	case SW_SERVICE_NOT_ENABLED:
		output = "SW SERVICE NOT ENABLED"
		break
	case SW_SERVICE_NOT_SUPPORTED:
		output = "SW SERVICE NOT SUPPORTED"
		break
	case SW_SIGN_FAILED:
		output = "SW SIGN FAILED"
		break
	case SW_VERIFY_FAILED:
		output = "SW VERIFY FAILED"
		break
	case SW_WRONG_ALGORITHM:
		output = "SW WRONG ALGORITHM"
		break
	case SW_WRONG_RANDOM:
		output = "SW WRONG RANDOM"
		break
	case SW_WRONG_SAM_TYPE:
		output = "SW WRONG SAM TYPE"
		break
	default:
		output = "Unknow Status"
	}
	return
}
