package scapi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	SAS_INT_AUTH_RMACKEY_ADMIN = 0
	SAS_INT_AUTH_FPKEY_ADMIN   = 1
	SAS_INT_AUTH_FPKEY_MOI     = 2
	SAS_INT_AUTH_FPKEY_HTH     = 3
	SAS_INT_AUTH_FPKEY_MOI2    = 4
	SAS_INT_AUTH_FPKEY_MOI3    = 5
	SAS_INT_AUTH_FPKEY_MOI4    = 6
	SAS_INT_AUTH_FPKEY_WVOT    = 7
	SAS_INT_AUTH_FPKEY_MOD     = 8
	SAS_INT_AUTH_FPKEY_MOA     = 9
	SAS_INT_AUTH_FPKEY_NHSO    = 10
)

const (
	CARD_UNKNOWN      = 20
	CARD_JCOP_20      = 21
	CARD_GXP_PROR3    = 22
	CARD_ST_DP66_A    = 23
	CARD_ST_DP66_B    = 24
	CARD_JCOP_31      = 26
	CARD_JCOP_31_27   = 27
	CARD_KONA26CC     = 28
	CARD_JCOP241R3    = 29
	CARD_JTOP20IDv2   = 30
	CARD_JTOP20IDv2L2 = 31
	CARD_SAM_SRV      = 99
)

// StringToCharPtr converts a Go string into pointer to a null-terminated cstring.
// This assumes the go string is already ANSI encoded.
func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

// StringToUTF16Ptr converts a Go string into a pointer to a null-terminated UTF-16 wide string.
// This assumes str is of a UTF-8 compatible encoding so that it can be re-encoded as UTF-16.
func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}

var (
	scapiDLL           = syscall.NewLazyDLL("scapi_ope.dll")
	procListReader2    = scapiDLL.NewProc("ListReader")
	procOpenReader     = scapiDLL.NewProc("OpenReader")
	procCloseReader    = scapiDLL.NewProc("CloseReader")
	procGetCardStatus  = scapiDLL.NewProc("GetCardStatus")
	procSelectApplet   = scapiDLL.NewProc("SelectApplet")
	procReadData       = scapiDLL.NewProc("ReadData")
	procGetCardInfo    = scapiDLL.NewProc("GetCardInfo")
	procGetInfoADM     = scapiDLL.NewProc("GetInfoADM")
	procVerifyPIN      = scapiDLL.NewProc("VerifyPIN")
	procGetMatchStatus = scapiDLL.NewProc("GetMatchStatus")
	procEnvelopeGMSx   = scapiDLL.NewProc("EnvelopeGMSx")
)

func ListReader() (readerList []string) {
	var status int32
	readerName := make([]byte, 200)
	ptr := unsafe.Pointer(&readerName[0])

	// Call Native DLL
	r, _, errno := procListReader2.Call(uintptr(ptr), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if r != 0 {
		return
	}
	readerList = splitReader(string(readerName))
	return
}

func OpenReader(readerName string) (err error) {
	var status int32
	b := append([]byte(readerName), 0)
	r, _, errno := procOpenReader.Call(uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&status)))
	_ = errno

	if r != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
	} else {
		err = nil
	}

	return
}

func CloseReader() (err error) {
	r, _, errno := procCloseReader.Call()
	_ = errno

	if r != 0 {
		e := fmt.Sprintf("[CloseReader error]")
		err = errors.New(e)
	} else {
		err = nil
	}

	return
}

func GetCardStatus() (err error) {
	var status int32
	atr := make([]byte, 1000)
	var atr_len int = 100
	var timeOut int = 200
	var cardType int = 0
	r, _, errno := procGetCardStatus.Call(uintptr(unsafe.Pointer(&atr[0])), uintptr(unsafe.Pointer(&atr_len)), uintptr(timeOut),
		uintptr(unsafe.Pointer(&cardType)), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if r != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
	} else {
		err = nil
	}
	return
}

