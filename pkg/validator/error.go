package validator

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/lib/pq"
)

var (
	ErrInvalidName  = InvalidArgErr(fmt.Errorf("name_min_3_max_200"))
	ErrInvalidPrice = InvalidArgErr(fmt.Errorf("price_min_0.1_max_50000"))
	ErrInvalidCode  = InvalidArgErr(fmt.Errorf("code_min_2_max_20"))
)

func ParseReadDbErrMsg(err error) error {
	if err == sql.ErrNoRows {
		return InvalidArgErr(fmt.Errorf("no_rows_found"))
	}
	return InternalErr(err)
}
func ParseWriteDbErrMsg(err *pq.Error) error {
	if strings.Contains(err.Message, "unique constraint") {
		return AlreadyExistsErr(fmt.Errorf("unique_constraint_%s", err.Constraint))
	}
	if strings.Contains(err.Message, "check") {
		return InvalidArgErr(fmt.Errorf("check_failed_%s", err.Constraint))
	}

	if strings.Contains(err.Message, "value too long") {
		return InvalidArgErr(fmt.Errorf("value_too_long_%s", err.Constraint))
	}

	return err
}
func ErrorMinMax(min float32, max float32) error {
	return fmt.Errorf("min_%f_max_%f", min, max)
}

func ErrorInvalid() error {
	return fmt.Errorf("invalid")
}

func AlreadyExistsErr(err error) error {
	return connect.NewError(connect.CodeAlreadyExists, err)
}
func InvalidArgErr(err error) error {
	return connect.NewError(connect.CodeInvalidArgument, err)
}

func InternalErr(err error) error {
	return connect.NewError(connect.CodeInternal, err)
}

func UnauthErr(err error) error {
	return connect.NewError(connect.CodeUnauthenticated, err)
}

func UnavilableErr(err error) error {
	return connect.NewError(connect.CodeUnavailable, err)
}
