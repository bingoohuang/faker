package faker

// Faker is a simple fake data generator for your own struct.
// Save your time, and Fake your data for your testing now.
import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	// Sets nil if the value type is struct or map and the size of it equals to zero.
	shouldSetNil = false
	//Sets random integer generation to zero for slice and maps
	testRandZero = false
	//Sets the default number of string when it is created randomly.
	randomStringLen = 25
	//Sets the boundary for random value generation. Boundaries can not exceed integer(4 byte...)
	nBoundary = numberBoundary{start: 0, end: 100}
	//Sets the random size for slices and maps.
	randomMinSize = 0
	randomMaxSize = 100
)

type numberBoundary struct {
	start int
	end   int
}

// Supported tags
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tagName       = "faker"
	keep          = "keep"

	ID                    = "uuid_digit"
	HyphenatedID          = "uuid_hyphenated"
	Regex                 = "regex"
	Enum                  = "enum"
	Snow                  = "snow"
	EmailTag              = "email"
	MacAddressTag         = "mac_address"
	DomainNameTag         = "domain_name"
	UserNameTag           = "username"
	URLTag                = "url"
	IPV4Tag               = "ipv4"
	IPV6Tag               = "ipv6"
	PASSWORD              = "password"
	LATITUDE              = "lat"
	LONGITUDE             = "long"
	CreditCardNumber      = "cc_number"
	CreditCardType        = "cc_type"
	PhoneNumber           = "phone_number"
	ChinaMobileNumber     = "china_mobile_number"
	TollFreeNumber        = "toll_free_number"
	E164PhoneNumberTag    = "e_164_phone_number"
	TitleMaleTag          = "title_male"
	TitleFemaleTag        = "title_female"
	FirstNameTag          = "first_name"
	FirstNameMaleTag      = "first_name_male"
	FirstNameFemaleTag    = "first_name_female"
	LastNameTag           = "last_name"
	NAME                  = "name"
	UnixTimeTag           = "unix_time"
	DATE                  = "date"
	TIME                  = "time"
	MonthNameTag          = "month_name"
	YEAR                  = "year"
	DayOfWeekTag          = "day_of_week"
	DayOfMonthTag         = "day_of_month"
	TIMESTAMP             = "timestamp"
	CENTURY               = "century"
	TIMEZONE              = "timezone"
	TimePeriodTag         = "time_period"
	WORD                  = "word"
	SENTENCE              = "sentence"
	PARAGRAPH             = "paragraph"
	CurrencyTag           = "currency"
	AmountTag             = "amount"
	AmountWithCurrencyTag = "amount_with_currency"
	SKIP                  = "-"
	Length                = "len"
	BoundaryStart         = "boundary_start"
	BoundaryEnd           = "boundary_end"
	Equals                = "="
	comma                 = ","
)

// TaggedFunction used as the standard layout function for tag providers in struct.
// This type also can be used for custom provider.
type TaggedFunction func(v reflect.Value) (interface{}, error)

// TaggedFunctionV2 used as the standard layout function for tag providers in struct.
// This type also can be used for custom provider.
type TaggedFunctionV2 func(v reflect.Value, tag Tag) (interface{}, error)