func GetCardInfo() (cid, chip, os, pre_perso, perso string, err error) {
	var status int32
	b_cid := make([]byte, 20)
	b_chip := make([]byte, 20)
	b_os := make([]byte, 20)
	b_pre_perso := make([]byte, 20)
	b_perso := make([]byte, 20)

	r, _, errno := procGetCardInfo.Call(
		uintptr(unsafe.Pointer(&b_cid[0])), uintptr(unsafe.Pointer(&b_chip[0])), uintptr(unsafe.Pointer(&b_os[0])),
		uintptr(unsafe.Pointer(&b_pre_perso[0])), uintptr(unsafe.Pointer(&b_perso[0])), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if r != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
	} else {
		err = nil
	}
	cid = decodeWindows874(string(b_cid))
	chip = decodeWindows874(string(b_chip))
	os = decodeWindows874(string(b_os))
	pre_perso = decodeWindows874(string(b_pre_perso))
	perso = decodeWindows874(string(b_perso))
	return
}

func GetInfoADM() (adm_version, laser_number string, err error) {
	var status int32
	b_version := make([]byte, 5)
	var adm_status int32 = 0
	var authorize int32 = 0
	b_laser_number := make([]byte, 33)

	r, _, errno := procGetInfoADM.Call(uintptr(unsafe.Pointer(&b_version[0])), uintptr(unsafe.Pointer(&adm_status)),
		uintptr(unsafe.Pointer(&authorize)), uintptr(unsafe.Pointer(&b_laser_number[0])), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if r != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
		return
	} else {
		err = nil
	}

	adm_version = decodeWindows874(string(b_version))
	laser_number = decodeWindows874(string(b_laser_number))

	adm_version = strings.Trim(adm_version, "\000")
	laser_number = strings.Trim(laser_number, "\000")
	laser_number = Convert_LaserNumber(laser_number)

	return
}

func GetLaserNumber() (laser_number string, err error) {
	err = OpenReader("")
	if err != nil {
		return
	}
	err = SelectApplet("ADM_AID")
	if err != nil {
		return
	}

	adm_version, laser_number, err := GetInfoADM()
	if err != nil {
		return
	}
	_ = adm_version
	return
}

func GetCID() (cid string, err error) {
	err = OpenReader("")
	if err != nil {
		return
	}
	err = SelectApplet("ADM_AID")
	if err != nil {
		return
	}
	cid, _, _, _, _, err = GetCardInfo()
	if err != nil {
		cid = ""
		return
	}
	cid = strings.Trim(cid, "\u0000")
	return
}

func Convert_LaserNumber(laser_number string) (output string) {
	// "4d453131323039313334313900000000" -> "ME1120913419"
	b := str2byte(laser_number)
	laser_number = decodeWindows874(string(b))
	laser_number = strings.Trim(laser_number, "\000")
	output = laser_number

	return output
}

func SelectApplet(AppletID string) (err error) {
	err = nil

	var MOI_AID string = "A000000054480001"
	var CM_AID string = "434D"
	var ADM_AID string = "A000000084060002"

	if AppletID == "MOI_AID" {
		AppletID = MOI_AID
	} else if AppletID == "CM_AID" {
		AppletID = CM_AID
	} else if AppletID == "ADM_AID" {
		AppletID = ADM_AID
	}

	var status int32
	aid := str2byte(AppletID)
	var aid_size int = len(aid)
	r, _, errno := procSelectApplet.Call(uintptr(unsafe.Pointer(&aid[0])), uintptr(aid_size), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if r != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
	} else {
		err = nil
	}

	return
}

func ReadData(block_id int, offset int, data_size int) (output []byte, err error) {
	var status int32 = 0
	if data_size >= 0 {
		output = make([]byte, data_size)
	} else {
		output = nil
		err = nil
		return nil, nil
	}

	ret, _, errno := procReadData.Call(uintptr(block_id), uintptr(offset), uintptr(data_size), uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if ret != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
	} else {
		err = nil
	}
	return
}

func VerifyPIN() (tryremain int, err error) {
	var status int32 = 0
	var pin_id int = 1
	var share_data int = 0
	var try_remain int = 0
	ret, _, errno := procVerifyPIN.Call(uintptr(pin_id), uintptr(share_data), uintptr(unsafe.Pointer(&try_remain)), uintptr(unsafe.Pointer(&status)))
	_ = errno
	if ret != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
	} else {
		err = nil
		tryremain = try_remain
	}
	return
}

