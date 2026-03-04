package objects

import "fmt"

func toUint8(value any) (uint8, error) {
	switch v := value.(type) {
	case uint8:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid uint8 type: %T", value)
	}
}

func toUint16(value any) (uint16, error) {
	switch v := value.(type) {
	case uint16:
		return v, nil
	}
	return 0, fmt.Errorf("invalid uint16 type: %T", value)
}

func toUint32(value any) (uint32, error) {
	switch v := value.(type) {
	case uint32:
		return v, nil
	}
	return 0, fmt.Errorf("invalid uint32 type: %T", value)
}

func toUint64(value any) (uint64, error) {
	switch v := value.(type) {
	case uint64:
		return v, nil
	}
	return 0, fmt.Errorf("invalid uint64 type: %T", value)
}

func toBool(value any) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	}
	return false, fmt.Errorf("invalid bool type: %T", value)
}

func toInt8Value(value any) (int8, error) {
	switch v := value.(type) {
	case int8:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid int8 type: %T", value)
	}
}

func toInt16Value(value any) (int16, error) {
	switch v := value.(type) {
	case int16:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid int16 type: %T", value)
	}
}
