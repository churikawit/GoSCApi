package scapi

import (
	"fmt"
	"strings"
	"golang.org/x/text/encoding/charmap"
)

type SmartCard struct {
	m_card_version string
    m_data []byte
    m_address []byte
    m_image []byte
}

func CreateSmartCard(card_version string, data []byte, address []byte, image []byte) *SmartCard {
	sc := new(SmartCard)

	sc.m_card_version = card_version
	sc.m_data = data
	sc.m_address = address
	sc.m_image = image

	return sc
}

func decodeWindows874(input string) string {
	dewin874 := charmap.Windows874.NewDecoder()
	output, err := dewin874.String(input)
	if err != nil {
		fmt.Printf("decodeWindows874() error: %v\r\n", err.Error())
		return ""
	}
	return output
}

func (sc *SmartCard) GetPID() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		return string(sc.m_data[4:4+13])
	} else {
		return ""
	}
}

func (sc *SmartCard) GetFullName() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		fullname := string(sc.m_data[17:17+100])
		f := decodeWindows874(fullname)
		return strings.Trim(f, " ")
	} else {
		return ""
	}
}

func (sc *SmartCard) GetTitle() (output string) { 
	output = ""

	fullname := sc.GetFullName()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 1 {
		output = token[0]
	}
	return
}

func (sc *SmartCard) GetFirstName() (output string) {
	output = ""

	fullname := sc.GetFullName()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 2 {
		output = token[1]
	}
	return
}

func (sc *SmartCard) GetMiddleName() (output string) {
	output = ""

	fullname := sc.GetFullName()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 3 {
		output = token[2]
	}
	return
}

func (sc *SmartCard) GetLastName() (output string) {
	output = ""

	fullname := sc.GetFullName()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 4 {
		output = token[3]
	}
	return
}

func (sc *SmartCard) GetFullName_En() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		fullname := string(sc.m_data[117:117+100])
		f := decodeWindows874(fullname)
		return strings.Trim(f, " ")
	} else {
		return ""
	}
}

func (sc *SmartCard) GetTitle_En() (output string) { 
	output = ""

	fullname := sc.GetFullName_En()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 1 {
		output = token[0]
	}
	return
}

func (sc *SmartCard) GetFirstName_En() (output string) {
	output = ""

	fullname := sc.GetFullName_En()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 2 {
		output = token[1]
	}
	return
}

func (sc *SmartCard) GetMiddleName_En() (output string) {
	output = ""

	fullname := sc.GetFullName_En()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 3 {
		output = token[2]
	}
	return
}

func (sc *SmartCard) GetLastName_En() (output string) {
	output = ""

	fullname := sc.GetFullName_En()
	separator := "#"
	var token []string = strings.Split(fullname, separator)
	if len(token) >= 4 {
		output = token[3]
	}
	return
}

func (sc *SmartCard) GetBirthDate() string {	
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		birthdate := string(sc.m_data[217:217+8])
		return birthdate
	} else {
		return ""
	}
}

func (sc *SmartCard) GetGender() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		gender := string(sc.m_data[225:225+1])
		return gender
	} else {
		return ""
	}
}

// หมายเลขคำร้อง
func (sc *SmartCard) GetRequestID() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		requestid := string(sc.m_data[226:226+20])
		return requestid
	} else {
		return ""
	}
}

// BP1NO
func (sc *SmartCard) GetBP1NO() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		requestid := string(sc.m_data[226:226+11])
		return requestid
	} else {
		return ""
	}
}

// สถานที่ออกบัตร
func (sc *SmartCard) GetIssueLocation() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		issuelocation := string(sc.m_data[246:246+100])
		issuelocation = decodeWindows874(issuelocation)
		return strings.Trim(issuelocation, " ")
	} else {
		return ""
	}
}

// ผู้ออกบัตร
func (sc *SmartCard) GetIssuePersonID() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		data := string(sc.m_data[346:346+13])
		data = decodeWindows874(data)
		return strings.Trim(data, " ")
	} else {
		return ""
	}
}

func (sc *SmartCard) GetIssueDate() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		data := string(sc.m_data[359:359+8])
		data = decodeWindows874(data)
		return strings.Trim(data, " ")
	} else {
		return ""
	}
}

func (sc *SmartCard) GetExpireDate() string {
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		data := string(sc.m_data[367:367+8])
		data = decodeWindows874(data)
		return strings.Trim(data, " ")
	} else {
		return ""
	}
}

func (sc *SmartCard) GetCardType() string {
	// Output: 01 = บัตรประชาชน
	if (sc.m_card_version == "0003" || sc.m_card_version == "0002") {
		data := string(sc.m_data[375:375+2])
		data = decodeWindows874(data)
		return strings.Trim(data, " ")
	} else {
		return ""
	}
}

func (sc *SmartCard) GetAddress() string {
	if (sc.m_card_version == "0003") {
		data := string(sc.m_address[0:160])
		data = decodeWindows874(data)
		return strings.Trim(data, " ")
	}else if (sc.m_card_version == "0002") {
			data := string(sc.m_address[0:150])
			data = decodeWindows874(data)
			return strings.Trim(data, " ")
	} else {
		return ""
	}
}