var mapperTag = map[string]interface{}{
	EmailTag:              internet.Email,
	MacAddressTag:         internet.MacAddress,
	DomainNameTag:         internet.DomainName,
	URLTag:                internet.URL,
	UserNameTag:           internet.UserName,
	IPV4Tag:               internet.IPv4,
	IPV6Tag:               internet.IPv6,
	PASSWORD:              internet.Password,
	CreditCardType:        pay.CreditCardType,
	CreditCardNumber:      pay.CreditCardNumber,
	LATITUDE:              address.Latitude,
	LONGITUDE:             address.Longitude,
	PhoneNumber:           phone.PhoneNumber,
	ChinaMobileNumber:     phone.ChinaMobileNumber,
	TollFreeNumber:        phone.TollFreePhoneNumber,
	E164PhoneNumberTag:    phone.E164PhoneNumber,
	TitleMaleTag:          person.TitleMale,
	TitleFemaleTag:        person.TitleFeMale,
	FirstNameTag:          person.FirstName,
	FirstNameMaleTag:      person.FirstNameMale,
	FirstNameFemaleTag:    person.FirstNameFemale,
	LastNameTag:           person.LastName,
	NAME:                  person.Name,
	UnixTimeTag:           date.UnixTime,
	DATE:                  date.Date,
	TIME:                  date.Time,
	MonthNameTag:          date.MonthName,
	YEAR:                  date.Year,
	DayOfWeekTag:          date.DayOfWeek,
	DayOfMonthTag:         date.DayOfMonth,
	TIMESTAMP:             date.Timestamp,
	CENTURY:               date.Century,
	TIMEZONE:              date.TimeZone,
	TimePeriodTag:         date.TimePeriod,
	WORD:                  lorem.Word,
	SENTENCE:              lorem.Sentence,
	PARAGRAPH:             lorem.Paragraph,
	CurrencyTag:           pri.Currency,
	AmountTag:             pri.Amount,
	AmountWithCurrencyTag: pri.AmountWithCurrency,
	ID:                    identifier.Digit,
	HyphenatedID:          identifier.Hyphenated,
	Regex:                 regexer.Gen,
	Enum:                  enumer.Gen,
	Snow:                  snow.Gen,
}

// Generic Error Messages for tags
// 		ErrValueNotPtr: Error when value is not pointer
// 		ErrTagNotSupported: Error when tag is not supported
// 		ErrTagAlreadyExists: Error when tag exists and call AddProvider
// 		ErrNotSupportedPointer: Error when passing unsupported pointer
const (
	ErrValueNotPtr             = "Not a pointer value"
	ErrTagNotSupported         = "Tag unsupported"
	ErrTagAlreadyExists        = "Tag exists"
	ErrNotSupportedPointer     = "Use sample:=new(%s)\n faker.Fake(sample) instead"
	ErrSmallerThanZero         = "Size:%d is smaller than zero."
	ErrStartValueBiggerThanEnd = "Start value can not be bigger than end value."
	ErrWrongFormattedTag       = "Tag \"%s\" is not written properly"
	ErrUnknownType             = "Unknown Type"
	ErrNotSupportedTypeForTag  = "Type is not supported by tag."
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SetNilIfLenIsZero allows to set nil for the slice and maps, if size is 0.
func SetNilIfLenIsZero(setNil bool) {
	shouldSetNil = setNil
}

// SetRandomStringLength sets a length for random string generation
func SetRandomStringLength(size int) error {
	if size < 0 {
		return fmt.Errorf(ErrSmallerThanZero, size)
	}
	randomStringLen = size
	return nil
}

// SetRandomMapAndSliceSize sets the size for maps and slices for random generation.
func SetRandomMapAndSliceSize(minSize, maxSize int) error {
	if maxSize <= 0 {
		return fmt.Errorf(ErrSmallerThanZero, maxSize)
	}
	if minSize < 0 {
		return fmt.Errorf(ErrSmallerThanZero, minSize)
	}

	randomMinSize = minSize
	randomMaxSize = maxSize
	return nil
}

// SetRandomNumberBoundaries sets boundary for random number generation
func SetRandomNumberBoundaries(start, end int) error {
	if start > end {
		return errors.New(ErrStartValueBiggerThanEnd)
	}
	nBoundary = numberBoundary{start: start, end: end}
	return nil
}

// Fake is the main function. Will generate a fake data based on your struct.  You can use this for automation testing, or anything that need automated data.
// You don't need to Create your own data for your testing.
func Fake(a interface{}) error {
	reflectType := reflect.TypeOf(a)

	if reflectType.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}

	if reflect.ValueOf(a).IsNil() {
		return fmt.Errorf(ErrNotSupportedPointer, reflectType.Elem().String())
	}

	rval := reflect.ValueOf(a)

	finalValue, err := getValue(a)
	if err != nil {
		return err
	}

	rval.Elem().Set(finalValue.Elem().Convert(reflectType.Elem()))
	return nil
}

