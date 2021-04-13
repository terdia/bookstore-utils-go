package dates

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestGetNowUTCFormatConstants(t *testing.T) {
	assert.EqualValues(t, "2006-01-02 15:04:05", dbDateLayout, "database date format incorrect")
	assert.EqualValues(t, "2006-01-02T15:04:05Z", defaultDateLayout, "default date format incorrect")
}

func TestGetNowUTC(t *testing.T) {
	nowUTC := GetNowUTC()
	assert.NotNil(t, nowUTC)
	assert.EqualValues(t, "UTC", nowUTC.Location().String())
}

func TestGetNowUtcDbFormat(t *testing.T) {
	nowString := GetNowUtcDbFormat()
	assert.NotNil(t, nowString)

	assert.EqualValues(t, "string", reflect.ValueOf(nowString).Kind().String())
	dbDate, err := time.Parse(dbDateLayout, nowString)

	assert.Nil(t, err)
	assert.NotNil(t, dbDate)
	assert.EqualValues(t, "UTC", dbDate.Location().String())

	_, wrongFormatErr := time.Parse(defaultDateLayout, nowString)
	assert.NotNil(t, wrongFormatErr)
}

func TestGetNowUtcDefault(t *testing.T) {
	nowString := GetNowUtcDefault()
	assert.NotNil(t, nowString)

	parsed, err := time.Parse(defaultDateLayout, nowString)
	assert.Nil(t, err)
	assert.EqualValues(t, "UTC", parsed.Location().String())
	assert.EqualValues(t, "string", reflect.ValueOf(nowString).Kind().String())

	_, wrongFormatErr := time.Parse(dbDateLayout, nowString)
	assert.NotNil(t, wrongFormatErr)
}
