## With Tag

Supported tag:

**Internet :**

* Email
* Mac address
* Domain name
* URL
* UserName
* IP Address (IPv4, IPv6 )
* Password


**Payment :**
* Credit Card Type (VISA, MASTERCARD, AMERICAN EXPRESS, DISCOVER, JCB, DINERS CLUB)
* Credit Card Number

**Address :**
* Latitude and Longitude

**Phone :**
* Phone number
* China mobile number
* Toll free phone number
* E164PhoneNumber

**Person :**
* Title male
* Title female
* FirstName
* FirstName male
* FirstName female
* LastName
* Name

**DateTime :**
* UnixTime
* Date
* Time
* MonthName
* Year
* DayOfWeek
* DayOfMonth
* Timestamp
* Century
* TimeZone
* TimePeriod

**Lorem :**
* Word
* Sentence
* Paragraph

**Price :**
* Currency
* Amount
* Amount with Currency

**UUID :**
* UUID Digit (32 bytes)
* UUID Hyphenated (36 bytes)

**Regex :**
* Regex expression

**Enum :**
* Enum expression

**Skip :**
* \-

```go

package main

import (
	"fmt"

	"github.com/bingoohuang/faker"
)

// SomeStruct ...
type SomeStruct struct {
    RegexArr		 []string  `faker:"regex=\\d{15}"`
    RegexArr2		 []string  `faker:"regex=\\d+,len=10"`
    Regex		       string  `faker:"regex=\\d{15}"`
    Regex2		       string  `faker:"regex=\\d+,len=10"`
    
    Enums1           []int     `faker:"enum=1/2/3"`
    Enums2           []string  `faker:"enum=x/y/z"`
    Enum1              int     `faker:"enum=4/5/6"`
    Enum2              string  `faker:"enum=a/b/c"`
    
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
	CreditCardNumber   string  `faker:"cc_number"`
	CreditCardType     string  `faker:"cc_type"`
	Email              string  `faker:"email"`
	DomainName		   string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	PhoneNumber        string  `faker:"phone_number"`
	ChinaMobileNumber  string  `faker:"china_mobile_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TollFreeNumber     string  `faker:"toll_free_number"`
	E164PhoneNumber    string  `faker:"e_164_phone_number"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`
	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated	   string  `faker:"uuid_hyphenated"`
	UUID	           string  `faker:"uuid_digit"`
	Skip		       string  `faker:"-"`
}

