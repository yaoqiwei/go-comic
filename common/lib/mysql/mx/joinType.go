package mx

type JoinType byte

func (t JoinType) String() string {
	return joinTypeName[t]
}

const (
	NO_JOIN JoinType = iota
	INNER_JOIN
	LEFT_JOIN
	RIGHT_JOIN
	OUTER_JOIN
)

var joinTypeName = map[JoinType]string{
	NO_JOIN:    "NO JOIN",
	INNER_JOIN: "JOIN",
	LEFT_JOIN:  "LEFT JOIN",
	RIGHT_JOIN: "RIGHT JOIN",
}