// AddProvider extend faker with tag to generate fake data with specified custom algoritm
// Example:
// 		type Gondoruwo struct {
// 			Name       string
// 			Locatadata int
// 		}
//
// 		type Sample struct {
// 			ID                 int64     `faker:"customIdFaker"`
// 			Gondoruwo          Gondoruwo `faker:"gondoruwo"`
// 			Danger             string    `faker:"danger"`
// 		}
//
// 		func CustomGenerator() {
// 			// explicit
// 			faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
// 			 	return int64(43), nil
// 			})
// 			// functional
// 			faker.AddProvider("danger", func() faker.TaggedFunction {
// 				return func(v reflect.Value) (interface{}, error) {
// 					return "danger-ranger", nil
// 				}
// 			}())
// 			faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
// 				obj := Gondoruwo{
// 					Name:       "Power",
// 					Locatadata: 324,
// 				}
// 				return obj, nil
// 			})
// 		}
//
// 		func main() {
// 			CustomGenerator()
// 			var sample Sample
// 			faker.Fake(&sample)
// 			fmt.Printf("%+v", sample)
// 		}
//
// Will print
// 		{ID:43 Gondoruwo:{Name:Power Locatadata:324} Danger:danger-ranger}
// Notes: when using a custom provider make sure to return the same type as the field
func AddProvider(tag string, provider TaggedFunction) error {
	if _, ok := mapperTag[tag]; ok {
		return errors.New(ErrTagAlreadyExists)
	}

	mapperTag[tag] = provider

	return nil
}

// AddProviderV2 extend faker with tag to generate fake data with specified custom algoritm.
// It is similar to AddProvider but with a differnt provider type.
func AddProviderV2(tag string, provider TaggedFunctionV2) error {
	if _, ok := mapperTag[tag]; ok {
		return errors.New(ErrTagAlreadyExists)
	}

	mapperTag[tag] = provider

	return nil
}

