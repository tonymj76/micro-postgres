package datastore

import (
	"database/sql/driver"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/pgtype"
	"github.com/tonymj76/micro-postgres/proto/user"
)

type roleWrapper user.Role

// Value implements database/sql/driver.Valuer for pbUsers.Role
func (rw roleWrapper) Value() (driver.Value, error) {
	switch user.Role(rw) {
	case user.Role_ADMIN:
		return "admin", nil
	case user.Role_GUEST:
		return "guest", nil
	case user.Role_MEMBER:
		return "member", nil
	default:
		return nil, fmt.Errorf("invalid Role: %q", rw)
	}
}

// Scan implements database/sql/driver.Scanner for pbUsers.Role
func (rw *roleWrapper) Scan(in interface{}) error {
	switch in.(string) {
	case "admin":
		*rw = roleWrapper(user.Role_ADMIN)
		return nil
	case "guest":
		*rw = roleWrapper(user.Role_GUEST)
		return nil
	case "member":
		*rw = roleWrapper(user.Role_MEMBER)
		return nil
	default:
		return fmt.Errorf("invalid Role: %q", in.(string))
	}
}

type timeWrapper timestamp.Timestamp

// Value implements database/sql/driver.Valuer for timestamp.Timestamp
func (tw *timeWrapper) Value() (driver.Value, error) {
	return ptypes.Timestamp((*timestamp.Timestamp)(tw))
}

// Scan implements database/sql/driver.Scanner for timestamp.Timestamp
func (tw *timeWrapper) Scan(in interface{}) error {
	var t pgtype.Timestamptz
	err := t.Scan(in)
	if err != nil {
		return err
	}
	tp, err := ptypes.TimestampProto(t.Time)
	if err != nil {
		return err
	}
	*tw = (timeWrapper)(*tp)
	return nil
}

type durationWrapper duration.Duration

// Value implements database/sql/driver.Valuer for duration.Duration
func (dw *durationWrapper) Value() (driver.Value, error) {
	d, err := ptypes.Duration((*duration.Duration)(dw))
	if err != nil {
		return nil, err
	}

	i := pgtype.Interval{
		Microseconds: int64(d) / 1000,
		Status:       pgtype.Present,
	}

	return i.Value()
}
