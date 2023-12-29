package main

import (
	"fmt"
	"strconv"
	"strings"
)

// packet to binary
// get version
//	4 = literal value
// 	else = operator
// get type id

type Input struct {
	Data    string
	Packets map[int][]string
}

var input = "60552F100693298A9EF0039D24B129BA56D67282E600A4B5857002439CE580E5E5AEF67803600D2E294B2FCE8AC489BAEF37FEACB31A678548034EA0086253B183F4F6BDDE864B13CBCFBC4C10066508E3F4B4B9965300470026E92DC2960691F7F3AB32CBE834C01A9B7A933E9D241003A520DF316647002E57C1331DFCE16A249802DA009CAD2117993CD2A253B33C8BA00277180390F60E45D30062354598AA4008641A8710FCC01492FB75004850EE5210ACEF68DE2A327B12500327D848028ED0046661A209986896041802DA0098002131621842300043E3C4168B12BCB6835C00B6033F480C493003C40080029F1400B70039808AC30024C009500208064C601674804E870025003AA400BED8024900066272D7A7F56A8FB0044B272B7C0E6F2392E3460094FAA5002512957B98717004A4779DAECC7E9188AB008B93B7B86CB5E47B2B48D7CAD3328FB76B40465243C8018F49CA561C979C182723D769642200412756271FC80460A00CC0401D8211A2270803D10A1645B947B3004A4BA55801494BC330A5BB6E28CCE60BE6012CB2A4A854A13CD34880572523898C7EDE1A9FA7EED53F1F38CD418080461B00440010A845152360803F0FA38C7798413005E4FB102D004E6492649CC017F004A448A44826AB9BFAB5E0AA8053306B0CE4D324BB2149ADDA2904028600021909E0AC7F0004221FC36826200FC3C8EB10940109DED1960CCE9A1008C731CB4FD0B8BD004872BC8C3A432BC8C3A4240231CF1C78028200F41485F100001098EB1F234900505224328612AF33A97367EA00CC4585F315073004E4C2B003530004363847889E200C45985F140C010A005565FD3F06C249F9E3BC8280804B234CA3C962E1F1C64ADED77D10C3002669A0C0109FB47D9EC58BC01391873141197DCBCEA401E2CE80D0052331E95F373798F4AF9B998802D3B64C9AB6617080"
var input1 = "D2FE28"
var input2 = "38006F45291200"

func BinaryToNum(s string) int64 {
	value, _ := strconv.ParseInt(s, 2, 64)
	return value
}

func Sections(s string) []string {
	final := []string{}

	for i := 0; i < len(s); i += 5 {
		if len(s[i:]) < 5 {
			final = append(final, s[i:])
		} else {
			final = append(final, s[i:i+5])
		}
	}

	return final
}

func NewInput(s string) *Input {
	b := []string{}

	for _, p := range s {
		switch string(p) {
		case "0":
			b = append(b, "0000")
		case "1":
			b = append(b, "0001")
		case "2":
			b = append(b, "0010")
		case "3":
			b = append(b, "0011")
		case "4":
			b = append(b, "0100")
		case "5":
			b = append(b, "0101")
		case "6":
			b = append(b, "0110")
		case "7":
			b = append(b, "0111")
		case "8":
			b = append(b, "1000")
		case "9":
			b = append(b, "1001")
		case "A":
			b = append(b, "1010")
		case "B":
			b = append(b, "1011")
		case "C":
			b = append(b, "1100")
		case "D":
			b = append(b, "1101")
		case "E":
			b = append(b, "1110")
		case "F":
			b = append(b, "1111")
		}
	}

	input := &Input{
		Data: strings.Join(b, ""),
	}
	input.Sections()

	return input
}

func (input *Input) Sections() error {
	input.Packets = make(map[int][]string)
	index := 0

	for i := 0; i < len(input.Data); i += 5 {

		if string(input.Data[i:][0]) == "0" {
			index++
			continue
		}

		if len(input.Data[i:]) < 5 {
			input.Packets[index] = append(input.Packets[index], input.Data[i:])
		} else {
			input.Packets[index] = append(input.Packets[index], input.Data[i:i+5])
		}
	}

	return fmt.Errorf("EOF")
}

func (input *Input) Version() string {
	return input.Data[:3]
}

func (input *Input) TypeId() string {
	return input.Data[3:6]
}

func (input *Input) Header() string {
	return input.Data[:6]
}

func (input *Input) Payload() string {
	return input.Data[6:]
}

func (input *Input) VersionValue() int64 {
	return BinaryToNum(input.Version())
}

func (input *Input) TypeIdValue() int64 {
	return BinaryToNum(input.TypeId())
}

func main() {
	p := NewInput(input)
	fmt.Println(p)
}