func main() {

	a := SomeStruct{}
	err := faker.Fake(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
	/*
		Output  :
		{
            Enums1: [2 1 3 2 2 3 3 3 1 2 3 1 2 1 3 1 3 3 1 2 3 3 1 2 1 3 3 1 1 3 2 2 1 1 1 1 1 2 2 2 1],
            Enums2: [x x y x z z x x z z y y x x z x z z z y y z x z x z z x z x z z y x x z x x z z y z y y x z x x x z x x x],
            Enum1: 6,
            Enum2: c,
            RegexArr2: [9864865028 36976087 59449005 527139 2640140661 70 927905 7286144762 8946 8 0186069926 6147 2 4845 2252796 0019347961 0 5171918958 8 91 1055563 3280 56 702777370 50891 668395885 65443 6119593 2 2049884 4202873662 238511 83208746 60 83058 8462 619896996 61618 66070895 101957857 5622 130836 756455828 19543 66999792 15483 704 7446656421 507784581 24 5136527603 8395901 2390570 9942861 6 609],
            RegexArr: [001128675026926 445978964025594 068132223916597 462316041047161 771085784627259 488403529033564 609793138572333 960422345322205 908809567990295 813303523077607 887867514181553 030061268872032 154217388108885 491745073149876 713629385179238 233525948771003 412213607665798 320184033609041 256686185787791 447378765235596 084299947458091 846127388688778 265998518516274 090375656128737 947411694665676 852898396037635 096436903218254 867711242978365 115048889560420 099790896463925 239320630803236 445263517293644 676500959207100 850066076120636 222813590933102 804217230469848 369255673775925 261400450852818 581646884168773 827036270450263 274686633265526 382447105281809 048653882806996 279408573779414 584787008992957 235730498507876 404046037668587 280275296013623 602337189892344 659209477489683 195501052679838 434674907857838 110716834622384 956395561936403 831895862018786 682262930660843 664343035799792 528550578609931 613959794461229 313859164392518 613267783417414 969137329394480 735805828920432 947312338904005 015756180579219 156153504986449 991705425170280 980750502854516 585550973120936 787966098770873 957201571855698 716225008717268 591528695938314 177594586339048 766958843992286 016808338626306 100634973098912 960772826518691 801160036487299 088213687680304 024976910879181 676946501492179 303456167583141 704495109932612 562296461430347 588169259242898 146182105077829 528498903218376 140047982110777 778505917647795 039153324619644 696864247595927 840006079828689 889103090654716 554955457617690 123862527151050 111923202751599 135489825271816 356851799038067],
            Regex: 086779516560007,
            Regex2: 9645228182,
			Latitude: 81.12195
			Longitude: -84.38158
			CreditCardType: American Express
			CreditCardNumber: 373641309057568
			Email: mJBJtbv@OSAaT.ru
			DomainName: FWZcaRE.ru,
			IPV4: 99.23.42.63
			IPV6: 975c:fb2c:2133:fbdd:beda:282e:1e0a:ec7d
			Password: dfJdyHGuVkHBgnHLQQgpINApynzexnRpgIKBpiIjpTPOmNyMFb
			PhoneNumber: 792-153-4861
			MacAddress: cd:65:e1:d4:76:c6
			URL: https://www.oEuqqAY.org/QgqfOhd
			UserName: lVxELHS
			TollFreeNumber: (777) 831-964572
			E164PhoneNumber: +724891571063
			TitleMale: Mr.
			TitleFemale: Queen
			FirstName: Whitney
			FirstNameMale: Kenny
			FirstNameFemale: Jana
			LastName: Rohan
			Name: Miss Casandra Kiehn
			UnixTime: 1197930901
			Date: 1982-02-27
			Time: 03:10:25
			MonthName: February
			Year: 1996
			DayOfWeek: Sunday
			DayOfMonth: 20
			Timestamp: 1973-06-21 14:50:46
			Century: IV
			TimeZone: Canada/Eastern
			TimePeriod: AM
			Word: nesciunt
			Sentence: Consequatur perferendis aut sit voluptatem accusantium.
			Paragraph: Aut consequatur sit perferendis accusantium voluptatem. Accusantium perferendis consequatur voluptatem sit aut. Aut sit accusantium consequatur voluptatem perferendis. Perferendis voluptatem aut accusantium consequatur sit.
			Currency: IRR,
			Amount: 88.990000,
			AmountWithCurrency: XBB 49257.100000,
			UUIDHypenated: 8f8e4463-9560-4a38-9b0c-ef24481e4e27,
			UUID: 90ea6479fd0e4940af741f0a87596b73,
			Skip:
		}
	*/
}

```

## Length And Boundary

---

You can set length for your random strings also set boundary for your integers.
```go
package main

import (
	"fmt"

	"github.com/bingoohuang/faker"
)

type SomeStruct struct {
	Inta  int   `faker:"boundary_start=5, boundary_end=10"`
	Int8  int8  `faker:"boundary_start=100, boundary_end=1000"`
	Int16 int16 `faker:"boundary_start=123, boundary_end=1123"`
	Int32 int32 `faker:"boundary_start=-10, boundary_end=8123"`
	Int64 int64 `faker:"boundary_start=31, boundary_end=88"`

	UInta  uint   `faker:"boundary_start=35, boundary_end=152"`
	UInt8  uint8  `faker:"boundary_start=5, boundary_end=1425"`
	UInt16 uint16 `faker:"boundary_start=245, boundary_end=2125"`
	UInt32 uint32 `faker:"boundary_start=0, boundary_end=40"`
	UInt64 uint64 `faker:"boundary_start=14, boundary_end=50"`

	ASString []string          `faker:"len=50"`
	SString  string            `faker:"len=25"`
	MSString map[string]string `faker:"len=30"`
	MIint    map[int]int       `faker:"boundary_start=5, boundary_end=10"`
}

func main(){
    faker.SetRandomMapAndSliceSize(10, 20) //Random generated map or array size between 10 and 20...
	a := SomeStruct{}
	err := faker.Fake(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}

```
Result:
```
{
    Inta:7
    Int8:-102
    Int16:556
    Int32:113
    Int64:70
    UInta:78
    UInt8:54
    UInt16:1797
    UInt32:8
    UInt64:34
    ASString:[
        geHYIpEoQhQdijFooVEAOyvtTwJOofbQPJdbHvEEdjueZaKIgI
        WVJBBtmrrVccyIydAiLSkMwWbFzFMEotEXsyUXqcmBTVORlkJK
        xYiRTRSZRuGDcMWYoPALVMZgIXoTQtmdGXQfbISKJiavLspuBV
        qsoiYlyRbXLDAMoIdQhgMriODYWCTEYepmjaldWLLjkulDGuQN
        GQXUlqNkVjPKodMebPIeoZZlxfhbQJOjHSRjUTrcgBFPeDZIxn
        MEeRkkLceDsqKLEJFEjJxHtYrYxQMxYcuRHEDGYSPbELDQLSsj
        tWUIACjQWeiUhbboGuuEQIhUJCRSBzVImpYwOlFbsjCRmxboZW
        ZDaAUZEgFKbMJoKIMpymTreeZGWLXNCfVzaEyWNdbkaZOcsfst
        uwlsZBMlEknIBsALpXRaplZWVtXTKzsWglRVBpmfsQfqraiEYA
        AXszbzsOzYPYeXHXRwoPmoPoBxopdFFvWMBTPCxESTepRpjlnB
        kTuOPHlUrSzUQRmZMYplWbyoBbWzQYCiydyzurOduhjuyiGrCE
        FZbeLMbelIeCMnixknIARZRbwALObGXADQqianJbkiEAqqpdnK
        TiQrZbnkvxEciyKXlliUDOGVdpMoAsHSalFbLcYyXxNFLAhqjy
        KlbjbloxkWKSqvUfJQPpFLoddWgeABfYUoaAnylKmEHwxgNsnO
        ]
    SString:VVcaPSFrOPYlEkpVyTRbSZneB
    MSString:map[
        ueFBFTTmqDwrXDoXAYTRhQRmLXhudA:AhQewvZfrlytbAROzGjpXUmNQzIoGl
        fZwrCsFfZwqMsDJXOUYIacflFIeyFU:VMufFCRRHTtuFthOrRAMbzbKVJHnvJ
        rHDQTyZqZVSPLwTtZfNSwKWrgmRghL:lRSXNHkhUyjDuBgoAfrQwOcHYilqRB
        BvCpQJMHzKXKbOoAnTXkLCNxKshwWr:tiNFrXAXUtdywkyygWBrEVrmAcAepD
        uWWKgHKTkUgAZiopAIUmgVWrkrceVy:GuuDNTUiaBtOKwWrMoZDiyaOPxywnq
        HohMjOdMDkAqimKPTgdjUorydpKkly:whAjmraukcZczskqycoJELlMJTghca
        umEgMBGUvBptdKImKsoWXMGJJoRbgT:tPpgHgLEyHmDOocOiSgTbXQHVduLxP
        SRQLHjBXCXKvbLIktdKeLwMnIFOmbi:IJBpLyTcraOxOUtwSKTisjElpulkTL
        dbnDeJZLqMXQGjbTSNxPSlfDHGCghU:JWrymovFwNWbIQBxPpQmlgJsgpXcui
        roraKNGnBXnrJlsxTnFgxHyZeTXdAC:XIcLWqUAQAbfkRrgfjrTVxZCvRJXyl
        TrvxqVVjXAboYDPvUglSJQrltPjzLx:nBhWdfNPybnNnCyQlSshWKOnwUMQzL
        dTHhWJWMwfVvKpIKTFCaoBJgKmnfbD:ixjNHsvSkRkFiNLpgUzIKPsheqhCeY
        lWyBrtfcGWiNbSTJZJXwOPvVngZZMk:kvlYeGgwguVtiafGKjHWsYWewbaXte
        bigsYNfVcNMGtnzgaqEjeRRlIcUdbR:hYOnJupEOvblTTEYzZYPuTVmvTmiit
        ]
    MIint:map[7:7 5:7 8:8 9:5 6:5]
}
```