func GetMatchStatus(in_buf []byte) (out_buf []byte, match_stt int, err error) {
	var status int32 = 0
	var req_type int = 1
	var req_mode int = 0
	var in_size int = 0
	var out_size int = 0

	str := decodeWindows874(string(in_buf))
	// s_in := bin2Str(str)
	s_in := fmt.Sprintf("%X", str)
	in_buf = []byte(s_in)

	in_size = len(in_buf)
	out_buf = make([]byte, in_size)
	ret, _, errno := procGetMatchStatus.Call(uintptr(req_type), uintptr(req_mode), uintptr(unsafe.Pointer(&in_buf[0])), uintptr(in_size),
		uintptr(unsafe.Pointer(&out_buf[0])), uintptr(unsafe.Pointer(&out_size)), uintptr(unsafe.Pointer(&match_stt)), uintptr(unsafe.Pointer(&status)))

	_ = errno
	if ret != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
		return
	} else {
		err = nil
	}

	return
}

func EnvelopeGMSx(in_buf []byte) (out_buf []byte, err error) {
	var status int32 = 0
	envelop_size := 255
	envelop := make([]byte, envelop_size)
	in_size := len(in_buf)

	ret, _, errno := procEnvelopeGMSx.Call(uintptr(SAS_INT_AUTH_FPKEY_ADMIN), uintptr(unsafe.Pointer(&in_buf[0])), uintptr(in_size),
		uintptr(unsafe.Pointer(&envelop[0])), uintptr(unsafe.Pointer(&envelop_size)), uintptr(unsafe.Pointer(&status)))

	_ = errno
	if ret != 0 {
		e := GetScapiStatus(status)
		err = errors.New(e)
		return
	} else {
		err = nil
		out_buf = envelop
	}
	return
}

func ReadCardVersion() (string, error) {
	data, err := ReadData(0, 0, 4)

	if err != nil {
		return "", err
	}
	return string(data), nil
}

func int2HexChar(d byte) byte {
	// input 9 -> '9', 10 -> 'A'
	if d >= 0 && d <= 9 {
		return (byte)('0' + d)
	} else if d >= 10 && d <= 15 {
		return (byte)('A' + d - 10)
	} else {
		return (byte)('X')
	}
}

func hexChar2Int(c byte) byte {
	// input 'A' -> 10, '1' -> 1
	if '0' <= c && c <= '9' {
		return (byte)(0 + c - '0')
	} else if c >= 'A' && c <= 'F' {
		return (byte)(10 + c - 'A')
	} else if c >= 'a' && c <= 'f' {
		return (byte)(10 + c - 'a')
	} else {
		return 0
	}
}

func bin2Str(bin string) string {
	// input: "5A" -> [0x35][0x41] -> "3541" -> output: [0x33][0x35][0x34][0x31] {double bytes}
	var output string = ""
	length := len(bin)
	for i := 0; i < length; i++ {
		temp := int(bin[i])
		high := byte(temp / 16)
		low := byte(temp % 16)

		output += string(int2HexChar(high))
		output += string(int2HexChar(low))
	}
	return output
}

func str2byte(str string) []byte {
	// input: "A000000054480001"
	// output: [0xA0][0x00][0x00][0x00][0x54][0x48][0x00][0x01] {reduce bytes}
	b_str := []byte(str)
	output_size := int((len(b_str) + 1) / 2)
	output := make([]byte, output_size)

	i := 0
	j := 0
	var tmp int
	for i = 0; i < len(b_str); i++ {
		tmp = int(hexChar2Int(b_str[i]))
		tmp = tmp * 16
		i++
		if str[i] == 0 {
			break
		}
		tmp = tmp + int(hexChar2Int(b_str[i]))
		output[j] = (byte)(tmp)
		j++
	}
	return output
}

func splitReader(buffer string) (reader_list []string) {
	reader_list = make([]string, 0)

	start_index := -1
	for i := 0; i < len(buffer); {
		if 0x30 <= buffer[i] && buffer[i] <= 0x39 {
			if start_index == -1 {
				start_index = i
			}
			i++
		} else if buffer[i] == 0 {
			break
		} else {
			if start_index == -1 {
				break
			}
			len, err := strconv.Atoi(buffer[start_index:i])
			if err != nil {
				return
			}
			reader := buffer[i : i+len]
			reader_list = append(reader_list, reader)

			start_index = -1
			i = i + len
		}
	}

	return reader_list
}

// ------------------------------------------------------------------------------
// Shortcut function
// ------------------------------------------------------------------------------

