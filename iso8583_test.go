package iso8583

import (
	"fmt"
	"testing"
)

func TestISOParse(t *testing.T) {
	// MTI = 0200
	// Field (3) = 000010
	// Field (4) = 1500
	// Field (7) = 1206041200
	// Field (11) = 000001
	// Field (41) = 12340001
	// Field (49) = 840
	isomsg := "02003220000000808000000010000000001500120604120000000112340001840"
	isostruct, err := NewISOStruct("spec1987.yml", true)
	if err != nil {
		fmt.Println(err)
		t.Errorf("initialize iso struct failed")
	}

	parsed, err := isostruct.Parse(isomsg)
	if err != nil {
		fmt.Println(err)
		t.Errorf("parse iso message failed")
	}

	isomsgUnpacked, err := parsed.ToString()
	if err != nil {
		fmt.Println(err)
		t.Errorf("failed to unpack valid isomsg")
	}
	if isomsgUnpacked != isomsg {
		t.Errorf("%s should be %s", isomsgUnpacked, isomsg)
	}
	// fmt.Printf("%#v, %#v\n%#v", parsed.Mti, parsed.Bitmap, parsed.Elements)
}

func TestEmpty(t *testing.T) {
	one, err := NewISOStruct("spec1987.yml", false)
	if err != nil {
		fmt.Println(err)
		t.Errorf("initialize iso struct failed")
	}

	if one.Mti.String() != "" {
		t.Errorf("Empty generates invalid MTI")
	}
	one.AddMTI("0200")
	one.AddField(3, "000010")
	one.AddField(4, "000000001500")
	one.AddField(7, "1206041200")
	one.AddField(11, "000001")
	one.AddField(41, "12340001")
	one.AddField(49, "840")

	expected := "02003220000000808000000010000000001500120604120000000112340001840"
	unpacked, _ := one.ToString()
	if unpacked != expected {
		t.Errorf("Manually constructed isostruct produced %s not %s", unpacked, expected)
	}
}

func BenchmarkISOParse(b *testing.B) {
	// MTI = 0200
	// Field (3) = 000010
	// Field (4) = 1500
	// Field (7) = 1206041200
	// Field (11) = 000001
	// Field (41) = 12340001
	// Field (49) = 840
	isomsg := "02003220000000808000000010000000001500120604120000000112340001840"
	isostruct, err := NewISOStruct("spec1987.yml", true)
	if err != nil {
		fmt.Println(err)
		b.Errorf("initialize iso struct failed")
	}

	for i := 0; i < b.N; i++ {
		_, err := isostruct.Parse(isomsg)
		if err != nil {
			fmt.Println(err)
			b.Errorf("parse iso message failed")
		}
	}

}

func BenchmarkISOPack(b *testing.B) {
	// MTI = 0200
	// Field (3) = 000010
	// Field (4) = 1500
	// Field (7) = 1206041200
	// Field (11) = 000001
	// Field (41) = 12340001
	// Field (49) = 840

	isostruct, err := NewISOStruct("spec1987.yml", true)
	if err != nil {
		fmt.Println(err)
		b.Errorf("initialize iso struct failed")
	}

	for i := 0; i < b.N; i++ {
		isostruct.AddField(3, "000010")
		isostruct.AddField(4, "1500")
		isostruct.AddField(7, "1206041200")
		isostruct.AddField(11, "000001")
		isostruct.AddField(41, "12340001")
		isostruct.AddField(49, "840")

		_, err := isostruct.ToString()
		if err != nil {
			fmt.Println(err)
			b.Errorf("pack iso message failed")
		}
	}

}
