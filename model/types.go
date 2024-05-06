package model

import "strconv"

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = Uint64(i)
	return
}

func (u64 Uint64) Value() uint64 {
	return uint64(u64)
}

func (u64 Uint64) String() string {
	return strconv.FormatUint(uint64(u64), 10)
}

// Int support string quoted number in json
type Int int

// UnmarshalJSON implement json Unmarshal interface
func (i *Int) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	v, _ := strconv.Atoi(string(b))
	*i = Int(v)
	return
}

func (i Int) Value() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// JSONUint64 support string quoted number in json and marshal to string
type JSONUint64 uint64

func JSONUint64FromUint64(v uint64) JSONUint64 {
	return JSONUint64(v)
}

// UnmarshalJSON implement json Unmarshal interface
func (u64 *JSONUint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = JSONUint64(i)
	return
}

func (u64 JSONUint64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(uint64(u64), 10) + `"`), nil
}

func (u64 JSONUint64) Value() uint64 {
	return uint64(u64)
}

// JSONInt64 support string quoted number in json and marshal to string
type JSONInt64 int64

func JSONInt64FromInt64(v int64) JSONInt64 {
	return JSONInt64(v)
}

// UnmarshalJSON implement json Unmarshal interface
func (i64 *JSONInt64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseInt(string(b), 10, 64)
	*i64 = JSONInt64(i)
	return
}

func (i64 JSONInt64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(int64(i64), 10) + `"`), nil
}

func (i64 JSONInt64) Value() int64 {
	return int64(i64)
}

// JSONInt support string quoted number in json and marshal to string
type JSONInt int

func JSONIntFromInt(v int) JSONInt {
	return JSONInt(v)
}

// UnmarshalJSON implement json Unmarshal interface
func (i *JSONInt) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	v, _ := strconv.Atoi(string(b))
	*i = JSONInt(v)
	return
}

func (i JSONInt) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.Itoa(int(i)) + `"`), nil
}

func (i JSONInt) Value() int {
	return int(i)
}