func ReadCardData() (smartcard *SmartCard, err error) {
	readerList := ListReader()
	if len(readerList) == 0 {
		e := fmt.Sprintf("[ListReader error: %v]", "no reader")
		err = errors.New(e)
		return
	}

	for i := 0; i < len(readerList); i++ {
		fmt.Printf("reader %d: %v\n", i, readerList[i])
	}

	// OpenReader
	if len(readerList) > 0 {
		err = OpenReader(readerList[0])
		if err != nil {
			e := fmt.Sprintf("[OpenReader error: %v]", err.Error())
			err = errors.New(e)
			return
		}
	}

	// GetCardStatus
	err = GetCardStatus()
	if err != nil {
		e := fmt.Sprintf("[GetCardStatus error: %v]", err.Error())
		err = errors.New(e)
		return
	}

	// SelectApplet
	err = SelectApplet("MOI_AID")
	if err != nil {
		e := fmt.Sprintf("[SelectApplet error: %v]", err.Error())
		err = errors.New(e)
		return
	}

	cardversion, err := ReadCardVersion()
	if err != nil {
		e := fmt.Sprintf("[ReadCardVersion error: %v]", err.Error())
		err = errors.New(e)
		return
	}

	// Read card data
	if cardversion == "0002" {
		data, err2_1 := ReadData(1, 0, 377)
		if err2_1 != nil {
			e := fmt.Sprintf("[ReadData err2_1: %v]", err.Error())
			err = errors.New(e)
			return
		}
		address, err2_2 := ReadData(0, 4, 150)
		if err2_2 != nil {
			e := fmt.Sprintf("[ReadData err2_2: %v]", err.Error())
			err = errors.New(e)
			return
		}
		image, err2_3 := ReadData(1, 377, 5120)
		if err2_3 != nil {
			e := fmt.Sprintf("[ReadData err2_3: %v]", err.Error())
			err = errors.New(e)
			return
		}

		smartcard = CreateSmartCard(cardversion, data, address, image)
	}else { // cardversion == "0003"+
		data, err3_1 := ReadData(0, 0, 377)
		if err3_1 != nil {
			e := fmt.Sprintf("[ReadData err3_1: %v]", err.Error())
			err = errors.New(e)
			return
		}
		address, err3_2 := ReadData(0, 5497, 160)
		if err3_2 != nil {
			e := fmt.Sprintf("[ReadData err3_2: %v]", err.Error())
			err = errors.New(e)
			return
		}
		image, err3_3 := ReadData(0, 377, 5120)
		if err3_3 != nil {
			e := fmt.Sprintf("[ReadData err3_3: %v]", err.Error())
			err = errors.New(e)
			return
		}

		smartcard = CreateSmartCard(cardversion, data, address, image)
	}

	return
}

func GetEnvelopeByVerifyPin(random string) (envelop string, err error) {
	err = OpenReader("")
	if err != nil {
		return
	}
	// SelectApplet
	err = SelectApplet("ADM_AID")
	if err != nil {
		fmt.Printf("SelectApplet error: %#v\n", err.Error())
		return
	}

	adm_version, laser_number, err := GetInfoADM()
	if err != nil {
		fmt.Printf("GetInfoADM error: %#v\n", err.Error())
		return
	}
	fmt.Printf("adm_version: %v, laser_number: %v\r\n", adm_version, laser_number)

	// VerifyPIN
	try_remain, err := VerifyPIN()
	if err != nil {
		fmt.Printf("VerifyPIN error: %v\n", err.Error())
		return
	}
	fmt.Printf("try_remain: %v\n", try_remain)

	// GetMatchStatus
	in_buf := []byte(random)
	out_buf, match_stt, err := GetMatchStatus(in_buf)
	if err != nil {
		fmt.Printf("GetMatchStatus error: %v\n", err.Error())
		return
	}
	if match_stt != 1 {
		e := fmt.Sprintf("[GetMatchStatus: status=%v, miss match]", match_stt)
		err = errors.New(e)
		return
	}
	// fmt.Printf("out_buf: %X\r\n", out_buf)
	// fmt.Printf("out_buf: %v\r\n", string(out_buf))
	fmt.Printf("match_stt: %v\r\n", match_stt)

	// EnvelopeGMSx
	b_envelop, err := EnvelopeGMSx(out_buf)
	if err != nil {
		fmt.Printf("EnvelopeGMSx error: %v\n", err.Error())
		return
	}

	envelop = string(b_envelop)
	return
}
