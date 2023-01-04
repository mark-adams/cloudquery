package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func msSQLType(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "bit"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "float(53)"
	case schema.TypeUUID:
		return "uniqueidentifier"
	case schema.TypeString:
		return "nvarchar(max)"
	case schema.TypeByteArray:
		return "varbinary(max)"
	case schema.TypeStringArray:
		return "nvarchar(max)"
	case schema.TypeTimestamp:
		return "datetime2"
	case schema.TypeJSON:
		return "nvarchar(max)"
	case schema.TypeUUIDArray:
		return "nvarchar(max)"
	case schema.TypeCIDR:
		return "nvarchar(max)"
	case schema.TypeCIDRArray:
		return "nvarchar(max)"
	case schema.TypeMacAddr:
		return "nvarchar(max)"
	case schema.TypeMacAddrArray:
		return "nvarchar(max)"
	case schema.TypeInet:
		return "nvarchar(max)"
	case schema.TypeInetArray:
		return "nvarchar(max)"
	case schema.TypeIntArray:
		return "nvarchar(max)"
	default:
		panic("unknown type " + t.String())
	}
}