func getValue(a interface{}) (reflect.Value, error) {
	t := reflect.TypeOf(a)
	if t == nil {
		return reflect.Value{}, fmt.Errorf("interface{} not allowed")
	}
	k := t.Kind()

	switch k {
	case reflect.Ptr:
		v := reflect.New(t.Elem())
		var val reflect.Value
		var err error
		if a != reflect.Zero(reflect.TypeOf(a)).Interface() {
			val, err = getValue(reflect.ValueOf(a).Elem().Interface())
			if err != nil {
				return reflect.Value{}, err
			}
		} else {
			val, err = getValue(v.Elem().Interface())
			if err != nil {
				return reflect.Value{}, err
			}
		}
		v.Elem().Set(val.Convert(t.Elem()))
		return v, nil
	case reflect.Struct:
		switch t.String() {
		case "time.Time":
			ft := time.Now().Add(time.Duration(rand.Int63()))
			return reflect.ValueOf(ft), nil
		default:
			originalDataVal := reflect.ValueOf(a)
			v := reflect.New(t).Elem()
			for i := 0; i < v.NumField(); i++ {
				if !v.Field(i).CanSet() {
					continue // to avoid panic to set on unexported field in struct
				}
				//tags := decodeTags(t, i)
				tags := decodeTags(t.Field(i).Tag.Get(tagName), t.Field(i).Type)
				switch {
				case tags.KeepOriginal:
					zero, err := isZero(reflect.ValueOf(a).Field(i))
					if err != nil {
						return reflect.Value{}, err
					}
					if zero {
						err := setDataWithTag(v.Field(i).Addr(), tags)
						if err != nil {
							return reflect.Value{}, err
						}
						continue
					}
					v.Field(i).Set(reflect.ValueOf(a).Field(i))
				case tags.Mapper == "":
					val, err := getValue(v.Field(i).Interface())
					if err != nil {
						return reflect.Value{}, err
					}
					val = val.Convert(v.Field(i).Type())
					v.Field(i).Set(val)
				case tags.Mapper == SKIP:
					item := originalDataVal.Field(i).Interface()
					if v.CanSet() && item != nil {
						v.Field(i).Set(reflect.ValueOf(item))
					}
				default:
					err := setDataWithTag(v.Field(i).Addr(), tags)
					if err != nil {
						return reflect.Value{}, err
					}
				}

			}
			return v, nil
		}

	case reflect.String:
		res := randomString(randomStringLen)
		return reflect.ValueOf(res), nil
	case reflect.Array, reflect.Slice:
		l := randomSliceAndMapSize()
		if shouldSetNil && l == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeSlice(t, l, l)
		for i := 0; i < v.Len(); i++ {
			val, err := getValue(v.Index(i).Interface())
			if err != nil {
				return reflect.Value{}, err
			}
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Int:
		return reflect.ValueOf(randomInteger()), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(randomInteger())), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(randomInteger())), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(randomInteger())), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(randomInteger())), nil
	case reflect.Float32:
		return reflect.ValueOf(rand.Float32()), nil
	case reflect.Float64:
		return reflect.ValueOf(rand.Float64()), nil
	case reflect.Bool:
		val := rand.Intn(2) > 0
		return reflect.ValueOf(val), nil
	case reflect.Uint:
		return reflect.ValueOf(uint(randomInteger())), nil
	case reflect.Uint8:
		return reflect.ValueOf(uint8(randomInteger())), nil
	case reflect.Uint16:
		return reflect.ValueOf(uint16(randomInteger())), nil
	case reflect.Uint32:
		return reflect.ValueOf(uint32(randomInteger())), nil
	case reflect.Uint64:
		return reflect.ValueOf(uint64(randomInteger())), nil

	case reflect.Map:
		l := randomSliceAndMapSize()
		if shouldSetNil && l == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeMap(t)
		for i := 0; i < l; i++ {
			keyInstance := reflect.New(t.Key()).Elem().Interface()
			key, err := getValue(keyInstance)
			if err != nil {
				return reflect.Value{}, err
			}

			valueInstance := reflect.New(t.Elem()).Elem().Interface()
			val, err := getValue(valueInstance)
			if err != nil {
				return reflect.Value{}, err
			}
			v.SetMapIndex(key, val)
		}
		return v, nil
	default:
		err := fmt.Errorf("no support for kind %+v", t)
		return reflect.Value{}, err
	}

}

func isZero(field reflect.Value) (bool, error) {
	for _, kind := range []reflect.Kind{reflect.Struct, reflect.Slice, reflect.Array, reflect.Map} {
		if kind == field.Kind() {
			return false, fmt.Errorf("keep not allowed on struct")
		}
	}
	return reflect.Zero(field.Type()).Interface() == field.Interface(), nil
}

//func decodeTags(typ reflect.Type, i int) Tag {
func decodeTags(rawTag string, typ reflect.Type) Tag {
	//field := typ.Field(i)
	//rawTag := field.Tag.Get(tagName)
	tags := strings.Split(rawTag, ",")

	keepOriginal := false
	mapper := ""
	opts := make(map[string]string)
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		if tag == keep {
			keepOriginal = true
		} else {
			kvs := strings.SplitN(tag, "=", 2)
			k := kvs[0]
			v := ""
			if mapper == "" {
				mapper = k
			}
			if len(kvs) == 1 {
				v = "true"
			} else {
				v = kvs[1]
			}

			opts[k] = v
		}
	}

	if mapper == "" && rawTag != "" {
		mapper = "none"
	}

	return Tag{
		RawTag:       rawTag,
		Type:         typ,
		Mapper:       mapper,
		Opts:         opts,
		KeepOriginal: keepOriginal,
	}
}

