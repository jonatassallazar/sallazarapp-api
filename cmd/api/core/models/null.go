package models

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type NullInt64 struct {
	sql.NullInt64
}

func (ni *NullInt64) CheckNullValue() int64 {
	if !ni.Valid {
		return 0
	}
	return ni.Int64
}

type NullBool struct {
	sql.NullBool
}

func (nb *NullBool) CheckNullValue() bool {
	if !nb.Valid {
		return false
	}
	return nb.Bool
}

type NullByte struct {
	sql.NullByte
}

func (nb *NullByte) CheckNullValue() []byte {

	if !nb.Valid {
		return []byte("")
	}

	return []byte{nb.Byte}
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) CheckNullValue() string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

type NullTime struct {
	mysql.NullTime
}

func (nt *NullTime) CheckNullValue() time.Time {
	if !nt.Valid {
		timeParsed, _ := time.Parse(time.RFC3339, "0000-00-00")
		return timeParsed
	}

	return nt.Time
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) CheckNullValue() float64 {
	if !nf.Valid {
		return 0.0
	}
	return nf.Float64
}