// Tag represents tag of faker related to the struct field
type Tag struct {
	Type         reflect.Type
	RawTag       string
	Mapper       string
	Opts         map[string]string
	KeepOriginal bool
}

func invoke(mapperFn interface{}, v reflect.Value, tag Tag) (interface{}, error) {
	fnValue := reflect.ValueOf(mapperFn)
	var args []reflect.Value
	switch fnValue.Type().NumIn() {
	case 1:
		args = []reflect.Value{reflect.ValueOf(v)}
	case 2:
		args = []reflect.Value{reflect.ValueOf(v), reflect.ValueOf(tag)}
	default:
		return nil, fmt.Errorf("unknown mapper func %T", mapperFn)
	}

	results := fnValue.Call(args)
	data := results[0].Interface()
	err := results[1].Interface()
	if err == nil {
		return data, nil
	}

	return data, err.(error)
}

func setDataWithTag(v reflect.Value, tag Tag) error {
	if v.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}
	v = reflect.Indirect(v)
	mapperName := tag.Mapper
	switch v.Kind() {
	case reflect.Ptr:
		mapper, exist := mapperTag[mapperName]
		if !exist {
			return errors.New(ErrTagNotSupported)
		}

		newv := reflect.New(v.Type().Elem())
		res, err := invoke(mapper, newv.Elem(), tag)
		if err != nil {
			return err
		}

		resV := reflect.ValueOf(res)
		if resV.Kind() == reflect.Ptr {
			v.Set(resV)
			return nil
		}

		newv.Elem().Set(resV)
		v.Set(newv)
		return nil
	case reflect.String:
		return userDefinedString(v, tag)
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return userDefinedNumber(v, tag)
	case reflect.Slice, reflect.Array:
		return userDefinedArray(v, tag)
	case reflect.Map:
		return userDefinedMap(v, tag)
	default:
		if _, exist := mapperTag[mapperName]; !exist {
			return errors.New(ErrTagNotSupported)
		}
		res, err := invoke(mapperTag[mapperName], v, tag)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(res))
	}
	return nil
}

func userDefinedMap(v reflect.Value, tag Tag) error {
	l := randomSliceAndMapSize()
	if shouldSetNil && l == 0 {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	definedMap := reflect.MakeMap(v.Type())
	for i := 0; i < l; i++ {
		key, err := getValueWithTag(v.Type().Key(), tag)
		if err != nil {
			return err
		}
		val, err := getValueWithTag(v.Type().Elem(), tag)
		if err != nil {
			return err
		}
		definedMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	v.Set(definedMap)
	return nil
}

func getValueWithTag(t reflect.Type, tag Tag) (interface{}, error) {
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res, err := extractNumberFromTag(tag, t)
		if err != nil {
			return nil, err
		}
		return res, nil
	case reflect.String:
		res, err := extractStringFromTag(tag)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return 0, errors.New(ErrUnknownType)
	}
}

func userDefinedArray(v reflect.Value, tag Tag) error {
	l := randomSliceAndMapSize()
	if shouldSetNil && l == 0 {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	array := reflect.MakeSlice(v.Type(), l, l)
	for i := 0; i < l; i++ {
		res, err := getValueWithTag(v.Type().Elem(), tag)
		if err != nil {
			return err
		}
		array.Index(i).Set(reflect.ValueOf(res))
	}
	v.Set(array)
	return nil
}

func userDefinedString(v reflect.Value, tag Tag) error {
	var res interface{}
	var err error

	if tagFunc, ok := mapperTag[tag.Mapper]; ok {
		res, err = invoke(tagFunc, v, tag)
		if err != nil {
			return err
		}
	} else {
		res, err = extractStringFromTag(tag)
		if err != nil {
			return err
		}
	}
	if res == nil {
		return errors.New(ErrTagNotSupported)
	}
	val, _ := res.(string)
	v.SetString(val)
	return nil
}

func userDefinedNumber(v reflect.Value, tag Tag) error {
	var res interface{}
	var err error

	if tagFunc, ok := mapperTag[tag.Mapper]; ok {
		res, err = invoke(tagFunc, v, tag)
		if err != nil {
			return err
		}
	} else {
		res, err = extractNumberFromTag(tag, v.Type())
		if err != nil {
			return err
		}
	}
	if res == nil {
		return errors.New(ErrTagNotSupported)
	}

	v.Set(reflect.ValueOf(res))
	return nil
}

func extractStringFromTag(tag Tag) (interface{}, error) {
	if fn, ok := mapperTag[tag.Mapper]; ok {
		res, err := invoke(fn, reflect.ValueOf(nil), tag)
		return res, err
	}

	if !strings.Contains(tag.RawTag, Length) {
		return nil, errors.New(ErrTagNotSupported)
	}
	l, err := extractNumberFromText(tag.RawTag)
	if err != nil {
		return nil, err
	}
	res := randomString(l)
	return res, nil
}

func extractNumberFromTag(Tag Tag, t reflect.Type) (interface{}, error) {
	if fn, ok := mapperTag[Tag.Mapper]; ok {
		res, err := invoke(fn, reflect.New(t), Tag)
		return res, err
	}

	tag := Tag.RawTag
	if !(strings.Contains(tag, BoundaryStart) && strings.Contains(tag, BoundaryEnd)) {
		return nil, errors.New(ErrTagNotSupported)
	}
	valuesStr := strings.SplitN(tag, comma, -1)
	if len(valuesStr) != 2 {
		return nil, fmt.Errorf(ErrWrongFormattedTag, tag)
	}
	startBoundary, err := extractNumberFromText(valuesStr[0])
	if err != nil {
		return nil, err
	}
	endBoundary, err := extractNumberFromText(valuesStr[1])
	if err != nil {
		return nil, err
	}
	boundary := numberBoundary{start: startBoundary, end: endBoundary}
	switch t.Kind() {
	case reflect.Uint:
		return uint(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint8:
		return uint8(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint16:
		return uint16(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint32:
		return uint32(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint64:
		return uint64(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int:
		return randomIntegerWithBoundary(boundary), nil
	case reflect.Int8:
		return int8(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int16:
		return int16(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int32:
		return int32(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int64:
		return int64(randomIntegerWithBoundary(boundary)), nil
	default:
		return nil, errors.New(ErrNotSupportedTypeForTag)
	}
}

func extractNumberFromText(text string) (int, error) {
	text = strings.TrimSpace(text)
	texts := strings.SplitN(text, Equals, -1)
	if len(texts) != 2 {
		return 0, fmt.Errorf(ErrWrongFormattedTag, text)
	}
	return strconv.Atoi(texts[1])
}

func randomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// randomIntegerWithBoundary returns a random integer between input start and end boundary. [start, end)
func randomIntegerWithBoundary(boundary numberBoundary) int {
	return rand.Intn(boundary.end-boundary.start) + boundary.start
}

// randomInteger returns a random integer between start and end boundary. [start, end)
func randomInteger() int {
	return rand.Intn(nBoundary.end-nBoundary.start) + nBoundary.start
}

// randomSliceAndMapSize returns a random integer between [0,randomSliceAndMapSize). If the testRandZero is set, returns 0
// Written for test purposes for shouldSetNil
func randomSliceAndMapSize() int {
	if testRandZero {
		return 0
	}
	return rand.Intn(randomMaxSize-randomMinSize) + randomMinSize
}

func randomElementFromSliceString(s []string) string {
	return s[rand.Int()%len(s)]
}
func randomStringNumber(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandomInt generates digits with len between minDigit and maxDigit
func RandomInt(minDigit, maxDigit int) []int {
	p := rand.Perm(maxDigit - minDigit + 1)

	for i := range p {
		p[i] += minDigit
	}
	return p
}
